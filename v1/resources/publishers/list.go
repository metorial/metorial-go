package publishers

import (
	"encoding/json"
	"time"
)

// PublishersListOutputItems represents the publishers list output items type.
type PublishersListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - URL of the publisher logo
	ImageUrl string `json:"image_url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PublishersListOutputPagination represents the publishers list output pagination type.
type PublishersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PublishersListOutput represents the publishers list output type.
type PublishersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                         `json:"object"`
	Items      []PublishersListOutputItems    `json:"items"`
	Pagination PublishersListOutputPagination `json:"pagination"`
}

// MapPublishersListOutputFromJSON deserializes JSON data into a PublishersListOutput.
func MapPublishersListOutputFromJSON(data []byte) (*PublishersListOutput, error) {
	var v PublishersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPublishersListOutputToJSON serializes a PublishersListOutput to JSON.
func MapPublishersListOutputToJSON(v *PublishersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PublishersListQuery represents the publishers list query type.
type PublishersListQuery struct {
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

// MapPublishersListQueryFromJSON deserializes JSON data into a PublishersListQuery.
func MapPublishersListQueryFromJSON(data []byte) (*PublishersListQuery, error) {
	var v PublishersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPublishersListQueryToJSON serializes a PublishersListQuery to JSON.
func MapPublishersListQueryToJSON(v *PublishersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
