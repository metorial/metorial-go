package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/monitors"
)

// MonitorsEndpoint provides access to monitors track automated observability checks for this instance.
type MonitorsEndpoint struct {
	client *endpoint.Client
}

// NewMonitorsEndpoint creates a new MonitorsEndpoint.
func NewMonitorsEndpoint(client *endpoint.Client) *MonitorsEndpoint {
	return &MonitorsEndpoint{client: client}
}

// MonitorsEndpointListParams contains optional query parameters for List.
type MonitorsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by monitor ID(s)
	Id                 *any    `json:"id,omitempty"`
	Target             *any    `json:"target,omitempty"`
	Status             *any    `json:"status,omitempty"`
	ProviderId         *any    `json:"provider_id,omitempty"`
	ProtoGuardFilterId *any    `json:"proto_guard_filter_id,omitempty"`
	Search             *string `json:"search,omitempty"`
	// CreatedAt - Filter monitor creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter monitor last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
	// FirstAlertAt - Filter first monitor alert time by date range
	FirstAlertAt *map[string]any `json:"first_alert_at,omitempty"`
	// LastAlertAt - Filter last monitor alert time by date range
	LastAlertAt *map[string]any `json:"last_alert_at,omitempty"`
}

// List returns a paginated list of monitors for this instance.
func (e *MonitorsEndpoint) List(instanceId string, params *MonitorsEndpointListParams) (*monitors.MonitorsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "monitors"},
		Query: query,
	}
	var result monitors.MonitorsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a monitor by ID.
func (e *MonitorsEndpoint) Get(instanceId string, monitorId string) (*monitors.MonitorsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "monitors", monitorId},
	}
	var result monitors.MonitorsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
