package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/items"
)

// SkillsItemsEndpoint provides access to skill items attach integrations and providers to skills.
type SkillsItemsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsItemsEndpoint creates a new SkillsItemsEndpoint.
func NewSkillsItemsEndpoint(client *endpoint.Client) *SkillsItemsEndpoint {
	return &SkillsItemsEndpoint{client: client}
}

// SkillsItemsEndpointListParams contains optional query parameters for List.
type SkillsItemsEndpointListParams struct {
	Limit         *float64 `json:"limit,omitempty"`
	After         *string  `json:"after,omitempty"`
	Before        *string  `json:"before,omitempty"`
	Cursor        *string  `json:"cursor,omitempty"`
	Order         *string  `json:"order,omitempty"`
	Status        *any     `json:"status,omitempty"`
	Type          *any     `json:"type,omitempty"`
	Id            *any     `json:"id,omitempty"`
	IntegrationId *any     `json:"integration_id,omitempty"`
	ProviderId    *any     `json:"provider_id,omitempty"`
	// CreatedAt - Filter skill item creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of items for a skill.
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

// Get retrieves a specific skill item.
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
