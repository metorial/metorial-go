package callbackevents

import (
	"encoding/json"
	"time"
)

// CallbackEventsGetOutputDetailsError - The processing failure, if this event failed
type CallbackEventsGetOutputDetailsError struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Code - Machine-readable reason this event failed to process
	Code string `json:"code"`
	// Message - Human-readable reason this event failed to process
	Message string `json:"message"`
}

// CallbackEventsGetOutputDetailsWebhookRequestBody - Raw body of the inbound request
type CallbackEventsGetOutputDetailsWebhookRequestBody struct {
	// Encoding - Encoding used for the request body
	Encoding string `json:"encoding"`
	// Content - Base64-encoded request body
	Content string `json:"content"`
}

// CallbackEventsGetOutputDetailsWebhookRequest - The inbound HTTP request as it was received. Null once the backend has aged the raw request out of hot storage.
type CallbackEventsGetOutputDetailsWebhookRequest struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Method - HTTP method of the inbound request
	Method string `json:"method"`
	// Url - URL the inbound request was sent to
	Url string `json:"url"`
	// Headers - Headers of the inbound request
	Headers map[string]string `json:"headers"`
	// Body - Raw body of the inbound request
	Body *CallbackEventsGetOutputDetailsWebhookRequestBody `json:"body,omitempty"`
}

// CallbackEventsGetOutputDetailsWebhook - The inbound webhook delivery behind this event, if there was one
type CallbackEventsGetOutputDetailsWebhook struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Status - Processing status of the inbound webhook, as reported by the provider backend
	Status string `json:"status"`
	// Request - The inbound HTTP request as it was received. Null once the backend has aged the raw request out of hot storage.
	Request *CallbackEventsGetOutputDetailsWebhookRequest `json:"request,omitempty"`
	// ReceivedAt - Timestamp when the webhook was received
	ReceivedAt time.Time `json:"received_at"`
}

// CallbackEventsGetOutputDetails - Enrichment fetched from the provider backend. Only present when fetching a single callback event.
type CallbackEventsGetOutputDetails struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Status - Processing status of the trigger invocation, as reported by the provider backend
	Status string `json:"status"`
	// Payload - Mapped payload the provider trigger produced for this event
	Payload *map[string]any `json:"payload,omitempty"`
	// Error - The processing failure, if this event failed
	Error *CallbackEventsGetOutputDetailsError `json:"error,omitempty"`
	// Webhook - The inbound webhook delivery behind this event, if there was one
	Webhook *CallbackEventsGetOutputDetailsWebhook `json:"webhook,omitempty"`
}

// CallbackEventsGetOutput represents the callback events get output type.
type CallbackEventsGetOutput struct {
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
	Details *CallbackEventsGetOutputDetails `json:"details,omitempty"`
	// OccurredAt - Timestamp when the underlying provider event occurred
	OccurredAt time.Time `json:"occurred_at"`
	// CreatedAt - Timestamp when the callback event was recorded
	CreatedAt time.Time `json:"created_at"`
}

// MapCallbackEventsGetOutputFromJSON deserializes JSON data into a CallbackEventsGetOutput.
func MapCallbackEventsGetOutputFromJSON(data []byte) (*CallbackEventsGetOutput, error) {
	var v CallbackEventsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackEventsGetOutputToJSON serializes a CallbackEventsGetOutput to JSON.
func MapCallbackEventsGetOutputToJSON(v *CallbackEventsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
