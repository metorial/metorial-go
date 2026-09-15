package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/protoguardconfig"
)

// ProtoGuardConfigEndpoint provides access to protoGuard config controls prompt-injection filters and alert thresholds.
type ProtoGuardConfigEndpoint struct {
	client *endpoint.Client
}

// NewProtoGuardConfigEndpoint creates a new ProtoGuardConfigEndpoint.
func NewProtoGuardConfigEndpoint(client *endpoint.Client) *ProtoGuardConfigEndpoint {
	return &ProtoGuardConfigEndpoint{client: client}
}

// ProtoGuardConfigEndpointUpdateFilterBody contains the request body for UpdateFilter.
type ProtoGuardConfigEndpointUpdateFilterBody struct {
	Enabled                  *bool    `json:"enabled,omitempty"`
	AlertConfidenceThreshold *float64 `json:"alert_confidence_threshold,omitempty"`
}

// ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody contains the request body for SetAlertFilterCountThreshold.
type ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody struct {
	Threshold *float64 `json:"threshold,omitempty"`
}

// Get retrieves ProtoGuard filter configuration for this instance.
func (e *ProtoGuardConfigEndpoint) Get(instanceId string) (*protoguardconfig.ProtoGuardConfigGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-config"},
	}
	var result protoguardconfig.ProtoGuardConfigGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateFilter updates ProtoGuard filter settings for this instance.
func (e *ProtoGuardConfigEndpoint) UpdateFilter(instanceId string, filterId string, body *ProtoGuardConfigEndpointUpdateFilterBody) (*protoguardconfig.ProtoGuardConfigUpdateFilterOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-config", "filters", filterId},
		Body: body,
	}
	var result protoguardconfig.ProtoGuardConfigUpdateFilterOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SetAlertFilterCountThreshold sets or clears the number of matching ProtoGuard filters required to create an alert.
func (e *ProtoGuardConfigEndpoint) SetAlertFilterCountThreshold(instanceId string, body *ProtoGuardConfigEndpointSetAlertFilterCountThresholdBody) (*protoguardconfig.ProtoGuardConfigSetAlertFilterCountThresholdOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-config", "alert-filter-count-threshold"},
		Body: body,
	}
	var result protoguardconfig.ProtoGuardConfigSetAlertFilterCountThresholdOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
