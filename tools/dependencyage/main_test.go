package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// registryStandIn plays both the module proxy and the module index. The proxy
// side serves .info answers for the request paths in info, each carrying the
// given commit time. The index side serves firstSeen as index entries ordered
// by timestamp, honoring since and limit, and adds an unrelated entry one hour
// in the future so a scan can see that the index has caught up. Every request
// path is recorded.
type registryStandIn struct {
	mu        sync.Mutex
	info      map[string]string
	firstSeen map[string]time.Time
	// lagging, when positive, is how many index requests are answered as if
	// the index had not yet caught up with the present.
	lagging int
	// busy, when positive, is how many index requests are refused with 429.
	busy     int
	requests []string
}

func (p *registryStandIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.requests = append(p.requests, r.URL.Path)
	if r.URL.Path == "/index" && p.busy > 0 {
		p.busy--
		http.Error(w, "slow down", http.StatusTooManyRequests)
		return
	}
	if r.URL.Path == "/index" {
		p.serveIndex(w, r)
		return
	}
	if strings.Contains(r.URL.Path, "broken") {
		http.Error(w, "boom", http.StatusInternalServerError)
		return
	}
	commit, ok := p.info[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	version := strings.TrimSuffix(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:], ".info")
	_ = json.NewEncoder(w).Encode(map[string]string{"Version": version, "Time": commit})
}

func (p *registryStandIn) serveIndex(w http.ResponseWriter, r *http.Request) {
	since, err := time.Parse(time.RFC3339Nano, r.URL.Query().Get("since"))
	if err != nil {
		http.Error(w, "bad since", http.StatusBadRequest)
		return
	}
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	entries := make([]indexEntry, 0, len(p.firstSeen)+1)
	for key, at := range p.firstSeen {
		path, version, _ := strings.Cut(key, "@")
		entries = append(entries, indexEntry{Path: path, Version: version, Timestamp: at})
	}
	if p.lagging > 0 {
		p.lagging--
	} else {
		entries = append(entries, indexEntry{Path: "example.com/unrelated", Version: "v9.9.9", Timestamp: time.Now().Add(time.Hour)})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Timestamp.Before(entries[j].Timestamp) })
	encoder := json.NewEncoder(w)
	written := 0
	for _, e := range entries {
		if e.Timestamp.Before(since) || written == limit {
			continue
		}
		_ = encoder.Encode(e)
		written++
	}
}

func daysAgo(days float64) time.Time {
	return time.Now().Add(-time.Duration(days * float64(day))).UTC()
}

func ago(days float64) string { return daysAgo(days).Format(time.RFC3339) }

// runGate writes each go.sum body to its own file, points the gate at the
// stand-in for both the proxy and the index, and returns the error plus both
// output streams.
func runGate(t *testing.T, registry *registryStandIn, gosums ...string) (error, string, string) {
	t.Helper()
	server := httptest.NewServer(registry)
	t.Cleanup(server.Close)
	args := []string{"-proxy", server.URL, "-index", server.URL + "/index"}
	for _, body := range gosums {
		path := filepath.Join(t.TempDir(), "go.sum")
		if err := os.WriteFile(path, []byte(body), 0o600); err != nil {
			t.Fatal(err)
		}
		args = append(args, "-gosum", path)
	}
	var out, errOut bytes.Buffer
	err := run(args, &out, &errOut)
	return err, out.String(), errOut.String()
}

func TestGateRefusesModuleFirstCachedInsideWindowAndSaysWhenItClears(t *testing.T) {
	firstSeen := daysAgo(2).Truncate(time.Second)
	registry := &registryStandIn{
		info:      map[string]string{"/example.com/fresh/@v/v1.0.0.info": firstSeen.Format(time.RFC3339)},
		firstSeen: map[string]time.Time{"example.com/fresh@v1.0.0": firstSeen},
	}
	err, _, errOut := runGate(t, registry, "example.com/fresh v1.0.0 h1:aaa=\nexample.com/fresh v1.0.0/go.mod h1:bbb=\n")
	if err == nil {
		t.Fatal("a module first cached two days ago must fail a seven day window")
	}
	want := "example.com/fresh@v1.0.0 first seen " + firstSeen.Format(time.RFC3339) + ", clears " + firstSeen.Add(7*day).Format(time.RFC3339)
	if !strings.Contains(errOut, want) {
		t.Fatalf("missing offender line %q in:\n%s", want, errOut)
	}
}

