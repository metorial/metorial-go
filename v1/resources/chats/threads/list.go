package threads

import (
	"encoding/json"
	"time"
)

// ChatsThreadsListOutputItemsContextAuthor represents the chats threads list output items context author type.
type ChatsThreadsListOutputItemsContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsThreadsListOutputItemsContextAssignee represents the chats threads list output items context assignee type.
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
	// Description - Description of the resource
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url      *string                                     `json:"url,omitempty"`
	Author   *ChatsThreadsListOutputItemsContextAuthor   `json:"author,omitempty"`
	Assignee *ChatsThreadsListOutputItemsContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsThreadsListOutputItems represents the chats threads list output items type.
type ChatsThreadsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat thread identifier
	Id string `json:"id"`
	// ChatId - The chat this thread belongs to
	ChatId string `json:"chat_id"`
	// ChannelId - The channel this thread was started in
	ChannelId string `json:"channel_id"`
	// Type - How the thread is scoped on the chat provider
	Type string `json:"type"`
	// ProviderType - The provider's own name for this kind of thread
	ProviderType string `json:"provider_type"`
	// ProviderThreadId - The thread's identifier on the chat provider
	ProviderThreadId string `json:"provider_thread_id"`
	// ProviderRootMessageId - Provider identifier of the message the thread was started from, if the provider reports one
	ProviderRootMessageId *string `json:"provider_root_message_id,omitempty"`
	// Subject - Subject of the thread
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatsThreadsListOutput represents the chats threads list output type.
type ChatsThreadsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ChannelId - The channel to list threads for
	ChannelId string `json:"channel_id"`
	// Type - Filter by thread type
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
