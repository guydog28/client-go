package conformancerunner

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// writeFixture writes name.json with the given content into dir.
func writeFixture(t *testing.T, dir, name, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dir, name+".json"), []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture %s: %v", name, err)
	}
}

// TestBuildPlanBodySchemeOperationUsesContext pins the Context-only contract:
// scheme operations take their scheme identity from the route's Context and
// therefore neither load nor depend on a scheme fixture.
func TestBuildPlanBodySchemeOperationUsesContext(t *testing.T) {
	dir := t.TempDir()
	path, body, err := buildPlanBody("process", []interface{}{"HELLO"}, "no-such-fixture", "", "", dir)
	if err != nil {
		t.Fatalf("process: err=%v", err)
	}
	if path != contextRouteBase("no-such-fixture", "", "")+"/scheme/process" {
		t.Errorf("path = %q, want Context scheme route", path)
	}
	if body["input"] != "HELLO" {
		t.Errorf("input = %v, want HELLO", body["input"])
	}
}

// TestBuildPlanBodyWireNamesPerRoute pins the per-route wire names for the
// plans' canonical output and surface options. Process and aggregate carry them as
// outputFormat and surface; render keeps format; text read by process or
// validate carries the plan's inputFormat. The recordings are keyed by
// plan args rather than by body, so a stale field name would otherwise pass
// replay while the real server rejects the request.
func TestBuildPlanBodyWireNamesPerRoute(t *testing.T) {
	dir := t.TempDir()
	assertions := []interface{}{map[string]interface{}{"conceptId": "c", "categoryId": "k"}}
	opts := map[string]interface{}{"outputFormat": "full", "surface": "part"}
	cases := []struct {
		method string
		args   []interface{}
		want   map[string]string
	}{
		{"process", []interface{}{assertions, opts}, map[string]string{"outputFormat": "full", "surface": "part"}},
		{"aggregateAssertions", []interface{}{[]interface{}{assertions}, opts}, map[string]string{"outputFormat": "full", "surface": "part"}},
		{"render", []interface{}{assertions, "full"}, map[string]string{"format": "full"}},
		{"process", []interface{}{"CONF/PARTNERS", map[string]interface{}{"inputFormat": "full_line_2019"}}, map[string]string{"input": "CONF/PARTNERS", "inputFormat": "full_line_2019"}},
		{"validate", []interface{}{"CONF", map[string]interface{}{"inputFormat": "sensitivity/short"}}, map[string]string{"input": "CONF", "inputFormat": "sensitivity/short"}},
	}
	for _, c := range cases {
		_, body, err := buildPlanBody(c.method, c.args, "no-such-fixture", "", "", dir)
		if err != nil {
			t.Fatalf("%s: err=%v", c.method, err)
		}
		for k, v := range c.want {
			if body[k] != v {
				t.Errorf("%s: body[%q] = %v, want %q", c.method, k, body[k], v)
			}
		}
		want := map[string]interface{}{}
		for key, value := range c.want {
			want[key] = value
		}
		switch c.method {
		case "aggregateAssertions":
			want["assertionSets"] = c.args[0]
		case "render":
			want["assertions"] = c.args[0]
		case "process":
			if facts, ok := c.args[0].([]interface{}); ok {
				want["assertions"] = facts
				want["input"] = ""
			}
		}
		if !reflect.DeepEqual(body, want) {
			t.Errorf("%s: body=%#v, want %#v", c.method, body, want)
		}
	}
}

// TestBuildPlanBodyCorruptGuideFixtureFails pins that an unparseable guide
// fixture fails the case with the real JSON error instead of skipping. The
// scheme fixture is valid so the failure is isolated to the guide load.
func TestBuildPlanBodyCorruptGuideFixtureFails(t *testing.T) {
	dir := t.TempDir()
	writeFixture(t, dir, "ref", `{"scheme":{"id":"ref-scheme"}}`)
	writeFixture(t, dir, "bad-guide", "{not json")
	_, _, err := buildPlanBody("evaluateGuide", nil, "ref", "bad-guide", "", dir)
	if err == nil {
		t.Fatal("expected an error for a corrupt guide fixture, got nil")
	}
	if !strings.Contains(err.Error(), "bad-guide") {
		t.Errorf("error %q does not name the corrupt fixture", err)
	}
}

