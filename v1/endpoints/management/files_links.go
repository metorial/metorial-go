package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/files/links"
)

// FilesLinksEndpoint provides access to files are private by default. If you want to share a file, you can create a link for it. Links are public and do not require authentication to access, so be careful with what you share.
type FilesLinksEndpoint struct {
	client *endpoint.Client
}

// NewFilesLinksEndpoint creates a new FilesLinksEndpoint.
func NewFilesLinksEndpoint(client *endpoint.Client) *FilesLinksEndpoint {
	return &FilesLinksEndpoint{client: client}
}

// FilesLinksEndpointListParams contains optional query parameters for List.
type FilesLinksEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// FileId - Filter by file ID
	FileId *string `json:"file_id,omitempty"`
}

// FilesLinksEndpointCreateBody contains the request body for Create.
type FilesLinksEndpointCreateBody struct {
	FileId    string  `json:"file_id"`
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// List returns a paginated list of file links owned by the instance organization.
func (e *FilesLinksEndpoint) List(instanceId string, params *FilesLinksEndpointListParams) (*links.FilesLinksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "file-links"},
		Query: query,
	}
	var result links.FilesLinksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves the details of a specific file link by its ID.
func (e *FilesLinksEndpoint) Get(instanceId string, linkId string) (*links.FilesLinksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "file-links", linkId},
	}
	var result links.FilesLinksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new link for a specific file.
func (e *FilesLinksEndpoint) Create(instanceId string, body *FilesLinksEndpointCreateBody) (*links.FilesLinksCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "file-links"},
		Body: body,
	}
	var result links.FilesLinksCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a specific file link by its ID.
func (e *FilesLinksEndpoint) Delete(instanceId string, linkId string) (*links.FilesLinksDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "file-links", linkId},
	}
	var result links.FilesLinksDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
