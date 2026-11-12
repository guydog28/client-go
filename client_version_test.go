package ezclassify_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
)

func TestEngineSpecVersion(t *testing.T) {
	c := ezclassify.New("http://localhost:9999")
	v := c.EngineSpecVersion()
	if v == "" {
		t.Fatal("EngineSpecVersion() returned empty string")
	}
	parts := strings.SplitN(v, ".", 3)
	if len(parts) < 2 {
		t.Fatalf("EngineSpecVersion() %q does not look like semver", v)
	}
}

func newHealthServer(t *testing.T, engineVersion, sha, tag string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/health" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		payload := map[string]string{
			"status":         "ok",
			"engineVersion":  engineVersion,
			"engineSha":      sha,
			"gitTag":         tag,
			"openapiVersion": engineVersion,
			"abiVersion":     "v1",
		}
		_ = json.NewEncoder(w).Encode(payload)
	}))
}

func TestPingReturnsEngineVersion(t *testing.T) {
	srv := newHealthServer(t, "1.0.0", "deadbeef", "v1.0.0")
	defer srv.Close()

	c := ezclassify.New(srv.URL)
	result, err := c.Ping(context.Background())
	if err != nil {
		t.Fatalf("Ping() error: %v", err)
	}
	if result.EngineVersion != "1.0.0" {
		t.Errorf("EngineVersion = %q, want 1.0.0", result.EngineVersion)
	}
	if result.EngineSHA != "deadbeef" {
		t.Errorf("EngineSHA = %q, want deadbeef", result.EngineSHA)
	}
	if result.GitTag != "v1.0.0" {
		t.Errorf("GitTag = %q, want v1.0.0", result.GitTag)
	}
}

// TestPingNoWarnOnMatchingVersions pins that Ping stays quiet when the server
// reports the same major.minor the SDK was built against. The redirected
// logger is process-global, so this test and its mismatch counterpart must not
// run in parallel.
func TestPingNoWarnOnMatchingVersions(t *testing.T) {
	c := ezclassify.New("http://localhost:9999")
	matching := c.EngineSpecVersion()
	srv := newHealthServer(t, matching, "", "")
	defer srv.Close()

	var logged bytes.Buffer
	log.SetOutput(&logged)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })

	c = ezclassify.New(srv.URL)
	if _, err := c.Ping(context.Background()); err != nil {
		t.Fatalf("Ping() error: %v", err)
	}
	if logged.Len() != 0 {
		t.Errorf("Ping logged %q on matching versions, want no warning", logged.String())
	}
}

func TestPingErrorOnUnreachableServer(t *testing.T) {
	c := ezclassify.New("http://127.0.0.1:19999", ezclassify.WithHTTPClient(&http.Client{Timeout: 1}))
	_, err := c.Ping(context.Background())
	if err == nil {
		t.Fatal("expected error pinging unreachable server, got nil")
	}
}
