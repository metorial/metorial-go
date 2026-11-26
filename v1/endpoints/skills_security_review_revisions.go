package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/securityreview/revisions"
)

// SkillsSecurityReviewRevisionsEndpoint provides access to a security review assesses what a skill makes an agent do and whether its files can be trusted. Everyone who can read a skill can read its review.
type SkillsSecurityReviewRevisionsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsSecurityReviewRevisionsEndpoint creates a new SkillsSecurityReviewRevisionsEndpoint.
func NewSkillsSecurityReviewRevisionsEndpoint(client *endpoint.Client) *SkillsSecurityReviewRevisionsEndpoint {
	return &SkillsSecurityReviewRevisionsEndpoint{client: client}
}

// SkillsSecurityReviewRevisionsEndpointListParams contains optional query parameters for List.
type SkillsSecurityReviewRevisionsEndpointListParams struct {
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

// List returns the most recent finished revisions of a skill's security review.
func (e *SkillsSecurityReviewRevisionsEndpoint) List(skillId string, params *SkillsSecurityReviewRevisionsEndpointListParams) (*revisions.SkillsSecurityReviewRevisionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "security-review", "revisions"},
		Query: query,
	}
	var result revisions.SkillsSecurityReviewRevisionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a single revision of a skill's security review.
func (e *SkillsSecurityReviewRevisionsEndpoint) Get(skillId string, revisionId string) (*revisions.SkillsSecurityReviewRevisionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "security-review", "revisions", revisionId},
	}
	var result revisions.SkillsSecurityReviewRevisionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
