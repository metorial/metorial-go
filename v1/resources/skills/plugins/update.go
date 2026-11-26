package plugins

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsPluginsUpdateOutputSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsUpdateOutputSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsUpdateOutputSourceUnsyncedSkills represents the skills plugins update output source unsynced skills type.
type SkillsPluginsUpdateOutputSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsUpdateOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsUpdateOutputSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsUpdateOutputSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsUpdateOutputSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsUpdateOutputSkillsSkillImage - Image shown for the skill.
type SkillsPluginsUpdateOutputSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsUpdateOutputSkillsSkill represents the skills plugins update output skills skill type.
type SkillsPluginsUpdateOutputSkillsSkill struct {
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
	Image SkillsPluginsUpdateOutputSkillsSkillImage `json:"image"`
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

// SkillsPluginsUpdateOutputSkills represents the skills plugins update output skills type.
type SkillsPluginsUpdateOutputSkills struct {
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
	Skill                SkillsPluginsUpdateOutputSkillsSkill `json:"skill"`
	CreatedAt            time.Time                            `json:"created_at"`
	UpdatedAt            time.Time                            `json:"updated_at"`
}

// SkillsPluginsUpdateOutput represents the skills plugins update output type.
type SkillsPluginsUpdateOutput struct {
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
	Source    SkillsPluginsUpdateOutputSource   `json:"source"`
	Skills    []SkillsPluginsUpdateOutputSkills `json:"skills"`
	CreatedAt time.Time                         `json:"created_at"`
	UpdatedAt time.Time                         `json:"updated_at"`
}

// MapSkillsPluginsUpdateOutputFromJSON deserializes JSON data into a SkillsPluginsUpdateOutput.
func MapSkillsPluginsUpdateOutputFromJSON(data []byte) (*SkillsPluginsUpdateOutput, error) {
	var v SkillsPluginsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsUpdateOutputToJSON serializes a SkillsPluginsUpdateOutput to JSON.
func MapSkillsPluginsUpdateOutputToJSON(v *SkillsPluginsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsUpdateBody represents the skills plugins update body type.
type SkillsPluginsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description. `null` clears it.
	Description *nullable.Nullable[string] `json:"description,omitempty"`
	// LongDescription - Longer description shown on the marketplace. `null` clears it.
	LongDescription *nullable.Nullable[string] `json:"long_description,omitempty"`
	// Category - Category for grouping plugins on the marketplace. `null` clears it.
	Category *nullable.Nullable[string] `json:"category,omitempty"`
	// ImageFileId - Image file for the plugin, by ID. `null` resets to the default image.
	ImageFileId *nullable.Nullable[string] `json:"image_file_id,omitempty"`
	// SkillConfigurationId - Skill configuration for the plugin, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
}

// MapSkillsPluginsUpdateBodyFromJSON deserializes JSON data into a SkillsPluginsUpdateBody.
func MapSkillsPluginsUpdateBodyFromJSON(data []byte) (*SkillsPluginsUpdateBody, error) {
	var v SkillsPluginsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsUpdateBodyToJSON serializes a SkillsPluginsUpdateBody to JSON.
func MapSkillsPluginsUpdateBodyToJSON(v *SkillsPluginsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
