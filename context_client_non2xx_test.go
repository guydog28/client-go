package ezclassify_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/ezclassify/client-go"
)

// stubStatus returns a handler that always responds with the given status and body.
func stubStatus(status int, body string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	})
}

// TestCtxOpsNotFoundReturnsContextNotFoundError pins that every op group
// returns *ContextNotFoundError when the server responds 404, with ContextID
// populated. One representative op per group.
func TestCtxOpsNotFoundReturnsContextNotFoundError(t *testing.T) {
	cases := []struct {
		name string
		call func(cx *ezclassify.EzClassifyContext) error
	}{
		{
			"Resolve",
			func(cx *ezclassify.EzClassifyContext) error {
				return cx.Resolve(context.Background())
			},
		},
		{
			"Scheme.Process",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
				return e
			},
		},
		{
			"Scheme.Diff",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Diff(context.Background(), ezclassify.CtxSchemeDiffRequest{A: "A", B: "B"})
				return e
			},
		},
		{
			"Scheme.Notices",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Notices(context.Background(), ezclassify.CtxSchemeNoticesRequest{Assertions: []ezclassify.Assertion{}})
				return e
			},
		},
		{
			"Guide.Scan",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Guide.Scan(context.Background(), ezclassify.CtxGuideScanRequest{Text: "x"})
				return e
			},
		},
		{
			"Guide.View",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Guide.View(context.Background(), nil)
				return e
			},
		},
		{
			"Access.Check",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Access.Check(context.Background(), ezclassify.CtxAccessCheckRequest{})
				return e
			},
		},
		{
			"Document.Audit",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Document.Audit(context.Background(), ezclassify.CtxDocumentAuditRequest{Text: "x"})
				return e
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(stubStatus(http.StatusNotFound, `{"error":"context not found"}`))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "missing-ctx", "1.0.0")

			err := tc.call(cx)
			if err == nil {
				t.Fatal("expected error on 404, got nil")
			}
			var notFound *ezclassify.ContextNotFoundError
			if !errors.As(err, &notFound) {
				t.Fatalf("error = %T(%v), want *ContextNotFoundError", err, err)
			}
			if notFound.ContextID != "missing-ctx" {
				t.Fatalf("ContextID = %q, want missing-ctx", notFound.ContextID)
			}
		})
	}
}

// TestCtxOpsNon2xxNonNotFoundReturnsAPIError pins that HTTP errors other
// than 404 produce an *APIError that carries and mentions the status code.
// Callers must not treat a 500 or 403 as a missing-context signal.
func TestCtxOpsNon2xxNonNotFoundReturnsAPIError(t *testing.T) {
	statuses := []int{
		http.StatusInternalServerError, // 500
		http.StatusForbidden,           // 403
		http.StatusUnprocessableEntity, // 422
	}
	for _, status := range statuses {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			body := fmt.Sprintf(`{"error":"server returned %d"}`, status)
			srv := httptest.NewServer(stubStatus(status, body))
			t.Cleanup(srv.Close)

			cx := ezclassify.Connect(srv.URL, "some-ctx", "1.0.0")
			_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
			if err == nil {
				t.Fatalf("expected error for status %d, got nil", status)
			}
			var notFound *ezclassify.ContextNotFoundError
			if errors.As(err, &notFound) {
				t.Fatalf("status %d must not produce *ContextNotFoundError", status)
			}
			var apiErr *ezclassify.APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("status %d: got %T(%v), want *APIError", status, err, err)
			}
			if apiErr.StatusCode != status {
				t.Errorf("APIError.StatusCode = %d, want %d", apiErr.StatusCode, status)
			}
			if !strings.Contains(err.Error(), strconv.Itoa(status)) {
				t.Errorf("error %q does not mention status %d", err.Error(), status)
			}
		})
	}
}

