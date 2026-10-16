package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/images/builds"
)

// SandboxImagesBuildsEndpoint provides access to images machines boot from. Metorial provides global images, and you can build your own from a Dockerfile.
type SandboxImagesBuildsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxImagesBuildsEndpoint creates a new SandboxImagesBuildsEndpoint.
func NewSandboxImagesBuildsEndpoint(client *endpoint.Client) *SandboxImagesBuildsEndpoint {
	return &SandboxImagesBuildsEndpoint{client: client}
}

// SandboxImagesBuildsEndpointListParams contains optional query parameters for List.
type SandboxImagesBuildsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ImageVersionId - Filter by image version
	ImageVersionId *any `json:"image_version_id,omitempty"`
	// Status - Filter by build status
	Status *any `json:"status,omitempty"`
}

// SandboxImagesBuildsEndpointCreateBody contains the request body for Create.
type SandboxImagesBuildsEndpointCreateBody struct {
	// FileId - The file holding the build context as a zip, tar, or tar.gz archive. Limited to 50 MB.
	FileId string `json:"file_id"`
	// DockerfilePath - The path of the Dockerfile inside the context
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// ContextPath - The directory inside the archive to build from
	ContextPath *string `json:"context_path,omitempty"`
	// Target - The build stage to build
	Target *string `json:"target,omitempty"`
	// BuildArgs - Build arguments passed to the Dockerfile
	BuildArgs *map[string]string `json:"build_args,omitempty"`
	// TimeoutSeconds - How long the build may run before it is cancelled
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
	// DiskSizeBytes - The size of disks created from the image
	DiskSizeBytes *float64 `json:"disk_size_bytes,omitempty"`
}

// List returns a paginated list of the builds of an image.
func (e *SandboxImagesBuildsEndpoint) List(instanceId string, imageId string, params *SandboxImagesBuildsEndpointListParams) (*builds.SandboxImagesBuildsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "images", imageId, "builds"},
		Query: query,
	}
	var result builds.SandboxImagesBuildsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create builds a new version of an image from a Dockerfile. Upload the build context as a file with the sandbox_image_build_context purpose first.
func (e *SandboxImagesBuildsEndpoint) Create(instanceId string, imageId string, body *SandboxImagesBuildsEndpointCreateBody) (*builds.SandboxImagesBuildsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "images", imageId, "builds"},
		Body: body,
	}
	var result builds.SandboxImagesBuildsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
