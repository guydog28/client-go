package ezclassify

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
)

type generationManifest struct {
	OpenAPISHA256 string `json:"openapiSha256"`
	Generators    struct {
		Go struct {
			Version string `json:"version"`
			SHA256  string `json:"sha256"`
		} `json:"go"`
	} `json:"generators"`
	GeneratedPaths []string `json:"generatedPaths"`
}

func TestGeneratedWireOwnership(t *testing.T) {
	manifestBytes, err := os.ReadFile("spec/MANIFEST.json")
	if err != nil {
		t.Fatal(err)
	}
	var manifest generationManifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		t.Fatalf("parse spec/MANIFEST.json: %v", err)
	}

	wantGenerated := []string{"internal/api/api.gen.go"}
	if !reflect.DeepEqual(manifest.GeneratedPaths, wantGenerated) {
		t.Fatalf("generated paths = %v, want %v", manifest.GeneratedPaths, wantGenerated)
	}

	generated, err := os.Open(wantGenerated[0])
	if err != nil {
		t.Fatal(err)
	}
	defer generated.Close()
	scanner := bufio.NewScanner(generated)
	if !scanner.Scan() {
		t.Fatalf("read generated ownership header: %v", scanner.Err())
	}
	wantHeader := fmt.Sprintf(
		"// Code generated from openapi.yaml@%s by oapi-codegen@%s@%s. DO NOT EDIT.",
		manifest.OpenAPISHA256,
		manifest.Generators.Go.Version,
		manifest.Generators.Go.SHA256,
	)
	if scanner.Text() != wantHeader {
		t.Fatalf("generated ownership header = %q, want %q", scanner.Text(), wantHeader)
	}

	assertPublicWireTypesAreAliases(t)
	assertRootStructBoundary(t)
}

func assertPublicWireTypesAreAliases(t *testing.T) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "types.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, declaration := range file.Decls {
		general, ok := declaration.(*ast.GenDecl)
		if !ok || general.Tok != token.TYPE {
			continue
		}
		for _, spec := range general.Specs {
			typeSpec := spec.(*ast.TypeSpec)
			selector, ok := typeSpec.Type.(*ast.SelectorExpr)
			if !typeSpec.Assign.IsValid() || !ok {
				t.Errorf("types.go declaration %s must remain an alias to internal/api", typeSpec.Name)
				continue
			}
			ident, fromAPI := selector.X.(*ast.Ident)
			if !fromAPI || ident.Name != "api" {
				t.Errorf("types.go declaration %s must remain an alias to internal/api", typeSpec.Name)
			}
		}
	}
}

func assertRootStructBoundary(t *testing.T) {
	t.Helper()
	allowed := map[string]bool{
		"APIError":               true,
		"Client":                 true,
		"ContextNotFoundError":   true,
		"CtxAccess":              true,
		"CtxCrosswalk":           true,
		"CtxDocument":            true,
		"CtxGuide":               true,
		"CtxScheme":              true,
		"CtxSchemes":             true,
		"EzClassifyContext":      true,
		"InvalidInputError":      true,
		"MalformedResponseError": true,
		"PingResult":             true,
		"RestClientCredentials":  true,
		"callResult":             true,
		"clientConfig":           true,
		"deadlineBody":           true,
		"fieldPlan":              true,
		"limitedBody":            true,
		"responseCache":          true,
		"responseCacheEntry":     true,
		"responseStatus":         true,
		"responseStatusKey":      true,
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	var unexpected []string
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), entry.Name(), nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, declaration := range file.Decls {
			general, ok := declaration.(*ast.GenDecl)
			if !ok || general.Tok != token.TYPE {
				continue
			}
			for _, spec := range general.Specs {
				typeSpec := spec.(*ast.TypeSpec)
				if _, isStruct := typeSpec.Type.(*ast.StructType); isStruct && !allowed[typeSpec.Name.Name] {
					unexpected = append(unexpected, filepath.Join(entry.Name(), typeSpec.Name.Name))
				}
			}
		}
	}
	sort.Strings(unexpected)
	if len(unexpected) > 0 {
		t.Fatalf("unexpected root structs could shadow generated wire DTOs: %v", unexpected)
	}
}

// TestEveryWireEnumDecodesAnyString guards response openness for enums at the
// declaration level, so it covers enums added by a future regen without a
// payload for each. The contract is open-response: a value a newer server adds
// to an enum must decode. A generated enum is a string-backed type with a
// Valid() membership check, which encoding/json fills with any string, so the
// guard is that every such type stays string-backed and that nothing, in the
// generated file or in handwritten code, gives one an UnmarshalJSON that could
// reject an unknown member.
func TestEveryWireEnumDecodesAnyString(t *testing.T) {
	fset := token.NewFileSet()
	underlying := map[string]string{}
	enums := map[string]bool{}
	unmarshalers := map[string]string{}
	for _, dir := range []string{".", "internal/api"} {
		paths, err := filepath.Glob(filepath.Join(dir, "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		for _, path := range paths {
			if strings.HasSuffix(path, "_test.go") {
				continue
			}
			file, err := parser.ParseFile(fset, path, nil, parser.SkipObjectResolution)
			if err != nil {
				t.Fatalf("parse %s: %v", path, err)
			}
			for _, decl := range file.Decls {
				switch d := decl.(type) {
				case *ast.GenDecl:
					for _, spec := range d.Specs {
						if ts, ok := spec.(*ast.TypeSpec); ok {
							if ident, ok := ts.Type.(*ast.Ident); ok {
								underlying[ts.Name.Name] = ident.Name
							}
						}
					}
				case *ast.FuncDecl:
					if d.Recv == nil || len(d.Recv.List) != 1 {
						continue
					}
					recv := d.Recv.List[0].Type
					if star, ok := recv.(*ast.StarExpr); ok {
						recv = star.X
					}
					ident, ok := recv.(*ast.Ident)
					if !ok {
						continue
					}
					switch d.Name.Name {
					case "Valid":
						enums[ident.Name] = true
					case "UnmarshalJSON", "UnmarshalText":
						unmarshalers[ident.Name] = fset.Position(d.Pos()).String()
					}
				}
			}
		}
	}

	if len(enums) == 0 || !enums["DocumentOverallVerdict"] {
		t.Fatalf("found %d enums and not DocumentOverallVerdict; the scan no longer sees the generated file", len(enums))
	}
	for name := range enums {
		if underlying[name] != "string" {
			t.Errorf("enum %s is backed by %q, want string so an unknown member decodes", name, underlying[name])
		}
		if at, ok := unmarshalers[name]; ok {
			t.Errorf("enum %s has a custom unmarshaler at %s that could reject a member a newer server adds", name, at)
		}
	}
}
