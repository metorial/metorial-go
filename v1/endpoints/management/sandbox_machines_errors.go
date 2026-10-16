package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/errors"
)

// SandboxMachinesErrorsEndpoint provides access to errors reported from inside a machine, such as a failing lifecycle hook or setup command.
type SandboxMachinesErrorsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesErrorsEndpoint creates a new SandboxMachinesErrorsEndpoint.
func NewSandboxMachinesErrorsEndpoint(client *endpoint.Client) *SandboxMachinesErrorsEndpoint {
	return &SandboxMachinesErrorsEndpoint{client: client}
}

// SandboxMachinesErrorsEndpointListParams contains optional query parameters for List.
type SandboxMachinesErrorsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// MachineInstanceId - Only errors of these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
}

// List returns a paginated list of the errors reported by a machine.
func (e *SandboxMachinesErrorsEndpoint) List(instanceId string, machineId string, params *SandboxMachinesErrorsEndpointListParams) (*errors.SandboxMachinesErrorsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "machines", machineId, "errors"},
		Query: query,
	}
	var result errors.SandboxMachinesErrorsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an error reported by a machine.
func (e *SandboxMachinesErrorsEndpoint) Get(instanceId string, machineId string, errorId string) (*errors.SandboxMachinesErrorsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "machines", machineId, "errors", errorId},
	}
	var result errors.SandboxMachinesErrorsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
