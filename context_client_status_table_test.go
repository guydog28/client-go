package ezclassify_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
)

// TestStatusToTypedErrorTable pins the complete HTTP status → error-type
// mapping for context operations. The mapping lives in ctxError
// (context_client.go):
//
//   - 404 → *ContextNotFoundError with ContextID populated, unless the body
//     classifies it as an unmatched route or a missing concept
//   - any other non-2xx → *APIError whose message includes the numeric
//     status code and the raw response body bytes
//
// The table covers 400, 401, 403, 404, 422, 429, 500, 502, and 503.
// Body-content propagation is verified here in addition to the status code.
//
// Bodies are always valid JSON so the generated response parser does not
// short-circuit with a decode error before ctxError is reached. See
// TestCtx404NonJsonBodyBehavior for the non-JSON case.
func TestStatusToTypedErrorTable(t *testing.T) {
	cases := []struct {
		status    int
		wantTyped bool // true = *ContextNotFoundError expected
	}{
		{http.StatusBadRequest, false},          // 400
		{http.StatusUnauthorized, false},        // 401
		{http.StatusForbidden, false},           // 403
		{http.StatusNotFound, true},             // 404 → *ContextNotFoundError
		{http.StatusUnprocessableEntity, false}, // 422
		{http.StatusTooManyRequests, false},     // 429
		{http.StatusInternalServerError, false}, // 500
		{http.StatusBadGateway, false},          // 502
		{http.StatusServiceUnavailable, false},  // 503
	}

	for _, tc := range cases {
		tc := tc
		t.Run(strconv.Itoa(tc.status), func(t *testing.T) {
			bodyStr := fmt.Sprintf(`{"error":"simulated-%d"}`, tc.status)
			srv := httptest.NewServer(stubStatus(tc.status, bodyStr))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "ctx-id", "1.0.0")

			_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			if err == nil {
				t.Fatalf("status %d: expected error, got nil", tc.status)
			}

			var notFound *ezclassify.ContextNotFoundError
			if tc.wantTyped {
				if !errors.As(err, &notFound) {
					t.Fatalf("status %d: got %T(%v), want *ContextNotFoundError", tc.status, err, err)
				}
				if notFound.ContextID != "ctx-id" {
					t.Fatalf("status %d: ContextID = %q, want ctx-id", tc.status, notFound.ContextID)
				}
			} else {
				if errors.As(err, &notFound) {
					t.Fatalf("status %d: must not produce *ContextNotFoundError", tc.status)
				}
				var apiErr *ezclassify.APIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("status %d: got %T(%v), want *APIError", tc.status, err, err)
				}
				if apiErr.StatusCode != tc.status {
					t.Errorf("status %d: APIError.StatusCode = %d", tc.status, apiErr.StatusCode)
				}
				if !strings.Contains(err.Error(), strconv.Itoa(tc.status)) {
					t.Errorf("status %d: error %q does not mention the status code", tc.status, err.Error())
				}
				if !strings.Contains(err.Error(), fmt.Sprintf("simulated-%d", tc.status)) {
					t.Errorf("status %d: error %q does not include response body", tc.status, err.Error())
				}
			}
		})
	}
}

// TestResolveStatusMapping verifies that Resolve applies the same ctxError
// status-to-error-type mapping as Scheme.Process. The op string differs
// ("resolve context" vs "scheme.process") but the mapping is shared.
func TestResolveStatusMapping(t *testing.T) {
	cases := []struct {
		status    int
		wantTyped bool
	}{
		{http.StatusBadRequest, false},          // 400
		{http.StatusUnauthorized, false},        // 401
		{http.StatusNotFound, true},             // 404 → *ContextNotFoundError
		{http.StatusInternalServerError, false}, // 500
		{http.StatusServiceUnavailable, false},  // 503
	}

	for _, tc := range cases {
		tc := tc
		t.Run(strconv.Itoa(tc.status), func(t *testing.T) {
			bodyStr := fmt.Sprintf(`{"error":"resolve-body-%d"}`, tc.status)
			srv := httptest.NewServer(stubStatus(tc.status, bodyStr))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "resolve-ctx", "1.0.0")

			err := cx.Resolve(context.Background())
			if err == nil {
				t.Fatalf("status %d: expected error, got nil", tc.status)
			}

			var notFound *ezclassify.ContextNotFoundError
			if tc.wantTyped {
				if !errors.As(err, &notFound) {
					t.Fatalf("status %d: got %T(%v), want *ContextNotFoundError", tc.status, err, err)
				}
				if notFound.ContextID != "resolve-ctx" {
					t.Fatalf("status %d: ContextID = %q, want resolve-ctx", tc.status, notFound.ContextID)
				}
			} else {
				if errors.As(err, &notFound) {
					t.Fatalf("status %d: must not produce *ContextNotFoundError", tc.status)
				}
				var apiErr *ezclassify.APIError
				if !errors.As(err, &apiErr) {
					t.Fatalf("status %d: got %T(%v), want *APIError", tc.status, err, err)
				}
				if apiErr.StatusCode != tc.status {
					t.Errorf("status %d: APIError.StatusCode = %d", tc.status, apiErr.StatusCode)
				}
				if !strings.Contains(err.Error(), strconv.Itoa(tc.status)) {
					t.Errorf("status %d: error %q does not mention the status code", tc.status, err.Error())
				}
				if !strings.Contains(err.Error(), fmt.Sprintf("resolve-body-%d", tc.status)) {
					t.Errorf("status %d: error %q does not include response body", tc.status, err.Error())
				}
			}
		})
	}
}

