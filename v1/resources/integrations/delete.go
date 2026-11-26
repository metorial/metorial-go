package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsDeleteOutputConfiguration - Which settings integration instances may customize.
type IntegrationsDeleteOutputConfiguration struct {
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

// IntegrationsDeleteOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsDeleteOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template this integration is built on.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server this integration is built on.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsDeleteOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsDeleteOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsDeleteOutputProvidersConfig represents the integrations delete output providers config type.
type IntegrationsDeleteOutputProvidersConfig struct {
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

// IntegrationsDeleteOutputProvidersCallbacks - Provider callback state for this integration provider
type IntegrationsDeleteOutputProvidersCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsDeleteOutputProviders represents the integrations delete output providers type.
type IntegrationsDeleteOutputProviders struct {
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
	ToolFilter *IntegrationsDeleteOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsDeleteOutputProvidersConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsDeleteOutputProvidersCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsDeleteOutput represents the integrations delete output type.
type IntegrationsDeleteOutput struct {
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
	Configuration IntegrationsDeleteOutputConfiguration `json:"configuration"`
	// Implementation - What this integration is built on. `null` for a standalone integration.
	Implementation *IntegrationsDeleteOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers in this integration, with the deployment, credentials, and tool filters each runs with.
	Providers []IntegrationsDeleteOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsDeleteOutputFromJSON deserializes JSON data into a IntegrationsDeleteOutput.
func MapIntegrationsDeleteOutputFromJSON(data []byte) (*IntegrationsDeleteOutput, error) {
	var v IntegrationsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsDeleteOutputToJSON serializes a IntegrationsDeleteOutput to JSON.
func MapIntegrationsDeleteOutputToJSON(v *IntegrationsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
