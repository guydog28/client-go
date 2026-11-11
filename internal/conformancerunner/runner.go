// Package conformancerunner provides the engine-slice conformance runner for
// EzClassify Go SDK integrations.
//
// It reads spec plans from a specDir, seeds a local HTTP replayer with matching
// recorded responses, sends each plan case's request to the replayer, and
// subset-matches the recorded response against the plan's expectation.
//
// Usage:
//
//	func TestConformance(t *testing.T) {
//	    _, thisFile, _, _ := runtime.Caller(0)
//	    specDir := filepath.Join(filepath.Dir(thisFile), "spec")
//	    conformancerunner.RunConformance(t, specDir)
//	}
package conformancerunner

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ezclassify/client-go/internal/replayer"
	"github.com/ezclassify/client-go/internal/subsetmatch"
)

// RunConformance runs the full engine-slice conformance suite against specDir.
// specDir must contain plans/v1/, recordings/v1/, and fixtures/v1/
// subdirectories, typically the spec/ directory of the consuming SDK repo.
func RunConformance(t *testing.T, specDir string) {
	t.Helper()

	fixtureDir := filepath.Join(specDir, "fixtures", "v1")
	plansDir := filepath.Join(specDir, "plans", "v1")
	recordingsDir := filepath.Join(specDir, "recordings", "v1")

	allRecordings := loadAllRecordings(t, recordingsDir)

	planEntries, err := os.ReadDir(plansDir)
	if err != nil {
		t.Fatalf("read plans dir %s: %v", plansDir, err)
	}

	type planFile struct {
		name  string
		slice string
		cases []map[string]interface{}
	}

	var plans []planFile
	for _, entry := range planEntries {
		if entry.IsDir() || entry.Name() == "plan-schema.json" {
			continue
		}
		b, err := os.ReadFile(filepath.Join(plansDir, entry.Name()))
		if err != nil {
			t.Fatalf("read plan %s: %v", entry.Name(), err)
		}
		var doc map[string]interface{}
		if err := json.Unmarshal(b, &doc); err != nil {
			t.Fatalf("parse plan %s: %v", entry.Name(), err)
		}
		slice, _ := doc["slice"].(string)
		var cases []map[string]interface{}
		if arr, ok := doc["cases"].([]interface{}); ok {
			for _, c := range arr {
				if m, ok := c.(map[string]interface{}); ok {
					cases = append(cases, m)
				}
			}
		}
		// Registry fan-out vectors exercise the embedded engine's ambient API,
		// which is deliberately absent from the Context-only HTTP contract.
		if entry.Name() == "registry.json" {
			if slice != "engine" || len(cases) != 5 {
				t.Fatalf("engine-only registry plan changed; review its HTTP ownership")
			}
			for _, tc := range cases {
				method, _ := tc["method"].(string)
				if method != "registryProcessAll" && method != "registrySuggestAcross" {
					t.Fatalf("unexpected method %q in engine-only registry plan", method)
				}
			}
			continue
		}
		plans = append(plans, planFile{name: entry.Name(), slice: slice, cases: cases})
	}

	// Pre-populate the replayer: for each engine-slice case, map
	// recording-key → replayer-key so the replayer returns the right response.
	replayerMap := make(map[string]replayer.Recording)
	for _, pf := range plans {
		if pf.slice != "engine" {
			continue
		}
		for _, tc := range pf.cases {
			method, _ := tc["method"].(string)
			args := toIfaceSlice(tc["args"])
			fixture := tc["fixture"]
			mode, _ := tc["mode"].(string)
			if mode == "" {
				mode = "fromJSON"
			}
			ctxID, _ := tc["ctxId"].(string)
			ctxVersion, _ := tc["ctxVersion"].(string)
			guideFixture, _ := tc["guideFixture"].(string)
			accessFixture, _ := tc["accessFixture"].(string)

			recKey := computeRecordingKey(method, args, fixture, mode, ctxID, ctxVersion, guideFixture, accessFixture)
			responseJSON, found := allRecordings[recKey]
			if !found {
				continue
			}

			apiPath, body, err := buildPlanBody(method, args, fixture, guideFixture, accessFixture, fixtureDir)
			if err != nil {
				// A broken fixture is spec drift, not an unsupported case:
				// fail the suite with the real cause instead of skipping.
				t.Fatalf("plan %s case %v: %v", pf.name, tc["name"], err)
			}

			rKey := replayer.RecordingKey(apiPath, body)
			if previous, exists := replayerMap[rKey]; exists {
				var previousBody, currentBody interface{}
				if err := json.Unmarshal(previous.Body, &previousBody); err != nil {
					t.Fatalf("parse previous recording: %v", err)
				}
				if err := json.Unmarshal(responseJSON, &currentBody); err != nil {
					t.Fatalf("parse current recording: %v", err)
				}
				if previous.StatusCode != 200 || !reflect.DeepEqual(previousBody, currentBody) {
					t.Fatalf("conflicting recordings for request: plan %s case %v", pf.name, tc["name"])
				}
			}
			replayerMap[rKey] = replayer.Recording{StatusCode: 200, Body: responseJSON}
		}
	}

	r := replayer.New(replayerMap)
	t.Cleanup(r.Close)

	for _, pf := range plans {
		pf := pf
		if pf.slice != "engine" {
			continue
		}
		for _, tc := range pf.cases {
			tc := tc
			name, _ := tc["name"].(string)
			t.Run(pf.name+"/"+name, func(t *testing.T) {
				t.Parallel()

				method, _ := tc["method"].(string)
				fixtureRaw := tc["fixture"]
				guideFixture, _ := tc["guideFixture"].(string)
				accessFixture, _ := tc["accessFixture"].(string)
				expect := tc["expect"]
				args := toIfaceSlice(tc["args"])

				actual, runErr := runCase(
					t, r.URL(), fixtureDir,
					method, fixtureRaw, guideFixture, accessFixture, args,
				)
				if runErr != nil {
					t.Fatalf("run: %v", runErr)
				}
				if !subsetmatch.Match(actual, expect) {
					expectJSON, _ := json.MarshalIndent(expect, "", "  ")
					actualJSON, _ := json.MarshalIndent(actual, "", "  ")
					t.Errorf("subset mismatch\n  expect: %s\n  actual: %s", expectJSON, actualJSON)
				}
			})
		}
	}
}
