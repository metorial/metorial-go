package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsDeleteOutputIntegrationConfiguration represents the skills items delete output integration configuration type.
type SkillsItemsDeleteOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsItemsDeleteOutputIntegration represents the skills items delete output integration type.
type SkillsItemsDeleteOutputIntegration struct {
	Object        string                                          `json:"object"`
	Id            string                                          `json:"id"`
	Slug          string                                          `json:"slug"`
	Name          string                                          `json:"name"`
	Description   *string                                         `json:"description,omitempty"`
	Metadata      *map[string]any                                 `json:"metadata,omitempty"`
	Configuration SkillsItemsDeleteOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                       `json:"created_at"`
	UpdatedAt     time.Time                                       `json:"updated_at"`
	ArchivedAt    *time.Time                                      `json:"archived_at,omitempty"`
}

// SkillsItemsDeleteOutputProvider represents the skills items delete output provider type.
type SkillsItemsDeleteOutputProvider struct {
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

// SkillsItemsDeleteOutput represents the skills items delete output type.
type SkillsItemsDeleteOutput struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Type        string                              `json:"type"`
	SkillId     string                              `json:"skill_id"`
	Integration *SkillsItemsDeleteOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsDeleteOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
}

// MapSkillsItemsDeleteOutputFromJSON deserializes JSON data into a SkillsItemsDeleteOutput.
func MapSkillsItemsDeleteOutputFromJSON(data []byte) (*SkillsItemsDeleteOutput, error) {
	var v SkillsItemsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsDeleteOutputToJSON serializes a SkillsItemsDeleteOutput to JSON.
func MapSkillsItemsDeleteOutputToJSON(v *SkillsItemsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
