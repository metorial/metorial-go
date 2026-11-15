package protoguardconfig

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
)

// ProtoGuardConfigUpdateOutputFilters represents the proto guard config update output filters type.
type ProtoGuardConfigUpdateOutputFilters struct {
	Object string `json:"object"`
	// Id - The unique identifier of the filter.
	Id string `json:"id"`
	// Key - The stable key of the filter.
	Key string `json:"key"`
	// Name - The display name of the filter.
	Name string `json:"name"`
	// Description - What the filter detects.
	Description *string `json:"description,omitempty"`
	IssueType   string  `json:"issue_type"`
	Severity    string  `json:"severity"`
	// ScoreWeight - How much a match in this filter contributes to the run score.
	ScoreWeight float64 `json:"score_weight"`
	// DefaultEnabled - Whether the filter is enabled when no override exists.
	DefaultEnabled bool `json:"default_enabled"`
	// Enabled - Whether the filter is enabled for this instance.
	Enabled bool `json:"enabled"`
	// DefaultAlertConfidenceThreshold - The default confidence a match needs for the filter to count as triggered.
	DefaultAlertConfidenceThreshold float64 `json:"default_alert_confidence_threshold"`
	// AlertConfidenceThreshold - The confidence a match needs for the filter to count as triggered.
	AlertConfidenceThreshold float64 `json:"alert_confidence_threshold"`
}

// ProtoGuardConfigUpdateOutput represents the proto guard config update output type.
type ProtoGuardConfigUpdateOutput struct {
	Object string `json:"object"`
	// Enabled - Whether ProtoGuard analyzes session messages for this instance. Enabled by default.
	Enabled bool `json:"enabled"`
	// AlertProbabilityThreshold - The combined injection probability, between 0 and 1, at which a message raises an alert.
	AlertProbabilityThreshold float64 `json:"alert_probability_threshold"`
	// AlertFilterCountThreshold - The number of triggered filters at which a message raises an alert.
	AlertFilterCountThreshold float64                               `json:"alert_filter_count_threshold"`
	Filters                   []ProtoGuardConfigUpdateOutputFilters `json:"filters"`
}

// MapProtoGuardConfigUpdateOutputFromJSON deserializes JSON data into a ProtoGuardConfigUpdateOutput.
func MapProtoGuardConfigUpdateOutputFromJSON(data []byte) (*ProtoGuardConfigUpdateOutput, error) {
	var v ProtoGuardConfigUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigUpdateOutputToJSON serializes a ProtoGuardConfigUpdateOutput to JSON.
func MapProtoGuardConfigUpdateOutputToJSON(v *ProtoGuardConfigUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardConfigUpdateBody represents the proto guard config update body type.
type ProtoGuardConfigUpdateBody struct {
	// Enabled - Whether ProtoGuard analyzes session messages for this instance.
	Enabled *bool `json:"enabled,omitempty"`
	// AlertProbabilityThreshold - The combined injection probability, above 0 and at most 1, at which a message raises an alert. Pass null to restore the default.
	AlertProbabilityThreshold *nullable.Nullable[float64] `json:"alert_probability_threshold,omitempty"`
	// AlertFilterCountThreshold - The number of triggered filters at which a message raises an alert. Pass null to restore the default.
	AlertFilterCountThreshold *nullable.Nullable[float64] `json:"alert_filter_count_threshold,omitempty"`
}

// MapProtoGuardConfigUpdateBodyFromJSON deserializes JSON data into a ProtoGuardConfigUpdateBody.
func MapProtoGuardConfigUpdateBodyFromJSON(data []byte) (*ProtoGuardConfigUpdateBody, error) {
	var v ProtoGuardConfigUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigUpdateBodyToJSON serializes a ProtoGuardConfigUpdateBody to JSON.
func MapProtoGuardConfigUpdateBodyToJSON(v *ProtoGuardConfigUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
