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
	Object                        string          `json:"object"`
	Id                            string          `json:"id"`
	Status                        string          `json:"status"`
	Name                          string          `json:"name"`
	Description                   *string         `json:"description,omitempty"`
	Metadata                      *map[string]any `json:"metadata,omitempty"`
	IntegrationId                 string          `json:"integration_id"`
	IntegrationInstanceGroupId    string          `json:"integration_instance_group_id"`
	IntegrationInstanceId         string          `json:"integration_instance_id"`
	IntegrationProviderId         *string         `json:"integration_provider_id,omitempty"`
	IntegrationInstanceProviderId string          `json:"integration_instance_provider_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstanceGroupsProvidersListOutputItemsToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                     `json:"created_at"`
	UpdatedAt            time.Time                                                     `json:"updated_at"`
	ArchivedAt           *time.Time                                                    `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListOutputPagination represents the integrations instance groups providers list output pagination type.
type IntegrationsInstanceGroupsProvidersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsInstanceGroupsProvidersListOutput represents the integrations instance groups providers list output type.
type IntegrationsInstanceGroupsProvidersListOutput struct {
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

// IntegrationsInstanceGroupsProvidersListQueryCreatedAt - Filter integration instance group provider creation time by date range
type IntegrationsInstanceGroupsProvidersListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration instance group provider creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance group provider creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListQueryUpdatedAt - Filter integration instance group provider last update time by date range
type IntegrationsInstanceGroupsProvidersListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration instance group provider last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance group provider last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsProvidersListQuery represents the integrations instance groups providers list query type.
type IntegrationsInstanceGroupsProvidersListQuery struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	IntegrationInstanceGroupId    *any     `json:"integration_instance_group_id,omitempty"`
	IntegrationId                 *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId         *any     `json:"integration_instance_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId             *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance group provider creation time by date range
	CreatedAt *IntegrationsInstanceGroupsProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance group provider last update time by date range
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
