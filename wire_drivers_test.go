package ezclassify_test

import (
	"context"
	"encoding/json"
	"fmt"
	ez "github.com/ezclassify/client-go"
	"maps"
	"net/url"
	"strings"
)

// publicDriver marks a corpus case sent through the public API. Every other
// entry is the reason no public call can send that recording, and the case is
// replayed as raw wire traffic instead. The reasons are few on purpose: a new
// one is a reviewed decision, not a place to park a missing driver.
const publicDriver = "public"

const (
	readinessProbe           = "readiness is an orchestrator probe; no public method sends it"
	undeclaredMember         = "a typed request cannot carry a member the contract does not declare"
	undeclaredParameter      = "no public method sends a query parameter the contract does not declare"
	undeclaredRoute          = "no public method sends a route or method the contract does not declare"
	requiredMemberAlwaysSent = "a typed request always sends the members the contract requires"
	refusedBeforeTransport   = "the public method refuses this request before transport with an InvalidInputError"
)

var rawWireReasons = map[string]bool{
	readinessProbe:           true,
	undeclaredMember:         true,
	undeclaredParameter:      true,
	undeclaredRoute:          true,
	requiredMemberAlwaysSent: true,
	refusedBeforeTransport:   true,
}

// wireDrivers registers every corpus case literally, so a recording added
// upstream fails the coverage test until it has a reviewed driver.
var wireDrivers = map[string]string{
	"health.ok":                                                      publicDriver,
	"ready.ok":                                                       readinessProbe,
	"ctx-list.ok":                                                    publicDriver,
	"ctx-versions.ok":                                                publicDriver,
	"ctx-describe.ok":                                                publicDriver,
	"ctx-describe.crosswalks":                                        publicDriver,
	"ctx-describe.marked":                                            publicDriver,
	"ctx-document-audit.ok":                                          publicDriver,
	"ctx-document-audit.empty-text":                                  publicDriver,
	"ctx-describe.presets":                                           publicDriver,
	"ctx-scheme-summary.ok":                                          publicDriver,
	"ctx-scheme-concepts.ok":                                         publicDriver,
	"ctx-scheme-categories.ok":                                       publicDriver,
	"ctx-scheme-summary.layered":                                     publicDriver,
	"ctx-scheme-concepts.layered":                                    publicDriver,
	"ctx-scheme-categories.layered":                                  publicDriver,
	"ctx-scheme-formats.ok":                                          publicDriver,
	"ctx-scheme-scopes.ok":                                           publicDriver,
	"ctx-scheme-scopes.none":                                         publicDriver,
	"ctx-scheme-label-presets.ok":                                    publicDriver,
	"ctx-scheme-apply-preset.ok":                                     publicDriver,
	"ctx-scheme-apply-preset.unknown":                                publicDriver,
	"ctx-scheme-summary.full":                                        publicDriver,
	"ctx-scheme-concepts.full":                                       publicDriver,
	"ctx-scheme-categories.full":                                     publicDriver,
	"ctx-scheme-formats.full":                                        publicDriver,
	"ctx-scheme-scopes.full":                                         publicDriver,
	"ctx-scheme-label-presets.full":                                  publicDriver,
	"ctx-scheme-process.ok":                                          publicDriver,
	"ctx-scheme-process.empty":                                       publicDriver,
	"ctx-scheme-process.outputs-and-options":                         publicDriver,
	"ctx-scheme-validate.ok":                                         publicDriver,
	"ctx-scheme-validate.assertions":                                 publicDriver,
	"ctx-scheme-validate.both-sources":                               refusedBeforeTransport,
	"ctx-scheme-validate.no-source":                                  refusedBeforeTransport,
	"ctx-scheme-process.undeclared-format":                           publicDriver,
	"ctx-scheme-aggregate.input-format-with-sets":                    refusedBeforeTransport,
	"ctx-scheme-explain.ok":                                          publicDriver,
	"ctx-scheme-render.ok":                                           publicDriver,
	"ctx-scheme-render.empty":                                        publicDriver,
	"ctx-scheme-diff.ok":                                             publicDriver,
	"ctx-scheme-view.ok":                                             publicDriver,
	"ctx-scheme-view.marked":                                         publicDriver,
	"ctx-scheme-notices.ok":                                          publicDriver,
	"ctx-scheme-scan.ok":                                             publicDriver,
	"ctx-scheme-aggregate.ok":                                        publicDriver,
	"ctx-scheme-aggregate.assertion-sets":                            publicDriver,
	"ctx-scheme-process.selected-scheme":                             publicDriver,
	"ctx-scheme-process.unknown-scheme":                              publicDriver,
	"ctx-scheme-validate.selected-scheme":                            publicDriver,
	"ctx-scheme-aggregate.selected-scheme":                           publicDriver,
	"ctx-scheme-suggest.ok":                                          publicDriver,
	"ctx-scheme-suggest.gated-hidden":                                publicDriver,
	"ctx-scheme-suggest.gated-included":                              publicDriver,
	"ctx-scheme-suggest.facts":                                       publicDriver,
	"ctx-scheme-suggest.both-sources":                                refusedBeforeTransport,
	"ctx-scheme-suggest.no-source":                                   refusedBeforeTransport,
	"ctx-scheme-suggest.input-format-with-facts":                     refusedBeforeTransport,
	"ctx-scheme-suggest-for-categories.ok":                           publicDriver,
	"ctx-scheme-suggest-for-categories.empty-baseline":               publicDriver,
	"ctx-scheme-check-coverage.ok":                                   publicDriver,
	"ctx-scheme-check-coverage.excess":                               publicDriver,
	"ctx-scheme-check-coverage.pipeline":                             publicDriver,
	"ctx-scheme-check-coverage.pipeline-narrowed":                    publicDriver,
	"ctx-scheme-describe-concept.unknown-concept":                    publicDriver,
	"ctx-scheme-describe-concept.ok":                                 publicDriver,
	"ctx-access-check.marking":                                       publicDriver,
	"ctx-access-check.assertions":                                    publicDriver,
	"ctx-access-check.two-attributes":                                publicDriver,
	"ctx-access-describe-requirements.ok":                            publicDriver,
	"ctx-access-rules.ok":                                            publicDriver,
	"ctx-access-rules.marked":                                        publicDriver,
	"ctx-access-rules.no-policy":                                     publicDriver,
	"ctx-guide-scan.ok":                                              publicDriver,
	"ctx-guide-evaluate.ok":                                          publicDriver,
	"ctx-guide-evaluate.selections":                                  publicDriver,
	"ctx-guide-view.by-id":                                           publicDriver,
	"ctx-guide-view.ok":                                              publicDriver,
	"ctx-guide-view.marked":                                          publicDriver,
	"ctx-crosswalk-translate.marking":                                publicDriver,
	"ctx-crosswalk-translate.assertions":                             publicDriver,
	"ctx-crosswalk-describe.ok":                                      publicDriver,
	"ctx-crosswalk-describe.marked":                                  publicDriver,
	"ctx-crosswalk-translate.translated":                             publicDriver,
	"ctx-crosswalk-translate.forbidden":                              publicDriver,
	"ctx-crosswalk-translate.not-selected":                           publicDriver,
	"ctx-crosswalk-translate.both-sources":                           refusedBeforeTransport,
	"ctx-crosswalk-translate.no-source":                              refusedBeforeTransport,
	"ctx-scheme-process.undeclared-request-field-fix":                undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-preserveDeprecated": undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-normalization":      undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-skipRepair":         undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-skipInference":      undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-tolerant":           undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-validateOnly":       undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-scope":              undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-format":             undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-repairStrategy":     undeclaredMember,
	"ctx-scheme-process.undeclared-request-field-unknownField":       undeclaredMember,
	"ctx-scheme-validate.undeclared-field":                           undeclaredMember,
	"ctx-scheme-explain.undeclared-field":                            undeclaredMember,
	"ctx-scheme-aggregate.undeclared-field":                          undeclaredMember,
	"ctx-scheme-suggest.undeclared-field":                            undeclaredMember,
	"ctx-scheme-suggest-for-categories.undeclared-field":             undeclaredMember,
	"ctx-scheme-render.undeclared-field":                             undeclaredMember,
	"ctx-scheme-scan.undeclared-field":                               undeclaredMember,
	"ctx-document-audit.undeclared-guide-field":                      undeclaredMember,
	"ctx-scheme-process.missing-input":                               requiredMemberAlwaysSent,
	"health.undeclared-query-parameter":                              undeclaredParameter,
	"ctx-scheme-label-presets.undeclared-query-parameter":            undeclaredParameter,
	"ctx-scheme-process.undeclared-query-parameter":                  undeclaredParameter,
	"denial.authenticated-no-credential":                             publicDriver,
	"denial.authenticated-bad-credential":                            publicDriver,
	"denial.authenticated-health-still-open":                         publicDriver,
	"denial.context-without-carrier":                                 publicDriver,
	"denial.ctx-scheme-without-carrier":                              publicDriver,
	"route-miss.route-not-found":                                     undeclaredRoute,
	"route-miss.ctx-operation-not-found":                             undeclaredRoute,
	"method-guard.allow-header":                                      undeclaredRoute,
}

