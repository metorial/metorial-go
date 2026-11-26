package models

import (
	"encoding/json"
	"time"
)

// AiGatewayModelsGetOutputLab - The lab that makes the model
type AiGatewayModelsGetOutputLab struct {
	// Id - The lab's unique identifier
	Id string `json:"id"`
	// Slug - The lab's slug, e.g. anthropic
	Slug string `json:"slug"`
	// Name - The lab's display name
	Name string `json:"name"`
}

// AiGatewayModelsGetOutputModalities - The kinds of content the model works with
type AiGatewayModelsGetOutputModalities struct {
	// Input - What the model accepts: text, image, audio, video, pdf or file
	Input []string `json:"input"`
	// Output - What the model produces: text, image, audio or video
	Output []string `json:"output"`
}

// AiGatewayModelsGetOutputCapabilitiesReasoning - How the model reasons. Null when it does not reason.
type AiGatewayModelsGetOutputCapabilitiesReasoning struct {
	// Toggle - Whether reasoning can be turned on and off
	Toggle bool `json:"toggle"`
	// Effort - The reasoning effort levels the model accepts, e.g. low, medium or high. Null when the effort cannot be set.
	Effort *[]string `json:"effort,omitempty"`
}

// AiGatewayModelsGetOutputCapabilitiesPromptCaching - How prompts are cached. Null when the model has no prompt cache.
type AiGatewayModelsGetOutputCapabilitiesPromptCaching struct {
	// Implicit - Whether repeated prompt prefixes are cached automatically
	Implicit bool `json:"implicit"`
	// Explicit - Whether parts of the prompt can be marked for caching
	Explicit bool `json:"explicit"`
	// Ttls - The cache lifetimes that can be requested, e.g. 5m or 1h
	Ttls []string `json:"ttls"`
}

// AiGatewayModelsGetOutputCapabilities - What the model supports
type AiGatewayModelsGetOutputCapabilities struct {
	// Tools - Whether the model can call tools
	Tools bool `json:"tools"`
	// ToolChoice - The tool_choice values the model accepts, e.g. auto, required or none. Null when unknown.
	ToolChoice *[]string `json:"tool_choice,omitempty"`
	// StructuredOutput - Whether the model can be constrained to a JSON schema
	StructuredOutput bool `json:"structured_output"`
	// Reasoning - How the model reasons. Null when it does not reason.
	Reasoning *AiGatewayModelsGetOutputCapabilitiesReasoning `json:"reasoning,omitempty"`
	// Temperature - Whether the temperature can be set
	Temperature bool `json:"temperature"`
	// PromptCaching - How prompts are cached. Null when the model has no prompt cache.
	PromptCaching *AiGatewayModelsGetOutputCapabilitiesPromptCaching `json:"prompt_caching,omitempty"`
	// Attachments - Whether files can be attached to messages
	Attachments bool `json:"attachments"`
}

// AiGatewayModelsGetOutput represents the ai gateway models get output type.
type AiGatewayModelsGetOutput struct {
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
	Lab AiGatewayModelsGetOutputLab `json:"lab"`
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
	Modalities AiGatewayModelsGetOutputModalities `json:"modalities"`
	// Capabilities - What the model supports
	Capabilities AiGatewayModelsGetOutputCapabilities `json:"capabilities"`
}

// MapAiGatewayModelsGetOutputFromJSON deserializes JSON data into a AiGatewayModelsGetOutput.
func MapAiGatewayModelsGetOutputFromJSON(data []byte) (*AiGatewayModelsGetOutput, error) {
	var v AiGatewayModelsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelsGetOutputToJSON serializes a AiGatewayModelsGetOutput to JSON.
func MapAiGatewayModelsGetOutputToJSON(v *AiGatewayModelsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
