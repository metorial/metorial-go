package setupsessions

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsSetupSessionsListOutputItemsError - The error reported during the OAuth step of this session, or null if none occurred
type ProviderDeploymentsSetupSessionsListOutputItemsError struct {
	// Code - Machine-readable error code reported by the OAuth provider or integration
	Code string `json:"code"`
	// Message - Human-readable error message
	Message string `json:"message"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodInputSchema represents the provider deployments setup sessions list output items auth method input schema type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodOutputSchema represents the provider deployments setup sessions list output items auth method output schema type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodScopes represents the provider deployments setup sessions list output items auth method scopes type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodScopes struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Scope - The scope value sent to the provider.
	Scope string `json:"scope"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthMethod represents the provider deployments setup sessions list output items auth method type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthMethod struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the user authenticates with the provider.
	Type string `json:"type"`
	// Key - Stable key for the auth method.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What this auth method supports.
	Capabilities map[string]any                                                         `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsListOutputItemsAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider this auth method belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification this auth method belongs to.
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsDeployment represents the provider deployments setup sessions list output items deployment type.
type ProviderDeploymentsSetupSessionsListOutputItemsDeployment struct {
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

// ProviderDeploymentsSetupSessionsListOutputItemsCredentials represents the provider deployments setup sessions list output items credentials type.
type ProviderDeploymentsSetupSessionsListOutputItemsCredentials struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Must be `oauth`.
	Type string `json:"type"`
	// Status - Archived once the credentials are archived. Archived credentials stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name credentials.
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether Metorial or the installation administrator manages these OAuth credentials.
	IsManaged bool `json:"is_managed"`
	// ManagedBy - Who manages these credentials: Metorial, or the administrator of this self-hosted installation. Null for credentials created in the project.
	ManagedBy *string `json:"managed_by,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider these credentials belong to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigDeployment represents the provider deployments setup sessions list output items auth config deployment type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigDeployment struct {
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

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigCredentials represents the provider deployments setup sessions list output items auth config credentials type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigCredentials struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Must be `oauth`.
	Type string `json:"type"`
	// Status - Archived once the credentials are archived. Archived credentials stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name credentials.
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether Metorial or the installation administrator manages these OAuth credentials.
	IsManaged bool `json:"is_managed"`
	// ManagedBy - Who manages these credentials: Metorial, or the administrator of this self-hosted installation. Null for credentials created in the project.
	ManagedBy *string `json:"managed_by,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider these credentials belong to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodInputSchema represents the provider deployments setup sessions list output items auth config auth method input schema type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodOutputSchema represents the provider deployments setup sessions list output items auth config auth method output schema type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodScopes represents the provider deployments setup sessions list output items auth config auth method scopes type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodScopes struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Scope - The scope value sent to the provider.
	Scope string `json:"scope"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethod represents the provider deployments setup sessions list output items auth config auth method type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethod struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the user authenticates with the provider.
	Type string `json:"type"`
	// Key - Stable key for the auth method.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Capabilities - What this auth method supports.
	Capabilities map[string]any                                                                   `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider this auth method belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification this auth method belongs to.
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsAuthConfig represents the provider deployments setup sessions list output items auth config type.
type ProviderDeploymentsSetupSessionsListOutputItemsAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the auth config was created. `oauth_automated` configs are created automatically and cannot be updated.
	Type string `json:"type"`
	// Source - Where the auth config was created.
	Source string `json:"source"`
	// Status - Archived once the auth config is archived. Archived auth configs stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter  ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsSetupSessionsListOutputItemsAuthConfigAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigDeployment represents the provider deployments setup sessions list output items config deployment type.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigDeployment struct {
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

// ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVaultDeployment represents the provider deployments setup sessions list output items config from vault deployment type.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVaultDeployment struct {
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

// ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVault represents the provider deployments setup sessions list output items config from vault type.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVault struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the vault is archived. Archived vaults stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this vault belongs to.
	ProviderId string                                                                    `json:"provider_id"`
	Deployment *ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsListOutputItemsConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsListOutputItemsConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutputItemsConfig represents the provider deployments setup sessions list output items config type.
type ProviderDeploymentsSetupSessionsListOutputItemsConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the config is archived. Archived configs stay listable.
	Status string `json:"status"`
	// IsDefault - The deployment falls back to this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsSetupSessionsListOutputItemsConfigToolFilter `json:"tool_filter"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification this config follows.
	SpecificationId string                                                           `json:"specification_id"`
	Deployment      *ProviderDeploymentsSetupSessionsListOutputItemsConfigDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsSetupSessionsListOutputItemsConfigFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsSetupSessionsListOutputItemsConfigAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsListOutputItems represents the provider deployments setup sessions list output items type.
type ProviderDeploymentsSetupSessionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Setup session type. `admin` sessions produce an admin consent link and complete without creating credentials.
	Type string `json:"type"`
	// Status - Session status. `pending` until the flow completes. Reads as `expired` once past `expires_at`. `archived` once the session is archived.
	Status string `json:"status"`
	// Url - URL where the user completes this session.
	Url string `json:"url"`
	// Error - The error reported during the OAuth step of this session, or null if none occurred
	Error *ProviderDeploymentsSetupSessionsListOutputItemsError `json:"error,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Options shaping the session flow, like provider search filters and UI layout.
	Configuration *map[string]any `json:"configuration,omitempty"`
	// ProviderId - Provider the session sets up, or `null` when none was chosen.
	ProviderId *string `json:"provider_id,omitempty"`
	// IdentityId - Identity the session is linked to. `null` when unlinked.
	IdentityId *string `json:"identity_id,omitempty"`
	// IdentityCredentialId - Identity credential created from this setup session
	IdentityCredentialId *string `json:"identity_credential_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access      ProviderDeploymentsSetupSessionsListOutputItemsAccess       `json:"access"`
	AuthMethod  *ProviderDeploymentsSetupSessionsListOutputItemsAuthMethod  `json:"auth_method,omitempty"`
	Deployment  *ProviderDeploymentsSetupSessionsListOutputItemsDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsSetupSessionsListOutputItemsCredentials `json:"credentials,omitempty"`
	AuthConfig  *ProviderDeploymentsSetupSessionsListOutputItemsAuthConfig  `json:"auth_config,omitempty"`
	Config      *ProviderDeploymentsSetupSessionsListOutputItemsConfig      `json:"config,omitempty"`
	// UiMode - How the session UI is presented.
	UiMode string `json:"ui_mode"`
	// RedirectUrl - URL to redirect after setup
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ExpiresAt - When the session expires. `pending` sessions read as `expired` past this time.
	ExpiresAt time.Time `json:"expires_at"`
}

