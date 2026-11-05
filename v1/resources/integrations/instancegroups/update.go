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
	Object         string                                                `json:"object"`
	Id             string                                                `json:"id"`
	Status         string                                                `json:"status"`
	Name           string                                                `json:"name"`
	Description    *string                                               `json:"description,omitempty"`
	Metadata       *map[string]any                                       `json:"metadata,omitempty"`
	Implementation *IntegrationsInstanceGroupsUpdateOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstanceGroupsUpdateOutputProviders     `json:"providers"`
	// Access - Resource ownership and access settings.
	Access     IntegrationsInstanceGroupsUpdateOutputAccess `json:"access"`
	CreatedAt  time.Time                                    `json:"created_at"`
	UpdatedAt  time.Time                                    `json:"updated_at"`
	ArchivedAt *time.Time                                   `json:"archived_at,omitempty"`
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
