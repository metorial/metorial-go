package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersDeleteOutputEndpoints represents the magic mcp servers delete output endpoints type.
type MagicMcpServersDeleteOutputEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpServersDeleteOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersDeleteOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersDeleteOutputProvidersProvider represents the magic mcp servers delete output providers provider type.
type MagicMcpServersDeleteOutputProvidersProvider struct {
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

// MagicMcpServersDeleteOutputProvidersDeployment represents the magic mcp servers delete output providers deployment type.
type MagicMcpServersDeleteOutputProvidersDeployment struct {
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

// MagicMcpServersDeleteOutputProvidersAuthMethodInputSchema represents the magic mcp servers delete output providers auth method input schema type.
type MagicMcpServersDeleteOutputProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersDeleteOutputProvidersAuthMethodOutputSchema represents the magic mcp servers delete output providers auth method output schema type.
type MagicMcpServersDeleteOutputProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersDeleteOutputProvidersAuthMethodScopes represents the magic mcp servers delete output providers auth method scopes type.
type MagicMcpServersDeleteOutputProvidersAuthMethodScopes struct {
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

// MagicMcpServersDeleteOutputProvidersAuthMethod represents the magic mcp servers delete output providers auth method type.
type MagicMcpServersDeleteOutputProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersDeleteOutputProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersDeleteOutputProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersDeleteOutputProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersDeleteOutputProvidersAuthCredentials represents the magic mcp servers delete output providers auth credentials type.
type MagicMcpServersDeleteOutputProvidersAuthCredentials struct {
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

// MagicMcpServersDeleteOutputProvidersConfig represents the magic mcp servers delete output providers config type.
type MagicMcpServersDeleteOutputProvidersConfig struct {
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

// MagicMcpServersDeleteOutputProvidersAuthConfig represents the magic mcp servers delete output providers auth config type.
type MagicMcpServersDeleteOutputProvidersAuthConfig struct {
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

// MagicMcpServersDeleteOutputProviders represents the magic mcp servers delete output providers type.
type MagicMcpServersDeleteOutputProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersDeleteOutputProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersDeleteOutputProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersDeleteOutputProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersDeleteOutputProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersDeleteOutputProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersDeleteOutputProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersDeleteOutputProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
}

// MagicMcpServersDeleteOutput represents the magic mcp servers delete output type.
type MagicMcpServersDeleteOutput struct {
	Object                 string                                 `json:"object"`
	Id                     string                                 `json:"id"`
	Status                 string                                 `json:"status"`
	Source                 string                                 `json:"source"`
	ProviderManagementMode string                                 `json:"provider_management_mode"`
	Endpoints              []MagicMcpServersDeleteOutputEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpServersDeleteOutputProviders `json:"providers"`
	Name                   *string                                `json:"name,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	Metadata               map[string]any                         `json:"metadata"`
	CreatedAt              time.Time                              `json:"created_at"`
	UpdatedAt              time.Time                              `json:"updated_at"`
}

// MapMagicMcpServersDeleteOutputFromJSON deserializes JSON data into a MagicMcpServersDeleteOutput.
func MapMagicMcpServersDeleteOutputFromJSON(data []byte) (*MagicMcpServersDeleteOutput, error) {
	var v MagicMcpServersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersDeleteOutputToJSON serializes a MagicMcpServersDeleteOutput to JSON.
func MapMagicMcpServersDeleteOutputToJSON(v *MagicMcpServersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
