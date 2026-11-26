package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersListOutputItemsConfig represents the integrations providers list output items config type.
type IntegrationsProvidersListOutputItemsConfig struct {
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

// IntegrationsProvidersListOutputItemsCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersListOutputItemsCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsProvidersListOutputItems represents the integrations providers list output items type.
type IntegrationsProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// IntegrationId - Integration this provider belongs to.
	IntegrationId string `json:"integration_id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsProvidersListOutputItemsToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                     `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersListOutputItemsConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsProvidersListOutputItemsCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsProvidersListOutputPagination represents the integrations providers list output pagination type.
type IntegrationsProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsProvidersListOutput represents the integrations providers list output type.
type IntegrationsProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []IntegrationsProvidersListOutputItems    `json:"items"`
	Pagination IntegrationsProvidersListOutputPagination `json:"pagination"`
}

// MapIntegrationsProvidersListOutputFromJSON deserializes JSON data into a IntegrationsProvidersListOutput.
func MapIntegrationsProvidersListOutputFromJSON(data []byte) (*IntegrationsProvidersListOutput, error) {
	var v IntegrationsProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersListOutputToJSON serializes a IntegrationsProvidersListOutput to JSON.
func MapIntegrationsProvidersListOutputToJSON(v *IntegrationsProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsProvidersListQueryCreatedAt - Only return results whose integration provider creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsProvidersListQueryUpdatedAt - Only return results whose integration provider last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsProvidersListQuery represents the integrations providers list query type.
type IntegrationsProvidersListQuery struct {
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
	// Search - Search integration providers by name, description, ID, or provider.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// ProviderId - Only return integration providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Only return integration providers using one of these auth methods.
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Only return integration providers using one of these credentials.
	ProviderAuthCredentialsId *any `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Only return integration providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// CreatedAt - Only return results whose integration provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsProvidersListQueryFromJSON deserializes JSON data into a IntegrationsProvidersListQuery.
func MapIntegrationsProvidersListQueryFromJSON(data []byte) (*IntegrationsProvidersListQuery, error) {
	var v IntegrationsProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersListQueryToJSON serializes a IntegrationsProvidersListQuery to JSON.
func MapIntegrationsProvidersListQueryToJSON(v *IntegrationsProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
