// Command notices gates the SDK's production dependency licenses and keeps
// THIRD-PARTY-NOTICES.md in sync with the vendored module set.
//
// It resolves the runtime closure with `go list -mod=vendor -deps .`
// (the same set a consuming build links), classifies each module's vendored
// license text, fails on any license outside the allowlist, and regenerates
// the notices file from that set. CI runs this command in check mode on every
// push, so the committed file cannot drift from the real dependency set.
// After a dependency change, regenerate deliberately with `make notices` and
// commit the result.
//
// A module whose license cannot be classified, or classifies outside the
// allowlist, is a release decision, not something to wave through: the
// command fails and names the module so a human can rule on it.
package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

const noticesPath = "THIRD-PARTY-NOTICES.md"

// ownModule is the module whose dependencies are being audited; it is
// excluded from its own closure.
const ownModule = "github.com/ezclassify/sdk-go"

// allowedLicenses is the set of licenses cleared for redistribution with the
// Apache-2.0-licensed SDK. It mirrors the permissive core of the allowlist used across the EzClassify
// repositories; anything else (including copyleft and anything the classifier
// cannot fingerprint) fails the build for human review.
var allowedLicenses = map[string]bool{
	"MIT":          true,
	"Apache-2.0":   true,
	"BSD-2-Clause": true,
	"BSD-3-Clause": true,
	"ISC":          true,
}

// moduleNotes carries the human-written one-liner for each known module,
// explaining why the SDK depends on it. A module without an entry gets a
// generic line; add a real one when a new dependency lands.
var moduleNotes = map[string]string{
	"github.com/oapi-codegen/runtime":    "The generated HTTP client (`internal/api`) uses this module for request/response parameter binding (path, query, and form styling).",
	"github.com/apapsch/go-jsonmerge/v2": "Pulled in transitively by `oapi-codegen/runtime` for JSON merge-patch handling.",
	"github.com/google/uuid":             "Pulled in transitively by `oapi-codegen/runtime` for UUID parameter types.",
}

// licenseFileNames are the file names probed, in order, for a module's
// license text under vendor/<module path>/.
var licenseFileNames = []string{"LICENSE", "LICENSE.md", "LICENSE.txt", "COPYING"}

type module struct {
	path        string
	version     string
	license     string
	licenseFile string // repo-relative path to the vendored license text
	copyright   string
}

func main() {
	write := flag.Bool("write", false, "regenerate "+noticesPath+" instead of checking against it")
	flag.Parse()

	mods, err := runtimeClosure()
	if err != nil {
		fail("%v", err)
	}

	var violations []string
	for i := range mods {
		if err := classify(&mods[i]); err != nil {
			violations = append(violations, err.Error())
			continue
		}
		if !allowedLicenses[mods[i].license] {
			violations = append(violations,
				fmt.Sprintf("%s@%s: license %s is not on the allowlist", mods[i].path, mods[i].version, mods[i].license))
		}
	}
	if len(violations) > 0 {
		fmt.Fprintln(os.Stderr, "notices: dependency license gate failed; these need a human ruling:")
		for _, v := range violations {
			fmt.Fprintf(os.Stderr, "  %s\n", v)
		}
		os.Exit(1)
	}

	generated := render(mods)

	if *write {
		if err := os.WriteFile(noticesPath, []byte(generated), 0o644); err != nil {
			fail("%v", err)
		}
		fmt.Printf("notices: wrote %s (%d production modules, all licenses allowed)\n", noticesPath, len(mods))
		return
	}

	committed, err := os.ReadFile(noticesPath)
	if err != nil {
		fail("%v (run 'make notices' to create it)", err)
	}
	if string(committed) != generated {
		fail("%s is stale for the current vendored module set; run 'make notices' and commit the result", noticesPath)
	}
	fmt.Printf("notices: OK, licenses allowed and %s matches the vendored module set (%d modules)\n", noticesPath, len(mods))
}

func fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "notices: "+format+"\n", args...)
	os.Exit(1)
}

// runtimeClosure lists every external module in the transitive dependency
// closure of the public SDK package, resolved strictly from vendor/ so the
// audited set is exactly what ships.
func runtimeClosure() ([]module, error) {
	cmd := exec.Command("go", "list", "-mod=vendor", "-deps",
		"-f", "{{if .Module}}{{.Module.Path}} {{.Module.Version}}{{end}}", ".")
	out, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("go list failed: %s", strings.TrimSpace(string(ee.Stderr)))
		}
		return nil, fmt.Errorf("go list failed: %v", err)
	}

	return parseClosure(string(out))
}

// parseClosure turns the `go list` output into the deduplicated, sorted
// module set. A non-empty line that is not exactly `path version` (for
// example a version-less local-replace module) is a hard error: a module the
// gate cannot attribute must fail the build, not slip past it.
func parseClosure(out string) ([]module, error) {
	seen := map[string]bool{}
	var mods []module
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		// The main module itself prints without a version; it is the audited
		// module, not a dependency, so it is exempt before the shape check.
		if len(fields) == 0 || fields[0] == ownModule {
			continue
		}
		if len(fields) != 2 {
			return nil, fmt.Errorf("go list line %q is not `path version`; a version-less module (e.g. a local replace) needs a human ruling", line)
		}
		if seen[fields[0]] {
			continue
		}
		seen[fields[0]] = true
		mods = append(mods, module{path: fields[0], version: fields[1]})
	}
	sort.Slice(mods, func(i, j int) bool { return mods[i].path < mods[j].path })
	return mods, nil
}

