package threads

import (
	"encoding/json"
	"time"
)

// ChatsThreadsGetOutputContextAuthor represents the chats threads get output context author type.
type ChatsThreadsGetOutputContextAuthor struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsThreadsGetOutputContextAssignee represents the chats threads get output context assignee type.
type ChatsThreadsGetOutputContextAssignee struct {
	// Id - The actor's identifier on the chat provider
	Id string `json:"id"`
	// Name - The actor's display name
	Name string `json:"name"`
}

// ChatsThreadsGetOutputContext - The external resource this conversation is attached to on the provider, such as an issue, ticket or page.
type ChatsThreadsGetOutputContext struct {
	// Type - What kind of external resource this is
	Type string `json:"type"`
	// Id - The resource's identifier on the provider
	Id string `json:"id"`
	// Description - Description of the resource
	Description *string `json:"description,omitempty"`
	// Status - Status of the resource, as reported by the provider
	Status *string `json:"status,omitempty"`
	// Url - Link to the resource
	Url      *string                               `json:"url,omitempty"`
	Author   *ChatsThreadsGetOutputContextAuthor   `json:"author,omitempty"`
	Assignee *ChatsThreadsGetOutputContextAssignee `json:"assignee,omitempty"`
	// Labels - Labels attached to the resource
	Labels *[]string `json:"labels,omitempty"`
}

// ChatsThreadsGetOutput represents the chats threads get output type.
type ChatsThreadsGetOutput struct {
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
	Context *ChatsThreadsGetOutputContext `json:"context,omitempty"`
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

// MapChatsThreadsGetOutputFromJSON deserializes JSON data into a ChatsThreadsGetOutput.
func MapChatsThreadsGetOutputFromJSON(data []byte) (*ChatsThreadsGetOutput, error) {
	var v ChatsThreadsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsThreadsGetOutputToJSON serializes a ChatsThreadsGetOutput to JSON.
func MapChatsThreadsGetOutputToJSON(v *ChatsThreadsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatsThreadsGetQuery represents the chats threads get query type.
type ChatsThreadsGetQuery struct {
	// ChannelId - The channel this thread belongs to
	ChannelId string `json:"channel_id"`
}

// MapChatsThreadsGetQueryFromJSON deserializes JSON data into a ChatsThreadsGetQuery.
func MapChatsThreadsGetQueryFromJSON(data []byte) (*ChatsThreadsGetQuery, error) {
	var v ChatsThreadsGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatsThreadsGetQueryToJSON serializes a ChatsThreadsGetQuery to JSON.
func MapChatsThreadsGetQueryToJSON(v *ChatsThreadsGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
