package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/skills/plugins/skills"
)

// SkillsPluginsSkillsEndpoint provides access to a skill plugin skill is one skill in a skill plugin.
type SkillsPluginsSkillsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsPluginsSkillsEndpoint creates a new SkillsPluginsSkillsEndpoint.
func NewSkillsPluginsSkillsEndpoint(client *endpoint.Client) *SkillsPluginsSkillsEndpoint {
	return &SkillsPluginsSkillsEndpoint{client: client}
}

// SkillsPluginsSkillsEndpointListParams contains optional query parameters for List.
type SkillsPluginsSkillsEndpointListParams struct {
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
	// Id - Only return plugin skills with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return links for one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// Status - Only return plugin skills with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return plugin skills using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Only return results whose skill plugin skill creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill plugin skill last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsPluginsSkillsEndpointAddBody contains the request body for Add.
type SkillsPluginsSkillsEndpointAddBody struct {
	// SkillId - Skill to add to the plugin, by ID.
	SkillId string `json:"skill_id"`
	// Identifier - Slug identifying the skill in the plugin. Defaults to the skill name.
	Identifier *string `json:"identifier,omitempty"`
	// ClientName - Client-facing skill name. `null` clears it.
	ClientName *nullable.Nullable[string] `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description. `null` clears it.
	ClientDescription *nullable.Nullable[string] `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill. `null` clears it.
	ClientMetadata *nullable.Nullable[map[string]any] `json:"client_metadata,omitempty"`
	// License - License terms for the skill. `null` clears it.
	License *nullable.Nullable[string] `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill. `null` clears it.
	Compatibility *nullable.Nullable[string] `json:"compatibility,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin skill, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
}

// SkillsPluginsSkillsEndpointUpdateBody contains the request body for Update.
type SkillsPluginsSkillsEndpointUpdateBody struct {
	// ClientName - Client-facing skill name. `null` clears it.
	ClientName *nullable.Nullable[string] `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description. `null` clears it.
	ClientDescription *nullable.Nullable[string] `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill. `null` clears it.
	ClientMetadata *nullable.Nullable[map[string]any] `json:"client_metadata,omitempty"`
	// License - License terms for the skill. `null` clears it.
	License *nullable.Nullable[string] `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill. `null` clears it.
	Compatibility *nullable.Nullable[string] `json:"compatibility,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin skill, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
}

// List returns skills linked to a skill plugin.
func (e *SkillsPluginsSkillsEndpoint) List(instanceId string, skillPluginId string, params *SkillsPluginsSkillsEndpointListParams) (*skills.SkillsPluginsSkillsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skill-plugins", skillPluginId, "skills"},
		Query: query,
	}
	var result skills.SkillsPluginsSkillsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Add adds a skill to a skill plugin. Adding a skill that was removed reactivates its link.
func (e *SkillsPluginsSkillsEndpoint) Add(instanceId string, skillPluginId string, body *SkillsPluginsSkillsEndpointAddBody) (*skills.SkillsPluginsSkillsAddOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "skills"},
		Body: body,
	}
	var result skills.SkillsPluginsSkillsAddOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a skill plugin skill link.
func (e *SkillsPluginsSkillsEndpoint) Get(instanceId string, skillPluginId string, skillPluginSkillId string) (*skills.SkillsPluginsSkillsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "skills", skillPluginSkillId},
	}
	var result skills.SkillsPluginsSkillsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a skill plugin skill link.
func (e *SkillsPluginsSkillsEndpoint) Update(instanceId string, skillPluginId string, skillPluginSkillId string, body *SkillsPluginsSkillsEndpointUpdateBody) (*skills.SkillsPluginsSkillsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "skills", skillPluginSkillId},
		Body: body,
	}
	var result skills.SkillsPluginsSkillsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Remove removes a skill from a skill plugin.
func (e *SkillsPluginsSkillsEndpoint) Remove(instanceId string, skillPluginId string, skillPluginSkillId string) (*skills.SkillsPluginsSkillsRemoveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skill-plugins", skillPluginId, "skills", skillPluginSkillId},
	}
	var result skills.SkillsPluginsSkillsRemoveOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
