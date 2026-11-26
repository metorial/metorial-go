package members

import (
	"encoding/json"
	"time"
)

// ChatsChannelsMembersListOutputItems represents the chats channels members list output items type.
type ChatsChannelsMembersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this author belongs to
	ChatId string `json:"chat_id"`
	// Type - What kind of account this author is on the chat provider
	Type string `json:"type"`
	// Role - The role this author holds in the chat workspace
	Role string `json:"role"`
	// ProviderType - The provider's own name for this kind of author
	ProviderType string `json:"provider_type"`
	// ProviderAuthorId - Identifier the chat provider uses for this author.
	ProviderAuthorId string `json:"provider_author_id"`
	// UserName - The author's handle on the chat provider
	UserName string `json:"user_name"`
	// FullName - Display name.
	FullName string `json:"full_name"`
	// Email - Email address of the author, if the provider exposes one
	Email *string `json:"email,omitempty"`
	// ImageUrl - URL of the author's avatar, if the provider exposes one
	ImageUrl *string `json:"image_url,omitempty"`
	// IsSelf - Whether this author is the account the integration itself is authenticated as
	IsSelf bool `json:"is_self"`
	// CreatedAt - Timestamp when the author was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the author was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw from this author
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// ChatsChannelsMembersListOutputPagination represents the chats channels members list output pagination type.
type ChatsChannelsMembersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatsChannelsMembersListOutput represents the chats channels members list output type.
type ChatsChannelsMembersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                   `json:"object"`
	Items      []ChatsChannelsMembersListOutputItems    `json:"items"`
	Pagination ChatsChannelsMembersListOutputPagination `json:"pagination"`
}

// MapChatsChannelsMembersListOutputFromJSON deserializes JSON data into a ChatsChannelsMembersListOutput.
func MapChatsChannelsMembersListOutputFromJSON(data []byte) (*ChatsChannelsMembersListOutput, error) {
	var v ChatsChannelsMembersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsMembersListOutputToJSON serializes a ChatsChannelsMembersListOutput to JSON.
func MapChatsChannelsMembersListOutputToJSON(v *ChatsChannelsMembersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsChannelsMembersListQuery represents the chats channels members list query type.
type ChatsChannelsMembersListQuery struct {
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
}

// MapChatsChannelsMembersListQueryFromJSON deserializes JSON data into a ChatsChannelsMembersListQuery.
func MapChatsChannelsMembersListQueryFromJSON(data []byte) (*ChatsChannelsMembersListQuery, error) {
	var v ChatsChannelsMembersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsChannelsMembersListQueryToJSON serializes a ChatsChannelsMembersListQuery to JSON.
func MapChatsChannelsMembersListQueryToJSON(v *ChatsChannelsMembersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
