package templates

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesDeleteOutputItemsIntegrationConfiguration represents the skills templates delete output items integration configuration type.
type SkillsTemplatesDeleteOutputItemsIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesDeleteOutputItemsIntegration represents the skills templates delete output items integration type.
type SkillsTemplatesDeleteOutputItemsIntegration struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	Description   *string                                                  `json:"description,omitempty"`
	Metadata      *map[string]any                                          `json:"metadata,omitempty"`
	Configuration SkillsTemplatesDeleteOutputItemsIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	ArchivedAt    *time.Time                                               `json:"archived_at,omitempty"`
}

// SkillsTemplatesDeleteOutputItemsProvider represents the skills templates delete output items provider type.
type SkillsTemplatesDeleteOutputItemsProvider struct {
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

// SkillsTemplatesDeleteOutputItems represents the skills templates delete output items type.
type SkillsTemplatesDeleteOutputItems struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesDeleteOutputItemsIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesDeleteOutputItemsProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// SkillsTemplatesDeleteOutput represents the skills templates delete output type.
type SkillsTemplatesDeleteOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Status      string                             `json:"status"`
	Owner       string                             `json:"owner"`
	Slug        string                             `json:"slug"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Metadata    map[string]any                     `json:"metadata"`
	StoreId     string                             `json:"store_id"`
	Items       []SkillsTemplatesDeleteOutputItems `json:"items"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
}

// MapSkillsTemplatesDeleteOutputFromJSON deserializes JSON data into a SkillsTemplatesDeleteOutput.
func MapSkillsTemplatesDeleteOutputFromJSON(data []byte) (*SkillsTemplatesDeleteOutput, error) {
	var v SkillsTemplatesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesDeleteOutputToJSON serializes a SkillsTemplatesDeleteOutput to JSON.
func MapSkillsTemplatesDeleteOutputToJSON(v *SkillsTemplatesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
