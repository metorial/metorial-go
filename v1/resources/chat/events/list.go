package events

import (
	"encoding/json"
	"time"
)

// ChatEventsListOutputItems represents the chat events list output items type.
type ChatEventsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What happened in the chat
	Type string `json:"type"`
	// Source - How Metorial learned about this event
	Source string `json:"source"`
	// ChatConnectionId - The chat connection this event was recorded for
	ChatConnectionId string `json:"chat_connection_id"`
	// ChatId - The chat this event was recorded for. Null for events recorded before any chat exists, such as a connection or instance lifecycle event.
	ChatId *string `json:"chat_id,omitempty"`
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
	// OccurredAt - Timestamp when the event happened on the chat provider
	OccurredAt time.Time `json:"occurred_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ChatEventsListOutputPagination represents the chat events list output pagination type.
type ChatEventsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ChatEventsListOutput represents the chat events list output type.
type ChatEventsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                         `json:"object"`
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

// ChatEventsListQueryCreatedAt - Only return results whose chat event creation time is in this range. Set `gt`, `lt`, or both.
type ChatEventsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatEventsListQueryOccurredAt - Only return results whose chat event occurrence time is in this range. Set `gt`, `lt`, or both.
type ChatEventsListQueryOccurredAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ChatEventsListQuery represents the chat events list query type.
type ChatEventsListQuery struct {
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
	// ChatId - Only return events for one of these chats.
	ChatId *any `json:"chat_id,omitempty"`
	// ChatConnectionId - Only return events for one of these chat connections.
	ChatConnectionId *any `json:"chat_connection_id,omitempty"`
	// ChatInstanceId - Only return events for one of these chat instances.
	ChatInstanceId *any `json:"chat_instance_id,omitempty"`
	// Type - Only return events with one of these types.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose chat event creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ChatEventsListQueryCreatedAt `json:"created_at,omitempty"`
	// OccurredAt - Only return results whose chat event occurrence time is in this range. Set `gt`, `lt`, or both.
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
