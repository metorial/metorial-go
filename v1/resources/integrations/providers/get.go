package providers

import (
	"encoding/json"
	"time"
)

// IntegrationsProvidersGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsProvidersGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsProvidersGetOutputConfig represents the integrations providers get output config type.
type IntegrationsProvidersGetOutputConfig struct {
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

// IntegrationsProvidersGetOutput represents the integrations providers get output type.
type IntegrationsProvidersGetOutput struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsProvidersGetOutputToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                    `json:"provider_id"`
	DeploymentId      string                                    `json:"deployment_id"`
	AuthMethodId      *string                                   `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                   `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsProvidersGetOutputConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                 `json:"created_at"`
	UpdatedAt         time.Time                                 `json:"updated_at"`
	ArchivedAt        *time.Time                                `json:"archived_at,omitempty"`
}

// MapIntegrationsProvidersGetOutputFromJSON deserializes JSON data into a IntegrationsProvidersGetOutput.
func MapIntegrationsProvidersGetOutputFromJSON(data []byte) (*IntegrationsProvidersGetOutput, error) {
	var v IntegrationsProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsProvidersGetOutputToJSON serializes a IntegrationsProvidersGetOutput to JSON.
func MapIntegrationsProvidersGetOutputToJSON(v *IntegrationsProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
