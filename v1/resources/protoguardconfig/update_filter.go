package protoguardconfig

import (
	"encoding/json"
)

// ProtoGuardConfigUpdateFilterOutputFilters represents the proto guard config update filter output filters type.
type ProtoGuardConfigUpdateFilterOutputFilters struct {
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

// ProtoGuardConfigUpdateFilterOutput represents the proto guard config update filter output type.
type ProtoGuardConfigUpdateFilterOutput struct {
	Object                    string                                      `json:"object"`
	AlertFilterCountThreshold float64                                     `json:"alert_filter_count_threshold"`
	Filters                   []ProtoGuardConfigUpdateFilterOutputFilters `json:"filters"`
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
	Enabled                  *bool    `json:"enabled,omitempty"`
	AlertConfidenceThreshold *float64 `json:"alert_confidence_threshold,omitempty"`
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
