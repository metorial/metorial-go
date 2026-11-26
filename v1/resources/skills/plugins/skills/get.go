package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsGetOutputSkillImage - Image shown for the skill.
type SkillsPluginsSkillsGetOutputSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsSkillsGetOutputSkill represents the skills plugins skills get output skill type.
type SkillsPluginsSkillsGetOutputSkill struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Whether the skill is available, archived, or deleted.
	Status string `json:"status"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// ImageUrl - Public URL of the skill image.
	ImageUrl string `json:"image_url"`
	// Image - Image shown for the skill.
	Image SkillsPluginsSkillsGetOutputSkillImage `json:"image"`
	// ClientName - Client-facing skill name.
	ClientName string `json:"client_name"`
	// ClientDescription - Client-facing description.
	ClientDescription *string `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill.
	ClientMetadata *map[string]any `json:"client_metadata,omitempty"`
	// License - License terms for the skill.
	License *string `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill.
	Compatibility *string `json:"compatibility,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsPluginsSkillsGetOutput represents the skills plugins skills get output type.
type SkillsPluginsSkillsGetOutput struct {
	Object               string                            `json:"object"`
	Id                   string                            `json:"id"`
	Identifier           string                            `json:"identifier"`
	Status               string                            `json:"status"`
	ClientName           *string                           `json:"client_name,omitempty"`
	ClientDescription    *string                           `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                   `json:"client_metadata,omitempty"`
	License              *string                           `json:"license,omitempty"`
	Compatibility        *string                           `json:"compatibility,omitempty"`
	SkillConfigurationId *string                           `json:"skill_configuration_id,omitempty"`
	SkillId              string                            `json:"skill_id"`
	Skill                SkillsPluginsSkillsGetOutputSkill `json:"skill"`
	CreatedAt            time.Time                         `json:"created_at"`
	UpdatedAt            time.Time                         `json:"updated_at"`
}

// MapSkillsPluginsSkillsGetOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsGetOutput.
func MapSkillsPluginsSkillsGetOutputFromJSON(data []byte) (*SkillsPluginsSkillsGetOutput, error) {
	var v SkillsPluginsSkillsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsGetOutputToJSON serializes a SkillsPluginsSkillsGetOutput to JSON.
func MapSkillsPluginsSkillsGetOutputToJSON(v *SkillsPluginsSkillsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
