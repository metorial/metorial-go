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

// IntegrationsProvidersCreateOutput represents the integrations providers create output type.
type IntegrationsProvidersCreateOutput struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsProvidersCreateOutputToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersCreateOutputConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                   `json:"archived_at,omitempty"`
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

// IntegrationsProvidersCreateBody represents the integrations providers create body type.
type IntegrationsProvidersCreateBody struct {
	IntegrationId             string          `json:"integration_id"`
	ProviderId                string          `json:"provider_id"`
	ProviderDeploymentId      string          `json:"provider_deployment_id"`
	ProviderAuthMethodId      *string         `json:"provider_auth_method_id,omitempty"`
	ProviderAuthCredentialsId *string         `json:"provider_auth_credentials_id,omitempty"`
	ProviderConfigId          *string         `json:"provider_config_id,omitempty"`
	Name                      *string         `json:"name,omitempty"`
	Description               *string         `json:"description,omitempty"`
	Metadata                  *map[string]any `json:"metadata,omitempty"`
	ToolFilters               *any            `json:"tool_filters,omitempty"`
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
