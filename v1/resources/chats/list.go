package chats

import (
	"encoding/json"
	"time"
)

// ChatsListOutputItems represents the chats list output items type.
type ChatsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once the chat or its connection is archived. Archived chats stay listable.
	Status string `json:"status"`
	// Name - Display name.
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
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ChatsListOutputPagination represents the chats list output pagination type.
type ChatsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatsListOutput represents the chats list output type.
type ChatsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                    `json:"object"`
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

// ChatsListQueryCreatedAt - Only return results whose chat creation time is in this range. Set `gt`, `lt`, or both.
type ChatsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatsListQueryUpdatedAt - Only return results whose chat last update time is in this range. Set `gt`, `lt`, or both.
type ChatsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatsListQuery represents the chats list query type.
type ChatsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Search - Search by name
	Search *string `json:"search,omitempty"`
	// Status - Only return chats with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return chats with one of these IDs.
	Id *any `json:"id,omitempty"`
	// ChatConnectionId - Only return chats for one of these chat connections.
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Only return chats for one of these chat instances.
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// ChatInstanceProviderId - Only return chats for one of these chat instance providers.
	ChatInstanceProviderId *any `json:"chat_instance_provider_id,omitempty"`
	// CreatedAt - Only return results whose chat creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ChatsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose chat last update time is in this range. Set `gt`, `lt`, or both.
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
