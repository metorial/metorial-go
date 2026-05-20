package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesUpdateOutputImplementation represents the integrations instances update output implementation type.
type IntegrationsInstancesUpdateOutputImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersProvider represents the integrations instances update output providers provider type.
type IntegrationsInstancesUpdateOutputProvidersProvider struct {
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

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion represents the integrations instances update output providers integration provider provider version type.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig represents the integrations instances update output providers integration provider config type.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesUpdateOutputProvidersIntegrationProvider represents the integrations instances update output providers integration provider type.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProvider struct {
	Object          string                                                                       `json:"object"`
	Id              string                                                                       `json:"id"`
	ProviderVersion IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                       `json:"status"`
	Name            string                                                                       `json:"name"`
	Description     *string                                                                      `json:"description,omitempty"`
	Metadata        *map[string]any                                                              `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                   `json:"provider_id"`
	DeploymentId      string                                                                   `json:"deployment_id"`
	AuthMethodId      *string                                                                  `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                `json:"created_at"`
	UpdatedAt         time.Time                                                                `json:"updated_at"`
	ArchivedAt        *time.Time                                                               `json:"archived_at,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersConfig represents the integrations instances update output providers config type.
type IntegrationsInstancesUpdateOutputProvidersConfig struct {
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

// IntegrationsInstancesUpdateOutputProvidersAuthConfig represents the integrations instances update output providers auth config type.
type IntegrationsInstancesUpdateOutputProvidersAuthConfig struct {
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

// IntegrationsInstancesUpdateOutputProviders represents the integrations instances update output providers type.
type IntegrationsInstancesUpdateOutputProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesUpdateOutputProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesUpdateOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesUpdateOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesUpdateOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesUpdateOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                     `json:"created_at"`
	UpdatedAt            time.Time                                                     `json:"updated_at"`
	ArchivedAt           *time.Time                                                    `json:"archived_at,omitempty"`
}

// IntegrationsInstancesUpdateOutput represents the integrations instances update output type.
type IntegrationsInstancesUpdateOutput struct {
	Object          string                                           `json:"object"`
	Id              string                                           `json:"id"`
	Status          string                                           `json:"status"`
	Name            string                                           `json:"name"`
	Description     *string                                          `json:"description,omitempty"`
	Metadata        *map[string]any                                  `json:"metadata,omitempty"`
	IntegrationId   string                                           `json:"integration_id"`
	IdentityActorId *string                                          `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                          `json:"identity_id,omitempty"`
	Implementation  *IntegrationsInstancesUpdateOutputImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsInstancesUpdateOutputProviders     `json:"providers"`
	CreatedAt       time.Time                                        `json:"created_at"`
	UpdatedAt       time.Time                                        `json:"updated_at"`
	ArchivedAt      *time.Time                                       `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesUpdateOutputFromJSON deserializes JSON data into a IntegrationsInstancesUpdateOutput.
func MapIntegrationsInstancesUpdateOutputFromJSON(data []byte) (*IntegrationsInstancesUpdateOutput, error) {
	var v IntegrationsInstancesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesUpdateOutputToJSON serializes a IntegrationsInstancesUpdateOutput to JSON.
func MapIntegrationsInstancesUpdateOutputToJSON(v *IntegrationsInstancesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesUpdateBodyProviders represents the integrations instances update body providers type.
type IntegrationsInstancesUpdateBodyProviders struct {
	ProviderId           string  `json:"provider_id"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
	IsOverrideToolFilter *bool   `json:"is_override_tool_filter,omitempty"`
}

// IntegrationsInstancesUpdateBody represents the integrations instances update body type.
type IntegrationsInstancesUpdateBody struct {
	Name            *string                                     `json:"name,omitempty"`
	Description     *string                                     `json:"description,omitempty"`
	Metadata        *map[string]any                             `json:"metadata,omitempty"`
	IdentityActorId *string                                     `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                     `json:"identity_id,omitempty"`
	Providers       *[]IntegrationsInstancesUpdateBodyProviders `json:"providers,omitempty"`
}

// MapIntegrationsInstancesUpdateBodyFromJSON deserializes JSON data into a IntegrationsInstancesUpdateBody.
func MapIntegrationsInstancesUpdateBodyFromJSON(data []byte) (*IntegrationsInstancesUpdateBody, error) {
	var v IntegrationsInstancesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesUpdateBodyToJSON serializes a IntegrationsInstancesUpdateBody to JSON.
func MapIntegrationsInstancesUpdateBodyToJSON(v *IntegrationsInstancesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
