package provider

import (
	"encoding/json"
	"time"
)

// ChatInstancesProviderGetOutputProvider represents the chat instances provider get output provider type.
type ChatInstancesProviderGetOutputProvider struct {
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

// ChatInstancesProviderGetOutputDeployment represents the chat instances provider get output deployment type.
type ChatInstancesProviderGetOutputDeployment struct {
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

// ChatInstancesProviderGetOutputConfig represents the chat instances provider get output config type.
type ChatInstancesProviderGetOutputConfig struct {
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

// ChatInstancesProviderGetOutputAuthConfig represents the chat instances provider get output auth config type.
type ChatInstancesProviderGetOutputAuthConfig struct {
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

// ChatInstancesProviderGetOutputIdentity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
type ChatInstancesProviderGetOutputIdentity struct {
	// Id - Unique chat author identifier
	Id string `json:"id"`
	// UserId - The provider-side user or app id this connection is authorized as
	UserId string `json:"user_id"`
	// Name - Display name of the authorized identity
	Name string `json:"name"`
	// Username - Username or handle of the authorized identity
	Username string `json:"username"`
	// ProviderType - The kind of identity as classified by the provider, e.g. a human user vs. an app/bot
	ProviderType string `json:"provider_type"`
	// Email - Email address of the authorized identity, if available
	Email *string `json:"email,omitempty"`
	// ImageUrl - Avatar or image URL of the authorized identity, if available
	ImageUrl *string `json:"image_url,omitempty"`
}

// ChatInstancesProviderGetOutput represents the chat instances provider get output type.
type ChatInstancesProviderGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat instance provider identifier
	Id string `json:"id"`
	// Status - The chat instance provider status
	Status string `json:"status"`
	// ChatInstanceId - The chat instance this provider belongs to
	ChatInstanceId string `json:"chat_instance_id"`
	// ChatConnectionProviderId - The connection-level provider this instance's provider was set up from
	ChatConnectionProviderId string                                    `json:"chat_connection_provider_id"`
	Provider                 ChatInstancesProviderGetOutputProvider    `json:"provider"`
	Deployment               *ChatInstancesProviderGetOutputDeployment `json:"deployment,omitempty"`
	Config                   *ChatInstancesProviderGetOutputConfig     `json:"config,omitempty"`
	AuthConfig               *ChatInstancesProviderGetOutputAuthConfig `json:"auth_config,omitempty"`
	// Name - Display name of the provider link
	Name string `json:"name"`
	// Description - Description of the provider link
	Description *string `json:"description,omitempty"`
	// Identity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
	Identity *ChatInstancesProviderGetOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Timestamp when the provider was set up
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider config was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatInstancesProviderGetOutputFromJSON deserializes JSON data into a ChatInstancesProviderGetOutput.
func MapChatInstancesProviderGetOutputFromJSON(data []byte) (*ChatInstancesProviderGetOutput, error) {
	var v ChatInstancesProviderGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesProviderGetOutputToJSON serializes a ChatInstancesProviderGetOutput to JSON.
func MapChatInstancesProviderGetOutputToJSON(v *ChatInstancesProviderGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
