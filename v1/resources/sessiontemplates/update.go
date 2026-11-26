package sessiontemplates

import (
	"encoding/json"
	"time"
)

// SessionTemplatesUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionTemplatesUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionTemplatesUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionTemplatesUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionTemplatesUpdateOutputAccessOwner - The actor that owns this resource, if any.
type SessionTemplatesUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionTemplatesUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionTemplatesUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionTemplatesUpdateOutputAccess - Resource ownership and access settings.
type SessionTemplatesUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionTemplatesUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// SessionTemplatesUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesUpdateOutputProvidersDeployment represents the session templates update output providers deployment type.
type SessionTemplatesUpdateOutputProvidersDeployment struct {
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

// SessionTemplatesUpdateOutputProvidersConfig represents the session templates update output providers config type.
type SessionTemplatesUpdateOutputProvidersConfig struct {
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

// SessionTemplatesUpdateOutputProvidersAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesUpdateOutputProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesUpdateOutputProviders represents the session templates update output providers type.
type SessionTemplatesUpdateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesUpdateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                          `json:"session_template_id"`
	Deployment        SessionTemplatesUpdateOutputProvidersDeployment `json:"deployment"`
	Config            SessionTemplatesUpdateOutputProvidersConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesUpdateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesUpdateOutput represents the session templates update output type.
type SessionTemplatesUpdateOutput struct {
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
	Access SessionTemplatesUpdateOutputAccess `json:"access"`
	// IdentityId - Identity linked to this template. `null` when there is none.
	IdentityId *string `json:"identity_id,omitempty"`
	// Providers - Providers that sessions created from this template connect to.
	Providers []SessionTemplatesUpdateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionTemplatesUpdateOutputFromJSON deserializes JSON data into a SessionTemplatesUpdateOutput.
func MapSessionTemplatesUpdateOutputFromJSON(data []byte) (*SessionTemplatesUpdateOutput, error) {
	var v SessionTemplatesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesUpdateOutputToJSON serializes a SessionTemplatesUpdateOutput to JSON.
func MapSessionTemplatesUpdateOutputToJSON(v *SessionTemplatesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionTemplatesUpdateBody represents the session templates update body type.
type SessionTemplatesUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the session template. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapSessionTemplatesUpdateBodyFromJSON deserializes JSON data into a SessionTemplatesUpdateBody.
func MapSessionTemplatesUpdateBodyFromJSON(data []byte) (*SessionTemplatesUpdateBody, error) {
	var v SessionTemplatesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesUpdateBodyToJSON serializes a SessionTemplatesUpdateBody to JSON.
func MapSessionTemplatesUpdateBodyToJSON(v *SessionTemplatesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
