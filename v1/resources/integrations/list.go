package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsListOutputItemsConfiguration - Which settings integration instances may customize.
type IntegrationsListOutputItemsConfiguration struct {
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters bool `json:"can_attach_custom_tool_filters"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters bool `json:"can_override_tool_filters"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools bool `json:"enable_callback_tools"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *bool `json:"use_integration_name_in_tool_names,omitempty"`
}

// IntegrationsListOutputItemsImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsListOutputItemsImplementation struct {
	// Type - Which kind of backing this is.
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template this integration is built on.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server this integration is built on.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsListOutputItemsProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsListOutputItemsProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsListOutputItemsProvidersConfig represents the integrations list output items providers config type.
type IntegrationsListOutputItemsProvidersConfig struct {
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

// IntegrationsListOutputItemsProvidersCallbacks - Provider callback state for this integration provider
type IntegrationsListOutputItemsProvidersCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsListOutputItemsProviders represents the integrations list output items providers type.
type IntegrationsListOutputItemsProviders struct {
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
	ToolFilter *IntegrationsListOutputItemsProvidersToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                     `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsListOutputItemsProvidersConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsListOutputItemsProvidersCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsListOutputItems represents the integrations list output items type.
type IntegrationsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the integration.
	Status string `json:"status"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Configuration - Which settings integration instances may customize.
	Configuration IntegrationsListOutputItemsConfiguration `json:"configuration"`
	// Implementation - What this integration is built on. `null` for a standalone integration.
	Implementation *IntegrationsListOutputItemsImplementation `json:"implementation,omitempty"`
	// Providers - Providers in this integration, with the deployment, credentials, and tool filters each runs with.
	Providers []IntegrationsListOutputItemsProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsListOutputPagination represents the integrations list output pagination type.
type IntegrationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IntegrationsListOutput represents the integrations list output type.
type IntegrationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
	Items      []IntegrationsListOutputItems    `json:"items"`
	Pagination IntegrationsListOutputPagination `json:"pagination"`
}

// MapIntegrationsListOutputFromJSON deserializes JSON data into a IntegrationsListOutput.
func MapIntegrationsListOutputFromJSON(data []byte) (*IntegrationsListOutput, error) {
	var v IntegrationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsListOutputToJSON serializes a IntegrationsListOutput to JSON.
func MapIntegrationsListOutputToJSON(v *IntegrationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsListQueryCreatedAt - Only return results whose integration creation time is in this range. Set `gt`, `lt`, or both.
type IntegrationsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsListQueryUpdatedAt - Only return results whose integration last update time is in this range. Set `gt`, `lt`, or both.
type IntegrationsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IntegrationsListQuery represents the integrations list query type.
type IntegrationsListQuery struct {
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
	// Search - Search integrations by name, description, ID, or slug.
	Search *string `json:"search,omitempty"`
	// Status - Only return integrations with one of these statuses. `deleted` integrations are never listed.
	Status *any `json:"status,omitempty"`
	// Id - Only return integrations with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ProviderId - Only return integrations using one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// IntegrationProviderId - Only return integrations using one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// CreatedAt - Only return results whose integration creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IntegrationsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose integration last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IntegrationsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIntegrationsListQueryFromJSON deserializes JSON data into a IntegrationsListQuery.
func MapIntegrationsListQueryFromJSON(data []byte) (*IntegrationsListQuery, error) {
	var v IntegrationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsListQueryToJSON serializes a IntegrationsListQuery to JSON.
func MapIntegrationsListQueryToJSON(v *IntegrationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
