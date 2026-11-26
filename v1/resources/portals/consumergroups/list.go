package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsListOutputItems represents the portals consumer groups list output items type.
type PortalsConsumerGroupsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `archived` once you delete the group. Archived groups stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IsDefault - Whether this is a default group on the portal.
	IsDefault bool `json:"is_default"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// PortalsConsumerGroupsListOutputPagination represents the portals consumer groups list output pagination type.
type PortalsConsumerGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsConsumerGroupsListOutput represents the portals consumer groups list output type.
type PortalsConsumerGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []PortalsConsumerGroupsListOutputItems    `json:"items"`
	Pagination PortalsConsumerGroupsListOutputPagination `json:"pagination"`
}

// MapPortalsConsumerGroupsListOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsListOutput.
func MapPortalsConsumerGroupsListOutputFromJSON(data []byte) (*PortalsConsumerGroupsListOutput, error) {
	var v PortalsConsumerGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsListOutputToJSON serializes a PortalsConsumerGroupsListOutput to JSON.
func MapPortalsConsumerGroupsListOutputToJSON(v *PortalsConsumerGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsListQuery represents the portals consumer groups list query type.
type PortalsConsumerGroupsListQuery struct {
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
	// Status - Only return groups with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Search - Search by name or description.
	Search *string `json:"search,omitempty"`
}

// MapPortalsConsumerGroupsListQueryFromJSON deserializes JSON data into a PortalsConsumerGroupsListQuery.
func MapPortalsConsumerGroupsListQueryFromJSON(data []byte) (*PortalsConsumerGroupsListQuery, error) {
	var v PortalsConsumerGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsListQueryToJSON serializes a PortalsConsumerGroupsListQuery to JSON.
func MapPortalsConsumerGroupsListQueryToJSON(v *PortalsConsumerGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
