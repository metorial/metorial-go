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

// MagicMcpServersProvidersCreateOutputProvider represents the magic mcp servers providers create output provider type.
type MagicMcpServersProvidersCreateOutputProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
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

// MagicMcpServersProvidersCreateOutputAuthMethodInputSchema represents the magic mcp servers providers create output auth method input schema type.
type MagicMcpServersProvidersCreateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema represents the magic mcp servers providers create output auth method output schema type.
type MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersCreateOutputAuthMethodScopes represents the magic mcp servers providers create output auth method scopes type.
type MagicMcpServersProvidersCreateOutputAuthMethodScopes struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique scope identifier
	Id string `json:"id"`
	// Scope - OAuth scope string
	Scope string `json:"scope"`
	// Name - Display name of the scope
	Name string `json:"name"`
	// Description - Scope description
	Description *string `json:"description,omitempty"`
}

// MagicMcpServersProvidersCreateOutputAuthMethod represents the magic mcp servers providers create output auth method type.
type MagicMcpServersProvidersCreateOutputAuthMethod struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth method identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Key - Auth method key
	Key string `json:"key"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Auth method description
	Description *string `json:"description,omitempty"`
	// Capabilities - Auth method capabilities
	Capabilities map[string]any                                              `json:"capabilities"`
	InputSchema  *MagicMcpServersProvidersCreateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersCreateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersCreateOutputAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersCreateOutputAuthCredentials represents the magic mcp servers providers create output auth credentials type.
type MagicMcpServersProvidersCreateOutputAuthCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// Status - Credentials status
	Status string `json:"status"`
	// IsDefault - Whether this is the default credentials for the provider
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether these credentials are managed by Metorial
	IsManaged bool `json:"is_managed"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
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

// MagicMcpServersProvidersCreateOutput represents the magic mcp servers providers create output type.
type MagicMcpServersProvidersCreateOutput struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersCreateOutputToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersCreateOutputProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersCreateOutputDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersCreateOutputAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersCreateOutputAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersCreateOutputConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersCreateOutputAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
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
	ProviderId           string  `json:"provider_id"`
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
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
