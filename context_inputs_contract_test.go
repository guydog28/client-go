package ezclassify_test

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
)

func TestSuggestExclusiveInputContract(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		valid      bool
	}{
		{"empty text", `{"input":"","inputFormat":"portion","outputFormat":"","surface":"portion","filter":"","includeGated":false}`, true},
		{"empty facts", `{"assertions":[],"outputFormat":"","surface":"portion","filter":"","includeGated":false}`, true},
		{"neither", `{}`, false},
		{"both", `{"input":"","assertions":[]}`, false},
		{"facts text format", `{"assertions":[],"inputFormat":""}`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var request ezclassify.CtxSchemeSuggestRequest
			if err := json.Unmarshal([]byte(tc.body), &request); err != nil {
				t.Fatal(err)
			}
			calls := 0
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				if r.URL.Path != "/v1/ctx/fixture/1.0.0/scheme/suggest" {
					t.Errorf("wrong route %s", r.URL.Path)
				}
				if r.Header.Get("Authorization") != "Bearer fixture-access" || r.Header.Get("x-ezclassify-context-token") != "fixture-context" {
					t.Error("credentials changed")
				}
				var got, expected map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
					t.Error(err)
				}
				if err := json.Unmarshal([]byte(tc.body), &expected); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("request=%v want=%v", got, expected)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{"resolution":{"contextId":"fixture","resolvedVersion":"1.0.0","artifacts":[]},"result":{"suggestions":[],"warnings":[]}}`)
			}))
			defer srv.Close()
			cx := ezclassify.ConnectWith(ezclassify.New(srv.URL, ezclassify.WithCredentials(ezclassify.RestClientCredentials{AccessToken: "fixture-access", ContextToken: "fixture-context"})), "fixture", "1.0.0")
			result, err := cx.Scheme.Suggest(context.Background(), request)
			if tc.valid {
				if err != nil || result == nil {
					t.Fatalf("result=%v error=%v", result, err)
				}
				if _, err = cx.Scheme.Suggest(context.Background(), request); err != nil {
					t.Fatal(err)
				}
				if calls != 1 {
					t.Fatalf("cache/request count=%d want1", calls)
				}
			} else if err == nil || result != nil || calls != 0 {
				t.Fatalf("invalid request: result=%v error=%v calls=%d", result, err, calls)
			}
		})
	}
}

// Every operation whose request carries its source as text or as facts,
// exactly one of the two, refuses a request that breaks that rule before it is
// sent, with the same typed error.
func TestSourceChoiceMistakesAreRefusedBeforeTransport(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { calls++; w.WriteHeader(400) }))
	defer srv.Close()
	cx := ezclassify.Connect(srv.URL, "fixture", "1.0.0")
	ctx := context.Background()
	facts := &[]ezclassify.Assertion{}
	sets := &[][]ezclassify.Assertion{{}}
	marking := ptr("INTERNAL")
	for _, tc := range []struct {
		name, op, detail string
		call             func() error
	}{
		{"validate facts with format", "scheme.validate", "input-format-with-assertions", func() error {
			_, err := cx.Scheme.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Assertions: facts, InputFormat: ptr("")})
			return err
		}},
		{"aggregate neither", "scheme.aggregate", "input-or-assertions", func() error {
			_, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{})
			return err
		}},
		{"aggregate both", "scheme.aggregate", "input-or-assertions", func() error {
			_, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: &[]string{"INTERNAL"}, AssertionSets: sets})
			return err
		}},
		{"aggregate sets with format", "scheme.aggregate", "input-format-with-assertions", func() error {
			_, err := cx.Scheme.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{AssertionSets: sets, InputFormat: ptr("default")})
			return err
		}},
		{"translate neither", "crosswalk.translate", "input-or-assertions", func() error {
			_, err := cx.Crosswalk.Translate(ctx, ezclassify.CtxCrosswalkTranslateRequest{CrosswalkID: "a-to-b"})
			return err
		}},
		{"translate both", "crosswalk.translate", "input-or-assertions", func() error {
			_, err := cx.Crosswalk.Translate(ctx, ezclassify.CtxCrosswalkTranslateRequest{CrosswalkID: "a-to-b", Marking: marking, Assertions: facts})
			return err
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call()
			var invalid *ezclassify.InvalidInputError
			if !errors.As(err, &invalid) || invalid.Op != tc.op || invalid.Detail != tc.detail {
				t.Fatalf("error = %T(%v), want an *InvalidInputError for %s with detail %q", err, err, tc.op, tc.detail)
			}
			if calls != 0 {
				t.Fatalf("%d requests sent, want none", calls)
			}
		})
	}
}

func TestDescribeConceptResource404IsNotContext404(t *testing.T) {
	for _, tc := range []struct {
		name, body string
		binding    bool
	}{
		{"resource", `{"error":"concept_not_found","code":"NOT_FOUND","detail":"conceptId: NOPE"}`, false},
		{"denial", `{"error":"not_found"}`, true},
		{"binding", `{"error":"context not found"}`, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(stubStatus(404, tc.body))
			defer srv.Close()
			cx := ezclassify.Connect(srv.URL, "fixture", "1.0.0")
			_, err := cx.Scheme.DescribeConcept(context.Background(), "NOPE", "level", nil)
			var apiErr *ezclassify.APIError
			var binding *ezclassify.ContextNotFoundError
			if !errors.As(err, &apiErr) || apiErr.StatusCode != 404 || string(apiErr.Body) != tc.body {
				t.Fatalf("error=%v", err)
			}
			if errors.As(err, &binding) != tc.binding {
				t.Fatalf("binding classification %T(%v),want%v", err, err, tc.binding)
			}
		})
	}
}

func TestSourceOpenLabelsAndValuesStayOnTheWire(t *testing.T) {
	for _, source := range []string{"parse", "user", "stated", "custom-author", "infer", "default", "context", "repair"} {
		t.Run(source, func(t *testing.T) {
			payload := `{"input":"","assertions":[{"categoryId":"level","conceptId":"H","confidence":"certain","source":"` + source + `","sourceRule":"fixture-rule","value":"20351231"}]}`
			var req ezclassify.CtxSchemeProcessRequest
			if err := json.Unmarshal([]byte(payload), &req); err != nil {
				t.Fatal(err)
			}
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var got, expected map[string]interface{}
				if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
					t.Error(err)
				}
				if err := json.Unmarshal([]byte(payload), &expected); err != nil {
					t.Error(err)
				}
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("wire=%v expected=%v", got, expected)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(400)
				_, _ = io.WriteString(w, `{"error":"Invalid input","code":"INVALID_INPUT","detail":"outputFormat: undeclared"}`)
			}))
			defer srv.Close()
			cx := ezclassify.Connect(srv.URL, "fixture", "1.0.0")
			result, err := cx.Scheme.Process(context.Background(), req)
			var apiErr *ezclassify.APIError
			if result != nil || !errors.As(err, &apiErr) || apiErr.StatusCode != 400 || !strings.Contains(string(apiErr.Body), `"detail":"outputFormat: undeclared"`) {
				t.Fatalf("result=%v error=%v", result, err)
			}
		})
	}
}
