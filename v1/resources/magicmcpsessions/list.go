package magicmcpsessions

import (
	"encoding/json"
	"time"
)

// MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints represents the magic mcp sessions list output items magic mcp server endpoints type.
type MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersProvider represents the magic mcp sessions list output items magic mcp server providers provider type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersProvider struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment represents the magic mcp sessions list output items magic mcp server providers deployment type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodInputSchema represents the magic mcp sessions list output items magic mcp server providers auth method input schema type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodOutputSchema represents the magic mcp sessions list output items magic mcp server providers auth method output schema type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodScopes represents the magic mcp sessions list output items magic mcp server providers auth method scopes type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodScopes struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod represents the magic mcp sessions list output items magic mcp server providers auth method type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod struct {
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
	Capabilities map[string]any                                                                `json:"capabilities"`
	InputSchema  *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials represents the magic mcp sessions list output items magic mcp server providers auth credentials type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig represents the magic mcp sessions list output items magic mcp server providers config type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig represents the magic mcp sessions list output items magic mcp server providers auth config type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProviders represents the magic mcp sessions list output items magic mcp server providers type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpSessionsListOutputItemsMagicMcpServerProvidersProvider         `json:"provider"`
	Deployment      MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                                              `json:"created_at"`
	UpdatedAt       time.Time                                                              `json:"updated_at"`
	ArchivedAt      *time.Time                                                             `json:"archived_at,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServer represents the magic mcp sessions list output items magic mcp server type.
type MagicMcpSessionsListOutputItemsMagicMcpServer struct {
	Object                 string                                                   `json:"object"`
	Id                     string                                                   `json:"id"`
	Status                 string                                                   `json:"status"`
	Source                 string                                                   `json:"source"`
	ProviderManagementMode string                                                   `json:"provider_management_mode"`
	Endpoints              []MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                                  `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpSessionsListOutputItemsMagicMcpServerProviders `json:"providers"`
	Name                   *string                                                  `json:"name,omitempty"`
	Description            *string                                                  `json:"description,omitempty"`
	Metadata               map[string]any                                           `json:"metadata"`
	CreatedAt              time.Time                                                `json:"created_at"`
	UpdatedAt              time.Time                                                `json:"updated_at"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpoint represents the magic mcp sessions list output items magic mcp endpoint type.
type MagicMcpSessionsListOutputItemsMagicMcpEndpoint struct {
	Object      string           `json:"object"`
	Id          string           `json:"id"`
	Status      string           `json:"status"`
	Slug        string           `json:"slug"`
	Url         string           `json:"url"`
	Servers     []map[string]any `json:"servers"`
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Metadata    map[string]any   `json:"metadata"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`
}

// MagicMcpSessionsListOutputItems represents the magic mcp sessions list output items type.
type MagicMcpSessionsListOutputItems struct {
	Object                 string                                           `json:"object"`
	Id                     string                                           `json:"id"`
	MagicMcpServer         *MagicMcpSessionsListOutputItemsMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	MagicMcpEndpoint       *MagicMcpSessionsListOutputItemsMagicMcpEndpoint `json:"magic_mcp_endpoint,omitempty"`
	ConsumerProfileId      *string                                          `json:"consumer_profile_id,omitempty"`
	ConsumerIntegrationIds []string                                         `json:"consumer_integration_ids"`
	SessionId              string                                           `json:"session_id"`
	ExpiresAt              *time.Time                                       `json:"expires_at,omitempty"`
	CreatedAt              time.Time                                        `json:"created_at"`
	UpdatedAt              time.Time                                        `json:"updated_at"`
}

// MagicMcpSessionsListOutputPagination represents the magic mcp sessions list output pagination type.
type MagicMcpSessionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpSessionsListOutput represents the magic mcp sessions list output type.
type MagicMcpSessionsListOutput struct {
	Items      []MagicMcpSessionsListOutputItems    `json:"items"`
	Pagination MagicMcpSessionsListOutputPagination `json:"pagination"`
}

// MapMagicMcpSessionsListOutputFromJSON deserializes JSON data into a MagicMcpSessionsListOutput.
func MapMagicMcpSessionsListOutputFromJSON(data []byte) (*MagicMcpSessionsListOutput, error) {
	var v MagicMcpSessionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpSessionsListOutputToJSON serializes a MagicMcpSessionsListOutput to JSON.
func MapMagicMcpSessionsListOutputToJSON(v *MagicMcpSessionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpSessionsListQuery represents the magic mcp sessions list query type.
type MagicMcpSessionsListQuery struct {
	Limit            *float64 `json:"limit,omitempty"`
	After            *string  `json:"after,omitempty"`
	Before           *string  `json:"before,omitempty"`
	Cursor           *string  `json:"cursor,omitempty"`
	Order            *string  `json:"order,omitempty"`
	MagicMcpServerId *any     `json:"magic_mcp_server_id,omitempty"`
}

// MapMagicMcpSessionsListQueryFromJSON deserializes JSON data into a MagicMcpSessionsListQuery.
func MapMagicMcpSessionsListQueryFromJSON(data []byte) (*MagicMcpSessionsListQuery, error) {
	var v MagicMcpSessionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpSessionsListQueryToJSON serializes a MagicMcpSessionsListQuery to JSON.
func MapMagicMcpSessionsListQueryToJSON(v *MagicMcpSessionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
