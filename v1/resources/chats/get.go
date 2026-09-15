package chats

import (
	"encoding/json"
	"time"
)

// ChatsGetOutput represents the chats get output type.
type ChatsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat identifier
	Id string `json:"id"`
	// Status - The chat status
	Status string `json:"status"`
	// Name - Display name of the chat
	Name string `json:"name"`
	// ChatConnectionId - The chat connection this chat was created through
	ChatConnectionId string `json:"chat_connection_id"`
	// ChatInstanceId - The chat instance this chat belongs to
	ChatInstanceId string `json:"chat_instance_id"`
	// ChatInstanceProviderId - The chat instance provider this chat is running on
	ChatInstanceProviderId string `json:"chat_instance_provider_id"`
	// ProviderId - The chat provider backing this chat
	ProviderId string `json:"provider_id"`
	// WorkspaceId - The workspace this chat is connected to, once one has been resolved
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// CreatedAt - Timestamp when the chat was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the chat was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// MapChatsGetOutputFromJSON deserializes JSON data into a ChatsGetOutput.
func MapChatsGetOutputFromJSON(data []byte) (*ChatsGetOutput, error) {
	var v ChatsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsGetOutputToJSON serializes a ChatsGetOutput to JSON.
func MapChatsGetOutputToJSON(v *ChatsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
