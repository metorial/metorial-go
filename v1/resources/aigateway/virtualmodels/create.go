package virtualmodels

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Null uses 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Null uses 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Null means true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Null means true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecision - The decision model that answers the signals. Null when the virtual model has no signals.
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection the decision model runs on. Null uses Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Null uses 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default skips the rules that read signals, fail rejects the request. Null means default.
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesModels represents the ai gateway virtual models create output current version definition rules models type.
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRules represents the ai gateway virtual models create output current version definition rules type.
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRules struct {
	// Id - The rule's identifier, unique within the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request, metadata, token, signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultModels represents the ai gateway virtual models create output current version definition default models type.
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefault - The models used when no rule matches
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Null uses 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersionDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsCreateOutputCurrentVersionDefinition struct {
	// Decision - The decision model that answers the signals. Null when the virtual model has no signals.
	Decision *AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by name. Rules read the answers as signals.<name>. Null when there are none.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionRules `json:"rules"`
	// Default - The models used when no rule matches
	Default AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
	Stickiness *AiGatewayVirtualModelsCreateOutputCurrentVersionDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsCreateOutputCurrentVersion represents the ai gateway virtual models create output current version type.
type AiGatewayVirtualModelsCreateOutputCurrentVersion struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The version's unique identifier
	Id string `json:"id"`
	// Number - The version number, counting up from 1 with every change to the definition
	Number float64 `json:"number"`
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsCreateOutputCurrentVersionDefinition `json:"definition"`
	// TargetModels - The catalog slugs of every model the definition can route to
	TargetModels []string `json:"target_models"`
	// CreatedAt - When the version was created
	CreatedAt time.Time `json:"created_at"`
}

// AiGatewayVirtualModelsCreateOutput represents the ai gateway virtual models create output type.
type AiGatewayVirtualModelsCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The virtual model's unique identifier
	Id string `json:"id"`
	// Slug - The model name to send in AI gateway requests to use this virtual model. Unique within the instance.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for
	Description    *string                                           `json:"description,omitempty"`
	CurrentVersion *AiGatewayVirtualModelsCreateOutputCurrentVersion `json:"current_version,omitempty"`
	// Warnings - Things to check about the definition that did not prevent saving it, e.g. a slug that also names a catalog model. Only returned when the virtual model is created, updated or rolled back.
	Warnings []string `json:"warnings"`
	// CreatedAt - When the virtual model was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the virtual model was last changed
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt - When the virtual model was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// MapAiGatewayVirtualModelsCreateOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsCreateOutput.
func MapAiGatewayVirtualModelsCreateOutputFromJSON(data []byte) (*AiGatewayVirtualModelsCreateOutput, error) {
	var v AiGatewayVirtualModelsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsCreateOutputToJSON serializes a AiGatewayVirtualModelsCreateOutput to JSON.
func MapAiGatewayVirtualModelsCreateOutputToJSON(v *AiGatewayVirtualModelsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsCreateBodyDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsCreateBodyDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Defaults to 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Defaults to 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Defaults to true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Defaults to true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionDecision - The decision model that answers the signals. Required with signals.
type AiGatewayVirtualModelsCreateBodyDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection to run the decision model on. Omit to use Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Defaults to 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default (the default) skips the rules that read signals, fail rejects the request
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsCreateBodyDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionRulesModels represents the ai gateway virtual models create body definition rules models type.
type AiGatewayVirtualModelsCreateBodyDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionRulesOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsCreateBodyDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionRules represents the ai gateway virtual models create body definition rules type.
type AiGatewayVirtualModelsCreateBodyDefinitionRules struct {
	// Id - An identifier for the rule, in lowercase kebab-case, unique in the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request (e.g. request.input_tokens, request.has_images, request.tool_count, request.last_user_text), metadata (the x-metorial-metadata header), token (id, consumer_ref, which is the ID of the consumer the token was created for, and metadata), signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateBodyDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsCreateBodyDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionDefaultModels represents the ai gateway virtual models create body definition default models type.
type AiGatewayVirtualModelsCreateBodyDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>. Cannot be another virtual model.
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Defaults to 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionDefaultOverrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
type AiGatewayVirtualModelsCreateBodyDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionDefault - The models to use when no rule matches
type AiGatewayVirtualModelsCreateBodyDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateBodyDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters to replace when these models are used. Omitted ones keep the request's values.
	Overrides *AiGatewayVirtualModelsCreateBodyDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
type AiGatewayVirtualModelsCreateBodyDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Defaults to 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsCreateBodyDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsCreateBodyDefinition struct {
	// Decision - The decision model that answers the signals. Required with signals.
	Decision *AiGatewayVirtualModelsCreateBodyDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by a name of lowercase letters, digits and underscores. Rules read the answers as signals.<name>. The decision model is only called once a rule that reads signals is reached.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsCreateBodyDefinitionRules `json:"rules"`
	// Default - The models to use when no rule matches
	Default AiGatewayVirtualModelsCreateBodyDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches
	Stickiness *AiGatewayVirtualModelsCreateBodyDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsCreateBody represents the ai gateway virtual models create body type.
type AiGatewayVirtualModelsCreateBody struct {
	// Slug - The model name to send in AI gateway requests to use the virtual model. Lowercase letters, digits, ".", "_" and "-", starting with a letter or digit.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for
	Description *string `json:"description,omitempty"`
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsCreateBodyDefinition `json:"definition"`
}

// MapAiGatewayVirtualModelsCreateBodyFromJSON deserializes JSON data into a AiGatewayVirtualModelsCreateBody.
func MapAiGatewayVirtualModelsCreateBodyFromJSON(data []byte) (*AiGatewayVirtualModelsCreateBody, error) {
	var v AiGatewayVirtualModelsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsCreateBodyToJSON serializes a AiGatewayVirtualModelsCreateBody to JSON.
func MapAiGatewayVirtualModelsCreateBodyToJSON(v *AiGatewayVirtualModelsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
