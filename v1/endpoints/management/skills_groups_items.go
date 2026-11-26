package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/groups/items"
)

// SkillsGroupsItemsEndpoint provides access to a skill group item is one skill in a skill group.
type SkillsGroupsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsGroupsItemsEndpoint creates a new SkillsGroupsItemsEndpoint.
func NewSkillsGroupsItemsEndpoint(client *endpoint.Client) *SkillsGroupsItemsEndpoint {
	return &SkillsGroupsItemsEndpoint{client: client}
}

// SkillsGroupsItemsEndpointListParams contains optional query parameters for List.
type SkillsGroupsItemsEndpointListParams struct {
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
	// Status - Only return skill group items with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return skill group items with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return skill group items for one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// CreatedAt - Only return results whose skill group item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// SkillsGroupsItemsEndpointCreateBody contains the request body for Create.
type SkillsGroupsItemsEndpointCreateBody struct {
	// SkillId - Skill to add to the group, by ID.
	SkillId string `json:"skill_id"`
}

// List returns a paginated list of skill group items.
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

// Get retrieves a skill group item.
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
