package monitors

import (
	"encoding/json"
	"time"
)

// MonitorsListOutputItems represents the monitors list output items type.
type MonitorsListOutputItems struct {
	Object             string     `json:"object"`
	Id                 string     `json:"id"`
	Name               string     `json:"name"`
	Description        *string    `json:"description,omitempty"`
	Target             string     `json:"target"`
	Status             string     `json:"status"`
	Owner              string     `json:"owner"`
	ProtoGuardFilterId *string    `json:"proto_guard_filter_id,omitempty"`
	ProviderId         *string    `json:"provider_id,omitempty"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	FirstAlertAt       *time.Time `json:"first_alert_at,omitempty"`
	LastAlertAt        *time.Time `json:"last_alert_at,omitempty"`
}

// MonitorsListOutputPagination represents the monitors list output pagination type.
type MonitorsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MonitorsListOutput represents the monitors list output type.
type MonitorsListOutput struct {
	Items      []MonitorsListOutputItems    `json:"items"`
	Pagination MonitorsListOutputPagination `json:"pagination"`
}

// MapMonitorsListOutputFromJSON deserializes JSON data into a MonitorsListOutput.
func MapMonitorsListOutputFromJSON(data []byte) (*MonitorsListOutput, error) {
	var v MonitorsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorsListOutputToJSON serializes a MonitorsListOutput to JSON.
func MapMonitorsListOutputToJSON(v *MonitorsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MonitorsListQueryCreatedAt - Filter monitor creation time by date range
type MonitorsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for monitor creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for monitor creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryUpdatedAt - Filter monitor last update time by date range
type MonitorsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for monitor last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for monitor last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryFirstAlertAt - Filter first monitor alert time by date range
type MonitorsListQueryFirstAlertAt struct {
	// Gt - Only include records after this timestamp for first monitor alert time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for first monitor alert time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryLastAlertAt - Filter last monitor alert time by date range
type MonitorsListQueryLastAlertAt struct {
	// Gt - Only include records after this timestamp for last monitor alert time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for last monitor alert time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQuery represents the monitors list query type.
type MonitorsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by monitor ID(s)
	Id                 *any    `json:"id,omitempty"`
	Target             *any    `json:"target,omitempty"`
	Status             *any    `json:"status,omitempty"`
	ProviderId         *any    `json:"provider_id,omitempty"`
	ProtoGuardFilterId *any    `json:"proto_guard_filter_id,omitempty"`
	Search             *string `json:"search,omitempty"`
	// CreatedAt - Filter monitor creation time by date range
	CreatedAt *MonitorsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter monitor last update time by date range
	UpdatedAt *MonitorsListQueryUpdatedAt `json:"updated_at,omitempty"`
	// FirstAlertAt - Filter first monitor alert time by date range
	FirstAlertAt *MonitorsListQueryFirstAlertAt `json:"first_alert_at,omitempty"`
	// LastAlertAt - Filter last monitor alert time by date range
	LastAlertAt *MonitorsListQueryLastAlertAt `json:"last_alert_at,omitempty"`
}

// MapMonitorsListQueryFromJSON deserializes JSON data into a MonitorsListQuery.
func MapMonitorsListQueryFromJSON(data []byte) (*MonitorsListQuery, error) {
	var v MonitorsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorsListQueryToJSON serializes a MonitorsListQuery to JSON.
func MapMonitorsListQueryToJSON(v *MonitorsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
