package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsProvidersDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersDeleteOutputConfig represents the integrations providers delete output config type.
type IntegrationsProvidersDeleteOutputConfig struct {
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

// IntegrationsProvidersDeleteOutput represents the integrations providers delete output type.
type IntegrationsProvidersDeleteOutput struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsProvidersDeleteOutputToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersDeleteOutputConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                   `json:"archived_at,omitempty"`
}

// MapIntegrationsProvidersDeleteOutputFromJSON deserializes JSON data into a IntegrationsProvidersDeleteOutput.
func MapIntegrationsProvidersDeleteOutputFromJSON(data []byte) (*IntegrationsProvidersDeleteOutput, error) {
	var v IntegrationsProvidersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersDeleteOutputToJSON serializes a IntegrationsProvidersDeleteOutput to JSON.
func MapIntegrationsProvidersDeleteOutputToJSON(v *IntegrationsProvidersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
