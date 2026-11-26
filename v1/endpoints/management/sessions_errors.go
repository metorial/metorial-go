package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sessions/errors"
)

// SessionsErrorsEndpoint provides access to a session error is an error recorded while running a session. Similar errors group together for review.
type SessionsErrorsEndpoint struct {
	client *endpoint.Client
}

// NewSessionsErrorsEndpoint creates a new SessionsErrorsEndpoint.
func NewSessionsErrorsEndpoint(client *endpoint.Client) *SessionsErrorsEndpoint {
	return &SessionsErrorsEndpoint{client: client}
}

// SessionsErrorsEndpointListParams contains optional query parameters for List.
type SessionsErrorsEndpointListParams struct {
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
	// Id - Filter by session error ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// SessionErrorGroupId - Filter by error group ID(s)
	SessionErrorGroupId *any `json:"session_error_group_id,omitempty"`
	// ProviderRunId - Filter by provider run ID(s)
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// IssueId - Filter by provider issue ID(s)
	IssueId *any `json:"issue_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// CreatedAt - Only return results whose session error creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session error last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
	// Search - Search by session error ID or error code, or by the name of a related provider, deployment, or auth config
	Search *string `json:"search,omitempty"`
}

// List returns a paginated list of errors across all sessions.
func (e *SessionsErrorsEndpoint) List(instanceId string, params *SessionsErrorsEndpointListParams) (*errors.SessionsErrorsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "session-errors"},
		Query: query,
	}
	var result errors.SessionsErrorsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a session error.
func (e *SessionsErrorsEndpoint) Get(instanceId string, sessionErrorId string) (*errors.SessionsErrorsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "session-errors", sessionErrorId},
	}
	var result errors.SessionsErrorsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
