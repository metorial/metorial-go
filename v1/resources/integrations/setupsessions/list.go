package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation represents the integrations setup sessions list output items integration instance implementation type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider represents the integrations setup sessions list output items integration instance providers provider type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion represents the integrations setup sessions list output items integration instance providers integration provider provider version type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions list output items integration instance providers integration provider config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions list output items integration instance providers integration provider type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider struct {
	Object          string                                                                                                 `json:"object"`
	Id              string                                                                                                 `json:"id"`
	ProviderVersion IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                                                 `json:"status"`
	Name            string                                                                                                 `json:"name"`
	Description     *string                                                                                                `json:"description,omitempty"`
	Metadata        *map[string]any                                                                                        `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                                             `json:"provider_id"`
	DeploymentId      string                                                                                             `json:"deployment_id"`
	AuthMethodId      *string                                                                                            `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                                            `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                                          `json:"created_at"`
	UpdatedAt         time.Time                                                                                          `json:"updated_at"`
	ArchivedAt        *time.Time                                                                                         `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig represents the integrations setup sessions list output items integration instance providers config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions list output items integration instance providers auth config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders represents the integrations setup sessions list output items integration instance providers type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                                                    `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                                               `json:"created_at"`
	UpdatedAt            time.Time                                                                               `json:"updated_at"`
	ArchivedAt           *time.Time                                                                              `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstance represents the integrations setup sessions list output items integration instance type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstance struct {
	Object          string                                                                     `json:"object"`
	Id              string                                                                     `json:"id"`
	Status          string                                                                     `json:"status"`
	Name            string                                                                     `json:"name"`
	Description     *string                                                                    `json:"description,omitempty"`
	Metadata        *map[string]any                                                            `json:"metadata,omitempty"`
	IntegrationId   string                                                                     `json:"integration_id"`
	IdentityActorId *string                                                                    `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                                                    `json:"identity_id,omitempty"`
	Implementation  *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders     `json:"providers"`
	CreatedAt       time.Time                                                                  `json:"created_at"`
	UpdatedAt       time.Time                                                                  `json:"updated_at"`
	ArchivedAt      *time.Time                                                                 `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItems represents the integrations setup sessions list output items type.
type IntegrationsSetupSessionsListOutputItems struct {
	Object              string                                                      `json:"object"`
	Id                  string                                                      `json:"id"`
	Status              string                                                      `json:"status"`
	Url                 string                                                      `json:"url"`
	Name                *string                                                     `json:"name,omitempty"`
	Description         *string                                                     `json:"description,omitempty"`
	Metadata            *map[string]any                                             `json:"metadata,omitempty"`
	Configuration       *map[string]any                                             `json:"configuration,omitempty"`
	RedirectUrl         *string                                                     `json:"redirect_url,omitempty"`
	IntegrationId       string                                                      `json:"integration_id"`
	IntegrationInstance IntegrationsSetupSessionsListOutputItemsIntegrationInstance `json:"integration_instance"`
	CreatedAt           time.Time                                                   `json:"created_at"`
	UpdatedAt           time.Time                                                   `json:"updated_at"`
	ExpiresAt           time.Time                                                   `json:"expires_at"`
}

// IntegrationsSetupSessionsListOutputPagination represents the integrations setup sessions list output pagination type.
type IntegrationsSetupSessionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsSetupSessionsListOutput represents the integrations setup sessions list output type.
type IntegrationsSetupSessionsListOutput struct {
	Items      []IntegrationsSetupSessionsListOutputItems    `json:"items"`
	Pagination IntegrationsSetupSessionsListOutputPagination `json:"pagination"`
}

// MapIntegrationsSetupSessionsListOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsListOutput.
func MapIntegrationsSetupSessionsListOutputFromJSON(data []byte) (*IntegrationsSetupSessionsListOutput, error) {
	var v IntegrationsSetupSessionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsListOutputToJSON serializes a IntegrationsSetupSessionsListOutput to JSON.
func MapIntegrationsSetupSessionsListOutputToJSON(v *IntegrationsSetupSessionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsSetupSessionsListQueryCreatedAt - Filter integration setup session creation time by date range
type IntegrationsSetupSessionsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration setup session creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration setup session creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsSetupSessionsListQueryUpdatedAt - Filter integration setup session last update time by date range
type IntegrationsSetupSessionsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration setup session last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration setup session last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsSetupSessionsListQuery represents the integrations setup sessions list query type.
type IntegrationsSetupSessionsListQuery struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId *any     `json:"integration_instance_id,omitempty"`
	// CreatedAt - Filter integration setup session creation time by date range
	CreatedAt *IntegrationsSetupSessionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration setup session last update time by date range
	UpdatedAt *IntegrationsSetupSessionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsSetupSessionsListQueryFromJSON deserializes JSON data into a IntegrationsSetupSessionsListQuery.
func MapIntegrationsSetupSessionsListQueryFromJSON(data []byte) (*IntegrationsSetupSessionsListQuery, error) {
	var v IntegrationsSetupSessionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsListQueryToJSON serializes a IntegrationsSetupSessionsListQuery to JSON.
func MapIntegrationsSetupSessionsListQueryToJSON(v *IntegrationsSetupSessionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
