package protoguardalerts

import (
	"encoding/json"
	"time"
)

// ProtoGuardAlertsGetOutputFilters - One filter match folded into the alert.
type ProtoGuardAlertsGetOutputFilters struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// FilterId - Filter that matched.
	FilterId string `json:"filter_id"`
	// Key - Stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// IssueType - What kind of prompt-injection pattern was found.
	IssueType string `json:"issue_type"`
	// Severity - How severe the finding is.
	Severity string `json:"severity"`
	// Confidence - Confidence of the match. `null` when not scored.
	Confidence *float64 `json:"confidence,omitempty"`
	// RuleKey - The rule within the filter that matched.
	RuleKey *string `json:"rule_key,omitempty"`
	// Direction - Whether the match was in the tool input, the tool output, or across the session.
	Direction *string `json:"direction,omitempty"`
	// ContentHash - A keyed hash of the matched text. Pass it to a suppression to silence this exact content.
	ContentHash *string `json:"content_hash,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ProtoGuardAlertsGetOutput represents the proto guard alerts get output type.
type ProtoGuardAlertsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// RunId - Run that produced this alert.
	RunId string `json:"run_id"`
	// SessionId - Session the finding came from. `null` when the finding has no session.
	SessionId *string `json:"session_id,omitempty"`
	// SessionMessageId - Message the finding came from. `null` when it spans the session.
	SessionMessageId *string `json:"session_message_id,omitempty"`
	// SessionConnectionId - Connection the finding came from. `null` when there is none.
	SessionConnectionId *string `json:"session_connection_id,omitempty"`
	// ProviderRunId - Provider run the finding came from. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// IssueId - Provider issue this alert is aggregated into.
	IssueId *string `json:"issue_id,omitempty"`
	// Filters - Filter matches folded into this alert.
	Filters []ProtoGuardAlertsGetOutputFilters `json:"filters"`
	// OccurrenceCount - How many times the same finding recurred and was folded into this alert.
	OccurrenceCount float64 `json:"occurrence_count"`
	// FirstSeenAt - When the finding was first seen.
	FirstSeenAt time.Time `json:"first_seen_at"`
	// LastSeenAt - When the finding was last seen.
	LastSeenAt time.Time `json:"last_seen_at"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
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