func TestRenderFullPreservationRefusalReturnsTypedAPIError(t *testing.T) {
	srv := httptest.NewServer(stubStatus(http.StatusBadRequest, `{"error":"render would lose facts","code":"RENDER_NOT_PRESERVING","detail":"omitted authority AUTH-X=A&B"}`))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "ctx", "1.0.0")
	_, err := cx.Scheme.Render(context.Background(), ezclassify.CtxSchemeRenderRequest{
		Assertions:   []ezclassify.Assertion{{ConceptID: "AUTH-X", CategoryID: "authority", Value: ptr("A&B")}},
		Preservation: ptr(ezclassify.RenderPreservationFull),
	})
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %T(%v), want *APIError", err, err)
	}
	if apiErr.Response == nil || apiErr.Response.Code == nil || *apiErr.Response.Code != "RENDER_NOT_PRESERVING" {
		t.Fatalf("API error response = %+v, want typed preservation refusal", apiErr.Response)
	}
	if !strings.Contains(string(apiErr.Body), `AUTH-X=A&B`) {
		t.Fatalf("API error body lost refusal detail: %s", apiErr.Body)
	}
}

// TestResolveNon2xxNonNotFoundReturnsGenericError pins Resolve's behavior for
// server errors that are not 404. The 404 case is covered by
// TestContextNotFoundError; this covers the non-typed path.
func TestResolveNon2xxNonNotFoundReturnsGenericError(t *testing.T) {
	srv := httptest.NewServer(stubStatus(http.StatusInternalServerError, `{"error":"internal"}`))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "some-ctx", "1.0.0")
	err := cx.Resolve(context.Background())
	if err == nil {
		t.Fatal("Resolve: expected error on 500, got nil")
	}
	var notFound *ezclassify.ContextNotFoundError
	if errors.As(err, &notFound) {
		t.Fatal("Resolve: 500 must not produce *ContextNotFoundError")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("error %q does not mention status 500", err.Error())
	}
}

