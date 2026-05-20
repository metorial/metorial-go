package assistants

import (
	"encoding/json"
	"time"
)

// AssistantsListOutputItemsDefaultModelProvider represents the assistants list output items default model provider type.
type AssistantsListOutputItemsDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// AssistantsListOutputItemsDefaultModel represents the assistants list output items default model type.
type AssistantsListOutputItemsDefaultModel struct {
	Object        string                                        `json:"object"`
	Id            string                                        `json:"id"`
	Slug          string                                        `json:"slug"`
	Name          string                                        `json:"name"`
	ContextWindow float64                                       `json:"context_window"`
	Provider      AssistantsListOutputItemsDefaultModelProvider `json:"provider"`
}

// AssistantsListOutputItemsAvailableModelsProvider represents the assistants list output items available models provider type.
type AssistantsListOutputItemsAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// AssistantsListOutputItemsAvailableModels represents the assistants list output items available models type.
type AssistantsListOutputItemsAvailableModels struct {
	Object        string                                           `json:"object"`
	Id            string                                           `json:"id"`
	Slug          string                                           `json:"slug"`
	Name          string                                           `json:"name"`
	ContextWindow float64                                          `json:"context_window"`
	Provider      AssistantsListOutputItemsAvailableModelsProvider `json:"provider"`
}

// AssistantsListOutputItems represents the assistants list output items type.
type AssistantsListOutputItems struct {
	Object          string                                     `json:"object"`
	Id              string                                     `json:"id"`
	Slug            string                                     `json:"slug"`
	Name            string                                     `json:"name"`
	OwnerType       string                                     `json:"owner_type"`
	OrganizationId  *string                                    `json:"organization_id,omitempty"`
	DefaultModel    *AssistantsListOutputItemsDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []AssistantsListOutputItemsAvailableModels `json:"available_models"`
	CreatedAt       time.Time                                  `json:"created_at"`
	UpdatedAt       time.Time                                  `json:"updated_at"`
}

// AssistantsListOutputPagination represents the assistants list output pagination type.
type AssistantsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// AssistantsListOutput represents the assistants list output type.
type AssistantsListOutput struct {
	Items      []AssistantsListOutputItems    `json:"items"`
	Pagination AssistantsListOutputPagination `json:"pagination"`
}

// MapAssistantsListOutputFromJSON deserializes JSON data into a AssistantsListOutput.
func MapAssistantsListOutputFromJSON(data []byte) (*AssistantsListOutput, error) {
	var v AssistantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAssistantsListOutputToJSON serializes a AssistantsListOutput to JSON.
func MapAssistantsListOutputToJSON(v *AssistantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AssistantsListQuery represents the assistants list query type.
type AssistantsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapAssistantsListQueryFromJSON deserializes JSON data into a AssistantsListQuery.
func MapAssistantsListQueryFromJSON(data []byte) (*AssistantsListQuery, error) {
	var v AssistantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAssistantsListQueryToJSON serializes a AssistantsListQuery to JSON.
func MapAssistantsListQueryToJSON(v *AssistantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
