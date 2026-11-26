package errors

import (
	"encoding/json"
	"time"
)

// SessionsErrorsListOutputItems represents the sessions errors list output items type.
type SessionsErrorsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Additional error details.
	Data map[string]any `json:"data"`
	// Status - Indicates whether the error is still being processed or has been fully processed and grouped.
	Status string `json:"status"`
	// SessionId - Session the error occurred in.
	SessionId string `json:"session_id"`
	// ProviderRunId - Provider run the error belongs to. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// ConnectionId - Connection the error belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// GroupId - Group this error belongs to. `null` while the error is still processing.
	GroupId *string `json:"group_id,omitempty"`
	// SimilarErrorCount - How many errors are in this group.
	SimilarErrorCount float64 `json:"similar_error_count"`
	// IssueId - Provider issue this error is aggregated into. `null` when there is none.
	IssueId *string `json:"issue_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// SessionsErrorsListOutputPagination represents the sessions errors list output pagination type.
type SessionsErrorsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SessionsErrorsListOutput represents the sessions errors list output type.
type SessionsErrorsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []SessionsErrorsListOutputItems    `json:"items"`
	Pagination SessionsErrorsListOutputPagination `json:"pagination"`
}

// MapSessionsErrorsListOutputFromJSON deserializes JSON data into a SessionsErrorsListOutput.
func MapSessionsErrorsListOutputFromJSON(data []byte) (*SessionsErrorsListOutput, error) {
	var v SessionsErrorsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorsListOutputToJSON serializes a SessionsErrorsListOutput to JSON.
func MapSessionsErrorsListOutputToJSON(v *SessionsErrorsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SessionsErrorsListQueryCreatedAt - Only return results whose session error creation time is in this range. Set `gt`, `lt`, or both.
type SessionsErrorsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsErrorsListQueryUpdatedAt - Only return results whose session error last update time is in this range. Set `gt`, `lt`, or both.
type SessionsErrorsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SessionsErrorsListQuery represents the sessions errors list query type.
type SessionsErrorsListQuery struct {
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
	// Id - Filter by session error ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// SessionErrorGroupId - Filter by error group ID(s)
	SessionErrorGroupId *any `json:"session_error_group_id,omitempty"`
	// ProviderRunId - Filter by provider run ID(s)
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// IssueId - Filter by provider issue ID(s)
	IssueId *any `json:"issue_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// SessionMessageId - Filter by session message ID(s)
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// CreatedAt - Only return results whose session error creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SessionsErrorsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose session error last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SessionsErrorsListQueryUpdatedAt `json:"updated_at,omitempty"`
	// Search - Search by session error ID or error code, or by the name of a related provider, deployment, or auth config
	Search *string `json:"search,omitempty"`
}

// MapSessionsErrorsListQueryFromJSON deserializes JSON data into a SessionsErrorsListQuery.
func MapSessionsErrorsListQueryFromJSON(data []byte) (*SessionsErrorsListQuery, error) {
	var v SessionsErrorsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorsListQueryToJSON serializes a SessionsErrorsListQuery to JSON.
func MapSessionsErrorsListQueryToJSON(v *SessionsErrorsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
