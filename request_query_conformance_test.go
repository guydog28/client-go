package ezclassify_test

// Closed-request gate for query parameters.
//
// The contract refuses a query parameter an operation does not declare with a
// 400, so a client that adds one turns a working call into a failing one.
// Every operation of the vendored contract is driven through the real client
// with every optional argument set, and the query string it puts on the wire is
// compared with the parameters the contract declares for that operation. The
// body half of the same rule lives in request_schema_conformance_test.go.

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"

	"go.yaml.in/yaml/v3"

	"github.com/ezclassify/client-go"
)

// specDocument is the vendored OpenAPI document as decoded JSON.
var specDocument = sync.OnceValues(func() (map[string]any, error) {
	raw, err := os.ReadFile(requestSchemaSpecPath)
	if err != nil {
		return nil, err
	}
	var root yaml.Node
	if err := yaml.Unmarshal(raw, &root); err != nil {
		return nil, fmt.Errorf("parse %s: %w", requestSchemaSpecPath, err)
	}
	document, err := yamlToJSON(&root)
	if err != nil {
		return nil, err
	}
	doc, _ := document.(map[string]any)
	if doc == nil {
		return nil, fmt.Errorf("%s is not an object", requestSchemaSpecPath)
	}
	return doc, nil
})

// declaredQueryParameters maps each operation of the contract to the query
// parameters it declares.
var declaredQueryParameters = sync.OnceValues(func() (map[operationKey]map[string]bool, error) {
	doc, err := specDocument()
	if err != nil {
		return nil, err
	}
	paths, _ := doc["paths"].(map[string]any)
	if paths == nil {
		return nil, fmt.Errorf("%s declares no paths", requestSchemaSpecPath)
	}
	operations := map[operationKey]map[string]bool{}
	for path, item := range paths {
		methods, _ := item.(map[string]any)
		for method, operation := range methods {
			op, ok := operation.(map[string]any)
			if !ok {
				continue
			}
			names := map[string]bool{}
			parameters, _ := op["parameters"].([]any)
			for _, entry := range parameters {
				parameter, _ := entry.(map[string]any)
				if ref, ok := parameter["$ref"].(string); ok {
					resolved, err := resolveParameter(doc, ref)
					if err != nil {
						return nil, err
					}
					parameter = resolved
				}
				if parameter["in"] == "query" {
					name, _ := parameter["name"].(string)
					names[name] = true
				}
			}
			operations[operationKey{method, path}] = names
		}
	}
	return operations, nil
})

func resolveParameter(doc map[string]any, ref string) (map[string]any, error) {
	const prefix = "#/components/parameters/"
	name, ok := strings.CutPrefix(ref, prefix)
	if !ok {
		return nil, fmt.Errorf("parameter reference %q is not under %s", ref, prefix)
	}
	parameter, _ := lookup(doc, "components", "parameters", name).(map[string]any)
	if parameter == nil {
		return nil, fmt.Errorf("parameter reference %q does not resolve", ref)
	}
	return parameter, nil
}

// readCase drives one client call and names the contract operation it sends.
type readCase struct {
	name     string
	method   string
	template string
	invoke   func(ctx context.Context, c *ezclassify.Client, cx *ezclassify.EzClassifyContext) error
}

// readCases holds one case per client call that sends no body, each with every
// optional argument set. Calls that send a body are covered by requestCases.
var readCases = []readCase{
	{"health", "get", "/v1/health", func(ctx context.Context, c *ezclassify.Client, _ *ezclassify.EzClassifyContext) error {
		_, err := c.Ping(ctx)
		return err
	}},
	{"list-contexts", "get", "/v1/ctx", func(ctx context.Context, c *ezclassify.Client, _ *ezclassify.EzClassifyContext) error {
		_, err := c.ListContexts(ctx)
		return err
	}},
	{"list-versions", "get", "/v1/ctx/{ctxId}", func(ctx context.Context, c *ezclassify.Client, _ *ezclassify.EzClassifyContext) error {
		_, err := c.ListContextVersions(ctx, requestCaseContext)
		return err
	}},
	{"resolve", "get", "/v1/ctx/{ctxId}/{version}", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		return cx.Resolve(ctx)
	}},
	{"scheme.summary", "get", requestRoutePrefix + "scheme", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Summary(ctx)
		return err
	}},
	{"scheme.concepts", "get", requestRoutePrefix + "scheme/concepts", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Concepts(ctx)
		return err
	}},
	{"scheme.categories", "get", requestRoutePrefix + "scheme/categories", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Categories(ctx)
		return err
	}},
	{"scheme.formats", "get", requestRoutePrefix + "scheme/formats", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Formats(ctx)
		return err
	}},
	{"scheme.grammar", "get", requestRoutePrefix + "scheme/grammar", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Grammar(ctx)
		return err
	}},
	{"scheme.scopes", "get", requestRoutePrefix + "scheme/scopes", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Scopes(ctx)
		return err
	}},
	{"scheme.labelPresets", "get", requestRoutePrefix + "scheme/label-presets", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").LabelPresets(ctx)
		return err
	}},
	{"scheme.describeConcept", "get", requestRoutePrefix + "scheme/describe-concept", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").DescribeConcept(ctx, "HIGH", "level", ptr("short"))
		return err
	}},
	{"access.rules", "get", requestRoutePrefix + "access/rules", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Access.Rules(ctx)
		return err
	}},
	{"crosswalk.describe", "get", requestRoutePrefix + "crosswalk/describe", func(ctx context.Context, _ *ezclassify.Client, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Crosswalk.Describe(ctx)
		return err
	}},
}

