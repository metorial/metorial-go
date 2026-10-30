package sessions

import (
	"encoding/json"
	"time"
)

// SessionsUpdateOutputUsage represents the sessions update output usage type.
type SessionsUpdateOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsUpdateOutputProvidersUsage - Usage statistics
type SessionsUpdateOutputProvidersUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsUpdateOutputProvidersDeployment represents the sessions update output providers deployment type.
type SessionsUpdateOutputProvidersDeployment struct {
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

// SessionsUpdateOutputProvidersConfig represents the sessions update output providers config type.
type SessionsUpdateOutputProvidersConfig struct {
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

// SessionsUpdateOutputProvidersAuthConfig represents the sessions update output providers auth config type.
type SessionsUpdateOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// SessionsUpdateOutputProviders represents the sessions update output providers type.
type SessionsUpdateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsUpdateOutputProvidersUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsUpdateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionId - Parent session ID
	SessionId string `json:"session_id"`
	// FromTemplateId - Source template ID
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Source template provider ID
	FromTemplateProviderId *string                                  `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsUpdateOutputProvidersDeployment  `json:"deployment"`
	Config                 SessionsUpdateOutputProvidersConfig      `json:"config"`
	AuthConfig             *SessionsUpdateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionsUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionsUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionsUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionsUpdateOutputAccessOwner - The actor that owns this resource, if any.
type SessionsUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionsUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionsUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionsUpdateOutputAccess - Resource ownership and access settings.
type SessionsUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionsUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// SessionsUpdateOutput represents the sessions update output type.
type SessionsUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session identifier
	Id string `json:"id"`
	// Status - Session status
	Status string `json:"status"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ConnectionState - Session connection state
	ConnectionState string `json:"connection_state"`
	// ConnectionUrl - MCP connection URL for this session
	ConnectionUrl string `json:"connection_url"`
	// ClientSecret - Session-scoped fine grained client secret token
	ClientSecret *string                   `json:"client_secret,omitempty"`
	Usage        SessionsUpdateOutputUsage `json:"usage"`
	// Providers - Session providers
	Providers []SessionsUpdateOutputProviders `json:"providers"`
	// FromTemplateIds - Template IDs this session was created from
	FromTemplateIds []string `json:"from_template_ids"`
	// HasErrors - Whether the session has any errors
	HasErrors bool `json:"has_errors"`
	// HasWarnings - Whether the session has any warnings
	HasWarnings     bool    `json:"has_warnings"`
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access     SessionsUpdateOutputAccess `json:"access"`
	IdentityId *string                    `json:"identity_id,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsUpdateOutputFromJSON deserializes JSON data into a SessionsUpdateOutput.
func MapSessionsUpdateOutputFromJSON(data []byte) (*SessionsUpdateOutput, error) {
	var v SessionsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsUpdateOutputToJSON serializes a SessionsUpdateOutput to JSON.
func MapSessionsUpdateOutputToJSON(v *SessionsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsUpdateBody represents the sessions update body type.
type SessionsUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapSessionsUpdateBodyFromJSON deserializes JSON data into a SessionsUpdateBody.
func MapSessionsUpdateBodyFromJSON(data []byte) (*SessionsUpdateBody, error) {
	var v SessionsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsUpdateBodyToJSON serializes a SessionsUpdateBody to JSON.
func MapSessionsUpdateBodyToJSON(v *SessionsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
