package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/commentthreads/comments"
)

// DocumentsCommentThreadsCommentsEndpoint provides access to a document comment is a message in a document comment thread.
type DocumentsCommentThreadsCommentsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsCommentThreadsCommentsEndpoint creates a new DocumentsCommentThreadsCommentsEndpoint.
func NewDocumentsCommentThreadsCommentsEndpoint(client *endpoint.Client) *DocumentsCommentThreadsCommentsEndpoint {
	return &DocumentsCommentThreadsCommentsEndpoint{client: client}
}

// DocumentsCommentThreadsCommentsEndpointCreateBody contains the request body for Create.
type DocumentsCommentThreadsCommentsEndpointCreateBody struct {
	// Body - The comment body
	Body string `json:"body"`
}

// DocumentsCommentThreadsCommentsEndpointUpdateBody contains the request body for Update.
type DocumentsCommentThreadsCommentsEndpointUpdateBody struct {
	// Body - The new comment body
	Body string `json:"body"`
}

// Create creates a comment in a document comment thread.
func (e *DocumentsCommentThreadsCommentsEndpoint) Create(instanceId string, documentId string, documentCommentThreadId string, body *DocumentsCommentThreadsCommentsEndpointCreateBody) (*comments.DocumentsCommentThreadsCommentsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "comment-threads", documentCommentThreadId, "comments"},
		Body: body,
	}
	var result comments.DocumentsCommentThreadsCommentsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update edits a comment. Only the comment author can edit it.
func (e *DocumentsCommentThreadsCommentsEndpoint) Update(instanceId string, documentId string, documentCommentThreadId string, documentCommentId string, body *DocumentsCommentThreadsCommentsEndpointUpdateBody) (*comments.DocumentsCommentThreadsCommentsUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "comment-threads", documentCommentThreadId, "comments", documentCommentId},
		Body: body,
	}
	var result comments.DocumentsCommentThreadsCommentsUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a comment. Only the comment author can delete it.
func (e *DocumentsCommentThreadsCommentsEndpoint) Delete(instanceId string, documentId string, documentCommentThreadId string, documentCommentId string) (*comments.DocumentsCommentThreadsCommentsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "comment-threads", documentCommentThreadId, "comments", documentCommentId},
	}
	var result comments.DocumentsCommentThreadsCommentsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
