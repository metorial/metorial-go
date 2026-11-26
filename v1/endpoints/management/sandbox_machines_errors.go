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
