package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/items"
)

// StoresItemsEndpoint provides access to a store is a shared collection of files and documents. Participants get read or write access to its contents.
type StoresItemsEndpoint struct {
	client *endpoint.Client
}

// NewStoresItemsEndpoint creates a new StoresItemsEndpoint.
func NewStoresItemsEndpoint(client *endpoint.Client) *StoresItemsEndpoint {
	return &StoresItemsEndpoint{client: client}
}

// StoresItemsEndpointModifyBody contains the request body for Modify.
type StoresItemsEndpointModifyBody struct {
	// Operations - Operations to apply to the store.
	Operations []map[string]any `json:"operations"`
}

// StoresItemsEndpointListParams contains optional query parameters for List.
type StoresItemsEndpointListParams struct {
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
	// Id - Only return items with one of these IDs.
	Id *any `json:"id,omitempty"`
	// FileId - Only return items holding one of these files.
	FileId *any `json:"file_id,omitempty"`
	// DocumentId - Only return items holding one of these documents.
	DocumentId *any `json:"document_id,omitempty"`
	// Type - Filter by store item type. Repeat `type` to include multiple values. Defaults to `file` and `document`.
	Type *any `json:"type,omitempty"`
	// CreatedAt - Only return results whose item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose item last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// Modify applies a list of add, modify, and remove operations to a store.
func (e *StoresItemsEndpoint) Modify(instanceId string, storeId string, body *StoresItemsEndpointModifyBody) (*items.StoresItemsModifyOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId, "items"},
		Body: body,
	}
	var result items.StoresItemsModifyOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a paginated list of items in a store.
func (e *StoresItemsEndpoint) List(instanceId string, storeId string, params *StoresItemsEndpointListParams) (*items.StoresItemsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "stores", storeId, "items"},
		Query: query,
	}
	var result items.StoresItemsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves an item in a store.
func (e *StoresItemsEndpoint) Get(instanceId string, storeId string, itemId string) (*items.StoresItemsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId, "items", itemId},
	}
	var result items.StoresItemsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
