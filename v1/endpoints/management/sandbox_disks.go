package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/disks"
)

// SandboxDisksEndpoint provides access to persistent disks that can be mounted into machines. Every machine also has a root filesystem disk that is managed for it.
type SandboxDisksEndpoint struct {
	client *endpoint.Client
}

// NewSandboxDisksEndpoint creates a new SandboxDisksEndpoint.
func NewSandboxDisksEndpoint(client *endpoint.Client) *SandboxDisksEndpoint {
	return &SandboxDisksEndpoint{client: client}
}

// SandboxDisksEndpointListParams contains optional query parameters for List.
type SandboxDisksEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Type - Filter by disk type
	Type *any `json:"type,omitempty"`
	// MachineId - Filter by the machines the disks are mounted into
	MachineId *any `json:"machine_id,omitempty"`
	// Region - Filter by the identifiers of the regions the disks live in
	Region *any `json:"region,omitempty"`
	// Attached - Only disks used by an active machine (true), as its root filesystem or a mount, or only disks no machine uses (false)
	Attached *bool `json:"attached,omitempty"`
	// Search - Search disks by name or ID
	Search *string `json:"search,omitempty"`
}

// SandboxDisksEndpointCreateBody contains the request body for Create.
type SandboxDisksEndpointCreateBody struct {
	// Region - The identifier of the region to create the disk in
	Region string `json:"region"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the disk
	Name *string `json:"name,omitempty"`
	// SizeBytes - The capacity of the disk in bytes
	SizeBytes *float64 `json:"size_bytes,omitempty"`
}

// SandboxDisksEndpointUpdateBody contains the request body for Update.
type SandboxDisksEndpointUpdateBody struct {
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the disk. Pass null to remove it.
	Name *string `json:"name,omitempty"`
}

// SandboxDisksEndpointCloneBody contains the request body for Clone.
type SandboxDisksEndpointCloneBody struct {
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the new disk
	Name *string `json:"name,omitempty"`
}

// SandboxDisksEndpointResizeBody contains the request body for Resize.
type SandboxDisksEndpointResizeBody struct {
	// SizeBytes - The new capacity of the disk in bytes
	SizeBytes float64 `json:"size_bytes"`
}

// List returns a paginated list of disks.
func (e *SandboxDisksEndpoint) List(instanceId string, params *SandboxDisksEndpointListParams) (*disks.SandboxDisksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "disks"},
		Query: query,
	}
	var result disks.SandboxDisksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a disk by ID.
func (e *SandboxDisksEndpoint) Get(instanceId string, diskId string) (*disks.SandboxDisksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks", diskId},
	}
	var result disks.SandboxDisksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an empty disk.
func (e *SandboxDisksEndpoint) Create(instanceId string, body *SandboxDisksEndpointCreateBody) (*disks.SandboxDisksCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks"},
		Body: body,
	}
	var result disks.SandboxDisksCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a disk name and metadata. Only metadata can be updated for root filesystem disks.
func (e *SandboxDisksEndpoint) Update(instanceId string, diskId string, body *SandboxDisksEndpointUpdateBody) (*disks.SandboxDisksUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks", diskId},
		Body: body,
	}
	var result disks.SandboxDisksUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a disk. Root filesystem disks cannot be deleted; they are removed with their machine.
func (e *SandboxDisksEndpoint) Delete(instanceId string, diskId string) (*disks.SandboxDisksDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks", diskId},
	}
	var result disks.SandboxDisksDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone creates a copy of a disk in the same region.
func (e *SandboxDisksEndpoint) Clone(instanceId string, diskId string, body *SandboxDisksEndpointCloneBody) (*disks.SandboxDisksCloneOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks", diskId, "clone"},
		Body: body,
	}
	var result disks.SandboxDisksCloneOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Resize starts growing a disk to a larger size.
func (e *SandboxDisksEndpoint) Resize(instanceId string, diskId string, body *SandboxDisksEndpointResizeBody) (*disks.SandboxDisksResizeOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "disks", diskId, "resize"},
		Body: body,
	}
	var result disks.SandboxDisksResizeOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