// TestBuildPlanBodyUnsupportedMethodFails prevents future contract methods
// from silently disappearing from conformance execution.
func TestBuildPlanBodyUnsupportedMethodFails(t *testing.T) {
	_, _, err := buildPlanBody("someFutureMethod", nil, "ref", "", "", t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "someFutureMethod") {
		t.Fatalf("unsupported method: err=%v", err)
	}
}

// TestBuildPlanBodyCheckAccessShortArgs pins that a checkAccess case with
// fewer than two args builds a body with nil placeholders instead of
// panicking the test binary on an out-of-range index.
func TestBuildPlanBodyCheckAccessShortArgs(t *testing.T) {
	dir := t.TempDir()
	writeFixture(t, dir, "ref", `{"scheme":{"id":"ref-scheme"}}`)
	path, body, err := buildPlanBody("checkAccess", nil, "ref", "", "", dir)
	if err != nil {
		t.Fatalf("checkAccess: err=%v", err)
	}
	if !strings.HasSuffix(path, "/access/check") {
		t.Errorf("path = %q, want an /access/check route", path)
	}
	for _, key := range []string{"assertions", "userAttributes"} {
		if v, ok := body[key]; !ok || v != nil {
			t.Errorf("body[%q] = %v (present=%v), want a nil placeholder", key, v, ok)
		}
	}
}

// TestBuildPlanBodyGuideEvaluate pins the happy path: the guide id in the
// request body comes from the fixture's guide document.
func TestBuildPlanBodyGuideEvaluate(t *testing.T) {
	dir := t.TempDir()
	writeFixture(t, dir, "ref-guide", `{"scheme":{"id":"s-1"},"guide":{"id":"g-1"}}`)
	path, body, err := buildPlanBody("evaluateGuide", nil, "ref-guide", "", "", dir)
	if err != nil {
		t.Fatalf("evaluateGuide: err=%v", err)
	}
	if !strings.HasSuffix(path, "/guide/evaluate") {
		t.Errorf("path = %q, want a /guide/evaluate route", path)
	}
	if body["guideId"] != "g-1" {
		t.Errorf("guideId = %v, want g-1", body["guideId"])
	}
}

// Facts retain their operand without becoming text. Process also requires the
// empty typed-text input; validate and suggest carry exactly one input kind.
func TestBuildPlanBodyFactsKeepTheirWireInput(t *testing.T) {
	for _, method := range []string{"process", "validate", "suggest"} {
		for _, facts := range [][]interface{}{{}, {map[string]interface{}{"conceptId": "c", "categoryId": "k"}}} {
			t.Run(fmt.Sprintf("%s/%d-facts", method, len(facts)), func(t *testing.T) {
				_, body, err := buildPlanBody(method, []interface{}{facts}, "ref", "", "", t.TempDir())
				want := map[string]interface{}{"assertions": facts}
				if method == "process" {
					want["input"] = ""
				}
				if err != nil || !reflect.DeepEqual(body, want) {
					t.Fatalf("body=%#v err=%v, want %#v", body, err, want)
				}
			})
		}
	}
}

func TestBuildPlanBodyRetainsContextAndReadOutputOptions(t *testing.T) {
	for _, method := range []string{"process", "validate", "aggregate"} {
		t.Run(method, func(t *testing.T) {
			context := map[string]interface{}{"role": "part", "partOf": map[string]interface{}{"rendered": "LOW", "assertions": []interface{}{}}}
			options := map[string]interface{}{"context": context, "inputFormat": "short", "outputFormat": "full", "surface": "part"}
			var value interface{} = "LOW"
			want := map[string]interface{}{"input": value, "context": context, "inputFormat": "short", "outputFormat": "full", "surface": "part"}
			if method == "aggregate" {
				value = []interface{}{"LOW"}
				delete(want, "input")
				want["inputs"] = []string{"LOW"}
			}
			_, body, err := buildPlanBody(method, []interface{}{value, options}, "ref", "", "", t.TempDir())
			if err != nil || !reflect.DeepEqual(body, want) {
				t.Fatalf("body=%#v err=%v, want %#v", body, err, want)
			}
		})
	}
}

