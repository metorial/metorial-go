package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions/errorgroups"
)

// SessionsErrorGroupsEndpoint provides access to a session error group is a set of similar session errors. Groups show how often each failure repeats.
type SessionsErrorGroupsEndpoint struct {
	client *endpoint.Client
}

// NewSessionsErrorGroupsEndpoint creates a new SessionsErrorGroupsEndpoint.
func NewSessionsErrorGroupsEndpoint(client *endpoint.Client) *SessionsErrorGroupsEndpoint {
	return &SessionsErrorGroupsEndpoint{client: client}
}

// SessionsErrorGroupsEndpointListParams contains optional query parameters for List.
type SessionsErrorGroupsEndpointListParams struct {
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
	// Type - Filter by error type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by error group ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search by error group ID or error code, or by the name of a related provider
	Search *string `json:"search,omitempty"`
}

// List returns a paginated list of error groups across all sessions.
func (e *SessionsErrorGroupsEndpoint) List(params *SessionsErrorGroupsEndpointListParams) (*errorgroups.SessionsErrorGroupsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"session-error-groups"},
		Query: query,
	}
	var result errorgroups.SessionsErrorGroupsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session error group.
func (e *SessionsErrorGroupsEndpoint) Get(sessionErrorGroupId string) (*errorgroups.SessionsErrorGroupsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"session-error-groups", sessionErrorGroupId},
	}
	var result errorgroups.SessionsErrorGroupsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
