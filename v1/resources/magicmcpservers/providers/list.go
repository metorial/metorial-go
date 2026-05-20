package providers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersProvidersListOutputItemsProvider represents the magic mcp servers providers list output items provider type.
type MagicMcpServersProvidersListOutputItemsProvider struct {
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

// MagicMcpServersProvidersListOutputItemsDeployment represents the magic mcp servers providers list output items deployment type.
type MagicMcpServersProvidersListOutputItemsDeployment struct {
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

// MagicMcpServersProvidersListOutputItemsAuthMethodInputSchema represents the magic mcp servers providers list output items auth method input schema type.
type MagicMcpServersProvidersListOutputItemsAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersListOutputItemsAuthMethodOutputSchema represents the magic mcp servers providers list output items auth method output schema type.
type MagicMcpServersProvidersListOutputItemsAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersProvidersListOutputItemsAuthMethodScopes represents the magic mcp servers providers list output items auth method scopes type.
type MagicMcpServersProvidersListOutputItemsAuthMethodScopes struct {
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

// MagicMcpServersProvidersListOutputItemsAuthMethod represents the magic mcp servers providers list output items auth method type.
type MagicMcpServersProvidersListOutputItemsAuthMethod struct {
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
	Capabilities map[string]any                                                 `json:"capabilities"`
	InputSchema  *MagicMcpServersProvidersListOutputItemsAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersListOutputItemsAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersListOutputItemsAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItemsAuthCredentials represents the magic mcp servers providers list output items auth credentials type.
type MagicMcpServersProvidersListOutputItemsAuthCredentials struct {
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

// MagicMcpServersProvidersListOutputItemsConfig represents the magic mcp servers providers list output items config type.
type MagicMcpServersProvidersListOutputItemsConfig struct {
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

// MagicMcpServersProvidersListOutputItemsAuthConfig represents the magic mcp servers providers list output items auth config type.
type MagicMcpServersProvidersListOutputItemsAuthConfig struct {
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

// MagicMcpServersProvidersListOutputItems represents the magic mcp servers providers list output items type.
type MagicMcpServersProvidersListOutputItems struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersListOutputItemsToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersListOutputItemsProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersListOutputItemsDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersListOutputItemsAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersListOutputItemsAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersListOutputItemsConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersListOutputItemsAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                               `json:"created_at"`
	UpdatedAt       time.Time                                               `json:"updated_at"`
	ArchivedAt      *time.Time                                              `json:"archived_at,omitempty"`
}

// MagicMcpServersProvidersListOutputPagination represents the magic mcp servers providers list output pagination type.
type MagicMcpServersProvidersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpServersProvidersListOutput represents the magic mcp servers providers list output type.
type MagicMcpServersProvidersListOutput struct {
	Items      []MagicMcpServersProvidersListOutputItems    `json:"items"`
	Pagination MagicMcpServersProvidersListOutputPagination `json:"pagination"`
}

// MapMagicMcpServersProvidersListOutputFromJSON deserializes JSON data into a MagicMcpServersProvidersListOutput.
func MapMagicMcpServersProvidersListOutputFromJSON(data []byte) (*MagicMcpServersProvidersListOutput, error) {
	var v MagicMcpServersProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersListOutputToJSON serializes a MagicMcpServersProvidersListOutput to JSON.
func MapMagicMcpServersProvidersListOutputToJSON(v *MagicMcpServersProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersProvidersListQueryCreatedAt - Filter magic MCP server provider creation time by date range
type MagicMcpServersProvidersListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQueryUpdatedAt - Filter magic MCP server provider last update time by date range
type MagicMcpServersProvidersListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for magic MCP server provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for magic MCP server provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQuery represents the magic mcp servers providers list query type.
type MagicMcpServersProvidersListQuery struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Filter magic MCP server provider creation time by date range
	CreatedAt *MagicMcpServersProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter magic MCP server provider last update time by date range
	UpdatedAt *MagicMcpServersProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapMagicMcpServersProvidersListQueryFromJSON deserializes JSON data into a MagicMcpServersProvidersListQuery.
func MapMagicMcpServersProvidersListQueryFromJSON(data []byte) (*MagicMcpServersProvidersListQuery, error) {
	var v MagicMcpServersProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersProvidersListQueryToJSON serializes a MagicMcpServersProvidersListQuery to JSON.
func MapMagicMcpServersProvidersListQueryToJSON(v *MagicMcpServersProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
