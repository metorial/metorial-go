package callbackevents

import (
	"encoding/json"
	"time"
)

// CallbackEventsListOutputItemsDetailsError - The processing failure, if this event failed
type CallbackEventsListOutputItemsDetailsError struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Code - Machine-readable reason this event failed to process
	Code string `json:"code"`
	// Message - Human-readable reason this event failed to process
	Message string `json:"message"`
}

// CallbackEventsListOutputItemsDetailsWebhookRequestBody - Raw body of the inbound request
type CallbackEventsListOutputItemsDetailsWebhookRequestBody struct {
	// Encoding - Encoding used for the request body
	Encoding string `json:"encoding"`
	// Content - Base64-encoded request body
	Content string `json:"content"`
}

// CallbackEventsListOutputItemsDetailsWebhookRequest - The inbound HTTP request as it was received. Null once the backend has aged the raw request out of hot storage.
type CallbackEventsListOutputItemsDetailsWebhookRequest struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Method - HTTP method of the inbound request
	Method string `json:"method"`
	// Url - URL the inbound request was sent to
	Url string `json:"url"`
	// Headers - Headers of the inbound request
	Headers map[string]string `json:"headers"`
	// Body - Raw body of the inbound request
	Body *CallbackEventsListOutputItemsDetailsWebhookRequestBody `json:"body,omitempty"`
}

// CallbackEventsListOutputItemsDetailsWebhook - The inbound webhook delivery behind this event, if there was one
type CallbackEventsListOutputItemsDetailsWebhook struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Status - Processing status of the inbound webhook, as reported by the provider backend
	Status string `json:"status"`
	// Request - The inbound HTTP request as it was received. Null once the backend has aged the raw request out of hot storage.
	Request *CallbackEventsListOutputItemsDetailsWebhookRequest `json:"request,omitempty"`
	// ReceivedAt - Timestamp when the webhook was received
	ReceivedAt time.Time `json:"received_at"`
}

// CallbackEventsListOutputItemsDetails - Enrichment fetched from the provider backend. Only present when fetching a single callback event.
type CallbackEventsListOutputItemsDetails struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Status - Processing status of the trigger invocation, as reported by the provider backend
	Status string `json:"status"`
	// Payload - Mapped payload the provider trigger produced for this event
	Payload *map[string]any `json:"payload,omitempty"`
	// Error - The processing failure, if this event failed
	Error *CallbackEventsListOutputItemsDetailsError `json:"error,omitempty"`
	// Webhook - The inbound webhook delivery behind this event, if there was one
	Webhook *CallbackEventsListOutputItemsDetailsWebhook `json:"webhook,omitempty"`
}

// CallbackEventsListOutputItems represents the callback events list output items type.
type CallbackEventsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback event identifier
	Id string `json:"id"`
	// Status - Processing status of this callback event within Metorial
	Status string `json:"status"`
	// Source - How the provider trigger that produced this event was invoked
	Source string `json:"source"`
	// ProviderTriggerKey - Key of the provider trigger this event was produced by
	ProviderTriggerKey string `json:"provider_trigger_key"`
	// MappedType - Type of the resource the provider mapped this event to, if any
	MappedType *string `json:"mapped_type,omitempty"`
	// MappedId - Identifier of the resource the provider mapped this event to, if any
	MappedId *string `json:"mapped_id,omitempty"`
	// CallbackId - Callback this event was recorded for
	CallbackId string `json:"callback_id"`
	// CallbackInstanceId - Callback instance this event was recorded for
	CallbackInstanceId string `json:"callback_instance_id"`
	// Details - Enrichment fetched from the provider backend. Only present when fetching a single callback event.
	Details *CallbackEventsListOutputItemsDetails `json:"details,omitempty"`
	// OccurredAt - Timestamp when the underlying provider event occurred
	OccurredAt time.Time `json:"occurred_at"`
	// CreatedAt - Timestamp when the callback event was recorded
	CreatedAt time.Time `json:"created_at"`
}

// CallbackEventsListOutputPagination represents the callback events list output pagination type.
type CallbackEventsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// CallbackEventsListOutput represents the callback events list output type.
type CallbackEventsListOutput struct {
	Items      []CallbackEventsListOutputItems    `json:"items"`
	Pagination CallbackEventsListOutputPagination `json:"pagination"`
}

// MapCallbackEventsListOutputFromJSON deserializes JSON data into a CallbackEventsListOutput.
func MapCallbackEventsListOutputFromJSON(data []byte) (*CallbackEventsListOutput, error) {
	var v CallbackEventsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackEventsListOutputToJSON serializes a CallbackEventsListOutput to JSON.
func MapCallbackEventsListOutputToJSON(v *CallbackEventsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbackEventsListQueryOccurredAt - Filter when the underlying provider event occurred by date range
type CallbackEventsListQueryOccurredAt struct {
	// Gt - Only include records after this timestamp for when the underlying provider event occurred
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for when the underlying provider event occurred
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackEventsListQueryCreatedAt - Filter callback event creation time by date range
type CallbackEventsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for callback event creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for callback event creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackEventsListQuery represents the callback events list query type.
type CallbackEventsListQuery struct {
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
	OccurredAt *CallbackEventsListQueryOccurredAt `json:"occurred_at,omitempty"`
	// CreatedAt - Filter callback event creation time by date range
	CreatedAt *CallbackEventsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapCallbackEventsListQueryFromJSON deserializes JSON data into a CallbackEventsListQuery.
func MapCallbackEventsListQueryFromJSON(data []byte) (*CallbackEventsListQuery, error) {
	var v CallbackEventsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackEventsListQueryToJSON serializes a CallbackEventsListQuery to JSON.
func MapCallbackEventsListQueryToJSON(v *CallbackEventsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