func TestBuildPlanBodyProcessKeepsPartsChosenFactsAndEmptyOutputs(t *testing.T) {
	facts := []interface{}{map[string]interface{}{"conceptId": "c", "categoryId": "k"}}
	options := map[string]interface{}{"parts": []interface{}{"LOW"}, "assertions": facts, "outputs": []interface{}{}, "suggestFor": []interface{}{}}
	_, body, err := buildPlanBody("process", []interface{}{"HIGH", options}, "ref", "", "", t.TempDir())
	want := map[string]interface{}{"input": "HIGH", "parts": options["parts"], "assertions": facts, "outputs": options["outputs"], "suggestFor": options["suggestFor"]}
	if err != nil || !reflect.DeepEqual(body, want) {
		t.Fatalf("body=%#v err=%v, want %#v", body, err, want)
	}
}

func TestBuildPlanBodySuggestPreservesExplicitFalseAndEmptyFilter(t *testing.T) {
	options := map[string]interface{}{"filter": []interface{}{}, "includeGated": false, "surface": "part", "outputFormat": "full"}
	_, body, err := buildPlanBody("suggest", []interface{}{"LOW", options}, "ref", "", "", t.TempDir())
	want := map[string]interface{}{"input": "LOW", "filter": options["filter"], "includeGated": false, "surface": "part", "outputFormat": "full"}
	if err != nil || !reflect.DeepEqual(body, want) {
		t.Fatalf("body=%#v err=%v, want %#v", body, err, want)
	}
}

func TestBuildPlanBodyConflictingOperandRefuses(t *testing.T) {
	for name, options := range map[string]map[string]interface{}{
		"input": {"input": "another input"},
	} {
		t.Run(name, func(t *testing.T) {
			_, _, err := buildPlanBody("process", []interface{}{"LOW", options}, "ref", "", "", t.TempDir())
			if err == nil || !strings.Contains(err.Error(), "conflicting plan option: "+name) {
				t.Fatalf("conflicting options %#v: err=%v", options, err)
			}
		})
	}
}

func requestAliasSpec(t *testing.T, secondResponse json.RawMessage) string {
	t.Helper()
	dir := t.TempDir()
	for _, name := range []string{"plans", "recordings", "fixtures"} {
		if err := os.MkdirAll(filepath.Join(dir, name, "v1"), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	cases := []map[string]interface{}{}
	responses := map[string]interface{}{}
	for i, args := range [][]interface{}{{"LOW"}, {"LOW", map[string]interface{}{}}} {
		cases = append(cases, map[string]interface{}{"name": fmt.Sprintf("alias-%d", i), "method": "validate", "args": args, "fixture": "ref", "expect": map[string]interface{}{"valid": true}})
		response := json.RawMessage(`{"valid":true,"rank":1}`)
		if i == 1 {
			response = secondResponse
		}
		responses[computeRecordingKey("validate", args, "ref", "fromJSON", "", "", "", "")] = map[string]interface{}{"response": response}
	}
	for path, value := range map[string]interface{}{
		"plans/v1/aliases.json":                 map[string]interface{}{"slice": "engine", "cases": cases},
		"recordings/v1/aliases.recordings.json": map[string]interface{}{"responses": responses},
	} {
		data, err := json.Marshal(value)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, path), data, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return dir
}

// The whole recorded response must agree, even when a subset expectation
// happens to accept both answers. RunConformance's refusal is fatal, so use
// the native test binary to observe the seed-phase error without hiding it.
func TestConformanceConflictingRequestRefuses(t *testing.T) {
	if os.Getenv("EZ_CONFORMANCE_COLLISION_CHILD") == "1" {
		RunConformance(t, requestAliasSpec(t, json.RawMessage(`{"valid":true,"rank":2}`)))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=^TestConformanceConflictingRequestRefuses$", "-test.v")
	cmd.Env = append(os.Environ(), "EZ_CONFORMANCE_COLLISION_CHILD=1")
	output, err := cmd.CombinedOutput()
	if err == nil || !strings.Contains(string(output), "conflicting recordings for request: plan aliases.json case alias-1") {
		t.Fatalf("conflicting request did not refuse in seed phase: err=%v\n%s", err, output)
	}
}

func TestConformanceEqualRequestAnswersRemainExecutable(t *testing.T) {
	RunConformance(t, requestAliasSpec(t, json.RawMessage(`{"rank":1.0,"valid":true}`)))
}
