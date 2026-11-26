package inferencelogs

import (
	"encoding/json"
	"time"
)

// AiGatewayInferenceLogsGetOutputConsumer - The consumer of the token at the time of the request. Null if the token had no consumer or the consumer no longer exists.
type AiGatewayInferenceLogsGetOutputConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayInferenceLogsGetOutputModel - The model that was used. Null if no model was resolved.
type AiGatewayInferenceLogsGetOutputModel struct {
	// Id - The model's unique identifier
	Id *string `json:"id,omitempty"`
	// Slug - The model's slug
	Slug string `json:"slug"`
	// LabSlug - The lab that makes the model
	LabSlug string `json:"lab_slug"`
}

// AiGatewayInferenceLogsGetOutputVirtualModelTrace represents the ai gateway inference logs get output virtual model trace type.
type AiGatewayInferenceLogsGetOutputVirtualModelTrace struct {
	// RuleId - The rule that was evaluated
	RuleId string `json:"rule_id"`
	// Outcome - What evaluating the rule resulted in
	Outcome string `json:"outcome"`
	// Error - Why evaluating the rule failed
	Error *string `json:"error,omitempty"`
}

// AiGatewayInferenceLogsGetOutputVirtualModelSkippedTargets represents the ai gateway inference logs get output virtual model skipped targets type.
type AiGatewayInferenceLogsGetOutputVirtualModelSkippedTargets struct {
	// Model - The model that was skipped
	Model string `json:"model"`
	// Reason - Why it was skipped
	Reason string `json:"reason"`
}

// AiGatewayInferenceLogsGetOutputVirtualModelDecisionError - Why the decision failed
type AiGatewayInferenceLogsGetOutputVirtualModelDecisionError struct {
	// Code - A machine-readable error code
	Code string `json:"code"`
	// Message - A human-readable description
	Message string `json:"message"`
}

// AiGatewayInferenceLogsGetOutputVirtualModelDecision - The model-made routing decision, when a rule asked for one
type AiGatewayInferenceLogsGetOutputVirtualModelDecision struct {
	// Model - The model that made the decision
	Model string `json:"model"`
	// Source - customer if the decision ran over your provider connection, platform for the Metorial provider
	Source *string `json:"source,omitempty"`
	// Cached - Whether a cached decision was reused
	Cached bool `json:"cached"`
	// LatencyMs - How long the decision took, in ms
	LatencyMs float64 `json:"latency_ms"`
	// InputTokens - Input tokens the decision used
	InputTokens float64 `json:"input_tokens"`
	// CostUsd - What the decision cost, in USD
	CostUsd float64 `json:"cost_usd"`
	// Signals - The signals the decision produced
	Signals *map[string]any `json:"signals,omitempty"`
	// Error - Why the decision failed
	Error *AiGatewayInferenceLogsGetOutputVirtualModelDecisionError `json:"error,omitempty"`
}

// AiGatewayInferenceLogsGetOutputVirtualModel - How the virtual model routed the request. Null for direct requests.
type AiGatewayInferenceLogsGetOutputVirtualModel struct {
	// Id - The virtual model's unique identifier
	Id string `json:"id"`
	// Slug - The virtual model's slug
	Slug string `json:"slug"`
	// VersionId - The virtual model version that was used
	VersionId string `json:"version_id"`
	// VersionNumber - The number of that version
	VersionNumber float64 `json:"version_number"`
	// RuleId - The rule that matched. Null when no rule matched.
	RuleId *string `json:"rule_id,omitempty"`
	// Reason - Why the targets were chosen: a rule matched, the version's defaults applied, or an earlier choice was reused
	Reason string `json:"reason"`
	// Trace - Every rule that was evaluated, in order
	Trace []AiGatewayInferenceLogsGetOutputVirtualModelTrace `json:"trace"`
	// Targets - The models the request could be routed to
	Targets []string `json:"targets"`
	// SkippedTargets - Targets that could not be used
	SkippedTargets []AiGatewayInferenceLogsGetOutputVirtualModelSkippedTargets `json:"skipped_targets"`
	// Decision - The model-made routing decision, when a rule asked for one
	Decision *AiGatewayInferenceLogsGetOutputVirtualModelDecision `json:"decision,omitempty"`
}

