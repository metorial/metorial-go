package errors

import (
	"encoding/json"
	"time"
)

// SessionsErrorsGetOutput represents the sessions errors get output type.
type SessionsErrorsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Code - Machine-readable error code.
	Code string `json:"code"`
	// Message - Human-readable summary of the failure.
	Message string `json:"message"`
	// Data - Additional error details.
	Data map[string]any `json:"data"`
	// Status - Indicates whether the error is still being processed or has been fully processed and grouped.
	Status string `json:"status"`
	// SessionId - Session the error occurred in.
	SessionId string `json:"session_id"`
	// ProviderRunId - Provider run the error belongs to. `null` when there is none.
	ProviderRunId *string `json:"provider_run_id,omitempty"`
	// ConnectionId - Connection the error belongs to. `null` when there is none.
	ConnectionId *string `json:"connection_id,omitempty"`
	// GroupId - Group this error belongs to. `null` while the error is still processing.
	GroupId *string `json:"group_id,omitempty"`
	// SimilarErrorCount - How many errors are in this group.
	SimilarErrorCount float64 `json:"similar_error_count"`
	// IssueId - Provider issue this error is aggregated into. `null` when there is none.
	IssueId *string `json:"issue_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapSessionsErrorsGetOutputFromJSON deserializes JSON data into a SessionsErrorsGetOutput.
func MapSessionsErrorsGetOutputFromJSON(data []byte) (*SessionsErrorsGetOutput, error) {
	var v SessionsErrorsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSessionsErrorsGetOutputToJSON serializes a SessionsErrorsGetOutput to JSON.
func MapSessionsErrorsGetOutputToJSON(v *SessionsErrorsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
