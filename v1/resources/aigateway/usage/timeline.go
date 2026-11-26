package usage

import (
	"encoding/json"
	"time"
)

// AiGatewayUsageTimelineOutputBuckets represents the ai gateway usage timeline output buckets type.
type AiGatewayUsageTimelineOutputBuckets struct {
	// Start - When the bucket starts
	Start time.Time `json:"start"`
	// Requests - All requests
	Requests float64 `json:"requests"`
	// Succeeded - Requests that succeeded
	Succeeded float64 `json:"succeeded"`
	// Failed - Requests that were refused or broken off by the provider
	Failed float64 `json:"failed"`
	// Cancelled - Requests the client cancelled
	Cancelled float64 `json:"cancelled"`
	// Rejected - Requests rejected before reaching a provider, e.g. by limits or invalid parameters
	Rejected float64 `json:"rejected"`
	// InputTokens - All input tokens
	InputTokens float64 `json:"input_tokens"`
	// CachedInputTokens - Input tokens read from the prompt cache
	CachedInputTokens float64 `json:"cached_input_tokens"`
	// OutputTokens - All output tokens
	OutputTokens float64 `json:"output_tokens"`
	// CostUsd - What the requests cost in USD, including model-made routing decisions
	CostUsd float64 `json:"cost_usd"`
	// MetorialCostUsd - The part of cost_usd billed by Metorial, for requests served by the Metorial provider
	MetorialCostUsd float64 `json:"metorial_cost_usd"`
	// AverageLatencyMs - The average request duration, in ms. Null when there were no requests.
	AverageLatencyMs *float64 `json:"average_latency_ms,omitempty"`
}

// AiGatewayUsageTimelineOutput represents the ai gateway usage timeline output type.
type AiGatewayUsageTimelineOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Interval - The width of every bucket
	Interval string `json:"interval"`
	// From - The start of the first bucket. Rounded down to a whole interval.
	From time.Time `json:"from"`
	// To - The end of the range, exclusive
	To time.Time `json:"to"`
	// Buckets - Usage per interval, oldest first. Intervals without requests are included with zeros.
	Buckets []AiGatewayUsageTimelineOutputBuckets `json:"buckets"`
}

// MapAiGatewayUsageTimelineOutputFromJSON deserializes JSON data into a AiGatewayUsageTimelineOutput.
func MapAiGatewayUsageTimelineOutputFromJSON(data []byte) (*AiGatewayUsageTimelineOutput, error) {
	var v AiGatewayUsageTimelineOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayUsageTimelineOutputToJSON serializes a AiGatewayUsageTimelineOutput to JSON.
func MapAiGatewayUsageTimelineOutputToJSON(v *AiGatewayUsageTimelineOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayUsageTimelineQuery represents the ai gateway usage timeline query type.
type AiGatewayUsageTimelineQuery struct {
	// From - Start of the range, inclusive. Rounded down to a whole interval.
	From time.Time `json:"from"`
	// To - End of the range, exclusive. Must be after from.
	To time.Time `json:"to"`
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

// MapAiGatewayUsageTimelineQueryFromJSON deserializes JSON data into a AiGatewayUsageTimelineQuery.
func MapAiGatewayUsageTimelineQueryFromJSON(data []byte) (*AiGatewayUsageTimelineQuery, error) {
	var v AiGatewayUsageTimelineQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayUsageTimelineQueryToJSON serializes a AiGatewayUsageTimelineQuery to JSON.
func MapAiGatewayUsageTimelineQueryToJSON(v *AiGatewayUsageTimelineQuery) ([]byte, error) {
	return json.Marshal(v)
}
