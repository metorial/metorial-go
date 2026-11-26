package magicmcpservers

import (
	"encoding/json"
	"time"
)

// MagicMcpServersListOutputItemsEndpoints represents the magic mcp servers list output items endpoints type.
type MagicMcpServersListOutputItemsEndpoints struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Alias - Connection slug for the server.
	Alias string `json:"alias"`
	// Url - Connection URL for the server.
	Url string `json:"url"`
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

// MagicMcpServersListOutputItemsProvidersDeployment represents the magic mcp servers list output items providers deployment type.
type MagicMcpServersListOutputItemsProvidersDeployment struct {
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

// MagicMcpServersListOutputItemsProvidersAuthMethod represents the magic mcp servers list output items providers auth method type.
type MagicMcpServersListOutputItemsProvidersAuthMethod struct {
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
	InputSchema  *MagicMcpServersListOutputItemsProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpServersListOutputItemsProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpServersListOutputItemsProvidersAuthMethodScopes `json:"scopes,omitempty"`
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

// MagicMcpServersListOutputItemsProvidersAuthCredentials represents the magic mcp servers list output items providers auth credentials type.
type MagicMcpServersListOutputItemsProvidersAuthCredentials struct {
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

// MagicMcpServersListOutputItemsProvidersConfig represents the magic mcp servers list output items providers config type.
type MagicMcpServersListOutputItemsProvidersConfig struct {
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

// MagicMcpServersListOutputItemsProvidersAuthConfig represents the magic mcp servers list output items providers auth config type.
type MagicMcpServersListOutputItemsProvidersAuthConfig struct {
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

// MagicMcpServersListOutputItemsProviders represents the magic mcp servers list output items providers type.
type MagicMcpServersListOutputItemsProviders struct {
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
	ToolFilter      *MagicMcpServersListOutputItemsProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpServersListOutputItemsProvidersProvider         `json:"provider"`
	Deployment      MagicMcpServersListOutputItemsProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpServersListOutputItemsProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpServersListOutputItemsProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpServersListOutputItemsProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpServersListOutputItemsProvidersAuthConfig      `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MagicMcpServersListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpServersListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpServersListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpServersListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpServersListOutputItemsAccessOwner - The actor that owns this resource, if any.
type MagicMcpServersListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpServersListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpServersListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpServersListOutputItemsAccess - Resource ownership and access settings.
type MagicMcpServersListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpServersListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// MagicMcpServersListOutputItems represents the magic mcp servers list output items type.
type MagicMcpServersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the server. Archived servers stay listable.
	Status string `json:"status"`
	// Source - How this Magic MCP server was created or configured
	Source string `json:"source"`
	// ProviderManagementMode - Whether providers are managed on the server or inherited.
	ProviderManagementMode string `json:"provider_management_mode"`
	// Endpoints - Connection targets for the server.
	Endpoints []MagicMcpServersListOutputItemsEndpoints `json:"endpoints"`
	// ProviderTemplateId - Provider template the server is built from, if any.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// Providers - Providers the server connects through.
	Providers []MagicMcpServersListOutputItemsProviders `json:"providers"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Access - Resource ownership and access settings.
	Access MagicMcpServersListOutputItemsAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpServersListOutputPagination represents the magic mcp servers list output pagination type.
type MagicMcpServersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpServersListOutput represents the magic mcp servers list output type.
type MagicMcpServersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
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
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'usable' for consumers, and otherwise to 'shared' when ownership is enforced for the project and 'all' when it isn't.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Only return servers with one of these statuses.
	Status *any `json:"status,omitempty"`
	// MagicMcpGroupId - Only return servers in these groups.
	MagicMcpGroupId *any `json:"magic_mcp_group_id,omitempty"`
	// ProviderTemplateId - Only return servers built from these provider templates.
	ProviderTemplateId *any `json:"provider_template_id,omitempty"`
	// IntegrationInstanceId - Only return servers built from these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// Owner - Only return organization-owned or consumer-owned servers.
	Owner *any `json:"owner,omitempty"`
	// ProviderId - Only return servers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ConsumerId - Only return servers owned by these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// ConsumerProfileId - Only return servers owned by these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
	// Id - Only return servers with these IDs.
	Id *any `json:"id,omitempty"`
	// PreconfiguredOnly - Only return preconfigured servers.
	PreconfiguredOnly *bool `json:"preconfigured_only,omitempty"`
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
