package virtualmodels

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsReplayOutputResultsOriginal - How the request was routed when it was made
type AiGatewayVirtualModelsReplayOutputResultsOriginal struct {
	// RuleId - The rule that picked the model. Null when the default was used.
	RuleId *string `json:"rule_id,omitempty"`
	// Model - The model that served the request
	Model *string `json:"model,omitempty"`
}

// AiGatewayVirtualModelsReplayOutputResultsReplayedTrace represents the ai gateway virtual models replay output results replayed trace type.
type AiGatewayVirtualModelsReplayOutputResultsReplayedTrace struct {
	// RuleId - The rule that was evaluated
	RuleId string `json:"rule_id"`
	// Outcome - The result. signals_unavailable means the rule reads signals but the decision model failed, so it was skipped.
	Outcome string `json:"outcome"`
	// Error - Why evaluating the expression failed, for error outcomes
	Error *string `json:"error,omitempty"`
}

// AiGatewayVirtualModelsReplayOutputResultsReplayedDecisionError - Why the decision failed. Null if it succeeded.
type AiGatewayVirtualModelsReplayOutputResultsReplayedDecisionError struct {
	// Code - Why the decision model failed
	Code string `json:"code"`
	// Message - A description of the failure
	Message string `json:"message"`
}

// AiGatewayVirtualModelsReplayOutputResultsReplayedDecision - The decision model call made to answer the signals. Null when no rule that reads signals was reached.
type AiGatewayVirtualModelsReplayOutputResultsReplayedDecision struct {
	// Model - The decision model that was asked
	Model string `json:"model"`
	// Source - Whose connection ran the decision model: customer for your own provider connection, platform for Metorial's, billed to you. Null when it was not called.
	Source *string `json:"source,omitempty"`
	// Cached - Whether earlier answers were reused instead of calling the decision model
	Cached bool `json:"cached"`
	// LatencyMs - How long the decision took, in milliseconds
	LatencyMs float64 `json:"latency_ms"`
	// InputTokens - The input tokens the decision model read
	InputTokens float64 `json:"input_tokens"`
	// CostUsd - What the decision cost, in USD
	CostUsd float64 `json:"cost_usd"`
	// Signals - The answer to every signal, keyed by signal name. Rules read them as signals.<name>.<field>. Null if the decision failed.
	Signals *map[string]any `json:"signals,omitempty"`
	// Error - Why the decision failed. Null if it succeeded.
	Error *AiGatewayVirtualModelsReplayOutputResultsReplayedDecisionError `json:"error,omitempty"`
}

// AiGatewayVirtualModelsReplayOutputResultsReplayed - How the definition routes the request now. Null if it failed.
type AiGatewayVirtualModelsReplayOutputResultsReplayed struct {
	// RuleId - The rule that picked the model. Null when the default was used.
	RuleId *string `json:"rule_id,omitempty"`
	// Model - The first model that would be tried
	Model *string `json:"model,omitempty"`
	// Trace - Every rule evaluated, in order, up to the one that matched
	Trace []AiGatewayVirtualModelsReplayOutputResultsReplayedTrace `json:"trace"`
	// Decision - The decision model call made to answer the signals. Null when no rule that reads signals was reached.
	Decision *AiGatewayVirtualModelsReplayOutputResultsReplayedDecision `json:"decision,omitempty"`
}

// AiGatewayVirtualModelsReplayOutputResults represents the ai gateway virtual models replay output results type.
type AiGatewayVirtualModelsReplayOutputResults struct {
	// InferenceLogId - The logged request that was replayed
	InferenceLogId string `json:"inference_log_id"`
	// ReceivedAt - When the logged request was received
	ReceivedAt time.Time `json:"received_at"`
	// Original - How the request was routed when it was made
	Original AiGatewayVirtualModelsReplayOutputResultsOriginal `json:"original"`
	// Replayed - How the definition routes the request now. Null if it failed.
	Replayed *AiGatewayVirtualModelsReplayOutputResultsReplayed `json:"replayed,omitempty"`
	// Changed - Whether the rule or the first model differs from the original
	Changed bool `json:"changed"`
	// Error - Why replaying the request failed
	Error *string `json:"error,omitempty"`
}

