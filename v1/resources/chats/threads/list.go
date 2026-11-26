package threads

import (
	"encoding/json"
	"time"
)

// ChatsThreadsListOutputItemsContextAuthor - Author of the external resource, as reported by the provider.
type ChatsThreadsListOutputItemsContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsThreadsListOutputItemsContextAssignee - Assignee of the external resource, as reported by the provider.
type ChatsThreadsListOutputItemsContextAssignee struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsThreadsListOutputItemsContext - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
type ChatsThreadsListOutputItemsContext struct {
	// Type - What kind of external resource this is
	Type string `json:"type"`
	// Id - The resource's identifier on the provider
	Id string `json:"id"`
	// Description - Description of the external resource, as reported by the provider.
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url *string `json:"url,omitempty"`
	// Author - Author of the external resource, as reported by the provider.
	Author *ChatsThreadsListOutputItemsContextAuthor `json:"author,omitempty"`
	// Assignee - Assignee of the external resource, as reported by the provider.
	Assignee *ChatsThreadsListOutputItemsContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsThreadsListOutputItems represents the chats threads list output items type.
type ChatsThreadsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ChatId - The chat this thread belongs to
	ChatId string `json:"chat_id"`
	// ChannelId - The channel this thread was started in
	ChannelId string `json:"channel_id"`
	// Type - How the thread is scoped on the chat provider
	Type string `json:"type"`
	// ProviderType - The provider's own name for this kind of thread
	ProviderType string `json:"provider_type"`
	// ProviderThreadId - Identifier the chat provider uses for this thread.
	ProviderThreadId string `json:"provider_thread_id"`
	// ProviderRootMessageId - Provider identifier of the message the thread was started from, if the provider reports one
	ProviderRootMessageId *string `json:"provider_root_message_id,omitempty"`
	// Subject - Subject of the thread, if the provider sets one.
	Subject *string `json:"subject,omitempty"`
	// Permalink - Link to the thread on the chat provider
	Permalink *string `json:"permalink,omitempty"`
	// Context - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
	Context *ChatsThreadsListOutputItemsContext `json:"context,omitempty"`
	// ReplyCount - Number of replies in the thread, if the provider reports it
	ReplyCount *float64 `json:"reply_count,omitempty"`
	// LastReplyAt - Timestamp of the most recent reply, if the provider reports it
	LastReplyAt *time.Time `json:"last_reply_at,omitempty"`
	// CreatedAt - Timestamp when the thread was first seen by Metorial
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the thread was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// LastInteractionAt - Timestamp of the last activity Metorial saw in this thread
	LastInteractionAt *time.Time `json:"last_interaction_at,omitempty"`
}

// ChatsThreadsListOutputPagination represents the chats threads list output pagination type.
type ChatsThreadsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatsThreadsListOutput represents the chats threads list output type.
type ChatsThreadsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
	Items      []ChatsThreadsListOutputItems    `json:"items"`
	Pagination ChatsThreadsListOutputPagination `json:"pagination"`
}

// MapChatsThreadsListOutputFromJSON deserializes JSON data into a ChatsThreadsListOutput.
func MapChatsThreadsListOutputFromJSON(data []byte) (*ChatsThreadsListOutput, error) {
	var v ChatsThreadsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsThreadsListOutputToJSON serializes a ChatsThreadsListOutput to JSON.
func MapChatsThreadsListOutputToJSON(v *ChatsThreadsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsThreadsListQuery represents the chats threads list query type.
type ChatsThreadsListQuery struct {
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
	// ChannelId - The channel to list threads for
	ChannelId string `json:"channel_id"`
	// Type - Only return threads of this type.
	Type *string `json:"type,omitempty"`
}

// MapChatsThreadsListQueryFromJSON deserializes JSON data into a ChatsThreadsListQuery.
func MapChatsThreadsListQueryFromJSON(data []byte) (*ChatsThreadsListQuery, error) {
	var v ChatsThreadsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsThreadsListQueryToJSON serializes a ChatsThreadsListQuery to JSON.
func MapChatsThreadsListQueryToJSON(v *ChatsThreadsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
