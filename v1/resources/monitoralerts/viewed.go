package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsViewedOutputMonitor represents the monitor alerts viewed output monitor type.
type MonitorAlertsViewedOutputMonitor struct {
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

// MonitorAlertsViewedOutputSpecificationChangeNotificationFromSpecification represents the monitor alerts viewed output specification change notification from specification type.
type MonitorAlertsViewedOutputSpecificationChangeNotificationFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsViewedOutputSpecificationChangeNotificationToSpecification represents the monitor alerts viewed output specification change notification to specification type.
type MonitorAlertsViewedOutputSpecificationChangeNotificationToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsViewedOutputSpecificationChangeNotificationFromProviderVersion represents the monitor alerts viewed output specification change notification from provider version type.
type MonitorAlertsViewedOutputSpecificationChangeNotificationFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsViewedOutputSpecificationChangeNotificationToProviderVersion represents the monitor alerts viewed output specification change notification to provider version type.
type MonitorAlertsViewedOutputSpecificationChangeNotificationToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsViewedOutputSpecificationChangeNotification represents the monitor alerts viewed output specification change notification type.
type MonitorAlertsViewedOutputSpecificationChangeNotification struct {
	Object              string                                                                       `json:"object"`
	Id                  string                                                                       `json:"id"`
	ProviderId          string                                                                       `json:"provider_id"`
	ProviderVersionId   string                                                                       `json:"provider_version_id"`
	FromSpecification   *MonitorAlertsViewedOutputSpecificationChangeNotificationFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *MonitorAlertsViewedOutputSpecificationChangeNotificationToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *MonitorAlertsViewedOutputSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *MonitorAlertsViewedOutputSpecificationChangeNotificationToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                    `json:"created_at"`
}

// MonitorAlertsViewedOutputRecipients represents the monitor alerts viewed output recipients type.
type MonitorAlertsViewedOutputRecipients struct {
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	RecipientId string     `json:"recipient_id"`
	ViewedAt    *time.Time `json:"viewed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// MonitorAlertsViewedOutputEvents represents the monitor alerts viewed output events type.
type MonitorAlertsViewedOutputEvents struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ActorId   *string   `json:"actor_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsViewedOutput represents the monitor alerts viewed output type.
type MonitorAlertsViewedOutput struct {
	Object                          string                                                    `json:"object"`
	Id                              string                                                    `json:"id"`
	Status                          string                                                    `json:"status"`
	Monitor                         MonitorAlertsViewedOutputMonitor                          `json:"monitor"`
	ProtoGuardAlertId               *string                                                   `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                 *string                                                   `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsViewedOutputSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	CreatedAt                       time.Time                                                 `json:"created_at"`
	ResolvedAt                      *time.Time                                                `json:"resolved_at,omitempty"`
	Recipients                      []MonitorAlertsViewedOutputRecipients                     `json:"recipients"`
	Events                          []MonitorAlertsViewedOutputEvents                         `json:"events"`
}

// MapMonitorAlertsViewedOutputFromJSON deserializes JSON data into a MonitorAlertsViewedOutput.
func MapMonitorAlertsViewedOutputFromJSON(data []byte) (*MonitorAlertsViewedOutput, error) {
	var v MonitorAlertsViewedOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsViewedOutputToJSON serializes a MonitorAlertsViewedOutput to JSON.
func MapMonitorAlertsViewedOutputToJSON(v *MonitorAlertsViewedOutput) ([]byte, error) {
	return json.Marshal(v)
}
