package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment represents the integrations instance groups create session template output providers deployment type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig represents the integrations instance groups create session template output providers config type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig represents the integrations instance groups create session template output providers auth config type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders represents the integrations instance groups create session template output providers type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session template provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Parent session template ID
	SessionTemplateId string                                                                    `json:"session_template_id"`
	Deployment        IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment  `json:"deployment"`
	Config            IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig      `json:"config"`
	AuthConfig        *IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutput represents the integrations instance groups create session template output type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session template identifier
	Id string `json:"id"`
	// Status - Status of the session template
	Status string `json:"status"`
	// Name - Template name
	Name string `json:"name"`
	// Description - Template description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs
	Metadata                   *map[string]any `json:"metadata,omitempty"`
	IntegrationInstanceId      *string         `json:"integration_instance_id,omitempty"`
	IntegrationInstanceGroupId *string         `json:"integration_instance_group_id,omitempty"`
	// Providers - Template providers
	Providers []IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionTemplateOutput.
func MapIntegrationsInstanceGroupsCreateSessionTemplateOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionTemplateOutput, error) {
	var v IntegrationsInstanceGroupsCreateSessionTemplateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateOutputToJSON serializes a IntegrationsInstanceGroupsCreateSessionTemplateOutput to JSON.
func MapIntegrationsInstanceGroupsCreateSessionTemplateOutputToJSON(v *IntegrationsInstanceGroupsCreateSessionTemplateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsCreateSessionTemplateBody represents the integrations instance groups create session template body type.
type IntegrationsInstanceGroupsCreateSessionTemplateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionTemplateBody.
func MapIntegrationsInstanceGroupsCreateSessionTemplateBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionTemplateBody, error) {
	var v IntegrationsInstanceGroupsCreateSessionTemplateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateBodyToJSON serializes a IntegrationsInstanceGroupsCreateSessionTemplateBody to JSON.
func MapIntegrationsInstanceGroupsCreateSessionTemplateBodyToJSON(v *IntegrationsInstanceGroupsCreateSessionTemplateBody) ([]byte, error) {
	return json.Marshal(v)
}
