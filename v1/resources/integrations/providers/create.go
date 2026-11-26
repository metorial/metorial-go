package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsProvidersCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersCreateOutputConfig represents the integrations providers create output config type.
type IntegrationsProvidersCreateOutputConfig struct {
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

// IntegrationsProvidersCreateOutputCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersCreateOutputCallbacks struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsProvidersCreateOutput represents the integrations providers create output type.
type IntegrationsProvidersCreateOutput struct {
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
	ToolFilter *IntegrationsProvidersCreateOutputToolFilter `json:"tool_filter,omitempty"`
	// ProviderId - Provider added to the integration.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Deployment the provider runs with.
	DeploymentId string `json:"deployment_id"`
	// AuthMethodId - Auth method the provider runs with. `null` when none is set.
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - Credentials the provider runs with. `null` when none are set.
	AuthCredentialsId *string                                  `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersCreateOutputConfig `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks IntegrationsProvidersCreateOutputCallbacks `json:"callbacks"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapIntegrationsProvidersCreateOutputFromJSON deserializes JSON data into a IntegrationsProvidersCreateOutput.
func MapIntegrationsProvidersCreateOutputFromJSON(data []byte) (*IntegrationsProvidersCreateOutput, error) {
	var v IntegrationsProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersCreateOutputToJSON serializes a IntegrationsProvidersCreateOutput to JSON.
func MapIntegrationsProvidersCreateOutputToJSON(v *IntegrationsProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsProvidersCreateBodyCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersCreateBodyCallbacks struct {
	// Status - Enabling registers a callback against every matching integration instance. Disabling tears it down again. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
}

// IntegrationsProvidersCreateBody represents the integrations providers create body type.
type IntegrationsProvidersCreateBody struct {
	// IntegrationId - Integration the provider belongs to.
	IntegrationId string `json:"integration_id"`
	// ProviderId - Provider to add to the integration.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the provider runs with.
	ProviderDeploymentId string `json:"provider_deployment_id"`
	// ProviderAuthMethodId - Auth method the provider runs with. Omitted leaves it unset.
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Credentials the provider runs with. Omitted leaves them unset.
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config the provider runs with. Omitted leaves it unset.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// Name - Display name for the integration provider. Defaults to the provider name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the integration provider.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks *IntegrationsProvidersCreateBodyCallbacks `json:"callbacks,omitempty"`
}

// MapIntegrationsProvidersCreateBodyFromJSON deserializes JSON data into a IntegrationsProvidersCreateBody.
func MapIntegrationsProvidersCreateBodyFromJSON(data []byte) (*IntegrationsProvidersCreateBody, error) {
	var v IntegrationsProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersCreateBodyToJSON serializes a IntegrationsProvidersCreateBody to JSON.
func MapIntegrationsProvidersCreateBodyToJSON(v *IntegrationsProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
