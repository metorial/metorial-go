package protoguardconfig

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
)

// ProtoGuardConfigUpdateFilterOutputFilters - One prompt-injection filter in the config.
type ProtoGuardConfigUpdateFilterOutputFilters struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - What the filter detects.
	Description *string `json:"description,omitempty"`
	// IssueType - What kind of prompt-injection pattern was found.
	IssueType string `json:"issue_type"`
	// Severity - How severe the finding is.
	Severity string `json:"severity"`
	// ScoreWeight - How much a match in this filter contributes to the run score.
	ScoreWeight float64 `json:"score_weight"`
	// DefaultEnabled - Whether the filter is enabled when no override exists.
	DefaultEnabled bool `json:"default_enabled"`
	// Enabled - The filter runs for this instance when `true`.
	Enabled bool `json:"enabled"`
	// DefaultAlertConfidenceThreshold - The default confidence a match needs for the filter to count as triggered.
	DefaultAlertConfidenceThreshold float64 `json:"default_alert_confidence_threshold"`
	// AlertConfidenceThreshold - The confidence a match needs for the filter to count as triggered.
	AlertConfidenceThreshold float64 `json:"alert_confidence_threshold"`
}

// ProtoGuardConfigUpdateFilterOutput represents the proto guard config update filter output type.
type ProtoGuardConfigUpdateFilterOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Enabled - Whether ProtoGuard analyzes session messages for this instance. Enabled by default.
	Enabled bool `json:"enabled"`
	// AlertProbabilityThreshold - The combined injection probability, between 0 and 1, at which a message raises an alert.
	AlertProbabilityThreshold float64 `json:"alert_probability_threshold"`
	// AlertFilterCountThreshold - The number of triggered filters at which a message raises an alert.
	AlertFilterCountThreshold float64 `json:"alert_filter_count_threshold"`
	// Filters - Prompt-injection filters and their per-instance settings.
	Filters []ProtoGuardConfigUpdateFilterOutputFilters `json:"filters"`
}

// MapProtoGuardConfigUpdateFilterOutputFromJSON deserializes JSON data into a ProtoGuardConfigUpdateFilterOutput.
func MapProtoGuardConfigUpdateFilterOutputFromJSON(data []byte) (*ProtoGuardConfigUpdateFilterOutput, error) {
	var v ProtoGuardConfigUpdateFilterOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigUpdateFilterOutputToJSON serializes a ProtoGuardConfigUpdateFilterOutput to JSON.
func MapProtoGuardConfigUpdateFilterOutputToJSON(v *ProtoGuardConfigUpdateFilterOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardConfigUpdateFilterBody represents the proto guard config update filter body type.
type ProtoGuardConfigUpdateFilterBody struct {
	// Enabled - Whether the filter runs for the instance.
	Enabled *bool `json:"enabled,omitempty"`
	// AlertConfidenceThreshold - Confidence a match needs for the filter to count as triggered. Pass `null` to restore the default.
	AlertConfidenceThreshold *nullable.Nullable[float64] `json:"alert_confidence_threshold,omitempty"`
}

// MapProtoGuardConfigUpdateFilterBodyFromJSON deserializes JSON data into a ProtoGuardConfigUpdateFilterBody.
func MapProtoGuardConfigUpdateFilterBodyFromJSON(data []byte) (*ProtoGuardConfigUpdateFilterBody, error) {
	var v ProtoGuardConfigUpdateFilterBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigUpdateFilterBodyToJSON serializes a ProtoGuardConfigUpdateFilterBody to JSON.
func MapProtoGuardConfigUpdateFilterBodyToJSON(v *ProtoGuardConfigUpdateFilterBody) ([]byte, error) {
	return json.Marshal(v)
}
