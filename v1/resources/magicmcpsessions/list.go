package magicmcpsessions

import (
	"encoding/json"
	"time"
)

// MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints represents the magic mcp sessions list output items magic mcp server endpoints type.
type MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints struct {
	// Id - Unique ID.
	Id string `json:"id"`
	// Alias - Connection slug for the server.
	Alias string `json:"alias"`
	// Url - Connection URL for the server.
	Url string `json:"url"`
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment represents the magic mcp sessions list output items magic mcp server providers deployment type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod represents the magic mcp sessions list output items magic mcp server providers auth method type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod struct {
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
	Capabilities map[string]any                                                                `json:"capabilities"`
	InputSchema  *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethodScopes `json:"scopes,omitempty"`
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials represents the magic mcp sessions list output items magic mcp server providers auth credentials type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig represents the magic mcp sessions list output items magic mcp server providers config type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig represents the magic mcp sessions list output items magic mcp server providers auth config type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig struct {
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

// MagicMcpSessionsListOutputItemsMagicMcpServerProviders represents the magic mcp sessions list output items magic mcp server providers type.
type MagicMcpSessionsListOutputItemsMagicMcpServerProviders struct {
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
	ToolFilter      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersToolFilter      `json:"tool_filter,omitempty"`
	Provider        MagicMcpSessionsListOutputItemsMagicMcpServerProvidersProvider         `json:"provider"`
	Deployment      MagicMcpSessionsListOutputItemsMagicMcpServerProvidersDeployment       `json:"deployment"`
	AuthMethod      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthMethod      `json:"auth_method,omitempty"`
	AuthCredentials *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthCredentials `json:"auth_credentials,omitempty"`
	Config          *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersConfig          `json:"config,omitempty"`
	AuthConfig      *MagicMcpSessionsListOutputItemsMagicMcpServerProvidersAuthConfig      `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwner - The actor that owns this resource, if any.
type MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServerAccess - Resource ownership and access settings.
type MagicMcpSessionsListOutputItemsMagicMcpServerAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpSessionsListOutputItemsMagicMcpServerAccessOwner `json:"owner,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpServer represents the magic mcp sessions list output items magic mcp server type.
type MagicMcpSessionsListOutputItemsMagicMcpServer struct {
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
	Endpoints []MagicMcpSessionsListOutputItemsMagicMcpServerEndpoints `json:"endpoints"`
	// ProviderTemplateId - Provider template the server is built from, if any.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// Providers - Providers the server connects through.
	Providers []MagicMcpSessionsListOutputItemsMagicMcpServerProviders `json:"providers"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Access - Resource ownership and access settings.
	Access MagicMcpSessionsListOutputItemsMagicMcpServerAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpointSkillSelection represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type MagicMcpSessionsListOutputItemsMagicMcpEndpointSkillSelection struct {
	// Type - Only the listed skills and the skills in the listed groups are exposed.
	Type *string `json:"type,omitempty"`
	// SkillIds - IDs of the selected skills.
	SkillIds *[]string `json:"skill_ids,omitempty"`
	// SkillGroupIds - IDs of the selected skill groups.
	SkillGroupIds *[]string `json:"skill_group_ids,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerOrganizationMember - Organization member associated with this owner.
type MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerConsumer - Consumer and profile associated with this owner.
type MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwner - The actor that owns this resource, if any.
type MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwnerConsumer `json:"consumer,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpointAccess - Resource ownership and access settings.
type MagicMcpSessionsListOutputItemsMagicMcpEndpointAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *MagicMcpSessionsListOutputItemsMagicMcpEndpointAccessOwner `json:"owner,omitempty"`
}

// MagicMcpSessionsListOutputItemsMagicMcpEndpoint represents the magic mcp sessions list output items magic mcp endpoint type.
type MagicMcpSessionsListOutputItemsMagicMcpEndpoint struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the endpoint. Archived endpoints stay listable.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Url - Connection URL for the endpoint.
	Url string `json:"url"`
	// Servers - Servers the endpoint routes to.
	Servers []map[string]any `json:"servers"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// EnableSkillEditing - Whether consumers connecting through this endpoint with a consumer magic MCP token get tools to read and edit the skills they have access to. The tools are hidden while skills are disabled for the instance.
	EnableSkillEditing bool `json:"enable_skill_editing"`
	// EnableSkillResources - Whether consumers connecting through this endpoint with a consumer magic MCP token can read the skills they have access to over MCP, as skill:// resources with skills/list and skills/get. Skills are hidden while skills are disabled for the instance.
	EnableSkillResources bool `json:"enable_skill_resources"`
	// SkillSelection - Which skills the endpoint exposes, both as skill resources and through the skill editing tools.
	SkillSelection MagicMcpSessionsListOutputItemsMagicMcpEndpointSkillSelection `json:"skill_selection"`
	// Access - Resource ownership and access settings.
	Access MagicMcpSessionsListOutputItemsMagicMcpEndpointAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpSessionsListOutputItems represents the magic mcp sessions list output items type.
type MagicMcpSessionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id               string                                           `json:"id"`
	MagicMcpServer   *MagicMcpSessionsListOutputItemsMagicMcpServer   `json:"magic_mcp_server,omitempty"`
	MagicMcpEndpoint *MagicMcpSessionsListOutputItemsMagicMcpEndpoint `json:"magic_mcp_endpoint,omitempty"`
	// ConsumerProfileId - Consumer profile the session connects as, if any.
	ConsumerProfileId *string `json:"consumer_profile_id,omitempty"`
	// ConsumerIntegrationIds - Consumer integrations the session connects with.
	ConsumerIntegrationIds []string `json:"consumer_integration_ids"`
	// SessionId - Session backing the connection.
	SessionId string `json:"session_id"`
	// ExpiresAt - When the session expires. `null` if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpSessionsListOutputPagination represents the magic mcp sessions list output pagination type.
type MagicMcpSessionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpSessionsListOutput represents the magic mcp sessions list output type.
type MagicMcpSessionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
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
	// MagicMcpServerId - Only return sessions for one of these servers.
	MagicMcpServerId *any `json:"magic_mcp_server_id,omitempty"`
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
