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

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage - Image shown for the skill.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkill represents the skills marketplaces plugins list output items skill plugin skills skill type.
type SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkill struct {
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
	Image SkillsMarketplacesPluginsListOutputItemsSkillPluginSkillsSkillImage `json:"image"`
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
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMarketplacesPluginsListOutput represents the skills marketplaces plugins list output type.
type SkillsMarketplacesPluginsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                        `json:"object"`
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

// SkillsMarketplacesPluginsListQueryCreatedAt - Only return results whose skill marketplace plugin creation time is in this range. Set `gt`, `lt`, or both.
type SkillsMarketplacesPluginsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQueryUpdatedAt - Only return results whose skill marketplace plugin last update time is in this range. Set `gt`, `lt`, or both.
type SkillsMarketplacesPluginsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesPluginsListQuery represents the skills marketplaces plugins list query type.
type SkillsMarketplacesPluginsListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Only return marketplace plugins with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillPluginId - Only return links for one of these skill plugins.
	SkillPluginId *any `json:"skill_plugin_id,omitempty"`
	// Status - Only return marketplace plugins with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return marketplace plugins using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Only return results whose skill marketplace plugin creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsMarketplacesPluginsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill marketplace plugin last update time is in this range. Set `gt`, `lt`, or both.
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
