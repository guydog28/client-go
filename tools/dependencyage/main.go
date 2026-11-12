// Command dependencyage is the dependency quarantine gate. It refuses
// third-party Go modules whose pinned version first reached the public module
// proxy inside a quarantine window, so a freshly compromised release has to sit
// on the proxy for a while before it can enter this repository's dependency
// tree.
//
// It reads go.sum rather than go.mod: go.sum names every module version the
// build graph resolved, including transitive ones, which is the set that
// matters. With no -gosum flag it checks every go.sum in the repository (the
// root module and each tool module) except those under vendor/ and testdata/.
// Modules under the ezclassify organization are first-party and exempt.
//
// A version's age is the time proxy.golang.org first cached it, as recorded by
// the module index (index.golang.org). The Time field of the proxy's
// <module>/@v/<version>.info answer is not used for age: it is the commit time,
// which whoever tags the version controls. The gate still asks the proxy for
// each .info, which confirms the proxy serves the version and makes the proxy
// cache it if it had not yet. It then reads the index from the start of the
// window until the index has caught up with the moment those lookups finished.
// A version the index lists inside that span is inside the window; a version it
// does not list was first cached before the window opened.
//
// The gate fails closed: a proxy or index that cannot be read, an answer that
// cannot be parsed, or an unreadable go.sum is a failure, never a silent pass.
// When the index rate limits or fails with a server error, the gate waits as
// the index asks (or backs off) and tries again. Every request has a timeout
// and the whole run has an overall bound, so a server that accepts
// connections and never answers, an index that keeps refusing, or an index
// that never catches up cannot hang the job.
//
// Usage:
//
//	go run ./tools/dependencyage [-days N] [-gosum PATH]... [-proxy URL] [-index URL]
//
// -days, -proxy and -index have environment forms:
// EZCLASSIFY_DEPENDENCY_AGE_DAYS, EZCLASSIFY_DEPENDENCY_AGE_PROXY and
// EZCLASSIFY_DEPENDENCY_AGE_INDEX.
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// defaultQuarantineDays is long enough that a malicious release has usually
	// been reported and retracted, and short enough that ordinary upgrades are
	// not held up. Dependabot's cooldown matches it.
	defaultQuarantineDays = 7

	// defaultProxy is the public Go module proxy.
	defaultProxy = "https://proxy.golang.org"

	// defaultIndex is the feed of versions in the order defaultProxy first
	// cached them.
	defaultIndex = "https://index.golang.org/index"

	// firstPartyPrefix is the module path prefix owned by this project.
	firstPartyPrefix = "github.com/ezclassify/"

	// requestTimeout bounds one proxy or index request.
	requestTimeout = 30 * time.Second

	// overallTimeout bounds the whole run; anything still unchecked at that
	// point is reported and fails the gate.
	overallTimeout = 15 * time.Minute

	// concurrency is the number of proxy lookups, and the number of index
	// segments read, in flight at once.
	concurrency = 4

	// indexPageSize is the largest page the index serves.
	indexPageSize = 2000

	// maxIndexPageBytes bounds one index page; a full page is about 300 KB.
	maxIndexPageBytes = 16 << 20

	day = 24 * time.Hour
)

// indexRetryPause is how long the gate waits before asking the index again
// when the index has not yet caught up with the end of the span it must cover.
// The index normally trails the proxy by seconds.
var indexRetryPause = 10 * time.Second

// maxIndexBackoff caps one wait between index requests when the index is
// rate limiting or failing.
const maxIndexBackoff = time.Minute

// module is one module version pinned by go.sum.
type module struct {
	Path    string
	Version string
}

func (m module) String() string { return m.Path + "@" + m.Version }

// finding is a module that failed the gate. Reason is set when its age could
// not be confirmed; otherwise FirstSeen is inside the window.
type finding struct {
	Module    module
	FirstSeen time.Time
	Clears    time.Time
	Reason    string
}

// pathList is a repeatable -gosum flag.
type pathList []string

func (p *pathList) String() string { return strings.Join(*p, ",") }

