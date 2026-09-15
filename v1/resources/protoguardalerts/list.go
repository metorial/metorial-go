package protoguardalerts

import (
	"encoding/json"
	"time"
)

// ProtoGuardAlertsListOutputItemsFilters represents the proto guard alerts list output items filters type.
type ProtoGuardAlertsListOutputItemsFilters struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	FilterId    string    `json:"filter_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IssueType   string    `json:"issue_type"`
	Severity    string    `json:"severity"`
	Confidence  *float64  `json:"confidence,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProtoGuardAlertsListOutputItems represents the proto guard alerts list output items type.
type ProtoGuardAlertsListOutputItems struct {
	Object              string                                   `json:"object"`
	Id                  string                                   `json:"id"`
	RunId               string                                   `json:"run_id"`
	SessionId           *string                                  `json:"session_id,omitempty"`
	SessionMessageId    *string                                  `json:"session_message_id,omitempty"`
	SessionConnectionId *string                                  `json:"session_connection_id,omitempty"`
	ProviderRunId       *string                                  `json:"provider_run_id,omitempty"`
	Filters             []ProtoGuardAlertsListOutputItemsFilters `json:"filters"`
	CreatedAt           time.Time                                `json:"created_at"`
}

// ProtoGuardAlertsListOutputPagination represents the proto guard alerts list output pagination type.
type ProtoGuardAlertsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProtoGuardAlertsListOutput represents the proto guard alerts list output type.
type ProtoGuardAlertsListOutput struct {
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

// ProtoGuardAlertsListQueryCreatedAt - Filter ProtoGuard alert creation time by date range
type ProtoGuardAlertsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for ProtoGuard alert creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for ProtoGuard alert creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// ProtoGuardAlertsListQuery represents the proto guard alerts list query type.
type ProtoGuardAlertsListQuery struct {
	Limit               *float64 `json:"limit,omitempty"`
	After               *string  `json:"after,omitempty"`
	Before              *string  `json:"before,omitempty"`
	Cursor              *string  `json:"cursor,omitempty"`
	Order               *string  `json:"order,omitempty"`
	Id                  *any     `json:"id,omitempty"`
	RunId               *any     `json:"run_id,omitempty"`
	FilterId            *any     `json:"filter_id,omitempty"`
	SessionId           *any     `json:"session_id,omitempty"`
	SessionMessageId    *any     `json:"session_message_id,omitempty"`
	SessionConnectionId *any     `json:"session_connection_id,omitempty"`
	ProviderRunId       *any     `json:"provider_run_id,omitempty"`
	// CreatedAt - Filter ProtoGuard alert creation time by date range
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
