package ezclassify_test

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	ez "github.com/ezclassify/sdk-go"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// The corpus is read from the vendored spec/ mirror only.
func TestWireCorpusCoverage(t *testing.T) {
	c, e := loadWireCorpus("spec")
	if e != nil {
		t.Fatal(e)
	}
	if e = checkWireDrivers(c, wireDrivers); e != nil {
		t.Fatal(e)
	}
}
func checkWireDrivers(c wireCorpus, drivers map[string]string) error {
	if len(c.Cases) != len(drivers) {
		return fmt.Errorf("wire driver count %d != %d", len(drivers), len(c.Cases))
	}
	for _, r := range c.Cases {
		mode, ok := drivers[r.ID]
		if !ok || (mode != publicDriver && !rawWireReasons[mode]) {
			return fmt.Errorf("unregistered wire id %q", r.ID)
		}
	}
	return nil
}
func TestWireCorpusReplay(t *testing.T)     { runWireCorpus(t, "spec", false) }
func TestWireCorpusHTTPReplay(t *testing.T) { runWireCorpus(t, "spec", true) }
func runWireCorpus(t *testing.T, specDir string, socket bool) {
	c, e := loadWireCorpus(specDir)
	if e != nil {
		t.Fatal(e)
	}
	if e = checkWireDrivers(c, wireDrivers); e != nil {
		t.Fatal(e)
	}
	for _, r := range c.Cases {
		r := r
		t.Run(r.ID, func(t *testing.T) {
			replay := &wireReplay{Case: r}
			baseURL := "http://wire.invalid"
			transport := wireTransport{replay: replay}
			if socket {
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					resp, e := replay.answer(req)
					if e != nil {
						replay.mu.Lock()
						replay.Failure = e
						replay.mu.Unlock()
						http.Error(w, e.Error(), 500)
						return
					}
					defer resp.Body.Close()
					for k, values := range resp.Header {
						for _, v := range values {
							w.Header().Add(k, v)
						}
					}
					w.WriteHeader(resp.StatusCode)
					io.Copy(w, resp.Body)
				}))
				defer server.Close()
				baseURL = server.URL
				transport.next = http.DefaultTransport
			}
			client := ez.New(baseURL, ez.WithHTTPClient(&http.Client{Transport: transport}), ez.WithCredentials(wireCredentials(r)))
			result, err := runWireDriver(context.Background(), client, r)
			replay.mu.Lock()
			defer replay.mu.Unlock()
			if replay.Failure != nil {
				t.Fatal(replay.Failure)
			}
			if replay.Hits != 1 {
				t.Fatalf("request count %d; error %v", replay.Hits, err)
			}
			if e := checkWireObserved(replay); e != nil {
				t.Fatal(e)
			}

			if wireDrivers[r.ID] != publicDriver {
				if err != nil {
					t.Fatal(err)
				}
				return
			}
			if r.Response.Status < 200 || r.Response.Status >= 300 {
				var apiErr *ez.APIError
				if !errors.As(err, &apiErr) || apiErr.StatusCode != r.Response.Status || !equalWireJSON(apiErr.Body, r.Response.Body) {
					t.Fatalf("missing exact API error: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			// Ping is a documented public projection, not a typed DTO return.
			if r.Operation == "getHealth" {
				v, ok := result.(ez.PingResult)
				if !ok {
					t.Fatalf("unexpected Ping result %T", result)
				}
				var b struct {
					EngineVersion string `json:"engineVersion"`
					EngineSHA     string `json:"engineSha"`
					GitTag        string `json:"gitTag"`
				}
				if e := json.Unmarshal(r.Response.Body, &b); e != nil {
					t.Fatal(e)
				}
				if v.EngineVersion != b.EngineVersion || v.EngineSHA != b.EngineSHA || v.GitTag != b.GitTag {
					t.Fatalf("health projection changed: %#v", v)
				}
				return
			}
			actual, e := json.Marshal(result)
			if e != nil {
				t.Fatal(e)
			}
			if !equalWireJSON(actual, r.Response.Body) {
				t.Fatalf("typed response lost or changed recorded fields:\ngot %s\nwant %s", actual, r.Response.Body)
			}
		})
	}
}
func TestWireMatcherControls(t *testing.T) {
	for _, r := range []struct {
		a, b  string
		equal bool
	}{
		{"/x/", "/x", true}, {"/x//", "/x", false}, {"/x?a=2&a=1", "/x?a=1&a=2", true}, {"/x?key=", "/x?key", true}, {"/x?a-=1&a=2", "/x?a=2&a-=1", true}, {"/a%2Fb", "/a/b", false}, {"/x?x=%2F", "/x?x=/", false},
	} {
		t.Run(r.a, func(t *testing.T) {
			got := canonicalWireTarget(r.a) == canonicalWireTarget(r.b)
			if got != r.equal {
				t.Fatalf("target equivalence %v != %v", got, r.equal)
			}
		})
	}
	for _, r := range []struct {
		a, b  string
		equal bool
	}{{`{"a":1,"b":2}`, `{"b":2,"a":1}`, true}, {`[1,2]`, `[2,1]`, false}, {``, `null`, false}, {`{}`, `null`, false}, {`{"x":true}`, `{"x":false}`, false}, {`{"__proto__":1}`, `{}`, false}} {
		t.Run("body:"+r.a, func(t *testing.T) {
			if equalWireJSON([]byte(r.a), []byte(r.b)) != r.equal {
				t.Fatal("wrong body equivalence")
			}
		})
	}
}
func TestWireReplayRefusesAlteredRequests(t *testing.T) {
	r := wireCase{ID: "synthetic"}
	r.Auth.Credential = "context-token"
	r.Request.Method = "POST"
	r.Request.Path = "/v1/ctx/a/1/scheme/process"
	r.Request.Body = json.RawMessage(`{"input":"A","outputs":["x","y"]}`)
	r.Request.Headers = map[string]string{"accept": "application/json", "content-type": "application/json"}
	r.Response.Status = 200
	r.Response.Body = json.RawMessage(`{"valid":true}`)
	for _, name := range []string{"body", "array", "path", "double-slash", "method", "carrier", "bearer", "content-type", "duplicate"} {
		t.Run(name, func(t *testing.T) {
			replay := &wireReplay{Case: r}
			method, path, body := r.Request.Method, r.Request.Path, r.Request.Body
			switch name {
			case "body":
				body = []byte(`{"input":"B","outputs":["x","y"]}`)
			case "array":
				body = []byte(`{"input":"A","outputs":["y","x"]}`)
			case "path":
				path += "-changed"
			case "double-slash":
				path += "//"
			case "method":
				method = "GET"
			}
			req, e := http.NewRequest(method, "http://wire.invalid"+path, bytes.NewReader(body))
			if e != nil {
				t.Fatal(e)
			}
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("x-ezclassify-context-token", "fixture-context-carrier")
			switch name {
			case "carrier":
				req.Header.Del("x-ezclassify-context-token")
			case "bearer":
				req.Header.Set("Authorization", "Bearer unexpected")
			case "content-type":
				req.Header.Set("Content-Type", "text/plain")
			case "duplicate":
				replay.Hits = 1
			}
			if _, e = replay.answer(req); e == nil {
				t.Fatal("altered request accepted")
			}
		})
	}
}
func TestWireDriverRegistryRefusesUnreviewedEntries(t *testing.T) {
	c := wireCorpus{Cases: []wireCase{{ID: "known"}}}
	if checkWireDrivers(c, map[string]string{"unknown": publicDriver}) == nil {
		t.Fatal("unknown driver accepted")
	}
	if checkWireDrivers(c, map[string]string{"known": "no driver yet"}) == nil {
		t.Fatal("a raw-wire entry without a reviewed reason was accepted")
	}
	if checkWireDrivers(c, map[string]string{"known": undeclaredRoute}) != nil {
		t.Fatal("a reviewed raw-wire reason was refused")
	}
}
func TestWireProvenanceControls(t *testing.T) {
	// Fixture files are disposable inputs to the actual maintained loader.
	yaml := []byte("canonical openapi")
	for _, name := range []string{"digest", "count", "duplicate", "divergence", "empty"} {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			wire := filepath.Join(dir, "wire/v1")
			if e := writeWireControl(wire, dir, yaml, name); e != nil {
				t.Fatal(e)
			}
			if _, e := loadWireCorpus(dir); e == nil {
				t.Fatal("invalid provenance accepted")
			}
		})
	}
}

