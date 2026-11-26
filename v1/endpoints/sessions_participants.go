package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions/participants"
)

// SessionsParticipantsEndpoint provides access to a session participant is a client, provider, or system actor in a session. Participants appear automatically as connections open.
type SessionsParticipantsEndpoint struct {
	client *endpoint.Client
}

// NewSessionsParticipantsEndpoint creates a new SessionsParticipantsEndpoint.
func NewSessionsParticipantsEndpoint(client *endpoint.Client) *SessionsParticipantsEndpoint {
	return &SessionsParticipantsEndpoint{client: client}
}

// SessionsParticipantsEndpointListParams contains optional query parameters for List.
type SessionsParticipantsEndpointListParams struct {
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
	// Type - Filter by participant type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by participant ID(s)
	Id *any `json:"id,omitempty"`
	// AgentId - Filter by connected agent ID(s)
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Filter by connected identity actor ID(s)
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerId - Filter by connected consumer ID(s)
	ConsumerId *any `json:"consumer_id,omitempty"`
	// IdentityId - Filter by connected identity ID(s)
	IdentityId *any `json:"identity_id,omitempty"`
	// AgentInstanceId - Filter by connected agent instance ID(s)
	AgentInstanceId *any `json:"agent_instance_id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// CreatedAt - Only return results whose session participant creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session participant last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of session participants.
func (e *SessionsParticipantsEndpoint) List(params *SessionsParticipantsEndpointListParams) (*participants.SessionsParticipantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"session-participants"},
		Query: query,
	}
	var result participants.SessionsParticipantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session participant.
func (e *SessionsParticipantsEndpoint) Get(sessionParticipantId string) (*participants.SessionsParticipantsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-participants", sessionParticipantId},
	}
	var result participants.SessionsParticipantsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
