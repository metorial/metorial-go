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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by callback instance ID(s)
	Id *any `json:"id,omitempty"`
	// CallbackId - Filter by callback ID(s)
	CallbackId *any `json:"callback_id,omitempty"`
	// IntegrationId - Filter by the integration the callback belongs to
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Filter by integration instance ID(s)
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Filter by integration instance provider ID(s)
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// Status - Filter by callback instance lifecycle status
	Status *any `json:"status,omitempty"`
	// CreatedAt - Filter callback instance creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter callback instance last update time by date range
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

// Get retrieves a specific callback instance by ID.
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
