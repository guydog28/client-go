package ezclassify_test

// Request-shape gate.
//
// Every Context operation that sends a body is driven through the real client
// with every optional field set, and the body it puts on the wire is validated
// against that operation's request schema in the vendored contract
// (spec/openapi.yaml). The request schemas close their objects
// (additionalProperties: false), so a field the contract renamed or removed
// fails here instead of as a 400 from a live server.
//
// The generated request structs make a renamed field unlikely in this SDK, but
// they do not make every field reachable: a field whose type the root package
// does not re-export cannot be set by a caller at all, because the type lives
// in internal/api. These cases are written against the public package only, so
// a field a caller cannot set is a field no case can send, and the
// every-field-reaches-the-wire check below names it.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"

	"go.yaml.in/yaml/v3"

	"github.com/ezclassify/sdk-go"
)

const (
	requestSchemaSpecPath = "spec/openapi.yaml"
	requestRoutePrefix    = "/v1/ctx/{ctxId}/{version}/"
	requestCaseContext    = "acme-marking"
	requestCaseVersion    = "1.0.0"
	requestCasePathPrefix = "/v1/ctx/" + requestCaseContext + "/" + requestCaseVersion + "/"
	componentSchemaPrefix = "#/components/schemas/"
)

// schemaAnnotations are keywords that describe a schema without constraining a
// value, so the validator may pass over them.
var schemaAnnotations = map[string]bool{
	"description": true,
	"title":       true,
	"default":     true,
	"example":     true,
	"examples":    true,
	"deprecated":  true,
	"readOnly":    true,
	"writeOnly":   true,
	"$comment":    true,
}

// operationKey names one operation of the contract by lowercase method and
// templated path, the way the OpenAPI document keys it.
type operationKey struct {
	method string
	path   string
}

// openAPIRequestSchemas holds the request schemas of the vendored OpenAPI
// document and validates bodies against them.
//
// The validator covers the JSON Schema subset those schemas use and fails
// closed: a keyword it does not implement is reported as an error rather than
// skipped, so a contract that starts using one forces this harness to learn it
// instead of silently passing everything that keyword would have rejected.
type openAPIRequestSchemas struct {
	components     map[string]any
	requestSchemas map[operationKey]any
}

var loadRequestSchemas = sync.OnceValues(func() (*openAPIRequestSchemas, error) {
	raw, err := os.ReadFile(requestSchemaSpecPath)
	if err != nil {
		return nil, err
	}
	var root yaml.Node
	if err := yaml.Unmarshal(raw, &root); err != nil {
		return nil, fmt.Errorf("parse %s: %w", requestSchemaSpecPath, err)
	}
	decoded, err := yamlToJSON(&root)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", requestSchemaSpecPath, err)
	}
	return newOpenAPIRequestSchemas(decoded)
})

func newOpenAPIRequestSchemas(document any) (*openAPIRequestSchemas, error) {
	doc, _ := document.(map[string]any)
	components, _ := lookup(doc, "components", "schemas").(map[string]any)
	paths, _ := doc["paths"].(map[string]any)
	if components == nil || paths == nil {
		return nil, errors.New("the OpenAPI document has no components.schemas or paths")
	}
	spec := &openAPIRequestSchemas{components: components, requestSchemas: map[operationKey]any{}}
	for path, item := range paths {
		operations, _ := item.(map[string]any)
		for method, operation := range operations {
			op, ok := operation.(map[string]any)
			if !ok {
				continue
			}
			if schema := lookup(op, "requestBody", "content", "application/json", "schema"); schema != nil {
				spec.requestSchemas[operationKey{method, path}] = schema
			}
		}
	}
	if len(spec.requestSchemas) == 0 {
		return nil, fmt.Errorf("read no request schemas out of %s", requestSchemaSpecPath)
	}
	return spec, nil
}

func requestSpec(t *testing.T) *openAPIRequestSchemas {
	t.Helper()
	spec, err := loadRequestSchemas()
	if err != nil {
		t.Fatalf("load request schemas: %v", err)
	}
	return spec
}

// lookup walks nested JSON objects by key and returns nil when any step is
// missing.
func lookup(value any, keys ...string) any {
	for _, key := range keys {
		object, ok := value.(map[string]any)
		if !ok {
			return nil
		}
		value = object[key]
	}
	return value
}