// AiGatewayInferenceLogsGetOutputRoute - Where the request was served. Null if it never reached a provider.
type AiGatewayInferenceLogsGetOutputRoute struct {
	// Provider - The provider that served the request, by slug. metorial for the Metorial provider.
	Provider string `json:"provider"`
	// ConnectionOwner - customer for one of your provider connections, platform for Metorial's own provider accounts
	ConnectionOwner string `json:"connection_owner"`
	// ConnectionId - Your provider connection. Null for the Metorial provider.
	ConnectionId *string `json:"connection_id,omitempty"`
	// VendorModelId - The provider's own model id. Null for the Metorial provider.
	VendorModelId *string `json:"vendor_model_id,omitempty"`
}

// AiGatewayInferenceLogsGetOutputAttemptsError - Why the request did not succeed
type AiGatewayInferenceLogsGetOutputAttemptsError struct {
	// Code - A machine-readable error code
	Code string `json:"code"`
	// Message - A human-readable description of the error
	Message string `json:"message"`
	// Stage - Where the request stopped: request (invalid or unsupported parameters), limits, routing, upstream (refused or broken off by the provider) or internal
	Stage *string `json:"stage,omitempty"`
	// Param - The request parameter that caused the error
	Param *string `json:"param,omitempty"`
	// HttpStatus - The HTTP status returned for the error
	HttpStatus *float64 `json:"http_status,omitempty"`
	// ProviderCode - The provider's own error code
	ProviderCode *string `json:"provider_code,omitempty"`
}

// AiGatewayInferenceLogsGetOutputAttempts represents the ai gateway inference logs get output attempts type.
type AiGatewayInferenceLogsGetOutputAttempts struct {
	// Provider - The provider that served the request, by slug. metorial for the Metorial provider.
	Provider string `json:"provider"`
	// ConnectionOwner - customer for one of your provider connections, platform for Metorial's own provider accounts
	ConnectionOwner string `json:"connection_owner"`
	// ConnectionId - Your provider connection. Null for the Metorial provider.
	ConnectionId *string `json:"connection_id,omitempty"`
	// VendorModelId - The provider's own model id. Null for the Metorial provider.
	VendorModelId *string `json:"vendor_model_id,omitempty"`
	// StartedAt - When the attempt started
	StartedAt time.Time `json:"started_at"`
	// DurationMs - How long the attempt took, in ms
	DurationMs float64 `json:"duration_ms"`
	// Outcome - How the attempt ended
	Outcome string `json:"outcome"`
	// Error - Why the request did not succeed
	Error *AiGatewayInferenceLogsGetOutputAttemptsError `json:"error,omitempty"`
}

// AiGatewayInferenceLogsGetOutputError - Why the request did not succeed
type AiGatewayInferenceLogsGetOutputError struct {
	// Code - A machine-readable error code
	Code string `json:"code"`
	// Message - A human-readable description of the error
	Message string `json:"message"`
	// Stage - Where the request stopped: request (invalid or unsupported parameters), limits, routing, upstream (refused or broken off by the provider) or internal
	Stage *string `json:"stage,omitempty"`
	// Param - The request parameter that caused the error
	Param *string `json:"param,omitempty"`
	// HttpStatus - The HTTP status returned for the error
	HttpStatus *float64 `json:"http_status,omitempty"`
	// ProviderCode - The provider's own error code
	ProviderCode *string `json:"provider_code,omitempty"`
}

// AiGatewayInferenceLogsGetOutputUsageInput - Input token counts
type AiGatewayInferenceLogsGetOutputUsageInput struct {
	// Total - All input tokens
	Total float64 `json:"total"`
	// NoCache - Input tokens not read from the cache
	NoCache float64 `json:"no_cache"`
	// CacheRead - Input tokens read from the prompt cache
	CacheRead float64 `json:"cache_read"`
	// CacheWrite - Input tokens written to the prompt cache
	CacheWrite float64 `json:"cache_write"`
}

// AiGatewayInferenceLogsGetOutputUsageOutput - Output token counts
type AiGatewayInferenceLogsGetOutputUsageOutput struct {
	// Total - All output tokens
	Total float64 `json:"total"`
	// Text - Output tokens of generated text
	Text float64 `json:"text"`
	// Reasoning - Output tokens used for reasoning
	Reasoning float64 `json:"reasoning"`
}

