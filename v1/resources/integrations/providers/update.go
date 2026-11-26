package providers

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// IntegrationsProvidersUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersUpdateOutputConfig represents the integrations providers update output config type.
type IntegrationsProvidersUpdateOutputConfig struct {
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

// IntegrationsProvidersUpdateOutputCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersUpdateOutputCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsProvidersUpdateOutput represents the integrations providers update output type.
type IntegrationsProvidersUpdateOutput struct {
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
	ToolFilter *IntegrationsProvidersUpdateOutputToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersUpdateOutputConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsProvidersUpdateOutputCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsProvidersUpdateOutputFromJSON deserializes JSON data into a IntegrationsProvidersUpdateOutput.
func MapIntegrationsProvidersUpdateOutputFromJSON(data []byte) (*IntegrationsProvidersUpdateOutput, error) {
	var v IntegrationsProvidersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersUpdateOutputToJSON serializes a IntegrationsProvidersUpdateOutput to JSON.
func MapIntegrationsProvidersUpdateOutputToJSON(v *IntegrationsProvidersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsProvidersUpdateBodyCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersUpdateBodyCallbacks struct {
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
}

// IntegrationsProvidersUpdateBody represents the integrations providers update body type.
type IntegrationsProvidersUpdateBody struct {
	// ProviderDeploymentId - Deployment the provider runs with. Omitted keeps the current one.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Auth method the provider runs with. `null` clears it. Omitted keeps the current one.
	ProviderAuthMethodId *nullable.Nullable[string] `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Credentials the provider runs with. `null` clears them. Omitted keeps the current ones.
	ProviderAuthCredentialsId *nullable.Nullable[string] `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config the provider runs with. `null` clears it. Omitted keeps the current one.
	ProviderConfigId *nullable.Nullable[string] `json:"provider_config_id,omitempty"`
	// Name - Display name for the integration provider. Omitted keeps the current one.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration provider. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks *IntegrationsProvidersUpdateBodyCallbacks `json:"callbacks,omitempty"`
}

// MapIntegrationsProvidersUpdateBodyFromJSON deserializes JSON data into a IntegrationsProvidersUpdateBody.
func MapIntegrationsProvidersUpdateBodyFromJSON(data []byte) (*IntegrationsProvidersUpdateBody, error) {
	var v IntegrationsProvidersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersUpdateBodyToJSON serializes a IntegrationsProvidersUpdateBody to JSON.
func MapIntegrationsProvidersUpdateBodyToJSON(v *IntegrationsProvidersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
