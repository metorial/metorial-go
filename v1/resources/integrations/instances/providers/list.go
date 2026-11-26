package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsInstancesProvidersListOutputItemsToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersListOutputItemsToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsProvider represents the integrations instances providers list output items provider type.
type IntegrationsInstancesProvidersListOutputItemsProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion - Version this snapshot was taken from.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Index - Position of this version, starting at 0.
	Index float64 `json:"index"`
}

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig represents the integrations instances providers list output items integration provider config type.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig struct {
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

// IntegrationsInstancesProvidersListOutputItemsIntegrationProvider represents the integrations instances providers list output items integration provider type.
type IntegrationsInstancesProvidersListOutputItemsIntegrationProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderVersion - Version this snapshot was taken from.
	ProviderVersion IntegrationsInstancesProvidersListOutputItemsIntegrationProviderProviderVersion `json:"provider_version"`
	// Status - `archived` once you archive the integration provider.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesProvidersListOutputItemsIntegrationProviderToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider this snapshot was taken from.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment this snapshot was taken from.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method at snapshot time. `null` when none was set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials at snapshot time. `null` when none were set.
	AuthCredentialsId *string                                                                 `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsInstancesProvidersListOutputItemsIntegrationProviderConfig `json:"config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersListOutputItemsConfig represents the integrations instances providers list output items config type.
type IntegrationsInstancesProvidersListOutputItemsConfig struct {
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

// IntegrationsInstancesProvidersListOutputItemsAuthConfig represents the integrations instances providers list output items auth config type.
type IntegrationsInstancesProvidersListOutputItemsAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstancesProvidersListOutputItems represents the integrations instances providers list output items type.
type IntegrationsInstancesProvidersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the entry is archived.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration the instance belongs to.
	IntegrationId string `json:"integration_id"`
	// IntegrationInstanceId - Integration instance this entry belongs to.
	IntegrationInstanceId string `json:"integration_instance_id"`
	// ToolFilter - Tool filter configuration
	ToolFilter *IntegrationsInstancesProvidersListOutputItemsToolFilter `json:"tool_filter,omitempty"`
	// IsOverrideToolFilter - Replaces the integration tool filters instead of combining with them.
	IsOverrideToolFilter bool                                                             `json:"is_override_tool_filter"`
	Provider             IntegrationsInstancesProvidersListOutputItemsProvider            `json:"provider"`
	IntegrationProvider  IntegrationsInstancesProvidersListOutputItemsIntegrationProvider `json:"integration_provider"`
	Config               *IntegrationsInstancesProvidersListOutputItemsConfig             `json:"config,omitempty"`
	AuthConfig           *IntegrationsInstancesProvidersListOutputItemsAuthConfig         `json:"auth_config,omitempty"`
	// CallbackInstanceId - The active callback instance registered for this integration instance provider, if its integration provider has callbacks enabled.
	CallbackInstanceId *string `json:"callback_instance_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsInstancesProvidersListOutputPagination represents the integrations instances providers list output pagination type.
type IntegrationsInstancesProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsInstancesProvidersListOutput represents the integrations instances providers list output type.
type IntegrationsInstancesProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                             `json:"object"`
	Items      []IntegrationsInstancesProvidersListOutputItems    `json:"items"`
	Pagination IntegrationsInstancesProvidersListOutputPagination `json:"pagination"`
}

// MapIntegrationsInstancesProvidersListOutputFromJSON deserializes JSON data into a IntegrationsInstancesProvidersListOutput.
func MapIntegrationsInstancesProvidersListOutputFromJSON(data []byte) (*IntegrationsInstancesProvidersListOutput, error) {
	var v IntegrationsInstancesProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersListOutputToJSON serializes a IntegrationsInstancesProvidersListOutput to JSON.
func MapIntegrationsInstancesProvidersListOutputToJSON(v *IntegrationsInstancesProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstancesProvidersListQueryCreatedAt - Only return results whose integration instance provider creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstancesProvidersListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesProvidersListQueryUpdatedAt - Only return results whose integration instance provider last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsInstancesProvidersListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsInstancesProvidersListQuery represents the integrations instances providers list query type.
type IntegrationsInstancesProvidersListQuery struct {
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
	// Search - Search integration instance providers by name, description, or ID.
	Search *string `json:"search,omitempty"`
	// Status - Only return integration instance providers with one of these statuses. `deleted` providers are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integration instance providers with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return integration instance providers for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return integration instance providers for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// ProviderId - Only return integration instance providers using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integration instance providers using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderDeploymentId - Only return integration instance providers running on one of these deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return integration instance providers using one of these configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return integration instance providers using one of these credentials.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// SessionTemplateId - Only return integration instance providers using one of these session templates.
	SessionTemplateId *any `json:"session_template_id,omitempty"`
	// CreatedAt - Only return results whose integration instance provider creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsInstancesProvidersListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration instance provider last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsInstancesProvidersListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsInstancesProvidersListQueryFromJSON deserializes JSON data into a IntegrationsInstancesProvidersListQuery.
func MapIntegrationsInstancesProvidersListQueryFromJSON(data []byte) (*IntegrationsInstancesProvidersListQuery, error) {
	var v IntegrationsInstancesProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstancesProvidersListQueryToJSON serializes a IntegrationsInstancesProvidersListQuery to JSON.
func MapIntegrationsInstancesProvidersListQueryToJSON(v *IntegrationsInstancesProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
