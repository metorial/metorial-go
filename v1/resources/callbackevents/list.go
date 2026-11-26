package callbackevents

import (
	"encoding/json"
	"time"
)

// CallbackEventsListOutputItems represents the callback events list output items type.
type CallbackEventsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Processing status of this callback event within Metorial. `failed` when processing failed; the reason is in `details.error`.
	Status string `json:"status"`
	// Source - How this event reached Metorial. `webhook` when the provider pushed it, `polling` when Metorial fetched it.
	Source string `json:"source"`
	// ProviderTriggerKey - Key of the provider trigger this event was produced by
	ProviderTriggerKey string `json:"provider_trigger_key"`
	// MappedType - Type of the resource the provider mapped this event to, if any
	MappedType *string `json:"mapped_type,omitempty"`
	// MappedId - Identifier of the resource the provider mapped this event to, if any
	MappedId *string `json:"mapped_id,omitempty"`
	// Attributes - Short key/value attributes the provider attached to this event (for example a file id or MIME type), mapping each key to its values. Empty for events without attributes.
	Attributes map[string][]string `json:"attributes"`
	// CallbackId - Callback this event was recorded for
	CallbackId string `json:"callback_id"`
	// CallbackInstanceId - Callback instance this event was recorded for
	CallbackInstanceId string `json:"callback_instance_id"`
	// OccurredAt - Timestamp when the underlying provider event occurred
	OccurredAt time.Time `json:"occurred_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// CallbackEventsListOutputPagination represents the callback events list output pagination type.
type CallbackEventsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CallbackEventsListOutput represents the callback events list output type.
type CallbackEventsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
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

// CallbackEventsListQueryOccurredAt - Only return results whose provider event time is in this range. Set `gt`, `lt`, or both.
type CallbackEventsListQueryOccurredAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackEventsListQueryCreatedAt - Only return results whose callback event creation time is in this range. Set `gt`, `lt`, or both.
type CallbackEventsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackEventsListQuery represents the callback events list query type.
type CallbackEventsListQuery struct {
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
	OccurredAt *CallbackEventsListQueryOccurredAt `json:"occurred_at,omitempty"`
	// CreatedAt - Only return results whose callback event creation time is in this range. Set `gt`, `lt`, or both.
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
