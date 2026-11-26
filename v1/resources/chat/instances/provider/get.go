package provider

import (
	"encoding/json"
	"time"
)

// ChatInstancesProviderGetOutputProvider represents the chat instances provider get output provider type.
type ChatInstancesProviderGetOutputProvider struct {
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

// ChatInstancesProviderGetOutputDeployment represents the chat instances provider get output deployment type.
type ChatInstancesProviderGetOutputDeployment struct {
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

// ChatInstancesProviderGetOutputConfig represents the chat instances provider get output config type.
type ChatInstancesProviderGetOutputConfig struct {
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

// ChatInstancesProviderGetOutputAuthConfig represents the chat instances provider get output auth config type.
type ChatInstancesProviderGetOutputAuthConfig struct {
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

// ChatInstancesProviderGetOutputIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesProviderGetOutputIdentity struct {
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

// ChatInstancesProviderGetOutput represents the chat instances provider get output type.
type ChatInstancesProviderGetOutput struct {
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
	Provider                 ChatInstancesProviderGetOutputProvider    `json:"provider"`
	Deployment               *ChatInstancesProviderGetOutputDeployment `json:"deployment,omitempty"`
	Config                   *ChatInstancesProviderGetOutputConfig     `json:"config,omitempty"`
	AuthConfig               *ChatInstancesProviderGetOutputAuthConfig `json:"auth_config,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Identity - Account this instance is authorized as on the chat provider. `null` until observed.
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
