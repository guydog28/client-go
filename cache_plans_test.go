package ezclassify_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	ez "github.com/ezclassify/client-go"
)

// The cache plans are the contract the other EzClassify clients run too:
// spec/cache/v1/plans.json drives one cache-enabled client per plan against the
// recorded wire corpus. This file is the Go harness for them. The plans decide
// what is measured; nothing here knows what a plan expects beyond how to
// perform its steps.

// Capabilities the Go client declares not applicable, each for the one reason
// plans.json allows for it. Every other capability is supplied by this harness.
var goAbsentCapabilities = map[string]bool{
	// An optional request member is a nil pointer with omitempty, so an unset
	// member is always omitted from the body and an explicit null cannot be sent.
	"explicitNullMember": true,
	// Go exposes one method per wire operation. Access.Check serves both
	// marking and assertion targets, so no two methods post the same body to
	// the same route.
	"sharedRouteOperations": true,
}

type cachePlanFile struct {
	DenialCase string `json:"denialCase"`
	Operations []struct {
		Name      string `json:"name"`
		Method    string `json:"method"`
		Route     string `json:"route"`
		Treatment string `json:"treatment"`
	} `json:"operations"`
	Checklist []struct {
		ID string `json:"id"`
	} `json:"checklist"`
	Capabilities map[string]struct {
		NotApplicable string `json:"notApplicable"`
	} `json:"capabilities"`
	KnownGaps []json.RawMessage `json:"knownGaps"`
	Plans     []cachePlan       `json:"plans"`
}

type cachePlan struct {
	ID           string   `json:"id"`
	Covers       []string `json:"covers"`
	Capabilities []string `json:"capabilities"`
	Client       struct {
		Binding string `json:"binding"`
		Context string `json:"context"`
	} `json:"client"`
	Steps []cacheStep `json:"steps"`
}

type cacheErrorSpec struct {
	Kind   string `json:"kind"`
	Status int    `json:"status"`
}

type cacheStep struct {
	Step          string   `json:"step"`
	Operation     string   `json:"operation"`
	Case          string   `json:"case"`
	Caller        string   `json:"caller"`
	Callers       []string `json:"callers"`
	Cancelable    []string `json:"cancelable"`
	Times         int      `json:"times"`
	Variant       string   `json:"variant"`
	Member        string   `json:"member"`
	PreCanceled   bool     `json:"preCanceled"`
	MS            int64    `json:"ms"`
	Synthetic     string   `json:"synthetic"`
	RemoveMembers []string `json:"removeMembers"`
	Version       string   `json:"version"`
	Count         int      `json:"count"`
	Aborted       bool     `json:"aborted"`
	Expect        *struct {
		Error cacheErrorSpec `json:"error"`
	} `json:"expect"`
	Error *cacheErrorSpec `json:"error"`
}

func loadCachePlans(t *testing.T) cachePlanFile {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join("spec", "cache", "v1", "plans.json"))
	if err != nil {
		t.Fatal(err)
	}
	var plans cachePlanFile
	if err := json.Unmarshal(raw, &plans); err != nil {
		t.Fatal(err)
	}
	return plans
}

// TestCachePlanDeclarations pins the capabilities Go declares not applicable
// against the omissions plans.json allows. Nothing is skipped silently: a
// checklist item is either covered by a plan Go runs or covered only by plans
// that need a capability Go declares absent, and adding a plan that does
// neither fails here.
func TestCachePlanDeclarations(t *testing.T) {
	plans := loadCachePlans(t)
	if len(plans.KnownGaps) != 0 {
		t.Fatalf("plans.json lists %d known gaps; each is a rule the corpus cannot measure yet", len(plans.KnownGaps))
	}
	for name := range goAbsentCapabilities {
		capability, ok := plans.Capabilities[name]
		if !ok || capability.NotApplicable == "" {
			t.Errorf("capability %q is declared absent but plans.json allows no such omission", name)
		}
	}
	covered := map[string]bool{}
	notApplicable := map[string]bool{}
	for _, plan := range plans.Plans {
		for _, item := range plan.Covers {
			if cachePlanRuns(plan) {
				covered[item] = true
			} else {
				notApplicable[item] = true
			}
		}
		for _, capability := range plan.Capabilities {
			if _, ok := plans.Capabilities[capability]; !ok {
				t.Errorf("plan %s names unknown capability %q", plan.ID, capability)
			}
		}
	}
	for _, item := range plans.Checklist {
		if !covered[item.ID] && !notApplicable[item.ID] {
			t.Errorf("no plan covers checklist item %s", item.ID)
		}
	}
}

