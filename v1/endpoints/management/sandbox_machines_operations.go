package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/operations"
)

// SandboxMachinesOperationsEndpoint provides access to operations record each start, pause, hibernate, restart, terminate, update, or upgrade a machine went through.
type SandboxMachinesOperationsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesOperationsEndpoint creates a new SandboxMachinesOperationsEndpoint.
func NewSandboxMachinesOperationsEndpoint(client *endpoint.Client) *SandboxMachinesOperationsEndpoint {
	return &SandboxMachinesOperationsEndpoint{client: client}
}

// SandboxMachinesOperationsEndpointListParams contains optional query parameters for List.
type SandboxMachinesOperationsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// MachineInstanceId - Only operations that involved these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	Kind              *any `json:"kind,omitempty"`
	Trigger           *any `json:"trigger,omitempty"`
	Outcome           *any `json:"outcome,omitempty"`
}

// List returns a paginated list of the operations of a machine.
func (e *SandboxMachinesOperationsEndpoint) List(instanceId string, machineId string, params *SandboxMachinesOperationsEndpointListParams) (*operations.SandboxMachinesOperationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "machines", machineId, "operations"},
		Query: query,
	}
	var result operations.SandboxMachinesOperationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an operation of a machine.
func (e *SandboxMachinesOperationsEndpoint) Get(instanceId string, machineId string, operationId string) (*operations.SandboxMachinesOperationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "machines", machineId, "operations", operationId},
	}
	var result operations.SandboxMachinesOperationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
