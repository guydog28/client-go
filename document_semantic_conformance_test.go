package ezclassify_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	ezclassify "github.com/ezclassify/sdk-go"
	"github.com/ezclassify/sdk-go/internal/subsetmatch"
)

var (
	lowercaseSHA1   = regexp.MustCompile(`^[0-9a-f]{40}$`)
	lowercaseSHA256 = regexp.MustCompile(`^[0-9a-f]{64}$`)
	semanticVersion = regexp.MustCompile(`^\d+\.\d+\.`)
)

// requiredSemanticSpecFiles are the mirror members this test cannot do its job
// without. Completeness of the mirror is checked against the publisher's own
// specBundleSha256 recomputed below rather than against a frozen file count or
// inventory hash, which a legitimate re-vendor of the bot-owned tree would turn
// red.
var requiredSemanticSpecFiles = []string{
	"spec/openapi.yaml",
	"spec/semantic/v1/document-assessment.contract.json",
	"spec/semantic/v1/fixtures/requests.json",
	"spec/semantic/v1/fixtures/recordings.json",
	"spec/semantic/v1/fixtures/clause-coverage.json",
}

type documentSemanticContract struct {
	Cases []documentSemanticCase `json:"cases"`
}

type documentSemanticCase struct {
	Assertions     []documentSemanticAssertion `json:"assertions"`
	Implementation string                      `json:"implementation"`
	Number         int                         `json:"number"`
	Operation      string                      `json:"operation"`
	RequestID      string                      `json:"requestId"`
	Required       bool                        `json:"required"`
	SemanticID     string                      `json:"semanticId"`
}

type documentSemanticAssertion struct {
	Operator string `json:"operator"`
	Path     string `json:"path"`
	Value    any    `json:"value"`
}

type documentSemanticRecordings struct {
	Status string                               `json:"status"`
	Cases  map[string]documentSemanticRecording `json:"cases"`
}

// documentSemanticFixtureContext is the published variant map in
// fixtures/context.json: the base fixtures and, per named variant, the ones it
// replaces.
type documentSemanticFixtureContext struct {
	SchemeFixture string                                    `json:"schemeFixture"`
	Variants      map[string]documentSemanticFixtureVariant `json:"variants"`
}

type documentSemanticFixtureVariant struct {
	SchemeFixture *string `json:"schemeFixture"`
}

// documentSemanticSchemeFixture is the part of a scheme fixture that names the
// Context it stands for.
type documentSemanticSchemeFixture struct {
	Scheme *struct {
		ID      string `json:"id"`
		Version string `json:"version"`
	} `json:"scheme"`
}

// semanticContextKey is the Context a replayed request is sent to.
type semanticContextKey struct {
	ID      string
	Version string
}

type documentSemanticRecording struct {
	Execution string          `json:"execution"`
	Output    json.RawMessage `json:"output"`
}

type documentSemanticManifest struct {
	ContractsSHA     string            `json:"contractsSha"`
	OpenAPIVersion   string            `json:"openapiVersion"`
	OpenAPISHA256    string            `json:"openapiSha256"`
	SpecBundleSHA256 string            `json:"specBundleSha256"`
	SpecDigest       string            `json:"specDigest"`
	SourcePackages   map[string]string `json:"sourcePackages"`
}

