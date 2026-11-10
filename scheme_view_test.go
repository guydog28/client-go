package ezclassify_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/ezclassify/client-go"
)

// schemeCall drives one CtxScheme method with a request the client sends
// without refusing it first.
type schemeCall struct {
	method string
	invoke func(ctx context.Context, s *ezclassify.CtxScheme) error
}

func schemeCalls() []schemeCall {
	input := "H"
	return []schemeCall{
		{"Summary", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Summary(ctx); return err }},
		{"Categories", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Categories(ctx); return err }},
		{"Concepts", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Concepts(ctx); return err }},
		{"Formats", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Formats(ctx); return err }},
		{"Grammar", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Grammar(ctx); return err }},
		{"Scopes", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.Scopes(ctx); return err }},
		{"LabelPresets", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.LabelPresets(ctx); return err }},
		{"ApplyPreset", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.ApplyPreset(ctx, ezclassify.ApplyPresetRequest{PresetID: "p"})
			return err
		}},
		{"Process", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: input})
			return err
		}},
		{"Validate", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Validate(ctx, ezclassify.CtxSchemeValidateRequest{Input: &input})
			return err
		}},
		{"Explain", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Explain(ctx, ezclassify.CtxSchemeExplainRequest{Input: input})
			return err
		}},
		{"Suggest", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Suggest(ctx, ezclassify.CtxSchemeSuggestRequest{Input: &input})
			return err
		}},
		{"SuggestForCategories", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.SuggestForCategories(ctx, ezclassify.CtxSchemeSuggestForCategoriesRequest{CategoryIds: []string{"c"}})
			return err
		}},
		{"Aggregate", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Aggregate(ctx, ezclassify.CtxSchemeAggregateRequest{Inputs: &[]string{input}})
			return err
		}},
		{"Scan", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Scan(ctx, ezclassify.CtxSchemeScanRequest{Text: input})
			return err
		}},
		{"Render", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Render(ctx, ezclassify.CtxSchemeRenderRequest{})
			return err
		}},
		{"Normalize", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Normalize(ctx, ezclassify.CtxSchemeNormalizeRequest{Input: input})
			return err
		}},
		{"CheckCoverage", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.CheckCoverage(ctx, ezclassify.CtxSchemeCheckCoverageRequest{})
			return err
		}},
		{"DescribeConcept", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.DescribeConcept(ctx, "C", "c", nil)
			return err
		}},
		{"Diff", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Diff(ctx, ezclassify.CtxSchemeDiffRequest{A: input, B: input})
			return err
		}},
		{"FilterConcepts", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.FilterConcepts(ctx, ezclassify.CtxSchemeFilterConceptsRequest{Query: input})
			return err
		}},
		{"Notices", func(ctx context.Context, s *ezclassify.CtxScheme) error {
			_, err := s.Notices(ctx, ezclassify.CtxSchemeNoticesRequest{})
			return err
		}},
		{"View", func(ctx context.Context, s *ezclassify.CtxScheme) error { _, err := s.View(ctx); return err }},
	}
}

// selectorRecorder answers every request with an empty object and records the
// scheme query parameter each one carried; absent is recorded as nil.
func selectorRecorder(t *testing.T) (*httptest.Server, func() []*string) {
	t.Helper()
	var mu sync.Mutex
	var sent []*string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var selector *string
		if values, ok := r.URL.Query()["scheme"]; ok {
			joined := values[0]
			if len(values) > 1 {
				joined = "<repeated>"
			}
			selector = &joined
		}
		mu.Lock()
		sent = append(sent, selector)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, "{}")
	}))
	t.Cleanup(srv.Close)
	return srv, func() []*string {
		mu.Lock()
		defer mu.Unlock()
		out := sent
		sent = nil
		return out
	}
}

// A new CtxScheme method cannot skip the selector check: the table must name
// every exported method.
func TestSchemeCallsCoverEveryCtxSchemeMethod(t *testing.T) {
	listed := map[string]bool{}
	for _, call := range schemeCalls() {
		listed[call.method] = true
	}
	var missing []string
	viewType := reflect.TypeOf(&ezclassify.CtxScheme{})
	for i := 0; i < viewType.NumMethod(); i++ {
		if name := viewType.Method(i).Name; !listed[name] {
			missing = append(missing, name)
		}
	}
	sort.Strings(missing)
	if len(missing) > 0 {
		t.Fatalf("schemeCalls does not drive: %v", missing)
	}
}