func (p *pathList) Set(v string) error {
	*p = append(*p, v)
	return nil
}

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("dependencyage", flag.ContinueOnError)
	flags.SetOutput(stderr)
	days := flags.String("days", envOr("EZCLASSIFY_DEPENDENCY_AGE_DAYS", strconv.Itoa(defaultQuarantineDays)), "quarantine window in days")
	var gosums pathList
	flags.Var(&gosums, "gosum", "path to a go.sum to check; repeat for several (default: every go.sum in the repository)")
	proxy := flags.String("proxy", envOr("EZCLASSIFY_DEPENDENCY_AGE_PROXY", defaultProxy), "module proxy base URL")
	index := flags.String("index", envOr("EZCLASSIFY_DEPENDENCY_AGE_INDEX", defaultIndex), "module index URL")
	if err := flags.Parse(args); err != nil {
		return err
	}
	window, err := strconv.ParseFloat(*days, 64)
	if err != nil || window < 0 {
		return fmt.Errorf("quarantine window must be a non-negative number of days, got %q", *days)
	}
	if len(gosums) == 0 {
		if gosums, err = findGoSums("."); err != nil {
			return err
		}
	}

	modules, err := readGoSums(gosums)
	if err != nil {
		return err
	}
	if window == 0 {
		fmt.Fprintf(stdout, "Dependency quarantine window is 0 days; nothing to check (%d module version(s)).\n", len(modules))
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), overallTimeout)
	defer cancel()
	findings := check(ctx, http.DefaultClient, strings.TrimRight(*proxy, "/"), *index, modules, time.Duration(window*float64(day)))
	if len(findings) == 0 {
		fmt.Fprintf(stdout, "Dependency quarantine passed: %d third-party module version(s) from %s all reached the module proxy more than %s day(s) ago.\n", len(modules), strings.Join(gosums, ", "), *days)
		return nil
	}

	var unknown, fresh []finding
	for _, f := range findings {
		if f.Reason != "" {
			unknown = append(unknown, f)
		} else {
			fresh = append(fresh, f)
		}
	}
	if len(unknown) > 0 {
		fmt.Fprintf(stderr, "Dependency quarantine: could not confirm the age of %d module version(s), so the gate fails closed.\n", len(unknown))
		for _, f := range unknown {
			fmt.Fprintf(stderr, "  %s: %s\n", f.Module, f.Reason)
		}
	}
	if len(fresh) > 0 {
		fmt.Fprintf(stderr, "Dependency quarantine: %d third-party module version(s) reached the module proxy less than %s day(s) ago.\n", len(fresh), *days)
		for _, f := range fresh {
			fmt.Fprintf(stderr, "  %s first seen %s, clears %s\n", f.Module, f.FirstSeen.UTC().Format(time.RFC3339), f.Clears.UTC().Format(time.RFC3339))
		}
		fmt.Fprintln(stderr, "Wait out the window, or pin an older version, before taking these in.")
	}
	return errors.New("dependency quarantine failed")
}

func envOr(name, fallback string) string {
	if v := os.Getenv(name); v != "" {
		return v
	}
	return fallback
}

// findGoSums returns every go.sum under root, sorted, skipping vendor/,
// testdata/ and hidden directories. Finding none is an error, so a gate run
// from the wrong directory cannot pass by checking nothing.
func findGoSums(root string) ([]string, error) {
	var found []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if path != root && (name == "vendor" || name == "testdata" || strings.HasPrefix(name, ".")) {
				return filepath.SkipDir
			}
			return nil
		}
		if d.Name() == "go.sum" {
			found = append(found, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("cannot search %s for go.sum files: %w", root, err)
	}
	if len(found) == 0 {
		return nil, fmt.Errorf("no go.sum found under %s", root)
	}
	sort.Strings(found)
	return found, nil
}

// readGoSums returns the distinct third-party module versions across the
// named go.sum files, sorted.
func readGoSums(paths []string) ([]module, error) {
	seen := map[module]bool{}
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("cannot read %s: %w", path, err)
		}
		err = thirdPartyModules(file, seen)
		file.Close()
		if err != nil {
			return nil, fmt.Errorf("cannot parse %s: %w", path, err)
		}
	}
	modules := make([]module, 0, len(seen))
	for m := range seen {
		modules = append(modules, m)
	}
	sort.Slice(modules, func(i, j int) bool { return modules[i].String() < modules[j].String() })
	return modules, nil
}

