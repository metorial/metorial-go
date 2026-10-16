package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/events"
)

// SandboxMachinesEventsEndpoint provides access to the timeline of everything that happened to a machine and its instances.
type SandboxMachinesEventsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesEventsEndpoint creates a new SandboxMachinesEventsEndpoint.
func NewSandboxMachinesEventsEndpoint(client *endpoint.Client) *SandboxMachinesEventsEndpoint {
	return &SandboxMachinesEventsEndpoint{client: client}
}

// SandboxMachinesEventsEndpointListParams contains optional query parameters for List.
type SandboxMachinesEventsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Scope - Whether to include machine-level events, instance-level events, or both
	Scope *string `json:"scope,omitempty"`
	// MachineInstanceId - Only events of these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	// Type - Filter by event type
	Type *any `json:"type,omitempty"`
	// Source - Filter by what caused the event
	Source *any `json:"source,omitempty"`
	// Reverse - List the oldest events first
	Reverse *bool `json:"reverse,omitempty"`
}

// List returns a paginated list of the events of a machine.
func (e *SandboxMachinesEventsEndpoint) List(instanceId string, machineId string, params *SandboxMachinesEventsEndpointListParams) (*events.SandboxMachinesEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "machines", machineId, "events"},
		Query: query,
	}
	var result events.SandboxMachinesEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an event of a machine.
func (e *SandboxMachinesEventsEndpoint) Get(instanceId string, machineId string, eventId string) (*events.SandboxMachinesEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "machines", machineId, "events", eventId},
	}
	var result events.SandboxMachinesEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
