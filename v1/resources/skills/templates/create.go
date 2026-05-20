package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesCreateOutputItemsIntegrationConfiguration represents the skills templates create output items integration configuration type.
type SkillsTemplatesCreateOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesCreateOutputItemsIntegration represents the skills templates create output items integration type.
type SkillsTemplatesCreateOutputItemsIntegration struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	Description   *string                                                  `json:"description,omitempty"`
	Metadata      *map[string]any                                          `json:"metadata,omitempty"`
	Configuration SkillsTemplatesCreateOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	ArchivedAt    *time.Time                                               `json:"archived_at,omitempty"`
}

// SkillsTemplatesCreateOutputItemsProvider represents the skills templates create output items provider type.
type SkillsTemplatesCreateOutputItemsProvider struct {
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

// SkillsTemplatesCreateOutputItems represents the skills templates create output items type.
type SkillsTemplatesCreateOutputItems struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesCreateOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesCreateOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// SkillsTemplatesCreateOutput represents the skills templates create output type.
type SkillsTemplatesCreateOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Status      string                             `json:"status"`
	Owner       string                             `json:"owner"`
	Slug        string                             `json:"slug"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Metadata    map[string]any                     `json:"metadata"`
	StoreId     string                             `json:"store_id"`
	Items       []SkillsTemplatesCreateOutputItems `json:"items"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
}

// MapSkillsTemplatesCreateOutputFromJSON deserializes JSON data into a SkillsTemplatesCreateOutput.
func MapSkillsTemplatesCreateOutputFromJSON(data []byte) (*SkillsTemplatesCreateOutput, error) {
	var v SkillsTemplatesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesCreateOutputToJSON serializes a SkillsTemplatesCreateOutput to JSON.
func MapSkillsTemplatesCreateOutputToJSON(v *SkillsTemplatesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsTemplatesCreateBody represents the skills templates create body type.
type SkillsTemplatesCreateBody struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	FromSkillId *string         `json:"from_skill_Id,omitempty"`
}

// MapSkillsTemplatesCreateBodyFromJSON deserializes JSON data into a SkillsTemplatesCreateBody.
func MapSkillsTemplatesCreateBodyFromJSON(data []byte) (*SkillsTemplatesCreateBody, error) {
	var v SkillsTemplatesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesCreateBodyToJSON serializes a SkillsTemplatesCreateBody to JSON.
func MapSkillsTemplatesCreateBodyToJSON(v *SkillsTemplatesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
