package providerruns

import (
	"encoding/json"
	"time"
)

// ProviderRunsListOutputItems represents the provider runs list output items type.
type ProviderRunsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - `running` while the provider handles messages. `stopped` once it finishes.
	Status string `json:"status"`
	// SessionId - Session the run belongs to.
	SessionId string `json:"session_id"`
	// SessionProviderId - Session provider that handled the run.
	SessionProviderId string `json:"session_provider_id"`
	// ProviderId - Provider that ran.
	ProviderId string `json:"provider_id"`
	// ConnectionId - Connection the run used.
	ConnectionId string `json:"connection_id"`
	// CompletedAt - When the run finished. `null` while it is still running.
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderRunsListOutputPagination represents the provider runs list output pagination type.
type ProviderRunsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderRunsListOutput represents the provider runs list output type.
type ProviderRunsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
	Items      []ProviderRunsListOutputItems    `json:"items"`
	Pagination ProviderRunsListOutputPagination `json:"pagination"`
}

// MapProviderRunsListOutputFromJSON deserializes JSON data into a ProviderRunsListOutput.
func MapProviderRunsListOutputFromJSON(data []byte) (*ProviderRunsListOutput, error) {
	var v ProviderRunsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderRunsListOutputToJSON serializes a ProviderRunsListOutput to JSON.
func MapProviderRunsListOutputToJSON(v *ProviderRunsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderRunsListQueryCreatedAt - Only return results whose provider run creation time is in this range. Set `gt`, `lt`, or both.
type ProviderRunsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderRunsListQueryUpdatedAt - Only return results whose provider run last update time is in this range. Set `gt`, `lt`, or both.
type ProviderRunsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderRunsListQuery represents the provider runs list query type.
type ProviderRunsListQuery struct {
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
	// Status - Filter by run status
	Status *any `json:"status,omitempty"`
	// Id - Filter by provider run ID(s)
	Id *any `json:"id,omitempty"`
	// SessionId - Filter by session ID(s)
	SessionId *any `json:"session_id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// SessionProviderId - Filter by session provider ID(s)
	SessionProviderId *any `json:"session_provider_id,omitempty"`
	// SessionConnectionId - Filter by session connection ID(s)
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderVersionId - Filter by provider version ID(s)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// CreatedAt - Only return results whose provider run creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderRunsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider run last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderRunsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderRunsListQueryFromJSON deserializes JSON data into a ProviderRunsListQuery.
func MapProviderRunsListQueryFromJSON(data []byte) (*ProviderRunsListQuery, error) {
	var v ProviderRunsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderRunsListQueryToJSON serializes a ProviderRunsListQuery to JSON.
func MapProviderRunsListQueryToJSON(v *ProviderRunsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
