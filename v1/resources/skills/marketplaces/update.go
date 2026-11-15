package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceUnsyncedSkills represents the skills marketplaces update output plugins skill plugin source unsynced skills type.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesUpdateOutputPluginsSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkillImage represents the skills marketplaces update output plugins skill plugin skills skill image type.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkill represents the skills marketplaces update output plugins skill plugin skills skill type.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkill struct {
	Object            string                                                           `json:"object"`
	Id                string                                                           `json:"id"`
	Status            string                                                           `json:"status"`
	Slug              string                                                           `json:"slug"`
	Name              string                                                           `json:"name"`
	Description       *string                                                          `json:"description,omitempty"`
	ImageUrl          string                                                           `json:"image_url"`
	Image             SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                           `json:"client_name"`
	ClientDescription *string                                                          `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                                  `json:"client_metadata,omitempty"`
	License           *string                                                          `json:"license,omitempty"`
	Compatibility     *string                                                          `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                                  `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                        `json:"created_at"`
	UpdatedAt         time.Time                                                        `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills represents the skills marketplaces update output plugins skill plugin skills type.
type SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills struct {
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
	Skill                SkillsMarketplacesUpdateOutputPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                   `json:"created_at"`
	UpdatedAt            time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutputPluginsSkillPlugin represents the skills marketplaces update output plugins skill plugin type.
type SkillsMarketplacesUpdateOutputPluginsSkillPlugin struct {
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
	Source    SkillsMarketplacesUpdateOutputPluginsSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesUpdateOutputPluginsSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                `json:"created_at"`
	UpdatedAt time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutputPlugins represents the skills marketplaces update output plugins type.
type SkillsMarketplacesUpdateOutputPlugins struct {
	Object               string                                            `json:"object"`
	Id                   string                                            `json:"id"`
	Status               string                                            `json:"status"`
	Identifier           string                                            `json:"identifier"`
	SkillConfigurationId *string                                           `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                           `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesUpdateOutputPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                         `json:"created_at"`
	UpdatedAt            time.Time                                         `json:"updated_at"`
}

// SkillsMarketplacesUpdateOutput represents the skills marketplaces update output type.
type SkillsMarketplacesUpdateOutput struct {
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
	Plugins               []SkillsMarketplacesUpdateOutputPlugins `json:"plugins"`
	CreatedAt             time.Time                               `json:"created_at"`
	UpdatedAt             time.Time                               `json:"updated_at"`
}

// MapSkillsMarketplacesUpdateOutputFromJSON deserializes JSON data into a SkillsMarketplacesUpdateOutput.
func MapSkillsMarketplacesUpdateOutputFromJSON(data []byte) (*SkillsMarketplacesUpdateOutput, error) {
	var v SkillsMarketplacesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesUpdateOutputToJSON serializes a SkillsMarketplacesUpdateOutput to JSON.
func MapSkillsMarketplacesUpdateOutputToJSON(v *SkillsMarketplacesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesUpdateBody represents the skills marketplaces update body type.
type SkillsMarketplacesUpdateBody struct {
	Name                  *string `json:"name,omitempty"`
	Description           *string `json:"description,omitempty"`
	ImageFileId           *string `json:"image_file_id,omitempty"`
	SkillConfigurationId  *string `json:"skill_configuration_id,omitempty"`
	RepositoryAccessMode  *string `json:"repository_access_mode,omitempty"`
	ForceMergeOrPush      *bool   `json:"force_merge_or_push,omitempty"`
	MergeBeforeChecksPass *bool   `json:"merge_before_checks_pass,omitempty"`
}

// MapSkillsMarketplacesUpdateBodyFromJSON deserializes JSON data into a SkillsMarketplacesUpdateBody.
func MapSkillsMarketplacesUpdateBodyFromJSON(data []byte) (*SkillsMarketplacesUpdateBody, error) {
	var v SkillsMarketplacesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesUpdateBodyToJSON serializes a SkillsMarketplacesUpdateBody to JSON.
func MapSkillsMarketplacesUpdateBodyToJSON(v *SkillsMarketplacesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