// TestDocumentSemanticConformance replays the frozen document contract through
// the public Context HTTP surface and evaluates every canonical assertion.
func TestDocumentSemanticConformance(t *testing.T) {
	verifyDocumentSemanticMirror(t)
	root := filepath.Join("spec", "semantic", "v1")
	contract := readSemanticFixture[documentSemanticContract](t, filepath.Join(root, "document-assessment.contract.json"))
	requests := readSemanticFixture[map[string]json.RawMessage](t, filepath.Join(root, "fixtures", "requests.json"))
	recordings := readSemanticFixture[documentSemanticRecordings](t, filepath.Join(root, "fixtures", "recordings.json"))
	coverage := readSemanticFixture[map[string][]json.RawMessage](t, filepath.Join(root, "fixtures", "clause-coverage.json"))
	fixtureContext := readSemanticFixture[documentSemanticFixtureContext](t, filepath.Join(root, "fixtures", "context.json"))

	if recordings.Status != "recorded-operation-output" {
		t.Fatalf("recordings status = %q, want recorded-operation-output", recordings.Status)
	}
	// The corpus is the count. Every inventory has to name exactly the
	// contract's cases, so a dropped or added case fails without a pinned
	// total to move.
	caseCount := len(contract.Cases)
	if caseCount == 0 || len(requests) != caseCount || len(recordings.Cases) != caseCount || len(coverage) != caseCount {
		t.Fatalf("canonical inventory sizes: contract=%d requests=%d recordings=%d coverage=%d, want the contract's count in each", caseCount, len(requests), len(recordings.Cases), len(coverage))
	}

	// Case numbers are stable identities: a retired case leaves a gap rather
	// than renumbering the rest. They still run in ascending order, and each
	// request id is the number's own.
	semanticIDs := make(map[string]struct{}, caseCount)
	previous := 0
	for index, testCase := range contract.Cases {
		wantID := fmt.Sprintf("case-%02d", testCase.Number)
		if testCase.Number <= previous || testCase.RequestID != wantID {
			t.Fatalf("case %d inventory entry = number %d request %q, want a number above %d and request %q", index, testCase.Number, testCase.RequestID, previous, wantID)
		}
		previous = testCase.Number
		if !testCase.Required || testCase.Implementation != "recorded" || len(testCase.Assertions) == 0 {
			t.Fatalf("%s is not a required recorded case with assertions", testCase.RequestID)
		}
		if _, duplicate := semanticIDs[testCase.SemanticID]; duplicate || testCase.SemanticID == "" {
			t.Fatalf("%s has empty or duplicate semanticId %q", testCase.RequestID, testCase.SemanticID)
		}
		semanticIDs[testCase.SemanticID] = struct{}{}
		if _, ok := requests[testCase.RequestID]; !ok {
			t.Fatalf("%s has no canonical request", testCase.RequestID)
		}
		if clauses, ok := coverage[testCase.RequestID]; !ok || len(clauses) == 0 {
			t.Fatalf("%s has no clause coverage", testCase.RequestID)
		}

		t.Run(testCase.RequestID+"/"+testCase.SemanticID, func(t *testing.T) {
			recording, ok := recordings.Cases[testCase.RequestID]
			if !ok || recording.Execution == "" || len(recording.Output) == 0 {
				t.Fatal("missing or malformed canonical recording")
			}
			key := semanticContextFor(t, root, fixtureContext, requests[testCase.RequestID])
			assertVariantMatchesRecording(t, requests[testCase.RequestID], recording.Output, key)
			output := replayThroughPublicContext(t, key, testCase.Operation, requests[testCase.RequestID], recording.Output)
			for _, assertion := range testCase.Assertions {
				if !evaluateDocumentSemanticAssertion(output, assertion) {
					t.Errorf("assertion failed: %s %s %v", assertion.Operator, assertion.Path, assertion.Value)
				}
			}
		})
	}
}

