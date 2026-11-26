package revisions

import (
	"encoding/json"
	"time"
)

// SkillsSecurityReviewRevisionsListOutputItemsFiles - What the review covered
type SkillsSecurityReviewRevisionsListOutputItemsFiles struct {
	// Total - Files in the reviewed version
	Total float64 `json:"total"`
	// Analyzed - Text files whose content was analysed
	Analyzed float64 `json:"analyzed"`
	// Kinds - File count per detected type. Binary files are classified by their signature only and never opened
	Kinds map[string]float64 `json:"kinds"`
}

// SkillsSecurityReviewRevisionsListOutputItemsScoresFindingCounts represents the skills security review revisions list output items scores finding counts type.
type SkillsSecurityReviewRevisionsListOutputItemsScoresFindingCounts struct {
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

// SkillsSecurityReviewRevisionsListOutputItemsScores - Computed scores; null until the review has an assessment
type SkillsSecurityReviewRevisionsListOutputItemsScores struct {
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
	FindingCounts SkillsSecurityReviewRevisionsListOutputItemsScoresFindingCounts `json:"finding_counts"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportSummary represents the skills security review revisions list output items report summary type.
type SkillsSecurityReviewRevisionsListOutputItemsReportSummary struct {
	// Text - Plain-language summary of the skill
	Text string `json:"text"`
	// TrustNote - Why the skill can or cannot be trusted
	TrustNote string `json:"trust_note"`
	// ExposureNote - What the skill could make an agent do
	ExposureNote string `json:"exposure_note"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportProvenance represents the skills security review revisions list output items report provenance type.
type SkillsSecurityReviewRevisionsListOutputItemsReportProvenance struct {
	// Class - Where the skill came from
	Class string `json:"class"`
	// Body - Explanation of the provenance
	Body string `json:"body"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportBehaviorsEvidence represents the skills security review revisions list output items report behaviors evidence type.
type SkillsSecurityReviewRevisionsListOutputItemsReportBehaviorsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportBehaviors represents the skills security review revisions list output items report behaviors type.
type SkillsSecurityReviewRevisionsListOutputItemsReportBehaviors struct {
	// Tag - What the skill makes an agent do
	Tag string `json:"tag"`
	// Body - How the skill does it
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsListOutputItemsReportBehaviorsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengthsEvidence represents the skills security review revisions list output items report implementation strengths evidence type.
type SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengthsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengths represents the skills security review revisions list output items report implementation strengths type.
type SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengths struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengthsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknessesEvidence represents the skills security review revisions list output items report implementation weaknesses evidence type.
type SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknessesEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknesses represents the skills security review revisions list output items report implementation weaknesses type.
type SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknesses struct {
	// Title - Short title of the observation
	Title string `json:"title"`
	// Body - Explanation of the observation
	Body string `json:"body"`
	// Evidence - Files of the skill the statement is based on
	Evidence []SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknessesEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportImplementation represents the skills security review revisions list output items report implementation type.
type SkillsSecurityReviewRevisionsListOutputItemsReportImplementation struct {
	Strengths  []SkillsSecurityReviewRevisionsListOutputItemsReportImplementationStrengths  `json:"strengths"`
	Weaknesses []SkillsSecurityReviewRevisionsListOutputItemsReportImplementationWeaknesses `json:"weaknesses"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportResolvedFindings represents the skills security review revisions list output items report resolved findings type.
type SkillsSecurityReviewRevisionsListOutputItemsReportResolvedFindings struct {
	// Key - Key of a finding from the previous revision
	Key string `json:"key"`
	// Reason - Why it no longer applies
	Reason string `json:"reason"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportChanges represents the skills security review revisions list output items report changes type.
type SkillsSecurityReviewRevisionsListOutputItemsReportChanges struct {
	// Body - What changed since the previous revision
	Body string `json:"body"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReportMethod represents the skills security review revisions list output items report method type.
type SkillsSecurityReviewRevisionsListOutputItemsReportMethod struct {
	// Included - What the review covered
	Included []string `json:"included"`
	// OutOfScope - What the review did not cover
	OutOfScope []string `json:"out_of_scope"`
	// Limitations - Known gaps in the review
	Limitations []string `json:"limitations"`
}

// SkillsSecurityReviewRevisionsListOutputItemsReport - Report sections; sections that are not written yet are null
type SkillsSecurityReviewRevisionsListOutputItemsReport struct {
	Summary    *SkillsSecurityReviewRevisionsListOutputItemsReportSummary    `json:"summary,omitempty"`
	Provenance *SkillsSecurityReviewRevisionsListOutputItemsReportProvenance `json:"provenance,omitempty"`
	// Behaviors - Capabilities the skill exercises
	Behaviors        []SkillsSecurityReviewRevisionsListOutputItemsReportBehaviors        `json:"behaviors"`
	Implementation   *SkillsSecurityReviewRevisionsListOutputItemsReportImplementation    `json:"implementation,omitempty"`
	ResolvedFindings []SkillsSecurityReviewRevisionsListOutputItemsReportResolvedFindings `json:"resolved_findings"`
	Changes          *SkillsSecurityReviewRevisionsListOutputItemsReportChanges           `json:"changes,omitempty"`
	Method           *SkillsSecurityReviewRevisionsListOutputItemsReportMethod            `json:"method,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutputItemsFindingsEvidence represents the skills security review revisions list output items findings evidence type.
type SkillsSecurityReviewRevisionsListOutputItemsFindingsEvidence struct {
	Type string `json:"type"`
	// Path - Path of the file within the skill
	Path string `json:"path"`
	// Lines - Start and end line of the cited range
	Lines *[]float64 `json:"lines,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutputItemsFindings represents the skills security review revisions list output items findings type.
type SkillsSecurityReviewRevisionsListOutputItemsFindings struct {
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
	Evidence []SkillsSecurityReviewRevisionsListOutputItemsFindingsEvidence `json:"evidence"`
}

// SkillsSecurityReviewRevisionsListOutputItems represents the skills security review revisions list output items type.
type SkillsSecurityReviewRevisionsListOutputItems struct {
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
	Files *SkillsSecurityReviewRevisionsListOutputItemsFiles `json:"files,omitempty"`
	// Scores - Computed scores; null until the review has an assessment
	Scores *SkillsSecurityReviewRevisionsListOutputItemsScores `json:"scores,omitempty"`
	// Report - Report sections; sections that are not written yet are null
	Report   SkillsSecurityReviewRevisionsListOutputItemsReport     `json:"report"`
	Findings []SkillsSecurityReviewRevisionsListOutputItemsFindings `json:"findings"`
	// StartedAt - When the review started running
	StartedAt *time.Time `json:"started_at,omitempty"`
	// CompletedAt - When the review finished
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - When the revision was created
	CreatedAt time.Time `json:"created_at"`
}

// SkillsSecurityReviewRevisionsListOutputPagination represents the skills security review revisions list output pagination type.
type SkillsSecurityReviewRevisionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsSecurityReviewRevisionsListOutput represents the skills security review revisions list output type.
type SkillsSecurityReviewRevisionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                            `json:"object"`
	Items      []SkillsSecurityReviewRevisionsListOutputItems    `json:"items"`
	Pagination SkillsSecurityReviewRevisionsListOutputPagination `json:"pagination"`
}

// MapSkillsSecurityReviewRevisionsListOutputFromJSON deserializes JSON data into a SkillsSecurityReviewRevisionsListOutput.
func MapSkillsSecurityReviewRevisionsListOutputFromJSON(data []byte) (*SkillsSecurityReviewRevisionsListOutput, error) {
	var v SkillsSecurityReviewRevisionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewRevisionsListOutputToJSON serializes a SkillsSecurityReviewRevisionsListOutput to JSON.
func MapSkillsSecurityReviewRevisionsListOutputToJSON(v *SkillsSecurityReviewRevisionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsSecurityReviewRevisionsListQuery represents the skills security review revisions list query type.
type SkillsSecurityReviewRevisionsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
}

// MapSkillsSecurityReviewRevisionsListQueryFromJSON deserializes JSON data into a SkillsSecurityReviewRevisionsListQuery.
func MapSkillsSecurityReviewRevisionsListQueryFromJSON(data []byte) (*SkillsSecurityReviewRevisionsListQuery, error) {
	var v SkillsSecurityReviewRevisionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsSecurityReviewRevisionsListQueryToJSON serializes a SkillsSecurityReviewRevisionsListQuery to JSON.
func MapSkillsSecurityReviewRevisionsListQueryToJSON(v *SkillsSecurityReviewRevisionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