// A version can carry an old commit time and still be new to the proxy: the
// tagger sets the commit date. Age must come from the index, not .info Time.
func TestGateRefusesBackdatedCommitFirstCachedInsideWindow(t *testing.T) {
	firstSeen := daysAgo(1).Truncate(time.Second)
	registry := &registryStandIn{
		info:      map[string]string{"/example.com/backdated/@v/v1.2.3.info": ago(400)},
		firstSeen: map[string]time.Time{"example.com/backdated@v1.2.3": firstSeen},
	}
	err, _, errOut := runGate(t, registry, "example.com/backdated v1.2.3 h1:aaa=\n")
	if err == nil {
		t.Fatal("a version first cached yesterday must fail even when its commit time is 400 days old")
	}
	if !strings.Contains(errOut, "example.com/backdated@v1.2.3 first seen "+firstSeen.Format(time.RFC3339)) {
		t.Fatalf("offender line must report the first-cache time, got:\n%s", errOut)
	}
}

func TestGatePassesModuleFirstCachedBeforeWindow(t *testing.T) {
	registry := &registryStandIn{
		info:      map[string]string{"/example.com/old/@v/v1.0.0.info": ago(30)},
		firstSeen: map[string]time.Time{"example.com/old@v1.0.0": daysAgo(30)},
	}
	err, out, errOut := runGate(t, registry, "example.com/old v1.0.0 h1:aaa=\n")
	if err != nil {
		t.Fatalf("a module first cached thirty days ago must pass: %v\n%s", err, errOut)
	}
	if !strings.Contains(out, "passed") {
		t.Fatalf("expected a pass line, got %q", out)
	}
}

func TestGateChecksEveryGoSumItIsGiven(t *testing.T) {
	registry := &registryStandIn{
		info: map[string]string{
			"/example.com/fresh/@v/v1.0.0.info": ago(400),
			"/example.com/old/@v/v1.0.0.info":   ago(400),
		},
		firstSeen: map[string]time.Time{"example.com/fresh@v1.0.0": daysAgo(1)},
	}
	err, _, errOut := runGate(t, registry, "example.com/fresh v1.0.0 h1:aaa=\n", "example.com/old v1.0.0 h1:bbb=\n")
	if err == nil || !strings.Contains(errOut, "example.com/fresh@v1.0.0 first seen") {
		t.Fatalf("a fresh module in the first of two go.sum files must fail the gate, got err=%v\n%s", err, errOut)
	}
}

func TestDefaultGoSumSetIsEveryGoSumOutsideVendorAndTestdata(t *testing.T) {
	root := t.TempDir()
	for _, path := range []string{"go.sum", "tools/codegen/go.sum", "tools/sbom/go.sum", "vendor/example.com/x/go.sum", "internal/testdata/go.sum", ".cache/go.sum"} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, nil, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	got, err := findGoSums(root)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{filepath.Join(root, "go.sum"), filepath.Join(root, "tools/codegen/go.sum"), filepath.Join(root, "tools/sbom/go.sum")}
	if strings.Join(got, "\n") != strings.Join(want, "\n") {
		t.Fatalf("go.sum set = %v, want %v", got, want)
	}
	if _, err := findGoSums(t.TempDir()); err == nil {
		t.Fatal("a tree with no go.sum must fail rather than check nothing")
	}
}

func TestGateExemptsFirstPartyModulesWithoutAsking(t *testing.T) {
	registry := &registryStandIn{}
	err, _, errOut := runGate(t, registry, "github.com/ezclassify/sdk-go v1.0.0 h1:aaa=\n")
	if err != nil {
		t.Fatalf("first-party module must be exempt: %v\n%s", err, errOut)
	}
	if len(registry.requests) != 0 {
		t.Fatalf("first-party module must not be looked up, got %v", registry.requests)
	}
}

func TestGateWaitsForALaggingIndexToCatchUp(t *testing.T) {
	defer func(pause time.Duration) { indexRetryPause = pause }(indexRetryPause)
	indexRetryPause = time.Millisecond
	// Four segments each see the lag once before the index catches up.
	registry := &registryStandIn{
		info:      map[string]string{"/example.com/old/@v/v1.0.0.info": ago(400)},
		firstSeen: map[string]time.Time{"example.com/old@v1.0.0": daysAgo(30)},
		lagging:   concurrency,
	}
	err, _, errOut := runGate(t, registry, "example.com/old v1.0.0 h1:aaa=\n")
	if err != nil {
		t.Fatalf("the gate must pass once the index catches up: %v\n%s", err, errOut)
	}
	if registry.lagging != 0 {
		t.Fatalf("the gate passed without waiting for the index (%d lagging answers unused)", registry.lagging)
	}
}

