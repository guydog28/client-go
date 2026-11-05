package ezclassify_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/ezclassify/sdk-go"
)

// TestContextIDSpecialCharEncoding verifies that a context ID containing
// characters requiring percent-encoding (space, slash) is properly escaped
// before placement in the request path. Mirrors the .NET client's
// ConnectAsync_EncodesContextIdSpecialChars test.
func TestContextIDSpecialCharEncoding(t *testing.T) {
	var rawPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// RawPath carries the pre-decode form the client actually transmitted.
		rawPath = r.URL.RawPath
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, processedFixture)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "tenant a/marking", "1.0.0")
	if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"}); err != nil {
		t.Fatalf("Process: %v", err)
	}

	// Space → %20, slash → %2F; "1.0.0" has no characters needing escaping.
	const want = "/v1/ctx/tenant%20a%2Fmarking/1.0.0/scheme/process"
	if rawPath != want {
		t.Fatalf("raw path = %q, want %q", rawPath, want)
	}
}

// TestContextNotFoundError verifies that Resolve returns a *ContextNotFoundError
// when the server responds with 404, allowing callers to distinguish a missing
// context from other server errors via errors.As.
func TestContextNotFoundError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":"context not found"}`)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "missing-ctx", "1.0.0")
	err := cx.Resolve(context.Background())
	if err == nil {
		t.Fatal("Resolve: expected error on 404, got nil")
	}

	var notFound *ezclassify.ContextNotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Resolve error = %T(%v), want *ContextNotFoundError", err, err)
	}
	if notFound.ContextID != "missing-ctx" {
		t.Fatalf("ContextNotFoundError.ContextID = %q, want %q", notFound.ContextID, "missing-ctx")
	}
}

// TestResolvedVersionUpdatesFromEveryOp verifies that ResolvedVersion() tracks
// the version header from every op response, not just the first one. Mirrors
// the .NET ResolvedVersion_UpdatesFromEveryOpResponse test.
func TestResolvedVersionUpdatesFromEveryOp(t *testing.T) {
	var n atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if n.Add(1) == 1 {
			w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		} else {
			w.Header().Set("X-EzClassify-Resolved-Version", "1.4.2")
		}
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, processedFixture)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "acme-marking", "latest")

	if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"}); err != nil {
		t.Fatalf("first Process: %v", err)
	}
	if got := cx.ResolvedVersion(); got != "1.0.0" {
		t.Fatalf("after first op: ResolvedVersion() = %q, want 1.0.0", got)
	}

	// A different request, since an identical one is answered from the cache.
	if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "L"}); err != nil {
		t.Fatalf("second Process: %v", err)
	}
	if got := cx.ResolvedVersion(); got != "1.4.2" {
		t.Fatalf("after second op: ResolvedVersion() = %q, want 1.4.2", got)
	}
}

// TestUndecodableSuccessIsAMalformedResponse verifies that a 2xx whose body
// cannot be decoded is refused as a *MalformedResponseError on the cached, the
// uncached and the client-level paths, while an undecodable error body keeps
// its decode error, since the status there already says the call failed.
func TestUndecodableSuccessIsAMalformedResponse(t *testing.T) {
	calls := map[string]func(*ezclassify.Client) error{
		"Scheme.Process": func(c *ezclassify.Client) error {
			_, err := ezclassify.ConnectWith(c, "ctx", "1.0.0").Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			return err
		},
		"Describe": func(c *ezclassify.Client) error {
			_, err := ezclassify.ConnectWith(c, "ctx", "1.0.0").Describe(context.Background())
			return err
		},
		"ListContexts": func(c *ezclassify.Client) error {
			_, err := c.ListContexts(context.Background())
			return err
		},
		"ListContextVersions": func(c *ezclassify.Client) error {
			_, err := c.ListContextVersions(context.Background(), "ctx")
			return err
		},
	}
	for _, body := range []string{"", `{not-valid`} {
		for name, call := range calls {
			for _, status := range []int{http.StatusOK, http.StatusBadRequest} {
				t.Run(fmt.Sprintf("%s/%d/%q", name, status, body), func(t *testing.T) {
					srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						w.Header().Set("Content-Type", "application/json")
						w.WriteHeader(status)
						_, _ = io.WriteString(w, body)
					}))
					t.Cleanup(srv.Close)
					err := call(ezclassify.New(srv.URL))
					var malformed *ezclassify.MalformedResponseError
					isMalformed := errors.As(err, &malformed)
					if status == http.StatusOK && (!isMalformed || malformed.StatusCode != status) {
						t.Fatalf("error = %T(%v), want a *MalformedResponseError for the 200", err, err)
					}
					if status != http.StatusOK && (err == nil || isMalformed) {
						t.Fatalf("error = %T(%v), want the decode error of a failed call", err, err)
					}
				})
			}
		}
	}
}

// TestRouteNotFoundIsNotAContextNotFound verifies that a 404 the server
// classifies as an unmatched route surfaces as a plain APIError. Reporting it
// as a missing context would send a caller whose base URL is wrong looking for
// a context that is present.
func TestRouteNotFoundIsNotAContextNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":"Not found: /v1/unknown","code":"route_not_found"}`)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "present-context", "1.0.0")
	_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	var notFound *ezclassify.ContextNotFoundError
	if errors.As(err, &notFound) {
		t.Fatalf("route 404 reported as a missing context: %v", err)
	}
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		t.Fatalf("error = %T(%v), want *APIError with status 404", err, err)
	}
}

// TestResourceNotFoundStaysAContextNotFound pins the other side of the code
// branch: a 404 without the route classifier still reports the context binding.
func TestResourceNotFoundStaysAContextNotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":"Context not found: missing-context"}`)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "missing-context", "1.0.0")
	_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	var notFound *ezclassify.ContextNotFoundError
	if !errors.As(err, &notFound) || notFound.ContextID != "missing-context" {
		t.Fatalf("error = %T(%v), want *ContextNotFoundError for missing-context", err, err)
	}
}

// TestConnectWithRejectsMissingDependencies pins that ConnectWith fails at the
// call that made the mistake rather than inside the first operation.
func TestConnectWithRejectsMissingDependencies(t *testing.T) {
	cases := map[string]func(){
		"nil client":       func() { ezclassify.ConnectWith(nil, "ctx", "1.0.0") },
		"empty context id": func() { ezclassify.ConnectWith(ezclassify.New("http://localhost:9999"), "", "1.0.0") },
	}
	for name, connect := range cases {
		t.Run(name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("ConnectWith returned a handle instead of rejecting the call")
				}
			}()
			connect()
		})
	}
}
