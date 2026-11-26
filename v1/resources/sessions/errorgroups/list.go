package errorgroups

import (
	"encoding/json"
	"time"
)

// SessionsErrorGroupsListOutputItems represents the sessions error groups list output items type.
type SessionsErrorGroupsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Error group data from the first occurrence.
	Data map[string]any `json:"data"`
	// ProviderId - Provider the errors belong to. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// OccurrenceCount - How many errors are in this group.
	OccurrenceCount float64 `json:"occurrence_count"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsErrorGroupsListOutputPagination represents the sessions error groups list output pagination type.
type SessionsErrorGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsErrorGroupsListOutput represents the sessions error groups list output type.
type SessionsErrorGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []SessionsErrorGroupsListOutputItems    `json:"items"`
	Pagination SessionsErrorGroupsListOutputPagination `json:"pagination"`
}

// MapSessionsErrorGroupsListOutputFromJSON deserializes JSON data into a SessionsErrorGroupsListOutput.
func MapSessionsErrorGroupsListOutputFromJSON(data []byte) (*SessionsErrorGroupsListOutput, error) {
	var v SessionsErrorGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorGroupsListOutputToJSON serializes a SessionsErrorGroupsListOutput to JSON.
func MapSessionsErrorGroupsListOutputToJSON(v *SessionsErrorGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsErrorGroupsListQuery represents the sessions error groups list query type.
type SessionsErrorGroupsListQuery struct {
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
	// Type - Filter by error type(s)
	Type *any `json:"type,omitempty"`
	// Id - Filter by error group ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Search by error group ID or error code, or by the name of a related provider
	Search *string `json:"search,omitempty"`
}

// MapSessionsErrorGroupsListQueryFromJSON deserializes JSON data into a SessionsErrorGroupsListQuery.
func MapSessionsErrorGroupsListQueryFromJSON(data []byte) (*SessionsErrorGroupsListQuery, error) {
	var v SessionsErrorGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorGroupsListQueryToJSON serializes a SessionsErrorGroupsListQuery to JSON.
func MapSessionsErrorGroupsListQueryToJSON(v *SessionsErrorGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
