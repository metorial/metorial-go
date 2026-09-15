package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/monitoralerts"
)

// MonitorAlertsEndpoint provides access to monitor alerts represent detected prompt-injection or schema-change events.
type MonitorAlertsEndpoint struct {
	client *endpoint.Client
}

// NewMonitorAlertsEndpoint creates a new MonitorAlertsEndpoint.
func NewMonitorAlertsEndpoint(client *endpoint.Client) *MonitorAlertsEndpoint {
	return &MonitorAlertsEndpoint{client: client}
}

// MonitorAlertsEndpointListParams contains optional query parameters for List.
type MonitorAlertsEndpointListParams struct {
	Limit                             *float64 `json:"limit,omitempty"`
	After                             *string  `json:"after,omitempty"`
	Before                            *string  `json:"before,omitempty"`
	Cursor                            *string  `json:"cursor,omitempty"`
	Order                             *string  `json:"order,omitempty"`
	Id                                *any     `json:"id,omitempty"`
	MonitorId                         *any     `json:"monitor_id,omitempty"`
	Status                            *any     `json:"status,omitempty"`
	Target                            *any     `json:"target,omitempty"`
	Source                            *any     `json:"source,omitempty"`
	ProviderId                        *any     `json:"provider_id,omitempty"`
	ProtoGuardAlertId                 *any     `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                   *any     `json:"proto_guard_run_id,omitempty"`
	ProtoGuardFilterId                *any     `json:"proto_guard_filter_id,omitempty"`
	SpecificationChangeNotificationId *any     `json:"specification_change_notification_id,omitempty"`
	SessionId                         *any     `json:"session_id,omitempty"`
	SessionMessageId                  *any     `json:"session_message_id,omitempty"`
	SessionConnectionId               *any     `json:"session_connection_id,omitempty"`
	ProviderRunId                     *any     `json:"provider_run_id,omitempty"`
	// CreatedAt - Filter monitor alert creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// ResolvedAt - Filter monitor alert resolution time by date range
	ResolvedAt *map[string]any `json:"resolved_at,omitempty"`
}

// List returns a paginated list of monitor alerts for this instance.
func (e *MonitorAlertsEndpoint) List(instanceId string, params *MonitorAlertsEndpointListParams) (*monitoralerts.MonitorAlertsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "monitor-alerts"},
		Query: query,
	}
	var result monitoralerts.MonitorAlertsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a monitor alert by ID.
func (e *MonitorAlertsEndpoint) Get(instanceId string, monitorAlertId string) (*monitoralerts.MonitorAlertsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "monitor-alerts", monitorAlertId},
	}
	var result monitoralerts.MonitorAlertsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Viewed marks a monitor alert as viewed by the current actor.
func (e *MonitorAlertsEndpoint) Viewed(instanceId string, monitorAlertId string) (*monitoralerts.MonitorAlertsViewedOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "monitor-alerts", monitorAlertId, "viewed"},
	}
	var result monitoralerts.MonitorAlertsViewedOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Resolve marks a monitor alert as resolved.
func (e *MonitorAlertsEndpoint) Resolve(instanceId string, monitorAlertId string) (*monitoralerts.MonitorAlertsResolveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "monitor-alerts", monitorAlertId, "resolve"},
	}
	var result monitoralerts.MonitorAlertsResolveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unresolve reopens a resolved monitor alert.
func (e *MonitorAlertsEndpoint) Unresolve(instanceId string, monitorAlertId string) (*monitoralerts.MonitorAlertsUnresolveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "monitor-alerts", monitorAlertId, "unresolve"},
	}
	var result monitoralerts.MonitorAlertsUnresolveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
