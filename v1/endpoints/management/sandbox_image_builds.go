package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/imagebuilds"
)

// SandboxImageBuildsEndpoint provides access to builds turn a Dockerfile and its context into an image version and replicate it to every region.
type SandboxImageBuildsEndpoint struct {
	client *endpoint.Client
}

// NewSandboxImageBuildsEndpoint creates a new SandboxImageBuildsEndpoint.
func NewSandboxImageBuildsEndpoint(client *endpoint.Client) *SandboxImageBuildsEndpoint {
	return &SandboxImageBuildsEndpoint{client: client}
}

// SandboxImageBuildsEndpointListParams contains optional query parameters for List.
type SandboxImageBuildsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ImageId - Filter by image
	ImageId *any `json:"image_id,omitempty"`
	// ImageVersionId - Filter by image version
	ImageVersionId *any `json:"image_version_id,omitempty"`
	// Status - Filter by build status
	Status *any `json:"status,omitempty"`
}

// SandboxImageBuildsEndpointLogsParams contains optional query parameters for Logs.
type SandboxImageBuildsEndpointLogsParams struct {
	// Direction - Whether to read from the cursor toward newer or older lines
	Direction *string `json:"direction,omitempty"`
	// Cursor - The line index to start from
	Cursor *float64 `json:"cursor,omitempty"`
	// Limit - The maximum number of lines to return
	Limit *float64 `json:"limit,omitempty"`
}

// List returns a paginated list of image builds across all images.
func (e *SandboxImageBuildsEndpoint) List(instanceId string, params *SandboxImageBuildsEndpointListParams) (*imagebuilds.SandboxImageBuildsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "image-builds"},
		Query: query,
	}
	var result imagebuilds.SandboxImageBuildsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an image build by ID.
func (e *SandboxImageBuildsEndpoint) Get(instanceId string, buildId string) (*imagebuilds.SandboxImageBuildsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "image-builds", buildId},
	}
	var result imagebuilds.SandboxImageBuildsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel cancels a build that is still running.
func (e *SandboxImageBuildsEndpoint) Cancel(instanceId string, buildId string) (*imagebuilds.SandboxImageBuildsCancelOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "sandbox", "image-builds", buildId, "cancel"},
	}
	var result imagebuilds.SandboxImageBuildsCancelOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Logs returns a page of the build output.
func (e *SandboxImageBuildsEndpoint) Logs(instanceId string, buildId string, params *SandboxImageBuildsEndpointLogsParams) (*imagebuilds.SandboxImageBuildsLogsOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "sandbox", "image-builds", buildId, "logs"},
		Query: query,
	}
	var result imagebuilds.SandboxImageBuildsLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