// yamlToJSON converts a YAML node into the values encoding/json produces, so a
// schema read from YAML compares like a body read from the wire. Only true and
// false are booleans and every number is a float64, which is the YAML 1.2 core
// reading JSON Schema expects. Anything that has no JSON form is an error.
func yamlToJSON(node *yaml.Node) (any, error) {
	switch node.Kind {
	case yaml.DocumentNode:
		if len(node.Content) != 1 {
			return nil, fmt.Errorf("line %d: expected one document", node.Line)
		}
		return yamlToJSON(node.Content[0])
	case yaml.AliasNode:
		return yamlToJSON(node.Alias)
	case yaml.MappingNode:
		object := make(map[string]any, len(node.Content)/2)
		for i := 0; i+1 < len(node.Content); i += 2 {
			key := node.Content[i]
			if key.Kind != yaml.ScalarNode {
				return nil, fmt.Errorf("line %d: mapping key is not a scalar", key.Line)
			}
			value, err := yamlToJSON(node.Content[i+1])
			if err != nil {
				return nil, err
			}
			object[key.Value] = value
		}
		return object, nil
	case yaml.SequenceNode:
		array := make([]any, 0, len(node.Content))
		for _, item := range node.Content {
			value, err := yamlToJSON(item)
			if err != nil {
				return nil, err
			}
			array = append(array, value)
		}
		return array, nil
	case yaml.ScalarNode:
		switch node.ShortTag() {
		case "!!null":
			return nil, nil
		case "!!bool":
			return strconv.ParseBool(strings.ToLower(node.Value))
		case "!!int", "!!float":
			number, err := strconv.ParseFloat(node.Value, 64)
			if err != nil {
				return nil, fmt.Errorf("line %d: number %q has no JSON form: %w", node.Line, node.Value, err)
			}
			return number, nil
		default:
			// Plain dates stay strings, as they are in JSON.
			return node.Value, nil
		}
	}
	return nil, fmt.Errorf("line %d: unsupported YAML node kind %d", node.Line, node.Kind)
}

// validate checks a request body against an operation's request schema. It
// returns one message per violation, empty when the body is valid, and an
// error when the schema uses something the validator cannot evaluate.
func (s *openAPIRequestSchemas) validate(method, path string, body any) ([]string, error) {
	schema, ok := s.requestSchemas[operationKey{method, path}]
	if !ok {
		return nil, fmt.Errorf("the contract declares no request body for %s %s", strings.ToUpper(method), path)
	}
	v := &schemaValidator{components: s.components}
	var violations []string
	v.check(body, schema, "$", &violations)
	if len(v.unsupported) > 0 {
		return violations, fmt.Errorf("the request-schema validator does not implement %s; teach schemaValidator.check the keyword rather than ignoring it", strings.Join(v.unsupported, ", "))
	}
	return violations, nil
}

// declaredProperties returns the top-level property names an operation's
// request schema declares. It follows $ref and the anyOf and if/then/else
// branches, so a field declared on any branch counts.
func (s *openAPIRequestSchemas) declaredProperties(method, path string) (map[string]bool, error) {
	root, ok := s.requestSchemas[operationKey{method, path}]
	if !ok {
		return nil, fmt.Errorf("the contract declares no request body for %s %s", strings.ToUpper(method), path)
	}
	return declaredPropertiesOf(s.components, root)
}

func declaredPropertiesOf(components map[string]any, root any) (map[string]bool, error) {
	names := map[string]bool{}
	pending := []any{root}
	for len(pending) > 0 {
		schema, ok := pending[len(pending)-1].(map[string]any)
		pending = pending[:len(pending)-1]
		if !ok {
			continue // a boolean schema declares nothing
		}
		if ref, ok := schema["$ref"].(string); ok {
			resolved, err := resolveComponent(components, ref)
			if err != nil {
				return nil, err
			}
			pending = append(pending, resolved)
		}
		if properties, ok := schema["properties"].(map[string]any); ok {
			for name := range properties {
				names[name] = true
			}
		}
		if branches, ok := schema["anyOf"].([]any); ok {
			pending = append(pending, branches...)
		}
		for _, key := range []string{"then", "else"} {
			if branch, ok := schema[key]; ok {
				pending = append(pending, branch)
			}
		}
	}
	return names, nil
}

