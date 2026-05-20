package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesUpdateOutputItemsIntegrationConfiguration represents the skills templates update output items integration configuration type.
type SkillsTemplatesUpdateOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesUpdateOutputItemsIntegration represents the skills templates update output items integration type.
type SkillsTemplatesUpdateOutputItemsIntegration struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	Description   *string                                                  `json:"description,omitempty"`
	Metadata      *map[string]any                                          `json:"metadata,omitempty"`
	Configuration SkillsTemplatesUpdateOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	ArchivedAt    *time.Time                                               `json:"archived_at,omitempty"`
}

// SkillsTemplatesUpdateOutputItemsProvider represents the skills templates update output items provider type.
type SkillsTemplatesUpdateOutputItemsProvider struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique provider identifier
	Id string `json:"id"`
	// Name - Display name of the provider
	Name string `json:"name"`
	// Description - Brief description of the provider
	Description *string `json:"description,omitempty"`
	// Slug - URL-friendly identifier
	Slug string `json:"slug"`
	// CreatedAt - Timestamp when the provider was created
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Timestamp when the provider was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsTemplatesUpdateOutputItems represents the skills templates update output items type.
type SkillsTemplatesUpdateOutputItems struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesUpdateOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesUpdateOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// SkillsTemplatesUpdateOutput represents the skills templates update output type.
type SkillsTemplatesUpdateOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Status      string                             `json:"status"`
	Owner       string                             `json:"owner"`
	Slug        string                             `json:"slug"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Metadata    map[string]any                     `json:"metadata"`
	StoreId     string                             `json:"store_id"`
	Items       []SkillsTemplatesUpdateOutputItems `json:"items"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
}

// MapSkillsTemplatesUpdateOutputFromJSON deserializes JSON data into a SkillsTemplatesUpdateOutput.
func MapSkillsTemplatesUpdateOutputFromJSON(data []byte) (*SkillsTemplatesUpdateOutput, error) {
	var v SkillsTemplatesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesUpdateOutputToJSON serializes a SkillsTemplatesUpdateOutput to JSON.
func MapSkillsTemplatesUpdateOutputToJSON(v *SkillsTemplatesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesUpdateBody represents the skills templates update body type.
type SkillsTemplatesUpdateBody struct {
	Name        *string         `json:"name,omitempty"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
}

// MapSkillsTemplatesUpdateBodyFromJSON deserializes JSON data into a SkillsTemplatesUpdateBody.
func MapSkillsTemplatesUpdateBodyFromJSON(data []byte) (*SkillsTemplatesUpdateBody, error) {
	var v SkillsTemplatesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesUpdateBodyToJSON serializes a SkillsTemplatesUpdateBody to JSON.
func MapSkillsTemplatesUpdateBodyToJSON(v *SkillsTemplatesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
