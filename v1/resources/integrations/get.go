package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsGetOutputConfiguration represents the integrations get output configuration type.
type IntegrationsGetOutputConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// IntegrationsGetOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsGetOutputImplementation struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
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

// IntegrationsGetOutputProviders represents the integrations get output providers type.
type IntegrationsGetOutputProviders struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsGetOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                    `json:"provider_id"`
	DeploymentId      string                                    `json:"deployment_id"`
	AuthMethodId      *string                                   `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                   `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsGetOutputProvidersConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                 `json:"created_at"`
	UpdatedAt         time.Time                                 `json:"updated_at"`
	ArchivedAt        *time.Time                                `json:"archived_at,omitempty"`
}

// IntegrationsGetOutput represents the integrations get output type.
type IntegrationsGetOutput struct {
	Object         string                               `json:"object"`
	Id             string                               `json:"id"`
	Status         string                               `json:"status"`
	Slug           string                               `json:"slug"`
	Name           string                               `json:"name"`
	Description    *string                              `json:"description,omitempty"`
	Metadata       *map[string]any                      `json:"metadata,omitempty"`
	Configuration  IntegrationsGetOutputConfiguration   `json:"configuration"`
	Implementation *IntegrationsGetOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsGetOutputProviders     `json:"providers"`
	CreatedAt      time.Time                            `json:"created_at"`
	UpdatedAt      time.Time                            `json:"updated_at"`
	ArchivedAt     *time.Time                           `json:"archived_at,omitempty"`
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
