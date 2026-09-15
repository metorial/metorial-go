package events

import (
	"encoding/json"
	"time"
)

// ChatEventsGetOutput represents the chat events get output type.
type ChatEventsGetOutput struct {
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

// MapChatEventsGetOutputFromJSON deserializes JSON data into a ChatEventsGetOutput.
func MapChatEventsGetOutputFromJSON(data []byte) (*ChatEventsGetOutput, error) {
	var v ChatEventsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapChatEventsGetOutputToJSON serializes a ChatEventsGetOutput to JSON.
func MapChatEventsGetOutputToJSON(v *ChatEventsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
