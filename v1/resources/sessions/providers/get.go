package providers

import (
	"encoding/json"
	"time"
)

// SessionsProvidersGetOutputUsage - Usage statistics
type SessionsProvidersGetOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsProvidersGetOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsProvidersGetOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsProvidersGetOutputDeployment represents the sessions providers get output deployment type.
type SessionsProvidersGetOutputDeployment struct {
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

// SessionsProvidersGetOutputConfig represents the sessions providers get output config type.
type SessionsProvidersGetOutputConfig struct {
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

// SessionsProvidersGetOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsProvidersGetOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsProvidersGetOutput represents the sessions providers get output type.
type SessionsProvidersGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsProvidersGetOutputUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsProvidersGetOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                              `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsProvidersGetOutputDeployment `json:"deployment"`
	Config                 SessionsProvidersGetOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsProvidersGetOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsProvidersGetOutputFromJSON deserializes JSON data into a SessionsProvidersGetOutput.
func MapSessionsProvidersGetOutputFromJSON(data []byte) (*SessionsProvidersGetOutput, error) {
	var v SessionsProvidersGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersGetOutputToJSON serializes a SessionsProvidersGetOutput to JSON.
func MapSessionsProvidersGetOutputToJSON(v *SessionsProvidersGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