func cachePlanRuns(plan cachePlan) bool {
	for _, capability := range plan.Capabilities {
		if goAbsentCapabilities[capability] {
			return false
		}
	}
	return true
}

func TestCachePlans(t *testing.T) {
	plans := loadCachePlans(t)
	corpus, err := loadWireCorpus("spec")
	if err != nil {
		t.Fatal(err)
	}
	cases := map[string]wireCase{}
	for _, wire := range corpus.Cases {
		cases[wire.ID] = wire
	}
	denial, ok := cases[plans.DenialCase]
	if !ok {
		t.Fatalf("denial case %q is not in the corpus", plans.DenialCase)
	}
	routes := map[string][2]string{}
	for _, operation := range plans.Operations {
		routes[operation.Name] = [2]string{operation.Method, operation.Route}
	}
	for _, plan := range plans.Plans {
		plan := plan
		if !cachePlanRuns(plan) {
			continue
		}
		t.Run(plan.ID, func(t *testing.T) {
			harness := newCacheHarness(t, plan, cases, routes, denial)
			harness.run()
		})
	}
}

// cacheTransport is the counting replay transport: it answers each request with
// the recorded response of the corpus case that matches its method, path and
// canonical body, and counts requests per case.
type cacheTransport struct {
	mu       sync.Mutex
	byKey    map[string]string
	cases    map[string]wireCase
	denial   wireCase
	counts   map[string]int
	gated    map[string]bool
	held     map[string][]chan struct{}
	aborted  map[string]bool
	scripted map[string][]cacheResponse
	version  string
	failure  error
}

type cacheResponse struct {
	status  int
	headers map[string]string
	body    []byte
}

// versionless replaces the version segment of a Context path, since the
// binding may carry its own version, the resolved one, or "latest".
func versionless(target string) string {
	path, query, hasQuery := strings.Cut(target, "?")
	parts := strings.Split(path, "/")
	if len(parts) > 4 && parts[1] == "v1" && parts[2] == "ctx" {
		parts[4] = "{version}"
	}
	path = strings.Join(parts, "/")
	if hasQuery {
		return path + "?" + query
	}
	return path
}

func newCacheTransport(cases map[string]wireCase, denial wireCase) *cacheTransport {
	tr := &cacheTransport{
		byKey:    map[string]string{},
		cases:    cases,
		denial:   denial,
		counts:   map[string]int{},
		gated:    map[string]bool{},
		held:     map[string][]chan struct{}{},
		aborted:  map[string]bool{},
		scripted: map[string][]cacheResponse{},
	}
	for id, wire := range cases {
		// The denial and routing recordings repeat a Context request with other
		// credentials or with no route behind it. Replaying them by request
		// alone would make the recorded Context cases ambiguous.
		if wire.Auth.Credential != "context-token" || !strings.HasPrefix(wire.Request.Path, "/v1/ctx/") || wire.Operation == "" {
			continue
		}
		key, err := wireKey(wire.Request.Method, versionless(wire.Request.Path), wire.Request.Body)
		if err != nil {
			continue
		}
		if other, dup := tr.byKey[key]; dup {
			// Two recordings with one request cannot be told apart by a replay.
			// Neither would be reachable by method, path and body.
			tr.byKey[key] = other + "|" + id
			continue
		}
		tr.byKey[key] = id
	}
	return tr
}

func (tr *cacheTransport) fail(err error) error {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	if tr.failure == nil {
		tr.failure = err
	}
	return err
}

