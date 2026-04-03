package provider

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProviderDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProviderDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProviderDeleteOutputDeployment represents the magic mcp servers provider delete output deployment type.
type MagicMcpServersProviderDeleteOutputDeployment struct {
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

// MagicMcpServersProviderDeleteOutputConfig represents the magic mcp servers provider delete output config type.
type MagicMcpServersProviderDeleteOutputConfig struct {
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

// MagicMcpServersProviderDeleteOutputAuthConfig represents the magic mcp servers provider delete output auth config type.
type MagicMcpServersProviderDeleteOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProviderDeleteOutput represents the magic mcp servers provider delete output type.
type MagicMcpServersProviderDeleteOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProviderDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                         `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProviderDeleteOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProviderDeleteOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProviderDeleteOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProviderDeleteOutputFromJSON deserializes JSON data into a MagicMcpServersProviderDeleteOutput.
func MapMagicMcpServersProviderDeleteOutputFromJSON(data []byte) (*MagicMcpServersProviderDeleteOutput, error) {
	var v MagicMcpServersProviderDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderDeleteOutputToJSON serializes a MagicMcpServersProviderDeleteOutput to JSON.
func MapMagicMcpServersProviderDeleteOutputToJSON(v *MagicMcpServersProviderDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