func resolveComponent(components map[string]any, ref string) (any, error) {
	name, ok := strings.CutPrefix(ref, componentSchemaPrefix)
	if !ok {
		return nil, fmt.Errorf("unsupported $ref %q", ref)
	}
	schema, ok := components[name]
	if !ok {
		return nil, fmt.Errorf("unresolved $ref %q", ref)
	}
	return schema, nil
}

type schemaValidator struct {
	components map[string]any
	// unsupported collects keywords the validator met but cannot evaluate. It
	// is kept apart from the violations so a keyword inside an anyOf branch or
	// an if condition still fails the run instead of just failing that branch.
	unsupported []string
}

func (v *schemaValidator) check(value, schema any, at string, violations *[]string) {
	if allowed, ok := schema.(bool); ok {
		if !allowed {
			*violations = append(*violations, at+": no value is allowed here")
		}
		return
	}
	object, ok := schema.(map[string]any)
	if !ok {
		v.unsupported = append(v.unsupported, fmt.Sprintf("a schema of Go type %T (at %s)", schema, at))
		return
	}

	keywords := make([]string, 0, len(object))
	for keyword := range object {
		keywords = append(keywords, keyword)
	}
	sort.Strings(keywords)

	for _, keyword := range keywords {
		argument := object[keyword]
		switch keyword {
		case "$ref":
			resolved, err := resolveComponent(v.components, fmt.Sprint(argument))
			if err != nil {
				v.unsupported = append(v.unsupported, fmt.Sprintf("%v (at %s)", err, at))
				continue
			}
			v.check(value, resolved, at, violations)
		case "type":
			accepted, _ := argument.([]any)
			if accepted == nil {
				accepted = []any{argument}
			}
			matched := false
			names := make([]string, 0, len(accepted))
			for _, jsonType := range accepted {
				name := fmt.Sprint(jsonType)
				names = append(names, name)
				is, known := isJSONType(value, name)
				if !known {
					v.unsupported = append(v.unsupported, fmt.Sprintf("type %q (at %s)", name, at))
				}
				matched = matched || is
			}
			if !matched {
				*violations = append(*violations, fmt.Sprintf("%s: %s is not of type %s", at, show(value), strings.Join(names, " | ")))
			}
		case "enum":
			options, _ := argument.([]any)
			found := false
			for _, option := range options {
				found = found || reflect.DeepEqual(option, value)
			}
			if !found {
				*violations = append(*violations, fmt.Sprintf("%s: %s is not one of %s", at, show(value), show(argument)))
			}
		case "const":
			if !reflect.DeepEqual(argument, value) {
				*violations = append(*violations, fmt.Sprintf("%s: %s is not %s", at, show(value), show(argument)))
			}
		case "required":
			if body, ok := value.(map[string]any); ok {
				names, _ := argument.([]any)
				for _, name := range names {
					if _, present := body[fmt.Sprint(name)]; !present {
						*violations = append(*violations, fmt.Sprintf("%s: required property '%v' is missing", at, name))
					}
				}
			}
		case "properties":
			if body, ok := value.(map[string]any); ok {
				properties, _ := argument.(map[string]any)
				for _, name := range sortedKeys(properties) {
					if item, present := body[name]; present {
						v.check(item, properties[name], at+"."+name, violations)
					}
				}
			}
		case "additionalProperties":
			if body, ok := value.(map[string]any); ok {
				declared, _ := object["properties"].(map[string]any)
				for _, name := range sortedKeys(body) {
					if _, isDeclared := declared[name]; isDeclared {
						continue
					}
					if allowed, isBool := argument.(bool); isBool {
						if !allowed {
							*violations = append(*violations, fmt.Sprintf("%s: property '%s' is not allowed by the contract", at, name))
						}
						continue
					}
					v.check(body[name], argument, at+"."+name, violations)
				}
			}
		case "items":
			if array, ok := value.([]any); ok {
				for index, item := range array {
					v.check(item, argument, fmt.Sprintf("%s[%d]", at, index), violations)
				}
			}
		case "anyOf":
			branches, _ := argument.([]any)
			matched := false
			for _, branch := range branches {
				matched = matched || v.passes(value, branch, at)
			}
			if !matched {
				*violations = append(*violations, fmt.Sprintf("%s: %s matches none of the anyOf branches", at, show(value)))
			}
		case "if":
			branch, present := object["else"]
			if v.passes(value, argument, at) {
				branch, present = object["then"]
			}
			if present {
				v.check(value, branch, at, violations)
			}
		case "then", "else":
			// Applied by "if".
		default:
			if !schemaAnnotations[keyword] {
				v.unsupported = append(v.unsupported, fmt.Sprintf("'%s' (at %s)", keyword, at))
			}
		}
	}
}

