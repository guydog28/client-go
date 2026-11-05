package ezclassify_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
)

// TestUnknownResponseFieldsAreTolerated is this SDK's conformance case for the
// response-openness contract: every schema reachable from a response is open,
// new fields are additive, and a client built against an older spec must read
// a newer server's payload without failing. The stub returns fields that are
// not in the vendored spec at the envelope, at the result, and inside a nested
// array element; every known value must still arrive intact.
//
// Go's encoding/json ignores unknown members by default, so this test exists
// to keep it that way: a DisallowUnknownFields decoder anywhere in the read
// path, generated or handwritten, fails here.
func TestUnknownResponseFieldsAreTolerated(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"H","inputWasValid":true,` +
			`"assertions":[{"conceptId":"L_H","categoryId":"level","confidence":"certain","source":"parse",` +
			`"futureAssertionField":{"nested":[1,2,3]}}],` +
			`"futureResultField":"added by a newer engine"}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	processed, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("Process against a newer payload: %v", err)
	}
	if !processed.Result.Valid || processed.Result.Rendered != "H" {
		t.Fatalf("known result fields = %+v, want valid with rendered H", processed.Result)
	}
	if len(processed.Result.Assertions) != 1 {
		t.Fatalf("len(Assertions) = %d, want 1", len(processed.Result.Assertions))
	}
	if got := processed.Result.Assertions[0].ConceptID; got != "L_H" {
		t.Errorf("Assertions[0].ConceptID = %q, want %q", got, "L_H")
	}
	if got := processed.Resolution.ResolvedVersion; got != ver {
		t.Errorf("Resolution.ResolvedVersion = %q, want %q", got, ver)
	}
}

// TestUnknownEnumMemberIsReadableAndReportsInvalid pins the other half of
// response openness: a value a newer server adds to a wire enum reaches the
// caller as its wire string rather than failing the decode, and Valid()
// reports that this build does not know it. The same payload carries a member
// name that is not in the contract at all, so the case also covers an unknown
// field arriving on a nested array element.
func TestUnknownEnumMemberIsReadableAndReportsInvalid(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/notices": `[{"conceptId":"N1","conceptName":"Notice","type":"future_kind",` +
			`"text":"t","futureNoticeField":"added by a newer engine","required":true}]`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	notices, err := cx.Scheme.Notices(ctx, ezclassify.CtxSchemeNoticesRequest{})
	if err != nil {
		t.Fatalf("Notices against a newer payload: %v", err)
	}
	if len(notices.Result) != 1 {
		t.Fatalf("len(Result) = %d, want 1", len(notices.Result))
	}
	if got := notices.Result[0].Text; got != "t" {
		t.Errorf("Text = %q, want the known field to survive alongside the unknown one", got)
	}
	kind := notices.Result[0].Type
	if string(kind) != "future_kind" {
		t.Errorf("Type = %q, want the raw wire value", kind)
	}
	if kind.Valid() {
		t.Error("Valid() must report an enum member this build does not know")
	}
}

