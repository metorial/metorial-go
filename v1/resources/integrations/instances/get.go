package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesGetOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstancesGetOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstancesGetOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstancesGetOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstancesGetOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstancesGetOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstancesGetOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstancesGetOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstancesGetOutputAccess - Resource ownership and access settings.
type IntegrationsInstancesGetOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstancesGetOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstancesGetOutputImplementation represents the integrations instances get output implementation type.
type IntegrationsInstancesGetOutputImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesGetOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesGetOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesGetOutputProvidersProvider represents the integrations instances get output providers provider type.
type IntegrationsInstancesGetOutputProvidersProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesGetOutputProvidersIntegrationProviderProviderVersion represents the integrations instances get output providers integration provider provider version type.
type IntegrationsInstancesGetOutputProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsInstancesGetOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesGetOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesGetOutputProvidersIntegrationProviderConfig represents the integrations instances get output providers integration provider config type.
type IntegrationsInstancesGetOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesGetOutputProvidersIntegrationProvider represents the integrations instances get output providers integration provider type.
type IntegrationsInstancesGetOutputProvidersIntegrationProvider struct {
	Object          string                                                                    `json:"object"`
	Id              string                                                                    `json:"id"`
	ProviderVersion IntegrationsInstancesGetOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                    `json:"status"`
	Name            string                                                                    `json:"name"`
	Description     *string                                                                   `json:"description,omitempty"`
	Metadata        *map[string]any                                                           `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsInstancesGetOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                `json:"provider_id"`
	DeploymentId      string                                                                `json:"deployment_id"`
	AuthMethodId      *string                                                               `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                               `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesGetOutputProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                             `json:"created_at"`
	UpdatedAt         time.Time                                                             `json:"updated_at"`
	ArchivedAt        *time.Time                                                            `json:"archived_at,omitempty"`
}

// IntegrationsInstancesGetOutputProvidersConfig represents the integrations instances get output providers config type.
type IntegrationsInstancesGetOutputProvidersConfig struct {
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

// IntegrationsInstancesGetOutputProvidersAuthConfig represents the integrations instances get output providers auth config type.
type IntegrationsInstancesGetOutputProvidersAuthConfig struct {
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

// IntegrationsInstancesGetOutputProviders represents the integrations instances get output providers type.
type IntegrationsInstancesGetOutputProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsInstancesGetOutputProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                       `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesGetOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesGetOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesGetOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesGetOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string    `json:"callback_instance_id,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	ArchivedAt         *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesGetOutput represents the integrations instances get output type.
type IntegrationsInstancesGetOutput struct {
	Object          string          `json:"object"`
	Id              string          `json:"id"`
	Status          string          `json:"status"`
	Name            string          `json:"name"`
	Description     *string         `json:"description,omitempty"`
	Metadata        *map[string]any `json:"metadata,omitempty"`
	IntegrationId   string          `json:"integration_id"`
	IdentityActorId *string         `json:"identity_actor_id,omitempty"`
	IdentityId      *string         `json:"identity_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access         IntegrationsInstancesGetOutputAccess          `json:"access"`
	Implementation *IntegrationsInstancesGetOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsInstancesGetOutputProviders     `json:"providers"`
	CreatedAt      time.Time                                     `json:"created_at"`
	UpdatedAt      time.Time                                     `json:"updated_at"`
	ArchivedAt     *time.Time                                    `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesGetOutputFromJSON deserializes JSON data into a IntegrationsInstancesGetOutput.
func MapIntegrationsInstancesGetOutputFromJSON(data []byte) (*IntegrationsInstancesGetOutput, error) {
	var v IntegrationsInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesGetOutputToJSON serializes a IntegrationsInstancesGetOutput to JSON.
func MapIntegrationsInstancesGetOutputToJSON(v *IntegrationsInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
