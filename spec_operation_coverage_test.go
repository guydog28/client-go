package ezclassify_test

// Operation coverage gate.
//
// spec/openapi.yaml decides what this client owes. Every operation it declares
// is listed here once: in specOperationMethods with the public method that
// sends it, or in specOperationsNotSent with the reason no public method does.
// A re-vendor that adds an operation fails until it is placed, and a row whose
// method was renamed or removed fails too, because each row is resolved
// against the real exported method set. The table declares the method; that
// the method sends its operation is proven by the wire corpus, whose public
// driver calls it for every recorded case of that operation.

import (
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"go.yaml.in/yaml/v3"

	ez "github.com/ezclassify/client-go"
)

// specOperationMethods maps each operationId to the public method that sends it.
var specOperationMethods = map[string]string{
	"getHealth":                     "Client.Ping",
	"ctxList":                       "Client.ListContexts",
	"ctxVersions":                   "Client.ListContextVersions",
	"ctxDescribe":                   "EzClassifyContext.Describe",
	"ctxSchemeSummary":              "CtxScheme.Summary",
	"ctxSchemeConcepts":             "CtxScheme.Concepts",
	"ctxSchemeCategories":           "CtxScheme.Categories",
	"ctxSchemeFormats":              "CtxScheme.Formats",
	"ctxSchemeGrammar":              "CtxScheme.Grammar",
	"ctxSchemeScopes":               "CtxScheme.Scopes",
	"ctxSchemeLabelPresets":         "CtxScheme.LabelPresets",
	"ctxSchemeApplyPreset":          "CtxScheme.ApplyPreset",
	"ctxSchemeProcess":              "CtxScheme.Process",
	"ctxSchemeValidate":             "CtxScheme.Validate",
	"ctxSchemeExplain":              "CtxScheme.Explain",
	"ctxSchemeNormalize":            "CtxScheme.Normalize",
	"ctxSchemeRender":               "CtxScheme.Render",
	"ctxSchemeDiff":                 "CtxScheme.Diff",
	"ctxSchemeView":                 "CtxScheme.View",
	"ctxSchemeNotices":              "CtxScheme.Notices",
	"ctxSchemeFilterConcepts":       "CtxScheme.FilterConcepts",
	"ctxSchemeScan":                 "CtxScheme.Scan",
	"ctxSchemeAggregate":            "CtxScheme.Aggregate",
	"ctxSchemeSuggest":              "CtxScheme.Suggest",
	"ctxSchemeSuggestForCategories": "CtxScheme.SuggestForCategories",
	"ctxSchemeCheckCoverage":        "CtxScheme.CheckCoverage",
	"ctxSchemeDescribeConcept":      "CtxScheme.DescribeConcept",
	"ctxAccessCheck":                "CtxAccess.Check",
	"ctxAccessDescribeRule":         "CtxAccess.DescribeRule",
	"ctxAccessDescribeRequirements": "CtxAccess.DescribeRequirements",
	"ctxAccessRules":                "CtxAccess.Rules",
	"ctxGuideScan":                  "CtxGuide.Scan",
	"ctxGuideEvaluate":              "CtxGuide.Evaluate",
	"ctxGuideDescribeRule":          "CtxGuide.DescribeRule",
	"ctxGuideFilterRules":           "CtxGuide.FilterRules",
	"ctxGuideView":                  "CtxGuide.View",
	"ctxDocumentAudit":              "CtxDocument.Audit",
	"ctxCrosswalkDescribe":          "CtxCrosswalk.Describe",
	"ctxCrosswalkTranslate":         "CtxCrosswalk.Translate",
}

// specOperationsNotSent maps each operationId no public method sends to the
// reason it is left out.
var specOperationsNotSent = map[string]string{
	"getReady": "readiness tells an orchestrator when to route traffic to a replica; Ping carries the engine version a client checks",
}

var publicTypes = map[string]reflect.Type{
	"Client":            reflect.TypeOf(&ez.Client{}),
	"EzClassifyContext": reflect.TypeOf(&ez.EzClassifyContext{}),
	"CtxScheme":         reflect.TypeOf(&ez.CtxScheme{}),
	"CtxGuide":          reflect.TypeOf(&ez.CtxGuide{}),
	"CtxAccess":         reflect.TypeOf(&ez.CtxAccess{}),
	"CtxDocument":       reflect.TypeOf(&ez.CtxDocument{}),
	"CtxCrosswalk":      reflect.TypeOf(&ez.CtxCrosswalk{}),
}

func specOperationIDs(t *testing.T) map[string]bool {
	t.Helper()
	raw, err := os.ReadFile(requestSchemaSpecPath)
	if err != nil {
		t.Fatal(err)
	}
	var document struct {
		Paths map[string]map[string]struct {
			OperationID string `yaml:"operationId"`
		} `yaml:"paths"`
	}
	if err := yaml.Unmarshal(raw, &document); err != nil {
		t.Fatal(err)
	}
	ids := map[string]bool{}
	for path, item := range document.Paths {
		for method, operation := range item {
			if operation.OperationID == "" {
				t.Fatalf("%s %s declares no operationId", strings.ToUpper(method), path)
			}
			if ids[operation.OperationID] {
				t.Fatalf("the contract declares operationId %q twice", operation.OperationID)
			}
			ids[operation.OperationID] = true
		}
	}
	if len(ids) == 0 {
		t.Fatal("no operations were read from the vendored contract")
	}
	return ids
}

func TestEverySpecOperationHasAPublicMethodOrAReason(t *testing.T) {
	ids := specOperationIDs(t)
	var unplaced []string
	for id := range ids {
		_, sent := specOperationMethods[id]
		_, excused := specOperationsNotSent[id]
		switch {
		case sent && excused:
			t.Errorf("%s is listed both as sent and as not sent", id)
		case !sent && !excused:
			unplaced = append(unplaced, id)
		}
	}
	sort.Strings(unplaced)
	for _, id := range unplaced {
		t.Errorf("the contract declares %s, which no public method sends and no reason excuses", id)
	}
	for _, table := range []map[string]string{specOperationMethods, specOperationsNotSent} {
		for id := range table {
			if !ids[id] {
				t.Errorf("%s is listed but the contract does not declare it", id)
			}
		}
	}
}

func TestEachListedMethodIsExported(t *testing.T) {
	for id, target := range specOperationMethods {
		typeName, methodName, _ := strings.Cut(target, ".")
		owner, ok := publicTypes[typeName]
		if !ok {
			t.Errorf("%s names %s, which is not a public client type", id, target)
			continue
		}
		if _, ok := owner.MethodByName(methodName); !ok {
			t.Errorf("%s names %s, which does not exist", id, target)
		}
	}
}
