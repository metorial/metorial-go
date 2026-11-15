package plugins

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceUnsyncedSkills represents the skills marketplaces plugins list output items skill plugin source unsynced skills type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesPluginsListOutputItemsSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage represents the skills marketplaces plugins list output items skill plugin skills skill image type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage struct {
	Object string  `json:"object"`
	Type   string  `json:"type"`
	Icon   *string `json:"icon,omitempty"`
	Color  *string `json:"color,omitempty"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkill represents the skills marketplaces plugins list output items skill plugin skills skill type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkill struct {
	Object            string                                                              `json:"object"`
	Id                string                                                              `json:"id"`
	Status            string                                                              `json:"status"`
	Slug              string                                                              `json:"slug"`
	Name              string                                                              `json:"name"`
	Description       *string                                                             `json:"description,omitempty"`
	ImageUrl          string                                                              `json:"image_url"`
	Image             SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage `json:"image"`
	ClientName        string                                                              `json:"client_name"`
	ClientDescription *string                                                             `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any                                                     `json:"client_metadata,omitempty"`
	License           *string                                                             `json:"license,omitempty"`
	Compatibility     *string                                                             `json:"compatibility,omitempty"`
	Metadata          *map[string]any                                                     `json:"metadata,omitempty"`
	CreatedAt         time.Time                                                           `json:"created_at"`
	UpdatedAt         time.Time                                                           `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills represents the skills marketplaces plugins list output items skill plugin skills type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills struct {
	Object               string                                                         `json:"object"`
	Id                   string                                                         `json:"id"`
	Identifier           string                                                         `json:"identifier"`
	Status               string                                                         `json:"status"`
	ClientName           *string                                                        `json:"client_name,omitempty"`
	ClientDescription    *string                                                        `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                                                `json:"client_metadata,omitempty"`
	License              *string                                                        `json:"license,omitempty"`
	Compatibility        *string                                                        `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                                        `json:"skill_configuration_id,omitempty"`
	SkillId              string                                                         `json:"skill_id"`
	Skill                SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                      `json:"created_at"`
	UpdatedAt            time.Time                                                      `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPlugin represents the skills marketplaces plugins list output items skill plugin type.
type SkillsMarketplacesPluginsListOutputItemsSkillPlugin struct {
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
	Source    SkillsMarketplacesPluginsListOutputItemsSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesPluginsListOutputItemsSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                   `json:"created_at"`
	UpdatedAt time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputItems represents the skills marketplaces plugins list output items type.
type SkillsMarketplacesPluginsListOutputItems struct {
	Object               string                                               `json:"object"`
	Id                   string                                               `json:"id"`
	Status               string                                               `json:"status"`
	Identifier           string                                               `json:"identifier"`
	SkillConfigurationId *string                                              `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                              `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesPluginsListOutputItemsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                            `json:"created_at"`
	UpdatedAt            time.Time                                            `json:"updated_at"`
}

// SkillsMarketplacesPluginsListOutputPagination represents the skills marketplaces plugins list output pagination type.
type SkillsMarketplacesPluginsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsMarketplacesPluginsListOutput represents the skills marketplaces plugins list output type.
type SkillsMarketplacesPluginsListOutput struct {
	Items      []SkillsMarketplacesPluginsListOutputItems    `json:"items"`
	Pagination SkillsMarketplacesPluginsListOutputPagination `json:"pagination"`
}

// MapSkillsMarketplacesPluginsListOutputFromJSON deserializes JSON data into a SkillsMarketplacesPluginsListOutput.
func MapSkillsMarketplacesPluginsListOutputFromJSON(data []byte) (*SkillsMarketplacesPluginsListOutput, error) {
	var v SkillsMarketplacesPluginsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsListOutputToJSON serializes a SkillsMarketplacesPluginsListOutput to JSON.
func MapSkillsMarketplacesPluginsListOutputToJSON(v *SkillsMarketplacesPluginsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesPluginsListQueryCreatedAt - Filter skill marketplace plugin creation time by date range
type SkillsMarketplacesPluginsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace plugin creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace plugin creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQueryUpdatedAt - Filter skill marketplace plugin last update time by date range
type SkillsMarketplacesPluginsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill marketplace plugin last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill marketplace plugin last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQuery represents the skills marketplaces plugins list query type.
type SkillsMarketplacesPluginsListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillPluginId        *any     `json:"skill_plugin_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill marketplace plugin creation time by date range
	CreatedAt *SkillsMarketplacesPluginsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill marketplace plugin last update time by date range
	UpdatedAt *SkillsMarketplacesPluginsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsMarketplacesPluginsListQueryFromJSON deserializes JSON data into a SkillsMarketplacesPluginsListQuery.
func MapSkillsMarketplacesPluginsListQueryFromJSON(data []byte) (*SkillsMarketplacesPluginsListQuery, error) {
	var v SkillsMarketplacesPluginsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesPluginsListQueryToJSON serializes a SkillsMarketplacesPluginsListQuery to JSON.
func MapSkillsMarketplacesPluginsListQueryToJSON(v *SkillsMarketplacesPluginsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
