package ezclassify

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

// TestMajorMinor verifies the major.minor extraction helper across well-formed
// semver, prerelease tags, single-segment strings, and empty input. The helper
// drives the version-skew warning in Ping, so correctness here matters at
// runtime.
func TestMajorMinor(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		// well-formed three-part semver
		{"1.2.3", "1.2"},
		{"0.1.0", "0.1"},
		{"10.0.1", "10.0"},
		// two-part semver, no patch segment
		{"1.2", "1.2"},
		// prerelease tag: extraction stops at the second dot; the suffix is ignored
		{"1.2.3-alpha.1", "1.2"},
		{"0.1.0-rc.2", "0.1"},
		// prerelease with build metadata
		{"0.1.0-rc.2+build.456", "0.1"},
		// only major present, minor defaults to "0"
		{"1", "1.0"},
		// empty or partial input falls back segment-by-segment, so a health
		// payload with no engineVersion compares as a clean "0.0"
		{"", "0.0"},
		{"1.", "1.0"},
		// no dots at all: whole string becomes major, minor defaults to "0"
		{"notver", "notver.0"},
	}
	for _, tc := range cases {
		got := majorMinor(tc.input)
		if got != tc.want {
			t.Errorf("majorMinor(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// TestPingNon2xxReturnsStatusAndBody confirms that a >=400 response from the
// health endpoint surfaces both the status code and the raw response body,
// and that the failure matches *APIError so one errors.As check covers Ping
// alongside every other operation.
func TestPingNon2xxReturnsStatusAndBody(t *testing.T) {
	const errBody = "something went wrong on the server"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = io.WriteString(w, errBody)
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL)
	_, err := c.Ping(context.Background())
	if err == nil {
		t.Fatal("Ping: expected error on 500 response, got nil")
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("Ping error %q does not mention status 500", err.Error())
	}
	if !strings.Contains(err.Error(), errBody) {
		t.Errorf("Ping error %q does not include the response body", err.Error())
	}
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Ping error is %T, want *APIError", err)
	}
	if apiErr.Op != "health" || apiErr.StatusCode != http.StatusInternalServerError {
		t.Errorf("APIError = {Op:%q StatusCode:%d}, want {Op:\"health\" StatusCode:500}", apiErr.Op, apiErr.StatusCode)
	}
	if string(apiErr.Body) != errBody {
		t.Errorf("APIError.Body = %q, want %q", apiErr.Body, errBody)
	}
}

// TestPingNon2xxBodyIsCapped confirms that Ping does not read an unbounded
// error body into memory: a huge non-2xx response is truncated at the
// maxHealthErrorBody cap.
func TestPingNon2xxBodyIsCapped(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		chunk := strings.Repeat("x", 8<<10)
		for written := 0; written < maxHealthErrorBody+64<<10; written += len(chunk) {
			if _, err := io.WriteString(w, chunk); err != nil {
				return
			}
		}
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL)
	_, err := c.Ping(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("Ping error is %T (%v), want *APIError", err, err)
	}
	if len(apiErr.Body) != maxHealthErrorBody {
		t.Errorf("len(Body) = %d, want the %d-byte cap", len(apiErr.Body), maxHealthErrorBody)
	}
}

// TestPingMalformedBodyIsAMalformedResponse confirms that a 200 response with
// a non-JSON body is refused as a *MalformedResponseError rather than silently
// returning a zero-value PingResult.
func TestPingMalformedBodyIsAMalformedResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "not valid json {{")
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL)
	_, err := c.Ping(context.Background())
	if err == nil {
		t.Fatal("Ping: expected decode error on malformed body, got nil")
	}
	var malformed *MalformedResponseError
	if !errors.As(err, &malformed) || malformed.Op != "health" || malformed.StatusCode != http.StatusOK || malformed.Member != "" {
		t.Errorf("Ping error = %T(%v), want a *MalformedResponseError for health naming no member", err, err)
	}
}

