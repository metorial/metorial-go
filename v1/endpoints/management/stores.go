package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores"
)

// StoresEndpoint provides access to create and manage instance stores backed by Cargo.
type StoresEndpoint struct {
	client *endpoint.Client
}

// NewStoresEndpoint creates a new StoresEndpoint.
func NewStoresEndpoint(client *endpoint.Client) *StoresEndpoint {
	return &StoresEndpoint{client: client}
}

// StoresEndpointListParams contains optional query parameters for List.
type StoresEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by store ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// StoresEndpointCreateBody contains the request body for Create.
type StoresEndpointCreateBody struct {
	Name       string  `json:"name"`
	Access     *string `json:"access,omitempty"`
	TemplateId *string `json:"template_id,omitempty"`
	ParentId   *string `json:"parent_id,omitempty"`
}

// StoresEndpointUpdateBody contains the request body for Update.
type StoresEndpointUpdateBody struct {
	Name   *string `json:"name,omitempty"`
	Access *string `json:"access,omitempty"`
}

// List returns a paginated list of stores owned by the instance.
func (e *StoresEndpoint) List(instanceId string, params *StoresEndpointListParams) (*stores.StoresListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "stores"},
		Query: query,
	}
	var result stores.StoresListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a new store for the instance.
func (e *StoresEndpoint) Create(instanceId string, body *StoresEndpointCreateBody) (*stores.StoresCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores"},
		Body: body,
	}
	var result stores.StoresCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a store by its ID.
func (e *StoresEndpoint) Get(instanceId string, storeId string) (*stores.StoresGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId},
	}
	var result stores.StoresGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a specific store.
func (e *StoresEndpoint) Update(instanceId string, storeId string, body *StoresEndpointUpdateBody) (*stores.StoresUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId},
		Body: body,
	}
	var result stores.StoresUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a specific store.
func (e *StoresEndpoint) Delete(instanceId string, storeId string) (*stores.StoresDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId},
	}
	var result stores.StoresDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