// semanticContextFor resolves the Context a semantic request runs against from
// the published variant map. The base fixtures apply unless the request names a
// fixtureVariant, whose scheme fixture then replaces the base one. The Context's
// identity is the one that scheme fixture declares.
func semanticContextFor(t *testing.T, root string, fixtureContext documentSemanticFixtureContext, rawRequest json.RawMessage) semanticContextKey {
	t.Helper()
	var request struct {
		FixtureVariant *string `json:"fixtureVariant"`
	}
	if err := json.Unmarshal(rawRequest, &request); err != nil {
		t.Fatalf("decode canonical request recipe: %v", err)
	}
	schemeFixture := fixtureContext.SchemeFixture
	if request.FixtureVariant != nil {
		variant, ok := fixtureContext.Variants[*request.FixtureVariant]
		if !ok {
			t.Fatalf("request names undeclared fixture variant %q", *request.FixtureVariant)
		}
		if variant.SchemeFixture != nil {
			schemeFixture = *variant.SchemeFixture
		}
	}
	path := filepath.Join(root, "fixtures", filepath.FromSlash(schemeFixture))
	fixture := readSemanticFixture[documentSemanticSchemeFixture](t, path)
	if fixture.Scheme == nil || fixture.Scheme.ID == "" || fixture.Scheme.Version == "" {
		t.Fatalf("%s declares no scheme identity", path)
	}
	return semanticContextKey{ID: fixture.Scheme.ID, Version: fixture.Scheme.Version}
}

// assertVariantMatchesRecording holds a variant case to the Context its
// recording names, which proves the runner read the variant map the way the
// recorder did.
func assertVariantMatchesRecording(t *testing.T, rawRequest, rawOutput json.RawMessage, key semanticContextKey) {
	t.Helper()
	var request struct {
		FixtureVariant *string `json:"fixtureVariant"`
	}
	var output struct {
		Resolution *struct {
			ContextID *string `json:"contextId"`
		} `json:"resolution"`
	}
	if err := json.Unmarshal(rawRequest, &request); err != nil || request.FixtureVariant == nil {
		return
	}
	if err := json.Unmarshal(rawOutput, &output); err != nil || output.Resolution == nil || output.Resolution.ContextID == nil {
		return
	}
	if *output.Resolution.ContextID != key.ID {
		t.Fatalf("variant %q ran against Context %q, recorded against %q", *request.FixtureVariant, key.ID, *output.Resolution.ContextID)
	}
}

func readSemanticFixture[T any](t *testing.T, path string) T {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
	return value
}

// replayThroughPublicContext invokes the exported Context operation, captures
// the real serialized body, and round-trips the concrete response or APIError.
func replayThroughPublicContext(t *testing.T, key semanticContextKey, operation string, rawRequest, rawOutput json.RawMessage) any {
	t.Helper()
	var canonical map[string]any
	if err := json.Unmarshal(rawOutput, &canonical); err != nil {
		t.Fatalf("decode recording output: %v", err)
	}
	payload := canonicalPayload(t, operation, rawRequest)
	if operation == "auditDocument" && payload["text"] == nil {
		payload["text"] = json.RawMessage(`" \t\n"`)
	}
	invalidUnicode := operation != "documentAccessRequirements" && strings.Contains(string(rawRequest), "invalidInputRecipes")
	if invalidUnicode {
		payload["text"] = json.RawMessage(`"\ufffd"`)
	}

	server := newSemanticContextServer(t, key, operation, canonical)
	cx := ezclassify.Connect(server.URL, key.ID, key.Version)
	var value any
	var err error
	switch operation {
	case "auditDocument", "cliDocument":
		var request ezclassify.CtxDocumentAuditRequest
		decodePayload(t, payload, &request)
		if invalidUnicode {
			request.Text = string([]byte{0xff})
		}
		value, err = cx.Document.Audit(context.Background(), request)
	case "documentAccessRequirements":
		var request ezclassify.CtxAccessDescribeRequirementsRequest
		decodePayload(t, payload, &request)
		value, err = cx.Access.DescribeRequirements(context.Background(), request)
	default:
		t.Fatalf("unsupported canonical document operation %q", operation)
	}
	if got := cx.ResolvedVersion(); !invalidUnicode && got != key.Version {
		t.Fatalf("ResolvedVersion() = %q, want %s", got, key.Version)
	}

	if canonicalError, isError := canonical["error"].(map[string]any); isError {
		if !invalidUnicode {
			assertCapturedPayload(t, payload, server.request())
		}
		return canonicalErrorFromPublicSurface(t, err, canonicalError, invalidUnicode)
	}
	if err != nil {
		t.Fatalf("public Context call: %v", err)
	}
	if invalidUnicode {
		t.Fatal("invalid Unicode unexpectedly succeeded")
	}
	assertCapturedPayload(t, payload, server.request())
	replayed := marshalSemanticValue(t, value)
	if !subsetmatch.Match(replayed, canonical) {
		t.Fatalf("%s public response lost canonical fields", operation)
	}
	return replayed
}

