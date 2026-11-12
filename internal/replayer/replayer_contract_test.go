package replayer_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/ezclassify/sdk-go/internal/replayer"
	"github.com/ezclassify/sdk-go/internal/subsetmatch"
)

// corpusPath returns the absolute path to the replayer-contract corpus file
// by resolving relative to this test source file.
func corpusPath(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	// thisFile = sdk-go/internal/replayer/replayer_contract_test.go
	// corpus is at sdk-go/spec/replayer-contract/v1/corpus.json
	p := filepath.Join(filepath.Dir(thisFile), "..", "..", "spec", "replayer-contract", "v1", "corpus.json")
	if _, err := os.Stat(p); err != nil {
		t.Fatalf("corpus not found at %s: %v", p, err)
	}
	return p
}

func TestCorpusRawBodySelection(t *testing.T) {
	corpusBytes, err := os.ReadFile(corpusPath(t))
	if err != nil {
		t.Fatalf("read corpus: %v", err)
	}
	var corpus struct {
		Cases []struct {
			ID      string `json:"id"`
			Request struct {
				Body    json.RawMessage `json:"body"`
				RawBody *string         `json:"rawBody"`
			} `json:"request"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(corpusBytes, &corpus); err != nil {
		t.Fatalf("parse corpus: %v", err)
	}
	for _, tc := range corpus.Cases {
		if tc.ID != "arg-order" {
			continue
		}
		if tc.Request.RawBody == nil {
			t.Fatal("arg-order requires rawBody")
		}
		var semantic, rawSemantic any
		if err := json.Unmarshal(tc.Request.Body, &semantic); err != nil {
			t.Fatalf("parse body: %v", err)
		}
		if err := json.Unmarshal([]byte(*tc.Request.RawBody), &rawSemantic); err != nil {
			t.Fatalf("parse rawBody: %v", err)
		}
		if !reflect.DeepEqual(semantic, rawSemantic) {
			t.Fatal("body and rawBody must have the same JSON meaning")
		}
		if bytes.Equal(tc.Request.Body, []byte(*tc.Request.RawBody)) {
			t.Fatal("arg-order must use a distinct wire key order")
		}
		if got := wireBody(tc.Request.Body, tc.Request.RawBody); !bytes.Equal(got, []byte(*tc.Request.RawBody)) {
			t.Fatalf("wire body changed: %s", got)
		}
		if got := wireBody(tc.Request.Body, nil); !bytes.Equal(got, tc.Request.Body) {
			t.Fatalf("default body changed: %s", got)
		}
		return
	}
	t.Fatal("arg-order case missing")
}

// wireBody preserves the corpus's deliberate HTTP wire order when rawBody exists.
func wireBody(body json.RawMessage, rawBody *string) []byte {
	if rawBody != nil {
		return []byte(*rawBody)
	}
	return body
}

// TestReplayerContract loads the replayer-contract corpus and runs every case
// against a fresh Replayer instance.  All cases must pass before the replayer
// is used for engine-slice conformance.
func TestReplayerContract(t *testing.T) {
	corpusBytes, err := os.ReadFile(corpusPath(t))
	if err != nil {
		t.Fatalf("read corpus: %v", err)
	}

	var corpus struct {
		Cases []struct {
			ID          string `json:"id"`
			Description string `json:"description"`

			ReplayerRecordings map[string]struct {
				StatusCode int             `json:"statusCode"`
				Body       json.RawMessage `json:"body"`
			} `json:"replayerRecordings"`

			Request struct {
				Method  string            `json:"method"`
				Path    string            `json:"path"`
				Headers map[string]string `json:"headers"`
				Body    json.RawMessage   `json:"body"`
				RawBody *string           `json:"rawBody"`
			} `json:"request"`

			ExpectedStatus       int             `json:"expectedStatus"`
			ExpectedBody         json.RawMessage `json:"expectedBody"`
			ExpectedMissEnvelope json.RawMessage `json:"expectedMissEnvelope"`
		} `json:"cases"`
	}
	if err := json.Unmarshal(corpusBytes, &corpus); err != nil {
		t.Fatalf("parse corpus: %v", err)
	}

	for _, tc := range corpus.Cases {
		tc := tc
		t.Run(tc.ID, func(t *testing.T) {
			// Build the recordings map for this case.
			recordings := make(map[string]replayer.Recording, len(tc.ReplayerRecordings))
			for hash, rec := range tc.ReplayerRecordings {
				recordings[hash] = replayer.Recording{
					StatusCode: rec.StatusCode,
					Body:       rec.Body,
				}
			}

			r := replayer.New(recordings)
			defer r.Close()

			// Build and send the HTTP request.
			var bodyReader io.Reader
			if body := wireBody(tc.Request.Body, tc.Request.RawBody); len(body) > 0 {
				bodyReader = bytes.NewReader(body)
			}
			reqURL := r.URL() + tc.Request.Path
			req, err := http.NewRequest(tc.Request.Method, reqURL, bodyReader)
			if err != nil {
				t.Fatalf("build request: %v", err)
			}
			for k, v := range tc.Request.Headers {
				req.Header.Set(k, v)
			}
			if req.Header.Get("Content-Type") == "" {
				req.Header.Set("Content-Type", "application/json")
			}
			if tc.Request.RawBody != nil {
				if req.GetBody == nil {
					t.Fatal("rawBody request has no inspectable HTTP body")
				}
				wire, err := req.GetBody()
				if err != nil {
					t.Fatalf("read request wire body: %v", err)
				}
				sentBytes, err := io.ReadAll(wire)
				wire.Close()
				if err != nil {
					t.Fatalf("read request wire body: %v", err)
				}
				if !bytes.Equal(sentBytes, []byte(*tc.Request.RawBody)) {
					t.Fatalf("request wire body changed: %s", sentBytes)
				}
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("send request: %v", err)
			}
			defer resp.Body.Close()
			respBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("read response: %v", err)
			}

			// Assert status code.
			if resp.StatusCode != tc.ExpectedStatus {
				t.Errorf("status: got %d, want %d\nbody: %s", resp.StatusCode, tc.ExpectedStatus, respBytes)
			}

			// Assert response body.
			var actual interface{}
			if err := json.Unmarshal(respBytes, &actual); err != nil {
				t.Fatalf("parse response body: %v", err)
			}

			if tc.ExpectedStatus == http.StatusOK && tc.ExpectedBody != nil {
				var expected interface{}
				if err := json.Unmarshal(tc.ExpectedBody, &expected); err != nil {
					t.Fatalf("parse expectedBody: %v", err)
				}
				if !subsetmatch.Match(actual, expected) {
					t.Errorf("body subset mismatch\n  want: %s\n  got:  %s", tc.ExpectedBody, respBytes)
				}
			}

			if tc.ExpectedStatus == http.StatusNotFound && tc.ExpectedMissEnvelope != nil {
				var expected interface{}
				if err := json.Unmarshal(tc.ExpectedMissEnvelope, &expected); err != nil {
					t.Fatalf("parse expectedMissEnvelope: %v", err)
				}
				if !subsetmatch.Match(actual, expected) {
					t.Errorf("miss envelope mismatch\n  want: %s\n  got:  %s", tc.ExpectedMissEnvelope, respBytes)
				}
			}
		})
	}
}
