package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/templates/items"
)

// SkillsTemplatesItemsEndpoint provides access to skill template items link template definitions to provider and integration items.
type SkillsTemplatesItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsTemplatesItemsEndpoint creates a new SkillsTemplatesItemsEndpoint.
func NewSkillsTemplatesItemsEndpoint(client *endpoint.Client) *SkillsTemplatesItemsEndpoint {
	return &SkillsTemplatesItemsEndpoint{client: client}
}

// SkillsTemplatesItemsEndpointListParams contains optional query parameters for List.
type SkillsTemplatesItemsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of items for a skill template.
func (e *SkillsTemplatesItemsEndpoint) List(skillTemplateId string, params *SkillsTemplatesItemsEndpointListParams) (*items.SkillsTemplatesItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-templates", skillTemplateId, "items"},
		Query: query,
	}
	var result items.SkillsTemplatesItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill template item.
func (e *SkillsTemplatesItemsEndpoint) Get(skillTemplateId string, skillTemplateItemId string) (*items.SkillsTemplatesItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId, "items", skillTemplateItemId},
	}
	var result items.SkillsTemplatesItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create adds a provider or integration item to a skill template.
func (e *SkillsTemplatesItemsEndpoint) Create(skillTemplateId string) (*items.SkillsTemplatesItemsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId, "items"},
	}
	var result items.SkillsTemplatesItemsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a skill template item.
func (e *SkillsTemplatesItemsEndpoint) Delete(skillTemplateId string, skillTemplateItemId string) (*items.SkillsTemplatesItemsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-templates", skillTemplateId, "items", skillTemplateItemId},
	}
	var result items.SkillsTemplatesItemsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
