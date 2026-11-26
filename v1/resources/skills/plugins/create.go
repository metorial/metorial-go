package plugins

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsPluginsCreateOutputSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsCreateOutputSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsCreateOutputSourceUnsyncedSkills represents the skills plugins create output source unsynced skills type.
type SkillsPluginsCreateOutputSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsCreateOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsCreateOutputSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsCreateOutputSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsCreateOutputSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsCreateOutputSkillsSkillImage - Image shown for the skill.
type SkillsPluginsCreateOutputSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsCreateOutputSkillsSkill represents the skills plugins create output skills skill type.
type SkillsPluginsCreateOutputSkillsSkill struct {
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
	Image SkillsPluginsCreateOutputSkillsSkillImage `json:"image"`
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

// SkillsPluginsCreateOutputSkills represents the skills plugins create output skills type.
type SkillsPluginsCreateOutputSkills struct {
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
	Skill                SkillsPluginsCreateOutputSkillsSkill `json:"skill"`
	CreatedAt            time.Time                            `json:"created_at"`
	UpdatedAt            time.Time                            `json:"updated_at"`
}

// SkillsPluginsCreateOutput represents the skills plugins create output type.
type SkillsPluginsCreateOutput struct {
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
	Source    SkillsPluginsCreateOutputSource   `json:"source"`
	Skills    []SkillsPluginsCreateOutputSkills `json:"skills"`
	CreatedAt time.Time                         `json:"created_at"`
	UpdatedAt time.Time                         `json:"updated_at"`
}

// MapSkillsPluginsCreateOutputFromJSON deserializes JSON data into a SkillsPluginsCreateOutput.
func MapSkillsPluginsCreateOutputFromJSON(data []byte) (*SkillsPluginsCreateOutput, error) {
	var v SkillsPluginsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsCreateOutputToJSON serializes a SkillsPluginsCreateOutput to JSON.
func MapSkillsPluginsCreateOutputToJSON(v *SkillsPluginsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsCreateBody represents the skills plugins create body type.
type SkillsPluginsCreateBody struct {
	// Name - The plugin name. Required unless skill_group_id is set, in which case it defaults to the group name.
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
	// SkillMarketplaceId - Marketplace to publish the plugin to, by ID. Required for consumers.
	SkillMarketplaceId *string `json:"skill_marketplace_id,omitempty"`
	// SkillGroupId - Creates the plugin from a skill group. The group's skills are synced into the plugin automatically and cannot be edited on the plugin directly. Name and description are copied from the group once and can be changed afterwards.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
}

// MapSkillsPluginsCreateBodyFromJSON deserializes JSON data into a SkillsPluginsCreateBody.
func MapSkillsPluginsCreateBodyFromJSON(data []byte) (*SkillsPluginsCreateBody, error) {
	var v SkillsPluginsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsCreateBodyToJSON serializes a SkillsPluginsCreateBody to JSON.
func MapSkillsPluginsCreateBodyToJSON(v *SkillsPluginsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