func wireCredentials(r wireCase) ez.RestClientCredentials {
	switch r.Auth.Credential {
	case "context-token":
		return ez.RestClientCredentials{ContextToken: "fixture-context-carrier"}
	case "invalid-bearer":
		return ez.RestClientCredentials{AccessToken: "fixture-refused-bearer", ContextToken: "fixture-context-carrier"}
	default:
		return ez.RestClientCredentials{}
	}
}
func decodeWireRequest[T any](body []byte) (T, error) {
	var v T
	e := json.Unmarshal(body, &v)
	return v, e
}
func wireCall[T any](body []byte, fn func(T) (interface{}, error)) (interface{}, error) {
	v, e := decodeWireRequest[T](body)
	if e != nil {
		return nil, e
	}
	return fn(v)
}
func wireOptionalQuery(u *url.URL, key string) *string {
	if !u.Query().Has(key) {
		return nil
	}
	v := u.Query().Get(key)
	return &v
}
func runWireDriver(ctx context.Context, c *ez.Client, r wireCase) (interface{}, error) {
	if wireDrivers[r.ID] != publicDriver {
		headers := maps.Clone(r.Request.Headers)
		if token := wireCredentials(r).ContextToken; token != "" {
			// Raw recordings can present a Context carrier outside the SDK's
			// automatic route scope. Honor their declared auth without changing it.
			if headers == nil {
				headers = make(map[string]string)
			}
			headers["x-ezclassify-context-token"] = token
		}
		b, status, e := ez.WireReplayRequestForTest(ctx, c, r.Request.Method, r.Request.Path, headers, r.Request.Body)
		if e != nil {
			return nil, e
		}
		if status != r.Response.Status || !equalWireJSON(b, r.Response.Body) {
			return nil, fmt.Errorf("private response mismatch")
		}
		return nil, nil
	}
	switch r.ID {
	case "denial.authenticated-no-credential", "denial.authenticated-bad-credential", "denial.context-without-carrier":
		r.Operation = "ctxSchemeProcess"
	case "denial.ctx-scheme-without-carrier":
		r.Operation = "ctxSchemeConcepts"
	}
	u, e := url.Parse(r.Request.Path)
	if e != nil {
		return nil, e
	}
	parts := strings.Split(strings.Trim(u.Path, "/"), "/")
	switch r.Operation {
	case "getHealth":
		return c.Ping(ctx)
	case "ctxList":
		return c.ListContexts(ctx)
	case "ctxVersions":
		return c.ListContextVersions(ctx, parts[2])
	}
	if len(parts) < 4 {
		return nil, fmt.Errorf("missing Context path")
	}
	return driveWire(ctx, ez.ConnectWith(c, parts[2], parts[3]), r, u)
}