// AiGatewayInferenceLogsGetOutputUsage - Token usage. Null if the request never reached a provider.
type AiGatewayInferenceLogsGetOutputUsage struct {
	// Input - Input token counts
	Input AiGatewayInferenceLogsGetOutputUsageInput `json:"input"`
	// Output - Output token counts
	Output AiGatewayInferenceLogsGetOutputUsageOutput `json:"output"`
}

// AiGatewayInferenceLogsGetOutputCostComponents - The cost broken down by token type
type AiGatewayInferenceLogsGetOutputCostComponents struct {
	// Input - USD for input tokens
	Input float64 `json:"input"`
	// CacheRead - USD for tokens read from the prompt cache
	CacheRead float64 `json:"cache_read"`
	// CacheWrite - USD for tokens written to the prompt cache
	CacheWrite float64 `json:"cache_write"`
	// Output - USD for output tokens
	Output float64 `json:"output"`
	// Reasoning - USD for reasoning tokens
	Reasoning float64 `json:"reasoning"`
}

// AiGatewayInferenceLogsGetOutputCost - What the request cost. Null if no cost was recorded, e.g. when it never reached a provider.
type AiGatewayInferenceLogsGetOutputCost struct {
	// Currency - The currency of every amount
	Currency string `json:"currency"`
	// Priced - Whether every part of the request could be priced. Parts without a known price count as 0.
	Priced bool `json:"priced"`
	// TotalUsd - What the request cost, in USD
	TotalUsd float64 `json:"total_usd"`
	// Components - The cost broken down by token type
	Components AiGatewayInferenceLogsGetOutputCostComponents `json:"components"`
	// PricingSource - Where the prices came from: the provider's list price, your connection's pricing override, or the Metorial provider's price
	PricingSource string `json:"pricing_source"`
}

// AiGatewayInferenceLogsGetOutputTiming - When the request ran and how long it took
type AiGatewayInferenceLogsGetOutputTiming struct {
	// ReceivedAt - When the gateway received the request
	ReceivedAt time.Time `json:"received_at"`
	// UpstreamStartedAt - When the request was sent to the provider
	UpstreamStartedAt *time.Time `json:"upstream_started_at,omitempty"`
	// FirstTokenAt - When the first token was received
	FirstTokenAt *time.Time `json:"first_token_at,omitempty"`
	// CompletedAt - When the request completed
	CompletedAt time.Time `json:"completed_at"`
	// TotalMs - How long the whole request took, in ms
	TotalMs float64 `json:"total_ms"`
	// UpstreamMs - How long the provider took, in ms
	UpstreamMs *float64 `json:"upstream_ms,omitempty"`
	// TimeToFirstTokenMs - How long until the first token arrived, in ms
	TimeToFirstTokenMs *float64 `json:"time_to_first_token_ms,omitempty"`
}

// AiGatewayInferenceLogsGetOutputRequest - A summary of the request, without its content
type AiGatewayInferenceLogsGetOutputRequest struct {
	// MessageCount - Messages in the request
	MessageCount *float64 `json:"message_count,omitempty"`
	// ToolCount - Tools offered in the request
	ToolCount *float64 `json:"tool_count,omitempty"`
	// MaxOutputTokens - The output token limit set in the request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// UserAgent - The client's user agent
	UserAgent *string `json:"user_agent,omitempty"`
}

// AiGatewayInferenceLogsGetOutputContent - The request and response bodies. Only returned when requested with include_content and still stored.
type AiGatewayInferenceLogsGetOutputContent struct {
	// Request - The request body as sent by the client
	Request map[string]any `json:"request"`
	// Response - The response body. Null if there was none.
	Response *map[string]any `json:"response,omitempty"`
}

