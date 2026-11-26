package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstancesCreateSessionTemplateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstancesCreateSessionTemplateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstancesCreateSessionTemplateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstancesCreateSessionTemplateOutputAccess - Resource ownership and access settings.
type IntegrationsInstancesCreateSessionTemplateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstancesCreateSessionTemplateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProvidersDeployment represents the integrations instances create session template output providers deployment type.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersDeployment struct {
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

// IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig represents the integrations instances create session template output providers config type.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig struct {
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

// IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig - Auth config the provider runs with. `null` when there is none.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProviders represents the integrations instances create session template output providers type.
type IntegrationsInstancesCreateSessionTemplateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstancesCreateSessionTemplateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                                              `json:"session_template_id"`
	Deployment        IntegrationsInstancesCreateSessionTemplateOutputProvidersDeployment `json:"deployment"`
	Config            IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesCreateSessionTemplateOutput represents the integrations instances create session template output type.
type IntegrationsInstancesCreateSessionTemplateOutput struct {
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
	Access IntegrationsInstancesCreateSessionTemplateOutputAccess `json:"access"`
	// IdentityId - Identity linked to this template. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// Providers - Providers that sessions created from this template connect to.
	Providers []IntegrationsInstancesCreateSessionTemplateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIntegrationsInstancesCreateSessionTemplateOutputFromJSON deserializes JSON data into a IntegrationsInstancesCreateSessionTemplateOutput.
func MapIntegrationsInstancesCreateSessionTemplateOutputFromJSON(data []byte) (*IntegrationsInstancesCreateSessionTemplateOutput, error) {
	var v IntegrationsInstancesCreateSessionTemplateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesCreateSessionTemplateOutputToJSON serializes a IntegrationsInstancesCreateSessionTemplateOutput to JSON.
func MapIntegrationsInstancesCreateSessionTemplateOutputToJSON(v *IntegrationsInstancesCreateSessionTemplateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesCreateSessionTemplateBody represents the integrations instances create session template body type.
type IntegrationsInstancesCreateSessionTemplateBody struct {
	// Name - Display name for the session template.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapIntegrationsInstancesCreateSessionTemplateBodyFromJSON deserializes JSON data into a IntegrationsInstancesCreateSessionTemplateBody.
func MapIntegrationsInstancesCreateSessionTemplateBodyFromJSON(data []byte) (*IntegrationsInstancesCreateSessionTemplateBody, error) {
	var v IntegrationsInstancesCreateSessionTemplateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesCreateSessionTemplateBodyToJSON serializes a IntegrationsInstancesCreateSessionTemplateBody to JSON.
func MapIntegrationsInstancesCreateSessionTemplateBodyToJSON(v *IntegrationsInstancesCreateSessionTemplateBody) ([]byte, error) {
	return json.Marshal(v)
}
