package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents"
)

// DocumentsEndpoint provides access to create and manage instance documents backed by Cargo.
type DocumentsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsEndpoint creates a new DocumentsEndpoint.
func NewDocumentsEndpoint(client *endpoint.Client) *DocumentsEndpoint {
	return &DocumentsEndpoint{client: client}
}

// DocumentsEndpointListParams contains optional query parameters for List.
type DocumentsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document ID
	Id *any `json:"id,omitempty"`
	// FileId - Filter by file ID
	FileId *any `json:"file_id,omitempty"`
	// StoreId - Filter by store ID
	StoreId *any `json:"store_id,omitempty"`
	// ParentDocumentId - Filter by parent document ID
	ParentDocumentId *any `json:"parent_document_id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// DocumentsEndpointCreateBody contains the request body for Create.
type DocumentsEndpointCreateBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// DocumentsEndpointUpdateBody contains the request body for Update.
type DocumentsEndpointUpdateBody struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
}

// DocumentsEndpointCloneBody contains the request body for Clone.
type DocumentsEndpointCloneBody struct {
	TargetDocumentId *string `json:"target_document_id,omitempty"`
	Title            *string `json:"title,omitempty"`
}

// List returns a paginated list of documents owned by the instance.
func (e *DocumentsEndpoint) List(instanceId string, params *DocumentsEndpointListParams) (*documents.DocumentsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "documents"},
		Query: query,
	}
	var result documents.DocumentsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new document for the instance.
func (e *DocumentsEndpoint) Create(instanceId string, body *DocumentsEndpointCreateBody) (*documents.DocumentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents"},
		Body: body,
	}
	var result documents.DocumentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a document by its ID.
func (e *DocumentsEndpoint) Get(instanceId string, documentId string) (*documents.DocumentsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId},
	}
	var result documents.DocumentsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific document.
func (e *DocumentsEndpoint) Update(instanceId string, documentId string, body *DocumentsEndpointUpdateBody) (*documents.DocumentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId},
		Body: body,
	}
	var result documents.DocumentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a specific document.
func (e *DocumentsEndpoint) Delete(instanceId string, documentId string) (*documents.DocumentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId},
	}
	var result documents.DocumentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Clone clones a specific document.
func (e *DocumentsEndpoint) Clone(instanceId string, documentId string, body *DocumentsEndpointCloneBody) (*documents.DocumentsCloneOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "clone"},
		Body: body,
	}
	var result documents.DocumentsCloneOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
