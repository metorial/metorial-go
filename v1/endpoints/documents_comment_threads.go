package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/commentthreads"
)

// DocumentsCommentThreadsEndpoint provides access to discuss selected text in a document through comment threads.
type DocumentsCommentThreadsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsCommentThreadsEndpoint creates a new DocumentsCommentThreadsEndpoint.
func NewDocumentsCommentThreadsEndpoint(client *endpoint.Client) *DocumentsCommentThreadsEndpoint {
	return &DocumentsCommentThreadsEndpoint{client: client}
}

// DocumentsCommentThreadsEndpointListParams contains optional query parameters for List.
type DocumentsCommentThreadsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by comment thread ID
	Id *any `json:"id,omitempty"`
	// Status - Filter by thread status
	Status *any `json:"status,omitempty"`
	// AnchorStatus - Filter by whether the commented text still exists
	AnchorStatus *any `json:"anchor_status,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// DocumentsCommentThreadsEndpointCreateBody contains the request body for Create.
type DocumentsCommentThreadsEndpointCreateBody struct {
	// Anchor - The text selection the thread is attached to
	Anchor map[string]any `json:"anchor"`
	// Body - The first comment in the thread
	Body string `json:"body"`
}

// DocumentsCommentThreadsEndpointUpdateBody contains the request body for Update.
type DocumentsCommentThreadsEndpointUpdateBody struct {
	Status *string `json:"status,omitempty"`
}

// List returns a paginated list of comment threads for a specific document.
func (e *DocumentsCommentThreadsEndpoint) List(documentId string, params *DocumentsCommentThreadsEndpointListParams) (*commentthreads.DocumentsCommentThreadsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"documents", documentId, "comment-threads"},
		Query: query,
	}
	var result commentthreads.DocumentsCommentThreadsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific comment thread with its comments.
func (e *DocumentsCommentThreadsEndpoint) Get(documentId string, documentCommentThreadId string) (*commentthreads.DocumentsCommentThreadsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "comment-threads", documentCommentThreadId},
	}
	var result commentthreads.DocumentsCommentThreadsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create starts a comment thread on a text selection. Anyone who can read the document can comment.
func (e *DocumentsCommentThreadsEndpoint) Create(documentId string, body *DocumentsCommentThreadsEndpointCreateBody) (*commentthreads.DocumentsCommentThreadsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "comment-threads"},
		Body: body,
	}
	var result commentthreads.DocumentsCommentThreadsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update resolves or reopens a comment thread. Allowed for the thread author and document editors.
func (e *DocumentsCommentThreadsEndpoint) Update(documentId string, documentCommentThreadId string, body *DocumentsCommentThreadsEndpointUpdateBody) (*commentthreads.DocumentsCommentThreadsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "comment-threads", documentCommentThreadId},
		Body: body,
	}
	var result commentthreads.DocumentsCommentThreadsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a comment thread and removes its highlight. Requires document write access.
func (e *DocumentsCommentThreadsEndpoint) Delete(documentId string, documentCommentThreadId string) (*commentthreads.DocumentsCommentThreadsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "comment-threads", documentCommentThreadId},
	}
	var result commentthreads.DocumentsCommentThreadsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
