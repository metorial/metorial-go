package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/plugins"
)

// SkillsPluginsEndpoint provides access to a skill plugin is a set of skills distributed through a skill marketplace.
type SkillsPluginsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsPluginsEndpoint creates a new SkillsPluginsEndpoint.
func NewSkillsPluginsEndpoint(client *endpoint.Client) *SkillsPluginsEndpoint {
	return &SkillsPluginsEndpoint{client: client}
}

// SkillsPluginsEndpointListParams contains optional query parameters for List.
type SkillsPluginsEndpointListParams struct {
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
	// Id - Only return skill plugins with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillMarketplaceId - Only return skill plugins in one of these marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// Status - Only return skill plugins with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Category - Only return skill plugins in this category.
	Category *string `json:"category,omitempty"`
	// Search - Only return skill plugins matching this search text.
	Search *string `json:"search,omitempty"`
	// SkillConfigurationId - Only return skill plugins using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// SkillGroupId - Only return plugins whose skills are synced from these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// CreatedAt - Only return results whose skill plugin creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill plugin last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsPluginsEndpointCreateBody contains the request body for Create.
type SkillsPluginsEndpointCreateBody struct {
	// Name - The plugin name. Required unless skill_group_id is set, in which case it defaults to the group name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// LongDescription - Longer description shown on the marketplace. `null` clears it.
	LongDescription *nullable.Nullable[string] `json:"long_description,omitempty"`
	// Category - Category for grouping plugins on the marketplace. `null` clears it.
	Category *nullable.Nullable[string] `json:"category,omitempty"`
	// ImageFileId - Image file for the plugin, by ID. `null` resets to the default image.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
	// SkillMarketplaceId - Marketplace to publish the plugin to, by ID. Required for consumers.
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
	// SkillGroupId - Creates the plugin from a skill group. The group's skills are synced into the plugin automatically and cannot be edited on the plugin directly. Name and description are copied from the group once and can be changed afterwards.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
}

// SkillsPluginsEndpointUpdateBody contains the request body for Update.
type SkillsPluginsEndpointUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// LongDescription - Longer description shown on the marketplace. `null` clears it.
	LongDescription *nullable.Nullable[string] `json:"long_description,omitempty"`
	// Category - Category for grouping plugins on the marketplace. `null` clears it.
	Category *nullable.Nullable[string] `json:"category,omitempty"`
	// ImageFileId - Image file for the plugin, by ID. `null` resets to the default image.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
}

// SkillsPluginsEndpointAttachSkillGroupBody contains the request body for AttachSkillGroup.
type SkillsPluginsEndpointAttachSkillGroupBody struct {
	// SkillGroupId - Skill group to sync the plugin from.
	SkillGroupId string `json:"skill_group_id"`
}

// List returns a paginated list of skill plugins.
func (e *SkillsPluginsEndpoint) List(instanceId string, params *SkillsPluginsEndpointListParams) (*plugins.SkillsPluginsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-plugins"},
		Query: query,
	}
	var result plugins.SkillsPluginsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill plugin.
func (e *SkillsPluginsEndpoint) Get(instanceId string, skillPluginId string) (*plugins.SkillsPluginsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId},
	}
	var result plugins.SkillsPluginsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a skill plugin.
func (e *SkillsPluginsEndpoint) Create(instanceId string, body *SkillsPluginsEndpointCreateBody) (*plugins.SkillsPluginsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins"},
		Body: body,
	}
	var result plugins.SkillsPluginsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill plugin. Omitted fields keep their current value.
func (e *SkillsPluginsEndpoint) Update(instanceId string, skillPluginId string, body *SkillsPluginsEndpointUpdateBody) (*plugins.SkillsPluginsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId},
		Body: body,
	}
	var result plugins.SkillsPluginsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a skill plugin.
func (e *SkillsPluginsEndpoint) Archive(instanceId string, skillPluginId string) (*plugins.SkillsPluginsArchiveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId},
	}
	var result plugins.SkillsPluginsArchiveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AttachSkillGroup starts syncing a plugin from a skill group. The plugin's skills are replaced with the group's skills right away and follow the group from then on.
func (e *SkillsPluginsEndpoint) AttachSkillGroup(instanceId string, skillPluginId string, body *SkillsPluginsEndpointAttachSkillGroupBody) (*plugins.SkillsPluginsAttachSkillGroupOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "attach-skill-group"},
		Body: body,
	}
	var result plugins.SkillsPluginsAttachSkillGroupOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DetachSkillGroup stops syncing a plugin from its skill group. The plugin keeps its current skills, which can then be edited directly.
func (e *SkillsPluginsEndpoint) DetachSkillGroup(instanceId string, skillPluginId string) (*plugins.SkillsPluginsDetachSkillGroupOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "detach-skill-group"},
	}
	var result plugins.SkillsPluginsDetachSkillGroupOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Sync starts a sync of the skill plugin to its linked repositories, even if nothing changed.
func (e *SkillsPluginsEndpoint) Sync(instanceId string, skillPluginId string) (*plugins.SkillsPluginsSyncOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "sync"},
	}
	var result plugins.SkillsPluginsSyncOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
