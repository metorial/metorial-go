package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsProvidersGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsProvidersGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsProvidersGetOutput represents the integrations instance groups providers get output type.
type IntegrationsInstanceGroupsProvidersGetOutput struct {
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
	ToolFilter *IntegrationsInstanceGroupsProvidersGetOutputToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsInstanceGroupsProvidersGetOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsProvidersGetOutput.
func MapIntegrationsInstanceGroupsProvidersGetOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsProvidersGetOutput, error) {
	var v IntegrationsInstanceGroupsProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsProvidersGetOutputToJSON serializes a IntegrationsInstanceGroupsProvidersGetOutput to JSON.
func MapIntegrationsInstanceGroupsProvidersGetOutputToJSON(v *IntegrationsInstanceGroupsProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
