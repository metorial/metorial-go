package callbackinstances

import (
	"encoding/json"
	"time"
)

// CallbackInstancesListOutputItems represents the callback instances list output items type.
type CallbackInstancesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique callback instance identifier
	Id string `json:"id"`
	// Status - Callback instance lifecycle status. Archived once the callback or the integration instance provider goes away.
	Status string `json:"status"`
	// CallbackId - Callback this instance belongs to
	CallbackId string `json:"callback_id"`
	// IntegrationInstanceId - Integration instance this callback instance listens for
	IntegrationInstanceId string `json:"integration_instance_id"`
	// IntegrationInstanceProviderId - Integration instance provider this callback instance is registered on
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// CreatedAt - Timestamp when the callback instance was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the callback instance was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbackInstancesListOutputPagination represents the callback instances list output pagination type.
type CallbackInstancesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// CallbackInstancesListOutput represents the callback instances list output type.
type CallbackInstancesListOutput struct {
	Items      []CallbackInstancesListOutputItems    `json:"items"`
	Pagination CallbackInstancesListOutputPagination `json:"pagination"`
}

// MapCallbackInstancesListOutputFromJSON deserializes JSON data into a CallbackInstancesListOutput.
func MapCallbackInstancesListOutputFromJSON(data []byte) (*CallbackInstancesListOutput, error) {
	var v CallbackInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackInstancesListOutputToJSON serializes a CallbackInstancesListOutput to JSON.
func MapCallbackInstancesListOutputToJSON(v *CallbackInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CallbackInstancesListQueryCreatedAt - Filter callback instance creation time by date range
type CallbackInstancesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for callback instance creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for callback instance creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackInstancesListQueryUpdatedAt - Filter callback instance last update time by date range
type CallbackInstancesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for callback instance last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for callback instance last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackInstancesListQuery represents the callback instances list query type.
type CallbackInstancesListQuery struct {
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
	CreatedAt *CallbackInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter callback instance last update time by date range
	UpdatedAt *CallbackInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCallbackInstancesListQueryFromJSON deserializes JSON data into a CallbackInstancesListQuery.
func MapCallbackInstancesListQueryFromJSON(data []byte) (*CallbackInstancesListQuery, error) {
	var v CallbackInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCallbackInstancesListQueryToJSON serializes a CallbackInstancesListQuery to JSON.
func MapCallbackInstancesListQueryToJSON(v *CallbackInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