func (v *schemaValidator) passes(value, schema any, at string) bool {
	var scratch []string
	v.check(value, schema, at, &scratch)
	return len(scratch) == 0
}

// isJSONType reports whether a decoded JSON value is of the named JSON Schema
// type, and whether the name is a type the validator knows.
func isJSONType(value any, jsonType string) (is, known bool) {
	switch jsonType {
	case "null":
		return value == nil, true
	case "object":
		_, ok := value.(map[string]any)
		return ok, true
	case "array":
		_, ok := value.([]any)
		return ok, true
	case "string":
		_, ok := value.(string)
		return ok, true
	case "boolean":
		_, ok := value.(bool)
		return ok, true
	case "number":
		_, ok := value.(float64)
		return ok, true
	case "integer":
		number, ok := value.(float64)
		return ok && !math.IsInf(number, 0) && number == math.Trunc(number), true
	}
	return false, false
}

func show(value any) string {
	encoded, err := json.Marshal(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	return string(encoded)
}

func sortedKeys[V any](object map[string]V) []string {
	keys := make([]string, 0, len(object))
	for key := range object {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

var (
	caseLevel = ezclassify.Assertion{
		CategoryID: "level",
		ConceptID:  "HIGH",
		Confidence: ezclassify.AssertionConfidenceCertain,
		Source:     "parse",
		Value:      ptr("H"),
	}

	caseDeterminations = ezclassify.GuideDeterminationInput{
		Findings: []ezclassify.GuideFindingDetermination{{
			FindingID:      "finding-1",
			DocumentDigest: "sha256:abc",
			Decision:       ezclassify.GuideFindingDecisionAccepted,
			Reason:         "Confirmed by analyst",
			ProposalOrigin: ptr(ezclassify.DecisionProposalOriginAnalystAuthored),
		}},
		ManualSelections: []ezclassify.ManualGuideSelectionInput{},
	}

	// casePartOf marks the input as a portion of a whole marking. PartOf is an
	// inline object in the contract, so the generator gives it no type name
	// and a caller spells the struct out.
	casePartOf = ezclassify.ProcessingContext{
		Role: ptr(ezclassify.ProcessingRolePart),
		PartOf: &struct {
			Assertions *[]ezclassify.Assertion `json:"assertions,omitempty"`
			Rendered   string                  `json:"rendered"`
		}{Rendered: "H//A"},
	}
)

// requestCase drives one client call that sends a body.
type requestCase struct {
	name   string
	route  string
	invoke func(ctx context.Context, cx *ezclassify.EzClassifyContext) error
}

// requestCases holds one case per client call that sends a body. Each sets
// every optional field its request takes, so the whole option surface reaches
// the wire. Calls that build a different body for the same route, like the two
// halves of aggregate, get their own case.
var requestCases = []requestCase{
	{"scheme.process", "scheme/process", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Process(ctx, ezclassify.CtxSchemeProcessRequest{
			Input:        "H//A",
			InputFormat:  ptr("marking/portion"),
			OutputFormat: ptr("marking/banner"),
			Surface:      ptr("portion"),
			Assertions:   &[]ezclassify.Assertion{caseLevel},
			Context:      &casePartOf,
			Parts:        &[]string{"H//A/PART"},
			Outputs:      &[]string{"marking/portion", "marking/banner"},
			SuggestFor:   &[]string{"level", "audience"},
		})
		return err
	}},
	{"scheme.explain", "scheme/explain", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Explain(ctx, ezclassify.CtxSchemeExplainRequest{
			Input:        "H//A",
			InputFormat:  ptr("marking/portion"),
			OutputFormat: ptr("marking/banner"),
			Surface:      ptr("portion"),
			Assertions:   &[]ezclassify.Assertion{caseLevel},
			Context:      &casePartOf,
			Parts:        &[]string{"H//A/PART"},
		})
		return err
	}},
	// The server takes the marking as text or as facts, never both, so each
	// gets a case.
	{"scheme.validate-text", "scheme/validate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Validate(ctx, ezclassify.CtxSchemeValidateRequest{
			Input:        ptr("H//A"),
			InputFormat:  ptr("marking/portion"),
			OutputFormat: ptr("marking/banner"),
			Surface:      ptr("portion"),
			Context:      &casePartOf,
			Parts:        &[]string{"H//A/PART"},
		})
		return err
	}},
	{"scheme.validate-assertions", "scheme/validate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Validate(ctx, ezclassify.CtxSchemeValidateRequest{
			Assertions:   &[]ezclassify.Assertion{caseLevel},
			OutputFormat: ptr("marking/banner"),
			Surface:      ptr("portion"),
			Context:      &casePartOf,
			Parts:        &[]string{"H//A/PART"},
		})
		return err
	}},
	{"scheme.suggest", "scheme/suggest", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Suggest(ctx, ezclassify.CtxSchemeSuggestRequest{
			Input:        ptr("H//"),
			InputFormat:  ptr("portion"),
			OutputFormat: ptr("banner"),
			Surface:      ptr("portion"),
			Filter:       ptr("hi"),
			IncludeGated: ptr(true),
		})
		return err
	}},
	{"scheme.suggest-facts", "scheme/suggest", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Suggest(ctx, ezclassify.CtxSchemeSuggestRequest{
			Assertions:   &[]ezclassify.Assertion{caseLevel},
			OutputFormat: ptr("banner"),
			Surface:      ptr("portion"),
			Filter:       ptr("hi"),
			IncludeGated: ptr(true),
		})
		return err
	}},
	{"scheme.suggest-for-categories", "scheme/suggest-for-categories", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").SuggestForCategories(ctx, ezclassify.CtxSchemeSuggestForCategoriesRequest{
			Assertions:   []ezclassify.Assertion{caseLevel},
			CategoryIds:  []string{"level", "audience"},
			OutputFormat: ptr("banner"),
			Surface:      ptr("portion"),
		})
		return err
	}},
	{"scheme.aggregate", "scheme/aggregate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{
			Inputs:       &[]string{"H", "L"},
			InputFormat:  ptr("portion"),
			OutputFormat: ptr("banner"),
			Surface:      ptr("portion"),
			Context:      &casePartOf,
		})
		return err
	}},
	{"scheme.aggregate-assertions", "scheme/aggregate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{
			AssertionSets: &[][]ezclassify.Assertion{{caseLevel}},
			OutputFormat:  ptr("banner"),
			Surface:       ptr("portion"),
			Context:       &casePartOf,
		})
		return err
	}},
	{"scheme.scan", "scheme/scan", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Scan(ctx, ezclassify.CtxSchemeScanRequest{
			Text:         "text H text",
			InputFormat:  ptr("portion"),
			OutputFormat: ptr("banner"),
		})
		return err
	}},
	{"scheme.render", "scheme/render", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Render(ctx, ezclassify.CtxSchemeRenderRequest{
			Assertions:   []ezclassify.Assertion{caseLevel},
			OutputFormat: ptr("banner"),
			Preservation: ptr(ezclassify.RenderPreservationFull),
		})
		return err
	}},
	{"scheme.check-coverage", "scheme/check-coverage", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").CheckCoverage(ctx, ezclassify.CtxSchemeCheckCoverageRequest{
			Children:     []string{"L"},
			Parent:       "H",
			InputFormat:  ptr("portion"),
			OutputFormat: ptr("banner"),
		})
		return err
	}},
	{"scheme.diff", "scheme/diff", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Diff(ctx, ezclassify.CtxSchemeDiffRequest{A: "H", B: "L", InputFormat: ptr("portion")})
		return err
	}},
	{"scheme.notices", "scheme/notices", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").Notices(ctx, ezclassify.CtxSchemeNoticesRequest{Assertions: []ezclassify.Assertion{caseLevel}})
		return err
	}},
	{"scheme.view", "scheme/view", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").View(ctx)
		return err
	}},
	{"scheme.apply-preset", "scheme/apply-preset", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Schemes.Get("acme").ApplyPreset(ctx, ezclassify.ApplyPresetRequest{PresetID: "default-label", OutputFormat: ptr("banner")})
		return err
	}},
	{"guide.scan", "guide/scan", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Guide.Scan(ctx, ezclassify.CtxGuideScanRequest{Text: "secret text", GuideID: ptr("g1")})
		return err
	}},
	{"guide.evaluate", "guide/evaluate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Guide.Evaluate(ctx, ezclassify.CtxGuideEvaluateRequest{GuideID: ptr("g1"), SelectedRuleIds: &[]string{"r1"}})
		return err
	}},
	{"guide.evaluateMany", "guide/evaluate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Guide.Evaluate(ctx, ezclassify.CtxGuideEvaluateRequest{Selections: &[]ezclassify.GuideRuleSelection{{GuideID: "g1", RuleID: "r1"}}})
		return err
	}},
	{"guide.view", "guide/view", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Guide.View(ctx, ptr("g1"))
		return err
	}},
	{"access.check", "access/check", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Access.Check(ctx, ezclassify.CtxAccessCheckRequest{
			Marking:        ptr("H"),
			Assertions:     &[]ezclassify.Assertion{caseLevel},
			UserAttributes: map[string]any{"clearance": "high"},
			Context:        &map[string]any{"orgs": []string{"org-1"}},
		})
		return err
	}},
	{"access.describe-requirements", "access/describe-requirements", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Access.DescribeRequirements(ctx, ezclassify.CtxAccessDescribeRequirementsRequest{
			Marking:    ptr("H"),
			Assertions: &[]ezclassify.Assertion{caseLevel},
		})
		return err
	}},
	// The server takes exactly one crosswalk source, so each gets a case.
	{"crosswalk.translate-marking", "crosswalk/translate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Crosswalk.Translate(ctx, ezclassify.CtxCrosswalkTranslateRequest{CrosswalkID: "cw-1", Marking: ptr("H")})
		return err
	}},
	{"crosswalk.translate-assertions", "crosswalk/translate", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Crosswalk.Translate(ctx, ezclassify.CtxCrosswalkTranslateRequest{
			CrosswalkID: "cw-1",
			Assertions:  &[]ezclassify.Assertion{caseLevel},
		})
		return err
	}},
	{"document.audit", "document/audit", func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
		_, err := cx.Document.Audit(ctx, ezclassify.CtxDocumentAuditRequest{
			Text:                "document text",
			GuideDeterminations: &caseDeterminations,
			ResponseProfile:     ptr(ezclassify.DocumentResponseProfileFull),
		})
		return err
	}},
}

