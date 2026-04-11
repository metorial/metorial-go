package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersListOutputItemsDeployment represents the magic mcp servers providers list output items deployment type.
type MagicMcpServersProvidersListOutputItemsDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// MagicMcpServersProvidersListOutputItemsConfig represents the magic mcp servers providers list output items config type.
type MagicMcpServersProvidersListOutputItemsConfig struct {
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

// MagicMcpServersProvidersListOutputItemsAuthConfig represents the magic mcp servers providers list output items auth config type.
type MagicMcpServersProvidersListOutputItemsAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProvidersListOutputItems represents the magic mcp servers providers list output items type.
type MagicMcpServersProvidersListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProvidersListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                             `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProvidersListOutputItemsDeployment  `json:"deployment"`
	Config           MagicMcpServersProvidersListOutputItemsConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProvidersListOutputItemsAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputPagination represents the magic mcp servers providers list output pagination type.
type MagicMcpServersProvidersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpServersProvidersListOutput represents the magic mcp servers providers list output type.
type MagicMcpServersProvidersListOutput struct {
	Items      []MagicMcpServersProvidersListOutputItems    `json:"items"`
	Pagination MagicMcpServersProvidersListOutputPagination `json:"pagination"`
}

// MapMagicMcpServersProvidersListOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersListOutput.
func MapMagicMcpServersProvidersListOutputFromJSON(data []byte) (*MagicMcpServersProvidersListOutput, error) {
	var v MagicMcpServersProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersListOutputToJSON serializes a MagicMcpServersProvidersListOutput to JSON.
func MapMagicMcpServersProvidersListOutputToJSON(v *MagicMcpServersProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProvidersListQueryCreatedAt - Filter magic MCP server provider creation time by date range
type MagicMcpServersProvidersListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQueryUpdatedAt - Filter magic MCP server provider last update time by date range
type MagicMcpServersProvidersListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQuery represents the magic mcp servers providers list query type.
type MagicMcpServersProvidersListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Status *any     `json:"status,omitempty"`
	// Id - Filter by magic MCP server provider ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Filter magic MCP server provider creation time by date range
	CreatedAt *MagicMcpServersProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter magic MCP server provider last update time by date range
	UpdatedAt *MagicMcpServersProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapMagicMcpServersProvidersListQueryFromJSON deserializes JSON data into a MagicMcpServersProvidersListQuery.
func MapMagicMcpServersProvidersListQueryFromJSON(data []byte) (*MagicMcpServersProvidersListQuery, error) {
	var v MagicMcpServersProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersListQueryToJSON serializes a MagicMcpServersProvidersListQuery to JSON.
func MapMagicMcpServersProvidersListQueryToJSON(v *MagicMcpServersProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
