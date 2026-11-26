package providers

import (
	"encoding/json"
	"time"
)

// SessionsProvidersCreateOutputUsage - Usage statistics
type SessionsProvidersCreateOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsProvidersCreateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsProvidersCreateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsProvidersCreateOutputDeployment represents the sessions providers create output deployment type.
type SessionsProvidersCreateOutputDeployment struct {
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

// SessionsProvidersCreateOutputConfig represents the sessions providers create output config type.
type SessionsProvidersCreateOutputConfig struct {
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

// SessionsProvidersCreateOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsProvidersCreateOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsProvidersCreateOutput represents the sessions providers create output type.
type SessionsProvidersCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsProvidersCreateOutputUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsProvidersCreateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                                 `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsProvidersCreateOutputDeployment `json:"deployment"`
	Config                 SessionsProvidersCreateOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsProvidersCreateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsProvidersCreateOutputFromJSON deserializes JSON data into a SessionsProvidersCreateOutput.
func MapSessionsProvidersCreateOutputFromJSON(data []byte) (*SessionsProvidersCreateOutput, error) {
	var v SessionsProvidersCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersCreateOutputToJSON serializes a SessionsProvidersCreateOutput to JSON.
func MapSessionsProvidersCreateOutputToJSON(v *SessionsProvidersCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsProvidersCreateBody represents the sessions providers create body type.
type SessionsProvidersCreateBody struct {
	SessionId string `json:"session_id"`
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapSessionsProvidersCreateBodyFromJSON deserializes JSON data into a SessionsProvidersCreateBody.
func MapSessionsProvidersCreateBodyFromJSON(data []byte) (*SessionsProvidersCreateBody, error) {
	var v SessionsProvidersCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersCreateBodyToJSON serializes a SessionsProvidersCreateBody to JSON.
func MapSessionsProvidersCreateBodyToJSON(v *SessionsProvidersCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
