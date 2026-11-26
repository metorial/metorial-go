package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/items"
)

// SkillsItemsEndpoint provides access to a skill item is a provider or integration attached to a skill.
type SkillsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsItemsEndpoint creates a new SkillsItemsEndpoint.
func NewSkillsItemsEndpoint(client *endpoint.Client) *SkillsItemsEndpoint {
	return &SkillsItemsEndpoint{client: client}
}

// SkillsItemsEndpointListParams contains optional query parameters for List.
type SkillsItemsEndpointListParams struct {
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
	// Status - Only return skill items with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return skill items of one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return skill items with these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return skill items for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return skill items for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// CreatedAt - Only return results whose skill item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of skill items.
func (e *SkillsItemsEndpoint) List(skillId string, params *SkillsItemsEndpointListParams) (*items.SkillsItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skills", skillId, "items"},
		Query: query,
	}
	var result items.SkillsItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill item.
func (e *SkillsItemsEndpoint) Get(skillId string, skillItemId string) (*items.SkillsItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "items", skillItemId},
	}
	var result items.SkillsItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new item on a skill.
func (e *SkillsItemsEndpoint) Create(skillId string) (*items.SkillsItemsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "items"},
	}
	var result items.SkillsItemsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a skill item.
func (e *SkillsItemsEndpoint) Delete(skillId string, skillItemId string) (*items.SkillsItemsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skills", skillId, "items", skillItemId},
	}
	var result items.SkillsItemsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
