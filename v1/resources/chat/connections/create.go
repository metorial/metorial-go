package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsCreateOutputProvidersProvider represents the chat connections create output providers provider type.
type ChatConnectionsCreateOutputProvidersProvider struct {
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

// ChatConnectionsCreateOutputProvidersDeployment represents the chat connections create output providers deployment type.
type ChatConnectionsCreateOutputProvidersDeployment struct {
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

// ChatConnectionsCreateOutputProvidersConfig represents the chat connections create output providers config type.
type ChatConnectionsCreateOutputProvidersConfig struct {
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

// ChatConnectionsCreateOutputProviders represents the chat connections create output providers type.
type ChatConnectionsCreateOutputProviders struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat connection provider identifier
	Id string `json:"id"`
	// Status - The chat connection provider status
	Status     string                                          `json:"status"`
	Provider   ChatConnectionsCreateOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsCreateOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsCreateOutputProvidersConfig     `json:"config,omitempty"`
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

// ChatConnectionsCreateOutput represents the chat connections create output type.
type ChatConnectionsCreateOutput struct {
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
	Providers []ChatConnectionsCreateOutputProviders `json:"providers"`
	// CreatedAt - Timestamp when the chat connection was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat connection was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatConnectionsCreateOutputFromJSON deserializes JSON data into a ChatConnectionsCreateOutput.
func MapChatConnectionsCreateOutputFromJSON(data []byte) (*ChatConnectionsCreateOutput, error) {
	var v ChatConnectionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsCreateOutputToJSON serializes a ChatConnectionsCreateOutput to JSON.
func MapChatConnectionsCreateOutputToJSON(v *ChatConnectionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatConnectionsCreateBodyProvider represents the chat connections create body provider type.
type ChatConnectionsCreateBodyProvider struct {
	// ProviderId - The chat provider to link this connection to
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Provider deployment to use for this connection
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Auth method to use for this connection
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Auth credentials to use for this connection
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config to use for this connection
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// Name - Display name of the provider link
	Name *string `json:"name,omitempty"`
	// Description - Description of the provider link
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata set on the provider link
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// ChatConnectionsCreateBody represents the chat connections create body type.
type ChatConnectionsCreateBody struct {
	// Name - Display name of the chat connection
	Name string `json:"name"`
	// Description - Description of the connection
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata for the connection
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata for the connection
	PrivateMetadata *map[string]any                   `json:"private_metadata,omitempty"`
	Provider        ChatConnectionsCreateBodyProvider `json:"provider"`
}

// MapChatConnectionsCreateBodyFromJSON deserializes JSON data into a ChatConnectionsCreateBody.
func MapChatConnectionsCreateBodyFromJSON(data []byte) (*ChatConnectionsCreateBody, error) {
	var v ChatConnectionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatConnectionsCreateBodyToJSON serializes a ChatConnectionsCreateBody to JSON.
func MapChatConnectionsCreateBodyToJSON(v *ChatConnectionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
