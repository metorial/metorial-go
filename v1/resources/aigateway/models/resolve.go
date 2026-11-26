package models

import (
	"encoding/json"
	"time"
)

// AiGatewayModelsResolveOutputRequest represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type AiGatewayModelsResolveOutputRequest struct {
	// Type - unified: a plain model name, which routes to your connections first and then to the Metorial provider. metorial: the name was prefixed with metorial:, which only routes to the Metorial provider.
	Type *string `json:"type,omitempty"`
	// Provider - The slug of the pinned provider
	Provider *string `json:"provider,omitempty"`
	// VendorModelId - The provider's id for the model
	VendorModelId *string `json:"vendor_model_id,omitempty"`
	// Name - The model name without any prefix
	Name *string `json:"name,omitempty"`
}

// AiGatewayModelsResolveOutputModelLab - The lab that makes the model
type AiGatewayModelsResolveOutputModelLab struct {
	// Id - The lab's unique identifier
	Id string `json:"id"`
	// Slug - The lab's slug, e.g. anthropic
	Slug string `json:"slug"`
	// Name - The lab's display name
	Name string `json:"name"`
}

// AiGatewayModelsResolveOutputModelModalities - The kinds of content the model works with
type AiGatewayModelsResolveOutputModelModalities struct {
	// Input - What the model accepts: text, image, audio, video, pdf or file
	Input []string `json:"input"`
	// Output - What the model produces: text, image, audio or video
	Output []string `json:"output"`
}

// AiGatewayModelsResolveOutputModelCapabilitiesReasoning - How the model reasons. Null when it does not reason.
type AiGatewayModelsResolveOutputModelCapabilitiesReasoning struct {
	// Toggle - Whether reasoning can be turned on and off
	Toggle bool `json:"toggle"`
	// Effort - The reasoning effort levels the model accepts, e.g. low, medium or high. Null when the effort cannot be set.
	Effort *[]string `json:"effort,omitempty"`
}

// AiGatewayModelsResolveOutputModelCapabilitiesPromptCaching - How prompts are cached. Null when the model has no prompt cache.
type AiGatewayModelsResolveOutputModelCapabilitiesPromptCaching struct {
	// Implicit - Whether repeated prompt prefixes are cached automatically
	Implicit bool `json:"implicit"`
	// Explicit - Whether parts of the prompt can be marked for caching
	Explicit bool `json:"explicit"`
	// Ttls - The cache lifetimes that can be requested, e.g. 5m or 1h
	Ttls []string `json:"ttls"`
}

// AiGatewayModelsResolveOutputModelCapabilities - What the model supports
type AiGatewayModelsResolveOutputModelCapabilities struct {
	// Tools - Whether the model can call tools
	Tools bool `json:"tools"`
	// ToolChoice - The tool_choice values the model accepts, e.g. auto, required or none. Null when unknown.
	ToolChoice *[]string `json:"tool_choice,omitempty"`
	// StructuredOutput - Whether the model can be constrained to a JSON schema
	StructuredOutput bool `json:"structured_output"`
	// Reasoning - How the model reasons. Null when it does not reason.
	Reasoning *AiGatewayModelsResolveOutputModelCapabilitiesReasoning `json:"reasoning,omitempty"`
	// Temperature - Whether the temperature can be set
	Temperature bool `json:"temperature"`
	// PromptCaching - How prompts are cached. Null when the model has no prompt cache.
	PromptCaching *AiGatewayModelsResolveOutputModelCapabilitiesPromptCaching `json:"prompt_caching,omitempty"`
	// Attachments - Whether files can be attached to messages
	Attachments bool `json:"attachments"`
}

// AiGatewayModelsResolveOutputModel represents the ai gateway models resolve output model type.
type AiGatewayModelsResolveOutputModel struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The model's unique identifier
	Id string `json:"id"`
	// Slug - The model's slug, made of the lab slug and the model name, e.g. anthropic/claude-sonnet-4. Use it as the model of AI gateway requests.
	Slug string `json:"slug"`
	// Name - The model's display name
	Name string `json:"name"`
	// Description - The model's description
	Description *string `json:"description,omitempty"`
	// Kind - What the model does: chat, embedding, reranking, image, video, speech, transcription, realtime, audio, decision or other
	Kind string `json:"kind"`
	// Status - Where the model is in its lifecycle. Deprecated models still work but will be retired by the lab.
	Status string `json:"status"`
	// Lab - The lab that makes the model
	Lab AiGatewayModelsResolveOutputModelLab `json:"lab"`
	// FamilySlug - The model family the model belongs to, e.g. claude-sonnet. Null if none.
	FamilySlug *string `json:"family_slug,omitempty"`
	// ReleaseDate - When the model was released
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	// OpenWeights - Whether the model's weights are published
	OpenWeights bool `json:"open_weights"`
	// ContextWindow - The most tokens a request and its response can use together. Null when unknown.
	ContextWindow *float64 `json:"context_window,omitempty"`
	// MaxOutputTokens - The most tokens a single response can have. Null when unknown.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Modalities - The kinds of content the model works with
	Modalities AiGatewayModelsResolveOutputModelModalities `json:"modalities"`
	// Capabilities - What the model supports
	Capabilities AiGatewayModelsResolveOutputModelCapabilities `json:"capabilities"`
}

// AiGatewayModelsResolveOutput represents the ai gateway models resolve output type.
type AiGatewayModelsResolveOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Requested - The model string that was resolved
	Requested string `json:"requested"`
	// Request - How the requested model string was parsed
	Request AiGatewayModelsResolveOutputRequest `json:"request"`
	// Via - How the name was matched to a catalog model. slug: it is the model's slug. alias: it is a name a provider, SDK or Metorial uses for the model, e.g. a dated snapshot id. key: it matches such a name once case, separators and date formats are normalized. offering: it pinned a provider and is one of that provider's ids for the model.
	Via   string                            `json:"via"`
	Model AiGatewayModelsResolveOutputModel `json:"model"`
}

// MapAiGatewayModelsResolveOutputFromJSON deserializes JSON data into a AiGatewayModelsResolveOutput.
func MapAiGatewayModelsResolveOutputFromJSON(data []byte) (*AiGatewayModelsResolveOutput, error) {
	var v AiGatewayModelsResolveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelsResolveOutputToJSON serializes a AiGatewayModelsResolveOutput to JSON.
func MapAiGatewayModelsResolveOutputToJSON(v *AiGatewayModelsResolveOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayModelsResolveBody represents the ai gateway models resolve body type.
type AiGatewayModelsResolveBody struct {
	// Model - The model string to resolve, as it would be sent to the gateway
	Model string `json:"model"`
}

// MapAiGatewayModelsResolveBodyFromJSON deserializes JSON data into a AiGatewayModelsResolveBody.
func MapAiGatewayModelsResolveBodyFromJSON(data []byte) (*AiGatewayModelsResolveBody, error) {
	var v AiGatewayModelsResolveBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelsResolveBodyToJSON serializes a AiGatewayModelsResolveBody to JSON.
func MapAiGatewayModelsResolveBodyToJSON(v *AiGatewayModelsResolveBody) ([]byte, error) {
	return json.Marshal(v)
}
