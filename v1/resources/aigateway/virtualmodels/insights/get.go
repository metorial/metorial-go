package insights

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsInsightsGetOutputRoutes represents the ai gateway virtual models insights get output routes type.
type AiGatewayVirtualModelsInsightsGetOutputRoutes struct {
	// RuleId - The rule that picked the model. Null when the default was used.
	RuleId *string `json:"rule_id,omitempty"`
	// Model - The model that served the requests. Null if none did.
	Model *string `json:"model,omitempty"`
	// Requests - The number of requests
	Requests float64 `json:"requests"`
	// Failed - How many of them failed
	Failed float64 `json:"failed"`
	// BilledUsd - What they were billed, in USD
	BilledUsd float64 `json:"billed_usd"`
	// InputTokens - Their input tokens
	InputTokens float64 `json:"input_tokens"`
	// OutputTokens - Their output tokens
	OutputTokens float64 `json:"output_tokens"`
	// AvgTotalMs - Their average total duration, in milliseconds
	AvgTotalMs *float64 `json:"avg_total_ms,omitempty"`
}

// AiGatewayVirtualModelsInsightsGetOutputDecision - Decision model usage in the period
type AiGatewayVirtualModelsInsightsGetOutputDecision struct {
	// Calls - Requests that needed the decision model
	Calls float64 `json:"calls"`
	// Cached - How many of them reused earlier answers
	Cached float64 `json:"cached"`
	// Failed - How many of them failed
	Failed float64 `json:"failed"`
	// CostUsd - What the decision model cost, in USD
	CostUsd float64 `json:"cost_usd"`
	// AvgLatencyMs - The average decision latency, in milliseconds
	AvgLatencyMs *float64 `json:"avg_latency_ms,omitempty"`
}

// AiGatewayVirtualModelsInsightsGetOutputSavings - An estimate of what was saved compared with sending every request to one model. Null unless a baseline model was given.
type AiGatewayVirtualModelsInsightsGetOutputSavings struct {
	// BaselineModel - The model compared against
	BaselineModel string `json:"baseline_model"`
	// BaselineUsd - An estimate of what the same tokens would have cost on it, in USD
	BaselineUsd float64 `json:"baseline_usd"`
	// SpentUsd - What was billed, including the decision model, in USD
	SpentUsd float64 `json:"spent_usd"`
	// SavedUsd - baseline_usd minus spent_usd. Negative when the baseline is cheaper.
	SavedUsd float64 `json:"saved_usd"`
	// Priced - Whether the baseline model has prices for every kind of token used
	Priced bool `json:"priced"`
}

// AiGatewayVirtualModelsInsightsGetOutput represents the ai gateway virtual models insights get output type.
type AiGatewayVirtualModelsInsightsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Since - The start of the period the insights cover
	Since time.Time `json:"since"`
	// Days - The length of the period, in days
	Days float64 `json:"days"`
	// Routes - Requests grouped by the rule and model that served them, most used first
	Routes []AiGatewayVirtualModelsInsightsGetOutputRoutes `json:"routes"`
	// Decision - Decision model usage in the period
	Decision AiGatewayVirtualModelsInsightsGetOutputDecision `json:"decision"`
	// BilledUsd - What all requests were billed, in USD, excluding the decision model
	BilledUsd float64 `json:"billed_usd"`
	// Savings - An estimate of what was saved compared with sending every request to one model. Null unless a baseline model was given.
	Savings *AiGatewayVirtualModelsInsightsGetOutputSavings `json:"savings,omitempty"`
}

// MapAiGatewayVirtualModelsInsightsGetOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsInsightsGetOutput.
func MapAiGatewayVirtualModelsInsightsGetOutputFromJSON(data []byte) (*AiGatewayVirtualModelsInsightsGetOutput, error) {
	var v AiGatewayVirtualModelsInsightsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsInsightsGetOutputToJSON serializes a AiGatewayVirtualModelsInsightsGetOutput to JSON.
func MapAiGatewayVirtualModelsInsightsGetOutputToJSON(v *AiGatewayVirtualModelsInsightsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsInsightsGetQuery represents the ai gateway virtual models insights get query type.
type AiGatewayVirtualModelsInsightsGetQuery struct {
	// Days - How many days back to look. Defaults to 7.
	Days *float64 `json:"days,omitempty"`
	// BaselineModel - A model to estimate the savings against, as if it had served every request
	BaselineModel *string `json:"baseline_model,omitempty"`
}

// MapAiGatewayVirtualModelsInsightsGetQueryFromJSON deserializes JSON data into a AiGatewayVirtualModelsInsightsGetQuery.
func MapAiGatewayVirtualModelsInsightsGetQueryFromJSON(data []byte) (*AiGatewayVirtualModelsInsightsGetQuery, error) {
	var v AiGatewayVirtualModelsInsightsGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsInsightsGetQueryToJSON serializes a AiGatewayVirtualModelsInsightsGetQuery to JSON.
func MapAiGatewayVirtualModelsInsightsGetQueryToJSON(v *AiGatewayVirtualModelsInsightsGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
