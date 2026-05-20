package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesListOutputItemsImplementation represents the integrations instances list output items implementation type.
type IntegrationsInstancesListOutputItemsImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersProvider represents the integrations instances list output items providers provider type.
type IntegrationsInstancesListOutputItemsProvidersProvider struct {
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

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion represents the integrations instances list output items providers integration provider provider version type.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig represents the integrations instances list output items providers integration provider config type.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesListOutputItemsProvidersIntegrationProvider represents the integrations instances list output items providers integration provider type.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProvider struct {
	Object          string                                                                          `json:"object"`
	Id              string                                                                          `json:"id"`
	ProviderVersion IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                          `json:"status"`
	Name            string                                                                          `json:"name"`
	Description     *string                                                                         `json:"description,omitempty"`
	Metadata        *map[string]any                                                                 `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                      `json:"provider_id"`
	DeploymentId      string                                                                      `json:"deployment_id"`
	AuthMethodId      *string                                                                     `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                     `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                   `json:"created_at"`
	UpdatedAt         time.Time                                                                   `json:"updated_at"`
	ArchivedAt        *time.Time                                                                  `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersConfig represents the integrations instances list output items providers config type.
type IntegrationsInstancesListOutputItemsProvidersConfig struct {
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

// IntegrationsInstancesListOutputItemsProvidersAuthConfig represents the integrations instances list output items providers auth config type.
type IntegrationsInstancesListOutputItemsProvidersAuthConfig struct {
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

// IntegrationsInstancesListOutputItemsProviders represents the integrations instances list output items providers type.
type IntegrationsInstancesListOutputItemsProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesListOutputItemsProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                             `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesListOutputItemsProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesListOutputItemsProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesListOutputItemsProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesListOutputItemsProvidersAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                        `json:"created_at"`
	UpdatedAt            time.Time                                                        `json:"updated_at"`
	ArchivedAt           *time.Time                                                       `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputItems represents the integrations instances list output items type.
type IntegrationsInstancesListOutputItems struct {
	Object          string                                              `json:"object"`
	Id              string                                              `json:"id"`
	Status          string                                              `json:"status"`
	Name            string                                              `json:"name"`
	Description     *string                                             `json:"description,omitempty"`
	Metadata        *map[string]any                                     `json:"metadata,omitempty"`
	IntegrationId   string                                              `json:"integration_id"`
	IdentityActorId *string                                             `json:"identity_actor_id,omitempty"`
	IdentityId      *string                                             `json:"identity_id,omitempty"`
	Implementation  *IntegrationsInstancesListOutputItemsImplementation `json:"implementation,omitempty"`
	Providers       []IntegrationsInstancesListOutputItemsProviders     `json:"providers"`
	CreatedAt       time.Time                                           `json:"created_at"`
	UpdatedAt       time.Time                                           `json:"updated_at"`
	ArchivedAt      *time.Time                                          `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputPagination represents the integrations instances list output pagination type.
type IntegrationsInstancesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsInstancesListOutput represents the integrations instances list output type.
type IntegrationsInstancesListOutput struct {
	Items      []IntegrationsInstancesListOutputItems    `json:"items"`
	Pagination IntegrationsInstancesListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstancesListOutputFromJSON deserializes JSON data into a IntegrationsInstancesListOutput.
func MapIntegrationsInstancesListOutputFromJSON(data []byte) (*IntegrationsInstancesListOutput, error) {
	var v IntegrationsInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesListOutputToJSON serializes a IntegrationsInstancesListOutput to JSON.
func MapIntegrationsInstancesListOutputToJSON(v *IntegrationsInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesListQueryCreatedAt - Filter integration instance creation time by date range
type IntegrationsInstancesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration instance creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesListQueryUpdatedAt - Filter integration instance last update time by date range
type IntegrationsInstancesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration instance last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesListQuery represents the integrations instances list query type.
type IntegrationsInstancesListQuery struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	IdentityId            *any     `json:"identity_id,omitempty"`
	IdentityCredentialId  *any     `json:"identity_credential_id,omitempty"`
	IdentityActorId       *any     `json:"identity_actor_id,omitempty"`
	ProviderDeploymentId  *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId  *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId     *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance creation time by date range
	CreatedAt *IntegrationsInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance last update time by date range
	UpdatedAt *IntegrationsInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstancesListQueryFromJSON deserializes JSON data into a IntegrationsInstancesListQuery.
func MapIntegrationsInstancesListQueryFromJSON(data []byte) (*IntegrationsInstancesListQuery, error) {
	var v IntegrationsInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesListQueryToJSON serializes a IntegrationsInstancesListQuery to JSON.
func MapIntegrationsInstancesListQueryToJSON(v *IntegrationsInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
