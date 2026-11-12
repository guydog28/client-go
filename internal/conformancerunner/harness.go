package conformancerunner

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go/internal/replayer"
)

// computeRecordingKey reproduces the canonical recording hash algorithm.
// Hash input: method | stableJSON(args) | stableJSON(fixture) | mode | ctxID | ctxVersion
func computeRecordingKey(method string, args []interface{}, fixture interface{}, mode, ctxID, ctxVersion, guideFixture, accessFixture string) string {
	raw := strings.Join([]string{
		method,
		replayer.StableJSON(args),
		replayer.StableJSON(fixture),
		mode,
		ctxID,
		ctxVersion,
		guideFixture,
		accessFixture,
	}, "|")
	return fmt.Sprintf("%x", sha256.Sum256([]byte(raw)))
}

// loadAllRecordings reads every *.recordings.json in dir and returns a flat map
// of recording-key → raw JSON response body.
func loadAllRecordings(t *testing.T, dir string) map[string]json.RawMessage {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("read recordings dir %s: %v", dir, err)
	}
	result := make(map[string]json.RawMessage)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".recordings.json") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			t.Fatalf("read %s: %v", entry.Name(), err)
		}
		var rec struct {
			Responses map[string]struct {
				Response json.RawMessage `json:"response"`
			} `json:"responses"`
		}
		if err := json.Unmarshal(b, &rec); err != nil {
			t.Fatalf("parse %s: %v", entry.Name(), err)
		}
		for hash, resp := range rec.Responses {
			result[hash] = resp.Response
		}
	}
	return result
}

// Guide and access operations are reachable only through a bound Context: a
// guide or access policy is meaningless without a scheme to bind it to, so
// the REST API serves no standalone guide/access routes. These cases
// route through /v1/ctx/{ctxId}/{version}/... with payload-only bodies (the
// scheme/guide/access identity comes from the named context).
//
// The replayer is keyed by (path, body) and never resolves a live context.
// Engine plans still identify their immutable composition by fixture name(s),
// so derive a deterministic Context id from that composition. This preserves
// the Context-only wire shape while preventing two distinct compositions with
// the same operation body from colliding in the replayer map.
const conformanceCtxVersion = "1.0.0"

func contextRouteBase(fixture interface{}, guideFixture, accessFixture string) string {
	digest := sha256.Sum256([]byte(replayer.StableJSON(map[string]interface{}{
		"fixture":       fixture,
		"guideFixture":  guideFixture,
		"accessFixture": accessFixture,
	})))
	return fmt.Sprintf("/v1/ctx/conformance-%x/%s", digest[:8], conformanceCtxVersion)
}

