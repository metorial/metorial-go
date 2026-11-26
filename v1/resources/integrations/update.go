package integrations

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// IntegrationsUpdateOutputConfiguration - Which settings integration instances may customize.
type IntegrationsUpdateOutputConfiguration struct {
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

// IntegrationsUpdateOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsUpdateOutputImplementation struct {
	// Type - Which kind of backing this is.
	Type *string `json:"type,omitempty"`
	// ProviderTemplateId - Provider template this integration is built on.
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	// MagicMcpServerId - Magic MCP server this integration is built on.
	MagicMcpServerId *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsUpdateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsUpdateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsUpdateOutputProvidersConfig represents the integrations update output providers config type.
type IntegrationsUpdateOutputProvidersConfig struct {
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

// IntegrationsUpdateOutputProvidersCallbacks - Provider callback state for this integration provider
type IntegrationsUpdateOutputProvidersCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsUpdateOutputProviders represents the integrations update output providers type.
type IntegrationsUpdateOutputProviders struct {
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
	ToolFilter *IntegrationsUpdateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsUpdateOutputProvidersConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsUpdateOutputProvidersCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// IntegrationsUpdateOutput represents the integrations update output type.
type IntegrationsUpdateOutput struct {
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
	Configuration IntegrationsUpdateOutputConfiguration `json:"configuration"`
	// Implementation - What this integration is built on. `null` for a standalone integration.
	Implementation *IntegrationsUpdateOutputImplementation `json:"implementation,omitempty"`
	// Providers - Providers in this integration, with the deployment, credentials, and tool filters each runs with.
	Providers []IntegrationsUpdateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsUpdateOutputFromJSON deserializes JSON data into a IntegrationsUpdateOutput.
func MapIntegrationsUpdateOutputFromJSON(data []byte) (*IntegrationsUpdateOutput, error) {
	var v IntegrationsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsUpdateOutputToJSON serializes a IntegrationsUpdateOutput to JSON.
func MapIntegrationsUpdateOutputToJSON(v *IntegrationsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsUpdateBody represents the integrations update body type.
type IntegrationsUpdateBody struct {
	// Name - Display name for the integration.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// UseIntegrationNameInToolNames - Names session providers after the integration instead of the provider. `null` uses the default.
	UseIntegrationNameInToolNames *nullable.Nullable[bool] `json:"use_integration_name_in_tool_names,omitempty"`
	// CanAttachCustomToolFilters - Lets integration instances attach their own tool filters.
	CanAttachCustomToolFilters *bool `json:"can_attach_custom_tool_filters,omitempty"`
	// CanAttachCustomProviderConfig - Lets integration instances use their own provider configs.
	CanAttachCustomProviderConfig *bool `json:"can_attach_custom_provider_config,omitempty"`
	// CanOverrideToolFilters - Lets integration instances ignore the integration tool filters.
	CanOverrideToolFilters *bool `json:"can_override_tool_filters,omitempty"`
	// EnableCallbackTools - Makes callback event tools available for this integration.
	EnableCallbackTools *bool `json:"enable_callback_tools,omitempty"`
}

// MapIntegrationsUpdateBodyFromJSON deserializes JSON data into a IntegrationsUpdateBody.
func MapIntegrationsUpdateBodyFromJSON(data []byte) (*IntegrationsUpdateBody, error) {
	var v IntegrationsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsUpdateBodyToJSON serializes a IntegrationsUpdateBody to JSON.
func MapIntegrationsUpdateBodyToJSON(v *IntegrationsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
