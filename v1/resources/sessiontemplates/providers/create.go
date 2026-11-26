package providers

import (
	"encoding/json"
	"time"
)

// SessionTemplatesProvidersCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesProvidersCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesProvidersCreateOutputDeployment represents the session templates providers create output deployment type.
type SessionTemplatesProvidersCreateOutputDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SessionTemplatesProvidersCreateOutputConfig represents the session templates providers create output config type.
type SessionTemplatesProvidersCreateOutputConfig struct {
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

// SessionTemplatesProvidersCreateOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesProvidersCreateOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesProvidersCreateOutput represents the session templates providers create output type.
type SessionTemplatesProvidersCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesProvidersCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                          `json:"session_template_id"`
	Deployment        SessionTemplatesProvidersCreateOutputDeployment `json:"deployment"`
	Config            SessionTemplatesProvidersCreateOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesProvidersCreateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionTemplatesProvidersCreateOutputFromJSON deserializes JSON data into a SessionTemplatesProvidersCreateOutput.
func MapSessionTemplatesProvidersCreateOutputFromJSON(data []byte) (*SessionTemplatesProvidersCreateOutput, error) {
	var v SessionTemplatesProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersCreateOutputToJSON serializes a SessionTemplatesProvidersCreateOutput to JSON.
func MapSessionTemplatesProvidersCreateOutputToJSON(v *SessionTemplatesProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionTemplatesProvidersCreateBody represents the session templates providers create body type.
type SessionTemplatesProvidersCreateBody struct {
	// SessionTemplateId - Session template to add the provider to.
	SessionTemplateId string `json:"session_template_id"`
	// ProviderDeploymentId - Deployment the template provider runs with.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Config the template provider runs with.
	ProviderConfigId      *string `json:"provider_config_id,omitempty"`
	ProviderConfigVaultId *string `json:"provider_config_vault_id,omitempty"`
	// ProviderAuthConfigId - Auth config the template provider runs with.
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapSessionTemplatesProvidersCreateBodyFromJSON deserializes JSON data into a SessionTemplatesProvidersCreateBody.
func MapSessionTemplatesProvidersCreateBodyFromJSON(data []byte) (*SessionTemplatesProvidersCreateBody, error) {
	var v SessionTemplatesProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersCreateBodyToJSON serializes a SessionTemplatesProvidersCreateBody to JSON.
func MapSessionTemplatesProvidersCreateBodyToJSON(v *SessionTemplatesProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
