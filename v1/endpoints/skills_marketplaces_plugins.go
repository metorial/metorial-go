package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/marketplaces/plugins"
)

// SkillsMarketplacesPluginsEndpoint provides access to a skill marketplace plugin is one plugin in a skill marketplace.
type SkillsMarketplacesPluginsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMarketplacesPluginsEndpoint creates a new SkillsMarketplacesPluginsEndpoint.
func NewSkillsMarketplacesPluginsEndpoint(client *endpoint.Client) *SkillsMarketplacesPluginsEndpoint {
	return &SkillsMarketplacesPluginsEndpoint{client: client}
}

// SkillsMarketplacesPluginsEndpointListParams contains optional query parameters for List.
type SkillsMarketplacesPluginsEndpointListParams struct {
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
	// Id - Only return marketplace plugins with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillPluginId - Only return links for one of these skill plugins.
	SkillPluginId *any `json:"skill_plugin_id,omitempty"`
	// Status - Only return marketplace plugins with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return marketplace plugins using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Only return results whose skill marketplace plugin creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill marketplace plugin last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsMarketplacesPluginsEndpointAddBody contains the request body for Add.
type SkillsMarketplacesPluginsEndpointAddBody struct {
	// SkillPluginId - Skill plugin to add to the marketplace, by ID.
	SkillPluginId string `json:"skill_plugin_id"`
	// SkillConfigurationId - Skill configuration for the plugin link, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
	// Identifier - Slug identifying the plugin in the marketplace. Defaults to the plugin name.
	Identifier *string `json:"identifier,omitempty"`
}

// List returns plugins linked to a skill marketplace.
func (e *SkillsMarketplacesPluginsEndpoint) List(skillMarketplaceId string, params *SkillsMarketplacesPluginsEndpointListParams) (*plugins.SkillsMarketplacesPluginsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"skill-marketplaces", skillMarketplaceId, "plugins"},
		Query: query,
	}
	var result plugins.SkillsMarketplacesPluginsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Add adds a skill plugin to a skill marketplace. Adding a plugin that was removed reactivates its link.
func (e *SkillsMarketplacesPluginsEndpoint) Add(skillMarketplaceId string, body *SkillsMarketplacesPluginsEndpointAddBody) (*plugins.SkillsMarketplacesPluginsAddOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId, "plugins"},
		Body: body,
	}
	var result plugins.SkillsMarketplacesPluginsAddOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill marketplace plugin link.
func (e *SkillsMarketplacesPluginsEndpoint) Get(skillMarketplaceId string, skillMarketplacePluginId string) (*plugins.SkillsMarketplacesPluginsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId, "plugins", skillMarketplacePluginId},
	}
	var result plugins.SkillsMarketplacesPluginsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Remove removes a skill plugin from a skill marketplace.
func (e *SkillsMarketplacesPluginsEndpoint) Remove(skillMarketplaceId string, skillMarketplacePluginId string) (*plugins.SkillsMarketplacesPluginsRemoveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"skill-marketplaces", skillMarketplaceId, "plugins", skillMarketplacePluginId},
	}
	var result plugins.SkillsMarketplacesPluginsRemoveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
