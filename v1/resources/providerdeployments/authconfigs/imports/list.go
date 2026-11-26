package imports

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigDeployment represents the provider deployments auth configs imports list output items auth config deployment type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigDeployment struct {
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

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigCredentials represents the provider deployments auth configs imports list output items auth config credentials type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigCredentials struct {
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

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodInputSchema represents the provider deployments auth configs imports list output items auth config auth method input schema type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodOutputSchema represents the provider deployments auth configs imports list output items auth config auth method output schema type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodScopes represents the provider deployments auth configs imports list output items auth config auth method scopes type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethod represents the provider deployments auth configs imports list output items auth config auth method type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethod struct {
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
	Capabilities map[string]any                                                                        `json:"capabilities"`
	InputSchema  *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfig represents the provider deployments auth configs imports list output items auth config type.
type ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfig struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfigAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsAuthConfigsImportsListOutputItems represents the provider deployments auth configs imports list output items type.
type ProviderDeploymentsAuthConfigsImportsListOutputItems struct {
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
	Metadata   *map[string]any                                                `json:"metadata,omitempty"`
	AuthConfig ProviderDeploymentsAuthConfigsImportsListOutputItemsAuthConfig `json:"auth_config"`
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

// ProviderDeploymentsAuthConfigsImportsListOutputPagination represents the provider deployments auth configs imports list output pagination type.
type ProviderDeploymentsAuthConfigsImportsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListOutput represents the provider deployments auth configs imports list output type.
type ProviderDeploymentsAuthConfigsImportsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                                    `json:"object"`
	Items      []ProviderDeploymentsAuthConfigsImportsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsAuthConfigsImportsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsAuthConfigsImportsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsImportsListOutput.
func MapProviderDeploymentsAuthConfigsImportsListOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsImportsListOutput, error) {
	var v ProviderDeploymentsAuthConfigsImportsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsImportsListOutputToJSON serializes a ProviderDeploymentsAuthConfigsImportsListOutput to JSON.
func MapProviderDeploymentsAuthConfigsImportsListOutputToJSON(v *ProviderDeploymentsAuthConfigsImportsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthConfigsImportsListQueryCreatedAt - Only return results whose provider auth import creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthConfigsImportsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListQueryUpdatedAt - Only return results whose provider auth import last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthConfigsImportsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthConfigsImportsListQuery represents the provider deployments auth configs imports list query type.
type ProviderDeploymentsAuthConfigsImportsListQuery struct {
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
	// Id - Filter by import ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderAuthConfigId - Filter by auth config ID(s)
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// CreatedAt - Only return results whose provider auth import creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsAuthConfigsImportsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth import last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsAuthConfigsImportsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsAuthConfigsImportsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsImportsListQuery.
func MapProviderDeploymentsAuthConfigsImportsListQueryFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsImportsListQuery, error) {
	var v ProviderDeploymentsAuthConfigsImportsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsImportsListQueryToJSON serializes a ProviderDeploymentsAuthConfigsImportsListQuery to JSON.
func MapProviderDeploymentsAuthConfigsImportsListQueryToJSON(v *ProviderDeploymentsAuthConfigsImportsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
