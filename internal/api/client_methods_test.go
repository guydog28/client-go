package api

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func recordingStub(t *testing.T) (*httptest.Server, func() (string, string)) {
	t.Helper()
	var mu sync.Mutex
	var method, path string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		method, path = r.Method, r.URL.Path
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "null")
	}))
	t.Cleanup(srv.Close)
	return srv, func() (string, string) {
		mu.Lock()
		defer mu.Unlock()
		return method, path
	}
}

// TestEndpointRoutes pins the generated client's complete public route family.
// The canonical OpenAPI intentionally contains only infrastructure, Context
// discovery, and Context-bound operations; a flat /v1/scheme, /v1/registry,
// /v1/guide, /v1/access, or /v1/crosswalk route cannot appear here.
func TestEndpointRoutes(t *testing.T) {
	srv, last := recordingStub(t)
	c, err := NewClientWithResponses(srv.URL)
	if err != nil {
		t.Fatalf("NewClientWithResponses: %v", err)
	}
	ctx := context.Background()
	const ctxID, version = "test-ctx", "1.0"
	base := "/v1/ctx/" + ctxID + "/" + version

	cases := []struct {
		name, method, path string
		call               func() error
	}{
		{"Health", "GET", "/v1/health", func() error { _, e := c.GetHealthWithResponse(ctx); return e }},
		{"Ready", "GET", "/v1/ready", func() error { _, e := c.GetReadyWithResponse(ctx); return e }},
		{"CtxList", "GET", "/v1/ctx", func() error { _, e := c.CtxListWithResponse(ctx); return e }},
		{"CtxVersions", "GET", "/v1/ctx/" + ctxID, func() error { _, e := c.CtxVersionsWithResponse(ctx, ctxID); return e }},
		{"CtxDescribe", "GET", base, func() error { _, e := c.CtxDescribeWithResponse(ctx, ctxID, version); return e }},
		{"AccessRules", "GET", base + "/access/rules", func() error { _, e := c.CtxAccessRulesWithResponse(ctx, ctxID, version); return e }},
		{"SchemeSummary", "GET", base + "/scheme", func() error { _, e := c.CtxSchemeSummaryWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemeCategories", "GET", base + "/scheme/categories", func() error { _, e := c.CtxSchemeCategoriesWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemeConcepts", "GET", base + "/scheme/concepts", func() error { _, e := c.CtxSchemeConceptsWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemeFormats", "GET", base + "/scheme/formats", func() error { _, e := c.CtxSchemeFormatsWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemeGrammar", "GET", base + "/scheme/grammar", func() error { _, e := c.CtxSchemeGrammarWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemeScopes", "GET", base + "/scheme/scopes", func() error { _, e := c.CtxSchemeScopesWithResponse(ctx, ctxID, version, nil); return e }},
		{"SchemePresets", "GET", base + "/scheme/label-presets", func() error { _, e := c.CtxSchemeLabelPresetsWithResponse(ctx, ctxID, version); return e }},
		{"DescribeConcept", "GET", base + "/scheme/describe-concept", func() error {
			_, e := c.CtxSchemeDescribeConceptWithResponse(ctx, ctxID, version, &CtxSchemeDescribeConceptParams{ConceptID: "c", CategoryID: "cat"})
			return e
		}},
		{"AccessCheck", "POST", base + "/access/check", func() error {
			_, e := c.CtxAccessCheckWithResponse(ctx, ctxID, version, CtxAccessCheckJSONRequestBody{})
			return e
		}},
		{"AccessRequirements", "POST", base + "/access/describe-requirements", func() error {
			_, e := c.CtxAccessDescribeRequirementsWithResponse(ctx, ctxID, version, CtxAccessDescribeRequirementsJSONRequestBody{})
			return e
		}},
		{"AccessRule", "POST", base + "/access/describe-rule", func() error {
			_, e := c.CtxAccessDescribeRuleWithResponse(ctx, ctxID, version, CtxAccessDescribeRuleJSONRequestBody{})
			return e
		}},
		{"Crosswalk", "POST", base + "/crosswalk/translate", func() error {
			_, e := c.CtxCrosswalkTranslateWithResponse(ctx, ctxID, version, CtxCrosswalkTranslateJSONRequestBody{})
			return e
		}},
		{"DocumentAudit", "POST", base + "/document/audit", func() error {
			_, e := c.CtxDocumentAuditWithResponse(ctx, ctxID, version, CtxDocumentAuditJSONRequestBody{})
			return e
		}},
		{"GuideScan", "POST", base + "/guide/scan", func() error {
			_, e := c.CtxGuideScanWithResponse(ctx, ctxID, version, CtxGuideScanJSONRequestBody{})
			return e
		}},
		{"GuideEvaluate", "POST", base + "/guide/evaluate", func() error {
			_, e := c.CtxGuideEvaluateWithResponse(ctx, ctxID, version, CtxGuideEvaluateJSONRequestBody{})
			return e
		}},
		{"GuideDescribe", "POST", base + "/guide/describe-rule", func() error {
			_, e := c.CtxGuideDescribeRuleWithResponse(ctx, ctxID, version, CtxGuideDescribeRuleJSONRequestBody{})
			return e
		}},
		{"GuideFilter", "POST", base + "/guide/filter-rules", func() error {
			_, e := c.CtxGuideFilterRulesWithResponse(ctx, ctxID, version, CtxGuideFilterRulesJSONRequestBody{})
			return e
		}},
		{"GuideView", "POST", base + "/guide/view", func() error {
			_, e := c.CtxGuideViewWithResponse(ctx, ctxID, version, CtxGuideViewJSONRequestBody{})
			return e
		}},
		{"SchemeProcess", "POST", base + "/scheme/process", func() error {
			_, e := c.CtxSchemeProcessWithResponse(ctx, ctxID, version, CtxSchemeProcessJSONRequestBody{})
			return e
		}},
		{"SchemeValidate", "POST", base + "/scheme/validate", func() error {
			_, e := c.CtxSchemeValidateWithResponse(ctx, ctxID, version, CtxSchemeValidateJSONRequestBody{})
			return e
		}},
		{"SchemeNormalize", "POST", base + "/scheme/normalize", func() error {
			_, e := c.CtxSchemeNormalizeWithResponse(ctx, ctxID, version, CtxSchemeNormalizeJSONRequestBody{})
			return e
		}},
		{"SchemeRender", "POST", base + "/scheme/render", func() error {
			_, e := c.CtxSchemeRenderWithResponse(ctx, ctxID, version, CtxSchemeRenderJSONRequestBody{})
			return e
		}},
		{"SchemeDiff", "POST", base + "/scheme/diff", func() error {
			_, e := c.CtxSchemeDiffWithResponse(ctx, ctxID, version, CtxSchemeDiffJSONRequestBody{})
			return e
		}},
		{"SchemeView", "POST", base + "/scheme/view", func() error {
			_, e := c.CtxSchemeViewWithResponse(ctx, ctxID, version, CtxSchemeViewJSONRequestBody{})
			return e
		}},
		{"SchemeNotices", "POST", base + "/scheme/notices", func() error {
			_, e := c.CtxSchemeNoticesWithResponse(ctx, ctxID, version, CtxSchemeNoticesJSONRequestBody{})
			return e
		}},
		{"SchemeFilter", "POST", base + "/scheme/filter-concepts", func() error {
			_, e := c.CtxSchemeFilterConceptsWithResponse(ctx, ctxID, version, CtxSchemeFilterConceptsJSONRequestBody{})
			return e
		}},
		{"SchemeScan", "POST", base + "/scheme/scan", func() error {
			_, e := c.CtxSchemeScanWithResponse(ctx, ctxID, version, CtxSchemeScanJSONRequestBody{})
			return e
		}},
		{"SchemeAggregate", "POST", base + "/scheme/aggregate", func() error {
			_, e := c.CtxSchemeAggregateWithResponse(ctx, ctxID, version, CtxSchemeAggregateJSONRequestBody{})
			return e
		}},
		{"SchemeSuggest", "POST", base + "/scheme/suggest", func() error {
			_, e := c.CtxSchemeSuggestWithResponse(ctx, ctxID, version, CtxSchemeSuggestJSONRequestBody{})
			return e
		}},
		{"SchemeSuggestForCategories", "POST", base + "/scheme/suggest-for-categories", func() error {
			_, e := c.CtxSchemeSuggestForCategoriesWithResponse(ctx, ctxID, version, CtxSchemeSuggestForCategoriesJSONRequestBody{})
			return e
		}},
		{"SchemeExplain", "POST", base + "/scheme/explain", func() error {
			_, e := c.CtxSchemeExplainWithResponse(ctx, ctxID, version, CtxSchemeExplainJSONRequestBody{})
			return e
		}},
		{"SchemeCoverage", "POST", base + "/scheme/check-coverage", func() error {
			_, e := c.CtxSchemeCheckCoverageWithResponse(ctx, ctxID, version, CtxSchemeCheckCoverageJSONRequestBody{})
			return e
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(); err != nil {
				t.Fatalf("call: %v", err)
			}
			method, path := last()
			if method != tc.method || path != tc.path {
				t.Fatalf("request = %s %s, want %s %s", method, path, tc.method, tc.path)
			}
		})
	}
}