// thirdPartyModules adds the third-party module versions in a go.sum stream to
// seen. Each line is "<module> <version>[/go.mod] <hash>".
func thirdPartyModules(r io.Reader, seen map[module]bool) error {
	scanner := bufio.NewScanner(r)
	for n := 1; scanner.Scan(); n++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return fmt.Errorf("line %d: expected \"<module> <version> <hash>\"", n)
		}
		if strings.HasPrefix(fields[0], firstPartyPrefix) {
			continue
		}
		seen[module{Path: fields[0], Version: strings.TrimSuffix(fields[1], "/go.mod")}] = true
	}
	return scanner.Err()
}

// check returns the modules that fail the gate: those the proxy does not
// confirm, those first cached inside the window, and, when the index cannot be
// read, every module the index was needed for.
func check(ctx context.Context, client *http.Client, proxy, index string, modules []module, window time.Duration) []finding {
	var findings []finding
	served := confirmServed(ctx, client, proxy, modules, &findings)
	if len(served) > 0 {
		// Every served version was cached by the proxy no later than now, so an
		// index that reaches now lists it if it was first cached in the window.
		until := time.Now()
		firstSeen, err := scanIndex(ctx, client, index, until.Add(-window), until, served)
		for _, m := range modules {
			if !served[m] {
				continue
			}
			switch seen, ok := firstSeen[m]; {
			case err != nil:
				findings = append(findings, finding{Module: m, Reason: err.Error()})
			case ok:
				findings = append(findings, finding{Module: m, FirstSeen: seen, Clears: seen.Add(window)})
			}
		}
	}
	sort.Slice(findings, func(i, j int) bool { return findings[i].Module.String() < findings[j].Module.String() })
	return findings
}

// confirmServed asks the proxy for every module's .info and returns the set it
// serves. Each module it does not confirm is added to findings.
func confirmServed(ctx context.Context, client *http.Client, proxy string, modules []module, findings *[]finding) map[module]bool {
	var (
		mu     sync.Mutex
		served = map[module]bool{}
		wg     sync.WaitGroup
		slots  = make(chan struct{}, concurrency)
	)
	fail := func(m module, reason string) {
		mu.Lock()
		*findings = append(*findings, finding{Module: m, Reason: reason})
		mu.Unlock()
	}
	for _, m := range modules {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case slots <- struct{}{}:
				defer func() { <-slots }()
			case <-ctx.Done():
				fail(m, "the run passed its overall time bound before it was checked")
				return
			}
			if err := proxyServes(ctx, client, proxy, m); err != nil {
				fail(m, err.Error())
				return
			}
			mu.Lock()
			served[m] = true
			mu.Unlock()
		}()
	}
	wg.Wait()
	return served
}

// proxyServes confirms the module proxy serves a version by fetching its
// .info. The answer's Time is the commit time and is not used.
func proxyServes(ctx context.Context, client *http.Client, proxy string, m module) error {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, proxy+"/"+escapePath(m.Path)+"/@v/"+escapePath(m.Version)+".info", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not reach the proxy: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("the proxy answered %d; cannot confirm it serves this version", resp.StatusCode)
	}
	var info struct{ Version string }
	if err := json.NewDecoder(io.LimitReader(resp.Body, 1<<20)).Decode(&info); err != nil {
		return fmt.Errorf("the proxy gave an unreadable answer: %w", err)
	}
	if info.Version != m.Version {
		return fmt.Errorf("the proxy answered for version %q", info.Version)
	}
	return nil
}

// busyError is an index answer that asks the caller to come back later: 429
// or a server error. retryAfter is the server's Retry-After, or zero.
type busyError struct {
	status     int
	retryAfter time.Duration
}

func (e *busyError) Error() string { return fmt.Sprintf("the index kept answering %d", e.status) }

// indexEntry is one line of the module index.
type indexEntry struct {
	Path      string
	Version   string
	Timestamp time.Time
}

