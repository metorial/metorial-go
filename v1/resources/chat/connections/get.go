package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsGetOutputProvidersProvider represents the chat connections get output providers provider type.
type ChatConnectionsGetOutputProvidersProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsGetOutputProvidersDeployment represents the chat connections get output providers deployment type.
type ChatConnectionsGetOutputProvidersDeployment struct {
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

// ChatConnectionsGetOutputProvidersConfig represents the chat connections get output providers config type.
type ChatConnectionsGetOutputProvidersConfig struct {
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

// ChatConnectionsGetOutputProviders represents the chat connections get output providers type.
type ChatConnectionsGetOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection provider identifier
	Id string `json:"id"`
	// Status - The chat connection provider status
	Status     string                                       `json:"status"`
	Provider   ChatConnectionsGetOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsGetOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsGetOutputProvidersConfig     `json:"config,omitempty"`
	// AuthMethodId - The auth method this connection authenticates with, if any
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - The auth credentials this connection authenticates with, if any
	AuthCredentialsId *string `json:"auth_credentials_id,omitempty"`
	// Name - Display name of the provider link
	Name string `json:"name"`
	// Description - Description of the provider link
	Description *string `json:"description,omitempty"`
	// CreatedAt - Timestamp when the provider was linked
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsGetOutput represents the chat connections get output type.
type ChatConnectionsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection identifier
	Id string `json:"id"`
	// Status - The chat connection status
	Status string `json:"status"`
	// Slug - URL-safe identifier for the chat connection
	Slug string `json:"slug"`
	// Name - Display name of the chat connection
	Name string `json:"name"`
	// Description - Description of the chat connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata set on the chat connection
	Metadata map[string]any `json:"metadata"`
	// Providers - The provider linked to this chat connection. Currently always zero or one items.
	Providers []ChatConnectionsGetOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when the chat connection was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat connection was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatConnectionsGetOutputFromJSON deserializes JSON data into a ChatConnectionsGetOutput.
func MapChatConnectionsGetOutputFromJSON(data []byte) (*ChatConnectionsGetOutput, error) {
	var v ChatConnectionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsGetOutputToJSON serializes a ChatConnectionsGetOutput to JSON.
func MapChatConnectionsGetOutputToJSON(v *ChatConnectionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
