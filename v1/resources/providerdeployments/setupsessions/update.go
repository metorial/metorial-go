package setupsessions

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsSetupSessionsUpdateOutputError - The error reported during the OAuth step of this session, or null if none occurred
type ProviderDeploymentsSetupSessionsUpdateOutputError struct {
	// Code - Machine-readable error code reported by the OAuth provider or integration
	Code string `json:"code"`
	// Message - Human-readable error message
	Message string `json:"message"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsUpdateOutputAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsUpdateOutputAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsUpdateOutputAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsUpdateOutputAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodInputSchema represents the provider deployments setup sessions update output auth method input schema type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodOutputSchema represents the provider deployments setup sessions update output auth method output schema type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodScopes represents the provider deployments setup sessions update output auth method scopes type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodScopes struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique scope identifier
	Id string `json:"id"`
	// Scope - OAuth scope string
	Scope string `json:"scope"`
	// Name - Display name of the scope
	Name string `json:"name"`
	// Description - Scope description
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthMethod represents the provider deployments setup sessions update output auth method type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthMethod struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth method identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Key - Auth method key
	Key string `json:"key"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Auth method description
	Description *string `json:"description,omitempty"`
	// Capabilities - Auth method capabilities
	Capabilities map[string]any                                                      `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsUpdateOutputAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputDeployment represents the provider deployments setup sessions update output deployment type.
type ProviderDeploymentsSetupSessionsUpdateOutputDeployment struct {
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

// ProviderDeploymentsSetupSessionsUpdateOutputCredentials represents the provider deployments setup sessions update output credentials type.
type ProviderDeploymentsSetupSessionsUpdateOutputCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// Status - Credentials status
	Status string `json:"status"`
	// IsDefault - Whether this is the default credentials for the provider
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether these credentials are managed by Metorial
	IsManaged bool `json:"is_managed"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigDeployment represents the provider deployments setup sessions update output auth config deployment type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigDeployment struct {
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

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigCredentials represents the provider deployments setup sessions update output auth config credentials type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigCredentials struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique credentials identifier
	Id   string `json:"id"`
	Type string `json:"type"`
	// Status - Credentials status
	Status string `json:"status"`
	// IsDefault - Whether this is the default credentials for the provider
	IsDefault bool `json:"is_default"`
	// IsManaged - Whether these credentials are managed by Metorial
	IsManaged bool `json:"is_managed"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Scopes - OAuth scopes requested by this credential
	Scopes *[]string `json:"scopes,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodInputSchema represents the provider deployments setup sessions update output auth config auth method input schema type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodOutputSchema represents the provider deployments setup sessions update output auth config auth method output schema type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodScopes represents the provider deployments setup sessions update output auth config auth method scopes type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodScopes struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique scope identifier
	Id string `json:"id"`
	// Scope - OAuth scope string
	Scope string `json:"scope"`
	// Name - Display name of the scope
	Name string `json:"name"`
	// Description - Scope description
	Description *string `json:"description,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethod represents the provider deployments setup sessions update output auth config auth method type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethod struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth method identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Key - Auth method key
	Key string `json:"key"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Auth method description
	Description *string `json:"description,omitempty"`
	// Capabilities - Auth method capabilities
	Capabilities map[string]any                                                                `json:"capabilities"`
	InputSchema  *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
	// Adapters - Adapter IDs this auth method can be used with
	Adapters *[]string `json:"adapters,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification ID
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputAuthConfig represents the provider deployments setup sessions update output auth config type.
type ProviderDeploymentsSetupSessionsUpdateOutputAuthConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique auth config identifier
	Id string `json:"id"`
	// Type - Authentication type
	Type string `json:"type"`
	// Source - Auth config source
	Source string `json:"source"`
	// Status - Auth config status
	Status string `json:"status"`
	// IsDefault - Whether this is the default auth config
	IsDefault bool `json:"is_default"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter  ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsSetupSessionsUpdateOutputAuthConfigAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigDeployment represents the provider deployments setup sessions update output config deployment type.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigDeployment struct {
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

// ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVaultDeployment represents the provider deployments setup sessions update output config from vault deployment type.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVaultDeployment struct {
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

// ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVault represents the provider deployments setup sessions update output config from vault type.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVault struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config vault identifier
	Id string `json:"id"`
	// Status - Status of the config vault
	Status string `json:"status"`
	// Name - Display name
	Name string `json:"name"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string                                                                 `json:"provider_id"`
	Deployment *ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsSetupSessionsUpdateOutputConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsSetupSessionsUpdateOutputConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsSetupSessionsUpdateOutputConfig represents the provider deployments setup sessions update output config type.
type ProviderDeploymentsSetupSessionsUpdateOutputConfig struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique config identifier
	Id string `json:"id"`
	// Status - Config status
	Status string `json:"status"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsSetupSessionsUpdateOutputConfigToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification ID
	SpecificationId string                                                        `json:"specification_id"`
	Deployment      *ProviderDeploymentsSetupSessionsUpdateOutputConfigDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsSetupSessionsUpdateOutputConfigFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsSetupSessionsUpdateOutputConfigAccess `json:"access"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsSetupSessionsUpdateOutput represents the provider deployments setup sessions update output type.
type ProviderDeploymentsSetupSessionsUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique setup session identifier
	Id string `json:"id"`
	// Type - Setup session type. `admin` sessions produce an admin consent link and complete without creating credentials.
	Type string `json:"type"`
	// Status - Session status
	Status string `json:"status"`
	// Url - URL where user completes authentication
	Url string `json:"url"`
	// Error - The error reported during the OAuth step of this session, or null if none occurred
	Error *ProviderDeploymentsSetupSessionsUpdateOutputError `json:"error,omitempty"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Setup session configuration
	Configuration *map[string]any `json:"configuration,omitempty"`
	// ProviderId - Provider ID
	ProviderId *string `json:"provider_id,omitempty"`
	// IdentityId - Linked identity ID
	IdentityId *string `json:"identity_id,omitempty"`
	// IdentityCredentialId - Identity credential created from this setup session
	IdentityCredentialId *string `json:"identity_credential_id,omitempty"`
	// Access - Resource ownership and access settings.
	Access      ProviderDeploymentsSetupSessionsUpdateOutputAccess       `json:"access"`
	AuthMethod  *ProviderDeploymentsSetupSessionsUpdateOutputAuthMethod  `json:"auth_method,omitempty"`
	Deployment  *ProviderDeploymentsSetupSessionsUpdateOutputDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsSetupSessionsUpdateOutputCredentials `json:"credentials,omitempty"`
	AuthConfig  *ProviderDeploymentsSetupSessionsUpdateOutputAuthConfig  `json:"auth_config,omitempty"`
	Config      *ProviderDeploymentsSetupSessionsUpdateOutputConfig      `json:"config,omitempty"`
	// UiMode - UI mode for setup
	UiMode string `json:"ui_mode"`
	// RedirectUrl - URL to redirect after setup
	RedirectUrl *string `json:"redirect_url,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
	// ExpiresAt - Timestamp when the session expires
	ExpiresAt time.Time `json:"expires_at"`
}

// MapProviderDeploymentsSetupSessionsUpdateOutputFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsUpdateOutput.
func MapProviderDeploymentsSetupSessionsUpdateOutputFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsUpdateOutput, error) {
	var v ProviderDeploymentsSetupSessionsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsUpdateOutputToJSON serializes a ProviderDeploymentsSetupSessionsUpdateOutput to JSON.
func MapProviderDeploymentsSetupSessionsUpdateOutputToJSON(v *ProviderDeploymentsSetupSessionsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsSetupSessionsUpdateBody represents the provider deployments setup sessions update body type.
type ProviderDeploymentsSetupSessionsUpdateBody struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IdentityId - Optional identity to link this setup session to
	IdentityId *string `json:"identity_id,omitempty"`
}

// MapProviderDeploymentsSetupSessionsUpdateBodyFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsUpdateBody.
func MapProviderDeploymentsSetupSessionsUpdateBodyFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsUpdateBody, error) {
	var v ProviderDeploymentsSetupSessionsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsUpdateBodyToJSON serializes a ProviderDeploymentsSetupSessionsUpdateBody to JSON.
func MapProviderDeploymentsSetupSessionsUpdateBodyToJSON(v *ProviderDeploymentsSetupSessionsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
