package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/configurations"
)

// SkillsConfigurationsEndpoint provides access to manage configuration profiles for skill execution.
type SkillsConfigurationsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsConfigurationsEndpoint creates a new SkillsConfigurationsEndpoint.
func NewSkillsConfigurationsEndpoint(client *endpoint.Client) *SkillsConfigurationsEndpoint {
	return &SkillsConfigurationsEndpoint{client: client}
}

// SkillsConfigurationsEndpointCreateBody contains the request body for Create.
type SkillsConfigurationsEndpointCreateBody struct {
	AllowScripts                *bool     `json:"allow_scripts,omitempty"`
	AllowedFileExtensions       *[]string `json:"allowed_file_extensions,omitempty"`
	AllowNonStandardDirectories *bool     `json:"allow_non_standard_directories,omitempty"`
}

// SkillsConfigurationsEndpointListParams contains optional query parameters for List.
type SkillsConfigurationsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// SkillsConfigurationsEndpointUpdateBody contains the request body for Update.
type SkillsConfigurationsEndpointUpdateBody struct {
	AllowScripts                *bool     `json:"allow_scripts,omitempty"`
	AllowedFileExtensions       *[]string `json:"allowed_file_extensions,omitempty"`
	AllowNonStandardDirectories *bool     `json:"allow_non_standard_directories,omitempty"`
}

// Create creates a new non-default skill configuration.
func (e *SkillsConfigurationsEndpoint) Create(instanceId string, body *SkillsConfigurationsEndpointCreateBody) (*configurations.SkillsConfigurationsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", "configurations"},
		Body: body,
	}
	var result configurations.SkillsConfigurationsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a paginated list of visible skill configurations.
func (e *SkillsConfigurationsEndpoint) List(instanceId string, params *SkillsConfigurationsEndpointListParams) (*configurations.SkillsConfigurationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "skills", "configurations"},
		Query: query,
	}
	var result configurations.SkillsConfigurationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific skill configuration by ID, or the default.
func (e *SkillsConfigurationsEndpoint) Get(instanceId string, skillConfigurationId string) (*configurations.SkillsConfigurationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", "configurations", skillConfigurationId},
	}
	var result configurations.SkillsConfigurationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific skill configuration. Updating default creates it first if needed.
func (e *SkillsConfigurationsEndpoint) Update(instanceId string, skillConfigurationId string, body *SkillsConfigurationsEndpointUpdateBody) (*configurations.SkillsConfigurationsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", "configurations", skillConfigurationId},
		Body: body,
	}
	var result configurations.SkillsConfigurationsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete soft deletes a specific non-internal skill configuration.
func (e *SkillsConfigurationsEndpoint) Delete(instanceId string, skillConfigurationId string) (*configurations.SkillsConfigurationsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "skills", "configurations", skillConfigurationId},
	}
	var result configurations.SkillsConfigurationsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