func (tr *cacheTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var body []byte
	if req.Body != nil {
		var err error
		if body, err = io.ReadAll(req.Body); err != nil {
			return nil, err
		}
	}
	key, err := wireKey(req.Method, versionless(req.URL.RequestURI()), body)
	if err != nil {
		return nil, tr.fail(err)
	}
	tr.mu.Lock()
	id, matched := tr.byKey[key]
	if !matched || strings.Contains(id, "|") {
		tr.mu.Unlock()
		return nil, tr.fail(fmt.Errorf("request has no single recorded case: %s (%s)", key, id))
	}
	tr.counts[id]++
	response := cacheResponse{status: tr.cases[id].Response.Status, headers: tr.cases[id].Response.Headers, body: tr.cases[id].Response.Body}
	if queue := tr.scripted[id]; len(queue) > 0 {
		response, tr.scripted[id] = queue[0], queue[1:]
	}
	var release chan struct{}
	if tr.gated[id] {
		release = make(chan struct{})
		tr.held[id] = append(tr.held[id], release)
	}
	version := tr.version
	tr.mu.Unlock()
	if release != nil {
		select {
		case <-release:
		case <-req.Context().Done():
			tr.mu.Lock()
			tr.aborted[id] = true
			tr.dropHold(id, release)
			tr.mu.Unlock()
			return nil, req.Context().Err()
		}
	}
	if version != "" {
		response = withResolvedVersion(response, version)
	}
	header := http.Header{}
	for name, value := range response.headers {
		header.Set(name, value)
	}
	return &http.Response{
		StatusCode:    response.status,
		Status:        fmt.Sprintf("%d %s", response.status, http.StatusText(response.status)),
		Header:        header,
		Body:          io.NopCloser(bytes.NewReader(response.body)),
		ContentLength: int64(len(response.body)),
		Request:       req,
	}, nil
}

func (tr *cacheTransport) dropHold(id string, release chan struct{}) {
	kept := tr.held[id][:0]
	for _, other := range tr.held[id] {
		if other != release {
			kept = append(kept, other)
		}
	}
	tr.held[id] = kept
}

func (tr *cacheTransport) count(id string) int {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	return tr.counts[id]
}

func (tr *cacheTransport) heldCount(id string) int {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	return len(tr.held[id])
}

func (tr *cacheTransport) wasAborted(id string) bool {
	tr.mu.Lock()
	defer tr.mu.Unlock()
	return tr.aborted[id]
}

func (tr *cacheTransport) gate(id string) {
	tr.mu.Lock()
	tr.gated[id] = true
	tr.mu.Unlock()
}

// releaseCase lifts the gate on a case and answers the requests it held. A
// request whose caller already canceled it left the hold list when it aborted.
func (tr *cacheTransport) releaseCase(id string) {
	tr.mu.Lock()
	tr.gated[id] = false
	held := tr.held[id]
	tr.held[id] = nil
	tr.mu.Unlock()
	for _, release := range held {
		close(release)
	}
}

func (tr *cacheTransport) script(id string, response cacheResponse) {
	tr.mu.Lock()
	tr.scripted[id] = append(tr.scripted[id], response)
	tr.mu.Unlock()
}

func (tr *cacheTransport) moveVersion(version string) {
	tr.mu.Lock()
	tr.version = version
	tr.mu.Unlock()
}

// withResolvedVersion rewrites the version a recorded response names, in the
// header and wherever the body repeats it: the resolution, or a Context
// description's own version.
func withResolvedVersion(response cacheResponse, version string) cacheResponse {
	headers := map[string]string{}
	for name, value := range response.headers {
		if strings.EqualFold(name, "x-ezclassify-resolved-version") {
			value = version
		}
		headers[name] = value
	}
	response.headers = headers
	var body map[string]any
	decoder := json.NewDecoder(bytes.NewReader(response.body))
	decoder.UseNumber()
	if decoder.Decode(&body) != nil {
		return response
	}
	if resolution, ok := body["resolution"].(map[string]any); ok {
		resolution["resolvedVersion"] = version
	}
	if _, described := body["schemes"]; described {
		body["version"] = version
	}
	if rewritten, err := json.Marshal(body); err == nil {
		response.body = rewritten
	}
	return response
}

// removeMembers derives a malformed success from recorded bytes by deleting the
// members the JSON pointers name.
func removeMembers(recorded []byte, pointers []string) ([]byte, error) {
	var document any
	decoder := json.NewDecoder(bytes.NewReader(recorded))
	decoder.UseNumber()
	if err := decoder.Decode(&document); err != nil {
		return nil, err
	}
	for _, pointer := range pointers {
		segments := strings.Split(strings.TrimPrefix(pointer, "/"), "/")
		node := document
		for _, segment := range segments[:len(segments)-1] {
			object, ok := node.(map[string]any)
			if !ok {
				return nil, fmt.Errorf("pointer %s leaves the object at %s", pointer, segment)
			}
			node = object[segment]
		}
		object, ok := node.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("pointer %s does not name an object member", pointer)
		}
		last := segments[len(segments)-1]
		if _, present := object[last]; !present {
			return nil, fmt.Errorf("pointer %s names a member the recording does not have", pointer)
		}
		delete(object, last)
	}
	return json.Marshal(document)
}

