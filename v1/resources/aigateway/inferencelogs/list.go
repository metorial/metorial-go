package inferencelogs

import (
	"encoding/json"
	"time"
)

// AiGatewayInferenceLogsListOutputItemsConsumer - The consumer of the token at the time of the request. Null if the token had no consumer or the consumer no longer exists.
type AiGatewayInferenceLogsListOutputItemsConsumer struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The consumer's unique identifier
	Id string `json:"id"`
	// Name - The consumer's name
	Name string `json:"name"`
	// Email - The consumer's email address
	Email string `json:"email"`
}

// AiGatewayInferenceLogsListOutputItemsModel - The model that was used. Null if no model was resolved.
type AiGatewayInferenceLogsListOutputItemsModel struct {
	// Id - The model's unique identifier
	Id *string `json:"id,omitempty"`
	// Slug - The model's slug
	Slug string `json:"slug"`
	// LabSlug - The lab that makes the model
	LabSlug string `json:"lab_slug"`
}

// AiGatewayInferenceLogsListOutputItemsVirtualModelTrace represents the ai gateway inference logs list output items virtual model trace type.
type AiGatewayInferenceLogsListOutputItemsVirtualModelTrace struct {
	// RuleId - The rule that was evaluated
	RuleId string `json:"rule_id"`
	// Outcome - What evaluating the rule resulted in
	Outcome string `json:"outcome"`
	// Error - Why evaluating the rule failed
	Error *string `json:"error,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsVirtualModelSkippedTargets represents the ai gateway inference logs list output items virtual model skipped targets type.
type AiGatewayInferenceLogsListOutputItemsVirtualModelSkippedTargets struct {
	// Model - The model that was skipped
	Model string `json:"model"`
	// Reason - Why it was skipped
	Reason string `json:"reason"`
}

// AiGatewayInferenceLogsListOutputItemsVirtualModelDecisionError - Why the decision failed
type AiGatewayInferenceLogsListOutputItemsVirtualModelDecisionError struct {
	// Code - A machine-readable error code
	Code string `json:"code"`
	// Message - A human-readable description
	Message string `json:"message"`
}

// AiGatewayInferenceLogsListOutputItemsVirtualModelDecision - The model-made routing decision, when a rule asked for one
type AiGatewayInferenceLogsListOutputItemsVirtualModelDecision struct {
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
	Error *AiGatewayInferenceLogsListOutputItemsVirtualModelDecisionError `json:"error,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsVirtualModel - How the virtual model routed the request. Null for direct requests.
type AiGatewayInferenceLogsListOutputItemsVirtualModel struct {
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
	Trace []AiGatewayInferenceLogsListOutputItemsVirtualModelTrace `json:"trace"`
	// Targets - The models the request could be routed to
	Targets []string `json:"targets"`
	// SkippedTargets - Targets that could not be used
	SkippedTargets []AiGatewayInferenceLogsListOutputItemsVirtualModelSkippedTargets `json:"skipped_targets"`
	// Decision - The model-made routing decision, when a rule asked for one
	Decision *AiGatewayInferenceLogsListOutputItemsVirtualModelDecision `json:"decision,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsRoute - Where the request was served. Null if it never reached a provider.
type AiGatewayInferenceLogsListOutputItemsRoute struct {
	// Provider - The provider that served the request, by slug. metorial for the Metorial provider.
	Provider string `json:"provider"`
	// ConnectionOwner - customer for one of your provider connections, platform for Metorial's own provider accounts
	ConnectionOwner string `json:"connection_owner"`
	// ConnectionId - Your provider connection. Null for the Metorial provider.
	ConnectionId *string `json:"connection_id,omitempty"`
	// VendorModelId - The provider's own model id. Null for the Metorial provider.
	VendorModelId *string `json:"vendor_model_id,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsAttemptsError - Why the request did not succeed
type AiGatewayInferenceLogsListOutputItemsAttemptsError struct {
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

// AiGatewayInferenceLogsListOutputItemsAttempts represents the ai gateway inference logs list output items attempts type.
type AiGatewayInferenceLogsListOutputItemsAttempts struct {
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
	Error *AiGatewayInferenceLogsListOutputItemsAttemptsError `json:"error,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsError - Why the request did not succeed
type AiGatewayInferenceLogsListOutputItemsError struct {
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

// AiGatewayInferenceLogsListOutputItemsUsageInput - Input token counts
type AiGatewayInferenceLogsListOutputItemsUsageInput struct {
	// Total - All input tokens
	Total float64 `json:"total"`
	// NoCache - Input tokens not read from the cache
	NoCache float64 `json:"no_cache"`
	// CacheRead - Input tokens read from the prompt cache
	CacheRead float64 `json:"cache_read"`
	// CacheWrite - Input tokens written to the prompt cache
	CacheWrite float64 `json:"cache_write"`
}

// AiGatewayInferenceLogsListOutputItemsUsageOutput - Output token counts
type AiGatewayInferenceLogsListOutputItemsUsageOutput struct {
	// Total - All output tokens
	Total float64 `json:"total"`
	// Text - Output tokens of generated text
	Text float64 `json:"text"`
	// Reasoning - Output tokens used for reasoning
	Reasoning float64 `json:"reasoning"`
}

// AiGatewayInferenceLogsListOutputItemsUsage - Token usage. Null if the request never reached a provider.
type AiGatewayInferenceLogsListOutputItemsUsage struct {
	// Input - Input token counts
	Input AiGatewayInferenceLogsListOutputItemsUsageInput `json:"input"`
	// Output - Output token counts
	Output AiGatewayInferenceLogsListOutputItemsUsageOutput `json:"output"`
}

// AiGatewayInferenceLogsListOutputItemsCostComponents - The cost broken down by token type
type AiGatewayInferenceLogsListOutputItemsCostComponents struct {
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

// AiGatewayInferenceLogsListOutputItemsCost - What the request cost. Null if no cost was recorded, e.g. when it never reached a provider.
type AiGatewayInferenceLogsListOutputItemsCost struct {
	// Currency - The currency of every amount
	Currency string `json:"currency"`
	// Priced - Whether every part of the request could be priced. Parts without a known price count as 0.
	Priced bool `json:"priced"`
	// TotalUsd - What the request cost, in USD
	TotalUsd float64 `json:"total_usd"`
	// Components - The cost broken down by token type
	Components AiGatewayInferenceLogsListOutputItemsCostComponents `json:"components"`
	// PricingSource - Where the prices came from: the provider's list price, your connection's pricing override, or the Metorial provider's price
	PricingSource string `json:"pricing_source"`
}

// AiGatewayInferenceLogsListOutputItemsTiming - When the request ran and how long it took
type AiGatewayInferenceLogsListOutputItemsTiming struct {
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

// AiGatewayInferenceLogsListOutputItemsRequest - A summary of the request, without its content
type AiGatewayInferenceLogsListOutputItemsRequest struct {
	// MessageCount - Messages in the request
	MessageCount *float64 `json:"message_count,omitempty"`
	// ToolCount - Tools offered in the request
	ToolCount *float64 `json:"tool_count,omitempty"`
	// MaxOutputTokens - The output token limit set in the request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// UserAgent - The client's user agent
	UserAgent *string `json:"user_agent,omitempty"`
}

// AiGatewayInferenceLogsListOutputItemsContent - The request and response bodies. Only returned when requested with include_content and still stored.
type AiGatewayInferenceLogsListOutputItemsContent struct {
	// Request - The request body as sent by the client
	Request map[string]any `json:"request"`
	// Response - The response body. Null if there was none.
	Response *map[string]any `json:"response,omitempty"`
}

// AiGatewayInferenceLogsListOutputItems represents the ai gateway inference logs list output items type.
type AiGatewayInferenceLogsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The inference log's unique identifier
	Id string `json:"id"`
	// TokenId - The AI gateway token the request was made with
	TokenId string `json:"token_id"`
	// Consumer - The consumer of the token at the time of the request. Null if the token had no consumer or the consumer no longer exists.
	Consumer *AiGatewayInferenceLogsListOutputItemsConsumer `json:"consumer,omitempty"`
	// Dialect - The API format the request was made in
	Dialect string `json:"dialect"`
	// Operation - What the request did
	Operation string `json:"operation"`
	// RequestId - The request id sent by the client, if any
	RequestId *string `json:"request_id,omitempty"`
	// RequestedModel - The model or virtual model as given in the request
	RequestedModel string `json:"requested_model"`
	// Model - The model that was used. Null if no model was resolved.
	Model *AiGatewayInferenceLogsListOutputItemsModel `json:"model,omitempty"`
	// VirtualModel - How the virtual model routed the request. Null for direct requests.
	VirtualModel *AiGatewayInferenceLogsListOutputItemsVirtualModel `json:"virtual_model,omitempty"`
	// Metadata - Metadata sent with the request
	Metadata map[string]string `json:"metadata"`
	// Route - Where the request was served. Null if it never reached a provider.
	Route *AiGatewayInferenceLogsListOutputItemsRoute `json:"route,omitempty"`
	// Attempts - Every provider the request was sent to, in order. More than one when fallbacks were used.
	Attempts []AiGatewayInferenceLogsListOutputItemsAttempts `json:"attempts"`
	// Status - How the request ended. rejected requests never reached a provider (invalid parameters, limits or routing); failed ones were refused or broken off by the provider.
	Status string `json:"status"`
	// FinishReason - Why the model stopped generating, as reported by it
	FinishReason *string `json:"finish_reason,omitempty"`
	// Error - Why the request did not succeed
	Error *AiGatewayInferenceLogsListOutputItemsError `json:"error,omitempty"`
	// Usage - Token usage. Null if the request never reached a provider.
	Usage *AiGatewayInferenceLogsListOutputItemsUsage `json:"usage,omitempty"`
	// Cost - What the request cost. Null if no cost was recorded, e.g. when it never reached a provider.
	Cost *AiGatewayInferenceLogsListOutputItemsCost `json:"cost,omitempty"`
	// Timing - When the request ran and how long it took
	Timing AiGatewayInferenceLogsListOutputItemsTiming `json:"timing"`
	// Request - A summary of the request, without its content
	Request AiGatewayInferenceLogsListOutputItemsRequest `json:"request"`
	// HasContent - Whether the request and response bodies were stored. They are only stored while content logging is enabled.
	HasContent bool `json:"has_content"`
	// Content - The request and response bodies. Only returned when requested with include_content and still stored.
	Content *AiGatewayInferenceLogsListOutputItemsContent `json:"content,omitempty"`
	// ExpiresAt - When the log is deleted
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt - When the gateway received the request
	CreatedAt time.Time `json:"created_at"`
}

// AiGatewayInferenceLogsListOutputPagination represents the ai gateway inference logs list output pagination type.
type AiGatewayInferenceLogsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayInferenceLogsListOutput represents the ai gateway inference logs list output type.
type AiGatewayInferenceLogsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                     `json:"object"`
	Items      []AiGatewayInferenceLogsListOutputItems    `json:"items"`
	Pagination AiGatewayInferenceLogsListOutputPagination `json:"pagination"`
}

// MapAiGatewayInferenceLogsListOutputFromJSON deserializes JSON data into a AiGatewayInferenceLogsListOutput.
func MapAiGatewayInferenceLogsListOutputFromJSON(data []byte) (*AiGatewayInferenceLogsListOutput, error) {
	var v AiGatewayInferenceLogsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayInferenceLogsListOutputToJSON serializes a AiGatewayInferenceLogsListOutput to JSON.
func MapAiGatewayInferenceLogsListOutputToJSON(v *AiGatewayInferenceLogsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayInferenceLogsListQuery represents the ai gateway inference logs list query type.
type AiGatewayInferenceLogsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// TokenId - Only requests made with this AI gateway token
	TokenId *string `json:"token_id,omitempty"`
	// ConsumerId - Only requests made with tokens of this consumer
	ConsumerId *string `json:"consumer_id,omitempty"`
	// Model - Only requests served by this model, by slug
	Model *string `json:"model,omitempty"`
	// VirtualModelId - Only requests routed by this virtual model
	VirtualModelId *string `json:"virtual_model_id,omitempty"`
	// Status - Only requests that ended this way
	Status *string `json:"status,omitempty"`
}

// MapAiGatewayInferenceLogsListQueryFromJSON deserializes JSON data into a AiGatewayInferenceLogsListQuery.
func MapAiGatewayInferenceLogsListQueryFromJSON(data []byte) (*AiGatewayInferenceLogsListQuery, error) {
	var v AiGatewayInferenceLogsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayInferenceLogsListQueryToJSON serializes a AiGatewayInferenceLogsListQuery to JSON.
func MapAiGatewayInferenceLogsListQueryToJSON(v *AiGatewayInferenceLogsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
