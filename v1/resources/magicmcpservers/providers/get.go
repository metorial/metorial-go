package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersGetOutputDeployment represents the magic mcp servers providers get output deployment type.
type MagicMcpServersProvidersGetOutputDeployment struct {
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

// MagicMcpServersProvidersGetOutputConfig represents the magic mcp servers providers get output config type.
type MagicMcpServersProvidersGetOutputConfig struct {
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

// MagicMcpServersProvidersGetOutputAuthConfig represents the magic mcp servers providers get output auth config type.
type MagicMcpServersProvidersGetOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProvidersGetOutput represents the magic mcp servers providers get output type.
type MagicMcpServersProvidersGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProvidersGetOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                       `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProvidersGetOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProvidersGetOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProvidersGetOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProvidersGetOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersGetOutput.
func MapMagicMcpServersProvidersGetOutputFromJSON(data []byte) (*MagicMcpServersProvidersGetOutput, error) {
	var v MagicMcpServersProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersGetOutputToJSON serializes a MagicMcpServersProvidersGetOutput to JSON.
func MapMagicMcpServersProvidersGetOutputToJSON(v *MagicMcpServersProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
