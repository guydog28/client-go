// Command dependencyage is the dependency quarantine gate. It refuses
// third-party Go modules whose pinned version was published inside a
// quarantine window, so a freshly compromised release has to sit on the module
// proxy for a while before it can enter this repository's dependency tree.
//
// It reads go.sum rather than go.mod: go.sum names every module version the
// build graph resolved, including transitive ones, which is the set that
// matters. Each version is checked against the publish time the module proxy
// records for it (<proxy>/<module>/@v/<version>.info). Modules under the
// ezclassify organization are first-party and exempt.
//
// The gate fails closed: a proxy that cannot be read, a response with no
// publish time, or an unreadable go.sum is a failure, never a silent pass.
// Every request has a timeout and the whole run has an overall bound, so a
// proxy that accepts connections and never answers cannot hang the job.
//
// Usage:
//
//	go run ./tools/dependencyage [-days N] [-gosum PATH] [-proxy URL]
//
// Each flag has an environment form: EZCLASSIFY_DEPENDENCY_AGE_DAYS,
// EZCLASSIFY_DEPENDENCY_AGE_GOSUM and EZCLASSIFY_DEPENDENCY_AGE_PROXY.
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// defaultQuarantineDays is long enough that a malicious release has usually
	// been reported and retracted, and short enough that ordinary upgrades are
	// not held up. Dependabot's cooldown matches it.
	defaultQuarantineDays = 7

	// defaultProxy is the public Go module proxy.
	defaultProxy = "https://proxy.golang.org"

	// firstPartyPrefix is the module path prefix owned by this project.
	firstPartyPrefix = "github.com/ezclassify/"

	// requestTimeout bounds one publish-time request.
	requestTimeout = 20 * time.Second

	// overallTimeout bounds the whole run; modules still unchecked at that
	// point are reported and fail the gate.
	overallTimeout = 15 * time.Minute

	// concurrency is the number of requests in flight at once.
	concurrency = 4

	day = 24 * time.Hour
)

// module is one module version pinned by go.sum.
type module struct {
	Path    string
	Version string
}

func (m module) String() string { return m.Path + "@" + m.Version }

// finding is a module that failed the gate. Reason is set when the publish time
// could not be confirmed; otherwise Published is inside the window.
type finding struct {
	Module    module
	Published time.Time
	Clears    time.Time
	Reason    string
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("dependencyage", flag.ContinueOnError)
	flags.SetOutput(stderr)
	days := flags.String("days", envOr("EZCLASSIFY_DEPENDENCY_AGE_DAYS", strconv.Itoa(defaultQuarantineDays)), "quarantine window in days")
	gosum := flags.String("gosum", envOr("EZCLASSIFY_DEPENDENCY_AGE_GOSUM", "go.sum"), "path to go.sum")
	proxy := flags.String("proxy", envOr("EZCLASSIFY_DEPENDENCY_AGE_PROXY", defaultProxy), "module proxy base URL")
	if err := flags.Parse(args); err != nil {
		return err
	}
	window, err := strconv.ParseFloat(*days, 64)
	if err != nil || window < 0 {
		return fmt.Errorf("quarantine window must be a non-negative number of days, got %q", *days)
	}

	file, err := os.Open(*gosum)
	if err != nil {
		return fmt.Errorf("cannot read %s: %w", *gosum, err)
	}
	defer file.Close()
	modules, err := thirdPartyModules(file)
	if err != nil {
		return fmt.Errorf("cannot parse %s: %w", *gosum, err)
	}
	if window == 0 {
		fmt.Fprintf(stdout, "Dependency quarantine window is 0 days; nothing to check (%d module version(s)).\n", len(modules))
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), overallTimeout)
	defer cancel()
	findings := check(ctx, http.DefaultClient, strings.TrimRight(*proxy, "/"), modules, time.Now(), time.Duration(window*float64(day)))
	if len(findings) == 0 {
		fmt.Fprintf(stdout, "Dependency quarantine passed: %d third-party module version(s) are all older than %s day(s).\n", len(modules), *days)
		return nil
	}

	var unknown, fresh []finding
	for _, f := range findings {
		if f.Reason != "" {
			unknown = append(unknown, f)
		} else {
			fresh = append(fresh, f)
		}
	}
	if len(unknown) > 0 {
		fmt.Fprintf(stderr, "Dependency quarantine: could not confirm the publish time of %d module version(s), so the gate fails closed.\n", len(unknown))
		for _, f := range unknown {
			fmt.Fprintf(stderr, "  %s: %s\n", f.Module, f.Reason)
		}
	}
	if len(fresh) > 0 {
		fmt.Fprintf(stderr, "Dependency quarantine: %d third-party module version(s) were published less than %s day(s) ago.\n", len(fresh), *days)
		for _, f := range fresh {
			fmt.Fprintf(stderr, "  %s published %s, clears %s\n", f.Module, f.Published.UTC().Format(time.RFC3339), f.Clears.UTC().Format(time.RFC3339))
		}
		fmt.Fprintln(stderr, "Wait out the window, or pin an older version, before taking these in.")
	}
	return errors.New("dependency quarantine failed")
}

