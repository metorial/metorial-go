package assistants

import (
	"encoding/json"
	"time"
)

// AssistantsGetOutputDefaultModelProvider represents the assistants get output default model provider type.
type AssistantsGetOutputDefaultModelProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// AssistantsGetOutputDefaultModel represents the assistants get output default model type.
type AssistantsGetOutputDefaultModel struct {
	Object        string                                  `json:"object"`
	Id            string                                  `json:"id"`
	Slug          string                                  `json:"slug"`
	Name          string                                  `json:"name"`
	ContextWindow float64                                 `json:"context_window"`
	Provider      AssistantsGetOutputDefaultModelProvider `json:"provider"`
}

// AssistantsGetOutputAvailableModelsProvider represents the assistants get output available models provider type.
type AssistantsGetOutputAvailableModelsProvider struct {
	Object   string `json:"object"`
	Id       string `json:"id"`
	Slug     string `json:"slug"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

// AssistantsGetOutputAvailableModels represents the assistants get output available models type.
type AssistantsGetOutputAvailableModels struct {
	Object        string                                     `json:"object"`
	Id            string                                     `json:"id"`
	Slug          string                                     `json:"slug"`
	Name          string                                     `json:"name"`
	ContextWindow float64                                    `json:"context_window"`
	Provider      AssistantsGetOutputAvailableModelsProvider `json:"provider"`
}

// AssistantsGetOutput represents the assistants get output type.
type AssistantsGetOutput struct {
	Object          string                               `json:"object"`
	Id              string                               `json:"id"`
	Slug            string                               `json:"slug"`
	Name            string                               `json:"name"`
	OwnerType       string                               `json:"owner_type"`
	OrganizationId  *string                              `json:"organization_id,omitempty"`
	DefaultModel    *AssistantsGetOutputDefaultModel     `json:"default_model,omitempty"`
	AvailableModels []AssistantsGetOutputAvailableModels `json:"available_models"`
	CreatedAt       time.Time                            `json:"created_at"`
	UpdatedAt       time.Time                            `json:"updated_at"`
}

// MapAssistantsGetOutputFromJSON deserializes JSON data into a AssistantsGetOutput.
func MapAssistantsGetOutputFromJSON(data []byte) (*AssistantsGetOutput, error) {
	var v AssistantsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAssistantsGetOutputToJSON serializes a AssistantsGetOutput to JSON.
func MapAssistantsGetOutputToJSON(v *AssistantsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