// reverseMembers writes a JSON document with the members of every object in the
// opposite order. Go maps and structs serialize in a fixed order, so the
// variant shows the key does not depend on how a caller built its arguments.
func reverseMembers(raw json.RawMessage) json.RawMessage {
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.UseNumber()
	token, err := decoder.Token()
	if err != nil {
		return raw
	}
	delim, isDelim := token.(json.Delim)
	if !isDelim {
		return raw
	}
	var out bytes.Buffer
	switch delim {
	case '{':
		type member struct {
			key   string
			value json.RawMessage
		}
		var members []member
		for decoder.More() {
			keyToken, _ := decoder.Token()
			var value json.RawMessage
			if err := decoder.Decode(&value); err != nil {
				return raw
			}
			members = append(members, member{keyToken.(string), reverseMembers(value)})
		}
		out.WriteByte('{')
		for i := len(members) - 1; i >= 0; i-- {
			key, _ := json.Marshal(members[i].key)
			out.Write(key)
			out.WriteByte(':')
			out.Write(members[i].value)
			if i > 0 {
				out.WriteByte(',')
			}
		}
		out.WriteByte('}')
	case '[':
		var items []json.RawMessage
		for decoder.More() {
			var item json.RawMessage
			if err := decoder.Decode(&item); err != nil {
				return raw
			}
			items = append(items, reverseMembers(item))
		}
		out.WriteByte('[')
		for i, item := range items {
			if i > 0 {
				out.WriteByte(',')
			}
			out.Write(item)
		}
		out.WriteByte(']')
	}
	return out.Bytes()
}

// cacheClock is the clock the response cache reads in these plans.
type cacheClock struct {
	mu  sync.Mutex
	now time.Time
}

func (c *cacheClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *cacheClock) advance(d time.Duration) {
	c.mu.Lock()
	c.now = c.now.Add(d)
	c.mu.Unlock()
}

type cacheOutcome struct {
	caseID string
	result any
	err    error
	done   chan struct{}
	cancel context.CancelFunc
}

type cacheHarness struct {
	t        *testing.T
	plan     cachePlan
	cases    map[string]wireCase
	routes   map[string][2]string
	tr       *cacheTransport
	cx       *ez.EzClassifyContext
	clock    *cacheClock
	outcomes map[string]*cacheOutcome
	started  []*cacheOutcome
}

func newCacheHarness(t *testing.T, plan cachePlan, cases map[string]wireCase, routes map[string][2]string, denial wireCase) *cacheHarness {
	t.Helper()
	return &cacheHarness{
		t: t, plan: plan, cases: cases, routes: routes,
		tr:       newCacheTransport(cases, denial),
		clock:    &cacheClock{now: time.Unix(1_700_000_000, 0)},
		outcomes: map[string]*cacheOutcome{},
	}
}

// binding builds the plan's one client and connects it. The Go client has no
// connect request: a handle only addresses a Context, so there is no connect
// traffic to leave out of the counts. A token-defaulted binding omits the
// version, which leaves the choice to the Context token; an unpinned binding
// names "latest". The two send the same request.
func (h *cacheHarness) binding() {
	t := h.t
	contextID := h.plan.Client.Context
	if contextID == "" {
		contextID = "ez-full"
	}
	pinned := "1.0.0"
	for _, step := range h.plan.Steps {
		if step.Case == "" {
			continue
		}
		parts := strings.Split(strings.Trim(h.cases[step.Case].Request.Path, "/"), "/")
		if len(parts) < 4 || parts[2] != contextID {
			t.Fatalf("step case %s is recorded against %v, not Context %s", step.Case, parts, contextID)
		}
		pinned = parts[3]
	}
	var version string
	switch h.plan.Client.Binding {
	case "pinned":
		version = pinned
	case "unpinned":
		version = "latest"
	case "token-defaulted":
		version = ""
	default:
		t.Fatalf("unknown binding %q", h.plan.Client.Binding)
	}
	options := []ez.Option{
		ez.WithHTTPClient(&http.Client{Transport: h.tr}),
		ez.WithCredentials(ez.RestClientCredentials{ContextToken: "fixture-context-carrier"}),
	}
	client := ez.New("http://cache.invalid", options...)
	h.cx = ez.ConnectWith(client, contextID, version)
	ez.SetCacheClockForTest(h.cx, h.clock.Now)
}

