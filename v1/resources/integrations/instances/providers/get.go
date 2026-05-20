package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesProvidersGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersGetOutputProvider represents the integrations instances providers get output provider type.
type IntegrationsInstancesProvidersGetOutputProvider struct {
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

// IntegrationsInstancesProvidersGetOutputIntegrationProviderProviderVersion represents the integrations instances providers get output integration provider provider version type.
type IntegrationsInstancesProvidersGetOutputIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesProvidersGetOutputIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersGetOutputIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersGetOutputIntegrationProviderConfig represents the integrations instances providers get output integration provider config type.
type IntegrationsInstancesProvidersGetOutputIntegrationProviderConfig struct {
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

// IntegrationsInstancesProvidersGetOutputIntegrationProvider represents the integrations instances providers get output integration provider type.
type IntegrationsInstancesProvidersGetOutputIntegrationProvider struct {
	Object          string                                                                    `json:"object"`
	Id              string                                                                    `json:"id"`
	ProviderVersion IntegrationsInstancesProvidersGetOutputIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                    `json:"status"`
	Name            string                                                                    `json:"name"`
	Description     *string                                                                   `json:"description,omitempty"`
	Metadata        *map[string]any                                                           `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesProvidersGetOutputIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                `json:"provider_id"`
	DeploymentId      string                                                                `json:"deployment_id"`
	AuthMethodId      *string                                                               `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                               `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesProvidersGetOutputIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                             `json:"created_at"`
	UpdatedAt         time.Time                                                             `json:"updated_at"`
	ArchivedAt        *time.Time                                                            `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersGetOutputConfig represents the integrations instances providers get output config type.
type IntegrationsInstancesProvidersGetOutputConfig struct {
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

// IntegrationsInstancesProvidersGetOutputAuthConfig represents the integrations instances providers get output auth config type.
type IntegrationsInstancesProvidersGetOutputAuthConfig struct {
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

// IntegrationsInstancesProvidersGetOutput represents the integrations instances providers get output type.
type IntegrationsInstancesProvidersGetOutput struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesProvidersGetOutputToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesProvidersGetOutputProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesProvidersGetOutputIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesProvidersGetOutputConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesProvidersGetOutputAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesProvidersGetOutputFromJSON deserializes JSON data into a IntegrationsInstancesProvidersGetOutput.
func MapIntegrationsInstancesProvidersGetOutputFromJSON(data []byte) (*IntegrationsInstancesProvidersGetOutput, error) {
	var v IntegrationsInstancesProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersGetOutputToJSON serializes a IntegrationsInstancesProvidersGetOutput to JSON.
func MapIntegrationsInstancesProvidersGetOutputToJSON(v *IntegrationsInstancesProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
