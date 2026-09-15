package providers

import (
	"encoding/json"
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
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Config ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default config
	IsDefault bool `json:"is_default"`
	// Name - Config name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsProvidersUpdateOutputCallbacks - Provider callback state for this integration provider
type IntegrationsProvidersUpdateOutputCallbacks struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Status - Whether this integration provider should receive provider callbacks. Enabling it creates a callback and registers it against every matching integration instance.
	Status string `json:"status"`
	// CallbackId - The active callback, once it has been created. Null while callbacks are disabled.
	CallbackId *string `json:"callback_id,omitempty"`
}

// IntegrationsProvidersUpdateOutput represents the integrations providers update output type.
type IntegrationsProvidersUpdateOutput struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsProvidersUpdateOutputToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersUpdateOutputConfig     `json:"config,omitempty"`
	// Callbacks - Provider callback state for this integration provider
	Callbacks  IntegrationsProvidersUpdateOutputCallbacks `json:"callbacks"`
	CreatedAt  time.Time                                  `json:"created_at"`
	UpdatedAt  time.Time                                  `json:"updated_at"`
	ArchivedAt *time.Time                                 `json:"archived_at,omitempty"`
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
	// Status - Enabling creates a callback and registers it against every matching integration instance. Disabling tears down the callback and every callback instance registered for it. Only providers whose type reports `triggers.status` as `enabled` can enable callbacks.
	Status string `json:"status"`
}

// IntegrationsProvidersUpdateBody represents the integrations providers update body type.
type IntegrationsProvidersUpdateBody struct {
	ProviderDeploymentId      *string         `json:"provider_deployment_id,omitempty"`
	ProviderAuthMethodId      *string         `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *string         `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *string         `json:"provider_config_id,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	Description               *string         `json:"description,omitempty"`
	Metadata                  *map[string]any `json:"metadata,omitempty"`
	ToolFilters               *any            `json:"tool_filters,omitempty"`
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
