package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsUpdateOutputProvidersProvider represents the chat connections update output providers provider type.
type ChatConnectionsUpdateOutputProvidersProvider struct {
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

// ChatConnectionsUpdateOutputProvidersDeployment represents the chat connections update output providers deployment type.
type ChatConnectionsUpdateOutputProvidersDeployment struct {
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

// ChatConnectionsUpdateOutputProvidersConfig represents the chat connections update output providers config type.
type ChatConnectionsUpdateOutputProvidersConfig struct {
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

// ChatConnectionsUpdateOutputProviders represents the chat connections update output providers type.
type ChatConnectionsUpdateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection provider identifier
	Id string `json:"id"`
	// Status - The chat connection provider status
	Status     string                                          `json:"status"`
	Provider   ChatConnectionsUpdateOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsUpdateOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsUpdateOutputProvidersConfig     `json:"config,omitempty"`
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

// ChatConnectionsUpdateOutput represents the chat connections update output type.
type ChatConnectionsUpdateOutput struct {
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
	Providers []ChatConnectionsUpdateOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when the chat connection was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat connection was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatConnectionsUpdateOutputFromJSON deserializes JSON data into a ChatConnectionsUpdateOutput.
func MapChatConnectionsUpdateOutputFromJSON(data []byte) (*ChatConnectionsUpdateOutput, error) {
	var v ChatConnectionsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsUpdateOutputToJSON serializes a ChatConnectionsUpdateOutput to JSON.
func MapChatConnectionsUpdateOutputToJSON(v *ChatConnectionsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatConnectionsUpdateBody represents the chat connections update body type.
type ChatConnectionsUpdateBody struct {
	// Name - Display name of the chat connection
	Name *string `json:"name,omitempty"`
	// Description - Description of the connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
}

// MapChatConnectionsUpdateBodyFromJSON deserializes JSON data into a ChatConnectionsUpdateBody.
func MapChatConnectionsUpdateBodyFromJSON(data []byte) (*ChatConnectionsUpdateBody, error) {
	var v ChatConnectionsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsUpdateBodyToJSON serializes a ChatConnectionsUpdateBody to JSON.
func MapChatConnectionsUpdateBodyToJSON(v *ChatConnectionsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
