package ezclassify_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ezclassify/client-go"
)

var cachedProcessBody = fixtureAs[ezclassify.ContextSchemeProcessResult](`{"resolution":{"contextId":"c","resolvedVersion":"1.0.0","artifacts":[]},"result":{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}}`)

func TestContextProcessReusesValidatedResponse(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	for i := 0; i < 2; i++ {
		got, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"})
		if err != nil || got == nil || got.Result.Rendered != "H" {
			t.Fatalf("Process returned %v, %v", got, err)
		}
	}
	if got := calls.Load(); got != 1 {
		t.Fatalf("repeated process used %d wire requests, want 1", got)
	}
}

// A failure and a malformed success are both errors, and neither is filed: the
// next call asks the server again and the first complete answer is the one that
// stays. A 200 that lacks a required member is refused wherever the member
// sits, so an envelope with an empty result object counts the same as a body of
// the wrong shape.
func TestContextCacheRetriesFailuresAndMalformedSuccess(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt := calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		switch attempt {
		case 1:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(w, `{"error":"retry"}`)
		case 2:
			_, _ = io.WriteString(w, `{"unexpected":"shape"}`)
		case 3:
			_, _ = io.WriteString(w, `{"resolution":{"contextId":"c","resolvedVersion":"1.0.0","artifacts":[]},"result":{}}`)
		default:
			_, _ = io.WriteString(w, cachedProcessBody)
		}
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	if _, err := cx.Scheme.Process(context.Background(), req); err == nil {
		t.Fatal("server error was not returned")
	}
	for _, member := range []string{"/resolution", "/result/"} {
		_, err := cx.Scheme.Process(context.Background(), req)
		var malformed *ezclassify.MalformedResponseError
		if !errors.As(err, &malformed) || !strings.HasPrefix(malformed.Member, member) {
			t.Fatalf("malformed success returned %T(%v), want *MalformedResponseError at %s", err, err, member)
		}
	}
	for i := 0; i < 2; i++ {
		got, err := cx.Scheme.Process(context.Background(), req)
		if err != nil || got == nil || got.Result.Rendered != "H" {
			t.Fatalf("recovery returned %v, %v", got, err)
		}
	}
	if got := calls.Load(); got != 4 {
		t.Fatalf("failure and malformed responses must be retried: %d wire calls, want 4", got)
	}
}

func TestContextCacheClearAndDisable(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	_, _ = cx.Scheme.Process(context.Background(), req)
	cx.ClearCache()
	_, _ = cx.Scheme.Process(context.Background(), req)
	if got := calls.Load(); got != 2 {
		t.Fatalf("clear cache: %d wire calls, want 2", got)
	}
	uncached := ezclassify.ConnectWith(ezclassify.New(srv.URL, ezclassify.WithResponseCacheDisabled()), "c", "1.0.0")
	_, _ = uncached.Scheme.Process(context.Background(), req)
	_, _ = uncached.Scheme.Process(context.Background(), req)
	if got := calls.Load(); got != 4 {
		t.Fatalf("disabled cache: %d wire calls, want 4", got)
	}
}

func TestContextCachePendingCallSurvivesOneCallerCancellation(t *testing.T) {
	var calls atomic.Int32
	entered := make(chan struct{})
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			close(entered)
		}
		<-release
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	first := make(chan error, 1)
	go func() {
		_, err := cx.Scheme.Process(context.Background(), req)
		first <- err
	}()
	<-entered
	secondCtx, cancel := context.WithCancel(context.Background())
	second := make(chan error, 1)
	go func() {
		_, err := cx.Scheme.Process(secondCtx, req)
		second <- err
	}()
	time.Sleep(20 * time.Millisecond)
	cancel()
	if err := <-second; err != context.Canceled {
		t.Fatalf("second caller: %v, want context.Canceled", err)
	}
	close(release)
	if err := <-first; err != nil {
		t.Fatalf("first caller was canceled by second: %v", err)
	}
	if got := calls.Load(); got != 1 {
		t.Fatalf("concurrent identical callers used %d wire calls, want 1", got)
	}
}

