package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceUnsyncedSkills represents the skills marketplaces archive output plugins skill plugin source unsynced skills type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesArchiveOutputPluginsSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkillImage represents the skills marketplaces archive output plugins skill plugin skills skill image type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill represents the skills marketplaces archive output plugins skill plugin skills skill type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill struct {
	Object            string                                                            `json:"object"`
	Id                string                                                            `json:"id"`
	Status            string                                                            `json:"status"`
	Slug              string                                                            `json:"slug"`
	Name              string                                                            `json:"name"`
	Description       *string                                                           `json:"description,omitempty"`
	ImageUrl          string                                                            `json:"image_url"`
	Image             SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                            `json:"client_name"`
	ClientDescription *string                                                           `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                                   `json:"client_metadata,omitempty"`
	License           *string                                                           `json:"license,omitempty"`
	Compatibility     *string                                                           `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                                   `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                         `json:"created_at"`
	UpdatedAt         time.Time                                                         `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills represents the skills marketplaces archive output plugins skill plugin skills type.
type SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills struct {
	Object               string                                                       `json:"object"`
	Id                   string                                                       `json:"id"`
	Identifier           string                                                       `json:"identifier"`
	Status               string                                                       `json:"status"`
	ClientName           *string                                                      `json:"client_name,omitempty"`
	ClientDescription    *string                                                      `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                              `json:"client_metadata,omitempty"`
	License              *string                                                      `json:"license,omitempty"`
	Compatibility        *string                                                      `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                      `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                       `json:"skill_id"`
	Skill                SkillsMarketplacesArchiveOutputPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                    `json:"created_at"`
	UpdatedAt            time.Time                                                    `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutputPluginsSkillPlugin represents the skills marketplaces archive output plugins skill plugin type.
type SkillsMarketplacesArchiveOutputPluginsSkillPlugin struct {
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
	Source    SkillsMarketplacesArchiveOutputPluginsSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesArchiveOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                 `json:"created_at"`
	UpdatedAt time.Time                                                 `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutputPlugins represents the skills marketplaces archive output plugins type.
type SkillsMarketplacesArchiveOutputPlugins struct {
	Object               string                                             `json:"object"`
	Id                   string                                             `json:"id"`
	Status               string                                             `json:"status"`
	Identifier           string                                             `json:"identifier"`
	SkillConfigurationId *string                                            `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                            `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesArchiveOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                          `json:"created_at"`
	UpdatedAt            time.Time                                          `json:"updated_at"`
}

// SkillsMarketplacesArchiveOutput represents the skills marketplaces archive output type.
type SkillsMarketplacesArchiveOutput struct {
	Object                string                                   `json:"object"`
	Id                    string                                   `json:"id"`
	Status                string                                   `json:"status"`
	RepositoryAccessMode  string                                   `json:"repository_access_mode"`
	ForceMergeOrPush      bool                                     `json:"force_merge_or_push"`
	MergeBeforeChecksPass bool                                     `json:"merge_before_checks_pass"`
	SyncStatus            string                                   `json:"sync_status"`
	ImageUrl              string                                   `json:"image_url"`
	Name                  string                                   `json:"name"`
	Description           *string                                  `json:"description,omitempty"`
	Slug                  string                                   `json:"slug"`
	SkillConfigurationId  *string                                  `json:"skill_configuration_id,omitempty"`
	Plugins               []SkillsMarketplacesArchiveOutputPlugins `json:"plugins"`
	CreatedAt             time.Time                                `json:"created_at"`
	UpdatedAt             time.Time                                `json:"updated_at"`
}

// MapSkillsMarketplacesArchiveOutputFromJSON deserializes JSON data into a SkillsMarketplacesArchiveOutput.
func MapSkillsMarketplacesArchiveOutputFromJSON(data []byte) (*SkillsMarketplacesArchiveOutput, error) {
	var v SkillsMarketplacesArchiveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesArchiveOutputToJSON serializes a SkillsMarketplacesArchiveOutput to JSON.
func MapSkillsMarketplacesArchiveOutputToJSON(v *SkillsMarketplacesArchiveOutput) ([]byte, error) {
	return json.Marshal(v)
}
