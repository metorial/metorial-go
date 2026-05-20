package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsProvider represents the integrations instances providers list output items provider type.
type IntegrationsInstancesProvidersListOutputItemsProvider struct {
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

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion represents the integrations instances providers list output items integration provider provider version type.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig represents the integrations instances providers list output items integration provider config type.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig struct {
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

// IntegrationsInstancesProvidersListOutputItemsIntegrationProvider represents the integrations instances providers list output items integration provider type.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProvider struct {
	Object          string                                                                          `json:"object"`
	Id              string                                                                          `json:"id"`
	ProviderVersion IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                          `json:"status"`
	Name            string                                                                          `json:"name"`
	Description     *string                                                                         `json:"description,omitempty"`
	Metadata        *map[string]any                                                                 `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                      `json:"provider_id"`
	DeploymentId      string                                                                      `json:"deployment_id"`
	AuthMethodId      *string                                                                     `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                     `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                   `json:"created_at"`
	UpdatedAt         time.Time                                                                   `json:"updated_at"`
	ArchivedAt        *time.Time                                                                  `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsConfig represents the integrations instances providers list output items config type.
type IntegrationsInstancesProvidersListOutputItemsConfig struct {
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

// IntegrationsInstancesProvidersListOutputItemsAuthConfig represents the integrations instances providers list output items auth config type.
type IntegrationsInstancesProvidersListOutputItemsAuthConfig struct {
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

// IntegrationsInstancesProvidersListOutputItems represents the integrations instances providers list output items type.
type IntegrationsInstancesProvidersListOutputItems struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesProvidersListOutputItemsToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                             `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesProvidersListOutputItemsProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesProvidersListOutputItemsIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesProvidersListOutputItemsConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesProvidersListOutputItemsAuthConfig         `json:"auth_config,omitempty"`
	CreatedAt            time.Time                                                        `json:"created_at"`
	UpdatedAt            time.Time                                                        `json:"updated_at"`
	ArchivedAt           *time.Time                                                       `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersListOutputPagination represents the integrations instances providers list output pagination type.
type IntegrationsInstancesProvidersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsInstancesProvidersListOutput represents the integrations instances providers list output type.
type IntegrationsInstancesProvidersListOutput struct {
	Items      []IntegrationsInstancesProvidersListOutputItems    `json:"items"`
	Pagination IntegrationsInstancesProvidersListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstancesProvidersListOutputFromJSON deserializes JSON data into a IntegrationsInstancesProvidersListOutput.
func MapIntegrationsInstancesProvidersListOutputFromJSON(data []byte) (*IntegrationsInstancesProvidersListOutput, error) {
	var v IntegrationsInstancesProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersListOutputToJSON serializes a IntegrationsInstancesProvidersListOutput to JSON.
func MapIntegrationsInstancesProvidersListOutputToJSON(v *IntegrationsInstancesProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesProvidersListQueryCreatedAt - Filter integration instance provider creation time by date range
type IntegrationsInstancesProvidersListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration instance provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesProvidersListQueryUpdatedAt - Filter integration instance provider last update time by date range
type IntegrationsInstancesProvidersListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration instance provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesProvidersListQuery represents the integrations instances providers list query type.
type IntegrationsInstancesProvidersListQuery struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	IntegrationId         *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId *any     `json:"integration_instance_id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId  *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId  *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId     *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance provider creation time by date range
	CreatedAt *IntegrationsInstancesProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance provider last update time by date range
	UpdatedAt *IntegrationsInstancesProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstancesProvidersListQueryFromJSON deserializes JSON data into a IntegrationsInstancesProvidersListQuery.
func MapIntegrationsInstancesProvidersListQueryFromJSON(data []byte) (*IntegrationsInstancesProvidersListQuery, error) {
	var v IntegrationsInstancesProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersListQueryToJSON serializes a IntegrationsInstancesProvidersListQuery to JSON.
func MapIntegrationsInstancesProvidersListQueryToJSON(v *IntegrationsInstancesProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
