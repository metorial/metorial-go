package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/permissions"
)

// DocumentsPermissionsEndpoint provides access to a document is rich text with comments, versions, and participants.
type DocumentsPermissionsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsPermissionsEndpoint creates a new DocumentsPermissionsEndpoint.
func NewDocumentsPermissionsEndpoint(client *endpoint.Client) *DocumentsPermissionsEndpoint {
	return &DocumentsPermissionsEndpoint{client: client}
}

// Get returns what the caller can do with a document.
func (e *DocumentsPermissionsEndpoint) Get(instanceId string, documentId string) (*permissions.DocumentsPermissionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "permissions"},
	}
	var result permissions.DocumentsPermissionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
