// Hand-written wrapper layer: constructors, Ping, and the spec-version
// mismatch warning. The flat operation methods come from the embedded
// generated client in internal/api. The package doc lives in doc.go.

package ezclassify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ezclassify/sdk-go/internal/api"
)

// PingResult carries the live version information returned by [Client.Ping].
type PingResult struct {
	// EngineVersion is the engine version currently running on the rest-api.
	EngineVersion string
	// EngineSHA is the git commit SHA of the deployed rest-api build
	// (empty string when unavailable).
	EngineSHA string
	// GitTag is the git tag of the deployed rest-api build
	// (empty string when unavailable).
	GitTag string
}

// RestClientCredentials carries caller identity and Context authorization.
// AccessToken is sent as a bearer token on every REST request. ContextToken is
// sent only to /v1/ctx and its descendants.
//
// Credentials are attached only to requests whose scheme and host match the
// client's base URL. A redirect to another origin is passed through without them.
//
// The zero value is valid and configures an unauthenticated client. A supplied
// token must contain at least one non-whitespace character and no control
// characters; the constructors panic on a token that fails either rule.
type RestClientCredentials struct {
	// AccessToken identifies the caller and is sent as a bearer token on every request.
	AccessToken string
	// ContextToken authorizes Context operations and is sent only on /v1/ctx routes.
	ContextToken string
}

const contextTokenHeader = "x-ezclassify-context-token"
const applicationIDHeader = "x-ezclassify-application-id"

func (c RestClientCredentials) validate() error {
	if c.AccessToken != "" && strings.TrimSpace(c.AccessToken) == "" {
		return fmt.Errorf("access token must be empty or non-blank")
	}
	if c.ContextToken != "" && strings.TrimSpace(c.ContextToken) == "" {
		return fmt.Errorf("context token must be empty or non-blank")
	}
	// A token read from a file or environment variable often arrives with a
	// stray newline. Catching it here names the credential; letting it through
	// turns every later call into an opaque transport error.
	if !validHeaderValue(c.AccessToken) {
		return fmt.Errorf("access token must not contain control characters")
	}
	if !validHeaderValue(c.ContextToken) {
		return fmt.Errorf("context token must not contain control characters")
	}
	return nil
}

