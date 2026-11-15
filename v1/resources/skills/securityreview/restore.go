package securityreview

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewRestoreOutputCurrentRevisionFiles - What the review covered
type SkillsSecurityReviewRestoreOutputCurrentRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionScoresFindingCounts represents the skills security review restore output current revision scores finding counts type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewRestoreOutputCurrentRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewRestoreOutputCurrentRevisionScores struct {
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
	FindingCounts SkillsSecurityReviewRestoreOutputCurrentRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportSummary represents the skills security review restore output current revision report summary type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportProvenance represents the skills security review restore output current revision report provenance type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviorsEvidence represents the skills security review restore output current revision report behaviors evidence type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviors represents the skills security review restore output current revision report behaviors type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengthsEvidence represents the skills security review restore output current revision report implementation strengths evidence type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengths represents the skills security review restore output current revision report implementation strengths type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknessesEvidence represents the skills security review restore output current revision report implementation weaknesses evidence type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknesses represents the skills security review restore output current revision report implementation weaknesses type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementation represents the skills security review restore output current revision report implementation type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportResolvedFindings represents the skills security review restore output current revision report resolved findings type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportChanges represents the skills security review restore output current revision report changes type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReportMethod represents the skills security review restore output current revision report method type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewRestoreOutputCurrentRevisionReport struct {
	Summary    *SkillsSecurityReviewRestoreOutputCurrentRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewRestoreOutputCurrentRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewRestoreOutputCurrentRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewRestoreOutputCurrentRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewRestoreOutputCurrentRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewRestoreOutputCurrentRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewRestoreOutputCurrentRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionFindingsEvidence represents the skills security review restore output current revision findings evidence type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevisionFindings represents the skills security review restore output current revision findings type.
type SkillsSecurityReviewRestoreOutputCurrentRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewRestoreOutputCurrentRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputCurrentRevision represents the skills security review restore output current revision type.
type SkillsSecurityReviewRestoreOutputCurrentRevision struct {
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
	Files *SkillsSecurityReviewRestoreOutputCurrentRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewRestoreOutputCurrentRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewRestoreOutputCurrentRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewRestoreOutputCurrentRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionFiles - What the review covered
type SkillsSecurityReviewRestoreOutputInProgressRevisionFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionScoresFindingCounts represents the skills security review restore output in progress revision scores finding counts type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionScoresFindingCounts struct {
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

// SkillsSecurityReviewRestoreOutputInProgressRevisionScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewRestoreOutputInProgressRevisionScores struct {
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
	FindingCounts SkillsSecurityReviewRestoreOutputInProgressRevisionScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportSummary represents the skills security review restore output in progress revision report summary type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportProvenance represents the skills security review restore output in progress revision report provenance type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviorsEvidence represents the skills security review restore output in progress revision report behaviors evidence type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviors represents the skills security review restore output in progress revision report behaviors type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengthsEvidence represents the skills security review restore output in progress revision report implementation strengths evidence type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengths represents the skills security review restore output in progress revision report implementation strengths type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknessesEvidence represents the skills security review restore output in progress revision report implementation weaknesses evidence type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknesses represents the skills security review restore output in progress revision report implementation weaknesses type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementation represents the skills security review restore output in progress revision report implementation type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementation struct {
	Strengths  []SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportResolvedFindings represents the skills security review restore output in progress revision report resolved findings type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportChanges represents the skills security review restore output in progress revision report changes type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReportMethod represents the skills security review restore output in progress revision report method type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewRestoreOutputInProgressRevisionReport struct {
	Summary    *SkillsSecurityReviewRestoreOutputInProgressRevisionReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewRestoreOutputInProgressRevisionReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewRestoreOutputInProgressRevisionReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewRestoreOutputInProgressRevisionReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewRestoreOutputInProgressRevisionReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewRestoreOutputInProgressRevisionReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewRestoreOutputInProgressRevisionReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionFindingsEvidence represents the skills security review restore output in progress revision findings evidence type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevisionFindings represents the skills security review restore output in progress revision findings type.
type SkillsSecurityReviewRestoreOutputInProgressRevisionFindings struct {
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
	Evidence []SkillsSecurityReviewRestoreOutputInProgressRevisionFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRestoreOutputInProgressRevision represents the skills security review restore output in progress revision type.
type SkillsSecurityReviewRestoreOutputInProgressRevision struct {
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
	Files *SkillsSecurityReviewRestoreOutputInProgressRevisionFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewRestoreOutputInProgressRevisionScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewRestoreOutputInProgressRevisionReport     `json:"report"`
	Findings []SkillsSecurityReviewRestoreOutputInProgressRevisionFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewRestoreOutput represents the skills security review restore output type.
type SkillsSecurityReviewRestoreOutput struct {
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
	CurrentRevision    *SkillsSecurityReviewRestoreOutputCurrentRevision    `json:"current_revision,omitempty"`
	InProgressRevision *SkillsSecurityReviewRestoreOutputInProgressRevision `json:"in_progress_revision,omitempty"`
	// CreatedAt - When the review was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the review was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSkillsSecurityReviewRestoreOutputFromJSON deserializes JSON data into a SkillsSecurityReviewRestoreOutput.
func MapSkillsSecurityReviewRestoreOutputFromJSON(data []byte) (*SkillsSecurityReviewRestoreOutput, error) {
	var v SkillsSecurityReviewRestoreOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewRestoreOutputToJSON serializes a SkillsSecurityReviewRestoreOutput to JSON.
func MapSkillsSecurityReviewRestoreOutputToJSON(v *SkillsSecurityReviewRestoreOutput) ([]byte, error) {
	return json.Marshal(v)
}