// sentRequest is what one case put on the wire.
type sentRequest struct {
	method string
	path   string
	query  url.Values
	body   []byte
}

// send drives one case through the real client against a stub that answers
// every request with an empty object, and returns the single request it sent.
// Most results reject an empty object, so the call's own error is ignored:
// only the request matters here.
func send(t *testing.T, c requestCase) sentRequest {
	t.Helper()
	var mu sync.Mutex
	var sent []sentRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		mu.Lock()
		sent = append(sent, sentRequest{method: r.Method, path: r.URL.Path, query: r.URL.Query(), body: body})
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{}")
	}))
	defer srv.Close()

	_ = c.invoke(context.Background(), ezclassify.Connect(srv.URL, requestCaseContext, requestCaseVersion))

	mu.Lock()
	defer mu.Unlock()
	if len(sent) != 1 {
		t.Fatalf("%s sent %d requests; expected exactly one", c.name, len(sent))
	}
	return sent[0]
}

func decodeBody(t *testing.T, name string, raw []byte) map[string]any {
	t.Helper()
	var body map[string]any
	if err := json.Unmarshal(raw, &body); err != nil {
		t.Fatalf("%s sent a body that is not a JSON object: %v: %s", name, err, raw)
	}
	return body
}

func TestRequestBodiesValidateAgainstTheContractRequestSchemas(t *testing.T) {
	spec := requestSpec(t)
	for _, c := range requestCases {
		t.Run(c.name, func(t *testing.T) {
			request := send(t, c)
			if request.method != http.MethodPost {
				t.Fatalf("%s used %s, not POST", c.name, request.method)
			}
			if want := requestCasePathPrefix + c.route; request.path != want {
				t.Fatalf("%s posted to %s, not %s", c.name, request.path, want)
			}
			violations, err := spec.validate("post", requestRoutePrefix+c.route, decodeBody(t, c.name, request.body))
			if err != nil {
				t.Fatal(err)
			}
			if len(violations) > 0 {
				t.Fatalf("%s sent a body the contract rejects:\n  %s\n  %s", c.name, request.body, strings.Join(violations, "\n  "))
			}
		})
	}
}

