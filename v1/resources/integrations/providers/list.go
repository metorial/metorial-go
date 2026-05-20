package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersListOutputItemsConfig represents the integrations providers list output items config type.
type IntegrationsProvidersListOutputItemsConfig struct {
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

// IntegrationsProvidersListOutputItems represents the integrations providers list output items type.
type IntegrationsProvidersListOutputItems struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsProvidersListOutputItemsToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                          `json:"provider_id"`
	DeploymentId      string                                          `json:"deployment_id"`
	AuthMethodId      *string                                         `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                         `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersListOutputItemsConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                       `json:"created_at"`
	UpdatedAt         time.Time                                       `json:"updated_at"`
	ArchivedAt        *time.Time                                      `json:"archived_at,omitempty"`
}

// IntegrationsProvidersListOutputPagination represents the integrations providers list output pagination type.
type IntegrationsProvidersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsProvidersListOutput represents the integrations providers list output type.
type IntegrationsProvidersListOutput struct {
	Items      []IntegrationsProvidersListOutputItems    `json:"items"`
	Pagination IntegrationsProvidersListOutputPagination `json:"pagination"`
}

// MapIntegrationsProvidersListOutputFromJSON deserializes JSON data into a IntegrationsProvidersListOutput.
func MapIntegrationsProvidersListOutputFromJSON(data []byte) (*IntegrationsProvidersListOutput, error) {
	var v IntegrationsProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersListOutputToJSON serializes a IntegrationsProvidersListOutput to JSON.
func MapIntegrationsProvidersListOutputToJSON(v *IntegrationsProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsProvidersListQueryCreatedAt - Filter integration provider creation time by date range
type IntegrationsProvidersListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsProvidersListQueryUpdatedAt - Filter integration provider last update time by date range
type IntegrationsProvidersListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsProvidersListQuery represents the integrations providers list query type.
type IntegrationsProvidersListQuery struct {
	Limit                     *float64 `json:"limit,omitempty"`
	After                     *string  `json:"after,omitempty"`
	Before                    *string  `json:"before,omitempty"`
	Cursor                    *string  `json:"cursor,omitempty"`
	Order                     *string  `json:"order,omitempty"`
	Search                    *string  `json:"search,omitempty"`
	Status                    *any     `json:"status,omitempty"`
	Id                        *any     `json:"id,omitempty"`
	IntegrationId             *any     `json:"integration_id,omitempty"`
	ProviderId                *any     `json:"provider_id,omitempty"`
	ProviderDeploymentId      *any     `json:"provider_deployment_id,omitempty"`
	ProviderAuthMethodId      *any     `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *any     `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *any     `json:"provider_config_id,omitempty"`
	// CreatedAt - Filter integration provider creation time by date range
	CreatedAt *IntegrationsProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration provider last update time by date range
	UpdatedAt *IntegrationsProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsProvidersListQueryFromJSON deserializes JSON data into a IntegrationsProvidersListQuery.
func MapIntegrationsProvidersListQueryFromJSON(data []byte) (*IntegrationsProvidersListQuery, error) {
	var v IntegrationsProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersListQueryToJSON serializes a IntegrationsProvidersListQuery to JSON.
func MapIntegrationsProvidersListQueryToJSON(v *IntegrationsProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
