package subsetmatch_test

import (
	"encoding/json"
	"testing"

	"github.com/ezclassify/sdk-go/internal/subsetmatch"
)

// mustJSON decodes a JSON literal into the interface{} shapes json.Unmarshal
// produces, which is what Match is specified against.
func mustJSON(t *testing.T, s string) interface{} {
	t.Helper()
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err != nil {
		t.Fatalf("bad test literal %q: %v", s, err)
	}
	return v
}

// TestMatchJSONValues pins the subset semantics for every JSON value kind:
// scalars must be equal, objects are compared key-by-key with extra actual
// keys allowed, and arrays are prefix-matched with extra actual elements allowed.
func TestMatchJSONValues(t *testing.T) {
	cases := []struct {
		name     string
		actual   string
		expected string
		want     bool
	}{
		{"equal_strings", `"a"`, `"a"`, true},
		{"unequal_strings", `"a"`, `"b"`, false},
		{"equal_numbers", `1.5`, `1.5`, true},
		{"unequal_numbers", `1`, `2`, false},
		{"equal_bools", `true`, `true`, true},
		{"unequal_bools", `true`, `false`, false},
		{"null_matches_null", `null`, `null`, true},
		{"null_vs_value", `"x"`, `null`, false},
		{"type_mismatch_string_number", `"1"`, `1`, false},
		{"type_mismatch_bool_string", `true`, `"true"`, false},

		{"object_exact", `{"a":1,"b":"x"}`, `{"a":1,"b":"x"}`, true},
		{"object_actual_superset", `{"a":1,"b":"x","extra":true}`, `{"a":1}`, true},
		{"object_missing_key", `{"a":1}`, `{"a":1,"b":2}`, false},
		{"object_wrong_value", `{"a":1}`, `{"a":2}`, false},
		{"object_vs_array", `[1]`, `{"a":1}`, false},
		{"nested_object", `{"outer":{"inner":{"k":"v"},"other":1}}`, `{"outer":{"inner":{"k":"v"}}}`, true},
		{"nested_object_mismatch", `{"outer":{"inner":{"k":"v"}}}`, `{"outer":{"inner":{"k":"w"}}}`, false},

		{"array_exact", `[1,2,3]`, `[1,2,3]`, true},
		{"array_actual_longer", `[1,2,3,4]`, `[1,2]`, true},
		{"array_actual_shorter", `[1]`, `[1,2]`, false},
		{"array_element_mismatch", `[1,3]`, `[1,2]`, false},
		{"array_vs_object", `{"0":1}`, `[1]`, false},
		{"array_of_objects", `[{"a":1,"b":2},{"c":3}]`, `[{"a":1},{"c":3}]`, true},
		{"empty_expected_object", `{"anything":1}`, `{}`, true},
		{"empty_expected_array", `[1,2]`, `[]`, true},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual := mustJSON(t, tc.actual)
			expected := mustJSON(t, tc.expected)
			if got := subsetmatch.Match(actual, expected); got != tc.want {
				t.Errorf("Match(%s, %s) = %v, want %v", tc.actual, tc.expected, got, tc.want)
			}
		})
	}
}

// TestMatchIntActualAgainstFloatExpected pins the numeric bridge: hand-built
// Go values may carry int where json.Unmarshal would produce float64, and an
// expected float64 must still match it.
func TestMatchIntActualAgainstFloatExpected(t *testing.T) {
	if !subsetmatch.Match(3, 3.0) {
		t.Error("Match(int 3, float64 3.0) = false, want true")
	}
	if subsetmatch.Match(4, 3.0) {
		t.Error("Match(int 4, float64 3.0) = true, want false")
	}
}

// TestMatchNonJSONExpectedFallsBackToDeepEqual pins the fallback arm for
// expected values that json.Unmarshal never produces (e.g. untyped ints in
// hand-built Go fixtures).
func TestMatchNonJSONExpectedFallsBackToDeepEqual(t *testing.T) {
	if !subsetmatch.Match(3, 3) {
		t.Error("Match(int 3, int 3) = false, want true")
	}
	if subsetmatch.Match(3, 4) {
		t.Error("Match(int 3, int 4) = true, want false")
	}
}
