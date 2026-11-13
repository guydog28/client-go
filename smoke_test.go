//go:build live

package ezclassify_test

// Live-server smoke test.
//
// Unlike the replayed conformance suite (which seeds a local HTTP replayer with
// canned responses), this test drives the SDK against a real, running
// EzClassify REST runtime over the wire. Its job is to catch a base-path
// or version break, e.g. the client emitting /api/* against a /v1 server, or a
// stale vendored spec, that a replayer can never surface, because the replayer
// answers whatever path the client happens to send.
//
// The live lane requires EZC_SMOKE_URL to point at a running, licensed and
// content-configured REST runtime (the ezclassify-rest-api binary, or the
// ghcr.io/ezclassify/rest-api image with its license and bundles mounted) that
// serves the Context below:
//
//	EZC_SMOKE_URL=http://localhost:3000 go test -tags live -run TestSmoke .

import (
	"context"
	"errors"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/ezclassify/sdk-go"
)

const (
	smokeContextID = "cedar-valley-protection"
	smokeInput     = "confidential"
	smokeWantOut   = "CONFIDENTIAL"
	// Two of the outputs cedar-valley-protection declares, comma-separated.
	// EZC_SMOKE_OUTPUTS overrides them for a server bound to another Context.
	smokeOutputs = "default,short"
)

func smokeURL(t *testing.T) string {
	t.Helper()
	url := os.Getenv("EZC_SMOKE_URL")
	if url == "" {
		t.Fatal("EZC_SMOKE_URL is required for the live-server smoke lane " +
			"(start the REST runtime and set EZC_SMOKE_URL=http://host:port)")
	}
	return url
}

// TestSmokeHealth pins that Ping reaches GET /v1/health on a live 1.0.0 server
// and reads back a non-empty wire version. A URL-prefix break shows up here as
// a 404 wrapped error, which a replayed test cannot reproduce.
func TestSmokeHealth(t *testing.T) {
	url := smokeURL(t)

	client := ezclassify.New(url)
	ping, err := client.Ping(context.Background())
	if err != nil {
		t.Fatalf("Ping(%s): %v", url, err)
	}
	if ping.EngineVersion == "" {
		t.Fatal("Ping returned an empty EngineVersion; health payload missing engineVersion")
	}
	t.Logf("live server engine version: %s", ping.EngineVersion)
}

// TestSmokeProcess runs real Context-bound process operations: one that names
// outputs, which must come back as exactly those renders, and one that names
// none, which must come back with no renders at all.
func TestSmokeProcess(t *testing.T) {
	url := smokeURL(t)
	ctx := context.Background()
	client := ezclassify.New(url, ezclassify.WithCredentials(smokeCredentials()))

	contextID := envOr("EZC_SMOKE_CONTEXT_ID", smokeContextID)
	input := envOr("EZC_SMOKE_INPUT", smokeInput)
	want := envOr("EZC_SMOKE_EXPECTED", smokeWantOut)
	outputs := splitList(envOr("EZC_SMOKE_OUTPUTS", smokeOutputs))
	cx := ezclassify.ConnectWith(client, contextID, "latest")
	res, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: input, Outputs: &outputs})
	if err != nil {
		t.Fatalf("ctx Scheme.Process: %v", err)
	}
	if !res.Result.Valid {
		t.Errorf("ctx Scheme.Process(%q): Valid = false, want true", input)
	}
	if res.Result.Rendered != want {
		t.Errorf("ctx Scheme.Process(%q): Rendered = %q, want %q", input, res.Result.Rendered, want)
	}
	if res.Result.Recognition == nil || res.Result.Recognition.Status != ezclassify.RecognitionStatusComplete {
		t.Errorf("ctx Scheme.Process(%q): Recognition = %+v, want complete", input, res.Result.Recognition)
	}
	if got := renderKeys(res.Result.Renders); strings.Join(got, ",") != strings.Join(sortedCopy(outputs), ",") {
		t.Errorf("ctx Scheme.Process(%q, Outputs %q): Renders keys = %q, want exactly the outputs named", input, outputs, got)
	}

	bare, err := cx.Scheme.Process(ctx, ezclassify.CtxSchemeProcessRequest{Input: input})
	if err != nil {
		t.Fatalf("ctx Scheme.Process with no outputs: %v", err)
	}
	if len(bare.Result.Renders) != 0 {
		t.Errorf("ctx Scheme.Process(%q) with no outputs: Renders keys = %q, want none", input, renderKeys(bare.Result.Renders))
	}

	format := "authority_block"
	preservation := ezclassify.RenderPreservationFull
	_, err = cx.Scheme.Render(ctx, ezclassify.CtxSchemeRenderRequest{
		Assertions:   res.Result.Assertions,
		OutputFormat: &format,
		Preservation: &preservation,
	})
	var refusal *ezclassify.APIError
	if !errors.As(err, &refusal) || refusal.StatusCode != 400 || refusal.Response == nil ||
		refusal.Response.Code == nil || *refusal.Response.Code != "INVALID_INPUT" {
		t.Fatalf("full-preservation refusal = %T(%v), want typed 400 INVALID_INPUT", err, err)
	}
	if cx.ResolvedVersion() == "" {
		t.Error("ctx Scheme.Process did not populate ResolvedVersion from the X-EzClassify-Resolved-Version header")
	}
}

func smokeCredentials() ezclassify.RestClientCredentials {
	return ezclassify.RestClientCredentials{
		AccessToken:  os.Getenv("EZC_SMOKE_ACCESS_TOKEN"),
		ContextToken: os.Getenv("EZC_SMOKE_CONTEXT_TOKEN"),
	}
}

func envOr(name, fallback string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return fallback
}

// splitList reads a comma-separated setting, dropping blanks.
func splitList(value string) []string {
	var items []string
	for _, item := range strings.Split(value, ",") {
		if item = strings.TrimSpace(item); item != "" {
			items = append(items, item)
		}
	}
	return items
}

func renderKeys(renders ezclassify.RenderResultsByFormat) []string {
	keys := make([]string, 0, len(renders))
	for key := range renders {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedCopy(items []string) []string {
	sorted := append([]string(nil), items...)
	sort.Strings(sorted)
	return sorted
}
