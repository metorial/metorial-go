package profiles

import (
	"encoding/json"
)

// ConsumersProfilesListOutputPagination represents the consumers profiles list output pagination type.
type ConsumersProfilesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConsumersProfilesListOutput represents the consumers profiles list output type.
type ConsumersProfilesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []map[string]any                      `json:"items"`
	Pagination ConsumersProfilesListOutputPagination `json:"pagination"`
}

// MapConsumersProfilesListOutputFromJSON deserializes JSON data into a ConsumersProfilesListOutput.
func MapConsumersProfilesListOutputFromJSON(data []byte) (*ConsumersProfilesListOutput, error) {
	var v ConsumersProfilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesListOutputToJSON serializes a ConsumersProfilesListOutput to JSON.
func MapConsumersProfilesListOutputToJSON(v *ConsumersProfilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersProfilesListQuery represents the consumers profiles list query type.
type ConsumersProfilesListQuery struct {
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
}

// MapConsumersProfilesListQueryFromJSON deserializes JSON data into a ConsumersProfilesListQuery.
func MapConsumersProfilesListQueryFromJSON(data []byte) (*ConsumersProfilesListQuery, error) {
	var v ConsumersProfilesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesListQueryToJSON serializes a ConsumersProfilesListQuery to JSON.
func MapConsumersProfilesListQueryToJSON(v *ConsumersProfilesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
