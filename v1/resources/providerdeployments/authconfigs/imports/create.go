package imports

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigDeployment represents the provider deployments auth configs imports create output auth config deployment type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigDeployment struct {
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

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigCredentials represents the provider deployments auth configs imports create output auth config credentials type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigCredentials struct {
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

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodInputSchema represents the provider deployments auth configs imports create output auth config auth method input schema type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodOutputSchema represents the provider deployments auth configs imports create output auth config auth method output schema type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodScopes represents the provider deployments auth configs imports create output auth config auth method scopes type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethod represents the provider deployments auth configs imports create output auth config auth method type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethod struct {
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
	Capabilities map[string]any                                                                     `json:"capabilities"`
	InputSchema  *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfig represents the provider deployments auth configs imports create output auth config type.
type ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfig struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfigAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsAuthConfigsImportsCreateOutput represents the provider deployments auth configs imports create output type.
type ProviderDeploymentsAuthConfigsImportsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Note - Note explaining the import
	Note string `json:"note"`
	// Ip - IP address of the import request
	Ip *string `json:"ip,omitempty"`
	// UserAgent - User agent of the import request
	UserAgent *string `json:"user_agent,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata   *map[string]any                                             `json:"metadata,omitempty"`
	AuthConfig ProviderDeploymentsAuthConfigsImportsCreateOutputAuthConfig `json:"auth_config"`
	// ProviderId - Provider this import belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the imported credentials were used with. `null` when the auth config isn’t tied to a deployment.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// AuthMethodId - Auth method the imported credentials belong to.
	AuthMethodId string `json:"auth_method_id"`
	// CredentialsId - Imported credentials, or `null` when the auth config has none.
	CredentialsId *string `json:"credentials_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - Always `null`. Imports do not expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapProviderDeploymentsAuthConfigsImportsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsImportsCreateOutput.
func MapProviderDeploymentsAuthConfigsImportsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsImportsCreateOutput, error) {
	var v ProviderDeploymentsAuthConfigsImportsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsImportsCreateOutputToJSON serializes a ProviderDeploymentsAuthConfigsImportsCreateOutput to JSON.
func MapProviderDeploymentsAuthConfigsImportsCreateOutputToJSON(v *ProviderDeploymentsAuthConfigsImportsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthConfigsImportsCreateBody represents the provider deployments auth configs imports create body type.
type ProviderDeploymentsAuthConfigsImportsCreateBody struct {
	// ProviderId - Provider the imported credentials belong to.
	ProviderId *string `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Deployment the imported credentials belong to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthConfigId - Auth config the imported credentials belong to.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ProviderAuthMethodId - The authentication method used by these credentials
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// Note - A note describing the import source or reason
	Note string `json:"note"`
	// Metadata - Custom data to identify or organize the provider auth import. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Value - The credential data to import
	Value map[string]any `json:"value"`
}

// MapProviderDeploymentsAuthConfigsImportsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsImportsCreateBody.
func MapProviderDeploymentsAuthConfigsImportsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsImportsCreateBody, error) {
	var v ProviderDeploymentsAuthConfigsImportsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsImportsCreateBodyToJSON serializes a ProviderDeploymentsAuthConfigsImportsCreateBody to JSON.
func MapProviderDeploymentsAuthConfigsImportsCreateBodyToJSON(v *ProviderDeploymentsAuthConfigsImportsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
