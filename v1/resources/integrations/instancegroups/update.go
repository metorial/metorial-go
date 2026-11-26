package instancegroups

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// IntegrationsInstanceGroupsUpdateOutputImplementation - What backs this group. `null` for a standalone group.
type IntegrationsInstanceGroupsUpdateOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpEndpointId - Magic MCP endpoint backing this group.
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
	ToolFilter *IntegrationsInstanceGroupsUpdateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool `json:"is_override_tool_filter"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsUpdateOutput represents the integrations instance groups update output type.
type IntegrationsInstanceGroupsUpdateOutput struct {
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
	Implementation *IntegrationsInstanceGroupsUpdateOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers linked into this group.
	Providers []IntegrationsInstanceGroupsUpdateOutputProviders `json:"providers"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsUpdateOutputAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
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
	// IntegrationInstanceProviderId - Integration instance provider to link into the group.
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// IntegrationsInstanceGroupsUpdateBody represents the integrations instance groups update body type.
type IntegrationsInstanceGroupsUpdateBody struct {
	// Name - Display name for the integration instance group.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance group. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// Providers - Providers to link or update. Existing providers not listed are left alone.
	Providers *[]IntegrationsInstanceGroupsUpdateBodyProviders `json:"providers,omitempty"`
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
