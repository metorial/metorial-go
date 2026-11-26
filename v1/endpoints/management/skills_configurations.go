package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/skills/configurations"
)

// SkillsConfigurationsEndpoint provides access to a skill configuration is a profile controlling how skills execute.
type SkillsConfigurationsEndpoint struct {
	client *endpoint.Client
}

// NewSkillsConfigurationsEndpoint creates a new SkillsConfigurationsEndpoint.
func NewSkillsConfigurationsEndpoint(client *endpoint.Client) *SkillsConfigurationsEndpoint {
	return &SkillsConfigurationsEndpoint{client: client}
}

// SkillsConfigurationsEndpointCreateBody contains the request body for Create.
type SkillsConfigurationsEndpointCreateBody struct {
	// AllowScripts - Allow skills to run scripts.
	AllowScripts *bool `json:"allow_scripts,omitempty"`
	// AllowedFileExtensions - File extensions skills may use.
	AllowedFileExtensions *[]string `json:"allowed_file_extensions,omitempty"`
	// AllowNonStandardDirectories - Allow skills to write outside the standard directories.
	AllowNonStandardDirectories *bool `json:"allow_non_standard_directories,omitempty"`
	// SecurityReviewPolicy - How security reviews affect distribution: `warn` flags failing skills, `block` withholds them from syncs and exports. The strictest policy of all applicable configurations wins.
	SecurityReviewPolicy *string `json:"security_review_policy,omitempty"`
}

// SkillsConfigurationsEndpointListParams contains optional query parameters for List.
type SkillsConfigurationsEndpointListParams struct {
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
}

// SkillsConfigurationsEndpointUpdateBody contains the request body for Update.
type SkillsConfigurationsEndpointUpdateBody struct {
	// AllowScripts - Allow skills to run scripts.
	AllowScripts *bool `json:"allow_scripts,omitempty"`
	// AllowedFileExtensions - File extensions skills may use.
	AllowedFileExtensions *[]string `json:"allowed_file_extensions,omitempty"`
	// AllowNonStandardDirectories - Allow skills to write outside the standard directories.
	AllowNonStandardDirectories *bool `json:"allow_non_standard_directories,omitempty"`
	// SecurityReviewPolicy - How security reviews affect distribution: `warn` flags failing skills, `block` withholds them from syncs and exports. The strictest policy of all applicable configurations wins.
	SecurityReviewPolicy *string `json:"security_review_policy,omitempty"`
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

// Get retrieves a skill configuration. Pass `default` for the default configuration.
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

// Update updates a skill configuration. Omitted fields keep their current value. Updating `default` creates it first if needed.
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

// Delete deletes a skill configuration. The default and internal configurations cannot be deleted.
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
