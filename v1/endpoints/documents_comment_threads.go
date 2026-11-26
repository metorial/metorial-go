package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/commentthreads"
)

// DocumentsCommentThreadsEndpoint provides access to a document comment thread is a discussion anchored to a text selection in a document.
type DocumentsCommentThreadsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsCommentThreadsEndpoint creates a new DocumentsCommentThreadsEndpoint.
func NewDocumentsCommentThreadsEndpoint(client *endpoint.Client) *DocumentsCommentThreadsEndpoint {
	return &DocumentsCommentThreadsEndpoint{client: client}
}

// DocumentsCommentThreadsEndpointListParams contains optional query parameters for List.
type DocumentsCommentThreadsEndpointListParams struct {
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
	// Id - Only return threads with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return threads with one of these statuses.
	Status *any `json:"status,omitempty"`
	// AnchorStatus - Filter by whether the commented text still exists
	AnchorStatus *any `json:"anchor_status,omitempty"`
	// CreatedAt - Only return results whose thread creation time is in this range. Set `gt`, `lt`, or both.
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
	// Status - New status for the thread.
	Status *string `json:"status,omitempty"`
}

// List returns a paginated list of comment threads in a document.
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

// Get retrieves a document comment thread with its comments.
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
