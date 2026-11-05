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
	Object         string                                             `json:"object"`
	Id             string                                             `json:"id"`
	Status         string                                             `json:"status"`
	Name           string                                             `json:"name"`
	Description    *string                                            `json:"description,omitempty"`
	Metadata       *map[string]any                                    `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsGetOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsGetOutputProviders     `json:"providers"`
	// Access - Resource ownership and access settings.
	Access     IntegrationsInstanceGroupsGetOutputAccess `json:"access"`
	CreatedAt  time.Time                                 `json:"created_at"`
	UpdatedAt  time.Time                                 `json:"updated_at"`
	ArchivedAt *time.Time                                `json:"archived_at,omitempty"`
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
