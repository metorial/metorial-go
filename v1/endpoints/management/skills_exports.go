package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/exports"
)

// SkillsExportsEndpoint provides access to export skills, skill plugins, and skill marketplaces as zip files.
type SkillsExportsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsExportsEndpoint creates a new SkillsExportsEndpoint.
func NewSkillsExportsEndpoint(client *endpoint.Client) *SkillsExportsEndpoint {
	return &SkillsExportsEndpoint{client: client}
}

// SkillsExportsEndpointListParams contains optional query parameters for List.
type SkillsExportsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Id     *any     `json:"id,omitempty"`
	Target *any     `json:"target,omitempty"`
	Status *any     `json:"status,omitempty"`
}

// SkillsExportsEndpointCreateBody contains the request body for Create.
type SkillsExportsEndpointCreateBody struct {
	Target             string  `json:"target"`
	SkillId            *string `json:"skill_id,omitempty"`
	SkillPluginId      *string `json:"skill_plugin_id,omitempty"`
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
}

// List returns a paginated list of skill exports.
func (e *SkillsExportsEndpoint) List(instanceId string, params *SkillsExportsEndpointListParams) (*exports.SkillsExportsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-exports"},
		Query: query,
	}
	var result exports.SkillsExportsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill export.
func (e *SkillsExportsEndpoint) Get(instanceId string, skillExportId string) (*exports.SkillsExportsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-exports", skillExportId},
	}
	var result exports.SkillsExportsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill, plugin, or marketplace export.
func (e *SkillsExportsEndpoint) Create(instanceId string, body *SkillsExportsEndpointCreateBody) (*exports.SkillsExportsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-exports"},
		Body: body,
	}
	var result exports.SkillsExportsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
