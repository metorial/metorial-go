package items

import (
	"encoding/json"
	"time"
)

// SkillsItemsGetOutputIntegrationConfiguration represents the skills items get output integration configuration type.
type SkillsItemsGetOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsItemsGetOutputIntegration represents the skills items get output integration type.
type SkillsItemsGetOutputIntegration struct {
	Object        string                                       `json:"object"`
	Id            string                                       `json:"id"`
	Slug          string                                       `json:"slug"`
	Name          string                                       `json:"name"`
	Description   *string                                      `json:"description,omitempty"`
	Metadata      *map[string]any                              `json:"metadata,omitempty"`
	Configuration SkillsItemsGetOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                    `json:"created_at"`
	UpdatedAt     time.Time                                    `json:"updated_at"`
	ArchivedAt    *time.Time                                   `json:"archived_at,omitempty"`
}

// SkillsItemsGetOutputProvider represents the skills items get output provider type.
type SkillsItemsGetOutputProvider struct {
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

// SkillsItemsGetOutput represents the skills items get output type.
type SkillsItemsGetOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Status      string                           `json:"status"`
	Type        string                           `json:"type"`
	SkillId     string                           `json:"skill_id"`
	Integration *SkillsItemsGetOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsItemsGetOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
}

// MapSkillsItemsGetOutputFromJSON deserializes JSON data into a SkillsItemsGetOutput.
func MapSkillsItemsGetOutputFromJSON(data []byte) (*SkillsItemsGetOutput, error) {
	var v SkillsItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsItemsGetOutputToJSON serializes a SkillsItemsGetOutput to JSON.
func MapSkillsItemsGetOutputToJSON(v *SkillsItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
