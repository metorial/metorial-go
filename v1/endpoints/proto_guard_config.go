package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/nullable"
	"github.com/metorial/metorial-go/v1/resources/protoguardconfig"
)

// ProtoGuardConfigEndpoint provides access to protoGuard config is the prompt-injection filter set for an instance. It controls which filters run and when they raise alerts.
type ProtoGuardConfigEndpoint struct {
	client *endpoint.Client
}

// NewProtoGuardConfigEndpoint creates a new ProtoGuardConfigEndpoint.
func NewProtoGuardConfigEndpoint(client *endpoint.Client) *ProtoGuardConfigEndpoint {
	return &ProtoGuardConfigEndpoint{client: client}
}

// ProtoGuardConfigEndpointUpdateBody contains the request body for Update.
type ProtoGuardConfigEndpointUpdateBody struct {
	// Enabled - Analyzes session messages for prompt injection when `true`. Set to `false` to stop analysis.
	Enabled *bool `json:"enabled,omitempty"`
	// AlertProbabilityThreshold - The combined injection probability, above 0 and at most 1, at which a message raises an alert. Pass null to restore the default.
	AlertProbabilityThreshold *nullable.Nullable[float64] `json:"alert_probability_threshold,omitempty"`
	// AlertFilterCountThreshold - The number of triggered filters at which a message raises an alert. Pass null to restore the default.
	AlertFilterCountThreshold *nullable.Nullable[float64] `json:"alert_filter_count_threshold,omitempty"`
}

// ProtoGuardConfigEndpointUpdateFilterBody contains the request body for UpdateFilter.
type ProtoGuardConfigEndpointUpdateFilterBody struct {
	// Enabled - Whether the filter runs for the instance.
	Enabled *bool `json:"enabled,omitempty"`
	// AlertConfidenceThreshold - Confidence a match needs for the filter to count as triggered. Pass `null` to restore the default.
	AlertConfidenceThreshold *nullable.Nullable[float64] `json:"alert_confidence_threshold,omitempty"`
}

// ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody contains the request body for SetAlertFilterCountThreshold.
type ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody struct {
	// Threshold - Number of matching ProtoGuard filters required to create an alert. Pass `null` to clear it.
	Threshold *float64 `json:"threshold,omitempty"`
}

// Get retrieves the ProtoGuard config for the instance.
func (e *ProtoGuardConfigEndpoint) Get() (*protoguardconfig.ProtoGuardConfigGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"protoguard-config"},
	}
	var result protoguardconfig.ProtoGuardConfigGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update enables or disables ProtoGuard for this instance and sets its alert thresholds.
func (e *ProtoGuardConfigEndpoint) Update(body *ProtoGuardConfigEndpointUpdateBody) (*protoguardconfig.ProtoGuardConfigUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"protoguard-config"},
		Body: body,
	}
	var result protoguardconfig.ProtoGuardConfigUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFilter updates ProtoGuard filter settings for this instance.
func (e *ProtoGuardConfigEndpoint) UpdateFilter(filterId string, body *ProtoGuardConfigEndpointUpdateFilterBody) (*protoguardconfig.ProtoGuardConfigUpdateFilterOutput, error) {
	req := &endpoint.Request{
		Path: []string{"protoguard-config", "filters", filterId},
		Body: body,
	}
	var result protoguardconfig.ProtoGuardConfigUpdateFilterOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetAlertFilterCountThreshold sets or clears the number of matching ProtoGuard filters required to create an alert.
func (e *ProtoGuardConfigEndpoint) SetAlertFilterCountThreshold(body *ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody) (*protoguardconfig.ProtoGuardConfigSetAlertFilterCountThresholdOutput, error) {
	req := &endpoint.Request{
		Path: []string{"protoguard-config", "alert-filter-count-threshold"},
		Body: body,
	}
	var result protoguardconfig.ProtoGuardConfigSetAlertFilterCountThresholdOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
