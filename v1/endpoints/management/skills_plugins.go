package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/plugins"
)

// SkillsPluginsEndpoint provides access to manage skill plugins for an instance.
type SkillsPluginsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsPluginsEndpoint creates a new SkillsPluginsEndpoint.
func NewSkillsPluginsEndpoint(client *endpoint.Client) *SkillsPluginsEndpoint {
	return &SkillsPluginsEndpoint{client: client}
}

// SkillsPluginsEndpointListParams contains optional query parameters for List.
type SkillsPluginsEndpointListParams struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillMarketplaceId   *any     `json:"skill_marketplace_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	Category             *string  `json:"category,omitempty"`
	Search               *string  `json:"search,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill plugin creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill plugin last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// SkillsPluginsEndpointCreateBody contains the request body for Create.
type SkillsPluginsEndpointCreateBody struct {
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
}

// SkillsPluginsEndpointUpdateBody contains the request body for Update.
type SkillsPluginsEndpointUpdateBody struct {
	Name                 *string `json:"name,omitempty"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	ImageFileId          *string `json:"image_file_id,omitempty"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
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

// Update updates a skill plugin.
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

// Sync forces a skill plugin sync.
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
