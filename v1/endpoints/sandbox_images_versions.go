package endpoints

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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns a paginated list of the versions of an image.
func (e *SandboxImagesVersionsEndpoint) List(imageId string, params *SandboxImagesVersionsEndpointListParams) (*versions.SandboxImagesVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "images", imageId, "versions"},
		Query: query,
	}
	var result versions.SandboxImagesVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a version of an image.
func (e *SandboxImagesVersionsEndpoint) Get(imageId string, versionId string) (*versions.SandboxImagesVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "images", imageId, "versions", versionId},
	}
	var result versions.SandboxImagesVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
