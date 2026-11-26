package provider

import (
	"encoding/json"
	"time"
)

// ChatInstancesProviderSetOutputProvider represents the chat instances provider set output provider type.
type ChatInstancesProviderSetOutputProvider struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the provider.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatInstancesProviderSetOutputDeployment represents the chat instances provider set output deployment type.
type ChatInstancesProviderSetOutputDeployment struct {
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

// ChatInstancesProviderSetOutputConfig represents the chat instances provider set output config type.
type ChatInstancesProviderSetOutputConfig struct {
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

// ChatInstancesProviderSetOutputAuthConfig represents the chat instances provider set output auth config type.
type ChatInstancesProviderSetOutputAuthConfig struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name an auth config.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this auth config belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatInstancesProviderSetOutputIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesProviderSetOutputIdentity struct {
	// Id - Unique ID.
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

// ChatInstancesProviderSetOutput represents the chat instances provider set output type.
type ChatInstancesProviderSetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the instance is archived.
	Status string `json:"status"`
	// ChatInstanceId - The chat instance this provider belongs to
	ChatInstanceId string `json:"chat_instance_id"`
	// ChatConnectionProviderId - The connection-level provider this instance's provider was set up from
	ChatConnectionProviderId string                                    `json:"chat_connection_provider_id"`
	Provider                 ChatInstancesProviderSetOutputProvider    `json:"provider"`
	Deployment               *ChatInstancesProviderSetOutputDeployment `json:"deployment,omitempty"`
	Config                   *ChatInstancesProviderSetOutputConfig     `json:"config,omitempty"`
	AuthConfig               *ChatInstancesProviderSetOutputAuthConfig `json:"auth_config,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Identity - Account this instance is authorized as on the chat provider. `null` until observed.
	Identity *ChatInstancesProviderSetOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Timestamp when the provider was set up
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider config was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatInstancesProviderSetOutputFromJSON deserializes JSON data into a ChatInstancesProviderSetOutput.
func MapChatInstancesProviderSetOutputFromJSON(data []byte) (*ChatInstancesProviderSetOutput, error) {
	var v ChatInstancesProviderSetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesProviderSetOutputToJSON serializes a ChatInstancesProviderSetOutput to JSON.
func MapChatInstancesProviderSetOutputToJSON(v *ChatInstancesProviderSetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatInstancesProviderSetBody represents the chat instances provider set body type.
type ChatInstancesProviderSetBody struct {
	// ProviderId - Chat provider to use for this instance.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Provider deployment to use
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Provider config to use
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Provider auth config to use
	ProviderAuthConfigId *string `json:"provider_auth_config_id,omitempty"`
}

// MapChatInstancesProviderSetBodyFromJSON deserializes JSON data into a ChatInstancesProviderSetBody.
func MapChatInstancesProviderSetBodyFromJSON(data []byte) (*ChatInstancesProviderSetBody, error) {
	var v ChatInstancesProviderSetBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesProviderSetBodyToJSON serializes a ChatInstancesProviderSetBody to JSON.
func MapChatInstancesProviderSetBodyToJSON(v *ChatInstancesProviderSetBody) ([]byte, error) {
	return json.Marshal(v)
}
