package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/edittoken"
)

// DocumentsEditTokenEndpoint provides access to create and manage instance documents backed by Cargo.
type DocumentsEditTokenEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsEditTokenEndpoint creates a new DocumentsEditTokenEndpoint.
func NewDocumentsEditTokenEndpoint(client *endpoint.Client) *DocumentsEditTokenEndpoint {
	return &DocumentsEditTokenEndpoint{client: client}
}

// Get returns a short-lived read or write token for establishing a live document session.
func (e *DocumentsEditTokenEndpoint) Get(instanceId string, documentId string) (*edittoken.DocumentsEditTokenGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "edit-token"},
	}
	var result edittoken.DocumentsEditTokenGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
