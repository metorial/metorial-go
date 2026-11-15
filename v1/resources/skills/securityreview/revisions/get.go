package revisions

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewRevisionsGetOutputFiles - What the review covered
type SkillsSecurityReviewRevisionsGetOutputFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewRevisionsGetOutputScoresFindingCounts represents the skills security review revisions get output scores finding counts type.
type SkillsSecurityReviewRevisionsGetOutputScoresFindingCounts struct {
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

// SkillsSecurityReviewRevisionsGetOutputScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewRevisionsGetOutputScores struct {
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
	Headline      []string                                                  `json:"headline"`
	FindingCounts SkillsSecurityReviewRevisionsGetOutputScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewRevisionsGetOutputReportSummary represents the skills security review revisions get output report summary type.
type SkillsSecurityReviewRevisionsGetOutputReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewRevisionsGetOutputReportProvenance represents the skills security review revisions get output report provenance type.
type SkillsSecurityReviewRevisionsGetOutputReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewRevisionsGetOutputReportBehaviorsEvidence represents the skills security review revisions get output report behaviors evidence type.
type SkillsSecurityReviewRevisionsGetOutputReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsGetOutputReportBehaviors represents the skills security review revisions get output report behaviors type.
type SkillsSecurityReviewRevisionsGetOutputReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsGetOutputReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengthsEvidence represents the skills security review revisions get output report implementation strengths evidence type.
type SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengths represents the skills security review revisions get output report implementation strengths type.
type SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknessesEvidence represents the skills security review revisions get output report implementation weaknesses evidence type.
type SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknesses represents the skills security review revisions get output report implementation weaknesses type.
type SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsGetOutputReportImplementation represents the skills security review revisions get output report implementation type.
type SkillsSecurityReviewRevisionsGetOutputReportImplementation struct {
	Strengths  []SkillsSecurityReviewRevisionsGetOutputReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewRevisionsGetOutputReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewRevisionsGetOutputReportResolvedFindings represents the skills security review revisions get output report resolved findings type.
type SkillsSecurityReviewRevisionsGetOutputReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewRevisionsGetOutputReportChanges represents the skills security review revisions get output report changes type.
type SkillsSecurityReviewRevisionsGetOutputReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewRevisionsGetOutputReportMethod represents the skills security review revisions get output report method type.
type SkillsSecurityReviewRevisionsGetOutputReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewRevisionsGetOutputReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewRevisionsGetOutputReport struct {
	Summary    *SkillsSecurityReviewRevisionsGetOutputReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewRevisionsGetOutputReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewRevisionsGetOutputReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewRevisionsGetOutputReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewRevisionsGetOutputReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewRevisionsGetOutputReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewRevisionsGetOutputReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewRevisionsGetOutputFindingsEvidence represents the skills security review revisions get output findings evidence type.
type SkillsSecurityReviewRevisionsGetOutputFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsGetOutputFindings represents the skills security review revisions get output findings type.
type SkillsSecurityReviewRevisionsGetOutputFindings struct {
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
	Evidence []SkillsSecurityReviewRevisionsGetOutputFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsGetOutput represents the skills security review revisions get output type.
type SkillsSecurityReviewRevisionsGetOutput struct {
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
	Files *SkillsSecurityReviewRevisionsGetOutputFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewRevisionsGetOutputScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewRevisionsGetOutputReport     `json:"report"`
	Findings []SkillsSecurityReviewRevisionsGetOutputFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// MapSkillsSecurityReviewRevisionsGetOutputFromJSON deserializes JSON data into a SkillsSecurityReviewRevisionsGetOutput.
func MapSkillsSecurityReviewRevisionsGetOutputFromJSON(data []byte) (*SkillsSecurityReviewRevisionsGetOutput, error) {
	var v SkillsSecurityReviewRevisionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewRevisionsGetOutputToJSON serializes a SkillsSecurityReviewRevisionsGetOutput to JSON.
func MapSkillsSecurityReviewRevisionsGetOutputToJSON(v *SkillsSecurityReviewRevisionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
