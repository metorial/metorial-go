package virtualmodeltemplates

import (
	"encoding/json"
)

// AiGatewayVirtualModelTemplatesListOutputItemsParameters represents the ai gateway virtual model templates list output items parameters type.
type AiGatewayVirtualModelTemplatesListOutputItemsParameters struct {
	// Key - The key to pass the model under when creating from the template
	Key string `json:"key"`
	// Name - The parameter's name
	Name string `json:"name"`
	// Description - What the model is used for
	Description string `json:"description"`
}

// AiGatewayVirtualModelTemplatesListOutputItems represents the ai gateway virtual model templates list output items type.
type AiGatewayVirtualModelTemplatesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Slug - The identifier to create a virtual model from the template with
	Slug string `json:"slug"`
	// Name - The template's name
	Name string `json:"name"`
	// Description - How virtual models made from it route requests
	Description string `json:"description"`
	// Parameters - The models to choose when creating a virtual model from the template
	Parameters []AiGatewayVirtualModelTemplatesListOutputItemsParameters `json:"parameters"`
	// UsesDecisionModel - Whether the template asks a decision model about each request, which is billed per call
	UsesDecisionModel bool `json:"uses_decision_model"`
}

// AiGatewayVirtualModelTemplatesListOutputPagination represents the ai gateway virtual model templates list output pagination type.
type AiGatewayVirtualModelTemplatesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AiGatewayVirtualModelTemplatesListOutput represents the ai gateway virtual model templates list output type.
type AiGatewayVirtualModelTemplatesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                             `json:"object"`
	Items      []AiGatewayVirtualModelTemplatesListOutputItems    `json:"items"`
	Pagination AiGatewayVirtualModelTemplatesListOutputPagination `json:"pagination"`
}

// MapAiGatewayVirtualModelTemplatesListOutputFromJSON deserializes JSON data into a AiGatewayVirtualModelTemplatesListOutput.
func MapAiGatewayVirtualModelTemplatesListOutputFromJSON(data []byte) (*AiGatewayVirtualModelTemplatesListOutput, error) {
	var v AiGatewayVirtualModelTemplatesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAiGatewayVirtualModelTemplatesListOutputToJSON serializes a AiGatewayVirtualModelTemplatesListOutput to JSON.
func MapAiGatewayVirtualModelTemplatesListOutputToJSON(v *AiGatewayVirtualModelTemplatesListOutput) ([]byte, error) {
	return json.Marshal(v)
}
