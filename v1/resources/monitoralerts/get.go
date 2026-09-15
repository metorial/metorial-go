package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsGetOutputMonitor represents the monitor alerts get output monitor type.
type MonitorAlertsGetOutputMonitor struct {
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

// MonitorAlertsGetOutputSpecificationChangeNotificationFromSpecification represents the monitor alerts get output specification change notification from specification type.
type MonitorAlertsGetOutputSpecificationChangeNotificationFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsGetOutputSpecificationChangeNotificationToSpecification represents the monitor alerts get output specification change notification to specification type.
type MonitorAlertsGetOutputSpecificationChangeNotificationToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsGetOutputSpecificationChangeNotificationFromProviderVersion represents the monitor alerts get output specification change notification from provider version type.
type MonitorAlertsGetOutputSpecificationChangeNotificationFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsGetOutputSpecificationChangeNotificationToProviderVersion represents the monitor alerts get output specification change notification to provider version type.
type MonitorAlertsGetOutputSpecificationChangeNotificationToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsGetOutputSpecificationChangeNotification represents the monitor alerts get output specification change notification type.
type MonitorAlertsGetOutputSpecificationChangeNotification struct {
	Object              string                                                                    `json:"object"`
	Id                  string                                                                    `json:"id"`
	ProviderId          string                                                                    `json:"provider_id"`
	ProviderVersionId   string                                                                    `json:"provider_version_id"`
	FromSpecification   *MonitorAlertsGetOutputSpecificationChangeNotificationFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *MonitorAlertsGetOutputSpecificationChangeNotificationToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *MonitorAlertsGetOutputSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *MonitorAlertsGetOutputSpecificationChangeNotificationToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                 `json:"created_at"`
}

// MonitorAlertsGetOutputRecipients represents the monitor alerts get output recipients type.
type MonitorAlertsGetOutputRecipients struct {
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	RecipientId string     `json:"recipient_id"`
	ViewedAt    *time.Time `json:"viewed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// MonitorAlertsGetOutputEvents represents the monitor alerts get output events type.
type MonitorAlertsGetOutputEvents struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ActorId   *string   `json:"actor_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsGetOutput represents the monitor alerts get output type.
type MonitorAlertsGetOutput struct {
	Object                          string                                                 `json:"object"`
	Id                              string                                                 `json:"id"`
	Status                          string                                                 `json:"status"`
	Monitor                         MonitorAlertsGetOutputMonitor                          `json:"monitor"`
	ProtoGuardAlertId               *string                                                `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                 *string                                                `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsGetOutputSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	CreatedAt                       time.Time                                              `json:"created_at"`
	ResolvedAt                      *time.Time                                             `json:"resolved_at,omitempty"`
	Recipients                      []MonitorAlertsGetOutputRecipients                     `json:"recipients"`
	Events                          []MonitorAlertsGetOutputEvents                         `json:"events"`
}

// MapMonitorAlertsGetOutputFromJSON deserializes JSON data into a MonitorAlertsGetOutput.
func MapMonitorAlertsGetOutputFromJSON(data []byte) (*MonitorAlertsGetOutput, error) {
	var v MonitorAlertsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsGetOutputToJSON serializes a MonitorAlertsGetOutput to JSON.
func MapMonitorAlertsGetOutputToJSON(v *MonitorAlertsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
