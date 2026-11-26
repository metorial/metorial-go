package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/protoguardalerts"
)

// ProtoGuardAlertsEndpoint provides access to a ProtoGuard alert is a record of related prompt-injection findings. Findings that recur are folded into the same alert.
type ProtoGuardAlertsEndpoint struct {
	client *endpoint.Client
}

// NewProtoGuardAlertsEndpoint creates a new ProtoGuardAlertsEndpoint.
func NewProtoGuardAlertsEndpoint(client *endpoint.Client) *ProtoGuardAlertsEndpoint {
	return &ProtoGuardAlertsEndpoint{client: client}
}

// ProtoGuardAlertsEndpointListParams contains optional query parameters for List.
type ProtoGuardAlertsEndpointListParams struct {
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
	// Id - Only return ProtoGuard alerts with one of these IDs.
	Id *any `json:"id,omitempty"`
	// RunId - Only return ProtoGuard alerts from one of these runs.
	RunId *any `json:"run_id,omitempty"`
	// FilterId - Only return ProtoGuard alerts matching one of these filters.
	FilterId *any `json:"filter_id,omitempty"`
	// SessionId - Only return ProtoGuard alerts linked to one of these sessions.
	SessionId *any `json:"session_id,omitempty"`
	// SessionMessageId - Only return ProtoGuard alerts linked to one of these session messages.
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionConnectionId - Only return ProtoGuard alerts linked to one of these session connections.
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Only return ProtoGuard alerts linked to one of these provider runs.
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// IssueId - Only return ProtoGuard alerts folded into one of these issues.
	IssueId *any `json:"issue_id,omitempty"`
	// CreatedAt - Only return results whose ProtoGuard alert creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of ProtoGuard alerts in the instance.
func (e *ProtoGuardAlertsEndpoint) List(instanceId string, params *ProtoGuardAlertsEndpointListParams) (*protoguardalerts.ProtoGuardAlertsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "protoguard-alerts"},
		Query: query,
	}
	var result protoguardalerts.ProtoGuardAlertsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a ProtoGuard alert.
func (e *ProtoGuardAlertsEndpoint) Get(instanceId string, protoGuardAlertId string) (*protoguardalerts.ProtoGuardAlertsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-alerts", protoGuardAlertId},
	}
	var result protoguardalerts.ProtoGuardAlertsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
