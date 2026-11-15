package securityreview

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewDismissOutputCurrentRevisionFiles - What the review covered
type SkillsSecurityReviewDismissOutputCurrentRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionScoresFindingCounts represents the skills security review dismiss output current revision scores finding counts type.
type SkillsSecurityReviewDismissOutputCurrentRevisionScoresFindingCounts struct {
	// Info - Informational findings
	Info float64 `json:"info"`
	// Low - Low severity findings
	Low float64 `json:"low"`
	// Medium - Medium severity findings
	Medium float64 `json:"medium"`
	// High - High severity findings
	High float64 `json:"high"`
	// Critical - Critical findings
	Critical float64 `json:"critical"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewDismissOutputCurrentRevisionScores struct {
	// SecurityScore - How trustworthy and well-built the skill is, from 0 to 100
	SecurityScore float64 `json:"security_score"`
	// Band - Label for the security score range
	Band     string `json:"band"`
	Exposure string `json:"exposure"`
	// Risk - Overall risk of using the skill, combining security and exposure
	Risk string `json:"risk"`
	// Confidence - How much of the skill the reviewer could verify
	Confidence string `json:"confidence"`
	// LethalTrifecta - Whether the skill combines access to private data or credentials with sending data out
	LethalTrifecta bool `json:"lethal_trifecta"`
	// Headline - Two-part headline: the verdict and the main reason
	Headline      []string                                                            `json:"headline"`
	FindingCounts SkillsSecurityReviewDismissOutputCurrentRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportSummary represents the skills security review dismiss output current revision report summary type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportProvenance represents the skills security review dismiss output current revision report provenance type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviorsEvidence represents the skills security review dismiss output current revision report behaviors evidence type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviors represents the skills security review dismiss output current revision report behaviors type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengthsEvidence represents the skills security review dismiss output current revision report implementation strengths evidence type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengths represents the skills security review dismiss output current revision report implementation strengths type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknessesEvidence represents the skills security review dismiss output current revision report implementation weaknesses evidence type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknesses represents the skills security review dismiss output current revision report implementation weaknesses type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementation represents the skills security review dismiss output current revision report implementation type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportResolvedFindings represents the skills security review dismiss output current revision report resolved findings type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportChanges represents the skills security review dismiss output current revision report changes type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReportMethod represents the skills security review dismiss output current revision report method type.
type SkillsSecurityReviewDismissOutputCurrentRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewDismissOutputCurrentRevisionReport struct {
	Summary    *SkillsSecurityReviewDismissOutputCurrentRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewDismissOutputCurrentRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewDismissOutputCurrentRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewDismissOutputCurrentRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewDismissOutputCurrentRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewDismissOutputCurrentRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewDismissOutputCurrentRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionFindingsEvidence represents the skills security review dismiss output current revision findings evidence type.
type SkillsSecurityReviewDismissOutputCurrentRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputCurrentRevisionFindings represents the skills security review dismiss output current revision findings type.
type SkillsSecurityReviewDismissOutputCurrentRevisionFindings struct {
	// Key - Stable key of the finding across revisions
	Key string `json:"key"`
	// Kind - Security issues lower the score, capabilities describe dangerous operations, positives are good practice
	Kind string `json:"kind"`
	// Category - Finding category
	Category string `json:"category"`
	Severity string `json:"severity"`
	// Change - How the finding relates to the previous revision
	Change string `json:"change"`
	// Title - Short title
	Title string `json:"title"`
	// Body - Explanation
	Body string `json:"body"`
	// Recommendation - Suggested fix
	Recommendation *string `json:"recommendation,omitempty"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputCurrentRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputCurrentRevision represents the skills security review dismiss output current revision type.
type SkillsSecurityReviewDismissOutputCurrentRevision struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique identifier of the revision
	Id string `json:"id"`
	// SecurityReviewId - The security review this revision belongs to
	SecurityReviewId string `json:"security_review_id"`
	// SkillVersionId - The skill version that was reviewed
	SkillVersionId *string `json:"skill_version_id,omitempty"`
	// Number - Sequential revision number within the review
	Number float64 `json:"number"`
	// Status - Lifecycle state; `completed_with_gaps` means some files could not be checked
	Status string `json:"status"`
	// Stage - Current stage while the revision is running
	Stage *string `json:"stage,omitempty"`
	// BlockReason - Why the AI review could not run, when only automated checks ran
	BlockReason *string `json:"block_reason,omitempty"`
	// AiGenerated - Whether an AI reviewer read the skill. Otherwise only automated pattern and file-type checks ran
	AiGenerated bool `json:"ai_generated"`
	// Trigger - What caused this revision to be created
	Trigger string `json:"trigger"`
	// Files - What the review covered
	Files *SkillsSecurityReviewDismissOutputCurrentRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewDismissOutputCurrentRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewDismissOutputCurrentRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewDismissOutputCurrentRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionFiles - What the review covered
type SkillsSecurityReviewDismissOutputInProgressRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionScoresFindingCounts represents the skills security review dismiss output in progress revision scores finding counts type.
type SkillsSecurityReviewDismissOutputInProgressRevisionScoresFindingCounts struct {
	// Info - Informational findings
	Info float64 `json:"info"`
	// Low - Low severity findings
	Low float64 `json:"low"`
	// Medium - Medium severity findings
	Medium float64 `json:"medium"`
	// High - High severity findings
	High float64 `json:"high"`
	// Critical - Critical findings
	Critical float64 `json:"critical"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewDismissOutputInProgressRevisionScores struct {
	// SecurityScore - How trustworthy and well-built the skill is, from 0 to 100
	SecurityScore float64 `json:"security_score"`
	// Band - Label for the security score range
	Band     string `json:"band"`
	Exposure string `json:"exposure"`
	// Risk - Overall risk of using the skill, combining security and exposure
	Risk string `json:"risk"`
	// Confidence - How much of the skill the reviewer could verify
	Confidence string `json:"confidence"`
	// LethalTrifecta - Whether the skill combines access to private data or credentials with sending data out
	LethalTrifecta bool `json:"lethal_trifecta"`
	// Headline - Two-part headline: the verdict and the main reason
	Headline      []string                                                               `json:"headline"`
	FindingCounts SkillsSecurityReviewDismissOutputInProgressRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportSummary represents the skills security review dismiss output in progress revision report summary type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportProvenance represents the skills security review dismiss output in progress revision report provenance type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviorsEvidence represents the skills security review dismiss output in progress revision report behaviors evidence type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviors represents the skills security review dismiss output in progress revision report behaviors type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengthsEvidence represents the skills security review dismiss output in progress revision report implementation strengths evidence type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengths represents the skills security review dismiss output in progress revision report implementation strengths type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknessesEvidence represents the skills security review dismiss output in progress revision report implementation weaknesses evidence type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknesses represents the skills security review dismiss output in progress revision report implementation weaknesses type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementation represents the skills security review dismiss output in progress revision report implementation type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportResolvedFindings represents the skills security review dismiss output in progress revision report resolved findings type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportChanges represents the skills security review dismiss output in progress revision report changes type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReportMethod represents the skills security review dismiss output in progress revision report method type.
type SkillsSecurityReviewDismissOutputInProgressRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewDismissOutputInProgressRevisionReport struct {
	Summary    *SkillsSecurityReviewDismissOutputInProgressRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewDismissOutputInProgressRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewDismissOutputInProgressRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewDismissOutputInProgressRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewDismissOutputInProgressRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewDismissOutputInProgressRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewDismissOutputInProgressRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionFindingsEvidence represents the skills security review dismiss output in progress revision findings evidence type.
type SkillsSecurityReviewDismissOutputInProgressRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewDismissOutputInProgressRevisionFindings represents the skills security review dismiss output in progress revision findings type.
type SkillsSecurityReviewDismissOutputInProgressRevisionFindings struct {
	// Key - Stable key of the finding across revisions
	Key string `json:"key"`
	// Kind - Security issues lower the score, capabilities describe dangerous operations, positives are good practice
	Kind string `json:"kind"`
	// Category - Finding category
	Category string `json:"category"`
	Severity string `json:"severity"`
	// Change - How the finding relates to the previous revision
	Change string `json:"change"`
	// Title - Short title
	Title string `json:"title"`
	// Body - Explanation
	Body string `json:"body"`
	// Recommendation - Suggested fix
	Recommendation *string `json:"recommendation,omitempty"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewDismissOutputInProgressRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewDismissOutputInProgressRevision represents the skills security review dismiss output in progress revision type.
type SkillsSecurityReviewDismissOutputInProgressRevision struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique identifier of the revision
	Id string `json:"id"`
	// SecurityReviewId - The security review this revision belongs to
	SecurityReviewId string `json:"security_review_id"`
	// SkillVersionId - The skill version that was reviewed
	SkillVersionId *string `json:"skill_version_id,omitempty"`
	// Number - Sequential revision number within the review
	Number float64 `json:"number"`
	// Status - Lifecycle state; `completed_with_gaps` means some files could not be checked
	Status string `json:"status"`
	// Stage - Current stage while the revision is running
	Stage *string `json:"stage,omitempty"`
	// BlockReason - Why the AI review could not run, when only automated checks ran
	BlockReason *string `json:"block_reason,omitempty"`
	// AiGenerated - Whether an AI reviewer read the skill. Otherwise only automated pattern and file-type checks ran
	AiGenerated bool `json:"ai_generated"`
	// Trigger - What caused this revision to be created
	Trigger string `json:"trigger"`
	// Files - What the review covered
	Files *SkillsSecurityReviewDismissOutputInProgressRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewDismissOutputInProgressRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewDismissOutputInProgressRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewDismissOutputInProgressRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewDismissOutput represents the skills security review dismiss output type.
type SkillsSecurityReviewDismissOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique identifier of the security review
	Id string `json:"id"`
	// SkillId - The reviewed skill
	SkillId *string `json:"skill_id,omitempty"`
	// Verdict - `pending` means the review, or the AI part of it, has not finished yet; `dismissed` means a member stopped reviewing this skill
	Verdict string `json:"verdict"`
	// DismissedAt - When the review was dismissed; null while it is active
	DismissedAt        *time.Time                                           `json:"dismissed_at,omitempty"`
	CurrentRevision    *SkillsSecurityReviewDismissOutputCurrentRevision    `json:"current_revision,omitempty"`
	InProgressRevision *SkillsSecurityReviewDismissOutputInProgressRevision `json:"in_progress_revision,omitempty"`
	// CreatedAt - When the review was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the review was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsSecurityReviewDismissOutputFromJSON deserializes JSON data into a SkillsSecurityReviewDismissOutput.
func MapSkillsSecurityReviewDismissOutputFromJSON(data []byte) (*SkillsSecurityReviewDismissOutput, error) {
	var v SkillsSecurityReviewDismissOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewDismissOutputToJSON serializes a SkillsSecurityReviewDismissOutput to JSON.
func MapSkillsSecurityReviewDismissOutputToJSON(v *SkillsSecurityReviewDismissOutput) ([]byte, error) {
	return json.Marshal(v)
}
