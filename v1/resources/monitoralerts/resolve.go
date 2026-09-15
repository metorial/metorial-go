package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsResolveOutputMonitor represents the monitor alerts resolve output monitor type.
type MonitorAlertsResolveOutputMonitor struct {
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

// MonitorAlertsResolveOutputSpecificationChangeNotificationFromSpecification represents the monitor alerts resolve output specification change notification from specification type.
type MonitorAlertsResolveOutputSpecificationChangeNotificationFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsResolveOutputSpecificationChangeNotificationToSpecification represents the monitor alerts resolve output specification change notification to specification type.
type MonitorAlertsResolveOutputSpecificationChangeNotificationToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsResolveOutputSpecificationChangeNotificationFromProviderVersion represents the monitor alerts resolve output specification change notification from provider version type.
type MonitorAlertsResolveOutputSpecificationChangeNotificationFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsResolveOutputSpecificationChangeNotificationToProviderVersion represents the monitor alerts resolve output specification change notification to provider version type.
type MonitorAlertsResolveOutputSpecificationChangeNotificationToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsResolveOutputSpecificationChangeNotification represents the monitor alerts resolve output specification change notification type.
type MonitorAlertsResolveOutputSpecificationChangeNotification struct {
	Object              string                                                                        `json:"object"`
	Id                  string                                                                        `json:"id"`
	ProviderId          string                                                                        `json:"provider_id"`
	ProviderVersionId   string                                                                        `json:"provider_version_id"`
	FromSpecification   *MonitorAlertsResolveOutputSpecificationChangeNotificationFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *MonitorAlertsResolveOutputSpecificationChangeNotificationToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *MonitorAlertsResolveOutputSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *MonitorAlertsResolveOutputSpecificationChangeNotificationToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                     `json:"created_at"`
}

// MonitorAlertsResolveOutputRecipients represents the monitor alerts resolve output recipients type.
type MonitorAlertsResolveOutputRecipients struct {
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	RecipientId string     `json:"recipient_id"`
	ViewedAt    *time.Time `json:"viewed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// MonitorAlertsResolveOutputEvents represents the monitor alerts resolve output events type.
type MonitorAlertsResolveOutputEvents struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ActorId   *string   `json:"actor_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsResolveOutput represents the monitor alerts resolve output type.
type MonitorAlertsResolveOutput struct {
	Object                          string                                                     `json:"object"`
	Id                              string                                                     `json:"id"`
	Status                          string                                                     `json:"status"`
	Monitor                         MonitorAlertsResolveOutputMonitor                          `json:"monitor"`
	ProtoGuardAlertId               *string                                                    `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                 *string                                                    `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsResolveOutputSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	CreatedAt                       time.Time                                                  `json:"created_at"`
	ResolvedAt                      *time.Time                                                 `json:"resolved_at,omitempty"`
	Recipients                      []MonitorAlertsResolveOutputRecipients                     `json:"recipients"`
	Events                          []MonitorAlertsResolveOutputEvents                         `json:"events"`
}

// MapMonitorAlertsResolveOutputFromJSON deserializes JSON data into a MonitorAlertsResolveOutput.
func MapMonitorAlertsResolveOutputFromJSON(data []byte) (*MonitorAlertsResolveOutput, error) {
	var v MonitorAlertsResolveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsResolveOutputToJSON serializes a MonitorAlertsResolveOutput to JSON.
func MapMonitorAlertsResolveOutputToJSON(v *MonitorAlertsResolveOutput) ([]byte, error) {
	return json.Marshal(v)
}
