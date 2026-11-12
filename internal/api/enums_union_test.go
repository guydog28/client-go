package api

import (
	"encoding/json"
	"testing"
)

// TestEnumsValid covers every Valid() method on every string-enum type in the
// generated package. Each test exercises every known constant (want true) plus
// one unknown value (want false) to hit the default arm.
func TestEnumsValid(t *testing.T) {
	t.Run("AssertionConfidence", func(t *testing.T) {
		for _, e := range []AssertionConfidence{
			Certain,
			Inferred,
			Repaired,
		} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if AssertionConfidence("bogus").Valid() {
			t.Error(`AssertionConfidence("bogus").Valid() = true`)
		}
	})

	t.Run("DocumentOverallVerdict", func(t *testing.T) {
		for _, e := range []DocumentOverallVerdict{
			DocumentOverallVerdictInadequateCoverage,
			DocumentOverallVerdictMixed,
			DocumentOverallVerdictPass,
			DocumentOverallVerdictProceduralGap,
		} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if DocumentOverallVerdict("bogus").Valid() {
			t.Error(`DocumentOverallVerdict("bogus").Valid() = true`)
		}
	})

	t.Run("FixAction", func(t *testing.T) {
		for _, e := range []FixAction{Add, Remove, Reorder, Replace} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if FixAction("bogus").Valid() {
			t.Error(`FixAction("bogus").Valid() = true`)
		}
	})

	t.Run("ConceptNoticeType", func(t *testing.T) {
		for _, e := range []ConceptNoticeType{
			ConceptNoticeTypeAdmonition, ConceptNoticeTypeWarning,
		} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if ConceptNoticeType("bogus").Valid() {
			t.Error(`ConceptNoticeType("bogus").Valid() = true`)
		}
	})

	t.Run("ProcessingRole", func(t *testing.T) {
		for _, e := range []ProcessingRole{
			Part, Whole,
		} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if ProcessingRole("bogus").Valid() {
			t.Error(`ProcessingRole("bogus").Valid() = true`)
		}
	})

	t.Run("RuleConditionTarget", func(t *testing.T) {
		// The generator prefixes these constants with their enum type.
		for _, e := range []RuleConditionTarget{RuleConditionTargetContext, RuleConditionTargetContributors, RuleConditionTargetSelf} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if RuleConditionTarget("bogus").Valid() {
			t.Error(`RuleConditionTarget("bogus").Valid() = true`)
		}
	})

	t.Run("SuggestionSource", func(t *testing.T) {
		for _, e := range []SuggestionSource{SuggestionSourceAlias, SuggestionSourceParse, SuggestionSourcePreset} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if SuggestionSource("bogus").Valid() {
			t.Error(`SuggestionSource("bogus").Valid() = true`)
		}
	})

	t.Run("WarningKind", func(t *testing.T) {
		for _, e := range []WarningKind{
			WarningKindDeprecation, WarningKindNotice, WarningKindRule,
		} {
			if !e.Valid() {
				t.Errorf("%q.Valid() = false", e)
			}
		}
		if WarningKind("bogus").Valid() {
			t.Error(`WarningKind("bogus").Valid() = true`)
		}
	})
}

