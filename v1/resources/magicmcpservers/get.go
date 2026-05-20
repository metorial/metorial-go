package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersGetOutputEndpoints represents the magic mcp servers get output endpoints type.
type MagicMcpServersGetOutputEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpServersGetOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersGetOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersGetOutputProvidersProvider represents the magic mcp servers get output providers provider type.
type MagicMcpServersGetOutputProvidersProvider struct {
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

// MagicMcpServersGetOutputProvidersDeployment represents the magic mcp servers get output providers deployment type.
type MagicMcpServersGetOutputProvidersDeployment struct {
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

// MagicMcpServersGetOutputProvidersAuthMethodInputSchema represents the magic mcp servers get output providers auth method input schema type.
type MagicMcpServersGetOutputProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersGetOutputProvidersAuthMethodOutputSchema represents the magic mcp servers get output providers auth method output schema type.
type MagicMcpServersGetOutputProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersGetOutputProvidersAuthMethodScopes represents the magic mcp servers get output providers auth method scopes type.
type MagicMcpServersGetOutputProvidersAuthMethodScopes struct {
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

// MagicMcpServersGetOutputProvidersAuthMethod represents the magic mcp servers get output providers auth method type.
type MagicMcpServersGetOutputProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersGetOutputProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersGetOutputProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersGetOutputProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersGetOutputProvidersAuthCredentials represents the magic mcp servers get output providers auth credentials type.
type MagicMcpServersGetOutputProvidersAuthCredentials struct {
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

// MagicMcpServersGetOutputProvidersConfig represents the magic mcp servers get output providers config type.
type MagicMcpServersGetOutputProvidersConfig struct {
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

// MagicMcpServersGetOutputProvidersAuthConfig represents the magic mcp servers get output providers auth config type.
type MagicMcpServersGetOutputProvidersAuthConfig struct {
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

// MagicMcpServersGetOutputProviders represents the magic mcp servers get output providers type.
type MagicMcpServersGetOutputProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersGetOutputProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersGetOutputProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersGetOutputProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersGetOutputProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersGetOutputProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersGetOutputProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersGetOutputProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                         `json:"created_at"`
	UpdatedAt       time.Time                                         `json:"updated_at"`
	ArchivedAt      *time.Time                                        `json:"archived_at,omitempty"`
}

// MagicMcpServersGetOutput represents the magic mcp servers get output type.
type MagicMcpServersGetOutput struct {
	Object                 string                              `json:"object"`
	Id                     string                              `json:"id"`
	Status                 string                              `json:"status"`
	Source                 string                              `json:"source"`
	ProviderManagementMode string                              `json:"provider_management_mode"`
	Endpoints              []MagicMcpServersGetOutputEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                             `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpServersGetOutputProviders `json:"providers"`
	Name                   *string                             `json:"name,omitempty"`
	Description            *string                             `json:"description,omitempty"`
	Metadata               map[string]any                      `json:"metadata"`
	CreatedAt              time.Time                           `json:"created_at"`
	UpdatedAt              time.Time                           `json:"updated_at"`
}

// MapMagicMcpServersGetOutputFromJSON deserializes JSON data into a MagicMcpServersGetOutput.
func MapMagicMcpServersGetOutputFromJSON(data []byte) (*MagicMcpServersGetOutput, error) {
	var v MagicMcpServersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersGetOutputToJSON serializes a MagicMcpServersGetOutput to JSON.
func MapMagicMcpServersGetOutputToJSON(v *MagicMcpServersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
