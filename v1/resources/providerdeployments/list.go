package providerdeployments

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProviderDeploymentsListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// ProviderDeploymentsListOutputItemsLockedVersion represents the provider deployments list output items locked version type.
type ProviderDeploymentsListOutputItemsLockedVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsListOutputItemsDefaultConfig represents the provider deployments list output items default config type.
type ProviderDeploymentsListOutputItemsDefaultConfig struct {
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

// ProviderDeploymentsListOutputItems represents the provider deployments list output items type.
type ProviderDeploymentsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the deployment is archived. Archived deployments stay listable.
	Status string `json:"status"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter ProviderDeploymentsListOutputItemsToolFilter `json:"tool_filter"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId    string                                           `json:"provider_id"`
	LockedVersion *ProviderDeploymentsListOutputItemsLockedVersion `json:"locked_version,omitempty"`
	DefaultConfig *ProviderDeploymentsListOutputItemsDefaultConfig `json:"default_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsListOutputPagination represents the provider deployments list output pagination type.
type ProviderDeploymentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsListOutput represents the provider deployments list output type.
type ProviderDeploymentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []ProviderDeploymentsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsListOutput.
func MapProviderDeploymentsListOutputFromJSON(data []byte) (*ProviderDeploymentsListOutput, error) {
	var v ProviderDeploymentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsListOutputToJSON serializes a ProviderDeploymentsListOutput to JSON.
func MapProviderDeploymentsListOutputToJSON(v *ProviderDeploymentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsListQueryCreatedAt - Only return results whose provider deployment creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsListQueryUpdatedAt - Only return results whose provider deployment last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsListQuery represents the provider deployments list query type.
type ProviderDeploymentsListQuery struct {
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
	// Id - Filter by deployment ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by version ID(s)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ActorId - Filter by actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// IdentityCredentialId - Filter by identity credential ID(s)
	IdentityCredentialId *any `json:"identity_credential_id,omitempty"`
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider deployment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider deployment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsListQuery.
func MapProviderDeploymentsListQueryFromJSON(data []byte) (*ProviderDeploymentsListQuery, error) {
	var v ProviderDeploymentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsListQueryToJSON serializes a ProviderDeploymentsListQuery to JSON.
func MapProviderDeploymentsListQueryToJSON(v *ProviderDeploymentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
