package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsListOutputItemsImplementation represents the integrations instance groups list output items implementation type.
type IntegrationsInstanceGroupsListOutputItemsImplementation struct {
	Type               string `json:"type"`
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
	ToolFilter           *IntegrationsInstanceGroupsListOutputItemsProvidersToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                     `json:"created_at"`
	UpdatedAt            time.Time                                                     `json:"updated_at"`
	ArchivedAt           *time.Time                                                    `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsListOutputItems represents the integrations instance groups list output items type.
type IntegrationsInstanceGroupsListOutputItems struct {
	Object         string                                                   `json:"object"`
	Id             string                                                   `json:"id"`
	Status         string                                                   `json:"status"`
	Name           string                                                   `json:"name"`
	Description    *string                                                  `json:"description,omitempty"`
	Metadata       *map[string]any                                          `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsListOutputItemsImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsListOutputItemsProviders     `json:"providers"`
	CreatedAt      time.Time                                                `json:"created_at"`
	UpdatedAt      time.Time                                                `json:"updated_at"`
	ArchivedAt     *time.Time                                               `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsListOutputPagination represents the integrations instance groups list output pagination type.
type IntegrationsInstanceGroupsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// IntegrationsInstanceGroupsListOutput represents the integrations instance groups list output type.
type IntegrationsInstanceGroupsListOutput struct {
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

// IntegrationsInstanceGroupsListQueryCreatedAt - Filter integration instance group creation time by date range
type IntegrationsInstanceGroupsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for integration instance group creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance group creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsListQueryUpdatedAt - Filter integration instance group last update time by date range
type IntegrationsInstanceGroupsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for integration instance group last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for integration instance group last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstanceGroupsListQuery represents the integrations instance groups list query type.
type IntegrationsInstanceGroupsListQuery struct {
	Limit                         *float64 `json:"limit,omitempty"`
	After                         *string  `json:"after,omitempty"`
	Before                        *string  `json:"before,omitempty"`
	Cursor                        *string  `json:"cursor,omitempty"`
	Order                         *string  `json:"order,omitempty"`
	Status                        *any     `json:"status,omitempty"`
	Id                            *any     `json:"id,omitempty"`
	IntegrationId                 *any     `json:"integration_id,omitempty"`
	IntegrationInstanceId         *any     `json:"integration_instance_id,omitempty"`
	IntegrationInstanceProviderId *any     `json:"integration_instance_provider_id,omitempty"`
	ProviderId                    *any     `json:"provider_id,omitempty"`
	IntegrationProviderId         *any     `json:"integration_provider_id,omitempty"`
	ProviderDeploymentId          *any     `json:"provider_deployment_id,omitempty"`
	ProviderConfigId              *any     `json:"provider_config_id,omitempty"`
	ProviderAuthConfigId          *any     `json:"provider_auth_config_id,omitempty"`
	SessionTemplateId             *any     `json:"session_template_id,omitempty"`
	// CreatedAt - Filter integration instance group creation time by date range
	CreatedAt *IntegrationsInstanceGroupsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter integration instance group last update time by date range
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
