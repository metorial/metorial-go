package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesDeleteOutputImplementation represents the integrations instances delete output implementation type.
type IntegrationsInstancesDeleteOutputImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesDeleteOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesDeleteOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersProvider represents the integrations instances delete output providers provider type.
type IntegrationsInstancesDeleteOutputProvidersProvider struct {
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

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion represents the integrations instances delete output providers integration provider provider version type.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig represents the integrations instances delete output providers integration provider config type.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesDeleteOutputProvidersIntegrationProvider represents the integrations instances delete output providers integration provider type.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProvider struct {
	Object          string                                                                       `json:"object"`
	Id              string                                                                       `json:"id"`
	ProviderVersion IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                       `json:"status"`
	Name            string                                                                       `json:"name"`
	Description     *string                                                                      `json:"description,omitempty"`
	Metadata        *map[string]any                                                              `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                   `json:"provider_id"`
	DeploymentId      string                                                                   `json:"deployment_id"`
	AuthMethodId      *string                                                                  `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                `json:"created_at"`
	UpdatedAt         time.Time                                                                `json:"updated_at"`
	ArchivedAt        *time.Time                                                               `json:"archived_at,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersConfig represents the integrations instances delete output providers config type.
type IntegrationsInstancesDeleteOutputProvidersConfig struct {
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

// IntegrationsInstancesDeleteOutputProvidersAuthConfig represents the integrations instances delete output providers auth config type.
type IntegrationsInstancesDeleteOutputProvidersAuthConfig struct {
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

// IntegrationsInstancesDeleteOutputProviders represents the integrations instances delete output providers type.
type IntegrationsInstancesDeleteOutputProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesDeleteOutputProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesDeleteOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesDeleteOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesDeleteOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesDeleteOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                     `json:"created_at"`
	UpdatedAt            time.Time                                                     `json:"updated_at"`
	ArchivedAt           *time.Time                                                    `json:"archived_at,omitempty"`
}

// IntegrationsInstancesDeleteOutput represents the integrations instances delete output type.
type IntegrationsInstancesDeleteOutput struct {
	Object          string                                           `json:"object"`
	Id              string                                           `json:"id"`
	Status          string                                           `json:"status"`
	Name            string                                           `json:"name"`
	Description     *string                                          `json:"description,omitempty"`
	Metadata        *map[string]any                                  `json:"metadata,omitempty"`
	IntegrationId   string                                           `json:"integration_id"`
	IdentityActorId *string                                          `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                          `json:"identity_id,omitempty"`
	Implementation  *IntegrationsInstancesDeleteOutputImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsInstancesDeleteOutputProviders     `json:"providers"`
	CreatedAt       time.Time                                        `json:"created_at"`
	UpdatedAt       time.Time                                        `json:"updated_at"`
	ArchivedAt      *time.Time                                       `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesDeleteOutputFromJSON deserializes JSON data into a IntegrationsInstancesDeleteOutput.
func MapIntegrationsInstancesDeleteOutputFromJSON(data []byte) (*IntegrationsInstancesDeleteOutput, error) {
	var v IntegrationsInstancesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesDeleteOutputToJSON serializes a IntegrationsInstancesDeleteOutput to JSON.
func MapIntegrationsInstancesDeleteOutputToJSON(v *IntegrationsInstancesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
