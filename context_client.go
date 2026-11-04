package ezclassify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/ezclassify/client-go/internal/api"
)

// resolvedVersionHeader carries the concrete version the rest-api resolved a
// context request to, even when the caller asked for "latest" or left the
// version to the Context token.
const resolvedVersionHeader = "X-EzClassify-Resolved-Version"

// latestVersion is the version segment that leaves the choice to the server:
// the Context token's default version under a token, otherwise the newest
// published version of the context.
const latestVersion = "latest"

// EzClassifyContext is a handle bound to one named, versioned EzClassify Context.
// It is named EzClassifyContext rather than Context so it never shadows
// [context.Context] in caller code.
// Every operation routes through /v1/ctx/{contextID}/{version}/*. The path
// selects the composed bundle (scheme + layers + guides + access policy), so
// request bodies carry payload only.
//
// Reach operations through the Scheme, Guide, Access, Document, and Crosswalk groups:
//
//	cx := ezclassify.Connect("http://localhost:3000", "my-ctx", "1.0")
//	res, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: "MARKING"})
//
// An EzClassifyContext is safe for concurrent use; the resolved version is tracked
// behind a mutex.
//
// Repeated identical calls to the cacheable operations are answered from a cache
// owned by the handle, and concurrent identical calls share one request. Every
// call returns its own result: changing a returned value never affects another
// caller or a later read. A 200 response that omits a member the contract
// requires is returned as a [MalformedResponseError] and is never cached. See
// the README's Caching section for the full behavior.
type EzClassifyContext struct {
	client  *Client
	id      string
	version string

	mu              sync.RWMutex
	resolvedVersion string
	cache           *responseCache

	// Scheme groups marking-scheme operations for the bound context.
	Scheme *CtxScheme
	// Guide groups guide operations for the bound context.
	Guide *CtxGuide
	// Access groups access-control operations for the bound context.
	Access *CtxAccess
	// Document groups document-level operations for the bound context.
	Document *CtxDocument
	// Crosswalk groups crosswalk operations for the bound context.
	Crosswalk *CtxCrosswalk
}

// Connect returns an [EzClassifyContext] bound to contextID at the given version.
// Pass an empty version to let the server choose: the Context token's default
// version when the client sends one, otherwise the newest published. The
// literal "latest" means the same. Either binding is unpinned, so its cached
// results expire. The concrete version each call resolves to is reported by
// [EzClassifyContext.ResolvedVersion] once any operation (or
// [EzClassifyContext.Resolve]) has run.
//
// Connect performs no I/O; it only constructs the handle. Use
// [EzClassifyContext.Resolve] to eagerly validate the binding against the server.
// An empty contextID is a programmer error, so Connect panics rather than
// deferring the failure into the first operation.
func Connect(baseURL, contextID, version string) *EzClassifyContext {
	return ConnectWith(New(baseURL), contextID, version)
}

// ConnectWith is like [Connect] but reuses an existing [Client] (and its HTTP
// transport, timeout, and auth middleware) instead of constructing a new one.
//
// A nil client or an empty contextID is a programmer error, so ConnectWith
// panics rather than deferring the failure into the first operation.
// It also panics when access-decision caching is enabled on a latest binding.
func ConnectWith(client *Client, contextID, version string) *EzClassifyContext {
	if client == nil {
		panic("ezclassify: client must not be nil")
	}
	if contextID == "" {
		panic("ezclassify: context id must not be empty")
	}
	if version == "" {
		version = latestVersion
	}
	if client.cacheConfig.cacheAccessDecisions && version == latestVersion && !client.cacheConfig.disableResponseCache {
		panic("ezclassify: access-decision caching requires a concrete context version")
	}
	cx := &EzClassifyContext{client: client, id: contextID, version: version, cache: newResponseCache(defaultResponseCacheEntries)}
	cx.Scheme = &CtxScheme{cx}
	cx.Guide = &CtxGuide{cx}
	cx.Access = &CtxAccess{cx}
	cx.Document = &CtxDocument{cx}
	cx.Crosswalk = &CtxCrosswalk{cx}
	return cx
}

// ID returns the context identifier this handle is bound to.
func (cx *EzClassifyContext) ID() string { return cx.id }

// Version returns the version segment used in request paths: a concrete
// version, or "latest" when the binding leaves the choice to the server.
func (cx *EzClassifyContext) Version() string { return cx.version }

// ResolvedVersion returns the concrete version the rest-api last resolved this
// context to, read from the X-EzClassify-Resolved-Version response header. It
// is empty until the first context operation (or [EzClassifyContext.Resolve]) completes.
func (cx *EzClassifyContext) ResolvedVersion() string {
	cx.mu.RLock()
	defer cx.mu.RUnlock()
	return cx.resolvedVersion
}

// ClearCache removes this Context handle's cached operation results, so the
// next call of each goes to the server. It does not disconnect the handle or
// change its last resolved version. The handle clears itself when an unpinned
// binding's resolved version changes and when the server answers 401 or its
// constant denial-shaped 404.
func (cx *EzClassifyContext) ClearCache() { cx.cache.clear() }

// Resolve validates the context binding against the server and populates
// [EzClassifyContext.ResolvedVersion]. It calls the context describe endpoint.
func (cx *EzClassifyContext) Resolve(ctx context.Context) error {
	_, err := cx.describe(ctx, "resolve context")
	return err
}

