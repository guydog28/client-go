package ezclassify_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/ezclassify/sdk-go"
)

type recordedReq struct {
	method string
	path   string
	query  string
	body   string
}

func TestCategoryViewPositionOverridesRoundTrip(t *testing.T) {
	positions := ezclassify.CategoryPositionOverrides{"first": -1.25, "second": 2.5}
	category := ezclassify.CategoryView{ID: "level", PositionOverrides: &positions}
	encoded, err := json.Marshal(category)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ezclassify.CategoryView
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.PositionOverrides == nil || (*decoded.PositionOverrides)["first"] != -1.25 || (*decoded.PositionOverrides)["second"] != 2.5 {
		t.Fatalf("positions did not round trip: %s", encoded)
	}
	var omitted ezclassify.CategoryView
	if err := json.Unmarshal([]byte(`{"id":"level"}`), &omitted); err != nil || omitted.PositionOverrides != nil {
		t.Fatalf("omitted positions = %+v, err = %v", omitted.PositionOverrides, err)
	}
	for _, malformed := range []string{`{"positionOverrides":[]}`, `{"positionOverrides":{"first":"1"}}`} {
		if err := json.Unmarshal([]byte(malformed), &omitted); err == nil {
			t.Fatalf("accepted malformed category: %s", malformed)
		}
	}
}

// ctxServer is an httptest stub for the /v1/ctx/* routes. It records every
// request and returns the canned body registered for the matching path suffix,
// always stamping the resolved-version header so the wrapper can surface it.
func ctxServer(t *testing.T, resolved string, routes map[string]string) (*httptest.Server, func() recordedReq) {
	t.Helper()
	var mu sync.Mutex
	var last recordedReq
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		mu.Lock()
		last = recordedReq{method: r.Method, path: r.URL.Path, query: r.URL.RawQuery, body: string(b)}
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		if resolved != "" {
			w.Header().Set("X-EzClassify-Resolved-Version", resolved)
		}
		for suffix, body := range routes {
			if strings.HasSuffix(r.URL.Path, suffix) {
				w.WriteHeader(http.StatusOK)
				if strings.Contains(r.URL.Path, "/scheme/") || strings.Contains(r.URL.Path, "/guide/") || strings.Contains(r.URL.Path, "/access/") || strings.Contains(r.URL.Path, "/document/") {
					projection := ""
					if strings.Contains(r.URL.Path, "/document/") {
						projection = `,"responseProjection":{"profile":"standard","detailsComplete":false,"omittedFields":[],"counts":{"nominations":0,"evidenceSpans":0,"guideUnits":0,"accessRuleResults":0,"contentRequirements":0,"contextRequirements":0}}`
					}
					body = contextFixture(r.URL.Path, `{"resolution":{"contextId":"c","resolvedVersion":"`+resolved+`","artifacts":[],"sources":[]},"operationLimitations":[],"result":`+body+projection+`}`)
				}
				_, _ = io.WriteString(w, body)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":"no recording for `+r.URL.Path+`"}`)
	}))
	t.Cleanup(srv.Close)
	getLast := func() recordedReq {
		mu.Lock()
		defer mu.Unlock()
		return last
	}
	return srv, getLast
}

func ptr[T any](v T) *T { return &v }

