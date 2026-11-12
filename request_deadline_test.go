package ezclassify_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ezclassify/sdk-go"
)

// budget is the client timeout these tests configure. A call must end within
// it plus scheduling slack; a call that is still running at failSafe has
// waited without a bound.
const (
	budget   = 200 * time.Millisecond
	slack    = 800 * time.Millisecond
	failSafe = 10 * time.Second
)

// publicCalls are the public operations every deadline test drives: the
// health probe, a discovery call, an uncached Context call and a cached one.
var publicCalls = map[string]func(context.Context, *ezclassify.Client) error{
	"Ping": func(ctx context.Context, c *ezclassify.Client) error {
		_, err := c.Ping(ctx)
		return err
	},
	"ListContexts": func(ctx context.Context, c *ezclassify.Client) error {
		_, err := c.ListContexts(ctx)
		return err
	},
	"Describe": func(ctx context.Context, c *ezclassify.Client) error {
		_, err := ezclassify.ConnectWith(c, "ctx", "1.0.0").Describe(ctx)
		return err
	},
	"Scheme.Process (cached)": func(ctx context.Context, c *ezclassify.Client) error {
		_, err := ezclassify.ConnectWith(c, "ctx", "1.0.0").Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "H"})
		return err
	},
}

// callWithin runs call and returns its error and how long it took. It fails
// the test, rather than hanging the suite, when the call outlives failSafe.
func callWithin(t *testing.T, call func() error) (error, time.Duration) {
	t.Helper()
	done := make(chan error, 1)
	start := time.Now()
	go func() { done <- call() }()
	select {
	case err := <-done:
		return err, time.Since(start)
	case <-time.After(failSafe):
		t.Fatalf("call still waiting after %v: nothing bounded it", failSafe)
		return nil, 0
	}
}

// stallingServer answers every request with a 200 and the first bytes of a
// JSON body, then sends nothing more until the client goes away.
func stallingServer(t *testing.T) *httptest.Server {
	t.Helper()
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, `{"result":`)
		w.(http.Flusher).Flush()
		select {
		case <-r.Context().Done():
		case <-release:
		}
	}))
	t.Cleanup(func() { close(release); srv.Close() })
	return srv
}

// deafTransport never answers and never looks at the request's context, the
// way a hand-written middleware that blocks on a lock or channel behaves.
type deafTransport struct{ stuck chan struct{} }

func (d deafTransport) RoundTrip(*http.Request) (*http.Response, error) {
	<-d.stuck
	return nil, errors.New("released")
}

// deafBody is a response body whose Read blocks until the test ends, ignoring
// both the request context and Close.
type deafBody struct{ stuck chan struct{} }

func (b deafBody) Read([]byte) (int, error) {
	<-b.stuck
	return 0, io.EOF
}

func (deafBody) Close() error { return nil }

// deafBodyTransport answers at once with a 200 whose body never delivers.
type deafBodyTransport struct{ stuck chan struct{} }

func (d deafBodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       deafBody(d),
		Request:    req,
	}, nil
}

func stuck(t *testing.T) chan struct{} {
	t.Helper()
	ch := make(chan struct{})
	t.Cleanup(func() { close(ch) })
	return ch
}

func requireTimedOut(t *testing.T, name string, err error, elapsed time.Duration) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s: succeeded against a server that never finished", name)
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("%s: error %v does not report a timeout (errors.Is context.DeadlineExceeded)", name, err)
	}
	if errors.Is(err, context.Canceled) {
		t.Errorf("%s: timeout reported as a caller cancellation: %v", name, err)
	}
	if elapsed > budget+slack {
		t.Errorf("%s: took %v, want about the %v budget", name, elapsed, budget)
	}
}

func TestEveryCallTimesOutWhenTheBodyStalls(t *testing.T) {
	srv := stallingServer(t)
	for name, call := range publicCalls {
		t.Run(name, func(t *testing.T) {
			c := ezclassify.New(srv.URL, ezclassify.WithTimeout(budget))
			err, elapsed := callWithin(t, func() error { return call(context.Background(), c) })
			requireTimedOut(t, name, err, elapsed)
		})
	}
}

func TestACallerTransportThatIgnoresTheContextStillTimesOut(t *testing.T) {
	for name, call := range publicCalls {
		t.Run(name+"/no answer", func(t *testing.T) {
			supplied := &http.Client{Transport: deafTransport{stuck(t)}}
			c := ezclassify.New("http://ezclassify.invalid", ezclassify.WithHTTPClient(supplied), ezclassify.WithTimeout(budget))
			err, elapsed := callWithin(t, func() error { return call(context.Background(), c) })
			requireTimedOut(t, name, err, elapsed)
		})
		t.Run(name+"/body never ends", func(t *testing.T) {
			supplied := &http.Client{Transport: deafBodyTransport{stuck(t)}}
			c := ezclassify.New("http://ezclassify.invalid", ezclassify.WithHTTPClient(supplied), ezclassify.WithTimeout(budget))
			err, elapsed := callWithin(t, func() error { return call(context.Background(), c) })
			requireTimedOut(t, name, err, elapsed)
		})
	}
}

