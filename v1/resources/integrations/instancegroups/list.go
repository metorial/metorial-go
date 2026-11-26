package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsListOutputItemsImplementation - What backs this group. `null` for a standalone group.
type IntegrationsInstanceGroupsListOutputItemsImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpEndpointId - Magic MCP endpoint backing this group.
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// IntegrationsInstanceGroupsListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItemsProviders represents the integrations instance groups list output items providers type.
type IntegrationsInstanceGroupsListOutputItemsProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the entry.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this entry belongs to.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceGroupId - Group this entry belongs to.
	IntegrationInstanceGroupId string `json:"integration_instance_group_id"`
	// IntegrationInstanceId - Instance this entry’s provider comes from.
	IntegrationInstanceId string `json:"integration_instance_id"`
	// IntegrationProviderId - Integration provider this entry was taken from. `null` when the entry has no integration provider.
	IntegrationProviderId *string `json:"integration_provider_id,omitempty"`
	// IntegrationInstanceProviderId - Integration instance provider this entry links.
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstanceGroupsListOutputItemsProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItemsAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItems represents the integrations instance groups list output items type.
type IntegrationsInstanceGroupsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `draft` when created. `archived` once you archive the group.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Implementation - What backs this group. `null` for a standalone group.
	Implementation *IntegrationsInstanceGroupsListOutputItemsImplementation `json:"implementation,omitempty"`
	// Providers - Providers linked into this group.
	Providers []IntegrationsInstanceGroupsListOutputItemsProviders `json:"providers"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsListOutputItemsAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsListOutputPagination represents the integrations instance groups list output pagination type.
type IntegrationsInstanceGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsInstanceGroupsListOutput represents the integrations instance groups list output type.
type IntegrationsInstanceGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                         `json:"object"`
	Items      []IntegrationsInstanceGroupsListOutputItems    `json:"items"`
	Pagination IntegrationsInstanceGroupsListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstanceGroupsListOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsListOutput.
func MapIntegrationsInstanceGroupsListOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsListOutput, error) {
	var v IntegrationsInstanceGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsListOutputToJSON serializes a IntegrationsInstanceGroupsListOutput to JSON.
func MapIntegrationsInstanceGroupsListOutputToJSON(v *IntegrationsInstanceGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsListQueryCreatedAt - Only return results whose integration instance group creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstanceGroupsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsListQueryUpdatedAt - Only return results whose integration instance group last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstanceGroupsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsListQuery represents the integrations instance groups list query type.
type IntegrationsInstanceGroupsListQuery struct {
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
	// Status - Only return integration instance groups with one of these statuses. `deleted` groups are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instance groups with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instance groups for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return integration instance groups containing one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return integration instance groups containing one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderId - Only return integration instance groups using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instance groups using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration instance groups running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instance groups using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instance groups using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instance groups using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance group creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsInstanceGroupsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance group last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsInstanceGroupsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstanceGroupsListQueryFromJSON deserializes JSON data into a IntegrationsInstanceGroupsListQuery.
func MapIntegrationsInstanceGroupsListQueryFromJSON(data []byte) (*IntegrationsInstanceGroupsListQuery, error) {
	var v IntegrationsInstanceGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsListQueryToJSON serializes a IntegrationsInstanceGroupsListQuery to JSON.
func MapIntegrationsInstanceGroupsListQueryToJSON(v *IntegrationsInstanceGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
