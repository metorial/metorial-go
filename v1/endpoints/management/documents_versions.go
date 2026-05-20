package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/versions"
)

// DocumentsVersionsEndpoint provides access to inspect document version history for an instance document.
type DocumentsVersionsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsVersionsEndpoint creates a new DocumentsVersionsEndpoint.
func NewDocumentsVersionsEndpoint(client *endpoint.Client) *DocumentsVersionsEndpoint {
	return &DocumentsVersionsEndpoint{client: client}
}

// DocumentsVersionsEndpointListParams contains optional query parameters for List.
type DocumentsVersionsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by document version ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// LastEditedAt - Filter Filter by last edit time by date range
	LastEditedAt *map[string]any `json:"last_edited_at,omitempty"`
}

// List returns a paginated list of versions for a specific document.
func (e *DocumentsVersionsEndpoint) List(instanceId string, documentId string, params *DocumentsVersionsEndpointListParams) (*versions.DocumentsVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "documents", documentId, "versions"},
		Query: query,
	}
	var result versions.DocumentsVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific document version by its ID.
func (e *DocumentsVersionsEndpoint) Get(instanceId string, documentId string, documentVersionId string) (*versions.DocumentsVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "documents", documentId, "versions", documentVersionId},
	}
	var result versions.DocumentsVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
