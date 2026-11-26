package skills

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsPluginsSkillsUpdateOutputSkillImage - Image shown for the skill.
type SkillsPluginsSkillsUpdateOutputSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsSkillsUpdateOutputSkill represents the skills plugins skills update output skill type.
type SkillsPluginsSkillsUpdateOutputSkill struct {
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
	Image SkillsPluginsSkillsUpdateOutputSkillImage `json:"image"`
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

// SkillsPluginsSkillsUpdateOutput represents the skills plugins skills update output type.
type SkillsPluginsSkillsUpdateOutput struct {
	Object               string                               `json:"object"`
	Id                   string                               `json:"id"`
	Identifier           string                               `json:"identifier"`
	Status               string                               `json:"status"`
	ClientName           *string                              `json:"client_name,omitempty"`
	ClientDescription    *string                              `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                      `json:"client_metadata,omitempty"`
	License              *string                              `json:"license,omitempty"`
	Compatibility        *string                              `json:"compatibility,omitempty"`
	SkillConfigurationId *string                              `json:"skill_configuration_id,omitempty"`
	SkillId              string                               `json:"skill_id"`
	Skill                SkillsPluginsSkillsUpdateOutputSkill `json:"skill"`
	CreatedAt            time.Time                            `json:"created_at"`
	UpdatedAt            time.Time                            `json:"updated_at"`
}

// MapSkillsPluginsSkillsUpdateOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsUpdateOutput.
func MapSkillsPluginsSkillsUpdateOutputFromJSON(data []byte) (*SkillsPluginsSkillsUpdateOutput, error) {
	var v SkillsPluginsSkillsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsUpdateOutputToJSON serializes a SkillsPluginsSkillsUpdateOutput to JSON.
func MapSkillsPluginsSkillsUpdateOutputToJSON(v *SkillsPluginsSkillsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSkillsUpdateBody represents the skills plugins skills update body type.
type SkillsPluginsSkillsUpdateBody struct {
	// ClientName - Client-facing skill name. `null` clears it.
	ClientName *nullable.Nullable[string] `json:"client_name,omitempty"`
	// ClientDescription - Client-facing description. `null` clears it.
	ClientDescription *nullable.Nullable[string] `json:"client_description,omitempty"`
	// ClientMetadata - Custom data for the client-facing skill. `null` clears it.
	ClientMetadata *nullable.Nullable[map[string]any] `json:"client_metadata,omitempty"`
	// License - License terms for the skill. `null` clears it.
	License *nullable.Nullable[string] `json:"license,omitempty"`
	// Compatibility - Compatibility notes for the skill. `null` clears it.
	Compatibility *nullable.Nullable[string] `json:"compatibility,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin skill, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsSkillsUpdateBodyFromJSON deserializes JSON data into a SkillsPluginsSkillsUpdateBody.
func MapSkillsPluginsSkillsUpdateBodyFromJSON(data []byte) (*SkillsPluginsSkillsUpdateBody, error) {
	var v SkillsPluginsSkillsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsUpdateBodyToJSON serializes a SkillsPluginsSkillsUpdateBody to JSON.
func MapSkillsPluginsSkillsUpdateBodyToJSON(v *SkillsPluginsSkillsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
