package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/callbackevents"
)

// CallbackEventsEndpoint provides access to a callback event is recorded every time a provider trigger behind one of your callbacks fires. Listing returns the events themselves; fetch a single event to enrich it with the payload the provider produced, and the inbound webhook behind it, if any.
type CallbackEventsEndpoint struct {
	client *endpoint.Client
}

// NewCallbackEventsEndpoint creates a new CallbackEventsEndpoint.
func NewCallbackEventsEndpoint(client *endpoint.Client) *CallbackEventsEndpoint {
	return &CallbackEventsEndpoint{client: client}
}

// CallbackEventsEndpointListParams contains optional query parameters for List.
type CallbackEventsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// CallbackId - Filter by callback ID(s)
	CallbackId *any `json:"callback_id,omitempty"`
	// CallbackInstanceId - Filter by callback instance ID(s)
	CallbackInstanceId *any `json:"callback_instance_id,omitempty"`
	// IntegrationId - Filter by the integration the callback belongs to
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationProviderId - Filter by the integration provider the callback belongs to
	IntegrationProviderId *any `json:"integration_provider_id,omitempty"`
	// ProviderId - Filter by the provider the callback belongs to
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderTriggerKey - Filter by the provider trigger key that produced the event
	ProviderTriggerKey *any `json:"provider_trigger_key,omitempty"`
	// Status - Filter by callback event processing status
	Status *any `json:"status,omitempty"`
	// Source - Filter by callback event source
	Source *any `json:"source,omitempty"`
	// OccurredAt - Filter when the underlying provider event occurred by date range
	OccurredAt *map[string]any `json:"occurred_at,omitempty"`
	// CreatedAt - Filter callback event creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// List returns a paginated list of callback events.
func (e *CallbackEventsEndpoint) List(instanceId string, params *CallbackEventsEndpointListParams) (*callbackevents.CallbackEventsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "callback-events"},
		Query: query,
	}
	var result callbackevents.CallbackEventsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific callback event by ID, enriched with the payload the provider produced for it and, if it came from a webhook, the inbound request behind it.
func (e *CallbackEventsEndpoint) Get(instanceId string, callbackEventId string) (*callbackevents.CallbackEventsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "callback-events", callbackEventId},
	}
	var result callbackevents.CallbackEventsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
