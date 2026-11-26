package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListOutputItems represents the integrations instance groups providers list output items type.
type IntegrationsInstanceGroupsProvidersListOutputItems struct {
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
	ToolFilter *IntegrationsInstanceGroupsProvidersListOutputItemsToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListOutputPagination represents the integrations instance groups providers list output pagination type.
type IntegrationsInstanceGroupsProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListOutput represents the integrations instance groups providers list output type.
type IntegrationsInstanceGroupsProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                                  `json:"object"`
	Items      []IntegrationsInstanceGroupsProvidersListOutputItems    `json:"items"`
	Pagination IntegrationsInstanceGroupsProvidersListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstanceGroupsProvidersListOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsProvidersListOutput.
func MapIntegrationsInstanceGroupsProvidersListOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsProvidersListOutput, error) {
	var v IntegrationsInstanceGroupsProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsProvidersListOutputToJSON serializes a IntegrationsInstanceGroupsProvidersListOutput to JSON.
func MapIntegrationsInstanceGroupsProvidersListOutputToJSON(v *IntegrationsInstanceGroupsProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsProvidersListQueryCreatedAt - Only return results whose integration instance group provider creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstanceGroupsProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListQueryUpdatedAt - Only return results whose integration instance group provider last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstanceGroupsProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListQuery represents the integrations instance groups providers list query type.
type IntegrationsInstanceGroupsProvidersListQuery struct {
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
	// Status - Only return group providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return group providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationInstanceGroupId - Only return group providers for one of these groups.
	IntegrationInstanceGroupId *any `json:"integration_instance_group_id,omitempty"`
	// IntegrationId - Only return group providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return group providers for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return group providers linking one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// ProviderId - Only return group providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return group providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return group providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return group providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return group providers using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return group providers using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance group provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsInstanceGroupsProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance group provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsInstanceGroupsProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstanceGroupsProvidersListQueryFromJSON deserializes JSON data into a IntegrationsInstanceGroupsProvidersListQuery.
func MapIntegrationsInstanceGroupsProvidersListQueryFromJSON(data []byte) (*IntegrationsInstanceGroupsProvidersListQuery, error) {
	var v IntegrationsInstanceGroupsProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsProvidersListQueryToJSON serializes a IntegrationsInstanceGroupsProvidersListQuery to JSON.
func MapIntegrationsInstanceGroupsProvidersListQueryToJSON(v *IntegrationsInstanceGroupsProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
