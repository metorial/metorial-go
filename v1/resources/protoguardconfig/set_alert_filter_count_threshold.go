package protoguardconfig

import (
	"encoding/json"
)

// ProtoGuardConfigSetAlertFilterCountThresholdOutputFilters represents the proto guard config set alert filter count threshold output filters type.
type ProtoGuardConfigSetAlertFilterCountThresholdOutputFilters struct {
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

// ProtoGuardConfigSetAlertFilterCountThresholdOutput represents the proto guard config set alert filter count threshold output type.
type ProtoGuardConfigSetAlertFilterCountThresholdOutput struct {
	Object string `json:"object"`
	// Enabled - Whether ProtoGuard analyzes session messages for this instance. Enabled by default.
	Enabled bool `json:"enabled"`
	// AlertProbabilityThreshold - The combined injection probability, between 0 and 1, at which a message raises an alert.
	AlertProbabilityThreshold float64 `json:"alert_probability_threshold"`
	// AlertFilterCountThreshold - The number of triggered filters at which a message raises an alert.
	AlertFilterCountThreshold float64                                                     `json:"alert_filter_count_threshold"`
	Filters                   []ProtoGuardConfigSetAlertFilterCountThresholdOutputFilters `json:"filters"`
}

// MapProtoGuardConfigSetAlertFilterCountThresholdOutputFromJSON deserializes JSON data into a ProtoGuardConfigSetAlertFilterCountThresholdOutput.
func MapProtoGuardConfigSetAlertFilterCountThresholdOutputFromJSON(data []byte) (*ProtoGuardConfigSetAlertFilterCountThresholdOutput, error) {
	var v ProtoGuardConfigSetAlertFilterCountThresholdOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigSetAlertFilterCountThresholdOutputToJSON serializes a ProtoGuardConfigSetAlertFilterCountThresholdOutput to JSON.
func MapProtoGuardConfigSetAlertFilterCountThresholdOutputToJSON(v *ProtoGuardConfigSetAlertFilterCountThresholdOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProtoGuardConfigSetAlertFilterCountThresholdBody represents the proto guard config set alert filter count threshold body type.
type ProtoGuardConfigSetAlertFilterCountThresholdBody struct {
	Threshold *float64 `json:"threshold,omitempty"`
}

// MapProtoGuardConfigSetAlertFilterCountThresholdBodyFromJSON deserializes JSON data into a ProtoGuardConfigSetAlertFilterCountThresholdBody.
func MapProtoGuardConfigSetAlertFilterCountThresholdBodyFromJSON(data []byte) (*ProtoGuardConfigSetAlertFilterCountThresholdBody, error) {
	var v ProtoGuardConfigSetAlertFilterCountThresholdBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigSetAlertFilterCountThresholdBodyToJSON serializes a ProtoGuardConfigSetAlertFilterCountThresholdBody to JSON.
func MapProtoGuardConfigSetAlertFilterCountThresholdBodyToJSON(v *ProtoGuardConfigSetAlertFilterCountThresholdBody) ([]byte, error) {
	return json.Marshal(v)
}
