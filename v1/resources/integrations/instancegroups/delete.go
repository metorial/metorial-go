package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsDeleteOutputImplementation - What backs this group. `null` for a standalone group.
type IntegrationsInstanceGroupsDeleteOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpEndpointId - Magic MCP endpoint backing this group.
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
	ToolFilter *IntegrationsInstanceGroupsDeleteOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsDeleteOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsDeleteOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsDeleteOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsDeleteOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsDeleteOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsDeleteOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsDeleteOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsDeleteOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsDeleteOutput represents the integrations instance groups delete output type.
type IntegrationsInstanceGroupsDeleteOutput struct {
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
	Implementation *IntegrationsInstanceGroupsDeleteOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers linked into this group.
	Providers []IntegrationsInstanceGroupsDeleteOutputProviders `json:"providers"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsDeleteOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
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
