package connections

import (
	"encoding/json"
	"time"
)

// ChatConnectionsCreateOutputProvidersProvider represents the chat connections create output providers provider type.
type ChatConnectionsCreateOutputProvidersProvider struct {
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

// ChatConnectionsCreateOutputProvidersDeployment represents the chat connections create output providers deployment type.
type ChatConnectionsCreateOutputProvidersDeployment struct {
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

// ChatConnectionsCreateOutputProvidersConfig represents the chat connections create output providers config type.
type ChatConnectionsCreateOutputProvidersConfig struct {
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

// ChatConnectionsCreateOutputProviders represents the chat connections create output providers type.
type ChatConnectionsCreateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the connection is archived.
	Status     string                                          `json:"status"`
	Provider   ChatConnectionsCreateOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsCreateOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsCreateOutputProvidersConfig     `json:"config,omitempty"`
	// AuthMethodId - The auth method this connection authenticates with, if any
	AuthMethodId *string `json:"auth_method_id,omitempty"`
	// AuthCredentialsId - The auth credentials this connection authenticates with, if any
	AuthCredentialsId *string `json:"auth_credentials_id,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Timestamp when the provider was linked
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider link was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatConnectionsCreateOutput represents the chat connections create output type.
type ChatConnectionsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the connection. Archived connections stay listable.
	Status string `json:"status"`
	// Slug - URL-safe name, unique across Metorial.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Providers - The provider linked to this chat connection. Currently always zero or one items.
	Providers []ChatConnectionsCreateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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

// ChatConnectionsCreateBodyProvider - Chat provider to link this connection to.
type ChatConnectionsCreateBodyProvider struct {
	// ProviderId - Chat provider to link this connection to.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Provider deployment to use for this connection.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// ProviderAuthMethodId - Auth method to use for this connection.
	ProviderAuthMethodId *string `json:"provider_auth_method_id,omitempty"`
	// ProviderAuthCredentialsId - Auth credentials to use for this connection.
	ProviderAuthCredentialsId *string `json:"provider_auth_credentials_id,omitempty"`
	// ProviderConfigId - Provider config to use for this connection.
	ProviderConfigId *string `json:"provider_config_id,omitempty"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider link.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// ChatConnectionsCreateBody represents the chat connections create body type.
type ChatConnectionsCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat connection.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the connection. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	// Provider - Chat provider to link this connection to.
	Provider ChatConnectionsCreateBodyProvider `json:"provider"`
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
