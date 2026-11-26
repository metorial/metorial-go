package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsGetOutputImplementation - What backs this group. `null` for a standalone group.
type IntegrationsInstanceGroupsGetOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpEndpointId - Magic MCP endpoint backing this group.
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
	ToolFilter *IntegrationsInstanceGroupsGetOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsGetOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsGetOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsGetOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsGetOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsGetOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsGetOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsGetOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsGetOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsGetOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsGetOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsGetOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsGetOutput represents the integrations instance groups get output type.
type IntegrationsInstanceGroupsGetOutput struct {
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
	Implementation *IntegrationsInstanceGroupsGetOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers linked into this group.
	Providers []IntegrationsInstanceGroupsGetOutputProviders `json:"providers"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsGetOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
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
