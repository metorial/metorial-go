package virtualmodels

import (
	"encoding/json"
	"time"
)

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecisionContext - What the decision model sees of each request
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecisionContext struct {
	// PriorTurns - How many earlier user turns the decision model sees. Null uses 3.
	PriorTurns *float64 `json:"prior_turns,omitempty"`
	// MaxChars - The most characters of conversation the decision model sees. Null uses 8000.
	MaxChars *float64 `json:"max_chars,omitempty"`
	// IncludeSystem - Whether the decision model sees the system prompt. Null means true.
	IncludeSystem *bool `json:"include_system,omitempty"`
	// IncludeToolNames - Whether the decision model sees the names of the tools in the request. Null means true.
	IncludeToolNames *bool `json:"include_tool_names,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecision - The decision model that answers the signals. Null when the virtual model has no signals.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecision struct {
	// Model - The decision model that answers the signals
	Model string `json:"model"`
	// ConnectionId - The ID of your own provider connection the decision model runs on. Null uses Metorial's, billed per call.
	ConnectionId *string `json:"connection_id,omitempty"`
	// TimeoutMs - How long to wait for the decision model, in milliseconds. Null uses 400.
	TimeoutMs *float64 `json:"timeout_ms,omitempty"`
	// OnError - What happens when the decision model fails or times out: default skips the rules that read signals, fail rejects the request. Null means default.
	OnError *string `json:"on_error,omitempty"`
	// Context - What the decision model sees of each request
	Context *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecisionContext `json:"context,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesModels represents the ai gateway virtual models create from template output current version definition rules models type.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRules represents the ai gateway virtual models create from template output current version definition rules type.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRules struct {
	// Id - The rule's identifier, unique within the definition
	Id string `json:"id"`
	// Description - What the rule is for
	Description *string `json:"description,omitempty"`
	// When - A CEL expression that selects the rule when it evaluates to true. It can read request, metadata, token, signals and now.
	When string `json:"when"`
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRulesOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultModels represents the ai gateway virtual models create from template output current version definition default models type.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultModels struct {
	// Model - The model, by catalog slug or name, or pinned to a provider as <provider>:<vendor model id>
	Model string `json:"model"`
	// Weight - The relative share of requests this model receives among the listed models. Null counts as 1.
	Weight *float64 `json:"weight,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultOverrides - Request parameters replaced when these models are used
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultOverrides struct {
	// ReasoningEffort - The reasoning effort to request. Null keeps what the request asks for.
	ReasoningEffort *string `json:"reasoning_effort,omitempty"`
	// MaxOutputTokens - The output token limit to request. Null keeps what the request asks for.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Temperature - The temperature to request. Null keeps what the request asks for.
	Temperature *float64 `json:"temperature,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefault - The models used when no rule matches
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefault struct {
	// Models - The models to use, picked by weight. Models that cannot handle a request (images, tools, context size) are skipped and the rest serve as fallbacks.
	Models []AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultModels `json:"models"`
	// Overrides - Request parameters replaced when these models are used
	Overrides *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefaultOverrides `json:"overrides,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionStickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionStickiness struct {
	// MetadataKey - The request metadata entry that identifies a conversation, e.g. a session id
	MetadataKey string `json:"metadata_key"`
	// TtlSeconds - How long a conversation stays on its model, in seconds. Null uses 3600.
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinition - How the virtual model picks a model for each request
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinition struct {
	// Decision - The decision model that answers the signals. Null when the virtual model has no signals.
	Decision *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDecision `json:"decision,omitempty"`
	// Signals - Questions a decision model answers about each request, keyed by name. Rules read the answers as signals.<name>. Null when there are none.
	Signals *map[string]any `json:"signals,omitempty"`
	// Rules - Rules evaluated in order. The first one that matches picks the models.
	Rules []AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionRules `json:"rules"`
	// Default - The models used when no rule matches
	Default AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionDefault `json:"default"`
	// Stickiness - Keeps every request of a conversation on the model its first request was routed to, which preserves prompt caches. Null when disabled.
	Stickiness *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinitionStickiness `json:"stickiness,omitempty"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersion represents the ai gateway virtual models create from template output current version type.
type AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersion struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The version's unique identifier
	Id string `json:"id"`
	// Number - The version number, counting up from 1 with every change to the definition
	Number float64 `json:"number"`
	// Definition - How the virtual model picks a model for each request
	Definition AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersionDefinition `json:"definition"`
	// TargetModels - The catalog slugs of every model the definition can route to
	TargetModels []string `json:"target_models"`
	// CreatedAt - When the version was created
	CreatedAt time.Time `json:"created_at"`
}

// AiGatewayVirtualModelsCreateFromTemplateOutput represents the ai gateway virtual models create from template output type.
type AiGatewayVirtualModelsCreateFromTemplateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The virtual model's unique identifier
	Id string `json:"id"`
	// Slug - The model name to send in AI gateway requests to use this virtual model. Unique within the instance.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for
	Description    *string                                                       `json:"description,omitempty"`
	CurrentVersion *AiGatewayVirtualModelsCreateFromTemplateOutputCurrentVersion `json:"current_version,omitempty"`
	// Warnings - Things to check about the definition that did not prevent saving it, e.g. a slug that also names a catalog model. Only returned when the virtual model is created, updated or rolled back.
	Warnings []string `json:"warnings"`
	// CreatedAt - When the virtual model was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - When the virtual model was last changed
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt - When the virtual model was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// MapAiGatewayVirtualModelsCreateFromTemplateOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelsCreateFromTemplateOutput.
func MapAiGatewayVirtualModelsCreateFromTemplateOutputFromJSON(data []byte) (*AiGatewayVirtualModelsCreateFromTemplateOutput, error) {
	var v AiGatewayVirtualModelsCreateFromTemplateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsCreateFromTemplateOutputToJSON serializes a AiGatewayVirtualModelsCreateFromTemplateOutput to JSON.
func MapAiGatewayVirtualModelsCreateFromTemplateOutputToJSON(v *AiGatewayVirtualModelsCreateFromTemplateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayVirtualModelsCreateFromTemplateBody represents the ai gateway virtual models create from template body type.
type AiGatewayVirtualModelsCreateFromTemplateBody struct {
	// Template - The slug of the template to use
	Template string `json:"template"`
	// Slug - The model name to send in AI gateway requests to use the virtual model. Lowercase letters, digits, ".", "_" and "-", starting with a letter or digit.
	Slug string `json:"slug"`
	// Name - The virtual model's name
	Name string `json:"name"`
	// Description - What the virtual model is for. Defaults to the template's description.
	Description *string `json:"description,omitempty"`
	// Parameters - The model for each of the template's parameters, keyed by parameter key
	Parameters map[string]string `json:"parameters"`
	// DecisionConnectionId - The ID of your own provider connection to run the template's decision model on. Omit to use Metorial's, billed per call.
	DecisionConnectionId *string `json:"decision_connection_id,omitempty"`
}

// MapAiGatewayVirtualModelsCreateFromTemplateBodyFromJSON deserializes JSON data into a AiGatewayVirtualModelsCreateFromTemplateBody.
func MapAiGatewayVirtualModelsCreateFromTemplateBodyFromJSON(data []byte) (*AiGatewayVirtualModelsCreateFromTemplateBody, error) {
	var v AiGatewayVirtualModelsCreateFromTemplateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelsCreateFromTemplateBodyToJSON serializes a AiGatewayVirtualModelsCreateFromTemplateBody to JSON.
func MapAiGatewayVirtualModelsCreateFromTemplateBodyToJSON(v *AiGatewayVirtualModelsCreateFromTemplateBody) ([]byte, error) {
	return json.Marshal(v)
}
