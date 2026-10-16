package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/instances"
)

// SandboxMachinesInstancesEndpoint provides access to instances are the concrete virtual machines a sandbox machine runs on over its lifetime.
type SandboxMachinesInstancesEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesInstancesEndpoint creates a new SandboxMachinesInstancesEndpoint.
func NewSandboxMachinesInstancesEndpoint(client *endpoint.Client) *SandboxMachinesInstancesEndpoint {
	return &SandboxMachinesInstancesEndpoint{client: client}
}

// SandboxMachinesInstancesEndpointListParams contains optional query parameters for List.
type SandboxMachinesInstancesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Status - Filter by instance status
	Status *any `json:"status,omitempty"`
}

// SandboxMachinesInstancesEndpointLogsParams contains optional query parameters for Logs.
type SandboxMachinesInstancesEndpointLogsParams struct {
	// Direction - Whether to read from the cursor toward newer or older lines
	Direction *string `json:"direction,omitempty"`
	// Cursor - The line index to start from
	Cursor *float64 `json:"cursor,omitempty"`
	// Limit - The maximum number of lines to return
	Limit *float64 `json:"limit,omitempty"`
}

// List returns a paginated list of the instances of a machine.
func (e *SandboxMachinesInstancesEndpoint) List(machineId string, params *SandboxMachinesInstancesEndpointListParams) (*instances.SandboxMachinesInstancesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines", machineId, "instances"},
		Query: query,
	}
	var result instances.SandboxMachinesInstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an instance of a machine.
func (e *SandboxMachinesInstancesEndpoint) Get(machineId string, machineInstanceId string) (*instances.SandboxMachinesInstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "instances", machineInstanceId},
	}
	var result instances.SandboxMachinesInstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Logs returns a page of the console output of an instance.
func (e *SandboxMachinesInstancesEndpoint) Logs(machineId string, machineInstanceId string, params *SandboxMachinesInstancesEndpointLogsParams) (*instances.SandboxMachinesInstancesLogsOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines", machineId, "instances", machineInstanceId, "logs"},
		Query: query,
	}
	var result instances.SandboxMachinesInstancesLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
