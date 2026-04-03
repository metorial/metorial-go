package provider

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProviderGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProviderGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProviderGetOutputDeployment represents the magic mcp servers provider get output deployment type.
type MagicMcpServersProviderGetOutputDeployment struct {
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

// MagicMcpServersProviderGetOutputConfig represents the magic mcp servers provider get output config type.
type MagicMcpServersProviderGetOutputConfig struct {
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

// MagicMcpServersProviderGetOutputAuthConfig represents the magic mcp servers provider get output auth config type.
type MagicMcpServersProviderGetOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProviderGetOutput represents the magic mcp servers provider get output type.
type MagicMcpServersProviderGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProviderGetOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                      `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProviderGetOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProviderGetOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProviderGetOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProviderGetOutputFromJSON deserializes JSON data into a MagicMcpServersProviderGetOutput.
func MapMagicMcpServersProviderGetOutputFromJSON(data []byte) (*MagicMcpServersProviderGetOutput, error) {
	var v MagicMcpServersProviderGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderGetOutputToJSON serializes a MagicMcpServersProviderGetOutput to JSON.
func MapMagicMcpServersProviderGetOutputToJSON(v *MagicMcpServersProviderGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
