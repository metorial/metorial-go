package callbacks

import (
	"encoding/json"
	"time"
)

// CallbacksListOutputItemsProvider represents the callbacks list output items provider type.
type CallbacksListOutputItemsProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksListOutputItems represents the callbacks list output items type.
type CallbacksListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback identifier
	Id string `json:"id"`
	// Status - Callback lifecycle status. Archived once callbacks are disabled on the integration provider.
	Status string `json:"status"`
	// Name - Display name for the callback
	Name string `json:"name"`
	// Description - Optional callback description
	Description *string `json:"description,omitempty"`
	// Metadata - Custom key-value pairs for storing additional callback metadata
	Metadata *map[string]any `json:"metadata,omitempty"`
	// IntegrationId - Integration this callback belongs to
	IntegrationId string `json:"integration_id"`
	// IntegrationProviderId - Integration provider this callback was created for
	IntegrationProviderId string                           `json:"integration_provider_id"`
	Provider              CallbacksListOutputItemsProvider `json:"provider"`
	// CreatedAt - Timestamp when the callback was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbacksListOutputPagination represents the callbacks list output pagination type.
type CallbacksListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// CallbacksListOutput represents the callbacks list output type.
type CallbacksListOutput struct {
	Items      []CallbacksListOutputItems    `json:"items"`
	Pagination CallbacksListOutputPagination `json:"pagination"`
}

// MapCallbacksListOutputFromJSON deserializes JSON data into a CallbacksListOutput.
func MapCallbacksListOutputFromJSON(data []byte) (*CallbacksListOutput, error) {
	var v CallbacksListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksListOutputToJSON serializes a CallbacksListOutput to JSON.
func MapCallbacksListOutputToJSON(v *CallbacksListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbacksListQueryCreatedAt - Filter callback creation time by date range
type CallbacksListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for callback creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for callback creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbacksListQueryUpdatedAt - Filter callback last update time by date range
type CallbacksListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for callback last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for callback last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbacksListQuery represents the callbacks list query type.
type CallbacksListQuery struct {
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
	CreatedAt *CallbacksListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter callback last update time by date range
	UpdatedAt *CallbacksListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCallbacksListQueryFromJSON deserializes JSON data into a CallbacksListQuery.
func MapCallbacksListQueryFromJSON(data []byte) (*CallbacksListQuery, error) {
	var v CallbacksListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbacksListQueryToJSON serializes a CallbacksListQuery to JSON.
func MapCallbacksListQueryToJSON(v *CallbacksListQuery) ([]byte, error) {
	return json.Marshal(v)
}
