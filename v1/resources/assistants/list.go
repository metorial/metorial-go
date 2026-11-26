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
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AssistantsListOutput represents the assistants list output type.
type AssistantsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                         `json:"object"`
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
