package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/config"
)

// AiGatewayConfigEndpoint provides access to settings, limits and log retention for every request made through the AI gateway in this instance.
type AiGatewayConfigEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayConfigEndpoint creates a new AiGatewayConfigEndpoint.
func NewAiGatewayConfigEndpoint(client *endpoint.Client) *AiGatewayConfigEndpoint {
	return &AiGatewayConfigEndpoint{client: client}
}

// AiGatewayConfigEndpointUpdateBody contains the request body for Update.
type AiGatewayConfigEndpointUpdateBody struct {
	// LogRetentionDays - How many days inference logs are kept, from 1 to 400. Usage timelines reach back no further. Applies to requests made after the change.
	LogRetentionDays *float64 `json:"log_retention_days,omitempty"`
	// Settings - The settings to change. Fields you omit keep their current value.
	Settings *map[string]any `json:"settings,omitempty"`
	// Limits - Limits that reject requests once reached. Replaces the existing limits when updated.
	Limits *[]map[string]any `json:"limits,omitempty"`
}

// Get retrieves the instance's AI gateway settings and limits, along with the limits and restrictions of your plan.
func (e *AiGatewayConfigEndpoint) Get() (*config.AiGatewayConfigGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "config"},
	}
	var result config.AiGatewayConfigGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the log retention, settings or limits. Fields you omit keep their current value.
func (e *AiGatewayConfigEndpoint) Update(body *AiGatewayConfigEndpointUpdateBody) (*config.AiGatewayConfigUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"ai-gateway", "config"},
		Body: body,
	}
	var result config.AiGatewayConfigUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
