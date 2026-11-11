package ezclassify_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ezclassify/client-go"
)

// TestAPIErrorMessage pins the formatted message returned by
// (*APIError).Error(), which callers may log or string-match.
func TestAPIErrorMessage(t *testing.T) {
	err := &ezclassify.APIError{
		Op:         "scheme.process",
		StatusCode: 500,
		Body:       []byte(`{"error":"boom"}`),
	}
	want := `ezclassify: scheme.process: rest-api returned 500: {"error":"boom"}`
	if got := err.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

// TestCtxOpsNon2xxReturnsAPIError pins that non-2xx context-op failures match
// *APIError via errors.As, with Op, StatusCode, and the raw body preserved.
func TestCtxOpsNon2xxReturnsAPIError(t *testing.T) {
	statuses := []int{
		http.StatusBadRequest,          // 400
		http.StatusForbidden,           // 403
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
	}
	for _, status := range statuses {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			body := fmt.Sprintf(`{"error":"simulated-%d"}`, status)
			srv := httptest.NewServer(stubStatus(status, body))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "some-ctx", "1.0.0")

			_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			if err == nil {
				t.Fatalf("status %d: expected error, got nil", status)
			}
			var apiErr *ezclassify.APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("status %d: got %T(%v), want *APIError", status, err, err)
			}
			if apiErr.Op != "scheme.process" {
				t.Errorf("Op = %q, want scheme.process", apiErr.Op)
			}
			if apiErr.StatusCode != status {
				t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, status)
			}
			if string(apiErr.Body) != body {
				t.Errorf("Body = %q, want %q", apiErr.Body, body)
			}
		})
	}
}

// TestContextNotFoundWrapsAPIError pins the 404 taxonomy: the returned error
// matches *ContextNotFoundError AND, through Unwrap, *APIError with the 404
// status and raw body. One errors.As check for *APIError therefore covers
// every non-2xx response.
func TestContextNotFoundWrapsAPIError(t *testing.T) {
	const body = `{"error":"context not found"}`
	srv := httptest.NewServer(stubStatus(http.StatusNotFound, body))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "missing-ctx", "1.0.0")

	_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err == nil {
		t.Fatal("expected error on 404, got nil")
	}

	var notFound *ezclassify.ContextNotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("got %T(%v), want *ContextNotFoundError", err, err)
	}
	if notFound.ContextID != "missing-ctx" {
		t.Errorf("ContextID = %q, want missing-ctx", notFound.ContextID)
	}

	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("errors.As(*APIError) = false for 404; ContextNotFoundError must wrap APIError")
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
	if apiErr.Op != "scheme.process" {
		t.Errorf("Op = %q, want scheme.process", apiErr.Op)
	}
	if string(apiErr.Body) != body {
		t.Errorf("Body = %q, want %q", apiErr.Body, body)
	}
	if notFound.APIError != apiErr {
		t.Error("notFound.APIError and the errors.As result are different values")
	}
}

// TestUnknownSchemeSelectorIsNotContextNotFound pins that a 404 refusing a
// scheme selector stays an APIError carrying its code: the Context exists, so
// a host that drops or re-resolves a missing Context must not act on it.
func TestUnknownSchemeSelectorIsNotContextNotFound(t *testing.T) {
	body := `{"error":"Scheme not found in context: other","code":"scheme_not_found"}`
	srv := httptest.NewServer(stubStatus(http.StatusNotFound, body))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "held-ctx", "1.0.0")

	_, err := cx.Schemes.Get("other").Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "X"})
	var notFound *ezclassify.ContextNotFoundError
	if errors.As(err, &notFound) {
		t.Fatalf("got ContextNotFoundError for a scheme the Context does not hold: %v", err)
	}
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("got %T(%v), want *APIError", err, err)
	}
	if apiErr.StatusCode != http.StatusNotFound || apiErr.Response == nil || apiErr.Response.Code == nil || *apiErr.Response.Code != "scheme_not_found" {
		t.Fatalf("APIError = %+v, want a 404 with code scheme_not_found", apiErr)
	}
}

// TestResolveNotFoundWrapsAPIError pins that Resolve's 404 path carries the
// same wrapped APIError, with Resolve's own op string.
func TestResolveNotFoundWrapsAPIError(t *testing.T) {
	srv := httptest.NewServer(stubStatus(http.StatusNotFound, `{"error":"nope"}`))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "missing-ctx", "1.0.0")

	err := cx.Resolve(context.Background())
	if err == nil {
		t.Fatal("expected error on 404, got nil")
	}
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("got %T(%v), want *APIError via ContextNotFoundError", err, err)
	}
	if apiErr.Op != "resolve context" {
		t.Errorf("Op = %q, want %q", apiErr.Op, "resolve context")
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
}

// TestContextNotFoundErrorUnwrapNil pins the defensive guard in Unwrap: a
// ContextNotFoundError constructed without an APIError (as a caller may do)
// unwraps to an untyped nil, so errors.As(*APIError) is
// false rather than matching a typed-nil pointer that would panic in Error().
func TestContextNotFoundErrorUnwrapNil(t *testing.T) {
	err := &ezclassify.ContextNotFoundError{ContextID: "bare"}
	if unwrapped := errors.Unwrap(err); unwrapped != nil {
		t.Errorf("Unwrap() = %v, want nil", unwrapped)
	}
	var apiErr *ezclassify.APIError
	if errors.As(err, &apiErr) {
		t.Error("errors.As(*APIError) matched a ContextNotFoundError with no APIError")
	}
}
