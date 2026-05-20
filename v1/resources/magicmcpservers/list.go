package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersListOutputItemsEndpoints represents the magic mcp servers list output items endpoints type.
type MagicMcpServersListOutputItemsEndpoints struct {
	Id    string `json:"id"`
	Alias string `json:"alias"`
	Url   string `json:"url"`
}

// MagicMcpServersListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpServersListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// MagicMcpServersListOutputItemsProvidersProvider represents the magic mcp servers list output items providers provider type.
type MagicMcpServersListOutputItemsProvidersProvider struct {
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

// MagicMcpServersListOutputItemsProvidersDeployment represents the magic mcp servers list output items providers deployment type.
type MagicMcpServersListOutputItemsProvidersDeployment struct {
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

// MagicMcpServersListOutputItemsProvidersAuthMethodInputSchema represents the magic mcp servers list output items providers auth method input schema type.
type MagicMcpServersListOutputItemsProvidersAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersListOutputItemsProvidersAuthMethodOutputSchema represents the magic mcp servers list output items providers auth method output schema type.
type MagicMcpServersListOutputItemsProvidersAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// MagicMcpServersListOutputItemsProvidersAuthMethodScopes represents the magic mcp servers list output items providers auth method scopes type.
type MagicMcpServersListOutputItemsProvidersAuthMethodScopes struct {
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

// MagicMcpServersListOutputItemsProvidersAuthMethod represents the magic mcp servers list output items providers auth method type.
type MagicMcpServersListOutputItemsProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersListOutputItemsProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersListOutputItemsProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersListOutputItemsProvidersAuthMethodScopes `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersListOutputItemsProvidersAuthCredentials represents the magic mcp servers list output items providers auth credentials type.
type MagicMcpServersListOutputItemsProvidersAuthCredentials struct {
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

// MagicMcpServersListOutputItemsProvidersConfig represents the magic mcp servers list output items providers config type.
type MagicMcpServersListOutputItemsProvidersConfig struct {
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

// MagicMcpServersListOutputItemsProvidersAuthConfig represents the magic mcp servers list output items providers auth config type.
type MagicMcpServersListOutputItemsProvidersAuthConfig struct {
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

// MagicMcpServersListOutputItemsProviders represents the magic mcp servers list output items providers type.
type MagicMcpServersListOutputItemsProviders struct {
	Object                 string          `json:"object"`
	Id                     string          `json:"id"`
	Status                 string          `json:"status"`
	MagicMcpServerId       string          `json:"magic_mcp_server_id"`
	ProviderManagementMode string          `json:"provider_management_mode"`
	Name                   string          `json:"name"`
	Description            *string         `json:"description,omitempty"`
	Metadata               *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersListOutputItemsProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersListOutputItemsProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersListOutputItemsProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersListOutputItemsProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersListOutputItemsProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersListOutputItemsProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersListOutputItemsProvidersAuthConfig      `json:"auth_config,omitempty"`
	CreatedAt       time.Time                                               `json:"created_at"`
	UpdatedAt       time.Time                                               `json:"updated_at"`
	ArchivedAt      *time.Time                                              `json:"archived_at,omitempty"`
}

// MagicMcpServersListOutputItems represents the magic mcp servers list output items type.
type MagicMcpServersListOutputItems struct {
	Object                 string                                    `json:"object"`
	Id                     string                                    `json:"id"`
	Status                 string                                    `json:"status"`
	Source                 string                                    `json:"source"`
	ProviderManagementMode string                                    `json:"provider_management_mode"`
	Endpoints              []MagicMcpServersListOutputItemsEndpoints `json:"endpoints"`
	ProviderTemplateId     *string                                   `json:"provider_template_id,omitempty"`
	Providers              []MagicMcpServersListOutputItemsProviders `json:"providers"`
	Name                   *string                                   `json:"name,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	Metadata               map[string]any                            `json:"metadata"`
	CreatedAt              time.Time                                 `json:"created_at"`
	UpdatedAt              time.Time                                 `json:"updated_at"`
}

// MagicMcpServersListOutputPagination represents the magic mcp servers list output pagination type.
type MagicMcpServersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MagicMcpServersListOutput represents the magic mcp servers list output type.
type MagicMcpServersListOutput struct {
	Items      []MagicMcpServersListOutputItems    `json:"items"`
	Pagination MagicMcpServersListOutputPagination `json:"pagination"`
}

// MapMagicMcpServersListOutputFromJSON deserializes JSON data into a MagicMcpServersListOutput.
func MapMagicMcpServersListOutputFromJSON(data []byte) (*MagicMcpServersListOutput, error) {
	var v MagicMcpServersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersListOutputToJSON serializes a MagicMcpServersListOutput to JSON.
func MapMagicMcpServersListOutputToJSON(v *MagicMcpServersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpServersListQuery represents the magic mcp servers list query type.
type MagicMcpServersListQuery struct {
	Limit              *float64 `json:"limit,omitempty"`
	After              *string  `json:"after,omitempty"`
	Before             *string  `json:"before,omitempty"`
	Cursor             *string  `json:"cursor,omitempty"`
	Order              *string  `json:"order,omitempty"`
	Status             *any     `json:"status,omitempty"`
	MagicMcpGroupId    *any     `json:"magic_mcp_group_id,omitempty"`
	ProviderTemplateId *any     `json:"provider_template_id,omitempty"`
	ConsumerId         *any     `json:"consumer_id,omitempty"`
	ConsumerProfileId  *any     `json:"consumer_profile_id,omitempty"`
	Search             *string  `json:"search,omitempty"`
	Id                 *any     `json:"id,omitempty"`
	PreconfiguredOnly  *bool    `json:"preconfigured_only,omitempty"`
}

// MapMagicMcpServersListQueryFromJSON deserializes JSON data into a MagicMcpServersListQuery.
func MapMagicMcpServersListQueryFromJSON(data []byte) (*MagicMcpServersListQuery, error) {
	var v MagicMcpServersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpServersListQueryToJSON serializes a MagicMcpServersListQuery to JSON.
func MapMagicMcpServersListQueryToJSON(v *MagicMcpServersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
