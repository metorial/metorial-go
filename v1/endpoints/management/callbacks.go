package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
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
	// Id - Only return callbacks with one of these IDs.
	Id *any `json:"id,omitempty"`
	// IntegrationId - Only return callbacks for one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationProviderId - Only return callbacks for one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderId - Only return callbacks for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search callbacks by name, description, integration, or provider.
	Search *string `json:"search,omitempty"`
	// Status - Only return callbacks with one of these statuses. `deleted` callbacks are never listed.
	Status *any `json:"status,omitempty"`
	// CreatedAt - Only return results whose callback creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose callback last update time is in this range. Set `gt`, `lt`, or both.
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
	// TriggerGroupSelections - Resource selections keyed by trigger group key. Each value must match that group’s selection schema. Limited to 64 groups and 64 KiB of JSON.
	TriggerGroupSelections *map[string]map[string]any `json:"trigger_group_selections,omitempty"`
}

// CallbacksEndpointUpdateBody contains the request body for Update.
type CallbacksEndpointUpdateBody struct {
	// Name - Updated display name
	Name *string `json:"name,omitempty"`
	// Description - Updated free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the callback. Replaces existing metadata. `null` clears it.
	Metadata *nullable.Nullable[map[string]any] `json:"metadata,omitempty"`
	// TriggerGroupSelections - Resource selections keyed by trigger group key. Each value must match that group’s selection schema. Limited to 64 groups and 64 KiB of JSON.
	TriggerGroupSelections *map[string]map[string]any `json:"trigger_group_selections,omitempty"`
}

// List returns a paginated list of callbacks.
func (e *CallbacksEndpoint) List(instanceId string, params *CallbacksEndpointListParams) (*callbacks.CallbacksListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "callbacks"},
		Query: query,
	}
	var result callbacks.CallbacksListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create enables callbacks for an integration provider and returns the callback it created. Only providers whose type reports `triggers.status` as `enabled` support this. Callback instances are then registered for every matching integration instance in the background.
func (e *CallbacksEndpoint) Create(instanceId string, body *CallbacksEndpointCreateBody) (*callbacks.CallbacksCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "callbacks"},
		Body: body,
	}
	var result callbacks.CallbacksCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a callback.
func (e *CallbacksEndpoint) Get(instanceId string, callbackId string) (*callbacks.CallbacksGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "callbacks", callbackId},
	}
	var result callbacks.CallbacksGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a callback. Omitted fields keep their current value. Changing the selections re-registers the affected triggers. Only an active callback can be updated.
func (e *CallbacksEndpoint) Update(instanceId string, callbackId string, body *CallbacksEndpointUpdateBody) (*callbacks.CallbacksUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "callbacks", callbackId},
		Body: body,
	}
	var result callbacks.CallbacksUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete archives a callback by disabling callbacks on its integration provider. Callback instances are torn down in the background. Only an active callback can be archived.
func (e *CallbacksEndpoint) Delete(instanceId string, callbackId string) (*callbacks.CallbacksDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "callbacks", callbackId},
	}
	var result callbacks.CallbacksDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
