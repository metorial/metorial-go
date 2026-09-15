package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/callbacks"
)

// CallbacksEndpoint provides access to a callback is what receives provider events for an integration provider. Creating one enables callbacks on the integration provider, and Metorial then registers the callback against every matching integration instance. Setting `callbacks.status` on the integration provider itself does the same thing.
type CallbacksEndpoint struct {
	client *endpoint.Client
}

// NewCallbacksEndpoint creates a new CallbacksEndpoint.
func NewCallbacksEndpoint(client *endpoint.Client) *CallbacksEndpoint {
	return &CallbacksEndpoint{client: client}
}

// CallbacksEndpointListParams contains optional query parameters for List.
type CallbacksEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by callback ID(s)
	Id *any `json:"id,omitempty"`
	// IntegrationId - Filter by integration ID(s)
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationProviderId - Filter by integration provider ID(s)
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// Status - Filter by callback lifecycle status
	Status *any `json:"status,omitempty"`
	// CreatedAt - Filter callback creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Filter callback last update time by date range
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// CallbacksEndpointCreateBody contains the request body for Create.
type CallbacksEndpointCreateBody struct {
	// IntegrationId - Integration the integration provider belongs to
	IntegrationId string `json:"integration_id"`
	// IntegrationProviderId - Integration provider to enable callbacks for
	IntegrationProviderId string `json:"integration_provider_id"`
	// Name - Display name for the callback. Defaults to the name of the integration provider.
	Name *string `json:"name,omitempty"`
	// Description - Description for the callback. Defaults to the description of the integration provider.
	Description *string `json:"description,omitempty"`
}

// CallbacksEndpointUpdateBody contains the request body for Update.
type CallbacksEndpointUpdateBody struct {
	// Name - Updated display name
	Name *string `json:"name,omitempty"`
	// Description - Updated description
	Description *string `json:"description,omitempty"`
	// Metadata - Updated custom metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// List returns a paginated list of callbacks.
func (e *CallbacksEndpoint) List(params *CallbacksEndpointListParams) (*callbacks.CallbacksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"callbacks"},
		Query: query,
	}
	var result callbacks.CallbacksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create enables callbacks for an integration provider and returns the callback it created. Only providers whose type reports `triggers.status` as `enabled` support this. Callback instances are then registered for every matching integration instance in the background.
func (e *CallbacksEndpoint) Create(body *CallbacksEndpointCreateBody) (*callbacks.CallbacksCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callbacks"},
		Body: body,
	}
	var result callbacks.CallbacksCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific callback by ID.
func (e *CallbacksEndpoint) Get(callbackId string) (*callbacks.CallbacksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callbacks", callbackId},
	}
	var result callbacks.CallbacksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the name, description or metadata of a callback. Everything else about a callback is derived from its integration provider - set `callbacks.status` to `disabled` there to tear it down.
func (e *CallbacksEndpoint) Update(callbackId string, body *CallbacksEndpointUpdateBody) (*callbacks.CallbacksUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callbacks", callbackId},
		Body: body,
	}
	var result callbacks.CallbacksUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete disables callbacks on the underlying integration provider, tearing down this callback and every callback instance registered for it.
func (e *CallbacksEndpoint) Delete(callbackId string) (*callbacks.CallbacksDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callbacks", callbackId},
	}
	var result callbacks.CallbacksDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
