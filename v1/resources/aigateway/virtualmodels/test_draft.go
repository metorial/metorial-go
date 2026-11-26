package virtualmodels

import (
	"encoding/json"
)

// AiGatewayVirtualModelsTestDraftOutputRequest - What the gateway read from the request. Rules see these values under request.
type AiGatewayVirtualModelsTestDraftOutputRequest struct {
	// Dialect - The API format the request was read as
	Dialect string `json:"dialect"`
	// InputTokens - The estimated input tokens, as rules read them in request.input_tokens
	InputTokens float64 `json:"input_tokens"`
	// MessageCount - The number of messages
	MessageCount float64 `json:"message_count"`
	// UserTurns - The number of user messages
	UserTurns float64 `json:"user_turns"`
	// IsContinuation - Whether the latest message is a tool result, as in an agent loop
	IsContinuation bool `json:"is_continuation"`
	// HasImages - Whether the request contains images
	HasImages bool `json:"has_images"`
	// HasFiles - Whether the request contains files
	HasFiles bool `json:"has_files"`
	// HasAudio - Whether the request contains audio
	HasAudio bool `json:"has_audio"`
	// ToolCount - The number of tools the request offers
	ToolCount float64 `json:"tool_count"`
}

// AiGatewayVirtualModelsTestDraftOutputRouteTrace represents the ai gateway virtual models test draft output route trace type.
type AiGatewayVirtualModelsTestDraftOutputRouteTrace struct {
	// RuleId - The rule that was evaluated
	RuleId string `json:"rule_id"`
	// Outcome - The result. signals_unavailable means the rule reads signals but the decision model failed, so it was skipped.
	Outcome string `json:"outcome"`
	// Error - Why evaluating the expression failed, for error outcomes
	Error *string `json:"error,omitempty"`
}

// AiGatewayVirtualModelsTestDraftOutputRouteSkippedModels represents the ai gateway virtual models test draft output route skipped models type.
type AiGatewayVirtualModelsTestDraftOutputRouteSkippedModels struct {
	// Model - The model that was left out
	Model string `json:"model"`
	// Reason - Why: the model cannot read images, audio or files, the request exceeds its context window, it cannot call tools, it is not in the catalog, or no provider can serve it
	Reason string `json:"reason"`
}

// AiGatewayVirtualModelsTestDraftOutputRouteDecisionError - Why the decision failed. Null if it succeeded.
type AiGatewayVirtualModelsTestDraftOutputRouteDecisionError struct {
	// Code - Why the decision model failed
	Code string `json:"code"`
	// Message - A description of the failure
	Message string `json:"message"`
}

// AiGatewayVirtualModelsTestDraftOutputRouteDecision - The decision model call made to answer the signals. Null when no rule that reads signals was reached.
type AiGatewayVirtualModelsTestDraftOutputRouteDecision struct {
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
	Error *AiGatewayVirtualModelsTestDraftOutputRouteDecisionError `json:"error,omitempty"`
}

// AiGatewayVirtualModelsTestDraftOutputRoute - How the virtual model routed the request
type AiGatewayVirtualModelsTestDraftOutputRoute struct {
	// VersionId - The version that routed the request. Null for a draft definition.
	VersionId *string `json:"version_id,omitempty"`
	// VersionNumber - The number of that version. Null for a draft definition.
	VersionNumber *float64 `json:"version_number,omitempty"`
	// RuleId - The rule that picked the model. Null when the default was used.
	RuleId *string `json:"rule_id,omitempty"`
	// Reason - Why the models were picked: a rule matched, no rule matched, or the conversation is pinned by stickiness
	Reason string `json:"reason"`
	// Trace - Every rule evaluated, in order, up to the one that matched
	Trace []AiGatewayVirtualModelsTestDraftOutputRouteTrace `json:"trace"`
	// Models - The models picked, in the order they are tried
	Models []string `json:"models"`
	// SkippedModels - Models of the picked rule that cannot handle the request
	SkippedModels []AiGatewayVirtualModelsTestDraftOutputRouteSkippedModels `json:"skipped_models"`
	// Decision - The decision model call made to answer the signals. Null when no rule that reads signals was reached.
	Decision *AiGatewayVirtualModelsTestDraftOutputRouteDecision `json:"decision,omitempty"`
}

