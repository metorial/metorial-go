package securityreview

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewReassessOutputCurrentRevisionFiles - What the review covered
type SkillsSecurityReviewReassessOutputCurrentRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionScoresFindingCounts represents the skills security review reassess output current revision scores finding counts type.
type SkillsSecurityReviewReassessOutputCurrentRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewReassessOutputCurrentRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewReassessOutputCurrentRevisionScores struct {
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
	Headline      []string                                                             `json:"headline"`
	FindingCounts SkillsSecurityReviewReassessOutputCurrentRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportSummary represents the skills security review reassess output current revision report summary type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportProvenance represents the skills security review reassess output current revision report provenance type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviorsEvidence represents the skills security review reassess output current revision report behaviors evidence type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviors represents the skills security review reassess output current revision report behaviors type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengthsEvidence represents the skills security review reassess output current revision report implementation strengths evidence type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengths represents the skills security review reassess output current revision report implementation strengths type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknessesEvidence represents the skills security review reassess output current revision report implementation weaknesses evidence type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknesses represents the skills security review reassess output current revision report implementation weaknesses type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementation represents the skills security review reassess output current revision report implementation type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportResolvedFindings represents the skills security review reassess output current revision report resolved findings type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportChanges represents the skills security review reassess output current revision report changes type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReportMethod represents the skills security review reassess output current revision report method type.
type SkillsSecurityReviewReassessOutputCurrentRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewReassessOutputCurrentRevisionReport struct {
	Summary    *SkillsSecurityReviewReassessOutputCurrentRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewReassessOutputCurrentRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewReassessOutputCurrentRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewReassessOutputCurrentRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewReassessOutputCurrentRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewReassessOutputCurrentRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewReassessOutputCurrentRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionFindingsEvidence represents the skills security review reassess output current revision findings evidence type.
type SkillsSecurityReviewReassessOutputCurrentRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputCurrentRevisionFindings represents the skills security review reassess output current revision findings type.
type SkillsSecurityReviewReassessOutputCurrentRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewReassessOutputCurrentRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputCurrentRevision represents the skills security review reassess output current revision type.
type SkillsSecurityReviewReassessOutputCurrentRevision struct {
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
	Files *SkillsSecurityReviewReassessOutputCurrentRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewReassessOutputCurrentRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewReassessOutputCurrentRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewReassessOutputCurrentRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionFiles - What the review covered
type SkillsSecurityReviewReassessOutputInProgressRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionScoresFindingCounts represents the skills security review reassess output in progress revision scores finding counts type.
type SkillsSecurityReviewReassessOutputInProgressRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewReassessOutputInProgressRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewReassessOutputInProgressRevisionScores struct {
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
	Headline      []string                                                                `json:"headline"`
	FindingCounts SkillsSecurityReviewReassessOutputInProgressRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportSummary represents the skills security review reassess output in progress revision report summary type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportProvenance represents the skills security review reassess output in progress revision report provenance type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviorsEvidence represents the skills security review reassess output in progress revision report behaviors evidence type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviors represents the skills security review reassess output in progress revision report behaviors type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengthsEvidence represents the skills security review reassess output in progress revision report implementation strengths evidence type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengths represents the skills security review reassess output in progress revision report implementation strengths type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknessesEvidence represents the skills security review reassess output in progress revision report implementation weaknesses evidence type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknesses represents the skills security review reassess output in progress revision report implementation weaknesses type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementation represents the skills security review reassess output in progress revision report implementation type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportResolvedFindings represents the skills security review reassess output in progress revision report resolved findings type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportChanges represents the skills security review reassess output in progress revision report changes type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReportMethod represents the skills security review reassess output in progress revision report method type.
type SkillsSecurityReviewReassessOutputInProgressRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewReassessOutputInProgressRevisionReport struct {
	Summary    *SkillsSecurityReviewReassessOutputInProgressRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewReassessOutputInProgressRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewReassessOutputInProgressRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewReassessOutputInProgressRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewReassessOutputInProgressRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewReassessOutputInProgressRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewReassessOutputInProgressRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionFindingsEvidence represents the skills security review reassess output in progress revision findings evidence type.
type SkillsSecurityReviewReassessOutputInProgressRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewReassessOutputInProgressRevisionFindings represents the skills security review reassess output in progress revision findings type.
type SkillsSecurityReviewReassessOutputInProgressRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewReassessOutputInProgressRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewReassessOutputInProgressRevision represents the skills security review reassess output in progress revision type.
type SkillsSecurityReviewReassessOutputInProgressRevision struct {
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
	Files *SkillsSecurityReviewReassessOutputInProgressRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewReassessOutputInProgressRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewReassessOutputInProgressRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewReassessOutputInProgressRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewReassessOutput represents the skills security review reassess output type.
type SkillsSecurityReviewReassessOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique identifier of the security review
	Id string `json:"id"`
	// SkillId - The reviewed skill
	SkillId *string `json:"skill_id,omitempty"`
	// Verdict - `pending` means the review, or the AI part of it, has not finished yet; `dismissed` means a member stopped reviewing this skill
	Verdict string `json:"verdict"`
	// DismissedAt - When the review was dismissed; null while it is active
	DismissedAt        *time.Time                                            `json:"dismissed_at,omitempty"`
	CurrentRevision    *SkillsSecurityReviewReassessOutputCurrentRevision    `json:"current_revision,omitempty"`
	InProgressRevision *SkillsSecurityReviewReassessOutputInProgressRevision `json:"in_progress_revision,omitempty"`
	// CreatedAt - When the review was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the review was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsSecurityReviewReassessOutputFromJSON deserializes JSON data into a SkillsSecurityReviewReassessOutput.
func MapSkillsSecurityReviewReassessOutputFromJSON(data []byte) (*SkillsSecurityReviewReassessOutput, error) {
	var v SkillsSecurityReviewReassessOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewReassessOutputToJSON serializes a SkillsSecurityReviewReassessOutput to JSON.
func MapSkillsSecurityReviewReassessOutputToJSON(v *SkillsSecurityReviewReassessOutput) ([]byte, error) {
	return json.Marshal(v)
}
