package connections

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// ChatConnectionsUpdateOutputProvidersProvider represents the chat connections update output providers provider type.
type ChatConnectionsUpdateOutputProvidersProvider struct {
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

// ChatConnectionsUpdateOutputProvidersDeployment represents the chat connections update output providers deployment type.
type ChatConnectionsUpdateOutputProvidersDeployment struct {
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

// ChatConnectionsUpdateOutputProvidersConfig represents the chat connections update output providers config type.
type ChatConnectionsUpdateOutputProvidersConfig struct {
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

// ChatConnectionsUpdateOutputProviders represents the chat connections update output providers type.
type ChatConnectionsUpdateOutputProviders struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the connection is archived.
	Status     string                                          `json:"status"`
	Provider   ChatConnectionsUpdateOutputProvidersProvider    `json:"provider"`
	Deployment *ChatConnectionsUpdateOutputProvidersDeployment `json:"deployment,omitempty"`
	Config     *ChatConnectionsUpdateOutputProvidersConfig     `json:"config,omitempty"`
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

// ChatConnectionsUpdateOutput represents the chat connections update output type.
type ChatConnectionsUpdateOutput struct {
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
	Providers []ChatConnectionsUpdateOutputProviders `json:"providers"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
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
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat connection. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the connection. Never returned by the API.
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
