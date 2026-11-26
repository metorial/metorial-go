package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/chat/events"
)

// ChatEventsEndpoint provides access to chat events record what happened on a chat, such as a new message, reaction, or membership change.
type ChatEventsEndpoint struct {
	client *endpoint.Client
}

// NewChatEventsEndpoint creates a new ChatEventsEndpoint.
func NewChatEventsEndpoint(client *endpoint.Client) *ChatEventsEndpoint {
	return &ChatEventsEndpoint{client: client}
}

// ChatEventsEndpointListParams contains optional query parameters for List.
type ChatEventsEndpointListParams struct {
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
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// OccurredAt - Only return results whose chat event occurrence time is in this range. Set `gt`, `lt`, or both.
	OccurredAt *map[string]any `json:"occurred_at,omitempty"`
}

// List returns a paginated list of chat events.
func (e *ChatEventsEndpoint) List(params *ChatEventsEndpointListParams) (*events.ChatEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"chat", "events"},
		Query: query,
	}
	var result events.ChatEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a chat event.
func (e *ChatEventsEndpoint) Get(chatEventId string) (*events.ChatEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"chat", "events", chatEventId},
	}
	var result events.ChatEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
