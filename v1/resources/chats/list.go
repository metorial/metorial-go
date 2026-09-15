package chats

import (
	"encoding/json"
	"time"
)

// ChatsListOutputItems represents the chats list output items type.
type ChatsListOutputItems struct {
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

// ChatsListOutputPagination represents the chats list output pagination type.
type ChatsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatsListOutput represents the chats list output type.
type ChatsListOutput struct {
	Items      []ChatsListOutputItems    `json:"items"`
	Pagination ChatsListOutputPagination `json:"pagination"`
}

// MapChatsListOutputFromJSON deserializes JSON data into a ChatsListOutput.
func MapChatsListOutputFromJSON(data []byte) (*ChatsListOutput, error) {
	var v ChatsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsListOutputToJSON serializes a ChatsListOutput to JSON.
func MapChatsListOutputToJSON(v *ChatsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsListQueryCreatedAt - Filter chat creation time by date range
type ChatsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for chat creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatsListQueryUpdatedAt - Filter chat last update time by date range
type ChatsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for chat last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatsListQuery represents the chats list query type.
type ChatsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Filter by status
	Status *any `json:"status,omitempty"`
	// Id - Filter by chat ID(s)
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Filter by chat connection ID(s)
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Filter by chat instance ID(s)
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// ChatInstanceProviderId - Filter by chat instance provider ID(s)
	ChatInstanceProviderId *any `json:"chat_instance_provider_id,omitempty"`
	// CreatedAt - Filter chat creation time by date range
	CreatedAt *ChatsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter chat last update time by date range
	UpdatedAt *ChatsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapChatsListQueryFromJSON deserializes JSON data into a ChatsListQuery.
func MapChatsListQueryFromJSON(data []byte) (*ChatsListQuery, error) {
	var v ChatsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsListQueryToJSON serializes a ChatsListQuery to JSON.
func MapChatsListQueryToJSON(v *ChatsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
