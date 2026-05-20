package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsListOutputItemsConfiguration represents the integrations list output items configuration type.
type IntegrationsListOutputItemsConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// IntegrationsListOutputItemsImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsListOutputItemsImplementation struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsListOutputItemsProvidersConfig represents the integrations list output items providers config type.
type IntegrationsListOutputItemsProvidersConfig struct {
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

// IntegrationsListOutputItemsProviders represents the integrations list output items providers type.
type IntegrationsListOutputItemsProviders struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsListOutputItemsProvidersToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                          `json:"provider_id"`
	DeploymentId      string                                          `json:"deployment_id"`
	AuthMethodId      *string                                         `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                         `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsListOutputItemsProvidersConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                       `json:"created_at"`
	UpdatedAt         time.Time                                       `json:"updated_at"`
	ArchivedAt        *time.Time                                      `json:"archived_at,omitempty"`
}

// IntegrationsListOutputItems represents the integrations list output items type.
type IntegrationsListOutputItems struct {
	Object         string                                     `json:"object"`
	Id             string                                     `json:"id"`
	Status         string                                     `json:"status"`
	Slug           string                                     `json:"slug"`
	Name           string                                     `json:"name"`
	Description    *string                                    `json:"description,omitempty"`
	Metadata       *map[string]any                            `json:"metadata,omitempty"`
	Configuration  IntegrationsListOutputItemsConfiguration   `json:"configuration"`
	Implementation *IntegrationsListOutputItemsImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsListOutputItemsProviders     `json:"providers"`
	CreatedAt      time.Time                                  `json:"created_at"`
	UpdatedAt      time.Time                                  `json:"updated_at"`
	ArchivedAt     *time.Time                                 `json:"archived_at,omitempty"`
}

// IntegrationsListOutputPagination represents the integrations list output pagination type.
type IntegrationsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsListOutput represents the integrations list output type.
type IntegrationsListOutput struct {
	Items      []IntegrationsListOutputItems    `json:"items"`
	Pagination IntegrationsListOutputPagination `json:"pagination"`
}

// MapIntegrationsListOutputFromJSON deserializes JSON data into a IntegrationsListOutput.
func MapIntegrationsListOutputFromJSON(data []byte) (*IntegrationsListOutput, error) {
	var v IntegrationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsListOutputToJSON serializes a IntegrationsListOutput to JSON.
func MapIntegrationsListOutputToJSON(v *IntegrationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsListQueryCreatedAt - Filter integration creation time by date range
type IntegrationsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsListQueryUpdatedAt - Filter integration last update time by date range
type IntegrationsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsListQuery represents the integrations list query type.
type IntegrationsListQuery struct {
	Limit                 *float64 `json:"limit,omitempty"`
	After                 *string  `json:"after,omitempty"`
	Before                *string  `json:"before,omitempty"`
	Cursor                *string  `json:"cursor,omitempty"`
	Order                 *string  `json:"order,omitempty"`
	Search                *string  `json:"search,omitempty"`
	Status                *any     `json:"status,omitempty"`
	Id                    *any     `json:"id,omitempty"`
	ProviderId            *any     `json:"provider_id,omitempty"`
	IntegrationProviderId *any     `json:"integration_provider_id,omitempty"`
	// CreatedAt - Filter integration creation time by date range
	CreatedAt *IntegrationsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration last update time by date range
	UpdatedAt *IntegrationsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsListQueryFromJSON deserializes JSON data into a IntegrationsListQuery.
func MapIntegrationsListQueryFromJSON(data []byte) (*IntegrationsListQuery, error) {
	var v IntegrationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsListQueryToJSON serializes a IntegrationsListQuery to JSON.
func MapIntegrationsListQueryToJSON(v *IntegrationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