// AiGatewayInferenceLogsGetOutput represents the ai gateway inference logs get output type.
type AiGatewayInferenceLogsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The inference log's unique identifier
	Id string `json:"id"`
	// TokenId - The AI gateway token the request was made with
	TokenId string `json:"token_id"`
	// Consumer - The consumer of the token at the time of the request. Null if the token had no consumer or the consumer no longer exists.
	Consumer *AiGatewayInferenceLogsGetOutputConsumer `json:"consumer,omitempty"`
	// Dialect - The API format the request was made in
	Dialect string `json:"dialect"`
	// Operation - What the request did
	Operation string `json:"operation"`
	// RequestId - The request id sent by the client, if any
	RequestId *string `json:"request_id,omitempty"`
	// RequestedModel - The model or virtual model as given in the request
	RequestedModel string `json:"requested_model"`
	// Model - The model that was used. Null if no model was resolved.
	Model *AiGatewayInferenceLogsGetOutputModel `json:"model,omitempty"`
	// VirtualModel - How the virtual model routed the request. Null for direct requests.
	VirtualModel *AiGatewayInferenceLogsGetOutputVirtualModel `json:"virtual_model,omitempty"`
	// Metadata - Metadata sent with the request
	Metadata map[string]string `json:"metadata"`
	// Route - Where the request was served. Null if it never reached a provider.
	Route *AiGatewayInferenceLogsGetOutputRoute `json:"route,omitempty"`
	// Attempts - Every provider the request was sent to, in order. More than one when fallbacks were used.
	Attempts []AiGatewayInferenceLogsGetOutputAttempts `json:"attempts"`
	// Status - How the request ended. rejected requests never reached a provider (invalid parameters, limits or routing); failed ones were refused or broken off by the provider.
	Status string `json:"status"`
	// FinishReason - Why the model stopped generating, as reported by it
	FinishReason *string `json:"finish_reason,omitempty"`
	// Error - Why the request did not succeed
	Error *AiGatewayInferenceLogsGetOutputError `json:"error,omitempty"`
	// Usage - Token usage. Null if the request never reached a provider.
	Usage *AiGatewayInferenceLogsGetOutputUsage `json:"usage,omitempty"`
	// Cost - What the request cost. Null if no cost was recorded, e.g. when it never reached a provider.
	Cost *AiGatewayInferenceLogsGetOutputCost `json:"cost,omitempty"`
	// Timing - When the request ran and how long it took
	Timing AiGatewayInferenceLogsGetOutputTiming `json:"timing"`
	// Request - A summary of the request, without its content
	Request AiGatewayInferenceLogsGetOutputRequest `json:"request"`
	// HasContent - Whether the request and response bodies were stored. They are only stored while content logging is enabled.
	HasContent bool `json:"has_content"`
	// Content - The request and response bodies. Only returned when requested with include_content and still stored.
	Content *AiGatewayInferenceLogsGetOutputContent `json:"content,omitempty"`
	// ExpiresAt - When the log is deleted
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - When the gateway received the request
	CreatedAt time.Time `json:"created_at"`
}

// MapAiGatewayInferenceLogsGetOutputFromJSON deserializes JSON data into a AiGatewayInferenceLogsGetOutput.
func MapAiGatewayInferenceLogsGetOutputFromJSON(data []byte) (*AiGatewayInferenceLogsGetOutput, error) {
	var v AiGatewayInferenceLogsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayInferenceLogsGetOutputToJSON serializes a AiGatewayInferenceLogsGetOutput to JSON.
func MapAiGatewayInferenceLogsGetOutputToJSON(v *AiGatewayInferenceLogsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayInferenceLogsGetQuery represents the ai gateway inference logs get query type.
type AiGatewayInferenceLogsGetQuery struct {
	// IncludeContent - Return the stored request and response bodies. They exist only for requests made while content logging was enabled, and until their retention ends.
	IncludeContent *bool `json:"include_content,omitempty"`
}

// MapAiGatewayInferenceLogsGetQueryFromJSON deserializes JSON data into a AiGatewayInferenceLogsGetQuery.
func MapAiGatewayInferenceLogsGetQueryFromJSON(data []byte) (*AiGatewayInferenceLogsGetQuery, error) {
	var v AiGatewayInferenceLogsGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayInferenceLogsGetQueryToJSON serializes a AiGatewayInferenceLogsGetQuery to JSON.
func MapAiGatewayInferenceLogsGetQueryToJSON(v *AiGatewayInferenceLogsGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
