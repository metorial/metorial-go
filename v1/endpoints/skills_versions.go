package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/versions"
)

// SkillsVersionsEndpoint provides access to a skill version is a snapshot of a skill at a point in time.
type SkillsVersionsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsVersionsEndpoint creates a new SkillsVersionsEndpoint.
func NewSkillsVersionsEndpoint(client *endpoint.Client) *SkillsVersionsEndpoint {
	return &SkillsVersionsEndpoint{client: client}
}

// SkillsVersionsEndpointListParams contains optional query parameters for List.
type SkillsVersionsEndpointListParams struct {
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

// List returns a paginated list of skill versions.
func (e *SkillsVersionsEndpoint) List(skillId string, params *SkillsVersionsEndpointListParams) (*versions.SkillsVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "versions"},
		Query: query,
	}
	var result versions.SkillsVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill version.
func (e *SkillsVersionsEndpoint) Get(skillId string, skillVersionId string) (*versions.SkillsVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "versions", skillVersionId},
	}
	var result versions.SkillsVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
