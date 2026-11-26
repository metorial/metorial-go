package stores

import (
	"encoding/json"
	"time"
)

// StoresListOutputItems represents the stores list output items type.
type StoresListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Access - Who else can access the store: `private`, `public_read`, or `public_write`.
	Access string `json:"access"`
	// ItemCount - Number of items in the store.
	ItemCount float64 `json:"item_count"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresListOutputPagination represents the stores list output pagination type.
type StoresListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// StoresListOutput represents the stores list output type.
type StoresListOutput struct {
	// Object - String representing the object’s type.
	Object     string                     `json:"object"`
	Items      []StoresListOutputItems    `json:"items"`
	Pagination StoresListOutputPagination `json:"pagination"`
}

// MapStoresListOutputFromJSON deserializes JSON data into a StoresListOutput.
func MapStoresListOutputFromJSON(data []byte) (*StoresListOutput, error) {
	var v StoresListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresListOutputToJSON serializes a StoresListOutput to JSON.
func MapStoresListOutputToJSON(v *StoresListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresListQueryCreatedAt - Only return results whose store creation time is in this range. Set `gt`, `lt`, or both.
type StoresListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresListQueryUpdatedAt - Only return results whose store last update time is in this range. Set `gt`, `lt`, or both.
type StoresListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresListQuery represents the stores list query type.
type StoresListQuery struct {
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
	// Id - Only return stores with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose store creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *StoresListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose store last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *StoresListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapStoresListQueryFromJSON deserializes JSON data into a StoresListQuery.
func MapStoresListQueryFromJSON(data []byte) (*StoresListQuery, error) {
	var v StoresListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresListQueryToJSON serializes a StoresListQuery to JSON.
func MapStoresListQueryToJSON(v *StoresListQuery) ([]byte, error) {
	return json.Marshal(v)
}