func TestGateRetriesARateLimitedIndex(t *testing.T) {
	defer func(pause time.Duration) { indexRetryPause = pause }(indexRetryPause)
	indexRetryPause = time.Millisecond
	registry := &registryStandIn{
		info:      map[string]string{"/example.com/fresh/@v/v1.0.0.info": ago(400)},
		firstSeen: map[string]time.Time{"example.com/fresh@v1.0.0": daysAgo(1)},
		busy:      3 * concurrency,
	}
	err, _, errOut := runGate(t, registry, "example.com/fresh v1.0.0 h1:aaa=\n")
	if err == nil || !strings.Contains(errOut, "example.com/fresh@v1.0.0 first seen") {
		t.Fatalf("the gate must read through 429 answers and still find the fresh version, got err=%v\n%s", err, errOut)
	}
	if registry.busy != 0 {
		t.Fatalf("%d refusals were never hit", registry.busy)
	}
}

func TestIndexScanFailsClosedWhenTheIndexKeepsRefusing(t *testing.T) {
	defer func(pause time.Duration) { indexRetryPause = pause }(indexRetryPause)
	indexRetryPause = time.Millisecond
	server := httptest.NewServer(&registryStandIn{busy: 1 << 30})
	t.Cleanup(server.Close)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	now := time.Now()
	want := map[module]bool{{Path: "example.com/old", Version: "v1.0.0"}: true}
	if _, err := scanIndex(ctx, server.Client(), server.URL+"/index", now.Add(-7*day), now, want); err == nil || !strings.Contains(err.Error(), "429") {
		t.Fatalf("an index that keeps answering 429 must fail the scan, got %v", err)
	}
}

func TestIndexScanFailsClosedWhenTheIndexNeverCatchesUp(t *testing.T) {
	defer func(pause time.Duration) { indexRetryPause = pause }(indexRetryPause)
	indexRetryPause = time.Millisecond
	server := httptest.NewServer(&registryStandIn{lagging: 1 << 30})
	t.Cleanup(server.Close)
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	now := time.Now()
	want := map[module]bool{{Path: "example.com/old", Version: "v1.0.0"}: true}
	if _, err := scanIndex(ctx, server.Client(), server.URL+"/index", now.Add(-7*day), now, want); err == nil {
		t.Fatal("an index that never reaches the end of the span must fail the scan")
	}
}

func TestGateFailsClosedWhenItCannotConfirmAnAge(t *testing.T) {
	cases := map[string]string{
		"server error":     "example.com/broken v1.0.0 h1:aaa=\n",
		"unknown module":   "example.com/missing v1.0.0 h1:aaa=\n",
		"unreadable index": "example.com/old v1.0.0 h1:aaa=\n",
	}
	for name, gosum := range cases {
		t.Run(name, func(t *testing.T) {
			registry := &registryStandIn{info: map[string]string{"/example.com/old/@v/v1.0.0.info": ago(400)}}
			server := httptest.NewServer(registry)
			t.Cleanup(server.Close)
			index := server.URL + "/index"
			if name == "unreadable index" {
				broken := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					_, _ = w.Write([]byte("not json\n"))
				}))
				t.Cleanup(broken.Close)
				index = broken.URL
			}
			path := filepath.Join(t.TempDir(), "go.sum")
			_ = os.WriteFile(path, []byte(gosum), 0o600)
			var out, errOut bytes.Buffer
			err := run([]string{"-gosum", path, "-proxy", server.URL, "-index", index}, &out, &errOut)
			if err == nil || !strings.Contains(errOut.String(), "fails closed") {
				t.Fatalf("expected a closed failure, got err=%v\n%s", err, errOut.String())
			}
		})
	}
	t.Run("unreachable proxy and index", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "go.sum")
		_ = os.WriteFile(path, []byte("example.com/old v1.0.0 h1:aaa=\n"), 0o600)
		var out, errOut bytes.Buffer
		// Port 1 refuses connections.
		if run([]string{"-gosum", path, "-proxy", "http://127.0.0.1:1", "-index", "http://127.0.0.1:1/index"}, &out, &errOut) == nil {
			t.Fatal("an unreachable proxy must fail the gate")
		}
	})
}

func TestGateFailsOnUnreadableGoSum(t *testing.T) {
	var out, errOut bytes.Buffer
	if run([]string{"-gosum", filepath.Join(t.TempDir(), "absent")}, &out, &errOut) == nil {
		t.Fatal("a missing go.sum must fail the gate")
	}
}

func TestProxyRequestsUseTheModuleProxyCaseEncoding(t *testing.T) {
	registry := &registryStandIn{info: map[string]string{"/github.com/!rave!no!x/go-jsoncommentstrip/@v/v1.0.0.info": ago(400)}}
	err, _, errOut := runGate(t, registry, "github.com/RaveNoX/go-jsoncommentstrip v1.0.0/go.mod h1:aaa=\n")
	if err != nil {
		t.Fatalf("encoded path must resolve: %v\n%s", err, errOut)
	}
}
