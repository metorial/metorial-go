package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSyncOutputSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsSyncOutputSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsSyncOutputSourceUnsyncedSkills represents the skills plugins sync output source unsynced skills type.
type SkillsPluginsSyncOutputSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsSyncOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsSyncOutputSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsSyncOutputSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsSyncOutputSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsSyncOutputSkillsSkillImage represents the skills plugins sync output skills skill image type.
type SkillsPluginsSyncOutputSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsPluginsSyncOutputSkillsSkill represents the skills plugins sync output skills skill type.
type SkillsPluginsSyncOutputSkillsSkill struct {
	Object            string                                  `json:"object"`
	Id                string                                  `json:"id"`
	Status            string                                  `json:"status"`
	Slug              string                                  `json:"slug"`
	Name              string                                  `json:"name"`
	Description       *string                                 `json:"description,omitempty"`
	ImageUrl          string                                  `json:"image_url"`
	Image             SkillsPluginsSyncOutputSkillsSkillImage `json:"image"`
	ClientName        string                                  `json:"client_name"`
	ClientDescription *string                                 `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                         `json:"client_metadata,omitempty"`
	License           *string                                 `json:"license,omitempty"`
	Compatibility     *string                                 `json:"compatibility,omitempty"`
	Metadata          *map[string]any                         `json:"metadata,omitempty"`
	CreatedAt         time.Time                               `json:"created_at"`
	UpdatedAt         time.Time                               `json:"updated_at"`
}

// SkillsPluginsSyncOutputSkills represents the skills plugins sync output skills type.
type SkillsPluginsSyncOutputSkills struct {
	Object               string                             `json:"object"`
	Id                   string                             `json:"id"`
	Identifier           string                             `json:"identifier"`
	Status               string                             `json:"status"`
	ClientName           *string                            `json:"client_name,omitempty"`
	ClientDescription    *string                            `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                    `json:"client_metadata,omitempty"`
	License              *string                            `json:"license,omitempty"`
	Compatibility        *string                            `json:"compatibility,omitempty"`
	SkillConfigurationId *string                            `json:"skill_configuration_id,omitempty"`
	SkillId              string                             `json:"skill_id"`
	Skill                SkillsPluginsSyncOutputSkillsSkill `json:"skill"`
	CreatedAt            time.Time                          `json:"created_at"`
	UpdatedAt            time.Time                          `json:"updated_at"`
}

// SkillsPluginsSyncOutput represents the skills plugins sync output type.
type SkillsPluginsSyncOutput struct {
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
	Source    SkillsPluginsSyncOutputSource   `json:"source"`
	Skills    []SkillsPluginsSyncOutputSkills `json:"skills"`
	CreatedAt time.Time                       `json:"created_at"`
	UpdatedAt time.Time                       `json:"updated_at"`
}

// MapSkillsPluginsSyncOutputFromJSON deserializes JSON data into a SkillsPluginsSyncOutput.
func MapSkillsPluginsSyncOutputFromJSON(data []byte) (*SkillsPluginsSyncOutput, error) {
	var v SkillsPluginsSyncOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSyncOutputToJSON serializes a SkillsPluginsSyncOutput to JSON.
func MapSkillsPluginsSyncOutputToJSON(v *SkillsPluginsSyncOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSyncBody represents the skills plugins sync body type.
type SkillsPluginsSyncBody struct{}

// MapSkillsPluginsSyncBodyFromJSON deserializes JSON data into a SkillsPluginsSyncBody.
func MapSkillsPluginsSyncBodyFromJSON(data []byte) (*SkillsPluginsSyncBody, error) {
	var v SkillsPluginsSyncBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSyncBodyToJSON serializes a SkillsPluginsSyncBody to JSON.
func MapSkillsPluginsSyncBodyToJSON(v *SkillsPluginsSyncBody) ([]byte, error) {
	return json.Marshal(v)
}
