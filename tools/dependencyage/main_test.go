package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// proxyStandIn serves module proxy .info answers from an in-memory table keyed
// by the request path, and records every path asked for.
type proxyStandIn struct {
	mu       sync.Mutex
	times    map[string]string
	requests []string
}

func (p *proxyStandIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	p.requests = append(p.requests, r.URL.Path)
	p.mu.Unlock()
	if strings.Contains(r.URL.Path, "broken") {
		http.Error(w, "boom", http.StatusInternalServerError)
		return
	}
	published, ok := p.times[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	body := map[string]string{"Version": "v0"}
	if published != "-" {
		body["Time"] = published
	}
	_ = json.NewEncoder(w).Encode(body)
}

func ago(days float64) string {
	return time.Now().Add(-time.Duration(days * float64(day))).UTC().Format(time.RFC3339)
}

// runGate writes a go.sum, points the gate at the stand-in and returns the
// error plus both output streams.
func runGate(t *testing.T, proxy *proxyStandIn, gosum string) (error, string, string) {
	t.Helper()
	server := httptest.NewServer(proxy)
	t.Cleanup(server.Close)
	path := filepath.Join(t.TempDir(), "go.sum")
	if err := os.WriteFile(path, []byte(gosum), 0o600); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer
	err := run([]string{"-gosum", path, "-proxy", server.URL}, &out, &errOut)
	return err, out.String(), errOut.String()
}

func TestGateRefusesModuleInsideWindowAndSaysWhenItClears(t *testing.T) {
	published := ago(2)
	proxy := &proxyStandIn{times: map[string]string{"/example.com/fresh/@v/v1.0.0.info": published}}
	err, _, errOut := runGate(t, proxy, "example.com/fresh v1.0.0 h1:aaa=\nexample.com/fresh v1.0.0/go.mod h1:bbb=\n")
	if err == nil {
		t.Fatal("a module published two days ago must fail a seven day window")
	}
	at, _ := time.Parse(time.RFC3339, published)
	want := "example.com/fresh@v1.0.0 published " + published + ", clears " + at.Add(7*day).Format(time.RFC3339)
	if !strings.Contains(errOut, want) {
		t.Fatalf("missing offender line %q in:\n%s", want, errOut)
	}
}

func TestGatePassesModuleOutsideWindow(t *testing.T) {
	proxy := &proxyStandIn{times: map[string]string{"/example.com/old/@v/v1.0.0.info": ago(30)}}
	err, out, errOut := runGate(t, proxy, "example.com/old v1.0.0 h1:aaa=\n")
	if err != nil {
		t.Fatalf("a module published thirty days ago must pass: %v\n%s", err, errOut)
	}
	if !strings.Contains(out, "passed") {
		t.Fatalf("expected a pass line, got %q", out)
	}
}

func TestGateExemptsFirstPartyModulesWithoutAsking(t *testing.T) {
	proxy := &proxyStandIn{times: map[string]string{}}
	err, _, errOut := runGate(t, proxy, "github.com/ezclassify/sdk-go v1.0.0 h1:aaa=\n")
	if err != nil {
		t.Fatalf("first-party module must be exempt: %v\n%s", err, errOut)
	}
	if len(proxy.requests) != 0 {
		t.Fatalf("first-party module must not be looked up, got %v", proxy.requests)
	}
}

func TestGateFailsClosedWhenTheProxyCannotConfirmAPublishTime(t *testing.T) {
	cases := map[string]string{
		"server error":      "example.com/broken v1.0.0 h1:aaa=\n",
		"unknown module":    "example.com/missing v1.0.0 h1:aaa=\n",
		"no publish time":   "example.com/untimed v1.0.0 h1:aaa=\n",
		"unreadable time":   "example.com/garbled v1.0.0 h1:aaa=\n",
		"unreachable proxy": "",
	}
	for name, gosum := range cases {
		t.Run(name, func(t *testing.T) {
			proxy := &proxyStandIn{times: map[string]string{
				"/example.com/untimed/@v/v1.0.0.info": "-",
				"/example.com/garbled/@v/v1.0.0.info": "yesterday",
			}}
			if name == "unreachable proxy" {
				path := filepath.Join(t.TempDir(), "go.sum")
				_ = os.WriteFile(path, []byte("example.com/old v1.0.0 h1:aaa=\n"), 0o600)
				var out, errOut bytes.Buffer
				// Port 1 refuses connections.
				if run([]string{"-gosum", path, "-proxy", "http://127.0.0.1:1"}, &out, &errOut) == nil {
					t.Fatal("an unreachable proxy must fail the gate")
				}
				return
			}
			err, _, errOut := runGate(t, proxy, gosum)
			if err == nil || !strings.Contains(errOut, "fails closed") {
				t.Fatalf("expected a closed failure, got err=%v\n%s", err, errOut)
			}
		})
	}
}

func TestGateFailsOnUnreadableGoSum(t *testing.T) {
	var out, errOut bytes.Buffer
	if run([]string{"-gosum", filepath.Join(t.TempDir(), "absent")}, &out, &errOut) == nil {
		t.Fatal("a missing go.sum must fail the gate")
	}
}

func TestProxyRequestsUseTheModuleProxyCaseEncoding(t *testing.T) {
	proxy := &proxyStandIn{times: map[string]string{"/github.com/!rave!no!x/go-jsoncommentstrip/@v/v1.0.0.info": ago(400)}}
	err, _, errOut := runGate(t, proxy, "github.com/RaveNoX/go-jsoncommentstrip v1.0.0/go.mod h1:aaa=\n")
	if err != nil {
		t.Fatalf("encoded path must resolve: %v\n%s", err, errOut)
	}
}
