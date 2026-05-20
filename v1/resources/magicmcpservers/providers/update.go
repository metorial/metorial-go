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

// MagicMcpServersProvidersUpdateOutputProvider represents the magic mcp servers providers update output provider type.
type MagicMcpServersProvidersUpdateOutputProvider struct {
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

// MagicMcpServersProvidersUpdateOutputAuthMethodInputSchema represents the magic mcp servers providers update output auth method input schema type.
type MagicMcpServersProvidersUpdateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersUpdateOutputAuthMethodOutputSchema represents the magic mcp servers providers update output auth method output schema type.
type MagicMcpServersProvidersUpdateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersUpdateOutputAuthMethodScopes represents the magic mcp servers providers update output auth method scopes type.
type MagicMcpServersProvidersUpdateOutputAuthMethodScopes struct {
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

// MagicMcpServersProvidersUpdateOutputAuthMethod represents the magic mcp servers providers update output auth method type.
type MagicMcpServersProvidersUpdateOutputAuthMethod struct {
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
	InputSchema  *MagicMcpServersProvidersUpdateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersUpdateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersUpdateOutputAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersUpdateOutputAuthCredentials represents the magic mcp servers providers update output auth credentials type.
type MagicMcpServersProvidersUpdateOutputAuthCredentials struct {
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

// MagicMcpServersProvidersUpdateOutput represents the magic mcp servers providers update output type.
type MagicMcpServersProvidersUpdateOutput struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersUpdateOutputToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersUpdateOutputProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersUpdateOutputDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersUpdateOutputAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersUpdateOutputAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersUpdateOutputConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersUpdateOutputAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                            `json:"created_at"`
	UpdatedAt       time.Time                                            `json:"updated_at"`
	ArchivedAt      *time.Time                                           `json:"archived_at,omitempty"`
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
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	ProviderConfigId     *string `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	ToolFilters          *any    `json:"tool_filters,omitempty"`
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
