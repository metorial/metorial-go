package provider

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProviderCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProviderCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProviderCreateOutputDeployment represents the magic mcp servers provider create output deployment type.
type MagicMcpServersProviderCreateOutputDeployment struct {
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

// MagicMcpServersProviderCreateOutputConfig represents the magic mcp servers provider create output config type.
type MagicMcpServersProviderCreateOutputConfig struct {
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

// MagicMcpServersProviderCreateOutputAuthConfig represents the magic mcp servers provider create output auth config type.
type MagicMcpServersProviderCreateOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProviderCreateOutput represents the magic mcp servers provider create output type.
type MagicMcpServersProviderCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProviderCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                         `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProviderCreateOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProviderCreateOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProviderCreateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProviderCreateOutputFromJSON deserializes JSON data into a MagicMcpServersProviderCreateOutput.
func MapMagicMcpServersProviderCreateOutputFromJSON(data []byte) (*MagicMcpServersProviderCreateOutput, error) {
	var v MagicMcpServersProviderCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderCreateOutputToJSON serializes a MagicMcpServersProviderCreateOutput to JSON.
func MapMagicMcpServersProviderCreateOutputToJSON(v *MagicMcpServersProviderCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProviderCreateBody represents the magic mcp servers provider create body type.
type MagicMcpServersProviderCreateBody struct {
	ProviderDeploymentId  *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *string `json:"provider_config_id,omitempty"`
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	ProviderAuthConfigId  *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters           *any    `json:"tool_filters,omitempty"`
}

// MapMagicMcpServersProviderCreateBodyFromJSON deserializes JSON data into a MagicMcpServersProviderCreateBody.
func MapMagicMcpServersProviderCreateBodyFromJSON(data []byte) (*MagicMcpServersProviderCreateBody, error) {
	var v MagicMcpServersProviderCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProviderCreateBodyToJSON serializes a MagicMcpServersProviderCreateBody to JSON.
func MapMagicMcpServersProviderCreateBodyToJSON(v *MagicMcpServersProviderCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
