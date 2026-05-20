package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/plugins/skills"
)

// SkillsPluginsSkillsEndpoint provides access to manage skill links for skill plugins.
type SkillsPluginsSkillsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsPluginsSkillsEndpoint creates a new SkillsPluginsSkillsEndpoint.
func NewSkillsPluginsSkillsEndpoint(client *endpoint.Client) *SkillsPluginsSkillsEndpoint {
	return &SkillsPluginsSkillsEndpoint{client: client}
}

// SkillsPluginsSkillsEndpointListParams contains optional query parameters for List.
type SkillsPluginsSkillsEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillId              *any     `json:"skill_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill plugin skill creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill plugin skill last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsPluginsSkillsEndpointAddBody contains the request body for Add.
type SkillsPluginsSkillsEndpointAddBody struct {
	SkillId              string          `json:"skill_id"`
	Identifier           *string         `json:"identifier,omitempty"`
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
}

// SkillsPluginsSkillsEndpointUpdateBody contains the request body for Update.
type SkillsPluginsSkillsEndpointUpdateBody struct {
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
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

// Add adds a skill to a skill plugin.
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
