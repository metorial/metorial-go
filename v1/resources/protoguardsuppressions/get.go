package protoguardsuppressions

import (
	"encoding/json"
	"time"
)

// ProtoGuardSuppressionsGetOutputFilter - Filter the suppression applies to.
type ProtoGuardSuppressionsGetOutputFilter struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
}

// ProtoGuardSuppressionsGetOutput represents the proto guard suppressions get output type.
type ProtoGuardSuppressionsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the suppression currently hides matching findings.
	Status string `json:"status"`
	// Reason - Why the findings are considered benign.
	Reason *string `json:"reason,omitempty"`
	// Filter - Filter the suppression applies to.
	Filter ProtoGuardSuppressionsGetOutputFilter `json:"filter"`
	// ProviderId - The provider whose findings are suppressed, or null to suppress across all providers.
	ProviderId *string `json:"provider_id,omitempty"`
	// ToolKey - The tool whose findings are suppressed, or null to suppress across all tools.
	ToolKey *string `json:"tool_key,omitempty"`
	// ContentHash - The keyed content hash of a specific matched text, taken from a ProtoGuard finding, or null to suppress any content.
	ContentHash *string `json:"content_hash,omitempty"`
	// ExpiresAt - When the suppression stops applying, or null if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProtoGuardSuppressionsGetOutputFromJSON deserializes JSON data into a ProtoGuardSuppressionsGetOutput.
func MapProtoGuardSuppressionsGetOutputFromJSON(data []byte) (*ProtoGuardSuppressionsGetOutput, error) {
	var v ProtoGuardSuppressionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardSuppressionsGetOutputToJSON serializes a ProtoGuardSuppressionsGetOutput to JSON.
func MapProtoGuardSuppressionsGetOutputToJSON(v *ProtoGuardSuppressionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
