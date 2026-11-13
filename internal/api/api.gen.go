// Code generated from openapi.yaml@b2bca608ea6203a0110078b19a8f566f9b572e9457b1847e390ccb9667bf85ff by oapi-codegen@v2.8.0@h1:s4hxMxuqtR8jPzXkBTtFwY/SBuj3gEAYikmbBSdtLMM=. DO NOT EDIT.
// Package api provides primitives to interact with the openapi HTTP API.
package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// Defines values for AccessOperandSource.
const (
	AccessOperandSourceAssertion AccessOperandSource = "assertion"
	AccessOperandSourceContext   AccessOperandSource = "context"
	AccessOperandSourceLiteral   AccessOperandSource = "literal"
	AccessOperandSourceSubject   AccessOperandSource = "subject"
)

// Valid indicates whether the value is a known member of the AccessOperandSource enum.
func (e AccessOperandSource) Valid() bool {
	switch e {
	case AccessOperandSourceAssertion:
		return true
	case AccessOperandSourceContext:
		return true
	case AccessOperandSourceLiteral:
		return true
	case AccessOperandSourceSubject:
		return true
	default:
		return false
	}
}

// Defines values for ArtifactKind.
const (
	ArtifactKindAccess    ArtifactKind = "access"
	ArtifactKindContext   ArtifactKind = "context"
	ArtifactKindCrosswalk ArtifactKind = "crosswalk"
	ArtifactKindGuide     ArtifactKind = "guide"
	ArtifactKindLayer     ArtifactKind = "layer"
	ArtifactKindScheme    ArtifactKind = "scheme"
)

// Valid indicates whether the value is a known member of the ArtifactKind enum.
func (e ArtifactKind) Valid() bool {
	switch e {
	case ArtifactKindAccess:
		return true
	case ArtifactKindContext:
		return true
	case ArtifactKindCrosswalk:
		return true
	case ArtifactKindGuide:
		return true
	case ArtifactKindLayer:
		return true
	case ArtifactKindScheme:
		return true
	default:
		return false
	}
}

// Defines values for AssertionConfidence.
const (
	Certain  AssertionConfidence = "certain"
	Inferred AssertionConfidence = "inferred"
	Repaired AssertionConfidence = "repaired"
)

// Valid indicates whether the value is a known member of the AssertionConfidence enum.
func (e AssertionConfidence) Valid() bool {
	switch e {
	case Certain:
		return true
	case Inferred:
		return true
	case Repaired:
		return true
	default:
		return false
	}
}

// Defines values for CalendarDateSpelling.
const (
	YYYYMMDD  CalendarDateSpelling = "YYYYMMDD"
	YYYYMMDD1 CalendarDateSpelling = "YYYY-MM-DD"
)

// Valid indicates whether the value is a known member of the CalendarDateSpelling enum.
func (e CalendarDateSpelling) Valid() bool {
	switch e {
	case YYYYMMDD:
		return true
	case YYYYMMDD1:
		return true
	default:
		return false
	}
}

// Defines values for CompositeEncodingFieldType.
const (
	Concept CompositeEncodingFieldType = "concept"
	Text    CompositeEncodingFieldType = "text"
)

// Valid indicates whether the value is a known member of the CompositeEncodingFieldType enum.
func (e CompositeEncodingFieldType) Valid() bool {
	switch e {
	case Concept:
		return true
	case Text:
		return true
	default:
		return false
	}
}

// Defines values for ConceptNoticeType.
const (
	ConceptNoticeTypeAdmonition ConceptNoticeType = "admonition"
	ConceptNoticeTypeWarning    ConceptNoticeType = "warning"
)

// Valid indicates whether the value is a known member of the ConceptNoticeType enum.
func (e ConceptNoticeType) Valid() bool {
	switch e {
	case ConceptNoticeTypeAdmonition:
		return true
	case ConceptNoticeTypeWarning:
		return true
	default:
		return false
	}
}

// Defines values for ContextSchemeRole.
const (
	ContextSchemeRoleAdditional      ContextSchemeRole = "additional"
	ContextSchemeRoleCrosswalkSource ContextSchemeRole = "crosswalk-source"
	ContextSchemeRoleCrosswalkTarget ContextSchemeRole = "crosswalk-target"
	ContextSchemeRolePrimary         ContextSchemeRole = "primary"
)

// Valid indicates whether the value is a known member of the ContextSchemeRole enum.
func (e ContextSchemeRole) Valid() bool {
	switch e {
	case ContextSchemeRoleAdditional:
		return true
	case ContextSchemeRoleCrosswalkSource:
		return true
	case ContextSchemeRoleCrosswalkTarget:
		return true
	case ContextSchemeRolePrimary:
		return true
	default:
		return false
	}
}

// Defines values for CrosswalkCompletedTranslationStatus.
const (
	CrosswalkCompletedTranslationStatusComplete      CrosswalkCompletedTranslationStatus = "complete"
	CrosswalkCompletedTranslationStatusInvalidSource CrosswalkCompletedTranslationStatus = "invalid-source"
	CrosswalkCompletedTranslationStatusPartial       CrosswalkCompletedTranslationStatus = "partial"
)

// Valid indicates whether the value is a known member of the CrosswalkCompletedTranslationStatus enum.
func (e CrosswalkCompletedTranslationStatus) Valid() bool {
	switch e {
	case CrosswalkCompletedTranslationStatusComplete:
		return true
	case CrosswalkCompletedTranslationStatusInvalidSource:
		return true
	case CrosswalkCompletedTranslationStatusPartial:
		return true
	default:
		return false
	}
}

// Defines values for CrosswalkDisposition.
const (
	Carry  CrosswalkDisposition = "carry"
	Forbid CrosswalkDisposition = "forbid"
	Omit   CrosswalkDisposition = "omit"
	Review CrosswalkDisposition = "review"
)

// Valid indicates whether the value is a known member of the CrosswalkDisposition enum.
func (e CrosswalkDisposition) Valid() bool {
	switch e {
	case Carry:
		return true
	case Forbid:
		return true
	case Omit:
		return true
	case Review:
		return true
	default:
		return false
	}
}

// Defines values for CrosswalkSelectionStrategy.
const (
	FirstMatch  CrosswalkSelectionStrategy = "first-match"
	MostMatches CrosswalkSelectionStrategy = "most-matches"
)

// Valid indicates whether the value is a known member of the CrosswalkSelectionStrategy enum.
func (e CrosswalkSelectionStrategy) Valid() bool {
	switch e {
	case FirstMatch:
		return true
	case MostMatches:
		return true
	default:
		return false
	}
}

// Defines values for DecisionProposalOrigin.
const (
	AgentDrafted    DecisionProposalOrigin = "agent-drafted"
	AnalystAuthored DecisionProposalOrigin = "analyst-authored"
	HostImported    DecisionProposalOrigin = "host-imported"
)

// Valid indicates whether the value is a known member of the DecisionProposalOrigin enum.
func (e DecisionProposalOrigin) Valid() bool {
	switch e {
	case AgentDrafted:
		return true
	case AnalystAuthored:
		return true
	case HostImported:
		return true
	default:
		return false
	}
}

// Defines values for DocumentInputErrorDetail.
const (
	ConflictingGuideDetermination    DocumentInputErrorDetail = "conflicting-guide-determination"
	ConflictingManualSelection       DocumentInputErrorDetail = "conflicting-manual-selection"
	EmptyDocument                    DocumentInputErrorDetail = "empty-document"
	InvalidUnicode                   DocumentInputErrorDetail = "invalid-unicode"
	ManualSelectionDuplicatesFinding DocumentInputErrorDetail = "manual-selection-duplicates-finding"
	StaleGuideDetermination          DocumentInputErrorDetail = "stale-guide-determination"
	StaleManualSelection             DocumentInputErrorDetail = "stale-manual-selection"
)

// Valid indicates whether the value is a known member of the DocumentInputErrorDetail enum.
func (e DocumentInputErrorDetail) Valid() bool {
	switch e {
	case ConflictingGuideDetermination:
		return true
	case ConflictingManualSelection:
		return true
	case EmptyDocument:
		return true
	case InvalidUnicode:
		return true
	case ManualSelectionDuplicatesFinding:
		return true
	case StaleGuideDetermination:
		return true
	case StaleManualSelection:
		return true
	default:
		return false
	}
}

// Defines values for DocumentOverallVerdict.
const (
	DocumentOverallVerdictInadequateCoverage DocumentOverallVerdict = "inadequate-coverage"
	DocumentOverallVerdictMixed              DocumentOverallVerdict = "mixed"
	DocumentOverallVerdictPass               DocumentOverallVerdict = "pass"
	DocumentOverallVerdictProceduralGap      DocumentOverallVerdict = "procedural-gap"
)

// Valid indicates whether the value is a known member of the DocumentOverallVerdict enum.
func (e DocumentOverallVerdict) Valid() bool {
	switch e {
	case DocumentOverallVerdictInadequateCoverage:
		return true
	case DocumentOverallVerdictMixed:
		return true
	case DocumentOverallVerdictPass:
		return true
	case DocumentOverallVerdictProceduralGap:
		return true
	default:
		return false
	}
}

// Defines values for DocumentResponseProfile.
const (
	DocumentResponseProfileFull     DocumentResponseProfile = "full"
	DocumentResponseProfileStandard DocumentResponseProfile = "standard"
	DocumentResponseProfileSummary  DocumentResponseProfile = "summary"
)

// Valid indicates whether the value is a known member of the DocumentResponseProfile enum.
func (e DocumentResponseProfile) Valid() bool {
	switch e {
	case DocumentResponseProfileFull:
		return true
	case DocumentResponseProfileStandard:
		return true
	case DocumentResponseProfileSummary:
		return true
	default:
		return false
	}
}

// Defines values for FixAction.
const (
	Add     FixAction = "add"
	Remove  FixAction = "remove"
	Reorder FixAction = "reorder"
	Replace FixAction = "replace"
)

// Valid indicates whether the value is a known member of the FixAction enum.
func (e FixAction) Valid() bool {
	switch e {
	case Add:
		return true
	case Remove:
		return true
	case Reorder:
		return true
	case Replace:
		return true
	default:
		return false
	}
}

// Defines values for FormatHeaderPosition.
const (
	Prefix FormatHeaderPosition = "prefix"
	Suffix FormatHeaderPosition = "suffix"
)

// Valid indicates whether the value is a known member of the FormatHeaderPosition enum.
func (e FormatHeaderPosition) Valid() bool {
	switch e {
	case Prefix:
		return true
	case Suffix:
		return true
	default:
		return false
	}
}

// Defines values for GuideCompilationStatus.
const (
	DeterminedInsufficient GuideCompilationStatus = "determined-insufficient"
	DeterminedSufficient   GuideCompilationStatus = "determined-sufficient"
	NoFindings             GuideCompilationStatus = "no-findings"
	NotEvaluated           GuideCompilationStatus = "not-evaluated"
	PendingDetermination   GuideCompilationStatus = "pending-determination"
)

// Valid indicates whether the value is a known member of the GuideCompilationStatus enum.
func (e GuideCompilationStatus) Valid() bool {
	switch e {
	case DeterminedInsufficient:
		return true
	case DeterminedSufficient:
		return true
	case NoFindings:
		return true
	case NotEvaluated:
		return true
	case PendingDetermination:
		return true
	default:
		return false
	}
}

// Defines values for GuideFactKind.
const (
	AcceptedFinding  GuideFactKind = "accepted-finding"
	ManualSelection  GuideFactKind = "manual-selection"
	ProjectedFinding GuideFactKind = "projected-finding"
)

// Valid indicates whether the value is a known member of the GuideFactKind enum.
func (e GuideFactKind) Valid() bool {
	switch e {
	case AcceptedFinding:
		return true
	case ManualSelection:
		return true
	case ProjectedFinding:
		return true
	default:
		return false
	}
}

// Defines values for GuideFindingDecision.
const (
	Accepted GuideFindingDecision = "accepted"
	Rejected GuideFindingDecision = "rejected"
)

// Valid indicates whether the value is a known member of the GuideFindingDecision enum.
func (e GuideFindingDecision) Valid() bool {
	switch e {
	case Accepted:
		return true
	case Rejected:
		return true
	default:
		return false
	}
}

// Defines values for GuideWhenOperator.
const (
	GuideWhenOperatorCountEq       GuideWhenOperator = "count_eq"
	GuideWhenOperatorCountGt       GuideWhenOperator = "count_gt"
	GuideWhenOperatorCountGte      GuideWhenOperator = "count_gte"
	GuideWhenOperatorCountLt       GuideWhenOperator = "count_lt"
	GuideWhenOperatorCountLte      GuideWhenOperator = "count_lte"
	GuideWhenOperatorHasConcept    GuideWhenOperator = "has_concept"
	GuideWhenOperatorHasContent    GuideWhenOperator = "has_content"
	GuideWhenOperatorNotHasConcept GuideWhenOperator = "not_has_concept"
)

// Valid indicates whether the value is a known member of the GuideWhenOperator enum.
func (e GuideWhenOperator) Valid() bool {
	switch e {
	case GuideWhenOperatorCountEq:
		return true
	case GuideWhenOperatorCountGt:
		return true
	case GuideWhenOperatorCountGte:
		return true
	case GuideWhenOperatorCountLt:
		return true
	case GuideWhenOperatorCountLte:
		return true
	case GuideWhenOperatorHasConcept:
		return true
	case GuideWhenOperatorHasContent:
		return true
	case GuideWhenOperatorNotHasConcept:
		return true
	default:
		return false
	}
}

// Defines values for MatchContextSide.
const (
	After  MatchContextSide = "after"
	Before MatchContextSide = "before"
	Both   MatchContextSide = "both"
)

// Valid indicates whether the value is a known member of the MatchContextSide enum.
func (e MatchContextSide) Valid() bool {
	switch e {
	case After:
		return true
	case Before:
		return true
	case Both:
		return true
	default:
		return false
	}
}

// Defines values for MatchNearEmit.
const (
	A        MatchNearEmit = "a"
	B        MatchNearEmit = "b"
	Envelope MatchNearEmit = "envelope"
)

// Valid indicates whether the value is a known member of the MatchNearEmit enum.
func (e MatchNearEmit) Valid() bool {
	switch e {
	case A:
		return true
	case B:
		return true
	case Envelope:
		return true
	default:
		return false
	}
}

// Defines values for MatchValidationKeep.
const (
	MatchValidationKeepFail MatchValidationKeep = "fail"
	MatchValidationKeepPass MatchValidationKeep = "pass"
)

// Valid indicates whether the value is a known member of the MatchValidationKeep enum.
func (e MatchValidationKeep) Valid() bool {
	switch e {
	case MatchValidationKeepFail:
		return true
	case MatchValidationKeepPass:
		return true
	default:
		return false
	}
}

// Defines values for ProcessingRole.
const (
	Part  ProcessingRole = "part"
	Whole ProcessingRole = "whole"
)

// Valid indicates whether the value is a known member of the ProcessingRole enum.
func (e ProcessingRole) Valid() bool {
	switch e {
	case Part:
		return true
	case Whole:
		return true
	default:
		return false
	}
}

// Defines values for ReadinessWarning.
const (
	ApplicationAllowanceExceeded ReadinessWarning = "application-allowance-exceeded"
	AuditDegraded                ReadinessWarning = "audit-degraded"
	DesiredStateStale            ReadinessWarning = "desired-state-stale"
	GenerationRefused            ReadinessWarning = "generation-refused"
)

// Valid indicates whether the value is a known member of the ReadinessWarning enum.
func (e ReadinessWarning) Valid() bool {
	switch e {
	case ApplicationAllowanceExceeded:
		return true
	case AuditDegraded:
		return true
	case DesiredStateStale:
		return true
	case GenerationRefused:
		return true
	default:
		return false
	}
}

// Defines values for ReadyStatus.
const (
	Ready    ReadyStatus = "ready"
	Starting ReadyStatus = "starting"
)

// Valid indicates whether the value is a known member of the ReadyStatus enum.
func (e ReadyStatus) Valid() bool {
	switch e {
	case Ready:
		return true
	case Starting:
		return true
	default:
		return false
	}
}

// Defines values for RecognitionCorrectionReason.
const (
	RecognitionCorrectionReasonAlias     RecognitionCorrectionReason = "alias"
	RecognitionCorrectionReasonCase      RecognitionCorrectionReason = "case"
	RecognitionCorrectionReasonOrder     RecognitionCorrectionReason = "order"
	RecognitionCorrectionReasonSeparator RecognitionCorrectionReason = "separator"
)

// Valid indicates whether the value is a known member of the RecognitionCorrectionReason enum.
func (e RecognitionCorrectionReason) Valid() bool {
	switch e {
	case RecognitionCorrectionReasonAlias:
		return true
	case RecognitionCorrectionReasonCase:
		return true
	case RecognitionCorrectionReasonOrder:
		return true
	case RecognitionCorrectionReasonSeparator:
		return true
	default:
		return false
	}
}

// Defines values for RecognitionLimit.
const (
	InputLength RecognitionLimit = "input-length"
	LiveStates  RecognitionLimit = "live-states"
	Lookups     RecognitionLimit = "lookups"
	Transitions RecognitionLimit = "transitions"
)

// Valid indicates whether the value is a known member of the RecognitionLimit enum.
func (e RecognitionLimit) Valid() bool {
	switch e {
	case InputLength:
		return true
	case LiveStates:
		return true
	case Lookups:
		return true
	case Transitions:
		return true
	default:
		return false
	}
}

// Defines values for RecognitionStatus.
const (
	RecognitionStatusAmbiguous     RecognitionStatus = "ambiguous"
	RecognitionStatusComplete      RecognitionStatus = "complete"
	RecognitionStatusPartial       RecognitionStatus = "partial"
	RecognitionStatusResourceLimit RecognitionStatus = "resource-limit"
	RecognitionStatusUnrecognized  RecognitionStatus = "unrecognized"
)

// Valid indicates whether the value is a known member of the RecognitionStatus enum.
func (e RecognitionStatus) Valid() bool {
	switch e {
	case RecognitionStatusAmbiguous:
		return true
	case RecognitionStatusComplete:
		return true
	case RecognitionStatusPartial:
		return true
	case RecognitionStatusResourceLimit:
		return true
	case RecognitionStatusUnrecognized:
		return true
	default:
		return false
	}
}

// Defines values for RenderFormatMode.
const (
	Block  RenderFormatMode = "block"
	Inline RenderFormatMode = "inline"
)

// Valid indicates whether the value is a known member of the RenderFormatMode enum.
func (e RenderFormatMode) Valid() bool {
	switch e {
	case Block:
		return true
	case Inline:
		return true
	default:
		return false
	}
}

// Defines values for RenderPreservation.
const (
	RenderPreservationFull       RenderPreservation = "full"
	RenderPreservationProjection RenderPreservation = "projection"
)

// Valid indicates whether the value is a known member of the RenderPreservation enum.
func (e RenderPreservation) Valid() bool {
	switch e {
	case RenderPreservationFull:
		return true
	case RenderPreservationProjection:
		return true
	default:
		return false
	}
}

// Defines values for Rollup.
const (
	First        Rollup = "first"
	Intersection Rollup = "intersection"
	Max          Rollup = "max"
	Min          Rollup = "min"
	Pipeline     Rollup = "pipeline"
	Union        Rollup = "union"
)

// Valid indicates whether the value is a known member of the Rollup enum.
func (e Rollup) Valid() bool {
	switch e {
	case First:
		return true
	case Intersection:
		return true
	case Max:
		return true
	case Min:
		return true
	case Pipeline:
		return true
	case Union:
		return true
	default:
		return false
	}
}

// Defines values for RuleConditionOperator.
const (
	RuleConditionOperatorAnyLacksConcept   RuleConditionOperator = "any_lacks_concept"
	RuleConditionOperatorCountEq           RuleConditionOperator = "count_eq"
	RuleConditionOperatorCountGt           RuleConditionOperator = "count_gt"
	RuleConditionOperatorCountGte          RuleConditionOperator = "count_gte"
	RuleConditionOperatorCountLt           RuleConditionOperator = "count_lt"
	RuleConditionOperatorCountLte          RuleConditionOperator = "count_lte"
	RuleConditionOperatorEveryHasConcept   RuleConditionOperator = "every_has_concept"
	RuleConditionOperatorHasConcept        RuleConditionOperator = "has_concept"
	RuleConditionOperatorHasContent        RuleConditionOperator = "has_content"
	RuleConditionOperatorMinRank           RuleConditionOperator = "min_rank"
	RuleConditionOperatorNotHasConcept     RuleConditionOperator = "not_has_concept"
	RuleConditionOperatorSharesConceptWith RuleConditionOperator = "shares_concept_with"
	RuleConditionOperatorSubsetOf          RuleConditionOperator = "subset_of"
)

// Valid indicates whether the value is a known member of the RuleConditionOperator enum.
func (e RuleConditionOperator) Valid() bool {
	switch e {
	case RuleConditionOperatorAnyLacksConcept:
		return true
	case RuleConditionOperatorCountEq:
		return true
	case RuleConditionOperatorCountGt:
		return true
	case RuleConditionOperatorCountGte:
		return true
	case RuleConditionOperatorCountLt:
		return true
	case RuleConditionOperatorCountLte:
		return true
	case RuleConditionOperatorEveryHasConcept:
		return true
	case RuleConditionOperatorHasConcept:
		return true
	case RuleConditionOperatorHasContent:
		return true
	case RuleConditionOperatorMinRank:
		return true
	case RuleConditionOperatorNotHasConcept:
		return true
	case RuleConditionOperatorSharesConceptWith:
		return true
	case RuleConditionOperatorSubsetOf:
		return true
	default:
		return false
	}
}

// Defines values for RuleConditionTarget.
const (
	RuleConditionTargetContext      RuleConditionTarget = "context"
	RuleConditionTargetContributors RuleConditionTarget = "contributors"
	RuleConditionTargetSelf         RuleConditionTarget = "self"
)

// Valid indicates whether the value is a known member of the RuleConditionTarget enum.
func (e RuleConditionTarget) Valid() bool {
	switch e {
	case RuleConditionTargetContext:
		return true
	case RuleConditionTargetContributors:
		return true
	case RuleConditionTargetSelf:
		return true
	default:
		return false
	}
}

// Defines values for RuleContributorQuantifier.
const (
	All  RuleContributorQuantifier = "all"
	Any  RuleContributorQuantifier = "any"
	None RuleContributorQuantifier = "none"
)

// Valid indicates whether the value is a known member of the RuleContributorQuantifier enum.
func (e RuleContributorQuantifier) Valid() bool {
	switch e {
	case All:
		return true
	case Any:
		return true
	case None:
		return true
	default:
		return false
	}
}

// Defines values for ScanStatus.
const (
	ScanStatusComplete      ScanStatus = "complete"
	ScanStatusResourceLimit ScanStatus = "resource-limit"
	ScanStatusTruncated     ScanStatus = "truncated"
)

// Valid indicates whether the value is a known member of the ScanStatus enum.
func (e ScanStatus) Valid() bool {
	switch e {
	case ScanStatusComplete:
		return true
	case ScanStatusResourceLimit:
		return true
	case ScanStatusTruncated:
		return true
	default:
		return false
	}
}

// Defines values for SortType.
const (
	Alphabetical SortType = "alphabetical"
	Position     SortType = "position"
)

// Valid indicates whether the value is a known member of the SortType enum.
func (e SortType) Valid() bool {
	switch e {
	case Alphabetical:
		return true
	case Position:
		return true
	default:
		return false
	}
}

// Defines values for SuggestionSource.
const (
	SuggestionSourceAlias  SuggestionSource = "alias"
	SuggestionSourceParse  SuggestionSource = "parse"
	SuggestionSourcePreset SuggestionSource = "preset"
)

// Valid indicates whether the value is a known member of the SuggestionSource enum.
func (e SuggestionSource) Valid() bool {
	switch e {
	case SuggestionSourceAlias:
		return true
	case SuggestionSourceParse:
		return true
	case SuggestionSourcePreset:
		return true
	default:
		return false
	}
}

// Defines values for UnrepresentableFactReason.
const (
	AmbiguousIdentity UnrepresentableFactReason = "ambiguous-identity"
	NoField           UnrepresentableFactReason = "no-field"
	NoView            UnrepresentableFactReason = "no-view"
	ReservedLiteral   UnrepresentableFactReason = "reserved-literal"
	Unframable        UnrepresentableFactReason = "unframable"
	ValueLossy        UnrepresentableFactReason = "value-lossy"
)

// Valid indicates whether the value is a known member of the UnrepresentableFactReason enum.
func (e UnrepresentableFactReason) Valid() bool {
	switch e {
	case AmbiguousIdentity:
		return true
	case NoField:
		return true
	case NoView:
		return true
	case ReservedLiteral:
		return true
	case Unframable:
		return true
	case ValueLossy:
		return true
	default:
		return false
	}
}

// Defines values for UnrepresentableOutputReason.
const (
	AmbiguousRepresentation UnrepresentableOutputReason = "ambiguous-representation"
	UnresolvedContext       UnrepresentableOutputReason = "unresolved-context"
)

// Valid indicates whether the value is a known member of the UnrepresentableOutputReason enum.
func (e UnrepresentableOutputReason) Valid() bool {
	switch e {
	case AmbiguousRepresentation:
		return true
	case UnresolvedContext:
		return true
	default:
		return false
	}
}

// Defines values for WarningKind.
const (
	WarningKindDeprecation WarningKind = "deprecation"
	WarningKindNotice      WarningKind = "notice"
	WarningKindRule        WarningKind = "rule"
)

// Valid indicates whether the value is a known member of the WarningKind enum.
func (e WarningKind) Valid() bool {
	switch e {
	case WarningKindDeprecation:
		return true
	case WarningKindNotice:
		return true
	case WarningKindRule:
		return true
	default:
		return false
	}
}

// AccessComparisonResult Machine-readable result of one access-rule comparison.
type AccessComparisonResult struct {
	// Left Typed operand captured while an access rule is evaluated.
	Left AccessOperand `json:"left"`

	// Matched Whether the comparison succeeded.
	Matched bool `json:"matched"`

	// Missing Required values missing from the evaluated operand, when applicable.
	Missing *[]interface{} `json:"missing,omitempty"`

	// Operator Policy operator applied to the operands.
	Operator string `json:"operator"`

	// Right Typed operand captured while an access rule is evaluated.
	Right AccessOperand `json:"right"`
}

// AccessDecision Overall access decision with per-rule breakdown.
//
// Returned by `AccessEngineAdapter.check()`. `granted` is `true` only when every non-skipped rule granted access. Inspect `rules` to surface per-category denial reasons in an authorization UI: each `AccessRuleResult` carries a human-readable `reason` field.
type AccessDecision struct {
	// Granted True when every rule that was not skipped granted access.
	Granted bool `json:"granted"`

	// Rules One result per rule in the policy, in policy order.
	Rules []AccessRuleResult `json:"rules"`

	// Warnings Warnings from preparing and evaluating the target, including automatic changes. Present on granted and denied decisions; empty when none apply.
	Warnings []Warning `json:"warnings"`
}

// AccessOperand Typed operand captured while an access rule is evaluated.
type AccessOperand struct {
	// Assertions Assertions addressed by an assertion operand.
	Assertions *[]AssertionReference `json:"assertions,omitempty"`

	// CategoryID Category addressed by an assertion operand.
	CategoryID *string `json:"categoryId,omitempty"`

	// Path Dot-path for subject or context operands.
	Path *string `json:"path,omitempty"`

	// Source Source of one operand in an access-policy comparison.
	Source AccessOperandSource `json:"source"`

	// Value Literal or resolved opaque value.
	Value interface{} `json:"value,omitempty"`
}

// AccessOperandSource Source of one operand in an access-policy comparison.
type AccessOperandSource string

// AccessRequirementsDescription Aggregated immutable-policy requirements for a trusted target.
type AccessRequirementsDescription struct {
	// Access Identifies an immutable artifact participating in a composed operation.
	Access ArtifactReference `json:"access"`

	// AppliedRules Qualified rules activated by the trusted target.
	AppliedRules []RuleReference `json:"appliedRules"`

	// AssertionDisplays Resolved operational display (rendered marking view) for each entry in `assertions`, index-aligned. A consumer that instantiates `requirementText` against the target assertions substitutes this rather than the raw concept id, so the requirement renders the concept's marking token instead of its internal id. `assertions` keeps the machine `conceptId` keys.
	AssertionDisplays []string `json:"assertionDisplays"`

	// Assertions Normalized effective target assertions, listed in print order: the order `marking` places each concept; a concept it leaves out appears where it would print.
	Assertions []Assertion `json:"assertions"`

	// CombinedAttributesNeeded Union of subject and context paths needed to evaluate the target.
	CombinedAttributesNeeded []string `json:"combinedAttributesNeeded"`

	// CombinedRequirementText Combined human-readable requirement expression.
	CombinedRequirementText string `json:"combinedRequirementText"`

	// Marking Normalized marking when the caller supplied or derived one.
	Marking *string `json:"marking,omitempty"`

	// Rules Per-rule immutable-policy descriptions.
	Rules []AccessRuleDescription `json:"rules"`

	// Scheme Identifies an immutable artifact participating in a composed operation.
	Scheme ArtifactReference `json:"scheme"`
}

// AccessRuleActivation Machine-readable record of why an access rule activated.
type AccessRuleActivation struct {
	// Assertions Assertions that activated the rule.
	Assertions []AssertionReference `json:"assertions"`

	// ContextPaths Context paths consulted by activation.
	ContextPaths []string `json:"contextPaths"`

	// Matched Whether the rule's activation predicate matched.
	Matched bool `json:"matched"`
}

// AccessRuleDescription Human-readable description of a single access rule.
//
// Returned by `AccessEngineAdapter.describeAllAccessRules`. Tooling (LLM agents, UI tooltips) can render `requirementText` directly or walk `attributesNeeded` to know which `userAttributes` values a caller has to provide.
type AccessRuleDescription struct {
	// Access Identifies an immutable artifact participating in a composed operation.
	Access ArtifactReference `json:"access"`

	// AttributesNeeded Subject or context paths required to evaluate the rule.
	AttributesNeeded []string `json:"attributesNeeded"`

	// DependsOn Qualified rules this rule depends on.
	DependsOn []RuleReference `json:"dependsOn"`

	// Description Optional policy-authored description.
	Description *string `json:"description,omitempty"`

	// Name Human-readable rule name.
	Name string `json:"name"`

	// Overrides Qualified rules this rule overrides.
	Overrides []RuleReference `json:"overrides"`

	// RequirementText Human-readable requirement expression.
	RequirementText string `json:"requirementText"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`

	// Scheme Identifies an immutable artifact participating in a composed operation.
	Scheme ArtifactReference `json:"scheme"`

	// WhenText Human-readable activation expression.
	WhenText *string `json:"whenText,omitempty"`
}

// AccessRuleResult Result of evaluating a single qualified access rule entry.
type AccessRuleResult struct {
	// Activation Machine-readable record of why an access rule activated.
	Activation AccessRuleActivation `json:"activation"`

	// Comparisons Machine-readable comparisons evaluated by the rule.
	Comparisons []AccessComparisonResult `json:"comparisons"`

	// Granted Whether this rule passed.
	Granted bool `json:"granted"`

	// ModernizedFrom Deprecated-to-current concept substitutions that processing applied to the data before this rule was evaluated.
	//
	// Present only when the data carried a deprecated concept. It completes the explanation of the decision: the marking the document bears, the current concept it resolved to, and the comparison this rule actually made. A consumer showing why access was granted or denied should surface it, so the verdict is never reported against a concept the reader never saw.
	ModernizedFrom *[]ConceptModernization `json:"modernizedFrom,omitempty"`

	// Name Rule name from the access policy (for display).
	Name string `json:"name"`

	// Reason Denial reason from the rule definition or generated explanation.
	Reason *string `json:"reason,omitempty"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`

	// RuleIndex Zero-based index of the rule in the access policy's evaluated rules.
	RuleIndex int `json:"ruleIndex"`

	// Skipped Whether the rule was skipped because its activation did not match.
	Skipped bool `json:"skipped"`
}

// AggregateResult Output of SchemeEngine.aggregate(): the combined marking produced from several input markings, always a whole. By design aggregate always produces the proper rollup: fixes are applied, never proposed. Every field it shares with `ValidateResult` means what it means there, `valid` included. It differs in that:
//
// - There is no `inputWasValid` (input is N markings, not a single string)
// - `proposedFixes` is always empty (aggregate applies what it can)
type AggregateResult struct {
	// AppliedFixes The changes the rollup made to reach the combined marking.
	AppliedFixes []Change `json:"appliedFixes"`

	// Assertions The assertions of the combined marking, listed in print order as a validate result's are: the order the rendered marking places each concept; a concept the rendering leaves out appears where it would print.
	Assertions []Assertion `json:"assertions"`

	// DisplayColor Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
	//
	// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
	DisplayColor *ConceptColor `json:"displayColor,omitempty"`

	// Errors The constraints the combined marking violates.
	Errors []ValidationError `json:"errors"`

	// Notices The notices the scheme attaches to the concepts in `assertions`, collected the same way `SchemeEngine.collectNotices()` collects them. Always present, empty when the aggregate carries none, so a consumer can read the field without checking that it exists.
	Notices []Notice `json:"notices"`

	// ProposedFixes Always empty: aggregate applies every fix it can.
	ProposedFixes []Fix `json:"proposedFixes"`

	// Rendered The combined marking, written in the output format.
	Rendered string `json:"rendered"`

	// Valid True when the combined marking stands once settled, as `ValidateResult.valid` defines it.
	Valid bool `json:"valid"`

	// Warnings Advisory signals about the combined marking.
	Warnings []Warning `json:"warnings"`
}

// AggregatedMarkings Aggregated guide output: one rolled-up marking per target scheme id.
type AggregatedMarkings map[string]string

// ApplyPresetRequest Context scheme apply-preset request.
type ApplyPresetRequest struct {
	// OutputFormat The format the expanded marking is rendered in. Omitted or empty, the scheme's default format. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// PresetID ID of the label preset to expand.
	PresetID string `json:"presetId"`
}

// ArtifactKind Closed artifact families participating in a composed operation.
type ArtifactKind string

// ArtifactReference Identifies an immutable artifact participating in a composed operation.
type ArtifactReference struct {
	// Digest Content digest when the resolver can supply one.
	Digest *string `json:"digest,omitempty"`

	// ID Artifact identifier.
	ID string `json:"id"`

	// Kind Closed artifact families participating in a composed operation.
	Kind ArtifactKind `json:"kind"`

	// Version Resolved immutable artifact version.
	Version string `json:"version"`
}

// Assertion A concept asserted in a category during processing, with how it was produced and where it came from.
//
// Assertions are the core data unit flowing through the engine pipeline. `confidence` records how the assertion was produced: `'certain'` when read or handed in as stated, `'inferred'` when an inference rule added it, and `'repaired'` when the engine added or rewrote it while repairing the marking. `source` says where the assertion came from.
type Assertion struct {
	// CategoryID The id of the category the concept is asserted in.
	CategoryID string `json:"categoryId"`

	// ConceptID The id of the concept.
	ConceptID string `json:"conceptId"`

	// Confidence Confidence assigned to an assertion by the processing pipeline.
	Confidence AssertionConfidence `json:"confidence"`

	// ReplacedConceptID The deprecated concept this assertion was modernized from.
	//
	// Set when processing rewrote a deprecated concept into the replacement its scheme declares. `conceptId` is the current concept every downstream stage evaluates; this field records what the data was actually marked with, so an explanation can show the substitution ("concept A was replaced by concept B; B is what your access was compared against") instead of silently presenting a concept the document never carried.
	//
	// It travels on the assertion rather than alongside it because assertion arrays are the engine's state transport: an access check made from transported assertions can still explain the substitution.
	ReplacedConceptID *string `json:"replacedConceptId,omitempty"`

	// Source Where this assertion came from. An open string: the engine labels the facts it adds, and a caller or a host may label the facts it hands in however it likes.
	//
	// Four labels mean the engine derived the fact, so it derives the fact again each time it settles a marking: `infer`, and `repair` with a `sourceRule`, a fact a scheme rule adds; `default`, a category's default concept; and `context`, a value a part took from its whole. When validate, process or explain settles a marking, a fact handed in under one of them is set aside and comes back only if the engine derives it again, so no label can make a removal silent.
	//
	// Every other label is the author's statement and is kept as given: `parse` (read from text), `user` (handed in by a caller), `aggregate` (a value a rollup states as its own, even one a contributor derived or took from its whole), every other label the engine writes on a result, and any label a caller or host chooses, such as `stated`.
	Source string `json:"source"`

	// SourceRule The id of the scheme rule that produced this assertion, when one did. Otherwise an engine-owned id beginning with `_`, set when the scheme's structure, not a rule, added the assertion. The engine writes two: `_tier_child_requires_parent` (a tier parent required by a stated child) and `_header_of_stated_member` (a header filled in for a stated member). A scheme rule id never begins with `_`.
	SourceRule *string `json:"sourceRule,omitempty"`

	// Value The value of a concept that takes one (`text_input`).
	//
	// A value handed in is read the way a value in text is: it settles in the spelling the concept's encoding stores. One the encoding cannot read is reported as the error `_invalid_assertion_value`, and a concept that takes a value, handed in with none, as `_missing_assertion_value`. See `ValidationError`.
	Value *string `json:"value,omitempty"`
}

// AssertionConfidence Confidence assigned to an assertion by the processing pipeline.
type AssertionConfidence string

// AssertionRef A reference to one assertion of the state being rendered.
type AssertionRef struct {
	// CategoryID The category it was asserted in.
	CategoryID string `json:"categoryId"`

	// ConceptID The concept asserted.
	ConceptID string `json:"conceptId"`

	// Value Its value, for a concept that carries free text.
	Value *string `json:"value,omitempty"`
}

// AssertionReference Compact reference to an assertion without engine-production metadata.
type AssertionReference struct {
	// CategoryID Category containing the asserted concept.
	CategoryID string `json:"categoryId"`

	// ConceptID Asserted concept identifier.
	ConceptID string `json:"conceptId"`

	// Value Captured value for value-bearing concepts.
	Value *string `json:"value,omitempty"`
}

// AssertionSupport De-duplicated content lineage for one compiled assertion.
type AssertionSupport struct {
	// Assertion Compact reference to an assertion without engine-production metadata.
	Assertion AssertionReference `json:"assertion"`

	// Evidence Canonically ordered evidence for the assertion.
	Evidence []ContentEvidenceSpan `json:"evidence"`

	// FactReferences Human-decision or projected facts supplying the assertion.
	FactReferences []GuideFactReference `json:"factReferences"`

	// Rules Direct and computed rules contributing to the assertion.
	Rules []RuleReference `json:"rules"`

	// SupportID Content-derived stable support identifier.
	SupportID string `json:"supportId"`
}

// AssessmentLimitation Explains why an assessment dimension or conclusion is constrained.
type AssessmentLimitation struct {
	// Code Stable machine-readable limitation code.
	Code string `json:"code"`

	// Reason Human-readable explanation safe for the selected response profile.
	Reason string `json:"reason"`

	// Scope Optional operation scope affected by the limitation.
	Scope *string `json:"scope,omitempty"`
}

// CalendarDateSpelling One spelling of a calendar date. The canonical spelling is the one a value is stored in; the others are read on input and written per view.
type CalendarDateSpelling string

// CalendarDateValueEncoding A value that is a calendar date: a real year, month and day.
type CalendarDateValueEncoding struct {
	// Accepts Spellings accepted on input.
	Accepts []CalendarDateSpelling `json:"accepts"`

	// Canonical The spelling the value is stored in.
	Canonical string `json:"canonical"`

	// Type A calendar date: a real year, month and day. Never a timestamp.
	Type string `json:"type"`

	// Views Output spelling per view name. A view not named here writes the canonical spelling.
	Views *CalendarDateViewSpellings `json:"views,omitempty"`
}

// CalendarDateViewSpellings Output spelling per view name. A view not named here writes the canonical spelling.
type CalendarDateViewSpellings map[string]CalendarDateSpelling

// CategoryEachAffixMap What is written around every single value of a category, keyed by category id.
type CategoryEachAffixMap map[string]EachAffix

// CategoryLeadingSeparatorMap What is written immediately before a category, keyed by category id.
type CategoryLeadingSeparatorMap map[string]string

// CategoryPositionOverrides Concept positions used by one category when position sorting is selected.
type CategoryPositionOverrides map[string]float32

// CategorySeparatorMap What is written between two values of one category, keyed by category id.
type CategorySeparatorMap map[string]string

// CategoryTierMap How a category's tiered concepts compose, keyed by category id.
type CategoryTierMap map[string][]TierComposition

// CategoryView Resolved view of a scheme category, suitable for display and selection UIs.
//
// Returned by `SchemeEngineAdapter.listCategories()`. Field names mirror the YAML `CategoryDoc` shape so consumers reading the source YAML and calling the API see identical names.
type CategoryView struct {
	// DefaultConcept Concept ID to assert automatically when no selection exists in this category. With `single: true`, this provides a default level. Mirrors the scheme's `default_concept` for the category.
	DefaultConcept *CategoryView_DefaultConcept `json:"defaultConcept,omitempty"`

	// HierarchyRoot When true, the category grades its concepts by severity `rank`, and a coverage comparison counts a whole-marking concept here as covering a part's concept of equal or lower rank. It does not change rendering.
	HierarchyRoot *bool `json:"hierarchyRoot,omitempty"`

	// ID Unique category identifier.
	ID string `json:"id"`

	// Label Resolved human-readable label (always a plain string).
	Label string `json:"label"`

	// Position Declared rendering position (lower = earlier). Not resolved from array order.
	Position *float32 `json:"position,omitempty"`

	// PositionOverrides Concept positions used by one category when position sorting is selected.
	PositionOverrides *CategoryPositionOverrides `json:"positionOverrides,omitempty"`

	// Required Whether at least one concept is required for a valid marking.
	Required *bool `json:"required,omitempty"`

	// Rollup Canonical rollup operator. `pipeline` requires a matching entry in the scheme's `rollup_pipelines`. Most merge operators are set-algebraic over each input's assertions; rank-based operators (max/min) require categories whose concepts declare `rank`.
	Rollup Rollup `json:"rollup"`

	// Scopes Scope IDs this category is visible in.
	Scopes *[]string `json:"scopes,omitempty"`

	// Single Whether only one concept can be selected in this category.
	Single *bool `json:"single,omitempty"`

	// Sort Sort order for concepts within this category.
	Sort *CategoryView_Sort `json:"sort,omitempty"`
}

// CategoryViewDefaultConcept0 defines model for CategoryView.DefaultConcept.0.
type CategoryViewDefaultConcept0 = string

// CategoryView_DefaultConcept Concept ID to assert automatically when no selection exists in this category. With `single: true`, this provides a default level. Mirrors the scheme's `default_concept` for the category.
type CategoryView_DefaultConcept struct {
	union json.RawMessage
}

// CategoryViewSort1 defines model for CategoryView.Sort.1.
type CategoryViewSort1 = []SortType

// CategoryView_Sort Sort order for concepts within this category.
type CategoryView_Sort struct {
	union json.RawMessage
}

// CategoryViewMap The concept view a category is written in, keyed by category id.
type CategoryViewMap map[string]string

// Change A change the engine actually made while processing. Lives in `appliedFixes`; a declined addition is not a change.
type Change struct {
	// Action Discrete mutation represented by an engine `Change`.
	Action FixAction `json:"action"`

	// Add Concepts the change added. Set for `add` and `replace`.
	Add *[]FixConceptReference `json:"add,omitempty"`

	// Message What the change did, in words.
	Message string `json:"message"`

	// Remove Concepts the change removed. Set for `remove` and `replace`.
	Remove *[]FixConceptReference `json:"remove,omitempty"`

	// RuleID Rule that made this change: a scheme rule's id, or an id the engine keeps for a change no scheme rule names.
	//
	// `_recognition_role` is one of them: the text was read as a view of one role, for a request whose context names the other. See `ValidateResult.inputWasValid`, and `ValidateResult.appliedFixes` for how it is shown.
	RuleID string `json:"ruleId"`

	// TriggerConceptIds Concept IDs whose presence triggered the rule that made this change.
	TriggerConceptIds *[]string `json:"triggerConceptIds,omitempty"`
}

// CharacterWindow Proximity measured in raw characters.
type CharacterWindow struct {
	// Chars The largest distance allowed between two spans, counted in characters (UTF-16 code units) between their nearest edges. Spans that touch or overlap are 0 apart.
	Chars int `json:"chars"`
}

// CheckCoverageResult Result of a coverage check. The engine is given a list of child markings and a single parent marking, aggregates the children in the scheme's format, then compares the parent with the children category by category.
//
// A parent can fail to cover its children in two ways, and each has its own field. `missing` is what the children need that the parent does not carry. `excess` is what the parent carries that the children do not account for. A caller tells "add this" from "remove this" by which field a concept is in, not by reading error text.
//
// How a category is judged comes from how the scheme rolls it up. A category that a rollup pipeline names is judged as a `pipeline` category, whatever rollup it declares. Every other category is judged by the rollup it declares.
//
// In a `union` category the parent must carry every value the children roll up to, and anything it lacks is missing. A value no child carries is allowed on the parent and is not excess.
//
// In an `intersection` category the parent may carry only values every child allows, and anything else it carries is excess. A shorter list, or none, is allowed and is not missing.
//
// In a `max` or `min` category the children roll up to one value, ordered by rank. The parent must sit at or above that value for `max`, and at or below it for `min`. A parent anywhere else is reported once, under `missing`, and never under `excess`. The value reported is the children's own value.
//
// In a `pipeline` category the parent covers when rolling it up with the children leaves it unchanged. This relies on the pipeline leaving a parent rolled up on its own unchanged, which the scheme author has to make true.
//
// The findings in `pipeline` categories are worked out for all the categories one pipeline names together. First they are narrowed to what the children roll up to on their own: a value the rollup of parent and children adds is missing only when the children's own rollup carries it, and a parent value that rollup drops is excess only when the children's own rollup does not carry it. The narrowed findings are then tested by acting on every finding of the report: the excess is taken out of the parent and the missing is put in, replacing what a single-value category held. They are reported only when that gives a parent that is valid under the scheme's rules and covers the children in every category, judged against the children's rollup. Otherwise, and whenever narrowing would leave nothing, the findings are not narrowed: every value the rollup of parent and children adds is missing, and every parent value it drops is excess, including one the children carry. Pipelines are decided one at a time, in the order the scheme lists them. Each is judged with the findings already decided for the pipelines before it and the full findings for those after it.
//
// In a `first` category the children roll up to one value and the rollup has no restrictive direction, so the parent must state that value. A parent that states a different one is reported once, under `missing`, as the value it should state, and never under `excess`. In such a category that holds several values, values the parent states beyond the picked one are not reported.
//
// In a category where the scheme has an `expand_group` rule, a parent that carries every member a group declares in that category states the group, and the group is not missing from it. This does not apply to a category whose rollup has no restrictive direction (`first`), where the parent must state what the children roll up to as it is.
//
// In a `union` or `pipeline` category the scheme marks as a hierarchy root, a parent value ranked at or above a needed value also covers it.
//
// With no non-empty children, a valid parent is covered and both lists are empty.
//
// `rendered` is the aggregate of the children rendered in the requested format; the UI can display it directly as the computed rollup without a separate aggregate() round-trip.
type CheckCoverageResult struct {
	// Covered Whether the parent covers the children. True when `missing` and `excess` are both empty and no rollup was refused. False otherwise; `missing`, `excess` and `refusals` together say why.
	Covered bool `json:"covered"`

	// Excess Concepts the parent carries that the children do not account for, judged per category as described on this type. In a `pipeline` category whose findings are not narrowed it means every parent value the rollup of parent and children drops, which can include values the children carry. Everywhere else it means values the children do not allow.
	//
	// Concepts appear as the author wrote them. When a written value brought others with it, such as a group that stands for its members, only the written value is listed, not the members. In a category where nothing in excess was written by the author, every excess value is listed. `excess` is empty only when the parent carries nothing the children do not account for.
	Excess []Assertion `json:"excess"`

	// ExcessLabels Display labels for the concepts in `excess`, for showing to a person. A label appears once even when several concepts display the same way, so the list can be shorter than `excess` and is not aligned with it by position. Where the operation takes an output format, labels are spelled in it. Empty when `excess` is empty.
	ExcessLabels []string `json:"excessLabels"`

	// Missing Concepts the children need that the parent does not carry, judged per category as described on this type. In a `pipeline` category whose findings are not narrowed, this can include a value the pipeline itself adds when it rolls the parent up with the children. Empty when the parent lacks nothing.
	Missing []Assertion `json:"missing"`

	// MissingLabels Display labels for the concepts in `missing`, for showing to a person. A label appears once even when several concepts display the same way, so the list can be shorter than `missing` and is not aligned with it by position. Where the operation takes an output format, labels are spelled in it. Empty when `missing` is empty.
	MissingLabels []string `json:"missingLabels"`

	// Refusals What the scheme said when it refused a rollup this check computed, in the scheme's own words: the children rolled up together, or the parent rolled up with them.
	//
	// Markings that are each valid can still combine into a state the scheme does not allow, and that is a defect in the document rather than in the request: it is the answer a coverage check exists to give, so it is reported here instead of raised. Absent when no rollup was refused. Missing and excess concepts are never reported here; they have their own fields.
	//
	// Each entry names the constraint and carries the message the rule wrote, which is where the contributing markings are named.
	Refusals *[]ValidationError `json:"refusals,omitempty"`

	// Rendered The children's aggregate, written in the requested format, so a UI can show the computed rollup without a separate aggregate call. Empty when there are no child markings.
	Rendered string `json:"rendered"`
}

// CompositeEncodingField One component of a resolved `composite` encoding.
type CompositeEncodingField struct {
	// Concepts `type: concept`: the concept ids this component may name.
	Concepts *[]string `json:"concepts,omitempty"`

	// Name Component name, used in diagnostics.
	Name string `json:"name"`

	// Prefix Written once before the component.
	Prefix *string `json:"prefix,omitempty"`

	// Required Whether the component must be present. Default true.
	Required *bool `json:"required,omitempty"`

	// Suffix Written once after the component.
	Suffix *string `json:"suffix,omitempty"`

	// Type What one component of a `composite` encoding stands for: free text, or a reference to known concepts.
	Type CompositeEncodingFieldType `json:"type"`

	// View `type: concept`: the view the component is written in.
	View *string `json:"view,omitempty"`
}

// CompositeEncodingFieldType What one component of a `composite` encoding stands for: free text, or a reference to known concepts.
type CompositeEncodingFieldType string

// CompositeValueEncoding An ordered set of components stored as one normalized string.
type CompositeValueEncoding struct {
	// Fields The components, in the order they are written.
	Fields []CompositeEncodingField `json:"fields"`

	// Separator Joins the components.
	Separator *string `json:"separator,omitempty"`

	// Type An ordered set of components stored as one normalized string.
	Type string `json:"type"`
}

// ConceptColor Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
//
// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
type ConceptColor struct {
	// Background Background color as a CSS hex string.
	Background string `json:"background"`

	// Text Text/foreground color as a CSS hex string.
	Text string `json:"text"`
}

// ConceptDescription Pre-formatted tooltip content for a concept: name, definition, guidance, and applicable rules.
type ConceptDescription struct {
	// CategoryID The category the concept belongs to.
	CategoryID string `json:"categoryId"`

	// ConceptID The id of the concept described.
	ConceptID string `json:"conceptId"`

	// Definition What the concept means, as the scheme defines it.
	Definition *string `json:"definition,omitempty"`

	// Guidance How to apply the concept, as the scheme advises.
	Guidance *string `json:"guidance,omitempty"`

	// Lines Pre-formatted lines for tooltip display (header, description, guidance, rules).
	Lines []string `json:"lines"`

	// Name The concept's display name.
	Name string `json:"name"`

	// Rules The scheme's rules that involve the concept, each as a sentence.
	Rules *[]string `json:"rules,omitempty"`
}

// ConceptModernization One deprecated-to-current concept substitution that processing applied before an access rule was evaluated.
//
// The id fields are machine keys; the `*Display` companions carry the rendered marking view, following the same split as `missing` / `missingDisplay`.
type ConceptModernization struct {
	// ConceptDisplay Resolved operational display for `conceptId`.
	ConceptDisplay string `json:"conceptDisplay"`

	// ConceptID Current concept the access rule was evaluated against.
	ConceptID string `json:"conceptId"`

	// ReplacedConceptDisplay Resolved operational display for `replacedConceptId`.
	ReplacedConceptDisplay string `json:"replacedConceptDisplay"`

	// ReplacedConceptID Deprecated concept the data was actually marked with.
	ReplacedConceptID string `json:"replacedConceptId"`
}

// ConceptNotice Regulatory or handling notice attached to a concept: the text the marking requires, what kind of notice it is, whether showing it is mandatory, and an optional condition that gates it.
//
// Where the text appears and how it is styled are the host's to decide. A scheme states what has to be said, not where the document says it.
type ConceptNotice struct {
	// Condition Only show when this concept is also present
	Condition *string `json:"condition,omitempty"`

	// Required Whether this notice MUST be displayed (regulatory requirement)
	Required bool `json:"required"`

	// Text The notice text the marking requires.
	Text string `json:"text"`

	// Type What kind of notice a concept carries.
	Type ConceptNoticeType `json:"type"`

	// Views The value a notice is registered under, keyed by view name, e.g. `{ ism: 'RD' }`.
	//
	// Structurally the same map as `ViewLabels` and deliberately not the same type. A view label is how a view spells something for a reader; this is the token a view's own vocabulary files the notice under, and one is never substitutable for the other.
	Views *NoticeViewValues `json:"views,omitempty"`
}

// ConceptNoticeType What kind of notice a concept carries.
type ConceptNoticeType string

// ConceptRef A reference to a concept in a dependent scheme. Resolved to a full Assertion during guide evaluation.
type ConceptRef struct {
	// Category Category ID: required when the concept belongs to multiple categories.
	Category *string `json:"category,omitempty"`

	// ID Concept ID: must exist in one of the guide's target schemes.
	ID string `json:"id"`

	// Value Value for text-input concepts (e.g., country codes, project names).
	Value *string `json:"value,omitempty"`
}

// ConceptView Resolved view of a concept, suitable for display and adapter consumers.
//
// Returned by `SchemeEngineAdapter.getConceptsInCategory()`. Field names mirror the YAML `ConceptDoc` shape. `position` and `rank` are concrete numbers (already resolved from relative specs by the loader).
//
// `rulesReferencing` is populated by `buildSchemeView` (Handbook) and omitted by `getConceptsInCategory`; treat it as absent unless explicitly set.
type ConceptView struct {
	// Aliases Alternative strings that parse to this concept.
	Aliases []string `json:"aliases"`

	// Categories Category IDs this concept belongs to.
	Categories []string `json:"categories"`

	// Citation Where this concept came from: the provenance text the scheme records against it.
	//
	// Inert. Nothing in the engine reads it and no decision turns on it; it is what a reference surface shows when a reader asks why a concept is spelled the way it is. Published here because this projection is the only one a remote client can reach, and a scheme that makes the field mandatory had nowhere to put it.
	Citation *string `json:"citation,omitempty"`

	// Color Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
	//
	// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
	Color *ConceptColor `json:"color,omitempty"`

	// Definition Definition text from the concept's documentation block.
	Definition *string `json:"definition,omitempty"`

	// Deprecated Whether this concept is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`

	// DeprecatedMessage Human-readable deprecation message.
	DeprecatedMessage *string `json:"deprecatedMessage,omitempty"`

	// DeprecatedReplacement Concept ID that should be used instead of this deprecated concept.
	DeprecatedReplacement *string `json:"deprecatedReplacement,omitempty"`

	// Guidance Usage guidance from the concept's documentation block.
	Guidance *string `json:"guidance,omitempty"`

	// ID Unique concept identifier.
	ID string `json:"id"`

	// Members For group concepts: the member concept IDs.
	Members *[]string `json:"members,omitempty"`

	// Name Human-readable display name.
	Name string `json:"name"`

	// Notices Regulatory or policy notices shown with this concept.
	Notices *[]ConceptNotice `json:"notices,omitempty"`

	// Parents Broader concept IDs forming the hierarchy chain.
	Parents []string `json:"parents"`

	// Position Resolved rendering position within a category (lower = earlier).
	Position float32 `json:"position"`

	// Rank Resolved severity rank for max/min rollup aggregation (higher = more severe).
	Rank *float32 `json:"rank,omitempty"`

	// RulesReferencing Human-readable summary of rules that reference this concept. Set by `buildSchemeView`; absent from `getConceptsInCategory`.
	RulesReferencing *[]string `json:"rulesReferencing,omitempty"`

	// Scopes Scope IDs this concept is visible in.
	Scopes *[]string `json:"scopes,omitempty"`

	// TextInput Text-input configuration for free-text concepts.
	TextInput *struct {
		// Encoding How a free-text value is spelled and stored: the resolved codec.
		//
		// A codec, not a validator. Every encoding states both directions, so a value this contract accepts is a value it can write back.
		Encoding *ValueEncoding `json:"encoding,omitempty"`

		// Label Label shown with the input.
		Label *string `json:"label,omitempty"`

		// Placeholder Hint text shown in the empty input.
		Placeholder *string `json:"placeholder,omitempty"`

		// Required Whether a value must be provided. Defaults to true.
		Required *bool `json:"required,omitempty"`
	} `json:"textInput,omitempty"`

	// Tier Tier label: marks this concept as a partial marking.
	Tier *string `json:"tier,omitempty"`

	// Views Display representations keyed by view name; includes at least `'default'`.
	Views map[string]string `json:"views"`
}

// ConditionalDefaultConcept A category default evaluated before any category defaults are added.
type ConditionalDefaultConcept struct {
	// Concept Concept ID to fill when the category is empty.
	Concept string `json:"concept"`

	// When Conditions that trigger a rule. Supports shorthand forms:
	//
	// - `string`: concept ID must be present (any category)
	// - `string[]`: any of these concept IDs must be present
	// - `object`: structured `{ any | all | none }` block, possibly nested
	When RuleWhen `json:"when"`
}

// ContentEvidenceSpan Evidence range in the exact accepted UTF-8 document bytes.
type ContentEvidenceSpan struct {
	// End Zero-based exclusive byte offset.
	End int `json:"end"`

	// Matched Matched source text, present only in the full response profile.
	Matched *string `json:"matched,omitempty"`

	// OffsetEncoding Frozen offset encoding for document evidence.
	OffsetEncoding string `json:"offsetEncoding"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`

	// Start Zero-based inclusive byte offset.
	Start int `json:"start"`
}

// ContextAccessDecisionResult Context envelope returned by trusted-target access checks.
type ContextAccessDecisionResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Overall access decision with per-rule breakdown.
	//
	// Returned by `AccessEngineAdapter.check()`. `granted` is `true` only when every non-skipped rule granted access. Inspect `rules` to surface per-category denial reasons in an authorization UI: each `AccessRuleResult` carries a human-readable `reason` field.
	Result AccessDecision `json:"result"`
}

// ContextAccessRequirementsResult Context envelope returned by immutable-policy requirement description.
type ContextAccessRequirementsResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Aggregated immutable-policy requirements for a trusted target.
	Result AccessRequirementsDescription `json:"result"`
}

// ContextCrosswalkCandidate One candidate of a cross-field decision a Context's crosswalk applies.
//
// A candidate document may leave its disposition out to mean `carry`. The description always states it, so a reader never applies that default itself.
type ContextCrosswalkCandidate struct {
	// Authority Authority behind the candidate decision, when one exists.
	Authority *string `json:"authority,omitempty"`

	// Citation Citation supporting the candidate decision.
	Citation *string `json:"citation,omitempty"`

	// Disposition Decision a reviewed crosswalk makes for one source assertion.
	Disposition CrosswalkDisposition `json:"disposition"`

	// Guidance Operator guidance for the candidate decision.
	Guidance *string `json:"guidance,omitempty"`

	// ID Stable candidate identifier within its selection.
	ID string `json:"id"`

	// Mappings Per-assertion mappings activated when this candidate is selected.
	Mappings *[]CrosswalkMapping `json:"mappings,omitempty"`

	// Reason Human-readable basis for selecting and applying this candidate.
	Reason string `json:"reason"`

	// Score Assertions counted by a `most-matches` selection.
	Score *[]AssertionReference `json:"score,omitempty"`

	// Sources Source assertions covered when this candidate is selected.
	Sources *[]AssertionReference `json:"sources,omitempty"`

	// Targets Targets emitted once when this candidate is selected.
	Targets *[]CrosswalkTarget `json:"targets,omitempty"`

	// When Boolean test over source assertions used to choose a crosswalk candidate.
	When *CrosswalkSelectionCondition `json:"when,omitempty"`
}

// ContextCrosswalkDescriptor One crosswalk a Context holds, with the reviewed decisions it applies.
//
// Read it as a sentence: this Context translates markings from the scheme `source` selects to the scheme `target` selects with crosswalk `id` at `version`, and these are its mappings and selections.
//
// `source` and `target` are scheme selectors, the `selector` values `describeSchemes` reports, not scheme ids. A view bound to one of the Context's schemes finds the crosswalk that reads from it by comparing selectors. The mappings and selections are the effective set after any crosswalk layer the Context stacks, which is the set a translation applies.
type ContextCrosswalkDescriptor struct {
	// Citation Primary citation for the crosswalk, when it declares one.
	Citation *string `json:"citation,omitempty"`

	// Guidance General usage and review guidance, when the crosswalk declares it.
	Guidance *string `json:"guidance,omitempty"`

	// ID Crosswalk id, from the crosswalk document; what `translate` takes.
	ID string `json:"id"`

	// Mappings How each source assertion translates, in declared order.
	Mappings []CrosswalkMapping `json:"mappings"`

	// Marking The crosswalk document's own overall marking, exactly as its `crosswalk.marking` writes it. Display only: never interpreted. Absent when the document declares none.
	Marking *string `json:"marking,omitempty"`

	// Name Human-readable crosswalk name.
	Name string `json:"name"`

	// Selections Cross-field decisions, in declared order. Empty when there are none.
	Selections []ContextCrosswalkSelection `json:"selections"`

	// Source Selector of the scheme this crosswalk translates from.
	Source string `json:"source"`

	// Target Selector of the scheme this crosswalk translates to.
	Target string `json:"target"`

	// Version Concrete crosswalk version that answers in this Context.
	Version string `json:"version"`
}

// ContextCrosswalkSelection One cross-field decision a Context's crosswalk applies.
//
// The document form of a selection, with every candidate's treatment written out.
type ContextCrosswalkSelection struct {
	// Candidates Candidates in declared order; order breaks ties, a fallback is last.
	Candidates []ContextCrosswalkCandidate `json:"candidates"`

	// ID Selection id, as reviewed.
	ID string `json:"id"`

	// Strategy Strategy for choosing one candidate from a reviewed cross-field decision.
	Strategy CrosswalkSelectionStrategy `json:"strategy"`
}

// ContextCrosswalkTranslationResult Context envelope returned by a directional crosswalk translation.
type ContextCrosswalkTranslationResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Result of applying a reviewed directional crosswalk.
	Result CrosswalkTranslationResult `json:"result"`
}

// ContextDocumentAuditResult Resolved Context envelope returned by document audit operations.
type ContextDocumentAuditResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Authored structural and analyst-determined guide audit of one document.
	Result DocumentAuditReport `json:"result"`
}

// ContextGuideEvaluationResult Context envelope returned by guide evaluation.
type ContextGuideEvaluationResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Output of GuideEngine.evaluate(): the aggregate marking produced by evaluating a set of selected guide rules. Parallels SchemeEngine.process() on the scheme side.
	Result EvaluationResult `json:"result"`
}

// ContextGuideScanResult Context envelope returned by guide text scanning.
type ContextGuideScanResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Typed operation result.
	Result []GuideScanResult `json:"result"`
}

// ContextGuideViewResult Context envelope returned by guide structural view.
type ContextGuideViewResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Render-ready Guidebook tree for a guide: its rules grouped under the guide's categories, with every compound rule in one `_compound` category at the end.
	Result GuideViewResponse `json:"result"`
}

// ContextResolution Resolution identity shared by every branch of one Context operation.
type ContextResolution struct {
	// Artifacts Ordered artifacts participating in the composition.
	Artifacts []ArtifactReference `json:"artifacts"`

	// ContextID Resolved Context identifier.
	ContextID string `json:"contextId"`

	// RequestedVersion The version the request named, an alias such as `latest` included, and `latest` when the request named none. `resolvedVersion` is the concrete version that answered. Absent only from a resolution produced in process without a request.
	RequestedVersion *string `json:"requestedVersion,omitempty"`

	// ResolvedVersion Concrete Context version used by the operation.
	ResolvedVersion string `json:"resolvedVersion"`
}

// ContextSchemeAggregateResult Context envelope returned by scheme aggregation.
type ContextSchemeAggregateResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Output of SchemeEngine.aggregate(): the combined marking produced from several input markings, always a whole. By design aggregate always produces the proper rollup: fixes are applied, never proposed. Every field it shares with `ValidateResult` means what it means there, `valid` included. It differs in that:
	//
	// - There is no `inputWasValid` (input is N markings, not a single string)
	// - `proposedFixes` is always empty (aggregate applies what it can)
	Result AggregateResult `json:"result"`
}

// ContextSchemeConceptDescriptionResult Context envelope returned by concept description.
type ContextSchemeConceptDescriptionResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Pre-formatted tooltip content for a concept: name, definition, guidance, and applicable rules.
	Result ConceptDescription `json:"result"`
}

// ContextSchemeCoverageResult Context envelope returned by scheme coverage checks.
type ContextSchemeCoverageResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Result of a coverage check. The engine is given a list of child markings and a single parent marking, aggregates the children in the scheme's format, then compares the parent with the children category by category.
	//
	// A parent can fail to cover its children in two ways, and each has its own field. `missing` is what the children need that the parent does not carry. `excess` is what the parent carries that the children do not account for. A caller tells "add this" from "remove this" by which field a concept is in, not by reading error text.
	//
	// How a category is judged comes from how the scheme rolls it up. A category that a rollup pipeline names is judged as a `pipeline` category, whatever rollup it declares. Every other category is judged by the rollup it declares.
	//
	// In a `union` category the parent must carry every value the children roll up to, and anything it lacks is missing. A value no child carries is allowed on the parent and is not excess.
	//
	// In an `intersection` category the parent may carry only values every child allows, and anything else it carries is excess. A shorter list, or none, is allowed and is not missing.
	//
	// In a `max` or `min` category the children roll up to one value, ordered by rank. The parent must sit at or above that value for `max`, and at or below it for `min`. A parent anywhere else is reported once, under `missing`, and never under `excess`. The value reported is the children's own value.
	//
	// In a `pipeline` category the parent covers when rolling it up with the children leaves it unchanged. This relies on the pipeline leaving a parent rolled up on its own unchanged, which the scheme author has to make true.
	//
	// The findings in `pipeline` categories are worked out for all the categories one pipeline names together. First they are narrowed to what the children roll up to on their own: a value the rollup of parent and children adds is missing only when the children's own rollup carries it, and a parent value that rollup drops is excess only when the children's own rollup does not carry it. The narrowed findings are then tested by acting on every finding of the report: the excess is taken out of the parent and the missing is put in, replacing what a single-value category held. They are reported only when that gives a parent that is valid under the scheme's rules and covers the children in every category, judged against the children's rollup. Otherwise, and whenever narrowing would leave nothing, the findings are not narrowed: every value the rollup of parent and children adds is missing, and every parent value it drops is excess, including one the children carry. Pipelines are decided one at a time, in the order the scheme lists them. Each is judged with the findings already decided for the pipelines before it and the full findings for those after it.
	//
	// In a `first` category the children roll up to one value and the rollup has no restrictive direction, so the parent must state that value. A parent that states a different one is reported once, under `missing`, as the value it should state, and never under `excess`. In such a category that holds several values, values the parent states beyond the picked one are not reported.
	//
	// In a category where the scheme has an `expand_group` rule, a parent that carries every member a group declares in that category states the group, and the group is not missing from it. This does not apply to a category whose rollup has no restrictive direction (`first`), where the parent must state what the children roll up to as it is.
	//
	// In a `union` or `pipeline` category the scheme marks as a hierarchy root, a parent value ranked at or above a needed value also covers it.
	//
	// With no non-empty children, a valid parent is covered and both lists are empty.
	//
	// `rendered` is the aggregate of the children rendered in the requested format; the UI can display it directly as the computed rollup without a separate aggregate() round-trip.
	Result CheckCoverageResult `json:"result"`
}

// ContextSchemeDescriptor One scheme a Context holds, and the selector that addresses it.
//
// Read it as a sentence: this Context holds a scheme you select with `selector`, it is `id` at `version`, and it plays these `roles`.
//
// `selector` is the whole point. A Context keys its schemes locally, and a caller reaches a scheme by its key: an adapter's `schemes.get` takes it, and so does `?scheme=` on the REST surface. A caller that cannot enumerate the keys cannot reach anything but the primary.
type ContextSchemeDescriptor struct {
	// ID Scheme id, from the scheme document itself.
	ID string `json:"id"`

	// Roles What this scheme is to the Context. Never empty.
	Roles []ContextSchemeRole `json:"roles"`

	// Selector Context-local key that selects this scheme.
	Selector string `json:"selector"`

	// Version Scheme version, from the scheme document itself.
	Version string `json:"version"`
}

// ContextSchemeDiffResult Context envelope returned by scheme diff.
type ContextSchemeDiffResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Comparison of two assertion sets: which assertions were added, removed, or shared.
	Result DiffResult `json:"result"`
}

// ContextSchemeExplainResult Context envelope returned by a scheme explanation.
type ContextSchemeExplainResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Full pipeline trace for one marking: parse input, repairs applied, inferences drawn, validation errors, and the final output.
	//
	// What `explain` returns, and nothing else. It takes what `process` takes and answers why the engine settled the marking the way it did, step by step. No other result carries it, because their callers want the answer rather than the working. Debug tools use this to explain why a marking was accepted or rejected. `repairs` lists each repair rule that fired and the before/after strings. `inferences` lists each inference rule that added a concept and the reason text.
	Result ExplanationTrace `json:"result"`
}

// ContextSchemeNoticesResult Context envelope returned by scheme notice collection.
type ContextSchemeNoticesResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Typed operation result.
	Result []Notice `json:"result"`
}

// ContextSchemeProcessResult Context envelope returned by scheme processing.
type ContextSchemeProcessResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result What process answers: everything validate answers, plus what an author needs to correct and finish the marking.
	//
	// Process is the author's call. A composer, a picker or an editor that helps someone write a marking calls it, often on every keystroke, and can send what was typed together with the facts already chosen. Beside the validate answer it carries:
	//
	// - the corrections that turn the text into canonical text, inside `recognition`;
	// - what was read before repair, when repair changed it;
	// - the fix paths, each a whole valid marking the engine has proven;
	// - the outputs the request named in `outputs`, and only those;
	// - the options for the categories the request named in `suggestFor`, only when it named some.
	//
	// The step-by-step trace of how the engine got here is not part of it. That is what explain answers, from the same request.
	Result ProcessResult `json:"result"`
}

// ContextSchemeRenderResult Context envelope returned by scheme rendering.
type ContextSchemeRenderResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Response body for Context rendering.
	Result RenderResponse `json:"result"`
}

// ContextSchemeRole What one scheme is to the Context that holds it.
//
// A scheme can be more than one of these at once. The common crosswalk carries the Context's own marking scheme as its source, so that scheme is `primary` and `crosswalk-source` together. `additional` is for a scheme a composition loaded without naming it as either, and it is the only entry that says nothing beyond "this Context holds it".
type ContextSchemeRole string

// ContextSchemeScanResult Context envelope returned by scheme text scanning.
type ContextSchemeScanResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result What a scan answers on the wire.
	//
	// An object rather than a bare array, because an array alone cannot tell a document that held no marking from a scan that stopped before the end of it. Both come back as zero entries, and the two mean opposite things to whoever decides what happens to the document next.
	Result ScanResponseBody `json:"result"`
}

// ContextSchemeSuggestForCategoriesResult Context envelope returned by per-category scheme suggestions: the options keyed by category id, the same map a process call returns in `suggestionsFor`.
type ContextSchemeSuggestForCategoriesResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Suggestions for several categories, keyed by category id.
	//
	// Each entry holds the options for its own category and no other, answered as if that category's own assertions were not yet chosen, so a picker can offer the alternatives to its current selection, not only what can still be added beside it. A label preset belongs to no category, so no entry offers one. `suggestForCategories` returns this, and so does a process call that names categories in `suggestFor`.
	Result SuggestResultsByCategory `json:"result"`
}

// ContextSchemeSuggestResult Context envelope returned by scheme suggestions.
type ContextSchemeSuggestResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result What suggest answers: the concepts a picker may offer, and nothing about the marking itself.
	//
	// How the current marking reads and whether it stands are what validate and process answer, and suggest does not repeat them. A caller that wants both in one round trip asks process for them with `suggestFor`.
	Result SuggestResult `json:"result"`
}

// ContextSchemeValidateResult Context envelope returned by scheme validation.
type ContextSchemeValidateResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Result ValidateResult `json:"result"`
}

// ContextSchemeViewResult Context envelope returned by scheme structural view.
type ContextSchemeViewResult struct {
	// Resolution Resolution identity shared by every branch of one Context operation.
	Resolution ContextResolution `json:"resolution"`

	// ResponseProjection Describes the explicit projection applied after full evaluation.
	ResponseProjection *DocumentResponseProjection `json:"responseProjection,omitempty"`

	// Result Render-ready Handbook tree for a scheme: its concepts grouped under the scheme's categories.
	Result SchemeViewResponse `json:"result"`
}

// CrosswalkCompletedTranslationResult Translation that reached target-scheme processing.
type CrosswalkCompletedTranslationResult struct {
	// Crosswalk Identifies an immutable artifact participating in a composed operation.
	Crosswalk ArtifactReference `json:"crosswalk"`

	// Outcomes Ordered treatment of every source assertion.
	Outcomes []CrosswalkOutcome `json:"outcomes"`

	// Selections Cross-field candidate chosen for every declared selection.
	Selections []CrosswalkSelectionOutcome `json:"selections"`

	// Source What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Source ValidateResult `json:"source"`

	// SourceScheme Identifies an immutable artifact participating in a composed operation.
	SourceScheme ArtifactReference `json:"sourceScheme"`

	// Status Completion state of a crosswalk that reached target-scheme processing.
	Status CrosswalkCompletedTranslationStatus `json:"status"`

	// Target What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Target ValidateResult `json:"target"`

	// TargetScheme Identifies an immutable artifact participating in a composed operation.
	TargetScheme ArtifactReference `json:"targetScheme"`
}

// CrosswalkCompletedTranslationStatus Completion state of a crosswalk that reached target-scheme processing.
type CrosswalkCompletedTranslationStatus string

// CrosswalkDisposition Decision a reviewed crosswalk makes for one source assertion.
type CrosswalkDisposition string

// CrosswalkForbidMatch A matched prohibition that aborted a crosswalk translation.
//
// Both branches are named types rather than anonymous object literals because the wire mirrors this union, and a generated SDK needs a name for each side.
type CrosswalkForbidMatch struct {
	union json.RawMessage
}

// CrosswalkForbiddenTranslationResult Translation aborted before target processing by one or more prohibitions.
type CrosswalkForbiddenTranslationResult struct {
	// Crosswalk Identifies an immutable artifact participating in a composed operation.
	Crosswalk ArtifactReference `json:"crosswalk"`

	// ForbiddenBy Every matched prohibition, in deterministic evaluation order.
	ForbiddenBy []CrosswalkForbidMatch `json:"forbiddenBy"`

	// Outcomes Ordered treatment of every source assertion.
	Outcomes []CrosswalkOutcome `json:"outcomes"`

	// Selections Cross-field candidate chosen for every declared selection.
	Selections []CrosswalkSelectionOutcome `json:"selections"`

	// Source What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Source ValidateResult `json:"source"`

	// SourceScheme Identifies an immutable artifact participating in a composed operation.
	SourceScheme ArtifactReference `json:"sourceScheme"`

	// Status Terminal hard-stop state.
	Status string `json:"status"`

	// TargetScheme Identifies an immutable artifact participating in a composed operation.
	TargetScheme ArtifactReference `json:"targetScheme"`
}

// CrosswalkMapping Reviewed treatment of one source assertion in a directional crosswalk.
type CrosswalkMapping struct {
	// Authority Authority or reviewed source behind the mapping, when one exists.
	Authority *string `json:"authority,omitempty"`

	// Citation Citation supporting this mapping decision.
	Citation *string `json:"citation,omitempty"`

	// Disposition Decision a reviewed crosswalk makes for one source assertion.
	Disposition CrosswalkDisposition `json:"disposition"`

	// Guidance Operator guidance for applying or reviewing this mapping.
	Guidance *string `json:"guidance,omitempty"`

	// ID Stable identifier for audit and review.
	ID string `json:"id"`

	// Reason Human-readable basis for the decision.
	Reason string `json:"reason"`

	// Source Compact reference to an assertion without engine-production metadata.
	Source AssertionReference `json:"source"`

	// Targets Target assertions. `carry` mappings emit these automatically; `review` mappings expose them only as proposals. `omit` and `forbid` mappings have no targets.
	Targets *[]CrosswalkTarget `json:"targets,omitempty"`
}

// CrosswalkMappingForbidMatch A per-assertion mapping that prohibited a crosswalk translation.
type CrosswalkMappingForbidMatch struct {
	// Authority Authority copied from the mapping, when declared.
	Authority *string `json:"authority,omitempty"`

	// Citation Supporting citation copied from the mapping, when declared.
	Citation *string `json:"citation,omitempty"`

	// Guidance Operator guidance copied from the mapping, when declared.
	Guidance *string `json:"guidance,omitempty"`

	// Kind A per-assertion mapping prohibited the translation.
	Kind string `json:"kind"`

	// MappingID Stable mapping id.
	MappingID string `json:"mappingId"`

	// Reason Human-readable basis for the prohibition.
	Reason string `json:"reason"`

	// Source Compact reference to an assertion without engine-production metadata.
	Source AssertionReference `json:"source"`
}

// CrosswalkOutcome Per-assertion decision produced by a crosswalk translation.
type CrosswalkOutcome struct {
	// Authority Authority copied from the mapping, when declared.
	Authority *string `json:"authority,omitempty"`

	// Citation Supporting citation copied from the mapping, when declared.
	Citation *string `json:"citation,omitempty"`

	// Disposition Mapping decision, `unmapped` when the artifact has no rule, or `not-selected` when the assertion is covered only by a selection candidate that lost.
	//
	// `not-selected` is not an `omit`. An `omit` is a reviewed decision to drop the assertion. `not-selected` says a selection chose a different candidate, so nothing decided this assertion on its own. Like an `omit` it emits no targets and leaves the translation settled.
	Disposition CrosswalkOutcome_Disposition `json:"disposition"`

	// Guidance Operator guidance copied from the mapping, when declared.
	Guidance *string `json:"guidance,omitempty"`

	// MappingID Matched mapping id, absent when no mapping covers the assertion. On a `not-selected` outcome, `<selection>:<candidate>` names the candidate that lost.
	MappingID *string `json:"mappingId,omitempty"`

	// Reason Reviewable explanation of the decision.
	Reason string `json:"reason"`

	// Source Compact reference to an assertion without engine-production metadata.
	Source AssertionReference `json:"source"`

	// Targets Emitted targets for `carry`, proposed targets for `review`.
	Targets []CrosswalkTarget `json:"targets"`
}

// CrosswalkOutcomeDisposition1 defines model for CrosswalkOutcome.Disposition.1.
type CrosswalkOutcomeDisposition1 = string

// CrosswalkOutcomeDisposition2 defines model for CrosswalkOutcome.Disposition.2.
type CrosswalkOutcomeDisposition2 = string

// CrosswalkOutcome_Disposition Mapping decision, `unmapped` when the artifact has no rule, or `not-selected` when the assertion is covered only by a selection candidate that lost.
//
// `not-selected` is not an `omit`. An `omit` is a reviewed decision to drop the assertion. `not-selected` says a selection chose a different candidate, so nothing decided this assertion on its own. Like an `omit` it emits no targets and leaves the translation settled.
type CrosswalkOutcome_Disposition struct {
	union json.RawMessage
}

// CrosswalkSelectionCondition Boolean test over source assertions used to choose a crosswalk candidate.
type CrosswalkSelectionCondition struct {
	// All Every referenced assertion must be present.
	All *[]AssertionReference `json:"all,omitempty"`

	// Any At least one referenced assertion must be present.
	Any *[]AssertionReference `json:"any,omitempty"`

	// None None of the referenced assertions may be present.
	None *[]AssertionReference `json:"none,omitempty"`
}

// CrosswalkSelectionForbidMatch A selected cross-field candidate that prohibited a crosswalk translation.
type CrosswalkSelectionForbidMatch struct {
	// Authority Authority copied from the candidate, when declared.
	Authority *string `json:"authority,omitempty"`

	// CandidateID Selected candidate id.
	CandidateID string `json:"candidateId"`

	// Citation Supporting citation copied from the candidate, when declared.
	Citation *string `json:"citation,omitempty"`

	// Guidance Operator guidance copied from the candidate, when declared.
	Guidance *string `json:"guidance,omitempty"`

	// Kind A selected cross-field candidate prohibited the translation.
	Kind string `json:"kind"`

	// Reason Human-readable basis for the prohibition.
	Reason string `json:"reason"`

	// SelectionID Selection definition id.
	SelectionID string `json:"selectionId"`

	// Sources Source assertions explicitly covered by the candidate.
	Sources []AssertionReference `json:"sources"`
}

// CrosswalkSelectionOutcome Selected branch of one cross-field crosswalk decision.
type CrosswalkSelectionOutcome struct {
	// Authority Authority copied from the selected candidate, when declared.
	Authority *string `json:"authority,omitempty"`

	// CandidateID Candidate chosen by the selection strategy.
	CandidateID string `json:"candidateId"`

	// Citation Supporting citation copied from the selected candidate, when declared.
	Citation *string `json:"citation,omitempty"`

	// Disposition Decision a reviewed crosswalk makes for one source assertion.
	Disposition CrosswalkDisposition `json:"disposition"`

	// Guidance Operator guidance copied from the selected candidate, when declared.
	Guidance *string `json:"guidance,omitempty"`

	// Reason Human-readable basis copied from the selected candidate.
	Reason string `json:"reason"`

	// SelectionID Selection definition id.
	SelectionID string `json:"selectionId"`

	// Targets Targets emitted for `carry`, or proposed for `review`.
	Targets []CrosswalkTarget `json:"targets"`
}

// CrosswalkSelectionStrategy Strategy for choosing one candidate from a reviewed cross-field decision.
type CrosswalkSelectionStrategy string

// CrosswalkTarget One target assertion a crosswalk can emit or propose.
type CrosswalkTarget struct {
	// CategoryID Category containing the asserted concept.
	CategoryID string `json:"categoryId"`

	// ConceptID Asserted concept identifier.
	ConceptID string `json:"conceptId"`

	// CopyValue Copy the source assertion's captured value onto this target.
	CopyValue *bool `json:"copyValue,omitempty"`

	// Value Captured value for value-bearing concepts.
	Value *string `json:"value,omitempty"`
}

// CrosswalkTranslationResult Result of applying a reviewed directional crosswalk.
type CrosswalkTranslationResult struct {
	union json.RawMessage
}

// CtxAccessCheckRequestBody POST /v1/ctx/{ctxId}/{version?}/access/check. Three inputs, matching the engine's `AccessEngine.check(dataAssertions, userAttributes, context)` contract. First, the TARGET being access-checked: either a rendered `marking` string or its structured `assertions` array form (a marking string is parsed to assertions first). Provide either or both; when both are present the assertions derived from `marking` are merged with the `assertions` array. Second, the SUBJECT: `userAttributes`, always an object of the user's attributes. Third, `context`: extra comparison data (e.g. distribution orgs) not encoded in the marking. When decision audit is enabled, deployment configuration identifies required user and resource ids inside `context`.
type CtxAccessCheckRequestBody struct {
	// Assertions TARGET as structured assertions: the alternative to the `marking` string for callers that already hold parsed assertions. Merged with any assertions derived from `marking`. This is the single structured target field.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Context Optional contextual attributes for access evaluation (third check argument).
	Context *map[string]interface{} `json:"context,omitempty"`

	// Marking TARGET marking string to check access against. Parsed to data assertions before evaluation. Provide this OR `assertions` (the structured alternative); both feed the same first argument of the access check.
	Marking *string `json:"marking,omitempty"`

	// UserAttributes User/principal attributes evaluated by the access policy. Attribute paths and value interpretation are defined by that access policy. This is the only user-attribute input.
	UserAttributes map[string]interface{} `json:"userAttributes"`
}

// CtxAccessDescribeRequirementsRequestBody POST /v1/ctx/{ctxId}/{version?}/access/describe-requirements
type CtxAccessDescribeRequirementsRequestBody struct {
	// Assertions Structured target assertions merged with any parsed marking.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Marking Marking parsed through the Context's bound scheme.
	Marking *string `json:"marking,omitempty"`
}

// CtxAccessRulesResponseBody Response body for GET /v1/ctx/{ctxId}/{version?}/access/rules: every rule in the context's access policy, with the policy's own marking.
type CtxAccessRulesResponseBody struct {
	// Marking The access policy document's own overall marking, exactly as its `access.marking` writes it. Display only: never interpreted. Omitted when the document declares none.
	Marking *string `json:"marking,omitempty"`

	// Rules Descriptions for every rule in the access policy, in declared order.
	Rules []AccessRuleDescription `json:"rules"`
}

// CtxComponentDescriptor Component reference echoed in a `CtxDescribeResponseBody`.
type CtxComponentDescriptor struct {
	// ID Component id from the context declaration.
	ID string `json:"id"`

	// Layers Layers nested beneath this component, in composition order.
	Layers *[]CtxComponentDescriptor `json:"layers,omitempty"`

	// ResolvedVersion Concrete version after `latest` resolution.
	ResolvedVersion string `json:"resolvedVersion"`

	// Version Requested version literal (concrete semver or 'latest').
	Version string `json:"version"`
}

// CtxCrosswalkDescribeResponseBody Response body for GET /v1/ctx/{ctxId}/{version?}/crosswalk/describe: every crosswalk the context holds.
//
// A dedicated read rather than a member of the context description, because a crosswalk's reviewed mappings run to tens of kilobytes and every client reads the description when it connects, translating or not.
type CtxCrosswalkDescribeResponseBody struct {
	// Crosswalks Every crosswalk in the context, in declaration order.
	Crosswalks []ContextCrosswalkDescriptor `json:"crosswalks"`
}

// CtxCrosswalkTranslateRequestBody POST /v1/ctx/{ctxId}/{version?}/crosswalk/translate. The adapter takes the SOURCE as one field holding either a marking string or its parsed assertions; the wire splits that into the same two fields `CtxAccessCheckRequestBody` uses for its target, because a field that is a string or an array is a shape the MCP catalog builder refuses to describe as a tool input. Exactly one of `marking` and `assertions` is required: neither is a request with no source, both is two sources with no rule for which one the crosswalk reads, and either is a 400.
type CtxCrosswalkTranslateRequestBody struct {
	// Assertions SOURCE as structured assertions, for callers that already hold the parsed form. Provide this OR `marking`, never both.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// CrosswalkID Crosswalk artifact id declared by the context.
	CrosswalkID string `json:"crosswalkId"`

	// Marking SOURCE marking text, parsed against the crosswalk's source scheme before translation. Provide this OR `assertions`, never both.
	Marking *string `json:"marking,omitempty"`
}

// CtxDescribeResponseBody Response body for GET /v1/ctx/{ctxId}/{version}: decomposed context.
type CtxDescribeResponseBody struct {
	// Access Component reference echoed in a `CtxDescribeResponseBody`.
	Access *CtxComponentDescriptor `json:"access,omitempty"`

	// Crosswalks Crosswalk components in declaration order, or omitted when the context declares none. A crosswalk the context follows at `latest` reports here the concrete version that answered. This is provenance; the source and target schemes and the reviewed mappings are on `crosswalk/describe`.
	Crosswalks *[]CtxComponentDescriptor `json:"crosswalks,omitempty"`

	// Description Optional human-readable description copied from the declaration.
	Description *string `json:"description,omitempty"`

	// Guides Guide components in precedence order, or omitted when none.
	Guides *[]CtxComponentDescriptor `json:"guides,omitempty"`

	// ID Context id.
	ID string `json:"id"`

	// Scheme Component reference echoed in a `CtxDescribeResponseBody`.
	Scheme CtxComponentDescriptor `json:"scheme"`

	// Schemes Every scheme this Context holds, primary first, each with the selector that addresses it.
	//
	// The scheme sub-resources under this Context take `?scheme={selector}`, and this is where the valid values come from. `scheme` above answers a different question: it names the component the Context resolved and the layers stacked on it, which is provenance rather than a selector.
	Schemes []ContextSchemeDescriptor `json:"schemes"`

	// Version Concrete context version after `latest` resolution.
	Version string `json:"version"`
}

// CtxDocumentAuditRequestBody POST /v1/ctx/{ctxId}/{version?}/document/audit
type CtxDocumentAuditRequestBody struct {
	// GuideDeterminations Set-semantic analyst input supplied to guide evaluation.
	GuideDeterminations *GuideDeterminationInput `json:"guideDeterminations,omitempty"`

	// ResponseProfile Controls how much evidence detail a document operation returns.
	ResponseProfile *DocumentResponseProfile `json:"responseProfile,omitempty"`

	// Text Full document text to scan for markings.
	Text string `json:"text"`
}

// CtxGuideEvaluateRequestBody POST /v1/ctx/{ctxId}/{version?}/guide/evaluate
//
// Two forms, exactly one per request: `guideId` with `selectedRuleIds` evaluates one guide, and `selections` rolls up a selection that spans guides. Any other combination is a 400. The rule is stated here rather than as a root `oneOf`, which would make every generated SDK model this request as a union of two request types.
type CtxGuideEvaluateRequestBody struct {
	// GuideID ID of the one guide to evaluate, with `selectedRuleIds`. Refused with `selections`.
	GuideID *string `json:"guideId,omitempty"`

	// SelectedRuleIds IDs of the rules selected in the guide `guideId` names. Required with `guideId`, refused with `selections`.
	SelectedRuleIds *[]string `json:"selectedRuleIds,omitempty"`

	// Selections The selected rules, each named by its guide, rolled up across guides: each guide evaluates the rules selected from it, and the outcomes are rolled up per target scheme. No guide overrides or depends on another, so every selected rule's outcome counts. Refused with `guideId` or `selectedRuleIds`.
	Selections *[]GuideRuleSelection `json:"selections,omitempty"`
}

// CtxGuideScanRequestBody POST /v1/ctx/{ctxId}/{version?}/guide/scan
type CtxGuideScanRequestBody struct {
	// GuideID Optional guide id to scope scanning to a single guide.
	GuideID *string `json:"guideId,omitempty"`

	// Text Text to scan against the context's guides.
	Text string `json:"text"`
}

// CtxGuideViewRequestBody POST /v1/ctx/{ctxId}/{version?}/guide/view
type CtxGuideViewRequestBody struct {
	// GuideID Guide id when the context has multiple guides. Defaults to the first guide.
	GuideID *string `json:"guideId,omitempty"`
}

// CtxListResponseBody Response body for GET /v1/ctx: the context ids this caller may use.
type CtxListResponseBody struct {
	// Contexts Context ids the caller may use, sorted alphabetically.
	Contexts []string `json:"contexts"`
}

// CtxSchemeAggregateRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/aggregate
type CtxSchemeAggregateRequestBody struct {
	// AssertionSets Already-parsed assertion sets to aggregate, one inner array per marking. Skips the string-parse stage; the rest of the aggregate pipeline runs identically to `inputs`. Transports marking state losslessly (scope-only assertions and capture values survive). Mutually exclusive with `inputs`.
	AssertionSets *[][]Assertion `json:"assertionSets,omitempty"`

	// Context Caller-supplied context naming the role the marking being read plays. No context, or `role: 'whole'`, means a whole: a marking that stands for everything it covers. `role: 'part'` means a part that belongs to a whole, named in `partOf` when it is known.
	//
	// On a rollup the context names the contributors' role, never the rollup's: the rollup is always a whole. So no context rolls wholes up into an aggregate, and `{ role: 'part', partOf }` rolls a whole's parts up. See `SchemeEngineAdapter.aggregate`.
	Context *ProcessingContext `json:"context,omitempty"`

	// InputFormat The format the strings in `inputs` are written in. For a scheme that declares `render.surfaces`, the view they are read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out. Refused with `assertionSets`, which carries no text to read.
	InputFormat *string `json:"inputFormat,omitempty"`

	// Inputs Array of marking strings to aggregate. Mutually exclusive with `assertionSets`: supply exactly one.
	Inputs *[]string `json:"inputs,omitempty"`

	// OutputFormat The format the combined marking is rendered in. Omitted or empty, the scheme's default format. It selects output only. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Surface Validation surface of the rollup, using the scheme's scope ids. Names what the combined marking is, not what any one contributor was. Omitted, the rollup is a whole marking and every required category belonging to the marking as a whole must be covered by the combined result: a contributor that omits one is not faulted for it. Given a scope id, the inputs are fragments of that one surface and only the required categories that surface owns are demanded of the rollup.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeCheckCoverageRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/check-coverage
type CtxSchemeCheckCoverageRequestBody struct {
	// Children Child marking strings that must be covered by the parent.
	Children []string `json:"children"`

	// InputFormat The format `children` and `parent` are written in. For a scheme that declares `render.surfaces`, the view they are read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The format `rendered` is written in. Omitted or empty, the scheme's default format. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Parent Parent marking string checked against the children's rollup. It covers them when it lacks nothing the rollup needs and carries nothing the rollup does not account for.
	Parent string `json:"parent"`
}

// CtxSchemeDiffRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/diff
type CtxSchemeDiffRequestBody struct {
	// A First marking string.
	A string `json:"a"`

	// B Second marking string.
	B string `json:"b"`

	// InputFormat The format `a` and `b` are written in. For a scheme that declares `render.surfaces`, the view they are read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out.
	//
	// There is no output format: the response is three assertion arrays and no string, so nothing here renders.
	InputFormat *string `json:"inputFormat,omitempty"`
}

// CtxSchemeExplainRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/explain: the trace of a process call.
//
// Takes what process takes, apart from `outputs` and `suggestFor`, which only add to a process result.
type CtxSchemeExplainRequestBody struct {
	// Assertions The facts already chosen, settled together with what `input` states as one marking.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Context Caller-supplied context naming the role the marking being read plays. No context, or `role: 'whole'`, means a whole: a marking that stands for everything it covers. `role: 'part'` means a part that belongs to a whole, named in `partOf` when it is known.
	//
	// On a rollup the context names the contributors' role, never the rollup's: the rollup is always a whole. So no context rolls wholes up into an aggregate, and `{ role: 'part', partOf }` rolls a whole's parts up. See `SchemeEngineAdapter.aggregate`.
	Context *ProcessingContext `json:"context,omitempty"`

	// Input The marking text, as typed.
	Input string `json:"input"`

	// InputFormat The format `input` is written in. For a scheme that declares `render.surfaces`, the view it is read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The output the trace's `rendered` is written in. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it. Omitted or empty, the scheme's default format, or with `render.surfaces`, each written surface's default for the marking's role, by the rule `render.surfaces` states for a request that names no output.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Parts Part markings the whole must cover. The whole is judged as covering them, as on the process request.
	Parts *[]string `json:"parts,omitempty"`

	// Surface Validation surface: names which surface of a marking `input` is, using the scheme's scope ids. Omitted, the input is a whole marking.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeNoticesRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/notices
type CtxSchemeNoticesRequestBody struct {
	// Assertions Assertions from a prior process() call whose concept notices should be collected.
	Assertions []Assertion `json:"assertions"`
}

// CtxSchemeProcessRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/process: the author's call.
//
// `input` is what was typed and `assertions` the facts already chosen, and a composer sends both on each keystroke. The two are settled together as one marking, so a conflict between them is the scheme's to repair or report, as it would be within one input. Either may be empty: an empty `input` with `assertions` processes the chosen facts alone.
type CtxSchemeProcessRequestBody struct {
	// Assertions The facts already chosen, settled together with what `input` states as one marking.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Context Caller-supplied context naming the role the marking being read plays. No context, or `role: 'whole'`, means a whole: a marking that stands for everything it covers. `role: 'part'` means a part that belongs to a whole, named in `partOf` when it is known.
	//
	// On a rollup the context names the contributors' role, never the rollup's: the rollup is always a whole. So no context rolls wholes up into an aggregate, and `{ role: 'part', partOf }` rolls a whole's parts up. See `SchemeEngineAdapter.aggregate`.
	Context *ProcessingContext `json:"context,omitempty"`

	// Input The marking text, as typed.
	Input string `json:"input"`

	// InputFormat The format `input` is written in. For a scheme that declares `render.surfaces`, the view it is read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out and reports what the text was read as on `recognition.inputFormat`.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The output the result is rendered in. It selects output only: it never narrows what may be read. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	//
	// Omitted or empty, the scheme's default format, or with `render.surfaces`, each written surface's default for the marking's role, by the rule `render.surfaces` states for a request that names no output.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Outputs The outputs to write into `renders`, each an output address spelled exactly as a key of `renders` is: `surface/view` for a scheme that declares `render.surfaces`, a format id otherwise. Omitted or empty, `renders` is empty. An address the scheme does not declare is refused.
	Outputs *[]string `json:"outputs,omitempty"`

	// Parts Part markings the whole must cover. The whole is judged as covering them: a part it does not cover is an error, and so is a part that does not read as a marking. Every rank floor is worked out against the whole and these parts together, so a raise never lands below what a part needs, and every proposed fix is proven with them, as `ProcessResult.proposedFixes` states. Each entry is a marking in one of the scheme's part formats, or any format the scheme reads. Empty or absent, the whole stands alone.
	Parts *[]string `json:"parts,omitempty"`

	// SuggestFor Category ids to answer picker options for, against the settled facts, in `suggestionsFor`. Omitted, nothing is suggested and the field is absent from the result. An id the scheme does not declare is refused.
	SuggestFor *[]string `json:"suggestFor,omitempty"`

	// Surface Validation surface: names which surface of a marking `input` is, using the scheme's scope ids. Omitted, the input is a whole marking and every required category belonging to the marking as a whole must be covered. Given a scope id, the input is that surface's fragment and only the required categories that surface owns are demanded.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeRenderRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/render
type CtxSchemeRenderRequestBody struct {
	// Assertions Assertions to render into a marking string.
	Assertions []Assertion `json:"assertions"`

	// OutputFormat The format the assertions are rendered in. Omitted or empty, the scheme's default format. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Preservation What a caller is asking a render for.
	//
	// `'projection'` is the ordinary human projection: the format writes the fields it declares and the assessment reports the rest. `'full'` is section 7.4's full-preservation request, which fails rather than returning a string when facts were omitted, conflicting, or unrepresentable. Naming the request is what makes the failure possible: a render that only ever returns a string has nowhere to say that the string is not the state.
	Preservation *RenderPreservation `json:"preservation,omitempty"`
}

// CtxSchemeScanRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/scan
type CtxSchemeScanRequestBody struct {
	// InputFormat The format the markings in `text` are written in. For a scheme that declares `render.surfaces`, the view they are read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The output each found marking is rendered back in. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it. Omitted or empty, each marking is written as `ScanResult.type` describes.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Text Text to scan for embedded markings.
	Text string `json:"text"`
}

// CtxSchemeSuggestForCategoriesRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/suggest-for-categories
type CtxSchemeSuggestForCategoriesRequestBody struct {
	// Assertions The caller's full committed assertion set. Each category is answered as if its own assertions were not present.
	Assertions []Assertion `json:"assertions"`

	// CategoryIds Category ids to produce suggestions for.
	CategoryIds []string `json:"categoryIds"`

	// OutputFormat The format suggestion previews are rendered in. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it. An empty one is the same as none.
	//
	// There is no input format: this takes a committed assertion set, never text.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Surface Suggestion surface, by scheme scope id: restricts suggestions to the concepts and categories that surface owns.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeSuggestRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/suggest: the options a picker may offer.
//
// The marking arrives as text in `input` or as facts in `assertions`, exactly one of the two; neither or both is a 400. Presence decides, as on validate: an empty `input` is text with nothing typed yet, and an empty `assertions` is facts with none chosen. Text typed beside facts already chosen is what process answers, with `suggestFor`.
type CtxSchemeSuggestRequestBody struct {
	// Assertions The marking as facts already chosen. Provide this or `input`, never both.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Filter Filter string to narrow suggestions.
	Filter *string `json:"filter,omitempty"`

	// IncludeGated Include concepts the active filter would otherwise hide because their prerequisites are unmet. Each comes back with `gated: true` and the `requires` ids that would unlock it.
	IncludeGated *bool `json:"includeGated,omitempty"`

	// Input The marking as text, which may be empty to ask for every option. Provide this or `assertions`, never both.
	Input *string `json:"input,omitempty"`

	// InputFormat The format `input` is written in, when it is a marking string. For a scheme that declares `render.surfaces`, the view it is read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out. Refused with `assertions`, which carries no text to read.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The format suggestion previews are rendered in. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it. An empty one is the same as none.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Surface Suggestion surface, by scheme scope id: restricts suggestions to the concepts and categories that surface owns.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeValidateRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/validate: the reader's call.
//
// The marking arrives as text in `input` or as facts in `assertions`, exactly one of the two; neither or both is a 400. They are two fields because the MCP catalog builder refuses a field that is a string or an array, and the rule is stated here rather than as a root `oneOf`, which would make every generated SDK model this request as a union of two request types.
//
// Nothing here changes how the engine settles the marking, so `valid` and `inputWasValid` mean what they mean on every other route. `surface`, `context` and `parts` say what the marking is judged as, as on process.
type CtxSchemeValidateRequestBody struct {
	// Assertions The marking as facts as stated: what a document or an author states, before any engine settled them. Never the `assertions` of an earlier result, which are settled already and can read valid where the text they came from did not. Provide this or `input`, never both.
	Assertions *[]Assertion `json:"assertions,omitempty"`

	// Context Caller-supplied context naming the role the marking being read plays. No context, or `role: 'whole'`, means a whole: a marking that stands for everything it covers. `role: 'part'` means a part that belongs to a whole, named in `partOf` when it is known.
	//
	// On a rollup the context names the contributors' role, never the rollup's: the rollup is always a whole. So no context rolls wholes up into an aggregate, and `{ role: 'part', partOf }` rolls a whole's parts up. See `SchemeEngineAdapter.aggregate`.
	Context *ProcessingContext `json:"context,omitempty"`

	// Input The marking as text. Provide this or `assertions`, never both.
	Input *string `json:"input,omitempty"`

	// InputFormat The format `input` is written in. For a scheme that declares `render.surfaces`, the view it is read as, or one of its forms, as the scheme document's `render.surfaces` defines them. Omitted or empty, recognition works it out and reports what the text was read as on `recognition.inputFormat`. Refused with `assertions`, which carries no text to read.
	InputFormat *string `json:"inputFormat,omitempty"`

	// OutputFormat The output the result is rendered in. It selects output only: it never narrows what may be read. For a scheme that declares `render.surfaces`, an output as the scheme document's `render.surfaces` defines it.
	//
	// Omitted or empty, the scheme's default format, or with `render.surfaces`, each written surface's default for the marking's role, by the rule `render.surfaces` states for a request that names no output.
	OutputFormat *string `json:"outputFormat,omitempty"`

	// Parts Part markings the whole must cover. The whole is judged as covering them: a part it does not cover is an error, and so is a part that does not read as a marking. Each entry is a marking in one of the scheme's part formats, or any format the scheme reads. Empty or absent, the whole stands alone.
	Parts *[]string `json:"parts,omitempty"`

	// Surface Validation surface: names which surface of a marking the input is, using the scheme's scope ids. Omitted, the input is a whole marking and every required category belonging to the marking as a whole must be covered. Given a scope id, the input is that surface's fragment and only the required categories that surface owns are demanded.
	Surface *string `json:"surface,omitempty"`
}

// CtxSchemeViewRequestBody POST /v1/ctx/{ctxId}/{version?}/scheme/view
type CtxSchemeViewRequestBody = map[string]interface{}

// CtxVersionsResponseBody Response body for GET /v1/ctx/{ctxId}: versions for a given context id.
type CtxVersionsResponseBody struct {
	// ID Echoed context id.
	ID string `json:"id"`

	// Versions Every installed version the caller may use, sorted highest-first.
	Versions []string `json:"versions"`
}

// DecisionProposalOrigin Describes how a decision proposal entered the analyst workflow.
type DecisionProposalOrigin string

// DiffResult Comparison of two assertion sets: which assertions were added, removed, or shared.
type DiffResult struct {
	// Added Assertions in the second marking and not the first.
	Added []Assertion `json:"added"`

	// Common Assertions in both markings, as the first marking holds them. Concept and category decide sameness; confidence, source and value are not compared.
	Common []Assertion `json:"common"`

	// Removed Assertions in the first marking and not the second.
	Removed []Assertion `json:"removed"`
}

// DocumentAuditReport Authored structural and analyst-determined guide audit of one document.
type DocumentAuditReport struct {
	// DocumentDigest SHA-256 digest of the exact accepted UTF-8 document bytes.
	DocumentDigest string `json:"documentDigest"`

	// GuideAssessment Guide compilation dimension, or null when no guide applies.
	GuideAssessment *DocumentGuideAssessment `json:"guideAssessment"`

	// Limitations Document-level assessment limitations.
	Limitations []AssessmentLimitation `json:"limitations"`

	// MarkingRequirements Canonical marking requirements evaluated for this document from the loaded scheme. Without declared conditions, whole defaults to false and part to true.
	MarkingRequirements DocumentMarkingRequirements `json:"markingRequirements"`

	// OverallVerdict Structural verdict derived from authored document markings.
	OverallVerdict DocumentOverallVerdict `json:"overallVerdict"`

	// Parts Authored part markings and their governed regions.
	Parts []DocumentPartAssessment `json:"parts"`

	// ProceduralFlags Stable authored-procedure findings.
	ProceduralFlags []string `json:"proceduralFlags"`

	// UnmarkedSpans Authored ranges not governed by a marking.
	UnmarkedSpans []DocumentUnmarkedSpan `json:"unmarkedSpans"`

	// Vocabulary Scheme-provided display vocabulary.
	Vocabulary *Vocabulary `json:"vocabulary"`

	// Whole First authored document-wide marking, or null when absent.
	Whole *string `json:"whole"`

	// WholeCoversParts Authored whole coverage of authored parts.
	WholeCoversParts *CheckCoverageResult `json:"wholeCoversParts"`

	// WholeValidation How the first authored whole was judged, as a whole that must cover the authored parts that read as markings, with the facts it states. Null when the document has no authored whole.
	WholeValidation *MarkingValidation `json:"wholeValidation"`

	// Wholes Every authored document-wide marking occurrence.
	Wholes []DocumentWholeOccurrence `json:"wholes"`
}

// DocumentGuideAssessment Guide dimension of a document audit.
type DocumentGuideAssessment struct {
	// AcceptedCount Number of accepted nominations.
	AcceptedCount int `json:"acceptedCount"`

	// DeterminedRequiredWhole Whole-document requirement from accepted/manual facts.
	DeterminedRequiredWhole *GuideCompiledResult `json:"determinedRequiredWhole"`

	// DeterminedWholeCoverage Authored whole coverage of the determined requirement.
	DeterminedWholeCoverage *CheckCoverageResult `json:"determinedWholeCoverage"`

	// Evaluation Complete evidence and compilation result for document guide evaluation.
	Evaluation *GuideContentEvaluation `json:"evaluation,omitempty"`

	// EvidenceSpanCount Number of evidence spans.
	EvidenceSpanCount int `json:"evidenceSpanCount"`

	// NominationCount Number of lexical nominations.
	NominationCount int `json:"nominationCount"`

	// ProjectedRequiredWhole Advisory whole-document requirement including unresolved facts.
	ProjectedRequiredWhole *GuideCompiledResult `json:"projectedRequiredWhole"`

	// ProjectedWholeCoverage Authored whole coverage of the projected requirement.
	ProjectedWholeCoverage *CheckCoverageResult `json:"projectedWholeCoverage"`

	// RejectedCount Number of rejected nominations.
	RejectedCount int `json:"rejectedCount"`

	// Status Closed status of the analyst-determined guide compilation dimension.
	Status GuideCompilationStatus `json:"status"`

	// UnitCoverage Determined/projected coverage for each evaluation unit.
	UnitCoverage []GuideScopeCoverage `json:"unitCoverage"`

	// UnresolvedCount Number of unresolved nominations.
	UnresolvedCount int `json:"unresolvedCount"`
}

// DocumentInputErrorDetail Stable INVALID_INPUT detail values emitted by document operations.
type DocumentInputErrorDetail string

// DocumentMarkingRequirements Canonical marking requirements evaluated for this document from the loaded scheme. Without declared conditions, whole defaults to false and part to true.
type DocumentMarkingRequirements struct {
	// Part Whether authored part markings are required for document text.
	Part bool `json:"part"`

	// Whole Whether an authored document-wide marking is required.
	Whole bool `json:"whole"`
}

// DocumentOverallVerdict Structural verdict derived from authored document markings.
type DocumentOverallVerdict string

// DocumentPartAssessment One authored part marking, the region it governs, and how it was judged, facts included.
type DocumentPartAssessment struct {
	// End Zero-based exclusive UTF-8 byte offset.
	End int `json:"end"`

	// Marking The part marking as the document writes it.
	Marking string `json:"marking"`

	// OffsetEncoding Frozen byte-offset encoding.
	OffsetEncoding string `json:"offsetEncoding"`

	// ScopeID The scope id the part is evaluated under, `part-` followed by its one-based position in the document. Guide results for the part carry the same id.
	ScopeID string `json:"scopeId"`

	// Start Zero-based inclusive UTF-8 byte offset.
	Start int `json:"start"`

	// Text The text the part marking governs. Omitted outside the full response profile.
	Text *string `json:"text,omitempty"`

	// Validation How an audit judged one authored marking, whole or part.
	//
	// `valid` and `inputWasValid` mean exactly what `ValidateResult` defines them to mean. An audit asks whether the marking was right as the author wrote it, and that is `inputWasValid`; `valid` says whether it stands once the engine has settled it, as it does everywhere else.
	//
	// It carries the marking's facts and how far recognition got with it, so a caller that rolls markings up, checks access or offers a correction works from the audit instead of reading each marking a second time.
	Validation MarkingValidation `json:"validation"`
}

// DocumentResponseCounts Counts the full, unprojected semantic result of a document operation.
type DocumentResponseCounts struct {
	// EvidenceSpans Number of evidence spans.
	EvidenceSpans int `json:"evidenceSpans"`

	// GuideUnits Number of evaluated guide units.
	GuideUnits int `json:"guideUnits"`

	// Nominations Number of lexical nominations.
	Nominations int `json:"nominations"`
}

// DocumentResponseProfile Controls how much evidence detail a document operation returns.
type DocumentResponseProfile string

// DocumentResponseProjection Describes the explicit projection applied after full evaluation.
type DocumentResponseProjection struct {
	// Counts Counts the full, unprojected semantic result of a document operation.
	Counts DocumentResponseCounts `json:"counts"`

	// DetailsComplete True only when the full profile returned every detail.
	DetailsComplete bool `json:"detailsComplete"`

	// OmittedFields JSON-pointer paths omitted by the selected profile.
	OmittedFields []string `json:"omittedFields"`

	// Profile Controls how much evidence detail a document operation returns.
	Profile DocumentResponseProfile `json:"profile"`
}

// DocumentUnmarkedSpan A document range not governed by an authored marking.
type DocumentUnmarkedSpan struct {
	// End Zero-based exclusive offset in the accepted document.
	End int `json:"end"`

	// OffsetEncoding Frozen byte-offset encoding.
	OffsetEncoding string `json:"offsetEncoding"`

	// Start Zero-based inclusive offset in the accepted document.
	Start int `json:"start"`

	// Text Authored text, omitted outside the full response profile.
	Text *string `json:"text,omitempty"`
}

// DocumentWholeOccurrence One authored document-wide marking occurrence.
type DocumentWholeOccurrence struct {
	// End Zero-based exclusive UTF-8 byte offset.
	End int `json:"end"`

	// Marking The whole marking as the document writes it.
	Marking string `json:"marking"`

	// OffsetEncoding Frozen byte-offset encoding.
	OffsetEncoding string `json:"offsetEncoding"`

	// Start Zero-based inclusive UTF-8 byte offset.
	Start int `json:"start"`
}

// EachAffix Affix written around every value of a field, rather than once around all of them.
type EachAffix struct {
	// Prefix Written before each value.
	Prefix *string `json:"prefix,omitempty"`

	// Suffix Written after each value.
	Suffix *string `json:"suffix,omitempty"`
}

// EmbeddedMarkingStyle One way a marking can appear embedded inside free text, e.g. a parenthesized short form like "(H)" at the start of a paragraph. A scheme may declare multiple styles; all are matched during scanning.
type EmbeddedMarkingStyle struct {
	// AutoInsert When a host UI should proactively prompt for an embedded marking.
	AutoInsert *EmbeddedMarkingStyleAutoInsert `json:"autoInsert,omitempty"`

	// Pattern RegExp source (no flags) whose first capture group is the raw marking text. Example: `\(([^)]+)\)` for a marking in parentheses.
	Pattern string `json:"pattern"`

	// Template Template used when inserting a new embedded marking. Must include the literal token "{marking}" where the rendered marking is substituted. Example: "({marking}) ".
	Template string `json:"template"`
}

// EmbeddedMarkingStyleAutoInsert When a host UI should proactively prompt for an embedded marking.
type EmbeddedMarkingStyleAutoInsert = string

// ErrorResponse Response body for error responses (4xx / 5xx).
type ErrorResponse struct {
	// Code Stable machine-readable classification of the failure, present whenever the server can name one. Set to `INVALID_INPUT` when the failure originated in engine input or the request carried something the operation does not declare, to `route_not_found` on a 404 that reached routing and matched no route, and to `scheme_not_found` on a 404 for a `?scheme=` key the Context does not hold. A request refused before routing carries the constant refusal body instead, which has no `code` at all by design, so the absence of a code on a 404 is itself uninformative.
	Code *string `json:"code,omitempty"`

	// Detail What the failure is about, beside `code`: a stable operation-specific detail used to classify it. For a request that names a format, a concept or a category the scheme does not declare, it is the request field and the id it named, written `field: id`, such as `outputFormat: nope` or `conceptId: NOPE`. For a request that carries a query parameter the operation does not declare, or names a declared one twice, it is `query: ` and the parameter's name; for an undeclared body field, `body`, the JSON pointer of the object that holds it, `: ` and the field, such as `body: extra` or `body/context: extra`. Never a scheme's internals.
	Detail *string `json:"detail,omitempty"`

	// Error Human-readable error message. Prose for an operator, not a contract: a client that branches on this string breaks the first time the wording is improved. Branch on `code` instead.
	Error string `json:"error"`
}

// EvaluationResult Output of GuideEngine.evaluate(): the aggregate marking produced by evaluating a set of selected guide rules. Parallels SchemeEngine.process() on the scheme side.
type EvaluationResult struct {
	// ActiveComputedRules The computed rules whose `when` condition the selection satisfied.
	ActiveComputedRules []GuideRule `json:"activeComputedRules"`

	// AggregatedMarking The effective markings rolled up into one marking by the target scheme. Empty when no rule contributed, or when the result spans more than one target scheme, in which case `aggregatedMarkings` holds one per scheme.
	AggregatedMarking string `json:"aggregatedMarking"`

	// AggregatedMarkings Aggregated guide output: one rolled-up marking per target scheme id.
	AggregatedMarkings *AggregatedMarkings `json:"aggregatedMarkings,omitempty"`

	// Breakdown One entry per evaluated rule: for each guide, its selected rules, then its computed rules.
	Breakdown []GuideBreakdownEntry `json:"breakdown"`

	// ContributingBreakdown Subset of `breakdown` that actually contributed to `aggregatedMarking` (i.e. `breakdown.filter(b => !b.wasOverridden)`). Precomputed so UI layers stop re-filtering the full breakdown.
	ContributingBreakdown []GuideBreakdownEntry `json:"contributingBreakdown"`

	// EffectiveMarkings The marking each contributing rule's outcome resolves to. Overridden rules contribute none.
	EffectiveMarkings []string `json:"effectiveMarkings"`

	// GuideID The id of the guide evaluated. Absent when the result combines several guides.
	GuideID *string `json:"guideId,omitempty"`

	// SelectedRules The selected rules that were evaluated. Rules with a `when` condition are computed, not selected, and are not listed here.
	SelectedRules []GuideRule `json:"selectedRules"`
}

// ExplanationTrace Full pipeline trace for one marking: parse input, repairs applied, inferences drawn, validation errors, and the final output.
//
// What `explain` returns, and nothing else. It takes what `process` takes and answers why the engine settled the marking the way it did, step by step. No other result carries it, because their callers want the answer rather than the working. Debug tools use this to explain why a marking was accepted or rejected. `repairs` lists each repair rule that fired and the before/after strings. `inferences` lists each inference rule that added a concept and the reason text.
type ExplanationTrace struct {
	// Entries Every step the pipeline took, in the order it took them.
	Entries []TraceEntry `json:"entries"`

	// FinalAssertions The assertions the pipeline settled on, listed in print order as a validate result's `assertions` are.
	FinalAssertions []Assertion `json:"finalAssertions"`

	// Inferences Each concept an inference rule added, in the order added.
	Inferences []TraceInference `json:"inferences"`

	// Input The marking text the explanation was asked for, as given.
	Input string `json:"input"`

	// Rendered The settled marking, written in the output format.
	Rendered string `json:"rendered"`

	// Repairs Each repair rule that rewrote the input, in the order applied.
	Repairs []TraceRepair `json:"repairs"`

	// SchemeID The id of the scheme the marking was read against.
	SchemeID string `json:"schemeId"`

	// ValidationErrors The constraints the settled marking violates.
	ValidationErrors []TraceValidationIssue `json:"validationErrors"`
}

// Fix One way to make the marking valid. Every change in it answers an error.
//
// A fix is a whole marking, not a patch: `rendered` is the complete marking the author would end up with. The engine proves each one before offering it: it reads the marking back, judges it the way it judges any input, and offers only those that read back valid.
type Fix struct {
	// Adds Concepts this fix adds because an error requires them.
	Adds []FixConceptReference `json:"adds"`

	// Keeps Concepts of the author's marking this fix keeps.
	Keeps []FixConceptReference `json:"keeps"`

	// Message What this fix does, in words.
	Message string `json:"message"`

	// Removes Concepts of the author's marking this fix removes because an error names them.
	Removes []FixConceptReference `json:"removes"`

	// Rendered The valid marking this fix produces, in the result's output format.
	//
	// It is written exactly as the result's own `rendered` is, so the two can be shown side by side. On an in-process engine with an `afterRender` plugin, that includes the plugin; the proof reads back the engine's own writing from before the plugin ran, so the plugin never decides which fixes are offered.
	Rendered string `json:"rendered"`
}

// FixAction Discrete mutation represented by an engine `Change`.
type FixAction string

// FixConceptReference Stable reference to a concept a fix or change touches.
type FixConceptReference struct {
	// CategoryID Owning category identifier.
	CategoryID string `json:"categoryId"`

	// ConceptID Concept identifier.
	ConceptID string `json:"conceptId"`
}

// FormatHeaderAssociation A concept that heads the values of another category, resolved to the list it heads.
//
// A format writes some categories as a named heading followed by the values it introduces. The pairing lives in the authored field tree, which stays server-side, so it is resolved here: which concept is the heading, which category that concept belongs to, and which category's values it heads. A UI that groups values under their heading needs all three; nothing about the grammar that produced them comes with it.
type FormatHeaderAssociation struct {
	// Category The category whose values the heading introduces.
	Category string `json:"category"`

	// Concept The concept written as the heading.
	Concept string `json:"concept"`

	// HeaderCategory The category the heading concept is asserted in.
	HeaderCategory string `json:"headerCategory"`

	// Position Whether a heading is written before or after the values it heads.
	Position FormatHeaderPosition `json:"position"`
}

// FormatHeaderPosition Whether a heading is written before or after the values it heads.
type FormatHeaderPosition string

// FormatSyntax A format's encoding: which framing it uses, and nothing about the settings that framing takes, which are authoring content.
type FormatSyntax struct {
	// Type How a format's text is framed: plain text or XML attributes on one element.
	Type FormatSyntaxType `json:"type"`
}

// FormatSyntaxType How a format's text is framed: plain text or XML attributes on one element.
type FormatSyntaxType struct {
	union json.RawMessage
}

// FormatSyntaxType0 Plain text.
type FormatSyntaxType0 = string

// FormatSyntaxType1 XML attributes on one element.
type FormatSyntaxType1 = string

// FormatView A format as the server describes it: what it writes and how it lays a marking out, not how it is built.
//
// A format's field tree is authoring content. It stays server-side, because the operations that need it (recognition, rendering, repair) are server operations, and a client that held the tree would be holding a second implementation of them. What a client legitimately needs is enough to name a format in a request, to label it in a picker, and to place a marking's parts itself: the resolved layout `RenderFormatLayout` describes, plus how the text is framed.
//
// The same shape whether a local engine or a remote client answers.
type FormatView struct {
	// CategoryEach What is written around every single value of a category, keyed by category id.
	CategoryEach *CategoryEachAffixMap `json:"categoryEach,omitempty"`

	// CategoryOrder The categories this format writes, in the order it writes them.
	CategoryOrder []string `json:"categoryOrder"`

	// CategorySeparators What is written between two values of one category, keyed by category id.
	CategorySeparators *CategorySeparatorMap `json:"categorySeparators,omitempty"`

	// CategoryTiers How a category's tiered concepts compose, keyed by category id.
	CategoryTiers *CategoryTierMap `json:"categoryTiers,omitempty"`

	// CategoryViews The concept view a category is written in, keyed by category id.
	CategoryViews *CategoryViewMap `json:"categoryViews,omitempty"`

	// DefaultFor For a scheme that declares `render.surfaces`, the roles this view is its surface's default for: the roles a marking is written in this view for when a request names none. A surface whose `default` is one view id lists both roles on that view. A surface with a default per role lists `whole` on the whole's view and `part` on the part's, both on one view when they are the same. Empty on the surface's other views, so a picker filling a part field finds the part default on any surface. Absent for a scheme whose formats are its outputs, which has one default, `defaultRenderFormat`.
	DefaultFor *[]ProcessingRole `json:"defaultFor,omitempty"`

	// Headers Headings this format writes, paired with the lists they head.
	Headers *[]FormatHeaderAssociation `json:"headers,omitempty"`

	// ID The output this entry describes, named by callers choosing an output format. For a scheme without `render.surfaces` it is a format id, which callers also name as an input format.
	//
	// For a scheme with them it is an output address, `surface/view`, and the layout described is the one that view's base format writes; a view with cases describes its base format. A view entry's `categoryOrder` and per-category maps cover only the categories its surface owns, so two views over one multi-surface format describe different category sets. A caller may also name the address as an input format, and `RenderConfigDoc.surfaces` says which forms that reads. Recognition reports the view it read the text as in `recognition.inputFormat`, the same address. These entries describe views; a view's forms are not entries of their own, and `contextFormats` lists them by role.
	ID string `json:"id"`

	// Mode How a format runs the fields it writes together.
	Mode RenderFormatMode `json:"mode"`

	// OutputView The view part of `id`, for a scheme that declares `render.surfaces`. It holds the plain view id, not the percent-encoded form `id` carries. Not a concept spelling view; that is `view`. Absent for a scheme whose formats are its outputs.
	OutputView *string `json:"outputView,omitempty"`

	// Prefix Written once before the nonempty whole, when it is written unconditionally.
	Prefix *string `json:"prefix,omitempty"`

	// Role Structural role of a marking within a processing operation.
	Role *ProcessingRole `json:"role,omitempty"`

	// Separator What this format writes between two adjacent categories by default.
	Separator *string `json:"separator,omitempty"`

	// SeparatorBefore What is written immediately before a category, keyed by category id.
	SeparatorBefore *CategoryLeadingSeparatorMap `json:"separatorBefore,omitempty"`

	// Suffix Written once after the nonempty whole, when it is written unconditionally.
	Suffix *string `json:"suffix,omitempty"`

	// Surface The semantic part of a marking this entry writes, a scheme scope id: the format's surface, or for a scheme with `render.surfaces`, the surface this output writes.
	Surface string `json:"surface"`

	// Syntax A format's encoding: which framing it uses, and nothing about the settings that framing takes, which are authoring content.
	Syntax *FormatSyntax `json:"syntax,omitempty"`

	// View Default concept view for fields that do not name one.
	View *string `json:"view,omitempty"`
}

// GuideBreakdownEntry Per-rule breakdown in a guide evaluation result.
type GuideBreakdownEntry struct {
	// IsComputed Whether the rule was activated from other guide facts.
	IsComputed bool `json:"isComputed"`

	// Outcome The effect of a guide rule: what it contributes to the marking.
	Outcome GuideRuleOutcome `json:"outcome"`

	// OverriddenBy Identifies the exact resolved semantics of one policy rule.
	OverriddenBy *RuleReference `json:"overriddenBy,omitempty"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`

	// WasOverridden Whether an active rule replaced this rule's outcome.
	WasOverridden bool `json:"wasOverridden"`
}

// GuideCompilationStatus Closed status of the analyst-determined guide compilation dimension.
type GuideCompilationStatus string

// GuideCompiledDetail Full guide and scheme breakdown retained by the full response profile.
type GuideCompiledDetail struct {
	// Evaluation Output of GuideEngine.evaluate(): the aggregate marking produced by evaluating a set of selected guide rules. Parallels SchemeEngine.process() on the scheme side.
	Evaluation EvaluationResult `json:"evaluation"`

	// Process What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Process ValidateResult `json:"process"`
}

// GuideCompiledResult Deterministic guide compilation for one authority basis and unit.
type GuideCompiledResult struct {
	// Assertions Effective compiled assertions, listed in print order as a validate result's are.
	Assertions []Assertion `json:"assertions"`

	// ComputedRules Computed rules contributing to the result.
	ComputedRules []RuleReference `json:"computedRules"`

	// Detail Full guide and scheme breakdown retained by the full response profile.
	Detail *GuideCompiledDetail `json:"detail,omitempty"`

	// DirectRules Direct rules contributing to the result.
	DirectRules []RuleReference `json:"directRules"`

	// Errors Validation errors from scheme processing.
	Errors []ValidationError `json:"errors"`

	// OverriddenRules Rules replaced by active overrides.
	OverriddenRules []RuleReference `json:"overriddenRules"`

	// Rendered Scheme-rendered marking for the compiled assertions.
	Rendered string `json:"rendered"`

	// SupportIds References into the evaluation support catalog.
	SupportIds []string `json:"supportIds"`

	// Valid Whether the compiled marking stands once settled, as `ValidateResult.valid` defines it.
	Valid bool `json:"valid"`

	// Warnings Non-fatal scheme-processing warnings.
	Warnings []Warning `json:"warnings"`
}

// GuideContentEvaluation Complete evidence and compilation result for document guide evaluation.
type GuideContentEvaluation struct {
	// AcceptedFindingIds Accepted nomination IDs.
	AcceptedFindingIds []string `json:"acceptedFindingIds"`

	// Determinations Validated analyst decisions.
	Determinations []GuideFindingDetermination `json:"determinations"`

	// DocumentDigest Digest of the exact evaluated document bytes.
	DocumentDigest string `json:"documentDigest"`

	// Limitations Guide-dimension limitations.
	Limitations []AssessmentLimitation `json:"limitations"`

	// ManualSelections Validated qualified manual selections.
	ManualSelections []ManualGuideSelection `json:"manualSelections"`

	// Nominations Canonically ordered lexical nominations.
	Nominations []GuideNomination `json:"nominations"`

	// RejectedFindingIds Rejected nomination IDs.
	RejectedFindingIds []string `json:"rejectedFindingIds"`

	// SupportCatalog De-duplicated assertion support graph.
	SupportCatalog []AssertionSupport `json:"supportCatalog"`

	// Units Compilation results for the document and authored parts.
	Units []GuideUnitEvaluation `json:"units"`

	// UnresolvedFindingIds Nomination IDs awaiting analyst decision.
	UnresolvedFindingIds []string `json:"unresolvedFindingIds"`
}

// GuideDeterminationInput Set-semantic analyst input supplied to guide evaluation.
type GuideDeterminationInput struct {
	// Findings Decisions for lexical nominations.
	Findings []GuideFindingDetermination `json:"findings"`

	// ManualSelections Qualified rules manually selected by the analyst.
	ManualSelections []ManualGuideSelectionInput `json:"manualSelections"`
}

// GuideFactKind Authority state of a fact used by guide compilation.
type GuideFactKind string

// GuideFactReference Joins assertion support to the guide fact that supplied its authority.
type GuideFactReference struct {
	// FactID Finding or manual-selection identifier.
	FactID string `json:"factId"`

	// Kind Authority state of a fact used by guide compilation.
	Kind GuideFactKind `json:"kind"`
}

// GuideFindingDecision Analyst determination for a nominated guide finding.
type GuideFindingDecision string

// GuideFindingDetermination Records an explicit analyst decision about one lexical nomination.
type GuideFindingDetermination struct {
	// Decision Analyst determination for a nominated guide finding.
	Decision GuideFindingDecision `json:"decision"`

	// DocumentDigest Digest of the exact document revision reviewed.
	DocumentDigest string `json:"documentDigest"`

	// FindingID Stable identifier of the exact nominated occurrence.
	FindingID string `json:"findingId"`

	// ProposalOrigin Describes how a decision proposal entered the analyst workflow.
	ProposalOrigin *DecisionProposalOrigin `json:"proposalOrigin,omitempty"`

	// Reason Analyst-supplied reason for the decision.
	Reason string `json:"reason"`
}

// GuideMatchSpec A named pattern matcher specification. Use built-in types like 'ipv4', 'email', 'coordinates', or 'regex' with a custom pattern.
type GuideMatchSpec struct {
	// CaseSensitive Case-sensitivity for a `type: 'regex'` leaf. When true the pattern compiles without the `i` flag, so an `[A-Z]` class means uppercase only. Absent or false, the pattern matches case-insensitively. Ignored on non-`regex` leaves.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`

	// Label Display label: overrides the matcher's built-in label.
	Label *string `json:"label,omitempty"`

	// Pattern Custom regex pattern: only used when type is 'regex'.
	Pattern *string `json:"pattern,omitempty"`

	// Runtime `runtime: true` marks a matcher provided by the host at runtime. Such a leaf is always valid at load; if the named matcher is not registered it matches nothing here, and resolves normally where it is registered.
	Runtime *bool `json:"runtime,omitempty"`

	// Type Matcher type: a built-in name (ipv4, cidr, ipv6, mac_address, email, hostname, url, coordinates, mgrs, phone, ssn, credit_card, date, uuid) or 'regex' for custom. Domain-specific packs (usgov, healthcare, financial, ...) add further names when registered.
	Type string `json:"type"`
}

// GuideNomination Evidence-preserving lexical nomination requiring analyst review.
type GuideNomination struct {
	// DocumentDigest Digest of the exact scanned document bytes.
	DocumentDigest string `json:"documentDigest"`

	// Evidence Evidence occurrences supporting the nomination.
	Evidence []ContentEvidenceSpan `json:"evidence"`

	// FindingID Deterministic identity of this evidence occurrence.
	FindingID string `json:"findingId"`

	// Origin Nomination source; lexical evidence is never authoritative itself.
	//
	// Always `lexical`, and fixed for this major version. A nomination is by definition a rule's evidence matched in the document text, and the counts, determinations and advisory results that refer to nominations all mean lexical ones. A proposal from any other source (an analyst, an agent, a host) reaches the analyst as a manual selection with its own `proposalOrigin`. A new nomination source would change what this list means, so it would arrive as a new field, not as a new value here.
	Origin string `json:"origin"`

	// ProposedAssertions Assertions the rule would propose if accepted.
	ProposedAssertions []AssertionReference `json:"proposedAssertions"`

	// Rationale Rule-authored explanation for the nomination.
	Rationale *string `json:"rationale,omitempty"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`
}

// GuideRule A rule in a marking guide: a data-element description with an outcome that contributes to the aggregate marking when activated.
//
// Rules are presented as checkboxes in the UI. `keywords` enables keyword-search filtering; `match` enables text-scan detection (e.g., finding IPv4 addresses in a document body). `outcome.rendered` or `outcome.concepts` provides the marking contribution. Rules with `when` are computed; they auto-activate based on sibling rule selection state and do not appear as user-visible checkboxes.
//
// This is the named runtime/read-result shape. A guide document authors its rules as `GuideRuleDoc` entries in `guide.rules`.
type GuideRule struct {
	// Category Guide category ID for grouping in the UI.
	Category *string `json:"category,omitempty"`

	// Description Description of the data element this rule covers.
	Description *string `json:"description,omitempty"`

	// ID Unique identifier within the guide.
	ID string `json:"id"`

	// Keywords Searchable terms for literal keyword filtering (absent on computed rules).
	Keywords *[]string `json:"keywords,omitempty"`

	// Match Named pattern matchers for content evaluation. Each entry is a `MatchNode`: a leaf spec referencing a built-in matcher (e.g. `{ type: 'ipv4' }`) or a custom regex (`{ type: 'regex', pattern: '...' }`), or a combinator node (`any` / `all` / `not` / `near` / `context` / `validate`) composed over child nodes. When content is evaluated, both `keywords` and `match` are checked.
	Match *[]MatchNode `json:"match,omitempty"`

	// Name Human-readable display name.
	Name string `json:"name"`

	// Outcome The effect of a guide rule: what it contributes to the marking.
	Outcome GuideRuleOutcome `json:"outcome"`

	// Overrides Rule IDs whose outcomes this rule replaces when active. When omitted (default), this rule's outcome is additive: it contributes alongside the rules it depends on.
	Overrides *[]string `json:"overrides,omitempty"`

	// Rationale Explanation of why this outcome applies.
	Rationale *string `json:"rationale,omitempty"`

	// When A block of conditions over the rules the user selected. Every list present must hold: `any` when one of its nodes holds, `all` when each does, `none` when none does.
	When *GuideWhen `json:"when,omitempty"`
}

// GuideRuleOutcome The effect of a guide rule: what it contributes to the marking.
type GuideRuleOutcome struct {
	// Concepts Direct concept references: converted to assertions, then engine.process()'d.
	Concepts *[]ConceptRef `json:"concepts,omitempty"`

	// Rendered Pre-rendered marking string: parsed by engine.process() during evaluation.
	Rendered *string `json:"rendered,omitempty"`
}

// GuideRuleSelection One selected rule, named by the guide that holds it.
//
// Guides in a Context are independent, so two of them may use the same rule id. A selection names the guide as well as the rule, and the two pairs are two rules.
type GuideRuleSelection struct {
	// GuideID Id of the guide that holds the rule.
	GuideID string `json:"guideId"`

	// RuleID Id of the rule within that guide.
	RuleID string `json:"ruleId"`
}

// GuideRuleView Item row rendered by the Guidebook / GuideChecklist (one per rule in a guide).
type GuideRuleView struct {
	// DependsOn Rule ids this compound rule references.
	DependsOn []string `json:"dependsOn"`

	// DependsOnNames Resolved display names for each id in `dependsOn`, index-aligned, so a "depends on" line reads the rule's name rather than its raw id. `dependsOn` stays as the machine keys.
	DependsOnNames []string `json:"dependsOnNames"`

	// Description The rule's description, when it declares one.
	Description *string `json:"description,omitempty"`

	// ID The rule id.
	ID string `json:"id"`

	// IsComputed True when the rule has a `when` condition and is activated by it.
	IsComputed bool `json:"isComputed"`

	// Keywords The keywords a text scan matches for this rule.
	Keywords []string `json:"keywords"`

	// MatcherLabels Labels like "IPv4 address", "CIDR block", already resolved via getMatcher.
	MatcherLabels []string `json:"matcherLabels"`

	// Name The rule's display name.
	Name string `json:"name"`

	// OutcomeText The rule's outcome for display: its rendered marking, or its concept ids listed in order when it has none. Empty when the outcome carries neither.
	OutcomeText string `json:"outcomeText"`

	// OverrideNames Resolved display names for each id in `overrides`, index-aligned, so the Guidebook "Overrides:" line reads rule names rather than raw ids, matching the sibling "When:" line which already resolves display. `overrides` stays as the machine keys.
	OverrideNames []string `json:"overrideNames"`

	// Overrides Rule ids this compound rule overrides when it fires.
	Overrides []string `json:"overrides"`

	// Rationale Optional rationale text from the rule definition.
	Rationale *string `json:"rationale,omitempty"`

	// WhenText The rule's `when` condition in words. Set only on a computed rule.
	WhenText *string `json:"whenText,omitempty"`
}

// GuideScanResult A guide rule matched by text scanning, with every span where the rule's keywords or matchers fired. Returned by the engine's `scanGuides` and the standalone `scanTextForRules` / `scanMultipleGuides` helpers.
//
// Parallels the scheme-side scan result: one entry per rule hit, with offsets into the scanned text.
type GuideScanResult struct {
	// GuideID The id of the guide the rule belongs to. Rule ids are unique only within a guide, so a selection built from this hit names this guide.
	GuideID string `json:"guideId"`

	// MatchedKeywords Each keyword or matcher label that fired, once each.
	MatchedKeywords []string `json:"matchedKeywords"`

	// Rule A rule in a marking guide: a data-element description with an outcome that contributes to the aggregate marking when activated.
	//
	// Rules are presented as checkboxes in the UI. `keywords` enables keyword-search filtering; `match` enables text-scan detection (e.g., finding IPv4 addresses in a document body). `outcome.rendered` or `outcome.concepts` provides the marking contribution. Rules with `when` are computed; they auto-activate based on sibling rule selection state and do not appear as user-visible checkboxes.
	//
	// This is the named runtime/read-result shape. A guide document authors its rules as `GuideRuleDoc` entries in `guide.rules`.
	Rule GuideRule `json:"rule"`

	// Spans Every place in the text where the rule fired, in the order found.
	Spans []GuideScanSpan `json:"spans"`
}

// GuideScanSpan A span, by offsets into the scanned text, where a guide rule keyword or pattern matched.
type GuideScanSpan struct {
	// End Offset in the scanned text just past the end of the match, in UTF-16 code units.
	End int `json:"end"`

	// Keyword The keyword that matched, or the label of the matcher that did.
	Keyword string `json:"keyword"`

	// Start Offset in the scanned text where the match starts, in UTF-16 code units.
	Start int `json:"start"`
}

// GuideScopeCoverage Coverage of one unit against determined and projected requirements.
type GuideScopeCoverage struct {
	// DeterminedCoverage Coverage against currently accepted/manual facts.
	DeterminedCoverage *CheckCoverageResult `json:"determinedCoverage"`

	// ProjectedCoverage Advisory coverage assuming unresolved nominations are accepted.
	ProjectedCoverage *CheckCoverageResult `json:"projectedCoverage"`

	// ScopeID Unit scope identifier.
	ScopeID string `json:"scopeId"`
}

// GuideUnitEvaluation Determined and projected compilation for one document unit.
type GuideUnitEvaluation struct {
	// Determined Compilation over accepted findings and manual selections.
	Determined *GuideCompiledResult `json:"determined"`

	// End Zero-based exclusive UTF-8 byte offset.
	End int `json:"end"`

	// FindingIds Nomination IDs whose evidence belongs to this unit.
	FindingIds []string `json:"findingIds"`

	// OffsetEncoding Frozen byte-offset encoding.
	OffsetEncoding string `json:"offsetEncoding"`

	// Projected Advisory compilation additionally assuming unresolved nominations.
	Projected *GuideCompiledResult `json:"projected"`

	// ScopeID Opaque stable scope identifier.
	ScopeID string `json:"scopeId"`

	// Start Zero-based inclusive UTF-8 byte offset.
	Start int `json:"start"`
}

// GuideViewResponse Render-ready Guidebook tree for a guide: its rules grouped under the guide's categories, with every compound rule in one `_compound` category at the end.
type GuideViewResponse struct {
	// Categories The guide's categories in declared order, each with its checkable rules, followed by the `_compound` category that holds every rule with a `when` condition. A category with no rules is left out.
	Categories []struct {
		// ID The category id, or `_compound` for the compound rules.
		ID string `json:"id"`

		// Items The category's rules, in declared order.
		Items []GuideRuleView `json:"items"`

		// Label The category's display label.
		Label string `json:"label"`
	} `json:"categories"`

	// Description The guide's description. Empty when the guide declares none.
	Description string `json:"description"`

	// ID The guide id.
	ID string `json:"id"`

	// Marking The guide document's own overall marking, exactly as its `guide.marking` writes it. Display only: never interpreted. Absent when the document declares none.
	Marking *string `json:"marking,omitempty"`

	// Name The guide's display name.
	Name string `json:"name"`

	// Uncategorized Checkable rules that name no category, or a category the guide does not declare.
	Uncategorized []GuideRuleView `json:"uncategorized"`

	// Vocabulary Author-facing vocabulary: an open map from a term key (e.g. "rule", "concept", "tag") to the display words. Term keys are not enumerated here; consumers pick the keys that matter to them.
	//
	// Purely a presentation hint: nothing in the engine's evaluation logic reads it. Each consumer owns its own default map and calls `resolveVocabulary(source, defaults)` to merge.
	Vocabulary *Vocabulary `json:"vocabulary,omitempty"`
}

// GuideWhen A block of conditions over the rules the user selected. Every list present must hold: `any` when one of its nodes holds, `all` when each does, `none` when none does.
type GuideWhen struct {
	// All Every one of these holds.
	All *[]GuideWhenNode `json:"all,omitempty"`

	// Any At least one of these holds.
	Any *[]GuideWhenNode `json:"any,omitempty"`

	// None None of these holds.
	None *[]GuideWhenNode `json:"none,omitempty"`
}

// GuideWhenHas A guide condition that holds when the user selected one rule.
type GuideWhenHas struct {
	// Has The id of a selectable rule of this guide.
	Has string `json:"has"`
}

// GuideWhenNode One node of a guide `when` block: a condition, or a nested block.
type GuideWhenNode struct {
	union json.RawMessage
}

// GuideWhenOperator The tests a guide `when` condition can apply to the rules the user selected. `has_concept` and `not_has_concept` name a rule by id in `value`; the `count_*` tests (greater than, greater than or equal, equal, less than, less than or equal) compare the number of selected rules with the number in `value`; `has_content` holds when any rule is selected.
type GuideWhenOperator string

// GuideWhenTest A guide condition that tests the user's selection with an operator.
type GuideWhenTest struct {
	// Operator The tests a guide `when` condition can apply to the rules the user selected. `has_concept` and `not_has_concept` name a rule by id in `value`; the `count_*` tests (greater than, greater than or equal, equal, less than, less than or equal) compare the number of selected rules with the number in `value`; `has_content` holds when any rule is selected.
	Operator GuideWhenOperator `json:"operator"`

	// Value A rule id for `has_concept` and `not_has_concept`, a count for the `count_*` tests, absent for `has_content`.
	Value *GuideWhenTest_Value `json:"value,omitempty"`
}

// GuideWhenTestValue0 defines model for GuideWhenTest.Value.0.
type GuideWhenTestValue0 = string

// GuideWhenTestValue1 defines model for GuideWhenTest.Value.1.
type GuideWhenTestValue1 = float32

// GuideWhenTest_Value A rule id for `has_concept` and `not_has_concept`, a count for the `count_*` tests, absent for `has_content`.
type GuideWhenTest_Value struct {
	union json.RawMessage
}

// HealthResponse Response body for GET /v1/health
type HealthResponse struct {
	// AbiVersion Version of the wire interface this server speaks, fixed at `v1`. A client can compare it against the version it was built for before making calls.
	AbiVersion string `json:"abiVersion"`

	// EngineSHA Git commit SHA of the deployed build (from GIT_SHA env var, or empty string).
	EngineSHA string `json:"engineSha"`

	// EngineVersion Engine package version exposed to clients.
	EngineVersion string `json:"engineVersion"`

	// GitTag Git tag of the deployed build (from GIT_TAG env var, or empty string).
	GitTag string `json:"gitTag"`

	// OpenapiVersion OpenAPI spec version from info.version in openapi.yaml. SDKs compare this against their vendored spec.
	OpenapiVersion string `json:"openapiVersion"`

	// Status Always `"ok"` while the server is running.
	Status string `json:"status"`
}

// InputSpan One span of the input, by offset and text.
type InputSpan struct {
	// End End offset in the input string.
	End int `json:"end"`

	// Start Start offset in the input string.
	Start int `json:"start"`

	// Text The text at that span.
	Text string `json:"text"`
}

// LabelPresetView A named shortcut marking, name only: the marking it stands for stays server-side and is expanded by `applyPreset`.
type LabelPresetView struct {
	// Description Longer explanation of when to use the preset.
	Description *string `json:"description,omitempty"`

	// ID Preset id, which `applyPreset` takes.
	ID string `json:"id"`

	// Label Display label for the preset.
	Label string `json:"label"`

	// Scope Id of the scope the preset belongs to. A preset without one belongs to the scheme's `defaultScope`.
	Scope *string `json:"scope,omitempty"`
}

// ManualGuideSelection Validated manual guide fact bound to a document revision.
type ManualGuideSelection struct {
	// DocumentDigest Digest of the exact document revision reviewed.
	DocumentDigest string `json:"documentDigest"`

	// Evidence Evidence in the submitted document supporting the selection.
	Evidence []ContentEvidenceSpan `json:"evidence"`

	// ProposalOrigin Describes how a decision proposal entered the analyst workflow.
	ProposalOrigin *DecisionProposalOrigin `json:"proposalOrigin,omitempty"`

	// Reason Analyst-supplied reason for the selection.
	Reason string `json:"reason"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`

	// SelectionID Deterministic selection identity.
	SelectionID string `json:"selectionId"`
}

// ManualGuideSelectionInput Input for an analyst-selected rule that lexical scanning missed.
type ManualGuideSelectionInput struct {
	// DocumentDigest Digest of the exact document revision reviewed.
	DocumentDigest string `json:"documentDigest"`

	// Evidence Evidence in the submitted document supporting the selection.
	Evidence []ContentEvidenceSpan `json:"evidence"`

	// ProposalOrigin Describes how a decision proposal entered the analyst workflow.
	ProposalOrigin *DecisionProposalOrigin `json:"proposalOrigin,omitempty"`

	// Reason Analyst-supplied reason for the selection.
	Reason string `json:"reason"`

	// Rule Identifies the exact resolved semantics of one policy rule.
	Rule RuleReference `json:"rule"`
}

// MarkingColorRule Ordered scheme-default or host display-color override.
type MarkingColorRule struct {
	// Color Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
	//
	// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
	Color ConceptColor `json:"color"`

	// When Conditions that trigger a rule. Supports shorthand forms:
	//
	// - `string`: concept ID must be present (any category)
	// - `string[]`: any of these concept IDs must be present
	// - `object`: structured `{ any | all | none }` block, possibly nested
	When *RuleWhen `json:"when,omitempty"`
}

// MarkingValidation How an audit judged one authored marking, whole or part.
//
// `valid` and `inputWasValid` mean exactly what `ValidateResult` defines them to mean. An audit asks whether the marking was right as the author wrote it, and that is `inputWasValid`; `valid` says whether it stands once the engine has settled it, as it does everywhere else.
//
// It carries the marking's facts and how far recognition got with it, so a caller that rolls markings up, checks access or offers a correction works from the audit instead of reading each marking a second time.
type MarkingValidation struct {
	// AppliedFixes Every change the engine made while settling the marking, as `ValidateResult.appliedFixes`. Empty when it made none, so a false `inputWasValid` is always explained here or in `errors`.
	AppliedFixes []Change `json:"appliedFixes"`

	// Assertions The facts the marking states, settled and listed in print order as a validate result's `assertions` are, in the role the marking has in this document.
	Assertions []Assertion `json:"assertions"`

	// Errors What the scheme's rules find wrong with the settled marking.
	Errors []ValidationError `json:"errors"`

	// InputWasValid Whether the marking needed no change as written, as `ValidateResult.inputWasValid`.
	InputWasValid bool `json:"inputWasValid"`

	// ProposedFixes The ways to make the marking valid, each a whole marking the engine has read back and found valid, as `ProcessResult.proposedFixes`. An audit offers them because its caller is usually the author about to fix the document.
	ProposedFixes []Fix `json:"proposedFixes"`

	// Recognition What recognition made of an input string: how far it got, what it read, and what it could not account for.
	//
	// Every character of the marking region is accounted for: `complete` means recognized content, declared syntax, or permitted whitespace, with nothing left over. Finding one field is never enough to stop reading. An input that supports two distinct readings is `ambiguous` rather than read as either, because picking one by declaration order, category order, rank, or assertion count is how a guess becomes a confident answer.
	//
	// This is what a reader needs and what validate returns. Process returns `RecognitionWithCorrections`, which adds how to correct the text.
	Recognition RecognitionResult `json:"recognition"`

	// Rendered The settled marking, written as validate writes it when the request names no output: in the scheme's default format, or for a scheme that declares `render.surfaces`, each written surface in its default for the role the marking has in this document, by the rule `RenderConfigDoc.surfaces` states.
	Rendered string `json:"rendered"`

	// Valid Whether the marking stands once settled, as `ValidateResult.valid`.
	Valid bool `json:"valid"`
}

// MatchAllNode The `all` combinator: every child must match, and all children must overlap within `within` of the first child's span. The emitted span is the first child's (the anchor). Used to require several signals co-occur.
type MatchAllNode struct {
	// All The signals that must all match. The first is the anchor whose span is kept.
	All []MatchNode `json:"all"`

	// Label Display label for this node's matches. Defaults to `all`.
	Label *string `json:"label,omitempty"`

	// Within A bounded proximity window. Exactly one token or character bound is present, so a windowed operation can never widen to the whole line accidentally.
	Within Win `json:"within"`
}

// MatchAnyNode The `any` combinator: a union of child spans under one label. Used to express a multi-format value as several small, individually screened leaves instead of one giant alternation.
type MatchAnyNode struct {
	// Any The alternatives. A span any of them matches is a match.
	Any []MatchNode `json:"any"`

	// Label Display label for this node's matches. Defaults to `any`.
	Label *string `json:"label,omitempty"`
}

// MatchContextExpression Cue relationship used by the context combinator.
type MatchContextExpression struct {
	// Candidate A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	Candidate MatchNode `json:"candidate"`

	// Cue A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	Cue MatchNode `json:"cue"`

	// Side Relative position accepted by a context-expression cue.
	Side *MatchContextSide `json:"side,omitempty"`

	// Within A bounded proximity window. Exactly one token or character bound is present, so a windowed operation can never widen to the whole line accidentally.
	Within Win `json:"within"`
}

// MatchContextNode The `context` combinator: keep each `candidate` span only when a `cue` span lies within `within` on the given `side` (default `both`). Used to require a positive cue (e.g. a ZIP only counts with a postal/state cue nearby).
type MatchContextNode struct {
	// Context Cue relationship used by the context combinator.
	Context MatchContextExpression `json:"context"`

	// Label Display label for this node's matches. Defaults to `context`.
	Label *string `json:"label,omitempty"`
}

// MatchContextSide Relative position accepted by a context-expression cue.
type MatchContextSide string

// MatchNearEmit Span emitted by a near-expression match.
type MatchNearEmit string

// MatchNearExpression Pairing expression used by the near combinator.
type MatchNearExpression struct {
	// A A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	A MatchNode `json:"a"`

	// B A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	B MatchNode `json:"b"`

	// Emit Span emitted by a near-expression match.
	Emit *MatchNearEmit `json:"emit,omitempty"`

	// Within A bounded proximity window. Exactly one token or character bound is present, so a windowed operation can never widen to the whole line accidentally.
	Within Win `json:"within"`
}

// MatchNearNode The `near` combinator: pair an `a` span with a `b` span that lies within `within`. `emit` selects which span survives: the `a` span (default), the `b` span, or the envelope spanning both.
type MatchNearNode struct {
	// Label Display label for this node's matches. Defaults to `near`.
	Label *string `json:"label,omitempty"`

	// Near Pairing expression used by the near combinator.
	Near MatchNearExpression `json:"near"`
}

// MatchNode A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
type MatchNode struct {
	union json.RawMessage
}

// MatchNotNode The `not` combinator: keep each `of` candidate span only when NO `not` span lies within `within` of it. Drops a candidate that has a competing cue nearby.
type MatchNotNode struct {
	// Label Display label for this node's matches. Defaults to `not`.
	Label *string `json:"label,omitempty"`

	// Not A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	Not MatchNode `json:"not"`

	// Of A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	Of MatchNode `json:"of"`

	// Within A bounded proximity window. Exactly one token or character bound is present, so a windowed operation can never widen to the whole line accidentally.
	Within Win `json:"within"`
}

// MatchValidateNode The `validate` combinator: run the named validator over each `candidate` span's matched text and keep the candidate when the validator passes (`keep: 'pass'`, the default) or when it fails (`keep: 'fail'`). The validator is content-free arithmetic or structure registered in the validator registry; an unknown name is a load-time error.
type MatchValidateNode struct {
	// Label Display label for this node's matches. Defaults to `validate`.
	Label *string `json:"label,omitempty"`

	// Validate Validator invocation used by the validate combinator.
	Validate MatchValidationExpression `json:"validate"`
}

// MatchValidationExpression Validator invocation used by the validate combinator.
type MatchValidationExpression struct {
	// Args Optional opaque arguments passed through to the validator.
	Args interface{} `json:"args,omitempty"`

	// Candidate A declarative match expression: a `GuideMatchSpec` leaf, or one of the combinator nodes composed over child nodes. A guide rule's `match` array is a list of these.
	Candidate MatchNode `json:"candidate"`

	// Keep Validator outcome retained by a validation expression.
	Keep *MatchValidationKeep `json:"keep,omitempty"`

	// With Registered validator name.
	With string `json:"with"`
}

// MatchValidationKeep Validator outcome retained by a validation expression.
type MatchValidationKeep string

// Notice A notice that was resolved and collected for a specific concept in a processed marking. Notices originate from the concept's `documentation.notices` definitions and may be conditional (only fire when another concept is simultaneously asserted) or unconditional.
//
// Returned by `SchemeEngine.collectNotices()`, and carried on every result that describes a marking, so reading the notices for a processed marking does not take a second call. The `required` flag indicates whether the notice represents a mandatory acknowledgment the UI must surface before the user can submit the marking.
type Notice struct {
	// ConceptID The concept the notice is attached to.
	ConceptID string `json:"conceptId"`

	// ConceptName The concept's display name, for showing beside the notice.
	ConceptName string `json:"conceptName"`

	// Required Whether the notice must be shown. A host surfaces a required notice before the marking is accepted.
	Required bool `json:"required"`

	// Text The notice text the marking requires.
	Text string `json:"text"`

	// Type What kind of notice a concept carries.
	Type ConceptNoticeType `json:"type"`

	// Views The value a notice is registered under, keyed by view name, e.g. `{ ism: 'RD' }`.
	//
	// Structurally the same map as `ViewLabels` and deliberately not the same type. A view label is how a view spells something for a reader; this is the token a view's own vocabulary files the notice under, and one is never substitutable for the other.
	Views *NoticeViewValues `json:"views,omitempty"`
}

// NoticeViewValues The value a notice is registered under, keyed by view name, e.g. `{ ism: 'RD' }`.
//
// Structurally the same map as `ViewLabels` and deliberately not the same type. A view label is how a view spells something for a reader; this is the token a view's own vocabulary files the notice under, and one is never substitutable for the other.
type NoticeViewValues map[string]string

// ProcessResult What process answers: everything validate answers, plus what an author needs to correct and finish the marking.
//
// Process is the author's call. A composer, a picker or an editor that helps someone write a marking calls it, often on every keystroke, and can send what was typed together with the facts already chosen. Beside the validate answer it carries:
//
// - the corrections that turn the text into canonical text, inside `recognition`;
// - what was read before repair, when repair changed it;
// - the fix paths, each a whole valid marking the engine has proven;
// - the outputs the request named in `outputs`, and only those;
// - the options for the categories the request named in `suggestFor`, only when it named some.
//
// The step-by-step trace of how the engine got here is not part of it. That is what explain answers, from the same request.
type ProcessResult struct {
	// AppliedFixes Every change the engine actually made while settling the marking: a departure from canonical spelling or order that recognition read through, a repair, a fact a rule required, a handed-in fact settling removed, or text read in the other role from the one the request names. A removed fact is named with the reason it was removed, including a fact handed in with a derived source label that the run does not derive again. A requirement that conflicts with an existing concept is not added and appears in errors and proposed fixes, with no applied change. Repair never removes an author-stated concept. These are the reasons a marking can be `valid` while `inputWasValid` is false.
	//
	// A departure from spelling or order is reported here without a position. On a process result, `recognition.corrections` reports the same departure with the span of the text it covers. They are one change seen two ways, so a display shows one of them, not both. The role departure, `_recognition_role`, is about the whole text: it has no span and no entry in `recognition.corrections`, so it is reported only here, and a display that shows corrections in place of this list still shows it.
	AppliedFixes []Change `json:"appliedFixes"`

	// Assertions The facts the marking states, as the engine settled them.
	//
	// Listed in print order: the order the rendered marking places each concept; a concept the rendering leaves out appears where it would print. A concept the rendering has no place for at all follows the others of its category, or comes last when its category has none, in the order the scheme declares.
	Assertions []Assertion `json:"assertions"`

	// Assessment What a render did to the state it was given, bound to that state and that format.
	//
	// Answered against the state actually rendered: what was emitted, what was left out, what collapsed into an equivalent, and what could not be written at all.
	//
	// A caller that needs full-state preservation reads this. A caller that wants a human projection can ignore it and still get a useful string.
	Assessment RenderAssessment `json:"assessment"`

	// DisplayColor Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
	//
	// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
	DisplayColor *ConceptColor `json:"displayColor,omitempty"`

	// Errors What the scheme's rules find wrong with the settled marking. Any entry makes `valid` false.
	Errors []ValidationError `json:"errors"`

	// InputWasValid True when the text or facts as given needed no change to stand.
	//
	// Defined here once, and judged after the same settling as `valid`. A marking can be `valid` without having been valid as given: the engine settled it into one that stands. A result whose `inputWasValid` is false always says why: each change the engine made is in `appliedFixes`, and whatever it could not settle is in `errors`.
	//
	// Right as written is judged against the view the text was read as, the one `recognition.inputFormat` reports, or for text read on more than one surface, the view each surface's part was read as. A fact settling adds, including a restored tier parent, counts against it only when that view writes the fact. A view that leaves the fact out writes the same text with the fact as without it, so the text was right already. Text read as no single view counts the fact when any view that read it writes it. Facts handed in with no text are judged against what a request that names no output writes them in, by the rule `rendered` follows.
	//
	// Text is also not right as written when the request's context names a role and recognition decided the text is a view of the other role, as `recognition.role` reports. The result carries that departure in `appliedFixes` under `_recognition_role`, and `valid` is judged as it would be without it. A request that names its input format has said which view the text is: the text is judged as that view, and no role departure is raised. On a scan region and on a marking a document audit judges, the role the marking's place gives it stands for the context's role: a region found inside a wrapper is a part, and an audit judges each marking in the role it has in the document.
	InputWasValid bool `json:"inputWasValid"`

	// Notices The notices the scheme attaches to the concepts in `assertions`, collected the same way `SchemeEngine.collectNotices()` collects them. Always present, empty when the marking carries none, so a consumer can read the field without checking that it exists.
	Notices []Notice `json:"notices"`

	// ProposedFixes The ways to make the resulting value valid, computed from that value after applied changes, each a whole marking the engine has read back and found valid before offering it. Empty when the marking is valid, or when no change the engine could prove would make it so.
	//
	// A request's `parts` go back in with each fix, as the caller will send them, so they are read again in the fix's whole and the fix is judged with them: every part that read as a valid part marking in the call still does, and the fix covers every part that does. A part that did not read as a valid part marking in the call is one no fix to the whole can mend, so it does not keep a fix from being offered, and it may still fail beside the fix.
	//
	// Proofs spend what is left of the call's work budget. When a proof runs out of it, or the caller cancels the call, the paths not yet proven are not offered, and `warnings` carries a notice whose `ruleId` is `_fixes_unproven`, so a short list or an empty one is not taken for every fix there is.
	ProposedFixes []Fix `json:"proposedFixes"`

	// Recognition What recognition made of an input string, with the corrections that turn it into canonical text.
	//
	// The corrections are the edit script recognition read through: each names a span of the input, what it said, and the canonical text it was read as. An author's editor uses them to show what changed and where. A reader has no use for them, which is why only process returns this form.
	Recognition *RecognitionWithCorrections `json:"recognition,omitempty"`

	// Rendered The settled marking, in one string.
	//
	// It is written in the output the request names. With none named, a scheme without `render.surfaces` writes its default format, and a scheme with them writes each surface in its default for the marking's role, by the rule `RenderConfigDoc.surfaces` states. Every operation that returns this shape follows that one rule.
	//
	// For a scheme that declares `render.surfaces`, the string is the output resolved on each surface the result writes, as `RenderConfigDoc.surfaces` describes, in declared surface order and joined by `render.surface_separator`. An output that writes nothing adds nothing, and outputs written as XML attributes in one namespace merge into one attribute list. One refused output refuses the whole string, and `assessment` says which one and why.
	Rendered string `json:"rendered"`

	// Renders Renders of one state, one per output, keyed by output address.
	//
	// For a scheme without `render.surfaces` the outputs are its formats, keyed by format id. For a scheme with them they are the views of its surfaces, keyed `surface/view` with each id percent-encoded, surfaces in declared order and each surface's views in theirs. Every view a surface's default names is one of its views, so a default has no entry of its own, and a format no view writes has none either.
	//
	// On a process result it holds exactly the outputs the request named in `outputs`. An output that was asked for and could not write the state says so in its assessment rather than being left out, because a missing key would mean nobody asked for it.
	Renders RenderResultsByFormat `json:"renders"`

	// SuggestionsFor Suggestions for several categories, keyed by category id.
	//
	// Each entry holds the options for its own category and no other, answered as if that category's own assertions were not yet chosen, so a picker can offer the alternatives to its current selection, not only what can still be added beside it. A label preset belongs to no category, so no entry offers one. `suggestForCategories` returns this, and so does a process call that names categories in `suggestFor`.
	SuggestionsFor *SuggestResultsByCategory `json:"suggestionsFor,omitempty"`

	// Valid True when the marking stands once the engine has settled it: read, repaired, completed by what the scheme's rules imply, and judged with no error left.
	//
	// Defined here once, and every result that reports `valid` means this, whichever operation produced it and whichever door it came through. Reading is tolerant on purpose, so the engine settles what it can. Whether the text was right as written is what `inputWasValid` and `appliedFixes` answer; no request asks the engine to settle less, so validate and process give the same `valid` for the same input, context, surface and parts.
	Valid bool `json:"valid"`

	// Warnings Advisories that do not affect `valid`, such as a deprecated concept. On a process result whose fix paths were not all proven, this carries a notice whose `ruleId` is `_fixes_unproven`, as `ProcessResult.proposedFixes` describes.
	Warnings []Warning `json:"warnings"`
}

// ProcessingContext Caller-supplied context naming the role the marking being read plays. No context, or `role: 'whole'`, means a whole: a marking that stands for everything it covers. `role: 'part'` means a part that belongs to a whole, named in `partOf` when it is known.
//
// On a rollup the context names the contributors' role, never the rollup's: the rollup is always a whole. So no context rolls wholes up into an aggregate, and `{ role: 'part', partOf }` rolls a whole's parts up. See `SchemeEngineAdapter.aggregate`.
type ProcessingContext struct {
	// PartOf The whole marking this is part of (when role='part').
	//
	// The whole does two jobs. In a role-aware scheme it is the reading context for the scheme's whole-aware rules: a rule condition with `target: 'context'` is answered from it. And in every scheme, role-aware or not, `process`, `validate` and `explain` compare a part that is valid on its own with its whole. That is the comparison a whole is judged by against its parts and the one `checkCoverage` gives, asked of this one part: category by category, in the direction each category's rollup declares. A category a rollup pipeline resolves is judged instead by what rolling the whole up with this part would add or drop. A category whose rollup has no direction one part can be judged by, such as `first`, is left to the whole.
	//
	// What the part carries beyond its whole is one `_exceeds_whole` error on the part. Its `affectedConcepts` names every finding: first what the whole does not cover, then what the whole carries that the part does not allow. Its `categoryId` is set when every finding is in one category. It offers no fix, because whether the part or the whole should change is the author's call. A part that breaks its own rules reports those errors and is not compared. Its own rules are the scheme's: an in-process engine's `afterValidate` plugin runs after the comparison and does not change whether it happens. A call's `parts` and a document scan or audit read each part in the whole without this comparison; they judge coverage on the whole instead.
	//
	// A part is compared with its whole as the engine settles the whole: read, modernized, completed by what the scheme's rules imply and judged, the same settling `valid` reports for the whole on its own. A whole spelled with a retired concept is compared as its replacement, and a fact the whole's rules add is one the part is compared with.
	//
	// The same holds wherever a part meets its whole: here; for each of a call's `parts`, however the call's whole came, as text or as facts; and for each input `SchemeEngineAdapter.aggregate` reads against a whole. A part that takes a value from its whole, such as a list it stands for, takes it from that settled whole.
	//
	// What the part gets depends on the state of its whole:
	//
	// - A whole that reads and is valid is verified. The part's rules read it, and `process`, `validate` and `explain` compare the part with it.
	// - A whole that reads but is invalid, or does not read, is not verified. The part is still read and judged by its own rules but is not compared with the whole, so it never carries `_exceeds_whole`. It always carries the `_whole_unverified` notice.
	// - With no whole given, the part is judged by its own rules alone. It carries the `_whole_unverified` notice only when one of its rules was left unresolved.
	//
	// Without a verified whole, what needs the whole is left open instead of guessed: a rule condition with `target: 'context'` is unresolved, so its rule never fires, as `RuleConditionTarget` describes. A stand-in is the exception, because it is shorthand for the whole's values and means nothing without them: a part that leaves a category to a concept that declares `stands_for_whole` carries a `_stand_in_requires_whole` error naming it.
	PartOf *struct {
		// Assertions Pre-resolved assertions of the whole marking (populated by the engine)
		Assertions *[]Assertion `json:"assertions,omitempty"`

		// Rendered Rendered string of the whole marking
		Rendered string `json:"rendered"`
	} `json:"partOf,omitempty"`

	// Role Structural role of a marking within a processing operation.
	Role *ProcessingRole `json:"role,omitempty"`
}

// ProcessingRole Structural role of a marking within a processing operation.
type ProcessingRole string

// RateLimitError Response body for 429 Too Many Requests.
type RateLimitError struct {
	// Error Always "Too Many Requests".
	Error string `json:"error"`

	// RetryAfter Seconds until the rate-limit window resets.
	RetryAfter float32 `json:"retryAfter"`
}

// ReadinessWarning Readiness warning codes a runtime replica reports on its readiness probe.
//
// A warning describes a degraded condition an operator should act on. It never changes readiness: a replica that is ready stays ready and keeps serving while it carries one. The set is closed so a probe consumer can alert on each code by name.
//
// `audit-degraded`: the runtime could not write one or more audit records and keeps answering requests without them until audit writes succeed again.
//
// `application-allowance-exceeded`: more distinct application ids have called the runtime than its license's application allowance. Every application is still served, registered or not: the warning informs and never refuses a request.
//
// `generation-refused`: a managed runtime refused what its control plane last offered, either the newest generation or the desired-state page as a whole, and keeps the generation it already holds, if it holds one.
//
// `desired-state-stale`: a managed runtime has no current signed statement from its control plane that the generation it serves is still the desired one. The control plane may be unreachable, or newer generations may be withheld on the way to the runtime. It keeps serving the generation it holds.
type ReadinessWarning string

// ReadyResponse Response body for GET /v1/ready
type ReadyResponse struct {
	// Status Replica readiness state reported by GET /v1/ready. `ready` once the replica may receive traffic after RuntimeHost initialization and required-artifact checks. License expiration never takes a mission runtime out of service.
	Status ReadyStatus `json:"status"`

	// Warnings Degraded conditions an operator should act on, in a stable order. Absent when there are none. A warning never changes `status`: a ready replica that carries one stays ready. Present only for an authenticated caller.
	Warnings *[]ReadinessWarning `json:"warnings,omitempty"`
}

// ReadyStatus Replica readiness state reported by GET /v1/ready. `ready` once the replica may receive traffic after RuntimeHost initialization and required-artifact checks. License expiration never takes a mission runtime out of service.
type ReadyStatus string

// RecognitionCorrection One span recognition rewrote, with what it was and what it became.
type RecognitionCorrection struct {
	// End End offset in the input string.
	End int `json:"end"`

	// Original The text as supplied.
	Original string `json:"original"`

	// Reason Why recognition rewrote one span of the input on the way to a reading.
	//
	// `alias` covers a concept read in a spelling the reading form does not write, the spelling another of the concept's views gives among them, and a prefix read in one of its `prefix_aliases`. The replacement is the spelling the form writes.
	Reason RecognitionCorrectionReason `json:"reason"`

	// Replacement The canonical text it was read as.
	Replacement string `json:"replacement"`

	// Start Start offset in the input string.
	Start int `json:"start"`
}

// RecognitionCorrectionReason Why recognition rewrote one span of the input on the way to a reading.
//
// `alias` covers a concept read in a spelling the reading form does not write, the spelling another of the concept's views gives among them, and a prefix read in one of its `prefix_aliases`. The replacement is the spelling the form writes.
type RecognitionCorrectionReason string

// RecognitionLimit The ceilings recognition can reach.
//
// Closed, because a caller telling one exhausted ceiling from another is writing a branch, not a message. `input-length` is the longest input the engine reads, `transitions` and `lookups` are the work it spends reading, and `live-states` is how many readings it keeps open at once.
type RecognitionLimit string

// RecognitionResult What recognition made of an input string: how far it got, what it read, and what it could not account for.
//
// Every character of the marking region is accounted for: `complete` means recognized content, declared syntax, or permitted whitespace, with nothing left over. Finding one field is never enough to stop reading. An input that supports two distinct readings is `ambiguous` rather than read as either, because picking one by declaration order, category order, rank, or assertion count is how a guess becomes a confident answer.
//
// This is what a reader needs and what validate returns. Process returns `RecognitionWithCorrections`, which adds how to correct the text.
type RecognitionResult struct {
	// InputFormat What the text was read as, when that is decided.
	//
	// For a scheme without `render.surfaces`, the format recognized, when exactly one was. For a scheme with them, the output address, `surface/view`, of the view the text was read as, decided as `RenderConfigDoc.surfaces` states: the same key that view's output has in `renders`. The form it happened to be read in is not reported. Absent when no one view is decided, as for text that carries the categories of more than one surface.
	InputFormat *string `json:"inputFormat,omitempty"`

	// Role Structural role of a marking within a processing operation.
	Role *ProcessingRole `json:"role,omitempty"`

	// Status How far recognition got.
	Status RecognitionStatus `json:"status"`

	// Surface The surface the input was read as, by scheme scope id.
	//
	// A caller that named a surface gets it back. A caller that named none gets the surface the recognized format declares, or with `render.surfaces` the surface whose categories the text carries, which tells it which part of a marking the text turned out to be. Absent for text that carries the categories of more than one surface, unless the caller named one.
	Surface *string `json:"surface,omitempty"`

	// Unmatched Input not accounted for as content, syntax, or permitted whitespace.
	Unmatched []InputSpan `json:"unmatched"`
}

// RecognitionStatus How far recognition got.
type RecognitionStatus string

// RecognitionWithCorrections What recognition made of an input string, with the corrections that turn it into canonical text.
//
// The corrections are the edit script recognition read through: each names a span of the input, what it said, and the canonical text it was read as. An author's editor uses them to show what changed and where. A reader has no use for them, which is why only process returns this form.
type RecognitionWithCorrections struct {
	// Corrections Departures from canonical that recognition read through.
	Corrections []RecognitionCorrection `json:"corrections"`

	// InputFormat What the text was read as, when that is decided.
	//
	// For a scheme without `render.surfaces`, the format recognized, when exactly one was. For a scheme with them, the output address, `surface/view`, of the view the text was read as, decided as `RenderConfigDoc.surfaces` states: the same key that view's output has in `renders`. The form it happened to be read in is not reported. Absent when no one view is decided, as for text that carries the categories of more than one surface.
	InputFormat *string `json:"inputFormat,omitempty"`

	// Role Structural role of a marking within a processing operation.
	Role *ProcessingRole `json:"role,omitempty"`

	// Status How far recognition got.
	Status RecognitionStatus `json:"status"`

	// Surface The surface the input was read as, by scheme scope id.
	//
	// A caller that named a surface gets it back. A caller that named none gets the surface the recognized format declares, or with `render.surfaces` the surface whose categories the text carries, which tells it which part of a marking the text turned out to be. Absent for text that carries the categories of more than one surface, unless the caller named one.
	Surface *string `json:"surface,omitempty"`

	// Unmatched Input not accounted for as content, syntax, or permitted whitespace.
	Unmatched []InputSpan `json:"unmatched"`
}

// ReferenceSurfaceLabel Resolved display label for an artifact's reference-surface chrome noun: the name its own reference component wears (a scheme's Handbook, a guide's Guidebook, an access policy's Access Handbook). Always fully populated: an artifact that declares no override resolves to the engine's generic noun for that surface, using the same singular/plural `Term` contract as the artifact vocabulary. Kept as its own type because this is an already-resolved reference-surface value rather than the complete author-facing vocabulary.
type ReferenceSurfaceLabel struct {
	// Plural The noun for more than one, such as `Handbooks`.
	Plural string `json:"plural"`

	// Singular The noun for one, such as `Handbook`.
	Singular string `json:"singular"`
}

// RenderAssessment What a render did to the state it was given, bound to that state and that format.
//
// Answered against the state actually rendered: what was emitted, what was left out, what collapsed into an equivalent, and what could not be written at all.
//
// A caller that needs full-state preservation reads this. A caller that wants a human projection can ignore it and still get a useful string.
type RenderAssessment struct {
	// Collapsed Groups of facts that collapsed into one written equivalent.
	Collapsed [][]AssertionRef `json:"collapsed"`

	// Conflicts Fields that two assertions both claim.
	Conflicts []RenderFieldConflict `json:"conflicts"`

	// Emitted Facts this render wrote.
	Emitted []AssertionRef `json:"emitted"`

	// Omitted Facts this render did not write.
	//
	// Some of them come back when the written text is read in the same context, because the scheme derives them from what was written: a fact a rule adds, a category's default, a value a part takes from its whole. The rest are lost to the string, such as a fact of a scope or a surface this render does not write, or one a field that writes only its own view's spellings has no spelling for (`view_only`). Either way the string is a projection of the facts, and a `'full'` render refuses any omission.
	Omitted []AssertionRef `json:"omitted"`

	// Unrepresentable Facts that cannot be written in this format at all, and outputs refused whole.
	Unrepresentable []UnrepresentableValue `json:"unrepresentable"`
}

// RenderFieldConflict Two fields of one format that both claim the same assertion.
type RenderFieldConflict struct {
	// Assertions The assertions in conflict.
	Assertions []AssertionRef `json:"assertions"`

	// Field The field, by `name` or category id.
	Field string `json:"field"`
}

// RenderFormatMode How a format runs the fields it writes together.
type RenderFormatMode string

// RenderPreservation What a caller is asking a render for.
//
// `'projection'` is the ordinary human projection: the format writes the fields it declares and the assessment reports the rest. `'full'` is section 7.4's full-preservation request, which fails rather than returning a string when facts were omitted, conflicting, or unrepresentable. Naming the request is what makes the failure possible: a render that only ever returns a string has nowhere to say that the string is not the state.
type RenderPreservation string

// RenderResponse Response body for Context rendering.
type RenderResponse struct {
	// Assessment What a render did to the state it was given, bound to that state and that format.
	//
	// Answered against the state actually rendered: what was emitted, what was left out, what collapsed into an equivalent, and what could not be written at all.
	//
	// A caller that needs full-state preservation reads this. A caller that wants a human projection can ignore it and still get a useful string.
	Assessment RenderAssessment `json:"assessment"`

	// Result The rendered marking string.
	Result string `json:"result"`
}

// RenderResult The output of a render: the string, and what producing it did to the state.
//
// The assessment travels with the string so that a convenience caller cannot accidentally treat a lossy projection as a transport. It is bound to this render of this state, not to the format in the abstract.
//
// For a scheme that declares `render.surfaces`, one result can be composed from the outputs of every surface it writes, and this is what such a composed result holds. The string is composed exactly as `ProcessResult.rendered` is. The assessment combines the surfaces written: it lists a fact as omitted only when no written surface wrote it, and an unrepresentable or conflicting claim from any surface stands. Such a claim is never cancelled because another surface wrote something.
type RenderResult struct {
	// Assessment What a render did to the state it was given, bound to that state and that format.
	//
	// Answered against the state actually rendered: what was emitted, what was left out, what collapsed into an equivalent, and what could not be written at all.
	//
	// A caller that needs full-state preservation reads this. A caller that wants a human projection can ignore it and still get a useful string.
	Assessment RenderAssessment `json:"assessment"`

	// Rendered The rendered marking.
	Rendered string `json:"rendered"`
}

// RenderResultsByFormat Renders of one state, one per output, keyed by output address.
//
// For a scheme without `render.surfaces` the outputs are its formats, keyed by format id. For a scheme with them they are the views of its surfaces, keyed `surface/view` with each id percent-encoded, surfaces in declared order and each surface's views in theirs. Every view a surface's default names is one of its views, so a default has no entry of its own, and a format no view writes has none either.
//
// On a process result it holds exactly the outputs the request named in `outputs`. An output that was asked for and could not write the state says so in its assessment rather than being left out, because a missing key would mean nobody asked for it.
type RenderResultsByFormat map[string]RenderResult

// Rollup Canonical rollup operator. `pipeline` requires a matching entry in the scheme's `rollup_pipelines`. Most merge operators are set-algebraic over each input's assertions; rank-based operators (max/min) require categories whose concepts declare `rank`.
type Rollup string

// RuleCondition A single condition in a rule's `when` block. Can test category contents, processing context role, or concept presence.
type RuleCondition struct {
	// Category Category to evaluate against. Also serves as a shorthand matching any concept in this category within when blocks.
	Category *string `json:"category,omitempty"`

	// Has Shorthand for has_concept: check if this concept is present
	Has *string `json:"has,omitempty"`

	// Operator The test a scheme condition applies to one category's assertions. `has_concept` and `not_has_concept` name a concept in `value`; `has_content` holds when the category holds anything; the `count_*` tests (greater than, greater than or equal, equal, less than, less than or equal) compare the number of concepts in the category with the number in `value`; `min_rank` holds when the category holds a concept ranked at or above the rank in `value`; `subset_of` holds when this category's concepts are a subset of the category named in `value`, and `shares_concept_with` when the two share a concept; `every_has_concept` and `any_lacks_concept` ask whether every assertion set, or at least one, holds the concept named in `value`.
	Operator *RuleConditionOperator `json:"operator,omitempty"`

	// Role Structural role of a marking within a processing operation.
	Role *ProcessingRole `json:"role,omitempty"`

	// Target Assertion set a rule condition is evaluated against.
	//
	// - `'self'` (default): the assertions currently under evaluation.
	// - `'context'`: the assertions of the whole marking this one is part of. Such a condition has an answer only when the part has a verified whole, one that reads and is valid. Without one (no whole given, a whole that does not read, or a whole that reads but is invalid) the condition is unresolved, and so is any `all`, `any` or `none` around it whose answer depends on it. A rule whose `when` is unresolved never fires, changes no value, reports no error and offers no fix; a check rule (`require`, `require_any`, `restrict`, `forbid`, `exclusive`, `cardinality`) that the marking already meets still passes, and any other unresolved rule gives the part the `_whole_unverified` notice. A part's stand-in for its whole's values is not left open this way: without a verified whole it is a `_stand_in_requires_whole` error. See `ProcessingContext.partOf`.
	// - `'contributors'`: each contributing assertion set that went into an aggregate, evaluated separately. The condition holds only when EVERY contributor satisfies it, and it is never satisfied outside aggregation, where there are no contributors to inspect. Use it to describe what must be true of every part of a document, which the merged set can no longer answer once the parts have been collapsed into one.
	Target *RuleConditionTarget `json:"target,omitempty"`

	// Value Explicit value for the operator
	Value *RuleCondition_Value `json:"value,omitempty"`
}

// RuleConditionValue0 defines model for RuleCondition.Value.0.
type RuleConditionValue0 = string

// RuleConditionValue1 defines model for RuleCondition.Value.1.
type RuleConditionValue1 = float32

// RuleCondition_Value Explicit value for the operator
type RuleCondition_Value struct {
	union json.RawMessage
}

// RuleConditionOperator The test a scheme condition applies to one category's assertions. `has_concept` and `not_has_concept` name a concept in `value`; `has_content` holds when the category holds anything; the `count_*` tests (greater than, greater than or equal, equal, less than, less than or equal) compare the number of concepts in the category with the number in `value`; `min_rank` holds when the category holds a concept ranked at or above the rank in `value`; `subset_of` holds when this category's concepts are a subset of the category named in `value`, and `shares_concept_with` when the two share a concept; `every_has_concept` and `any_lacks_concept` ask whether every assertion set, or at least one, holds the concept named in `value`.
type RuleConditionOperator string

// RuleConditionTarget Assertion set a rule condition is evaluated against.
//
// - `'self'` (default): the assertions currently under evaluation.
// - `'context'`: the assertions of the whole marking this one is part of. Such a condition has an answer only when the part has a verified whole, one that reads and is valid. Without one (no whole given, a whole that does not read, or a whole that reads but is invalid) the condition is unresolved, and so is any `all`, `any` or `none` around it whose answer depends on it. A rule whose `when` is unresolved never fires, changes no value, reports no error and offers no fix; a check rule (`require`, `require_any`, `restrict`, `forbid`, `exclusive`, `cardinality`) that the marking already meets still passes, and any other unresolved rule gives the part the `_whole_unverified` notice. A part's stand-in for its whole's values is not left open this way: without a verified whole it is a `_stand_in_requires_whole` error. See `ProcessingContext.partOf`.
// - `'contributors'`: each contributing assertion set that went into an aggregate, evaluated separately. The condition holds only when EVERY contributor satisfies it, and it is never satisfied outside aggregation, where there are no contributors to inspect. Use it to describe what must be true of every part of a document, which the merged set can no longer answer once the parts have been collapsed into one.
type RuleConditionTarget string

// RuleContributorQuantifier Quantifies a condition block over contributing assertion sets.
type RuleContributorQuantifier string

// RuleReference Identifies the exact resolved semantics of one policy rule.
type RuleReference struct {
	// Artifact Identifies an immutable artifact participating in a composed operation.
	Artifact ArtifactReference `json:"artifact"`

	// RuleDigest RFC 8785/SHA-256 digest of the canonical resolved rule content.
	RuleDigest string `json:"ruleDigest"`

	// RuleID Rule identifier within the owning artifact.
	RuleID string `json:"ruleId"`
}

// RuleWhen Conditions that trigger a rule. Supports shorthand forms:
//
// - `string`: concept ID must be present (any category)
// - `string[]`: any of these concept IDs must be present
// - `object`: structured `{ any | all | none }` block, possibly nested
type RuleWhen struct {
	union json.RawMessage
}

// RuleWhen0 defines model for RuleWhen.0.
type RuleWhen0 = string

// RuleWhen1 defines model for RuleWhen.1.
type RuleWhen1 = []string

// RuleWhenNode A leaf condition or recursively nested boolean quantifier in a rule tree.
type RuleWhenNode struct {
	union json.RawMessage
}

// RuleWhenQuantifier A node in a rule `when`-tree: either a leaf `RuleCondition` or a nested quantifier (`any` / `all` / `none`) over further nodes.
//
// The tree structure allows arbitrarily deep boolean logic within a rule's trigger condition. Leaf nodes test assertion or context state; inner nodes combine them with OR (`any`), AND (`all`), or NOT-OR (`none`).
type RuleWhenQuantifier struct {
	// All Satisfied when every one of these nodes is.
	All *[]RuleWhenNode `json:"all,omitempty"`

	// Any Satisfied when at least one of these nodes is.
	Any *[]RuleWhenNode `json:"any,omitempty"`

	// None Satisfied when none of these nodes is.
	None *[]RuleWhenNode `json:"none,omitempty"`

	// Quantifier Quantifies a condition block over contributing assertion sets.
	Quantifier *RuleContributorQuantifier `json:"quantifier,omitempty"`

	// Target Evaluate this entire boolean block within each contributor.
	Target *string `json:"target,omitempty"`
}

// ScanLimits How far a scan got, and what stopped it.
//
// A scan spends one work budget over a whole document: finding the regions a marking could occupy costs against it, and so does reading each one. A document large enough exhausts it, and a document that offers more regions than the scanner will hold is capped. Either way the answer is a shorter list of markings than the document holds, and a shorter list is indistinguishable from a clean one unless the scanner says so. For a scanner whose answer decides how a document is handled, that is the difference between "this document carries no marking" and "this document was not read to the end", so it is reported rather than inferred.
//
// `complete` means the whole document was offered to recognition. `resource-limit` means the budget was spent or the caller aborted, and `limitHit` names which. `truncated` means the region cap was reached with budget to spare. `regionsFound` is how many candidate regions the scan considered, which is the number that makes a zero result legible.
type ScanLimits struct {
	// LimitHit The ceilings recognition can reach.
	//
	// Closed, because a caller telling one exhausted ceiling from another is writing a branch, not a message. `input-length` is the longest input the engine reads, `transitions` and `lookups` are the work it spends reading, and `live-states` is how many readings it keeps open at once.
	LimitHit *RecognitionLimit `json:"limitHit,omitempty"`

	// RegionsCapped Whether the per-document region cap was reached.
	RegionsCapped *bool `json:"regionsCapped,omitempty"`

	// RegionsFound Candidate regions the scan considered.
	RegionsFound int `json:"regionsFound"`

	// Status How far a scan got.
	//
	// `complete` means the whole document was offered to recognition. `truncated` means the region cap was reached with budget to spare. `resource-limit` means the budget was spent or the caller aborted.
	Status ScanStatus `json:"status"`
}

// ScanResponseBody What a scan answers on the wire.
//
// An object rather than a bare array, because an array alone cannot tell a document that held no marking from a scan that stopped before the end of it. Both come back as zero entries, and the two mean opposite things to whoever decides what happens to the document next.
type ScanResponseBody struct {
	// Limits How far a scan got, and what stopped it.
	//
	// A scan spends one work budget over a whole document: finding the regions a marking could occupy costs against it, and so does reading each one. A document large enough exhausts it, and a document that offers more regions than the scanner will hold is capped. Either way the answer is a shorter list of markings than the document holds, and a shorter list is indistinguishable from a clean one unless the scanner says so. For a scanner whose answer decides how a document is handled, that is the difference between "this document carries no marking" and "this document was not read to the end", so it is reported rather than inferred.
	//
	// `complete` means the whole document was offered to recognition. `resource-limit` means the budget was spent or the caller aborted, and `limitHit` names which. `truncated` means the region cap was reached with budget to spare. `regionsFound` is how many candidate regions the scan considered, which is the number that makes a zero result legible.
	Limits ScanLimits `json:"limits"`

	// Results The markings the scan found, in document order.
	Results []ScanResult `json:"results"`

	// Unread The wrapped marks the scan found but could not read, in document order. Empty when every wrapped mark it reached was read. Text past where the scan stopped was never offered to recognition, so none of it is listed.
	Unread []UnreadMark `json:"unread"`
}

// ScanResult A single marking found during a scheme scan of free text.
//
// One entry of `ScanOutcome.results` per detected marking occurrence in the document. `start`/`end` are offsets in UTF-16 code units in the input text, suitable for overlaying highlights in a rich-text editor. `result` carries what validate answers for the marking, so its facts and validity are available without a second call. `wrapped` is `true` when the marking was found inside an `embeddedMarkingStyle` wrapper pattern (e.g., a marking in parentheses).
type ScanResult struct {
	// End End position in the (possibly stripped) input text
	End int `json:"end"`

	// Inner The marking text inside the wrapper: the capture group of the scheme's `embeddedMarking` pattern. Present only when `wrapped` is true. Consumers that need to re-process the marking must use this rather than stripping delimiters from `marking`, because the wrapper shape is scheme-declared and is not always a pair of parentheses.
	Inner *string `json:"inner,omitempty"`

	// Marking The matched marking string (as found in the document)
	Marking string `json:"marking"`

	// Result What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
	//
	// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
	//
	// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
	//
	// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
	Result ValidateResult `json:"result"`

	// Role Structural role of a marking within a processing operation.
	Role ProcessingRole `json:"role"`

	// Start Start position in the (possibly stripped) input text
	Start int `json:"start"`

	// Type The output this marking is written in when the scan names no output format: `result.rendered` is written in it.
	//
	// Without `render.surfaces`, the format id that read it, or when no single format read it, the scheme's default output, the `defaultRenderFormat` its metadata reports. With them, the view it was read as, the address `result.recognition.inputFormat` reports, when the form it was read in is one that view writes: its `format` or a case's format. A marking read in one form that carries more than one surface, such as an attribute list, is read as no one view; when the views that name that form all write it under one view id, it is written in them, and its type is that bare view id, the output format that writes each surface in its view by that id. Any other marking, one read in a form its view is only read in or read as no one view, is written in the default for its `role`, the `defaultWholeRenderFormat` or `defaultPartRenderFormat` its metadata reports, and that default is its type.
	Type string `json:"type"`

	// Wrapped Whether the marking was found inside a wrapper pattern (e.g., parentheses)
	Wrapped bool `json:"wrapped"`
}

// ScanStatus How far a scan got.
//
// `complete` means the whole document was offered to recognition. `truncated` means the region cap was reached with budget to spare. `resource-limit` means the budget was spent or the caller aborted.
type ScanStatus string

// SchemeScope One scope a scheme declares, as the scopes resource lists it.
type SchemeScope struct {
	// ID Scope identifier, referenced by `scopes` on concepts and categories.
	ID string `json:"id"`

	// Label Display label for the scope selector.
	Label string `json:"label"`
}

// SchemeScopesResponse Context scheme scopes response: the declared scopes in declaration order. Empty when the scheme declares none.
//
// An array rather than a map, because the order is the order a selector offers them in and a JSON object does not carry order. Which scope is the default is the summary's `defaultScope`; no entry here marks it.
type SchemeScopesResponse = []SchemeScope

// SchemeSummary Context scheme summary response.
type SchemeSummary struct {
	// ColorFrom Single-category scheme-default color source, when category mode is used.
	ColorFrom *string `json:"colorFrom,omitempty"`

	// ColorRules Ordered rule-driven color outcomes, when rule mode is used.
	ColorRules *[]MarkingColorRule `json:"colorRules,omitempty"`

	// ContextFormats Render format IDs grouped by role. For a scheme without `render.surfaces`, the lists `role_detection` gives, which the scanner reads a marking's role from. For a scheme with them, the scanner takes the role from the view a marking was read as, and these lists are derived from the views for a caller that picks a format per role: each role lists the forms of the views that state it, `format`, case formats and `also_read_in` in that order, surfaces and views in declared order and each format once.
	ContextFormats *struct {
		// Part Format ids a part marking is read in.
		Part []string `json:"part"`

		// Whole Format ids a whole marking is read in.
		Whole []string `json:"whole"`
	} `json:"contextFormats,omitempty"`

	// DefaultPartRenderFormat Present only for a scheme that declares `render.surfaces`: the output address a part marking is written in by default. It is the `surface/view` of the default scope's surface's default for a part, from that surface's `default`, and does not depend on `contextFormats`, which lists the formats each role is read in.
	DefaultPartRenderFormat *string `json:"defaultPartRenderFormat,omitempty"`

	// DefaultRenderFormat Render format applied when a call names none. A remote adapter has no way to derive this from the formats sub-resource: format declaration order is not default order, so a client that guesses the first entry renders in the wrong format whenever the two differ.
	//
	// Omitted when the scheme names no default. A client may then fall back to the first declared format, which is the same rule the engine applies, so both adapters render such a scheme the same way. The field is not a place to encode that absence: `default` is an ordinary format id, and a scheme declaring one would be indistinguishable from a scheme naming nothing.
	//
	// For a scheme that declares `render.surfaces`, this is the `surface/view` address of the view the default scope's surface writes a marking with no role in, which is its default for a whole.
	DefaultRenderFormat *string `json:"defaultRenderFormat,omitempty"`

	// DefaultScope Default scope id applied when a marking asserts no scope explicitly. Present exactly when the scheme declares scopes. Pairs with the per-category `scopes` lists on the categories sub-resource so a remote adapter can reconstruct which categories a scope-only assertion set targets: the metadata the assertion-set aggregate path needs to render scoped categories that the whole-marking format would otherwise drop.
	DefaultScope *string `json:"defaultScope,omitempty"`

	// DefaultWholeRenderFormat Present only for a scheme that declares `render.surfaces`: the output address a whole marking is written in by default, the default scope's surface's default for a whole, resolved the same way `defaultPartRenderFormat` is for a part.
	DefaultWholeRenderFormat *string `json:"defaultWholeRenderFormat,omitempty"`

	// Definition Human-readable definition of the scheme, for tooltips and about panels.
	Definition *string `json:"definition,omitempty"`

	// Description Free-text description from the scheme metadata. Omitted when the scheme has none.
	Description *string `json:"description,omitempty"`

	// EmbeddedMarking Ways a marking may appear embedded in free text. Omitted when the scheme declares none, which is how a client tells that inline marking is off.
	EmbeddedMarking *[]EmbeddedMarkingStyle `json:"embeddedMarking,omitempty"`

	// Guidance Usage guidance for the scheme, shown in reference views.
	Guidance *string `json:"guidance,omitempty"`

	// ID Scheme identifier from the scheme metadata.
	ID string `json:"id"`

	// Marking The scheme document's own overall marking, exactly as its `scheme.marking` writes it. Display only: never interpreted. Omitted when the document declares none.
	Marking *string `json:"marking,omitempty"`

	// Name Display name from the scheme metadata.
	Name string `json:"name"`

	// ReferenceSurfaceLabel Resolved display label for an artifact's reference-surface chrome noun: the name its own reference component wears (a scheme's Handbook, a guide's Guidebook, an access policy's Access Handbook). Always fully populated: an artifact that declares no override resolves to the engine's generic noun for that surface, using the same singular/plural `Term` contract as the artifact vocabulary. Kept as its own type because this is an already-resolved reference-surface value rather than the complete author-facing vocabulary.
	ReferenceSurfaceLabel ReferenceSurfaceLabel `json:"referenceSurfaceLabel"`

	// RoleAware Whether the scheme uses role-aware processing. Omitted when unset.
	RoleAware *bool `json:"roleAware,omitempty"`

	// ScopeLabel Display label for the scope selector shown in UIs. Omitted when the scheme declares none. Carried here so `getSchemeMeta` answers the same for a remote adapter as it does for a local one; without it the selector is labeled from the scheme document in one mode and left unlabeled in the other.
	ScopeLabel *string `json:"scopeLabel,omitempty"`

	// Version Scheme semantic version.
	Version string `json:"version"`

	// Vocabulary Author-facing vocabulary: an open map from a term key (e.g. "rule", "concept", "tag") to the display words. Term keys are not enumerated here; consumers pick the keys that matter to them.
	//
	// Purely a presentation hint: nothing in the engine's evaluation logic reads it. Each consumer owns its own default map and calls `resolveVocabulary(source, defaults)` to merge.
	Vocabulary *Vocabulary `json:"vocabulary,omitempty"`
}

// SchemeViewResponse Render-ready Handbook tree for a scheme: its concepts grouped under the scheme's categories.
type SchemeViewResponse struct {
	// Categories The scheme's categories in declared order, each with the concepts that belong to it. A category with no concepts is left out.
	Categories []struct {
		// ID The category id.
		ID string `json:"id"`

		// Items The category's concepts, in declared order.
		Items []ConceptView `json:"items"`

		// Label The category's display label.
		Label string `json:"label"`
	} `json:"categories"`

	// Description The scheme's description. Empty when the scheme declares none.
	Description string `json:"description"`

	// ID The scheme id.
	ID string `json:"id"`

	// Marking The scheme document's own overall marking, exactly as its `scheme.marking` writes it. Display only: never interpreted. Absent when the document declares none.
	Marking *string `json:"marking,omitempty"`

	// Name The scheme's display name.
	Name string `json:"name"`

	// Uncategorized Concepts that belong to none of the scheme's loaded categories. Usually empty.
	Uncategorized []ConceptView `json:"uncategorized"`

	// Vocabulary Author-facing vocabulary: an open map from a term key (e.g. "rule", "concept", "tag") to the display words. Term keys are not enumerated here; consumers pick the keys that matter to them.
	//
	// Purely a presentation hint: nothing in the engine's evaluation logic reads it. Each consumer owns its own default map and calls `resolveVocabulary(source, defaults)` to merge.
	Vocabulary *Vocabulary `json:"vocabulary,omitempty"`
}

// SortType How concepts within a category are sorted during rendering.
type SortType string

// SuggestResult What suggest answers: the concepts a picker may offer, and nothing about the marking itself.
//
// How the current marking reads and whether it stands are what validate and process answer, and suggest does not repeat them. A caller that wants both in one round trip asks process for them with `suggestFor`.
type SuggestResult struct {
	// Suggestions The concepts on offer, in display order.
	Suggestions []Suggestion `json:"suggestions"`

	// Warnings Advisories raised while building the options, such as a concept source that could not answer in full. Empty when there were none. Advisories about the marking itself belong to validate and process.
	Warnings []Warning `json:"warnings"`
}

// SuggestResultsByCategory Suggestions for several categories, keyed by category id.
//
// Each entry holds the options for its own category and no other, answered as if that category's own assertions were not yet chosen, so a picker can offer the alternatives to its current selection, not only what can still be added beside it. A label preset belongs to no category, so no entry offers one. `suggestForCategories` returns this, and so does a process call that names categories in `suggestFor`.
type SuggestResultsByCategory map[string]SuggestResult

// Suggestion A concept offered to the user in the suggest UI, with label, description, group, and current selection state.
//
// Returned in `SuggestResult.suggestions`. `selected: true` means the concept is already in the current marking; clicking it again removes it. `order` controls display position within the group. `views` carries all render-format representations so the UI can switch formats without a second query. `source` distinguishes concepts from parsing the input (`'parse'`), preset shortcuts (`'preset'`), or alias matches (`'alias'`).
type Suggestion struct {
	// CategoryID The category the concept belongs to.
	CategoryID string `json:"categoryId"`

	// ConceptID The id of the suggested concept.
	ConceptID string `json:"conceptId"`

	// Description Secondary label (longer name or alternate view).
	Description string `json:"description"`

	// Gated True when this concept matches the active filter but is currently gated by the scheme's rules (a rule prerequisite is unmet, or a constraint forbids it). Gated suggestions are surfaced as non-selectable, explanatory rows so the user can see why the concept is unavailable. Only populated when the caller passes `includeGated: true` and a filter is active.
	Gated *bool `json:"gated,omitempty"`

	// Group Category label for UI grouping.
	Group string `json:"group"`

	// Label Primary display label in the current render format.
	Label string `json:"label"`

	// Name Human-readable concept name.
	Name string `json:"name"`

	// Order Sort order.
	Order int `json:"order"`

	// Requires When `gated` is true, the unmet prerequisite concept ids that would unlock this concept. Empty array when the concept is gated by a constraint with no specific prerequisite (e.g. mutual exclusion).
	Requires *[]string `json:"requires,omitempty"`

	// RequiresDisplay Resolved operational display (rendered marking view) for each id in `requires`, index-aligned, so the unlock row can read "add X first" with the marking rather than the raw id. `requires` stays as the machine keys.
	RequiresDisplay *[]string `json:"requiresDisplay,omitempty"`

	// Selected True if this concept is already in the current marking.
	Selected bool `json:"selected"`

	// Source `preset` on a row that offers a label preset, whose `conceptId` is `preset:` followed by the preset id. A concept row leaves it unset.
	Source *SuggestionSource `json:"source,omitempty"`

	// Views All views for this concept.
	Views map[string]string `json:"views"`
}

// SuggestionSource `preset` on a row that offers a label preset, whose `conceptId` is `preset:` followed by the preset id. A concept row leaves it unset.
type SuggestionSource string

// Term A singular/plural display-word pair.
type Term struct {
	// Plural The word for more than one, such as `rules`.
	Plural *string `json:"plural,omitempty"`

	// Singular The word for one, such as `rule`.
	Singular *string `json:"singular,omitempty"`
}

// TextValueEncoding A free-text value stored exactly as supplied.
type TextValueEncoding struct {
	// MaxLength Reject a value longer than this many characters.
	MaxLength *int `json:"maxLength,omitempty"`

	// Pattern A validation constraint: the value is rejected unless the whole of it matches. Never a recognition device, and nothing is extracted from it.
	Pattern *string `json:"pattern,omitempty"`

	// Type Plain text, stored exactly as supplied.
	Type string `json:"type"`
}

// TierComposition How a tiered concept composes inside one field: what attaches a tier level to the one above it, and the view each half of the composed token is written in.
//
// The head is the level being composed with and the child is the level this entry governs. They need not agree: a format that spells its concepts out can still want the head short inside a composed token, where the composed form is a name in its own right. Either falls back to the field's view, and neither reaches a head written on its own.
type TierComposition struct {
	// AttachedBy What attaches this tier level to the one it composes with.
	AttachedBy string `json:"attachedBy"`

	// HeadView View for the composed token's head; falls back to the field's view.
	HeadView *string `json:"headView,omitempty"`

	// Tier The tier label this entry governs.
	Tier string `json:"tier"`

	// View View for the composed token's child; falls back to the field's view.
	View *string `json:"view,omitempty"`
}

// TokenWindow Proximity measured in whitespace-delimited tokens.
type TokenWindow struct {
	// Tokens The largest distance allowed between two spans, counted in whitespace-delimited tokens between their nearest edges.
	Tokens int `json:"tokens"`
}

// TraceEntry One step in the engine's explanation trace: which stage acted, what it did, and which concept/rule was involved.
//
// Lives in `ExplanationTrace.entries`. Consumers use the trace to render a "why was this marking produced?" breakdown. `stage` is one of `parse | repair | infer | validate | render`; `conceptId` and `ruleId` identify the specific element that drove the action.
type TraceEntry struct {
	// Action What the stage did, as a short machine-readable token, such as `asserted`, `inferred`, `removed` or `unmatched`. An open string.
	Action string `json:"action"`

	// CategoryID Category the action landed in, when the acting rule declared one. This is the structured channel for the category: `detail` is human-facing prose and must never be parsed to recover it.
	CategoryID *string `json:"categoryId,omitempty"`

	// ConceptID The concept the step acted on, when it acted on one.
	ConceptID *string `json:"conceptId,omitempty"`

	// ConceptName Resolved operational display for `conceptId` (the concept's rendered marking view), so a "why" breakdown UI can show the marking rather than the raw id. `conceptId` stays as the machine key.
	ConceptName *string `json:"conceptName,omitempty"`

	// Detail A sentence describing the step for a person. It is prose and must not be parsed; the structured fields carry the ids.
	Detail string `json:"detail"`

	// RuleID The scheme rule that drove the step, when one did.
	RuleID *string `json:"ruleId,omitempty"`

	// Stage The pipeline stage that acted, such as `parse`, `normalize`, `repair`, `infer`, `validate`, `render`, `aggregate` or `context`. An open string.
	Stage string `json:"stage"`

	// TriggerConceptIds Concept IDs that triggered the rule producing this entry.
	TriggerConceptIds *[]string `json:"triggerConceptIds,omitempty"`
}

// TraceInference Inference recorded in an explanation trace.
type TraceInference struct {
	// ConceptID The concept the rule added.
	ConceptID string `json:"conceptId"`

	// ConceptName Resolved operational display for `conceptId` (the concept's rendered marking view), so a "why" breakdown UI can show the inferred marking rather than the raw id. `conceptId` stays as the machine key.
	ConceptName *string `json:"conceptName,omitempty"`

	// Explanation Why the rule added the concept, in words.
	Explanation string `json:"explanation"`

	// RuleID The id of the inference rule that fired.
	RuleID string `json:"ruleId"`
}

// TraceRepair Repair recorded in an explanation trace.
type TraceRepair struct {
	// After The text after the repair.
	After string `json:"after"`

	// Before The text before the repair.
	Before string `json:"before"`

	// Cosmetic True when the repair changed only formatting and left the assertions as they were.
	Cosmetic *bool `json:"cosmetic,omitempty"`

	// Description What the repair does, as the scheme describes it.
	Description string `json:"description"`

	// RuleID The id of the repair rule that fired.
	RuleID string `json:"ruleId"`
}

// TraceValidationIssue Validation issue recorded in an explanation trace: the constraint that failed and its message, and nothing else. The category, severity and concepts of the error it came from belong to the error, which a validate or process result reports in `errors`.
type TraceValidationIssue struct {
	// ConstraintID The id of the constraint that failed.
	ConstraintID string `json:"constraintId"`

	// Message What is wrong, in words.
	Message string `json:"message"`
}

// UnreadMark A wrapped mark a scan found but could not read.
//
// The text matched the scheme's `embeddedMarking` pattern, which is how a part marking is written, but recognition read no marking in it, and it has no entry in `ScanOutcome.results`. Text that looks like a marking and carries none is reported rather than dropped, so a reader can see it and correct it. `start`/`end` are offsets in UTF-16 code units in the input text, as on `ScanResult`.
type UnreadMark struct {
	// End End position in the input text
	End int `json:"end"`

	// Inner The text inside the wrapper: the capture group of the scheme's `embeddedMarking` pattern.
	Inner string `json:"inner"`

	// Leading Whether the mark stands where a part mark goes: at the start of its paragraph. Both of these hold for a leading mark.
	//
	// Its line starts a paragraph. The line starts at the beginning of the text or just after the last line feed before `start`. It starts a paragraph when it is the first line, when the line before it is blank or ends in `.`, `!`, `?`, `:` or `;`, or when the line before it holds nothing but a marking read in the whole role, such as a banner.
	//
	// Only whitespace and outline labels come before it on its line: at most three labels, each followed by whitespace, in no more than 64 UTF-16 code units after the line's leading whitespace. A line feed is not whitespace here. An outline label is one or more groups of one to four letters or digits, each closed by `.` or `)` and optionally opened by `(` (`1.`, `a)`, `(1)`, `iv.`, `1.a.`, `(a)(1)`); a dotted numeral (`1.2`, `1.2.3.`); one bullet (`-`, `*`, `•`, `‣`, `⁃`, `◦`, `▪` or `–`); or one to six `#`. So the mark in `(QQ) text`, `1. (QQ) text` and `a) 2. (QQ) text` leads, and the one in `See (QQ) text` does not.
	//
	// A mark that follows another mark on its line does not lead, whether the earlier mark was read or not: in `(S)(QQ)` and in `(S) (QQ)` only `(S)` leads. Wrapped text before a mark is a label rather than a mark when its text is one lowercase letter, that letter two or three times, one to three digits, or a lowercase roman numeral, so the mark in `(a) (QQ) text` leads.
	Leading bool `json:"leading"`

	// Start Start position in the input text
	Start int `json:"start"`

	// Text The wrapped text as found in the document, wrapper included.
	Text string `json:"text"`
}

// UnrepresentableFact A fact this format has no way to write.
type UnrepresentableFact struct {
	// ConceptID The concept that cannot be written.
	ConceptID string `json:"conceptId"`

	// Detail Human-readable detail, for a message rather than a branch.
	Detail *string `json:"detail,omitempty"`

	// Reason Why this format cannot write a fact it was given.
	//
	// Closed, like the recognition-side reasons. `unframable` is a value carrying a character the syntax cannot frame; `reserved-literal` a free value equal to a fixed instruction beside it, which would read back as the other concept; `ambiguous-identity` a value indistinguishable from a differently typed one; `no-field` a fact this format's tree has no field for; `no-view` a concept with no spelling in this format's view; `value-lossy` a concept this format writes without its captured text.
	Reason UnrepresentableFactReason `json:"reason"`
}

// UnrepresentableFactReason Why this format cannot write a fact it was given.
//
// Closed, like the recognition-side reasons. `unframable` is a value carrying a character the syntax cannot frame; `reserved-literal` a free value equal to a fixed instruction beside it, which would read back as the other concept; `ambiguous-identity` a value indistinguishable from a differently typed one; `no-field` a fact this format's tree has no field for; `no-view` a concept with no spelling in this format's view; `value-lossy` a concept this format writes without its captured text.
type UnrepresentableFactReason string

// UnrepresentableOutput An output refused whole. Only a scheme that declares `render.surfaces` refuses one, because only its views choose a format by condition, and a choice that cannot be made is not a fact some format failed to write.
type UnrepresentableOutput struct {
	// Detail Human-readable detail, for a message rather than a branch.
	Detail *string `json:"detail,omitempty"`

	// Output The output refused, by its `surface/view` address.
	Output string `json:"output"`

	// Reason Why an output was refused whole, before any of its facts was written.
	//
	// Closed. `unresolved-context` is an output whose choice of format turns on a condition the operation cannot decide, such as one that reads context the caller did not supply. `ambiguous-representation` is an output with more than one representation in force and nothing to choose between them, such as two of its cases holding at once.
	Reason UnrepresentableOutputReason `json:"reason"`
}

// UnrepresentableOutputReason Why an output was refused whole, before any of its facts was written.
//
// Closed. `unresolved-context` is an output whose choice of format turns on a condition the operation cannot decide, such as one that reads context the caller did not supply. `ambiguous-representation` is an output with more than one representation in force and nothing to choose between them, such as two of its cases holding at once.
type UnrepresentableOutputReason string

// UnrepresentableValue Something a render could not write: one fact, or a whole output. `reason` tells the two apart, and only a fact carries a `conceptId`.
type UnrepresentableValue struct {
	union json.RawMessage
}

// ValidateResult What validate answers about one marking: the facts it states once settled, whether they stand, and what a reader needs to show why not.
//
// Validate is the reader's call. A host that displays a marking, gates on it, stores it, or hands its facts to an access check or a rollup calls it. The result carries no fix paths, no corrections to the text and no output beyond the one rendering, because a reader is not writing the marking. The author's call is process, and a `ProcessResult` is this result with the authoring answers added.
//
// Each field means the same wherever it appears: on a validate or process result, on each region of a scan, in-process or over a wire.
//
// Scheme and engine metadata (schemeVersion, schemeId) are not on the result; read them from `engine.version`, `engine.scheme.id` and `engine.scheme.version` instead.
type ValidateResult struct {
	// AppliedFixes Every change the engine actually made while settling the marking: a departure from canonical spelling or order that recognition read through, a repair, a fact a rule required, a handed-in fact settling removed, or text read in the other role from the one the request names. A removed fact is named with the reason it was removed, including a fact handed in with a derived source label that the run does not derive again. A requirement that conflicts with an existing concept is not added and appears in errors and proposed fixes, with no applied change. Repair never removes an author-stated concept. These are the reasons a marking can be `valid` while `inputWasValid` is false.
	//
	// A departure from spelling or order is reported here without a position. On a process result, `recognition.corrections` reports the same departure with the span of the text it covers. They are one change seen two ways, so a display shows one of them, not both. The role departure, `_recognition_role`, is about the whole text: it has no span and no entry in `recognition.corrections`, so it is reported only here, and a display that shows corrections in place of this list still shows it.
	AppliedFixes []Change `json:"appliedFixes"`

	// Assertions The facts the marking states, as the engine settled them.
	//
	// Listed in print order: the order the rendered marking places each concept; a concept the rendering leaves out appears where it would print. A concept the rendering has no place for at all follows the others of its category, or comes last when its category has none, in the order the scheme declares.
	Assertions []Assertion `json:"assertions"`

	// Assessment What a render did to the state it was given, bound to that state and that format.
	//
	// Answered against the state actually rendered: what was emitted, what was left out, what collapsed into an equivalent, and what could not be written at all.
	//
	// A caller that needs full-state preservation reads this. A caller that wants a human projection can ignore it and still get a useful string.
	Assessment RenderAssessment `json:"assessment"`

	// DisplayColor Portable scheme-default display colors, carried as a text/background pair so contrast is authored together. In category mode they are declared only on concepts in the category named by `render.color_from`; consuming hosts may override them.
	//
	// Both values are CSS hex (`#rgb` or `#rrggbb`). The pair travels together so a host never composes one authored color against an unknown surface (the white-on-white failure mode).
	DisplayColor *ConceptColor `json:"displayColor,omitempty"`

	// Errors What the scheme's rules find wrong with the settled marking. Any entry makes `valid` false.
	Errors []ValidationError `json:"errors"`

	// InputWasValid True when the text or facts as given needed no change to stand.
	//
	// Defined here once, and judged after the same settling as `valid`. A marking can be `valid` without having been valid as given: the engine settled it into one that stands. A result whose `inputWasValid` is false always says why: each change the engine made is in `appliedFixes`, and whatever it could not settle is in `errors`.
	//
	// Right as written is judged against the view the text was read as, the one `recognition.inputFormat` reports, or for text read on more than one surface, the view each surface's part was read as. A fact settling adds, including a restored tier parent, counts against it only when that view writes the fact. A view that leaves the fact out writes the same text with the fact as without it, so the text was right already. Text read as no single view counts the fact when any view that read it writes it. Facts handed in with no text are judged against what a request that names no output writes them in, by the rule `rendered` follows.
	//
	// Text is also not right as written when the request's context names a role and recognition decided the text is a view of the other role, as `recognition.role` reports. The result carries that departure in `appliedFixes` under `_recognition_role`, and `valid` is judged as it would be without it. A request that names its input format has said which view the text is: the text is judged as that view, and no role departure is raised. On a scan region and on a marking a document audit judges, the role the marking's place gives it stands for the context's role: a region found inside a wrapper is a part, and an audit judges each marking in the role it has in the document.
	InputWasValid bool `json:"inputWasValid"`

	// Notices The notices the scheme attaches to the concepts in `assertions`, collected the same way `SchemeEngine.collectNotices()` collects them. Always present, empty when the marking carries none, so a consumer can read the field without checking that it exists.
	Notices []Notice `json:"notices"`

	// Recognition What recognition made of an input string: how far it got, what it read, and what it could not account for.
	//
	// Every character of the marking region is accounted for: `complete` means recognized content, declared syntax, or permitted whitespace, with nothing left over. Finding one field is never enough to stop reading. An input that supports two distinct readings is `ambiguous` rather than read as either, because picking one by declaration order, category order, rank, or assertion count is how a guess becomes a confident answer.
	//
	// This is what a reader needs and what validate returns. Process returns `RecognitionWithCorrections`, which adds how to correct the text.
	Recognition *RecognitionResult `json:"recognition,omitempty"`

	// Rendered The settled marking, in one string.
	//
	// It is written in the output the request names. With none named, a scheme without `render.surfaces` writes its default format, and a scheme with them writes each surface in its default for the marking's role, by the rule `RenderConfigDoc.surfaces` states. Every operation that returns this shape follows that one rule.
	//
	// For a scheme that declares `render.surfaces`, the string is the output resolved on each surface the result writes, as `RenderConfigDoc.surfaces` describes, in declared surface order and joined by `render.surface_separator`. An output that writes nothing adds nothing, and outputs written as XML attributes in one namespace merge into one attribute list. One refused output refuses the whole string, and `assessment` says which one and why.
	Rendered string `json:"rendered"`

	// Valid True when the marking stands once the engine has settled it: read, repaired, completed by what the scheme's rules imply, and judged with no error left.
	//
	// Defined here once, and every result that reports `valid` means this, whichever operation produced it and whichever door it came through. Reading is tolerant on purpose, so the engine settles what it can. Whether the text was right as written is what `inputWasValid` and `appliedFixes` answer; no request asks the engine to settle less, so validate and process give the same `valid` for the same input, context, surface and parts.
	Valid bool `json:"valid"`

	// Warnings Advisories that do not affect `valid`, such as a deprecated concept. On a process result whose fix paths were not all proven, this carries a notice whose `ruleId` is `_fixes_unproven`, as `ProcessResult.proposedFixes` describes.
	Warnings []Warning `json:"warnings"`
}

// ValidationError A constraint violation detected during validation, referencing the constraint ID and a human-readable message.
//
// Carried in the `errors` of every result that judges a marking. `constraintId` identifies the rule that fired; `message` is human-readable. Every entry sets `valid` to `false`; advisories are warnings, not errors. Use `categoryId` and `affectedConcepts` to highlight the specific UI elements involved.
//
// Two of the ids the engine keeps for itself are for a fact a caller hands in that cannot be settled as given: `_missing_assertion_value`, a concept that takes a value handed in with none, and `_invalid_assertion_value`, a value the concept's encoding cannot read. Each is an error on the result, not a refused request: the fact is left out of the settled marking and named in `affectedConcepts`.
//
// Another, `_exceeds_whole`, is a part that carries more than the whole it is read in: the whole does not cover a value the part holds, carries one the part does not allow, or, in a category a rollup pipeline resolves, would change if rolled up with the part. It is given only for a part that is valid on its own, read in a whole that reads and is valid. `affectedConcepts` names every finding, and no fix answers it. See `ProcessingContext.partOf`.
//
// And `_stand_in_requires_whole` is a part that leaves a category to a stand-in, a concept that declares `stands_for_whole`, with no verified whole to take the values from. `affectedConcepts` names the stand-in and `categoryId` its category.
type ValidationError struct {
	// AffectedConcepts Concept IDs involved in this validation failure.
	AffectedConcepts *[]string `json:"affectedConcepts,omitempty"`

	// CategoryID Category ID this error relates to (when determinable from the constraint).
	CategoryID *string `json:"categoryId,omitempty"`

	// ConstraintID The id of the constraint that was violated: a scheme rule's id, or an engine-reserved id beginning with `_`.
	ConstraintID string `json:"constraintId"`

	// DidYouMean The ids of the concepts the unread token was likely meant as. Never empty. They are ids for a program to act on, not spellings to show an author: `message` gives the same hint in words, naming the spellings.
	//
	// Set only on an `_unmatched_token` error, which is about one token the engine could not read, and only when that token reads as a misspelling of some concept's spelling. The concepts named are tied for nearest, so their order ranks none above another.
	//
	// Its absence is no evidence that the token is stray text. A token too short to judge carries none, and so does one equally near to too many concepts. A token that already spells a concept, apart from case, carries none even when it is also near another concept: it was left unread for another reason, such as where it stands, and naming a different concept would point the author away from a word they spelled right.
	//
	// It is not a fix: nothing here has been proven to make the marking valid, and the result's fix paths do not depend on it.
	DidYouMean *[]string `json:"didYouMean,omitempty"`

	// Message What is wrong, in words a person can act on.
	Message string `json:"message"`
}

// ValueEncoding How a free-text value is spelled and stored: the resolved codec.
//
// A codec, not a validator. Every encoding states both directions, so a value this contract accepts is a value it can write back.
type ValueEncoding struct {
	union json.RawMessage
}

// Vocabulary Author-facing vocabulary: an open map from a term key (e.g. "rule", "concept", "tag") to the display words. Term keys are not enumerated here; consumers pick the keys that matter to them.
//
// Purely a presentation hint: nothing in the engine's evaluation logic reads it. Each consumer owns its own default map and calls `resolveVocabulary(source, defaults)` to merge.
type Vocabulary map[string]Term

// Warning A non-blocking informational signal surfaced in a result's `warnings`.
//
// Unlike `ValidationError`, warnings never affect `valid`. `kind: 'deprecation'` fires when a deprecated concept is asserted, and `replacement` names a suggested successor concept. `kind: 'rule'` fires when a scheme's authored `warn` rule matches the marking, with the rule's id as `ruleId`. `kind: 'notice'` is the engine speaking about the operation itself rather than about anything the scheme declares; its `ruleId`, when present, is an engine-reserved id beginning with `_`, the same convention a `ValidationError` follows.
//
// A marking's warnings come in a fixed order: deprecations first, then the scheme's rule warnings in the order the scheme declares its rules, then the engine's notices. An aggregate's warnings are about the rolled-up marking, not about each marking that fed it.
type Warning struct {
	// CategoryID Category this warning relates to (when applicable)
	CategoryID *string `json:"categoryId,omitempty"`

	// ConceptID Concept this warning relates to (when applicable). A rule warning names one when the rule's `when` names exactly one concept the marking holds.
	ConceptID *string `json:"conceptId,omitempty"`

	// ConceptName Resolved operational display for `conceptId` (the concept's rendered marking view), for a consumer that shows the warning next to the concept. `conceptId` stays as the machine key; this is the human-visible string.
	ConceptName *string `json:"conceptName,omitempty"`

	// Kind Non-blocking warning category: a deprecated concept, a scheme's authored `warn` rule, or a notice the engine gives about the operation itself.
	Kind WarningKind `json:"kind"`

	// Message Human-readable message
	Message string `json:"message"`

	// Replacement When kind=deprecation, the suggested replacement concept id (if any)
	Replacement *string `json:"replacement,omitempty"`

	// ReplacementName When kind=deprecation, the resolved operational display for `replacement` (the successor concept's rendered marking view). `replacement` stays as the machine key; this is the human-visible string.
	ReplacementName *string `json:"replacementName,omitempty"`

	// RuleID The scheme rule that gave a `rule` warning, or the engine-reserved id of a `notice`. Absent on a deprecation warning.
	RuleID *string `json:"ruleId,omitempty"`
}

// WarningKind Non-blocking warning category: a deprecated concept, a scheme's authored `warn` rule, or a notice the engine gives about the operation itself.
type WarningKind string

// Win A bounded proximity window. Exactly one token or character bound is present, so a windowed operation can never widen to the whole line accidentally.
type Win struct {
	union json.RawMessage
}

// ApplicationID defines model for ApplicationId.
type ApplicationID = string

// BadRequest Response body for error responses (4xx / 5xx).
type BadRequest = ErrorResponse

// Conflict Response body for error responses (4xx / 5xx).
type Conflict = ErrorResponse

// InternalError Response body for error responses (4xx / 5xx).
type InternalError = ErrorResponse

// NotFound Response body for error responses (4xx / 5xx).
type NotFound = ErrorResponse

// NotReady Response body for GET /v1/ready
type NotReady = ReadyResponse

// TooManyRequests Response body for 429 Too Many Requests.
type TooManyRequests = RateLimitError

// UnprocessableEntity Response body for error responses (4xx / 5xx).
type UnprocessableEntity = ErrorResponse

// CtxListParams defines parameters for CtxList.
type CtxListParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxVersionsParams defines parameters for CtxVersions.
type CtxVersionsParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxDescribeParams defines parameters for CtxDescribe.
type CtxDescribeParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxAccessCheckParams defines parameters for CtxAccessCheck.
type CtxAccessCheckParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxAccessDescribeRequirementsParams defines parameters for CtxAccessDescribeRequirements.
type CtxAccessDescribeRequirementsParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxAccessRulesParams defines parameters for CtxAccessRules.
type CtxAccessRulesParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxCrosswalkDescribeParams defines parameters for CtxCrosswalkDescribe.
type CtxCrosswalkDescribeParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxCrosswalkTranslateParams defines parameters for CtxCrosswalkTranslate.
type CtxCrosswalkTranslateParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxDocumentAuditParams defines parameters for CtxDocumentAudit.
type CtxDocumentAuditParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxGuideEvaluateParams defines parameters for CtxGuideEvaluate.
type CtxGuideEvaluateParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxGuideScanParams defines parameters for CtxGuideScan.
type CtxGuideScanParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxGuideViewParams defines parameters for CtxGuideView.
type CtxGuideViewParams struct {
	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeSummaryParams defines parameters for CtxSchemeSummary.
type CtxSchemeSummaryParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeAggregateParams defines parameters for CtxSchemeAggregate.
type CtxSchemeAggregateParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeApplyPresetParams defines parameters for CtxSchemeApplyPreset.
type CtxSchemeApplyPresetParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeCategoriesParams defines parameters for CtxSchemeCategories.
type CtxSchemeCategoriesParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeCheckCoverageParams defines parameters for CtxSchemeCheckCoverage.
type CtxSchemeCheckCoverageParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeConceptsParams defines parameters for CtxSchemeConcepts.
type CtxSchemeConceptsParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeDescribeConceptParams defines parameters for CtxSchemeDescribeConcept.
type CtxSchemeDescribeConceptParams struct {
	// ConceptID Concept identifier to describe.
	ConceptID string `form:"conceptId" json:"conceptId"`

	// CategoryID Category identifier the concept belongs to.
	CategoryID string `form:"categoryId" json:"categoryId"`

	// View The view to name the concept in: a key of the concept's `views`. Omitted, its default view.
	View *string `form:"view,omitempty" json:"view,omitempty"`

	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeDiffParams defines parameters for CtxSchemeDiff.
type CtxSchemeDiffParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeExplainParams defines parameters for CtxSchemeExplain.
type CtxSchemeExplainParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeFormatsParams defines parameters for CtxSchemeFormats.
type CtxSchemeFormatsParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeLabelPresetsParams defines parameters for CtxSchemeLabelPresets.
type CtxSchemeLabelPresetsParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeNoticesParams defines parameters for CtxSchemeNotices.
type CtxSchemeNoticesParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeProcessParams defines parameters for CtxSchemeProcess.
type CtxSchemeProcessParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeRenderParams defines parameters for CtxSchemeRender.
type CtxSchemeRenderParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeScanParams defines parameters for CtxSchemeScan.
type CtxSchemeScanParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeScopesParams defines parameters for CtxSchemeScopes.
type CtxSchemeScopesParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeSuggestParams defines parameters for CtxSchemeSuggest.
type CtxSchemeSuggestParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeSuggestForCategoriesParams defines parameters for CtxSchemeSuggestForCategories.
type CtxSchemeSuggestForCategoriesParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeValidateParams defines parameters for CtxSchemeValidate.
type CtxSchemeValidateParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxSchemeViewParams defines parameters for CtxSchemeView.
type CtxSchemeViewParams struct {
	// Scheme Context-local key of the scheme the request is answered in, as `GET /v1/ctx/{ctxId}/{version}` lists it. Without it the context's primary scheme answers. Any scheme the context holds may be named, a crosswalk's source or target among them. A key the context does not hold answers 404 with code `scheme_not_found`; it never falls back to the primary. Name it once: a repeated selector answers 400.
	Scheme *string `form:"scheme,omitempty" json:"scheme,omitempty"`

	// XEzclassifyApplicationID Identifies the calling application for per-application usage reporting: the id the application is registered under. It is not a credential and authorizes nothing. A request without it, or with an id the deployment has not registered, is served the same way. An id is 1 to 128 printable ASCII characters (0x21 to 0x7E), with no spaces. A value that does not match is served and counted as unidentified.
	XEzclassifyApplicationID *ApplicationID `json:"x-ezclassify-application-id,omitempty"`
}

// CtxAccessCheckJSONRequestBody defines body for CtxAccessCheck for application/json ContentType.
type CtxAccessCheckJSONRequestBody = CtxAccessCheckRequestBody

// CtxAccessDescribeRequirementsJSONRequestBody defines body for CtxAccessDescribeRequirements for application/json ContentType.
type CtxAccessDescribeRequirementsJSONRequestBody = CtxAccessDescribeRequirementsRequestBody

// CtxCrosswalkTranslateJSONRequestBody defines body for CtxCrosswalkTranslate for application/json ContentType.
type CtxCrosswalkTranslateJSONRequestBody = CtxCrosswalkTranslateRequestBody

// CtxDocumentAuditJSONRequestBody defines body for CtxDocumentAudit for application/json ContentType.
type CtxDocumentAuditJSONRequestBody = CtxDocumentAuditRequestBody

// CtxGuideEvaluateJSONRequestBody defines body for CtxGuideEvaluate for application/json ContentType.
type CtxGuideEvaluateJSONRequestBody = CtxGuideEvaluateRequestBody

// CtxGuideScanJSONRequestBody defines body for CtxGuideScan for application/json ContentType.
type CtxGuideScanJSONRequestBody = CtxGuideScanRequestBody

// CtxGuideViewJSONRequestBody defines body for CtxGuideView for application/json ContentType.
type CtxGuideViewJSONRequestBody = CtxGuideViewRequestBody

// CtxSchemeAggregateJSONRequestBody defines body for CtxSchemeAggregate for application/json ContentType.
type CtxSchemeAggregateJSONRequestBody = CtxSchemeAggregateRequestBody

// CtxSchemeApplyPresetJSONRequestBody defines body for CtxSchemeApplyPreset for application/json ContentType.
type CtxSchemeApplyPresetJSONRequestBody = ApplyPresetRequest

// CtxSchemeCheckCoverageJSONRequestBody defines body for CtxSchemeCheckCoverage for application/json ContentType.
type CtxSchemeCheckCoverageJSONRequestBody = CtxSchemeCheckCoverageRequestBody

// CtxSchemeDiffJSONRequestBody defines body for CtxSchemeDiff for application/json ContentType.
type CtxSchemeDiffJSONRequestBody = CtxSchemeDiffRequestBody

// CtxSchemeExplainJSONRequestBody defines body for CtxSchemeExplain for application/json ContentType.
type CtxSchemeExplainJSONRequestBody = CtxSchemeExplainRequestBody

// CtxSchemeNoticesJSONRequestBody defines body for CtxSchemeNotices for application/json ContentType.
type CtxSchemeNoticesJSONRequestBody = CtxSchemeNoticesRequestBody

// CtxSchemeProcessJSONRequestBody defines body for CtxSchemeProcess for application/json ContentType.
type CtxSchemeProcessJSONRequestBody = CtxSchemeProcessRequestBody

// CtxSchemeRenderJSONRequestBody defines body for CtxSchemeRender for application/json ContentType.
type CtxSchemeRenderJSONRequestBody = CtxSchemeRenderRequestBody

// CtxSchemeScanJSONRequestBody defines body for CtxSchemeScan for application/json ContentType.
type CtxSchemeScanJSONRequestBody = CtxSchemeScanRequestBody

// CtxSchemeSuggestJSONRequestBody defines body for CtxSchemeSuggest for application/json ContentType.
type CtxSchemeSuggestJSONRequestBody = CtxSchemeSuggestRequestBody

// CtxSchemeSuggestForCategoriesJSONRequestBody defines body for CtxSchemeSuggestForCategories for application/json ContentType.
type CtxSchemeSuggestForCategoriesJSONRequestBody = CtxSchemeSuggestForCategoriesRequestBody

// CtxSchemeValidateJSONRequestBody defines body for CtxSchemeValidate for application/json ContentType.
type CtxSchemeValidateJSONRequestBody = CtxSchemeValidateRequestBody

// CtxSchemeViewJSONRequestBody defines body for CtxSchemeView for application/json ContentType.
type CtxSchemeViewJSONRequestBody = CtxSchemeViewRequestBody

// AsCategoryViewDefaultConcept0 returns the union data inside the CategoryView_DefaultConcept as a CategoryViewDefaultConcept0
func (t CategoryView_DefaultConcept) AsCategoryViewDefaultConcept0() (CategoryViewDefaultConcept0, error) {
	var body CategoryViewDefaultConcept0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCategoryViewDefaultConcept0 overwrites any union data inside the CategoryView_DefaultConcept as the provided CategoryViewDefaultConcept0
func (t *CategoryView_DefaultConcept) FromCategoryViewDefaultConcept0(v CategoryViewDefaultConcept0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCategoryViewDefaultConcept0 performs a merge with any union data inside the CategoryView_DefaultConcept, using the provided CategoryViewDefaultConcept0
func (t *CategoryView_DefaultConcept) MergeCategoryViewDefaultConcept0(v CategoryViewDefaultConcept0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsConditionalDefaultConcept returns the union data inside the CategoryView_DefaultConcept as a ConditionalDefaultConcept
func (t CategoryView_DefaultConcept) AsConditionalDefaultConcept() (ConditionalDefaultConcept, error) {
	var body ConditionalDefaultConcept
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromConditionalDefaultConcept overwrites any union data inside the CategoryView_DefaultConcept as the provided ConditionalDefaultConcept
func (t *CategoryView_DefaultConcept) FromConditionalDefaultConcept(v ConditionalDefaultConcept) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeConditionalDefaultConcept performs a merge with any union data inside the CategoryView_DefaultConcept, using the provided ConditionalDefaultConcept
func (t *CategoryView_DefaultConcept) MergeConditionalDefaultConcept(v ConditionalDefaultConcept) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CategoryView_DefaultConcept) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CategoryView_DefaultConcept) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSortType returns the union data inside the CategoryView_Sort as a SortType
func (t CategoryView_Sort) AsSortType() (SortType, error) {
	var body SortType
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSortType overwrites any union data inside the CategoryView_Sort as the provided SortType
func (t *CategoryView_Sort) FromSortType(v SortType) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSortType performs a merge with any union data inside the CategoryView_Sort, using the provided SortType
func (t *CategoryView_Sort) MergeSortType(v SortType) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCategoryViewSort1 returns the union data inside the CategoryView_Sort as a CategoryViewSort1
func (t CategoryView_Sort) AsCategoryViewSort1() (CategoryViewSort1, error) {
	var body CategoryViewSort1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCategoryViewSort1 overwrites any union data inside the CategoryView_Sort as the provided CategoryViewSort1
func (t *CategoryView_Sort) FromCategoryViewSort1(v CategoryViewSort1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCategoryViewSort1 performs a merge with any union data inside the CategoryView_Sort, using the provided CategoryViewSort1
func (t *CategoryView_Sort) MergeCategoryViewSort1(v CategoryViewSort1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CategoryView_Sort) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CategoryView_Sort) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCrosswalkMappingForbidMatch returns the union data inside the CrosswalkForbidMatch as a CrosswalkMappingForbidMatch
func (t CrosswalkForbidMatch) AsCrosswalkMappingForbidMatch() (CrosswalkMappingForbidMatch, error) {
	var body CrosswalkMappingForbidMatch
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkMappingForbidMatch overwrites any union data inside the CrosswalkForbidMatch as the provided CrosswalkMappingForbidMatch
func (t *CrosswalkForbidMatch) FromCrosswalkMappingForbidMatch(v CrosswalkMappingForbidMatch) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkMappingForbidMatch performs a merge with any union data inside the CrosswalkForbidMatch, using the provided CrosswalkMappingForbidMatch
func (t *CrosswalkForbidMatch) MergeCrosswalkMappingForbidMatch(v CrosswalkMappingForbidMatch) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCrosswalkSelectionForbidMatch returns the union data inside the CrosswalkForbidMatch as a CrosswalkSelectionForbidMatch
func (t CrosswalkForbidMatch) AsCrosswalkSelectionForbidMatch() (CrosswalkSelectionForbidMatch, error) {
	var body CrosswalkSelectionForbidMatch
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkSelectionForbidMatch overwrites any union data inside the CrosswalkForbidMatch as the provided CrosswalkSelectionForbidMatch
func (t *CrosswalkForbidMatch) FromCrosswalkSelectionForbidMatch(v CrosswalkSelectionForbidMatch) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkSelectionForbidMatch performs a merge with any union data inside the CrosswalkForbidMatch, using the provided CrosswalkSelectionForbidMatch
func (t *CrosswalkForbidMatch) MergeCrosswalkSelectionForbidMatch(v CrosswalkSelectionForbidMatch) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CrosswalkForbidMatch) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CrosswalkForbidMatch) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCrosswalkDisposition returns the union data inside the CrosswalkOutcome_Disposition as a CrosswalkDisposition
func (t CrosswalkOutcome_Disposition) AsCrosswalkDisposition() (CrosswalkDisposition, error) {
	var body CrosswalkDisposition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkDisposition overwrites any union data inside the CrosswalkOutcome_Disposition as the provided CrosswalkDisposition
func (t *CrosswalkOutcome_Disposition) FromCrosswalkDisposition(v CrosswalkDisposition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkDisposition performs a merge with any union data inside the CrosswalkOutcome_Disposition, using the provided CrosswalkDisposition
func (t *CrosswalkOutcome_Disposition) MergeCrosswalkDisposition(v CrosswalkDisposition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCrosswalkOutcomeDisposition1 returns the union data inside the CrosswalkOutcome_Disposition as a CrosswalkOutcomeDisposition1
func (t CrosswalkOutcome_Disposition) AsCrosswalkOutcomeDisposition1() (CrosswalkOutcomeDisposition1, error) {
	var body CrosswalkOutcomeDisposition1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkOutcomeDisposition1 overwrites any union data inside the CrosswalkOutcome_Disposition as the provided CrosswalkOutcomeDisposition1
func (t *CrosswalkOutcome_Disposition) FromCrosswalkOutcomeDisposition1(v CrosswalkOutcomeDisposition1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkOutcomeDisposition1 performs a merge with any union data inside the CrosswalkOutcome_Disposition, using the provided CrosswalkOutcomeDisposition1
func (t *CrosswalkOutcome_Disposition) MergeCrosswalkOutcomeDisposition1(v CrosswalkOutcomeDisposition1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCrosswalkOutcomeDisposition2 returns the union data inside the CrosswalkOutcome_Disposition as a CrosswalkOutcomeDisposition2
func (t CrosswalkOutcome_Disposition) AsCrosswalkOutcomeDisposition2() (CrosswalkOutcomeDisposition2, error) {
	var body CrosswalkOutcomeDisposition2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkOutcomeDisposition2 overwrites any union data inside the CrosswalkOutcome_Disposition as the provided CrosswalkOutcomeDisposition2
func (t *CrosswalkOutcome_Disposition) FromCrosswalkOutcomeDisposition2(v CrosswalkOutcomeDisposition2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkOutcomeDisposition2 performs a merge with any union data inside the CrosswalkOutcome_Disposition, using the provided CrosswalkOutcomeDisposition2
func (t *CrosswalkOutcome_Disposition) MergeCrosswalkOutcomeDisposition2(v CrosswalkOutcomeDisposition2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CrosswalkOutcome_Disposition) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CrosswalkOutcome_Disposition) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCrosswalkCompletedTranslationResult returns the union data inside the CrosswalkTranslationResult as a CrosswalkCompletedTranslationResult
func (t CrosswalkTranslationResult) AsCrosswalkCompletedTranslationResult() (CrosswalkCompletedTranslationResult, error) {
	var body CrosswalkCompletedTranslationResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkCompletedTranslationResult overwrites any union data inside the CrosswalkTranslationResult as the provided CrosswalkCompletedTranslationResult
func (t *CrosswalkTranslationResult) FromCrosswalkCompletedTranslationResult(v CrosswalkCompletedTranslationResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkCompletedTranslationResult performs a merge with any union data inside the CrosswalkTranslationResult, using the provided CrosswalkCompletedTranslationResult
func (t *CrosswalkTranslationResult) MergeCrosswalkCompletedTranslationResult(v CrosswalkCompletedTranslationResult) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCrosswalkForbiddenTranslationResult returns the union data inside the CrosswalkTranslationResult as a CrosswalkForbiddenTranslationResult
func (t CrosswalkTranslationResult) AsCrosswalkForbiddenTranslationResult() (CrosswalkForbiddenTranslationResult, error) {
	var body CrosswalkForbiddenTranslationResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCrosswalkForbiddenTranslationResult overwrites any union data inside the CrosswalkTranslationResult as the provided CrosswalkForbiddenTranslationResult
func (t *CrosswalkTranslationResult) FromCrosswalkForbiddenTranslationResult(v CrosswalkForbiddenTranslationResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCrosswalkForbiddenTranslationResult performs a merge with any union data inside the CrosswalkTranslationResult, using the provided CrosswalkForbiddenTranslationResult
func (t *CrosswalkTranslationResult) MergeCrosswalkForbiddenTranslationResult(v CrosswalkForbiddenTranslationResult) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t CrosswalkTranslationResult) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CrosswalkTranslationResult) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFormatSyntaxType0 returns the union data inside the FormatSyntaxType as a FormatSyntaxType0
func (t FormatSyntaxType) AsFormatSyntaxType0() (FormatSyntaxType0, error) {
	var body FormatSyntaxType0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFormatSyntaxType0 overwrites any union data inside the FormatSyntaxType as the provided FormatSyntaxType0
func (t *FormatSyntaxType) FromFormatSyntaxType0(v FormatSyntaxType0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFormatSyntaxType0 performs a merge with any union data inside the FormatSyntaxType, using the provided FormatSyntaxType0
func (t *FormatSyntaxType) MergeFormatSyntaxType0(v FormatSyntaxType0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFormatSyntaxType1 returns the union data inside the FormatSyntaxType as a FormatSyntaxType1
func (t FormatSyntaxType) AsFormatSyntaxType1() (FormatSyntaxType1, error) {
	var body FormatSyntaxType1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFormatSyntaxType1 overwrites any union data inside the FormatSyntaxType as the provided FormatSyntaxType1
func (t *FormatSyntaxType) FromFormatSyntaxType1(v FormatSyntaxType1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFormatSyntaxType1 performs a merge with any union data inside the FormatSyntaxType, using the provided FormatSyntaxType1
func (t *FormatSyntaxType) MergeFormatSyntaxType1(v FormatSyntaxType1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FormatSyntaxType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FormatSyntaxType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGuideWhenHas returns the union data inside the GuideWhenNode as a GuideWhenHas
func (t GuideWhenNode) AsGuideWhenHas() (GuideWhenHas, error) {
	var body GuideWhenHas
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideWhenHas overwrites any union data inside the GuideWhenNode as the provided GuideWhenHas
func (t *GuideWhenNode) FromGuideWhenHas(v GuideWhenHas) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideWhenHas performs a merge with any union data inside the GuideWhenNode, using the provided GuideWhenHas
func (t *GuideWhenNode) MergeGuideWhenHas(v GuideWhenHas) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGuideWhenTest returns the union data inside the GuideWhenNode as a GuideWhenTest
func (t GuideWhenNode) AsGuideWhenTest() (GuideWhenTest, error) {
	var body GuideWhenTest
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideWhenTest overwrites any union data inside the GuideWhenNode as the provided GuideWhenTest
func (t *GuideWhenNode) FromGuideWhenTest(v GuideWhenTest) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideWhenTest performs a merge with any union data inside the GuideWhenNode, using the provided GuideWhenTest
func (t *GuideWhenNode) MergeGuideWhenTest(v GuideWhenTest) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGuideWhen returns the union data inside the GuideWhenNode as a GuideWhen
func (t GuideWhenNode) AsGuideWhen() (GuideWhen, error) {
	var body GuideWhen
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideWhen overwrites any union data inside the GuideWhenNode as the provided GuideWhen
func (t *GuideWhenNode) FromGuideWhen(v GuideWhen) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideWhen performs a merge with any union data inside the GuideWhenNode, using the provided GuideWhen
func (t *GuideWhenNode) MergeGuideWhen(v GuideWhen) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t GuideWhenNode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GuideWhenNode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGuideWhenTestValue0 returns the union data inside the GuideWhenTest_Value as a GuideWhenTestValue0
func (t GuideWhenTest_Value) AsGuideWhenTestValue0() (GuideWhenTestValue0, error) {
	var body GuideWhenTestValue0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideWhenTestValue0 overwrites any union data inside the GuideWhenTest_Value as the provided GuideWhenTestValue0
func (t *GuideWhenTest_Value) FromGuideWhenTestValue0(v GuideWhenTestValue0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideWhenTestValue0 performs a merge with any union data inside the GuideWhenTest_Value, using the provided GuideWhenTestValue0
func (t *GuideWhenTest_Value) MergeGuideWhenTestValue0(v GuideWhenTestValue0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsGuideWhenTestValue1 returns the union data inside the GuideWhenTest_Value as a GuideWhenTestValue1
func (t GuideWhenTest_Value) AsGuideWhenTestValue1() (GuideWhenTestValue1, error) {
	var body GuideWhenTestValue1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideWhenTestValue1 overwrites any union data inside the GuideWhenTest_Value as the provided GuideWhenTestValue1
func (t *GuideWhenTest_Value) FromGuideWhenTestValue1(v GuideWhenTestValue1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideWhenTestValue1 performs a merge with any union data inside the GuideWhenTest_Value, using the provided GuideWhenTestValue1
func (t *GuideWhenTest_Value) MergeGuideWhenTestValue1(v GuideWhenTestValue1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t GuideWhenTest_Value) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GuideWhenTest_Value) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGuideMatchSpec returns the union data inside the MatchNode as a GuideMatchSpec
func (t MatchNode) AsGuideMatchSpec() (GuideMatchSpec, error) {
	var body GuideMatchSpec
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGuideMatchSpec overwrites any union data inside the MatchNode as the provided GuideMatchSpec
func (t *MatchNode) FromGuideMatchSpec(v GuideMatchSpec) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGuideMatchSpec performs a merge with any union data inside the MatchNode, using the provided GuideMatchSpec
func (t *MatchNode) MergeGuideMatchSpec(v GuideMatchSpec) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchAnyNode returns the union data inside the MatchNode as a MatchAnyNode
func (t MatchNode) AsMatchAnyNode() (MatchAnyNode, error) {
	var body MatchAnyNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchAnyNode overwrites any union data inside the MatchNode as the provided MatchAnyNode
func (t *MatchNode) FromMatchAnyNode(v MatchAnyNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchAnyNode performs a merge with any union data inside the MatchNode, using the provided MatchAnyNode
func (t *MatchNode) MergeMatchAnyNode(v MatchAnyNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchAllNode returns the union data inside the MatchNode as a MatchAllNode
func (t MatchNode) AsMatchAllNode() (MatchAllNode, error) {
	var body MatchAllNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchAllNode overwrites any union data inside the MatchNode as the provided MatchAllNode
func (t *MatchNode) FromMatchAllNode(v MatchAllNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchAllNode performs a merge with any union data inside the MatchNode, using the provided MatchAllNode
func (t *MatchNode) MergeMatchAllNode(v MatchAllNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchNotNode returns the union data inside the MatchNode as a MatchNotNode
func (t MatchNode) AsMatchNotNode() (MatchNotNode, error) {
	var body MatchNotNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchNotNode overwrites any union data inside the MatchNode as the provided MatchNotNode
func (t *MatchNode) FromMatchNotNode(v MatchNotNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchNotNode performs a merge with any union data inside the MatchNode, using the provided MatchNotNode
func (t *MatchNode) MergeMatchNotNode(v MatchNotNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchNearNode returns the union data inside the MatchNode as a MatchNearNode
func (t MatchNode) AsMatchNearNode() (MatchNearNode, error) {
	var body MatchNearNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchNearNode overwrites any union data inside the MatchNode as the provided MatchNearNode
func (t *MatchNode) FromMatchNearNode(v MatchNearNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchNearNode performs a merge with any union data inside the MatchNode, using the provided MatchNearNode
func (t *MatchNode) MergeMatchNearNode(v MatchNearNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchContextNode returns the union data inside the MatchNode as a MatchContextNode
func (t MatchNode) AsMatchContextNode() (MatchContextNode, error) {
	var body MatchContextNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchContextNode overwrites any union data inside the MatchNode as the provided MatchContextNode
func (t *MatchNode) FromMatchContextNode(v MatchContextNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchContextNode performs a merge with any union data inside the MatchNode, using the provided MatchContextNode
func (t *MatchNode) MergeMatchContextNode(v MatchContextNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsMatchValidateNode returns the union data inside the MatchNode as a MatchValidateNode
func (t MatchNode) AsMatchValidateNode() (MatchValidateNode, error) {
	var body MatchValidateNode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromMatchValidateNode overwrites any union data inside the MatchNode as the provided MatchValidateNode
func (t *MatchNode) FromMatchValidateNode(v MatchValidateNode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeMatchValidateNode performs a merge with any union data inside the MatchNode, using the provided MatchValidateNode
func (t *MatchNode) MergeMatchValidateNode(v MatchValidateNode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t MatchNode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *MatchNode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsRuleConditionValue0 returns the union data inside the RuleCondition_Value as a RuleConditionValue0
func (t RuleCondition_Value) AsRuleConditionValue0() (RuleConditionValue0, error) {
	var body RuleConditionValue0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleConditionValue0 overwrites any union data inside the RuleCondition_Value as the provided RuleConditionValue0
func (t *RuleCondition_Value) FromRuleConditionValue0(v RuleConditionValue0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleConditionValue0 performs a merge with any union data inside the RuleCondition_Value, using the provided RuleConditionValue0
func (t *RuleCondition_Value) MergeRuleConditionValue0(v RuleConditionValue0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRuleConditionValue1 returns the union data inside the RuleCondition_Value as a RuleConditionValue1
func (t RuleCondition_Value) AsRuleConditionValue1() (RuleConditionValue1, error) {
	var body RuleConditionValue1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleConditionValue1 overwrites any union data inside the RuleCondition_Value as the provided RuleConditionValue1
func (t *RuleCondition_Value) FromRuleConditionValue1(v RuleConditionValue1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleConditionValue1 performs a merge with any union data inside the RuleCondition_Value, using the provided RuleConditionValue1
func (t *RuleCondition_Value) MergeRuleConditionValue1(v RuleConditionValue1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t RuleCondition_Value) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *RuleCondition_Value) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsRuleWhen0 returns the union data inside the RuleWhen as a RuleWhen0
func (t RuleWhen) AsRuleWhen0() (RuleWhen0, error) {
	var body RuleWhen0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleWhen0 overwrites any union data inside the RuleWhen as the provided RuleWhen0
func (t *RuleWhen) FromRuleWhen0(v RuleWhen0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleWhen0 performs a merge with any union data inside the RuleWhen, using the provided RuleWhen0
func (t *RuleWhen) MergeRuleWhen0(v RuleWhen0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRuleWhen1 returns the union data inside the RuleWhen as a RuleWhen1
func (t RuleWhen) AsRuleWhen1() (RuleWhen1, error) {
	var body RuleWhen1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleWhen1 overwrites any union data inside the RuleWhen as the provided RuleWhen1
func (t *RuleWhen) FromRuleWhen1(v RuleWhen1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleWhen1 performs a merge with any union data inside the RuleWhen, using the provided RuleWhen1
func (t *RuleWhen) MergeRuleWhen1(v RuleWhen1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRuleWhenQuantifier returns the union data inside the RuleWhen as a RuleWhenQuantifier
func (t RuleWhen) AsRuleWhenQuantifier() (RuleWhenQuantifier, error) {
	var body RuleWhenQuantifier
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleWhenQuantifier overwrites any union data inside the RuleWhen as the provided RuleWhenQuantifier
func (t *RuleWhen) FromRuleWhenQuantifier(v RuleWhenQuantifier) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleWhenQuantifier performs a merge with any union data inside the RuleWhen, using the provided RuleWhenQuantifier
func (t *RuleWhen) MergeRuleWhenQuantifier(v RuleWhenQuantifier) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t RuleWhen) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *RuleWhen) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsRuleCondition returns the union data inside the RuleWhenNode as a RuleCondition
func (t RuleWhenNode) AsRuleCondition() (RuleCondition, error) {
	var body RuleCondition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleCondition overwrites any union data inside the RuleWhenNode as the provided RuleCondition
func (t *RuleWhenNode) FromRuleCondition(v RuleCondition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleCondition performs a merge with any union data inside the RuleWhenNode, using the provided RuleCondition
func (t *RuleWhenNode) MergeRuleCondition(v RuleCondition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsRuleWhenQuantifier returns the union data inside the RuleWhenNode as a RuleWhenQuantifier
func (t RuleWhenNode) AsRuleWhenQuantifier() (RuleWhenQuantifier, error) {
	var body RuleWhenQuantifier
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRuleWhenQuantifier overwrites any union data inside the RuleWhenNode as the provided RuleWhenQuantifier
func (t *RuleWhenNode) FromRuleWhenQuantifier(v RuleWhenQuantifier) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRuleWhenQuantifier performs a merge with any union data inside the RuleWhenNode, using the provided RuleWhenQuantifier
func (t *RuleWhenNode) MergeRuleWhenQuantifier(v RuleWhenQuantifier) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t RuleWhenNode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *RuleWhenNode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsUnrepresentableFact returns the union data inside the UnrepresentableValue as a UnrepresentableFact
func (t UnrepresentableValue) AsUnrepresentableFact() (UnrepresentableFact, error) {
	var body UnrepresentableFact
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUnrepresentableFact overwrites any union data inside the UnrepresentableValue as the provided UnrepresentableFact
func (t *UnrepresentableValue) FromUnrepresentableFact(v UnrepresentableFact) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUnrepresentableFact performs a merge with any union data inside the UnrepresentableValue, using the provided UnrepresentableFact
func (t *UnrepresentableValue) MergeUnrepresentableFact(v UnrepresentableFact) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsUnrepresentableOutput returns the union data inside the UnrepresentableValue as a UnrepresentableOutput
func (t UnrepresentableValue) AsUnrepresentableOutput() (UnrepresentableOutput, error) {
	var body UnrepresentableOutput
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromUnrepresentableOutput overwrites any union data inside the UnrepresentableValue as the provided UnrepresentableOutput
func (t *UnrepresentableValue) FromUnrepresentableOutput(v UnrepresentableOutput) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeUnrepresentableOutput performs a merge with any union data inside the UnrepresentableValue, using the provided UnrepresentableOutput
func (t *UnrepresentableValue) MergeUnrepresentableOutput(v UnrepresentableOutput) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t UnrepresentableValue) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *UnrepresentableValue) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTextValueEncoding returns the union data inside the ValueEncoding as a TextValueEncoding
func (t ValueEncoding) AsTextValueEncoding() (TextValueEncoding, error) {
	var body TextValueEncoding
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTextValueEncoding overwrites any union data inside the ValueEncoding as the provided TextValueEncoding
func (t *ValueEncoding) FromTextValueEncoding(v TextValueEncoding) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTextValueEncoding performs a merge with any union data inside the ValueEncoding, using the provided TextValueEncoding
func (t *ValueEncoding) MergeTextValueEncoding(v TextValueEncoding) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCalendarDateValueEncoding returns the union data inside the ValueEncoding as a CalendarDateValueEncoding
func (t ValueEncoding) AsCalendarDateValueEncoding() (CalendarDateValueEncoding, error) {
	var body CalendarDateValueEncoding
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCalendarDateValueEncoding overwrites any union data inside the ValueEncoding as the provided CalendarDateValueEncoding
func (t *ValueEncoding) FromCalendarDateValueEncoding(v CalendarDateValueEncoding) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCalendarDateValueEncoding performs a merge with any union data inside the ValueEncoding, using the provided CalendarDateValueEncoding
func (t *ValueEncoding) MergeCalendarDateValueEncoding(v CalendarDateValueEncoding) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCompositeValueEncoding returns the union data inside the ValueEncoding as a CompositeValueEncoding
func (t ValueEncoding) AsCompositeValueEncoding() (CompositeValueEncoding, error) {
	var body CompositeValueEncoding
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCompositeValueEncoding overwrites any union data inside the ValueEncoding as the provided CompositeValueEncoding
func (t *ValueEncoding) FromCompositeValueEncoding(v CompositeValueEncoding) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCompositeValueEncoding performs a merge with any union data inside the ValueEncoding, using the provided CompositeValueEncoding
func (t *ValueEncoding) MergeCompositeValueEncoding(v CompositeValueEncoding) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ValueEncoding) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ValueEncoding) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsTokenWindow returns the union data inside the Win as a TokenWindow
func (t Win) AsTokenWindow() (TokenWindow, error) {
	var body TokenWindow
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromTokenWindow overwrites any union data inside the Win as the provided TokenWindow
func (t *Win) FromTokenWindow(v TokenWindow) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeTokenWindow performs a merge with any union data inside the Win, using the provided TokenWindow
func (t *Win) MergeTokenWindow(v TokenWindow) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsCharacterWindow returns the union data inside the Win as a CharacterWindow
func (t Win) AsCharacterWindow() (CharacterWindow, error) {
	var body CharacterWindow
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCharacterWindow overwrites any union data inside the Win as the provided CharacterWindow
func (t *Win) FromCharacterWindow(v CharacterWindow) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCharacterWindow performs a merge with any union data inside the Win, using the provided CharacterWindow
func (t *Win) MergeCharacterWindow(v CharacterWindow) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Win) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Win) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {

	// CtxList List the context ids the caller may use
	//
	// Lists the installed context ids the caller may use. On a managed deployment, a Context token lists the one context it names, and an API-key caller without a token lists the contexts the deployment approves for its organization, when the deployment accepts API keys. A standalone runtime takes no Context token: installation is the grant, and it lists every installed context it can serve now, leaving out one whose latest version does not resolve or whose content the license does not cover. The list is unpaginated. On a managed deployment, returns 404 when authorization is denied.
	//
	// Corresponds with GET /v1/ctx (the `CtxList` operationId).
	CtxList(ctx context.Context, params *CtxListParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxVersions List available versions for a context id
	//
	// Lists every installed version of one context id that the caller may use, highest semver first. On a managed deployment, these are the versions its Context token or API key admits. A standalone runtime takes no Context token: installation is the grant, and it lists every installed version it can serve, leaving out one whose content the license does not cover. Returns 404 when the id is not installed or authorization is denied.
	//
	// Corresponds with GET /v1/ctx/{ctxId} (the `CtxVersions` operationId).
	CtxVersions(ctx context.Context, ctxID string, params *CtxVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxDescribe Describe a resolved context
	//
	// Returns the component breakdown for the resolved context: scheme, ordered layers, ordered guides, optional access policy, and ordered crosswalks, each with the version it declares and the concrete version that answered. The :version segment may be a concrete semver or the literal "latest"; the resolved version is reported in the X-EzClassify-Resolved-Version response header and echoed in the response body.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version} (the `CtxDescribe` operationId).
	CtxDescribe(ctx context.Context, ctxID string, version string, params *CtxDescribeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxAccessCheckWithBody Check access decision in a context
	//
	// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
	CtxAccessCheckWithBody(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxAccessCheck Check access decision in a context
	//
	// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
	CtxAccessCheck(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, body CtxAccessCheckJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxAccessDescribeRequirementsWithBody Describe access requirements for a target in a context
	//
	// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
	CtxAccessDescribeRequirementsWithBody(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxAccessDescribeRequirements Describe access requirements for a target in a context
	//
	// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
	CtxAccessDescribeRequirements(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, body CtxAccessDescribeRequirementsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxAccessRules Describe every access rule in a context
	//
	// Returns the human-readable predicate, the `userAttributes` paths, and the data categories each rule depends on, for every rule in the context's access policy. A caller that only needs the rule ids reads them from the same response. The :version segment may be omitted to default to "latest".
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/access/rules (the `CtxAccessRules` operationId).
	CtxAccessRules(ctx context.Context, ctxID string, version string, params *CtxAccessRulesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxCrosswalkDescribe Describe every crosswalk in a context
	//
	// Returns every crosswalk the context holds, in declaration order: its id and the concrete version that answers, its name, guidance and citation, the source and target schemes as the selectors the context description lists under `schemes`, and the mappings and selections a translation applies, each with its disposition and reason. A context that declares no crosswalk answers with an empty list. The :version segment may be omitted to default to "latest".
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/crosswalk/describe (the `CtxCrosswalkDescribe` operationId).
	CtxCrosswalkDescribe(ctx context.Context, ctxID string, version string, params *CtxCrosswalkDescribeParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxCrosswalkTranslateWithBody Translate a marking through a crosswalk in a context
	//
	// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
	CtxCrosswalkTranslateWithBody(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxCrosswalkTranslate Translate a marking through a crosswalk in a context
	//
	// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
	CtxCrosswalkTranslate(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, body CtxCrosswalkTranslateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxDocumentAuditWithBody Audit a document for marking coverage in a context
	//
	// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
	CtxDocumentAuditWithBody(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxDocumentAudit Audit a document for marking coverage in a context
	//
	// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
	CtxDocumentAudit(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, body CtxDocumentAuditJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideEvaluateWithBody Evaluate selected guide rules in a context
	//
	// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
	CtxGuideEvaluateWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideEvaluate Evaluate selected guide rules in a context
	//
	// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
	CtxGuideEvaluate(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, body CtxGuideEvaluateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideScanWithBody Scan text against context guides
	//
	// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
	CtxGuideScanWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideScan Scan text against context guides
	//
	// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
	CtxGuideScan(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, body CtxGuideScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideViewWithBody Build guide structural view in a context
	//
	// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
	CtxGuideViewWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxGuideView Build guide structural view in a context
	//
	// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
	CtxGuideView(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, body CtxGuideViewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeSummary Get the scheme summary in a context
	//
	// Returns high-level metadata for the context's composed scheme (base plus its declared layers): id, version, name, description, presentation configuration, default render format, and scope label. The scheme identity and layers come from the Context. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme (the `CtxSchemeSummary` operationId).
	CtxSchemeSummary(ctx context.Context, ctxID string, version string, params *CtxSchemeSummaryParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeAggregateWithBody Aggregate marking strings in a context
	//
	// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
	CtxSchemeAggregateWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeAggregate Aggregate marking strings in a context
	//
	// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
	CtxSchemeAggregate(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, body CtxSchemeAggregateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeApplyPresetWithBody Apply a label preset in a context
	//
	// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
	CtxSchemeApplyPresetWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeApplyPreset Apply a label preset in a context
	//
	// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
	CtxSchemeApplyPreset(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, body CtxSchemeApplyPresetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeCategories Get the scheme categories in a context
	//
	// Returns every category of the context's composed scheme as a category view, in declared order: label, scopes, default concept, rollup, and concept sort order, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/categories (the `CtxSchemeCategories` operationId).
	CtxSchemeCategories(ctx context.Context, ctxID string, version string, params *CtxSchemeCategoriesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeCheckCoverageWithBody Check marking coverage in a context
	//
	// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
	CtxSchemeCheckCoverageWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeCheckCoverage Check marking coverage in a context
	//
	// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
	CtxSchemeCheckCoverage(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, body CtxSchemeCheckCoverageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeConcepts Get the scheme concepts in a context
	//
	// Returns every concept of the context's composed scheme as a concept view: id, name, views, categories, parents, position, rank, and display metadata, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/concepts (the `CtxSchemeConcepts` operationId).
	CtxSchemeConcepts(ctx context.Context, ctxID string, version string, params *CtxSchemeConceptsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeDescribeConcept Describe a concept in a context
	//
	// Returns structured description and metadata for the specified concept within a category, scoped to the context's composed scheme. The :version segment may be omitted to default to "latest".
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/describe-concept (the `CtxSchemeDescribeConcept` operationId).
	CtxSchemeDescribeConcept(ctx context.Context, ctxID string, version string, params *CtxSchemeDescribeConceptParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeDiffWithBody Diff two markings in a context
	//
	// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
	CtxSchemeDiffWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeDiff Diff two markings in a context
	//
	// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
	CtxSchemeDiff(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, body CtxSchemeDiffJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeExplainWithBody Explain how a marking is processed in a context
	//
	// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
	CtxSchemeExplainWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeExplain Explain how a marking is processed in a context
	//
	// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
	CtxSchemeExplain(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, body CtxSchemeExplainJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeFormats Get the scheme render formats in a context
	//
	// Lists the formats the context's composed scheme writes as format views: each format's id, the surface it writes, its resolved layout, and how its text is framed, the same shape a local engine's scheme view carries. Not the field trees, which stay server-side. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/formats (the `CtxSchemeFormats` operationId).
	CtxSchemeFormats(ctx context.Context, ctxID string, version string, params *CtxSchemeFormatsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeLabelPresets List the label presets in a context
	//
	// Returns the label presets of the selected scheme as preset views: id, label, description, and scope. Empty when the scheme declares none. The marking string and tag list stay server-side; expand a preset with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset under the same selector. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/label-presets (the `CtxSchemeLabelPresets` operationId).
	CtxSchemeLabelPresets(ctx context.Context, ctxID string, version string, params *CtxSchemeLabelPresetsParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeNoticesWithBody Collect notices for assertions in a context
	//
	// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
	CtxSchemeNoticesWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeNotices Collect notices for assertions in a context
	//
	// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
	CtxSchemeNotices(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, body CtxSchemeNoticesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeProcessWithBody Process a marking in a context
	//
	// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
	CtxSchemeProcessWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeProcess Process a marking in a context
	//
	// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
	CtxSchemeProcess(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, body CtxSchemeProcessJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeRenderWithBody Render assertions in a context
	//
	// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
	CtxSchemeRenderWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeRender Render assertions in a context
	//
	// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
	CtxSchemeRender(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, body CtxSchemeRenderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeScanWithBody Scan text for markings in a context
	//
	// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
	CtxSchemeScanWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeScan Scan text for markings in a context
	//
	// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
	CtxSchemeScan(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, body CtxSchemeScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeScopes Get the scheme scopes in a context
	//
	// Returns the scope definitions for the context's composed scheme in declaration order. Empty when it declares none. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/scopes (the `CtxSchemeScopes` operationId).
	CtxSchemeScopes(ctx context.Context, ctxID string, version string, params *CtxSchemeScopesParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeSuggestWithBody Suggest marking completions in a context
	//
	// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
	CtxSchemeSuggestWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeSuggest Suggest marking completions in a context
	//
	// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
	CtxSchemeSuggest(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, body CtxSchemeSuggestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeSuggestForCategoriesWithBody Suggest completions for several categories in a context
	//
	// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
	CtxSchemeSuggestForCategoriesWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeSuggestForCategories Suggest completions for several categories in a context
	//
	// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
	CtxSchemeSuggestForCategories(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, body CtxSchemeSuggestForCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeValidateWithBody Validate a marking in a context
	//
	// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
	CtxSchemeValidateWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeValidate Validate a marking in a context
	//
	// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
	CtxSchemeValidate(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, body CtxSchemeValidateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeViewWithBody Build scheme structural view in a context
	//
	// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
	CtxSchemeViewWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CtxSchemeView Build scheme structural view in a context
	//
	// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type.
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
	CtxSchemeView(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, body CtxSchemeViewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetHealth Health check
	//
	// Returns service liveness status and the engine version string.
	//
	// Corresponds with GET /v1/health (the `GetHealth` operationId).
	GetHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetReady Readiness check
	//
	// Reports whether the shared runtime host admitted the configured license and all explicitly required local artifacts are available. Expiration never makes a replica unready.
	//
	// Corresponds with GET /v1/ready (the `GetReady` operationId).
	GetReady(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)
}

// CtxList List the context ids the caller may use
//
// Lists the installed context ids the caller may use. On a managed deployment, a Context token lists the one context it names, and an API-key caller without a token lists the contexts the deployment approves for its organization, when the deployment accepts API keys. A standalone runtime takes no Context token: installation is the grant, and it lists every installed context it can serve now, leaving out one whose latest version does not resolve or whose content the license does not cover. The list is unpaginated. On a managed deployment, returns 404 when authorization is denied.
//
// Corresponds with GET /v1/ctx (the `CtxList` operationId).
func (c *Client) CtxList(ctx context.Context, params *CtxListParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxListRequest(c.Server, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxVersions List available versions for a context id
//
// Lists every installed version of one context id that the caller may use, highest semver first. On a managed deployment, these are the versions its Context token or API key admits. A standalone runtime takes no Context token: installation is the grant, and it lists every installed version it can serve, leaving out one whose content the license does not cover. Returns 404 when the id is not installed or authorization is denied.
//
// Corresponds with GET /v1/ctx/{ctxId} (the `CtxVersions` operationId).
func (c *Client) CtxVersions(ctx context.Context, ctxID string, params *CtxVersionsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxVersionsRequest(c.Server, ctxID, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxDescribe Describe a resolved context
//
// Returns the component breakdown for the resolved context: scheme, ordered layers, ordered guides, optional access policy, and ordered crosswalks, each with the version it declares and the concrete version that answered. The :version segment may be a concrete semver or the literal "latest"; the resolved version is reported in the X-EzClassify-Resolved-Version response header and echoed in the response body.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version} (the `CtxDescribe` operationId).
func (c *Client) CtxDescribe(ctx context.Context, ctxID string, version string, params *CtxDescribeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxDescribeRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxAccessCheckWithBody Check access decision in a context
//
// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
func (c *Client) CtxAccessCheckWithBody(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxAccessCheckRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxAccessCheck Check access decision in a context
//
// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
func (c *Client) CtxAccessCheck(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, body CtxAccessCheckJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxAccessCheckRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxAccessDescribeRequirementsWithBody Describe access requirements for a target in a context
//
// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
func (c *Client) CtxAccessDescribeRequirementsWithBody(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxAccessDescribeRequirementsRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxAccessDescribeRequirements Describe access requirements for a target in a context
//
// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
func (c *Client) CtxAccessDescribeRequirements(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, body CtxAccessDescribeRequirementsJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxAccessDescribeRequirementsRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxAccessRules Describe every access rule in a context
//
// Returns the human-readable predicate, the `userAttributes` paths, and the data categories each rule depends on, for every rule in the context's access policy. A caller that only needs the rule ids reads them from the same response. The :version segment may be omitted to default to "latest".
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/access/rules (the `CtxAccessRules` operationId).
func (c *Client) CtxAccessRules(ctx context.Context, ctxID string, version string, params *CtxAccessRulesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxAccessRulesRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxCrosswalkDescribe Describe every crosswalk in a context
//
// Returns every crosswalk the context holds, in declaration order: its id and the concrete version that answers, its name, guidance and citation, the source and target schemes as the selectors the context description lists under `schemes`, and the mappings and selections a translation applies, each with its disposition and reason. A context that declares no crosswalk answers with an empty list. The :version segment may be omitted to default to "latest".
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/crosswalk/describe (the `CtxCrosswalkDescribe` operationId).
func (c *Client) CtxCrosswalkDescribe(ctx context.Context, ctxID string, version string, params *CtxCrosswalkDescribeParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxCrosswalkDescribeRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxCrosswalkTranslateWithBody Translate a marking through a crosswalk in a context
//
// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
func (c *Client) CtxCrosswalkTranslateWithBody(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxCrosswalkTranslateRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxCrosswalkTranslate Translate a marking through a crosswalk in a context
//
// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
func (c *Client) CtxCrosswalkTranslate(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, body CtxCrosswalkTranslateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxCrosswalkTranslateRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxDocumentAuditWithBody Audit a document for marking coverage in a context
//
// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
func (c *Client) CtxDocumentAuditWithBody(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxDocumentAuditRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxDocumentAudit Audit a document for marking coverage in a context
//
// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
func (c *Client) CtxDocumentAudit(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, body CtxDocumentAuditJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxDocumentAuditRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideEvaluateWithBody Evaluate selected guide rules in a context
//
// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
func (c *Client) CtxGuideEvaluateWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideEvaluateRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideEvaluate Evaluate selected guide rules in a context
//
// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
func (c *Client) CtxGuideEvaluate(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, body CtxGuideEvaluateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideEvaluateRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideScanWithBody Scan text against context guides
//
// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
func (c *Client) CtxGuideScanWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideScanRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideScan Scan text against context guides
//
// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
func (c *Client) CtxGuideScan(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, body CtxGuideScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideScanRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideViewWithBody Build guide structural view in a context
//
// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
func (c *Client) CtxGuideViewWithBody(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideViewRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxGuideView Build guide structural view in a context
//
// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
func (c *Client) CtxGuideView(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, body CtxGuideViewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxGuideViewRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeSummary Get the scheme summary in a context
//
// Returns high-level metadata for the context's composed scheme (base plus its declared layers): id, version, name, description, presentation configuration, default render format, and scope label. The scheme identity and layers come from the Context. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme (the `CtxSchemeSummary` operationId).
func (c *Client) CtxSchemeSummary(ctx context.Context, ctxID string, version string, params *CtxSchemeSummaryParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeSummaryRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeAggregateWithBody Aggregate marking strings in a context
//
// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
func (c *Client) CtxSchemeAggregateWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeAggregateRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeAggregate Aggregate marking strings in a context
//
// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
func (c *Client) CtxSchemeAggregate(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, body CtxSchemeAggregateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeAggregateRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeApplyPresetWithBody Apply a label preset in a context
//
// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
func (c *Client) CtxSchemeApplyPresetWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeApplyPresetRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeApplyPreset Apply a label preset in a context
//
// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
func (c *Client) CtxSchemeApplyPreset(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, body CtxSchemeApplyPresetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeApplyPresetRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeCategories Get the scheme categories in a context
//
// Returns every category of the context's composed scheme as a category view, in declared order: label, scopes, default concept, rollup, and concept sort order, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/categories (the `CtxSchemeCategories` operationId).
func (c *Client) CtxSchemeCategories(ctx context.Context, ctxID string, version string, params *CtxSchemeCategoriesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeCategoriesRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeCheckCoverageWithBody Check marking coverage in a context
//
// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
func (c *Client) CtxSchemeCheckCoverageWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeCheckCoverageRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeCheckCoverage Check marking coverage in a context
//
// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
func (c *Client) CtxSchemeCheckCoverage(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, body CtxSchemeCheckCoverageJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeCheckCoverageRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeConcepts Get the scheme concepts in a context
//
// Returns every concept of the context's composed scheme as a concept view: id, name, views, categories, parents, position, rank, and display metadata, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/concepts (the `CtxSchemeConcepts` operationId).
func (c *Client) CtxSchemeConcepts(ctx context.Context, ctxID string, version string, params *CtxSchemeConceptsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeConceptsRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeDescribeConcept Describe a concept in a context
//
// Returns structured description and metadata for the specified concept within a category, scoped to the context's composed scheme. The :version segment may be omitted to default to "latest".
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/describe-concept (the `CtxSchemeDescribeConcept` operationId).
func (c *Client) CtxSchemeDescribeConcept(ctx context.Context, ctxID string, version string, params *CtxSchemeDescribeConceptParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeDescribeConceptRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeDiffWithBody Diff two markings in a context
//
// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
func (c *Client) CtxSchemeDiffWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeDiffRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeDiff Diff two markings in a context
//
// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
func (c *Client) CtxSchemeDiff(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, body CtxSchemeDiffJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeDiffRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeExplainWithBody Explain how a marking is processed in a context
//
// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
func (c *Client) CtxSchemeExplainWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeExplainRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeExplain Explain how a marking is processed in a context
//
// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
func (c *Client) CtxSchemeExplain(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, body CtxSchemeExplainJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeExplainRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeFormats Get the scheme render formats in a context
//
// Lists the formats the context's composed scheme writes as format views: each format's id, the surface it writes, its resolved layout, and how its text is framed, the same shape a local engine's scheme view carries. Not the field trees, which stay server-side. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/formats (the `CtxSchemeFormats` operationId).
func (c *Client) CtxSchemeFormats(ctx context.Context, ctxID string, version string, params *CtxSchemeFormatsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeFormatsRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeLabelPresets List the label presets in a context
//
// Returns the label presets of the selected scheme as preset views: id, label, description, and scope. Empty when the scheme declares none. The marking string and tag list stay server-side; expand a preset with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset under the same selector. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/label-presets (the `CtxSchemeLabelPresets` operationId).
func (c *Client) CtxSchemeLabelPresets(ctx context.Context, ctxID string, version string, params *CtxSchemeLabelPresetsParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeLabelPresetsRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeNoticesWithBody Collect notices for assertions in a context
//
// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
func (c *Client) CtxSchemeNoticesWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeNoticesRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeNotices Collect notices for assertions in a context
//
// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
func (c *Client) CtxSchemeNotices(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, body CtxSchemeNoticesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeNoticesRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeProcessWithBody Process a marking in a context
//
// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
func (c *Client) CtxSchemeProcessWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeProcessRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeProcess Process a marking in a context
//
// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
func (c *Client) CtxSchemeProcess(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, body CtxSchemeProcessJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeProcessRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeRenderWithBody Render assertions in a context
//
// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
func (c *Client) CtxSchemeRenderWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeRenderRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeRender Render assertions in a context
//
// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
func (c *Client) CtxSchemeRender(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, body CtxSchemeRenderJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeRenderRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeScanWithBody Scan text for markings in a context
//
// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
func (c *Client) CtxSchemeScanWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeScanRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeScan Scan text for markings in a context
//
// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
func (c *Client) CtxSchemeScan(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, body CtxSchemeScanJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeScanRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeScopes Get the scheme scopes in a context
//
// Returns the scope definitions for the context's composed scheme in declaration order. Empty when it declares none. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/scopes (the `CtxSchemeScopes` operationId).
func (c *Client) CtxSchemeScopes(ctx context.Context, ctxID string, version string, params *CtxSchemeScopesParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeScopesRequest(c.Server, ctxID, version, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeSuggestWithBody Suggest marking completions in a context
//
// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
func (c *Client) CtxSchemeSuggestWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeSuggestRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeSuggest Suggest marking completions in a context
//
// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
func (c *Client) CtxSchemeSuggest(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, body CtxSchemeSuggestJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeSuggestRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeSuggestForCategoriesWithBody Suggest completions for several categories in a context
//
// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
func (c *Client) CtxSchemeSuggestForCategoriesWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeSuggestForCategoriesRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeSuggestForCategories Suggest completions for several categories in a context
//
// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
func (c *Client) CtxSchemeSuggestForCategories(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, body CtxSchemeSuggestForCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeSuggestForCategoriesRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeValidateWithBody Validate a marking in a context
//
// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
func (c *Client) CtxSchemeValidateWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeValidateRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeValidate Validate a marking in a context
//
// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
func (c *Client) CtxSchemeValidate(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, body CtxSchemeValidateJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeValidateRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeViewWithBody Build scheme structural view in a context
//
// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
func (c *Client) CtxSchemeViewWithBody(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeViewRequestWithBody(c.Server, ctxID, version, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// CtxSchemeView Build scheme structural view in a context
//
// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type.
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
func (c *Client) CtxSchemeView(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, body CtxSchemeViewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCtxSchemeViewRequest(c.Server, ctxID, version, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetHealth Health check
//
// Returns service liveness status and the engine version string.
//
// Corresponds with GET /v1/health (the `GetHealth` operationId).
func (c *Client) GetHealth(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetHealthRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// GetReady Readiness check
//
// Reports whether the shared runtime host admitted the configured license and all explicitly required local artifacts are available. Expiration never makes a replica unready.
//
// Corresponds with GET /v1/ready (the `GetReady` operationId).
func (c *Client) GetReady(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetReadyRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCtxListRequest constructs an http.Request for the CtxList method
func NewCtxListRequest(server string, params *CtxListParams) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxVersionsRequest constructs an http.Request for the CtxVersions method
func NewCtxVersionsRequest(server string, ctxID string, params *CtxVersionsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxDescribeRequest constructs an http.Request for the CtxDescribe method
func NewCtxDescribeRequest(server string, ctxID string, version string, params *CtxDescribeParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxAccessCheckRequest calls the generic CtxAccessCheck builder with application/json body
func NewCtxAccessCheckRequest(server string, ctxID string, version string, params *CtxAccessCheckParams, body CtxAccessCheckJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxAccessCheckRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxAccessCheckRequestWithBody constructs an http.Request for the CtxAccessCheck method, with any body, and a specified content type
func NewCtxAccessCheckRequestWithBody(server string, ctxID string, version string, params *CtxAccessCheckParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/access/check", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxAccessDescribeRequirementsRequest calls the generic CtxAccessDescribeRequirements builder with application/json body
func NewCtxAccessDescribeRequirementsRequest(server string, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, body CtxAccessDescribeRequirementsJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxAccessDescribeRequirementsRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxAccessDescribeRequirementsRequestWithBody constructs an http.Request for the CtxAccessDescribeRequirements method, with any body, and a specified content type
func NewCtxAccessDescribeRequirementsRequestWithBody(server string, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/access/describe-requirements", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxAccessRulesRequest constructs an http.Request for the CtxAccessRules method
func NewCtxAccessRulesRequest(server string, ctxID string, version string, params *CtxAccessRulesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/access/rules", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxCrosswalkDescribeRequest constructs an http.Request for the CtxCrosswalkDescribe method
func NewCtxCrosswalkDescribeRequest(server string, ctxID string, version string, params *CtxCrosswalkDescribeParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/crosswalk/describe", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxCrosswalkTranslateRequest calls the generic CtxCrosswalkTranslate builder with application/json body
func NewCtxCrosswalkTranslateRequest(server string, ctxID string, version string, params *CtxCrosswalkTranslateParams, body CtxCrosswalkTranslateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxCrosswalkTranslateRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxCrosswalkTranslateRequestWithBody constructs an http.Request for the CtxCrosswalkTranslate method, with any body, and a specified content type
func NewCtxCrosswalkTranslateRequestWithBody(server string, ctxID string, version string, params *CtxCrosswalkTranslateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/crosswalk/translate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxDocumentAuditRequest calls the generic CtxDocumentAudit builder with application/json body
func NewCtxDocumentAuditRequest(server string, ctxID string, version string, params *CtxDocumentAuditParams, body CtxDocumentAuditJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxDocumentAuditRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxDocumentAuditRequestWithBody constructs an http.Request for the CtxDocumentAudit method, with any body, and a specified content type
func NewCtxDocumentAuditRequestWithBody(server string, ctxID string, version string, params *CtxDocumentAuditParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/document/audit", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxGuideEvaluateRequest calls the generic CtxGuideEvaluate builder with application/json body
func NewCtxGuideEvaluateRequest(server string, ctxID string, version string, params *CtxGuideEvaluateParams, body CtxGuideEvaluateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxGuideEvaluateRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxGuideEvaluateRequestWithBody constructs an http.Request for the CtxGuideEvaluate method, with any body, and a specified content type
func NewCtxGuideEvaluateRequestWithBody(server string, ctxID string, version string, params *CtxGuideEvaluateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/guide/evaluate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxGuideScanRequest calls the generic CtxGuideScan builder with application/json body
func NewCtxGuideScanRequest(server string, ctxID string, version string, params *CtxGuideScanParams, body CtxGuideScanJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxGuideScanRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxGuideScanRequestWithBody constructs an http.Request for the CtxGuideScan method, with any body, and a specified content type
func NewCtxGuideScanRequestWithBody(server string, ctxID string, version string, params *CtxGuideScanParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/guide/scan", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxGuideViewRequest calls the generic CtxGuideView builder with application/json body
func NewCtxGuideViewRequest(server string, ctxID string, version string, params *CtxGuideViewParams, body CtxGuideViewJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxGuideViewRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxGuideViewRequestWithBody constructs an http.Request for the CtxGuideView method, with any body, and a specified content type
func NewCtxGuideViewRequestWithBody(server string, ctxID string, version string, params *CtxGuideViewParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/guide/view", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeSummaryRequest constructs an http.Request for the CtxSchemeSummary method
func NewCtxSchemeSummaryRequest(server string, ctxID string, version string, params *CtxSchemeSummaryParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeAggregateRequest calls the generic CtxSchemeAggregate builder with application/json body
func NewCtxSchemeAggregateRequest(server string, ctxID string, version string, params *CtxSchemeAggregateParams, body CtxSchemeAggregateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeAggregateRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeAggregateRequestWithBody constructs an http.Request for the CtxSchemeAggregate method, with any body, and a specified content type
func NewCtxSchemeAggregateRequestWithBody(server string, ctxID string, version string, params *CtxSchemeAggregateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/aggregate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeApplyPresetRequest calls the generic CtxSchemeApplyPreset builder with application/json body
func NewCtxSchemeApplyPresetRequest(server string, ctxID string, version string, params *CtxSchemeApplyPresetParams, body CtxSchemeApplyPresetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeApplyPresetRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeApplyPresetRequestWithBody constructs an http.Request for the CtxSchemeApplyPreset method, with any body, and a specified content type
func NewCtxSchemeApplyPresetRequestWithBody(server string, ctxID string, version string, params *CtxSchemeApplyPresetParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/apply-preset", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeCategoriesRequest constructs an http.Request for the CtxSchemeCategories method
func NewCtxSchemeCategoriesRequest(server string, ctxID string, version string, params *CtxSchemeCategoriesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/categories", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeCheckCoverageRequest calls the generic CtxSchemeCheckCoverage builder with application/json body
func NewCtxSchemeCheckCoverageRequest(server string, ctxID string, version string, params *CtxSchemeCheckCoverageParams, body CtxSchemeCheckCoverageJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeCheckCoverageRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeCheckCoverageRequestWithBody constructs an http.Request for the CtxSchemeCheckCoverage method, with any body, and a specified content type
func NewCtxSchemeCheckCoverageRequestWithBody(server string, ctxID string, version string, params *CtxSchemeCheckCoverageParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/check-coverage", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeConceptsRequest constructs an http.Request for the CtxSchemeConcepts method
func NewCtxSchemeConceptsRequest(server string, ctxID string, version string, params *CtxSchemeConceptsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/concepts", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeDescribeConceptRequest constructs an http.Request for the CtxSchemeDescribeConcept method
func NewCtxSchemeDescribeConceptRequest(server string, ctxID string, version string, params *CtxSchemeDescribeConceptParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/describe-concept", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "conceptId", params.ConceptID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if queryFrag, err := runtime.StyleParamWithOptions("form", true, "categoryId", params.CategoryID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
			return nil, err
		} else {
			for _, qp := range strings.Split(queryFrag, "&") {
				rawQueryFragments = append(rawQueryFragments, qp)
			}
		}

		if params.View != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "view", *params.View, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeDiffRequest calls the generic CtxSchemeDiff builder with application/json body
func NewCtxSchemeDiffRequest(server string, ctxID string, version string, params *CtxSchemeDiffParams, body CtxSchemeDiffJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeDiffRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeDiffRequestWithBody constructs an http.Request for the CtxSchemeDiff method, with any body, and a specified content type
func NewCtxSchemeDiffRequestWithBody(server string, ctxID string, version string, params *CtxSchemeDiffParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/diff", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeExplainRequest calls the generic CtxSchemeExplain builder with application/json body
func NewCtxSchemeExplainRequest(server string, ctxID string, version string, params *CtxSchemeExplainParams, body CtxSchemeExplainJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeExplainRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeExplainRequestWithBody constructs an http.Request for the CtxSchemeExplain method, with any body, and a specified content type
func NewCtxSchemeExplainRequestWithBody(server string, ctxID string, version string, params *CtxSchemeExplainParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/explain", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeFormatsRequest constructs an http.Request for the CtxSchemeFormats method
func NewCtxSchemeFormatsRequest(server string, ctxID string, version string, params *CtxSchemeFormatsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/formats", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeLabelPresetsRequest constructs an http.Request for the CtxSchemeLabelPresets method
func NewCtxSchemeLabelPresetsRequest(server string, ctxID string, version string, params *CtxSchemeLabelPresetsParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/label-presets", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeNoticesRequest calls the generic CtxSchemeNotices builder with application/json body
func NewCtxSchemeNoticesRequest(server string, ctxID string, version string, params *CtxSchemeNoticesParams, body CtxSchemeNoticesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeNoticesRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeNoticesRequestWithBody constructs an http.Request for the CtxSchemeNotices method, with any body, and a specified content type
func NewCtxSchemeNoticesRequestWithBody(server string, ctxID string, version string, params *CtxSchemeNoticesParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/notices", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeProcessRequest calls the generic CtxSchemeProcess builder with application/json body
func NewCtxSchemeProcessRequest(server string, ctxID string, version string, params *CtxSchemeProcessParams, body CtxSchemeProcessJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeProcessRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeProcessRequestWithBody constructs an http.Request for the CtxSchemeProcess method, with any body, and a specified content type
func NewCtxSchemeProcessRequestWithBody(server string, ctxID string, version string, params *CtxSchemeProcessParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/process", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeRenderRequest calls the generic CtxSchemeRender builder with application/json body
func NewCtxSchemeRenderRequest(server string, ctxID string, version string, params *CtxSchemeRenderParams, body CtxSchemeRenderJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeRenderRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeRenderRequestWithBody constructs an http.Request for the CtxSchemeRender method, with any body, and a specified content type
func NewCtxSchemeRenderRequestWithBody(server string, ctxID string, version string, params *CtxSchemeRenderParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/render", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeScanRequest calls the generic CtxSchemeScan builder with application/json body
func NewCtxSchemeScanRequest(server string, ctxID string, version string, params *CtxSchemeScanParams, body CtxSchemeScanJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeScanRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeScanRequestWithBody constructs an http.Request for the CtxSchemeScan method, with any body, and a specified content type
func NewCtxSchemeScanRequestWithBody(server string, ctxID string, version string, params *CtxSchemeScanParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/scan", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeScopesRequest constructs an http.Request for the CtxSchemeScopes method
func NewCtxSchemeScopesRequest(server string, ctxID string, version string, params *CtxSchemeScopesParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/scopes", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeSuggestRequest calls the generic CtxSchemeSuggest builder with application/json body
func NewCtxSchemeSuggestRequest(server string, ctxID string, version string, params *CtxSchemeSuggestParams, body CtxSchemeSuggestJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeSuggestRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeSuggestRequestWithBody constructs an http.Request for the CtxSchemeSuggest method, with any body, and a specified content type
func NewCtxSchemeSuggestRequestWithBody(server string, ctxID string, version string, params *CtxSchemeSuggestParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/suggest", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeSuggestForCategoriesRequest calls the generic CtxSchemeSuggestForCategories builder with application/json body
func NewCtxSchemeSuggestForCategoriesRequest(server string, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, body CtxSchemeSuggestForCategoriesJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeSuggestForCategoriesRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeSuggestForCategoriesRequestWithBody constructs an http.Request for the CtxSchemeSuggestForCategories method, with any body, and a specified content type
func NewCtxSchemeSuggestForCategoriesRequestWithBody(server string, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/suggest-for-categories", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeValidateRequest calls the generic CtxSchemeValidate builder with application/json body
func NewCtxSchemeValidateRequest(server string, ctxID string, version string, params *CtxSchemeValidateParams, body CtxSchemeValidateJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeValidateRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeValidateRequestWithBody constructs an http.Request for the CtxSchemeValidate method, with any body, and a specified content type
func NewCtxSchemeValidateRequestWithBody(server string, ctxID string, version string, params *CtxSchemeValidateParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/validate", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewCtxSchemeViewRequest calls the generic CtxSchemeView builder with application/json body
func NewCtxSchemeViewRequest(server string, ctxID string, version string, params *CtxSchemeViewParams, body CtxSchemeViewJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCtxSchemeViewRequestWithBody(server, ctxID, version, params, "application/json", bodyReader)
}

// NewCtxSchemeViewRequestWithBody constructs an http.Request for the CtxSchemeView method, with any body, and a specified content type
func NewCtxSchemeViewRequestWithBody(server string, ctxID string, version string, params *CtxSchemeViewParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithOptions("simple", false, "ctxId", ctxID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithOptions("simple", false, "version", version, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationPath, Type: "string", Format: ""})
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ctx/%s/%s/scheme/view", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		// queryValues collects non-styled parameters (passthrough, JSON)
		// that are safe to round-trip through url.Values.Encode().
		queryValues := queryURL.Query()
		// rawQueryFragments collects pre-encoded query fragments from
		// styled parameters, preserving literal commas as delimiters
		// per the OpenAPI spec (e.g. "color=blue,black,brown").
		var rawQueryFragments []string

		if params.Scheme != nil {

			if queryFrag, err := runtime.StyleParamWithOptions("form", true, "scheme", *params.Scheme, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationQuery, Type: "string", Format: ""}); err != nil {
				return nil, err
			} else {
				for _, qp := range strings.Split(queryFrag, "&") {
					rawQueryFragments = append(rawQueryFragments, qp)
				}
			}

		}

		if encoded := queryValues.Encode(); encoded != "" {
			rawQueryFragments = append(rawQueryFragments, encoded)
		}
		queryURL.RawQuery = strings.Join(rawQueryFragments, "&")
	}

	req, err := http.NewRequest(http.MethodPost, queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		if params.XEzclassifyApplicationID != nil {
			var headerParam0 string

			headerParam0, err = runtime.StyleParamWithOptions("simple", false, "x-ezclassify-application-id", *params.XEzclassifyApplicationID, runtime.StyleParamOptions{ParamLocation: runtime.ParamLocationHeader, Type: "string", Format: ""})
			if err != nil {
				return nil, err
			}

			req.Header.Set("x-ezclassify-application-id", headerParam0)
		}

	}

	return req, nil
}

// NewGetHealthRequest constructs an http.Request for the GetHealth method
func NewGetHealthRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/health")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetReadyRequest constructs an http.Request for the GetReady method
func NewGetReadyRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/v1/ready")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {

	// CtxListWithResponse List the context ids the caller may use
	//
	// Lists the installed context ids the caller may use. On a managed deployment, a Context token lists the one context it names, and an API-key caller without a token lists the contexts the deployment approves for its organization, when the deployment accepts API keys. A standalone runtime takes no Context token: installation is the grant, and it lists every installed context it can serve now, leaving out one whose latest version does not resolve or whose content the license does not cover. The list is unpaginated. On a managed deployment, returns 404 when authorization is denied.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx (the `CtxList` operationId).
	CtxListWithResponse(ctx context.Context, params *CtxListParams, reqEditors ...RequestEditorFn) (*CtxListHTTP, error)

	// CtxVersionsWithResponse List available versions for a context id
	//
	// Lists every installed version of one context id that the caller may use, highest semver first. On a managed deployment, these are the versions its Context token or API key admits. A standalone runtime takes no Context token: installation is the grant, and it lists every installed version it can serve, leaving out one whose content the license does not cover. Returns 404 when the id is not installed or authorization is denied.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId} (the `CtxVersions` operationId).
	CtxVersionsWithResponse(ctx context.Context, ctxID string, params *CtxVersionsParams, reqEditors ...RequestEditorFn) (*CtxVersionsHTTP, error)

	// CtxDescribeWithResponse Describe a resolved context
	//
	// Returns the component breakdown for the resolved context: scheme, ordered layers, ordered guides, optional access policy, and ordered crosswalks, each with the version it declares and the concrete version that answered. The :version segment may be a concrete semver or the literal "latest"; the resolved version is reported in the X-EzClassify-Resolved-Version response header and echoed in the response body.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version} (the `CtxDescribe` operationId).
	CtxDescribeWithResponse(ctx context.Context, ctxID string, version string, params *CtxDescribeParams, reqEditors ...RequestEditorFn) (*CtxDescribeHTTP, error)

	// CtxAccessCheckWithBodyWithResponse Check access decision in a context
	//
	// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
	CtxAccessCheckWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxAccessCheckHTTP, error)

	// CtxAccessCheckWithResponse Check access decision in a context
	//
	// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
	CtxAccessCheckWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, body CtxAccessCheckJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxAccessCheckHTTP, error)

	// CtxAccessDescribeRequirementsWithBodyWithResponse Describe access requirements for a target in a context
	//
	// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
	CtxAccessDescribeRequirementsWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxAccessDescribeRequirementsHTTP, error)

	// CtxAccessDescribeRequirementsWithResponse Describe access requirements for a target in a context
	//
	// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
	CtxAccessDescribeRequirementsWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, body CtxAccessDescribeRequirementsJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxAccessDescribeRequirementsHTTP, error)

	// CtxAccessRulesWithResponse Describe every access rule in a context
	//
	// Returns the human-readable predicate, the `userAttributes` paths, and the data categories each rule depends on, for every rule in the context's access policy. A caller that only needs the rule ids reads them from the same response. The :version segment may be omitted to default to "latest".
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/access/rules (the `CtxAccessRules` operationId).
	CtxAccessRulesWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessRulesParams, reqEditors ...RequestEditorFn) (*CtxAccessRulesHTTP, error)

	// CtxCrosswalkDescribeWithResponse Describe every crosswalk in a context
	//
	// Returns every crosswalk the context holds, in declaration order: its id and the concrete version that answers, its name, guidance and citation, the source and target schemes as the selectors the context description lists under `schemes`, and the mappings and selections a translation applies, each with its disposition and reason. A context that declares no crosswalk answers with an empty list. The :version segment may be omitted to default to "latest".
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/crosswalk/describe (the `CtxCrosswalkDescribe` operationId).
	CtxCrosswalkDescribeWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkDescribeParams, reqEditors ...RequestEditorFn) (*CtxCrosswalkDescribeHTTP, error)

	// CtxCrosswalkTranslateWithBodyWithResponse Translate a marking through a crosswalk in a context
	//
	// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
	CtxCrosswalkTranslateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxCrosswalkTranslateHTTP, error)

	// CtxCrosswalkTranslateWithResponse Translate a marking through a crosswalk in a context
	//
	// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
	CtxCrosswalkTranslateWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, body CtxCrosswalkTranslateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxCrosswalkTranslateHTTP, error)

	// CtxDocumentAuditWithBodyWithResponse Audit a document for marking coverage in a context
	//
	// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
	CtxDocumentAuditWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxDocumentAuditHTTP, error)

	// CtxDocumentAuditWithResponse Audit a document for marking coverage in a context
	//
	// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
	CtxDocumentAuditWithResponse(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, body CtxDocumentAuditJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxDocumentAuditHTTP, error)

	// CtxGuideEvaluateWithBodyWithResponse Evaluate selected guide rules in a context
	//
	// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
	CtxGuideEvaluateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideEvaluateHTTP, error)

	// CtxGuideEvaluateWithResponse Evaluate selected guide rules in a context
	//
	// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
	CtxGuideEvaluateWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, body CtxGuideEvaluateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideEvaluateHTTP, error)

	// CtxGuideScanWithBodyWithResponse Scan text against context guides
	//
	// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
	CtxGuideScanWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideScanHTTP, error)

	// CtxGuideScanWithResponse Scan text against context guides
	//
	// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
	CtxGuideScanWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, body CtxGuideScanJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideScanHTTP, error)

	// CtxGuideViewWithBodyWithResponse Build guide structural view in a context
	//
	// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
	CtxGuideViewWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideViewHTTP, error)

	// CtxGuideViewWithResponse Build guide structural view in a context
	//
	// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
	CtxGuideViewWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, body CtxGuideViewJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideViewHTTP, error)

	// CtxSchemeSummaryWithResponse Get the scheme summary in a context
	//
	// Returns high-level metadata for the context's composed scheme (base plus its declared layers): id, version, name, description, presentation configuration, default render format, and scope label. The scheme identity and layers come from the Context. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme (the `CtxSchemeSummary` operationId).
	CtxSchemeSummaryWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSummaryParams, reqEditors ...RequestEditorFn) (*CtxSchemeSummaryHTTP, error)

	// CtxSchemeAggregateWithBodyWithResponse Aggregate marking strings in a context
	//
	// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
	CtxSchemeAggregateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeAggregateHTTP, error)

	// CtxSchemeAggregateWithResponse Aggregate marking strings in a context
	//
	// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
	CtxSchemeAggregateWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, body CtxSchemeAggregateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeAggregateHTTP, error)

	// CtxSchemeApplyPresetWithBodyWithResponse Apply a label preset in a context
	//
	// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
	CtxSchemeApplyPresetWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeApplyPresetHTTP, error)

	// CtxSchemeApplyPresetWithResponse Apply a label preset in a context
	//
	// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
	CtxSchemeApplyPresetWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, body CtxSchemeApplyPresetJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeApplyPresetHTTP, error)

	// CtxSchemeCategoriesWithResponse Get the scheme categories in a context
	//
	// Returns every category of the context's composed scheme as a category view, in declared order: label, scopes, default concept, rollup, and concept sort order, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/categories (the `CtxSchemeCategories` operationId).
	CtxSchemeCategoriesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCategoriesParams, reqEditors ...RequestEditorFn) (*CtxSchemeCategoriesHTTP, error)

	// CtxSchemeCheckCoverageWithBodyWithResponse Check marking coverage in a context
	//
	// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
	CtxSchemeCheckCoverageWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeCheckCoverageHTTP, error)

	// CtxSchemeCheckCoverageWithResponse Check marking coverage in a context
	//
	// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
	CtxSchemeCheckCoverageWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, body CtxSchemeCheckCoverageJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeCheckCoverageHTTP, error)

	// CtxSchemeConceptsWithResponse Get the scheme concepts in a context
	//
	// Returns every concept of the context's composed scheme as a concept view: id, name, views, categories, parents, position, rank, and display metadata, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/concepts (the `CtxSchemeConcepts` operationId).
	CtxSchemeConceptsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeConceptsParams, reqEditors ...RequestEditorFn) (*CtxSchemeConceptsHTTP, error)

	// CtxSchemeDescribeConceptWithResponse Describe a concept in a context
	//
	// Returns structured description and metadata for the specified concept within a category, scoped to the context's composed scheme. The :version segment may be omitted to default to "latest".
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/describe-concept (the `CtxSchemeDescribeConcept` operationId).
	CtxSchemeDescribeConceptWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDescribeConceptParams, reqEditors ...RequestEditorFn) (*CtxSchemeDescribeConceptHTTP, error)

	// CtxSchemeDiffWithBodyWithResponse Diff two markings in a context
	//
	// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
	CtxSchemeDiffWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeDiffHTTP, error)

	// CtxSchemeDiffWithResponse Diff two markings in a context
	//
	// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
	CtxSchemeDiffWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, body CtxSchemeDiffJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeDiffHTTP, error)

	// CtxSchemeExplainWithBodyWithResponse Explain how a marking is processed in a context
	//
	// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
	CtxSchemeExplainWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeExplainHTTP, error)

	// CtxSchemeExplainWithResponse Explain how a marking is processed in a context
	//
	// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
	CtxSchemeExplainWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, body CtxSchemeExplainJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeExplainHTTP, error)

	// CtxSchemeFormatsWithResponse Get the scheme render formats in a context
	//
	// Lists the formats the context's composed scheme writes as format views: each format's id, the surface it writes, its resolved layout, and how its text is framed, the same shape a local engine's scheme view carries. Not the field trees, which stay server-side. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/formats (the `CtxSchemeFormats` operationId).
	CtxSchemeFormatsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeFormatsParams, reqEditors ...RequestEditorFn) (*CtxSchemeFormatsHTTP, error)

	// CtxSchemeLabelPresetsWithResponse List the label presets in a context
	//
	// Returns the label presets of the selected scheme as preset views: id, label, description, and scope. Empty when the scheme declares none. The marking string and tag list stay server-side; expand a preset with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset under the same selector. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/label-presets (the `CtxSchemeLabelPresets` operationId).
	CtxSchemeLabelPresetsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeLabelPresetsParams, reqEditors ...RequestEditorFn) (*CtxSchemeLabelPresetsHTTP, error)

	// CtxSchemeNoticesWithBodyWithResponse Collect notices for assertions in a context
	//
	// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
	CtxSchemeNoticesWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeNoticesHTTP, error)

	// CtxSchemeNoticesWithResponse Collect notices for assertions in a context
	//
	// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
	CtxSchemeNoticesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, body CtxSchemeNoticesJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeNoticesHTTP, error)

	// CtxSchemeProcessWithBodyWithResponse Process a marking in a context
	//
	// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
	CtxSchemeProcessWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeProcessHTTP, error)

	// CtxSchemeProcessWithResponse Process a marking in a context
	//
	// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
	CtxSchemeProcessWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, body CtxSchemeProcessJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeProcessHTTP, error)

	// CtxSchemeRenderWithBodyWithResponse Render assertions in a context
	//
	// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
	CtxSchemeRenderWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeRenderHTTP, error)

	// CtxSchemeRenderWithResponse Render assertions in a context
	//
	// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
	CtxSchemeRenderWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, body CtxSchemeRenderJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeRenderHTTP, error)

	// CtxSchemeScanWithBodyWithResponse Scan text for markings in a context
	//
	// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
	CtxSchemeScanWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeScanHTTP, error)

	// CtxSchemeScanWithResponse Scan text for markings in a context
	//
	// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
	CtxSchemeScanWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, body CtxSchemeScanJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeScanHTTP, error)

	// CtxSchemeScopesWithResponse Get the scheme scopes in a context
	//
	// Returns the scope definitions for the context's composed scheme in declaration order. Empty when it declares none. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/scopes (the `CtxSchemeScopes` operationId).
	CtxSchemeScopesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScopesParams, reqEditors ...RequestEditorFn) (*CtxSchemeScopesHTTP, error)

	// CtxSchemeSuggestWithBodyWithResponse Suggest marking completions in a context
	//
	// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
	CtxSchemeSuggestWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestHTTP, error)

	// CtxSchemeSuggestWithResponse Suggest marking completions in a context
	//
	// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
	CtxSchemeSuggestWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, body CtxSchemeSuggestJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestHTTP, error)

	// CtxSchemeSuggestForCategoriesWithBodyWithResponse Suggest completions for several categories in a context
	//
	// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
	CtxSchemeSuggestForCategoriesWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestForCategoriesHTTP, error)

	// CtxSchemeSuggestForCategoriesWithResponse Suggest completions for several categories in a context
	//
	// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
	CtxSchemeSuggestForCategoriesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, body CtxSchemeSuggestForCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestForCategoriesHTTP, error)

	// CtxSchemeValidateWithBodyWithResponse Validate a marking in a context
	//
	// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
	CtxSchemeValidateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeValidateHTTP, error)

	// CtxSchemeValidateWithResponse Validate a marking in a context
	//
	// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
	CtxSchemeValidateWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, body CtxSchemeValidateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeValidateHTTP, error)

	// CtxSchemeViewWithBodyWithResponse Build scheme structural view in a context
	//
	// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
	CtxSchemeViewWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeViewHTTP, error)

	// CtxSchemeViewWithResponse Build scheme structural view in a context
	//
	// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
	//
	// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
	//
	// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
	CtxSchemeViewWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, body CtxSchemeViewJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeViewHTTP, error)

	// GetHealthWithResponse Health check
	//
	// Returns service liveness status and the engine version string.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/health (the `GetHealth` operationId).
	GetHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthHTTP, error)

	// GetReadyWithResponse Readiness check
	//
	// Reports whether the shared runtime host admitted the configured license and all explicitly required local artifacts are available. Expiration never makes a replica unready.
	//
	// Returns a wrapper object for the known response body format(s).
	//
	// Corresponds with GET /v1/ready (the `GetReady` operationId).
	GetReadyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetReadyHTTP, error)
}

// CtxListHTTP429Headers the declared response headers of an HTTP 429 response for CtxList
type CtxListHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxListHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *CtxListResponseBody
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxListHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxListHTTP) GetJSON200() *CtxListResponseBody {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxListHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxListHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxListHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxListHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxListHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxListHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxListHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxListHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxVersionsHTTP429Headers the declared response headers of an HTTP 429 response for CtxVersions
type CtxVersionsHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxVersionsHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *CtxVersionsResponseBody
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxVersionsHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxVersionsHTTP) GetJSON200() *CtxVersionsResponseBody {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxVersionsHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxVersionsHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxVersionsHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxVersionsHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxVersionsHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxVersionsHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxVersionsHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxVersionsHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxDescribeHTTP200Headers the declared response headers of an HTTP 200 response for CtxDescribe
type CtxDescribeHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxDescribeHTTP429Headers the declared response headers of an HTTP 429 response for CtxDescribe
type CtxDescribeHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxDescribeHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *CtxDescribeResponseBody
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxDescribeHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxDescribeHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxDescribeHTTP) GetJSON200() *CtxDescribeResponseBody {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxDescribeHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxDescribeHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxDescribeHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxDescribeHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxDescribeHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxDescribeHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxDescribeHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxDescribeHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxAccessCheckHTTP200Headers the declared response headers of an HTTP 200 response for CtxAccessCheck
type CtxAccessCheckHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxAccessCheckHTTP429Headers the declared response headers of an HTTP 429 response for CtxAccessCheck
type CtxAccessCheckHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxAccessCheckHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextAccessDecisionResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxAccessCheckHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxAccessCheckHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON200() *ContextAccessDecisionResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxAccessCheckHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxAccessCheckHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxAccessCheckHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxAccessCheckHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxAccessCheckHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxAccessDescribeRequirementsHTTP200Headers the declared response headers of an HTTP 200 response for CtxAccessDescribeRequirements
type CtxAccessDescribeRequirementsHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxAccessDescribeRequirementsHTTP429Headers the declared response headers of an HTTP 429 response for CtxAccessDescribeRequirements
type CtxAccessDescribeRequirementsHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxAccessDescribeRequirementsHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextAccessRequirementsResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxAccessDescribeRequirementsHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxAccessDescribeRequirementsHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON200() *ContextAccessRequirementsResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxAccessDescribeRequirementsHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxAccessDescribeRequirementsHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxAccessDescribeRequirementsHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxAccessDescribeRequirementsHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxAccessDescribeRequirementsHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxAccessRulesHTTP200Headers the declared response headers of an HTTP 200 response for CtxAccessRules
type CtxAccessRulesHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxAccessRulesHTTP429Headers the declared response headers of an HTTP 429 response for CtxAccessRules
type CtxAccessRulesHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxAccessRulesHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *CtxAccessRulesResponseBody
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxAccessRulesHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxAccessRulesHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON200() *CtxAccessRulesResponseBody {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxAccessRulesHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxAccessRulesHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxAccessRulesHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxAccessRulesHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxAccessRulesHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxCrosswalkDescribeHTTP200Headers the declared response headers of an HTTP 200 response for CtxCrosswalkDescribe
type CtxCrosswalkDescribeHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxCrosswalkDescribeHTTP429Headers the declared response headers of an HTTP 429 response for CtxCrosswalkDescribe
type CtxCrosswalkDescribeHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxCrosswalkDescribeHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *CtxCrosswalkDescribeResponseBody
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxCrosswalkDescribeHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxCrosswalkDescribeHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxCrosswalkDescribeHTTP) GetJSON200() *CtxCrosswalkDescribeResponseBody {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxCrosswalkDescribeHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxCrosswalkDescribeHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxCrosswalkDescribeHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxCrosswalkDescribeHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxCrosswalkDescribeHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxCrosswalkDescribeHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxCrosswalkDescribeHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxCrosswalkDescribeHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxCrosswalkTranslateHTTP200Headers the declared response headers of an HTTP 200 response for CtxCrosswalkTranslate
type CtxCrosswalkTranslateHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxCrosswalkTranslateHTTP429Headers the declared response headers of an HTTP 429 response for CtxCrosswalkTranslate
type CtxCrosswalkTranslateHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxCrosswalkTranslateHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextCrosswalkTranslationResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxCrosswalkTranslateHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxCrosswalkTranslateHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON200() *ContextCrosswalkTranslationResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxCrosswalkTranslateHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxCrosswalkTranslateHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxCrosswalkTranslateHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxCrosswalkTranslateHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxCrosswalkTranslateHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxDocumentAuditHTTP200Headers the declared response headers of an HTTP 200 response for CtxDocumentAudit
type CtxDocumentAuditHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxDocumentAuditHTTP429Headers the declared response headers of an HTTP 429 response for CtxDocumentAudit
type CtxDocumentAuditHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxDocumentAuditHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextDocumentAuditResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON409 the response for an HTTP 409 `application/json` response
	JSON409 *Conflict
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxDocumentAuditHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxDocumentAuditHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON200() *ContextDocumentAuditResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON409 returns the response for an HTTP 409 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON409() *Conflict {
	return r.JSON409
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxDocumentAuditHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxDocumentAuditHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxDocumentAuditHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxDocumentAuditHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxDocumentAuditHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxGuideEvaluateHTTP200Headers the declared response headers of an HTTP 200 response for CtxGuideEvaluate
type CtxGuideEvaluateHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxGuideEvaluateHTTP429Headers the declared response headers of an HTTP 429 response for CtxGuideEvaluate
type CtxGuideEvaluateHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxGuideEvaluateHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextGuideEvaluationResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxGuideEvaluateHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxGuideEvaluateHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON200() *ContextGuideEvaluationResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxGuideEvaluateHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxGuideEvaluateHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxGuideEvaluateHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxGuideEvaluateHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxGuideEvaluateHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxGuideScanHTTP200Headers the declared response headers of an HTTP 200 response for CtxGuideScan
type CtxGuideScanHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxGuideScanHTTP429Headers the declared response headers of an HTTP 429 response for CtxGuideScan
type CtxGuideScanHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxGuideScanHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextGuideScanResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxGuideScanHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxGuideScanHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxGuideScanHTTP) GetJSON200() *ContextGuideScanResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxGuideScanHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxGuideScanHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxGuideScanHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxGuideScanHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxGuideScanHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxGuideScanHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxGuideScanHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxGuideScanHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxGuideScanHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxGuideViewHTTP200Headers the declared response headers of an HTTP 200 response for CtxGuideView
type CtxGuideViewHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxGuideViewHTTP429Headers the declared response headers of an HTTP 429 response for CtxGuideView
type CtxGuideViewHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxGuideViewHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextGuideViewResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON422 the response for an HTTP 422 `application/json` response
	JSON422 *UnprocessableEntity
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxGuideViewHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxGuideViewHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxGuideViewHTTP) GetJSON200() *ContextGuideViewResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxGuideViewHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxGuideViewHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON422 returns the response for an HTTP 422 `application/json` response
func (r CtxGuideViewHTTP) GetJSON422() *UnprocessableEntity {
	return r.JSON422
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxGuideViewHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxGuideViewHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxGuideViewHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxGuideViewHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxGuideViewHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxGuideViewHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeSummaryHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeSummary
type CtxSchemeSummaryHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeSummaryHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeSummary
type CtxSchemeSummaryHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeSummaryHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *SchemeSummary
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeSummaryHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeSummaryHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeSummaryHTTP) GetJSON200() *SchemeSummary {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeSummaryHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeSummaryHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeSummaryHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeSummaryHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeSummaryHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeSummaryHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeSummaryHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeSummaryHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeAggregateHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeAggregate
type CtxSchemeAggregateHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeAggregateHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeAggregate
type CtxSchemeAggregateHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeAggregateHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeAggregateResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeAggregateHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeAggregateHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeAggregateHTTP) GetJSON200() *ContextSchemeAggregateResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeAggregateHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeAggregateHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeAggregateHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeAggregateHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeAggregateHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeAggregateHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeAggregateHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeAggregateHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeApplyPresetHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeApplyPreset
type CtxSchemeApplyPresetHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeApplyPresetHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeApplyPreset
type CtxSchemeApplyPresetHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeApplyPresetHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeValidateResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeApplyPresetHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeApplyPresetHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeApplyPresetHTTP) GetJSON200() *ContextSchemeValidateResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeApplyPresetHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeApplyPresetHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeApplyPresetHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeApplyPresetHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeApplyPresetHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeApplyPresetHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeApplyPresetHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeApplyPresetHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeCategoriesHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeCategories
type CtxSchemeCategoriesHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeCategoriesHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeCategories
type CtxSchemeCategoriesHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeCategoriesHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *[]CategoryView
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeCategoriesHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeCategoriesHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeCategoriesHTTP) GetJSON200() *[]CategoryView {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeCategoriesHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeCategoriesHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeCategoriesHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeCategoriesHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeCategoriesHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeCategoriesHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeCategoriesHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeCategoriesHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeCheckCoverageHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeCheckCoverage
type CtxSchemeCheckCoverageHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeCheckCoverageHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeCheckCoverage
type CtxSchemeCheckCoverageHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeCheckCoverageHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeCoverageResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeCheckCoverageHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeCheckCoverageHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeCheckCoverageHTTP) GetJSON200() *ContextSchemeCoverageResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeCheckCoverageHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeCheckCoverageHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeCheckCoverageHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeCheckCoverageHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeCheckCoverageHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeCheckCoverageHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeCheckCoverageHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeCheckCoverageHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeConceptsHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeConcepts
type CtxSchemeConceptsHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeConceptsHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeConcepts
type CtxSchemeConceptsHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeConceptsHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *[]ConceptView
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeConceptsHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeConceptsHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeConceptsHTTP) GetJSON200() *[]ConceptView {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeConceptsHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeConceptsHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeConceptsHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeConceptsHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeConceptsHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeConceptsHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeConceptsHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeConceptsHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeDescribeConceptHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeDescribeConcept
type CtxSchemeDescribeConceptHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeDescribeConceptHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeDescribeConcept
type CtxSchemeDescribeConceptHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeDescribeConceptHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeConceptDescriptionResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeDescribeConceptHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeDescribeConceptHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeDescribeConceptHTTP) GetJSON200() *ContextSchemeConceptDescriptionResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeDescribeConceptHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeDescribeConceptHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeDescribeConceptHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeDescribeConceptHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeDescribeConceptHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeDescribeConceptHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeDescribeConceptHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeDescribeConceptHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeDiffHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeDiff
type CtxSchemeDiffHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeDiffHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeDiff
type CtxSchemeDiffHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeDiffHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeDiffResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeDiffHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeDiffHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeDiffHTTP) GetJSON200() *ContextSchemeDiffResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeDiffHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeDiffHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeDiffHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeDiffHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeDiffHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeDiffHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeDiffHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeDiffHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeExplainHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeExplain
type CtxSchemeExplainHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeExplainHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeExplain
type CtxSchemeExplainHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeExplainHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeExplainResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeExplainHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeExplainHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeExplainHTTP) GetJSON200() *ContextSchemeExplainResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeExplainHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeExplainHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeExplainHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeExplainHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeExplainHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeExplainHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeExplainHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeExplainHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeFormatsHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeFormats
type CtxSchemeFormatsHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeFormatsHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeFormats
type CtxSchemeFormatsHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeFormatsHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *[]FormatView
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeFormatsHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeFormatsHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeFormatsHTTP) GetJSON200() *[]FormatView {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeFormatsHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeFormatsHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeFormatsHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeFormatsHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeFormatsHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeFormatsHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeFormatsHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeFormatsHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeLabelPresetsHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeLabelPresets
type CtxSchemeLabelPresetsHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeLabelPresetsHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeLabelPresets
type CtxSchemeLabelPresetsHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeLabelPresetsHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *[]LabelPresetView
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeLabelPresetsHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeLabelPresetsHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeLabelPresetsHTTP) GetJSON200() *[]LabelPresetView {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeLabelPresetsHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeLabelPresetsHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeLabelPresetsHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeLabelPresetsHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeLabelPresetsHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeLabelPresetsHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeLabelPresetsHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeLabelPresetsHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeNoticesHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeNotices
type CtxSchemeNoticesHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeNoticesHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeNotices
type CtxSchemeNoticesHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeNoticesHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeNoticesResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeNoticesHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeNoticesHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeNoticesHTTP) GetJSON200() *ContextSchemeNoticesResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeNoticesHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeNoticesHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeNoticesHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeNoticesHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeNoticesHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeNoticesHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeNoticesHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeNoticesHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeProcessHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeProcess
type CtxSchemeProcessHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeProcessHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeProcess
type CtxSchemeProcessHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeProcessHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeProcessResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeProcessHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeProcessHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeProcessHTTP) GetJSON200() *ContextSchemeProcessResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeProcessHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeProcessHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeProcessHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeProcessHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeProcessHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeProcessHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeProcessHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeProcessHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeRenderHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeRender
type CtxSchemeRenderHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeRenderHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeRender
type CtxSchemeRenderHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeRenderHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeRenderResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeRenderHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeRenderHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeRenderHTTP) GetJSON200() *ContextSchemeRenderResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeRenderHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeRenderHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeRenderHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeRenderHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeRenderHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeRenderHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeRenderHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeRenderHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeScanHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeScan
type CtxSchemeScanHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeScanHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeScan
type CtxSchemeScanHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeScanHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeScanResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeScanHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeScanHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeScanHTTP) GetJSON200() *ContextSchemeScanResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeScanHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeScanHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeScanHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeScanHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeScanHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeScanHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeScanHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeScanHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeScopesHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeScopes
type CtxSchemeScopesHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeScopesHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeScopes
type CtxSchemeScopesHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeScopesHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *SchemeScopesResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeScopesHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeScopesHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeScopesHTTP) GetJSON200() *SchemeScopesResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeScopesHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeScopesHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeScopesHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeScopesHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeScopesHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeScopesHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeScopesHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeScopesHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeSuggestHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeSuggest
type CtxSchemeSuggestHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeSuggestHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeSuggest
type CtxSchemeSuggestHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeSuggestHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeSuggestResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeSuggestHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeSuggestHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeSuggestHTTP) GetJSON200() *ContextSchemeSuggestResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeSuggestHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeSuggestHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeSuggestHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeSuggestHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeSuggestHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeSuggestHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeSuggestHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeSuggestHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeSuggestForCategoriesHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeSuggestForCategories
type CtxSchemeSuggestForCategoriesHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeSuggestForCategoriesHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeSuggestForCategories
type CtxSchemeSuggestForCategoriesHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeSuggestForCategoriesHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeSuggestForCategoriesResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeSuggestForCategoriesHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeSuggestForCategoriesHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeSuggestForCategoriesHTTP) GetJSON200() *ContextSchemeSuggestForCategoriesResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeSuggestForCategoriesHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeSuggestForCategoriesHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeSuggestForCategoriesHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeSuggestForCategoriesHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeSuggestForCategoriesHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeSuggestForCategoriesHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeSuggestForCategoriesHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeSuggestForCategoriesHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeValidateHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeValidate
type CtxSchemeValidateHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeValidateHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeValidate
type CtxSchemeValidateHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeValidateHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeValidateResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeValidateHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeValidateHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeValidateHTTP) GetJSON200() *ContextSchemeValidateResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeValidateHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeValidateHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeValidateHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeValidateHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeValidateHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeValidateHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeValidateHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeValidateHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxSchemeViewHTTP200Headers the declared response headers of an HTTP 200 response for CtxSchemeView
type CtxSchemeViewHTTP200Headers struct {
	XEzClassifyResolvedVersion *string
}

// CtxSchemeViewHTTP429Headers the declared response headers of an HTTP 429 response for CtxSchemeView
type CtxSchemeViewHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type CtxSchemeViewHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ContextSchemeViewResult
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON404 the response for an HTTP 404 `application/json` response
	JSON404 *NotFound
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON500 the response for an HTTP 500 `application/json` response
	JSON500 *InternalError
	// Headers200 the parsed response headers for an HTTP 200 response
	Headers200 *CtxSchemeViewHTTP200Headers
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *CtxSchemeViewHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r CtxSchemeViewHTTP) GetJSON200() *ContextSchemeViewResult {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r CtxSchemeViewHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON404 returns the response for an HTTP 404 `application/json` response
func (r CtxSchemeViewHTTP) GetJSON404() *NotFound {
	return r.JSON404
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r CtxSchemeViewHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON500 returns the response for an HTTP 500 `application/json` response
func (r CtxSchemeViewHTTP) GetJSON500() *InternalError {
	return r.JSON500
}

// GetBody returns the raw response body bytes
func (r CtxSchemeViewHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r CtxSchemeViewHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CtxSchemeViewHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r CtxSchemeViewHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// GetHealthHTTP429Headers the declared response headers of an HTTP 429 response for GetHealth
type GetHealthHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type GetHealthHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *HealthResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *GetHealthHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetHealthHTTP) GetJSON200() *HealthResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetHealthHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetHealthHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetBody returns the raw response body bytes
func (r GetHealthHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetHealthHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetHealthHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetHealthHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// GetReadyHTTP429Headers the declared response headers of an HTTP 429 response for GetReady
type GetReadyHTTP429Headers struct {
	RetryAfter          *int
	XRateLimitLimit     *int
	XRateLimitRemaining *int
	XRateLimitReset     *int
}

type GetReadyHTTP struct {
	Body         []byte
	HTTPResponse *http.Response
	// JSON200 the response for an HTTP 200 `application/json` response
	JSON200 *ReadyResponse
	// JSON400 the response for an HTTP 400 `application/json` response
	JSON400 *BadRequest
	// JSON429 the response for an HTTP 429 `application/json` response
	JSON429 *TooManyRequests
	// JSON503 the response for an HTTP 503 `application/json` response
	JSON503 *NotReady
	// Headers429 the parsed response headers for an HTTP 429 response
	Headers429 *GetReadyHTTP429Headers
}

// GetJSON200 returns the response for an HTTP 200 `application/json` response
func (r GetReadyHTTP) GetJSON200() *ReadyResponse {
	return r.JSON200
}

// GetJSON400 returns the response for an HTTP 400 `application/json` response
func (r GetReadyHTTP) GetJSON400() *BadRequest {
	return r.JSON400
}

// GetJSON429 returns the response for an HTTP 429 `application/json` response
func (r GetReadyHTTP) GetJSON429() *TooManyRequests {
	return r.JSON429
}

// GetJSON503 returns the response for an HTTP 503 `application/json` response
func (r GetReadyHTTP) GetJSON503() *NotReady {
	return r.JSON503
}

// GetBody returns the raw response body bytes
func (r GetReadyHTTP) GetBody() []byte {
	return r.Body
}

// Status returns HTTPResponse.Status
func (r GetReadyHTTP) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetReadyHTTP) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ContentType is a convenience method to retrieve the Content-Type value from the HTTP response headers
func (r GetReadyHTTP) ContentType() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Header.Get("Content-Type")
	}
	return ""
}

// CtxListWithResponse List the context ids the caller may use
//
// Lists the installed context ids the caller may use. On a managed deployment, a Context token lists the one context it names, and an API-key caller without a token lists the contexts the deployment approves for its organization, when the deployment accepts API keys. A standalone runtime takes no Context token: installation is the grant, and it lists every installed context it can serve now, leaving out one whose latest version does not resolve or whose content the license does not cover. The list is unpaginated. On a managed deployment, returns 404 when authorization is denied.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx (the `CtxList` operationId).
func (c *ClientWithResponses) CtxListWithResponse(ctx context.Context, params *CtxListParams, reqEditors ...RequestEditorFn) (*CtxListHTTP, error) {
	rsp, err := c.CtxList(ctx, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxListHTTP(rsp)
}

// CtxVersionsWithResponse List available versions for a context id
//
// Lists every installed version of one context id that the caller may use, highest semver first. On a managed deployment, these are the versions its Context token or API key admits. A standalone runtime takes no Context token: installation is the grant, and it lists every installed version it can serve, leaving out one whose content the license does not cover. Returns 404 when the id is not installed or authorization is denied.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId} (the `CtxVersions` operationId).
func (c *ClientWithResponses) CtxVersionsWithResponse(ctx context.Context, ctxID string, params *CtxVersionsParams, reqEditors ...RequestEditorFn) (*CtxVersionsHTTP, error) {
	rsp, err := c.CtxVersions(ctx, ctxID, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxVersionsHTTP(rsp)
}

// CtxDescribeWithResponse Describe a resolved context
//
// Returns the component breakdown for the resolved context: scheme, ordered layers, ordered guides, optional access policy, and ordered crosswalks, each with the version it declares and the concrete version that answered. The :version segment may be a concrete semver or the literal "latest"; the resolved version is reported in the X-EzClassify-Resolved-Version response header and echoed in the response body.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version} (the `CtxDescribe` operationId).
func (c *ClientWithResponses) CtxDescribeWithResponse(ctx context.Context, ctxID string, version string, params *CtxDescribeParams, reqEditors ...RequestEditorFn) (*CtxDescribeHTTP, error) {
	rsp, err := c.CtxDescribe(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxDescribeHTTP(rsp)
}

// CtxAccessCheckWithBodyWithResponse Check access decision in a context
//
// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
func (c *ClientWithResponses) CtxAccessCheckWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxAccessCheckHTTP, error) {
	rsp, err := c.CtxAccessCheckWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxAccessCheckHTTP(rsp)
}

// CtxAccessCheckWithResponse Check access decision in a context
//
// Evaluates the context's access policy against a TARGET (a `marking` string or its structured `assertions` form, merged when both are given) and the SUBJECT (`userAttributes`), with optional comparison `context`. When decision audit is enabled, configured JSON Pointer paths select required user and resource identifiers from `context`. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/check (the `CtxAccessCheck` operationId).
func (c *ClientWithResponses) CtxAccessCheckWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessCheckParams, body CtxAccessCheckJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxAccessCheckHTTP, error) {
	rsp, err := c.CtxAccessCheck(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxAccessCheckHTTP(rsp)
}

// CtxAccessDescribeRequirementsWithBodyWithResponse Describe access requirements for a target in a context
//
// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
func (c *ClientWithResponses) CtxAccessDescribeRequirementsWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxAccessDescribeRequirementsHTTP, error) {
	rsp, err := c.CtxAccessDescribeRequirementsWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxAccessDescribeRequirementsHTTP(rsp)
}

// CtxAccessDescribeRequirementsWithResponse Describe access requirements for a target in a context
//
// Describes requirements imposed by immutable policy for a trusted target supplied as `marking`, structured `assertions`, or both. Both target forms are normalized through the same Context scheme path used by access check. These are policy requirements for the supplied target, not remediation instructions and not advice to change immutable `userAttributes`.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/access/describe-requirements (the `CtxAccessDescribeRequirements` operationId).
func (c *ClientWithResponses) CtxAccessDescribeRequirementsWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessDescribeRequirementsParams, body CtxAccessDescribeRequirementsJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxAccessDescribeRequirementsHTTP, error) {
	rsp, err := c.CtxAccessDescribeRequirements(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxAccessDescribeRequirementsHTTP(rsp)
}

// CtxAccessRulesWithResponse Describe every access rule in a context
//
// Returns the human-readable predicate, the `userAttributes` paths, and the data categories each rule depends on, for every rule in the context's access policy. A caller that only needs the rule ids reads them from the same response. The :version segment may be omitted to default to "latest".
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/access/rules (the `CtxAccessRules` operationId).
func (c *ClientWithResponses) CtxAccessRulesWithResponse(ctx context.Context, ctxID string, version string, params *CtxAccessRulesParams, reqEditors ...RequestEditorFn) (*CtxAccessRulesHTTP, error) {
	rsp, err := c.CtxAccessRules(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxAccessRulesHTTP(rsp)
}

// CtxCrosswalkDescribeWithResponse Describe every crosswalk in a context
//
// Returns every crosswalk the context holds, in declaration order: its id and the concrete version that answers, its name, guidance and citation, the source and target schemes as the selectors the context description lists under `schemes`, and the mappings and selections a translation applies, each with its disposition and reason. A context that declares no crosswalk answers with an empty list. The :version segment may be omitted to default to "latest".
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/crosswalk/describe (the `CtxCrosswalkDescribe` operationId).
func (c *ClientWithResponses) CtxCrosswalkDescribeWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkDescribeParams, reqEditors ...RequestEditorFn) (*CtxCrosswalkDescribeHTTP, error) {
	rsp, err := c.CtxCrosswalkDescribe(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxCrosswalkDescribeHTTP(rsp)
}

// CtxCrosswalkTranslateWithBodyWithResponse Translate a marking through a crosswalk in a context
//
// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
func (c *ClientWithResponses) CtxCrosswalkTranslateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxCrosswalkTranslateHTTP, error) {
	rsp, err := c.CtxCrosswalkTranslateWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxCrosswalkTranslateHTTP(rsp)
}

// CtxCrosswalkTranslateWithResponse Translate a marking through a crosswalk in a context
//
// Applies the named directional crosswalk to a SOURCE and returns the treatment of every source assertion: what was carried, what needs review, and what no mapping covers. A prohibition stops the translation before target processing and the response says which one. The source is either a rendered `marking` string or its structured `assertions` form: send exactly one of the two, because a request with neither has no source and a request with both has two. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/crosswalk/translate (the `CtxCrosswalkTranslate` operationId).
func (c *ClientWithResponses) CtxCrosswalkTranslateWithResponse(ctx context.Context, ctxID string, version string, params *CtxCrosswalkTranslateParams, body CtxCrosswalkTranslateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxCrosswalkTranslateHTTP, error) {
	rsp, err := c.CtxCrosswalkTranslate(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxCrosswalkTranslateHTTP(rsp)
}

// CtxDocumentAuditWithBodyWithResponse Audit a document for marking coverage in a context
//
// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
func (c *ClientWithResponses) CtxDocumentAuditWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxDocumentAuditHTTP, error) {
	rsp, err := c.CtxDocumentAuditWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxDocumentAuditHTTP(rsp)
}

// CtxDocumentAuditWithResponse Audit a document for marking coverage in a context
//
// Audits authored document markings and evaluates guide nominations without treating lexical matches as authoritative facts. Never accept a nomination without an explicit analyst decision bound to its finding ID and document digest; revised text requires new decisions. Determined results include only accepted findings and manual selections, while projected results assume unresolved nominations and remain advisory. This operation answers authored correctness and guide review. The standard response profile is the default; summary suits only an initial large-document overview, standard is needed before analyst decisions or remediation, and full only when matched text or complete breakdowns are necessary. responseProjection counts and omittedFields state what the profile left out; omitted detail is not absent evidence.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/document/audit (the `CtxDocumentAudit` operationId).
func (c *ClientWithResponses) CtxDocumentAuditWithResponse(ctx context.Context, ctxID string, version string, params *CtxDocumentAuditParams, body CtxDocumentAuditJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxDocumentAuditHTTP, error) {
	rsp, err := c.CtxDocumentAudit(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxDocumentAuditHTTP(rsp)
}

// CtxGuideEvaluateWithBodyWithResponse Evaluate selected guide rules in a context
//
// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
func (c *ClientWithResponses) CtxGuideEvaluateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideEvaluateHTTP, error) {
	rsp, err := c.CtxGuideEvaluateWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideEvaluateHTTP(rsp)
}

// CtxGuideEvaluateWithResponse Evaluate selected guide rules in a context
//
// Send `guideId` with `selectedRuleIds` to evaluate that one guide, or `selections` to roll up rules selected from several guides. Each selection names its guide, so two guides that use the same rule ID hold two rules. Each guide evaluates its own selected rules, with its own dependencies and overrides, and the outcomes are rolled up per target scheme. No guide overrides or depends on another. Returns the engine's structured evaluation result.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/evaluate (the `CtxGuideEvaluate` operationId).
func (c *ClientWithResponses) CtxGuideEvaluateWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideEvaluateParams, body CtxGuideEvaluateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideEvaluateHTTP, error) {
	rsp, err := c.CtxGuideEvaluate(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideEvaluateHTTP(rsp)
}

// CtxGuideScanWithBodyWithResponse Scan text against context guides
//
// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
func (c *ClientWithResponses) CtxGuideScanWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideScanHTTP, error) {
	rsp, err := c.CtxGuideScanWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideScanHTTP(rsp)
}

// CtxGuideScanWithResponse Scan text against context guides
//
// Scans the provided text against the guides attached to the context and returns per-guide scan results. The optional `guideId` body field restricts the scan to a single guide; otherwise every guide in the context is consulted.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/scan (the `CtxGuideScan` operationId).
func (c *ClientWithResponses) CtxGuideScanWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideScanParams, body CtxGuideScanJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideScanHTTP, error) {
	rsp, err := c.CtxGuideScan(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideScanHTTP(rsp)
}

// CtxGuideViewWithBodyWithResponse Build guide structural view in a context
//
// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
func (c *ClientWithResponses) CtxGuideViewWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxGuideViewHTTP, error) {
	rsp, err := c.CtxGuideViewWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideViewHTTP(rsp)
}

// CtxGuideViewWithResponse Build guide structural view in a context
//
// Returns a structural view of the context-bound guide (categories, rule index, outcomes) suitable for rendering a picker UI. Pass guideId to disambiguate when the context has multiple guides.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/guide/view (the `CtxGuideView` operationId).
func (c *ClientWithResponses) CtxGuideViewWithResponse(ctx context.Context, ctxID string, version string, params *CtxGuideViewParams, body CtxGuideViewJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxGuideViewHTTP, error) {
	rsp, err := c.CtxGuideView(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxGuideViewHTTP(rsp)
}

// CtxSchemeSummaryWithResponse Get the scheme summary in a context
//
// Returns high-level metadata for the context's composed scheme (base plus its declared layers): id, version, name, description, presentation configuration, default render format, and scope label. The scheme identity and layers come from the Context. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme (the `CtxSchemeSummary` operationId).
func (c *ClientWithResponses) CtxSchemeSummaryWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSummaryParams, reqEditors ...RequestEditorFn) (*CtxSchemeSummaryHTTP, error) {
	rsp, err := c.CtxSchemeSummary(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeSummaryHTTP(rsp)
}

// CtxSchemeAggregateWithBodyWithResponse Aggregate marking strings in a context
//
// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
func (c *ClientWithResponses) CtxSchemeAggregateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeAggregateHTTP, error) {
	rsp, err := c.CtxSchemeAggregateWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeAggregateHTTP(rsp)
}

// CtxSchemeAggregateWithResponse Aggregate marking strings in a context
//
// Rolls markings up into the one whole marking that covers them, using the context's composed scheme rollup rules. With no `context` the inputs are read as wholes, as in an aggregate such as the documents in a folder; `context` names the role the inputs are read in, and `{ role: 'part', partOf }` reads each as a part of that whole, as in a whole's parts rollup. Supply either `inputs` (marking strings) or `assertionSets` (already-parsed assertion arrays for lossless transport of scope-only assertions and capture values): exactly one, not both. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/aggregate (the `CtxSchemeAggregate` operationId).
func (c *ClientWithResponses) CtxSchemeAggregateWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeAggregateParams, body CtxSchemeAggregateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeAggregateHTTP, error) {
	rsp, err := c.CtxSchemeAggregate(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeAggregateHTTP(rsp)
}

// CtxSchemeApplyPresetWithBodyWithResponse Apply a label preset in a context
//
// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
func (c *ClientWithResponses) CtxSchemeApplyPresetWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeApplyPresetHTTP, error) {
	rsp, err := c.CtxSchemeApplyPresetWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeApplyPresetHTTP(rsp)
}

// CtxSchemeApplyPresetWithResponse Apply a label preset in a context
//
// Expands the named label preset against the context's composed scheme and returns what validate answers for the marking it stands for, in the standard context operation envelope every other context operation answers with. A preset is a stored marking, so the answer carries no fix paths or corrections. The label-presets resource keeps each preset's marking string server-side, so this is the only way to order one of its items. Returns 404 when the scheme declares no preset by that id. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset (the `CtxSchemeApplyPreset` operationId).
func (c *ClientWithResponses) CtxSchemeApplyPresetWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeApplyPresetParams, body CtxSchemeApplyPresetJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeApplyPresetHTTP, error) {
	rsp, err := c.CtxSchemeApplyPreset(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeApplyPresetHTTP(rsp)
}

// CtxSchemeCategoriesWithResponse Get the scheme categories in a context
//
// Returns every category of the context's composed scheme as a category view, in declared order: label, scopes, default concept, rollup, and concept sort order, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/categories (the `CtxSchemeCategories` operationId).
func (c *ClientWithResponses) CtxSchemeCategoriesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCategoriesParams, reqEditors ...RequestEditorFn) (*CtxSchemeCategoriesHTTP, error) {
	rsp, err := c.CtxSchemeCategories(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeCategoriesHTTP(rsp)
}

// CtxSchemeCheckCoverageWithBodyWithResponse Check marking coverage in a context
//
// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
func (c *ClientWithResponses) CtxSchemeCheckCoverageWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeCheckCoverageHTTP, error) {
	rsp, err := c.CtxSchemeCheckCoverageWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeCheckCoverageHTTP(rsp)
}

// CtxSchemeCheckCoverageWithResponse Check marking coverage in a context
//
// Aggregates the child markings using the context's composed scheme and checks whether the result is covered by the parent marking. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/check-coverage (the `CtxSchemeCheckCoverage` operationId).
func (c *ClientWithResponses) CtxSchemeCheckCoverageWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeCheckCoverageParams, body CtxSchemeCheckCoverageJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeCheckCoverageHTTP, error) {
	rsp, err := c.CtxSchemeCheckCoverage(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeCheckCoverageHTTP(rsp)
}

// CtxSchemeConceptsWithResponse Get the scheme concepts in a context
//
// Returns every concept of the context's composed scheme as a concept view: id, name, views, categories, parents, position, rank, and display metadata, the same shape a local engine's scheme view carries. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/concepts (the `CtxSchemeConcepts` operationId).
func (c *ClientWithResponses) CtxSchemeConceptsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeConceptsParams, reqEditors ...RequestEditorFn) (*CtxSchemeConceptsHTTP, error) {
	rsp, err := c.CtxSchemeConcepts(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeConceptsHTTP(rsp)
}

// CtxSchemeDescribeConceptWithResponse Describe a concept in a context
//
// Returns structured description and metadata for the specified concept within a category, scoped to the context's composed scheme. The :version segment may be omitted to default to "latest".
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/describe-concept (the `CtxSchemeDescribeConcept` operationId).
func (c *ClientWithResponses) CtxSchemeDescribeConceptWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDescribeConceptParams, reqEditors ...RequestEditorFn) (*CtxSchemeDescribeConceptHTTP, error) {
	rsp, err := c.CtxSchemeDescribeConcept(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeDescribeConceptHTTP(rsp)
}

// CtxSchemeDiffWithBodyWithResponse Diff two markings in a context
//
// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
func (c *ClientWithResponses) CtxSchemeDiffWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeDiffHTTP, error) {
	rsp, err := c.CtxSchemeDiffWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeDiffHTTP(rsp)
}

// CtxSchemeDiffWithResponse Diff two markings in a context
//
// Returns the structural diff between two marking strings as evaluated against the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/diff (the `CtxSchemeDiff` operationId).
func (c *ClientWithResponses) CtxSchemeDiffWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeDiffParams, body CtxSchemeDiffJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeDiffHTTP, error) {
	rsp, err := c.CtxSchemeDiff(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeDiffHTTP(rsp)
}

// CtxSchemeExplainWithBodyWithResponse Explain how a marking is processed in a context
//
// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
func (c *ClientWithResponses) CtxSchemeExplainWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeExplainHTTP, error) {
	rsp, err := c.CtxSchemeExplainWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeExplainHTTP(rsp)
}

// CtxSchemeExplainWithResponse Explain how a marking is processed in a context
//
// Processes the input exactly as the process operation does and returns the trace of that processing instead of its result: what was read, what was repaired, what the scheme's rules inferred, what was found wrong, and what was written. Use it to answer why a marking came out the way it did. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/explain (the `CtxSchemeExplain` operationId).
func (c *ClientWithResponses) CtxSchemeExplainWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeExplainParams, body CtxSchemeExplainJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeExplainHTTP, error) {
	rsp, err := c.CtxSchemeExplain(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeExplainHTTP(rsp)
}

// CtxSchemeFormatsWithResponse Get the scheme render formats in a context
//
// Lists the formats the context's composed scheme writes as format views: each format's id, the surface it writes, its resolved layout, and how its text is framed, the same shape a local engine's scheme view carries. Not the field trees, which stay server-side. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/formats (the `CtxSchemeFormats` operationId).
func (c *ClientWithResponses) CtxSchemeFormatsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeFormatsParams, reqEditors ...RequestEditorFn) (*CtxSchemeFormatsHTTP, error) {
	rsp, err := c.CtxSchemeFormats(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeFormatsHTTP(rsp)
}

// CtxSchemeLabelPresetsWithResponse List the label presets in a context
//
// Returns the label presets of the selected scheme as preset views: id, label, description, and scope. Empty when the scheme declares none. The marking string and tag list stay server-side; expand a preset with POST /v1/ctx/{ctxId}/{version}/scheme/apply-preset under the same selector. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/label-presets (the `CtxSchemeLabelPresets` operationId).
func (c *ClientWithResponses) CtxSchemeLabelPresetsWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeLabelPresetsParams, reqEditors ...RequestEditorFn) (*CtxSchemeLabelPresetsHTTP, error) {
	rsp, err := c.CtxSchemeLabelPresets(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeLabelPresetsHTTP(rsp)
}

// CtxSchemeNoticesWithBodyWithResponse Collect notices for assertions in a context
//
// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
func (c *ClientWithResponses) CtxSchemeNoticesWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeNoticesHTTP, error) {
	rsp, err := c.CtxSchemeNoticesWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeNoticesHTTP(rsp)
}

// CtxSchemeNoticesWithResponse Collect notices for assertions in a context
//
// Collects non-blocking notices attached to the supplied assertions' concepts under the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header. Operations that answer with a processed marking, such as process, validate and aggregate, already carry the same list on the result's notices field.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/notices (the `CtxSchemeNotices` operationId).
func (c *ClientWithResponses) CtxSchemeNoticesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeNoticesParams, body CtxSchemeNoticesJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeNoticesHTTP, error) {
	rsp, err := c.CtxSchemeNotices(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeNoticesHTTP(rsp)
}

// CtxSchemeProcessWithBodyWithResponse Process a marking in a context
//
// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
func (c *ClientWithResponses) CtxSchemeProcessWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeProcessHTTP, error) {
	rsp, err := c.CtxSchemeProcessWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeProcessHTTP(rsp)
}

// CtxSchemeProcessWithResponse Process a marking in a context
//
// The author's call. Reads the input marking text, together with any facts already chosen, against the context's composed scheme (base + layers), settles it, and returns everything validate returns plus what an author needs: the corrections to the text, proven fix paths, the outputs named in `outputs` and, when `suggestFor` names categories, the picker options for them. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/process (the `CtxSchemeProcess` operationId).
func (c *ClientWithResponses) CtxSchemeProcessWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeProcessParams, body CtxSchemeProcessJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeProcessHTTP, error) {
	rsp, err := c.CtxSchemeProcess(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeProcessHTTP(rsp)
}

// CtxSchemeRenderWithBodyWithResponse Render assertions in a context
//
// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
func (c *ClientWithResponses) CtxSchemeRenderWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeRenderHTTP, error) {
	rsp, err := c.CtxSchemeRenderWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeRenderHTTP(rsp)
}

// CtxSchemeRenderWithResponse Render assertions in a context
//
// Renders assertions to a marking string using the context's composed scheme. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/render (the `CtxSchemeRender` operationId).
func (c *ClientWithResponses) CtxSchemeRenderWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeRenderParams, body CtxSchemeRenderJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeRenderHTTP, error) {
	rsp, err := c.CtxSchemeRender(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeRenderHTTP(rsp)
}

// CtxSchemeScanWithBodyWithResponse Scan text for markings in a context
//
// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
func (c *ClientWithResponses) CtxSchemeScanWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeScanHTTP, error) {
	rsp, err := c.CtxSchemeScanWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeScanHTTP(rsp)
}

// CtxSchemeScanWithResponse Scan text for markings in a context
//
// Scans the provided text for marking occurrences against the context's composed scheme and returns the results in a resolved Context envelope. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/scan (the `CtxSchemeScan` operationId).
func (c *ClientWithResponses) CtxSchemeScanWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScanParams, body CtxSchemeScanJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeScanHTTP, error) {
	rsp, err := c.CtxSchemeScan(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeScanHTTP(rsp)
}

// CtxSchemeScopesWithResponse Get the scheme scopes in a context
//
// Returns the scope definitions for the context's composed scheme in declaration order. Empty when it declares none. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ctx/{ctxId}/{version}/scheme/scopes (the `CtxSchemeScopes` operationId).
func (c *ClientWithResponses) CtxSchemeScopesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeScopesParams, reqEditors ...RequestEditorFn) (*CtxSchemeScopesHTTP, error) {
	rsp, err := c.CtxSchemeScopes(ctx, ctxID, version, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeScopesHTTP(rsp)
}

// CtxSchemeSuggestWithBodyWithResponse Suggest marking completions in a context
//
// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
func (c *ClientWithResponses) CtxSchemeSuggestWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestHTTP, error) {
	rsp, err := c.CtxSchemeSuggestWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeSuggestHTTP(rsp)
}

// CtxSchemeSuggestWithResponse Suggest marking completions in a context
//
// Returns concept suggestions for a partial marking input within the context's composed scheme: the options only. Whether the marking stands is what validate and process answer. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest (the `CtxSchemeSuggest` operationId).
func (c *ClientWithResponses) CtxSchemeSuggestWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestParams, body CtxSchemeSuggestJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestHTTP, error) {
	rsp, err := c.CtxSchemeSuggest(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeSuggestHTTP(rsp)
}

// CtxSchemeSuggestForCategoriesWithBodyWithResponse Suggest completions for several categories in a context
//
// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
func (c *ClientWithResponses) CtxSchemeSuggestForCategoriesWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestForCategoriesHTTP, error) {
	rsp, err := c.CtxSchemeSuggestForCategoriesWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeSuggestForCategoriesHTTP(rsp)
}

// CtxSchemeSuggestForCategoriesWithResponse Suggest completions for several categories in a context
//
// Returns the suggestion result for each requested category within the context's composed scheme, keyed by category id, each computed as if that category's own assertions were not present. The map is the same one a process request that names `suggestFor` returns in `suggestionsFor`. Each result carries the options only; whether the marking stands is what validate and process answer. An unknown category id is refused. The :version segment may be omitted to default to "latest".
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/suggest-for-categories (the `CtxSchemeSuggestForCategories` operationId).
func (c *ClientWithResponses) CtxSchemeSuggestForCategoriesWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeSuggestForCategoriesParams, body CtxSchemeSuggestForCategoriesJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeSuggestForCategoriesHTTP, error) {
	rsp, err := c.CtxSchemeSuggestForCategories(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeSuggestForCategoriesHTTP(rsp)
}

// CtxSchemeValidateWithBodyWithResponse Validate a marking in a context
//
// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
func (c *ClientWithResponses) CtxSchemeValidateWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeValidateHTTP, error) {
	rsp, err := c.CtxSchemeValidateWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeValidateHTTP(rsp)
}

// CtxSchemeValidateWithResponse Validate a marking in a context
//
// The reader's call. Judges a marking, given as text or as the facts it states, against the context's composed scheme (base + layers). The marking is settled exactly as process settles it, so `valid` and `inputWasValid` mean what they mean on every route. `surface`, `context` and `parts` say what the marking is judged as, as they do on process. Returns the settled facts, the verdicts, errors and warnings, the changes settling made, the marking written back with what that writing left out, how recognition read it, the notices it requires, and its display color. Fix paths and corrections to the text belong to process. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/validate (the `CtxSchemeValidate` operationId).
func (c *ClientWithResponses) CtxSchemeValidateWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeValidateParams, body CtxSchemeValidateJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeValidateHTTP, error) {
	rsp, err := c.CtxSchemeValidate(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeValidateHTTP(rsp)
}

// CtxSchemeViewWithBodyWithResponse Build scheme structural view in a context
//
// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes any type of body and a specified content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
func (c *ClientWithResponses) CtxSchemeViewWithBodyWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CtxSchemeViewHTTP, error) {
	rsp, err := c.CtxSchemeViewWithBody(ctx, ctxID, version, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeViewHTTP(rsp)
}

// CtxSchemeViewWithResponse Build scheme structural view in a context
//
// Returns a categorized structural view of the context's composed scheme (base + layers), suitable for rendering a picker UI. The :version segment may be omitted to default to "latest"; the resolved version is returned in the X-EzClassify-Resolved-Version response header.
//
// Takes a body of the `application/json` content type, and returns a wrapper object for the known response body format(s).
//
// Corresponds with POST /v1/ctx/{ctxId}/{version}/scheme/view (the `CtxSchemeView` operationId).
func (c *ClientWithResponses) CtxSchemeViewWithResponse(ctx context.Context, ctxID string, version string, params *CtxSchemeViewParams, body CtxSchemeViewJSONRequestBody, reqEditors ...RequestEditorFn) (*CtxSchemeViewHTTP, error) {
	rsp, err := c.CtxSchemeView(ctx, ctxID, version, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCtxSchemeViewHTTP(rsp)
}

// GetHealthWithResponse Health check
//
// Returns service liveness status and the engine version string.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/health (the `GetHealth` operationId).
func (c *ClientWithResponses) GetHealthWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetHealthHTTP, error) {
	rsp, err := c.GetHealth(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetHealthHTTP(rsp)
}

// GetReadyWithResponse Readiness check
//
// Reports whether the shared runtime host admitted the configured license and all explicitly required local artifacts are available. Expiration never makes a replica unready.
//
// Returns a wrapper object for the known response body format(s).
//
// Corresponds with GET /v1/ready (the `GetReady` operationId).
func (c *ClientWithResponses) GetReadyWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetReadyHTTP, error) {
	rsp, err := c.GetReady(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetReadyHTTP(rsp)
}

// ParseCtxListHTTP parses an HTTP response from a CtxListWithResponse call
func ParseCtxListHTTP(rsp *http.Response) (*CtxListHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxListHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CtxListResponseBody
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 429:
		var headers CtxListHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxVersionsHTTP parses an HTTP response from a CtxVersionsWithResponse call
func ParseCtxVersionsHTTP(rsp *http.Response) (*CtxVersionsHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxVersionsHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CtxVersionsResponseBody
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 429:
		var headers CtxVersionsHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxDescribeHTTP parses an HTTP response from a CtxDescribeWithResponse call
func ParseCtxDescribeHTTP(rsp *http.Response) (*CtxDescribeHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxDescribeHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CtxDescribeResponseBody
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxDescribeHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxDescribeHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxAccessCheckHTTP parses an HTTP response from a CtxAccessCheckWithResponse call
func ParseCtxAccessCheckHTTP(rsp *http.Response) (*CtxAccessCheckHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxAccessCheckHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextAccessDecisionResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxAccessCheckHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxAccessCheckHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxAccessDescribeRequirementsHTTP parses an HTTP response from a CtxAccessDescribeRequirementsWithResponse call
func ParseCtxAccessDescribeRequirementsHTTP(rsp *http.Response) (*CtxAccessDescribeRequirementsHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxAccessDescribeRequirementsHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextAccessRequirementsResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxAccessDescribeRequirementsHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxAccessDescribeRequirementsHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxAccessRulesHTTP parses an HTTP response from a CtxAccessRulesWithResponse call
func ParseCtxAccessRulesHTTP(rsp *http.Response) (*CtxAccessRulesHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxAccessRulesHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CtxAccessRulesResponseBody
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxAccessRulesHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxAccessRulesHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxCrosswalkDescribeHTTP parses an HTTP response from a CtxCrosswalkDescribeWithResponse call
func ParseCtxCrosswalkDescribeHTTP(rsp *http.Response) (*CtxCrosswalkDescribeHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxCrosswalkDescribeHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest CtxCrosswalkDescribeResponseBody
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxCrosswalkDescribeHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxCrosswalkDescribeHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxCrosswalkTranslateHTTP parses an HTTP response from a CtxCrosswalkTranslateWithResponse call
func ParseCtxCrosswalkTranslateHTTP(rsp *http.Response) (*CtxCrosswalkTranslateHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxCrosswalkTranslateHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextCrosswalkTranslationResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxCrosswalkTranslateHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxCrosswalkTranslateHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxDocumentAuditHTTP parses an HTTP response from a CtxDocumentAuditWithResponse call
func ParseCtxDocumentAuditHTTP(rsp *http.Response) (*CtxDocumentAuditHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxDocumentAuditHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextDocumentAuditResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 409:
		var dest Conflict
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON409 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxDocumentAuditHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxDocumentAuditHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxGuideEvaluateHTTP parses an HTTP response from a CtxGuideEvaluateWithResponse call
func ParseCtxGuideEvaluateHTTP(rsp *http.Response) (*CtxGuideEvaluateHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxGuideEvaluateHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextGuideEvaluationResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxGuideEvaluateHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxGuideEvaluateHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxGuideScanHTTP parses an HTTP response from a CtxGuideScanWithResponse call
func ParseCtxGuideScanHTTP(rsp *http.Response) (*CtxGuideScanHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxGuideScanHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextGuideScanResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxGuideScanHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxGuideScanHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxGuideViewHTTP parses an HTTP response from a CtxGuideViewWithResponse call
func ParseCtxGuideViewHTTP(rsp *http.Response) (*CtxGuideViewHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxGuideViewHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextGuideViewResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest UnprocessableEntity
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxGuideViewHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxGuideViewHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeSummaryHTTP parses an HTTP response from a CtxSchemeSummaryWithResponse call
func ParseCtxSchemeSummaryHTTP(rsp *http.Response) (*CtxSchemeSummaryHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeSummaryHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SchemeSummary
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeSummaryHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeSummaryHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeAggregateHTTP parses an HTTP response from a CtxSchemeAggregateWithResponse call
func ParseCtxSchemeAggregateHTTP(rsp *http.Response) (*CtxSchemeAggregateHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeAggregateHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeAggregateResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeAggregateHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeAggregateHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeApplyPresetHTTP parses an HTTP response from a CtxSchemeApplyPresetWithResponse call
func ParseCtxSchemeApplyPresetHTTP(rsp *http.Response) (*CtxSchemeApplyPresetHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeApplyPresetHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeValidateResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeApplyPresetHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeApplyPresetHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeCategoriesHTTP parses an HTTP response from a CtxSchemeCategoriesWithResponse call
func ParseCtxSchemeCategoriesHTTP(rsp *http.Response) (*CtxSchemeCategoriesHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeCategoriesHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []CategoryView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeCategoriesHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeCategoriesHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeCheckCoverageHTTP parses an HTTP response from a CtxSchemeCheckCoverageWithResponse call
func ParseCtxSchemeCheckCoverageHTTP(rsp *http.Response) (*CtxSchemeCheckCoverageHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeCheckCoverageHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeCoverageResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeCheckCoverageHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeCheckCoverageHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeConceptsHTTP parses an HTTP response from a CtxSchemeConceptsWithResponse call
func ParseCtxSchemeConceptsHTTP(rsp *http.Response) (*CtxSchemeConceptsHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeConceptsHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []ConceptView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeConceptsHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeConceptsHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeDescribeConceptHTTP parses an HTTP response from a CtxSchemeDescribeConceptWithResponse call
func ParseCtxSchemeDescribeConceptHTTP(rsp *http.Response) (*CtxSchemeDescribeConceptHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeDescribeConceptHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeConceptDescriptionResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeDescribeConceptHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeDescribeConceptHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeDiffHTTP parses an HTTP response from a CtxSchemeDiffWithResponse call
func ParseCtxSchemeDiffHTTP(rsp *http.Response) (*CtxSchemeDiffHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeDiffHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeDiffResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeDiffHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeDiffHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeExplainHTTP parses an HTTP response from a CtxSchemeExplainWithResponse call
func ParseCtxSchemeExplainHTTP(rsp *http.Response) (*CtxSchemeExplainHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeExplainHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeExplainResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeExplainHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeExplainHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeFormatsHTTP parses an HTTP response from a CtxSchemeFormatsWithResponse call
func ParseCtxSchemeFormatsHTTP(rsp *http.Response) (*CtxSchemeFormatsHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeFormatsHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []FormatView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeFormatsHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeFormatsHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeLabelPresetsHTTP parses an HTTP response from a CtxSchemeLabelPresetsWithResponse call
func ParseCtxSchemeLabelPresetsHTTP(rsp *http.Response) (*CtxSchemeLabelPresetsHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeLabelPresetsHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []LabelPresetView
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeLabelPresetsHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeLabelPresetsHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeNoticesHTTP parses an HTTP response from a CtxSchemeNoticesWithResponse call
func ParseCtxSchemeNoticesHTTP(rsp *http.Response) (*CtxSchemeNoticesHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeNoticesHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeNoticesResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeNoticesHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeNoticesHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeProcessHTTP parses an HTTP response from a CtxSchemeProcessWithResponse call
func ParseCtxSchemeProcessHTTP(rsp *http.Response) (*CtxSchemeProcessHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeProcessHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeProcessResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeProcessHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeProcessHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeRenderHTTP parses an HTTP response from a CtxSchemeRenderWithResponse call
func ParseCtxSchemeRenderHTTP(rsp *http.Response) (*CtxSchemeRenderHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeRenderHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeRenderResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeRenderHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeRenderHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeScanHTTP parses an HTTP response from a CtxSchemeScanWithResponse call
func ParseCtxSchemeScanHTTP(rsp *http.Response) (*CtxSchemeScanHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeScanHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeScanResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeScanHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeScanHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeScopesHTTP parses an HTTP response from a CtxSchemeScopesWithResponse call
func ParseCtxSchemeScopesHTTP(rsp *http.Response) (*CtxSchemeScopesHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeScopesHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest SchemeScopesResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeScopesHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeScopesHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeSuggestHTTP parses an HTTP response from a CtxSchemeSuggestWithResponse call
func ParseCtxSchemeSuggestHTTP(rsp *http.Response) (*CtxSchemeSuggestHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeSuggestHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeSuggestResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeSuggestHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeSuggestHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeSuggestForCategoriesHTTP parses an HTTP response from a CtxSchemeSuggestForCategoriesWithResponse call
func ParseCtxSchemeSuggestForCategoriesHTTP(rsp *http.Response) (*CtxSchemeSuggestForCategoriesHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeSuggestForCategoriesHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeSuggestForCategoriesResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeSuggestForCategoriesHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeSuggestForCategoriesHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeValidateHTTP parses an HTTP response from a CtxSchemeValidateWithResponse call
func ParseCtxSchemeValidateHTTP(rsp *http.Response) (*CtxSchemeValidateHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeValidateHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeValidateResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeValidateHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeValidateHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseCtxSchemeViewHTTP parses an HTTP response from a CtxSchemeViewWithResponse call
func ParseCtxSchemeViewHTTP(rsp *http.Response) (*CtxSchemeViewHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CtxSchemeViewHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ContextSchemeViewResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest NotFound
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest InternalError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	switch {
	case rsp.StatusCode == 200:
		var headers CtxSchemeViewHTTP200Headers
		if values := rsp.Header.Values("X-EzClassify-Resolved-Version"); len(values) > 0 {
			var value string
			if err := runtime.BindStyledParameterWithOptions("simple", "X-EzClassify-Resolved-Version", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "string", Format: ""}); err != nil {
				return nil, err
			}
			headers.XEzClassifyResolvedVersion = &value
		}
		response.Headers200 = &headers
	case rsp.StatusCode == 429:
		var headers CtxSchemeViewHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseGetHealthHTTP parses an HTTP response from a GetHealthWithResponse call
func ParseGetHealthHTTP(rsp *http.Response) (*GetHealthHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetHealthHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest HealthResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	}

	switch {
	case rsp.StatusCode == 429:
		var headers GetHealthHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}

// ParseGetReadyHTTP parses an HTTP response from a GetReadyWithResponse call
func ParseGetReadyHTTP(rsp *http.Response) (*GetReadyHTTP, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetReadyHTTP{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ReadyResponse
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest BadRequest
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 429:
		var dest TooManyRequests
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON429 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 503:
		var dest NotReady
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON503 = &dest

	}

	switch {
	case rsp.StatusCode == 429:
		var headers GetReadyHTTP429Headers
		if values := rsp.Header.Values("Retry-After"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "Retry-After", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.RetryAfter = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Limit"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Limit", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitLimit = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Remaining"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Remaining", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitRemaining = &value
		}
		if values := rsp.Header.Values("X-RateLimit-Reset"); len(values) > 0 {
			var value int
			if err := runtime.BindStyledParameterWithOptions("simple", "X-RateLimit-Reset", values[0], &value, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: false, Type: "integer", Format: ""}); err != nil {
				return nil, err
			}
			headers.XRateLimitReset = &value
		}
		response.Headers429 = &headers
	}

	return response, nil
}
