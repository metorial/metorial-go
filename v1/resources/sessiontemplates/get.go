package sessiontemplates

import (
	"encoding/json"
	"time"
)

// SessionTemplatesGetOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionTemplatesGetOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionTemplatesGetOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionTemplatesGetOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionTemplatesGetOutputAccessOwner - The actor that owns this resource, if any.
type SessionTemplatesGetOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionTemplatesGetOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionTemplatesGetOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionTemplatesGetOutputAccess - Resource ownership and access settings.
type SessionTemplatesGetOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionTemplatesGetOutputAccessOwner `json:"owner,omitempty"`
}

// SessionTemplatesGetOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesGetOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesGetOutputProvidersDeployment represents the session templates get output providers deployment type.
type SessionTemplatesGetOutputProvidersDeployment struct {
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

// SessionTemplatesGetOutputProvidersConfig represents the session templates get output providers config type.
type SessionTemplatesGetOutputProvidersConfig struct {
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

// SessionTemplatesGetOutputProvidersAuthConfig represents the session templates get output providers auth config type.
type SessionTemplatesGetOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// SessionTemplatesGetOutputProviders represents the session templates get output providers type.
type SessionTemplatesGetOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session template provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesGetOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Parent session template ID
	SessionTemplateId string                                        `json:"session_template_id"`
	Deployment        SessionTemplatesGetOutputProvidersDeployment  `json:"deployment"`
	Config            SessionTemplatesGetOutputProvidersConfig      `json:"config"`
	AuthConfig        *SessionTemplatesGetOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesGetOutput represents the session templates get output type.
type SessionTemplatesGetOutput struct {
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
	Access     SessionTemplatesGetOutputAccess `json:"access"`
	IdentityId *string                         `json:"identity_id,omitempty"`
	// Providers - Template providers
	Providers []SessionTemplatesGetOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionTemplatesGetOutputFromJSON deserializes JSON data into a SessionTemplatesGetOutput.
func MapSessionTemplatesGetOutputFromJSON(data []byte) (*SessionTemplatesGetOutput, error) {
	var v SessionTemplatesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesGetOutputToJSON serializes a SessionTemplatesGetOutput to JSON.
func MapSessionTemplatesGetOutputToJSON(v *SessionTemplatesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
