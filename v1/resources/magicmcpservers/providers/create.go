package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersCreateOutputDeployment represents the magic mcp servers providers create output deployment type.
type MagicMcpServersProvidersCreateOutputDeployment struct {
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

// MagicMcpServersProvidersCreateOutputConfig represents the magic mcp servers providers create output config type.
type MagicMcpServersProvidersCreateOutputConfig struct {
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

// MagicMcpServersProvidersCreateOutputAuthConfig represents the magic mcp servers providers create output auth config type.
type MagicMcpServersProvidersCreateOutputAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// MagicMcpServersProvidersCreateOutput represents the magic mcp servers providers create output type.
type MagicMcpServersProvidersCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique magic MCP server provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter MagicMcpServersProvidersCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// MagicMcpServerId - Parent magic MCP server ID
	MagicMcpServerId string                                          `json:"magic_mcp_server_id"`
	Deployment       MagicMcpServersProvidersCreateOutputDeployment  `json:"deployment"`
	Config           MagicMcpServersProvidersCreateOutputConfig      `json:"config"`
	AuthConfig       *MagicMcpServersProvidersCreateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapMagicMcpServersProvidersCreateOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersCreateOutput.
func MapMagicMcpServersProvidersCreateOutputFromJSON(data []byte) (*MagicMcpServersProvidersCreateOutput, error) {
	var v MagicMcpServersProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersCreateOutputToJSON serializes a MagicMcpServersProvidersCreateOutput to JSON.
func MapMagicMcpServersProvidersCreateOutputToJSON(v *MagicMcpServersProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProvidersCreateBody represents the magic mcp servers providers create body type.
type MagicMcpServersProvidersCreateBody struct {
	ProviderDeploymentId  *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId      *string `json:"provider_config_id,omitempty"`
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	ProviderAuthConfigId  *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters           *any    `json:"tool_filters,omitempty"`
}

// MapMagicMcpServersProvidersCreateBodyFromJSON deserializes JSON data into a MagicMcpServersProvidersCreateBody.
func MapMagicMcpServersProvidersCreateBodyFromJSON(data []byte) (*MagicMcpServersProvidersCreateBody, error) {
	var v MagicMcpServersProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersCreateBodyToJSON serializes a MagicMcpServersProvidersCreateBody to JSON.
func MapMagicMcpServersProvidersCreateBodyToJSON(v *MagicMcpServersProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