// validHeaderValue reports whether every byte of v is legal in an HTTP header
// field value: printable ASCII, horizontal tab, or a high byte, which is what
// net/http accepts when it writes the header.
func validHeaderValue(v string) bool {
	for i := 0; i < len(v); i++ {
		b := v[i]
		if b < 0x20 && b != '\t' {
			return false
		}
		if b == 0x7f {
			return false
		}
	}
	return true
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

// roundTripWithin sends req through rt and returns no later than the end of
// req's context, even when rt, or the body of the response it returns, never
// looks at that context. net/http enforces Client.Timeout and the caller's
// deadline by ending the request context and trusts the transport to notice;
// a caller-supplied transport that blocks without watching it would otherwise
// hold the call past every deadline.
//
// The goroutine sending the request ends when rt returns. If the call has
// already given up by then, it closes the response nobody will read.
func roundTripWithin(rt http.RoundTripper, req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	type outcome struct {
		resp *http.Response
		err  error
	}
	// Unbuffered, so a response is either received here or, once the context
	// has ended and nobody is receiving, closed by the sender.
	arrived := make(chan outcome)
	go func() {
		resp, err := rt.RoundTrip(req)
		select {
		case arrived <- outcome{resp, err}:
		case <-ctx.Done():
			if resp != nil && resp.Body != nil {
				_ = resp.Body.Close()
			}
		}
	}()
	select {
	case result := <-arrived:
		if result.resp != nil && result.resp.Body != nil && result.resp.Body != http.NoBody {
			result.resp.Body = newDeadlineBody(ctx, result.resp.Body)
		}
		return result.resp, result.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// deadlineBody serves a response body until the request context ends, then
// fails every read with the context's error, even when the underlying body
// ignores both the context and Close. A pump goroutine copies the body into a
// pipe; ending the context closes the pipe, which releases a blocked Read at
// once. The pump never writes into a caller's buffer, so an abandoned read
// cannot race a later one.
type deadlineBody struct {
	pipe      *io.PipeReader
	stop      func() bool
	closeBody func() error
}

func newDeadlineBody(ctx context.Context, body io.ReadCloser) io.ReadCloser {
	pipe, pump := io.Pipe()
	// Close can block just as Read can. Keep its one invocation outside the
	// caller's goroutine, while preserving its result until the context ends.
	closed := make(chan struct{})
	var closeErr error
	startClose := sync.OnceFunc(func() {
		go func() {
			closeErr = body.Close()
			close(closed)
		}()
	})
	closeBody := func() error {
		startClose()
		select {
		case <-closed:
			return closeErr
		case <-ctx.Done():
			// net/http also cancels its timer context after a completed
			// response. Keep a Close result already obtained in that case.
			select {
			case <-closed:
				return closeErr
			default:
				return ctx.Err()
			}
		}
	}
	go func() {
		_, err := io.Copy(pump, body)
		_ = pump.CloseWithError(err)
	}()
	// Ending the write side hands the reader the context's error. Closing the
	// body as well frees the pump from a body that honors Close but not the
	// context.
	stop := context.AfterFunc(ctx, func() {
		_ = pump.CloseWithError(ctx.Err())
		_ = closeBody()
	})
	return &deadlineBody{pipe: pipe, stop: stop, closeBody: closeBody}
}

func (b *deadlineBody) Read(p []byte) (int, error) { return b.pipe.Read(p) }

func (b *deadlineBody) Close() error {
	b.stop()
	_ = b.pipe.Close()
	return b.closeBody()
}

// sameOrigin builds the origin test that gates credential attachment. Scheme
// and host must both match the configured base URL; the comparison is on the
// URL's Host, so an explicit default port counts as a distinct origin the same
// way the standard library treats it.
func sameOrigin(base *url.URL) func(*url.URL) bool {
	scheme := strings.ToLower(base.Scheme)
	host := strings.ToLower(base.Host)
	return func(target *url.URL) bool {
		if target == nil {
			return false
		}
		return strings.EqualFold(target.Scheme, scheme) && strings.EqualFold(target.Host, host)
	}
}

// isContextPath reports whether an outgoing request targets a Context route.
// routePrefix is the base URL's own path, which the request path carries when
// the REST API is mounted under a sub-path behind an ingress or proxy.
func isContextPath(routePrefix, path string) bool {
	root := routePrefix + "/v1/ctx"
	return path == root || strings.HasPrefix(path, root+"/")
}

// Client is the EzClassify Go SDK deployment client. Use [ConnectWith] to bind
// a published Context and access its grouped operations.
type Client struct {
	wire            *api.ClientWithResponses
	baseURL         string
	httpClient      *http.Client
	versionWarnOnce sync.Once
	cacheConfig     clientConfig
}

// New creates a Client pointing at the given REST API base URL. The URL should
// not have a trailing slash. Every request is bounded by a 30 second timeout
// unless [WithTimeout] or [WithHTTPClient] sets another bound. Options supply
// credentials, the HTTP client, the timeout, and response handling; see
// [Option].
//
// baseURL is a programmer-supplied constant in normal use, so a malformed value
// is treated as a programming error: New panics rather than returning an
// error. Validate untrusted input with [net/url.ParseRequestURI] before
// passing it here.
//
// New never mutates the client given to [WithHTTPClient]; it sends requests
// through a copy that keeps the supplied timeout, redirect policy, cookie jar,
// and transport, and adds the route-scoped credentials and the response
// ceiling.
func New(baseURL string, opts ...Option) *Client {
	config := resolveOptions(opts)
	credentials := config.credentials
	// The generated client accepts any string as a server URL, so this is the
	// only place the documented panic-on-malformed-baseURL contract can hold.
	parsedBaseURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		panic("ezclassify: malformed base url: " + err.Error())
	}
	contextRoot := strings.TrimRight(parsedBaseURL.Path, "/")
	credentialOrigin := sameOrigin(parsedBaseURL)
	var configuredHTTPClient http.Client
	if config.httpClient != nil {
		configuredHTTPClient = *config.httpClient
	}
	// A zero or negative Timeout means no timeout to net/http, so only an
	// explicit positive bound survives. Applying the default to the copy keeps
	// the caller's client untouched.
	switch {
	case config.timeout > 0:
		configuredHTTPClient.Timeout = config.timeout
	case configuredHTTPClient.Timeout <= 0:
		configuredHTTPClient.Timeout = defaultTimeout
	}
	baseTransport := configuredHTTPClient.Transport
	if baseTransport == nil {
		baseTransport = http.DefaultTransport
	}
	capResponse := func(resp *http.Response, err error) (*http.Response, error) {
		if err != nil || resp == nil {
			return resp, err
		}
		recordResponseStatus(resp.Request.Context(), resp.StatusCode)
		if resp.Body == nil {
			return resp, nil
		}
		resp.Body = newLimitedBody(resp.Body, maxResponseBytes)
		return resp, nil
	}
	configuredHTTPClient.Transport = roundTripperFunc(func(req *http.Request) (*http.Response, error) {
		// Clone before mutating: RoundTrip must not modify the caller's
		// request. Clone already deep-copies the header map.
		req = req.Clone(req.Context())
		// A redirect sends the transport round again with a new URL. The
		// standard library strips sensitive headers before following one to
		// another origin, so re-adding credentials off-origin would undo that.
		if !credentialOrigin(req.URL) {
			return capResponse(roundTripWithin(baseTransport, req))
		}
		if credentials.AccessToken != "" {
			req.Header.Set("Authorization", "Bearer "+credentials.AccessToken)
		}
		if isContextPath(contextRoot, req.URL.Path) && credentials.ContextToken != "" {
			req.Header.Set(contextTokenHeader, credentials.ContextToken)
		}
		if isContextPath(contextRoot, req.URL.Path) && config.applicationID != "" {
			req.Header.Set(applicationIDHeader, config.applicationID)
		}
		return capResponse(roundTripWithin(baseTransport, req))
	})

	apiClient, err := api.NewClientWithResponses(baseURL, api.WithHTTPClient(&configuredHTTPClient))
	if err != nil {
		// Only a failing client option can reach here; none of ours fail.
		panic("ezclassify: " + err.Error())
	}
	return &Client{
		wire:        apiClient,
		baseURL:     strings.TrimRight(baseURL, "/"),
		httpClient:  &configuredHTTPClient,
		cacheConfig: config,
	}
}

// EngineSpecVersion returns the engine spec version this client was built
// against. This is a compile-time constant set at release; its major.minor is
// the rest-api line this client is intended to communicate with.
func (c *Client) EngineSpecVersion() string {
	return clientSpecVersion
}

// ListContexts returns the ids of the Contexts the caller may use on this
// deployment. A non-2xx response is an *[APIError]; a 2xx body that lacks a
// member the contract requires is a *[MalformedResponseError].
func (c *Client) ListContexts(ctx context.Context) (*ContextListResponse, error) {
	ctx, status := withResponseStatus(ctx)
	resp, err := c.wire.CtxListWithResponse(ctx, nil)
	if err != nil {
		return nil, decodeFailure("list contexts", status, fmt.Errorf("ezclassify: list contexts: %w", err))
	}
	return successResult("list contexts", resp.HTTPResponse, resp.StatusCode(), resp.Body, resp.JSON200)
}

// ListContextVersions returns the published versions of a Context, newest
// first as ordered by the server. Errors follow [Client.ListContexts].
func (c *Client) ListContextVersions(ctx context.Context, contextID string) (*ContextVersionsResponse, error) {
	ctx, status := withResponseStatus(ctx)
	resp, err := c.wire.CtxVersionsWithResponse(ctx, contextID, nil)
	if err != nil {
		return nil, decodeFailure("list context versions", status, fmt.Errorf("ezclassify: list context versions: %w", err))
	}
	return successResult("list context versions", resp.HTTPResponse, resp.StatusCode(), resp.Body, resp.JSON200)
}

// maxHealthErrorBody caps how much of a failed health response body Ping
// reads into the returned error, so a misconfigured baseURL pointing at a
// host that streams an enormous error page cannot balloon memory or logs.
const maxHealthErrorBody = 64 << 10

// Ping calls GET /v1/health on the connected rest-api and returns a
// [PingResult] containing the live engine version, git SHA, and git tag.
//
// A non-2xx response is returned as an *[APIError], matching the error
// taxonomy of every other operation, with the body capped at 64 KiB. A
// success body is bounded by the client's response ceiling like any other
// operation's. A 2xx body that cannot be decoded, or that lacks a member the
// contract requires, is a *[MalformedResponseError].
//
// On the first call, if the rest-api major.minor version differs from
// [EngineSpecVersion], a warning is logged via the standard [log] package.
func (c *Client) Ping(ctx context.Context) (PingResult, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/v1/health", nil)
	if err != nil {
		return PingResult{}, fmt.Errorf("ezclassify: building health request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PingResult{}, fmt.Errorf("ezclassify: health request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// A read failure here only truncates the diagnostic body; the call is
		// already failing with the HTTP status, so the partial body is kept.
		body, _ := io.ReadAll(io.LimitReader(resp.Body, maxHealthErrorBody))
		return PingResult{}, newAPIError("health", resp.StatusCode, resp.Header, body)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PingResult{}, fmt.Errorf("ezclassify: health response: %w", err)
	}
	var payload api.HealthResponse
	if json.Unmarshal(body, &payload) != nil {
		return PingResult{}, &MalformedResponseError{Op: "health", StatusCode: resp.StatusCode}
	}
	if err := requireMembers[api.HealthResponse]("health", resp.StatusCode, body); err != nil {
		return PingResult{}, err
	}

	result := PingResult{
		EngineVersion: payload.EngineVersion,
		EngineSHA:     payload.EngineSHA,
		GitTag:        payload.GitTag,
	}

	c.versionWarnOnce.Do(func() {
		if majorMinor(clientSpecVersion) != majorMinor(result.EngineVersion) {
			engineVersion := result.EngineVersion
			if engineVersion == "" {
				engineVersion = "(unreported)"
			}
			log.Printf(
				"ezclassify-go v%s targets rest-api %s.x. "+
					"Connected to rest-api v%s. Mixed minor versions are unsupported.",
				clientSpecVersion, majorMinor(clientSpecVersion), engineVersion,
			)
		}
	})

	return result, nil
}

func newAPIError(op string, statusCode int, header http.Header, body []byte) *APIError {
	apiErr := &APIError{Op: op, StatusCode: statusCode, Body: body}
	var response ErrorResponse
	if err := json.Unmarshal(body, &response); err == nil && response.Error != "" {
		apiErr.Response = &response
	}
	if statusCode == http.StatusTooManyRequests {
		apiErr.RetryAfter = retryAfter(header, body)
	}
	return apiErr
}

// retryAfter reads how long a 429 asks the caller to wait. The Retry-After
// header, whole seconds, is the authority; the body's retryAfter carries the
// same value and answers when a proxy has dropped the header. A missing,
// malformed, or negative value reads as zero, and an enormous one saturates at
// the largest representable duration instead of wrapping negative.
func retryAfter(header http.Header, body []byte) time.Duration {
	if seconds, err := strconv.ParseInt(strings.TrimSpace(header.Get("Retry-After")), 10, 64); err == nil {
		return secondsToDuration(float64(seconds))
	}
	var limit api.RateLimitError
	if err := json.Unmarshal(body, &limit); err != nil {
		return 0
	}
	return secondsToDuration(float64(limit.RetryAfter))
}

func secondsToDuration(seconds float64) time.Duration {
	switch {
	case math.IsNaN(seconds) || seconds <= 0:
		return 0
	case seconds >= float64(math.MaxInt64)/float64(time.Second):
		return time.Duration(math.MaxInt64)
	}
	return time.Duration(seconds * float64(time.Second))
}

func majorMinor(v string) string {
	parts := strings.SplitN(v, ".", 3)
	maj := "0"
	min := "0"
	// strings.SplitN never returns an empty slice, so the defaults survive
	// only when a segment is absent or empty (e.g. "" or "1.").
	if parts[0] != "" {
		maj = parts[0]
	}
	if len(parts) > 1 && parts[1] != "" {
		min = parts[1]
	}
	return maj + "." + min
}
