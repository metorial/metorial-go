package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesProvidersSetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersSetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputProvider represents the integrations instances providers set output provider type.
type IntegrationsInstancesProvidersSetOutputProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion represents the integrations instances providers set output integration provider provider version type.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig represents the integrations instances providers set output integration provider config type.
type IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig struct {
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

// IntegrationsInstancesProvidersSetOutputIntegrationProvider represents the integrations instances providers set output integration provider type.
type IntegrationsInstancesProvidersSetOutputIntegrationProvider struct {
	Object          string                                                                    `json:"object"`
	Id              string                                                                    `json:"id"`
	ProviderVersion IntegrationsInstancesProvidersSetOutputIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                    `json:"status"`
	Name            string                                                                    `json:"name"`
	Description     *string                                                                   `json:"description,omitempty"`
	Metadata        *map[string]any                                                           `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesProvidersSetOutputIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                `json:"provider_id"`
	DeploymentId      string                                                                `json:"deployment_id"`
	AuthMethodId      *string                                                               `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                               `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesProvidersSetOutputIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                             `json:"created_at"`
	UpdatedAt         time.Time                                                             `json:"updated_at"`
	ArchivedAt        *time.Time                                                            `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersSetOutputConfig represents the integrations instances providers set output config type.
type IntegrationsInstancesProvidersSetOutputConfig struct {
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

// IntegrationsInstancesProvidersSetOutputAuthConfig represents the integrations instances providers set output auth config type.
type IntegrationsInstancesProvidersSetOutputAuthConfig struct {
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

// IntegrationsInstancesProvidersSetOutput represents the integrations instances providers set output type.
type IntegrationsInstancesProvidersSetOutput struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesProvidersSetOutputToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesProvidersSetOutputProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesProvidersSetOutputIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesProvidersSetOutputConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesProvidersSetOutputAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesProvidersSetOutputFromJSON deserializes JSON data into a IntegrationsInstancesProvidersSetOutput.
func MapIntegrationsInstancesProvidersSetOutputFromJSON(data []byte) (*IntegrationsInstancesProvidersSetOutput, error) {
	var v IntegrationsInstancesProvidersSetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersSetOutputToJSON serializes a IntegrationsInstancesProvidersSetOutput to JSON.
func MapIntegrationsInstancesProvidersSetOutputToJSON(v *IntegrationsInstancesProvidersSetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesProvidersSetBody represents the integrations instances providers set body type.
type IntegrationsInstancesProvidersSetBody struct {
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
	IsOverrideToolFilter *bool   `json:"is_override_tool_filter,omitempty"`
}

// MapIntegrationsInstancesProvidersSetBodyFromJSON deserializes JSON data into a IntegrationsInstancesProvidersSetBody.
func MapIntegrationsInstancesProvidersSetBodyFromJSON(data []byte) (*IntegrationsInstancesProvidersSetBody, error) {
	var v IntegrationsInstancesProvidersSetBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersSetBodyToJSON serializes a IntegrationsInstancesProvidersSetBody to JSON.
func MapIntegrationsInstancesProvidersSetBodyToJSON(v *IntegrationsInstancesProvidersSetBody) ([]byte, error) {
	return json.Marshal(v)
}