func (h *cacheHarness) run() {
	t := h.t
	h.binding()
	for i, step := range h.plan.Steps {
		h.perform(i, step)
		h.tr.mu.Lock()
		failure := h.tr.failure
		h.tr.mu.Unlock()
		if failure != nil {
			t.Fatalf("step %d (%s): %v", i, step.Step, failure)
		}
	}
	// Every caller a plan started must have settled, or it leaked.
	for _, outcome := range h.started {
		h.settled(outcome)
	}
}

func (h *cacheHarness) wire(step cacheStep) wireCase {
	h.t.Helper()
	wire, ok := h.cases[step.Case]
	if !ok {
		h.t.Fatalf("case %q is not in the corpus", step.Case)
	}
	return wire
}

// invoke performs the plan operation by the public method the case's recorded
// operation names, sending exactly the case's recorded request.
func (h *cacheHarness) invoke(ctx context.Context, step cacheStep) (any, error) {
	h.t.Helper()
	wire := h.wire(step)
	route, known := h.routes[step.Operation]
	path, _, _ := strings.Cut(wire.Request.Path, "?")
	if !known || route[0] != wire.Request.Method || !strings.HasSuffix(path, "/"+route[1]) {
		h.t.Fatalf("operation %s (%v) does not describe case %s (%s %s)", step.Operation, route, step.Case, wire.Request.Method, wire.Request.Path)
	}
	switch step.Variant {
	case "":
	case "reorderedMembers":
		if len(wire.Request.Body) > 0 {
			wire.Request.Body = reverseMembers(wire.Request.Body)
		}
	case "unsetMember":
		var members map[string]json.RawMessage
		if err := json.Unmarshal(wire.Request.Body, &members); err != nil {
			h.t.Fatal(err)
		}
		if _, present := members[step.Member]; present {
			h.t.Fatalf("case %s sets %s, so it cannot stand for the member passed as unset", step.Case, step.Member)
		}
	default:
		h.t.Fatalf("variant %q is not supported by this harness", step.Variant)
	}
	parsed, err := url.Parse(wire.Request.Path)
	if err != nil {
		h.t.Fatal(err)
	}
	return driveWire(ctx, h.cx, wire, parsed)
}

func errorMatches(err error, want cacheErrorSpec) error {
	switch want.Kind {
	case "http":
		var apiErr *ez.APIError
		if !errors.As(err, &apiErr) || apiErr.StatusCode != want.Status {
			return fmt.Errorf("want HTTP %d, got %T: %v", want.Status, err, err)
		}
	case "aborted":
		if !errors.Is(err, context.Canceled) {
			return fmt.Errorf("want a canceled call, got %T: %v", err, err)
		}
	case "malformed-response":
		var malformed *ez.MalformedResponseError
		if !errors.As(err, &malformed) {
			return fmt.Errorf("want a malformed response error, got %T: %v", err, err)
		}
	default:
		return fmt.Errorf("unknown error kind %q", want.Kind)
	}
	return nil
}

