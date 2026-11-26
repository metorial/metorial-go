package magicmcpgroups

import (
	"encoding/json"
	"time"
)

// MagicMcpGroupsListOutputItems represents the magic mcp groups list output items type.
type MagicMcpGroupsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `deleted` once you delete the group.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata map[string]any `json:"metadata"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MagicMcpGroupsListOutputPagination represents the magic mcp groups list output pagination type.
type MagicMcpGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MagicMcpGroupsListOutput represents the magic mcp groups list output type.
type MagicMcpGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []MagicMcpGroupsListOutputItems    `json:"items"`
	Pagination MagicMcpGroupsListOutputPagination `json:"pagination"`
}

// MapMagicMcpGroupsListOutputFromJSON deserializes JSON data into a MagicMcpGroupsListOutput.
func MapMagicMcpGroupsListOutputFromJSON(data []byte) (*MagicMcpGroupsListOutput, error) {
	var v MagicMcpGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsListOutputToJSON serializes a MagicMcpGroupsListOutput to JSON.
func MapMagicMcpGroupsListOutputToJSON(v *MagicMcpGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MagicMcpGroupsListQuery represents the magic mcp groups list query type.
type MagicMcpGroupsListQuery struct {
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
	// Id - Only return groups with these IDs.
	Id *any `json:"id,omitempty"`
}

// MapMagicMcpGroupsListQueryFromJSON deserializes JSON data into a MagicMcpGroupsListQuery.
func MapMagicMcpGroupsListQueryFromJSON(data []byte) (*MagicMcpGroupsListQuery, error) {
	var v MagicMcpGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMagicMcpGroupsListQueryToJSON serializes a MagicMcpGroupsListQuery to JSON.
func MapMagicMcpGroupsListQueryToJSON(v *MagicMcpGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
