package plugins

import (
	"encoding/json"
	"github.com/metorial/metorial-go/v1/nullable"
	"time"
)

// SkillsMarketplacesPluginsAddOutputSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesPluginsAddOutputSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesPluginsAddOutputSkillPluginSourceUnsyncedSkills represents the skills marketplaces plugins add output skill plugin source unsynced skills type.
type SkillsMarketplacesPluginsAddOutputSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesPluginsAddOutputSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesPluginsAddOutputSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesPluginsAddOutputSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesPluginsAddOutputSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkillImage - Image shown for the skill.
type SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkill represents the skills marketplaces plugins add output skill plugin skills skill type.
type SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkill struct {
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
	Image SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkillImage `json:"image"`
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

// SkillsMarketplacesPluginsAddOutputSkillPluginSkills represents the skills marketplaces plugins add output skill plugin skills type.
type SkillsMarketplacesPluginsAddOutputSkillPluginSkills struct {
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
	Skill                SkillsMarketplacesPluginsAddOutputSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                `json:"created_at"`
	UpdatedAt            time.Time                                                `json:"updated_at"`
}

// SkillsMarketplacesPluginsAddOutputSkillPlugin represents the skills marketplaces plugins add output skill plugin type.
type SkillsMarketplacesPluginsAddOutputSkillPlugin struct {
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
	Source    SkillsMarketplacesPluginsAddOutputSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesPluginsAddOutputSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                             `json:"created_at"`
	UpdatedAt time.Time                                             `json:"updated_at"`
}

// SkillsMarketplacesPluginsAddOutput represents the skills marketplaces plugins add output type.
type SkillsMarketplacesPluginsAddOutput struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Status               string                                         `json:"status"`
	Identifier           string                                         `json:"identifier"`
	SkillConfigurationId *string                                        `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                        `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsAddOutputSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                      `json:"created_at"`
	UpdatedAt            time.Time                                      `json:"updated_at"`
}

// MapSkillsMarketplacesPluginsAddOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsAddOutput.
func MapSkillsMarketplacesPluginsAddOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsAddOutput, error) {
	var v SkillsMarketplacesPluginsAddOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsAddOutputToJSON serializes a SkillsMarketplacesPluginsAddOutput to JSON.
func MapSkillsMarketplacesPluginsAddOutputToJSON(v *SkillsMarketplacesPluginsAddOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesPluginsAddBody represents the skills marketplaces plugins add body type.
type SkillsMarketplacesPluginsAddBody struct {
	// SkillPluginId - Skill plugin to add to the marketplace, by ID.
	SkillPluginId string `json:"skill_plugin_id"`
	// SkillConfigurationId - Skill configuration for the plugin link, by ID. `null` uses the default configuration.
	SkillConfigurationId *nullable.Nullable[string] `json:"skill_configuration_id,omitempty"`
	// Identifier - Slug identifying the plugin in the marketplace. Defaults to the plugin name.
	Identifier *string `json:"identifier,omitempty"`
}

// MapSkillsMarketplacesPluginsAddBodyFromJSON deserializes JSON data into a SkillsMarketplacesPluginsAddBody.
func MapSkillsMarketplacesPluginsAddBodyFromJSON(data []byte) (*SkillsMarketplacesPluginsAddBody, error) {
	var v SkillsMarketplacesPluginsAddBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsAddBodyToJSON serializes a SkillsMarketplacesPluginsAddBody to JSON.
func MapSkillsMarketplacesPluginsAddBodyToJSON(v *SkillsMarketplacesPluginsAddBody) ([]byte, error) {
	return json.Marshal(v)
}
