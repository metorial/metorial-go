package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsArchiveOutputSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsArchiveOutputSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsArchiveOutputSourceUnsyncedSkills represents the skills plugins archive output source unsynced skills type.
type SkillsPluginsArchiveOutputSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsArchiveOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsArchiveOutputSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsArchiveOutputSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsArchiveOutputSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsArchiveOutputSkillsSkillImage - Image shown for the skill.
type SkillsPluginsArchiveOutputSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsArchiveOutputSkillsSkill represents the skills plugins archive output skills skill type.
type SkillsPluginsArchiveOutputSkillsSkill struct {
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
	Image SkillsPluginsArchiveOutputSkillsSkillImage `json:"image"`
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

// SkillsPluginsArchiveOutputSkills represents the skills plugins archive output skills type.
type SkillsPluginsArchiveOutputSkills struct {
	Object               string                                `json:"object"`
	Id                   string                                `json:"id"`
	Identifier           string                                `json:"identifier"`
	Status               string                                `json:"status"`
	ClientName           *string                               `json:"client_name,omitempty"`
	ClientDescription    *string                               `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                       `json:"client_metadata,omitempty"`
	License              *string                               `json:"license,omitempty"`
	Compatibility        *string                               `json:"compatibility,omitempty"`
	SkillConfigurationId *string                               `json:"skill_configuration_id,omitempty"`
	SkillId              string                                `json:"skill_id"`
	Skill                SkillsPluginsArchiveOutputSkillsSkill `json:"skill"`
	CreatedAt            time.Time                             `json:"created_at"`
	UpdatedAt            time.Time                             `json:"updated_at"`
}

// SkillsPluginsArchiveOutput represents the skills plugins archive output type.
type SkillsPluginsArchiveOutput struct {
	Object               string  `json:"object"`
	Id                   string  `json:"id"`
	Status               string  `json:"status"`
	SyncStatus           string  `json:"sync_status"`
	ImageUrl             string  `json:"image_url"`
	Name                 string  `json:"name"`
	Description          *string `json:"description,omitempty"`
	LongDescription      *string `json:"long_description,omitempty"`
	Category             *string `json:"category,omitempty"`
	Slug                 string  `json:"slug"`
	SkillConfigurationId *string `json:"skill_configuration_id,omitempty"`
	// Source - Where the plugin's skills come from.
	Source    SkillsPluginsArchiveOutputSource   `json:"source"`
	Skills    []SkillsPluginsArchiveOutputSkills `json:"skills"`
	CreatedAt time.Time                          `json:"created_at"`
	UpdatedAt time.Time                          `json:"updated_at"`
}

// MapSkillsPluginsArchiveOutputFromJSON deserializes JSON data into a SkillsPluginsArchiveOutput.
func MapSkillsPluginsArchiveOutputFromJSON(data []byte) (*SkillsPluginsArchiveOutput, error) {
	var v SkillsPluginsArchiveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsArchiveOutputToJSON serializes a SkillsPluginsArchiveOutput to JSON.
func MapSkillsPluginsArchiveOutputToJSON(v *SkillsPluginsArchiveOutput) ([]byte, error) {
	return json.Marshal(v)
}
