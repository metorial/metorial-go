package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersCreateOutputEndpoints represents the magic mcp servers create output endpoints type.
type MagicMcpServersCreateOutputEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpServersCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersCreateOutputProvidersProvider represents the magic mcp servers create output providers provider type.
type MagicMcpServersCreateOutputProvidersProvider struct {
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

// MagicMcpServersCreateOutputProvidersDeployment represents the magic mcp servers create output providers deployment type.
type MagicMcpServersCreateOutputProvidersDeployment struct {
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

// MagicMcpServersCreateOutputProvidersAuthMethodInputSchema represents the magic mcp servers create output providers auth method input schema type.
type MagicMcpServersCreateOutputProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema represents the magic mcp servers create output providers auth method output schema type.
type MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersCreateOutputProvidersAuthMethodScopes represents the magic mcp servers create output providers auth method scopes type.
type MagicMcpServersCreateOutputProvidersAuthMethodScopes struct {
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

// MagicMcpServersCreateOutputProvidersAuthMethod represents the magic mcp servers create output providers auth method type.
type MagicMcpServersCreateOutputProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersCreateOutputProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersCreateOutputProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersCreateOutputProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersCreateOutputProvidersAuthCredentials represents the magic mcp servers create output providers auth credentials type.
type MagicMcpServersCreateOutputProvidersAuthCredentials struct {
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

// MagicMcpServersCreateOutputProvidersConfig represents the magic mcp servers create output providers config type.
type MagicMcpServersCreateOutputProvidersConfig struct {
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

// MagicMcpServersCreateOutputProvidersAuthConfig represents the magic mcp servers create output providers auth config type.
type MagicMcpServersCreateOutputProvidersAuthConfig struct {
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

// MagicMcpServersCreateOutputProviders represents the magic mcp servers create output providers type.
type MagicMcpServersCreateOutputProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersCreateOutputProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersCreateOutputProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersCreateOutputProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersCreateOutputProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersCreateOutputProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersCreateOutputProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersCreateOutputProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
}

// MagicMcpServersCreateOutput represents the magic mcp servers create output type.
type MagicMcpServersCreateOutput struct {
	Object                 string                                 `json:"object"`
	Id                     string                                 `json:"id"`
	Status                 string                                 `json:"status"`
	Source                 string                                 `json:"source"`
	ProviderManagementMode string                                 `json:"provider_management_mode"`
	Endpoints              []MagicMcpServersCreateOutputEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpServersCreateOutputProviders `json:"providers"`
	Name                   *string                                `json:"name,omitempty"`
	Description            *string                                `json:"description,omitempty"`
	Metadata               map[string]any                         `json:"metadata"`
	CreatedAt              time.Time                              `json:"created_at"`
	UpdatedAt              time.Time                              `json:"updated_at"`
}

// MapMagicMcpServersCreateOutputFromJSON deserializes JSON data into a MagicMcpServersCreateOutput.
func MapMagicMcpServersCreateOutputFromJSON(data []byte) (*MagicMcpServersCreateOutput, error) {
	var v MagicMcpServersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersCreateOutputToJSON serializes a MagicMcpServersCreateOutput to JSON.
func MapMagicMcpServersCreateOutputToJSON(v *MagicMcpServersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersCreateBody represents the magic mcp servers create body type.
type MagicMcpServersCreateBody struct {
	Name               *string         `json:"name,omitempty"`
	Description        *string         `json:"description,omitempty"`
	Metadata           *map[string]any `json:"metadata,omitempty"`
	ProviderTemplateId *string         `json:"provider_template_id,omitempty"`
	ConsumerProfileId  *string         `json:"consumer_profile_id,omitempty"`
}

// MapMagicMcpServersCreateBodyFromJSON deserializes JSON data into a MagicMcpServersCreateBody.
func MapMagicMcpServersCreateBodyFromJSON(data []byte) (*MagicMcpServersCreateBody, error) {
	var v MagicMcpServersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersCreateBodyToJSON serializes a MagicMcpServersCreateBody to JSON.
func MapMagicMcpServersCreateBodyToJSON(v *MagicMcpServersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
