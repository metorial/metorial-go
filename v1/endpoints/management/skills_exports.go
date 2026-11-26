package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/exports"
)

// SkillsExportsEndpoint provides access to a skill export is a zip file containing a skill, plugin, or marketplace.
type SkillsExportsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsExportsEndpoint creates a new SkillsExportsEndpoint.
func NewSkillsExportsEndpoint(client *endpoint.Client) *SkillsExportsEndpoint {
	return &SkillsExportsEndpoint{client: client}
}

// SkillsExportsEndpointListParams contains optional query parameters for List.
type SkillsExportsEndpointListParams struct {
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
	// Id - Only return skill exports with these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return skill exports of one of these types.
	Target *any `json:"target,omitempty"`
	// Status - Only return skill exports with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// SkillsExportsEndpointCreateBody contains the request body for Create.
type SkillsExportsEndpointCreateBody struct {
	// Target - What to export: a skill, plugin, or marketplace.
	Target string `json:"target"`
	// SkillId - Skill to export, by ID. Required when `target` is `skill`.
	SkillId *string `json:"skill_id,omitempty"`
	// SkillPluginId - Skill plugin to export, by ID. Required when `target` is `plugin`.
	SkillPluginId *string `json:"skill_plugin_id,omitempty"`
	// SkillMarketplaceId - Skill marketplace to export, by ID. Required when `target` is `marketplace`.
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
