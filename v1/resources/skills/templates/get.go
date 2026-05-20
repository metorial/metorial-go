package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesGetOutputItemsIntegrationConfiguration represents the skills templates get output items integration configuration type.
type SkillsTemplatesGetOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesGetOutputItemsIntegration represents the skills templates get output items integration type.
type SkillsTemplatesGetOutputItemsIntegration struct {
	Object        string                                                `json:"object"`
	Id            string                                                `json:"id"`
	Slug          string                                                `json:"slug"`
	Name          string                                                `json:"name"`
	Description   *string                                               `json:"description,omitempty"`
	Metadata      *map[string]any                                       `json:"metadata,omitempty"`
	Configuration SkillsTemplatesGetOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                             `json:"created_at"`
	UpdatedAt     time.Time                                             `json:"updated_at"`
	ArchivedAt    *time.Time                                            `json:"archived_at,omitempty"`
}

// SkillsTemplatesGetOutputItemsProvider represents the skills templates get output items provider type.
type SkillsTemplatesGetOutputItemsProvider struct {
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

// SkillsTemplatesGetOutputItems represents the skills templates get output items type.
type SkillsTemplatesGetOutputItems struct {
	Object      string                                    `json:"object"`
	Id          string                                    `json:"id"`
	Type        string                                    `json:"type"`
	Integration *SkillsTemplatesGetOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesGetOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                 `json:"created_at"`
	UpdatedAt   time.Time                                 `json:"updated_at"`
}

// SkillsTemplatesGetOutput represents the skills templates get output type.
type SkillsTemplatesGetOutput struct {
	Object      string                          `json:"object"`
	Id          string                          `json:"id"`
	Status      string                          `json:"status"`
	Owner       string                          `json:"owner"`
	Slug        string                          `json:"slug"`
	Name        string                          `json:"name"`
	Description *string                         `json:"description,omitempty"`
	Metadata    map[string]any                  `json:"metadata"`
	StoreId     string                          `json:"store_id"`
	Items       []SkillsTemplatesGetOutputItems `json:"items"`
	CreatedAt   time.Time                       `json:"created_at"`
	UpdatedAt   time.Time                       `json:"updated_at"`
}

// MapSkillsTemplatesGetOutputFromJSON deserializes JSON data into a SkillsTemplatesGetOutput.
func MapSkillsTemplatesGetOutputFromJSON(data []byte) (*SkillsTemplatesGetOutput, error) {
	var v SkillsTemplatesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesGetOutputToJSON serializes a SkillsTemplatesGetOutput to JSON.
func MapSkillsTemplatesGetOutputToJSON(v *SkillsTemplatesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