// TestRuleConditionValueUnion exercises the full From/As/Merge/Marshal/Unmarshal
// surface for the RuleCondition_Value discriminated union (string | float32).
func TestRuleConditionValueUnion(t *testing.T) {
	t.Run("string/FromAs", func(t *testing.T) {
		var v RuleCondition_Value
		if err := v.FromRuleConditionValue0("hello"); err != nil {
			t.Fatalf("From: %v", err)
		}
		got, err := v.AsRuleConditionValue0()
		if err != nil {
			t.Fatalf("As: %v", err)
		}
		if got != "hello" {
			t.Errorf("got %q, want hello", got)
		}
	})

	t.Run("float32/FromAs", func(t *testing.T) {
		var v RuleCondition_Value
		if err := v.FromRuleConditionValue1(2.5); err != nil {
			t.Fatalf("From: %v", err)
		}
		got, err := v.AsRuleConditionValue1()
		if err != nil {
			t.Fatalf("As: %v", err)
		}
		if got != 2.5 {
			t.Errorf("got %v, want 2.5", got)
		}
	})

	t.Run("string/Merge", func(t *testing.T) {
		var v RuleCondition_Value
		if err := v.MergeRuleConditionValue0("x"); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("float32/Merge", func(t *testing.T) {
		var v RuleCondition_Value
		_ = v.FromRuleConditionValue1(1.0)
		if err := v.MergeRuleConditionValue1(2.0); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		var v RuleCondition_Value
		_ = v.FromRuleConditionValue0("abc")
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if string(b) != `"abc"` {
			t.Errorf("got %s, want %q", b, `"abc"`)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var v RuleCondition_Value
		if err := json.Unmarshal([]byte(`"world"`), &v); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		got, err := v.AsRuleConditionValue0()
		if err != nil {
			t.Fatalf("As after Unmarshal: %v", err)
		}
		if got != "world" {
			t.Errorf("got %q, want world", got)
		}
	})
}

// TestRuleWhenUnion exercises the full From/As/Merge/Marshal/Unmarshal surface
// for the RuleWhen discriminated union (string | []string | struct).
func TestRuleWhenUnion(t *testing.T) {
	t.Run("string/FromAs", func(t *testing.T) {
		var w RuleWhen
		if err := w.FromRuleWhen0("concept-a"); err != nil {
			t.Fatalf("From: %v", err)
		}
		got, err := w.AsRuleWhen0()
		if err != nil {
			t.Fatalf("As: %v", err)
		}
		if got != "concept-a" {
			t.Errorf("got %q, want concept-a", got)
		}
	})

	t.Run("slice/FromAs", func(t *testing.T) {
		var w RuleWhen
		if err := w.FromRuleWhen1([]string{"a", "b"}); err != nil {
			t.Fatalf("From: %v", err)
		}
		got, err := w.AsRuleWhen1()
		if err != nil {
			t.Fatalf("As: %v", err)
		}
		if len(got) != 2 || got[0] != "a" {
			t.Errorf("got %v, want [a b]", got)
		}
	})

	t.Run("struct/FromAs", func(t *testing.T) {
		var w RuleWhen
		if err := w.FromRuleWhenQuantifier(RuleWhenQuantifier{}); err != nil {
			t.Fatalf("From: %v", err)
		}
		if _, err := w.AsRuleWhenQuantifier(); err != nil {
			t.Fatalf("As: %v", err)
		}
	})

	t.Run("string/Merge", func(t *testing.T) {
		var w RuleWhen
		if err := w.MergeRuleWhen0("y"); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("slice/Merge", func(t *testing.T) {
		var w RuleWhen
		_ = w.FromRuleWhen1([]string{"z"})
		if err := w.MergeRuleWhen1([]string{"z2"}); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("struct/Merge", func(t *testing.T) {
		var w RuleWhen
		if err := w.MergeRuleWhenQuantifier(RuleWhenQuantifier{}); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		var w RuleWhen
		_ = w.FromRuleWhen0("m")
		b, err := json.Marshal(w)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if string(b) != `"m"` {
			t.Errorf("got %s, want %q", b, `"m"`)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var w RuleWhen
		if err := json.Unmarshal([]byte(`"concept-b"`), &w); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		got, err := w.AsRuleWhen0()
		if err != nil {
			t.Fatalf("As: %v", err)
		}
		if got != "concept-b" {
			t.Errorf("got %q, want concept-b", got)
		}
	})
}

// TestRuleWhenNodeUnion exercises the full From/As/Merge/Marshal/Unmarshal surface
// for the RuleWhenNode discriminated union (RuleCondition | RuleWhenQuantifier).
func TestRuleWhenNodeUnion(t *testing.T) {
	t.Run("condition/FromAs", func(t *testing.T) {
		var n RuleWhenNode
		if err := n.FromRuleCondition(RuleCondition{}); err != nil {
			t.Fatalf("From: %v", err)
		}
		if _, err := n.AsRuleCondition(); err != nil {
			t.Fatalf("As: %v", err)
		}
	})

	t.Run("node1/FromAs", func(t *testing.T) {
		var n RuleWhenNode
		if err := n.FromRuleWhenQuantifier(RuleWhenQuantifier{}); err != nil {
			t.Fatalf("From: %v", err)
		}
		if _, err := n.AsRuleWhenQuantifier(); err != nil {
			t.Fatalf("As: %v", err)
		}
	})

	t.Run("condition/Merge", func(t *testing.T) {
		var n RuleWhenNode
		if err := n.MergeRuleCondition(RuleCondition{}); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("node1/Merge", func(t *testing.T) {
		var n RuleWhenNode
		_ = n.FromRuleWhenQuantifier(RuleWhenQuantifier{})
		if err := n.MergeRuleWhenQuantifier(RuleWhenQuantifier{}); err != nil {
			t.Fatalf("Merge: %v", err)
		}
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		var n RuleWhenNode
		_ = n.FromRuleCondition(RuleCondition{})
		if _, err := json.Marshal(n); err != nil {
			t.Fatalf("Marshal: %v", err)
		}
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var n RuleWhenNode
		if err := json.Unmarshal([]byte(`{}`), &n); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if _, err := n.AsRuleCondition(); err != nil {
			t.Fatalf("As: %v", err)
		}
	})
}