// ProviderDeploymentsSetupSessionsListOutputPagination represents the provider deployments setup sessions list output pagination type.
type ProviderDeploymentsSetupSessionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsSetupSessionsListOutput represents the provider deployments setup sessions list output type.
type ProviderDeploymentsSetupSessionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                               `json:"object"`
	Items      []ProviderDeploymentsSetupSessionsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsSetupSessionsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsSetupSessionsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsListOutput.
func MapProviderDeploymentsSetupSessionsListOutputFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsListOutput, error) {
	var v ProviderDeploymentsSetupSessionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsListOutputToJSON serializes a ProviderDeploymentsSetupSessionsListOutput to JSON.
func MapProviderDeploymentsSetupSessionsListOutputToJSON(v *ProviderDeploymentsSetupSessionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsSetupSessionsListQueryCreatedAt - Only return results whose provider setup session creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsSetupSessionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsSetupSessionsListQueryUpdatedAt - Only return results whose provider setup session last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsSetupSessionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsSetupSessionsListQuery represents the provider deployments setup sessions list query type.
type ProviderDeploymentsSetupSessionsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Filter by setup session ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Filter by auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthConfigId - Filter by provider auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by provider auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// Status - Filter by session status (archived, failed, completed, expired, pending)
	Status *any `json:"status,omitempty"`
	// Type - Filter by session type (auth_only, config_only, auth_and_config, admin), matching either the requested or the resolved type
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose provider setup session creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsSetupSessionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider setup session last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsSetupSessionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsSetupSessionsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsListQuery.
func MapProviderDeploymentsSetupSessionsListQueryFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsListQuery, error) {
	var v ProviderDeploymentsSetupSessionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsListQueryToJSON serializes a ProviderDeploymentsSetupSessionsListQuery to JSON.
func MapProviderDeploymentsSetupSessionsListQueryToJSON(v *ProviderDeploymentsSetupSessionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
