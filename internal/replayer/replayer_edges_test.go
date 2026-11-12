package replayer_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/ezclassify/sdk-go/internal/replayer"
)

// TestHandleInvalidJSONBody verifies the replayer returns HTTP 400 with an
// error envelope when the request body is not valid JSON.
func TestHandleInvalidJSONBody(t *testing.T) {
	r := replayer.New(nil)
	defer r.Close()

	req, err := http.NewRequest(http.MethodPost, r.URL()+"/some/path", bytes.NewBufferString("{bad json"))
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("status: got %d, want 400\nbody: %s", resp.StatusCode, body)
	}

	var envelope map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		t.Fatalf("parse envelope: %v", err)
	}
	if envelope["error"] != "invalid_json" {
		t.Errorf("error field: got %q, want \"invalid_json\"", envelope["error"])
	}
	if _, ok := envelope["message"]; !ok {
		t.Error("envelope is missing \"message\" field")
	}
}

// TestHandleEmptyBodyIsInvalidJSON verifies that a request with no body
// (EOF on decode) also returns 400, exercising the same invalid-body arm.
func TestHandleEmptyBodyIsInvalidJSON(t *testing.T) {
	r := replayer.New(nil)
	defer r.Close()

	req, err := http.NewRequest(http.MethodPost, r.URL()+"/v1/ctx/c/resolve", bytes.NewReader(nil))
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusBadRequest {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("status: got %d, want 400\nbody: %s", resp.StatusCode, body)
	}
}

// TestStableJSONArray verifies that StableJSON correctly serializes a
// []interface{} value as a JSON array with elements in original order.
func TestStableJSONArray(t *testing.T) {
	cases := []struct {
		input interface{}
		want  string
	}{
		{[]interface{}{"a", "b", "c"}, `["a","b","c"]`},
		{[]interface{}{1.0, 2.0, 3.0}, `[1,2,3]`},
		{[]interface{}{}, `[]`},
		// nested: array inside an object
		{map[string]interface{}{
			"tags": []interface{}{"x", "y"},
			"n":    1.0,
		}, `{"n":1,"tags":["x","y"]}`},
		// array of objects: each sub-object gets sorted keys
		{[]interface{}{
			map[string]interface{}{"z": 1.0, "a": 2.0},
			map[string]interface{}{"m": "v"},
		}, `[{"a":2,"z":1},{"m":"v"}]`},
	}

	for _, tc := range cases {
		got := replayer.StableJSON(tc.input)
		if got != tc.want {
			t.Errorf("StableJSON(%v) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// TestRecordingKeyMatchesReplayerLookup verifies that RecordingKey produces
// the same hash the replayer uses internally, so a recording pre-populated
// with RecordingKey can be hit by the matching request.
func TestRecordingKeyMatchesReplayerLookup(t *testing.T) {
	path := "/v1/ctx/acme/resolve"
	body := map[string]interface{}{"version": "1.0"}

	key := replayer.RecordingKey(path, body)
	if len(key) != 64 {
		t.Fatalf("RecordingKey returned %q (len %d), want 64-char hex sha256", key, len(key))
	}

	recordings := map[string]replayer.Recording{
		key: {StatusCode: 200, Body: json.RawMessage(`{"ok":true}`)},
	}
	r := replayer.New(recordings)
	defer r.Close()

	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, r.URL()+path, bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatalf("build request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		t.Fatalf("status: got %d, want 200\nbody: %s", resp.StatusCode, body)
	}
}

// TestRecordingKeyNormalizesTrailingSlash verifies that RecordingKey strips
// trailing slashes from the path, consistent with the replayer's handle function.
func TestRecordingKeyNormalizesTrailingSlash(t *testing.T) {
	body := map[string]interface{}{"x": "y"}
	keyWithSlash := replayer.RecordingKey("/ctx/c/resolve/", body)
	keyWithout := replayer.RecordingKey("/ctx/c/resolve", body)
	if keyWithSlash != keyWithout {
		t.Errorf("RecordingKey trailing-slash normalization failed: %q != %q", keyWithSlash, keyWithout)
	}
}

// TestMissEnvelopeIsValidJSONForQuotedPath verifies the 404 miss envelope
// stays valid JSON when the request path contains characters that need JSON
// escaping. req.URL.Path is percent-decoded, so a quote or backslash in the
// URL reaches the message string.
func TestMissEnvelopeIsValidJSONForQuotedPath(t *testing.T) {
	r := replayer.New(map[string]replayer.Recording{})
	defer r.Close()

	resp, err := http.Post(r.URL()+`/v1/%22quote%5C`, "application/json", bytes.NewReader([]byte(`{}`)))
	if err != nil {
		t.Fatalf("POST: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want 404", resp.StatusCode)
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var envelope struct {
		Error   string `json:"error"`
		Message string `json:"message"`
		Path    string `json:"path"`
	}
	if err := json.Unmarshal(raw, &envelope); err != nil {
		t.Fatalf("miss envelope is not valid JSON: %v\nbody: %s", err, raw)
	}
	if envelope.Error != "no_recording" {
		t.Errorf("error = %q, want no_recording", envelope.Error)
	}
	if envelope.Path != `/v1/"quote\` {
		t.Errorf("path = %q, want the decoded request path", envelope.Path)
	}
	if envelope.Message != `no recording for path=/v1/"quote\` {
		t.Errorf("message = %q, want it to carry the decoded path", envelope.Message)
	}
}
