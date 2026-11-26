package triggers

import (
	"encoding/json"
	"time"
)

// ProvidersTriggersListOutputItemsInputSchema represents the providers triggers list output items input schema type.
type ProvidersTriggersListOutputItemsInputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the trigger payload input shape
	Schema map[string]any `json:"schema"`
}

// ProvidersTriggersListOutputItemsOutputSchema represents the providers triggers list output items output schema type.
type ProvidersTriggersListOutputItemsOutputSchema struct {
	Type string `json:"type"`
	// Schema - JSON Schema defining the trigger delivery output shape. Fields marked `x-metorial-event-attribute` are stored with each event and can be used to filter the event list with its `attribute` parameter.
	Schema map[string]any `json:"schema"`
}

// ProvidersTriggersListOutputItemsInvocationAutoRegistration represents the providers triggers list output items invocation auto registration type.
type ProvidersTriggersListOutputItemsInvocationAutoRegistration struct {
	// Status - Whether Metorial registers the webhook for you.
	Status string `json:"status"`
}

// ProvidersTriggersListOutputItemsInvocationAutoUnregistration represents the providers triggers list output items invocation auto unregistration type.
type ProvidersTriggersListOutputItemsInvocationAutoUnregistration struct {
	// Status - Whether Metorial removes the webhook for you.
	Status string `json:"status"`
}

// ProvidersTriggersListOutputItemsInvocation represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type ProvidersTriggersListOutputItemsInvocation struct {
	Type *string `json:"type,omitempty"`
	// IntervalSeconds - Polling interval in seconds for polling-based triggers
	IntervalSeconds    *float64                                                      `json:"interval_seconds,omitempty"`
	AutoRegistration   *ProvidersTriggersListOutputItemsInvocationAutoRegistration   `json:"auto_registration,omitempty"`
	AutoUnregistration *ProvidersTriggersListOutputItemsInvocationAutoUnregistration `json:"auto_unregistration,omitempty"`
}

// ProvidersTriggersListOutputItems represents the providers triggers list output items type.
type ProvidersTriggersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - Trigger key used when subscribing callbacks
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description  *string                                       `json:"description,omitempty"`
	InputSchema  *ProvidersTriggersListOutputItemsInputSchema  `json:"input_schema,omitempty"`
	OutputSchema *ProvidersTriggersListOutputItemsOutputSchema `json:"output_schema,omitempty"`
	Invocation   ProvidersTriggersListOutputItemsInvocation    `json:"invocation"`
	// ProviderId - Provider this trigger belongs to.
	ProviderId string `json:"provider_id"`
	// ProviderSpecificationId - Specification this trigger belongs to.
	ProviderSpecificationId string `json:"provider_specification_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProvidersTriggersListOutputPagination represents the providers triggers list output pagination type.
type ProvidersTriggersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersTriggersListOutput represents the providers triggers list output type.
type ProvidersTriggersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []ProvidersTriggersListOutputItems    `json:"items"`
	Pagination ProvidersTriggersListOutputPagination `json:"pagination"`
}

// MapProvidersTriggersListOutputFromJSON deserializes JSON data into a ProvidersTriggersListOutput.
func MapProvidersTriggersListOutputFromJSON(data []byte) (*ProvidersTriggersListOutput, error) {
	var v ProvidersTriggersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersTriggersListOutputToJSON serializes a ProvidersTriggersListOutput to JSON.
func MapProvidersTriggersListOutputToJSON(v *ProvidersTriggersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersTriggersListQuery represents the providers triggers list query type.
type ProvidersTriggersListQuery struct {
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
	// ProviderVersionId - Provider version to list triggers for
	ProviderVersionId string `json:"provider_version_id"`
}

// MapProvidersTriggersListQueryFromJSON deserializes JSON data into a ProvidersTriggersListQuery.
func MapProvidersTriggersListQueryFromJSON(data []byte) (*ProvidersTriggersListQuery, error) {
	var v ProvidersTriggersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersTriggersListQueryToJSON serializes a ProvidersTriggersListQuery to JSON.
func MapProvidersTriggersListQueryToJSON(v *ProvidersTriggersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
