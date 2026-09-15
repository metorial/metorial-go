package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsListOutputItemsMonitor represents the monitor alerts list output items monitor type.
type MonitorAlertsListOutputItemsMonitor struct {
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

// MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification represents the monitor alerts list output items specification change notification from specification type.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification represents the monitor alerts list output items specification change notification to specification type.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion represents the monitor alerts list output items specification change notification from provider version type.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion represents the monitor alerts list output items specification change notification to provider version type.
type MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsListOutputItemsSpecificationChangeNotification represents the monitor alerts list output items specification change notification type.
type MonitorAlertsListOutputItemsSpecificationChangeNotification struct {
	Object              string                                                                          `json:"object"`
	Id                  string                                                                          `json:"id"`
	ProviderId          string                                                                          `json:"provider_id"`
	ProviderVersionId   string                                                                          `json:"provider_version_id"`
	FromSpecification   *MonitorAlertsListOutputItemsSpecificationChangeNotificationFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *MonitorAlertsListOutputItemsSpecificationChangeNotificationToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *MonitorAlertsListOutputItemsSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *MonitorAlertsListOutputItemsSpecificationChangeNotificationToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                       `json:"created_at"`
}

// MonitorAlertsListOutputItemsRecipients represents the monitor alerts list output items recipients type.
type MonitorAlertsListOutputItemsRecipients struct {
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	RecipientId string     `json:"recipient_id"`
	ViewedAt    *time.Time `json:"viewed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// MonitorAlertsListOutputItemsEvents represents the monitor alerts list output items events type.
type MonitorAlertsListOutputItemsEvents struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ActorId   *string   `json:"actor_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsListOutputItems represents the monitor alerts list output items type.
type MonitorAlertsListOutputItems struct {
	Object                          string                                                       `json:"object"`
	Id                              string                                                       `json:"id"`
	Status                          string                                                       `json:"status"`
	Monitor                         MonitorAlertsListOutputItemsMonitor                          `json:"monitor"`
	ProtoGuardAlertId               *string                                                      `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                 *string                                                      `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsListOutputItemsSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	CreatedAt                       time.Time                                                    `json:"created_at"`
	ResolvedAt                      *time.Time                                                   `json:"resolved_at,omitempty"`
	Recipients                      []MonitorAlertsListOutputItemsRecipients                     `json:"recipients"`
	Events                          []MonitorAlertsListOutputItemsEvents                         `json:"events"`
}

// MonitorAlertsListOutputPagination represents the monitor alerts list output pagination type.
type MonitorAlertsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// MonitorAlertsListOutput represents the monitor alerts list output type.
type MonitorAlertsListOutput struct {
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

// MonitorAlertsListQueryCreatedAt - Filter monitor alert creation time by date range
type MonitorAlertsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for monitor alert creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for monitor alert creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorAlertsListQueryResolvedAt - Filter monitor alert resolution time by date range
type MonitorAlertsListQueryResolvedAt struct {
	// Gt - Only include records after this timestamp for monitor alert resolution time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for monitor alert resolution time
	Lt *time.Time `json:"lt,omitempty"`
}

// MonitorAlertsListQuery represents the monitor alerts list query type.
type MonitorAlertsListQuery struct {
	Limit                             *float64 `json:"limit,omitempty"`
	After                             *string  `json:"after,omitempty"`
	Before                            *string  `json:"before,omitempty"`
	Cursor                            *string  `json:"cursor,omitempty"`
	Order                             *string  `json:"order,omitempty"`
	Id                                *any     `json:"id,omitempty"`
	MonitorId                         *any     `json:"monitor_id,omitempty"`
	Status                            *any     `json:"status,omitempty"`
	Target                            *any     `json:"target,omitempty"`
	Source                            *any     `json:"source,omitempty"`
	ProviderId                        *any     `json:"provider_id,omitempty"`
	ProtoGuardAlertId                 *any     `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                   *any     `json:"proto_guard_run_id,omitempty"`
	ProtoGuardFilterId                *any     `json:"proto_guard_filter_id,omitempty"`
	SpecificationChangeNotificationId *any     `json:"specification_change_notification_id,omitempty"`
	SessionId                         *any     `json:"session_id,omitempty"`
	SessionMessageId                  *any     `json:"session_message_id,omitempty"`
	SessionConnectionId               *any     `json:"session_connection_id,omitempty"`
	ProviderRunId                     *any     `json:"provider_run_id,omitempty"`
	// CreatedAt - Filter monitor alert creation time by date range
	CreatedAt *MonitorAlertsListQueryCreatedAt `json:"created_at,omitempty"`
	// ResolvedAt - Filter monitor alert resolution time by date range
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
