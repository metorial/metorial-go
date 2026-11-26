package protoguardsuppressions

import (
	"encoding/json"
	"time"
)

// ProtoGuardSuppressionsListOutputItemsFilter - Filter the suppression applies to.
type ProtoGuardSuppressionsListOutputItemsFilter struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
}

// ProtoGuardSuppressionsListOutputItems represents the proto guard suppressions list output items type.
type ProtoGuardSuppressionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the suppression currently hides matching findings.
	Status string `json:"status"`
	// Reason - Why the findings are considered benign.
	Reason *string `json:"reason,omitempty"`
	// Filter - Filter the suppression applies to.
	Filter ProtoGuardSuppressionsListOutputItemsFilter `json:"filter"`
	// ProviderId - The provider whose findings are suppressed, or null to suppress across all providers.
	ProviderId *string `json:"provider_id,omitempty"`
	// ToolKey - The tool whose findings are suppressed, or null to suppress across all tools.
	ToolKey *string `json:"tool_key,omitempty"`
	// ContentHash - The keyed content hash of a specific matched text, taken from a ProtoGuard finding, or null to suppress any content.
	ContentHash *string `json:"content_hash,omitempty"`
	// ExpiresAt - When the suppression stops applying, or null if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProtoGuardSuppressionsListOutputPagination represents the proto guard suppressions list output pagination type.
type ProtoGuardSuppressionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProtoGuardSuppressionsListOutput represents the proto guard suppressions list output type.
type ProtoGuardSuppressionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                     `json:"object"`
	Items      []ProtoGuardSuppressionsListOutputItems    `json:"items"`
	Pagination ProtoGuardSuppressionsListOutputPagination `json:"pagination"`
}

// MapProtoGuardSuppressionsListOutputFromJSON deserializes JSON data into a ProtoGuardSuppressionsListOutput.
func MapProtoGuardSuppressionsListOutputFromJSON(data []byte) (*ProtoGuardSuppressionsListOutput, error) {
	var v ProtoGuardSuppressionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardSuppressionsListOutputToJSON serializes a ProtoGuardSuppressionsListOutput to JSON.
func MapProtoGuardSuppressionsListOutputToJSON(v *ProtoGuardSuppressionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardSuppressionsListQuery represents the proto guard suppressions list query type.
type ProtoGuardSuppressionsListQuery struct {
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
	// Status - Only return suppressions with these statuses. Defaults to active.
	Status *any `json:"status,omitempty"`
	// FilterId - Only return suppressions of these filters, by id or key.
	FilterId *any `json:"filter_id,omitempty"`
	// ProviderId - Only return suppressions scoped to these providers.
	ProviderId *any `json:"provider_id,omitempty"`
}

// MapProtoGuardSuppressionsListQueryFromJSON deserializes JSON data into a ProtoGuardSuppressionsListQuery.
func MapProtoGuardSuppressionsListQueryFromJSON(data []byte) (*ProtoGuardSuppressionsListQuery, error) {
	var v ProtoGuardSuppressionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardSuppressionsListQueryToJSON serializes a ProtoGuardSuppressionsListQuery to JSON.
func MapProtoGuardSuppressionsListQueryToJSON(v *ProtoGuardSuppressionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
