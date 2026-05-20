package integrations

import (
	"encoding/json"
	"time"
)

// IntegrationsUpdateOutputConfiguration represents the integrations update output configuration type.
type IntegrationsUpdateOutputConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// IntegrationsUpdateOutputImplementation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsUpdateOutputImplementation struct {
	Type               *string `json:"type,omitempty"`
	ProviderTemplateId *string `json:"provider_template_id,omitempty"`
	MagicMcpServerId   *string `json:"magic_mcp_server_id,omitempty"`
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

// IntegrationsUpdateOutputProviders represents the integrations update output providers type.
type IntegrationsUpdateOutputProviders struct {
	Object        string          `json:"object"`
	Id            string          `json:"id"`
	Status        string          `json:"status"`
	IntegrationId string          `json:"integration_id"`
	Name          string          `json:"name"`
	Description   *string         `json:"description,omitempty"`
	Metadata      *map[string]any `json:"metadata,omitempty"`
	// ToolFilter - Tool filter configuration
	ToolFilter        *IntegrationsUpdateOutputProvidersToolFilter `json:"tool_filter,omitempty"`
	ProviderId        string                                       `json:"provider_id"`
	DeploymentId      string                                       `json:"deployment_id"`
	AuthMethodId      *string                                      `json:"auth_method_id,omitempty"`
	AuthCredentialsId *string                                      `json:"auth_credentials_id,omitempty"`
	Config            *IntegrationsUpdateOutputProvidersConfig     `json:"config,omitempty"`
	CreatedAt         time.Time                                    `json:"created_at"`
	UpdatedAt         time.Time                                    `json:"updated_at"`
	ArchivedAt        *time.Time                                   `json:"archived_at,omitempty"`
}

// IntegrationsUpdateOutput represents the integrations update output type.
type IntegrationsUpdateOutput struct {
	Object         string                                  `json:"object"`
	Id             string                                  `json:"id"`
	Status         string                                  `json:"status"`
	Slug           string                                  `json:"slug"`
	Name           string                                  `json:"name"`
	Description    *string                                 `json:"description,omitempty"`
	Metadata       *map[string]any                         `json:"metadata,omitempty"`
	Configuration  IntegrationsUpdateOutputConfiguration   `json:"configuration"`
	Implementation *IntegrationsUpdateOutputImplementation `json:"implementation,omitempty"`
	Providers      []IntegrationsUpdateOutputProviders     `json:"providers"`
	CreatedAt      time.Time                               `json:"created_at"`
	UpdatedAt      time.Time                               `json:"updated_at"`
	ArchivedAt     *time.Time                              `json:"archived_at,omitempty"`
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
	Name                          *string         `json:"name,omitempty"`
	Description                   *string         `json:"description,omitempty"`
	Metadata                      *map[string]any `json:"metadata,omitempty"`
	CanAttachCustomToolFilters    *bool           `json:"can_attach_custom_tool_filters,omitempty"`
	CanAttachCustomProviderConfig *bool           `json:"can_attach_custom_provider_config,omitempty"`
	CanOverrideToolFilters        *bool           `json:"can_override_tool_filters,omitempty"`
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