// TestPingClosedServerReturnsNetworkError confirms that a transport-level
// failure (connection refused) results in a non-nil error carrying the
// "ezclassify:" wrapper prefix. The server is closed before Ping is called so
// the TCP connection is refused immediately.
func TestPingClosedServerReturnsNetworkError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	// Close before Ping so the dial fails with connection refused.
	srv.Close()

	c := New(srv.URL)
	_, err := c.Ping(context.Background())
	if err == nil {
		t.Fatal("Ping: expected error against closed server, got nil")
	}
	// Transport errors are wrapped as "ezclassify: health request: <cause>".
	if !strings.Contains(err.Error(), "ezclassify:") {
		t.Errorf("Ping error %q is missing the ezclassify wrapper", err.Error())
	}
}

// TestPingRequestBuildError pins that a client whose baseURL cannot be
// embedded into a valid HTTP request URL propagates a "building health request"
// error rather than panicking.
//
// New is bypassed here (it would panic on the same bad URL) by
// constructing the Client struct directly, which is allowed within the package.
func TestPingRequestBuildError(t *testing.T) {
	// "%zz" is not valid percent-encoding; http.NewRequestWithContext rejects
	// it via url.Parse, returning a non-nil error before any network I/O.
	c := &Client{
		httpClient: &http.Client{},
		baseURL:    "http://host/%zz",
	}
	_, err := c.Ping(context.Background())
	if err == nil {
		t.Fatal("Ping: expected error for invalid URL, got nil")
	}
	if !strings.Contains(err.Error(), "building health request") {
		t.Errorf("Ping error %q does not identify the request-build step", err.Error())
	}
}

