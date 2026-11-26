package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/callbackevents"
)

// CallbackEventsEndpoint provides access to a callback event is recorded every time a provider trigger behind one of your callbacks fires.
type CallbackEventsEndpoint struct {
	client *endpoint.Client
}

// NewCallbackEventsEndpoint creates a new CallbackEventsEndpoint.
func NewCallbackEventsEndpoint(client *endpoint.Client) *CallbackEventsEndpoint {
	return &CallbackEventsEndpoint{client: client}
}

// CallbackEventsEndpointListParams contains optional query parameters for List.
type CallbackEventsEndpointListParams struct {
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
	// CallbackId - Only return events for one of these callbacks.
	CallbackId *any `json:"callback_id,omitempty"`
	// CallbackInstanceId - Only return events for one of these callback instances.
	CallbackInstanceId *any `json:"callback_instance_id,omitempty"`
	// IntegrationId - Only return events for callbacks belonging to one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationProviderId - Only return events for callbacks belonging to one of these integration providers.
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderId - Only return events for callbacks belonging to one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderTriggerKey - Only return events produced by one of these provider triggers.
	ProviderTriggerKey *any `json:"provider_trigger_key,omitempty"`
	// Status - Only return events with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Source - Only return events from one of these sources.
	Source *any `json:"source,omitempty"`
	// OccurredAt - Only return results whose provider event time is in this range. Set `gt`, `lt`, or both.
	OccurredAt *map[string]any `json:"occurred_at,omitempty"`
	// CreatedAt - Only return results whose callback event creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of callback events.
func (e *CallbackEventsEndpoint) List(params *CallbackEventsEndpointListParams) (*callbackevents.CallbackEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"callback-events"},
		Query: query,
	}
	var result callbackevents.CallbackEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a callback event, with the payload the provider produced and the inbound webhook behind it when there is one. The inbound request is `null` once it is past its retention period.
func (e *CallbackEventsEndpoint) Get(callbackEventId string) (*callbackevents.CallbackEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"callback-events", callbackEventId},
	}
	var result callbackevents.CallbackEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
