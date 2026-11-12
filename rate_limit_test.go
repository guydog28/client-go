package ezclassify_test

// The rate limiter answers ahead of every route, so any operation can return
// 429 with a Retry-After header and a RateLimitError body. These tests read the
// answer's shape from the vendored contract and drive every client call that
// reaches an operation through a stub that returns it.

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ezclassify/sdk-go"
)

// rateLimitExample is the contract's own 429 body: the error text and the wait
// in seconds.
func rateLimitExample(t *testing.T) (body string, seconds int) {
	t.Helper()
	doc, err := specDocument()
	if err != nil {
		t.Fatal(err)
	}
	example, _ := lookup(doc, "components", "responses", "TooManyRequests", "content", "application/json", "example").(map[string]any)
	wait, ok := example["retryAfter"].(float64)
	if !ok || example["error"] != "Too Many Requests" {
		t.Fatalf("the contract's 429 example is %v", example)
	}
	return fmt.Sprintf(`{"error":%q,"retryAfter":%d}`, example["error"], int(wait)), int(wait)
}

func rateLimited(header, body string, status int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if header != "" {
			w.Header().Set("Retry-After", header)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	})
}

func retryAfterOf(t *testing.T, err error) (*ezclassify.APIError, time.Duration) {
	t.Helper()
	var apiErr *ezclassify.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error = %T(%v), want *APIError", err, err)
	}
	return apiErr, apiErr.RetryAfter
}

// The contract states that every operation can answer 429. If it stops saying
// so, the cases below no longer describe the wire.
func TestContractDeclares429OnEveryDrivenOperation(t *testing.T) {
	doc, err := specDocument()
	if err != nil {
		t.Fatal(err)
	}
	paths, _ := doc["paths"].(map[string]any)
	check := func(method, template string) {
		if lookup(paths, template, method, "responses", "429") == nil {
			t.Errorf("the contract does not declare 429 for %s %s", method, template)
		}
	}
	for _, c := range readCases {
		check(c.method, c.template)
	}
	for _, c := range requestCases {
		check("post", requestRoutePrefix+c.route)
	}
}

func TestEveryOperationSurfacesRetryAfterFrom429(t *testing.T) {
	body, seconds := rateLimitExample(t)
	want := time.Duration(seconds) * time.Second
	srv := httptest.NewServer(rateLimited(strconv.Itoa(seconds), body, http.StatusTooManyRequests))
	t.Cleanup(srv.Close)
	client := ezclassify.New(srv.URL)
	cx := ezclassify.ConnectWith(client, requestCaseContext, requestCaseVersion)

	for _, c := range readCases {
		if c.name == "ready" {
			continue // the generated probe call; the SDK wraps no error around it
		}
		t.Run(c.name, func(t *testing.T) {
			apiErr, got := retryAfterOf(t, c.invoke(context.Background(), client, cx))
			if apiErr.StatusCode != http.StatusTooManyRequests || got != want {
				t.Fatalf("status %d, RetryAfter %v; want 429, %v", apiErr.StatusCode, got, want)
			}
			if apiErr.Response == nil || apiErr.Response.Error != "Too Many Requests" {
				t.Fatalf("Response = %+v, want the decoded 429 body", apiErr.Response)
			}
		})
	}
	for _, c := range requestCases {
		t.Run(c.name, func(t *testing.T) {
			_, got := retryAfterOf(t, c.invoke(context.Background(), cx))
			if got != want {
				t.Fatalf("RetryAfter = %v, want %v", got, want)
			}
		})
	}
}

func TestRetryAfterPrefersTheHeaderAndFallsBackToTheBody(t *testing.T) {
	for _, tc := range []struct {
		name, header, body string
		want               time.Duration
	}{
		{"header wins over the body", "3", `{"error":"Too Many Requests","retryAfter":9}`, 3 * time.Second},
		{"body answers when the header is absent", "", `{"error":"Too Many Requests","retryAfter":9}`, 9 * time.Second},
		{"fractional body seconds", "", `{"error":"Too Many Requests","retryAfter":1.5}`, 1500 * time.Millisecond},
		{"neither source states a wait", "", `{"error":"Too Many Requests"}`, 0},
		{"a negative wait reads as zero", "-5", `{"error":"Too Many Requests"}`, 0},
		{"an enormous wait saturates instead of wrapping", "9000000000000000000", `{"error":"Too Many Requests"}`, time.Duration(math.MaxInt64)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(rateLimited(tc.header, tc.body, http.StatusTooManyRequests))
			t.Cleanup(srv.Close)
			_, err := ezclassify.New(srv.URL).ListContexts(context.Background())
			if _, got := retryAfterOf(t, err); got != tc.want {
				t.Fatalf("RetryAfter = %v, want %v", got, tc.want)
			}
		})
	}
}

// Ping reads its answer directly, so an intermediary's 429 that does not follow
// the contract still reaches the caller as a rate-limit error with the wait the
// body states.
func TestPingReadsRetryAfterFromAnIntermediarys429(t *testing.T) {
	for _, tc := range []struct {
		name, header, body string
		want               time.Duration
	}{
		{"a date in the header falls back to the body", "Wed, 21 Oct 2015 07:28:00 GMT", `{"error":"Too Many Requests","retryAfter":4}`, 4 * time.Second},
		{"a body that is not JSON states no wait", "", `slow down`, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(rateLimited(tc.header, tc.body, http.StatusTooManyRequests))
			t.Cleanup(srv.Close)
			_, err := ezclassify.New(srv.URL).Ping(context.Background())
			if _, got := retryAfterOf(t, err); got != tc.want {
				t.Fatalf("RetryAfter = %v, want %v", got, tc.want)
			}
		})
	}
}

// Retry-After on another status is not the rate limiter's wait: the field
// describes the 429 the contract declares and nothing else.
func TestRetryAfterIsZeroOffA429(t *testing.T) {
	for _, status := range []int{http.StatusBadRequest, http.StatusServiceUnavailable} {
		srv := httptest.NewServer(rateLimited("30", `{"error":"x","retryAfter":30}`, status))
		_, err := ezclassify.New(srv.URL).ListContexts(context.Background())
		srv.Close()
		if apiErr, got := retryAfterOf(t, err); got != 0 || apiErr.StatusCode != status {
			t.Errorf("status %d: RetryAfter %v", status, got)
		}
	}
}

// A 429 is not an answer: the next identical call must reach the server again
// instead of replaying the refusal from the cache.
func TestRateLimitedAnswerIsNotCached(t *testing.T) {
	var mu sync.Mutex
	hits := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		hits++
		mu.Unlock()
		rateLimited("1", `{"error":"Too Many Requests","retryAfter":1}`, http.StatusTooManyRequests).ServeHTTP(w, r)
	}))
	t.Cleanup(srv.Close)
	cx := ezclassify.Connect(srv.URL, requestCaseContext, requestCaseVersion)
	for range 2 {
		if _, err := cx.Scheme.Process(context.Background(), ezclassify.CtxSchemeProcessRequest{Input: "H"}); err == nil {
			t.Fatal("expected a rate-limit error")
		}
	}
	mu.Lock()
	defer mu.Unlock()
	if hits != 2 {
		t.Fatalf("server saw %d requests, want 2", hits)
	}
}
