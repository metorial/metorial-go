package instances

import (
	"encoding/json"
	"time"
)

// ChatInstancesCreateOutputIdentity - Account this instance is authorized as on the chat provider. `null` until observed.
type ChatInstancesCreateOutputIdentity struct {
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

// ChatInstancesCreateOutput represents the chat instances create output type.
type ChatInstancesCreateOutput struct {
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
	Identity *ChatInstancesCreateOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatInstancesCreateOutputFromJSON deserializes JSON data into a ChatInstancesCreateOutput.
func MapChatInstancesCreateOutputFromJSON(data []byte) (*ChatInstancesCreateOutput, error) {
	var v ChatInstancesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesCreateOutputToJSON serializes a ChatInstancesCreateOutput to JSON.
func MapChatInstancesCreateOutputToJSON(v *ChatInstancesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatInstancesCreateBody represents the chat instances create body type.
type ChatInstancesCreateBody struct {
	// ChatConnectionId - Chat connection to create this instance for.
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the chat instance.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private data stored with the instance. Never returned by the API.
	PrivateMetadata *map[string]any `json:"private_metadata,omitempty"`
	// IdentityActorId - Identity actor to run this instance as
	IdentityActorId *string `json:"identity_actor_id,omitempty"`
	// IdentityId - Identity to run this instance as
	IdentityId *string `json:"identity_id,omitempty"`
}

// MapChatInstancesCreateBodyFromJSON deserializes JSON data into a ChatInstancesCreateBody.
func MapChatInstancesCreateBodyFromJSON(data []byte) (*ChatInstancesCreateBody, error) {
	var v ChatInstancesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatInstancesCreateBodyToJSON serializes a ChatInstancesCreateBody to JSON.
func MapChatInstancesCreateBodyToJSON(v *ChatInstancesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
