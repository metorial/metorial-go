package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/monitors"
)

// MonitorsEndpoint provides access to a monitor is an automated observability check for an instance. It raises alerts when its target misbehaves.
type MonitorsEndpoint struct {
	client *endpoint.Client
}

// NewMonitorsEndpoint creates a new MonitorsEndpoint.
func NewMonitorsEndpoint(client *endpoint.Client) *MonitorsEndpoint {
	return &MonitorsEndpoint{client: client}
}

// MonitorsEndpointListParams contains optional query parameters for List.
type MonitorsEndpointListParams struct {
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
	// Id - Only return monitors with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return monitors watching one of these targets.
	Target *any `json:"target,omitempty"`
	// Status - Only return monitors with one of these statuses.
	Status *any `json:"status,omitempty"`
	// ProviderId - Only return monitors watching one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProtoGuardFilterId - Only return monitors watching one of these ProtoGuard filters.
	ProtoGuardFilterId *any `json:"proto_guard_filter_id,omitempty"`
	// Search - Filter monitors by name or description.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose monitor creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose monitor last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
	// FirstAlertAt - Only return results whose first monitor alert time is in this range. Set `gt`, `lt`, or both.
	FirstAlertAt *map[string]any `json:"first_alert_at,omitempty"`
	// LastAlertAt - Only return results whose last monitor alert time is in this range. Set `gt`, `lt`, or both.
	LastAlertAt *map[string]any `json:"last_alert_at,omitempty"`
}

// List returns a paginated list of monitors in the instance.
func (e *MonitorsEndpoint) List(params *MonitorsEndpointListParams) (*monitors.MonitorsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"monitors"},
		Query: query,
	}
	var result monitors.MonitorsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a monitor.
func (e *MonitorsEndpoint) Get(monitorId string) (*monitors.MonitorsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"monitors", monitorId},
	}
	var result monitors.MonitorsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
