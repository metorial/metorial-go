package instances

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesDeleteOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstancesDeleteOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstancesDeleteOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstancesDeleteOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstancesDeleteOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstancesDeleteOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstancesDeleteOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstancesDeleteOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstancesDeleteOutputAccess - Resource ownership and access settings.
type IntegrationsInstancesDeleteOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstancesDeleteOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstancesDeleteOutputImplementation - What backs this instance. `null` for a standalone instance.
type IntegrationsInstancesDeleteOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpServerId - Magic MCP server backing this instance.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesDeleteOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesDeleteOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersProvider represents the integrations instances delete output providers provider type.
type IntegrationsInstancesDeleteOutputProvidersProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig represents the integrations instances delete output providers integration provider config type.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesDeleteOutputProvidersIntegrationProvider represents the integrations instances delete output providers integration provider type.
type IntegrationsInstancesDeleteOutputProvidersIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsInstancesDeleteOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesDeleteOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                              `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesDeleteOutputProvidersIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesDeleteOutputProvidersConfig represents the integrations instances delete output providers config type.
type IntegrationsInstancesDeleteOutputProvidersConfig struct {
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

// IntegrationsInstancesDeleteOutputProvidersAuthConfig represents the integrations instances delete output providers auth config type.
type IntegrationsInstancesDeleteOutputProvidersAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesDeleteOutputProviders represents the integrations instances delete output providers type.
type IntegrationsInstancesDeleteOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the entry is archived.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Integration instance this entry belongs to.
	IntegrationInstanceId string `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesDeleteOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesDeleteOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesDeleteOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesDeleteOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesDeleteOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesDeleteOutput represents the integrations instances delete output type.
type IntegrationsInstancesDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `draft` when created. `archived` once you archive the instance.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IdentityActorId - Identity actor the instance connects as. `null` when none is set.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance connects as. `null` when none is set.
	IdentityId *string `json:"identity_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access IntegrationsInstancesDeleteOutputAccess `json:"access"`
	// Implementation - What backs this instance. `null` for a standalone instance.
	Implementation *IntegrationsInstancesDeleteOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers set up on this instance.
	Providers []IntegrationsInstancesDeleteOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesDeleteOutputFromJSON deserializes JSON data into a IntegrationsInstancesDeleteOutput.
func MapIntegrationsInstancesDeleteOutputFromJSON(data []byte) (*IntegrationsInstancesDeleteOutput, error) {
	var v IntegrationsInstancesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesDeleteOutputToJSON serializes a IntegrationsInstancesDeleteOutput to JSON.
func MapIntegrationsInstancesDeleteOutputToJSON(v *IntegrationsInstancesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
