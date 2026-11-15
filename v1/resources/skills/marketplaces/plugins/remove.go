package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceUnsyncedSkills represents the skills marketplaces plugins remove output skill plugin source unsynced skills type.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesPluginsRemoveOutputSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkillImage represents the skills marketplaces plugins remove output skill plugin skills skill image type.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkill represents the skills marketplaces plugins remove output skill plugin skills skill type.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkill struct {
	Object            string                                                           `json:"object"`
	Id                string                                                           `json:"id"`
	Status            string                                                           `json:"status"`
	Slug              string                                                           `json:"slug"`
	Name              string                                                           `json:"name"`
	Description       *string                                                          `json:"description,omitempty"`
	ImageUrl          string                                                           `json:"image_url"`
	Image             SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                           `json:"client_name"`
	ClientDescription *string                                                          `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                                  `json:"client_metadata,omitempty"`
	License           *string                                                          `json:"license,omitempty"`
	Compatibility     *string                                                          `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                                  `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                        `json:"created_at"`
	UpdatedAt         time.Time                                                        `json:"updated_at"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills represents the skills marketplaces plugins remove output skill plugin skills type.
type SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills struct {
	Object               string                                                      `json:"object"`
	Id                   string                                                      `json:"id"`
	Identifier           string                                                      `json:"identifier"`
	Status               string                                                      `json:"status"`
	ClientName           *string                                                     `json:"client_name,omitempty"`
	ClientDescription    *string                                                     `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                             `json:"client_metadata,omitempty"`
	License              *string                                                     `json:"license,omitempty"`
	Compatibility        *string                                                     `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                     `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                      `json:"skill_id"`
	Skill                SkillsMarketplacesPluginsRemoveOutputSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                   `json:"created_at"`
	UpdatedAt            time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesPluginsRemoveOutputSkillPlugin represents the skills marketplaces plugins remove output skill plugin type.
type SkillsMarketplacesPluginsRemoveOutputSkillPlugin struct {
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
	Source    SkillsMarketplacesPluginsRemoveOutputSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesPluginsRemoveOutputSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                `json:"created_at"`
	UpdatedAt time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesPluginsRemoveOutput represents the skills marketplaces plugins remove output type.
type SkillsMarketplacesPluginsRemoveOutput struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsRemoveOutputSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// MapSkillsMarketplacesPluginsRemoveOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsRemoveOutput.
func MapSkillsMarketplacesPluginsRemoveOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsRemoveOutput, error) {
	var v SkillsMarketplacesPluginsRemoveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsRemoveOutputToJSON serializes a SkillsMarketplacesPluginsRemoveOutput to JSON.
func MapSkillsMarketplacesPluginsRemoveOutputToJSON(v *SkillsMarketplacesPluginsRemoveOutput) ([]byte, error) {
	return json.Marshal(v)
}
