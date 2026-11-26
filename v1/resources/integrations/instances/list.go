package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstancesListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstancesListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstancesListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstancesListOutputItemsAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstancesListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstancesListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstancesListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstancesListOutputItemsAccess - Resource ownership and access settings.
type IntegrationsInstancesListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstancesListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstancesListOutputItemsImplementation - What backs this instance. `null` for a standalone instance.
type IntegrationsInstancesListOutputItemsImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpServerId - Magic MCP server backing this instance.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersProvider represents the integrations instances list output items providers provider type.
type IntegrationsInstancesListOutputItemsProvidersProvider struct {
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

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig represents the integrations instances list output items providers integration provider config type.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesListOutputItemsProvidersIntegrationProvider represents the integrations instances list output items providers integration provider type.
type IntegrationsInstancesListOutputItemsProvidersIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsInstancesListOutputItemsProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesListOutputItemsProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                                 `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesListOutputItemsProvidersIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputItemsProvidersConfig represents the integrations instances list output items providers config type.
type IntegrationsInstancesListOutputItemsProvidersConfig struct {
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

// IntegrationsInstancesListOutputItemsProvidersAuthConfig represents the integrations instances list output items providers auth config type.
type IntegrationsInstancesListOutputItemsProvidersAuthConfig struct {
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

// IntegrationsInstancesListOutputItemsProviders represents the integrations instances list output items providers type.
type IntegrationsInstancesListOutputItemsProviders struct {
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
	ToolFilter *IntegrationsInstancesListOutputItemsProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                             `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesListOutputItemsProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesListOutputItemsProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesListOutputItemsProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesListOutputItemsProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputItems represents the integrations instances list output items type.
type IntegrationsInstancesListOutputItems struct {
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
	Access IntegrationsInstancesListOutputItemsAccess `json:"access"`
	// Implementation - What backs this instance. `null` for a standalone instance.
	Implementation *IntegrationsInstancesListOutputItemsImplementation `json:"implementation,omitempty"`
	// Providers - Providers set up on this instance.
	Providers []IntegrationsInstancesListOutputItemsProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesListOutputPagination represents the integrations instances list output pagination type.
type IntegrationsInstancesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsInstancesListOutput represents the integrations instances list output type.
type IntegrationsInstancesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []IntegrationsInstancesListOutputItems    `json:"items"`
	Pagination IntegrationsInstancesListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstancesListOutputFromJSON deserializes JSON data into a IntegrationsInstancesListOutput.
func MapIntegrationsInstancesListOutputFromJSON(data []byte) (*IntegrationsInstancesListOutput, error) {
	var v IntegrationsInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesListOutputToJSON serializes a IntegrationsInstancesListOutput to JSON.
func MapIntegrationsInstancesListOutputToJSON(v *IntegrationsInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesListQueryCreatedAt - Only return results whose integration instance creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstancesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesListQueryUpdatedAt - Only return results whose integration instance last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstancesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesListQuery represents the integrations instances list query type.
type IntegrationsInstancesListQuery struct {
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
	// Ownership - Ownership scope: `shared` (not owned by someone else), `usable` (shared plus ones you may use), or `all`. Defaults to `shared` when ownership is enforced, `all` otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Search - Search integration instances by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration instances with one of these statuses. `deleted` instances are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instances for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return integration instances using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instances using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// IdentityId - Only return integration instances for one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Only return integration instances using one of these identity credentials.
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// IdentityActorId - Only return integration instances for one of these identity actors.
	IdentityActorId *any `json:"identity_actor_id,omitempty"`
	// ProviderDeploymentId - Only return integration instances running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instances using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instances using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instances using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstancesListQueryFromJSON deserializes JSON data into a IntegrationsInstancesListQuery.
func MapIntegrationsInstancesListQueryFromJSON(data []byte) (*IntegrationsInstancesListQuery, error) {
	var v IntegrationsInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesListQueryToJSON serializes a IntegrationsInstancesListQuery to JSON.
func MapIntegrationsInstancesListQueryToJSON(v *IntegrationsInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
