package providers

import (
	"encoding/json"
)

// AiGatewayProvidersListOutputItemsFeatures - Features a provider connection for the provider can enable
type AiGatewayProvidersListOutputItemsFeatures struct {
	// Zdr - Whether connections to the provider can enable zero data retention
	Zdr bool `json:"zdr"`
}

// AiGatewayProvidersListOutputItems represents the ai gateway providers list output items type.
type AiGatewayProvidersListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The provider's unique identifier
	Id string `json:"id"`
	// Slug - The provider's slug, e.g. openai or bedrock
	Slug string `json:"slug"`
	// Name - The provider's display name
	Name string `json:"name"`
	// Kind - What kind of provider it is. metorial: Metorial's own provider, billed by Metorial. first_party: a lab's own API. cloud: a cloud platform such as Bedrock or Vertex AI. inference: an inference host. aggregator: a service that resells other providers. subscription: a subscription plan. decision: a classifier used by virtual model rules.
	Kind string `json:"kind"`
	// Website - The provider's website
	Website *string `json:"website,omitempty"`
	// Docs - The URL of the provider's documentation
	Docs *string `json:"docs,omitempty"`
	// Logo - The URL of the provider's logo
	Logo *string `json:"logo,omitempty"`
	// Status - retired providers no longer serve models
	Status string `json:"status"`
	// Supported - Whether the gateway can send requests to the provider, i.e. whether you can create a provider connection for it
	Supported bool `json:"supported"`
	// Features - Features a provider connection for the provider can enable
	Features AiGatewayProvidersListOutputItemsFeatures `json:"features"`
}

// AiGatewayProvidersListOutputPagination represents the ai gateway providers list output pagination type.
type AiGatewayProvidersListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayProvidersListOutput represents the ai gateway providers list output type.
type AiGatewayProvidersListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []AiGatewayProvidersListOutputItems    `json:"items"`
	Pagination AiGatewayProvidersListOutputPagination `json:"pagination"`
}

// MapAiGatewayProvidersListOutputFromJSON deserializes JSON data into a AiGatewayProvidersListOutput.
func MapAiGatewayProvidersListOutputFromJSON(data []byte) (*AiGatewayProvidersListOutput, error) {
	var v AiGatewayProvidersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayProvidersListOutputToJSON serializes a AiGatewayProvidersListOutput to JSON.
func MapAiGatewayProvidersListOutputToJSON(v *AiGatewayProvidersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayProvidersListQuery represents the ai gateway providers list query type.
type AiGatewayProvidersListQuery struct {
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
}

// MapAiGatewayProvidersListQueryFromJSON deserializes JSON data into a AiGatewayProvidersListQuery.
func MapAiGatewayProvidersListQueryFromJSON(data []byte) (*AiGatewayProvidersListQuery, error) {
	var v AiGatewayProvidersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayProvidersListQueryToJSON serializes a AiGatewayProvidersListQuery to JSON.
func MapAiGatewayProvidersListQueryToJSON(v *AiGatewayProvidersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