func TestContextAccessDecisionIsNeverServedFromClientMemory(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"resolution":{"contextId":"c","resolvedVersion":"1.0.0","artifacts":[]},"result":{"granted":false,"rules":[],"warnings":[]}}`)
	}))
	t.Cleanup(srv.Close)
	req := ezclassify.CtxAccessCheckRequest{}
	plain := ezclassify.Connect(srv.URL, "c", "1.0.0")
	for i := 0; i < 2; i++ {
		if _, err := plain.Access.Check(context.Background(), req); err != nil {
			t.Fatal(err)
		}
	}
	if got := calls.Load(); got != 2 {
		t.Fatalf("access check used %d wire calls, want 2", got)
	}
}

func TestContextCacheClearsOnUnauthorizedAndVersionChange(t *testing.T) {
	var calls atomic.Int32
	var version atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v1/ctx/c/latest" {
			version.Store(2)
			w.Header().Set("X-EzClassify-Resolved-Version", "2.0.0")
			_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextDescription](`{"id":"c","version":"2.0.0"}`))
			return
		}
		if r.URL.Path == "/v1/ctx/c/latest/scheme/validate" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = io.WriteString(w, `{"error":"expired"}`)
			return
		}
		if version.Load() == 1 {
			w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		} else {
			w.Header().Set("X-EzClassify-Resolved-Version", "2.0.0")
		}
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	version.Store(1)
	cx := ezclassify.Connect(srv.URL, "c", "latest")
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	_, _ = cx.Scheme.Process(context.Background(), req)
	_, _ = cx.Scheme.Process(context.Background(), req)
	if got := calls.Load(); got != 1 {
		t.Fatalf("warm cache used %d calls, want 1", got)
	}
	if _, err := cx.Scheme.Validate(context.Background(), ezclassify.CtxSchemeValidateRequest{Input: ptr("H")}); err == nil {
		t.Fatal("401 should fail")
	}
	_, _ = cx.Scheme.Process(context.Background(), req)
	if got := calls.Load(); got != 3 {
		t.Fatalf("401 did not clear cached process: %d calls, want 3", got)
	}
	if err := cx.Resolve(context.Background()); err != nil {
		t.Fatalf("Resolve: %v", err)
	}
	_, _ = cx.Scheme.Process(context.Background(), req)
	if got := calls.Load(); got != 5 {
		t.Fatalf("version change did not clear cache: %d calls, want 5", got)
	}
}

// countingContextServer answers every Context route with body inside a
// resolution envelope and counts the requests it receives per path.
func countingContextServer(t *testing.T, body func(path string, attempt int32) string) (*httptest.Server, func(path string) int32) {
	t.Helper()
	var mu sync.Mutex
	calls := map[string]*atomic.Int32{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		counter := calls[r.URL.Path]
		if counter == nil {
			counter = &atomic.Int32{}
			calls[r.URL.Path] = counter
		}
		mu.Unlock()
		attempt := counter.Add(1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		_, _ = io.WriteString(w, contextFixture(r.URL.Path, `{"resolution":{"contextId":"c","resolvedVersion":"1.0.0","artifacts":[]},"result":`+body(r.URL.Path, attempt)+`}`))
	}))
	t.Cleanup(srv.Close)
	return srv, func(path string) int32 {
		mu.Lock()
		defer mu.Unlock()
		if counter := calls[path]; counter != nil {
			return counter.Load()
		}
		return 0
	}
}

// requireEveryFieldVaried fails when variants does not name every field of
// the request type, so a field the contract adds later gets a variant of its
// own instead of going unexamined.
func requireEveryFieldVaried(t *testing.T, request any, variants map[string]any) {
	t.Helper()
	requestType := reflect.TypeOf(request)
	for i := 0; i < requestType.NumField(); i++ {
		if _, ok := variants[requestType.Field(i).Name]; !ok {
			t.Errorf("%s.%s has no variant: every request field must be shown to reach the cache key", requestType.Name(), requestType.Field(i).Name)
		}
	}
	if len(variants) != requestType.NumField() {
		t.Errorf("%s has %d fields and %d variants", requestType.Name(), requestType.NumField(), len(variants))
	}
}

// Every request field takes part in the key, including the ones that only add
// to an answer, so a request that differs in any one of them is a new wire
// request and repeating the base request is not.
func TestContextCacheKeysEveryRequestField(t *testing.T) {
	stated := []ezclassify.Assertion{{ConceptID: "L_H", CategoryID: "level", Confidence: ezclassify.AssertionConfidenceCertain, Source: "stated"}}
	part := &ezclassify.ProcessingContext{Role: ptr(ezclassify.ProcessingRolePart)}
	srv, calls := countingContextServer(t, func(path string, _ int32) string {
		if strings.HasSuffix(path, "/scheme/explain") {
			return `{"input":"H","schemeId":"s","entries":[],"repairs":[],"validationErrors":[],"inferences":[],"finalAssertions":[],"rendered":"H"}`
		}
		return `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()

	t.Run("process", func(t *testing.T) {
		base := ezclassify.CtxSchemeProcessRequest{Input: "H"}
		variants := map[string]any{
			"Input":        ezclassify.CtxSchemeProcessRequest{Input: "L"},
			"Assertions":   ezclassify.CtxSchemeProcessRequest{Input: "H", Assertions: &stated},
			"Context":      ezclassify.CtxSchemeProcessRequest{Input: "H", Context: part},
			"InputFormat":  ezclassify.CtxSchemeProcessRequest{Input: "H", InputFormat: ptr("marking/banner")},
			"OutputFormat": ezclassify.CtxSchemeProcessRequest{Input: "H", OutputFormat: ptr("marking/portion")},
			"Outputs":      ezclassify.CtxSchemeProcessRequest{Input: "H", Outputs: &[]string{"marking/portion"}},
			"Parts":        ezclassify.CtxSchemeProcessRequest{Input: "H", Parts: &[]string{"(H)"}},
			"SuggestFor":   ezclassify.CtxSchemeProcessRequest{Input: "H", SuggestFor: &[]string{"level"}},
			"Surface":      ezclassify.CtxSchemeProcessRequest{Input: "H", Surface: ptr("marking")},
		}
		requireEveryFieldVaried(t, base, variants)
		const path = "/v1/ctx/c/1.0.0/scheme/process"
		if _, err := cx.Scheme.Process(ctx, base); err != nil {
			t.Fatal(err)
		}
		for _, name := range sortedKeys(variants) {
			before := calls(path)
			if _, err := cx.Scheme.Process(ctx, variants[name].(ezclassify.CtxSchemeProcessRequest)); err != nil {
				t.Fatal(err)
			}
			if calls(path) != before+1 {
				t.Errorf("a process request differing only in %s was answered from the cache", name)
			}
		}
		before := calls(path)
		if _, err := cx.Scheme.Process(ctx, base); err != nil {
			t.Fatal(err)
		}
		if calls(path) != before {
			t.Error("repeating the base process request went to the wire")
		}
	})

	t.Run("validate", func(t *testing.T) {
		base := ezclassify.CtxSchemeValidateRequest{Input: ptr("H")}
		variants := map[string]any{
			"Input":        ezclassify.CtxSchemeValidateRequest{Input: ptr("L")},
			"Assertions":   ezclassify.CtxSchemeValidateRequest{Assertions: &stated},
			"Context":      ezclassify.CtxSchemeValidateRequest{Input: ptr("H"), Context: part},
			"InputFormat":  ezclassify.CtxSchemeValidateRequest{Input: ptr("H"), InputFormat: ptr("marking/banner")},
			"OutputFormat": ezclassify.CtxSchemeValidateRequest{Input: ptr("H"), OutputFormat: ptr("marking/portion")},
			"Parts":        ezclassify.CtxSchemeValidateRequest{Input: ptr("H"), Parts: &[]string{"(H)"}},
			"Surface":      ezclassify.CtxSchemeValidateRequest{Input: ptr("H"), Surface: ptr("marking")},
		}
		requireEveryFieldVaried(t, base, variants)
		const path = "/v1/ctx/c/1.0.0/scheme/validate"
		if _, err := cx.Scheme.Validate(ctx, base); err != nil {
			t.Fatal(err)
		}
		for _, name := range sortedKeys(variants) {
			before := calls(path)
			if _, err := cx.Scheme.Validate(ctx, variants[name].(ezclassify.CtxSchemeValidateRequest)); err != nil {
				t.Fatal(err)
			}
			if calls(path) != before+1 {
				t.Errorf("a validate request differing only in %s was answered from the cache", name)
			}
		}
		before := calls(path)
		if _, err := cx.Scheme.Validate(ctx, base); err != nil {
			t.Fatal(err)
		}
		if calls(path) != before {
			t.Error("repeating the base validate request went to the wire")
		}
	})

	t.Run("explain", func(t *testing.T) {
		base := ezclassify.CtxSchemeExplainRequest{Input: "H"}
		variants := map[string]any{
			"Input":        ezclassify.CtxSchemeExplainRequest{Input: "L"},
			"Assertions":   ezclassify.CtxSchemeExplainRequest{Input: "H", Assertions: &stated},
			"Context":      ezclassify.CtxSchemeExplainRequest{Input: "H", Context: part},
			"InputFormat":  ezclassify.CtxSchemeExplainRequest{Input: "H", InputFormat: ptr("marking/banner")},
			"OutputFormat": ezclassify.CtxSchemeExplainRequest{Input: "H", OutputFormat: ptr("marking/portion")},
			"Parts":        ezclassify.CtxSchemeExplainRequest{Input: "H", Parts: &[]string{"(H)"}},
			"Surface":      ezclassify.CtxSchemeExplainRequest{Input: "H", Surface: ptr("marking")},
		}
		requireEveryFieldVaried(t, base, variants)
		const path = "/v1/ctx/c/1.0.0/scheme/explain"
		if _, err := cx.Scheme.Explain(ctx, base); err != nil {
			t.Fatal(err)
		}
		for _, name := range sortedKeys(variants) {
			before := calls(path)
			if _, err := cx.Scheme.Explain(ctx, variants[name].(ezclassify.CtxSchemeExplainRequest)); err != nil {
				t.Fatal(err)
			}
			if calls(path) != before+1 {
				t.Errorf("an explain request differing only in %s was answered from the cache", name)
			}
		}
		before := calls(path)
		if _, err := cx.Scheme.Explain(ctx, base); err != nil {
			t.Fatal(err)
		}
		if calls(path) != before {
			t.Error("repeating the base explain request went to the wire")
		}
	})
}