func canonicalPayload(t *testing.T, operation string, raw json.RawMessage) map[string]json.RawMessage {
	t.Helper()
	var recipe map[string]json.RawMessage
	if err := json.Unmarshal(raw, &recipe); err != nil {
		t.Fatalf("decode canonical request recipe: %v", err)
	}
	keys := map[string][]string{"auditDocument": {"text", "guideDeterminations", "responseProfile"}, "cliDocument": {"text", "guideDeterminations", "responseProfile"}, "documentAccessRequirements": {"marking", "assertions"}}[operation]
	payload := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		if raw, ok := recipe[key]; ok {
			payload[key] = raw
		}
	}
	return payload
}
func decodePayload(t *testing.T, payload map[string]json.RawMessage, target any) {
	t.Helper()
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(data, target); err != nil {
		t.Fatalf("decode public request DTO: %v", err)
	}
}
func marshalSemanticValue(t *testing.T, value any) any {
	t.Helper()
	raw, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("encode public DTO: %v", err)
	}
	var output any
	if err := json.Unmarshal(raw, &output); err != nil {
		t.Fatalf("decode public DTO: %v", err)
	}
	return output
}

type semanticContextServer struct {
	*httptest.Server
	mu   sync.Mutex
	body []byte
}

func newSemanticContextServer(t *testing.T, key semanticContextKey, operation string, canonical map[string]any) *semanticContextServer {
	t.Helper()
	server := &semanticContextServer{}
	server.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		server.mu.Lock()
		server.body = body
		server.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", key.Version)
		expectedSuffix := map[string]string{"auditDocument": "/document/audit", "cliDocument": "/document/audit", "documentAccessRequirements": "/access/describe-requirements"}[operation]
		if !strings.HasPrefix(r.URL.Path, "/v1/ctx/"+key.ID+"/"+key.Version+"/") || !strings.HasSuffix(r.URL.Path, expectedSuffix) {
			t.Errorf("unexpected route %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if errorObject, isError := canonical["error"].(map[string]any); isError {
			code, detail, message := errorObject["code"].(string), errorObject["detail"].(string), errorObject["message"].(string)
			status := http.StatusBadRequest
			if strings.HasPrefix(detail, "stale-") {
				status = http.StatusConflict
			} else if code == "NO_ACCESS_POLICY" {
				status = http.StatusUnprocessableEntity
			}
			w.WriteHeader(status)
			_, _ = io.WriteString(w, fmt.Sprintf(`{"error":%q,"code":%q,"detail":%q}`, message, code, detail))
			return
		}
		raw, err := json.Marshal(canonical)
		if err != nil {
			t.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(raw)
	}))
	t.Cleanup(server.Close)
	return server
}
func (s *semanticContextServer) request() []byte {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]byte(nil), s.body...)
}

