package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/edittoken"
)

// DocumentsEditTokenEndpoint provides access to a document is rich text with comments, versions, and participants.
type DocumentsEditTokenEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsEditTokenEndpoint creates a new DocumentsEditTokenEndpoint.
func NewDocumentsEditTokenEndpoint(client *endpoint.Client) *DocumentsEditTokenEndpoint {
	return &DocumentsEditTokenEndpoint{client: client}
}

// Get returns a short-lived read or write token for establishing a live document session. Requires write scopes, even when the issued token only grants read access.
func (e *DocumentsEditTokenEndpoint) Get(documentId string) (*edittoken.DocumentsEditTokenGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "edit-token"},
	}
	var result edittoken.DocumentsEditTokenGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
