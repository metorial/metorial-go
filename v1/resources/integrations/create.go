package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsCreateOutputConfiguration represents the integrations create output configuration type.
type IntegrationsCreateOutputConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// IntegrationsCreateOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsCreateOutputImplementation struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
}

// IntegrationsCreateOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsCreateOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsCreateOutputProvidersConfig represents the integrations create output providers config type.
type IntegrationsCreateOutputProvidersConfig struct {
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

// IntegrationsCreateOutputProviders represents the integrations create output providers type.
type IntegrationsCreateOutputProviders struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsCreateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsCreateOutputProvidersConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                   `json:"archived_at,omitempty"`
}

// IntegrationsCreateOutput represents the integrations create output type.
type IntegrationsCreateOutput struct {
	Object         string                                  `json:"object"`
	Id             string                                  `json:"id"`
	Status         string                                  `json:"status"`
	Slug           string                                  `json:"slug"`
	Name           string                                  `json:"name"`
	Description    *string                                 `json:"description,omitempty"`
	Metadata       *map[string]any                         `json:"metadata,omitempty"`
	Configuration  IntegrationsCreateOutputConfiguration   `json:"configuration"`
	Implementation *IntegrationsCreateOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsCreateOutputProviders     `json:"providers"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
	ArchivedAt     *time.Time                              `json:"archived_at,omitempty"`
}

// MapIntegrationsCreateOutputFromJSON deserializes JSON data into a IntegrationsCreateOutput.
func MapIntegrationsCreateOutputFromJSON(data []byte) (*IntegrationsCreateOutput, error) {
	var v IntegrationsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsCreateOutputToJSON serializes a IntegrationsCreateOutput to JSON.
func MapIntegrationsCreateOutputToJSON(v *IntegrationsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsCreateBody represents the integrations create body type.
type IntegrationsCreateBody struct {
	Name                          string          `json:"name"`
	Description                   *string         `json:"description,omitempty"`
	Metadata                      *map[string]any `json:"metadata,omitempty"`
	CanAttachCustomToolFilters    *bool           `json:"can_attach_custom_tool_filters,omitempty"`
	CanAttachCustomProviderConfig *bool           `json:"can_attach_custom_provider_config,omitempty"`
	CanOverrideToolFilters        *bool           `json:"can_override_tool_filters,omitempty"`
}

// MapIntegrationsCreateBodyFromJSON deserializes JSON data into a IntegrationsCreateBody.
func MapIntegrationsCreateBodyFromJSON(data []byte) (*IntegrationsCreateBody, error) {
	var v IntegrationsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsCreateBodyToJSON serializes a IntegrationsCreateBody to JSON.
func MapIntegrationsCreateBodyToJSON(v *IntegrationsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
