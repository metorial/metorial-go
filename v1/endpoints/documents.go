package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents"
)

// DocumentsEndpoint provides access to a document is rich text with comments, versions, and participants.
type DocumentsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsEndpoint creates a new DocumentsEndpoint.
func NewDocumentsEndpoint(client *endpoint.Client) *DocumentsEndpoint {
	return &DocumentsEndpoint{client: client}
}

// DocumentsEndpointListParams contains optional query parameters for List.
type DocumentsEndpointListParams struct {
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
	// Id - Only return documents with one of these IDs.
	Id *any `json:"id,omitempty"`
	// FileId - Only return documents stored in one of these files.
	FileId *any `json:"file_id,omitempty"`
	// StoreId - Only return documents in one of these stores.
	StoreId *any `json:"store_id,omitempty"`
	// ParentDocumentId - Only return documents cloned from one of these documents.
	ParentDocumentId *any `json:"parent_document_id,omitempty"`
	// CreatedAt - Only return results whose document creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose document last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// DocumentsEndpointCreateBody contains the request body for Create.
type DocumentsEndpointCreateBody struct {
	// Title - Title of the document.
	Title string `json:"title"`
	// Content - Content of the document.
	Content string `json:"content"`
}

// DocumentsEndpointUpdateBody contains the request body for Update.
type DocumentsEndpointUpdateBody struct {
	// Title - Title of the document.
	Title *string `json:"title,omitempty"`
	// Content - Content of the document.
	Content *string `json:"content,omitempty"`
}

// DocumentsEndpointCloneBody contains the request body for Clone.
type DocumentsEndpointCloneBody struct {
	// TargetDocumentId - ID to use for the cloned document.
	TargetDocumentId *string `json:"target_document_id,omitempty"`
	// Title - Title of the cloned document.
	Title *string `json:"title,omitempty"`
}

// List returns a paginated list of documents.
func (e *DocumentsEndpoint) List(params *DocumentsEndpointListParams) (*documents.DocumentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"documents"},
		Query: query,
	}
	var result documents.DocumentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a document.
func (e *DocumentsEndpoint) Create(body *DocumentsEndpointCreateBody) (*documents.DocumentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents"},
		Body: body,
	}
	var result documents.DocumentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a document.
func (e *DocumentsEndpoint) Get(documentId string) (*documents.DocumentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId},
	}
	var result documents.DocumentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a document.
func (e *DocumentsEndpoint) Update(documentId string, body *DocumentsEndpointUpdateBody) (*documents.DocumentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId},
		Body: body,
	}
	var result documents.DocumentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a document.
func (e *DocumentsEndpoint) Delete(documentId string) (*documents.DocumentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId},
	}
	var result documents.DocumentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone creates a copy of a document.
func (e *DocumentsEndpoint) Clone(documentId string, body *DocumentsEndpointCloneBody) (*documents.DocumentsCloneOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "clone"},
		Body: body,
	}
	var result documents.DocumentsCloneOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
