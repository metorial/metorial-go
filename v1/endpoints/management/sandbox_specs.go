package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/specs"
)

// SandboxSpecsEndpoint provides access to the CPU and memory sizes sandbox machines can be created with.
type SandboxSpecsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxSpecsEndpoint creates a new SandboxSpecsEndpoint.
func NewSandboxSpecsEndpoint(client *endpoint.Client) *SandboxSpecsEndpoint {
	return &SandboxSpecsEndpoint{client: client}
}

// SandboxSpecsEndpointListParams contains optional query parameters for List.
type SandboxSpecsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of machine specs. The available field shows whether your plan allows creating machines of that size.
func (e *SandboxSpecsEndpoint) List(instanceId string, params *SandboxSpecsEndpointListParams) (*specs.SandboxSpecsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "specs"},
		Query: query,
	}
	var result specs.SandboxSpecsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a machine spec by ID.
func (e *SandboxSpecsEndpoint) Get(instanceId string, specId string) (*specs.SandboxSpecsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "specs", specId},
	}
	var result specs.SandboxSpecsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
