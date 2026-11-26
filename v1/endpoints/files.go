package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/files"
)

// FilesEndpoint provides access to a file is content uploaded to an instance, such as an image or an attachment. Its purpose controls what it can be linked to.
type FilesEndpoint struct {
	client *endpoint.Client
}

// NewFilesEndpoint creates a new FilesEndpoint.
func NewFilesEndpoint(client *endpoint.Client) *FilesEndpoint {
	return &FilesEndpoint{client: client}
}

// FilesEndpointListParams contains optional query parameters for List.
type FilesEndpointListParams struct {
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
	// Id - Only return files with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Purpose - Only return files with one of these purposes.
	Purpose *any `json:"purpose,omitempty"`
	// StoreId - Only return files in one of these stores.
	StoreId *any `json:"store_id,omitempty"`
	// DocumentId - Only return files attached to one of these documents.
	DocumentId *any `json:"document_id,omitempty"`
	// FileLinkId - Only return files shared through one of these file links.
	FileLinkId *any `json:"file_link_id,omitempty"`
	// CreatedAt - Only return results whose file creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose file last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of files.
func (e *FilesEndpoint) List(params *FilesEndpointListParams) (*files.FilesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"files"},
		Query: query,
	}
	var result files.FilesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a file.
func (e *FilesEndpoint) Get(fileId string) (*files.FilesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"files", fileId},
	}
	var result files.FilesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a file.
func (e *FilesEndpoint) Delete(fileId string) (*files.FilesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"files", fileId},
	}
	var result files.FilesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
