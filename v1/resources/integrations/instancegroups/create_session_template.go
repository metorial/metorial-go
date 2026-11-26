package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsCreateSessionTemplateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment represents the integrations instance groups create session template output providers deployment type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig represents the integrations instance groups create session template output providers config type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig - Auth config the provider runs with. `null` when there is none.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders represents the integrations instance groups create session template output providers type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                                                   `json:"session_template_id"`
	Deployment        IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersDeployment `json:"deployment"`
	Config            IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *IntegrationsInstanceGroupsCreateSessionTemplateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionTemplateOutput represents the integrations instance groups create session template output type.
type IntegrationsInstanceGroupsCreateSessionTemplateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template is deleted. Archived templates stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationInstanceId - Integration instance linked to this template. `null` when there is none.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceGroupId - Integration instance group linked to this template. `null` when there is none.
	IntegrationInstanceGroupId *string `json:"integration_instance_group_id,omitempty"`
	// IdentityActorId - Identity actor linked to this template. `null` when there is none.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstanceGroupsCreateSessionTemplateOutputAccess `json:"access"`
	// IdentityId - Identity linked to this template. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// Providers - Providers that sessions created from this template connect to.
	Providers []IntegrationsInstanceGroupsCreateSessionTemplateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionTemplateOutput.
func MapIntegrationsInstanceGroupsCreateSessionTemplateOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionTemplateOutput, error) {
	var v IntegrationsInstanceGroupsCreateSessionTemplateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateOutputToJSON serializes a IntegrationsInstanceGroupsCreateSessionTemplateOutput to JSON.
func MapIntegrationsInstanceGroupsCreateSessionTemplateOutputToJSON(v *IntegrationsInstanceGroupsCreateSessionTemplateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsCreateSessionTemplateBody represents the integrations instance groups create session template body type.
type IntegrationsInstanceGroupsCreateSessionTemplateBody struct {
	// Name - Display name for the session template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionTemplateBody.
func MapIntegrationsInstanceGroupsCreateSessionTemplateBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionTemplateBody, error) {
	var v IntegrationsInstanceGroupsCreateSessionTemplateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionTemplateBodyToJSON serializes a IntegrationsInstanceGroupsCreateSessionTemplateBody to JSON.
func MapIntegrationsInstanceGroupsCreateSessionTemplateBodyToJSON(v *IntegrationsInstanceGroupsCreateSessionTemplateBody) ([]byte, error) {
	return json.Marshal(v)
}
