package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/monitoralerts"
)

// MonitorAlertsEndpoint provides access to a monitor alert is one firing of a monitor. It records a detected prompt-injection or schema-change event.
type MonitorAlertsEndpoint struct {
	client *endpoint.Client
}

// NewMonitorAlertsEndpoint creates a new MonitorAlertsEndpoint.
func NewMonitorAlertsEndpoint(client *endpoint.Client) *MonitorAlertsEndpoint {
	return &MonitorAlertsEndpoint{client: client}
}

// MonitorAlertsEndpointListParams contains optional query parameters for List.
type MonitorAlertsEndpointListParams struct {
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
	// Id - Only return monitor alerts with one of these IDs.
	Id *any `json:"id,omitempty"`
	// MonitorId - Only return alerts from one of these monitors.
	MonitorId *any `json:"monitor_id,omitempty"`
	// Status - Only return monitor alerts with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Target - Only return monitor alerts watching one of these targets.
	Target *any `json:"target,omitempty"`
	// Source - Only return monitor alerts from one of these sources.
	Source *any `json:"source,omitempty"`
	// ProviderId - Only return monitor alerts for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProtoGuardAlertId - Only return monitor alerts linked to one of these ProtoGuard alerts.
	ProtoGuardAlertId *any `json:"proto_guard_alert_id,omitempty"`
	// ProtoGuardRunId - Only return monitor alerts linked to one of these ProtoGuard runs.
	ProtoGuardRunId *any `json:"proto_guard_run_id,omitempty"`
	// ProtoGuardFilterId - Only return monitor alerts linked to one of these ProtoGuard filters.
	ProtoGuardFilterId *any `json:"proto_guard_filter_id,omitempty"`
	// SpecificationChangeNotificationId - Only return monitor alerts linked to one of these specification change notifications.
	SpecificationChangeNotificationId *any `json:"specification_change_notification_id,omitempty"`
	// SessionId - Only return monitor alerts linked to one of these sessions.
	SessionId *any `json:"session_id,omitempty"`
	// SessionMessageId - Only return monitor alerts linked to one of these session messages.
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionConnectionId - Only return monitor alerts linked to one of these session connections.
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Only return monitor alerts linked to one of these provider runs.
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// CreatedAt - Only return results whose monitor alert creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// ResolvedAt - Only return results whose monitor alert resolution time is in this range. Set `gt`, `lt`, or both.
	ResolvedAt *map[string]any `json:"resolved_at,omitempty"`
}

// List returns a paginated list of monitor alerts in the instance.
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

// Get retrieves a monitor alert.
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
