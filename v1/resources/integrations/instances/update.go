package instances

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// IntegrationsInstancesUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsInstancesUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsInstancesUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsInstancesUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsInstancesUpdateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsInstancesUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsInstancesUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsInstancesUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsInstancesUpdateOutputAccess - Resource ownership and access settings.
type IntegrationsInstancesUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsInstancesUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsInstancesUpdateOutputImplementation - What backs this instance. `null` for a standalone instance.
type IntegrationsInstancesUpdateOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpServerId - Magic MCP server backing this instance.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsInstancesUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersProvider represents the integrations instances update output providers provider type.
type IntegrationsInstancesUpdateOutputProvidersProvider struct {
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

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig represents the integrations instances update output providers integration provider config type.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig struct {
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

// IntegrationsInstancesUpdateOutputProvidersIntegrationProvider represents the integrations instances update output providers integration provider type.
type IntegrationsInstancesUpdateOutputProvidersIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsInstancesUpdateOutputProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesUpdateOutputProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                              `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesUpdateOutputProvidersIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesUpdateOutputProvidersConfig represents the integrations instances update output providers config type.
type IntegrationsInstancesUpdateOutputProvidersConfig struct {
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

// IntegrationsInstancesUpdateOutputProvidersAuthConfig represents the integrations instances update output providers auth config type.
type IntegrationsInstancesUpdateOutputProvidersAuthConfig struct {
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

// IntegrationsInstancesUpdateOutputProviders represents the integrations instances update output providers type.
type IntegrationsInstancesUpdateOutputProviders struct {
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
	ToolFilter *IntegrationsInstancesUpdateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                          `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesUpdateOutputProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesUpdateOutputProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesUpdateOutputProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesUpdateOutputProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesUpdateOutput represents the integrations instances update output type.
type IntegrationsInstancesUpdateOutput struct {
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
	Access IntegrationsInstancesUpdateOutputAccess `json:"access"`
	// Implementation - What backs this instance. `null` for a standalone instance.
	Implementation *IntegrationsInstancesUpdateOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers set up on this instance.
	Providers []IntegrationsInstancesUpdateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsInstancesUpdateOutputFromJSON deserializes JSON data into a IntegrationsInstancesUpdateOutput.
func MapIntegrationsInstancesUpdateOutputFromJSON(data []byte) (*IntegrationsInstancesUpdateOutput, error) {
	var v IntegrationsInstancesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesUpdateOutputToJSON serializes a IntegrationsInstancesUpdateOutput to JSON.
func MapIntegrationsInstancesUpdateOutputToJSON(v *IntegrationsInstancesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesUpdateBodyProviders represents the integrations instances update body providers type.
type IntegrationsInstancesUpdateBodyProviders struct {
	// ProviderId - Provider to bind, by provider ID or integration provider ID.
	ProviderId string `json:"provider_id"`
	// ProviderConfigId - Provider config the entry runs with. Omitted uses the shared config.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Credentials the entry runs with. Omitted leaves them unset.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter *bool `json:"is_override_tool_filter,omitempty"`
}

// IntegrationsInstancesUpdateBody represents the integrations instances update body type.
type IntegrationsInstancesUpdateBody struct {
	// Name - Display name for the integration instance.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration instance. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// IdentityActorId - Identity actor the instance connects as. `null` clears it. Omitted keeps the current one.
	IdentityActorId *nullable.Nullable[string] `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance connects as. `null` clears it. Omitted keeps the current one.
	IdentityId *nullable.Nullable[string] `json:"identity_id,omitempty"`
	// Providers - Providers to create or update. Existing providers not listed are left alone.
	Providers *[]IntegrationsInstancesUpdateBodyProviders `json:"providers,omitempty"`
}

// MapIntegrationsInstancesUpdateBodyFromJSON deserializes JSON data into a IntegrationsInstancesUpdateBody.
func MapIntegrationsInstancesUpdateBodyFromJSON(data []byte) (*IntegrationsInstancesUpdateBody, error) {
	var v IntegrationsInstancesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesUpdateBodyToJSON serializes a IntegrationsInstancesUpdateBody to JSON.
func MapIntegrationsInstancesUpdateBodyToJSON(v *IntegrationsInstancesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