func TestConnectDefaultsToLatest(t *testing.T) {
	srv, last := ctxServer(t, "2.4.0", map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`,
	})
	cx := ezclassify.Connect(srv.URL, "acme-ctx", "")
	if cx.Version() != "latest" {
		t.Fatalf("Version() = %q, want latest", cx.Version())
	}
	if cx.ID() != "acme-ctx" {
		t.Fatalf("ID() = %q, want acme-ctx", cx.ID())
	}
	if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"}); err != nil {
		t.Fatalf("Process: %v", err)
	}
	if got := last().path; got != "/v1/ctx/acme-ctx/latest/scheme/process" {
		t.Fatalf("path = %q", got)
	}
	if cx.ResolvedVersion() != "2.4.0" {
		t.Fatalf("ResolvedVersion() = %q, want 2.4.0", cx.ResolvedVersion())
	}
}

// An omitted version leaves the choice to the server: under a Context token
// the runtime answers "latest" with the token's default version. Describe
// sends that request with the token and reports the version the server chose,
// and the binding stays unpinned, so access-decision caching is still refused.
func TestOmittedVersionLetsTheServerChooseUnderAContextToken(t *testing.T) {
	var method, path, token string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method, path, token = r.Method, r.URL.Path, r.Header.Get("x-ezclassify-context-token")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.4.0")
		_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextDescription](`{"id":"acme-ctx","version":"1.4.0"}`))
	}))
	t.Cleanup(srv.Close)
	client := ezclassify.New(srv.URL, ezclassify.WithCredentials(ezclassify.RestClientCredentials{ContextToken: "signed-context"}))
	cx := ezclassify.ConnectWith(client, "acme-ctx", "")
	description, err := cx.Describe(context.Background())
	if err != nil {
		t.Fatalf("Describe: %v", err)
	}
	if method != http.MethodGet || path != "/v1/ctx/acme-ctx/latest" || token != "signed-context" {
		t.Fatalf("describe sent %s %s with token %q, want GET /v1/ctx/acme-ctx/latest with the Context token", method, path, token)
	}
	if description.Version != "1.4.0" || cx.ResolvedVersion() != "1.4.0" {
		t.Fatalf("description version %q, resolved %q, want the token default 1.4.0", description.Version, cx.ResolvedVersion())
	}
	defer func() {
		if recover() == nil {
			t.Fatal("an omitted version must stay unpinned and refuse access-decision caching")
		}
	}()
	_ = ezclassify.ConnectWith(ezclassify.New(srv.URL, ezclassify.WithCacheAccessDecisions()), "acme-ctx", "")
}

func TestResolveSurfacesVersion(t *testing.T) {
	// The describe endpoint has no op suffix; match the bare context path.
	srv, last := ctxServer(t, "1.0.3", map[string]string{
		"/v1/ctx/acme-ctx/latest": fixtureAs[ezclassify.ContextDescription](`{"id":"acme-ctx","version":"1.0.3"}`),
	})
	cx := ezclassify.Connect(srv.URL, "acme-ctx", "latest")
	if err := cx.Resolve(context.Background()); err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	if last().method != http.MethodGet {
		t.Fatalf("describe method = %q, want GET", last().method)
	}
	if cx.ResolvedVersion() != "1.0.3" {
		t.Fatalf("ResolvedVersion() = %q, want 1.0.3", cx.ResolvedVersion())
	}
}

func TestSchemeOps(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, last := ctxServer(t, ver, map[string]string{
		"/scheme/process":                `{"valid":true,"rendered":"H","renders":{"marking/banner":{"rendered":"H","assessment":{"emitted":[{"conceptId":"L_H","categoryId":"level","value":"A&B"}],"omitted":[],"collapsed":[],"conflicts":[],"unrepresentable":[]}},"marking/portion":{"rendered":"(H)","assessment":{"emitted":[],"omitted":[],"collapsed":[],"conflicts":[],"unrepresentable":[]}}},"recognition":{"status":"complete","inputFormat":"marking/banner","role":"whole","surface":"marking","unmatched":[],"corrections":[{"start":0,"end":1,"original":"h","replacement":"H","reason":"case"}]},"inputWasValid":true,"assertions":[{"conceptId":"L_H","categoryId":"level","confidence":"certain","source":"parse","value":"A&B"}],"notices":[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true}],"suggestionsFor":{"level":{"suggestions":[{"conceptId":"L_H","label":"H","categoryId":"level"}],"warnings":[]}}}`,
		"/scheme/validate":               `{"valid":false,"rendered":"H","inputWasValid":false,"errors":[{"constraintId":"_unmatched_token","message":"Could not read HGH. Did you mean HIGH?","didYouMean":["L_H"]}],"recognition":{"status":"partial","inputFormat":"marking/banner","role":"whole","surface":"marking","unmatched":[{"start":2,"end":5,"text":"HGH"}]},"notices":[]}`,
		"/scheme/explain":                `{"input":"h","schemeId":"s","entries":[{"stage":"parse","action":"add","detail":"Read h as H.","conceptId":"L_H"}],"repairs":[],"validationErrors":[],"inferences":[],"finalAssertions":[],"rendered":"H"}`,
		"/scheme/suggest":                `{"suggestions":[{"conceptId":"L_H","label":"H","categoryId":"level"}],"warnings":[]}`,
		"/scheme/suggest-for-categories": `{"level":{"suggestions":[{"conceptId":"L_H","label":"H","categoryId":"level","selected":true}],"warnings":[]}}`,
		"/scheme/aggregate":              `{"valid":true,"rendered":"H","assertions":[],"notices":[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true}]}`,
		"/scheme/scan":                   `{"results":[{"marking":"H","start":5,"end":6,"type":"marking/banner","role":"whole","wrapped":false,"result":{"valid":true,"inputWasValid":true,"rendered":"H","recognition":{"status":"complete","inputFormat":"marking/banner","role":"whole","surface":"marking","unmatched":[]}}}],"limits":{"status":"complete","regionsFound":1}}`,
		"/scheme/render":                 `{"result":"LEVEL-HIGH","assessment":{"emitted":[],"omitted":[],"collapsed":[],"conflicts":[],"unrepresentable":[]}}`,
		"/scheme/normalize":              `{"result":"LEVEL-HIGH"}`,
		"/scheme/check-coverage":         `{"covered":false,"missing":[],"missingLabels":["HIGH"],"excess":[],"excessLabels":[],"rendered":"HIGH"}`,
		"/scheme/describe-concept":       `{"conceptId":"L_H","categoryId":"level","name":"High","lines":["High"]}`,
		"/scheme/diff":                   `{"added":[{"conceptId":"A","categoryId":"c","confidence":"certain","source":"parse"}],"removed":[],"common":[]}`,
		"/scheme/filter-concepts":        `{"visibleIds":["a","b"],"categoryCounts":{},"total":{"visible":2,"total":9}}`,
		"/scheme/notices":                `[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true}]`,
		"/scheme/view":                   `{"id":"s","name":"Scheme","description":"d","categories":[],"uncategorized":[]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	proc, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{
		Input:      "h",
		Outputs:    &[]string{"marking/banner", "marking/portion"},
		SuggestFor: &[]string{"level"},
	})
	if err != nil || !proc.Result.Valid || proc.Result.Rendered != "H" {
		t.Fatalf("Process: %+v err=%v", proc, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/process")
	if got, want := last().body, `{"input":"h","outputs":["marking/banner","marking/portion"],"suggestFor":["level"]}`; got != want {
		t.Fatalf("Process body = %s, want %s", got, want)
	}
	if len(proc.Result.Notices) != 1 || !proc.Result.Notices[0].Required {
		t.Fatalf("Process notices = %+v, want the one required notice off the result", proc.Result.Notices)
	}
	recognition := proc.Result.Recognition
	if recognition == nil || recognition.InputFormat == nil || *recognition.InputFormat != "marking/banner" ||
		recognition.Role == nil || *recognition.Role != ezclassify.ProcessingRoleWhole {
		t.Fatalf("Process recognition = %+v, want the banner view of a whole", recognition)
	}
	if len(recognition.Corrections) != 1 || recognition.Corrections[0].Reason != ezclassify.RecognitionCorrectionReasonCase {
		t.Fatalf("Process corrections = %+v, want the one case correction", recognition.Corrections)
	}
	if len(proc.Result.Renders) != 2 || proc.Result.Renders["marking/banner"].Assessment.Emitted[0].Value == nil || *proc.Result.Renders["marking/banner"].Assessment.Emitted[0].Value != "A&B" {
		t.Fatalf("Process renders = %+v, want both named outputs and typed assertion value", proc.Result.Renders)
	}
	if proc.Result.SuggestionsFor == nil || len((*proc.Result.SuggestionsFor)["level"].Suggestions) != 1 {
		t.Fatalf("Process suggestionsFor = %+v, want the level picker", proc.Result.SuggestionsFor)
	}

	val, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: ptr("H HGH")})
	if err != nil || val.Result.Valid || val.Result.InputWasValid {
		t.Fatalf("Validate: %+v err=%v", val, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/validate")
	if len(val.Result.Errors) != 1 || val.Result.Errors[0].DidYouMean == nil || (*val.Result.Errors[0].DidYouMean)[0] != "L_H" {
		t.Fatalf("Validate errors = %+v, want the unread token with its did-you-mean concept", val.Result.Errors)
	}
	if val.Result.Recognition == nil || len(val.Result.Recognition.Unmatched) != 1 || val.Result.Recognition.Unmatched[0].Text != "HGH" {
		t.Fatalf("Validate recognition = %+v, want the unread span", val.Result.Recognition)
	}

	explained, err := cx.Scheme.Explain(ctx, ezclassify.CtxSchemeExplainRequest{Input: "h"})
	if err != nil || explained.Result.Rendered != "H" || len(explained.Result.Entries) != 1 {
		t.Fatalf("Explain: %+v err=%v", explained, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/explain")

	sug, err := cx.Scheme.Suggest(ctx, ezclassify.CtxSchemeSuggestRequest{Input: ptr("L")})
	if err != nil || len(sug.Result.Suggestions) != 1 || sug.Result.Warnings == nil {
		t.Fatalf("Suggest: %+v err=%v", sug, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/suggest")

	byCategory, err := cx.Scheme.SuggestForCategories(ctx, ezclassify.CtxSchemeSuggestForCategoriesRequest{
		Assertions:  []ezclassify.Assertion{{ConceptID: "L_H", CategoryID: "level", Confidence: ezclassify.AssertionConfidenceCertain, Source: "stated"}},
		CategoryIds: []string{"level"},
	})
	if err != nil {
		t.Fatalf("SuggestForCategories: %v", err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/suggest-for-categories")
	level, ok := byCategory.Result["level"]
	if !ok || len(level.Suggestions) != 1 || !level.Suggestions[0].Selected {
		t.Fatalf("SuggestForCategories = %+v, want the selected level option under its category", byCategory.Result)
	}

	agg, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: ptr([]string{"L", "H"})})
	if err != nil || agg.Result.Rendered != "H" {
		t.Fatalf("Aggregate: %+v err=%v", agg, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/aggregate")
	if len(agg.Result.Notices) != 1 || agg.Result.Notices[0].ConceptID != "H_CTRL" {
		t.Fatalf("Aggregate notices = %+v, want the one notice off the result", agg.Result.Notices)
	}

	scan, err := cx.Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{Text: "doc H"})
	if err != nil ||
		len(scan.Result.Results) != 1 ||
		scan.Result.Results[0].Marking != "H" ||
		string(scan.Result.Limits.Status) != "complete" {
		t.Fatalf("Scan: %+v err=%v", scan, err)
	}
	region := scan.Result.Results[0]
	if region.Type != "marking/banner" || region.Role != ezclassify.ProcessingRoleWhole ||
		region.Result.Recognition == nil || region.Result.Recognition.InputFormat == nil || *region.Result.Recognition.InputFormat != "marking/banner" {
		t.Fatalf("Scan region = %+v, want the banner view it was read as, written in it", region)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/scan")

	rendered, err := cx.Scheme.Render(ctx, ezclassify.CtxSchemeRenderRequest{Assertions: []ezclassify.Assertion{}, OutputFormat: ptr("banner"), Preservation: ptr(ezclassify.RenderPreservationFull)})
	if err != nil || rendered.Result.Result != "LEVEL-HIGH" {
		t.Fatalf("Render: %+v err=%v", rendered, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/render")
	if !strings.Contains(last().body, `"preservation":"full"`) || !strings.Contains(last().body, `"outputFormat":"banner"`) {
		t.Fatalf("Render request = %s, want format and full preservation", last().body)
	}

	norm, err := cx.Scheme.Normalize(ctx, ezclassify.CtxSchemeNormalizeRequest{Input: "high"})
	if err != nil || norm.Result.Result != "LEVEL-HIGH" {
		t.Fatalf("Normalize: %+v err=%v", norm, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/normalize")

	cov, err := cx.Scheme.CheckCoverage(ctx, ezclassify.CtxSchemeCheckCoverageRequest{Children: []string{"HIGH"}, Parent: "HIGH//H_CTRL"})
	if err != nil || cov.Result.Covered || len(cov.Result.MissingLabels) != 1 || cov.Result.Excess == nil || cov.Result.ExcessLabels == nil || len(cov.Result.ExcessLabels) != 0 {
		t.Fatalf("CheckCoverage: %+v err=%v", cov, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/check-coverage")

	desc, err := cx.Scheme.DescribeConcept(ctx, "L_H", "level", nil)
	if err != nil || desc.Result.Name != "High" {
		t.Fatalf("DescribeConcept: %+v err=%v", desc, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/describe-concept")
	if q := last().query; !strings.Contains(q, "conceptId=L_H") || !strings.Contains(q, "categoryId=level") {
		t.Fatalf("DescribeConcept query = %q", q)
	}

	diff, err := cx.Scheme.Diff(ctx, ezclassify.CtxSchemeDiffRequest{A: "A", B: "B"})
	if err != nil || len(diff.Result.Added) != 1 {
		t.Fatalf("Diff: %+v err=%v", diff, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/diff")

	filt, err := cx.Scheme.FilterConcepts(ctx, ezclassify.CtxSchemeFilterConceptsRequest{Query: "hi"})
	if err != nil || len(filt.Result.VisibleIds) != 2 || filt.Result.Total.Total != 9 {
		t.Fatalf("FilterConcepts: %+v err=%v", filt, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/filter-concepts")

	notices, err := cx.Scheme.Notices(ctx, ezclassify.CtxSchemeNoticesRequest{Assertions: []ezclassify.Assertion{}})
	if err != nil || len(notices.Result) != 1 || !notices.Result[0].Required {
		t.Fatalf("Notices: %+v err=%v", notices, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/notices")
	// The standalone call stays for assertions a caller assembled itself. For a
	// marking it already processed, the notices are on the process result above.

	view, err := cx.Scheme.View(ctx)
	if err != nil || view.Result.Name != "Scheme" {
		t.Fatalf("View: %+v err=%v", view, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/scheme/view")
}

func TestGuideOps(t *testing.T) {
	ctx := context.Background()
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/guide/scan":          `[{"rule":{"id":"r1","name":"Rule"},"matchedKeywords":["secret"],"spans":[],"guideId":"g1"}]`,
		"/guide/evaluate":      `{"aggregatedMarking":"H","selectedRules":[],"activeComputedRules":[],"effectiveMarkings":["H"],"breakdown":[],"contributingBreakdown":[],"guideId":"g1"}`,
		"/guide/filter-rules":  `{"visibleIds":["r1"],"categoryCounts":{},"total":{"visible":1,"total":3}}`,
		"/guide/describe-rule": `{"ruleId":"r1","guideId":"g1","name":"Rule","outcomeText":"marks H","isComputed":false,"dependsOn":[],"overrides":[],"lines":[]}`,
		"/guide/view":          `{"id":"g1","name":"Guide","description":"","categories":[],"uncategorized":[]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")

	scan, err := cx.Guide.Scan(ctx, ezclassify.CtxGuideScanRequest{Text: "secret"})
	if err != nil || len(scan.Result) != 1 || scan.Result[0].Rule.ID != "r1" {
		t.Fatalf("Guide.Scan: %+v err=%v", scan, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/guide/scan")
	assertNoSchemeID(t, last())

	guideID := "g1"
	ev, err := cx.Guide.Evaluate(ctx, ezclassify.CtxGuideEvaluateRequest{GuideID: &guideID, SelectedRuleIds: []string{"r1"}})
	if err != nil || ev.Result.AggregatedMarking != "H" {
		t.Fatalf("Guide.Evaluate: %+v err=%v", ev, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/guide/evaluate")
	assertNoSchemeID(t, last())
	allGuides := true
	ev, err = cx.Guide.Evaluate(ctx, ezclassify.CtxGuideEvaluateRequest{AllGuides: &allGuides, SelectedRuleIds: []string{"r2"}})
	if err != nil || ev.Result.AggregatedMarking != "H" {
		t.Fatalf("Guide.Evaluate all guides: %+v err=%v", ev, err)
	}
	if got := last().body; !strings.Contains(got, `"allGuides":true`) || !strings.Contains(got, `"selectedRuleIds":["r2"]`) || strings.Contains(got, `"guideId"`) {
		t.Fatalf("aggregate guide request body = %s", got)
	}

	fr, err := cx.Guide.FilterRules(ctx, ezclassify.CtxGuideFilterRulesRequest{Query: "sec"})
	if err != nil || len(fr.Result.VisibleIds) != 1 {
		t.Fatalf("Guide.FilterRules: %+v err=%v", fr, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/guide/filter-rules")

	dr, err := cx.Guide.DescribeRule(ctx, ezclassify.CtxGuideDescribeRuleRequest{RuleID: "r1"})
	if err != nil || dr.Result.OutcomeText != "marks H" {
		t.Fatalf("Guide.DescribeRule: %+v err=%v", dr, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/guide/describe-rule")

	gv, err := cx.Guide.View(ctx, ptr("g1"))
	if err != nil || gv.Result.Name != "Guide" {
		t.Fatalf("Guide.View: %+v err=%v", gv, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/guide/view")
}

func TestAccessOps(t *testing.T) {
	ctx := context.Background()
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/access/check":                 `{"granted":true,"rules":[{"id":"r1","name":"Clr","granted":true,"skipped":false,"ruleIndex":0}]}`,
		"/access/describe-requirements": `{"marking":"TS//NF","accessId":"ap","schemeId":"s","appliedRuleIds":["r1"],"rules":[],"combinedRequirementText":"US only","combinedAttributesNeeded":["nationality"],"lines":[]}`,
		"/access/describe-rule":         `{"ruleId":"r1","accessId":"ap","schemeId":"s","name":"NF","requirementText":"US only","attributesNeeded":["nationality"],"dependsOn":[],"overrides":[],"lines":[]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")

	dec, err := cx.Access.Check(ctx, ezclassify.CtxAccessCheckRequest{
		UserAttributes: map[string]interface{}{"clearance": "HIGH"},
		Assertions:     &[]ezclassify.Assertion{{ConceptID: "L_H", CategoryID: "level", Confidence: "certain", Source: "parse"}},
	})
	if err != nil || !dec.Result.Granted || len(dec.Result.Rules) != 1 {
		t.Fatalf("Access.Check: %+v err=%v", dec, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/access/check")
	assertNoSchemeID(t, last())
	if !strings.Contains(last().body, "userAttributes") {
		t.Fatalf("Access.Check body missing userAttributes: %s", last().body)
	}

	req, err := cx.Access.DescribeRequirements(ctx, ezclassify.CtxAccessDescribeRequirementsRequest{Marking: ptr("TS//NF")})
	if err != nil || len(req.Result.CombinedAttributesNeeded) != 1 {
		t.Fatalf("Access.DescribeRequirements: %+v err=%v", req, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/access/describe-requirements")

	rule, err := cx.Access.DescribeRule(ctx, ezclassify.CtxAccessDescribeRuleRequest{RuleID: "r1"})
	if err != nil || rule.Result.Name != "NF" {
		t.Fatalf("Access.DescribeRule: %+v err=%v", rule, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/access/describe-rule")
}

func TestDocumentAudit(t *testing.T) {
	ctx := context.Background()
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/document/audit": `{"whole":"TS//NF","wholes":[],"parts":[],"unmarkedSpans":[],"proceduralFlags":["unmarked-content"],"overallVerdict":"mixed"}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")

	rep, err := cx.Document.Audit(ctx, ezclassify.CtxDocumentAuditRequest{Text: "a doc"})
	if err != nil || rep.Result.Whole == nil || *rep.Result.Whole != "TS//NF" || string(rep.Result.OverallVerdict) != "mixed" {
		t.Fatalf("Document.Audit: %+v err=%v", rep, err)
	}
	assertPath(t, last(), "/v1/ctx/c/1.0.0/document/audit")
	if !strings.Contains(last().body, `"text":"a doc"`) {
		t.Fatalf("audit body = %s", last().body)
	}
}

func TestDocumentErrorsPreserveStatusAndStableDetail(t *testing.T) {
	tests := []struct {
		name   string
		status int
		call   func(*ezclassify.EzClassifyContext) error
		detail string
	}{
		{"audit stale", http.StatusConflict, func(cx *ezclassify.EzClassifyContext) error {
			_, err := cx.Document.Audit(context.Background(), ezclassify.CtxDocumentAuditRequest{Text: "document"})
			return err
		}, "stale-guide-determination"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(test.status)
				_, _ = io.WriteString(w, `{"error":"request failed","code":"INVALID_INPUT","detail":"`+test.detail+`"}`)
			}))
			t.Cleanup(srv.Close)
			err := test.call(ezclassify.Connect(srv.URL, "ctx", "1.0.0"))
			var apiErr *ezclassify.APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("error = %T(%v), want *APIError", err, err)
			}
			if apiErr.StatusCode != test.status || apiErr.Response == nil || apiErr.Response.Detail == nil || *apiErr.Response.Detail != test.detail {
				t.Fatalf("typed error = %+v", apiErr)
			}
		})
	}
}

func TestTextOperationsRejectInvalidUTF8BeforeSerialization(t *testing.T) {
	requests := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{}`)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "ctx", "1.0.0")
	invalid := string([]byte{0xff})

	tests := []struct {
		name string
		call func() error
		op   string
	}{
		{"audit", func() error {
			_, err := cx.Document.Audit(context.Background(), ezclassify.CtxDocumentAuditRequest{Text: invalid})
			return err
		}, "document.audit"},
		{"scheme scan", func() error {
			_, err := cx.Scheme.Scan(context.Background(), ezclassify.CtxSchemeScanRequest{Text: invalid})
			return err
		}, "scheme.scan"},
		{"guide scan", func() error {
			_, err := cx.Guide.Scan(context.Background(), ezclassify.CtxGuideScanRequest{Text: invalid})
			return err
		}, "guide.scan"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.call()
			var inputErr *ezclassify.InvalidInputError
			if !errors.As(err, &inputErr) || inputErr.Op != test.op || inputErr.Detail != "invalid-unicode" {
				t.Fatalf("error = %T(%v), want typed invalid-unicode for %s", err, err, test.op)
			}
			if strings.Contains(err.Error(), invalid) {
				t.Fatal("invalid input error echoed document content")
			}
		})
	}
	if requests != 0 {
		t.Fatalf("server received %d requests, want 0", requests)
	}
}

func TestCtxOpPropagatesServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = io.WriteString(w, `{"error":"context not found"}`)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "missing", "1.0.0")
	if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"}); err == nil {
		t.Fatal("expected error on 404, got nil")
	}
}

func TestContextProcessRoute(t *testing.T) {
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextSchemeProcessResult](`{"resolution":{"contextId":"tlp","requestedVersion":"latest","resolvedVersion":"1.0.0","engineVersion":"1.0.0"},"operationLimitations":[],"result":{"valid":true,"rendered":"H","inputWasValid":true,"assertions":[],"errors":[],"warnings":[],"appliedFixes":[],"proposedFixes":[],"notices":[],"assessment":{"emitted":[],"omitted":[],"collapsed":[],"conflicts":[],"unrepresentable":[]},"renders":{}}}`))
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "tlp", "latest")
	resp, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("context Scheme.Process: %v", err)
	}
	if !resp.Result.Valid {
		t.Fatalf("context Scheme.Process result = %+v", resp.Result)
	}
	if gotPath != "/v1/ctx/tlp/latest/scheme/process" {
		t.Fatalf("path = %q", gotPath)
	}
}

// An unset option must leave its field off the wire so the server applies its
// default, and a set one must reach it.
func TestScanSendsSetOptionsAndOmitsUnsetOnes(t *testing.T) {
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/scheme/scan": `{"results":[],"limits":{"status":"complete","regionsFound":0}}`,
	})
	cx := ezclassify.Connect(srv.URL, "tlp", "1.0.0")

	_, _ = cx.Scheme.Scan(context.Background(), ezclassify.CtxSchemeScanRequest{Text: "H", OutputFormat: ptr("banner")})
	if got, want := last().body, `{"outputFormat":"banner","text":"H"}`; got != want {
		t.Errorf("body with OutputFormat = %s, want %s", got, want)
	}

	_, _ = cx.Scheme.Scan(context.Background(), ezclassify.CtxSchemeScanRequest{Text: "H"})
	if got, want := last().body, `{"text":"H"}`; got != want {
		t.Errorf("body with no options = %s, want %s", got, want)
	}
}

// A process call names the outputs it wants rendered and the categories it
// wants picker options for. Neither is sent when the caller names none, and
// the answer carries no renders and no suggestions then, so an author asking
// for neither pays for neither.
func TestProcessNamesOutputsAndCategoriesOrOmitsThem(t *testing.T) {
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[],"renders":{}}`,
	})
	cx := ezclassify.Connect(srv.URL, "tlp", "1.0.0")

	plain, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if got, want := last().body, `{"input":"H"}`; got != want {
		t.Errorf("body with no outputs or categories = %s, want %s", got, want)
	}
	if len(plain.Result.Renders) != 0 || plain.Result.SuggestionsFor != nil {
		t.Errorf("result = renders %v, suggestionsFor %v, want neither", plain.Result.Renders, plain.Result.SuggestionsFor)
	}

	_, _ = cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{
		Input:      "H",
		Assertions: &[]ezclassify.Assertion{{ConceptID: "L_H", CategoryID: "level", Confidence: ezclassify.AssertionConfidenceCertain, Source: "stated"}},
		Outputs:    &[]string{"marking/portion"},
		SuggestFor: &[]string{"level", "audience"},
	})
	want := `{"assertions":[{"categoryId":"level","conceptId":"L_H","confidence":"certain","source":"stated"}],"input":"H","outputs":["marking/portion"],"suggestFor":["level","audience"]}`
	if got := last().body; got != want {
		t.Errorf("body with typed text, chosen facts, outputs and categories = %s, want %s", got, want)
	}
}

// Validate takes the marking as text or as stated facts. Each goes out alone,
// so the server sees exactly what the caller stated and nothing the client
// filled in.
func TestValidateSendsTextOrFactsAsGiven(t *testing.T) {
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/scheme/validate": `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`,
	})
	cx := ezclassify.Connect(srv.URL, "tlp", "1.0.0")

	_, _ = cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{
		Input:   ptr("(H)"),
		Context: &ezclassify.ProcessingContext{Role: ptr(ezclassify.ProcessingRolePart)},
	})
	if got, want := last().body, `{"context":{"role":"part"},"input":"(H)"}`; got != want {
		t.Errorf("validate body for text = %s, want %s", got, want)
	}

	_, _ = cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{
		Assertions: &[]ezclassify.Assertion{{ConceptID: "L_H", CategoryID: "level", Confidence: ezclassify.AssertionConfidenceCertain, Source: "stated"}},
	})
	if got, want := last().body, `{"assertions":[{"categoryId":"level","conceptId":"L_H","confidence":"certain","source":"stated"}]}`; got != want {
		t.Errorf("validate body for facts = %s, want %s", got, want)
	}
}

// TestProcessAndValidateSendPartsOrOmitThem pins that Parts reaches the wire
// as a plain "parts" array when set, and that a nil or empty Parts leaves the
// key off the body entirely rather than sending "parts":null or "parts":[].
func TestProcessAndValidateSendPartsOrOmitThem(t *testing.T) {
	srv, last := ctxServer(t, "1.0.0", map[string]string{
		"/scheme/process":  `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`,
		"/scheme/validate": `{"valid":true,"rendered":"H","inputWasValid":true}`,
	})
	cx := ezclassify.Connect(srv.URL, "tlp", "1.0.0")

	_, _ = cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{
		Input: "H", Parts: &[]string{"H//A", "H//B"},
	})
	if got, want := last().body, `{"input":"H","parts":["H//A","H//B"]}`; got != want {
		t.Errorf("process body with Parts set = %s, want %s", got, want)
	}

	_, _ = cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if got, want := last().body, `{"input":"H"}`; got != want {
		t.Errorf("process body with nil Parts = %s, want %s", got, want)
	}

	_, _ = cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{
		Input: ptr("H"), Parts: &[]string{"H//A"},
	})
	if got, want := last().body, `{"input":"H","parts":["H//A"]}`; got != want {
		t.Errorf("validate body with Parts set = %s, want %s", got, want)
	}

	_, _ = cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{Input: ptr("H")})
	if got, want := last().body, `{"input":"H"}`; got != want {
		t.Errorf("validate body with nil Parts = %s, want %s", got, want)
	}
}

func TestContextLabelPresetsWithResponse(t *testing.T) {
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `[{"id":"state-assessment-staff","label":"State Assessment Staff"}]`)
	}))
	t.Cleanup(srv.Close)
	c := ezclassify.New(srv.URL)
	resp, err := c.CtxSchemeLabelPresetsWithResponse(context.Background(), "cedar-valley-protection", "latest")
	if err != nil {
		t.Fatalf("CtxSchemeLabelPresetsWithResponse: %v", err)
	}
	if resp.JSON200 == nil || len(*resp.JSON200) != 1 || (*resp.JSON200)[0].ID != "state-assessment-staff" {
		t.Fatalf("CtxSchemeLabelPresetsWithResponse JSON200 = %+v", resp.JSON200)
	}
	if gotPath != "/v1/ctx/cedar-valley-protection/latest/scheme/label-presets" {
		t.Fatalf("path = %q", gotPath)
	}
}

// TestContextSchemeScopesWithResponse covers the Context-bound scopes wire
// surface. The response is an ordered array of scope
// entries, not a map, because a selector offers scopes in declaration order
// and a JSON object carries no order. This pins that shape and the ordering.
func TestContextSchemeScopesWithResponse(t *testing.T) {
	var gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `[{"id":"public","label":"Public"},{"id":"internal","label":"Internal"}]`)
	}))
	t.Cleanup(srv.Close)
	c := ezclassify.New(srv.URL)
	resp, err := c.CtxSchemeScopesWithResponse(context.Background(), "cedar-valley-protection", "latest", nil)
	if err != nil {
		t.Fatalf("CtxSchemeScopesWithResponse: %v", err)
	}
	if resp.JSON200 == nil {
		t.Fatalf("CtxSchemeScopesWithResponse JSON200 = nil")
	}
	scopes := *resp.JSON200
	if len(scopes) != 2 {
		t.Fatalf("scopes length = %d, want 2", len(scopes))
	}
	if scopes[0].ID != "public" || scopes[0].Label != "Public" {
		t.Fatalf("scopes[0] = %+v, want id=public label=Public", scopes[0])
	}
	if scopes[1].ID != "internal" {
		t.Fatalf("scopes[1].ID = %q, want internal", scopes[1].ID)
	}
	if gotPath != "/v1/ctx/cedar-valley-protection/latest/scheme/scopes" {
		t.Fatalf("path = %q", gotPath)
	}
}

// TestRichResultsCarryNotices pins the collected notices as a first-class
// member of the process and aggregate results. A caller that already processed
// a marking reads them off the result it has; the standalone notices call is
// for assertions it assembled some other way. The field is always sent, so the
// empty case has to decode as an empty list rather than a missing one.
func TestRichResultsCarryNotices(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"H","inputWasValid":true,"assertions":[],` +
			`"notices":[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true},` +
			`{"conceptId":"H_ADV","conceptName":"Advisory","type":"info","text":"advisory","required":false}]}`,
		"/scheme/validate": `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`,
		"/scheme/aggregate": `{"valid":true,"rendered":"H","assertions":[],` +
			`"notices":[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true}]}`,
		"/scheme/scan": `{"results":[{"marking":"H","start":0,"end":1,"type":"marking","wrapped":false,` +
			`"result":{"valid":true,"rendered":"H","inputWasValid":true,` +
			`"notices":[{"conceptId":"H_CTRL","conceptName":"Handling Controlled","type":"handling","text":"controlled","required":true}]}}],` +
			`"limits":{"status":"complete","regionsFound":1}}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	proc, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if len(proc.Result.Notices) != 2 {
		t.Fatalf("len(Process notices) = %d, want 2", len(proc.Result.Notices))
	}
	if !proc.Result.Notices[0].Required || proc.Result.Notices[1].Required {
		t.Errorf("required flags = %v, %v, want true then false",
			proc.Result.Notices[0].Required, proc.Result.Notices[1].Required)
	}
	if got := proc.Result.Notices[0].Type; got != ezclassify.ConceptNoticeTypeHandling {
		t.Errorf("Notices[0].Type = %q, want the handling member", got)
	}

	val, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: ptr("H")})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	if val.Result.Notices == nil || len(val.Result.Notices) != 0 {
		t.Errorf("Validate notices = %+v, want an empty list when the marking carries none", val.Result.Notices)
	}

	agg, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: ptr([]string{"L", "H"})})
	if err != nil {
		t.Fatalf("Aggregate: %v", err)
	}
	if len(agg.Result.Notices) != 1 || agg.Result.Notices[0].ConceptID != "H_CTRL" {
		t.Errorf("Aggregate notices = %+v, want the one collected notice", agg.Result.Notices)
	}

	// A scan match carries what validate answers for it, so its notices come along too.
	scan, err := cx.Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{Text: "H"})
	if err != nil {
		t.Fatalf("Scan: %v", err)
	}
	if len(scan.Result.Results) != 1 ||
		len(scan.Result.Results[0].Result.Notices) != 1 {
		t.Errorf("Scan match notices = %+v, want the one notice off the embedded result", scan.Result.Results)
	}
}

