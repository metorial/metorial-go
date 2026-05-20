package instancegroups

import (
	"encoding/json"
	"time"
)

// IntegrationsInstanceGroupsCreateSessionOutputUsage represents the integrations instance groups create session output usage type.
type IntegrationsInstanceGroupsCreateSessionOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage - Usage statistics
type IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment represents the integrations instance groups create session output providers deployment type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Deployment ID
	Id string `json:"id"`
	// IsDefault - Whether this is the default deployment
	IsDefault bool `json:"is_default"`
	// Name - Deployment name
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

// IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig represents the integrations instance groups create session output providers config type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig struct {
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

// IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig represents the integrations instance groups create session output providers auth config type.
type IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig struct {
	Object string `json:"object"`
	Id     string `json:"id"`
}

// IntegrationsInstanceGroupsCreateSessionOutputProviders represents the integrations instance groups create session output providers type.
type IntegrationsInstanceGroupsCreateSessionOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session provider identifier
	Id string `json:"id"`
	// Status - Provider status
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage IntegrationsInstanceGroupsCreateSessionOutputProvidersUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter IntegrationsInstanceGroupsCreateSessionOutputProvidersToolFilter `json:"tool_filter"`
	// ProviderId - Provider ID
	ProviderId string `json:"provider_id"`
	// SessionId - Parent session ID
	SessionId string `json:"session_id"`
	// FromTemplateId - Source template ID
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Source template provider ID
	FromTemplateProviderId *string                                                           `json:"from_template_provider_id,omitempty"`
	Deployment             IntegrationsInstanceGroupsCreateSessionOutputProvidersDeployment  `json:"deployment"`
	Config                 IntegrationsInstanceGroupsCreateSessionOutputProvidersConfig      `json:"config"`
	AuthConfig             *IntegrationsInstanceGroupsCreateSessionOutputProvidersAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationsInstanceGroupsCreateSessionOutput represents the integrations instance groups create session output type.
type IntegrationsInstanceGroupsCreateSessionOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique session identifier
	Id string `json:"id"`
	// Status - Session status
	Status string `json:"status"`
	// Name - Display name
	Name *string `json:"name,omitempty"`
	// Description - Description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional information
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ConnectionState - Session connection state
	ConnectionState string `json:"connection_state"`
	// ConnectionUrl - MCP connection URL for this session
	ConnectionUrl string `json:"connection_url"`
	// ClientSecret - Session-scoped fine grained client secret token
	ClientSecret *string                                            `json:"client_secret,omitempty"`
	Usage        IntegrationsInstanceGroupsCreateSessionOutputUsage `json:"usage"`
	// Providers - Session providers
	Providers []IntegrationsInstanceGroupsCreateSessionOutputProviders `json:"providers"`
	// FromTemplatesIds - Template IDs this session was created from
	FromTemplatesIds []string `json:"from_templates_ids"`
	// HasErrors - Whether the session has any errors
	HasErrors bool `json:"has_errors"`
	// HasWarnings - Whether the session has any warnings
	HasWarnings bool `json:"has_warnings"`
	// CreatedAt - Timestamp when created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIntegrationsInstanceGroupsCreateSessionOutputFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionOutput.
func MapIntegrationsInstanceGroupsCreateSessionOutputFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionOutput, error) {
	var v IntegrationsInstanceGroupsCreateSessionOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionOutputToJSON serializes a IntegrationsInstanceGroupsCreateSessionOutput to JSON.
func MapIntegrationsInstanceGroupsCreateSessionOutputToJSON(v *IntegrationsInstanceGroupsCreateSessionOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationsInstanceGroupsCreateSessionBody represents the integrations instance groups create session body type.
type IntegrationsInstanceGroupsCreateSessionBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MapIntegrationsInstanceGroupsCreateSessionBodyFromJSON deserializes JSON data into a IntegrationsInstanceGroupsCreateSessionBody.
func MapIntegrationsInstanceGroupsCreateSessionBodyFromJSON(data []byte) (*IntegrationsInstanceGroupsCreateSessionBody, error) {
	var v IntegrationsInstanceGroupsCreateSessionBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationsInstanceGroupsCreateSessionBodyToJSON serializes a IntegrationsInstanceGroupsCreateSessionBody to JSON.
func MapIntegrationsInstanceGroupsCreateSessionBodyToJSON(v *IntegrationsInstanceGroupsCreateSessionBody) ([]byte, error) {
	return json.Marshal(v)
}