// TestPingVersionMismatchLogsWarning pins that Ping fires its versionWarnOnce
// log when the wire version differs in major.minor from the compiled-in spec
// version. The redirected logger is process-global, so this test and its
// matching-version counterpart must not run in parallel.
func TestPingVersionMismatchLogsWarning(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// "99.0.0" will never match any plausible clientSpecVersion.
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"99.0.0","engineSha":"","gitTag":"","openapiVersion":"99.0.0","abiVersion":"v1"}`)
	}))
	t.Cleanup(srv.Close)

	logged := captureLog(t)

	c := New(srv.URL)
	result, err := c.Ping(context.Background())
	if err != nil {
		t.Fatalf("Ping: %v", err)
	}
	if result.EngineVersion != "99.0.0" {
		t.Errorf("EngineVersion = %q, want 99.0.0", result.EngineVersion)
	}
	warning := logged.String()
	if !strings.Contains(warning, "Connected to rest-api v99.0.0") {
		t.Errorf("version-skew warning = %q, want it to name the connected version", warning)
	}
	if !strings.Contains(warning, clientSpecVersion) {
		t.Errorf("version-skew warning = %q, want it to name the SDK spec version", warning)
	}

	// The warning is once per client, not once per call.
	logged.Reset()
	if _, err := c.Ping(context.Background()); err != nil {
		t.Fatalf("second Ping: %v", err)
	}
	if logged.Len() != 0 {
		t.Errorf("second Ping logged %q, want the warning only once", logged.String())
	}
}

// captureLog redirects the standard logger into a buffer for the duration of
// the test and restores the default destination afterwards.
func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	t.Cleanup(func() { log.SetOutput(os.Stderr) })
	return &buf
}

// TestPingNon2xxOutside4xxIsAnAPIError pins Ping's documented contract that
// any non-2xx response is an *APIError, not only the 4xx and 5xx range. A 3xx
// reaches here whenever the caller's redirect policy stops the client.
func TestPingNon2xxOutside4xxIsAnAPIError(t *testing.T) {
	for _, status := range []int{http.StatusMultipleChoices, http.StatusFound} {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_, _ = io.WriteString(w, `{"error":"redirected"}`)
		}))

		c := New(srv.URL, WithHTTPClient(&http.Client{
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		}))
		_, err := c.Ping(context.Background())
		var apiErr *APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != status {
			t.Errorf("status %d: error = %T(%v), want *APIError carrying that status", status, err, err)
		}
		srv.Close()
	}
}

// TestResponseCeilingRejectsAnOversizedBody covers the transport-level
// response cap on a success path: the body is well-formed JSON, so nothing
// but the ceiling can reject it.
func TestResponseCeilingRejectsAnOversizedBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		padding := strings.Repeat("p", maxResponseBytes+1)
		_, _ = io.WriteString(w, `{"status":"ok","engineVersion":"1.0.0","engineSha":"","openapiVersion":"1.0.0","abiVersion":"v1","gitTag":"`+padding+`"}`)
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL)
	if _, err := c.Ping(context.Background()); !errors.Is(err, ErrResponseTooLarge) {
		t.Fatalf("Ping error = %v, want one wrapping ErrResponseTooLarge", err)
	}
}

// TestResponseCeilingAcceptsABodyOnTheLimit pins the boundary: a body whose
// length is exactly the ceiling is delivered whole, not truncated.
//
// It runs through ListContexts, whose generated parser reads the body with
// io.ReadAll, which keeps asking the body for more until it reports the end.
//
// The two cases differ in how the body ends. A buffered response carries a
// Content-Length, so its last read returns the final bytes and io.EOF
// together. A streamed response is chunked and has no Content-Length, so the
// last read returns the bytes with no error and the end of the body only
// shows up on the read after it, with the ceiling already spent. Real
// rest-api responses take the streamed path, so both belong here.
func TestResponseCeilingAcceptsABodyOnTheLimit(t *testing.T) {
	const prefix, suffix = `{"contexts":["`, `"]}`
	contextID := strings.Repeat("c", maxResponseBytes-len(prefix)-len(suffix))
	payload := prefix + contextID + suffix
	if len(payload) != maxResponseBytes {
		t.Fatalf("payload is %d bytes, want exactly the %d byte ceiling", len(payload), maxResponseBytes)
	}

	cases := []struct {
		name  string
		write func(w http.ResponseWriter)
	}{
		{
			name:  "content length",
			write: func(w http.ResponseWriter) { _, _ = io.WriteString(w, payload) },
		},
		{
			name: "chunked",
			write: func(w http.ResponseWriter) {
				// Flushing mid-body commits the response to chunked transfer
				// encoding, which drops the Content-Length. The pause after
				// the final flush keeps the terminating chunk off the wire
				// until the client has already consumed the last of the
				// payload, which is what leaves the client holding a spent
				// budget and no end-of-body yet.
				_, _ = io.WriteString(w, payload[:8])
				w.(http.Flusher).Flush()
				time.Sleep(20 * time.Millisecond)
				_, _ = io.WriteString(w, payload[8:])
				w.(http.Flusher).Flush()
				time.Sleep(20 * time.Millisecond)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				tc.write(w)
			}))
			t.Cleanup(srv.Close)

			c := New(srv.URL)
			result, err := c.ListContexts(context.Background())
			if err != nil {
				t.Fatalf("ListContexts: %v", err)
			}
			if len(result.Contexts) != 1 || result.Contexts[0] != contextID {
				t.Errorf("Contexts = %v, want [%q]", result.Contexts, contextID)
			}
		})
	}
}

// TestResponseCeilingCoversGeneratedOperations proves the cap sits in the
// transport, where every generated operation inherits it, rather than in the
// handwritten Ping path alone.
func TestResponseCeilingCoversGeneratedOperations(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"contexts":["`+strings.Repeat("c", maxResponseBytes)+`"]}`)
	}))
	t.Cleanup(srv.Close)

	c := New(srv.URL)
	if _, err := c.ListContexts(context.Background()); !errors.Is(err, ErrResponseTooLarge) {
		t.Fatalf("ListContexts error = %v, want one wrapping ErrResponseTooLarge", err)
	}
}
