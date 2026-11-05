// Conformance runner for the EzClassify Go SDK.
//
// Reads every engine-slice plan from spec/plans/v1/, pre-populates a
// local HTTP replayer with the matching recorded responses, and runs each plan
// case against the replayer. Server-slice plans (ctx* methods) belong to a separate lane;
// those run against a live rest-api via `make conformance-server URL=...`.
//
// Run with:
//
//	cd sdk-go && go test -run TestConformance ./...
package ezclassify_test

import (
	"path/filepath"
	"runtime"
	"testing"

	conformancerunner "github.com/ezclassify/sdk-go/internal/conformancerunner"
)

// TestConformance runs the engine-slice conformance suite against the local
// spec/ directory.
func TestConformance(t *testing.T) {
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	specDir := filepath.Join(filepath.Dir(thisFile), "spec")
	conformancerunner.RunConformance(t, specDir)
}
