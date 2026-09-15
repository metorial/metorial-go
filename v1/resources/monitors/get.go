package monitors

import (
	"encoding/json"
	"time"
)

// MonitorsGetOutput represents the monitors get output type.
type MonitorsGetOutput struct {
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
