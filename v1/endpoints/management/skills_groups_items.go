package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/groups/items"
)

// SkillsGroupsItemsEndpoint provides access to skill group items link groups to skills.
type SkillsGroupsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsGroupsItemsEndpoint creates a new SkillsGroupsItemsEndpoint.
func NewSkillsGroupsItemsEndpoint(client *endpoint.Client) *SkillsGroupsItemsEndpoint {
	return &SkillsGroupsItemsEndpoint{client: client}
}

// SkillsGroupsItemsEndpointListParams contains optional query parameters for List.
type SkillsGroupsItemsEndpointListParams struct {
	Limit   *float64 `json:"limit,omitempty"`
	After   *string  `json:"after,omitempty"`
	Before  *string  `json:"before,omitempty"`
	Cursor  *string  `json:"cursor,omitempty"`
	Order   *string  `json:"order,omitempty"`
	Status  *any     `json:"status,omitempty"`
	Id      *any     `json:"id,omitempty"`
	SkillId *any     `json:"skill_id,omitempty"`
	// CreatedAt - Filter skill group item creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// SkillsGroupsItemsEndpointCreateBody contains the request body for Create.
type SkillsGroupsItemsEndpointCreateBody struct {
	SkillId string `json:"skill_id"`
}

// List returns a paginated list of items for a skill group.
func (e *SkillsGroupsItemsEndpoint) List(instanceId string, skillGroupId string, params *SkillsGroupsItemsEndpointListParams) (*items.SkillsGroupsItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-groups", skillGroupId, "items"},
		Query: query,
	}
	var result items.SkillsGroupsItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill group item.
func (e *SkillsGroupsItemsEndpoint) Get(instanceId string, skillGroupId string, skillGroupItemId string) (*items.SkillsGroupsItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-groups", skillGroupId, "items", skillGroupItemId},
	}
	var result items.SkillsGroupsItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a skill to a skill group.
func (e *SkillsGroupsItemsEndpoint) Create(instanceId string, skillGroupId string, body *SkillsGroupsItemsEndpointCreateBody) (*items.SkillsGroupsItemsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-groups", skillGroupId, "items"},
		Body: body,
	}
	var result items.SkillsGroupsItemsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill group item.
func (e *SkillsGroupsItemsEndpoint) Delete(instanceId string, skillGroupId string, skillGroupItemId string) (*items.SkillsGroupsItemsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-groups", skillGroupId, "items", skillGroupItemId},
	}
	var result items.SkillsGroupsItemsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