func writeWireControl(wire, dir string, yaml []byte, mutation string) error {
	if e := os.MkdirAll(wire, 0700); e != nil {
		return e
	}
	if e := os.WriteFile(filepath.Join(dir, "openapi.yaml"), yaml, 0600); e != nil {
		return e
	}
	c := wireCorpus{Version: "v1", Cases: []wireCase{{ID: "one", Operation: "op"}}}
	m := wireManifest{Version: "v1", Divergences: []json.RawMessage{}}
	m.Spec.Digest = fmt.Sprintf("sha256:%x", sha256.Sum256(yaml))
	m.Coverage.Cases = 1
	m.Coverage.Operations = 1
	switch mutation {
	case "digest":
		m.Spec.Digest = "sha256:stale"
	case "count":
		m.Coverage.Cases = 2
	case "duplicate":
		c.Cases = append(c.Cases, c.Cases[0])
		m.Coverage.Cases = 2
	case "divergence":
		m.Divergences = []json.RawMessage{json.RawMessage(`{"unresolved":true}`)}
	case "empty":
		c.Cases = nil
		m.Coverage.Cases = 0
	}
	for name, v := range map[string]interface{}{"corpus.json": c, "manifest.json": m} {
		b, e := json.Marshal(v)
		if e != nil {
			return e
		}
		if e = os.WriteFile(filepath.Join(wire, name), b, 0600); e != nil {
			return e
		}
	}
	return nil
}