func assertCapturedPayload(t *testing.T, expected map[string]json.RawMessage, actual []byte) {
	t.Helper()
	var want, got any
	wantRaw, _ := json.Marshal(expected)
	if err := json.Unmarshal(wantRaw, &want); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(actual, &got); err != nil {
		t.Fatalf("decode captured request: %v", err)
	}
	if !subsetmatch.Match(got, want) {
		t.Fatalf("public Context request changed canonical payload: got %s want %s", actual, wantRaw)
	}
}
func canonicalErrorFromPublicSurface(t *testing.T, err error, canonical map[string]any, invalidUnicode bool) any {
	t.Helper()
	if invalidUnicode {
		var input *ezclassify.InvalidInputError
		if !errors.As(err, &input) || input.Detail != canonical["detail"] {
			t.Fatalf("local invalid Unicode error = %T(%v), want InvalidInputError detail %v", err, err, canonical["detail"])
		}
		return map[string]any{"error": map[string]any{"code": "INVALID_INPUT", "detail": input.Detail, "message": canonical["message"]}}
	}
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) || apiErr.Response == nil {
		t.Fatalf("HTTP operation error = %T(%v), want APIError with ErrorResponse", err, err)
	}
	if apiErr.Response.Code == nil || apiErr.Response.Detail == nil {
		t.Fatalf("APIError.Response lost code/detail: %+v", apiErr.Response)
	}
	if *apiErr.Response.Code != canonical["code"] || *apiErr.Response.Detail != canonical["detail"] || apiErr.Response.Error != canonical["message"] {
		t.Fatalf("APIError.Response = %+v, want %v", apiErr.Response, canonical)
	}
	return map[string]any{"error": map[string]any{"code": *apiErr.Response.Code, "detail": *apiErr.Response.Detail, "message": apiErr.Response.Error}}
}

