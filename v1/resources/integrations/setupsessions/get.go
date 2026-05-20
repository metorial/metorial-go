package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation represents the integrations setup sessions get output integration instance implementation type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider represents the integrations setup sessions get output integration instance providers provider type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion represents the integrations setup sessions get output integration instance providers integration provider provider version type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions get output integration instance providers integration provider config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions get output integration instance providers integration provider type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider struct {
	Object          string                                                                                           `json:"object"`
	Id              string                                                                                           `json:"id"`
	ProviderVersion IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                                           `json:"status"`
	Name            string                                                                                           `json:"name"`
	Description     *string                                                                                          `json:"description,omitempty"`
	Metadata        *map[string]any                                                                                  `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                                       `json:"provider_id"`
	DeploymentId      string                                                                                       `json:"deployment_id"`
	AuthMethodId      *string                                                                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                                    `json:"created_at"`
	UpdatedAt         time.Time                                                                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                                                                   `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig represents the integrations setup sessions get output integration instance providers config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions get output integration instance providers auth config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders represents the integrations setup sessions get output integration instance providers type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                                              `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                                         `json:"created_at"`
	UpdatedAt            time.Time                                                                         `json:"updated_at"`
	ArchivedAt           *time.Time                                                                        `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstance represents the integrations setup sessions get output integration instance type.
type IntegrationsSetupSessionsGetOutputIntegrationInstance struct {
	Object          string                                                               `json:"object"`
	Id              string                                                               `json:"id"`
	Status          string                                                               `json:"status"`
	Name            string                                                               `json:"name"`
	Description     *string                                                              `json:"description,omitempty"`
	Metadata        *map[string]any                                                      `json:"metadata,omitempty"`
	IntegrationId   string                                                               `json:"integration_id"`
	IdentityActorId *string                                                              `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                                              `json:"identity_id,omitempty"`
	Implementation  *IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders     `json:"providers"`
	CreatedAt       time.Time                                                            `json:"created_at"`
	UpdatedAt       time.Time                                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                                           `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutput represents the integrations setup sessions get output type.
type IntegrationsSetupSessionsGetOutput struct {
	Object              string                                                `json:"object"`
	Id                  string                                                `json:"id"`
	Status              string                                                `json:"status"`
	Url                 string                                                `json:"url"`
	Name                *string                                               `json:"name,omitempty"`
	Description         *string                                               `json:"description,omitempty"`
	Metadata            *map[string]any                                       `json:"metadata,omitempty"`
	Configuration       *map[string]any                                       `json:"configuration,omitempty"`
	RedirectUrl         *string                                               `json:"redirect_url,omitempty"`
	IntegrationId       string                                                `json:"integration_id"`
	IntegrationInstance IntegrationsSetupSessionsGetOutputIntegrationInstance `json:"integration_instance"`
	CreatedAt           time.Time                                             `json:"created_at"`
	UpdatedAt           time.Time                                             `json:"updated_at"`
	ExpiresAt           time.Time                                             `json:"expires_at"`
}

// MapIntegrationsSetupSessionsGetOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsGetOutput.
func MapIntegrationsSetupSessionsGetOutputFromJSON(data []byte) (*IntegrationsSetupSessionsGetOutput, error) {
	var v IntegrationsSetupSessionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsGetOutputToJSON serializes a IntegrationsSetupSessionsGetOutput to JSON.
func MapIntegrationsSetupSessionsGetOutputToJSON(v *IntegrationsSetupSessionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