// driveWire sends the case's recorded request through the public method of the
// Context handle that the case's operation names. u is the parsed recorded
// path, whose query carries the arguments of the GET operations and the
// scheme selector of every scheme operation. Passing the
// handle in lets the cache plans reuse one handle, and so one cache, across
// calls.
func driveWire(ctx context.Context, h *ez.EzClassifyContext, r wireCase, u *url.URL) (interface{}, error) {
	scheme := h.Scheme
	if selector := wireOptionalQuery(u, "scheme"); selector != nil {
		scheme = h.Schemes.Get(*selector)
	}
	switch r.Operation {
	case "ctxDescribe":
		description, e := h.Describe(ctx)
		if e != nil {
			return nil, e
		}
		if h.ResolvedVersion() != r.Response.Headers["x-ezclassify-resolved-version"] {
			return nil, fmt.Errorf("changed resolved version")
		}
		return description, nil
	case "ctxSchemeSummary":
		return scheme.Summary(ctx)
	case "ctxSchemeConcepts":
		return scheme.Concepts(ctx)
	case "ctxSchemeCategories":
		return scheme.Categories(ctx)
	case "ctxSchemeFormats":
		return scheme.Formats(ctx)
	case "ctxSchemeScopes":
		return scheme.Scopes(ctx)
	case "ctxSchemeLabelPresets":
		return scheme.LabelPresets(ctx)
	case "ctxSchemeDescribeConcept":
		return scheme.DescribeConcept(ctx, u.Query().Get("conceptId"), u.Query().Get("categoryId"), wireOptionalQuery(u, "view"))
	case "ctxAccessRules":
		return h.Access.Rules(ctx)
	case "ctxCrosswalkDescribe":
		return h.Crosswalk.Describe(ctx)
	case "ctxSchemeView":
		return scheme.View(ctx)
	case "ctxSchemeProcess":
		return wireCall[ez.CtxSchemeProcessRequest](r.Request.Body, func(v ez.CtxSchemeProcessRequest) (interface{}, error) { return scheme.Process(ctx, v) })
	case "ctxSchemeValidate":
		return wireCall[ez.CtxSchemeValidateRequest](r.Request.Body, func(v ez.CtxSchemeValidateRequest) (interface{}, error) { return scheme.Validate(ctx, v) })
	case "ctxSchemeApplyPreset":
		return wireCall[ez.ApplyPresetRequest](r.Request.Body, func(v ez.ApplyPresetRequest) (interface{}, error) { return scheme.ApplyPreset(ctx, v) })
	case "ctxSchemeExplain":
		return wireCall[ez.CtxSchemeExplainRequest](r.Request.Body, func(v ez.CtxSchemeExplainRequest) (interface{}, error) { return scheme.Explain(ctx, v) })
	case "ctxSchemeRender":
		return wireCall[ez.CtxSchemeRenderRequest](r.Request.Body, func(v ez.CtxSchemeRenderRequest) (interface{}, error) { return scheme.Render(ctx, v) })
	case "ctxSchemeDiff":
		return wireCall[ez.CtxSchemeDiffRequest](r.Request.Body, func(v ez.CtxSchemeDiffRequest) (interface{}, error) { return scheme.Diff(ctx, v) })
	case "ctxSchemeNotices":
		return wireCall[ez.CtxSchemeNoticesRequest](r.Request.Body, func(v ez.CtxSchemeNoticesRequest) (interface{}, error) { return scheme.Notices(ctx, v) })
	case "ctxSchemeScan":
		return wireCall[ez.CtxSchemeScanRequest](r.Request.Body, func(v ez.CtxSchemeScanRequest) (interface{}, error) { return scheme.Scan(ctx, v) })
	case "ctxSchemeAggregate":
		return wireCall[ez.CtxSchemeAggregateRequest](r.Request.Body, func(v ez.CtxSchemeAggregateRequest) (interface{}, error) { return scheme.Aggregate(ctx, v) })
	case "ctxSchemeSuggest":
		return wireCall[ez.CtxSchemeSuggestRequest](r.Request.Body, func(v ez.CtxSchemeSuggestRequest) (interface{}, error) { return scheme.Suggest(ctx, v) })
	case "ctxSchemeSuggestForCategories":
		return wireCall[ez.CtxSchemeSuggestForCategoriesRequest](r.Request.Body, func(v ez.CtxSchemeSuggestForCategoriesRequest) (interface{}, error) {
			return scheme.SuggestForCategories(ctx, v)
		})
	case "ctxSchemeCheckCoverage":
		return wireCall[ez.CtxSchemeCheckCoverageRequest](r.Request.Body, func(v ez.CtxSchemeCheckCoverageRequest) (interface{}, error) { return scheme.CheckCoverage(ctx, v) })
	case "ctxAccessCheck":
		return wireCall[ez.CtxAccessCheckRequest](r.Request.Body, func(v ez.CtxAccessCheckRequest) (interface{}, error) { return h.Access.Check(ctx, v) })
	case "ctxAccessDescribeRequirements":
		return wireCall[ez.CtxAccessDescribeRequirementsRequest](r.Request.Body, func(v ez.CtxAccessDescribeRequirementsRequest) (interface{}, error) {
			return h.Access.DescribeRequirements(ctx, v)
		})
	case "ctxGuideScan":
		return wireCall[ez.CtxGuideScanRequest](r.Request.Body, func(v ez.CtxGuideScanRequest) (interface{}, error) { return h.Guide.Scan(ctx, v) })
	case "ctxGuideEvaluate":
		return wireCall[ez.CtxGuideEvaluateRequest](r.Request.Body, func(v ez.CtxGuideEvaluateRequest) (interface{}, error) { return h.Guide.Evaluate(ctx, v) })
	case "ctxDocumentAudit":
		return wireCall[ez.CtxDocumentAuditRequest](r.Request.Body, func(v ez.CtxDocumentAuditRequest) (interface{}, error) { return h.Document.Audit(ctx, v) })
	case "ctxCrosswalkTranslate":
		return wireCall[ez.CtxCrosswalkTranslateRequest](r.Request.Body, func(v ez.CtxCrosswalkTranslateRequest) (interface{}, error) { return h.Crosswalk.Translate(ctx, v) })
	case "ctxGuideView":
		v, e := decodeWireRequest[ez.CtxGuideViewRequest](r.Request.Body)
		if e != nil {
			return nil, e
		}
		return h.Guide.View(ctx, v.GuideID)
	default:
		return nil, fmt.Errorf("unmapped public operation %q", r.Operation)
	}
}
