package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesCreateOutputImplementation represents the integrations instances create output implementation type.
type IntegrationsInstancesCreateOutputImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesCreateOutputProvidersProvider represents the integrations instances create output providers provider type.
type IntegrationsInstancesCreateOutputProvidersProvider struct {
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

// IntegrationsInstancesCreateOutputProvidersIntegrationProviderProviderVersion represents the integrations instances create output providers integration provider provider version type.
type IntegrationsInstancesCreateOutputProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesCreateOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesCreateOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesCreateOutputProvidersIntegrationProviderConfig represents the integrations instances create output providers integration provider config type.
type IntegrationsInstancesCreateOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesCreateOutputProvidersIntegrationProvider represents the integrations instances create output providers integration provider type.
type IntegrationsInstancesCreateOutputProvidersIntegrationProvider struct {
	Object          string                                                                       `json:"object"`
	Id              string                                                                       `json:"id"`
	ProviderVersion IntegrationsInstancesCreateOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                       `json:"status"`
	Name            string                                                                       `json:"name"`
	Description     *string                                                                      `json:"description,omitempty"`
	Metadata        *map[string]any                                                              `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesCreateOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                   `json:"provider_id"`
	DeploymentId      string                                                                   `json:"deployment_id"`
	AuthMethodId      *string                                                                  `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesCreateOutputProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                `json:"created_at"`
	UpdatedAt         time.Time                                                                `json:"updated_at"`
	ArchivedAt        *time.Time                                                               `json:"archived_at,omitempty"`
}

// IntegrationsInstancesCreateOutputProvidersConfig represents the integrations instances create output providers config type.
type IntegrationsInstancesCreateOutputProvidersConfig struct {
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

// IntegrationsInstancesCreateOutputProvidersAuthConfig represents the integrations instances create output providers auth config type.
type IntegrationsInstancesCreateOutputProvidersAuthConfig struct {
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

// IntegrationsInstancesCreateOutputProviders represents the integrations instances create output providers type.
type IntegrationsInstancesCreateOutputProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesCreateOutputProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesCreateOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesCreateOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesCreateOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesCreateOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                     `json:"created_at"`
	UpdatedAt            time.Time                                                     `json:"updated_at"`
	ArchivedAt           *time.Time                                                    `json:"archived_at,omitempty"`
}

// IntegrationsInstancesCreateOutput represents the integrations instances create output type.
type IntegrationsInstancesCreateOutput struct {
	Object          string                                           `json:"object"`
	Id              string                                           `json:"id"`
	Status          string                                           `json:"status"`
	Name            string                                           `json:"name"`
	Description     *string                                          `json:"description,omitempty"`
	Metadata        *map[string]any                                  `json:"metadata,omitempty"`
	IntegrationId   string                                           `json:"integration_id"`
	IdentityActorId *string                                          `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                          `json:"identity_id,omitempty"`
	Implementation  *IntegrationsInstancesCreateOutputImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsInstancesCreateOutputProviders     `json:"providers"`
	CreatedAt       time.Time                                        `json:"created_at"`
	UpdatedAt       time.Time                                        `json:"updated_at"`
	ArchivedAt      *time.Time                                       `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesCreateOutputFromJSON deserializes JSON data into a IntegrationsInstancesCreateOutput.
func MapIntegrationsInstancesCreateOutputFromJSON(data []byte) (*IntegrationsInstancesCreateOutput, error) {
	var v IntegrationsInstancesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesCreateOutputToJSON serializes a IntegrationsInstancesCreateOutput to JSON.
func MapIntegrationsInstancesCreateOutputToJSON(v *IntegrationsInstancesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesCreateBodyProviders represents the integrations instances create body providers type.
type IntegrationsInstancesCreateBodyProviders struct {
	ProviderId           string  `json:"provider_id"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
	IsOverrideToolFilter *bool   `json:"is_override_tool_filter,omitempty"`
}

// IntegrationsInstancesCreateBody represents the integrations instances create body type.
type IntegrationsInstancesCreateBody struct {
	IntegrationId   string                                      `json:"integration_id"`
	Name            string                                      `json:"name"`
	Description     *string                                     `json:"description,omitempty"`
	Metadata        *map[string]any                             `json:"metadata,omitempty"`
	IdentityActorId *string                                     `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                     `json:"identity_id,omitempty"`
	Providers       *[]IntegrationsInstancesCreateBodyProviders `json:"providers,omitempty"`
}

// MapIntegrationsInstancesCreateBodyFromJSON deserializes JSON data into a IntegrationsInstancesCreateBody.
func MapIntegrationsInstancesCreateBodyFromJSON(data []byte) (*IntegrationsInstancesCreateBody, error) {
	var v IntegrationsInstancesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesCreateBodyToJSON serializes a IntegrationsInstancesCreateBody to JSON.
func MapIntegrationsInstancesCreateBodyToJSON(v *IntegrationsInstancesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
