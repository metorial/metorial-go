package provider

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProviderUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProviderUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProviderUpdateOutputDeployment represents the magic mcp servers provider update output deployment type.
type MagicMcpServersProviderUpdateOutputDeployment struct {
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

// MagicMcpServersProviderUpdateOutputConfig represents the magic mcp servers provider update output config type.
type MagicMcpServersProviderUpdateOutputConfig struct {
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

// MagicMcpServersProviderUpdateOutputAuthConfig represents the magic mcp servers provider update output auth config type.
type MagicMcpServersProviderUpdateOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProviderUpdateOutput represents the magic mcp servers provider update output type.
type MagicMcpServersProviderUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProviderUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                         `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProviderUpdateOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProviderUpdateOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProviderUpdateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProviderUpdateOutputFromJSON deserializes JSON data into a MagicMcpServersProviderUpdateOutput.
func MapMagicMcpServersProviderUpdateOutputFromJSON(data []byte) (*MagicMcpServersProviderUpdateOutput, error) {
	var v MagicMcpServersProviderUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderUpdateOutputToJSON serializes a MagicMcpServersProviderUpdateOutput to JSON.
func MapMagicMcpServersProviderUpdateOutputToJSON(v *MagicMcpServersProviderUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProviderUpdateBody represents the magic mcp servers provider update body type.
type MagicMcpServersProviderUpdateBody struct {
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapMagicMcpServersProviderUpdateBodyFromJSON deserializes JSON data into a MagicMcpServersProviderUpdateBody.
func MapMagicMcpServersProviderUpdateBodyFromJSON(data []byte) (*MagicMcpServersProviderUpdateBody, error) {
	var v MagicMcpServersProviderUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderUpdateBodyToJSON serializes a MagicMcpServersProviderUpdateBody to JSON.
func MapMagicMcpServersProviderUpdateBodyToJSON(v *MagicMcpServersProviderUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
