package setupsessions

import (
	"encoding/json"
	"time"
)

// IntegrationsSetupSessionsCreateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsCreateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsCreateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsCreateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsCreateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsCreateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsCreateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsCreateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsCreateOutputAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerOrganizationMember - Organization member associated with this owner.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerConsumer - Consumer and profile associated with this owner.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwner - The actor that owns this resource, if any.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwnerConsumer `json:"consumer,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccess - Resource ownership and access settings.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccessOwner `json:"owner,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation - What backs this instance. `null` for a standalone instance.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation struct {
	// Type - Which kind of backing this is.
	Type string `json:"type"`
	// MagicMcpServerId - Magic MCP server backing this instance.
	MagicMcpServerId string `json:"magic_mcp_server_id"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider represents the integrations setup sessions create output integration instance providers provider type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider struct {
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

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig represents the integrations setup sessions create output integration instance providers integration provider config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig struct {
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

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider represents the integrations setup sessions create output integration instance providers integration provider type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                                                     `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig represents the integrations setup sessions create output integration instance providers config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig struct {
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

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig represents the integrations setup sessions create output integration instance providers auth config type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig struct {
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

// IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders represents the integrations setup sessions create output integration instance providers type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders struct {
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
	ToolFilter *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                                                 `json:"is_override_tool_filter"`
	Provider             IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersProvider            `json:"provider"`
	IntegrationProvider  IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsSetupSessionsCreateOutputIntegrationInstanceProvidersAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutputIntegrationInstance represents the integrations setup sessions create output integration instance type.
type IntegrationsSetupSessionsCreateOutputIntegrationInstance struct {
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
	Access IntegrationsSetupSessionsCreateOutputIntegrationInstanceAccess `json:"access"`
	// Implementation - What backs this instance. `null` for a standalone instance.
	Implementation *IntegrationsSetupSessionsCreateOutputIntegrationInstanceImplementation `json:"implementation,omitempty"`
	// Providers - Providers set up on this instance.
	Providers []IntegrationsSetupSessionsCreateOutputIntegrationInstanceProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsSetupSessionsCreateOutput represents the integrations setup sessions create output type.
type IntegrationsSetupSessionsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `pending` until every provider is configured. `successful` once every provider is configured. `expired` past `expires_at` without completion.
	Status string `json:"status"`
	// Url - Hosted setup page URL.
	Url string `json:"url"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Configuration for the setup flow. `null` when none was set.
	Configuration *map[string]any `json:"configuration,omitempty"`
	// RedirectUrl - Where to send the user once setup completes. `null` when none was given.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// IntegrationId - Integration being set up.
	IntegrationId string `json:"integration_id"`
	// Access - Resource ownership and access settings.
	Access              IntegrationsSetupSessionsCreateOutputAccess              `json:"access"`
	IntegrationInstance IntegrationsSetupSessionsCreateOutputIntegrationInstance `json:"integration_instance"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ExpiresAt - When the session expires. Past this time the session reads as `expired`.
	ExpiresAt time.Time `json:"expires_at"`
}

// MapIntegrationsSetupSessionsCreateOutputFromJSON deserializes JSON data into a IntegrationsSetupSessionsCreateOutput.
func MapIntegrationsSetupSessionsCreateOutputFromJSON(data []byte) (*IntegrationsSetupSessionsCreateOutput, error) {
	var v IntegrationsSetupSessionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsCreateOutputToJSON serializes a IntegrationsSetupSessionsCreateOutput to JSON.
func MapIntegrationsSetupSessionsCreateOutputToJSON(v *IntegrationsSetupSessionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups represents the integrations setup sessions create body configuration provider search groups type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups struct {
	// GroupId - Provider group to include in the search.
	GroupId string `json:"group_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections represents the integrations setup sessions create body configuration provider search collections type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections struct {
	// CollectionId - Provider collection to include in the search.
	CollectionId string `json:"collection_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories represents the integrations setup sessions create body configuration provider search categories type.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories struct {
	// CategoryId - Provider category to include in the search.
	CategoryId string `json:"category_id"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch - Narrows the setup flow’s provider search to these catalog entries.
type IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch struct {
	// Groups - Provider groups to include in the search.
	Groups *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchGroups `json:"groups,omitempty"`
	// Collections - Provider collections to include in the search.
	Collections *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCollections `json:"collections,omitempty"`
	// Categories - Provider categories to include in the search.
	Categories *[]IntegrationsSetupSessionsCreateBodyConfigurationProviderSearchCategories `json:"categories,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationToolFilters - Tool filter selection for the setup flow.
type IntegrationsSetupSessionsCreateBodyConfigurationToolFilters struct {
	// Enabled - Lets the user pick tool filters during setup. Only applies when the integration allows custom tool filters.
	Enabled *bool `json:"enabled,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfigurationUi - Presentation of the setup flow.
type IntegrationsSetupSessionsCreateBodyConfigurationUi struct {
	// Layout - Layout of the setup flow. Defaults to `box`.
	Layout *string `json:"layout,omitempty"`
}

// IntegrationsSetupSessionsCreateBodyConfiguration - Configuration for the setup flow.
type IntegrationsSetupSessionsCreateBodyConfiguration struct {
	// ProviderSearch - Narrows the setup flow’s provider search to these catalog entries.
	ProviderSearch *IntegrationsSetupSessionsCreateBodyConfigurationProviderSearch `json:"provider_search,omitempty"`
	// ToolFilters - Tool filter selection for the setup flow.
	ToolFilters *IntegrationsSetupSessionsCreateBodyConfigurationToolFilters `json:"tool_filters,omitempty"`
	// Ui - Presentation of the setup flow.
	Ui *IntegrationsSetupSessionsCreateBodyConfigurationUi `json:"ui,omitempty"`
}

// IntegrationsSetupSessionsCreateBody represents the integrations setup sessions create body type.
type IntegrationsSetupSessionsCreateBody struct {
	// IntegrationId - Integration to set up.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Existing instance to reuse. Omitted creates a draft instance.
	IntegrationInstanceId *string `json:"integration_instance_id,omitempty"`
	// Name - Display name for the setup session.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the setup session.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IdentityActorId - Identity actor the instance is set up for. Cannot be set when `integration_instance_id` is given.
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity the instance is set up for. Cannot be set when `integration_instance_id` is given.
	IdentityId *string `json:"identity_id,omitempty"`
	// ExpiresAt - When the session expires. Defaults to 30 minutes after creation.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// RedirectUrl - Where to send the user once setup completes.
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// Configuration - Configuration for the setup flow.
	Configuration *IntegrationsSetupSessionsCreateBodyConfiguration `json:"configuration,omitempty"`
}

// MapIntegrationsSetupSessionsCreateBodyFromJSON deserializes JSON data into a IntegrationsSetupSessionsCreateBody.
func MapIntegrationsSetupSessionsCreateBodyFromJSON(data []byte) (*IntegrationsSetupSessionsCreateBody, error) {
	var v IntegrationsSetupSessionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsSetupSessionsCreateBodyToJSON serializes a IntegrationsSetupSessionsCreateBody to JSON.
func MapIntegrationsSetupSessionsCreateBodyToJSON(v *IntegrationsSetupSessionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
