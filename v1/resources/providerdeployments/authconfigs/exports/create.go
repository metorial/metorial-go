package exports

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigDeployment represents the provider deployments auth configs exports create output auth config deployment type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigDeployment struct {
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

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigCredentials represents the provider deployments auth configs exports create output auth config credentials type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigCredentials struct {
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

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodInputSchema represents the provider deployments auth configs exports create output auth config auth method input schema type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodOutputSchema represents the provider deployments auth configs exports create output auth config auth method output schema type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodScopes represents the provider deployments auth configs exports create output auth config auth method scopes type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethod represents the provider deployments auth configs exports create output auth config auth method type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethod struct {
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
	InputSchema  *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfig represents the provider deployments auth configs exports create output auth config type.
type ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfig struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfigAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsAuthConfigsExportsCreateOutput represents the provider deployments auth configs exports create output type.
type ProviderDeploymentsAuthConfigsExportsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Note - Note explaining the export reason
	Note string `json:"note"`
	// Ip - IP address of the export request
	Ip *string `json:"ip,omitempty"`
	// UserAgent - User agent of the export request
	UserAgent *string `json:"user_agent,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata   *map[string]any                                             `json:"metadata,omitempty"`
	AuthConfig ProviderDeploymentsAuthConfigsExportsCreateOutputAuthConfig `json:"auth_config"`
	// ProviderId - Provider this export belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the exported credentials were used with. `null` when the auth config isn’t tied to a deployment.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// AuthMethodId - Auth method the exported credentials belong to.
	AuthMethodId string `json:"auth_method_id"`
	// CredentialsId - Exported credentials, or `null` when the auth config has none.
	CredentialsId *string `json:"credentials_id,omitempty"`
	// Value - Exported credential data. Only returned when creating the export.
	Value *map[string]any `json:"value,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - Set by the backend when the export was created.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapProviderDeploymentsAuthConfigsExportsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsExportsCreateOutput.
func MapProviderDeploymentsAuthConfigsExportsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsExportsCreateOutput, error) {
	var v ProviderDeploymentsAuthConfigsExportsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsExportsCreateOutputToJSON serializes a ProviderDeploymentsAuthConfigsExportsCreateOutput to JSON.
func MapProviderDeploymentsAuthConfigsExportsCreateOutputToJSON(v *ProviderDeploymentsAuthConfigsExportsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthConfigsExportsCreateBody represents the provider deployments auth configs exports create body type.
type ProviderDeploymentsAuthConfigsExportsCreateBody struct {
	// ProviderAuthConfigId - Auth config to export credentials from.
	ProviderAuthConfigId string `json:"provider_auth_config_id"`
	// Note - Why the credentials are exported.
	Note string `json:"note"`
	// Metadata - Custom data to identify or organize the provider auth export. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapProviderDeploymentsAuthConfigsExportsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsExportsCreateBody.
func MapProviderDeploymentsAuthConfigsExportsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsExportsCreateBody, error) {
	var v ProviderDeploymentsAuthConfigsExportsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsExportsCreateBodyToJSON serializes a ProviderDeploymentsAuthConfigsExportsCreateBody to JSON.
func MapProviderDeploymentsAuthConfigsExportsCreateBodyToJSON(v *ProviderDeploymentsAuthConfigsExportsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
