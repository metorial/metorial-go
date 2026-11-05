package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateSessionOutputUsage represents the integrations instance groups create session output usage type.
type IntegrationsInstanceGroupsCreateSessionOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage - Usage statistics
type IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment represents the integrations instance groups create session output providers deployment type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment struct {
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

// IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig represents the integrations instance groups create session output providers config type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig struct {
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

// IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig represents the integrations instance groups create session output providers auth config type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProviders represents the integrations instance groups create session output providers type.
type IntegrationsInstanceGroupsCreateSessionOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionId - Parent session ID
	SessionId string `json:"session_id"`
	// FromTemplateId - Source template ID
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Source template provider ID
	FromTemplateProviderId *string                                                           `json:"from_template_provider_id,omitempty"`
	Deployment             IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment  `json:"deployment"`
	Config                 IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig      `json:"config"`
	AuthConfig             *IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstanceGroupsCreateSessionOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstanceGroupsCreateSessionOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionOutputAccess - Resource ownership and access settings.
type IntegrationsInstanceGroupsCreateSessionOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstanceGroupsCreateSessionOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionOutput represents the integrations instance groups create session output type.
type IntegrationsInstanceGroupsCreateSessionOutput struct {
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
	ClientSecret *string                                            `json:"client_secret,omitempty"`
	Usage        IntegrationsInstanceGroupsCreateSessionOutputUsage `json:"usage"`
	// Providers - Session providers
	Providers []IntegrationsInstanceGroupsCreateSessionOutputProviders `json:"providers"`
	// FromTemplateIds - Template IDs this session was created from
	FromTemplateIds []string `json:"from_template_ids"`
	// HasErrors - Whether the session has any errors
	HasErrors bool `json:"has_errors"`
	// HasWarnings - Whether the session has any warnings
	HasWarnings     bool    `json:"has_warnings"`
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access     IntegrationsInstanceGroupsCreateSessionOutputAccess `json:"access"`
	IdentityId *string                                             `json:"identity_id,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIntegrationsInstanceGroupsCreateSessionOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionOutput.
func MapIntegrationsInstanceGroupsCreateSessionOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionOutput, error) {
	var v IntegrationsInstanceGroupsCreateSessionOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionOutputToJSON serializes a IntegrationsInstanceGroupsCreateSessionOutput to JSON.
func MapIntegrationsInstanceGroupsCreateSessionOutputToJSON(v *IntegrationsInstanceGroupsCreateSessionOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsCreateSessionBody represents the integrations instance groups create session body type.
type IntegrationsInstanceGroupsCreateSessionBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateSessionBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionBody.
func MapIntegrationsInstanceGroupsCreateSessionBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionBody, error) {
	var v IntegrationsInstanceGroupsCreateSessionBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionBodyToJSON serializes a IntegrationsInstanceGroupsCreateSessionBody to JSON.
func MapIntegrationsInstanceGroupsCreateSessionBodyToJSON(v *IntegrationsInstanceGroupsCreateSessionBody) ([]byte, error) {
	return json.Marshal(v)
}
