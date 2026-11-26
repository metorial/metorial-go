package errorgroups

import (
	"encoding/json"
	"time"
)

// SessionsErrorGroupsGetOutput represents the sessions error groups get output type.
type SessionsErrorGroupsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Error group data from the first occurrence.
	Data map[string]any `json:"data"`
	// ProviderId - Provider the errors belong to. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// OccurrenceCount - How many errors are in this group.
	OccurrenceCount float64 `json:"occurrence_count"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapSessionsErrorGroupsGetOutputFromJSON deserializes JSON data into a SessionsErrorGroupsGetOutput.
func MapSessionsErrorGroupsGetOutputFromJSON(data []byte) (*SessionsErrorGroupsGetOutput, error) {
	var v SessionsErrorGroupsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorGroupsGetOutputToJSON serializes a SessionsErrorGroupsGetOutput to JSON.
func MapSessionsErrorGroupsGetOutputToJSON(v *SessionsErrorGroupsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