// A new operation with a body cannot ship without its body being checked.
func TestEveryContractRequestBodyHasACase(t *testing.T) {
	spec := requestSpec(t)
	covered := map[string]bool{}
	for _, c := range requestCases {
		covered[requestRoutePrefix+c.route] = true
	}
	var missing, nonPost []string
	for key := range spec.requestSchemas {
		switch {
		case key.method != "post":
			nonPost = append(nonPost, strings.ToUpper(key.method)+" "+key.path)
		case !covered[key.path]:
			missing = append(missing, "POST "+key.path)
		}
	}
	sort.Strings(missing)
	sort.Strings(nonPost)
	if len(missing) > 0 {
		t.Errorf("no request-schema case for: %s", strings.Join(missing, ", "))
	}
	if len(nonPost) > 0 {
		t.Errorf("the contract now has request bodies on non-POST operations; extend the cases to cover them: %s", strings.Join(nonPost, ", "))
	}
}

// sentFieldsByRoute drives every case and collects the top-level body fields
// actually sent to each route.
func sentFieldsByRoute(t *testing.T, cases []requestCase) map[string]map[string]bool {
	t.Helper()
	sent := map[string]map[string]bool{}
	for _, c := range cases {
		if sent[c.route] == nil {
			sent[c.route] = map[string]bool{}
		}
		for field := range decodeBody(t, c.name, send(t, c).body) {
			sent[c.route][field] = true
		}
	}
	return sent
}

