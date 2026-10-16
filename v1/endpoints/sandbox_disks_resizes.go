package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/disks/resizes"
)

// SandboxDisksResizesEndpoint provides access to resizing a disk briefly stops the machines it is mounted into. Each resize is tracked until it completes.
type SandboxDisksResizesEndpoint struct {
	client *endpoint.Client
}

// NewSandboxDisksResizesEndpoint creates a new SandboxDisksResizesEndpoint.
func NewSandboxDisksResizesEndpoint(client *endpoint.Client) *SandboxDisksResizesEndpoint {
	return &SandboxDisksResizesEndpoint{client: client}
}

// SandboxDisksResizesEndpointListParams contains optional query parameters for List.
type SandboxDisksResizesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Status - Filter by resize status
	Status *any `json:"status,omitempty"`
}

// List returns a paginated list of the resizes of a disk.
func (e *SandboxDisksResizesEndpoint) List(diskId string, params *SandboxDisksResizesEndpointListParams) (*resizes.SandboxDisksResizesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "disks", diskId, "resizes"},
		Query: query,
	}
	var result resizes.SandboxDisksResizesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a resize of a disk.
func (e *SandboxDisksResizesEndpoint) Get(diskId string, resizeId string) (*resizes.SandboxDisksResizesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "disks", diskId, "resizes", resizeId},
	}
	var result resizes.SandboxDisksResizesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Retry retries a resize that failed.
func (e *SandboxDisksResizesEndpoint) Retry(diskId string, resizeId string) (*resizes.SandboxDisksResizesRetryOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "disks", diskId, "resizes", resizeId, "retry"},
	}
	var result resizes.SandboxDisksResizesRetryOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
