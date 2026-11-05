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
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig represents the integrations instances create session template output providers config type.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig represents the integrations instances create session template output providers auth config type.
type IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// IntegrationsInstancesCreateSessionTemplateOutputProviders represents the integrations instances create session template output providers type.
type IntegrationsInstancesCreateSessionTemplateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session template provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstancesCreateSessionTemplateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Parent session template ID
	SessionTemplateId string                                                               `json:"session_template_id"`
	Deployment        IntegrationsInstancesCreateSessionTemplateOutputProvidersDeployment  `json:"deployment"`
	Config            IntegrationsInstancesCreateSessionTemplateOutputProvidersConfig      `json:"config"`
	AuthConfig        *IntegrationsInstancesCreateSessionTemplateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesCreateSessionTemplateOutput represents the integrations instances create session template output type.
type IntegrationsInstancesCreateSessionTemplateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session template identifier
	Id string `json:"id"`
	// Status - Status of the session template
	Status string `json:"status"`
	// Name - Template name
	Name string `json:"name"`
	// Description - Template description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs
	Metadata                   *map[string]any `json:"metadata,omitempty"`
	IntegrationInstanceId      *string         `json:"integration_instance_id,omitempty"`
	IntegrationInstanceGroupId *string         `json:"integration_instance_group_id,omitempty"`
	IdentityActorId            *string         `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access     IntegrationsInstancesCreateSessionTemplateOutputAccess `json:"access"`
	IdentityId *string                                                `json:"identity_id,omitempty"`
	// Providers - Template providers
	Providers []IntegrationsInstancesCreateSessionTemplateOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
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
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
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
