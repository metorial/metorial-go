package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsDeleteOutputProvidersProvider represents the chat connections delete output providers provider type.
type ChatConnectionsDeleteOutputProvidersProvider struct {
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

// ChatConnectionsDeleteOutputProvidersDeployment represents the chat connections delete output providers deployment type.
type ChatConnectionsDeleteOutputProvidersDeployment struct {
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

// ChatConnectionsDeleteOutputProvidersConfig represents the chat connections delete output providers config type.
type ChatConnectionsDeleteOutputProvidersConfig struct {
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

// ChatConnectionsDeleteOutputProviders represents the chat connections delete output providers type.
type ChatConnectionsDeleteOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection provider identifier
	Id string `json:"id"`
	// Status - The chat connection provider status
	Status     string                                          `json:"status"`
	Provider   ChatConnectionsDeleteOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsDeleteOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsDeleteOutputProvidersConfig     `json:"config,omitempty"`
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

// ChatConnectionsDeleteOutput represents the chat connections delete output type.
type ChatConnectionsDeleteOutput struct {
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
	Providers []ChatConnectionsDeleteOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when the chat connection was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat connection was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatConnectionsDeleteOutputFromJSON deserializes JSON data into a ChatConnectionsDeleteOutput.
func MapChatConnectionsDeleteOutputFromJSON(data []byte) (*ChatConnectionsDeleteOutput, error) {
	var v ChatConnectionsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsDeleteOutputToJSON serializes a ChatConnectionsDeleteOutput to JSON.
func MapChatConnectionsDeleteOutputToJSON(v *ChatConnectionsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
