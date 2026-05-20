package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/groups"
)

// SkillsGroupsEndpoint provides access to skill groups organize skills into reusable collections.
type SkillsGroupsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsGroupsEndpoint creates a new SkillsGroupsEndpoint.
func NewSkillsGroupsEndpoint(client *endpoint.Client) *SkillsGroupsEndpoint {
	return &SkillsGroupsEndpoint{client: client}
}

// SkillsGroupsEndpointListParams contains optional query parameters for List.
type SkillsGroupsEndpointListParams struct {
	Limit   *float64 `json:"limit,omitempty"`
	After   *string  `json:"after,omitempty"`
	Before  *string  `json:"before,omitempty"`
	Cursor  *string  `json:"cursor,omitempty"`
	Order   *string  `json:"order,omitempty"`
	Search  *string  `json:"search,omitempty"`
	Status  *any     `json:"status,omitempty"`
	Id      *any     `json:"id,omitempty"`
	SkillId *any     `json:"skill_id,omitempty"`
	// CreatedAt - Filter skill group creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill group last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsGroupsEndpointCreateBody contains the request body for Create.
type SkillsGroupsEndpointCreateBody struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	SkillIds    *[]string       `json:"skill_ids,omitempty"`
}

// SkillsGroupsEndpointUpdateBody contains the request body for Update.
type SkillsGroupsEndpointUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	SkillIds    *[]string       `json:"skill_ids,omitempty"`
}

// List returns a paginated list of skill groups.
func (e *SkillsGroupsEndpoint) List(params *SkillsGroupsEndpointListParams) (*groups.SkillsGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-groups"},
		Query: query,
	}
	var result groups.SkillsGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill group.
func (e *SkillsGroupsEndpoint) Get(skillGroupId string) (*groups.SkillsGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
	}
	var result groups.SkillsGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill group.
func (e *SkillsGroupsEndpoint) Create(body *SkillsGroupsEndpointCreateBody) (*groups.SkillsGroupsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups"},
		Body: body,
	}
	var result groups.SkillsGroupsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill group.
func (e *SkillsGroupsEndpoint) Update(skillGroupId string, body *SkillsGroupsEndpointUpdateBody) (*groups.SkillsGroupsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
		Body: body,
	}
	var result groups.SkillsGroupsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill group.
func (e *SkillsGroupsEndpoint) Delete(skillGroupId string) (*groups.SkillsGroupsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-groups", skillGroupId},
	}
	var result groups.SkillsGroupsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
