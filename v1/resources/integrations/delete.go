package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsDeleteOutputConfiguration represents the integrations delete output configuration type.
type IntegrationsDeleteOutputConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// IntegrationsDeleteOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsDeleteOutputImplementation struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
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

// IntegrationsDeleteOutputProviders represents the integrations delete output providers type.
type IntegrationsDeleteOutputProviders struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsDeleteOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsDeleteOutputProvidersConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                   `json:"archived_at,omitempty"`
}

// IntegrationsDeleteOutput represents the integrations delete output type.
type IntegrationsDeleteOutput struct {
	Object         string                                  `json:"object"`
	Id             string                                  `json:"id"`
	Status         string                                  `json:"status"`
	Slug           string                                  `json:"slug"`
	Name           string                                  `json:"name"`
	Description    *string                                 `json:"description,omitempty"`
	Metadata       *map[string]any                         `json:"metadata,omitempty"`
	Configuration  IntegrationsDeleteOutputConfiguration   `json:"configuration"`
	Implementation *IntegrationsDeleteOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsDeleteOutputProviders     `json:"providers"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
	ArchivedAt     *time.Time                              `json:"archived_at,omitempty"`
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
