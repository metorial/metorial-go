package items

import (
	"encoding/json"
	"time"
)

// SkillsTemplatesItemsGetOutputIntegrationConfiguration represents the skills templates items get output integration configuration type.
type SkillsTemplatesItemsGetOutputIntegrationConfiguration struct {
	CanAttachCustomToolFilters    bool `json:"can_attach_custom_tool_filters"`
	CanAttachCustomProviderConfig bool `json:"can_attach_custom_provider_config"`
	CanOverrideToolFilters        bool `json:"can_override_tool_filters"`
}

// SkillsTemplatesItemsGetOutputIntegration represents the skills templates items get output integration type.
type SkillsTemplatesItemsGetOutputIntegration struct {
	Object        string                                                `json:"object"`
	Id            string                                                `json:"id"`
	Slug          string                                                `json:"slug"`
	Name          string                                                `json:"name"`
	Description   *string                                               `json:"description,omitempty"`
	Metadata      *map[string]any                                       `json:"metadata,omitempty"`
	Configuration SkillsTemplatesItemsGetOutputIntegrationConfiguration `json:"configuration"`
	CreatedAt     time.Time                                             `json:"created_at"`
	UpdatedAt     time.Time                                             `json:"updated_at"`
	ArchivedAt    *time.Time                                            `json:"archived_at,omitempty"`
}

// SkillsTemplatesItemsGetOutputProvider represents the skills templates items get output provider type.
type SkillsTemplatesItemsGetOutputProvider struct {
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

// SkillsTemplatesItemsGetOutput represents the skills templates items get output type.
type SkillsTemplatesItemsGetOutput struct {
	Object      string                                    `json:"object"`
	Id          string                                    `json:"id"`
	Type        string                                    `json:"type"`
	Integration *SkillsTemplatesItemsGetOutputIntegration `json:"integration,omitempty"`
	Provider    *SkillsTemplatesItemsGetOutputProvider    `json:"provider,omitempty"`
	CreatedAt   time.Time                                 `json:"created_at"`
	UpdatedAt   time.Time                                 `json:"updated_at"`
}

// MapSkillsTemplatesItemsGetOutputFromJSON deserializes JSON data into a SkillsTemplatesItemsGetOutput.
func MapSkillsTemplatesItemsGetOutputFromJSON(data []byte) (*SkillsTemplatesItemsGetOutput, error) {
	var v SkillsTemplatesItemsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsTemplatesItemsGetOutputToJSON serializes a SkillsTemplatesItemsGetOutput to JSON.
func MapSkillsTemplatesItemsGetOutputToJSON(v *SkillsTemplatesItemsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