// Describe returns the Context's description: the scheme, guides, access policy
// and crosswalks it composes, each with the version it declares and the
// concrete version that answered, plus every scheme it holds with the selector
// that addresses it. Crosswalks is nil when the Context declares none.
//
// Read it for an audit trail: it records exactly which artifact versions a
// Context version, or "latest" at the moment of the call, resolved to. Every
// call goes to the server so the answer is current, and a description that
// shows the composition moved on an unpinned binding clears this handle's
// cached operation results like any other response. Describe also records the
// resolved version, as [EzClassifyContext.Resolve] does.
func (cx *EzClassifyContext) Describe(ctx context.Context) (*ContextDescription, error) {
	return cx.describe(ctx, "describe context")
}

// describe is the one request behind [EzClassifyContext.Describe] and
// [EzClassifyContext.Resolve]; op names the caller in errors.
func (cx *EzClassifyContext) describe(ctx context.Context, op string) (*ContextDescription, error) {
	return do(cx, ctx, op, func(wireCtx context.Context) (callResult[ContextDescription], error) {
		r, err := cx.client.CtxDescribeWithResponse(wireCtx, cx.id, cx.version)
		if err != nil {
			return callResult[ContextDescription]{}, fmt.Errorf("ezclassify: %s %q: %w", op, cx.id, err)
		}
		return callResult[ContextDescription]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

func (cx *EzClassifyContext) captureResolved(resp *http.Response) {
	if resp == nil {
		return
	}
	if v := resp.Header.Get(resolvedVersionHeader); v != "" {
		cx.mu.Lock()
		changed := cx.resolvedVersion != "" && cx.resolvedVersion != v
		cx.resolvedVersion = v
		cx.mu.Unlock()
		if changed {
			cx.cache.clear()
		}
	}
}

// APIError is returned by context operations and [Client.Ping] when the
// rest-api responds with a non-2xx status. It carries the failed operation,
// the HTTP status code, and the raw response body, so callers can branch on
// status and read the API's error payload programmatically instead of
// parsing the message:
//
//	var apiErr *ezclassify.APIError
//	if errors.As(err, &apiErr) && apiErr.StatusCode == http.StatusTooManyRequests {
//	    // the rate limiter answers ahead of every route; apiErr.RetryAfter is
//	    // how long it asks you to wait
//	}
//
// A 404 the server attributes to the context binding is reported as
// [ContextNotFoundError], which wraps an APIError. An unmatched-route 404 or
// a describeConcept concept_not_found refusal stays a plain APIError. An errors.As check for
// *APIError matches every non-2xx response either way. A 2xx that breaks the
// response contract is a [MalformedResponseError]. Transport failures, and an
// error body the client could not decode, are returned as-is.
type APIError struct {
	// Op names the SDK operation that failed, e.g. "scheme.process" or
	// "resolve context".
	Op string
	// StatusCode is the HTTP status the rest-api returned.
	StatusCode int
	// Body is the raw response body, normally the rest-api's JSON error
	// payload.
	Body []byte
	// Response is the decoded structured error body when the server returned a
	// conforming JSON error response. It is nil for non-JSON or malformed bodies.
	Response *ErrorResponse
	// RetryAfter is how long a 429 response asks the caller to wait before
	// sending the request again, read from the Retry-After header and, when
	// the header is absent, from the body's retryAfter. It is zero on every
	// other status and when a 429 states no usable wait. The client never
	// sleeps or retries on its own; the caller chooses whether to wait, and
	// should bound that wait, since the value comes from the server.
	RetryAfter time.Duration
}

// Error formats the failure as
// "ezclassify: <op>: rest-api returned <status>: <body>".
func (e *APIError) Error() string {
	return fmt.Sprintf("ezclassify: %s: rest-api returned %d: %s", e.Op, e.StatusCode, e.Body)
}

// ContextNotFoundError is returned by context operations when the rest-api
// reports a 404, unless it identifies an unmatched route or a describeConcept
// concept_not_found refusal. Those resource errors stay plain [APIError] values.
// The named Context otherwise does not exist or is not visible to the caller.
// Use errors.As to distinguish this from other non-2xx failures. It wraps the
// underlying [APIError], so a single errors.As check for *APIError also matches
// the not-found case.
type ContextNotFoundError struct {
	// ContextID is the identifier that was not found.
	ContextID string
	// APIError carries the underlying HTTP failure: the operation, the 404
	// status, and the raw response body.
	APIError *APIError
}

// InvalidInputError is returned when the Go client refuses a request before
// sending it: a request that names neither or both of its alternative sources,
// or text that cannot be serialized without changing its meaning. Detail is a
// stable classifier such as "invalid-unicode" or "input-or-assertions"; the
// error never includes the rejected value. It is distinct from
// [MalformedResponseError], which reports a bad answer from the server.
type InvalidInputError struct {
	// Op names the SDK operation that rejected the input.
	Op string
	// Detail classifies the invalid input without echoing it.
	Detail string
}

// Error returns a content-free description of the invalid input.
func (e *InvalidInputError) Error() string {
	return fmt.Sprintf("ezclassify: %s: invalid input: %s", e.Op, e.Detail)
}

// Error names the context that the server did not find.
func (e *ContextNotFoundError) Error() string {
	return fmt.Sprintf("ezclassify: context %q not found", e.ContextID)
}

// Unwrap returns the underlying [APIError] so [errors.As] can reach it.
func (e *ContextNotFoundError) Unwrap() error {
	// Return an untyped nil when unset: a typed-nil *APIError inside the error
	// interface would satisfy errors.As and then panic in Error().
	if e.APIError == nil {
		return nil
	}
	return e.APIError
}

// headerOf returns the response headers, or none when there is no response.
func headerOf(resp *http.Response) http.Header {
	if resp == nil {
		return nil
	}
	return resp.Header
}

func (cx *EzClassifyContext) ctxError(op string, status int, header http.Header, body []byte) error {
	// A 2xx can reach here when the generated parser produced no typed value
	// (for example a 200 with an unexpected Content-Type leaves JSON200 nil
	// without an error). That is a decode failure, not an API failure: the
	// APIError contract covers non-2xx responses only.
	if status >= 200 && status < 300 {
		return &MalformedResponseError{Op: op, StatusCode: status}
	}
	apiErr := newAPIError(op, status, header, body)
	if status == http.StatusNotFound && !isRouteNotFound(apiErr) && !isConceptNotFound(apiErr) {
		return &ContextNotFoundError{ContextID: cx.id, APIError: apiErr}
	}
	return apiErr
}

// routeNotFoundCode is the wire contract's classifier for a 404 that reached
// routing and matched no route, as opposed to a route that resolved and found
// no resource.
const routeNotFoundCode = "route_not_found"

// isRouteNotFound reports whether a 404 came from routing rather than from the
// context binding. The contract says to branch on the error code and never on
// the message, so a 404 with no code stays classified as a missing context.
func isRouteNotFound(apiErr *APIError) bool {
	if apiErr == nil || apiErr.Response == nil || apiErr.Response.Code == nil {
		return false
	}
	return *apiErr.Response.Code == routeNotFoundCode
}

// isConceptNotFound uses the server's structured resource classifier. A
// credential or missing Context still carries its original binding error.
func isConceptNotFound(apiErr *APIError) bool {
	return apiErr != nil && apiErr.Op == "scheme.describeConcept" && apiErr.Response != nil && apiErr.Response.Error == "concept_not_found"
}

// callResult is the slice of a generated *...HTTP response that the context
// operations care about, extracted at the call site. The per-call closures own
// the concrete wrapper struct (CtxSchemeProcessHTTP, ...); generics cannot read
// a struct field off a type parameter, so the closure hands the pieces back.
type callResult[T any] struct {
	val        *T
	httpResp   *http.Response
	statusCode int
	body       []byte
}

// do centralizes the capture/nil-check/ctxError chokepoint shared by every
// schema-backed context operation. The closure performs the request with the
// context it is handed and returns the extracted pieces (the typed JSON200
// pointer, the HTTP response, the status code, and the raw body). do then
// records the resolved version, and routes a nil typed value through ctxError:
// an APIError for a non-2xx status, a MalformedResponseError for a 2xx with no
// typed value. A 2xx body the generated decoder refused, and one that lacks a
// member the contract requires, are MalformedResponseErrors too.
//
// A single marking-policy change (new error type, capture behavior, logging)
// touches this one function rather than every method body.
func do[T any](cx *EzClassifyContext, ctx context.Context, op string, call func(context.Context) (callResult[T], error)) (*T, error) {
	wireCtx, status := withResponseStatus(ctx)
	r, err := call(wireCtx)
	if err != nil {
		return nil, decodeFailure(op, status, err)
	}
	cx.captureResolved(r.httpResp)
	if r.val == nil {
		if refusesCallerIdentity(r.statusCode, r.body) {
			cx.cache.clear()
		}
		return nil, cx.ctxError(op, r.statusCode, headerOf(r.httpResp), r.body)
	}
	if err := requireMembers[T](op, r.statusCode, r.body); err != nil {
		return nil, err
	}
	return r.val, nil
}

// denialErrorCode is the error member of the runtime's constant denial-shaped
// 404. It is the answer to a credential that may not see the Context, and it
// carries nothing else, unlike a 404 that names a missing resource.
const denialErrorCode = "not_found"

// refusesCallerIdentity reports whether a failed response says the credential
// behind it is not the caller the held entries were computed for: an HTTP 401,
// or the denial-shaped 404 whose whole body is {"error":"not_found"}.
func refusesCallerIdentity(status int, body []byte) bool {
	switch status {
	case http.StatusUnauthorized:
		return true
	case http.StatusNotFound:
		var fields map[string]any
		return json.Unmarshal(body, &fields) == nil && len(fields) == 1 && fields["error"] == denialErrorCode
	}
	return false
}

// contextResult returns the members of a Context envelope's result object, or
// false when the body is not an envelope that names its resolution and carries
// a result object.
func contextResult(body []byte) (map[string]json.RawMessage, bool) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, false
	}
	var resolution map[string]json.RawMessage
	if json.Unmarshal(fields["resolution"], &resolution) != nil ||
		resolution["contextId"] == nil || resolution["resolvedVersion"] == nil {
		return nil, false
	}
	var result map[string]json.RawMessage
	if json.Unmarshal(fields["result"], &result) != nil || result == nil {
		return nil, false
	}
	return result, true
}

// answersEveryCategory reports whether a suggest-for-categories body holds a
// result object for each requested category. That answer is the map itself, so
// a request naming no category is answered by an empty map, and that is a
// complete answer too.
func answersEveryCategory(categoryIDs []string) func([]byte) bool {
	return func(body []byte) bool {
		byCategory, ok := contextResult(body)
		if !ok {
			return false
		}
		for _, id := range categoryIDs {
			var entry map[string]json.RawMessage
			if json.Unmarshal(byCategory[id], &entry) != nil || entry == nil {
				return false
			}
		}
		return true
	}
}

// cachedDo runs a cached Context operation. A response that fails [do]'s check
// for required members is an error, so it is never filed.
func cachedDo[T any](cx *EzClassifyContext, ctx context.Context, op string, body any, call func(context.Context) (callResult[T], error)) (*T, error) {
	return cachedDoAccepting(cx, ctx, op, body, func([]byte) bool { return true }, call)
}

// cachedDoAccepting is [cachedDo] for an operation that decides for itself
// which well-formed answers are complete enough to file. The key is the
// operation, the binding and the whole request body, so every field that
// changes the answer is part of it.
//
// The cache holds the validated response bytes, not the decoded value. Every
// caller, including the ones that joined an in-flight request, decodes its own
// *T, so nothing a caller does to its result reaches another caller or a later
// cached read.
func cachedDoAccepting[T any](cx *EzClassifyContext, ctx context.Context, op string, body any, cacheable func([]byte) bool, call func(context.Context) (callResult[T], error)) (*T, error) {
	if cx.client.cacheConfig.disableResponseCache {
		return do(cx, ctx, op, call)
	}
	serialized, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("ezclassify: %s cache key: %w", op, err)
	}
	ttl := 5 * time.Minute
	switch op {
	case "scheme.scan":
		ttl = 2 * time.Minute
	case "scheme.describeConcept", "access.rules":
		ttl = time.Hour
	}
	key := op + "|ctx:" + cx.id + "@" + cx.version + "|" + string(serialized)
	value, err := cx.cache.load(ctx, key, ttl, cx.version != latestVersion, func(shared context.Context) (any, bool, error) {
		var responseBody []byte
		_, callErr := do(cx, shared, op, func(wireCtx context.Context) (callResult[T], error) {
			response, requestErr := call(wireCtx)
			responseBody = response.body
			return response, requestErr
		})
		if callErr != nil {
			return nil, false, callErr
		}
		return responseBody, cacheable(responseBody), nil
	})
	if err != nil {
		return nil, err
	}
	var result T
	if err := json.Unmarshal(value.([]byte), &result); err != nil {
		return nil, fmt.Errorf("ezclassify: %s: %w", op, err)
	}
	return &result, nil
}