// scanIndex reads the module index over [from, until) and returns the time
// each wanted version was first cached, for those cached in that span. The
// span is split into segments read in parallel. It fails if any segment
// cannot be read completely before ctx ends.
func scanIndex(ctx context.Context, client *http.Client, index string, from, until time.Time, want map[module]bool) (map[module]time.Time, error) {
	var (
		mu        sync.Mutex
		firstSeen = map[module]time.Time{}
		errs      []error
		wg        sync.WaitGroup
	)
	step := until.Sub(from) / concurrency
	for i := range concurrency {
		start := from.Add(time.Duration(i) * step)
		end := start.Add(step)
		if i == concurrency-1 {
			end = until
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := scanSegment(ctx, client, index, start, end, func(e indexEntry) {
				m := module{Path: e.Path, Version: e.Version}
				if !want[m] {
					return
				}
				mu.Lock()
				if seen, ok := firstSeen[m]; !ok || e.Timestamp.Before(seen) {
					firstSeen[m] = e.Timestamp
				}
				mu.Unlock()
			})
			if err != nil {
				mu.Lock()
				errs = append(errs, err)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	if len(errs) > 0 {
		return nil, fmt.Errorf("could not read the module index, so first-cache times are unconfirmed: %w", errs[0])
	}
	return firstSeen, nil
}

// scanSegment pages through the index from start and hands each entry with a
// timestamp in [start, end) to visit. It stops at the first entry at or after
// end. When the index ends before reaching end, it waits and asks again, until
// ctx ends.
func scanSegment(ctx context.Context, client *http.Client, index string, start, end time.Time, visit func(indexEntry)) error {
	since := start
	backoff := indexRetryPause
	for {
		page, err := indexPage(ctx, client, index, since)
		var busy *busyError
		if errors.As(err, &busy) {
			// The index is shedding load. Wait as long as it asks, or back off,
			// and ask again until ctx ends.
			wait := busy.retryAfter
			if wait <= 0 {
				wait = backoff
				backoff = min(2*backoff, maxIndexBackoff)
			}
			select {
			case <-ctx.Done():
				return fmt.Errorf("%w until the run ended", err)
			case <-time.After(min(wait, maxIndexBackoff)):
			}
			continue
		}
		if err != nil {
			return err
		}
		backoff = indexRetryPause
		for _, e := range page {
			if !e.Timestamp.Before(end) {
				return nil
			}
			if !e.Timestamp.Before(start) {
				visit(e)
			}
		}
		if len(page) == indexPageSize {
			last := page[len(page)-1].Timestamp
			if !last.After(since) {
				return fmt.Errorf("the index did not advance past %s", since.UTC().Format(time.RFC3339Nano))
			}
			since = last
			continue
		}
		// The index has nothing newer yet. It trails the proxy by seconds, so
		// wait briefly and ask again from the same point.
		if len(page) > 0 {
			since = page[len(page)-1].Timestamp
		}
		select {
		case <-ctx.Done():
			return fmt.Errorf("the index had not reached %s when the run ended", end.UTC().Format(time.RFC3339))
		case <-time.After(indexRetryPause):
		}
	}
}

// indexPage fetches one page of the index starting at since (inclusive).
func indexPage(ctx context.Context, client *http.Client, index string, since time.Time) ([]indexEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()
	query := url.Values{"since": {since.UTC().Format(time.RFC3339Nano)}, "limit": {strconv.Itoa(indexPageSize)}}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, index+"?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not reach the index: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= http.StatusInternalServerError {
		seconds, _ := strconv.Atoi(resp.Header.Get("Retry-After"))
		return nil, &busyError{status: resp.StatusCode, retryAfter: time.Duration(seconds) * time.Second}
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("the index answered %d", resp.StatusCode)
	}
	var page []indexEntry
	decoder := json.NewDecoder(io.LimitReader(resp.Body, maxIndexPageBytes))
	for {
		var e indexEntry
		err := decoder.Decode(&e)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("the index gave an unreadable answer: %w", err)
		}
		if e.Path == "" || e.Version == "" || e.Timestamp.IsZero() {
			return nil, errors.New("the index gave an entry without a path, version or timestamp")
		}
		page = append(page, e)
	}
	if len(page) > indexPageSize {
		return nil, fmt.Errorf("the index returned %d entries for a page of %d", len(page), indexPageSize)
	}
	return page, nil
}

// escapePath applies the module proxy's case encoding: each upper-case letter
// becomes "!" followed by its lower-case form.
func escapePath(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			b.WriteByte('!')
			b.WriteRune(r + ('a' - 'A'))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
