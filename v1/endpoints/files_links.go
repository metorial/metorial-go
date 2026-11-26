package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/files/links"
)

// FilesLinksEndpoint provides access to a file link is a public URL for a file. Anyone with the link can download the file without authenticating, until the link expires or is deleted.
type FilesLinksEndpoint struct {
	client *endpoint.Client
}

// NewFilesLinksEndpoint creates a new FilesLinksEndpoint.
func NewFilesLinksEndpoint(client *endpoint.Client) *FilesLinksEndpoint {
	return &FilesLinksEndpoint{client: client}
}

// FilesLinksEndpointListParams contains optional query parameters for List.
type FilesLinksEndpointListParams struct {
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
	// FileId - Only return links for this file.
	FileId *string `json:"file_id,omitempty"`
}

// FilesLinksEndpointCreateBody contains the request body for Create.
type FilesLinksEndpointCreateBody struct {
	// FileId - File to share, by ID.
	FileId string `json:"file_id"`
	// ExpiresAt - When the link stops working. The link never expires when omitted.
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// List returns a paginated list of file links.
func (e *FilesLinksEndpoint) List(params *FilesLinksEndpointListParams) (*links.FilesLinksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"file-links"},
		Query: query,
	}
	var result links.FilesLinksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a file link.
func (e *FilesLinksEndpoint) Get(linkId string) (*links.FilesLinksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"file-links", linkId},
	}
	var result links.FilesLinksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a link for a file.
func (e *FilesLinksEndpoint) Create(body *FilesLinksEndpointCreateBody) (*links.FilesLinksCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"file-links"},
		Body: body,
	}
	var result links.FilesLinksCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a file link. The file itself is kept.
func (e *FilesLinksEndpoint) Delete(linkId string) (*links.FilesLinksDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"file-links", linkId},
	}
	var result links.FilesLinksDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
