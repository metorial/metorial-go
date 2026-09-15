package protoguardalerts

import (
	"encoding/json"
	"time"
)

// ProtoGuardAlertsGetOutputFilters represents the proto guard alerts get output filters type.
type ProtoGuardAlertsGetOutputFilters struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	FilterId    string    `json:"filter_id"`
	Key         string    `json:"key"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IssueType   string    `json:"issue_type"`
	Severity    string    `json:"severity"`
	Confidence  *float64  `json:"confidence,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProtoGuardAlertsGetOutput represents the proto guard alerts get output type.
type ProtoGuardAlertsGetOutput struct {
	Object              string                             `json:"object"`
	Id                  string                             `json:"id"`
	RunId               string                             `json:"run_id"`
	SessionId           *string                            `json:"session_id,omitempty"`
	SessionMessageId    *string                            `json:"session_message_id,omitempty"`
	SessionConnectionId *string                            `json:"session_connection_id,omitempty"`
	ProviderRunId       *string                            `json:"provider_run_id,omitempty"`
	Filters             []ProtoGuardAlertsGetOutputFilters `json:"filters"`
	CreatedAt           time.Time                          `json:"created_at"`
}

// MapProtoGuardAlertsGetOutputFromJSON deserializes JSON data into a ProtoGuardAlertsGetOutput.
func MapProtoGuardAlertsGetOutputFromJSON(data []byte) (*ProtoGuardAlertsGetOutput, error) {
	var v ProtoGuardAlertsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardAlertsGetOutputToJSON serializes a ProtoGuardAlertsGetOutput to JSON.
func MapProtoGuardAlertsGetOutputToJSON(v *ProtoGuardAlertsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
