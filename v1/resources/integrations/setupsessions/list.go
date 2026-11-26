package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation - What backs this instance. `null` for a standalone instance.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpServerId - Magic MCP server backing this instance.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider represents the integrations setup sessions list output items integration instance providers provider type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions list output items integration instance providers integration provider config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions list output items integration instance providers integration provider type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                                                        `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig represents the integrations setup sessions list output items integration instance providers config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions list output items integration instance providers auth config type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig struct {
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

// IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders represents the integrations setup sessions list output items integration instance providers type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the entry is archived.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Integration instance this entry belongs to.
	IntegrationInstanceId string `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                                                    `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItemsIntegrationInstance represents the integrations setup sessions list output items integration instance type.
type IntegrationsSetupSessionsListOutputItemsIntegrationInstance struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `draft` when created. `archived` once you archive the instance.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IdentityActorId - Identity actor the instance connects as. `null` when none is set.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance connects as. `null` when none is set.
	IdentityId *string `json:"identity_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access IntegrationsSetupSessionsListOutputItemsIntegrationInstanceAccess `json:"access"`
	// Implementation - What backs this instance. `null` for a standalone instance.
	Implementation *IntegrationsSetupSessionsListOutputItemsIntegrationInstanceImplementation `json:"implementation,omitempty"`
	// Providers - Providers set up on this instance.
	Providers []IntegrationsSetupSessionsListOutputItemsIntegrationInstanceProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsListOutputItems represents the integrations setup sessions list output items type.
type IntegrationsSetupSessionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `pending` until every provider is configured. `successful` once every provider is configured. `expired` past `expires_at` without completion.
	Status string `json:"status"`
	// Url - Hosted setup page URL.
	Url string `json:"url"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Configuration for the setup flow. `null` when none was set.
	Configuration *map[string]any `json:"configuration,omitempty"`
	// RedirectUrl - Where to send the user once setup completes. `null` when none was given.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// IntegrationId - Integration being set up.
	IntegrationId string `json:"integration_id"`
	// Access - Resource ownership and access settings.
	Access              IntegrationsSetupSessionsListOutputItemsAccess              `json:"access"`
	IntegrationInstance IntegrationsSetupSessionsListOutputItemsIntegrationInstance `json:"integration_instance"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ExpiresAt - When the session expires. Past this time the session reads as `expired`.
	ExpiresAt time.Time `json:"expires_at"`
}

// IntegrationsSetupSessionsListOutputPagination represents the integrations setup sessions list output pagination type.
type IntegrationsSetupSessionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsSetupSessionsListOutput represents the integrations setup sessions list output type.
type IntegrationsSetupSessionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                        `json:"object"`
	Items      []IntegrationsSetupSessionsListOutputItems    `json:"items"`
	Pagination IntegrationsSetupSessionsListOutputPagination `json:"pagination"`
}

// MapIntegrationsSetupSessionsListOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsListOutput.
func MapIntegrationsSetupSessionsListOutputFromJSON(data []byte) (*IntegrationsSetupSessionsListOutput, error) {
	var v IntegrationsSetupSessionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsListOutputToJSON serializes a IntegrationsSetupSessionsListOutput to JSON.
func MapIntegrationsSetupSessionsListOutputToJSON(v *IntegrationsSetupSessionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsSetupSessionsListQueryCreatedAt - Only return results whose integration setup session creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsSetupSessionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsSetupSessionsListQueryUpdatedAt - Only return results whose integration setup session last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsSetupSessionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsSetupSessionsListQuery represents the integrations setup sessions list query type.
type IntegrationsSetupSessionsListQuery struct {
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
	// Status - Only return setup sessions with one of these statuses. `deleted` sessions are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return setup sessions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return setup sessions for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return setup sessions for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// CreatedAt - Only return results whose integration setup session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsSetupSessionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration setup session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsSetupSessionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsSetupSessionsListQueryFromJSON deserializes JSON data into a IntegrationsSetupSessionsListQuery.
func MapIntegrationsSetupSessionsListQueryFromJSON(data []byte) (*IntegrationsSetupSessionsListQuery, error) {
	var v IntegrationsSetupSessionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsListQueryToJSON serializes a IntegrationsSetupSessionsListQuery to JSON.
func MapIntegrationsSetupSessionsListQueryToJSON(v *IntegrationsSetupSessionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
