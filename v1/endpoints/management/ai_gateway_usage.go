package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/usage"
)

// AiGatewayUsageEndpoint provides access to requests, tokens and cost of the AI gateway over time. Built from inference logs, so it covers the configured log retention.
type AiGatewayUsageEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayUsageEndpoint creates a new AiGatewayUsageEndpoint.
func NewAiGatewayUsageEndpoint(client *endpoint.Client) *AiGatewayUsageEndpoint {
	return &AiGatewayUsageEndpoint{client: client}
}

// AiGatewayUsageEndpointTimelineParams contains optional query parameters for Timeline.
type AiGatewayUsageEndpointTimelineParams struct {
	// From - Start of the range, inclusive. Rounded down to a whole interval.
	From string `json:"from"`
	// To - End of the range, exclusive. Must be after from.
	To string `json:"to"`
	// Interval - The width of every bucket. Omit to pick the smallest interval that gives at most 500 buckets.
	Interval *string `json:"interval,omitempty"`
	// TokenId - Only requests made with this AI gateway token
	TokenId *string `json:"token_id,omitempty"`
	// ConsumerId - Only requests made with tokens of this consumer
	ConsumerId *string `json:"consumer_id,omitempty"`
	// Model - Only requests served by this model, by slug
	Model *string `json:"model,omitempty"`
	// VirtualModelId - Only requests routed by this virtual model
	VirtualModelId *string `json:"virtual_model_id,omitempty"`
}

// Timeline returns usage in fixed-width buckets aligned to UTC, oldest first. A range can span at most 1500 buckets.
func (e *AiGatewayUsageEndpoint) Timeline(instanceId string, params *AiGatewayUsageEndpointTimelineParams) (*usage.AiGatewayUsageTimelineOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "ai-gateway", "usage", "timeline"},
		Query: query,
	}
	var result usage.AiGatewayUsageTimelineOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
