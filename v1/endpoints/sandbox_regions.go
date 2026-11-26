package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/regions"
)

// SandboxRegionsEndpoint provides access to regions in which sandbox machines, disks, and images are available.
type SandboxRegionsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxRegionsEndpoint creates a new SandboxRegionsEndpoint.
func NewSandboxRegionsEndpoint(client *endpoint.Client) *SandboxRegionsEndpoint {
	return &SandboxRegionsEndpoint{client: client}
}

// SandboxRegionsEndpointListParams contains optional query parameters for List.
type SandboxRegionsEndpointListParams struct {
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
}

// List returns a paginated list of regions sandboxes can run in.
func (e *SandboxRegionsEndpoint) List(params *SandboxRegionsEndpointListParams) (*regions.SandboxRegionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "regions"},
		Query: query,
	}
	var result regions.SandboxRegionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a sandbox region by its identifier.
func (e *SandboxRegionsEndpoint) Get(regionId string) (*regions.SandboxRegionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "regions", regionId},
	}
	var result regions.SandboxRegionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
