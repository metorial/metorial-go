package providers

import (
	"encoding/json"
	"time"
)

// SessionsProvidersUpdateOutputUsage - Usage statistics
type SessionsProvidersUpdateOutputUsage struct {
	// TotalProductiveClientMessageCount - Total productive client messages
	TotalProductiveClientMessageCount float64 `json:"total_productive_client_message_count"`
	// TotalProductiveProviderMessageCount - Total productive provider messages
	TotalProductiveProviderMessageCount float64 `json:"total_productive_provider_message_count"`
}

// SessionsProvidersUpdateOutputToolFilter represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SessionsProvidersUpdateOutputToolFilter struct {
	Type                *string `json:"type,omitempty"`
	IgnoreParentFilters *bool   `json:"ignore_parent_filters,omitempty"`
	Filters             *[]any  `json:"filters,omitempty"`
}

// SessionsProvidersUpdateOutputDeployment represents the sessions providers update output deployment type.
type SessionsProvidersUpdateOutputDeployment struct {
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

// SessionsProvidersUpdateOutputConfig represents the sessions providers update output config type.
type SessionsProvidersUpdateOutputConfig struct {
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

// SessionsProvidersUpdateOutputAuthConfig - Auth config the provider runs with. `null` when there is none.
type SessionsProvidersUpdateOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
}

// SessionsProvidersUpdateOutput represents the sessions providers update output type.
type SessionsProvidersUpdateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the session provider is deleted. Archived session providers stay listable.
	Status string `json:"status"`
	// Usage - Usage statistics
	Usage SessionsProvidersUpdateOutputUsage `json:"usage"`
	// ToolFilter - Tool filter configuration
	ToolFilter SessionsProvidersUpdateOutputToolFilter `json:"tool_filter"`
	// ProviderId - Provider this session provider connects to.
	ProviderId string `json:"provider_id"`
	// SessionId - Session this provider is connected to.
	SessionId string `json:"session_id"`
	// FromTemplateId - Template this provider was created from. `null` when it was added directly.
	FromTemplateId *string `json:"from_template_id,omitempty"`
	// FromTemplateProviderId - Template provider this was created from. `null` when added directly.
	FromTemplateProviderId *string                                 `json:"from_template_provider_id,omitempty"`
	Deployment             SessionsProvidersUpdateOutputDeployment `json:"deployment"`
	Config                 SessionsProvidersUpdateOutputConfig     `json:"config"`
	// AuthConfig - Auth config the provider runs with. `null` when there is none.
	AuthConfig *SessionsProvidersUpdateOutputAuthConfig `json:"auth_config,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapSessionsProvidersUpdateOutputFromJSON deserializes JSON data into a SessionsProvidersUpdateOutput.
func MapSessionsProvidersUpdateOutputFromJSON(data []byte) (*SessionsProvidersUpdateOutput, error) {
	var v SessionsProvidersUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersUpdateOutputToJSON serializes a SessionsProvidersUpdateOutput to JSON.
func MapSessionsProvidersUpdateOutputToJSON(v *SessionsProvidersUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsProvidersUpdateBody represents the sessions providers update body type.
type SessionsProvidersUpdateBody struct {
	// ToolFilters - Limits which tools the provider can call. Omitted allows all tools.
	ToolFilters *any `json:"tool_filters,omitempty"`
}

// MapSessionsProvidersUpdateBodyFromJSON deserializes JSON data into a SessionsProvidersUpdateBody.
func MapSessionsProvidersUpdateBodyFromJSON(data []byte) (*SessionsProvidersUpdateBody, error) {
	var v SessionsProvidersUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsProvidersUpdateBodyToJSON serializes a SessionsProvidersUpdateBody to JSON.
func MapSessionsProvidersUpdateBodyToJSON(v *SessionsProvidersUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
