package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/callbackinstances"
)

// CallbackInstancesEndpoint provides access to a callback instance is a callback as it applies to one integration instance provider. Metorial reconciles one for every matching integration instance, so these are read-only.
type CallbackInstancesEndpoint struct {
	client *endpoint.Client
}

// NewCallbackInstancesEndpoint creates a new CallbackInstancesEndpoint.
func NewCallbackInstancesEndpoint(client *endpoint.Client) *CallbackInstancesEndpoint {
	return &CallbackInstancesEndpoint{client: client}
}

// CallbackInstancesEndpointListParams contains optional query parameters for List.
type CallbackInstancesEndpointListParams struct {
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
	// Id - Only return callback instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CallbackId - Only return instances of one of these callbacks.
	CallbackId *any `json:"callback_id,omitempty"`
	// IntegrationId - Only return instances of callbacks belonging to one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return instances listening for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return instances registered on one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// Status - Only return callback instances with one of these statuses. `deleted` instances are never listed.
	Status *any `json:"status,omitempty"`
	// Search - Match a callback instance by its exact ID
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose callback instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose callback instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of callback instances.
func (e *CallbackInstancesEndpoint) List(params *CallbackInstancesEndpointListParams) (*callbackinstances.CallbackInstancesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"callback-instances"},
		Query: query,
	}
	var result callbackinstances.CallbackInstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a callback instance.
func (e *CallbackInstancesEndpoint) Get(callbackInstanceId string) (*callbackinstances.CallbackInstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callback-instances", callbackInstanceId},
	}
	var result callbackinstances.CallbackInstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
