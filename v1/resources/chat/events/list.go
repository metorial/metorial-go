package events

import (
	"encoding/json"
	"time"
)

// ChatEventsListOutputItems represents the chat events list output items type.
type ChatEventsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique chat event identifier
	Id string `json:"id"`
	// Type - What happened in the chat
	Type string `json:"type"`
	// Source - How Metorial learned about this event
	Source string `json:"source"`
	// ChatId - The chat this event was recorded for
	ChatId string `json:"chat_id"`
	// ChannelId - The channel this event happened in. Null when the event is not channel-scoped, or once the channel has been removed.
	ChannelId *string `json:"channel_id,omitempty"`
	// ThreadId - The thread this event happened in. Null when the event is not thread-scoped, or once the thread has been removed.
	ThreadId *string `json:"thread_id,omitempty"`
	// MessageId - The message this event is about. Null when the event is not about a message, or once the message has been removed.
	MessageId *string `json:"message_id,omitempty"`
	// AuthorId - The author who caused this event. Null when no author was involved, or once the author has been removed.
	AuthorId *string `json:"author_id,omitempty"`
	// ProviderEventId - The event's identifier on the chat provider, where the provider supplies one
	ProviderEventId *string `json:"provider_event_id,omitempty"`
	// ProviderChannelId - Provider identifier of the channel this event happened in. Stays readable after the channel has been removed.
	ProviderChannelId *string `json:"provider_channel_id,omitempty"`
	// ProviderThreadId - Provider identifier of the thread this event happened in. Stays readable after the thread has been removed.
	ProviderThreadId *string `json:"provider_thread_id,omitempty"`
	// ProviderMessageId - Provider identifier of the message this event is about. Stays readable after the message has been removed.
	ProviderMessageId *string `json:"provider_message_id,omitempty"`
	// ProviderAuthorId - Provider identifier of the author who caused this event. Stays readable after the author has been removed.
	ProviderAuthorId *string `json:"provider_author_id,omitempty"`
	// Payload - The event body, shaped like the chat resource the event is about. Null while the body is unavailable.
	Payload *map[string]any `json:"payload,omitempty"`
	// OccurredAt - Timestamp when the event happened on the chat provider
	OccurredAt time.Time `json:"occurred_at"`
	// CreatedAt - Timestamp when the event was recorded by Metorial
	CreatedAt time.Time `json:"created_at"`
}

// ChatEventsListOutputPagination represents the chat events list output pagination type.
type ChatEventsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ChatEventsListOutput represents the chat events list output type.
type ChatEventsListOutput struct {
	Items      []ChatEventsListOutputItems    `json:"items"`
	Pagination ChatEventsListOutputPagination `json:"pagination"`
}

// MapChatEventsListOutputFromJSON deserializes JSON data into a ChatEventsListOutput.
func MapChatEventsListOutputFromJSON(data []byte) (*ChatEventsListOutput, error) {
	var v ChatEventsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatEventsListOutputToJSON serializes a ChatEventsListOutput to JSON.
func MapChatEventsListOutputToJSON(v *ChatEventsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ChatEventsListQueryCreatedAt - Filter chat event creation time by date range
type ChatEventsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for chat event creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat event creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatEventsListQueryOccurredAt - Filter chat event occurrence time by date range
type ChatEventsListQueryOccurredAt struct {
	// Gt - Only include records after this timestamp for chat event occurrence time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for chat event occurrence time
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatEventsListQuery represents the chat events list query type.
type ChatEventsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ChatId - Filter by chat ID(s)
	ChatId *any `json:"chat_id,omitempty"`
	// ChatConnectionId - Filter by chat connection ID(s)
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Filter by chat instance ID(s)
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// Type - Filter by event type(s)
	Type *any `json:"type,omitempty"`
	// CreatedAt - Filter chat event creation time by date range
	CreatedAt *ChatEventsListQueryCreatedAt `json:"created_at,omitempty"`
	// OccurredAt - Filter chat event occurrence time by date range
	OccurredAt *ChatEventsListQueryOccurredAt `json:"occurred_at,omitempty"`
}

// MapChatEventsListQueryFromJSON deserializes JSON data into a ChatEventsListQuery.
func MapChatEventsListQueryFromJSON(data []byte) (*ChatEventsListQuery, error) {
	var v ChatEventsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatEventsListQueryToJSON serializes a ChatEventsListQuery to JSON.
func MapChatEventsListQueryToJSON(v *ChatEventsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
