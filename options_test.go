package ezclassify

import (
	"context"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// headerRecorder is a transport that records the headers of every request it
// carries and answers with a canned health payload, so a test can tell which
// client configuration a request actually went through.
type headerRecorder struct {
	seen []http.Header
}

func (h *headerRecorder) RoundTrip(req *http.Request) (*http.Response, error) {
	h.seen = append(h.seen, req.Header.Clone())
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(`{"status":"ok","engineVersion":"1.0.0","engineSha":"sha","gitTag":"tag","openapiVersion":"1.0.0","abiVersion":"v1"}`)),
		Request:    req,
	}, nil
}

func TestNewBoundsWaitingByDefault(t *testing.T) {
	if got := New("http://localhost").httpClient.Timeout; got != 30*time.Second {
		t.Fatalf("default timeout = %v, want 30s", got)
	}
}

func TestWithTimeoutSetsTheClientTimeout(t *testing.T) {
	if got := New("http://localhost", WithTimeout(5*time.Second)).httpClient.Timeout; got != 5*time.Second {
		t.Fatalf("timeout = %v, want 5s", got)
	}
}

func TestWithTimeoutOverridesASuppliedClientsTimeoutInEitherOrder(t *testing.T) {
	supplied := &http.Client{Timeout: time.Minute}
	for name, opts := range map[string][]Option{
		"timeout first": {WithTimeout(5 * time.Second), WithHTTPClient(supplied)},
		"client first":  {WithHTTPClient(supplied), WithTimeout(5 * time.Second)},
	} {
		if got := New("http://localhost", opts...).httpClient.Timeout; got != 5*time.Second {
			t.Errorf("%s: timeout = %v, want the explicit 5s", name, got)
		}
	}
}

func TestWithTimeoutRejectsANonPositiveDuration(t *testing.T) {
	for _, d := range []time.Duration{0, -time.Second} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("WithTimeout(%v) was accepted, which would allow an unbounded wait", d)
				}
			}()
			_ = WithTimeout(d)
		}()
	}
}

func TestWithHTTPClientRejectsNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("WithHTTPClient(nil) was accepted")
		}
	}()
	_ = WithHTTPClient(nil)
}

func TestWithHTTPClientAppliesATimeoutWhenTheSuppliedClientHasNone(t *testing.T) {
	transport := &headerRecorder{}
	for _, timeout := range []time.Duration{0, -time.Second} {
		c := New("http://localhost", WithHTTPClient(&http.Client{Transport: transport, Timeout: timeout}))
		if got := c.httpClient.Timeout; got != 30*time.Second {
			t.Errorf("supplied Timeout %v: client timeout = %v, want the 30s default", timeout, got)
		}
		if _, err := c.Ping(context.Background()); err != nil {
			t.Fatal(err)
		}
	}
	if len(transport.seen) != 2 {
		t.Fatalf("supplied transport carried %d requests, want 2", len(transport.seen))
	}
}

func TestWithHTTPClientKeepsASuppliedPositiveTimeoutAndPolicies(t *testing.T) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}
	supplied := &http.Client{Timeout: 3 * time.Second, Jar: jar}
	c := New("http://localhost", WithHTTPClient(supplied))
	if c.httpClient.Timeout != 3*time.Second {
		t.Errorf("timeout = %v, want the supplied 3s", c.httpClient.Timeout)
	}
	if c.httpClient.Jar != jar {
		t.Error("the supplied cookie jar was not kept")
	}
}

func TestWithHTTPClientDoesNotMutateTheCallersClient(t *testing.T) {
	transport := &headerRecorder{}
	supplied := &http.Client{Transport: transport}
	c := New("http://localhost", WithHTTPClient(supplied), WithCredentials(RestClientCredentials{AccessToken: "token"}))
	if supplied.Timeout != 0 {
		t.Errorf("caller's Timeout became %v; the default must apply to a copy", supplied.Timeout)
	}
	if supplied.Transport != http.RoundTripper(transport) {
		t.Error("caller's Transport was replaced")
	}
	if c.httpClient == supplied {
		t.Error("client retained the caller's mutable http.Client pointer")
	}
}

func TestWithCredentialsComposesWithASuppliedTransport(t *testing.T) {
	transport := &headerRecorder{}
	c := New("http://localhost",
		WithHTTPClient(&http.Client{Transport: transport}),
		WithCredentials(RestClientCredentials{AccessToken: "caller-access"}))
	if _, err := c.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(transport.seen) != 1 || transport.seen[0].Get("Authorization") != "Bearer caller-access" {
		t.Fatalf("supplied transport saw %v, want one request with the bearer credential", transport.seen)
	}
}

func TestOptionsApplyAgainstALiveServer(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer caller-access" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"1.0.0","engineSha":"sha","gitTag":"tag","openapiVersion":"1.0.0","abiVersion":"v1"}`)
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL, WithTimeout(time.Second), WithCredentials(RestClientCredentials{AccessToken: "caller-access"}))
	if _, err := c.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}
}
