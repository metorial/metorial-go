package labs

import (
	"encoding/json"
)

// AiGatewayLabsListOutputItems represents the ai gateway labs list output items type.
type AiGatewayLabsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The lab's unique identifier
	Id string `json:"id"`
	// Slug - The lab's slug, e.g. anthropic or openai
	Slug string `json:"slug"`
	// Name - The lab's display name
	Name string `json:"name"`
	// Description - The lab's description
	Description *string `json:"description,omitempty"`
	// Website - The lab's website
	Website *string `json:"website,omitempty"`
	// Logo - The URL of the lab's logo
	Logo *string `json:"logo,omitempty"`
	// Status - retired labs no longer release or serve models
	Status string `json:"status"`
}

// AiGatewayLabsListOutputPagination represents the ai gateway labs list output pagination type.
type AiGatewayLabsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayLabsListOutput represents the ai gateway labs list output type.
type AiGatewayLabsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []AiGatewayLabsListOutputItems    `json:"items"`
	Pagination AiGatewayLabsListOutputPagination `json:"pagination"`
}

// MapAiGatewayLabsListOutputFromJSON deserializes JSON data into a AiGatewayLabsListOutput.
func MapAiGatewayLabsListOutputFromJSON(data []byte) (*AiGatewayLabsListOutput, error) {
	var v AiGatewayLabsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayLabsListOutputToJSON serializes a AiGatewayLabsListOutput to JSON.
func MapAiGatewayLabsListOutputToJSON(v *AiGatewayLabsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AiGatewayLabsListQuery represents the ai gateway labs list query type.
type AiGatewayLabsListQuery struct {
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

// MapAiGatewayLabsListQueryFromJSON deserializes JSON data into a AiGatewayLabsListQuery.
func MapAiGatewayLabsListQueryFromJSON(data []byte) (*AiGatewayLabsListQuery, error) {
	var v AiGatewayLabsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayLabsListQueryToJSON serializes a AiGatewayLabsListQuery to JSON.
func MapAiGatewayLabsListQueryToJSON(v *AiGatewayLabsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
