package monitoralerts

import (
	"encoding/json"
	"time"
)

// MonitorAlertsUnresolveOutputMonitor represents the monitor alerts unresolve output monitor type.
type MonitorAlertsUnresolveOutputMonitor struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification - Specification on one side of the change.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification - Specification on one side of the change.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion - Provider version on one side of the change.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion - Provider version on one side of the change.
type MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion struct {
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

// MonitorAlertsUnresolveOutputSpecificationChangeNotification represents the monitor alerts unresolve output specification change notification type.
type MonitorAlertsUnresolveOutputSpecificationChangeNotification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderId - Provider whose specification changed.
	ProviderId string `json:"provider_id"`
	// ProviderVersionId - Provider version whose specification changed.
	ProviderVersionId string `json:"provider_version_id"`
	// FromSpecification - Specification on one side of the change.
	FromSpecification *MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromSpecification `json:"from_specification,omitempty"`
	// ToSpecification - Specification on one side of the change.
	ToSpecification *MonitorAlertsUnresolveOutputSpecificationChangeNotificationToSpecification `json:"to_specification,omitempty"`
	// FromProviderVersion - Provider version on one side of the change.
	FromProviderVersion *MonitorAlertsUnresolveOutputSpecificationChangeNotificationFromProviderVersion `json:"from_provider_version,omitempty"`
	// ToProviderVersion - Provider version on one side of the change.
	ToProviderVersion *MonitorAlertsUnresolveOutputSpecificationChangeNotificationToProviderVersion `json:"to_provider_version,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MonitorAlertsUnresolveOutputRecipients - One actor notified about the alert.
type MonitorAlertsUnresolveOutputRecipients struct {
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

// MonitorAlertsUnresolveOutputEvents - One entry in the history of the alert.
type MonitorAlertsUnresolveOutputEvents struct {
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

// MonitorAlertsUnresolveOutput represents the monitor alerts unresolve output type.
type MonitorAlertsUnresolveOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current status of the alert. `pending` until you resolve or ignore it.
	Status  string                              `json:"status"`
	Monitor MonitorAlertsUnresolveOutputMonitor `json:"monitor"`
	// ProtoGuardAlertId - ProtoGuard alert behind this alert. `null` for schema-change alerts.
	ProtoGuardAlertId *string `json:"proto_guard_alert_id,omitempty"`
	// ProtoGuardRunId - ProtoGuard run behind this alert. `null` for schema-change alerts.
	ProtoGuardRunId                 *string                                                      `json:"proto_guard_run_id,omitempty"`
	SpecificationChangeNotification *MonitorAlertsUnresolveOutputSpecificationChangeNotification `json:"specification_change_notification,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ResolvedAt - When the alert was resolved. `null` while pending.
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Recipients - Actors notified about the alert.
	Recipients []MonitorAlertsUnresolveOutputRecipients `json:"recipients"`
	// Events - History of the alert.
	Events []MonitorAlertsUnresolveOutputEvents `json:"events"`
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