func (h *cacheHarness) perform(index int, step cacheStep) {
	t := h.t
	t.Helper()
	switch step.Step {
	case "call":
		times := step.Times
		if times == 0 {
			times = 1
		}
		for n := 0; n < times; n++ {
			ctx, cancel := context.WithCancel(context.Background())
			if step.PreCanceled {
				cancel()
			}
			result, err := h.invoke(ctx, step)
			cancel()
			if step.Expect != nil {
				if err == nil {
					t.Fatalf("step %d: %s succeeded, want %v", index, step.Operation, step.Expect.Error)
				}
				if mismatch := errorMatches(err, step.Expect.Error); mismatch != nil {
					t.Fatalf("step %d: %s: %v", index, step.Operation, mismatch)
				}
			} else if err != nil {
				t.Fatalf("step %d: %s: %v", index, step.Operation, err)
			}
			if step.Caller != "" {
				done := make(chan struct{})
				close(done)
				h.outcomes[step.Caller] = &cacheOutcome{caseID: step.Case, result: result, err: err, done: done}
			}
		}
	case "concurrent":
		h.concurrent(step)
	case "gate":
		h.tr.gate(step.Case)
	case "release":
		h.tr.releaseCase(step.Case)
		for _, outcome := range h.started {
			if outcome.caseID == step.Case {
				h.settled(outcome)
			}
		}
	case "cancel":
		outcome := h.outcome(step.Caller)
		if outcome.cancel == nil {
			t.Fatalf("caller %s was not started cancelable", step.Caller)
		}
		outcome.cancel()
	case "advanceClock":
		h.clock.advance(time.Duration(step.MS) * time.Millisecond)
	case "respond":
		h.respond(step)
	case "quiesce":
		h.waitFor("the cache to have no waiting callers", func() bool { return ez.CacheWaitersForTest(h.cx) == 0 })
	case "mutateResult":
		outcome := h.outcome(step.Caller)
		h.settled(outcome)
		if outcome.err != nil {
			t.Fatalf("caller %s has no result to mutate: %v", step.Caller, outcome.err)
		}
		if mutateInPlace(reflect.ValueOf(outcome.result)) == 0 {
			t.Fatalf("caller %s: the result held nothing a caller could change", step.Caller)
		}
	case "expectOk":
		for _, name := range step.Callers {
			outcome := h.outcome(name)
			h.settled(outcome)
			if outcome.err != nil {
				t.Fatalf("caller %s: %v", name, outcome.err)
			}
		}
	case "expectError":
		for _, name := range step.Callers {
			outcome := h.outcome(name)
			h.settled(outcome)
			if outcome.err == nil {
				t.Fatalf("caller %s succeeded, want %v", name, *step.Error)
			}
			if mismatch := errorMatches(outcome.err, *step.Error); mismatch != nil {
				t.Fatalf("caller %s: %v", name, mismatch)
			}
		}
	case "expectWireCount":
		if got := h.tr.count(step.Case); got != step.Count {
			t.Fatalf("step %d: %s reached the wire %d times, want %d", index, step.Case, got, step.Count)
		}
	case "expectWireAborted":
		h.expectAborted(step)
	case "expectIndependentResults":
		h.expectIndependent(step)
	default:
		t.Fatalf("step %q is not supported by this harness", step.Step)
	}
}

func (h *cacheHarness) outcome(name string) *cacheOutcome {
	outcome, ok := h.outcomes[name]
	if !ok {
		h.t.Fatalf("no caller named %q", name)
	}
	return outcome
}

func (h *cacheHarness) waitFor(what string, condition func() bool) {
	h.t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for !condition() {
		if time.Now().After(deadline) {
			h.t.Fatalf("timed out waiting for %s", what)
		}
		time.Sleep(time.Millisecond)
	}
}

func (h *cacheHarness) settled(outcome *cacheOutcome) {
	h.t.Helper()
	select {
	case <-outcome.done:
	case <-time.After(10 * time.Second):
		h.t.Fatalf("a caller of %s never settled", outcome.caseID)
	}
}

// concurrent starts one call per caller at once. With the case gated it returns
// once the held request has reached the transport and every new caller has
// joined the in-flight entry, which the cache reports as its waiter count.
func (h *cacheHarness) concurrent(step cacheStep) {
	waitingBefore := ez.CacheWaitersForTest(h.cx)
	cancelable := map[string]bool{}
	for _, name := range step.Cancelable {
		cancelable[name] = true
	}
	var batch []*cacheOutcome
	for _, name := range step.Callers {
		ctx, cancel := context.WithCancel(context.Background())
		outcome := &cacheOutcome{caseID: step.Case, done: make(chan struct{})}
		if cancelable[name] {
			outcome.cancel = cancel
		}
		h.outcomes[name] = outcome
		h.started = append(h.started, outcome)
		batch = append(batch, outcome)
		go func() {
			defer close(outcome.done)
			defer cancel()
			outcome.result, outcome.err = h.invoke(ctx, step)
		}()
	}
	h.tr.mu.Lock()
	gated := h.tr.gated[step.Case]
	h.tr.mu.Unlock()
	if gated {
		h.waitFor("the held request to reach the transport", func() bool { return h.tr.heldCount(step.Case) > 0 })
		h.waitFor("every caller to join the in-flight entry", func() bool {
			return ez.CacheWaitersForTest(h.cx) >= waitingBefore+len(step.Callers)
		})
		return
	}
	for _, outcome := range batch {
		h.settled(outcome)
	}
}

