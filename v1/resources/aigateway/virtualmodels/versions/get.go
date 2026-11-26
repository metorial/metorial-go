package versions

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsVersionsGetOutputDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsVersionsGetOutputDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Null uses 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Null uses 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Null means true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Null means true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionDecision - The decision model that answers the signals. Null when the virtual model has no signals.
type AiGatewayVirtualModelsVersionsGetOutputDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection the decision model runs on. Null uses Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Null uses 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default skips the rules that read signals, fail rejects the request. Null means default.
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsVersionsGetOutputDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesModels represents the ai gateway virtual models versions get output definition rules models type.
type AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionRules represents the ai gateway virtual models versions get output definition rules type.
type AiGatewayVirtualModelsVersionsGetOutputDefinitionRules struct {
	// Id - The rule's identifier, unique within the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request, metadata, token, signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsVersionsGetOutputDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultModels represents the ai gateway virtual models versions get output definition default models type.
type AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionDefault - The models used when no rule matches
type AiGatewayVirtualModelsVersionsGetOutputDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsVersionsGetOutputDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
type AiGatewayVirtualModelsVersionsGetOutputDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Null uses 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutputDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsVersionsGetOutputDefinition struct {
	// Decision - The decision model that answers the signals. Null when the virtual model has no signals.
	Decision *AiGatewayVirtualModelsVersionsGetOutputDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by name. Rules read the answers as signals.<name>. Null when there are none.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsVersionsGetOutputDefinitionRules `json:"rules"`
	// Default - The models used when no rule matches
	Default AiGatewayVirtualModelsVersionsGetOutputDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
	Stickiness *AiGatewayVirtualModelsVersionsGetOutputDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsVersionsGetOutput represents the ai gateway virtual models versions get output type.
type AiGatewayVirtualModelsVersionsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The version's unique identifier
	Id string `json:"id"`
	// Number - The version number, counting up from 1 with every change to the definition
	Number float64 `json:"number"`
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsVersionsGetOutputDefinition `json:"definition"`
	// TargetModels - The catalog slugs of every model the definition can route to
	TargetModels []string `json:"target_models"`
	// CreatedAt - When the version was created
	CreatedAt time.Time `json:"created_at"`
}

// MapAiGatewayVirtualModelsVersionsGetOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsVersionsGetOutput.
func MapAiGatewayVirtualModelsVersionsGetOutputFromJSON(data []byte) (*AiGatewayVirtualModelsVersionsGetOutput, error) {
	var v AiGatewayVirtualModelsVersionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsVersionsGetOutputToJSON serializes a AiGatewayVirtualModelsVersionsGetOutput to JSON.
func MapAiGatewayVirtualModelsVersionsGetOutputToJSON(v *AiGatewayVirtualModelsVersionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
