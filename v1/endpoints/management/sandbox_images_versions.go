package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/images/versions"
)

// SandboxImagesVersionsEndpoint provides access to images machines boot from. Metorial provides global images, and you can build your own from a Dockerfile.
type SandboxImagesVersionsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxImagesVersionsEndpoint creates a new SandboxImagesVersionsEndpoint.
func NewSandboxImagesVersionsEndpoint(client *endpoint.Client) *SandboxImagesVersionsEndpoint {
	return &SandboxImagesVersionsEndpoint{client: client}
}

// SandboxImagesVersionsEndpointListParams contains optional query parameters for List.
type SandboxImagesVersionsEndpointListParams struct {
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

// List returns a paginated list of the versions of an image.
func (e *SandboxImagesVersionsEndpoint) List(instanceId string, imageId string, params *SandboxImagesVersionsEndpointListParams) (*versions.SandboxImagesVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "images", imageId, "versions"},
		Query: query,
	}
	var result versions.SandboxImagesVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a version of an image.
func (e *SandboxImagesVersionsEndpoint) Get(instanceId string, imageId string, versionId string) (*versions.SandboxImagesVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images", imageId, "versions", versionId},
	}
	var result versions.SandboxImagesVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
