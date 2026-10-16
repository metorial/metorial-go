package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines"
)

// SandboxMachinesEndpoint provides access to create and operate sandbox machines. A machine boots from an image, can mount disks, and moves between running, paused, and hibernated states.
type SandboxMachinesEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesEndpoint creates a new SandboxMachinesEndpoint.
func NewSandboxMachinesEndpoint(client *endpoint.Client) *SandboxMachinesEndpoint {
	return &SandboxMachinesEndpoint{client: client}
}

// SandboxMachinesEndpointListParams contains optional query parameters for List.
type SandboxMachinesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Status - Filter by whether the machine still exists (active) or was terminated (deleted). Defaults to active.
	Status *any `json:"status,omitempty"`
	// State - Filter by the machine's current runtime status, the status field of the machine. Terminated machines are only listed together with status=deleted.
	State *any `json:"state,omitempty"`
	// ImageId - Filter by the image the machines boot from
	ImageId *any `json:"image_id,omitempty"`
	// Region - Filter by the identifiers of the regions the machines run in
	Region *any `json:"region,omitempty"`
	// Search - Search machines by name, slug, or ID
	Search *string `json:"search,omitempty"`
}

// SandboxMachinesEndpointCreateBody contains the request body for Create.
type SandboxMachinesEndpointCreateBody struct {
	// Region - The identifier of the region to create the machine in
	Region string `json:"region"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the machine
	Name *string `json:"name,omitempty"`
	// ImageId - The image to boot the machine from
	ImageId string `json:"image_id"`
	// SpecId - The ID of the spec that sets the size of the machine
	SpecId *string `json:"spec_id,omitempty"`
	// SpecIdentifier - The identifier of the spec that sets the size of the machine
	SpecIdentifier *string `json:"spec_identifier,omitempty"`
	// MemoryGib - The memory of the machine in GiB, in steps of 0.5. Use together with cpus instead of a spec ID.
	MemoryGib *float64 `json:"memory_gib,omitempty"`
	// Cpus - The CPUs of the machine, in steps of 0.5. Use together with memory_gib instead of a spec ID.
	Cpus            *float64        `json:"cpus,omitempty"`
	LifecyclePolicy *map[string]any `json:"lifecycle_policy,omitempty"`
	Runtime         *map[string]any `json:"runtime,omitempty"`
	LifecycleHooks  *[]any          `json:"lifecycle_hooks,omitempty"`
	// IgnoreImageHooks - Ignore the image's own lifecycle hooks
	IgnoreImageHooks *bool `json:"ignore_image_hooks,omitempty"`
	// EntrypointWaitSeconds - How long startup waits for the entrypoint
	EntrypointWaitSeconds *float64          `json:"entrypoint_wait_seconds,omitempty"`
	Mounts                *[]map[string]any `json:"mounts,omitempty"`
	// FirewallIds - Firewalls to apply to the machine
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// PublicPorts - Ports reachable without an access token
	PublicPorts *[]float64 `json:"public_ports,omitempty"`
	// Start - Whether to start the machine right away
	Start *bool `json:"start,omitempty"`
}

// SandboxMachinesEndpointUpdateBody contains the request body for Update.
type SandboxMachinesEndpointUpdateBody struct {
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the machine
	Name *string `json:"name,omitempty"`
	// SpecId - The ID of the spec that sets the size of the machine
	SpecId *string `json:"spec_id,omitempty"`
	// SpecIdentifier - The identifier of the spec that sets the size of the machine
	SpecIdentifier *string `json:"spec_identifier,omitempty"`
	// MemoryGib - The memory of the machine in GiB, in steps of 0.5. Use together with cpus instead of a spec ID.
	MemoryGib *float64 `json:"memory_gib,omitempty"`
	// Cpus - The CPUs of the machine, in steps of 0.5. Use together with memory_gib instead of a spec ID.
	Cpus            *float64        `json:"cpus,omitempty"`
	LifecyclePolicy *map[string]any `json:"lifecycle_policy,omitempty"`
	Runtime         *map[string]any `json:"runtime,omitempty"`
	LifecycleHooks  *[]any          `json:"lifecycle_hooks,omitempty"`
	// IgnoreImageHooks - Ignore the image's own lifecycle hooks
	IgnoreImageHooks *bool `json:"ignore_image_hooks,omitempty"`
	// EntrypointWaitSeconds - How long startup waits for the entrypoint
	EntrypointWaitSeconds *float64          `json:"entrypoint_wait_seconds,omitempty"`
	Mounts                *[]map[string]any `json:"mounts,omitempty"`
	// FirewallIds - Firewalls to apply to the machine
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// PublicPorts - Ports reachable without an access token
	PublicPorts *[]float64 `json:"public_ports,omitempty"`
}

// SandboxMachinesEndpointUsageParams contains optional query parameters for Usage.
type SandboxMachinesEndpointUsageParams struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	MachineInstanceId *any     `json:"machine_instance_id,omitempty"`
}

// List returns a paginated list of machines.
func (e *SandboxMachinesEndpoint) List(params *SandboxMachinesEndpointListParams) (*machines.SandboxMachinesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines"},
		Query: query,
	}
	var result machines.SandboxMachinesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a machine by ID.
func (e *SandboxMachinesEndpoint) Get(machineId string) (*machines.SandboxMachinesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId},
	}
	var result machines.SandboxMachinesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a machine and, unless disabled, starts it.
func (e *SandboxMachinesEndpoint) Create(body *SandboxMachinesEndpointCreateBody) (*machines.SandboxMachinesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines"},
		Body: body,
	}
	var result machines.SandboxMachinesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a machine. Changes to its size, runtime, or mounts take effect once the machine restarts, which happens on its own for a running machine, and the machine reports restart_required until then.
func (e *SandboxMachinesEndpoint) Update(machineId string, body *SandboxMachinesEndpointUpdateBody) (*machines.SandboxMachinesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId},
		Body: body,
	}
	var result machines.SandboxMachinesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete terminates a machine and deletes its root disk.
func (e *SandboxMachinesEndpoint) Delete(machineId string) (*machines.SandboxMachinesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId},
	}
	var result machines.SandboxMachinesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Start starts a paused or hibernated machine.
func (e *SandboxMachinesEndpoint) Start(machineId string) (*machines.SandboxMachinesStartOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "start"},
	}
	var result machines.SandboxMachinesStartOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Pause pauses a running machine and keeps its memory state.
func (e *SandboxMachinesEndpoint) Pause(machineId string) (*machines.SandboxMachinesPauseOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "pause"},
	}
	var result machines.SandboxMachinesPauseOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Hibernate hibernates a machine so it stops consuming compute resources.
func (e *SandboxMachinesEndpoint) Hibernate(machineId string) (*machines.SandboxMachinesHibernateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "hibernate"},
	}
	var result machines.SandboxMachinesHibernateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Restart restarts a machine on a fresh instance.
func (e *SandboxMachinesEndpoint) Restart(machineId string) (*machines.SandboxMachinesRestartOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "restart"},
	}
	var result machines.SandboxMachinesRestartOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Touch records activity so the machine is not paused for being idle.
func (e *SandboxMachinesEndpoint) Touch(machineId string) (*machines.SandboxMachinesTouchOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "touch"},
	}
	var result machines.SandboxMachinesTouchOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Usage returns resource usage samples of a machine, newest first.
func (e *SandboxMachinesEndpoint) Usage(machineId string, params *SandboxMachinesEndpointUsageParams) (*machines.SandboxMachinesUsageOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines", machineId, "usage"},
		Query: query,
	}
	var result machines.SandboxMachinesUsageOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
