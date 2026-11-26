package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores"
)

// StoresEndpoint provides access to a store is a shared collection of files and documents. Participants get read or write access to its contents.
type StoresEndpoint struct {
	client *endpoint.Client
}

// NewStoresEndpoint creates a new StoresEndpoint.
func NewStoresEndpoint(client *endpoint.Client) *StoresEndpoint {
	return &StoresEndpoint{client: client}
}

// StoresEndpointListParams contains optional query parameters for List.
type StoresEndpointListParams struct {
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
	// Id - Only return stores with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose store creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose store last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// StoresEndpointCreateBody contains the request body for Create.
type StoresEndpointCreateBody struct {
	// Name - Display name for the store.
	Name string `json:"name"`
	// Access - Who else can access the store: `private`, `public_read`, or `public_write`.
	Access *string `json:"access,omitempty"`
	// TemplateId - Template to create the store from, by ID. Cannot be combined with `parent_id`.
	TemplateId *string `json:"template_id,omitempty"`
	// ParentId - Store to copy, by ID. Cannot be combined with `template_id`.
	ParentId *string `json:"parent_id,omitempty"`
}

// StoresEndpointUpdateBody contains the request body for Update.
type StoresEndpointUpdateBody struct {
	// Name - Display name for the store.
	Name *string `json:"name,omitempty"`
	// Access - Who else can access the store: `private`, `public_read`, or `public_write`.
	Access *string `json:"access,omitempty"`
}

// List returns a paginated list of stores.
func (e *StoresEndpoint) List(params *StoresEndpointListParams) (*stores.StoresListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"stores"},
		Query: query,
	}
	var result stores.StoresListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a store.
func (e *StoresEndpoint) Create(body *StoresEndpointCreateBody) (*stores.StoresCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores"},
		Body: body,
	}
	var result stores.StoresCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a store.
func (e *StoresEndpoint) Get(storeId string) (*stores.StoresGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId},
	}
	var result stores.StoresGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a store.
func (e *StoresEndpoint) Update(storeId string, body *StoresEndpointUpdateBody) (*stores.StoresUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId},
		Body: body,
	}
	var result stores.StoresUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a store.
func (e *StoresEndpoint) Delete(storeId string) (*stores.StoresDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId},
	}
	var result stores.StoresDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