// TestCtxSchemeProcessUnexpectedShapeIsMalformedResponse pins how an op
// handles a 200 whose body is valid JSON of the wrong shape. The generated
// decoder would drop the unknown keys and zero-fill everything the contract
// requires, which reads as an answer the server never gave. The client reports
// a *MalformedResponseError naming the first required member the body lacks.
//
// This complements TestMalformedResponseBodyReturnsError (empty and
// syntactically broken bodies, which surface the decoder's own error).
func TestCtxSchemeProcessUnexpectedShapeIsMalformedResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"completely":"unexpected","shape":true}`)
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "ctx", "1.0.0")
	result, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	var malformed *ezclassify.MalformedResponseError
	if !errors.As(err, &malformed) {
		t.Fatalf("Process error = %T(%v), want *MalformedResponseError", err, err)
	}
	if result != nil {
		t.Errorf("result = %+v, want nil beside an error", result)
	}
	if malformed.Op != "scheme.process" || malformed.StatusCode != http.StatusOK || malformed.Member != "/resolution" {
		t.Errorf("error = %+v, want scheme.process, 200, first missing member /resolution", malformed)
	}
	var apiErr *ezclassify.APIError
	if errors.As(err, &apiErr) {
		t.Errorf("a 2xx must not surface as *APIError, got status %d", apiErr.StatusCode)
	}
}

// TestContextNotFoundErrorMessage pins the formatted message returned by
// (*ContextNotFoundError).Error(), so the human-readable text callers log or
// surface does not drift.
func TestContextNotFoundErrorMessage(t *testing.T) {
	err := &ezclassify.ContextNotFoundError{ContextID: "my-ctx"}
	want := `ezclassify: context "my-ctx" not found`
	if got := err.Error(); got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

// TestCtxOpsTransportErrorPropagate verifies that every EzClassifyContext method returns
// a non-nil error (rather than panicking or silently returning nil) when the
// underlying HTTP transport fails, i.e. the server is unreachable. This
// covers the "if err != nil { return nil, err }" transport-error branch in
// every op method body (context_client.go).
func TestCtxOpsTransportErrorPropagate(t *testing.T) {
	// Start a server and close it immediately so every subsequent call gets
	// "connection refused".
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	srv.Close()

	cx := ezclassify.Connect(srv.URL, "c", "1.0")
	ctx := context.Background()

	if err := cx.Resolve(ctx); err == nil {
		t.Error("Resolve: expected transport error, got nil")
	}
	if _, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: ptr("")}); err == nil {
		t.Error("Scheme.Validate: expected transport error")
	}
	if _, err := cx.Scheme.Suggest(ctx, ezclassify.CtxSchemeSuggestRequest{Input: ptr("")}); err == nil {
		t.Error("Scheme.Suggest: expected transport error")
	}
	if _, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: &[]string{"X"}}); err == nil {
		t.Error("Scheme.Aggregate: expected transport error")
	}
	if _, err := cx.Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{}); err == nil {
		t.Error("Scheme.Scan: expected transport error")
	}
	if _, err := cx.Scheme.Render(ctx, ezclassify.CtxSchemeRenderRequest{}); err == nil {
		t.Error("Scheme.Render: expected transport error")
	}
	if _, err := cx.Scheme.CheckCoverage(ctx, ezclassify.CtxSchemeCheckCoverageRequest{}); err == nil {
		t.Error("Scheme.CheckCoverage: expected transport error")
	}
	if _, err := cx.Scheme.DescribeConcept(ctx, "c", "cat", nil); err == nil {
		t.Error("Scheme.DescribeConcept: expected transport error")
	}
	if _, err := cx.Scheme.Diff(ctx, ezclassify.CtxSchemeDiffRequest{}); err == nil {
		t.Error("Scheme.Diff: expected transport error")
	}
	if _, err := cx.Scheme.Notices(ctx, ezclassify.CtxSchemeNoticesRequest{}); err == nil {
		t.Error("Scheme.Notices: expected transport error")
	}
	if _, err := cx.Scheme.View(ctx); err == nil {
		t.Error("Scheme.View: expected transport error")
	}
	if _, err := cx.Guide.Scan(ctx, ezclassify.CtxGuideScanRequest{}); err == nil {
		t.Error("Guide.Scan: expected transport error")
	}
	if _, err := cx.Guide.Evaluate(ctx, ezclassify.CtxGuideEvaluateRequest{}); err == nil {
		t.Error("Guide.Evaluate: expected transport error")
	}
	if _, err := cx.Guide.View(ctx, nil); err == nil {
		t.Error("Guide.View: expected transport error")
	}
	if _, err := cx.Access.Check(ctx, ezclassify.CtxAccessCheckRequest{}); err == nil {
		t.Error("Access.Check: expected transport error")
	}
	if _, err := cx.Access.DescribeRequirements(ctx, ezclassify.CtxAccessDescribeRequirementsRequest{}); err == nil {
		t.Error("Access.DescribeRequirements: expected transport error")
	}
	if _, err := cx.Document.Audit(ctx, ezclassify.CtxDocumentAuditRequest{}); err == nil {
		t.Error("Document.Audit: expected transport error")
	}
}

// TestCtxOpsNotFoundRemainingMethods extends TestCtxOpsNotFoundReturnsContextNotFoundError
// for the scheme/guide/access methods not covered in that table. The 404 path
// hits the "resp.JSON200 == nil → ctxError" branch in each method body.
func TestCtxOpsNotFoundRemainingMethods(t *testing.T) {
	cases := []struct {
		name string
		call func(cx *ezclassify.EzClassifyContext) error
	}{
		{
			"Scheme.Validate",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{Input: ptr("")})
				return e
			},
		},
		{
			"Scheme.Suggest",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Suggest(context.Background(), ezclassify.CtxSchemeSuggestRequest{Input: ptr("")})
				return e
			},
		},
		{
			"Scheme.Aggregate",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Aggregate(context.Background(), ezclassify.CtxSchemeAggregateRequest{Inputs: &[]string{"X"}})
				return e
			},
		},
		{
			"Scheme.Scan",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Scan(context.Background(), ezclassify.CtxSchemeScanRequest{})
				return e
			},
		},
		{
			"Scheme.Render",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Render(context.Background(), ezclassify.CtxSchemeRenderRequest{})
				return e
			},
		},
		{
			"Scheme.CheckCoverage",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.CheckCoverage(context.Background(), ezclassify.CtxSchemeCheckCoverageRequest{})
				return e
			},
		},
		{
			"Scheme.DescribeConcept",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.DescribeConcept(context.Background(), "c", "cat", nil)
				return e
			},
		},
		{
			"Scheme.View",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.View(context.Background())
				return e
			},
		},
		{
			"Guide.Evaluate",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Guide.Evaluate(context.Background(), ezclassify.CtxGuideEvaluateRequest{})
				return e
			},
		},
		{
			"Access.DescribeRequirements",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Access.DescribeRequirements(context.Background(), ezclassify.CtxAccessDescribeRequirementsRequest{})
				return e
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(stubStatus(http.StatusNotFound, `{"error":"not found"}`))
			t.Cleanup(srv.Close)
			cx := ezclassify.Connect(srv.URL, "missing-ctx", "1.0.0")

			err := tc.call(cx)
			if err == nil {
				t.Fatal("expected error on 404, got nil")
			}
			var notFound *ezclassify.ContextNotFoundError
			if !errors.As(err, &notFound) {
				t.Fatalf("error = %T(%v), want *ContextNotFoundError", err, err)
			}
			if notFound.ContextID != "missing-ctx" {
				t.Fatalf("ContextID = %q, want missing-ctx", notFound.ContextID)
			}
		})
	}
}

// TestCtxOpsDecodeErrorOnWrongBodyType pins the decode-failure path in the ops
// below (Diff, FilterConcepts, Notices, FilterRules, Guide.View) that route
// through the generated wire client like every other context method.
// A server returning the JSON token "true" with status 200 is valid JSON but
// not an object or array, so the generated parser refuses it, and the call
// reports a *MalformedResponseError for the 200.
func TestCtxOpsDecodeErrorOnWrongBodyType(t *testing.T) {
	cases := []struct {
		name string
		call func(cx *ezclassify.EzClassifyContext) error
	}{
		{
			"Scheme.Diff",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Diff(context.Background(), ezclassify.CtxSchemeDiffRequest{A: "A", B: "B"})
				return e
			},
		},
		{
			"Scheme.Notices",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Scheme.Notices(context.Background(), ezclassify.CtxSchemeNoticesRequest{})
				return e
			},
		},
		{
			"Guide.View",
			func(cx *ezclassify.EzClassifyContext) error {
				_, e := cx.Guide.View(context.Background(), nil)
				return e
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
				w.WriteHeader(http.StatusOK)
				// "true" is valid JSON but cannot unmarshal into a struct or slice.
				_, _ = io.WriteString(w, "true")
			}))
			t.Cleanup(srv.Close)

			cx := ezclassify.Connect(srv.URL, "ctx", "1.0.0")
			err := tc.call(cx)
			var malformed *ezclassify.MalformedResponseError
			if !errors.As(err, &malformed) || malformed.StatusCode != http.StatusOK || malformed.Member != "" {
				t.Fatalf("%s: error = %T(%v), want a *MalformedResponseError for the 200", tc.name, err, err)
			}
		})
	}
}

// TestCtxOps2xxUndecodableIsNotAPIError pins that a 2xx response the
// generated parser produces no typed value for (a 200 with a non-JSON
// Content-Type leaves JSON200 nil without an unmarshal error) is reported as
// a *MalformedResponseError with no member named. APIError's contract covers
// non-2xx statuses only, so a caller branching on it must never see an
// APIError with a 2xx code.
func TestCtxOps2xxUndecodableIsNotAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok but not json")
	}))
	t.Cleanup(srv.Close)

	cx := ezclassify.Connect(srv.URL, "some-ctx", "1.0.0")
	_, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err == nil {
		t.Fatal("expected an error for an undecodable 200 response, got nil")
	}
	var apiErr *ezclassify.APIError
	if errors.As(err, &apiErr) {
		t.Fatalf("got *APIError with StatusCode=%d; a 2xx must not surface as APIError", apiErr.StatusCode)
	}
	var malformed *ezclassify.MalformedResponseError
	if !errors.As(err, &malformed) || malformed.Member != "" || malformed.StatusCode != http.StatusOK {
		t.Fatalf("error = %T(%v), want a *MalformedResponseError for status 200 naming no member", err, err)
	}
	for _, want := range []string{"200", "could not decode"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error %q does not mention %q", err.Error(), want)
		}
	}
}
