package sessions

import (
	"encoding/json"
	"time"
)

// SessionsCreateOutputUsage represents the sessions create output usage type.
type SessionsCreateOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsCreateOutputProvidersUsage - Usage statistics
type SessionsCreateOutputProvidersUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsCreateOutputProvidersDeployment represents the sessions create output providers deployment type.
type SessionsCreateOutputProvidersDeployment struct {
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

// SessionsCreateOutputProvidersConfig represents the sessions create output providers config type.
type SessionsCreateOutputProvidersConfig struct {
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

// SessionsCreateOutputProvidersAuthConfig represents the sessions create output providers auth config type.
type SessionsCreateOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// SessionsCreateOutputProviders represents the sessions create output providers type.
type SessionsCreateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsCreateOutputProvidersUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsCreateOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionId - Parent session ID
	SessionId string `json:"session_id"`
	// FromTemplateId - Source template ID
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Source template provider ID
	FromTemplateProviderId *string                                  `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsCreateOutputProvidersDeployment  `json:"deployment"`
	Config                 SessionsCreateOutputProvidersConfig      `json:"config"`
	AuthConfig             *SessionsCreateOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type SessionsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// SessionsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type SessionsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// SessionsCreateOutputAccessOwner - The actor that owns this resource, if any.
type SessionsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *SessionsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *SessionsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// SessionsCreateOutputAccess - Resource ownership and access settings.
type SessionsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *SessionsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// SessionsCreateOutput represents the sessions create output type.
type SessionsCreateOutput struct {
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
	Usage        SessionsCreateOutputUsage `json:"usage"`
	// Providers - Session providers
	Providers []SessionsCreateOutputProviders `json:"providers"`
	// FromTemplateIds - Template IDs this session was created from
	FromTemplateIds []string `json:"from_template_ids"`
	// HasErrors - Whether the session has any errors
	HasErrors bool `json:"has_errors"`
	// HasWarnings - Whether the session has any warnings
	HasWarnings     bool    `json:"has_warnings"`
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access     SessionsCreateOutputAccess `json:"access"`
	IdentityId *string                    `json:"identity_id,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsCreateOutputFromJSON deserializes JSON data into a SessionsCreateOutput.
func MapSessionsCreateOutputFromJSON(data []byte) (*SessionsCreateOutput, error) {
	var v SessionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsCreateOutputToJSON serializes a SessionsCreateOutput to JSON.
func MapSessionsCreateOutputToJSON(v *SessionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsCreateBody represents the sessions create body type.
type SessionsCreateBody struct {
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Metadata    *map[string]any  `json:"metadata,omitempty"`
	Providers   []map[string]any `json:"providers"`
}

// MapSessionsCreateBodyFromJSON deserializes JSON data into a SessionsCreateBody.
func MapSessionsCreateBodyFromJSON(data []byte) (*SessionsCreateBody, error) {
	var v SessionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsCreateBodyToJSON serializes a SessionsCreateBody to JSON.
func MapSessionsCreateBodyToJSON(v *SessionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
