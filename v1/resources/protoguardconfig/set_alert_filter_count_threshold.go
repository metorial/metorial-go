package protoguardconfig

import (
	"encoding/json"
)

// ProtoGuardConfigSetAlertFilterCountThresholdOutputFilters represents the proto guard config set alert filter count threshold output filters type.
type ProtoGuardConfigSetAlertFilterCountThresholdOutputFilters struct {
	Object                          string  `json:"object"`
	Id                              string  `json:"id"`
	Key                             string  `json:"key"`
	Name                            string  `json:"name"`
	Description                     *string `json:"description,omitempty"`
	IssueType                       string  `json:"issue_type"`
	Severity                        string  `json:"severity"`
	ScoreWeight                     float64 `json:"score_weight"`
	DefaultEnabled                  bool    `json:"default_enabled"`
	Enabled                         bool    `json:"enabled"`
	DefaultAlertConfidenceThreshold float64 `json:"default_alert_confidence_threshold"`
	AlertConfidenceThreshold        float64 `json:"alert_confidence_threshold"`
}

// ProtoGuardConfigSetAlertFilterCountThresholdOutput represents the proto guard config set alert filter count threshold output type.
type ProtoGuardConfigSetAlertFilterCountThresholdOutput struct {
	Object                    string                                                      `json:"object"`
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
