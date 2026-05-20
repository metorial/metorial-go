package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/permissions"
)

// DocumentsPermissionsEndpoint provides access to create and manage instance documents backed by Cargo.
type DocumentsPermissionsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsPermissionsEndpoint creates a new DocumentsPermissionsEndpoint.
func NewDocumentsPermissionsEndpoint(client *endpoint.Client) *DocumentsPermissionsEndpoint {
	return &DocumentsPermissionsEndpoint{client: client}
}

// Get returns the effective Cargo permissions for the current actor on a specific document.
func (e *DocumentsPermissionsEndpoint) Get(documentId string) (*permissions.DocumentsPermissionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "permissions"},
	}
	var result permissions.DocumentsPermissionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
