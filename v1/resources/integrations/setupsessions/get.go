package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsGetOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsGetOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsGetOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsGetOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsGetOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsGetOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsGetOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsGetOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsGetOutputAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsGetOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsGetOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsGetOutputIntegrationInstanceAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation represents the integrations setup sessions get output integration instance implementation type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation struct {
	Type             string `json:"type"`
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider represents the integrations setup sessions get output integration instance providers provider type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion represents the integrations setup sessions get output integration instance providers integration provider provider version type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	Object string  `json:"object"`
	Id     string  `json:"id"`
	Index  float64 `json:"index"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions get output integration instance providers integration provider config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions get output integration instance providers integration provider type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider struct {
	Object          string                                                                                           `json:"object"`
	Id              string                                                                                           `json:"id"`
	ProviderVersion IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	Status          string                                                                                           `json:"status"`
	Name            string                                                                                           `json:"name"`
	Description     *string                                                                                          `json:"description,omitempty"`
	Metadata        *map[string]any                                                                                  `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                                                                       `json:"provider_id"`
	DeploymentId      string                                                                                       `json:"deployment_id"`
	AuthMethodId      *string                                                                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                                                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProviderConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                                                                    `json:"created_at"`
	UpdatedAt         time.Time                                                                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                                                                   `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig represents the integrations setup sessions get output integration instance providers config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions get output integration instance providers auth config type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig struct {
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

// IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders represents the integrations setup sessions get output integration instance providers type.
type IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders struct {
	Object                string          `json:"object"`
	Id                    string          `json:"id"`
	Status                string          `json:"status"`
	Name                  string          `json:"name"`
	Description           *string         `json:"description,omitempty"`
	Metadata              *map[string]any `json:"metadata,omitempty"`
	IntegrationId         string          `json:"integration_id"`
	IntegrationInstanceId string          `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter           *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersToolFilter         `json:"tool_filter,omitempty"`
	IsOverrideToolFilter bool                                                                              `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsGetOutputIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string    `json:"callback_instance_id,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	ArchivedAt         *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutputIntegrationInstance represents the integrations setup sessions get output integration instance type.
type IntegrationsSetupSessionsGetOutputIntegrationInstance struct {
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
	Access         IntegrationsSetupSessionsGetOutputIntegrationInstanceAccess          `json:"access"`
	Implementation *IntegrationsSetupSessionsGetOutputIntegrationInstanceImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsSetupSessionsGetOutputIntegrationInstanceProviders     `json:"providers"`
	CreatedAt      time.Time                                                            `json:"created_at"`
	UpdatedAt      time.Time                                                            `json:"updated_at"`
	ArchivedAt     *time.Time                                                           `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsGetOutput represents the integrations setup sessions get output type.
type IntegrationsSetupSessionsGetOutput struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	Url           string          `json:"url"`
	Name          *string         `json:"name,omitempty"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	Configuration *map[string]any `json:"configuration,omitempty"`
	RedirectUrl   *string         `json:"redirect_url,omitempty"`
	IntegrationId string          `json:"integration_id"`
	// Access - Resource ownership and access settings.
	Access              IntegrationsSetupSessionsGetOutputAccess              `json:"access"`
	IntegrationInstance IntegrationsSetupSessionsGetOutputIntegrationInstance `json:"integration_instance"`
	CreatedAt           time.Time                                             `json:"created_at"`
	UpdatedAt           time.Time                                             `json:"updated_at"`
	ExpiresAt           time.Time                                             `json:"expires_at"`
}

// MapIntegrationsSetupSessionsGetOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsGetOutput.
func MapIntegrationsSetupSessionsGetOutputFromJSON(data []byte) (*IntegrationsSetupSessionsGetOutput, error) {
	var v IntegrationsSetupSessionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsGetOutputToJSON serializes a IntegrationsSetupSessionsGetOutput to JSON.
func MapIntegrationsSetupSessionsGetOutputToJSON(v *IntegrationsSetupSessionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
