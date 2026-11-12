// Client options: the variadic seam every constructor accepts. The package
// doc lives in doc.go.

package ezclassify

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"
)

// maxResponseBytes is the ceiling a [Client] puts on a single response
// body. It is sized for the largest payload the engine produces in normal use (a full
// document assessment over a large document), with room to spare, and exists
// so a misconfigured base URL or a compromised endpoint cannot drive
// unbounded allocation in the caller's process.
const maxResponseBytes = 32 << 20

// ErrResponseTooLarge is returned by an operation whose response body exceeds
// the client's response ceiling. It is never returned for a body within the
// ceiling, so callers can test for it with [errors.Is] rather than guess at a
// transport failure.
var ErrResponseTooLarge = errors.New("ezclassify: response body exceeds the client response limit")

// Option configures a [Client] at construction. [New] accepts a variadic list
// of them; when two options set the same value, the later one wins.
type Option func(*clientConfig)

// clientConfig holds the resolved option values. It is deliberately unexported:
// the option functions are the public contract, not the struct.
type clientConfig struct {
	credentials          RestClientCredentials
	httpClient           *http.Client
	timeout              time.Duration
	disableResponseCache bool
	applicationID        string
}

func resolveOptions(opts []Option) clientConfig {
	config := clientConfig{}
	for _, opt := range opts {
		if opt != nil {
			opt(&config)
		}
	}
	return config
}

// defaultTimeout bounds every request of a [Client] whose timeout the caller
// did not set. A request must never wait indefinitely on a server that accepts
// the connection and then stays silent.
const defaultTimeout = 30 * time.Second

// WithCredentials sets caller and Context credentials or one managed application
// credential; see [RestClientCredentials] for route scoping and mutually exclusive
// modes. The zero value sends no credentials, which suits an unauthenticated
// local runtime.
//
// WithCredentials panics on mixed credential modes, a blank non-empty value,
// or a value containing control characters. These are programming errors.
func WithCredentials(credentials RestClientCredentials) Option {
	if err := credentials.validate(); err != nil {
		panic("ezclassify: " + err.Error())
	}
	return func(config *clientConfig) { config.credentials = credentials }
}

// WithTimeout bounds each call to d: connecting, waiting for the answer,
// reading the whole response body, the redirects the client follows and the
// retries a supplied transport makes, which all share the one deadline. It
// holds even when a transport supplied with [WithHTTPClient] ignores the
// request context. A call that runs out of time fails with an error matching
// [context.DeadlineExceeded]; a shorter deadline on the call's
// [context.Context] wins. WithTimeout overrides the Timeout of a client
// supplied with [WithHTTPClient], whichever option comes first. Without it a
// client allows 30 seconds.
//
// There is no way to wait indefinitely: WithTimeout panics on a duration of
// zero or less, which is a programming error. Pass a long duration for slow
// operations.
func WithTimeout(d time.Duration) Option {
	if d <= 0 {
		panic("ezclassify: timeout must be positive")
	}
	return func(config *clientConfig) { config.timeout = d }
}

// WithHTTPClient makes the [Client] send requests through a copy of httpClient,
// for callers that need a custom transport, retry middleware, cookie jar, or
// redirect policy. The copy keeps the supplied client's Transport, Jar, and
// CheckRedirect; EzClassify's credential and response-size handling wraps the
// transport. The caller's http.Client is never mutated.
//
// A supplied client whose Timeout is zero or negative would wait forever, so
// the copy gets the 30 second default instead. Set Timeout on the supplied
// client, or pass [WithTimeout], to choose another bound. The bound holds
// whatever the supplied transport does: a call returns at its deadline even if
// the transport, or the response body it returns, never watches the request
// context. At the deadline the SDK closes that body, possibly while a read of
// it is still in progress, as net/http does with its own bodies.
//
// WithHTTPClient panics on a nil client, which is a programming error.
func WithHTTPClient(httpClient *http.Client) Option {
	if httpClient == nil {
		panic("ezclassify: http client must not be nil")
	}
	return func(config *clientConfig) { config.httpClient = httpClient }
}

// WithResponseCacheDisabled makes every Context operation call the server, with
// no caching and no sharing of concurrent identical calls.
func WithResponseCacheDisabled() Option {
	return func(config *clientConfig) { config.disableResponseCache = true }
}

// WithApplicationID names the calling application on every Context request,
// as the header x-ezclassify-application-id. The runtime counts requests per
// application id and reports the usage to Studio, so use the id the application
// is registered under. The id is not a credential and authorizes nothing: a
// request without it, or with an id the deployment has not registered, is
// served the same way. It is sent only to Context routes on the configured
// origin.
//
// WithApplicationID panics on a blank id, an id with leading or trailing
// whitespace, or an id containing control characters. These are programming
// errors.
func WithApplicationID(id string) Option {
	if strings.TrimSpace(id) == "" {
		panic("ezclassify: application id must not be blank")
	}
	if id != strings.TrimSpace(id) {
		panic("ezclassify: application id must not have leading or trailing whitespace")
	}
	if strings.ContainsRune(id, '\t') || !validHeaderValue(id) {
		panic("ezclassify: application id must not contain control characters")
	}
	return func(config *clientConfig) { config.applicationID = id }
}

// limitedBody caps how much of a response body is readable. The cap belongs in
// the transport rather than in each operation because every generated
// operation reads its body with io.ReadAll.
type limitedBody struct {
	body      io.ReadCloser
	remaining int64
}

func newLimitedBody(body io.ReadCloser, limit int64) io.ReadCloser {
	return &limitedBody{body: body, remaining: limit}
}

func (b *limitedBody) Read(p []byte) (int, error) {
	if b.remaining <= 0 {
		return 0, b.finishedOrOverLimit()
	}
	// Read one byte past the remaining budget so a body that lands exactly on
	// the limit still succeeds and the first byte over it is detected here,
	// rather than on a later read that may never come.
	if int64(len(p)) > b.remaining+1 {
		p = p[:b.remaining+1]
	}
	n, err := b.body.Read(p)
	if int64(n) > b.remaining {
		b.remaining = 0
		return 0, ErrResponseTooLarge
	}
	b.remaining -= int64(n)
	return n, err
}

// maxEmptyReads bounds how many times finishedOrOverLimit will tolerate a
// reader that returns neither a byte nor an error, matching what the standard
// library allows before it gives up on such a reader.
const maxEmptyReads = 100

// finishedOrOverLimit decides what an exhausted budget means. A body whose
// length lands exactly on the ceiling is complete and must report io.EOF; a
// body with one more byte waiting is over the ceiling. The two look identical
// from here, and a response without a Content-Length reaches this path on
// every exact-limit body, because its last read hands back the final bytes
// with no error attached and the end of the body only turns up on the read
// after that. So ask the underlying body for one more byte instead of
// assuming.
func (b *limitedBody) finishedOrOverLimit() error {
	var scratch [1]byte
	for attempt := 0; attempt < maxEmptyReads; attempt++ {
		n, err := b.body.Read(scratch[:])
		if n > 0 {
			return ErrResponseTooLarge
		}
		if err != nil {
			return err
		}
	}
	return io.ErrNoProgress
}

func (b *limitedBody) Close() error { return b.body.Close() }
