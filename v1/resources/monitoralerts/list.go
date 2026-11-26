package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsListOutputItemsMonitor represents the monitor alerts list output items monitor type.
type MonitorAlertsListOutputItemsMonitor struct {
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

// MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification - Specification on one side of the change.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification - Specification on one side of the change.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion - Provider version on one side of the change.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion - Provider version on one side of the change.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotification represents the monitor alerts list output items specification change notification type.
type MonitorAlertsListOutputItemsSpecificationChangeNotification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderId - Provider whose specification changed.
	ProviderId string `json:"provider_id"`
	// ProviderVersionId - Provider version whose specification changed.
	ProviderVersionId string `json:"provider_version_id"`
	// FromSpecification - Specification on one side of the change.
	FromSpecification *MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification `json:"from_specification,omitempty"`
	// ToSpecification - Specification on one side of the change.
	ToSpecification *MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification `json:"to_specification,omitempty"`
	// FromProviderVersion - Provider version on one side of the change.
	FromProviderVersion *MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	// ToProviderVersion - Provider version on one side of the change.
	ToProviderVersion *MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion `json:"to_provider_version,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsListOutputItemsRecipients - One actor notified about the alert.
type MonitorAlertsListOutputItemsRecipients struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// RecipientId - Notified actor.
	RecipientId string `json:"recipient_id"`
	// ViewedAt - When the recipient viewed the alert. `null` until viewed.
	ViewedAt *time.Time `json:"viewed_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsListOutputItemsEvents - One entry in the history of the alert.
type MonitorAlertsListOutputItemsEvents struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - What happened.
	Type string `json:"type"`
	// ActorId - Actor behind the event. `null` when no actor performed it.
	ActorId *string `json:"actor_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsListOutputItems represents the monitor alerts list output items type.
type MonitorAlertsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the alert. `pending` until you resolve or ignore it.
	Status  string                              `json:"status"`
	Monitor MonitorAlertsListOutputItemsMonitor `json:"monitor"`
	// ProtoGuardAlertId - ProtoGuard alert behind this alert. `null` for schema-change alerts.
	ProtoGuardAlertId *string `json:"proto_guard_alert_id,omitempty"`
	// ProtoGuardRunId - ProtoGuard run behind this alert. `null` for schema-change alerts.
	ProtoGuardRunId                 *string                                                      `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsListOutputItemsSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ResolvedAt - When the alert was resolved. `null` while pending.
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Recipients - Actors notified about the alert.
	Recipients []MonitorAlertsListOutputItemsRecipients `json:"recipients"`
	// Events - History of the alert.
	Events []MonitorAlertsListOutputItemsEvents `json:"events"`
}

// MonitorAlertsListOutputPagination represents the monitor alerts list output pagination type.
type MonitorAlertsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// MonitorAlertsListOutput represents the monitor alerts list output type.
type MonitorAlertsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []MonitorAlertsListOutputItems    `json:"items"`
	Pagination MonitorAlertsListOutputPagination `json:"pagination"`
}

// MapMonitorAlertsListOutputFromJSON deserializes JSON data into a MonitorAlertsListOutput.
func MapMonitorAlertsListOutputFromJSON(data []byte) (*MonitorAlertsListOutput, error) {
	var v MonitorAlertsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsListOutputToJSON serializes a MonitorAlertsListOutput to JSON.
func MapMonitorAlertsListOutputToJSON(v *MonitorAlertsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// MonitorAlertsListQueryCreatedAt - Only return results whose monitor alert creation time is in this range. Set `gt`, `lt`, or both.
type MonitorAlertsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorAlertsListQueryResolvedAt - Only return results whose monitor alert resolution time is in this range. Set `gt`, `lt`, or both.
type MonitorAlertsListQueryResolvedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorAlertsListQuery represents the monitor alerts list query type.
type MonitorAlertsListQuery struct {
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
	// Id - Only return monitor alerts with one of these IDs.
	Id *any `json:"id,omitempty"`
	// MonitorId - Only return alerts from one of these monitors.
	MonitorId *any `json:"monitor_id,omitempty"`
	// Status - Only return monitor alerts with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Target - Only return monitor alerts watching one of these targets.
	Target *any `json:"target,omitempty"`
	// Source - Only return monitor alerts from one of these sources.
	Source *any `json:"source,omitempty"`
	// ProviderId - Only return monitor alerts for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProtoGuardAlertId - Only return monitor alerts linked to one of these ProtoGuard alerts.
	ProtoGuardAlertId *any `json:"proto_guard_alert_id,omitempty"`
	// ProtoGuardRunId - Only return monitor alerts linked to one of these ProtoGuard runs.
	ProtoGuardRunId *any `json:"proto_guard_run_id,omitempty"`
	// ProtoGuardFilterId - Only return monitor alerts linked to one of these ProtoGuard filters.
	ProtoGuardFilterId *any `json:"proto_guard_filter_id,omitempty"`
	// SpecificationChangeNotificationId - Only return monitor alerts linked to one of these specification change notifications.
	SpecificationChangeNotificationId *any `json:"specification_change_notification_id,omitempty"`
	// SessionId - Only return monitor alerts linked to one of these sessions.
	SessionId *any `json:"session_id,omitempty"`
	// SessionMessageId - Only return monitor alerts linked to one of these session messages.
	SessionMessageId *any `json:"session_message_id,omitempty"`
	// SessionConnectionId - Only return monitor alerts linked to one of these session connections.
	SessionConnectionId *any `json:"session_connection_id,omitempty"`
	// ProviderRunId - Only return monitor alerts linked to one of these provider runs.
	ProviderRunId *any `json:"provider_run_id,omitempty"`
	// CreatedAt - Only return results whose monitor alert creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *MonitorAlertsListQueryCreatedAt `json:"created_at,omitempty"`
	// ResolvedAt - Only return results whose monitor alert resolution time is in this range. Set `gt`, `lt`, or both.
	ResolvedAt *MonitorAlertsListQueryResolvedAt `json:"resolved_at,omitempty"`
}

// MapMonitorAlertsListQueryFromJSON deserializes JSON data into a MonitorAlertsListQuery.
func MapMonitorAlertsListQueryFromJSON(data []byte) (*MonitorAlertsListQuery, error) {
	var v MonitorAlertsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsListQueryToJSON serializes a MonitorAlertsListQuery to JSON.
func MapMonitorAlertsListQueryToJSON(v *MonitorAlertsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