// AiGatewayVirtualModelsReplayOutputSummary represents the ai gateway virtual models replay output summary type.
type AiGatewayVirtualModelsReplayOutputSummary struct {
	// Replayed - How many requests were replayed
	Replayed float64 `json:"replayed"`
	// Changed - How many would be routed differently
	Changed float64 `json:"changed"`
	// DecisionCostUsd - What the decision model calls of the replay cost, in USD. Logged answers are reused when the signals are unchanged.
	DecisionCostUsd float64 `json:"decision_cost_usd"`
}

// AiGatewayVirtualModelsReplayOutput represents the ai gateway virtual models replay output type.
type AiGatewayVirtualModelsReplayOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Results - One result per replayed request, newest first
	Results []AiGatewayVirtualModelsReplayOutputResults `json:"results"`
	Summary AiGatewayVirtualModelsReplayOutputSummary   `json:"summary"`
}

// MapAiGatewayVirtualModelsReplayOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsReplayOutput.
func MapAiGatewayVirtualModelsReplayOutputFromJSON(data []byte) (*AiGatewayVirtualModelsReplayOutput, error) {
	var v AiGatewayVirtualModelsReplayOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsReplayOutputToJSON serializes a AiGatewayVirtualModelsReplayOutput to JSON.
func MapAiGatewayVirtualModelsReplayOutputToJSON(v *AiGatewayVirtualModelsReplayOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsReplayBodyDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsReplayBodyDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Defaults to 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Defaults to 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Defaults to true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Defaults to true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionDecision - The decision model that answers the signals. Required with signals.
type AiGatewayVirtualModelsReplayBodyDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection to run the decision model on. Omit to use Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Defaults to 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default (the default) skips the rules that read signals, fail rejects the request
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsReplayBodyDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionRulesModels represents the ai gateway virtual models replay body definition rules models type.
type AiGatewayVirtualModelsReplayBodyDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionRulesOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsReplayBodyDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionRules represents the ai gateway virtual models replay body definition rules type.
type AiGatewayVirtualModelsReplayBodyDefinitionRules struct {
	// Id - An identifier for the rule, in lowercase kebab-case, unique in the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request (e.g. request.input_tokens, request.has_images, request.tool_count, request.last_user_text), metadata (the x-metorial-metadata header), token (id, consumer_ref, which is the ID of the consumer the token was created for, and metadata), signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsReplayBodyDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsReplayBodyDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionDefaultModels represents the ai gateway virtual models replay body definition default models type.
type AiGatewayVirtualModelsReplayBodyDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionDefaultOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsReplayBodyDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionDefault - The models to use when no rule matches
type AiGatewayVirtualModelsReplayBodyDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsReplayBodyDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsReplayBodyDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
type AiGatewayVirtualModelsReplayBodyDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Defaults to 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsReplayBodyDefinition - A definition to replay instead of the current one, without saving it
type AiGatewayVirtualModelsReplayBodyDefinition struct {
	// Decision - The decision model that answers the signals. Required with signals.
	Decision *AiGatewayVirtualModelsReplayBodyDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by a name of lowercase letters, digits and underscores. Rules read the answers as signals.<name>. The decision model is only called once a rule that reads signals is reached.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsReplayBodyDefinitionRules `json:"rules"`
	// Default - The models to use when no rule matches
	Default AiGatewayVirtualModelsReplayBodyDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
	Stickiness *AiGatewayVirtualModelsReplayBodyDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsReplayBody represents the ai gateway virtual models replay body type.
type AiGatewayVirtualModelsReplayBody struct {
	// Definition - A definition to replay instead of the current one, without saving it
	Definition *AiGatewayVirtualModelsReplayBodyDefinition `json:"definition,omitempty"`
	// Limit - How many of the most recent requests to replay. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
}

// MapAiGatewayVirtualModelsReplayBodyFromJSON deserializes JSON data into a AiGatewayVirtualModelsReplayBody.
func MapAiGatewayVirtualModelsReplayBodyFromJSON(data []byte) (*AiGatewayVirtualModelsReplayBody, error) {
	var v AiGatewayVirtualModelsReplayBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsReplayBodyToJSON serializes a AiGatewayVirtualModelsReplayBody to JSON.
func MapAiGatewayVirtualModelsReplayBodyToJSON(v *AiGatewayVirtualModelsReplayBody) ([]byte, error) {
	return json.Marshal(v)
}
