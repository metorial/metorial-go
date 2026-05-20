package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/marketplaces/plugins"
)

// SkillsMarketplacesPluginsEndpoint provides access to manage plugin links for skill marketplaces.
type SkillsMarketplacesPluginsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsMarketplacesPluginsEndpoint creates a new SkillsMarketplacesPluginsEndpoint.
func NewSkillsMarketplacesPluginsEndpoint(client *endpoint.Client) *SkillsMarketplacesPluginsEndpoint {
	return &SkillsMarketplacesPluginsEndpoint{client: client}
}

// SkillsMarketplacesPluginsEndpointListParams contains optional query parameters for List.
type SkillsMarketplacesPluginsEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillPluginId        *any     `json:"skill_plugin_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill marketplace plugin creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill marketplace plugin last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsMarketplacesPluginsEndpointAddBody contains the request body for Add.
type SkillsMarketplacesPluginsEndpointAddBody struct {
	SkillPluginId        string  `json:"skill_plugin_id"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
	Identifier           *string `json:"identifier,omitempty"`
}

// List returns plugins linked to a skill marketplace.
func (e *SkillsMarketplacesPluginsEndpoint) List(instanceId string, skillMarketplaceId string, params *SkillsMarketplacesPluginsEndpointListParams) (*plugins.SkillsMarketplacesPluginsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId, "plugins"},
		Query: query,
	}
	var result plugins.SkillsMarketplacesPluginsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Add adds a skill plugin to a skill marketplace.
func (e *SkillsMarketplacesPluginsEndpoint) Add(instanceId string, skillMarketplaceId string, body *SkillsMarketplacesPluginsEndpointAddBody) (*plugins.SkillsMarketplacesPluginsAddOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId, "plugins"},
		Body: body,
	}
	var result plugins.SkillsMarketplacesPluginsAddOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill marketplace plugin link.
func (e *SkillsMarketplacesPluginsEndpoint) Get(instanceId string, skillMarketplaceId string, skillMarketplacePluginId string) (*plugins.SkillsMarketplacesPluginsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId, "plugins", skillMarketplacePluginId},
	}
	var result plugins.SkillsMarketplacesPluginsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Remove removes a skill plugin from a skill marketplace.
func (e *SkillsMarketplacesPluginsEndpoint) Remove(instanceId string, skillMarketplaceId string, skillMarketplacePluginId string) (*plugins.SkillsMarketplacesPluginsRemoveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-marketplaces", skillMarketplaceId, "plugins", skillMarketplacePluginId},
	}
	var result plugins.SkillsMarketplacesPluginsRemoveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
