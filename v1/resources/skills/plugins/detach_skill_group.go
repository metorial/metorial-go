package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsDetachSkillGroupOutputSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsDetachSkillGroupOutputSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsDetachSkillGroupOutputSourceUnsyncedSkills represents the skills plugins detach skill group output source unsynced skills type.
type SkillsPluginsDetachSkillGroupOutputSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsDetachSkillGroupOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsDetachSkillGroupOutputSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsDetachSkillGroupOutputSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsDetachSkillGroupOutputSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsDetachSkillGroupOutputSkillsSkillImage represents the skills plugins detach skill group output skills skill image type.
type SkillsPluginsDetachSkillGroupOutputSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsPluginsDetachSkillGroupOutputSkillsSkill represents the skills plugins detach skill group output skills skill type.
type SkillsPluginsDetachSkillGroupOutputSkillsSkill struct {
	Object            string                                              `json:"object"`
	Id                string                                              `json:"id"`
	Status            string                                              `json:"status"`
	Slug              string                                              `json:"slug"`
	Name              string                                              `json:"name"`
	Description       *string                                             `json:"description,omitempty"`
	ImageUrl          string                                              `json:"image_url"`
	Image             SkillsPluginsDetachSkillGroupOutputSkillsSkillImage `json:"image"`
	ClientName        string                                              `json:"client_name"`
	ClientDescription *string                                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                     `json:"client_metadata,omitempty"`
	License           *string                                             `json:"license,omitempty"`
	Compatibility     *string                                             `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                     `json:"metadata,omitempty"`
	CreatedAt         time.Time                                           `json:"created_at"`
	UpdatedAt         time.Time                                           `json:"updated_at"`
}

// SkillsPluginsDetachSkillGroupOutputSkills represents the skills plugins detach skill group output skills type.
type SkillsPluginsDetachSkillGroupOutputSkills struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Identifier           string                                         `json:"identifier"`
	Status               string                                         `json:"status"`
	ClientName           *string                                        `json:"client_name,omitempty"`
	ClientDescription    *string                                        `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                `json:"client_metadata,omitempty"`
	License              *string                                        `json:"license,omitempty"`
	Compatibility        *string                                        `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                        `json:"skill_configuration_id,omitempty"`
	SkillId              string                                         `json:"skill_id"`
	Skill                SkillsPluginsDetachSkillGroupOutputSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                      `json:"created_at"`
	UpdatedAt            time.Time                                      `json:"updated_at"`
}

// SkillsPluginsDetachSkillGroupOutput represents the skills plugins detach skill group output type.
type SkillsPluginsDetachSkillGroupOutput struct {
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
	Source    SkillsPluginsDetachSkillGroupOutputSource   `json:"source"`
	Skills    []SkillsPluginsDetachSkillGroupOutputSkills `json:"skills"`
	CreatedAt time.Time                                   `json:"created_at"`
	UpdatedAt time.Time                                   `json:"updated_at"`
}

// MapSkillsPluginsDetachSkillGroupOutputFromJSON deserializes JSON data into a SkillsPluginsDetachSkillGroupOutput.
func MapSkillsPluginsDetachSkillGroupOutputFromJSON(data []byte) (*SkillsPluginsDetachSkillGroupOutput, error) {
	var v SkillsPluginsDetachSkillGroupOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsDetachSkillGroupOutputToJSON serializes a SkillsPluginsDetachSkillGroupOutput to JSON.
func MapSkillsPluginsDetachSkillGroupOutputToJSON(v *SkillsPluginsDetachSkillGroupOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsDetachSkillGroupBody represents the skills plugins detach skill group body type.
type SkillsPluginsDetachSkillGroupBody struct{}

// MapSkillsPluginsDetachSkillGroupBodyFromJSON deserializes JSON data into a SkillsPluginsDetachSkillGroupBody.
func MapSkillsPluginsDetachSkillGroupBodyFromJSON(data []byte) (*SkillsPluginsDetachSkillGroupBody, error) {
	var v SkillsPluginsDetachSkillGroupBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsDetachSkillGroupBodyToJSON serializes a SkillsPluginsDetachSkillGroupBody to JSON.
func MapSkillsPluginsDetachSkillGroupBodyToJSON(v *SkillsPluginsDetachSkillGroupBody) ([]byte, error) {
	return json.Marshal(v)
}
