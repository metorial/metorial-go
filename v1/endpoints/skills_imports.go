package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/imports"
)

// SkillsImportsEndpoint provides access to a skill import is a queued import of skills from a repository or uploaded file.
type SkillsImportsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsImportsEndpoint creates a new SkillsImportsEndpoint.
func NewSkillsImportsEndpoint(client *endpoint.Client) *SkillsImportsEndpoint {
	return &SkillsImportsEndpoint{client: client}
}

// SkillsImportsEndpointListParams contains optional query parameters for List.
type SkillsImportsEndpointListParams struct {
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
	// Id - Only return skill imports with these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return skill imports with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// SkillsImportsEndpointCreateBody contains the request body for Create.
type SkillsImportsEndpointCreateBody struct {
	// Source - Where to import skills from.
	Source any `json:"source"`
	// SkillGroupId - The ID of a skill group that every imported skill is added to.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
}

// List returns a paginated list of skill imports.
func (e *SkillsImportsEndpoint) List(params *SkillsImportsEndpointListParams) (*imports.SkillsImportsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-imports"},
		Query: query,
	}
	var result imports.SkillsImportsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill import and its results.
func (e *SkillsImportsEndpoint) Get(skillImportId string) (*imports.SkillsImportsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-imports", skillImportId},
	}
	var result imports.SkillsImportsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create queues a skill import from a repository or uploaded file.
func (e *SkillsImportsEndpoint) Create(body *SkillsImportsEndpointCreateBody) (*imports.SkillsImportsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-imports"},
		Body: body,
	}
	var result imports.SkillsImportsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
