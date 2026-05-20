package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsUpdateOutputImplementation represents the integrations instance groups update output implementation type.
type IntegrationsInstanceGroupsUpdateOutputImplementation struct {
	Type               string `json:"type"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// IntegrationsInstanceGroupsUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutputProviders represents the integrations instance groups update output providers type.
type IntegrationsInstanceGroupsUpdateOutputProviders struct {
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
	ToolFilter           *IntegrationsInstanceGroupsUpdateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutput represents the integrations instance groups update output type.
type IntegrationsInstanceGroupsUpdateOutput struct {
	Object         string                                                `json:"object"`
	Id             string                                                `json:"id"`
	Status         string                                                `json:"status"`
	Name           string                                                `json:"name"`
	Description    *string                                               `json:"description,omitempty"`
	Metadata       *map[string]any                                       `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsUpdateOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsUpdateOutputProviders     `json:"providers"`
	CreatedAt      time.Time                                             `json:"created_at"`
	UpdatedAt      time.Time                                             `json:"updated_at"`
	ArchivedAt     *time.Time                                            `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsUpdateOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsUpdateOutput.
func MapIntegrationsInstanceGroupsUpdateOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsUpdateOutput, error) {
	var v IntegrationsInstanceGroupsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsUpdateOutputToJSON serializes a IntegrationsInstanceGroupsUpdateOutput to JSON.
func MapIntegrationsInstanceGroupsUpdateOutputToJSON(v *IntegrationsInstanceGroupsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsUpdateBodyProviders represents the integrations instance groups update body providers type.
type IntegrationsInstanceGroupsUpdateBodyProviders struct {
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	ToolFilters                   *any   `json:"tool_filters,omitempty"`
}

// IntegrationsInstanceGroupsUpdateBody represents the integrations instance groups update body type.
type IntegrationsInstanceGroupsUpdateBody struct {
	Name        *string                                          `json:"name,omitempty"`
	Description *string                                          `json:"description,omitempty"`
	Metadata    *map[string]any                                  `json:"metadata,omitempty"`
	Providers   *[]IntegrationsInstanceGroupsUpdateBodyProviders `json:"providers,omitempty"`
}

// MapIntegrationsInstanceGroupsUpdateBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsUpdateBody.
func MapIntegrationsInstanceGroupsUpdateBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsUpdateBody, error) {
	var v IntegrationsInstanceGroupsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsUpdateBodyToJSON serializes a IntegrationsInstanceGroupsUpdateBody to JSON.
func MapIntegrationsInstanceGroupsUpdateBodyToJSON(v *IntegrationsInstanceGroupsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
