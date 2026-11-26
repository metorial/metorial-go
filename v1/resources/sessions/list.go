package sessions

import (
	"encoding/json"
	"time"
)

// SessionsListOutputItemsUsage - Usage statistics.
type SessionsListOutputItemsUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsListOutputItemsProvidersUsage - Usage statistics
type SessionsListOutputItemsProvidersUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsListOutputItemsProvidersDeployment represents the sessions list output items providers deployment type.
type SessionsListOutputItemsProvidersDeployment struct {
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

// SessionsListOutputItemsProvidersConfig represents the sessions list output items providers config type.
type SessionsListOutputItemsProvidersConfig struct {
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

// SessionsListOutputItemsProvidersAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsListOutputItemsProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsListOutputItemsProviders represents the sessions list output items providers type.
type SessionsListOutputItemsProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsListOutputItemsProvidersUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsListOutputItemsProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                                    `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsListOutputItemsProvidersDeployment `json:"deployment"`
	Config                 SessionsListOutputItemsProvidersConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsListOutputItemsProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type SessionsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionsListOutputItemsAccess - Resource ownership and access settings.
type SessionsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// SessionsListOutputItems represents the sessions list output items type.
type SessionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session is deleted. Archived sessions stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ConnectionState - Whether the session currently holds an open connection.
	ConnectionState string `json:"connection_state"`
	// ConnectionUrl - MCP connection URL for this session.
	ConnectionUrl string `json:"connection_url"`
	// ClientSecret - Secret for scoped client access. `null` when the caller cannot use this session.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Usage - Usage statistics.
	Usage SessionsListOutputItemsUsage `json:"usage"`
	// Providers - Providers connected to this session.
	Providers []SessionsListOutputItemsProviders `json:"providers"`
	// FromTemplateIds - Templates the session was created from. Empty when created without one.
	FromTemplateIds []string `json:"from_template_ids"`
	// HasErrors - `true` when the session has errors.
	HasErrors bool `json:"has_errors"`
	// HasWarnings - `true` when the session has warnings.
	HasWarnings bool `json:"has_warnings"`
	// IdentityActorId - Identity actor connected to the session. `null` when there is none.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access SessionsListOutputItemsAccess `json:"access"`
	// IdentityId - Identity connected to the session. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsListOutputPagination represents the sessions list output pagination type.
type SessionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsListOutput represents the sessions list output type.
type SessionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                       `json:"object"`
	Items      []SessionsListOutputItems    `json:"items"`
	Pagination SessionsListOutputPagination `json:"pagination"`
}

// MapSessionsListOutputFromJSON deserializes JSON data into a SessionsListOutput.
func MapSessionsListOutputFromJSON(data []byte) (*SessionsListOutput, error) {
	var v SessionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsListOutputToJSON serializes a SessionsListOutput to JSON.
func MapSessionsListOutputToJSON(v *SessionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsListQueryCreatedAt - Only return results whose session creation time is in this range. Set `gt`, `lt`, or both.
type SessionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsListQueryUpdatedAt - Only return results whose session last update time is in this range. Set `gt`, `lt`, or both.
type SessionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsListQuery represents the sessions list query type.
type SessionsListQuery struct {
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
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by session ID(s)
	Id *any `json:"id,omitempty"`
	// AgentId - Filter by connected agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by connected identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by connected consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by connected identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// SessionTemplateId - Filter by session template ID(s)
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsListQueryUpdatedAt `json:"updated_at,omitempty"`
	// Search - Search by session ID, or by the name of a related provider, auth method, deployment, config, auth config, agent, actor, identity, or consumer
	Search *string `json:"search,omitempty"`
}

// MapSessionsListQueryFromJSON deserializes JSON data into a SessionsListQuery.
func MapSessionsListQueryFromJSON(data []byte) (*SessionsListQuery, error) {
	var v SessionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsListQueryToJSON serializes a SessionsListQuery to JSON.
func MapSessionsListQueryToJSON(v *SessionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
