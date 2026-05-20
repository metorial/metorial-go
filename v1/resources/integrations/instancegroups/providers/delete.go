package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsProvidersDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsProvidersDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsProvidersDeleteOutput represents the integrations instance groups providers delete output type.
type IntegrationsInstanceGroupsProvidersDeleteOutput struct {
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
	ToolFilter           *IntegrationsInstanceGroupsProvidersDeleteOutputToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsProvidersDeleteOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsProvidersDeleteOutput.
func MapIntegrationsInstanceGroupsProvidersDeleteOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsProvidersDeleteOutput, error) {
	var v IntegrationsInstanceGroupsProvidersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsProvidersDeleteOutputToJSON serializes a IntegrationsInstanceGroupsProvidersDeleteOutput to JSON.
func MapIntegrationsInstanceGroupsProvidersDeleteOutputToJSON(v *IntegrationsInstanceGroupsProvidersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
