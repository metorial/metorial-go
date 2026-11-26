package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/aigateway/virtualmodels/insights"
)

// AiGatewayVirtualModelsInsightsEndpoint provides access to virtual models are model names that pick a real model for every request, using rules, optional signals answered by a decision model, a default, and stickiness. Every change to the definition creates a new version.
type AiGatewayVirtualModelsInsightsEndpoint struct {
	client *endpoint.Client
}

// NewAiGatewayVirtualModelsInsightsEndpoint creates a new AiGatewayVirtualModelsInsightsEndpoint.
func NewAiGatewayVirtualModelsInsightsEndpoint(client *endpoint.Client) *AiGatewayVirtualModelsInsightsEndpoint {
	return &AiGatewayVirtualModelsInsightsEndpoint{client: client}
}

// AiGatewayVirtualModelsInsightsEndpointGetParams contains optional query parameters for Get.
type AiGatewayVirtualModelsInsightsEndpointGetParams struct {
	// Days - How many days back to look. Defaults to 7.
	Days *float64 `json:"days,omitempty"`
	// BaselineModel - A model to estimate the savings against, as if it had served every request
	BaselineModel *string `json:"baseline_model,omitempty"`
}

// Get summarizes how the virtual model routed requests over the last days: requests, failures, cost and latency per rule and model, decision model usage, and optionally the savings compared with sending everything to one model.
func (e *AiGatewayVirtualModelsInsightsEndpoint) Get(virtualModelId string, params *AiGatewayVirtualModelsInsightsEndpointGetParams) (*insights.AiGatewayVirtualModelsInsightsGetOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"ai-gateway", "virtual-models", virtualModelId, "insights"},
		Query: query,
	}
	var result insights.AiGatewayVirtualModelsInsightsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
