package instances

import (
	"encoding/json"
	"time"
)

// ChatInstancesUpdateOutputIdentity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
type ChatInstancesUpdateOutputIdentity struct {
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

// ChatInstancesUpdateOutput represents the chat instances update output type.
type ChatInstancesUpdateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat instance identifier
	Id string `json:"id"`
	// Status - The chat instance status
	Status string `json:"status"`
	// ChatConnectionId - The chat connection this instance belongs to
	ChatConnectionId string `json:"chat_connection_id"`
	// Name - Display name of the chat instance
	Name string `json:"name"`
	// Description - Description of the chat instance
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata set on the chat instance
	Metadata map[string]any `json:"metadata"`
	// Identity - The identity this connection is authorized as on the chat provider (e.g. the connected bot or user). Null until this account has been observed on the provider (e.g. through a synced message or channel membership).
	Identity *ChatInstancesUpdateOutputIdentity `json:"identity,omitempty"`
	// CreatedAt - Timestamp when the chat instance was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat instance was last updated
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
	// Name - Display name of the chat instance
	Name *string `json:"name,omitempty"`
	// Description - Description of the instance
	Description *string `json:"description,omitempty"`
	// Metadata - Metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// PrivateMetadata - Private metadata
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
