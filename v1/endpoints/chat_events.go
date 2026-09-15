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
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// OccurredAt - Filter chat event occurrence time by date range
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

// Get retrieves a specific chat event.
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
