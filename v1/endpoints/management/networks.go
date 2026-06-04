package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/networks"
)

// NetworksEndpoint provides access to read network records for an instance environment.
type NetworksEndpoint struct {
	client *endpoint.Client
}

// NewNetworksEndpoint creates a new NetworksEndpoint.
func NewNetworksEndpoint(client *endpoint.Client) *NetworksEndpoint {
	return &NetworksEndpoint{client: client}
}

// NetworksEndpointListParams contains optional query parameters for List.
type NetworksEndpointListParams struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	EnclaveId  *any     `json:"enclave_id,omitempty"`
	// CreatedAt - Filter network creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter network last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// NetworksEndpointListNetworkLogsParams contains optional query parameters for ListNetworkLogs.
type NetworksEndpointListNetworkLogsParams struct {
	Direction       string   `json:"direction"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	Hostname        *any     `json:"hostname,omitempty"`
	Ip              *any     `json:"ip,omitempty"`
	From            *string  `json:"from,omitempty"`
	To              *string  `json:"to,omitempty"`
	IntervalMinutes *float64 `json:"interval_minutes,omitempty"`
}

// List returns a paginated list of networks.
func (e *NetworksEndpoint) List(instanceId string, params *NetworksEndpointListParams) (*networks.NetworksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "networks"},
		Query: query,
	}
	var result networks.NetworksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific network by ID.
func (e *NetworksEndpoint) Get(instanceId string, networkId string) (*networks.NetworksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "networks", networkId},
	}
	var result networks.NetworksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListNetworkLogs returns ingress or egress network logs for enclaves in the instance environment.
func (e *NetworksEndpoint) ListNetworkLogs(instanceId string, params *NetworksEndpointListNetworkLogsParams) (*networks.NetworksListNetworkLogsOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "network-logs"},
		Query: query,
	}
	var result networks.NetworksListNetworkLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