func envOr(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}

// thirdPartyModules returns the distinct third-party module versions in a
// go.sum stream, sorted. Each line is "<module> <version>[/go.mod] <hash>".
func thirdPartyModules(r io.Reader) ([]module, error) {
	seen := map[module]bool{}
	scanner := bufio.NewScanner(r)
	for n := 1; scanner.Scan(); n++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("line %d: expected \"<module> <version> <hash>\"", n)
		}
		if strings.HasPrefix(fields[0], firstPartyPrefix) {
			continue
		}
		seen[module{Path: fields[0], Version: strings.TrimSuffix(fields[1], "/go.mod")}] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	modules := make([]module, 0, len(seen))
	for m := range seen {
		modules = append(modules, m)
	}
	sort.Slice(modules, func(i, j int) bool { return modules[i].String() < modules[j].String() })
	return modules, nil
}

// check looks up every module's publish time and returns the failures: modules
// published inside the window and modules whose publish time is unconfirmed.
func check(ctx context.Context, client *http.Client, proxy string, modules []module, now time.Time, window time.Duration) []finding {
	var (
		mu       sync.Mutex
		findings []finding
		wg       sync.WaitGroup
		slots    = make(chan struct{}, concurrency)
	)
	for _, m := range modules {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case slots <- struct{}{}:
				defer func() { <-slots }()
			case <-ctx.Done():
				mu.Lock()
				findings = append(findings, finding{Module: m, Reason: "the run passed its overall time bound before it was checked"})
				mu.Unlock()
				return
			}
			published, err := publishTime(ctx, client, proxy, m)
			mu.Lock()
			defer mu.Unlock()
			switch {
			case err != nil:
				findings = append(findings, finding{Module: m, Reason: err.Error()})
			case now.Sub(published) < window:
				findings = append(findings, finding{Module: m, Published: published, Clears: published.Add(window)})
			}
		}()
	}
	wg.Wait()
	sort.Slice(findings, func(i, j int) bool { return findings[i].Module.String() < findings[j].Module.String() })
	return findings
}

// publishTime asks the module proxy when a version was published.
func publishTime(ctx context.Context, client *http.Client, proxy string, m module) (time.Time, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	url := proxy + "/" + escapePath(m.Path) + "/@v/" + escapePath(m.Version) + ".info"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return time.Time{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return time.Time{}, fmt.Errorf("could not reach the proxy: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("the proxy answered %d; cannot confirm the publish time", resp.StatusCode)
	}
	var info struct{ Time string }
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&info); err != nil {
		return time.Time{}, fmt.Errorf("the proxy gave an unreadable answer: %w", err)
	}
	if info.Time == "" {
		return time.Time{}, errors.New("the proxy returned no publish time")
	}
	published, err := time.Parse(time.RFC3339, info.Time)
	if err != nil {
		return time.Time{}, fmt.Errorf("the proxy gave an unreadable publish time %q", info.Time)
	}
	return published, nil
}

// escapePath applies the module proxy's case encoding: each upper-case letter
// becomes "!" followed by its lower-case form.
func escapePath(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			b.WriteByte('!')
			b.WriteRune(r + ('a' - 'A'))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
