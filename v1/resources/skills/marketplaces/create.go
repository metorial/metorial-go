package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesCreateOutputPluginsSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPluginSourceUnsyncedSkills represents the skills marketplaces create output plugins skill plugin source unsynced skills type.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesCreateOutputPluginsSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesCreateOutputPluginsSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkillImage represents the skills marketplaces create output plugins skill plugin skills skill image type.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkill represents the skills marketplaces create output plugins skill plugin skills skill type.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkill struct {
	Object            string                                                           `json:"object"`
	Id                string                                                           `json:"id"`
	Status            string                                                           `json:"status"`
	Slug              string                                                           `json:"slug"`
	Name              string                                                           `json:"name"`
	Description       *string                                                          `json:"description,omitempty"`
	ImageUrl          string                                                           `json:"image_url"`
	Image             SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                           `json:"client_name"`
	ClientDescription *string                                                          `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                                  `json:"client_metadata,omitempty"`
	License           *string                                                          `json:"license,omitempty"`
	Compatibility     *string                                                          `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                                  `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                        `json:"created_at"`
	UpdatedAt         time.Time                                                        `json:"updated_at"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPluginSkills represents the skills marketplaces create output plugins skill plugin skills type.
type SkillsMarketplacesCreateOutputPluginsSkillPluginSkills struct {
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
	Skill                SkillsMarketplacesCreateOutputPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                   `json:"created_at"`
	UpdatedAt            time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesCreateOutputPluginsSkillPlugin represents the skills marketplaces create output plugins skill plugin type.
type SkillsMarketplacesCreateOutputPluginsSkillPlugin struct {
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
	Source    SkillsMarketplacesCreateOutputPluginsSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesCreateOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                `json:"created_at"`
	UpdatedAt time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesCreateOutputPlugins represents the skills marketplaces create output plugins type.
type SkillsMarketplacesCreateOutputPlugins struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesCreateOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// SkillsMarketplacesCreateOutput represents the skills marketplaces create output type.
type SkillsMarketplacesCreateOutput struct {
	Object                string                                  `json:"object"`
	Id                    string                                  `json:"id"`
	Status                string                                  `json:"status"`
	RepositoryAccessMode  string                                  `json:"repository_access_mode"`
	ForceMergeOrPush      bool                                    `json:"force_merge_or_push"`
	MergeBeforeChecksPass bool                                    `json:"merge_before_checks_pass"`
	SyncStatus            string                                  `json:"sync_status"`
	ImageUrl              string                                  `json:"image_url"`
	Name                  string                                  `json:"name"`
	Description           *string                                 `json:"description,omitempty"`
	Slug                  string                                  `json:"slug"`
	SkillConfigurationId  *string                                 `json:"skill_configuration_id,omitempty"`
	Plugins               []SkillsMarketplacesCreateOutputPlugins `json:"plugins"`
	CreatedAt             time.Time                               `json:"created_at"`
	UpdatedAt             time.Time                               `json:"updated_at"`
}

// MapSkillsMarketplacesCreateOutputFromJSON deserializes JSON data into a SkillsMarketplacesCreateOutput.
func MapSkillsMarketplacesCreateOutputFromJSON(data []byte) (*SkillsMarketplacesCreateOutput, error) {
	var v SkillsMarketplacesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesCreateOutputToJSON serializes a SkillsMarketplacesCreateOutput to JSON.
func MapSkillsMarketplacesCreateOutputToJSON(v *SkillsMarketplacesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesCreateBody represents the skills marketplaces create body type.
type SkillsMarketplacesCreateBody struct {
	Name                  string  `json:"name"`
	Description           *string `json:"description,omitempty"`
	ImageFileId           *string `json:"image_file_id,omitempty"`
	SkillConfigurationId  *string `json:"skill_configuration_id,omitempty"`
	RepositoryAccessMode  *string `json:"repository_access_mode,omitempty"`
	ForceMergeOrPush      *bool   `json:"force_merge_or_push,omitempty"`
	MergeBeforeChecksPass *bool   `json:"merge_before_checks_pass,omitempty"`
}

// MapSkillsMarketplacesCreateBodyFromJSON deserializes JSON data into a SkillsMarketplacesCreateBody.
func MapSkillsMarketplacesCreateBodyFromJSON(data []byte) (*SkillsMarketplacesCreateBody, error) {
	var v SkillsMarketplacesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesCreateBodyToJSON serializes a SkillsMarketplacesCreateBody to JSON.
func MapSkillsMarketplacesCreateBodyToJSON(v *SkillsMarketplacesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