// Explain and process send the same body to different routes. The operation is
// part of the key, so neither is answered from the other's entry, and each is
// cached on its own.
func TestContextExplainAndProcessKeepSeparateEntries(t *testing.T) {
	srv, calls := countingContextServer(t, func(path string, _ int32) string {
		if strings.HasSuffix(path, "/scheme/explain") {
			return `{"input":"H","schemeId":"s","entries":[],"repairs":[],"validationErrors":[],"inferences":[],"finalAssertions":[],"rendered":"H"}`
		}
		return `{"valid":true,"rendered":"H","inputWasValid":true,"notices":[]}`
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	for i := 0; i < 2; i++ {
		processed, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
		if err != nil || !processed.Result.Valid {
			t.Fatalf("Process = %+v, %v", processed, err)
		}
		explained, err := cx.Scheme.Explain(ctx, ezclassify.CtxSchemeExplainRequest{Input: "H"})
		if err != nil || explained.Result.SchemeID != "s" {
			t.Fatalf("Explain = %+v, %v", explained, err)
		}
	}
	if got := calls("/v1/ctx/c/1.0.0/scheme/process"); got != 1 {
		t.Errorf("process used %d wire requests, want 1", got)
	}
	if got := calls("/v1/ctx/c/1.0.0/scheme/explain"); got != 1 {
		t.Errorf("explain used %d wire requests, want 1", got)
	}
}

// The suggest-for-categories answer is the map itself. An answer that leaves
// out a requested category still reaches the caller who received it, but it is
// not filed for the next one; a complete answer is, and so is the empty map
// that answers a request naming no category.
func TestContextSuggestForCategoriesCachesOnlyCompleteAnswers(t *testing.T) {
	const level = `"level":{"suggestions":[{"conceptId":"L_H","categoryId":"level","label":"H","description":"","group":"","name":"","order":0,"selected":true,"views":{}}],"warnings":[]}`
	const audience = `"audience":{"suggestions":[],"warnings":[]}`
	srv, calls := countingContextServer(t, func(_ string, attempt int32) string {
		switch attempt {
		case 1:
			return `{` + level + `}`
		case 2:
			return `{` + level + `,` + audience + `}`
		default:
			return `{}`
		}
	})
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	const path = "/v1/ctx/c/1.0.0/scheme/suggest-for-categories"
	req := ezclassify.CtxSchemeSuggestForCategoriesRequest{Assertions: []ezclassify.Assertion{}, CategoryIds: []string{"level", "audience"}}

	partial, err := cx.Scheme.SuggestForCategories(ctx, req)
	if err != nil {
		t.Fatalf("partial answer: %v", err)
	}
	if _, ok := partial.Result["audience"]; ok || len(partial.Result["level"].Suggestions) != 1 {
		t.Fatalf("partial answer = %+v, want level alone", partial.Result)
	}
	for i := 0; i < 2; i++ {
		complete, err := cx.Scheme.SuggestForCategories(ctx, req)
		if err != nil {
			t.Fatal(err)
		}
		if _, ok := complete.Result["audience"]; !ok {
			t.Fatalf("complete answer = %+v, want both categories", complete.Result)
		}
	}
	if got := calls(path); got != 2 {
		t.Fatalf("partial then complete answers used %d wire requests, want 2", got)
	}

	none := ezclassify.CtxSchemeSuggestForCategoriesRequest{Assertions: []ezclassify.Assertion{}, CategoryIds: []string{}}
	for i := 0; i < 2; i++ {
		empty, err := cx.Scheme.SuggestForCategories(ctx, none)
		if err != nil || empty.Result == nil || len(empty.Result) != 0 {
			t.Fatalf("no categories = %+v, %v, want an empty map", empty, err)
		}
	}
	if got := calls(path); got != 3 {
		t.Fatalf("the empty answer was not cached: %d wire requests, want 3", got)
	}
}

// The runtime answers a credential that may not see a Context with one constant
// 404. The entries held for the previous caller must go. A 404 that names a
// missing resource says nothing about the caller, so it leaves them in place.
func TestContextCacheClearsOnDenialShapedNotFoundOnly(t *testing.T) {
	var status atomic.Int32
	var body atomic.Value
	var processCalls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.HasSuffix(r.URL.Path, "/scheme/process") {
			processCalls.Add(1)
			_, _ = io.WriteString(w, cachedProcessBody)
			return
		}
		w.WriteHeader(int(status.Load()))
		_, _ = io.WriteString(w, body.Load().(string))
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	process := func() {
		t.Helper()
		if _, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"}); err != nil {
			t.Fatal(err)
		}
	}
	for _, refusal := range []struct {
		name  string
		body  string
		keeps bool
	}{
		{"names a missing resource", `{"error":"concept_not_found","detail":"NOPE"}`, true},
		{"denial-shaped", `{"error":"not_found"}`, false},
		{"denial code with another member", `{"error":"not_found","detail":"x"}`, true},
	} {
		t.Run(refusal.name, func(t *testing.T) {
			process()
			before := processCalls.Load()
			status.Store(http.StatusNotFound)
			body.Store(refusal.body)
			if _, err := cx.Scheme.Concepts(ctx); err == nil {
				t.Fatal("a 404 returned no error")
			}
			process()
			reached := processCalls.Load() != before
			if refusal.keeps && reached {
				t.Error("a 404 that is not the denial cleared the cache")
			}
			if !refusal.keeps && !reached {
				t.Error("the denial-shaped 404 left the previous caller's entries in place")
			}
			cx.ClearCache()
		})
	}
}

// Callers hold their own results. Changing one, whether it came from the wire,
// from the cache or from a request it joined, reaches no other caller and no
// later cached read.
func TestContextCachedResultsAreIndependentCopies(t *testing.T) {
	var calls atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, fixtureAs[ezclassify.ContextSchemeProcessResult](`{"resolution":{"contextId":"c","resolvedVersion":"1.0.0","artifacts":[{"kind":"scheme","id":"s","version":"1.0.0"}]},"result":{"valid":true,"rendered":"H","inputWasValid":true,"assertions":[{"conceptId":"L_H","categoryId":"level","confidence":"certain","source":"parse"}]}}`))
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	first, err := cx.Scheme.Process(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	first.Result.Rendered = "changed"
	first.Result.Assertions[0].ConceptID = "changed"
	first.Resolution.Artifacts[0].ID = "changed"
	second, err := cx.Scheme.Process(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if second.Result.Rendered != "H" || second.Result.Assertions[0].ConceptID != "L_H" || second.Resolution.Artifacts[0].ID != "s" {
		t.Fatalf("a later read saw the first caller's changes: %+v", second.Result)
	}
	if calls.Load() != 1 {
		t.Fatalf("the repeat used %d wire requests, want 1", calls.Load())
	}
}