// stalledTransports are the two ways a request can stall: a real server that
// stops partway through the body, and a caller transport that never answers
// and never watches the request context.
var stalledTransports = map[string]func(*testing.T) (string, *http.Client){
	"stalled body": func(t *testing.T) (string, *http.Client) { return stallingServer(t).URL, &http.Client{} },
	"deaf transport": func(t *testing.T) (string, *http.Client) {
		return "http://ezclassify.invalid", &http.Client{Transport: deafTransport{stuck(t)}}
	},
}

func TestAShorterCallerDeadlineWins(t *testing.T) {
	for transportName, transport := range stalledTransports {
		for name, call := range publicCalls {
			t.Run(transportName+"/"+name, func(t *testing.T) {
				baseURL, supplied := transport(t)
				c := ezclassify.New(baseURL, ezclassify.WithHTTPClient(supplied), ezclassify.WithTimeout(time.Minute))
				ctx, cancel := context.WithTimeout(context.Background(), budget)
				defer cancel()
				err, elapsed := callWithin(t, func() error { return call(ctx, c) })
				requireTimedOut(t, name, err, elapsed)
			})
		}
	}
}

func TestCallerCancellationIsReportedAsCancellation(t *testing.T) {
	for transportName, transport := range stalledTransports {
		for name, call := range publicCalls {
			t.Run(transportName+"/"+name, func(t *testing.T) {
				baseURL, supplied := transport(t)
				c := ezclassify.New(baseURL, ezclassify.WithHTTPClient(supplied), ezclassify.WithTimeout(time.Minute))
				ctx, cancel := context.WithCancel(context.Background())
				timer := time.AfterFunc(budget, cancel)
				defer timer.Stop()
				err, elapsed := callWithin(t, func() error { return call(ctx, c) })
				if !errors.Is(err, context.Canceled) {
					t.Fatalf("error %v does not report the caller's cancellation", err)
				}
				if errors.Is(err, context.DeadlineExceeded) {
					t.Errorf("cancellation reported as a timeout: %v", err)
				}
				if elapsed > budget+slack {
					t.Errorf("took %v after a cancellation at %v", elapsed, budget)
				}
			})
		}
	}
}

// retryingTransport resends a request that failed, as retry middleware handed
// to WithHTTPClient does. It records the deadline each attempt carried.
type retryingTransport struct {
	mu        sync.Mutex
	deadlines []time.Time
	finished  chan struct{}
}

func (r *retryingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	defer close(r.finished)
	var lastErr error
	for range 3 {
		deadline, _ := req.Context().Deadline()
		r.mu.Lock()
		r.deadlines = append(r.deadlines, deadline)
		r.mu.Unlock()
		resp, err := http.DefaultTransport.RoundTrip(req)
		if err == nil {
			return resp, nil
		}
		lastErr = err
	}
	return nil, lastErr
}

func TestARetrySharesTheBudget(t *testing.T) {
	// The first attempt gets no answer; any later attempt is answered at once.
	// A retry with a fresh budget would succeed, so success here means the
	// budget was not shared.
	var attempts atomic.Int32
	release := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if attempts.Add(1) == 1 {
			select {
			case <-r.Context().Done():
			case <-release:
			}
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"1.0.0","engineSha":"sha","gitTag":"tag","openapiVersion":"1.0.0","abiVersion":"v1"}`)
	}))
	t.Cleanup(func() { close(release); srv.Close() })

	retry := &retryingTransport{finished: make(chan struct{})}
	c := ezclassify.New(srv.URL, ezclassify.WithHTTPClient(&http.Client{Transport: retry}), ezclassify.WithTimeout(budget))
	err, elapsed := callWithin(t, func() error {
		_, err := c.Ping(context.Background())
		return err
	})
	requireTimedOut(t, "Ping", err, elapsed)

	// The call returns at the deadline; the middleware's last attempt fails
	// right after it, on the context it already carries.
	select {
	case <-retry.finished:
	case <-time.After(failSafe):
		t.Fatal("retrying transport never returned")
	}
	retry.mu.Lock()
	defer retry.mu.Unlock()
	if len(retry.deadlines) < 2 {
		t.Fatalf("transport made %d attempts, want a retry", len(retry.deadlines))
	}
	for i, deadline := range retry.deadlines {
		if deadline.IsZero() || !deadline.Equal(retry.deadlines[0]) {
			t.Errorf("attempt %d carried deadline %v, want the first attempt's %v", i+1, deadline, retry.deadlines[0])
		}
	}
}

func TestARedirectSharesTheBudget(t *testing.T) {
	// The first hop spends half the budget before redirecting, and the target
	// answers after three quarters of it. A redirect that started a fresh
	// budget would succeed; one that shares it times out.
	target := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			return
		case <-time.After(budget * 3 / 4):
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"1.0.0","engineSha":"sha","gitTag":"tag","openapiVersion":"1.0.0","abiVersion":"v1"}`)
	}))
	t.Cleanup(target.Close)
	first := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(budget / 2)
		http.Redirect(w, r, target.URL+r.URL.Path, http.StatusTemporaryRedirect)
	}))
	t.Cleanup(first.Close)

	c := ezclassify.New(first.URL, ezclassify.WithTimeout(budget))
	err, elapsed := callWithin(t, func() error {
		_, err := c.Ping(context.Background())
		return err
	})
	requireTimedOut(t, "Ping", err, elapsed)
}

