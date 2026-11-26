package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsGetOutputConfiguration - Which settings integration instances may customize.
type IntegrationsGetOutputConfiguration struct {
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

// IntegrationsGetOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsGetOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template this integration is built on.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server this integration is built on.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsGetOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsGetOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsGetOutputProvidersConfig represents the integrations get output providers config type.
type IntegrationsGetOutputProvidersConfig struct {
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

// IntegrationsGetOutputProvidersCallbacks - Provider callback state for this integration provider
type IntegrationsGetOutputProvidersCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsGetOutputProviders represents the integrations get output providers type.
type IntegrationsGetOutputProviders struct {
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
	ToolFilter *IntegrationsGetOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                               `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsGetOutputProvidersConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsGetOutputProvidersCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsGetOutput represents the integrations get output type.
type IntegrationsGetOutput struct {
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
	Configuration IntegrationsGetOutputConfiguration `json:"configuration"`
	// Implementation - What this integration is built on. `null` for a standalone integration.
	Implementation *IntegrationsGetOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers in this integration, with the deployment, credentials, and tool filters each runs with.
	Providers []IntegrationsGetOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsGetOutputFromJSON deserializes JSON data into a IntegrationsGetOutput.
func MapIntegrationsGetOutputFromJSON(data []byte) (*IntegrationsGetOutput, error) {
	var v IntegrationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsGetOutputToJSON serializes a IntegrationsGetOutput to JSON.
func MapIntegrationsGetOutputToJSON(v *IntegrationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
