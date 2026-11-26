package monitors

import (
	"encoding/json"
	"time"
)

// MonitorsListOutputItems represents the monitors list output items type.
type MonitorsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Target - What the monitor watches.
	Target string `json:"target"`
	// Status - Current status of the monitor. `inactive` monitors raise no alerts.
	Status string `json:"status"`
	// Owner - Who manages the monitor. `system` monitors are managed by Metorial.
	Owner string `json:"owner"`
	// ProtoGuardFilterId - ProtoGuard filter the monitor watches. `null` unless the target is `protoguard_filter`.
	ProtoGuardFilterId *string `json:"proto_guard_filter_id,omitempty"`
	// ProviderId - Provider the monitor watches. `null` unless scoped to a provider.
	ProviderId *string `json:"provider_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// FirstAlertAt - When the monitor first raised an alert. `null` when it never has.
	FirstAlertAt *time.Time `json:"first_alert_at,omitempty"`
	// LastAlertAt - When the monitor last raised an alert. `null` when it never has.
	LastAlertAt *time.Time `json:"last_alert_at,omitempty"`
}

// MonitorsListOutputPagination represents the monitors list output pagination type.
type MonitorsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MonitorsListOutput represents the monitors list output type.
type MonitorsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                       `json:"object"`
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

// MonitorsListQueryCreatedAt - Only return results whose monitor creation time is in this range. Set `gt`, `lt`, or both.
type MonitorsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryUpdatedAt - Only return results whose monitor last update time is in this range. Set `gt`, `lt`, or both.
type MonitorsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryFirstAlertAt - Only return results whose first monitor alert time is in this range. Set `gt`, `lt`, or both.
type MonitorsListQueryFirstAlertAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQueryLastAlertAt - Only return results whose last monitor alert time is in this range. Set `gt`, `lt`, or both.
type MonitorsListQueryLastAlertAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorsListQuery represents the monitors list query type.
type MonitorsListQuery struct {
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
	// Id - Only return monitors with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return monitors watching one of these targets.
	Target *any `json:"target,omitempty"`
	// Status - Only return monitors with one of these statuses.
	Status *any `json:"status,omitempty"`
	// ProviderId - Only return monitors watching one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProtoGuardFilterId - Only return monitors watching one of these ProtoGuard filters.
	ProtoGuardFilterId *any `json:"proto_guard_filter_id,omitempty"`
	// Search - Filter monitors by name or description.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose monitor creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *MonitorsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose monitor last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *MonitorsListQueryUpdatedAt `json:"updated_at,omitempty"`
	// FirstAlertAt - Only return results whose first monitor alert time is in this range. Set `gt`, `lt`, or both.
	FirstAlertAt *MonitorsListQueryFirstAlertAt `json:"first_alert_at,omitempty"`
	// LastAlertAt - Only return results whose last monitor alert time is in this range. Set `gt`, `lt`, or both.
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