// AiGatewayVirtualModelsTestDraftOutputOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsTestDraftOutputOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsTestDraftOutput represents the ai gateway virtual models test draft output type.
type AiGatewayVirtualModelsTestDraftOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Request - What the gateway read from the request. Rules see these values under request.
	Request AiGatewayVirtualModelsTestDraftOutputRequest `json:"request"`
	// Route - How the virtual model routed the request
	Route *AiGatewayVirtualModelsTestDraftOutputRoute `json:"route,omitempty"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsTestDraftOutputOverrides `json:"overrides,omitempty"`
	// Candidates - Where the request would be sent, in the order tried. Nothing is sent to a provider during a test.
	Candidates []any `json:"candidates"`
}

// MapAiGatewayVirtualModelsTestDraftOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsTestDraftOutput.
func MapAiGatewayVirtualModelsTestDraftOutputFromJSON(data []byte) (*AiGatewayVirtualModelsTestDraftOutput, error) {
	var v AiGatewayVirtualModelsTestDraftOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsTestDraftOutputToJSON serializes a AiGatewayVirtualModelsTestDraftOutput to JSON.
func MapAiGatewayVirtualModelsTestDraftOutputToJSON(v *AiGatewayVirtualModelsTestDraftOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsTestDraftBodyDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Defaults to 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Defaults to 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Defaults to true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Defaults to true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionDecision - The decision model that answers the signals. Required with signals.
type AiGatewayVirtualModelsTestDraftBodyDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection to run the decision model on. Omit to use Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Defaults to 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default (the default) skips the rules that read signals, fail rejects the request
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsTestDraftBodyDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionRulesModels represents the ai gateway virtual models test draft body definition rules models type.
type AiGatewayVirtualModelsTestDraftBodyDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionRulesOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsTestDraftBodyDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionRules represents the ai gateway virtual models test draft body definition rules type.
type AiGatewayVirtualModelsTestDraftBodyDefinitionRules struct {
	// Id - An identifier for the rule, in lowercase kebab-case, unique in the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request (e.g. request.input_tokens, request.has_images, request.tool_count, request.last_user_text), metadata (the x-metorial-metadata header), token (id, consumer_ref, which is the ID of the consumer the token was created for, and metadata), signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsTestDraftBodyDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsTestDraftBodyDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultModels represents the ai gateway virtual models test draft body definition default models type.
type AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionDefault - The models to use when no rule matches
type AiGatewayVirtualModelsTestDraftBodyDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsTestDraftBodyDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
type AiGatewayVirtualModelsTestDraftBodyDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Defaults to 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBodyDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsTestDraftBodyDefinition struct {
	// Decision - The decision model that answers the signals. Required with signals.
	Decision *AiGatewayVirtualModelsTestDraftBodyDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by a name of lowercase letters, digits and underscores. Rules read the answers as signals.<name>. The decision model is only called once a rule that reads signals is reached.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsTestDraftBodyDefinitionRules `json:"rules"`
	// Default - The models to use when no rule matches
	Default AiGatewayVirtualModelsTestDraftBodyDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
	Stickiness *AiGatewayVirtualModelsTestDraftBodyDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsTestDraftBody represents the ai gateway virtual models test draft body type.
type AiGatewayVirtualModelsTestDraftBody struct {
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsTestDraftBodyDefinition `json:"definition"`
	// TokenId - The AI gateway token to test as, so rules see its consumer and metadata and its allowed models apply. Omit to test without a token's restrictions.
	TokenId *string `json:"token_id,omitempty"`
	// Dialect - The API format of the request: openai_chat (Chat Completions), anthropic_messages (Messages) or open_responses (Responses)
	Dialect string `json:"dialect"`
	// Request - The request body to route, in the format given by dialect
	Request map[string]any `json:"request"`
	// Metadata - Request metadata as sent in the x-metorial-metadata header. Rules read it as metadata.
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// MapAiGatewayVirtualModelsTestDraftBodyFromJSON deserializes JSON data into a AiGatewayVirtualModelsTestDraftBody.
func MapAiGatewayVirtualModelsTestDraftBodyFromJSON(data []byte) (*AiGatewayVirtualModelsTestDraftBody, error) {
	var v AiGatewayVirtualModelsTestDraftBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsTestDraftBodyToJSON serializes a AiGatewayVirtualModelsTestDraftBody to JSON.
func MapAiGatewayVirtualModelsTestDraftBodyToJSON(v *AiGatewayVirtualModelsTestDraftBody) ([]byte, error) {
	return json.Marshal(v)
}
