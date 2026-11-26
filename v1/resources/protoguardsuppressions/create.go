package protoguardsuppressions

import (
	"encoding/json"
	"time"
)

// ProtoGuardSuppressionsCreateOutputFilter - Filter the suppression applies to.
type ProtoGuardSuppressionsCreateOutputFilter struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
}

// ProtoGuardSuppressionsCreateOutput represents the proto guard suppressions create output type.
type ProtoGuardSuppressionsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the suppression currently hides matching findings.
	Status string `json:"status"`
	// Reason - Why the findings are considered benign.
	Reason *string `json:"reason,omitempty"`
	// Filter - Filter the suppression applies to.
	Filter ProtoGuardSuppressionsCreateOutputFilter `json:"filter"`
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

// MapProtoGuardSuppressionsCreateOutputFromJSON deserializes JSON data into a ProtoGuardSuppressionsCreateOutput.
func MapProtoGuardSuppressionsCreateOutputFromJSON(data []byte) (*ProtoGuardSuppressionsCreateOutput, error) {
	var v ProtoGuardSuppressionsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardSuppressionsCreateOutputToJSON serializes a ProtoGuardSuppressionsCreateOutput to JSON.
func MapProtoGuardSuppressionsCreateOutputToJSON(v *ProtoGuardSuppressionsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardSuppressionsCreateBody represents the proto guard suppressions create body type.
type ProtoGuardSuppressionsCreateBody struct {
	// FilterId - The filter to suppress, by id or key.
	FilterId string `json:"filter_id"`
	// ProviderId - Only suppress findings from this provider.
	ProviderId *string `json:"provider_id,omitempty"`
	// ToolKey - Only suppress findings from this tool.
	ToolKey *string `json:"tool_key,omitempty"`
	// ContentHash - Only suppress findings with this content hash, as reported on a ProtoGuard alert finding.
	ContentHash *string `json:"content_hash,omitempty"`
	// Reason - Why the findings are considered benign.
	Reason *string `json:"reason,omitempty"`
	// ExpiresAt - When the suppression stops applying. Omit to never expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapProtoGuardSuppressionsCreateBodyFromJSON deserializes JSON data into a ProtoGuardSuppressionsCreateBody.
func MapProtoGuardSuppressionsCreateBodyFromJSON(data []byte) (*ProtoGuardSuppressionsCreateBody, error) {
	var v ProtoGuardSuppressionsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardSuppressionsCreateBodyToJSON serializes a ProtoGuardSuppressionsCreateBody to JSON.
func MapProtoGuardSuppressionsCreateBodyToJSON(v *ProtoGuardSuppressionsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
