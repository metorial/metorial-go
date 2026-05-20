package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateOutputImplementation represents the integrations instance groups create output implementation type.
type IntegrationsInstanceGroupsCreateOutputImplementation struct {
	Type               string `json:"type"`
	MagicMcpEndpointId string `json:"magic_mcp_endpoint_id"`
}

// IntegrationsInstanceGroupsCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutputProviders represents the integrations instance groups create output providers type.
type IntegrationsInstanceGroupsCreateOutputProviders struct {
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
	ToolFilter           *IntegrationsInstanceGroupsCreateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	CreatedAt            time.Time                                                  `json:"created_at"`
	UpdatedAt            time.Time                                                  `json:"updated_at"`
	ArchivedAt           *time.Time                                                 `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutput represents the integrations instance groups create output type.
type IntegrationsInstanceGroupsCreateOutput struct {
	Object         string                                                `json:"object"`
	Id             string                                                `json:"id"`
	Status         string                                                `json:"status"`
	Name           string                                                `json:"name"`
	Description    *string                                               `json:"description,omitempty"`
	Metadata       *map[string]any                                       `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsCreateOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsCreateOutputProviders     `json:"providers"`
	CreatedAt      time.Time                                             `json:"created_at"`
	UpdatedAt      time.Time                                             `json:"updated_at"`
	ArchivedAt     *time.Time                                            `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateOutput.
func MapIntegrationsInstanceGroupsCreateOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateOutput, error) {
	var v IntegrationsInstanceGroupsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateOutputToJSON serializes a IntegrationsInstanceGroupsCreateOutput to JSON.
func MapIntegrationsInstanceGroupsCreateOutputToJSON(v *IntegrationsInstanceGroupsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsCreateBodyProviders represents the integrations instance groups create body providers type.
type IntegrationsInstanceGroupsCreateBodyProviders struct {
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	ToolFilters                   *any   `json:"tool_filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateBody represents the integrations instance groups create body type.
type IntegrationsInstanceGroupsCreateBody struct {
	Name        string                                           `json:"name"`
	Description *string                                          `json:"description,omitempty"`
	Metadata    *map[string]any                                  `json:"metadata,omitempty"`
	Providers   *[]IntegrationsInstanceGroupsCreateBodyProviders `json:"providers,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateBody.
func MapIntegrationsInstanceGroupsCreateBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateBody, error) {
	var v IntegrationsInstanceGroupsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateBodyToJSON serializes a IntegrationsInstanceGroupsCreateBody to JSON.
func MapIntegrationsInstanceGroupsCreateBodyToJSON(v *IntegrationsInstanceGroupsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
