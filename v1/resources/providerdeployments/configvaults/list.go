package configvaults

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigVaultsListOutputItemsDeployment represents the provider deployments config vaults list output items deployment type.
type ProviderDeploymentsConfigVaultsListOutputItemsDeployment struct {
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

// ProviderDeploymentsConfigVaultsListOutputItems represents the provider deployments config vaults list output items type.
type ProviderDeploymentsConfigVaultsListOutputItems struct {
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
	ProviderId string                                                    `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigVaultsListOutputItemsDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigVaultsListOutputPagination represents the provider deployments config vaults list output pagination type.
type ProviderDeploymentsConfigVaultsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderDeploymentsConfigVaultsListOutput represents the provider deployments config vaults list output type.
type ProviderDeploymentsConfigVaultsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                              `json:"object"`
	Items      []ProviderDeploymentsConfigVaultsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsConfigVaultsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsConfigVaultsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsListOutput.
func MapProviderDeploymentsConfigVaultsListOutputFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsListOutput, error) {
	var v ProviderDeploymentsConfigVaultsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsListOutputToJSON serializes a ProviderDeploymentsConfigVaultsListOutput to JSON.
func MapProviderDeploymentsConfigVaultsListOutputToJSON(v *ProviderDeploymentsConfigVaultsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigVaultsListQueryCreatedAt - Only return results whose provider config vault creation time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsConfigVaultsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsConfigVaultsListQueryUpdatedAt - Only return results whose provider config vault last update time is in this range. Set `gt`, `lt`, or both.
type ProviderDeploymentsConfigVaultsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderDeploymentsConfigVaultsListQuery represents the provider deployments config vaults list query type.
type ProviderDeploymentsConfigVaultsListQuery struct {
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
	// Status - Filter by status (active, archived)
	Status *any `json:"status,omitempty"`
	// Id - Filter by config vault ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderConfigVaultId - Filter by config vault ID(s)
	ProviderConfigVaultId *any `json:"provider_config_vault_id,omitempty"`
	// Search - Search by name or description
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider config vault creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderDeploymentsConfigVaultsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider config vault last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderDeploymentsConfigVaultsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderDeploymentsConfigVaultsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsListQuery.
func MapProviderDeploymentsConfigVaultsListQueryFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsListQuery, error) {
	var v ProviderDeploymentsConfigVaultsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsListQueryToJSON serializes a ProviderDeploymentsConfigVaultsListQuery to JSON.
func MapProviderDeploymentsConfigVaultsListQueryToJSON(v *ProviderDeploymentsConfigVaultsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
