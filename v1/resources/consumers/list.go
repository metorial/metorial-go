package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersListOutputItems represents the consumers list output items type.
type ConsumersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Email - Email address.
	Email string `json:"email"`
	// ImageUrl - Avatar image URL.
	ImageUrl string `json:"image_url"`
	// UserId - User this consumer is linked to. `null` when unlinked.
	UserId *string `json:"user_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ConsumersListOutputPagination represents the consumers list output pagination type.
type ConsumersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ConsumersListOutput represents the consumers list output type.
type ConsumersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
	Items      []ConsumersListOutputItems    `json:"items"`
	Pagination ConsumersListOutputPagination `json:"pagination"`
}

// MapConsumersListOutputFromJSON deserializes JSON data into a ConsumersListOutput.
func MapConsumersListOutputFromJSON(data []byte) (*ConsumersListOutput, error) {
	var v ConsumersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersListOutputToJSON serializes a ConsumersListOutput to JSON.
func MapConsumersListOutputToJSON(v *ConsumersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersListQuery represents the consumers list query type.
type ConsumersListQuery struct {
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
	// Search - Search by name or email.
	Search *string `json:"search,omitempty"`
	// Email - Only return consumers with one of these emails.
	Email *any `json:"email,omitempty"`
	// Id - Only return consumers with one of these IDs.
	Id *any `json:"id,omitempty"`
}

// MapConsumersListQueryFromJSON deserializes JSON data into a ConsumersListQuery.
func MapConsumersListQueryFromJSON(data []byte) (*ConsumersListQuery, error) {
	var v ConsumersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersListQueryToJSON serializes a ConsumersListQuery to JSON.
func MapConsumersListQueryToJSON(v *ConsumersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
