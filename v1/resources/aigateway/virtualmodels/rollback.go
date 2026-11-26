package virtualmodels

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Null uses 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Null uses 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Null means true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Null means true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecision - The decision model that answers the signals. Null when the virtual model has no signals.
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection the decision model runs on. Null uses Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Null uses 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default skips the rules that read signals, fail rejects the request. Null means default.
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesModels represents the ai gateway virtual models rollback output current version definition rules models type.
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRules represents the ai gateway virtual models rollback output current version definition rules type.
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRules struct {
	// Id - The rule's identifier, unique within the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request, metadata, token, signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultModels represents the ai gateway virtual models rollback output current version definition default models type.
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefault - The models used when no rule matches
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Null uses 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinition struct {
	// Decision - The decision model that answers the signals. Null when the virtual model has no signals.
	Decision *AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by name. Rules read the answers as signals.<name>. Null when there are none.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionRules `json:"rules"`
	// Default - The models used when no rule matches
	Default AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
	Stickiness *AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsRollbackOutputCurrentVersion represents the ai gateway virtual models rollback output current version type.
type AiGatewayVirtualModelsRollbackOutputCurrentVersion struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The version's unique identifier
	Id string `json:"id"`
	// Number - The version number, counting up from 1 with every change to the definition
	Number float64 `json:"number"`
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsRollbackOutputCurrentVersionDefinition `json:"definition"`
	// TargetModels - The catalog slugs of every model the definition can route to
	TargetModels []string `json:"target_models"`
	// CreatedAt - When the version was created
	CreatedAt time.Time `json:"created_at"`
}

// AiGatewayVirtualModelsRollbackOutput represents the ai gateway virtual models rollback output type.
type AiGatewayVirtualModelsRollbackOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The virtual model's unique identifier
	Id string `json:"id"`
	// Slug - The model name to send in AI gateway requests to use this virtual model. Unique within the instance.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for
	Description    *string                                             `json:"description,omitempty"`
	CurrentVersion *AiGatewayVirtualModelsRollbackOutputCurrentVersion `json:"current_version,omitempty"`
	// Warnings - Things to check about the definition that did not prevent saving it, e.g. a slug that also names a catalog model. Only returned when the virtual model is created, updated or rolled back.
	Warnings []string `json:"warnings"`
	// CreatedAt - When the virtual model was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the virtual model was last changed
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt - When the virtual model was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// MapAiGatewayVirtualModelsRollbackOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsRollbackOutput.
func MapAiGatewayVirtualModelsRollbackOutputFromJSON(data []byte) (*AiGatewayVirtualModelsRollbackOutput, error) {
	var v AiGatewayVirtualModelsRollbackOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsRollbackOutputToJSON serializes a AiGatewayVirtualModelsRollbackOutput to JSON.
func MapAiGatewayVirtualModelsRollbackOutputToJSON(v *AiGatewayVirtualModelsRollbackOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsRollbackBody represents the ai gateway virtual models rollback body type.
type AiGatewayVirtualModelsRollbackBody struct {
	// VersionId - The ID of the version to roll back to
	VersionId string `json:"version_id"`
}

// MapAiGatewayVirtualModelsRollbackBodyFromJSON deserializes JSON data into a AiGatewayVirtualModelsRollbackBody.
func MapAiGatewayVirtualModelsRollbackBodyFromJSON(data []byte) (*AiGatewayVirtualModelsRollbackBody, error) {
	var v AiGatewayVirtualModelsRollbackBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsRollbackBodyToJSON serializes a AiGatewayVirtualModelsRollbackBody to JSON.
func MapAiGatewayVirtualModelsRollbackBodyToJSON(v *AiGatewayVirtualModelsRollbackBody) ([]byte, error) {
	return json.Marshal(v)
}
