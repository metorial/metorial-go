package protoguardconfig

import (
	"encoding/json"
)

// ProtoGuardConfigGetOutputFilters represents the proto guard config get output filters type.
type ProtoGuardConfigGetOutputFilters struct {
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

// ProtoGuardConfigGetOutput represents the proto guard config get output type.
type ProtoGuardConfigGetOutput struct {
	Object                    string                             `json:"object"`
	AlertFilterCountThreshold float64                            `json:"alert_filter_count_threshold"`
	Filters                   []ProtoGuardConfigGetOutputFilters `json:"filters"`
}

// MapProtoGuardConfigGetOutputFromJSON deserializes JSON data into a ProtoGuardConfigGetOutput.
func MapProtoGuardConfigGetOutputFromJSON(data []byte) (*ProtoGuardConfigGetOutput, error) {
	var v ProtoGuardConfigGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProtoGuardConfigGetOutputToJSON serializes a ProtoGuardConfigGetOutput to JSON.
func MapProtoGuardConfigGetOutputToJSON(v *ProtoGuardConfigGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
