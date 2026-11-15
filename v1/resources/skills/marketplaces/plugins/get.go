package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsGetOutputSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesPluginsGetOutputSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesPluginsGetOutputSkillPluginSourceUnsyncedSkills represents the skills marketplaces plugins get output skill plugin source unsynced skills type.
type SkillsMarketplacesPluginsGetOutputSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesPluginsGetOutputSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesPluginsGetOutputSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesPluginsGetOutputSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesPluginsGetOutputSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkillImage represents the skills marketplaces plugins get output skill plugin skills skill image type.
type SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkill represents the skills marketplaces plugins get output skill plugin skills skill type.
type SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkill struct {
	Object            string                                                        `json:"object"`
	Id                string                                                        `json:"id"`
	Status            string                                                        `json:"status"`
	Slug              string                                                        `json:"slug"`
	Name              string                                                        `json:"name"`
	Description       *string                                                       `json:"description,omitempty"`
	ImageUrl          string                                                        `json:"image_url"`
	Image             SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                        `json:"client_name"`
	ClientDescription *string                                                       `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                               `json:"client_metadata,omitempty"`
	License           *string                                                       `json:"license,omitempty"`
	Compatibility     *string                                                       `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                               `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                     `json:"created_at"`
	UpdatedAt         time.Time                                                     `json:"updated_at"`
}

// SkillsMarketplacesPluginsGetOutputSkillPluginSkills represents the skills marketplaces plugins get output skill plugin skills type.
type SkillsMarketplacesPluginsGetOutputSkillPluginSkills struct {
	Object               string                                                   `json:"object"`
	Id                   string                                                   `json:"id"`
	Identifier           string                                                   `json:"identifier"`
	Status               string                                                   `json:"status"`
	ClientName           *string                                                  `json:"client_name,omitempty"`
	ClientDescription    *string                                                  `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                          `json:"client_metadata,omitempty"`
	License              *string                                                  `json:"license,omitempty"`
	Compatibility        *string                                                  `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                  `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                   `json:"skill_id"`
	Skill                SkillsMarketplacesPluginsGetOutputSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesPluginsGetOutputSkillPlugin represents the skills marketplaces plugins get output skill plugin type.
type SkillsMarketplacesPluginsGetOutputSkillPlugin struct {
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
	Source    SkillsMarketplacesPluginsGetOutputSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesPluginsGetOutputSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                             `json:"created_at"`
	UpdatedAt time.Time                                             `json:"updated_at"`
}

// SkillsMarketplacesPluginsGetOutput represents the skills marketplaces plugins get output type.
type SkillsMarketplacesPluginsGetOutput struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Status               string                                         `json:"status"`
	Identifier           string                                         `json:"identifier"`
	SkillConfigurationId *string                                        `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                        `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsGetOutputSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                      `json:"created_at"`
	UpdatedAt            time.Time                                      `json:"updated_at"`
}

// MapSkillsMarketplacesPluginsGetOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsGetOutput.
func MapSkillsMarketplacesPluginsGetOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsGetOutput, error) {
	var v SkillsMarketplacesPluginsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsGetOutputToJSON serializes a SkillsMarketplacesPluginsGetOutput to JSON.
func MapSkillsMarketplacesPluginsGetOutputToJSON(v *SkillsMarketplacesPluginsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
