package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsDeleteOutputIntegrationConfiguration represents the skills templates items delete output integration configuration type.
type SkillsTemplatesItemsDeleteOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesItemsDeleteOutputIntegration represents the skills templates items delete output integration type.
type SkillsTemplatesItemsDeleteOutputIntegration struct {
	Object        string                                                   `json:"object"`
	Id            string                                                   `json:"id"`
	Slug          string                                                   `json:"slug"`
	Name          string                                                   `json:"name"`
	Description   *string                                                  `json:"description,omitempty"`
	Metadata      *map[string]any                                          `json:"metadata,omitempty"`
	Configuration SkillsTemplatesItemsDeleteOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                                `json:"created_at"`
	UpdatedAt     time.Time                                                `json:"updated_at"`
	ArchivedAt    *time.Time                                               `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsDeleteOutputProvider represents the skills templates items delete output provider type.
type SkillsTemplatesItemsDeleteOutputProvider struct {
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

// SkillsTemplatesItemsDeleteOutput represents the skills templates items delete output type.
type SkillsTemplatesItemsDeleteOutput struct {
	Object      string                                       `json:"object"`
	Id          string                                       `json:"id"`
	Type        string                                       `json:"type"`
	Integration *SkillsTemplatesItemsDeleteOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsDeleteOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                    `json:"created_at"`
	UpdatedAt   time.Time                                    `json:"updated_at"`
}

// MapSkillsTemplatesItemsDeleteOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsDeleteOutput.
func MapSkillsTemplatesItemsDeleteOutputFromJSON(data []byte) (*SkillsTemplatesItemsDeleteOutput, error) {
	var v SkillsTemplatesItemsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsDeleteOutputToJSON serializes a SkillsTemplatesItemsDeleteOutput to JSON.
func MapSkillsTemplatesItemsDeleteOutputToJSON(v *SkillsTemplatesItemsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
