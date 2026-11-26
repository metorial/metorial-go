package configs

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigsListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsConfigsListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsConfigsListOutputItemsDeployment represents the provider deployments configs list output items deployment type.
type ProviderDeploymentsConfigsListOutputItemsDeployment struct {
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

// ProviderDeploymentsConfigsListOutputItemsFromVaultDeployment represents the provider deployments configs list output items from vault deployment type.
type ProviderDeploymentsConfigsListOutputItemsFromVaultDeployment struct {
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

// ProviderDeploymentsConfigsListOutputItemsFromVault represents the provider deployments configs list output items from vault type.
type ProviderDeploymentsConfigsListOutputItemsFromVault struct {
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
	ProviderId string                                                        `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigsListOutputItemsFromVaultDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsListOutputItemsAccessOwnerOrganizationMember - Organization member associated with this owner.
type ProviderDeploymentsConfigsListOutputItemsAccessOwnerOrganizationMember struct {
	// Id - Identifier of the organization member who owns the resource.
	Id string `json:"id"`
}

// ProviderDeploymentsConfigsListOutputItemsAccessOwnerConsumer - Consumer and profile associated with this owner.
type ProviderDeploymentsConfigsListOutputItemsAccessOwnerConsumer struct {
	// Id - Identifier of the owning consumer.
	Id string `json:"id"`
	// ProfileId - Identifier of the owning consumer profile.
	ProfileId string `json:"profile_id"`
	// PortalId - Identifier of the consumer profile’s portal, if any.
	PortalId *string `json:"portal_id,omitempty"`
}

// ProviderDeploymentsConfigsListOutputItemsAccessOwner - The actor that owns this resource, if any.
type ProviderDeploymentsConfigsListOutputItemsAccessOwner struct {
	// Type - Type of the owning actor.
	Type string `json:"type"`
	// Name - Display name of the owning actor.
	Name *string `json:"name,omitempty"`
	// ImageUrl - Avatar image URL of the owning actor, if available.
	ImageUrl *string `json:"image_url,omitempty"`
	// Email - Email address of the owning actor, if available.
	Email *string `json:"email,omitempty"`
	// OrganizationMember - Organization member associated with this owner.
	OrganizationMember *ProviderDeploymentsConfigsListOutputItemsAccessOwnerOrganizationMember `json:"organization_member,omitempty"`
	// Consumer - Consumer and profile associated with this owner.
	Consumer *ProviderDeploymentsConfigsListOutputItemsAccessOwnerConsumer `json:"consumer,omitempty"`
}

// ProviderDeploymentsConfigsListOutputItemsAccess - Resource ownership and access settings.
type ProviderDeploymentsConfigsListOutputItemsAccess struct {
	// Source - Integration resource source, or null when the resource is not from Explorer.
	Source *string `json:"source,omitempty"`
	// Type - Whether usage is public or locked to the owner under the project ownership rules.
	Type string `json:"type"`
	// Visibility - Whether visibility is unrestricted or limited to the owner under the project ownership rules.
	Visibility string `json:"visibility"`
	// Owner - The actor that owns this resource, if any.
	Owner *ProviderDeploymentsConfigsListOutputItemsAccessOwner `json:"owner,omitempty"`
}

// ProviderDeploymentsConfigsListOutputItems represents the provider deployments configs list output items type.
type ProviderDeploymentsConfigsListOutputItems struct {
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
	ToolFilter ProviderDeploymentsConfigsListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider this config belongs to.
	ProviderId string `json:"provider_id"`
	// SpecificationId - Specification this config follows.
	SpecificationId string                                               `json:"specification_id"`
	Deployment      *ProviderDeploymentsConfigsListOutputItemsDeployment `json:"deployment,omitempty"`
	FromVault       *ProviderDeploymentsConfigsListOutputItemsFromVault  `json:"from_vault,omitempty"`
	// Access - Resource ownership and access settings.
	Access ProviderDeploymentsConfigsListOutputItemsAccess `json:"access"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigsListOutputPagination represents the provider deployments configs list output pagination type.
type ProviderDeploymentsConfigsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsConfigsListOutput represents the provider deployments configs list output type.
type ProviderDeploymentsConfigsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                         `json:"object"`
	Items      []ProviderDeploymentsConfigsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsConfigsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsConfigsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigsListOutput.
func MapProviderDeploymentsConfigsListOutputFromJSON(data []byte) (*ProviderDeploymentsConfigsListOutput, error) {
	var v ProviderDeploymentsConfigsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsListOutputToJSON serializes a ProviderDeploymentsConfigsListOutput to JSON.
func MapProviderDeploymentsConfigsListOutputToJSON(v *ProviderDeploymentsConfigsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigsListQueryCreatedAt - Only return results whose provider config creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsConfigsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsConfigsListQueryUpdatedAt - Only return results whose provider config last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsConfigsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsConfigsListQuery represents the provider deployments configs list query type.
type ProviderDeploymentsConfigsListQuery struct {
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
	// Ownership - Which resources to return under the resource ownership rules. 'shared' returns only resources that aren't owned, 'usable' returns shared resources plus owned resources the caller can use, 'all' returns every resource the caller can see. Defaults to 'usable' for consumers, and otherwise to 'shared' when ownership is enforced for the project and 'all' when it isn't.
	Ownership *string `json:"ownership,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by config ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderSpecificationId - Filter by provider specification ID(s)
	ProviderSpecificationId *any `json:"provider_specification_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// AvailableForUse - Only return configs that are not owned by another integration instance provider.
	AvailableForUse *bool `json:"available_for_use,omitempty"`
	// AvailableForProviderDeploymentId - Only return configs that are not locked to a different provider deployment.
	AvailableForProviderDeploymentId *string `json:"available_for_provider_deployment_id,omitempty"`
	// ProviderConfigVaultId - Filter by config vault ID(s)
	ProviderConfigVaultId *any `json:"provider_config_vault_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsConfigsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsConfigsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsConfigsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsConfigsListQuery.
func MapProviderDeploymentsConfigsListQueryFromJSON(data []byte) (*ProviderDeploymentsConfigsListQuery, error) {
	var v ProviderDeploymentsConfigsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigsListQueryToJSON serializes a ProviderDeploymentsConfigsListQuery to JSON.
func MapProviderDeploymentsConfigsListQueryToJSON(v *ProviderDeploymentsConfigsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
