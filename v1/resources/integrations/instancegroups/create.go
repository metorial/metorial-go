package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateOutputImplementation - What backs this group. `null` for a standalone group.
type IntegrationsInstanceGroupsCreateOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpEndpointId - Magic MCP endpoint backing this group.
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
	ToolFilter *IntegrationsInstanceGroupsCreateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsCreateOutput represents the integrations instance groups create output type.
type IntegrationsInstanceGroupsCreateOutput struct {
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
	Implementation *IntegrationsInstanceGroupsCreateOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers linked into this group.
	Providers []IntegrationsInstanceGroupsCreateOutputProviders `json:"providers"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsCreateOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
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
	// IntegrationInstanceProviderId - Integration instance provider to link into the group.
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateBody represents the integrations instance groups create body type.
type IntegrationsInstanceGroupsCreateBody struct {
	// Name - Display name for the integration instance group.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance group.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Providers - Providers to link into the group.
	Providers *[]IntegrationsInstanceGroupsCreateBodyProviders `json:"providers,omitempty"`
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
