package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/templates/items"
)

// SkillsTemplatesItemsEndpoint provides access to a skill template item is a provider or integration attached to a skill template.
type SkillsTemplatesItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsTemplatesItemsEndpoint creates a new SkillsTemplatesItemsEndpoint.
func NewSkillsTemplatesItemsEndpoint(client *endpoint.Client) *SkillsTemplatesItemsEndpoint {
	return &SkillsTemplatesItemsEndpoint{client: client}
}

// SkillsTemplatesItemsEndpointListParams contains optional query parameters for List.
type SkillsTemplatesItemsEndpointListParams struct {
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

// List returns a paginated list of skill template items.
func (e *SkillsTemplatesItemsEndpoint) List(instanceId string, skillTemplateId string, params *SkillsTemplatesItemsEndpointListParams) (*items.SkillsTemplatesItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-templates", skillTemplateId, "items"},
		Query: query,
	}
	var result items.SkillsTemplatesItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill template item.
func (e *SkillsTemplatesItemsEndpoint) Get(instanceId string, skillTemplateId string, skillTemplateItemId string) (*items.SkillsTemplatesItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId, "items", skillTemplateItemId},
	}
	var result items.SkillsTemplatesItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a provider or integration item to a skill template.
func (e *SkillsTemplatesItemsEndpoint) Create(instanceId string, skillTemplateId string) (*items.SkillsTemplatesItemsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId, "items"},
	}
	var result items.SkillsTemplatesItemsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a skill template item.
func (e *SkillsTemplatesItemsEndpoint) Delete(instanceId string, skillTemplateId string, skillTemplateItemId string) (*items.SkillsTemplatesItemsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-templates", skillTemplateId, "items", skillTemplateItemId},
	}
	var result items.SkillsTemplatesItemsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
