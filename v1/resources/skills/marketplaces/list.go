package marketplaces

import (
	"encoding/json"
	"time"
)

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceUnsyncedSkills represents the skills marketplaces list output items plugins skill plugin source unsynced skills type.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsMarketplacesListOutputItemsPluginsSkillPluginSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkillImage - Image shown for the skill.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkill represents the skills marketplaces list output items plugins skill plugin skills skill type.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkill struct {
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
	Image SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkillImage `json:"image"`
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

// SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills represents the skills marketplaces list output items plugins skill plugin skills type.
type SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills struct {
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
	Skill                SkillsMarketplacesListOutputItemsPluginsSkillPluginSkillsSkill `json:"skill"`
	CreatedAt            time.Time                                                      `json:"created_at"`
	UpdatedAt            time.Time                                                      `json:"updated_at"`
}

// SkillsMarketplacesListOutputItemsPluginsSkillPlugin represents the skills marketplaces list output items plugins skill plugin type.
type SkillsMarketplacesListOutputItemsPluginsSkillPlugin struct {
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
	Source    SkillsMarketplacesListOutputItemsPluginsSkillPluginSource   `json:"source"`
	Skills    []SkillsMarketplacesListOutputItemsPluginsSkillPluginSkills `json:"skills"`
	CreatedAt time.Time                                                   `json:"created_at"`
	UpdatedAt time.Time                                                   `json:"updated_at"`
}

// SkillsMarketplacesListOutputItemsPlugins represents the skills marketplaces list output items plugins type.
type SkillsMarketplacesListOutputItemsPlugins struct {
	Object               string                                               `json:"object"`
	Id                   string                                               `json:"id"`
	Status               string                                               `json:"status"`
	Identifier           string                                               `json:"identifier"`
	SkillConfigurationId *string                                              `json:"skill_configuration_id,omitempty"`
	SkillMarketplaceId   *string                                              `json:"skill_marketplace_id,omitempty"`
	SkillPlugin          *SkillsMarketplacesListOutputItemsPluginsSkillPlugin `json:"skill_plugin,omitempty"`
	CreatedAt            time.Time                                            `json:"created_at"`
	UpdatedAt            time.Time                                            `json:"updated_at"`
}

// SkillsMarketplacesListOutputItems represents the skills marketplaces list output items type.
type SkillsMarketplacesListOutputItems struct {
	Object                string                                     `json:"object"`
	Id                    string                                     `json:"id"`
	Status                string                                     `json:"status"`
	RepositoryAccessMode  string                                     `json:"repository_access_mode"`
	ForceMergeOrPush      bool                                       `json:"force_merge_or_push"`
	MergeBeforeChecksPass bool                                       `json:"merge_before_checks_pass"`
	SyncStatus            string                                     `json:"sync_status"`
	ImageUrl              string                                     `json:"image_url"`
	Name                  string                                     `json:"name"`
	Description           *string                                    `json:"description,omitempty"`
	Slug                  string                                     `json:"slug"`
	SkillConfigurationId  *string                                    `json:"skill_configuration_id,omitempty"`
	Plugins               []SkillsMarketplacesListOutputItemsPlugins `json:"plugins"`
	CreatedAt             time.Time                                  `json:"created_at"`
	UpdatedAt             time.Time                                  `json:"updated_at"`
}

// SkillsMarketplacesListOutputPagination represents the skills marketplaces list output pagination type.
type SkillsMarketplacesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsMarketplacesListOutput represents the skills marketplaces list output type.
type SkillsMarketplacesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []SkillsMarketplacesListOutputItems    `json:"items"`
	Pagination SkillsMarketplacesListOutputPagination `json:"pagination"`
}

// MapSkillsMarketplacesListOutputFromJSON deserializes JSON data into a SkillsMarketplacesListOutput.
func MapSkillsMarketplacesListOutputFromJSON(data []byte) (*SkillsMarketplacesListOutput, error) {
	var v SkillsMarketplacesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesListOutputToJSON serializes a SkillsMarketplacesListOutput to JSON.
func MapSkillsMarketplacesListOutputToJSON(v *SkillsMarketplacesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsMarketplacesListQueryCreatedAt - Only return results whose skill marketplace creation time is in this range. Set `gt`, `lt`, or both.
type SkillsMarketplacesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesListQueryUpdatedAt - Only return results whose skill marketplace last update time is in this range. Set `gt`, `lt`, or both.
type SkillsMarketplacesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsMarketplacesListQuery represents the skills marketplaces list query type.
type SkillsMarketplacesListQuery struct {
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
	// Id - Only return skill marketplaces with these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return skill marketplaces with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return skill marketplaces using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// SkillGroupId - Only return marketplaces that contain an active plugin synced from these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// Search - Only return skill marketplaces matching this search text.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose skill marketplace creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsMarketplacesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill marketplace last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SkillsMarketplacesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsMarketplacesListQueryFromJSON deserializes JSON data into a SkillsMarketplacesListQuery.
func MapSkillsMarketplacesListQueryFromJSON(data []byte) (*SkillsMarketplacesListQuery, error) {
	var v SkillsMarketplacesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsMarketplacesListQueryToJSON serializes a SkillsMarketplacesListQuery to JSON.
func MapSkillsMarketplacesListQueryToJSON(v *SkillsMarketplacesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
