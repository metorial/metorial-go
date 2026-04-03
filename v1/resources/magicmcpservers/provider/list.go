package provider

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProviderListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProviderListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProviderListOutputItemsDeployment represents the magic mcp servers provider list output items deployment type.
type MagicMcpServersProviderListOutputItemsDeployment struct {
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

// MagicMcpServersProviderListOutputItemsConfig represents the magic mcp servers provider list output items config type.
type MagicMcpServersProviderListOutputItemsConfig struct {
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

// MagicMcpServersProviderListOutputItemsAuthConfig represents the magic mcp servers provider list output items auth config type.
type MagicMcpServersProviderListOutputItemsAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProviderListOutputItems represents the magic mcp servers provider list output items type.
type MagicMcpServersProviderListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProviderListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                            `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProviderListOutputItemsDeployment  `json:"deployment"`
	Config           MagicMcpServersProviderListOutputItemsConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProviderListOutputItemsAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProviderListOutputPagination represents the magic mcp servers provider list output pagination type.
type MagicMcpServersProviderListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpServersProviderListOutput represents the magic mcp servers provider list output type.
type MagicMcpServersProviderListOutput struct {
	Items      []MagicMcpServersProviderListOutputItems    `json:"items"`
	Pagination MagicMcpServersProviderListOutputPagination `json:"pagination"`
}

// MapMagicMcpServersProviderListOutputFromJSON deserializes JSON data into a MagicMcpServersProviderListOutput.
func MapMagicMcpServersProviderListOutputFromJSON(data []byte) (*MagicMcpServersProviderListOutput, error) {
	var v MagicMcpServersProviderListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderListOutputToJSON serializes a MagicMcpServersProviderListOutput to JSON.
func MapMagicMcpServersProviderListOutputToJSON(v *MagicMcpServersProviderListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProviderListQueryCreatedAt - Filter magic MCP server provider creation time by date range
type MagicMcpServersProviderListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProviderListQueryUpdatedAt - Filter magic MCP server provider last update time by date range
type MagicMcpServersProviderListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProviderListQuery represents the magic mcp servers provider list query type.
type MagicMcpServersProviderListQuery struct {
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
	CreatedAt *MagicMcpServersProviderListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter magic MCP server provider last update time by date range
	UpdatedAt *MagicMcpServersProviderListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapMagicMcpServersProviderListQueryFromJSON deserializes JSON data into a MagicMcpServersProviderListQuery.
func MapMagicMcpServersProviderListQueryFromJSON(data []byte) (*MagicMcpServersProviderListQuery, error) {
	var v MagicMcpServersProviderListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderListQueryToJSON serializes a MagicMcpServersProviderListQuery to JSON.
func MapMagicMcpServersProviderListQueryToJSON(v *MagicMcpServersProviderListQuery) ([]byte, error) {
	return json.Marshal(v)
}
