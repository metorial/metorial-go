package callbackinstances

import (
	"encoding/json"
	"time"
)

// CallbackInstancesListOutputItems represents the callback instances list output items type.
type CallbackInstancesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Callback instance lifecycle status. Archived once the callback or the integration instance provider goes away.
	Status string `json:"status"`
	// CallbackId - Callback this instance belongs to
	CallbackId string `json:"callback_id"`
	// IntegrationInstanceId - Integration instance this callback instance listens for
	IntegrationInstanceId string `json:"integration_instance_id"`
	// IntegrationInstanceProviderId - Integration instance provider this callback instance is registered on
	IntegrationInstanceProviderId string `json:"integration_instance_provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CallbackInstancesListOutputPagination represents the callback instances list output pagination type.
type CallbackInstancesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CallbackInstancesListOutput represents the callback instances list output type.
type CallbackInstancesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
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

// CallbackInstancesListQueryCreatedAt - Only return results whose callback instance creation time is in this range. Set `gt`, `lt`, or both.
type CallbackInstancesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackInstancesListQueryUpdatedAt - Only return results whose callback instance last update time is in this range. Set `gt`, `lt`, or both.
type CallbackInstancesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CallbackInstancesListQuery represents the callback instances list query type.
type CallbackInstancesListQuery struct {
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
	// Id - Only return callback instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CallbackId - Only return instances of one of these callbacks.
	CallbackId *any `json:"callback_id,omitempty"`
	// IntegrationId - Only return instances of callbacks belonging to one of these integrations.
	IntegrationId *any `json:"integration_id,omitempty"`
	// IntegrationInstanceId - Only return instances listening for one of these integration instances.
	IntegrationInstanceId *any `json:"integration_instance_id,omitempty"`
	// IntegrationInstanceProviderId - Only return instances registered on one of these integration instance providers.
	IntegrationInstanceProviderId *any `json:"integration_instance_provider_id,omitempty"`
	// Status - Only return callback instances with one of these statuses. `deleted` instances are never listed.
	Status *any `json:"status,omitempty"`
	// Search - Match a callback instance by its exact ID
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose callback instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CallbackInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose callback instance last update time is in this range. Set `gt`, `lt`, or both.
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
