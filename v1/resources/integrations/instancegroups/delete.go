package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsDeleteOutputImplementation represents the integrations instance groups delete output implementation type.
type IntegrationsInstanceGroupsDeleteOutputImplementation struct {
	Type               string `json:"type"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// IntegrationsInstanceGroupsDeleteOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsDeleteOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutputProviders represents the integrations instance groups delete output providers type.
type IntegrationsInstanceGroupsDeleteOutputProviders struct {
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
	ToolFilter           *IntegrationsInstanceGroupsDeleteOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutput represents the integrations instance groups delete output type.
type IntegrationsInstanceGroupsDeleteOutput struct {
	Object         string                                                `json:"object"`
	Id             string                                                `json:"id"`
	Status         string                                                `json:"status"`
	Name           string                                                `json:"name"`
	Description    *string                                               `json:"description,omitempty"`
	Metadata       *map[string]any                                       `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsDeleteOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsDeleteOutputProviders     `json:"providers"`
	CreatedAt      time.Time                                             `json:"created_at"`
	UpdatedAt      time.Time                                             `json:"updated_at"`
	ArchivedAt     *time.Time                                            `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsDeleteOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsDeleteOutput.
func MapIntegrationsInstanceGroupsDeleteOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsDeleteOutput, error) {
	var v IntegrationsInstanceGroupsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsDeleteOutputToJSON serializes a IntegrationsInstanceGroupsDeleteOutput to JSON.
func MapIntegrationsInstanceGroupsDeleteOutputToJSON(v *IntegrationsInstanceGroupsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
