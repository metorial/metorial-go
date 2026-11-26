package providers

import (
	"encoding/json"
	"time"
)

// SessionTemplatesProvidersUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionTemplatesProvidersUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionTemplatesProvidersUpdateOutputDeployment represents the session templates providers update output deployment type.
type SessionTemplatesProvidersUpdateOutputDeployment struct {
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

// SessionTemplatesProvidersUpdateOutputConfig represents the session templates providers update output config type.
type SessionTemplatesProvidersUpdateOutputConfig struct {
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

// SessionTemplatesProvidersUpdateOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionTemplatesProvidersUpdateOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionTemplatesProvidersUpdateOutput represents the session templates providers update output type.
type SessionTemplatesProvidersUpdateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the template provider is removed. Removed template providers stay listable.
	Status string `json:"status"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionTemplatesProvidersUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this template provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionTemplateId - Template this provider belongs to.
	SessionTemplateId string                                          `json:"session_template_id"`
	Deployment        SessionTemplatesProvidersUpdateOutputDeployment `json:"deployment"`
	Config            SessionTemplatesProvidersUpdateOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionTemplatesProvidersUpdateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionTemplatesProvidersUpdateOutputFromJSON deserializes JSON data into a SessionTemplatesProvidersUpdateOutput.
func MapSessionTemplatesProvidersUpdateOutputFromJSON(data []byte) (*SessionTemplatesProvidersUpdateOutput, error) {
	var v SessionTemplatesProvidersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersUpdateOutputToJSON serializes a SessionTemplatesProvidersUpdateOutput to JSON.
func MapSessionTemplatesProvidersUpdateOutputToJSON(v *SessionTemplatesProvidersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionTemplatesProvidersUpdateBody represents the session templates providers update body type.
type SessionTemplatesProvidersUpdateBody struct {
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapSessionTemplatesProvidersUpdateBodyFromJSON deserializes JSON data into a SessionTemplatesProvidersUpdateBody.
func MapSessionTemplatesProvidersUpdateBodyFromJSON(data []byte) (*SessionTemplatesProvidersUpdateBody, error) {
	var v SessionTemplatesProvidersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionTemplatesProvidersUpdateBodyToJSON serializes a SessionTemplatesProvidersUpdateBody to JSON.
func MapSessionTemplatesProvidersUpdateBodyToJSON(v *SessionTemplatesProvidersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
