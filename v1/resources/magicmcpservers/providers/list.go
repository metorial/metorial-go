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
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItemsDeployment represents the magic mcp servers providers list output items deployment type.
type MagicMcpServersProvidersListOutputItemsDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Scope - The scope value sent to the provider.
	Scope string `json:"scope"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// MagicMcpServersProvidersListOutputItemsAuthMethod represents the magic mcp servers providers list output items auth method type.
type MagicMcpServersProvidersListOutputItemsAuthMethod struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the user authenticates with the provider.
	Type string `json:"type"`
	// Key - Stable key for the auth method.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - Capabilities of this auth method in the provider version the deployment currently uses: its locked version, otherwise the provider's current version. The other fields describe the pinned auth method and may come from an older specification.
	Capabilities map[string]any                                                 `json:"capabilities"`
	InputSchema  *MagicMcpServersProvidersListOutputItemsAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersProvidersListOutputItemsAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersProvidersListOutputItemsAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider this auth method belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification this auth method belongs to.
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItemsAuthCredentials represents the magic mcp servers providers list output items auth credentials type.
type MagicMcpServersProvidersListOutputItemsAuthCredentials struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Must be `oauth`.
	Type string `json:"type"`
	// Status - Archived once the credentials are archived. Archived credentials stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name credentials.
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether Metorial or the installation administrator manages these OAuth credentials.
	IsManaged bool `json:"is_managed"`
	// ManagedBy - Who manages these credentials: Metorial, or the administrator of this self-hosted installation. Null for credentials created in the project.
	ManagedBy *string `json:"managed_by,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider these credentials belong to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItemsConfig represents the magic mcp servers providers list output items config type.
type MagicMcpServersProvidersListOutputItemsConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItemsAuthConfig represents the magic mcp servers providers list output items auth config type.
type MagicMcpServersProvidersListOutputItemsAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersProvidersListOutputItems represents the magic mcp servers providers list output items type.
type MagicMcpServersProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you remove the provider from the server.
	Status string `json:"status"`
	// MagicMcpServerId - Server the provider belongs to.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
	// ProviderManagementMode - Whether the provider is managed here or inherited.
	ProviderManagementMode string `json:"provider_management_mode"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter      *MagicMcpServersProvidersListOutputItemsToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersProvidersListOutputItemsProvider         `json:"provider"`
	Deployment      MagicMcpServersProvidersListOutputItemsDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersProvidersListOutputItemsAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersProvidersListOutputItemsAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersProvidersListOutputItemsConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersProvidersListOutputItemsAuthConfig      `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MagicMcpServersProvidersListOutputPagination represents the magic mcp servers providers list output pagination type.
type MagicMcpServersProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpServersProvidersListOutput represents the magic mcp servers providers list output type.
type MagicMcpServersProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                       `json:"object"`
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

// MagicMcpServersProvidersListQueryCreatedAt - Only return results whose magic MCP server provider creation time is in this range. Set `gt`, `lt`, or both.
type MagicMcpServersProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQueryUpdatedAt - Only return results whose magic MCP server provider last update time is in this range. Set `gt`, `lt`, or both.
type MagicMcpServersProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MagicMcpServersProvidersListQuery represents the magic mcp servers providers list query type.
type MagicMcpServersProvidersListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Status - Only return providers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return providers with these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// IntegrationInstanceProviderId - Only return providers using one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderDeploymentId - Only return providers using one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return providers using one of these configurations.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return providers using one of these auth configurations.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose magic MCP server provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *MagicMcpServersProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose magic MCP server provider last update time is in this range. Set `gt`, `lt`, or both.
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