// Every operation and read of a scheme view names the view's scheme on the
// wire, and the primary view names none, so the server answers the primary.
func TestSchemeViewSendsItsSelectorOnEveryCall(t *testing.T) {
	srv, sent := selectorRecorder(t)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	for _, call := range schemeCalls() {
		t.Run(call.method, func(t *testing.T) {
			_ = call.invoke(ctx, cx.SchemeView("second"))
			got := sent()
			if len(got) != 1 || got[0] == nil || *got[0] != "second" {
				t.Fatalf("view call sent scheme %v, want exactly one request naming second", describeSelectors(got))
			}
			_ = call.invoke(ctx, cx.Scheme)
			if got := sent(); len(got) != 1 || got[0] != nil {
				t.Fatalf("primary call sent scheme %v, want one request naming none", describeSelectors(got))
			}
		})
	}
}

func describeSelectors(selectors []*string) []string {
	out := make([]string, len(selectors))
	for i, selector := range selectors {
		if selector == nil {
			out[i] = "<absent>"
		} else {
			out[i] = *selector
		}
	}
	return out
}

// One view per selector: the same selector returns the same view, also when
// asked for concurrently, and a different selector or the primary is another.
func TestSchemeViewIsOnePerSelector(t *testing.T) {
	cx := ezclassify.Connect("http://127.0.0.1:1", "c", "1.0.0")
	const callers = 16
	views := make([]*ezclassify.CtxScheme, callers)
	var wg sync.WaitGroup
	for i := range views {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			views[i] = cx.SchemeView("second")
		}(i)
	}
	wg.Wait()
	for i, view := range views {
		if view == nil || view != views[0] {
			t.Fatalf("SchemeView(second) call %d returned a different view", i)
		}
	}
	if cx.SchemeView("third") == views[0] {
		t.Fatal("two selectors share one view")
	}
	if cx.Scheme == views[0] {
		t.Fatal("a selector's view is the primary view")
	}
	if ezclassify.Connect("http://127.0.0.1:1", "c", "1.0.0").SchemeView("second") == views[0] {
		t.Fatal("two handles share one view")
	}
}

// Cached results are filed under the scheme that answered them. The same
// request through the primary and two views reaches the server once for each
// scheme, and a repeat through a view is answered from that view's entry.
func TestSchemeViewCachesPerScheme(t *testing.T) {
	var mu sync.Mutex
	requests := map[string]int{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		requests[r.URL.Query().Get("scheme")]++
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	for _, view := range []*ezclassify.CtxScheme{cx.Scheme, cx.SchemeView("second"), cx.SchemeView("second"), cx.SchemeView("third"), cx.Scheme} {
		if _, err := view.Process(ctx, req); err != nil {
			t.Fatal(err)
		}
	}
	mu.Lock()
	defer mu.Unlock()
	want := map[string]int{"": 1, "second": 1, "third": 1}
	if !reflect.DeepEqual(requests, want) {
		t.Fatalf("wire requests by scheme = %v, want %v", requests, want)
	}
}

// A selector the Context does not hold is an error. The client neither
// retries without the selector nor answers from the primary scheme's cached
// result for the same request.
func TestSchemeViewUnknownSelectorIsAnError(t *testing.T) {
	var mu sync.Mutex
	var selectors []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		selector := r.URL.Query().Get("scheme")
		mu.Lock()
		selectors = append(selectors, selector)
		mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-EzClassify-Resolved-Version", "1.0.0")
		if selector != "" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = io.WriteString(w, `{"error":"Scheme not found in context: `+selector+`"}`)
			return
		}
		_, _ = io.WriteString(w, cachedProcessBody)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, "c", "1.0.0")
	ctx := context.Background()
	req := ezclassify.CtxSchemeProcessRequest{Input: "H"}
	if _, err := cx.Scheme.Process(ctx, req); err != nil {
		t.Fatal(err)
	}
	got, err := cx.SchemeView("unheld").Process(ctx, req)
	if got != nil {
		t.Fatalf("unknown selector answered %+v", got)
	}
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		t.Fatalf("unknown selector error = %v, want a 404 APIError", err)
	}
	if _, err := cx.SchemeView("unheld").Summary(ctx); !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusNotFound {
		t.Fatalf("unknown selector read error = %v, want a 404 APIError", err)
	}
	mu.Lock()
	defer mu.Unlock()
	if want := []string{"", "unheld", "unheld"}; !reflect.DeepEqual(selectors, want) {
		t.Fatalf("requests named schemes %q, want %q", selectors, want)
	}
}

func TestSchemeViewRefusesAnEmptySelector(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("SchemeView(\"\") did not panic")
		}
	}()
	ezclassify.Connect("http://127.0.0.1:1", "c", "1.0.0").SchemeView("")
}
