package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersUpdateOutputDeployment represents the magic mcp servers providers update output deployment type.
type MagicMcpServersProvidersUpdateOutputDeployment struct {
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

// MagicMcpServersProvidersUpdateOutputConfig represents the magic mcp servers providers update output config type.
type MagicMcpServersProvidersUpdateOutputConfig struct {
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

// MagicMcpServersProvidersUpdateOutputAuthConfig represents the magic mcp servers providers update output auth config type.
type MagicMcpServersProvidersUpdateOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProvidersUpdateOutput represents the magic mcp servers providers update output type.
type MagicMcpServersProvidersUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProvidersUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                          `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProvidersUpdateOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProvidersUpdateOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProvidersUpdateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProvidersUpdateOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersUpdateOutput.
func MapMagicMcpServersProvidersUpdateOutputFromJSON(data []byte) (*MagicMcpServersProvidersUpdateOutput, error) {
	var v MagicMcpServersProvidersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersUpdateOutputToJSON serializes a MagicMcpServersProvidersUpdateOutput to JSON.
func MapMagicMcpServersProvidersUpdateOutputToJSON(v *MagicMcpServersProvidersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProvidersUpdateBody represents the magic mcp servers providers update body type.
type MagicMcpServersProvidersUpdateBody struct {
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapMagicMcpServersProvidersUpdateBodyFromJSON deserializes JSON data into a MagicMcpServersProvidersUpdateBody.
func MapMagicMcpServersProvidersUpdateBodyFromJSON(data []byte) (*MagicMcpServersProvidersUpdateBody, error) {
	var v MagicMcpServersProvidersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersUpdateBodyToJSON serializes a MagicMcpServersProvidersUpdateBody to JSON.
func MapMagicMcpServersProvidersUpdateBodyToJSON(v *MagicMcpServersProvidersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