// rejectInvalidText guards an operation whose request carries caller text that
// the server maps back onto by position. JSON encoding replaces each invalid
// UTF-8 byte with a three-byte replacement character, which would shift every
// returned offset relative to the bytes the caller holds, so the SDK refuses
// the call instead of returning positions that do not line up.
func rejectInvalidText(op, text string) error {
	if !utf8.ValidString(text) {
		return &InvalidInputError{Op: op, Detail: "invalid-unicode"}
	}
	return nil
}

// CtxScheme groups marking-scheme operations bound to an [EzClassifyContext].
type CtxScheme struct{ cx *EzClassifyContext }

// Summary returns high-level metadata for the context's composed scheme: id,
// version, name, description, presentation configuration, default render
// format, and scope label. Pass nil for scheme to read the context's primary
// scheme, or a context-local key to name another one it holds.
func (s *CtxScheme) Summary(ctx context.Context, scheme *string) (*SchemeSummary, error) {
	return do(s.cx, ctx, "scheme.summary", func(wireCtx context.Context) (callResult[SchemeSummary], error) {
		r, err := s.cx.client.CtxSchemeSummaryWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeSummaryParams{Scheme: scheme})
		if err != nil {
			return callResult[SchemeSummary]{}, err
		}
		return callResult[SchemeSummary]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Categories returns the categories declared by the context's composed scheme.
// Pass nil for scheme to read the context's primary scheme, or a context-local
// key to name another one it holds.
func (s *CtxScheme) Categories(ctx context.Context, scheme *string) (*[]CategoryView, error) {
	return do(s.cx, ctx, "scheme.categories", func(wireCtx context.Context) (callResult[[]CategoryView], error) {
		r, err := s.cx.client.CtxSchemeCategoriesWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeCategoriesParams{Scheme: scheme})
		if err != nil {
			return callResult[[]CategoryView]{}, err
		}
		return callResult[[]CategoryView]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Concepts returns the concepts declared by the context's composed scheme.
// Pass nil for scheme to read the context's primary scheme, or a context-local
// key to name another one it holds.
func (s *CtxScheme) Concepts(ctx context.Context, scheme *string) (*[]ConceptView, error) {
	return do(s.cx, ctx, "scheme.concepts", func(wireCtx context.Context) (callResult[[]ConceptView], error) {
		r, err := s.cx.client.CtxSchemeConceptsWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeConceptsParams{Scheme: scheme})
		if err != nil {
			return callResult[[]ConceptView]{}, err
		}
		return callResult[[]ConceptView]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Formats returns the render formats the context's composed scheme supports.
// Pass nil for scheme to read the context's primary scheme, or a context-local
// key to name another one it holds.
func (s *CtxScheme) Formats(ctx context.Context, scheme *string) (*[]FormatView, error) {
	return do(s.cx, ctx, "scheme.formats", func(wireCtx context.Context) (callResult[[]FormatView], error) {
		r, err := s.cx.client.CtxSchemeFormatsWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeFormatsParams{Scheme: scheme})
		if err != nil {
			return callResult[[]FormatView]{}, err
		}
		return callResult[[]FormatView]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Grammar returns the parse grammar of the context's composed scheme. Pass nil
// for scheme to read the context's primary scheme, or a context-local key to
// name another one it holds.
func (s *CtxScheme) Grammar(ctx context.Context, scheme *string) (*SchemeGrammar, error) {
	return do(s.cx, ctx, "scheme.grammar", func(wireCtx context.Context) (callResult[SchemeGrammar], error) {
		r, err := s.cx.client.CtxSchemeGrammarWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeGrammarParams{Scheme: scheme})
		if err != nil {
			return callResult[SchemeGrammar]{}, err
		}
		return callResult[SchemeGrammar]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// LabelPresets lists the label presets declared by the context's primary
// scheme, the scheme [CtxScheme.ApplyPreset] expands them in. The list is empty
// when the scheme declares none. The preset marking strings stay server-side,
// so ApplyPreset is the way to order one.
func (s *CtxScheme) LabelPresets(ctx context.Context) (*[]LabelPresetView, error) {
	return do(s.cx, ctx, "scheme.labelPresets", func(wireCtx context.Context) (callResult[[]LabelPresetView], error) {
		r, err := s.cx.client.CtxSchemeLabelPresetsWithResponse(wireCtx, s.cx.id, s.cx.version)
		if err != nil {
			return callResult[[]LabelPresetView]{}, err
		}
		return callResult[[]LabelPresetView]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Scopes returns the scopes declared by the context's composed scheme. Pass nil
// for scheme to read the context's primary scheme, or a context-local key to
// name another one it holds.
func (s *CtxScheme) Scopes(ctx context.Context, scheme *string) (*[]SchemeScope, error) {
	return do(s.cx, ctx, "scheme.scopes", func(wireCtx context.Context) (callResult[[]SchemeScope], error) {
		r, err := s.cx.client.CtxSchemeScopesWithResponse(wireCtx, s.cx.id, s.cx.version, &api.CtxSchemeScopesParams{Scheme: scheme})
		if err != nil {
			return callResult[[]SchemeScope]{}, err
		}
		return callResult[[]SchemeScope]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// ApplyPreset expands the named label preset against the context's composed
// scheme and returns what [CtxScheme.Validate] answers for the marking it
// produces, in the format the request's OutputFormat names or the scheme's
// default. The preset comes from the context's primary scheme; a preset id it
// does not declare answers 404.
func (s *CtxScheme) ApplyPreset(ctx context.Context, req ApplyPresetRequest) (*ContextSchemeValidateResult, error) {
	return cachedDo(s.cx, ctx, "scheme.applyPreset", req, func(wireCtx context.Context) (callResult[ContextSchemeValidateResult], error) {
		r, err := s.cx.client.CtxSchemeApplyPresetWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeValidateResult]{}, err
		}
		return callResult[ContextSchemeValidateResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Process is the author's call: it settles what was typed in Input together
// with the facts already chosen in Assertions, and answers everything
// [CtxScheme.Validate] answers plus what an author needs to correct and finish
// the marking: the corrections to the text, the proven fix paths, a render for
// each output named in Outputs, and picker options for each category named in
// SuggestFor. Renders is empty unless Outputs names something, and
// SuggestionsFor is nil unless SuggestFor does.
func (s *CtxScheme) Process(ctx context.Context, req CtxSchemeProcessRequest) (*ContextSchemeProcessResult, error) {
	return cachedDo(s.cx, ctx, "scheme.process", req, func(wireCtx context.Context) (callResult[ContextSchemeProcessResult], error) {
		r, err := s.cx.client.CtxSchemeProcessWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeProcessResult]{}, err
		}
		return callResult[ContextSchemeProcessResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Validate is the reader's call: it answers the facts a marking states once
// settled, whether they stand (Valid), whether it was right as written
// (InputWasValid), and the one rendering. Send the marking as text in Input or
// as stated facts in Assertions, exactly one of the two. Neither or both, or
// InputFormat alongside facts, is refused before transport with an
// [InvalidInputError]. Context, Surface and Parts say what the marking is
// judged as.
func (s *CtxScheme) Validate(ctx context.Context, req CtxSchemeValidateRequest) (*ContextSchemeValidateResult, error) {
	if err := exclusiveInputError("scheme.validate", req.Input != nil, req.Assertions != nil, req.InputFormat); err != nil {
		return nil, err
	}
	return cachedDo(s.cx, ctx, "scheme.validate", req, func(wireCtx context.Context) (callResult[ContextSchemeValidateResult], error) {
		r, err := s.cx.client.CtxSchemeValidateWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeValidateResult]{}, err
		}
		return callResult[ContextSchemeValidateResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Explain returns the step-by-step trace of how the engine settles a marking:
// what it read, the repairs and inferences it applied, and the errors left. It
// takes what [CtxScheme.Process] takes apart from the outputs and categories
// that only add to a process result. Explain and Process keep separate cache
// entries even for the same request.
func (s *CtxScheme) Explain(ctx context.Context, req CtxSchemeExplainRequest) (*ContextSchemeExplainResult, error) {
	return cachedDo(s.cx, ctx, "scheme.explain", req, func(wireCtx context.Context) (callResult[ContextSchemeExplainResult], error) {
		r, err := s.cx.client.CtxSchemeExplainWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeExplainResult]{}, err
		}
		return callResult[ContextSchemeExplainResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Suggest returns the concepts a picker may offer for text or stated facts,
// exactly one of Input and Assertions. An empty pointed-to string or slice is
// present input; nil is absent. InputFormat cannot accompany facts. A request
// that breaks either rule is refused before transport with an
// [InvalidInputError]. It says nothing about the marking itself; ask
// [CtxScheme.Validate] or [CtxScheme.Process] for that, or name the categories
// in a process request's SuggestFor to get both in one call.
func (s *CtxScheme) Suggest(ctx context.Context, req CtxSchemeSuggestRequest) (*ContextSchemeSuggestResult, error) {
	if err := exclusiveInputError("scheme.suggest", req.Input != nil, req.Assertions != nil, req.InputFormat); err != nil {
		return nil, err
	}
	return cachedDo(s.cx, ctx, "scheme.suggest", req, func(wireCtx context.Context) (callResult[ContextSchemeSuggestResult], error) {
		r, err := s.cx.client.CtxSchemeSuggestWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeSuggestResult]{}, err
		}
		return callResult[ContextSchemeSuggestResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// SuggestForCategories returns picker options for each requested category,
// keyed by category id. Each category is answered as if its own assertions
// were not yet chosen, so a picker can offer the alternatives to the current
// selection. A request naming no category is answered by an empty map. An
// answer that leaves out a requested category is returned to the caller but
// not cached, so the next call asks the server again.
func (s *CtxScheme) SuggestForCategories(ctx context.Context, req CtxSchemeSuggestForCategoriesRequest) (*ContextSchemeSuggestForCategoriesResult, error) {
	return cachedDoAccepting(s.cx, ctx, "scheme.suggestForCategories", req, answersEveryCategory(req.CategoryIds), func(wireCtx context.Context) (callResult[ContextSchemeSuggestForCategoriesResult], error) {
		r, err := s.cx.client.CtxSchemeSuggestForCategoriesWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeSuggestForCategoriesResult]{}, err
		}
		return callResult[ContextSchemeSuggestForCategoriesResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Aggregate combines multiple markings into a single rolled-up marking. Send
// the markings as text in Inputs or as assertion sets in AssertionSets,
// exactly one of the two; InputFormat applies to text only. A request that
// breaks either rule is refused before transport with an [InvalidInputError].
// Like a process result, the aggregate result carries the collected notices
// for the marking it produced.
func (s *CtxScheme) Aggregate(ctx context.Context, req CtxSchemeAggregateRequest) (*ContextSchemeAggregateResult, error) {
	if err := exclusiveInputError("scheme.aggregate", req.Inputs != nil, req.AssertionSets != nil, req.InputFormat); err != nil {
		return nil, err
	}
	return cachedDo(s.cx, ctx, "scheme.aggregate", req, func(wireCtx context.Context) (callResult[ContextSchemeAggregateResult], error) {
		r, err := s.cx.client.CtxSchemeAggregateWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeAggregateResult]{}, err
		}
		return callResult[ContextSchemeAggregateResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Scan finds markings embedded in free text. Each match carries what
// [CtxScheme.Validate] answers for that marking, read in the role it was found
// in. Text must be valid UTF-8; Scan returns an [InvalidInputError] otherwise,
// because encoding an invalid byte would shift the returned positions away from
// the caller's own text.
func (s *CtxScheme) Scan(ctx context.Context, req CtxSchemeScanRequest) (*ContextSchemeScanResult, error) {
	if err := rejectInvalidText("scheme.scan", req.Text); err != nil {
		return nil, err
	}
	return cachedDo(s.cx, ctx, "scheme.scan", req, func(wireCtx context.Context) (callResult[ContextSchemeScanResult], error) {
		r, err := s.cx.client.CtxSchemeScanWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeScanResult]{}, err
		}
		return callResult[ContextSchemeScanResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Render turns assertions into a marking string.
func (s *CtxScheme) Render(ctx context.Context, req CtxSchemeRenderRequest) (*ContextSchemeRenderResult, error) {
	return cachedDo(s.cx, ctx, "scheme.render", req, func(wireCtx context.Context) (callResult[ContextSchemeRenderResult], error) {
		r, err := s.cx.client.CtxSchemeRenderWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeRenderResult]{}, err
		}
		return callResult[ContextSchemeRenderResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Normalize returns the canonical normalized form of a marking.
func (s *CtxScheme) Normalize(ctx context.Context, req CtxSchemeNormalizeRequest) (*ContextSchemeNormalizeResult, error) {
	return cachedDo(s.cx, ctx, "scheme.normalize", req, func(wireCtx context.Context) (callResult[ContextSchemeNormalizeResult], error) {
		r, err := s.cx.client.CtxSchemeNormalizeWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeNormalizeResult]{}, err
		}
		return callResult[ContextSchemeNormalizeResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// CheckCoverage reports whether a set of child markings is covered by a parent.
func (s *CtxScheme) CheckCoverage(ctx context.Context, req CtxSchemeCheckCoverageRequest) (*ContextSchemeCoverageResult, error) {
	return do(s.cx, ctx, "scheme.checkCoverage", func(wireCtx context.Context) (callResult[ContextSchemeCoverageResult], error) {
		r, err := s.cx.client.CtxSchemeCheckCoverageWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeCoverageResult]{}, err
		}
		return callResult[ContextSchemeCoverageResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// DescribeConcept returns a concept's definition, guidance, and applicable
// rules. view optionally selects the concept property used for the display
// name; pass nil for the default.
func (s *CtxScheme) DescribeConcept(ctx context.Context, conceptID, categoryID string, view *string) (*ContextSchemeConceptDescriptionResult, error) {
	return cachedDo(s.cx, ctx, "scheme.describeConcept", map[string]any{"conceptId": conceptID, "categoryId": categoryID, "view": view}, func(wireCtx context.Context) (callResult[ContextSchemeConceptDescriptionResult], error) {
		params := &api.CtxSchemeDescribeConceptParams{
			ConceptID:  conceptID,
			CategoryID: categoryID,
			View:       view,
		}
		r, err := s.cx.client.CtxSchemeDescribeConceptWithResponse(wireCtx, s.cx.id, s.cx.version, params)
		if err != nil {
			return callResult[ContextSchemeConceptDescriptionResult]{}, err
		}
		return callResult[ContextSchemeConceptDescriptionResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Diff compares two marking strings.
func (s *CtxScheme) Diff(ctx context.Context, req CtxSchemeDiffRequest) (*ContextSchemeDiffResult, error) {
	return cachedDo(s.cx, ctx, "scheme.diff", req, func(wireCtx context.Context) (callResult[ContextSchemeDiffResult], error) {
		r, err := s.cx.client.CtxSchemeDiffWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeDiffResult]{}, err
		}
		return callResult[ContextSchemeDiffResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// FilterConcepts filters the context's concepts by a text query.
func (s *CtxScheme) FilterConcepts(ctx context.Context, req CtxSchemeFilterConceptsRequest) (*ContextSchemeFilterConceptsResult, error) {
	return do(s.cx, ctx, "scheme.filterConcepts", func(wireCtx context.Context) (callResult[ContextSchemeFilterConceptsResult], error) {
		r, err := s.cx.client.CtxSchemeFilterConceptsWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeFilterConceptsResult]{}, err
		}
		return callResult[ContextSchemeFilterConceptsResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Notices collects the concept notices for a set of assertions. Use it when the
// assertions came from somewhere other than a call on this SDK. Every result
// that describes a marking already carries the same list on its Notices field,
// so reading the notices for a marking this context processed takes no second
// call.
func (s *CtxScheme) Notices(ctx context.Context, req CtxSchemeNoticesRequest) (*ContextSchemeNoticesResult, error) {
	return do(s.cx, ctx, "scheme.notices", func(wireCtx context.Context) (callResult[ContextSchemeNoticesResult], error) {
		r, err := s.cx.client.CtxSchemeNoticesWithResponse(wireCtx, s.cx.id, s.cx.version, req)
		if err != nil {
			return callResult[ContextSchemeNoticesResult]{}, err
		}
		return callResult[ContextSchemeNoticesResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// View returns a render-ready structural view of the context's composed scheme.
func (s *CtxScheme) View(ctx context.Context) (*ContextSchemeViewResult, error) {
	return do(s.cx, ctx, "scheme.view", func(wireCtx context.Context) (callResult[ContextSchemeViewResult], error) {
		r, err := s.cx.client.CtxSchemeViewWithResponse(wireCtx, s.cx.id, s.cx.version, CtxSchemeViewRequest{})
		if err != nil {
			return callResult[ContextSchemeViewResult]{}, err
		}
		return callResult[ContextSchemeViewResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// CtxGuide groups guide operations bound to an [EzClassifyContext].
type CtxGuide struct{ cx *EzClassifyContext }

// Scan finds guide rule matches in free text. Text must be valid UTF-8; Scan
// returns an [InvalidInputError] otherwise, because encoding an invalid byte
// would shift the returned positions away from the caller's own text.
func (g *CtxGuide) Scan(ctx context.Context, req CtxGuideScanRequest) (*ContextGuideScanResult, error) {
	if err := rejectInvalidText("guide.scan", req.Text); err != nil {
		return nil, err
	}
	return do(g.cx, ctx, "guide.scan", func(wireCtx context.Context) (callResult[ContextGuideScanResult], error) {
		r, err := g.cx.client.CtxGuideScanWithResponse(wireCtx, g.cx.id, g.cx.version, req)
		if err != nil {
			return callResult[ContextGuideScanResult]{}, err
		}
		return callResult[ContextGuideScanResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Evaluate evaluates a guide with the given rule selection.
func (g *CtxGuide) Evaluate(ctx context.Context, req CtxGuideEvaluateRequest) (*ContextGuideEvaluationResult, error) {
	return do(g.cx, ctx, "guide.evaluate", func(wireCtx context.Context) (callResult[ContextGuideEvaluationResult], error) {
		r, err := g.cx.client.CtxGuideEvaluateWithResponse(wireCtx, g.cx.id, g.cx.version, req)
		if err != nil {
			return callResult[ContextGuideEvaluationResult]{}, err
		}
		return callResult[ContextGuideEvaluationResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// FilterRules filters guide rules by a text query.
func (g *CtxGuide) FilterRules(ctx context.Context, req CtxGuideFilterRulesRequest) (*ContextGuideFilterRulesResult, error) {
	return do(g.cx, ctx, "guide.filterRules", func(wireCtx context.Context) (callResult[ContextGuideFilterRulesResult], error) {
		r, err := g.cx.client.CtxGuideFilterRulesWithResponse(wireCtx, g.cx.id, g.cx.version, req)
		if err != nil {
			return callResult[ContextGuideFilterRulesResult]{}, err
		}
		return callResult[ContextGuideFilterRulesResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// DescribeRule returns a guide rule's name, outcome, and dependencies.
func (g *CtxGuide) DescribeRule(ctx context.Context, req CtxGuideDescribeRuleRequest) (*ContextGuideRuleDescriptionResult, error) {
	return do(g.cx, ctx, "guide.describeRule", func(wireCtx context.Context) (callResult[ContextGuideRuleDescriptionResult], error) {
		r, err := g.cx.client.CtxGuideDescribeRuleWithResponse(wireCtx, g.cx.id, g.cx.version, req)
		if err != nil {
			return callResult[ContextGuideRuleDescriptionResult]{}, err
		}
		return callResult[ContextGuideRuleDescriptionResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// View returns a render-ready structural view of a guide. guideID is optional;
// pass nil to use the first guide in the context.
func (g *CtxGuide) View(ctx context.Context, guideID *string) (*ContextGuideViewResult, error) {
	return do(g.cx, ctx, "guide.view", func(wireCtx context.Context) (callResult[ContextGuideViewResult], error) {
		r, err := g.cx.client.CtxGuideViewWithResponse(wireCtx, g.cx.id, g.cx.version, CtxGuideViewRequest{GuideID: guideID})
		if err != nil {
			return callResult[ContextGuideViewResult]{}, err
		}
		return callResult[ContextGuideViewResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// CtxAccess groups access-control operations bound to an [EzClassifyContext].
type CtxAccess struct{ cx *EzClassifyContext }

// Check evaluates whether a subject may access tagged data. The request carries
// payload only, with no schemeID; the context selects the access policy.
func (a *CtxAccess) Check(ctx context.Context, req CtxAccessCheckRequest) (*ContextAccessDecisionResult, error) {
	call := func(wireCtx context.Context) (callResult[ContextAccessDecisionResult], error) {
		r, err := a.cx.client.CtxAccessCheckWithResponse(wireCtx, a.cx.id, a.cx.version, req)
		if err != nil {
			return callResult[ContextAccessDecisionResult]{}, err
		}
		return callResult[ContextAccessDecisionResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	}
	if a.cx.client.cacheConfig.cacheAccessDecisions {
		return cachedDo(a.cx, ctx, "access.check", req, call)
	}
	return do(a.cx, ctx, "access.check", call)
}

// DescribeRequirements summarizes the access requirements for a marking.
func (a *CtxAccess) DescribeRequirements(ctx context.Context, req CtxAccessDescribeRequirementsRequest) (*ContextAccessRequirementsResult, error) {
	return do(a.cx, ctx, "access.describeRequirements", func(wireCtx context.Context) (callResult[ContextAccessRequirementsResult], error) {
		r, err := a.cx.client.CtxAccessDescribeRequirementsWithResponse(wireCtx, a.cx.id, a.cx.version, req)
		if err != nil {
			return callResult[ContextAccessRequirementsResult]{}, err
		}
		return callResult[ContextAccessRequirementsResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// DescribeRule returns the description of a single access rule.
func (a *CtxAccess) DescribeRule(ctx context.Context, req CtxAccessDescribeRuleRequest) (*ContextAccessRuleDescriptionResult, error) {
	return do(a.cx, ctx, "access.describeRule", func(wireCtx context.Context) (callResult[ContextAccessRuleDescriptionResult], error) {
		r, err := a.cx.client.CtxAccessDescribeRuleWithResponse(wireCtx, a.cx.id, a.cx.version, req)
		if err != nil {
			return callResult[ContextAccessRuleDescriptionResult]{}, err
		}
		return callResult[ContextAccessRuleDescriptionResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// Rules describes every rule in the context's access policy: the
// human-readable predicate, the user-attribute paths, and the data categories
// each rule depends on.
func (a *CtxAccess) Rules(ctx context.Context) (*AccessRulesResponse, error) {
	return cachedDo(a.cx, ctx, "access.rules", nil, func(wireCtx context.Context) (callResult[AccessRulesResponse], error) {
		r, err := a.cx.client.CtxAccessRulesWithResponse(wireCtx, a.cx.id, a.cx.version)
		if err != nil {
			return callResult[AccessRulesResponse]{}, err
		}
		return callResult[AccessRulesResponse]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// CtxCrosswalk groups crosswalk operations bound to an [EzClassifyContext].
type CtxCrosswalk struct{ cx *EzClassifyContext }

// Translate applies the named directional crosswalk to a source marking and
// reports the treatment of every source assertion: what carried, what needs
// review, and what no mapping covers. Send exactly one of the request's
// rendered marking or its structured assertions; neither or both is refused
// before transport with an [InvalidInputError].
func (c *CtxCrosswalk) Translate(ctx context.Context, req CtxCrosswalkTranslateRequest) (*ContextCrosswalkTranslationResult, error) {
	if err := exclusiveInputError("crosswalk.translate", req.Marking != nil, req.Assertions != nil, nil); err != nil {
		return nil, err
	}
	return do(c.cx, ctx, "crosswalk.translate", func(wireCtx context.Context) (callResult[ContextCrosswalkTranslationResult], error) {
		r, err := c.cx.client.CtxCrosswalkTranslateWithResponse(wireCtx, c.cx.id, c.cx.version, req)
		if err != nil {
			return callResult[ContextCrosswalkTranslationResult]{}, err
		}
		return callResult[ContextCrosswalkTranslationResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}

// CtxDocument groups document-level operations bound to an [EzClassifyContext].
type CtxDocument struct{ cx *EzClassifyContext }

// Audit runs a full document audit: whole-marking adequacy, per-part coverage,
// and unmarked-content detection.
func (d *CtxDocument) Audit(ctx context.Context, req CtxDocumentAuditRequest) (*ContextDocumentAuditResult, error) {
	if err := rejectInvalidText("document.audit", req.Text); err != nil {
		return nil, err
	}
	return do(d.cx, ctx, "document.audit", func(wireCtx context.Context) (callResult[ContextDocumentAuditResult], error) {
		r, err := d.cx.client.CtxDocumentAuditWithResponse(wireCtx, d.cx.id, d.cx.version, req)
		if err != nil {
			return callResult[ContextDocumentAuditResult]{}, err
		}
		return callResult[ContextDocumentAuditResult]{r.JSON200, r.HTTPResponse, r.StatusCode(), r.Body}, nil
	})
}