func checkWireObserved(replay *wireReplay) error {
	r := replay.Case
	if replay.ObservedStatus != r.Response.Status || !equalWireJSON(replay.ObservedBody, r.Response.Body) {
		return fmt.Errorf("recorded response changed: status %d body %s", replay.ObservedStatus, replay.ObservedBody)
	}
	for k, v := range r.Response.Headers {
		if replay.ObservedHeaders.Get(k) != v {
			return fmt.Errorf("received response header %s: %q != %q", k, replay.ObservedHeaders.Get(k), v)
		}
	}
	return nil
}
func TestWireResponseMutationControls(t *testing.T) {
	for _, name := range []string{"payload", "status", "header"} {
		t.Run(name, func(t *testing.T) {
			r := &wireReplay{}
			r.Case.Response.Status = 400
			r.Case.Response.Body = json.RawMessage(`{"error":"invalid_input","detail":"field"}`)
			r.Case.Response.Headers = map[string]string{"allow": "POST", "content-type": "application/json"}
			r.ObservedStatus = 400
			r.ObservedBody = append([]byte(nil), r.Case.Response.Body...)
			r.ObservedHeaders = http.Header{"Allow": []string{"POST"}, "Content-Type": []string{"application/json"}}
			if e := checkWireObserved(r); e != nil {
				t.Fatal(e)
			}
			switch name {
			case "payload":
				r.ObservedBody = []byte(`{"error":"invalid_input","detail":"other"}`)
			case "status":
				r.ObservedStatus = 200
			case "header":
				r.ObservedHeaders.Del("Allow")
			}
			if e := checkWireObserved(r); e == nil {
				t.Fatal("changed received response accepted")
			}
		})
	}
}

// Raw cases have no public operation to attach route-scoped credentials for
// them. Their auth metadata still describes the exact recorded HTTP exchange.
func TestRawWireRecordedCredentials(t *testing.T) {
	for _, credential := range []string{"none", "context-token", "invalid-bearer"} {
		t.Run(credential, func(t *testing.T) {
			r := wireCase{ID: "route-miss.route-not-found"}
			r.Auth.Credential = credential
			r.Request.Method = http.MethodGet
			r.Request.Path = "/v1/no-such-route"
			if credential != "context-token" {
				r.Request.Headers = map[string]string{"accept": "application/json"}
			}
			r.Response.Status = http.StatusNotFound
			r.Response.Headers = map[string]string{"content-type": "application/json"}
			r.Response.Body = json.RawMessage(`{"code":"route_not_found"}`)
			originalHeaders, err := json.Marshal(r.Request.Headers)
			if err != nil {
				t.Fatal(err)
			}
			replay := &wireReplay{Case: r}
			client := ez.New("http://wire.invalid", ez.WithHTTPClient(&http.Client{Transport: wireTransport{replay: replay}}), ez.WithCredentials(wireCredentials(r)))
			if _, err := runWireDriver(context.Background(), client, r); err != nil {
				t.Fatal(err)
			}
			if replay.Hits != 1 {
				t.Fatalf("request count %d, want one", replay.Hits)
			}
			if err := checkWireObserved(replay); err != nil {
				t.Fatal(err)
			}
			afterHeaders, err := json.Marshal(r.Request.Headers)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(originalHeaders, afterHeaders) {
				t.Fatal("raw driver mutated recorded headers")
			}
		})
	}
}
