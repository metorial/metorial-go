package modelofferings

import (
	"encoding/json"
)

// AiGatewayModelOfferingsListOutputItemsProvider - The provider that serves the model
type AiGatewayModelOfferingsListOutputItemsProvider struct {
	// Id - The provider's unique identifier
	Id string `json:"id"`
	// Slug - The provider's slug, e.g. openai
	Slug string `json:"slug"`
}

// AiGatewayModelOfferingsListOutputItemsModel - The catalog model the offering serves. Null when it has not been matched to one.
type AiGatewayModelOfferingsListOutputItemsModel struct {
	// Id - The model's unique identifier
	Id string `json:"id"`
	// Slug - The model's slug, e.g. openai/gpt-4o
	Slug string `json:"slug"`
}

// AiGatewayModelOfferingsListOutputItemsPricingTextCacheWrite represents the ai gateway model offerings list output items pricing text cache write type.
type AiGatewayModelOfferingsListOutputItemsPricingTextCacheWrite struct {
	// Ttl - The cache lifetime this rate applies to, e.g. 5m or 1h
	Ttl *string `json:"ttl,omitempty"`
	// Price - USD per million tokens written to the prompt cache
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingText - Token rates for text. Null when the modality is not priced per token.
type AiGatewayModelOfferingsListOutputItemsPricingText struct {
	// Input - USD per million input tokens
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million input tokens read from the prompt cache
	CacheRead *float64 `json:"cache_read,omitempty"`
	// CacheWrite - Prompt cache write rates, by cache lifetime
	CacheWrite []AiGatewayModelOfferingsListOutputItemsPricingTextCacheWrite `json:"cache_write"`
	// Reasoning - USD per million reasoning tokens. Null when reasoning is billed as output.
	Reasoning *float64 `json:"reasoning,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingImageCacheWrite represents the ai gateway model offerings list output items pricing image cache write type.
type AiGatewayModelOfferingsListOutputItemsPricingImageCacheWrite struct {
	// Ttl - The cache lifetime this rate applies to, e.g. 5m or 1h
	Ttl *string `json:"ttl,omitempty"`
	// Price - USD per million tokens written to the prompt cache
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingImage - Token rates for images. Null when the modality is not priced per token.
type AiGatewayModelOfferingsListOutputItemsPricingImage struct {
	// Input - USD per million input tokens
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million input tokens read from the prompt cache
	CacheRead *float64 `json:"cache_read,omitempty"`
	// CacheWrite - Prompt cache write rates, by cache lifetime
	CacheWrite []AiGatewayModelOfferingsListOutputItemsPricingImageCacheWrite `json:"cache_write"`
	// Reasoning - USD per million reasoning tokens. Null when reasoning is billed as output.
	Reasoning *float64 `json:"reasoning,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingAudioCacheWrite represents the ai gateway model offerings list output items pricing audio cache write type.
type AiGatewayModelOfferingsListOutputItemsPricingAudioCacheWrite struct {
	// Ttl - The cache lifetime this rate applies to, e.g. 5m or 1h
	Ttl *string `json:"ttl,omitempty"`
	// Price - USD per million tokens written to the prompt cache
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingAudio - Token rates for audio. Null when the modality is not priced per token.
type AiGatewayModelOfferingsListOutputItemsPricingAudio struct {
	// Input - USD per million input tokens
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million input tokens read from the prompt cache
	CacheRead *float64 `json:"cache_read,omitempty"`
	// CacheWrite - Prompt cache write rates, by cache lifetime
	CacheWrite []AiGatewayModelOfferingsListOutputItemsPricingAudioCacheWrite `json:"cache_write"`
	// Reasoning - USD per million reasoning tokens. Null when reasoning is billed as output.
	Reasoning *float64 `json:"reasoning,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingVideoCacheWrite represents the ai gateway model offerings list output items pricing video cache write type.
type AiGatewayModelOfferingsListOutputItemsPricingVideoCacheWrite struct {
	// Ttl - The cache lifetime this rate applies to, e.g. 5m or 1h
	Ttl *string `json:"ttl,omitempty"`
	// Price - USD per million tokens written to the prompt cache
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingVideo - Token rates for video. Null when the modality is not priced per token.
type AiGatewayModelOfferingsListOutputItemsPricingVideo struct {
	// Input - USD per million input tokens
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million input tokens read from the prompt cache
	CacheRead *float64 `json:"cache_read,omitempty"`
	// CacheWrite - Prompt cache write rates, by cache lifetime
	CacheWrite []AiGatewayModelOfferingsListOutputItemsPricingVideoCacheWrite `json:"cache_write"`
	// Reasoning - USD per million reasoning tokens. Null when reasoning is billed as output.
	Reasoning *float64 `json:"reasoning,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingContextTiers represents the ai gateway model offerings list output items pricing context tiers type.
type AiGatewayModelOfferingsListOutputItemsPricingContextTiers struct {
	// MinInputTokens - The tier applies to requests with at least this many input tokens
	MinInputTokens float64 `json:"min_input_tokens"`
	// Input - USD per million input tokens in this tier
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens in this tier
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million cached input tokens in this tier
	CacheRead *float64 `json:"cache_read,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingUnits represents the ai gateway model offerings list output items pricing units type.
type AiGatewayModelOfferingsListOutputItemsPricingUnits struct {
	Modality  string `json:"modality"`
	Direction string `json:"direction"`
	// Unit - What is counted, e.g. image or second
	Unit string `json:"unit"`
	// Price - USD per unit
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingConditionsRatesCacheWrite represents the ai gateway model offerings list output items pricing conditions rates cache write type.
type AiGatewayModelOfferingsListOutputItemsPricingConditionsRatesCacheWrite struct {
	// Ttl - The cache lifetime this rate applies to, e.g. 5m or 1h
	Ttl *string `json:"ttl,omitempty"`
	// Price - USD per million tokens written to the prompt cache
	Price float64 `json:"price"`
}

// AiGatewayModelOfferingsListOutputItemsPricingConditionsRates represents the ai gateway model offerings list output items pricing conditions rates type.
type AiGatewayModelOfferingsListOutputItemsPricingConditionsRates struct {
	// Input - USD per million input tokens
	Input *float64 `json:"input,omitempty"`
	// Output - USD per million output tokens
	Output *float64 `json:"output,omitempty"`
	// CacheRead - USD per million input tokens read from the prompt cache
	CacheRead *float64 `json:"cache_read,omitempty"`
	// CacheWrite - Prompt cache write rates, by cache lifetime
	CacheWrite []AiGatewayModelOfferingsListOutputItemsPricingConditionsRatesCacheWrite `json:"cache_write"`
	// Reasoning - USD per million reasoning tokens. Null when reasoning is billed as output.
	Reasoning *float64 `json:"reasoning,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsPricingConditions represents the ai gateway model offerings list output items pricing conditions type.
type AiGatewayModelOfferingsListOutputItemsPricingConditions struct {
	// Label - A short name for the condition
	Label string `json:"label"`
	// Description - When the condition applies
	Description string                                                       `json:"description"`
	Modality    string                                                       `json:"modality"`
	Rates       AiGatewayModelOfferingsListOutputItemsPricingConditionsRates `json:"rates"`
}

// AiGatewayModelOfferingsListOutputItemsPricing - Prices in USD. Token rates are per million tokens.
type AiGatewayModelOfferingsListOutputItemsPricing struct {
	// Currency - The currency of every price
	Currency string `json:"currency"`
	// Text - Token rates for text. Null when the modality is not priced per token.
	Text *AiGatewayModelOfferingsListOutputItemsPricingText `json:"text,omitempty"`
	// Image - Token rates for images. Null when the modality is not priced per token.
	Image *AiGatewayModelOfferingsListOutputItemsPricingImage `json:"image,omitempty"`
	// Audio - Token rates for audio. Null when the modality is not priced per token.
	Audio *AiGatewayModelOfferingsListOutputItemsPricingAudio `json:"audio,omitempty"`
	// Video - Token rates for video. Null when the modality is not priced per token.
	Video *AiGatewayModelOfferingsListOutputItemsPricingVideo `json:"video,omitempty"`
	// ContextTiers - Rates that replace the text rates for long prompts
	ContextTiers []AiGatewayModelOfferingsListOutputItemsPricingContextTiers `json:"context_tiers"`
	// Units - Prices charged per unit instead of per token
	Units []AiGatewayModelOfferingsListOutputItemsPricingUnits `json:"units"`
	// WebSearchPer1K - USD per thousand web searches
	WebSearchPer1K *float64 `json:"web_search_per_1k,omitempty"`
	// Conditions - Rates that apply only under the described conditions
	Conditions []AiGatewayModelOfferingsListOutputItemsPricingConditions `json:"conditions"`
}

// AiGatewayModelOfferingsListOutputItemsModalities - The kinds of content the model works with
type AiGatewayModelOfferingsListOutputItemsModalities struct {
	// Input - What the model accepts: text, image, audio, video, pdf or file
	Input []string `json:"input"`
	// Output - What the model produces: text, image, audio or video
	Output []string `json:"output"`
}

// AiGatewayModelOfferingsListOutputItemsCapabilitiesReasoning - How the model reasons. Null when it does not reason.
type AiGatewayModelOfferingsListOutputItemsCapabilitiesReasoning struct {
	// Toggle - Whether reasoning can be turned on and off
	Toggle bool `json:"toggle"`
	// Effort - The reasoning effort levels the model accepts, e.g. low, medium or high. Null when the effort cannot be set.
	Effort *[]string `json:"effort,omitempty"`
}

// AiGatewayModelOfferingsListOutputItemsCapabilitiesPromptCaching - How prompts are cached. Null when the model has no prompt cache.
type AiGatewayModelOfferingsListOutputItemsCapabilitiesPromptCaching struct {
	// Implicit - Whether repeated prompt prefixes are cached automatically
	Implicit bool `json:"implicit"`
	// Explicit - Whether parts of the prompt can be marked for caching
	Explicit bool `json:"explicit"`
	// Ttls - The cache lifetimes that can be requested, e.g. 5m or 1h
	Ttls []string `json:"ttls"`
}

// AiGatewayModelOfferingsListOutputItemsCapabilities - What the model supports
type AiGatewayModelOfferingsListOutputItemsCapabilities struct {
	// Tools - Whether the model can call tools
	Tools bool `json:"tools"`
	// ToolChoice - The tool_choice values the model accepts, e.g. auto, required or none. Null when unknown.
	ToolChoice *[]string `json:"tool_choice,omitempty"`
	// StructuredOutput - Whether the model can be constrained to a JSON schema
	StructuredOutput bool `json:"structured_output"`
	// Reasoning - How the model reasons. Null when it does not reason.
	Reasoning *AiGatewayModelOfferingsListOutputItemsCapabilitiesReasoning `json:"reasoning,omitempty"`
	// Temperature - Whether the temperature can be set
	Temperature bool `json:"temperature"`
	// PromptCaching - How prompts are cached. Null when the model has no prompt cache.
	PromptCaching *AiGatewayModelOfferingsListOutputItemsCapabilitiesPromptCaching `json:"prompt_caching,omitempty"`
	// Attachments - Whether files can be attached to messages
	Attachments bool `json:"attachments"`
}

// AiGatewayModelOfferingsListOutputItemsVariant - How this offering differs from the standard serving of the model
type AiGatewayModelOfferingsListOutputItemsVariant struct {
	// Serving - A special serving mode, e.g. batch, priority or highspeed. Null for standard serving.
	Serving *string `json:"serving,omitempty"`
	// Reasoning - Whether this variant always or never reasons. Null when reasoning is not fixed.
	Reasoning *string `json:"reasoning,omitempty"`
	// Quantization - The quantization the model is served at, e.g. fp8. Null if unknown.
	Quantization *string `json:"quantization,omitempty"`
	// Free - Whether the provider serves this variant for free
	Free bool `json:"free"`
	// Region - The region the variant is served from. Null if global.
	Region *string `json:"region,omitempty"`
}

// AiGatewayModelOfferingsListOutputItems represents the ai gateway model offerings list output items type.
type AiGatewayModelOfferingsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The offering's unique identifier
	Id string `json:"id"`
	// PinnedModel - The model name that pins requests to this provider and vendor model id, e.g. openai:gpt-4o. Pinned requests only route to your connections for the provider, so this works for supported providers only.
	PinnedModel string `json:"pinned_model"`
	// VendorModelId - The provider's id for the model, e.g. gpt-4o-2024-08-06
	VendorModelId string `json:"vendor_model_id"`
	// DisplayName - The name the provider lists the model under
	DisplayName string `json:"display_name"`
	// Status - Where the offering is in its lifecycle at this provider
	Status string `json:"status"`
	// Provider - The provider that serves the model
	Provider AiGatewayModelOfferingsListOutputItemsProvider `json:"provider"`
	// Model - The catalog model the offering serves. Null when it has not been matched to one.
	Model *AiGatewayModelOfferingsListOutputItemsModel `json:"model,omitempty"`
	// Apis - The APIs the provider serves the model through, e.g. chat-completions or responses
	Apis []string `json:"apis"`
	// Pricing - Prices in USD. Token rates are per million tokens.
	Pricing *AiGatewayModelOfferingsListOutputItemsPricing `json:"pricing,omitempty"`
	// ContextWindow - The most tokens a request and its response can use together. Null when unknown.
	ContextWindow *float64 `json:"context_window,omitempty"`
	// MaxOutputTokens - The most tokens a single response can have. Null when unknown.
	MaxOutputTokens *float64 `json:"max_output_tokens,omitempty"`
	// Modalities - The kinds of content the model works with
	Modalities AiGatewayModelOfferingsListOutputItemsModalities `json:"modalities"`
	// Capabilities - What the model supports
	Capabilities AiGatewayModelOfferingsListOutputItemsCapabilities `json:"capabilities"`
	// Variant - How this offering differs from the standard serving of the model
	Variant AiGatewayModelOfferingsListOutputItemsVariant `json:"variant"`
}

// AiGatewayModelOfferingsListOutputPagination represents the ai gateway model offerings list output pagination type.
type AiGatewayModelOfferingsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayModelOfferingsListOutput represents the ai gateway model offerings list output type.
type AiGatewayModelOfferingsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                      `json:"object"`
	Items      []AiGatewayModelOfferingsListOutputItems    `json:"items"`
	Pagination AiGatewayModelOfferingsListOutputPagination `json:"pagination"`
}

// MapAiGatewayModelOfferingsListOutputFromJSON deserializes JSON data into a AiGatewayModelOfferingsListOutput.
func MapAiGatewayModelOfferingsListOutputFromJSON(data []byte) (*AiGatewayModelOfferingsListOutput, error) {
	var v AiGatewayModelOfferingsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelOfferingsListOutputToJSON serializes a AiGatewayModelOfferingsListOutput to JSON.
func MapAiGatewayModelOfferingsListOutputToJSON(v *AiGatewayModelOfferingsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayModelOfferingsListQuery represents the ai gateway model offerings list query type.
type AiGatewayModelOfferingsListQuery struct {
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
	// ModelId - Only offerings of the model with this ID or slug
	ModelId *string `json:"model_id,omitempty"`
	// Provider - Only offerings of the provider with this slug
	Provider *string `json:"provider,omitempty"`
}

// MapAiGatewayModelOfferingsListQueryFromJSON deserializes JSON data into a AiGatewayModelOfferingsListQuery.
func MapAiGatewayModelOfferingsListQueryFromJSON(data []byte) (*AiGatewayModelOfferingsListQuery, error) {
	var v AiGatewayModelOfferingsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayModelOfferingsListQueryToJSON serializes a AiGatewayModelOfferingsListQuery to JSON.
func MapAiGatewayModelOfferingsListQueryToJSON(v *AiGatewayModelOfferingsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
