package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions/events"
)

// SessionsEventsEndpoint provides access to a session event records something that happened in a session, like a connection opening or an error occurring.
type SessionsEventsEndpoint struct {
	client *endpoint.Client
}

// NewSessionsEventsEndpoint creates a new SessionsEventsEndpoint.
func NewSessionsEventsEndpoint(client *endpoint.Client) *SessionsEventsEndpoint {
	return &SessionsEventsEndpoint{client: client}
}

// SessionsEventsEndpointListParams contains optional query parameters for List.
type SessionsEventsEndpointListParams struct {
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
	// Type - Filter by event type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by session event ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Filter by provider run ID(s)
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionErrorId - Filter by session error ID(s)
	SessionErrorId *any `json:"session_error_id,omitempty"`
	// CreatedAt - Only return results whose session event creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session event last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of session events.
func (e *SessionsEventsEndpoint) List(params *SessionsEventsEndpointListParams) (*events.SessionsEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"session-events"},
		Query: query,
	}
	var result events.SessionsEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session event.
func (e *SessionsEventsEndpoint) Get(sessionEventId string) (*events.SessionsEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-events", sessionEventId},
	}
	var result events.SessionsEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
