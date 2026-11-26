package monitors

import (
	"encoding/json"
	"time"
)

// MonitorsGetOutput represents the monitors get output type.
type MonitorsGetOutput struct {
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

// MapMonitorsGetOutputFromJSON deserializes JSON data into a MonitorsGetOutput.
func MapMonitorsGetOutputFromJSON(data []byte) (*MonitorsGetOutput, error) {
	var v MonitorsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapMonitorsGetOutputToJSON serializes a MonitorsGetOutput to JSON.
func MapMonitorsGetOutputToJSON(v *MonitorsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
