package models

import (
	"encoding/json"
	"time"
)

// AiGatewayModelsListOutputItemsLab - The lab that makes the model
type AiGatewayModelsListOutputItemsLab struct {
	// Id - The lab's unique identifier
	Id string `json:"id"`
	// Slug - The lab's slug, e.g. anthropic
	Slug string `json:"slug"`
	// Name - The lab's display name
	Name string `json:"name"`
}

// AiGatewayModelsListOutputItemsModalities - The kinds of content the model works with
type AiGatewayModelsListOutputItemsModalities struct {
	// Input - What the model accepts: text, image, audio, video, pdf or file
	Input []string `json:"input"`
	// Output - What the model produces: text, image, audio or video
	Output []string `json:"output"`
}

// AiGatewayModelsListOutputItemsCapabilitiesReasoning - How the model reasons. Null when it does not reason.
type AiGatewayModelsListOutputItemsCapabilitiesReasoning struct {
	// Toggle - Whether reasoning can be turned on and off
	Toggle bool `json:"toggle"`
	// Effort - The reasoning effort levels the model accepts, e.g. low, medium or high. Null when the effort cannot be set.
	Effort *[]string `json:"effort,omitempty"`
}

// AiGatewayModelsListOutputItemsCapabilitiesPromptCaching - How prompts are cached. Null when the model has no prompt cache.
type AiGatewayModelsListOutputItemsCapabilitiesPromptCaching struct {
	// Implicit - Whether repeated prompt prefixes are cached automatically
	Implicit bool `json:"implicit"`
	// Explicit - Whether parts of the prompt can be marked for caching
	Explicit bool `json:"explicit"`
	// Ttls - The cache lifetimes that can be requested, e.g. 5m or 1h
	Ttls []string `json:"ttls"`
}

// AiGatewayModelsListOutputItemsCapabilities - What the model supports
type AiGatewayModelsListOutputItemsCapabilities struct {
	// Tools - Whether the model can call tools
	Tools bool `json:"tools"`
	// ToolChoice - The tool_choice values the model accepts, e.g. auto, required or none. Null when unknown.
	ToolChoice *[]string `json:"tool_choice,omitempty"`
	// StructuredOutput - Whether the model can be constrained to a JSON schema
	StructuredOutput bool `json:"structured_output"`
	// Reasoning - How the model reasons. Null when it does not reason.
	Reasoning *AiGatewayModelsListOutputItemsCapabilitiesReasoning `json:"reasoning,omitempty"`
	// Temperature - Whether the temperature can be set
	Temperature bool `json:"temperature"`
	// PromptCaching - How prompts are cached. Null when the model has no prompt cache.
	PromptCaching *AiGatewayModelsListOutputItemsCapabilitiesPromptCaching `json:"prompt_caching,omitempty"`
	// Attachments - Whether files can be attached to messages
	Attachments bool `json:"attachments"`
}

// AiGatewayModelsListOutputItems represents the ai gateway models list output items type.
type AiGatewayModelsListOutputItems struct {
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
	Lab AiGatewayModelsListOutputItemsLab `json:"lab"`
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
	Modalities AiGatewayModelsListOutputItemsModalities `json:"modalities"`
	// Capabilities - What the model supports
	Capabilities AiGatewayModelsListOutputItemsCapabilities `json:"capabilities"`
}

// AiGatewayModelsListOutputPagination represents the ai gateway models list output pagination type.
type AiGatewayModelsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayModelsListOutput represents the ai gateway models list output type.
type AiGatewayModelsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []AiGatewayModelsListOutputItems    `json:"items"`
	Pagination AiGatewayModelsListOutputPagination `json:"pagination"`
}

// MapAiGatewayModelsListOutputFromJSON deserializes JSON data into a AiGatewayModelsListOutput.
func MapAiGatewayModelsListOutputFromJSON(data []byte) (*AiGatewayModelsListOutput, error) {
	var v AiGatewayModelsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelsListOutputToJSON serializes a AiGatewayModelsListOutput to JSON.
func MapAiGatewayModelsListOutputToJSON(v *AiGatewayModelsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayModelsListQuery represents the ai gateway models list query type.
type AiGatewayModelsListQuery struct {
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
	// Lab - Only models made by the lab with this slug
	Lab *string `json:"lab,omitempty"`
	// Kind - Only models of this kind: chat, embedding, reranking, image, video, speech, transcription, realtime, audio, decision or other
	Kind *string `json:"kind,omitempty"`
	// Search - Only models whose slug or name contains this text
	Search *string `json:"search,omitempty"`
}

// MapAiGatewayModelsListQueryFromJSON deserializes JSON data into a AiGatewayModelsListQuery.
func MapAiGatewayModelsListQueryFromJSON(data []byte) (*AiGatewayModelsListQuery, error) {
	var v AiGatewayModelsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelsListQueryToJSON serializes a AiGatewayModelsListQuery to JSON.
func MapAiGatewayModelsListQueryToJSON(v *AiGatewayModelsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