func assertPath(t *testing.T, r recordedReq, want string) {
	t.Helper()
	if r.path != want {
		t.Fatalf("path = %q, want %q", r.path, want)
	}
}

func assertNoSchemeID(t *testing.T, r recordedReq) {
	t.Helper()
	if strings.Contains(r.body, "schemeId") || strings.Contains(r.body, "schemeID") {
		t.Fatalf("ctx body should not carry a scheme id, got: %s", r.body)
	}
}

// TestProcessCarriesWhatWasReadBeforeRepair pins that the reading a server
// sends when repair changed the marking reaches the caller. A fix for a rule
// repair enforced is applied to that reading rather than to Assertions, so a
// decode that dropped it would apply the fix to the wrong set. Absent on the
// wire, the field is nil.
func TestProcessCarriesWhatWasReadBeforeRepair(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	assertion := func(id string) string {
		return `{"conceptId":"` + id + `","categoryId":"level","confidence":"certain","source":"user"}`
	}
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":false,"rendered":"SENSITIVE","inputWasValid":false,` +
			`"assertions":[` + assertion("SENSITIVE") + `],` +
			`"assertionsBeforeRepair":[` + assertion("OPEN") + `,` + assertion("SENSITIVE") + `]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	processed, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "OPEN/SENSITIVE"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	read := processed.Result.AssertionsBeforeRepair
	if read == nil || len(*read) != 2 || (*read)[0].ConceptID != "OPEN" || (*read)[1].ConceptID != "SENSITIVE" {
		t.Fatalf("AssertionsBeforeRepair = %+v, want OPEN then SENSITIVE", read)
	}

	unrepairedServer, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"OPEN","inputWasValid":true,` +
			`"assertions":[` + assertion("OPEN") + `]}`,
	})
	unrepaired, err := ezclassify.Connect(unrepairedServer.URL, "c", ver).Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "OPEN"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	if unrepaired.Result.AssertionsBeforeRepair != nil {
		t.Errorf("AssertionsBeforeRepair = %+v, want nil when the wire omits it", unrepaired.Result.AssertionsBeforeRepair)
	}
}