// verifyDocumentSemanticMirror checks what the vendored spec bundle claims
// about where it came from, and whether the claim is true.
//
// The assertions are about SHAPE and INTERNAL CONSISTENCY, never about one
// blessed commit. spec/ is bot-owned: the spec bot re-vendors it and rewrites
// MANIFEST.json, and it writes nothing outside spec/. A literal digest pinned
// in this file could therefore only be repaired by a human, so it would go red
// on the next legitimate propagation. `specDigest` is computed over the spec
// bot's own bytes, so a change to the bot re-stales it without a single spec
// byte moving. Everything below survives a legitimate re-vendor and only goes
// red when the provenance is actually wrong.
func verifyDocumentSemanticMirror(t *testing.T) {
	t.Helper()
	manifestPath := filepath.Join("spec", "MANIFEST.json")
	manifest := readSemanticFixture[documentSemanticManifest](t, manifestPath)

	// An empty field means the manifest key was renamed out from under this
	// struct rather than that the value drifted. Say which, so the next
	// provenance rename reads as a rename instead of a mysterious mismatch.
	for field, value := range map[string]string{
		"contractsSha":     manifest.ContractsSHA,
		"openapiVersion":   manifest.OpenAPIVersion,
		"specDigest":       manifest.SpecDigest,
		"specBundleSha256": manifest.SpecBundleSHA256,
		"openapiSha256":    manifest.OpenAPISHA256,
	} {
		if value == "" {
			t.Fatalf("spec/MANIFEST.json has no %q field; the publisher renamed it and this test still reads the old key", field)
		}
	}
	if !lowercaseSHA1.MatchString(manifest.ContractsSHA) {
		t.Fatalf("contractsSha = %q, want a 40-character lowercase hex commit", manifest.ContractsSHA)
	}
	for field, value := range map[string]string{
		"specDigest":       manifest.SpecDigest,
		"specBundleSha256": manifest.SpecBundleSHA256,
		"openapiSha256":    manifest.OpenAPISHA256,
	} {
		if !lowercaseSHA256.MatchString(value) {
			t.Fatalf("%s = %q, want a 64-character lowercase hex SHA-256", field, value)
		}
	}
	if len(manifest.SourcePackages) == 0 {
		t.Fatal("spec/MANIFEST.json has no sourcePackages entries, so it does not say which packages the bundle was built from")
	}
	for name, version := range manifest.SourcePackages {
		if !strings.HasPrefix(name, "@ezclassify/") {
			t.Errorf("sourcePackages entry %q is not an @ezclassify package", name)
		}
		if !semanticVersion.MatchString(version) {
			t.Errorf("sourcePackages[%q] = %q, want a version string", name, version)
		}
	}

	// Two files, one fact. They are written in the same pass, so a
	// disagreement means the writer half-applied a change.
	if row := semanticSpecProvenanceRow(t, "contracts@"); row != manifest.ContractsSHA {
		t.Fatalf("SPEC.md contracts@ = %q, MANIFEST.json contractsSha = %q", row, manifest.ContractsSHA)
	}
	if url := semanticSpecProvenanceRow(t, "specSourceUrl"); !strings.Contains(url, "github.com/ezclassify/contracts") || !strings.Contains(url, manifest.ContractsSHA) {
		t.Fatalf("SPEC.md specSourceUrl = %q, want the contracts repo at %s", url, manifest.ContractsSHA)
	}

	files := publisherSpecFiles(t)
	present := make(map[string]struct{}, len(files))
	for _, path := range files {
		present[filepath.ToSlash(path)] = struct{}{}
	}
	for _, required := range requiredSemanticSpecFiles {
		if _, ok := present[required]; !ok {
			t.Fatalf("publisher spec inventory is missing %s; the mirror is incomplete", required)
		}
	}

	// Reproduce the publisher's path-plus-lowercase-file-hash bundle digest
	// over the whole inventory. This is what makes the mirror's completeness
	// checkable without pinning a file count: a file added, removed or edited
	// under spec/ that the publisher did not record lands here.
	hash := sha256.New()
	for _, path := range files {
		contents, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		contentHash := sha256.Sum256(contents)
		_, _ = io.WriteString(hash, filepath.ToSlash(path))
		_, _ = hash.Write([]byte{0})
		_, _ = io.WriteString(hash, hex.EncodeToString(contentHash[:]))
		_, _ = hash.Write([]byte{'\n'})
	}
	if actual := hex.EncodeToString(hash.Sum(nil)); actual != manifest.SpecBundleSHA256 {
		t.Fatalf("publisher-compatible spec bundle SHA over %d files = %s, manifest = %s", len(files), actual, manifest.SpecBundleSHA256)
	}
	openapi, err := os.ReadFile(filepath.Join("spec", "openapi.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if actual := sha256.Sum256(openapi); hex.EncodeToString(actual[:]) != manifest.OpenAPISHA256 {
		t.Fatal("openapi SHA disagrees with MANIFEST.json")
	}
	declared := regexp.MustCompile(`(?m)^ {2}version:\s+(\S+)`).FindSubmatch(openapi)
	if declared == nil {
		t.Fatal("spec/openapi.yaml declares no info.version")
	}
	if version := string(declared[1]); version != manifest.OpenAPIVersion {
		t.Fatalf("openapi.yaml info.version = %q, MANIFEST.json openapiVersion = %q", version, manifest.OpenAPIVersion)
	}
}

// semanticSpecProvenanceRow returns the value cell of a row in the SPEC.md
// provenance table, failing the test when the row is absent.
func semanticSpecProvenanceRow(t *testing.T, field string) string {
	t.Helper()
	value, present := semanticSpecProvenanceRowIfPresent(t, field)
	if !present {
		t.Fatalf("spec/SPEC.md has no %q provenance row", field)
	}
	return value
}

// semanticSpecProvenanceRowIfPresent returns the value cell of a row in the
// SPEC.md provenance table and whether that row exists. The column widths are
// re-padded on every publish, so match the row structure rather than a fixed
// number of spaces.
func semanticSpecProvenanceRowIfPresent(t *testing.T, field string) (string, bool) {
	t.Helper()
	specMD, err := os.ReadFile(filepath.Join("spec", "SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	row := regexp.MustCompile(`(?m)^\|\s*` + regexp.QuoteMeta("`"+field+"`") + `\s*\|([^|]*)\|`).FindSubmatch(specMD)
	if row == nil {
		return "", false
	}
	return strings.Trim(strings.TrimSpace(string(row[1])), "`"), true
}
func publisherSpecFiles(t *testing.T) []string {
	t.Helper()
	var files []string
	var add func(string, func(string) bool)
	add = func(root string, predicate func(string) bool) {
		entries, err := os.ReadDir(root)
		if err != nil {
			t.Fatal(err)
		}
		for _, entry := range entries {
			path := filepath.Join(root, entry.Name())
			if entry.IsDir() {
				add(path, predicate)
			} else if predicate(path) {
				files = append(files, path)
			}
		}
	}
	files = append(files, filepath.Join("spec", "openapi.yaml"))
	add(filepath.Join("spec", "json-schemas"), func(path string) bool { return strings.HasSuffix(path, ".schema.json") })
	for _, directory := range []string{"plans", "fixtures", "recordings"} {
		add(filepath.Join("spec", directory, "v1"), func(path string) bool { return strings.HasSuffix(path, ".json") })
	}
	add(filepath.Join("spec", "wire", "v1"), func(path string) bool { return strings.HasSuffix(path, ".json") })
	add(filepath.Join("spec", "cache", "v1"), func(string) bool { return true })
	add(filepath.Join("spec", "replayer-contract", "v1"), func(string) bool { return true })
	add(filepath.Join("spec", "semantic", "v1"), func(string) bool { return true })
	// The spec bot uses JavaScript localeCompare. Its source paths are
	// lowercase except the replayer README; case-fold before the byte tie-break
	// so Go produces the same corpus-before-README ordering.
	sort.Slice(files, func(i, j int) bool {
		left, right := filepath.ToSlash(files[i]), filepath.ToSlash(files[j])
		if foldedLeft, foldedRight := strings.ToLower(left), strings.ToLower(right); foldedLeft != foldedRight {
			return foldedLeft < foldedRight
		}
		return left < right
	})
	return files
}

func evaluateDocumentSemanticAssertion(recording any, assertion documentSemanticAssertion) bool {
	values := documentSemanticValuesAtPath(recording, assertion.Path)
	switch assertion.Operator {
	case "present":
		return len(values) > 0
	case "absent":
		return len(values) == 0
	case "equals":
		return len(values) > 0 && everySemanticValue(values, func(value any) bool { return reflect.DeepEqual(value, assertion.Value) })
	case "minItems":
		minimum, ok := assertion.Value.(float64)
		return ok && len(values) > 0 && everySemanticValue(values, func(value any) bool {
			items, ok := value.([]any)
			return ok && len(items) > 0 && len(items) >= int(minimum)
		})
	case "contains":
		return len(values) > 0 && everySemanticValue(values, func(value any) bool {
			items, ok := value.([]any)
			if !ok {
				return false
			}
			for _, item := range items {
				if reflect.DeepEqual(item, assertion.Value) {
					return true
				}
			}
			return false
		})
	default:
		return false
	}
}
func everySemanticValue(values []any, predicate func(any) bool) bool {
	for _, value := range values {
		if !predicate(value) {
			return false
		}
	}
	return true
}
func documentSemanticValuesAtPath(value any, path string) []any {
	tokens := strings.Split(strings.NewReplacer("[", ".", "]", "").Replace(path), ".")
	current := []any{value}
	for _, token := range tokens {
		if token == "" {
			continue
		}
		next := make([]any, 0)
		for _, item := range current {
			if token == "*" {
				if array, ok := item.([]any); ok {
					next = append(next, array...)
				}
				continue
			}
			if object, ok := item.(map[string]any); ok {
				if child, exists := object[token]; exists {
					next = append(next, child)
				}
				continue
			}
			if array, ok := item.([]any); ok {
				index, err := strconv.Atoi(token)
				if err == nil && index >= 0 && index < len(array) {
					next = append(next, array[index])
				}
			}
		}
		current = next
	}
	return current
}