// sendRead drives one read case against a stub that answers every request with
// an empty object and returns the single request it sent.
func sendRead(t *testing.T, c readCase) sentRequest {
	t.Helper()
	var mu sync.Mutex
	var sent []sentRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		sent = append(sent, sentRequest{method: r.Method, path: r.URL.Path, query: r.URL.Query(), body: body})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{}")
	}))
	defer srv.Close()

	client := ezclassify.New(srv.URL)
	_ = c.invoke(context.Background(), client, ezclassify.ConnectWith(client, requestCaseContext, requestCaseVersion))

	mu.Lock()
	defer mu.Unlock()
	if len(sent) != 1 {
		t.Fatalf("%s sent %d requests; expected exactly one", c.name, len(sent))
	}
	return sent[0]
}

// unexpectedQuery returns the sent query parameters the operation does not
// declare, sorted.
func unexpectedQuery(sent url.Values, declared map[string]bool) []string {
	var extra []string
	for name := range sent {
		if !declared[name] {
			extra = append(extra, name)
		}
	}
	sort.Strings(extra)
	return extra
}

func TestReadRequestsSendOnlyDeclaredQueryParameters(t *testing.T) {
	declared, err := declaredQueryParameters()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range readCases {
		t.Run(c.name, func(t *testing.T) {
			operation, ok := declared[operationKey{c.method, c.template}]
			if !ok {
				t.Fatalf("the contract declares no %s %s", strings.ToUpper(c.method), c.template)
			}
			request := sendRead(t, c)
			if extra := unexpectedQuery(request.query, operation); len(extra) > 0 {
				t.Fatalf("%s sent query parameters the contract does not declare, which a server refuses with 400: %s", c.name, strings.Join(extra, ", "))
			}
			var missing []string
			for name := range operation {
				if !request.query.Has(name) {
					missing = append(missing, name)
				}
			}
			sort.Strings(missing)
			if len(missing) > 0 {
				t.Fatalf("%s never sent declared query parameters: %s", c.name, strings.Join(missing, ", "))
			}
		})
	}
}

// Every operation that takes a body takes its input there alone.
func TestBodyRequestsSendOnlyDeclaredQueryParameters(t *testing.T) {
	declared, err := declaredQueryParameters()
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range requestCases {
		t.Run(c.name, func(t *testing.T) {
			request := send(t, c)
			operation := declared[operationKey{"post", requestRoutePrefix + c.route}]
			if extra := unexpectedQuery(request.query, operation); len(extra) > 0 {
				t.Fatalf("%s sent query parameters the contract does not declare, which a server refuses with 400: %s", c.name, strings.Join(extra, ", "))
			}
		})
	}
}

// A new operation cannot ship without its query string being checked.
func TestEveryContractReadOperationHasACase(t *testing.T) {
	declared, err := declaredQueryParameters()
	if err != nil {
		t.Fatal(err)
	}
	covered := map[operationKey]bool{}
	for _, c := range readCases {
		covered[operationKey{c.method, c.template}] = true
	}
	for _, c := range requestCases {
		covered[operationKey{"post", requestRoutePrefix + c.route}] = true
	}
	var missing []string
	for key := range declared {
		if !covered[key] {
			missing = append(missing, strings.ToUpper(key.method)+" "+key.path)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Errorf("no query-string case for: %s", strings.Join(missing, ", "))
	}
}

// The gate must be able to fail: a stub request that adds a parameter the
// contract does not declare is reported by name.
func TestUnexpectedQueryNamesAnUndeclaredParameter(t *testing.T) {
	sent := url.Values{"scheme": {"acme"}, "verbose": {"1"}}
	got := unexpectedQuery(sent, map[string]bool{"scheme": true})
	if len(got) != 1 || got[0] != "verbose" {
		t.Fatalf("unexpectedQuery = %v, want [verbose]", got)
	}
}