// TestContextNotFoundForUnclassified404Bodies pins that a 404 response with
// valid JSON produces *ContextNotFoundError for bodies that do not carry the
// route_not_found or concept_not_found classifiers.
//
// Only valid-JSON bodies are tested here. A 404 with a non-JSON body returns
// a JSON parse error from the generated response parser before ctxError is
// reached. That case is covered separately in TestCtx404NonJsonBodyBehavior.
func TestContextNotFoundForUnclassified404Bodies(t *testing.T) {
	bodies := []struct {
		name string
		body string
	}{
		{"standard_message", `{"error":"context not found"}`},
		{"empty_object", `{}`},
		{"unrelated_message", `{"error":"something completely different"}`},
		{"null", `null`},
		{"numeric_error", `{"error":"12345"}`},
	}
	for _, tc := range bodies {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(stubStatus(http.StatusNotFound, tc.body))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "cx", "1.0.0")

			_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			var notFound *ezclassify.ContextNotFoundError
			if !errors.As(err, &notFound) {
				t.Fatalf("body=%q: got %T(%v), want *ContextNotFoundError", tc.body, err, err)
			}
		})
	}
}

// TestCtx404NonJsonBodyBehavior pins that a 404 response with a non-JSON body
// returns a JSON parse error rather than *ContextNotFoundError. The generated
// response parser tries to unmarshal the body before ctxError is called; if
// that unmarshal fails the parse error propagates directly.
//
// The real API always sends JSON on error responses, so this edge case should
// not occur in practice. It is pinned here so that any future hardening of the
// parse path, e.g. falling through to ctxError when JSON parse fails, shows
// up as a deliberate test change rather than a silent regression.
func TestCtx404NonJsonBodyBehavior(t *testing.T) {
	cases := []struct {
		name string
		body string
	}{
		{"plain_text", "not json at all"},
		{"empty_body", ""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(stubStatus(http.StatusNotFound, tc.body))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "cx", "1.0.0")

			_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			if err == nil {
				t.Fatal("expected error, got nil")
			}
			// With a non-JSON body, the generated response parser short-circuits
			// with a decode error. *ContextNotFoundError is NOT returned.
			var notFound *ezclassify.ContextNotFoundError
			if errors.As(err, &notFound) {
				t.Fatalf("body=%q: got *ContextNotFoundError, expected a parse error for non-JSON 404 body", tc.body)
			}
		})
	}
}

// TestGenericErrorContainsStatusAndBody pins that non-404 context-op errors
// carry both the numeric status code and the raw response body in their
// message. The body content is important for debugging: callers should be
// able to read the API's error payload from the error string.
func TestGenericErrorContainsStatusAndBody(t *testing.T) {
	const errBody = `{"code":"RATE_LIMITED","retry_after":30}`
	srv := httptest.NewServer(stubStatus(http.StatusTooManyRequests, errBody))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "acme", "1.0.0")

	_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "429") {
		t.Errorf("error %q does not contain status 429", msg)
	}
	if !strings.Contains(msg, "RATE_LIMITED") {
		t.Errorf("error %q does not contain body field RATE_LIMITED", msg)
	}
	if !strings.Contains(msg, "retry_after") {
		t.Errorf("error %q does not contain body field retry_after", msg)
	}
}
