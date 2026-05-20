package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsGetOutputImplementation represents the integrations instance groups get output implementation type.
type IntegrationsInstanceGroupsGetOutputImplementation struct {
	Type               string `json:"type"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// IntegrationsInstanceGroupsGetOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsGetOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsGetOutputProviders represents the integrations instance groups get output providers type.
type IntegrationsInstanceGroupsGetOutputProviders struct {
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
	ToolFilter           *IntegrationsInstanceGroupsGetOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                    `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                               `json:"created_at"`
	UpdatedAt            time.Time                                               `json:"updated_at"`
	ArchivedAt           *time.Time                                              `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsGetOutput represents the integrations instance groups get output type.
type IntegrationsInstanceGroupsGetOutput struct {
	Object         string                                             `json:"object"`
	Id             string                                             `json:"id"`
	Status         string                                             `json:"status"`
	Name           string                                             `json:"name"`
	Description    *string                                            `json:"description,omitempty"`
	Metadata       *map[string]any                                    `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsGetOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsGetOutputProviders     `json:"providers"`
	CreatedAt      time.Time                                          `json:"created_at"`
	UpdatedAt      time.Time                                          `json:"updated_at"`
	ArchivedAt     *time.Time                                         `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsGetOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsGetOutput.
func MapIntegrationsInstanceGroupsGetOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsGetOutput, error) {
	var v IntegrationsInstanceGroupsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsGetOutputToJSON serializes a IntegrationsInstanceGroupsGetOutput to JSON.
func MapIntegrationsInstanceGroupsGetOutputToJSON(v *IntegrationsInstanceGroupsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
