package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/documents/versions"
)

// DocumentsVersionsEndpoint provides access to a document version is a snapshot of a document's content.
type DocumentsVersionsEndpoint struct {
	client *endpoint.Client
}

// NewDocumentsVersionsEndpoint creates a new DocumentsVersionsEndpoint.
func NewDocumentsVersionsEndpoint(client *endpoint.Client) *DocumentsVersionsEndpoint {
	return &DocumentsVersionsEndpoint{client: client}
}

// DocumentsVersionsEndpointListParams contains optional query parameters for List.
type DocumentsVersionsEndpointListParams struct {
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
	// Id - Only return versions with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose version creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// LastEditedAt - Only return results whose version last edit time is in this range. Set `gt`, `lt`, or both.
	LastEditedAt *map[string]any `json:"last_edited_at,omitempty"`
}

// List returns a paginated list of versions of a document.
func (e *DocumentsVersionsEndpoint) List(documentId string, params *DocumentsVersionsEndpointListParams) (*versions.DocumentsVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"documents", documentId, "versions"},
		Query: query,
	}
	var result versions.DocumentsVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a version of a document.
func (e *DocumentsVersionsEndpoint) Get(documentId string, documentVersionId string) (*versions.DocumentsVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"documents", documentId, "versions", documentVersionId},
	}
	var result versions.DocumentsVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