// classify reads the module's vendored license text, fingerprints the license,
// and extracts the copyright line. Both a missing license file and an
// unrecognized text are hard errors: they mean the closure contains something
// nobody has ruled on.
func classify(m *module) error {
	var text string
	for _, name := range licenseFileNames {
		p := filepath.Join("vendor", filepath.FromSlash(m.path), name)
		b, err := os.ReadFile(p)
		if err == nil {
			text = string(b)
			m.licenseFile = "vendor/" + m.path + "/" + name
			break
		}
	}
	if text == "" {
		return fmt.Errorf("%s@%s: no license file found under vendor/%s/", m.path, m.version, m.path)
	}

	switch {
	case strings.Contains(text, "Apache License") && strings.Contains(text, "Version 2.0"):
		m.license = "Apache-2.0"
	case strings.Contains(text, "Permission is hereby granted, free of charge"):
		m.license = "MIT"
	case strings.Contains(text, "Permission to use, copy, modify, and/or distribute this software for any purpose"):
		m.license = "ISC"
	case strings.Contains(text, "Redistribution and use in source and binary forms"):
		if strings.Contains(text, "Neither the name") {
			m.license = "BSD-3-Clause"
		} else {
			m.license = "BSD-2-Clause"
		}
	default:
		return fmt.Errorf("%s@%s: unrecognized license text in %s", m.path, m.version, m.licenseFile)
	}

	m.copyright = copyrightLine(text, m.path)
	return nil
}

// copyrightLine returns the first real copyright line from a license text,
// skipping the bracketed "[yyyy] [name of copyright owner]" placeholder that
// the verbatim Apache-2.0 template carries in its appendix. Falls back to an
// attribution derived from the module path when the text names no owner.
func copyrightLine(text, modPath string) string {
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Copyright") && !strings.Contains(line, "[yyyy]") {
			return line
		}
	}
	parts := strings.Split(modPath, "/")
	owner := parts[len(parts)-1]
	if len(parts) >= 2 {
		owner = parts[1]
	}
	return "Copyright the " + owner + " authors"
}

// render produces the full THIRD-PARTY-NOTICES.md content for the given
// module set. Output is deterministic: modules are sorted by path and all
// prose is fixed, so check mode can diff byte for byte.
func render(mods []module) string {
	var b strings.Builder

	b.WriteString(`# Third-Party Notices

The EzClassify Go SDK (` + "`" + ownModule + "`" + `) is licensed under the
Apache License 2.0 (see ` + "`LICENSE`" + `). It depends at runtime on, and vendors
the source of, the third-party Go modules listed below. Each remains under
its own license; the texts vendored under ` + "`vendor/`" + ` govern your use of
those components.

This file covers **runtime** dependencies only. The codegen-only ` + "`oapi-codegen`" + `
v2 command used by ` + "`go generate`" + ` is not imported by the SDK, not redistributed
in a consuming build, and is omitted.

The runtime set is the transitive dependency closure of the public SDK
package, resolved from ` + "`vendor/`" + `:

` + "```" + `
go list -mod=vendor -deps -f '{{if .Module}}{{.Module.Path}}{{end}}' .
` + "```" + `

This file is generated by ` + "`tools/notices`" + ` from that closure and the vendored
license texts; CI fails if it goes stale. Do not edit by hand; regenerate with
` + "`make notices`" + ` after a dependency change.

License texts live alongside each module under ` + "`vendor/`" + `:

`)
	for _, m := range mods {
		b.WriteString("- `" + m.licenseFile + "`\n")
	}

	b.WriteString(`
## Runtime dependencies

| Module | Version | License | Copyright |
| --- | --- | --- | --- |
`)
	for _, m := range mods {
		b.WriteString("| `" + m.path + "` | " + m.version + " | " + m.license + " | " + m.copyright + " |\n")
	}

	for _, m := range mods {
		note, ok := moduleNotes[m.path]
		if !ok {
			note = "Production dependency of the SDK."
		}
		b.WriteString("\n### " + m.path + "\n\n")
		b.WriteString(note + "\n")
		b.WriteString("Licensed under the **" + spelledOut(m.license) + "** (see `" + m.licenseFile + "`).\n")
	}

	return b.String()
}

// spelledOut expands an SPDX identifier to the license's conventional long
// name for the per-module prose sections.
func spelledOut(spdx string) string {
	switch spdx {
	case "MIT":
		return "MIT License"
	case "Apache-2.0":
		return "Apache License, Version 2.0"
	case "BSD-2-Clause":
		return "BSD 2-Clause License"
	case "BSD-3-Clause":
		return "BSD 3-Clause License"
	case "ISC":
		return "ISC License"
	default:
		return spdx
	}
}
