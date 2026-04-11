package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersDeleteOutputDeployment represents the magic mcp servers providers delete output deployment type.
type MagicMcpServersProvidersDeleteOutputDeployment struct {
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

// MagicMcpServersProvidersDeleteOutputConfig represents the magic mcp servers providers delete output config type.
type MagicMcpServersProvidersDeleteOutputConfig struct {
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

// MagicMcpServersProvidersDeleteOutputAuthConfig represents the magic mcp servers providers delete output auth config type.
type MagicMcpServersProvidersDeleteOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProvidersDeleteOutput represents the magic mcp servers providers delete output type.
type MagicMcpServersProvidersDeleteOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProvidersDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                          `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProvidersDeleteOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProvidersDeleteOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProvidersDeleteOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProvidersDeleteOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersDeleteOutput.
func MapMagicMcpServersProvidersDeleteOutputFromJSON(data []byte) (*MagicMcpServersProvidersDeleteOutput, error) {
	var v MagicMcpServersProvidersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersDeleteOutputToJSON serializes a MagicMcpServersProvidersDeleteOutput to JSON.
func MapMagicMcpServersProvidersDeleteOutputToJSON(v *MagicMcpServersProvidersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