// unsentContractFields returns, per route, the fields its request schema
// declares that no case sent.
func unsentContractFields(t *testing.T, spec *openAPIRequestSchemas, sent map[string]map[string]bool) map[string][]string {
	t.Helper()
	unsent := map[string][]string{}
	for route, fields := range sent {
		declared, err := spec.declaredProperties("post", requestRoutePrefix+route)
		if err != nil {
			t.Fatal(err)
		}
		for name := range declared {
			if !fields[name] {
				unsent[route] = append(unsent[route], name)
			}
		}
		sort.Strings(unsent[route])
	}
	return unsent
}

func formatUnsent(unsent map[string][]string) string {
	parts := make([]string, 0, len(unsent))
	for _, route := range sortedKeys(unsent) {
		parts = append(parts, route+": "+strings.Join(unsent[route], ", "))
	}
	return strings.Join(parts, "; ")
}

// Every field a request schema declares is sent by at least one case. Each
// case sets every optional field its request takes, so a declared field no
// case sends is one a caller has no way to set. Without this, the contract can
// add an option and the SDK can ship without a way to reach it while every
// body it does send still validates.
func TestEveryContractRequestPropertyReachesTheWire(t *testing.T) {
	spec := requestSpec(t)
	if unsent := unsentContractFields(t, spec, sentFieldsByRoute(t, requestCases)); len(unsent) > 0 {
		t.Fatalf("the client never sends these contract fields: %s", formatUnsent(unsent))
	}
}

// Negative controls. A gate that cannot fail proves nothing, so each failure
// mode the gate exists for is pinned here.

func TestUnsentFieldCheckNamesAFieldNoCaseSends(t *testing.T) {
	spec := requestSpec(t)
	cases := make([]requestCase, 0, len(requestCases))
	for _, c := range requestCases {
		if c.name == "scheme.scan" {
			c.invoke = func(ctx context.Context, cx *ezclassify.EzClassifyContext) error {
				_, err := cx.Scheme.Scan(ctx, ezclassify.CtxSchemeScanRequest{
					Text:        "text H text",
					InputFormat: ptr("portion"),
				})
				return err
			}
		}
		cases = append(cases, c)
	}
	unsent := unsentContractFields(t, spec, sentFieldsByRoute(t, cases))
	want := map[string][]string{"scheme/scan": {"outputFormat"}}
	if !reflect.DeepEqual(unsent, want) {
		t.Fatalf("dropping outputFormat from the scan case reported %q, want %q", formatUnsent(unsent), formatUnsent(want))
	}
}

