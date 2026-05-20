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

// MagicMcpServersProvidersGetOutputProvider represents the magic mcp servers providers get output provider type.
type MagicMcpServersProvidersGetOutputProvider struct {
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

// MagicMcpServersProvidersGetOutputAuthMethodInputSchema represents the magic mcp servers providers get output auth method input schema type.
type MagicMcpServersProvidersGetOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersGetOutputAuthMethodOutputSchema represents the magic mcp servers providers get output auth method output schema type.
type MagicMcpServersProvidersGetOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersGetOutputAuthMethodScopes represents the magic mcp servers providers get output auth method scopes type.
type MagicMcpServersProvidersGetOutputAuthMethodScopes struct {
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

// MagicMcpServersProvidersGetOutputAuthMethod represents the magic mcp servers providers get output auth method type.
type MagicMcpServersProvidersGetOutputAuthMethod struct {
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
	Capabilities map[string]any                                           `json:"capabilities"`
	InputSchema  *MagicMcpServersProvidersGetOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersGetOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersGetOutputAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersGetOutputAuthCredentials represents the magic mcp servers providers get output auth credentials type.
type MagicMcpServersProvidersGetOutputAuthCredentials struct {
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

// MagicMcpServersProvidersGetOutput represents the magic mcp servers providers get output type.
type MagicMcpServersProvidersGetOutput struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersGetOutputToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersGetOutputProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersGetOutputDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersGetOutputAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersGetOutputAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersGetOutputConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersGetOutputAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                         `json:"created_at"`
	UpdatedAt       time.Time                                         `json:"updated_at"`
	ArchivedAt      *time.Time                                        `json:"archived_at,omitempty"`
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