// buildPlanBody constructs the HTTP request body map and API path for a plan
// case. It returns a non-nil error for a method the Go runner does not
// implement and when a fixture the case depends on cannot be loaded: a missing
// or corrupt fixture must fail the suite, never skip it.
func buildPlanBody(method string, args []interface{}, fixture interface{}, guideFixture, accessFixture string, fixtureDir string) (path string, body map[string]interface{}, err error) {
	ctxBase := contextRouteBase(fixture, guideFixture, accessFixture)
	strArg := func(i int) string {
		if i < len(args) {
			if s, ok := args[i].(string); ok {
				return s
			}
		}
		return ""
	}
	strSliceArg := func(i int) []string {
		if i < len(args) {
			if arr, ok := args[i].([]interface{}); ok {
				out := make([]string, 0, len(arr))
				for _, v := range arr {
					if s, ok := v.(string); ok {
						out = append(out, s)
					}
				}
				return out
			}
		}
		return nil
	}
	// rawArg returns arg i verbatim (nil if absent), preserving the JSON shape
	// so assertion-array args travel to the request body untouched.
	rawArg := func(i int) interface{} {
		if i < len(args) {
			return args[i]
		}
		return nil
	}
	// Preserve canonical options without overwriting an operand.
	withOptions := func(m map[string]interface{}, i int) (map[string]interface{}, error) {
		if i < len(args) {
			if options, ok := args[i].(map[string]interface{}); ok {
				for key, value := range options {
					if prior, exists := m[key]; exists && replayer.StableJSON(prior) != replayer.StableJSON(value) {
						return nil, fmt.Errorf("conflicting plan option: %s", key)
					}
					m[key] = value
				}
			}
		}
		return m, nil
	}

	if method == "registryProcessAll" {
		return "", nil, fmt.Errorf("unsupported plan method %q", method)
	}
	if method == "registrySuggestAcross" {
		return "", nil, fmt.Errorf("unsupported plan method %q", method)
	}

	switch method {
	case "process":
		// Facts alone still supply the required empty typed-text input.
		body = map[string]interface{}{"input": strArg(0)}
		if _, isArray := rawArg(0).([]interface{}); isArray {
			body["assertions"] = rawArg(0)
		}
		body, err = withOptions(body, 1)
		return ctxBase + "/scheme/process", body, err
	case "validate", "suggest":
		body = map[string]interface{}{"input": strArg(0)}
		if _, isArray := rawArg(0).([]interface{}); isArray {
			body = map[string]interface{}{"assertions": rawArg(0)}
		}
		body, err = withOptions(body, 1)
		return ctxBase + "/scheme/" + method, body, err
	case "normalize":
		body, err = withOptions(map[string]interface{}{"input": strArg(0)}, 1)
		return ctxBase + "/scheme/normalize", body, err
	case "scan":
		body, err = withOptions(map[string]interface{}{"text": strArg(0)}, 1)
		return ctxBase + "/scheme/scan", body, err
	case "aggregate":
		body, err = withOptions(map[string]interface{}{"inputs": strSliceArg(0)}, 1)
		return ctxBase + "/scheme/aggregate", body, err
	case "aggregateAssertions":
		// Assertion sets preserve facts no rendered input can carry.
		body, err = withOptions(map[string]interface{}{"assertionSets": rawArg(0)}, 1)
		return ctxBase + "/scheme/aggregate", body, err
	case "diff":
		body, err = withOptions(map[string]interface{}{"a": strArg(0), "b": strArg(1)}, 2)
		return ctxBase + "/scheme/diff", body, err
	case "checkAccess":
		// The access policy is bound to the context; the body carries only the
		// data assertions and the subject's user attributes.
		return ctxBase + "/access/check", map[string]interface{}{
			"assertions":     rawArg(0),
			"userAttributes": rawArg(1),
		}, nil
	case "evaluateMany":
		// The plan's selections already name each rule's guide, so they go out
		// as they are.
		return ctxBase + "/guide/evaluate", map[string]interface{}{
			"selections": rawArg(0),
		}, nil
	case "evaluateGuide":
		// The guide is bound to the context; the body carries only the guide
		// id and the selected rules.
		gf := guideFixture
		if gf == "" {
			if s, ok := fixture.(string); ok {
				gf = s
			}
		}
		parsed, err := loadGuideFixture(fixtureDir, gf)
		if err != nil {
			return "", nil, fmt.Errorf("guide fixture for method %q: %w", method, err)
		}
		return ctxBase + "/guide/evaluate", map[string]interface{}{
			"guideId":         guideIDFromParsed(parsed, gf),
			"selectedRuleIds": strSliceArg(0),
		}, nil
	case "guideScan":
		gf := guideFixture
		if gf == "" {
			if s, ok := fixture.(string); ok {
				gf = s
			}
		}
		parsed, err := loadGuideFixture(fixtureDir, gf)
		if err != nil {
			return "", nil, fmt.Errorf("guide fixture for method %q: %w", method, err)
		}
		return ctxBase + "/guide/scan", map[string]interface{}{
			"text":    strArg(0),
			"guideId": guideIDFromParsed(parsed, gf),
		}, nil
	case "render":
		body = map[string]interface{}{"assertions": rawArg(0)}
		if format, ok := rawArg(1).(string); ok && format != "" {
			body["format"] = format
		}
		return ctxBase + "/scheme/render", body, nil
	default:
		return "", nil, fmt.Errorf("unsupported plan method %q", method)
	}
}

// runCase sends one plan case to the replayer and returns the parsed response.
func runCase(
	t *testing.T,
	replayerURL string,
	fixtureDir string,
	method string,
	fixtureRaw interface{},
	guideFixture string,
	accessFixture string,
	args []interface{},
) (interface{}, error) {
	t.Helper()

	apiPath, bodyMap, err := buildPlanBody(method, args, fixtureRaw, guideFixture, accessFixture, fixtureDir)
	if err != nil {
		return nil, fmt.Errorf("build plan body: %w", err)
	}

	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return nil, fmt.Errorf("marshal body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, replayerURL+apiPath, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		var miss map[string]interface{}
		if json.Unmarshal(respBytes, &miss) == nil && miss["error"] == "no_recording" {
			return nil, fmt.Errorf("no recording for path=%s body=%s", apiPath, bodyBytes)
		}
		return nil, fmt.Errorf("replayer 404: %s", respBytes)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected 200, got %d: %s", resp.StatusCode, respBytes)
	}

	var actual interface{}
	if err := json.Unmarshal(respBytes, &actual); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
	return actual, nil
}

// loadGuideFixture reads a guide fixture file and returns the parsed JSON.
func loadGuideFixture(fixtureDir, fixtureName string) (interface{}, error) {
	fp := filepath.Join(fixtureDir, fixtureName+".json")
	b, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read guide fixture %s: %w", fixtureName, err)
	}
	var m interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("parse guide fixture %s: %w", fixtureName, err)
	}
	return m, nil
}

// guideIDFromParsed extracts the guide ID from an already-parsed guide
// fixture, falling back to the fixture name when the document names none.
func guideIDFromParsed(fixture interface{}, fixtureName string) string {
	m, ok := fixture.(map[string]interface{})
	if !ok {
		return fixtureName
	}
	if g, ok := m["guide"].(map[string]interface{}); ok {
		if id, ok := g["id"].(string); ok {
			return id
		}
	}
	if id, ok := m["id"].(string); ok {
		return id
	}
	return fixtureName
}

// toIfaceSlice coerces v to []interface{}; returns nil if v is not a JSON array.
func toIfaceSlice(v interface{}) []interface{} {
	if arr, ok := v.([]interface{}); ok {
		return arr
	}
	return nil
}
