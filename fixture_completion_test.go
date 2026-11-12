package ezclassify_test

import (
	"reflect"
	"strings"

	ez "github.com/ezclassify/sdk-go"
)

// contextResponseTypes maps a Context route's suffix to the response type the
// client decodes it into.
var contextResponseTypes = map[string]reflect.Type{
	"/scheme/process":                reflect.TypeOf(ez.ContextSchemeProcessResult{}),
	"/scheme/validate":               reflect.TypeOf(ez.ContextSchemeValidateResult{}),
	"/scheme/explain":                reflect.TypeOf(ez.ContextSchemeExplainResult{}),
	"/scheme/suggest":                reflect.TypeOf(ez.ContextSchemeSuggestResult{}),
	"/scheme/suggest-for-categories": reflect.TypeOf(ez.ContextSchemeSuggestForCategoriesResult{}),
	"/scheme/aggregate":              reflect.TypeOf(ez.ContextSchemeAggregateResult{}),
	"/scheme/scan":                   reflect.TypeOf(ez.ContextSchemeScanResult{}),
	"/scheme/render":                 reflect.TypeOf(ez.ContextSchemeRenderResult{}),
	"/scheme/check-coverage":         reflect.TypeOf(ez.ContextSchemeCoverageResult{}),
	"/scheme/describe-concept":       reflect.TypeOf(ez.ContextSchemeConceptDescriptionResult{}),
	"/scheme/diff":                   reflect.TypeOf(ez.ContextSchemeDiffResult{}),
	"/scheme/notices":                reflect.TypeOf(ez.ContextSchemeNoticesResult{}),
	"/scheme/view":                   reflect.TypeOf(ez.ContextSchemeViewResult{}),
	"/guide/scan":                    reflect.TypeOf(ez.ContextGuideScanResult{}),
	"/guide/evaluate":                reflect.TypeOf(ez.ContextGuideEvaluationResult{}),
	"/guide/view":                    reflect.TypeOf(ez.ContextGuideViewResult{}),
	"/access/check":                  reflect.TypeOf(ez.ContextAccessDecisionResult{}),
	"/access/describe-requirements":  reflect.TypeOf(ez.ContextAccessRequirementsResult{}),
	"/document/audit":                reflect.TypeOf(ez.ContextDocumentAuditResult{}),
	"/crosswalk/translate":           reflect.TypeOf(ez.ContextCrosswalkTranslationResult{}),
}

// contextFixture returns body, a stubbed Context response, with every member
// the response contract requires and the stub omits filled with its zero
// value. Stubs state only what a test examines; this keeps them valid
// responses without restating the rest.
func contextFixture(path, body string) string {
	for suffix, responseType := range contextResponseTypes {
		if strings.HasSuffix(path, suffix) {
			return string(ez.ZeroFillRequiredForTest(responseType, []byte(body)))
		}
	}
	return body
}

// fixtureAs is contextFixture for a stub that knows the response type it
// stands for.
func fixtureAs[T any](body string) string {
	return string(ez.ZeroFillRequiredForTest(reflect.TypeOf((*T)(nil)).Elem(), []byte(body)))
}

// processedFixture is a Process response whose result states only that the
// marking was valid.
var processedFixture = fixtureAs[ez.ContextSchemeProcessResult](`{"result":{"valid":true,"rendered":"H","inputWasValid":true}}`)
