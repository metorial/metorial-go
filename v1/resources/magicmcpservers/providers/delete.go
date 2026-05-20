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

// MagicMcpServersProvidersDeleteOutputProvider represents the magic mcp servers providers delete output provider type.
type MagicMcpServersProvidersDeleteOutputProvider struct {
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

// MagicMcpServersProvidersDeleteOutputAuthMethodInputSchema represents the magic mcp servers providers delete output auth method input schema type.
type MagicMcpServersProvidersDeleteOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersDeleteOutputAuthMethodOutputSchema represents the magic mcp servers providers delete output auth method output schema type.
type MagicMcpServersProvidersDeleteOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersDeleteOutputAuthMethodScopes represents the magic mcp servers providers delete output auth method scopes type.
type MagicMcpServersProvidersDeleteOutputAuthMethodScopes struct {
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

// MagicMcpServersProvidersDeleteOutputAuthMethod represents the magic mcp servers providers delete output auth method type.
type MagicMcpServersProvidersDeleteOutputAuthMethod struct {
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
	InputSchema  *MagicMcpServersProvidersDeleteOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersDeleteOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersDeleteOutputAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersDeleteOutputAuthCredentials represents the magic mcp servers providers delete output auth credentials type.
type MagicMcpServersProvidersDeleteOutputAuthCredentials struct {
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

// MagicMcpServersProvidersDeleteOutput represents the magic mcp servers providers delete output type.
type MagicMcpServersProvidersDeleteOutput struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersDeleteOutputToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersDeleteOutputProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersDeleteOutputDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersDeleteOutputAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersDeleteOutputAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersDeleteOutputConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersDeleteOutputAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
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
