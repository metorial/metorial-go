package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsUnresolveOutputMonitor represents the monitor alerts unresolve output monitor type.
type MonitorAlertsUnresolveOutputMonitor struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification represents the monitor alerts unresolve output specification change notification from specification type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification represents the monitor alerts unresolve output specification change notification to specification type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion represents the monitor alerts unresolve output specification change notification from provider version type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion represents the monitor alerts unresolve output specification change notification to provider version type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Version     string    `json:"version"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// MonitorAlertsUnresolveOutputSpecificationChangeNotification represents the monitor alerts unresolve output specification change notification type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotification struct {
	Object              string                                                                          `json:"object"`
	Id                  string                                                                          `json:"id"`
	ProviderId          string                                                                          `json:"provider_id"`
	ProviderVersionId   string                                                                          `json:"provider_version_id"`
	FromSpecification   *MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification   `json:"from_specification,omitempty"`
	ToSpecification     *MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification     `json:"to_specification,omitempty"`
	FromProviderVersion *MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	ToProviderVersion   *MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion   `json:"to_provider_version,omitempty"`
	CreatedAt           time.Time                                                                       `json:"created_at"`
}

// MonitorAlertsUnresolveOutputRecipients represents the monitor alerts unresolve output recipients type.
type MonitorAlertsUnresolveOutputRecipients struct {
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	RecipientId string     `json:"recipient_id"`
	ViewedAt    *time.Time `json:"viewed_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}

// MonitorAlertsUnresolveOutputEvents represents the monitor alerts unresolve output events type.
type MonitorAlertsUnresolveOutputEvents struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	ActorId   *string   `json:"actor_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsUnresolveOutput represents the monitor alerts unresolve output type.
type MonitorAlertsUnresolveOutput struct {
	Object                          string                                                       `json:"object"`
	Id                              string                                                       `json:"id"`
	Status                          string                                                       `json:"status"`
	Monitor                         MonitorAlertsUnresolveOutputMonitor                          `json:"monitor"`
	ProtoGuardAlertId               *string                                                      `json:"proto_guard_alert_id,omitempty"`
	ProtoGuardRunId                 *string                                                      `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsUnresolveOutputSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	CreatedAt                       time.Time                                                    `json:"created_at"`
	ResolvedAt                      *time.Time                                                   `json:"resolved_at,omitempty"`
	Recipients                      []MonitorAlertsUnresolveOutputRecipients                     `json:"recipients"`
	Events                          []MonitorAlertsUnresolveOutputEvents                         `json:"events"`
}

// MapMonitorAlertsUnresolveOutputFromJSON deserializes JSON data into a MonitorAlertsUnresolveOutput.
func MapMonitorAlertsUnresolveOutputFromJSON(data []byte) (*MonitorAlertsUnresolveOutput, error) {
	var v MonitorAlertsUnresolveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorAlertsUnresolveOutputToJSON serializes a MonitorAlertsUnresolveOutput to JSON.
func MapMonitorAlertsUnresolveOutputToJSON(v *MonitorAlertsUnresolveOutput) ([]byte, error) {
	return json.Marshal(v)
}
