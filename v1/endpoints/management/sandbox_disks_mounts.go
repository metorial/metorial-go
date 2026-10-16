package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/disks/mounts"
)

// SandboxDisksMountsEndpoint provides access to persistent disks that can be mounted into machines. Every machine also has a root filesystem disk that is managed for it.
type SandboxDisksMountsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxDisksMountsEndpoint creates a new SandboxDisksMountsEndpoint.
func NewSandboxDisksMountsEndpoint(client *endpoint.Client) *SandboxDisksMountsEndpoint {
	return &SandboxDisksMountsEndpoint{client: client}
}

// SandboxDisksMountsEndpointListParams contains optional query parameters for List.
type SandboxDisksMountsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of the times a disk was mounted into a machine.
func (e *SandboxDisksMountsEndpoint) List(instanceId string, diskId string, params *SandboxDisksMountsEndpointListParams) (*mounts.SandboxDisksMountsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "disks", diskId, "mounts"},
		Query: query,
	}
	var result mounts.SandboxDisksMountsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
