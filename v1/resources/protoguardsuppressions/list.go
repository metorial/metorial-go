package protoguardsuppressions

import (
	"encoding/json"
	"time"
)

// ProtoGuardSuppressionsListOutputItemsFilter represents the proto guard suppressions list output items filter type.
type ProtoGuardSuppressionsListOutputItemsFilter struct {
	Object string `json:"object"`
	// Id - The unique identifier of the filter.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - The display name of the filter.
	Name string `json:"name"`
}

// ProtoGuardSuppressionsListOutputItems represents the proto guard suppressions list output items type.
type ProtoGuardSuppressionsListOutputItems struct {
	Object string `json:"object"`
	// Id - The unique identifier of the suppression.
	Id string `json:"id"`
	// Status - Whether the suppression currently hides matching findings.
	Status string `json:"status"`
	// Reason - Why the findings are considered benign.
	Reason *string                                     `json:"reason,omitempty"`
	Filter ProtoGuardSuppressionsListOutputItemsFilter `json:"filter"`
	// ProviderId - The provider whose findings are suppressed, or null to suppress across all providers.
	ProviderId *string `json:"provider_id,omitempty"`
	// ToolKey - The tool whose findings are suppressed, or null to suppress across all tools.
	ToolKey *string `json:"tool_key,omitempty"`
	// ContentHash - The keyed content hash of a specific matched text, taken from a ProtoGuard finding, or null to suppress any content.
	ContentHash *string `json:"content_hash,omitempty"`
	// ExpiresAt - When the suppression stops applying, or null if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ArchivedAt - When the suppression was archived.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// CreatedAt - When the suppression was created.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the suppression was last updated.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProtoGuardSuppressionsListOutputPagination represents the proto guard suppressions list output pagination type.
type ProtoGuardSuppressionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProtoGuardSuppressionsListOutput represents the proto guard suppressions list output type.
type ProtoGuardSuppressionsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
