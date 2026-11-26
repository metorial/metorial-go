package providers

import (
	"encoding/json"
	"time"
)

// SessionTemplatesProvidersDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesProvidersDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesProvidersDeleteOutputDeployment represents the session templates providers delete output deployment type.
type SessionTemplatesProvidersDeleteOutputDeployment struct {
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

// SessionTemplatesProvidersDeleteOutputConfig represents the session templates providers delete output config type.
type SessionTemplatesProvidersDeleteOutputConfig struct {
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

// SessionTemplatesProvidersDeleteOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesProvidersDeleteOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesProvidersDeleteOutput represents the session templates providers delete output type.
type SessionTemplatesProvidersDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesProvidersDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                          `json:"session_template_id"`
	Deployment        SessionTemplatesProvidersDeleteOutputDeployment `json:"deployment"`
	Config            SessionTemplatesProvidersDeleteOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesProvidersDeleteOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionTemplatesProvidersDeleteOutputFromJSON deserializes JSON data into a SessionTemplatesProvidersDeleteOutput.
func MapSessionTemplatesProvidersDeleteOutputFromJSON(data []byte) (*SessionTemplatesProvidersDeleteOutput, error) {
	var v SessionTemplatesProvidersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersDeleteOutputToJSON serializes a SessionTemplatesProvidersDeleteOutput to JSON.
func MapSessionTemplatesProvidersDeleteOutputToJSON(v *SessionTemplatesProvidersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