// TestFixDeserializesKeepsAddsRemovesAndRendered pins the Fix shape
// (proposedFixes only): every field is required on the wire, and a fix that
// keeps two concepts, adds another, and removes a fourth decodes each list
// into its own field alongside the rendered marking the fix would produce.
func TestFixDeserializesKeepsAddsRemovesAndRendered(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":false,"rendered":"H","inputWasValid":false,` +
			`"proposedFixes":[` +
			`{"message":"add required tag","rendered":"H//A/TAG",` +
			`"keeps":[{"categoryId":"level","conceptId":"H"},{"categoryId":"scope","conceptId":"A"}],` +
			`"adds":[{"categoryId":"tag","conceptId":"TAG"}],` +
			`"removes":[{"categoryId":"level","conceptId":"OLD"}]}` +
			`]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	processed, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	proposed := processed.Result.ProposedFixes
	if len(proposed) != 1 {
		t.Fatalf("ProposedFixes = %+v, want 1 fix", proposed)
	}
	fix := proposed[0]
	if fix.Message != "add required tag" || fix.Rendered != "H//A/TAG" {
		t.Errorf("Fix message/rendered = %q/%q, want %q/%q", fix.Message, fix.Rendered, "add required tag", "H//A/TAG")
	}
	if len(fix.Keeps) != 2 || fix.Keeps[0].ConceptID != "H" || fix.Keeps[1].ConceptID != "A" {
		t.Errorf("Fix.Keeps = %+v, want [level H, scope A]", fix.Keeps)
	}
	if len(fix.Adds) != 1 || fix.Adds[0].ConceptID != "TAG" {
		t.Errorf("Fix.Adds = %+v, want [tag TAG]", fix.Adds)
	}
	if len(fix.Removes) != 1 || fix.Removes[0].ConceptID != "OLD" {
		t.Errorf("Fix.Removes = %+v, want [level OLD]", fix.Removes)
	}
}

// TestChangeDeserializesEachAction pins that appliedFixes decodes as Change,
// with each of the four FixAction values carrying the fields its own doc
// promises: 'add' only Add, 'remove' only Remove, 'replace' both, and
// 'reorder' neither (the canonical form is in ProcessResult.rendered).
func TestChangeDeserializesEachAction(t *testing.T) {
	ctx := context.Background()
	const ver = "1.0.0"
	srv, _ := ctxServer(t, ver, map[string]string{
		"/scheme/process": `{"valid":true,"rendered":"H","inputWasValid":true,` +
			`"appliedFixes":[` +
			`{"ruleId":"r-add","action":"add","message":"add level",` +
			`"add":[{"categoryId":"level","conceptId":"H"}]},` +
			`{"ruleId":"r-remove","action":"remove","message":"remove superseded",` +
			`"remove":[{"categoryId":"level","conceptId":"OLD"}]},` +
			`{"ruleId":"r-replace","action":"replace","message":"replace level",` +
			`"add":[{"categoryId":"level","conceptId":"H"}],` +
			`"remove":[{"categoryId":"level","conceptId":"OLD"}]},` +
			`{"ruleId":"r-reorder","action":"reorder","message":"canonical order"}` +
			`]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", ver)

	processed, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
	if err != nil {
		t.Fatalf("Process: %v", err)
	}
	changes := processed.Result.AppliedFixes
	if len(changes) != 4 {
		t.Fatalf("AppliedFixes = %+v, want 4 changes", changes)
	}

	add := changes[0]
	if add.Action != ezclassify.FixActionAdd || add.Add == nil || len(*add.Add) != 1 || add.Remove != nil {
		t.Errorf("add change = %+v, want only Add set", add)
	}

	remove := changes[1]
	if remove.Action != ezclassify.FixActionRemove || remove.Remove == nil || len(*remove.Remove) != 1 || remove.Add != nil {
		t.Errorf("remove change = %+v, want only Remove set", remove)
	}

	replace := changes[2]
	if replace.Action != ezclassify.FixActionReplace ||
		replace.Add == nil || len(*replace.Add) != 1 ||
		replace.Remove == nil || len(*replace.Remove) != 1 {
		t.Errorf("replace change = %+v, want both Add and Remove set", replace)
	}

	reorder := changes[3]
	if reorder.Action != ezclassify.FixActionReorder || reorder.Add != nil || reorder.Remove != nil {
		t.Errorf("reorder change = %+v, want neither Add nor Remove set", reorder)
	}
}

// TestAuditValidationCarriesFactsRecognitionAndBothVerdicts pins that each
// authored marking an audit judges carries the facts it states, how far
// recognition got with it, and both verdicts: Valid, whether it stands once
// settled, and InputWasValid, whether it was right as the author wrote it. A
// caller that rolls markings up or checks access works from these instead of
// reading each marking again. AppliedFixes explains a false InputWasValid and
// is always sent, so a marking that needed no change decodes an empty list.
func TestAuditValidationCarriesFactsRecognitionAndBothVerdicts(t *testing.T) {
	ctx := context.Background()
	assertion := `{"conceptId":"L_H","categoryId":"level","confidence":"certain","source":"parse"}`
	srv, _ := ctxServer(t, "1.0.0", map[string]string{
		"/document/audit": `{"whole":"h","wholes":[],"unmarkedSpans":[],"proceduralFlags":[],"overallVerdict":"pass",` +
			`"wholeValidation":{"assertions":[` + assertion + `],` +
			`"recognition":{"status":"complete","inputFormat":"marking/banner","role":"whole","surface":"marking","unmatched":[]},` +
			`"valid":true,"inputWasValid":false,"errors":[],"proposedFixes":[],"rendered":"H",` +
			`"appliedFixes":[{"ruleId":"_recognition_case","action":"reorder","message":"Read h as H (case)."}]},` +
			`"parts":[{"marking":"(H)","start":4,"end":7,"offsetEncoding":"utf-8-bytes","scopeId":"part-1",` +
			`"validation":{"assertions":[` + assertion + `],` +
			`"recognition":{"status":"complete","inputFormat":"marking/portion","role":"part","surface":"marking","unmatched":[]},` +
			`"valid":true,"inputWasValid":true,"errors":[],"proposedFixes":[],"appliedFixes":[],"rendered":"H"}}]}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")

	rep, err := cx.Document.Audit(ctx, ezclassify.CtxDocumentAuditRequest{Text: "h\n\n(H) body"})
	if err != nil {
		t.Fatalf("Document.Audit: %v", err)
	}
	whole := rep.Result.WholeValidation
	if whole == nil {
		t.Fatalf("WholeValidation = nil, want a validation")
	}
	if !whole.Valid || whole.InputWasValid {
		t.Errorf("whole verdicts = valid %v, inputWasValid %v, want a marking that stands but was not right as written", whole.Valid, whole.InputWasValid)
	}
	if len(whole.Assertions) != 1 || whole.Assertions[0].ConceptID != "L_H" {
		t.Errorf("whole assertions = %+v, want the one level fact", whole.Assertions)
	}
	if whole.Recognition.InputFormat == nil || *whole.Recognition.InputFormat != "marking/banner" ||
		whole.Recognition.Role == nil || *whole.Recognition.Role != ezclassify.ProcessingRoleWhole {
		t.Errorf("whole recognition = %+v, want the banner view of a whole", whole.Recognition)
	}
	if len(whole.AppliedFixes) != 1 || whole.AppliedFixes[0].Action != ezclassify.FixActionReorder {
		t.Errorf("whole appliedFixes = %+v, want the change that explains inputWasValid false", whole.AppliedFixes)
	}

	if len(rep.Result.Parts) != 1 {
		t.Fatalf("parts = %+v, want one", rep.Result.Parts)
	}
	part := rep.Result.Parts[0].Validation
	if !part.Valid || !part.InputWasValid || part.Recognition.Role == nil || *part.Recognition.Role != ezclassify.ProcessingRolePart {
		t.Errorf("part validation = %+v, want a part right as written", part)
	}
	if part.AppliedFixes == nil || len(part.AppliedFixes) != 0 {
		t.Errorf("part appliedFixes = %+v, want an empty list", part.AppliedFixes)
	}
}

// TestAssessmentTellsAnUnwritableFactFromARefusedOutput pins how a caller
// reads Assessment.Unrepresentable, whose entries are either one fact a
// format cannot write or a whole output refused before anything was written.
// Only a fact carries a concept id and only a refused output names an output,
// so each entry reads as exactly one of the two. This is the reading the
// README and the UnrepresentableValue doc teach.
func TestAssessmentTellsAnUnwritableFactFromARefusedOutput(t *testing.T) {
	srv, _ := ctxServer(t, "1.0.0", map[string]string{
		"/scheme/validate": `{"valid":false,"rendered":"","inputWasValid":false,"notices":[],` +
			`"assessment":{"emitted":[],"omitted":[],"collapsed":[],"conflicts":[],"unrepresentable":[` +
			`{"conceptId":"ORIG","reason":"value-lossy","detail":"written without its captured text"},` +
			`{"output":"marking/portion","reason":"unresolved-context"}]}}`,
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	read, err := cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{Input: ptr("H")})
	if err != nil {
		t.Fatalf("Validate: %v", err)
	}
	entries := read.Result.Assessment.Unrepresentable
	if len(entries) != 2 {
		t.Fatalf("unrepresentable = %d entries, want 2", len(entries))
	}

	fact, err := entries[0].AsUnrepresentableFact()
	if err != nil || fact.ConceptID != "ORIG" || fact.Reason != ezclassify.UnrepresentableFactReasonValueLossy || !fact.Reason.Valid() {
		t.Errorf("first entry as a fact = %+v, %v, want ORIG written lossily", fact, err)
	}
	if output, err := entries[0].AsUnrepresentableOutput(); err != nil || output.Output != "" {
		t.Errorf("first entry as an output = %+v, %v, want no output named", output, err)
	}

	output, err := entries[1].AsUnrepresentableOutput()
	if err != nil || output.Output != "marking/portion" || output.Reason != ezclassify.UnrepresentableOutputReasonUnresolvedContext || !output.Reason.Valid() {
		t.Errorf("second entry as an output = %+v, %v, want the portion output refused", output, err)
	}
	if fact, err := entries[1].AsUnrepresentableFact(); err != nil || fact.ConceptID != "" {
		t.Errorf("second entry as a fact = %+v, %v, want no concept named", fact, err)
	}
}

// Describe is the audit-trail read: the Context's description with each
// component's declared version beside the one that answered, crosswalks
// included. It always goes to the server and records the resolved version.
func TestDescribeReturnsCrosswalksWithDeclaredAndResolvedVersions(t *testing.T) {
	description := fixtureAs[ezclassify.ContextDescription](`{
		"id":"c","version":"1.0.4",
		"scheme":{"id":"s","version":"latest","resolvedVersion":"1.0.0"},
		"crosswalks":[{"id":"s-to-t","version":"latest","resolvedVersion":"2.1.0"}],
		"schemes":[{"id":"s","version":"1.0.0","selector":"s","roles":["primary","crosswalk-source"]}]
	}`)
	var calls atomic.Int32
	var gotPath, gotMethod string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		gotPath, gotMethod = r.URL.Path, r.Method
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.4")
		_, _ = io.WriteString(w, description)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "latest")

	for i := 0; i < 2; i++ {
		got, err := cx.Describe(context.Background())
		if err != nil {
			t.Fatalf("Describe: %v", err)
		}
		if got.ID != "c" || got.Version != "1.0.4" || got.Scheme.ResolvedVersion != "1.0.0" {
			t.Fatalf("description = %+v", got)
		}
		if got.Crosswalks == nil || len(*got.Crosswalks) != 1 {
			t.Fatalf("Crosswalks = %v, want the one declared crosswalk", got.Crosswalks)
		}
		if crosswalk := (*got.Crosswalks)[0]; crosswalk.ID != "s-to-t" || crosswalk.Version != "latest" || crosswalk.ResolvedVersion != "2.1.0" {
			t.Fatalf("crosswalk = %+v, want declared latest beside resolved 2.1.0", crosswalk)
		}
	}
	if gotMethod != http.MethodGet || gotPath != "/v1/ctx/c/latest" {
		t.Fatalf("Describe sent %s %s, want GET /v1/ctx/c/latest", gotMethod, gotPath)
	}
	if calls.Load() != 2 {
		t.Fatalf("Describe used %d wire requests for two calls, want every call to reach the server", calls.Load())
	}
	if cx.ResolvedVersion() != "1.0.4" {
		t.Fatalf("ResolvedVersion() = %q, want the version Describe observed", cx.ResolvedVersion())
	}
}

func TestDescribeOmitsCrosswalksWhenNoneAreDeclared(t *testing.T) {
	srv, _ := ctxServer(t, "1.0.0", map[string]string{
		"/v1/ctx/c/1.0.0": fixtureAs[ezclassify.ContextDescription](`{"id":"c","version":"1.0.0"}`),
	})
	got, err := ezclassify.Connect(srv.URL, "c", "1.0.0").Describe(context.Background())
	if err != nil || got.Crosswalks != nil {
		t.Fatalf("Describe = %+v, %v; want no Crosswalks", got, err)
	}
}

func TestDescribeFailuresCarryTheirOperation(t *testing.T) {
	missing := httptest.NewServer(stubStatus(http.StatusNotFound, `{"error":"context not found"}`))
	t.Cleanup(missing.Close)
	_, err := ezclassify.Connect(missing.URL, "gone", "1.0.0").Describe(context.Background())
	var notFound *ezclassify.ContextNotFoundError
	if !errors.As(err, &notFound) || notFound.APIError.Op != "describe context" {
		t.Fatalf("Describe on 404 = %T(%v), want *ContextNotFoundError for describe context", err, err)
	}

	partial := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"id":"c","version":"1.0.0"}`)
	}))
	t.Cleanup(partial.Close)
	_, err = ezclassify.Connect(partial.URL, "c", "1.0.0").Describe(context.Background())
	var malformed *ezclassify.MalformedResponseError
	if !errors.As(err, &malformed) || malformed.Op != "describe context" || malformed.Member != "/scheme" {
		t.Fatalf("Describe on a description without a scheme = %T(%v), want a *MalformedResponseError at /scheme", err, err)
	}
}
