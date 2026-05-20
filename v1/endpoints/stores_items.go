package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/items"
)

// StoresItemsEndpoint provides access to create and manage instance stores backed by Cargo.
type StoresItemsEndpoint struct {
	client *endpoint.Client
}

// NewStoresItemsEndpoint creates a new StoresItemsEndpoint.
func NewStoresItemsEndpoint(client *endpoint.Client) *StoresItemsEndpoint {
	return &StoresItemsEndpoint{client: client}
}

// StoresItemsEndpointModifyBody contains the request body for Modify.
type StoresItemsEndpointModifyBody struct {
	Operations []map[string]any `json:"operations"`
}

// StoresItemsEndpointListParams contains optional query parameters for List.
type StoresItemsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by store item ID
	Id *any `json:"id,omitempty"`
	// FileId - Filter by file ID
	FileId *any `json:"file_id,omitempty"`
	// DocumentId - Filter by document ID
	DocumentId *any `json:"document_id,omitempty"`
	// Type - Filter by store item type. Repeat `type` to include multiple values. Defaults to `file` and `document`.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// Modify applies bulk item operations to a specific store.
func (e *StoresItemsEndpoint) Modify(storeId string, body *StoresItemsEndpointModifyBody) (*items.StoresItemsModifyOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId, "items"},
		Body: body,
	}
	var result items.StoresItemsModifyOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a paginated list of items for a specific store.
func (e *StoresItemsEndpoint) List(storeId string, params *StoresItemsEndpointListParams) (*items.StoresItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"stores", storeId, "items"},
		Query: query,
	}
	var result items.StoresItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific item within a store.
func (e *StoresItemsEndpoint) Get(storeId string, itemId string) (*items.StoresItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId, "items", itemId},
	}
	var result items.StoresItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
