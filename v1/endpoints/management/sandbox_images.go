package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/images"
)

// SandboxImagesEndpoint provides access to images machines boot from. Metorial provides global images, and you can build your own from a Dockerfile.
type SandboxImagesEndpoint struct {
	client *endpoint.Client
}

// NewSandboxImagesEndpoint creates a new SandboxImagesEndpoint.
func NewSandboxImagesEndpoint(client *endpoint.Client) *SandboxImagesEndpoint {
	return &SandboxImagesEndpoint{client: client}
}

// SandboxImagesEndpointListParams contains optional query parameters for List.
type SandboxImagesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Owner - Whether to list your own images, global images, or both (the default)
	Owner *any `json:"owner,omitempty"`
	// Status - Filter by image status
	Status *any `json:"status,omitempty"`
	// Search - Search by name or identifier
	Search *string `json:"search,omitempty"`
}

// SandboxImagesEndpointCreateBody contains the request body for Create.
type SandboxImagesEndpointCreateBody struct {
	// Name - A name for the image
	Name *string `json:"name,omitempty"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// SandboxImagesEndpointUpdateBody contains the request body for Update.
type SandboxImagesEndpointUpdateBody struct {
	// Name - A name for the image
	Name *string `json:"name,omitempty"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of the images you can create machines from.
func (e *SandboxImagesEndpoint) List(instanceId string, params *SandboxImagesEndpointListParams) (*images.SandboxImagesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "images"},
		Query: query,
	}
	var result images.SandboxImagesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an image by ID.
func (e *SandboxImagesEndpoint) Get(instanceId string, imageId string) (*images.SandboxImagesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images", imageId},
	}
	var result images.SandboxImagesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an empty image. Build a version of it to make it available to machines.
func (e *SandboxImagesEndpoint) Create(instanceId string, body *SandboxImagesEndpointCreateBody) (*images.SandboxImagesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images"},
		Body: body,
	}
	var result images.SandboxImagesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the name and metadata of one of your images.
func (e *SandboxImagesEndpoint) Update(instanceId string, imageId string, body *SandboxImagesEndpointUpdateBody) (*images.SandboxImagesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images", imageId},
		Body: body,
	}
	var result images.SandboxImagesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes one of your images and cancels its running builds.
func (e *SandboxImagesEndpoint) Delete(instanceId string, imageId string) (*images.SandboxImagesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images", imageId},
	}
	var result images.SandboxImagesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
