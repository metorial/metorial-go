package securityreview

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewGetOutputCurrentRevisionFiles - What the review covered
type SkillsSecurityReviewGetOutputCurrentRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionScoresFindingCounts represents the skills security review get output current revision scores finding counts type.
type SkillsSecurityReviewGetOutputCurrentRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewGetOutputCurrentRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewGetOutputCurrentRevisionScores struct {
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
	Headline      []string                                                        `json:"headline"`
	FindingCounts SkillsSecurityReviewGetOutputCurrentRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportSummary represents the skills security review get output current revision report summary type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportProvenance represents the skills security review get output current revision report provenance type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviorsEvidence represents the skills security review get output current revision report behaviors evidence type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviors represents the skills security review get output current revision report behaviors type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengthsEvidence represents the skills security review get output current revision report implementation strengths evidence type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengths represents the skills security review get output current revision report implementation strengths type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknessesEvidence represents the skills security review get output current revision report implementation weaknesses evidence type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknesses represents the skills security review get output current revision report implementation weaknesses type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportImplementation represents the skills security review get output current revision report implementation type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewGetOutputCurrentRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportResolvedFindings represents the skills security review get output current revision report resolved findings type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportChanges represents the skills security review get output current revision report changes type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReportMethod represents the skills security review get output current revision report method type.
type SkillsSecurityReviewGetOutputCurrentRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewGetOutputCurrentRevisionReport struct {
	Summary    *SkillsSecurityReviewGetOutputCurrentRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewGetOutputCurrentRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewGetOutputCurrentRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewGetOutputCurrentRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewGetOutputCurrentRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewGetOutputCurrentRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewGetOutputCurrentRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionFindingsEvidence represents the skills security review get output current revision findings evidence type.
type SkillsSecurityReviewGetOutputCurrentRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputCurrentRevisionFindings represents the skills security review get output current revision findings type.
type SkillsSecurityReviewGetOutputCurrentRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewGetOutputCurrentRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputCurrentRevision represents the skills security review get output current revision type.
type SkillsSecurityReviewGetOutputCurrentRevision struct {
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
	Files *SkillsSecurityReviewGetOutputCurrentRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewGetOutputCurrentRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewGetOutputCurrentRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewGetOutputCurrentRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionFiles - What the review covered
type SkillsSecurityReviewGetOutputInProgressRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionScoresFindingCounts represents the skills security review get output in progress revision scores finding counts type.
type SkillsSecurityReviewGetOutputInProgressRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewGetOutputInProgressRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewGetOutputInProgressRevisionScores struct {
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
	Headline      []string                                                           `json:"headline"`
	FindingCounts SkillsSecurityReviewGetOutputInProgressRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportSummary represents the skills security review get output in progress revision report summary type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportProvenance represents the skills security review get output in progress revision report provenance type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviorsEvidence represents the skills security review get output in progress revision report behaviors evidence type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviors represents the skills security review get output in progress revision report behaviors type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengthsEvidence represents the skills security review get output in progress revision report implementation strengths evidence type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengths represents the skills security review get output in progress revision report implementation strengths type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknessesEvidence represents the skills security review get output in progress revision report implementation weaknesses evidence type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknesses represents the skills security review get output in progress revision report implementation weaknesses type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportImplementation represents the skills security review get output in progress revision report implementation type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewGetOutputInProgressRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportResolvedFindings represents the skills security review get output in progress revision report resolved findings type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportChanges represents the skills security review get output in progress revision report changes type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReportMethod represents the skills security review get output in progress revision report method type.
type SkillsSecurityReviewGetOutputInProgressRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewGetOutputInProgressRevisionReport struct {
	Summary    *SkillsSecurityReviewGetOutputInProgressRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewGetOutputInProgressRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewGetOutputInProgressRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewGetOutputInProgressRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewGetOutputInProgressRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewGetOutputInProgressRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewGetOutputInProgressRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionFindingsEvidence represents the skills security review get output in progress revision findings evidence type.
type SkillsSecurityReviewGetOutputInProgressRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewGetOutputInProgressRevisionFindings represents the skills security review get output in progress revision findings type.
type SkillsSecurityReviewGetOutputInProgressRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewGetOutputInProgressRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewGetOutputInProgressRevision represents the skills security review get output in progress revision type.
type SkillsSecurityReviewGetOutputInProgressRevision struct {
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
	Files *SkillsSecurityReviewGetOutputInProgressRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewGetOutputInProgressRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewGetOutputInProgressRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewGetOutputInProgressRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewGetOutput represents the skills security review get output type.
type SkillsSecurityReviewGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique identifier of the security review
	Id string `json:"id"`
	// SkillId - The reviewed skill
	SkillId *string `json:"skill_id,omitempty"`
	// Verdict - `pending` means the review, or the AI part of it, has not finished yet; `dismissed` means a member stopped reviewing this skill
	Verdict string `json:"verdict"`
	// DismissedAt - When the review was dismissed; null while it is active
	DismissedAt        *time.Time                                       `json:"dismissed_at,omitempty"`
	CurrentRevision    *SkillsSecurityReviewGetOutputCurrentRevision    `json:"current_revision,omitempty"`
	InProgressRevision *SkillsSecurityReviewGetOutputInProgressRevision `json:"in_progress_revision,omitempty"`
	// CreatedAt - When the review was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the review was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsSecurityReviewGetOutputFromJSON deserializes JSON data into a SkillsSecurityReviewGetOutput.
func MapSkillsSecurityReviewGetOutputFromJSON(data []byte) (*SkillsSecurityReviewGetOutput, error) {
	var v SkillsSecurityReviewGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewGetOutputToJSON serializes a SkillsSecurityReviewGetOutput to JSON.
func MapSkillsSecurityReviewGetOutputToJSON(v *SkillsSecurityReviewGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
