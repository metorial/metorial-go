package instances

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// ChatInstancesUpdateOutputIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesUpdateOutputIdentity struct {
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

// ChatInstancesUpdateOutput represents the chat instances update output type.
type ChatInstancesUpdateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you archive the instance. Archived instances stay listable.
	Status string `json:"status"`
	// ChatConnectionId - The chat connection this instance belongs to
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// Identity - Account this instance is authorized as on the chat provider. `null` until observed.
	Identity *ChatInstancesUpdateOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatInstancesUpdateOutputFromJSON deserializes JSON data into a ChatInstancesUpdateOutput.
func MapChatInstancesUpdateOutputFromJSON(data []byte) (*ChatInstancesUpdateOutput, error) {
	var v ChatInstancesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesUpdateOutputToJSON serializes a ChatInstancesUpdateOutput to JSON.
func MapChatInstancesUpdateOutputToJSON(v *ChatInstancesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatInstancesUpdateBody represents the chat instances update body type.
type ChatInstancesUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat instance. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the instance. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
}

// MapChatInstancesUpdateBodyFromJSON deserializes JSON data into a ChatInstancesUpdateBody.
func MapChatInstancesUpdateBodyFromJSON(data []byte) (*ChatInstancesUpdateBody, error) {
	var v ChatInstancesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesUpdateBodyToJSON serializes a ChatInstancesUpdateBody to JSON.
func MapChatInstancesUpdateBodyToJSON(v *ChatInstancesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