func (h *cacheHarness) respond(step cacheStep) {
	t := h.t
	switch step.Synthetic {
	case "malformed-success":
		recorded := h.wire(step).Response
		body, err := removeMembers(recorded.Body, step.RemoveMembers)
		if err != nil {
			t.Fatal(err)
		}
		h.tr.script(step.Case, cacheResponse{status: recorded.Status, headers: recorded.Headers, body: body})
	case "denial-404":
		recorded := h.tr.denial.Response
		h.tr.script(step.Case, cacheResponse{status: recorded.Status, headers: recorded.Headers, body: recorded.Body})
	case "resolved-version":
		h.tr.moveVersion(step.Version)
	default:
		t.Fatalf("synthetic response %q is not supported by this harness", step.Synthetic)
	}
}

func (h *cacheHarness) expectAborted(step cacheStep) {
	if step.Aborted {
		h.waitFor("the wire request to be canceled", func() bool { return h.tr.wasAborted(step.Case) })
		return
	}
	// Absence cannot be waited for. Give a wrongly forwarded cancellation time
	// to reach the transport before concluding it did not.
	time.Sleep(25 * time.Millisecond)
	if h.tr.wasAborted(step.Case) {
		h.t.Fatalf("the wire request for %s was canceled while a caller still waited on it", step.Case)
	}
}

// expectIndependent holds each caller's result to a fresh decode of the
// recorded response it was answered from.
func (h *cacheHarness) expectIndependent(step cacheStep) {
	for _, name := range step.Callers {
		outcome := h.outcome(name)
		h.settled(outcome)
		if outcome.err != nil {
			h.t.Fatalf("caller %s: %v", name, outcome.err)
		}
		fresh := reflect.New(reflect.TypeOf(outcome.result).Elem())
		if err := json.Unmarshal(h.cases[outcome.caseID].Response.Body, fresh.Interface()); err != nil {
			h.t.Fatal(err)
		}
		if !reflect.DeepEqual(fresh.Interface(), outcome.result) {
			h.t.Fatalf("caller %s holds a result that differs from a fresh decode of its response: another caller's change reached it", name)
		}
	}
}

// mutateInPlace changes every value reachable from v that a caller could
// change: scalars through pointers, slice elements and map entries. It returns
// how many it changed.
func mutateInPlace(v reflect.Value) int {
	changed := 0
	switch v.Kind() {
	case reflect.Pointer, reflect.Interface:
		if !v.IsNil() {
			changed += mutateInPlace(v.Elem())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Type().Field(i).IsExported() {
				changed += mutateInPlace(v.Field(i))
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			changed += mutateInPlace(v.Index(i))
		}
	case reflect.Map:
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool { return fmt.Sprint(keys[i]) < fmt.Sprint(keys[j]) })
		for _, key := range keys {
			entry := reflect.New(v.Type().Elem()).Elem()
			entry.Set(v.MapIndex(key))
			if entry.Kind() == reflect.Interface && !entry.IsNil() {
				// An interface holding a scalar or composite cannot be changed
				// in place through the map; replace it with a changed copy.
				inner := reflect.New(entry.Elem().Type()).Elem()
				inner.Set(entry.Elem())
				if n := mutateInPlace(inner); n > 0 {
					entry.Set(inner)
					changed += n
				}
			} else if n := mutateInPlace(entry); n > 0 {
				changed += n
			}
			v.SetMapIndex(key, entry)
		}
	case reflect.String:
		if v.CanSet() {
			v.SetString(v.String() + "-changed")
			changed++
		}
	case reflect.Bool:
		if v.CanSet() {
			v.SetBool(!v.Bool())
			changed++
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.CanSet() {
			v.SetInt(v.Int() + 1)
			changed++
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.CanSet() {
			v.SetUint(v.Uint() + 1)
			changed++
		}
	case reflect.Float32, reflect.Float64:
		if v.CanSet() {
			v.SetFloat(v.Float() + 1)
			changed++
		}
	}
	return changed
}
