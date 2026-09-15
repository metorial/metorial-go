package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/protoguardalerts"
)

// ProtoGuardAlertsEndpoint provides access to protoGuard alerts describe prompt-injection detections.
type ProtoGuardAlertsEndpoint struct {
	client *endpoint.Client
}

// NewProtoGuardAlertsEndpoint creates a new ProtoGuardAlertsEndpoint.
func NewProtoGuardAlertsEndpoint(client *endpoint.Client) *ProtoGuardAlertsEndpoint {
	return &ProtoGuardAlertsEndpoint{client: client}
}

// ProtoGuardAlertsEndpointListParams contains optional query parameters for List.
type ProtoGuardAlertsEndpointListParams struct {
	Limit               *float64 `json:"limit,omitempty"`
	After               *string  `json:"after,omitempty"`
	Before              *string  `json:"before,omitempty"`
	Cursor              *string  `json:"cursor,omitempty"`
	Order               *string  `json:"order,omitempty"`
	Id                  *any     `json:"id,omitempty"`
	RunId               *any     `json:"run_id,omitempty"`
	FilterId            *any     `json:"filter_id,omitempty"`
	SessionId           *any     `json:"session_id,omitempty"`
	SessionMessageId    *any     `json:"session_message_id,omitempty"`
	SessionConnectionId *any     `json:"session_connection_id,omitempty"`
	ProviderRunId       *any     `json:"provider_run_id,omitempty"`
	// CreatedAt - Filter ProtoGuard alert creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of ProtoGuard alerts for this instance.
func (e *ProtoGuardAlertsEndpoint) List(params *ProtoGuardAlertsEndpointListParams) (*protoguardalerts.ProtoGuardAlertsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"protoguard-alerts"},
		Query: query,
	}
	var result protoguardalerts.ProtoGuardAlertsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a ProtoGuard alert by ID.
func (e *ProtoGuardAlertsEndpoint) Get(protoGuardAlertId string) (*protoguardalerts.ProtoGuardAlertsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"protoguard-alerts", protoGuardAlertId},
	}
	var result protoguardalerts.ProtoGuardAlertsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
