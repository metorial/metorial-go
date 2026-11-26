package authconfigs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsAuthConfigsListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsAuthConfigsListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsDeployment represents the provider deployments auth configs list output items deployment type.
type ProviderDeploymentsAuthConfigsListOutputItemsDeployment struct {
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

// ProviderDeploymentsAuthConfigsListOutputItemsCredentials represents the provider deployments auth configs list output items credentials type.
type ProviderDeploymentsAuthConfigsListOutputItemsCredentials struct {
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

// ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodInputSchema represents the provider deployments auth configs list output items auth method input schema type.
type ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the required auth input fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodOutputSchema represents the provider deployments auth configs list output items auth method output schema type.
type ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the auth output fields
	Schema map[string]any `json:"schema"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodScopes represents the provider deployments auth configs list output items auth method scopes type.
type ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodScopes struct {
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

// ProviderDeploymentsAuthConfigsListOutputItemsAuthMethod represents the provider deployments auth configs list output items auth method type.
type ProviderDeploymentsAuthConfigsListOutputItemsAuthMethod struct {
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
	Capabilities map[string]any                                                       `json:"capabilities"`
	InputSchema  *ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodOutputSchema `json:"output_schema,omitempty"`
	// Scopes - Available OAuth scopes
	Scopes *[]ProviderDeploymentsAuthConfigsListOutputItemsAuthMethodScopes `json:"scopes,omitempty"`
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

// ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsAuthConfigsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsAuthConfigsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsAuthConfigsListOutputItemsAccess - Resource ownership and access settings.
type ProviderDeploymentsAuthConfigsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsAuthConfigsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsAuthConfigsListOutputItems represents the provider deployments auth configs list output items type.
type ProviderDeploymentsAuthConfigsListOutputItems struct {
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
	ToolFilter  ProviderDeploymentsAuthConfigsListOutputItemsToolFilter   `json:"tool_filter"`
	Deployment  *ProviderDeploymentsAuthConfigsListOutputItemsDeployment  `json:"deployment,omitempty"`
	Credentials *ProviderDeploymentsAuthConfigsListOutputItemsCredentials `json:"credentials,omitempty"`
	AuthMethod  ProviderDeploymentsAuthConfigsListOutputItemsAuthMethod   `json:"auth_method"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsAuthConfigsListOutputItemsAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsAuthConfigsListOutputPagination represents the provider deployments auth configs list output pagination type.
type ProviderDeploymentsAuthConfigsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsAuthConfigsListOutput represents the provider deployments auth configs list output type.
type ProviderDeploymentsAuthConfigsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                             `json:"object"`
	Items      []ProviderDeploymentsAuthConfigsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsAuthConfigsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsAuthConfigsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsListOutput.
func MapProviderDeploymentsAuthConfigsListOutputFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsListOutput, error) {
	var v ProviderDeploymentsAuthConfigsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsListOutputToJSON serializes a ProviderDeploymentsAuthConfigsListOutput to JSON.
func MapProviderDeploymentsAuthConfigsListOutputToJSON(v *ProviderDeploymentsAuthConfigsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsAuthConfigsListQueryCreatedAt - Only return results whose provider auth config creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthConfigsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthConfigsListQueryUpdatedAt - Only return results whose provider auth config last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsAuthConfigsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsAuthConfigsListQuery represents the provider deployments auth configs list query type.
type ProviderDeploymentsAuthConfigsListQuery struct {
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
	// Ownership - Which resources to return under the resource ownership rules: shared (resources without an owner), usable (shared resources plus owned resources the caller may use), or all (everything the caller can see). Defaults to shared when ownership is enforced for the project, and all otherwise.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by auth config ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// AvailableForUse - Only return auth configs that are not owned by another integration instance provider.
	AvailableForUse *bool `json:"available_for_use,omitempty"`
	// AvailableForProviderDeploymentId - Only return auth configs that are not locked to a different provider deployment.
	AvailableForProviderDeploymentId *string `json:"available_for_provider_deployment_id,omitempty"`
	// ProviderAuthCredentialsId - Filter by auth credentials ID(s)
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderAuthMethodId - Filter by auth method ID(s)
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Adapter - Only return auth configs whose auth method supports this adapter
	Adapter *string `json:"adapter,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider auth config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsAuthConfigsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider auth config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsAuthConfigsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsAuthConfigsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsAuthConfigsListQuery.
func MapProviderDeploymentsAuthConfigsListQueryFromJSON(data []byte) (*ProviderDeploymentsAuthConfigsListQuery, error) {
	var v ProviderDeploymentsAuthConfigsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsAuthConfigsListQueryToJSON serializes a ProviderDeploymentsAuthConfigsListQuery to JSON.
func MapProviderDeploymentsAuthConfigsListQueryToJSON(v *ProviderDeploymentsAuthConfigsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
