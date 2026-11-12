package ezclassify

import "testing"

// The declared wire enum values are a 1.0-locked contract: both the public
// constant names and the wire strings they carry must stay stable. A spec
// regen that renames a generated constant breaks this test instead of
// silently changing the public surface.
func TestWireEnumConstantsCarryTheLockedWireValues(t *testing.T) {
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"AssertionConfidenceCertain", string(AssertionConfidenceCertain), "certain"},
		{"AssertionConfidenceInferred", string(AssertionConfidenceInferred), "inferred"},
		{"AssertionConfidenceRepaired", string(AssertionConfidenceRepaired), "repaired"},
		{"ConceptNoticeTypeAdmonition", string(ConceptNoticeTypeAdmonition), "admonition"},
		{"ConceptNoticeTypeWarning", string(ConceptNoticeTypeWarning), "warning"},
		{"ProcessingRolePart", string(ProcessingRolePart), "part"},
		{"ProcessingRoleWhole", string(ProcessingRoleWhole), "whole"},
		{"WarningKindDeprecation", string(WarningKindDeprecation), "deprecation"},
		{"WarningKindRule", string(WarningKindRule), "rule"},
		{"OverallVerdictPass", string(OverallVerdictPass), "pass"},
		{"OverallVerdictMixed", string(OverallVerdictMixed), "mixed"},
		{"OverallVerdictInadequateCoverage", string(OverallVerdictInadequateCoverage), "inadequate-coverage"},
		{"OverallVerdictProceduralGap", string(OverallVerdictProceduralGap), "procedural-gap"},
		{"DocumentResponseProfileSummary", string(DocumentResponseProfileSummary), "summary"},
		{"DocumentResponseProfileStandard", string(DocumentResponseProfileStandard), "standard"},
		{"DocumentResponseProfileFull", string(DocumentResponseProfileFull), "full"},
		{"GuideCompilationStatusDeterminedInsufficient", string(GuideCompilationStatusDeterminedInsufficient), "determined-insufficient"},
		{"GuideCompilationStatusDeterminedSufficient", string(GuideCompilationStatusDeterminedSufficient), "determined-sufficient"},
		{"GuideCompilationStatusNoFindings", string(GuideCompilationStatusNoFindings), "no-findings"},
		{"GuideCompilationStatusNotEvaluated", string(GuideCompilationStatusNotEvaluated), "not-evaluated"},
		{"GuideCompilationStatusPendingDetermination", string(GuideCompilationStatusPendingDetermination), "pending-determination"},
		{"UnrepresentableFactReasonAmbiguousIdentity", string(UnrepresentableFactReasonAmbiguousIdentity), "ambiguous-identity"},
		{"UnrepresentableFactReasonNoField", string(UnrepresentableFactReasonNoField), "no-field"},
		{"UnrepresentableFactReasonNoView", string(UnrepresentableFactReasonNoView), "no-view"},
		{"UnrepresentableFactReasonReservedLiteral", string(UnrepresentableFactReasonReservedLiteral), "reserved-literal"},
		{"UnrepresentableFactReasonUnframable", string(UnrepresentableFactReasonUnframable), "unframable"},
		{"UnrepresentableFactReasonValueLossy", string(UnrepresentableFactReasonValueLossy), "value-lossy"},
		{"UnrepresentableOutputReasonAmbiguousRepresentation", string(UnrepresentableOutputReasonAmbiguousRepresentation), "ambiguous-representation"},
		{"UnrepresentableOutputReasonUnresolvedContext", string(UnrepresentableOutputReasonUnresolvedContext), "unresolved-context"},
		{"DecisionProposalOriginAgentDrafted", string(DecisionProposalOriginAgentDrafted), "agent-drafted"},
		{"DecisionProposalOriginAnalystAuthored", string(DecisionProposalOriginAnalystAuthored), "analyst-authored"},
		{"DecisionProposalOriginHostImported", string(DecisionProposalOriginHostImported), "host-imported"},
		{"FixActionAdd", string(FixActionAdd), "add"},
		{"FixActionRemove", string(FixActionRemove), "remove"},
		{"FixActionReplace", string(FixActionReplace), "replace"},
		{"FixActionReorder", string(FixActionReorder), "reorder"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %q, want %q", c.name, c.got, c.want)
		}
	}

	if !ConceptNoticeTypeWarning.Valid() || ConceptNoticeType("bogus").Valid() {
		t.Error("ConceptNoticeType.Valid() must accept locked members and reject unknown values")
	}

	// The membership check travels with the alias.
	if !ProcessingRoleWhole.Valid() || ProcessingRole("bogus").Valid() {
		t.Error("Valid() must accept locked members and reject unknown values")
	}
	if !DocumentResponseProfileStandard.Valid() || DocumentResponseProfile("bogus").Valid() {
		t.Error("DocumentResponseProfile.Valid() must accept locked members and reject unknown values")
	}
	if !GuideCompilationStatusPendingDetermination.Valid() || GuideCompilationStatus("bogus").Valid() {
		t.Error("GuideCompilationStatus.Valid() must accept locked members and reject unknown values")
	}
	if !UnrepresentableFactReasonNoView.Valid() || UnrepresentableFactReason("unresolved-context").Valid() {
		t.Error("UnrepresentableFactReason.Valid() must accept its own members and reject an output reason")
	}
	if !UnrepresentableOutputReasonUnresolvedContext.Valid() || UnrepresentableOutputReason("no-view").Valid() {
		t.Error("UnrepresentableOutputReason.Valid() must accept its own members and reject a fact reason")
	}
	if !DecisionProposalOriginHostImported.Valid() || DecisionProposalOrigin("bogus").Valid() {
		t.Error("DecisionProposalOrigin.Valid() must accept locked members and reject unknown values")
	}
	if !FixActionReplace.Valid() || FixAction("bogus").Valid() {
		t.Error("FixAction.Valid() must accept locked members and reject unknown values")
	}
}