func TestDeclaredPropertiesFollowRefsAndBranches(t *testing.T) {
	components := map[string]any{
		"Base": map[string]any{"properties": map[string]any{"fromRef": true}},
	}
	schema := map[string]any{
		"$ref":       componentSchemaPrefix + "Base",
		"properties": map[string]any{"direct": true},
		"anyOf":      []any{map[string]any{"properties": map[string]any{"fromAnyOf": true}}},
		"if":         map[string]any{"properties": map[string]any{"fromIf": true}},
		"then":       map[string]any{"properties": map[string]any{"fromThen": true}},
		"else":       map[string]any{"properties": map[string]any{"fromElse": true}},
	}
	got, err := declaredPropertiesOf(components, schema)
	if err != nil {
		t.Fatal(err)
	}
	// The if condition only selects a branch; it declares nothing of its own.
	want := map[string]bool{"direct": true, "fromRef": true, "fromAnyOf": true, "fromThen": true, "fromElse": true}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("declared properties = %v, want %v", sortedKeys(got), sortedKeys(want))
	}
}

func TestValidatorRejectsRemovedFields(t *testing.T) {
	spec := requestSpec(t)
	cases := []struct {
		route string
		body  map[string]any
		field string
	}{
		{"scheme/process", map[string]any{"input": "H", "format": "banner"}, "format"},
		{"scheme/process", map[string]any{"input": "H", "scope": "portion"}, "scope"},
		{"scheme/validate", map[string]any{"input": "H", "scope": "portion"}, "scope"},
		{"scheme/suggest", map[string]any{"input": "H", "format": "banner"}, "format"},
		{"scheme/aggregate", map[string]any{"inputs": []any{"H"}, "format": "banner", "scope": "portion"}, "format"},
		{"scheme/aggregate", map[string]any{"assertionSets": []any{}, "format": "banner"}, "format"},
		{"scheme/diff", map[string]any{"a": "H", "b": "L", "format": "banner"}, "format"},
		{"scheme/process", map[string]any{"input": "H", "tolerant": true}, "tolerant"},
		{"scheme/process", map[string]any{"input": "H", "normalization": "loose"}, "normalization"},
		{"scheme/explain", map[string]any{"input": "H", "outputs": []any{"marking/banner"}}, "outputs"},
		{"scheme/explain", map[string]any{"input": "H", "suggestFor": []any{"level"}}, "suggestFor"},
		{"scheme/suggest-for-categories", map[string]any{"assertions": []any{}, "categoryIds": []any{}, "input": "H"}, "input"},
	}
	for _, c := range cases {
		violations, err := spec.validate("post", requestRoutePrefix+c.route, c.body)
		if err != nil {
			t.Fatal(err)
		}
		if !containsMention(violations, "'"+c.field+"'") {
			t.Errorf("%s with %s: want a violation naming '%s', got %q", c.route, show(c.body), c.field, violations)
		}
	}
}

func TestValidatorRejectsAMissingRequiredFieldAndAWrongType(t *testing.T) {
	spec := requestSpec(t)
	missing, err := spec.validate("post", requestRoutePrefix+"scheme/process", map[string]any{"surface": "portion"})
	if err != nil {
		t.Fatal(err)
	}
	if !containsMention(missing, "'input'") {
		t.Errorf("a process body without input: want a violation naming 'input', got %q", missing)
	}

	wrongType, err := spec.validate("post", requestRoutePrefix+"scheme/render", map[string]any{
		"assertions": []any{map[string]any{
			"categoryId": "level",
			"conceptId":  "HIGH",
			"confidence": "sure",
			"source":     "parse",
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(wrongType) == 0 {
		t.Error("a render body with an unknown confidence validated")
	}
}

func TestValidatorAcceptsTheRenamedFields(t *testing.T) {
	spec := requestSpec(t)
	violations, err := spec.validate("post", requestRoutePrefix+"scheme/process", map[string]any{
		"input":        "H",
		"inputFormat":  "portion",
		"outputFormat": "banner",
		"surface":      "portion",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(violations) > 0 {
		t.Errorf("a process body with the current field names was rejected: %q", violations)
	}
}

func TestValidatorRefusesAKeywordItDoesNotImplement(t *testing.T) {
	v := &schemaValidator{}
	var violations []string
	v.check("x", map[string]any{"type": "string", "pattern": "^x$"}, "$", &violations)
	if !containsMention(v.unsupported, "'pattern'") {
		t.Fatalf("an unimplemented keyword was not reported: unsupported = %q", v.unsupported)
	}
}

func containsMention(messages []string, needle string) bool {
	for _, message := range messages {
		if strings.Contains(message, needle) {
			return true
		}
	}
	return false
}