// TestUnknownEnumValuesDecodeAtEveryPosition extends the enum half of
// response openness to every kind of place an enum sits in a response: a
// field of the result, an element of an enum list and of a list of models, a
// model nested inside the result, and the discriminator of a source-reference
// union. Every generated enum is a string-backed type, so each unknown value
// reaches the caller as its wire string, Valid() reports it as unknown, and
// it marshals back unchanged. A custom UnmarshalJSON that rejected unknown
// members, generated or handwritten, fails here.
func TestUnknownEnumValuesDecodeAtEveryPosition(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"

	t.Run("top-level field", func(t *testing.T) {
		srv, _ := ctxServer(t, ver, map[string]string{
			"/document/audit": `{"wholes":[],"parts":[],"unmarkedSpans":[],"proceduralFlags":[],` +
				`"overallVerdict":"needs-review"}`,
		})
		rep, err := ezclassify.Connect(srv.URL, "c", ver).Document.Audit(ctx, ezclassify.CtxDocumentAuditRequest{Text: "d"})
		if err != nil {
			t.Fatalf("Document.Audit against a newer verdict: %v", err)
		}
		verdict := rep.Result.OverallVerdict
		if verdict != "needs-review" || verdict.Valid() {
			t.Errorf("OverallVerdict = %q (Valid=%v), want the raw unknown value", verdict, verdict.Valid())
		}
	})

	t.Run("list elements", func(t *testing.T) {
		srv, _ := ctxServer(t, ver, map[string]string{
			"/scheme/process": `{"valid":false,"rendered":"H","notices":[],` +
				`"assertions":[{"conceptId":"L_H","categoryId":"level","confidence":"probable","source":"parse"}],` +
				`"errors":[{"constraintId":"c1","message":"m","severity":"fatal"}],` +
				`"warnings":[{"kind":"advisory","message":"w"}]}`,
		})
		processed, err := ezclassify.Connect(srv.URL, "c", ver).Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
		if err != nil {
			t.Fatalf("Scheme.Process against newer enum values: %v", err)
		}
		r := processed.Result
		if len(r.Assertions) != 1 || len(r.Errors) != 1 || len(r.Warnings) != 1 {
			t.Fatalf("result lists = %+v, want one of each", r)
		}
		if c := r.Assertions[0].Confidence; c != "probable" || c.Valid() {
			t.Errorf("Assertions[0].Confidence = %v, want the raw unknown value", c)
		}
		if s := r.Errors[0].Severity; s == nil || *s != "fatal" || s.Valid() {
			t.Errorf("Errors[0].Severity = %v, want the raw unknown value", s)
		}
		if k := r.Warnings[0].Kind; k != "advisory" || k.Valid() {
			t.Errorf("Warnings[0].Kind = %q, want the raw unknown value", k)
		}

		// An enum list itself: the roles a Context gives each scheme.
		describe := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextDescription](`{"id":"c","version":"1.0.0","scheme":{"id":"s","version":"1.0.0"},`+
				`"schemes":[{"id":"s","version":"1.0.0","selector":"s","roles":["primary","reviewer"]}]}`))
		}))
		t.Cleanup(describe.Close)
		if err := ezclassify.Connect(describe.URL, "c", ver).Resolve(ctx); err != nil {
			t.Fatalf("Resolve against a newer scheme role: %v", err)
		}
		described, err := ezclassify.New(describe.URL).CtxDescribeWithResponse(ctx, "c", ver)
		if err != nil || described.JSON200 == nil {
			t.Fatalf("describe decode: err=%v status=%d", err, described.StatusCode())
		}
		roles := described.JSON200.Schemes[0].Roles
		if len(roles) != 2 || !roles[0].Valid() || roles[1] != "reviewer" || roles[1].Valid() {
			t.Errorf("Roles = %v, want [primary reviewer] with the unknown role kept", roles)
		}
	})

	t.Run("nested model", func(t *testing.T) {
		srv, _ := ctxServer(t, ver, map[string]string{
			"/scheme/scan": `{"results":[],"limits":{"status":"paused","limitHit":"memory","regionsFound":0}}`,
		})
		scanned, err := ezclassify.Connect(srv.URL, "c", ver).Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{Text: "t"})
		if err != nil {
			t.Fatalf("Scheme.Scan against newer limit values: %v", err)
		}
		limits := scanned.Result.Limits
		if limits.Status != "paused" || limits.Status.Valid() {
			t.Errorf("Limits.Status = %q, want the raw unknown value", limits.Status)
		}
		if limits.LimitHit == nil || *limits.LimitHit != "memory" || limits.LimitHit.Valid() {
			t.Errorf("Limits.LimitHit = %v, want the raw unknown value", limits.LimitHit)
		}
	})

	t.Run("discriminator", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("X-EzClassify-Resolved-Version", ver)
			_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextSchemeProcessResult](`{"resolution":{"contextId":"c","resolvedVersion":"1.0.0",`+
				`"artifacts":[{"kind":"policy","id":"p","version":"1.0.0"}],`+
				`"sources":[{"providerId":"idp","reproducibility":"pinned","snapshotId":"snap-1",`+
				`"attribution":{"name":"IdP","license":"internal"}}]},`+
				`"operationLimitations":[],"result":{"valid":true,"rendered":"H","notices":[]}}`))
		}))
		t.Cleanup(srv.Close)
		processed, err := ezclassify.Connect(srv.URL, "c", ver).Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
		if err != nil {
			t.Fatalf("Scheme.Process with a newer source variant: %v", err)
		}
		res := processed.Resolution
		if kind := res.Artifacts[0].Kind; kind != "policy" || kind.Valid() {
			t.Errorf("Artifacts[0].Kind = %q, want the raw unknown value", kind)
		}
		source, err := res.Sources[0].AsReproducibleSourceReference()
		if err != nil {
			t.Fatalf("reading the unknown source variant: %v", err)
		}
		if source.Reproducibility != "pinned" || source.SnapshotID != "snap-1" {
			t.Errorf("source = %+v, want the unknown discriminator and its fields kept", source)
		}
		wire, err := json.Marshal(res)
		if err != nil {
			t.Fatalf("re-encoding the resolution: %v", err)
		}
		for _, want := range []string{`"kind":"policy"`, `"reproducibility":"pinned"`} {
			if !strings.Contains(string(wire), want) {
				t.Errorf("re-encoded resolution %s lacks %s", wire, want)
			}
		}
	})
}