// requestGoroutines lists the goroutines running this package's request
// plumbing, so a test can tell that none outlives the call it served.
func requestGoroutines() []string {
	buf := make([]byte, 1<<20)
	buf = buf[:runtime.Stack(buf, true)]
	var found []string
	for _, stack := range strings.Split(string(buf), "\n\n") {
		if strings.Contains(stack, "sdk-go.roundTripWithin") || strings.Contains(stack, "sdk-go.newDeadlineBody") {
			found = append(found, stack)
		}
	}
	return found
}

func TestNothingOutlivesAFinishedCall(t *testing.T) {
	ok := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"1.0.0","engineSha":"sha","gitTag":"tag","openapiVersion":"1.0.0","abiVersion":"v1"}`)
	}))
	t.Cleanup(ok.Close)
	release := make(chan struct{})
	silent := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-release:
		}
	}))
	t.Cleanup(func() { close(release); silent.Close() })
	stalledBody := stallingServer(t)

	for range 5 {
		if _, err := ezclassify.New(ok.URL, ezclassify.WithTimeout(budget)).Ping(context.Background()); err != nil {
			t.Fatal(err)
		}
		for _, stalled := range []string{silent.URL, stalledBody.URL} {
			err, _ := callWithin(t, func() error {
				_, err := ezclassify.New(stalled, ezclassify.WithTimeout(budget)).Ping(context.Background())
				return err
			})
			if err == nil {
				t.Fatal("stalled Ping succeeded")
			}
		}
	}
	deadline := time.Now().Add(2 * time.Second)
	for {
		left := requestGoroutines()
		if len(left) == 0 {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("%d request goroutines outlived their calls; first:\n%s", len(left), left[0])
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// heldBody keeps both I/O methods blocked until the test releases them.
// Closing it must not become an unbounded wait on the caller's deadline path.
type heldBody struct {
	readRelease, closeRelease chan struct{}
	readDone, closeDone       chan struct{}
	closeStarted              chan struct{}
	closes                    atomic.Int32
}

func (b *heldBody) Read([]byte) (int, error) {
	defer close(b.readDone)
	<-b.readRelease
	return 0, io.EOF
}

func (b *heldBody) Close() error {
	if b.closes.Add(1) == 1 {
		close(b.closeStarted)
	}
	<-b.closeRelease
	close(b.closeDone)
	return nil
}

type bodyTransport struct{ body io.ReadCloser }

func (t bodyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       t.body, Request: req,
	}, nil
}

func TestPublicDeadlineDoesNotWaitForBodyClose(t *testing.T) {
	for name, call := range publicCalls {
		t.Run(name, func(t *testing.T) {
			body := &heldBody{
				readRelease: make(chan struct{}), closeRelease: make(chan struct{}),
				readDone: make(chan struct{}), closeDone: make(chan struct{}),
				closeStarted: make(chan struct{}),
			}
			t.Cleanup(func() {
				close(body.readRelease)
				close(body.closeRelease)
				for _, done := range []chan struct{}{body.readDone, body.closeDone} {
					select {
					case <-done:
					case <-time.After(time.Second):
						t.Error("body fixture did not finish after release")
					}
				}
			})
			client := ezclassify.New("http://ezclassify.invalid",
				ezclassify.WithHTTPClient(&http.Client{Transport: bodyTransport{body}}),
				ezclassify.WithTimeout(50*time.Millisecond))
			done := make(chan error, 1)
			go func() { done <- call(context.Background(), client) }()
			select {
			case err := <-done:
				if !errors.Is(err, context.DeadlineExceeded) {
					t.Fatalf("error = %v, want deadline exceeded", err)
				}
			case <-time.After(500 * time.Millisecond):
				t.Fatal("50ms public timeout waited for blocked Read and Close")
			}
			select {
			case <-body.closeStarted:
			case <-time.After(time.Second):
				t.Fatal("timed-out response body was never closed")
			}
			if got := body.closes.Load(); got != 1 {
				t.Errorf("underlying Close called %d times, want once", got)
			}
		})
	}
}
