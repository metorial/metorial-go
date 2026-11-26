package providers

import (
	"encoding/json"
	"time"
)

// SessionsProvidersDeleteOutputUsage - Usage statistics
type SessionsProvidersDeleteOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsProvidersDeleteOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsProvidersDeleteOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsProvidersDeleteOutputDeployment represents the sessions providers delete output deployment type.
type SessionsProvidersDeleteOutputDeployment struct {
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

// SessionsProvidersDeleteOutputConfig represents the sessions providers delete output config type.
type SessionsProvidersDeleteOutputConfig struct {
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

// SessionsProvidersDeleteOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsProvidersDeleteOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsProvidersDeleteOutput represents the sessions providers delete output type.
type SessionsProvidersDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsProvidersDeleteOutputUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsProvidersDeleteOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                                 `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsProvidersDeleteOutputDeployment `json:"deployment"`
	Config                 SessionsProvidersDeleteOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsProvidersDeleteOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsProvidersDeleteOutputFromJSON deserializes JSON data into a SessionsProvidersDeleteOutput.
func MapSessionsProvidersDeleteOutputFromJSON(data []byte) (*SessionsProvidersDeleteOutput, error) {
	var v SessionsProvidersDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersDeleteOutputToJSON serializes a SessionsProvidersDeleteOutput to JSON.
func MapSessionsProvidersDeleteOutputToJSON(v *SessionsProvidersDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
