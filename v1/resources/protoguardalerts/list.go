package protoguardalerts

import (
	"encoding/json"
	"time"
)

// ProtoGuardAlertsListOutputItemsFilters - One filter match folded into the alert.
type ProtoGuardAlertsListOutputItemsFilters struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// FilterId - Filter that matched.
	FilterId string `json:"filter_id"`
	// Key - Stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IssueType - What kind of prompt-injection pattern was found.
	IssueType string `json:"issue_type"`
	// Severity - How severe the finding is.
	Severity string `json:"severity"`
	// Confidence - Confidence of the match. `null` when not scored.
	Confidence *float64 `json:"confidence,omitempty"`
	// RuleKey - The rule within the filter that matched.
	RuleKey *string `json:"rule_key,omitempty"`
	// Direction - Whether the match was in the tool input, the tool output, or across the session.
	Direction *string `json:"direction,omitempty"`
	// ContentHash - A keyed hash of the matched text. Pass it to a suppression to silence this exact content.
	ContentHash *string `json:"content_hash,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ProtoGuardAlertsListOutputItems represents the proto guard alerts list output items type.
type ProtoGuardAlertsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// RunId - Run that produced this alert.
	RunId string `json:"run_id"`
	// SessionId - Session the finding came from. `null` when the finding has no session.
	SessionId *string `json:"session_id,omitempty"`
	// SessionMessageId - Message the finding came from. `null` when it spans the session.
	SessionMessageId *string `json:"session_message_id,omitempty"`
	// SessionConnectionId - Connection the finding came from. `null` when there is none.
	SessionConnectionId *string `json:"session_connection_id,omitempty"`
	// ProviderRunId - Provider run the finding came from. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// IssueId - Provider issue this alert is aggregated into.
	IssueId *string `json:"issue_id,omitempty"`
	// Filters - Filter matches folded into this alert.
	Filters []ProtoGuardAlertsListOutputItemsFilters `json:"filters"`
	// OccurrenceCount - How many times the same finding recurred and was folded into this alert.
	OccurrenceCount float64 `json:"occurrence_count"`
	// FirstSeenAt - When the finding was first seen.
	FirstSeenAt time.Time `json:"first_seen_at"`
	// LastSeenAt - When the finding was last seen.
	LastSeenAt time.Time `json:"last_seen_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ProtoGuardAlertsListOutputPagination represents the proto guard alerts list output pagination type.
type ProtoGuardAlertsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProtoGuardAlertsListOutput represents the proto guard alerts list output type.
type ProtoGuardAlertsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                               `json:"object"`
	Items      []ProtoGuardAlertsListOutputItems    `json:"items"`
	Pagination ProtoGuardAlertsListOutputPagination `json:"pagination"`
}

// MapProtoGuardAlertsListOutputFromJSON deserializes JSON data into a ProtoGuardAlertsListOutput.
func MapProtoGuardAlertsListOutputFromJSON(data []byte) (*ProtoGuardAlertsListOutput, error) {
	var v ProtoGuardAlertsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardAlertsListOutputToJSON serializes a ProtoGuardAlertsListOutput to JSON.
func MapProtoGuardAlertsListOutputToJSON(v *ProtoGuardAlertsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardAlertsListQueryCreatedAt - Only return results whose ProtoGuard alert creation time is in this range. Set `gt`, `lt`, or both.
type ProtoGuardAlertsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProtoGuardAlertsListQuery represents the proto guard alerts list query type.
type ProtoGuardAlertsListQuery struct {
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
	// Id - Only return ProtoGuard alerts with one of these IDs.
	Id *any `json:"id,omitempty"`
	// RunId - Only return ProtoGuard alerts from one of these runs.
	RunId *any `json:"run_id,omitempty"`
	// FilterId - Only return ProtoGuard alerts matching one of these filters.
	FilterId *any `json:"filter_id,omitempty"`
	// SessionId - Only return ProtoGuard alerts linked to one of these sessions.
	SessionId *any `json:"session_id,omitempty"`
	// SessionMessageId - Only return ProtoGuard alerts linked to one of these session messages.
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionConnectionId - Only return ProtoGuard alerts linked to one of these session connections.
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Only return ProtoGuard alerts linked to one of these provider runs.
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// IssueId - Only return ProtoGuard alerts folded into one of these issues.
	IssueId *any `json:"issue_id,omitempty"`
	// CreatedAt - Only return results whose ProtoGuard alert creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProtoGuardAlertsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapProtoGuardAlertsListQueryFromJSON deserializes JSON data into a ProtoGuardAlertsListQuery.
func MapProtoGuardAlertsListQueryFromJSON(data []byte) (*ProtoGuardAlertsListQuery, error) {
	var v ProtoGuardAlertsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardAlertsListQueryToJSON serializes a ProtoGuardAlertsListQuery to JSON.
func MapProtoGuardAlertsListQueryToJSON(v *ProtoGuardAlertsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
