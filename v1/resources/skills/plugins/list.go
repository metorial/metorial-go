package plugins

import (
	"encoding/json"
	"time"
)

// SkillsPluginsListOutputItemsSourceSkillGroup - The skill group the plugin's skills are synced from.
type SkillsPluginsListOutputItemsSourceSkillGroup struct {
	// Id - The ID of the skill group.
	Id string `json:"id"`
	// Name - The name of the skill group.
	Name string `json:"name"`
}

// SkillsPluginsListOutputItemsSourceUnsyncedSkills represents the skills plugins list output items source unsynced skills type.
type SkillsPluginsListOutputItemsSourceUnsyncedSkills struct {
	// SkillId - The ID of the skill that was not synced.
	SkillId string `json:"skill_id"`
	// Reason - Why the skill was not synced.
	Reason string `json:"reason"`
	// Message - A human-readable explanation.
	Message string `json:"message"`
}

// SkillsPluginsListOutputItemsSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsPluginsListOutputItemsSource struct {
	// Type - The plugin's skills are synced from a skill group and cannot be edited on the plugin directly.
	Type *string `json:"type,omitempty"`
	// SkillGroup - The skill group the plugin's skills are synced from.
	SkillGroup *SkillsPluginsListOutputItemsSourceSkillGroup `json:"skill_group,omitempty"`
	// UnsyncedSkills - Group skills that are not in the plugin, for example because a plugin or marketplace limit was reached.
	UnsyncedSkills *[]SkillsPluginsListOutputItemsSourceUnsyncedSkills `json:"unsynced_skills,omitempty"`
	// LastSyncedAt - When the plugin was last reconciled with the skill group.
	LastSyncedAt *time.Time `json:"last_synced_at,omitempty"`
}

// SkillsPluginsListOutputItemsSkillsSkillImage - Image shown for the skill.
type SkillsPluginsListOutputItemsSkillsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsListOutputItemsSkillsSkill represents the skills plugins list output items skills skill type.
type SkillsPluginsListOutputItemsSkillsSkill struct {
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
	Image SkillsPluginsListOutputItemsSkillsSkillImage `json:"image"`
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

// SkillsPluginsListOutputItemsSkills represents the skills plugins list output items skills type.
type SkillsPluginsListOutputItemsSkills struct {
	Object               string                                  `json:"object"`
	Id                   string                                  `json:"id"`
	Identifier           string                                  `json:"identifier"`
	Status               string                                  `json:"status"`
	ClientName           *string                                 `json:"client_name,omitempty"`
	ClientDescription    *string                                 `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any                         `json:"client_metadata,omitempty"`
	License              *string                                 `json:"license,omitempty"`
	Compatibility        *string                                 `json:"compatibility,omitempty"`
	SkillConfigurationId *string                                 `json:"skill_configuration_id,omitempty"`
	SkillId              string                                  `json:"skill_id"`
	Skill                SkillsPluginsListOutputItemsSkillsSkill `json:"skill"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            time.Time                               `json:"updated_at"`
}

// SkillsPluginsListOutputItems represents the skills plugins list output items type.
type SkillsPluginsListOutputItems struct {
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
	Source    SkillsPluginsListOutputItemsSource   `json:"source"`
	Skills    []SkillsPluginsListOutputItemsSkills `json:"skills"`
	CreatedAt time.Time                            `json:"created_at"`
	UpdatedAt time.Time                            `json:"updated_at"`
}

// SkillsPluginsListOutputPagination represents the skills plugins list output pagination type.
type SkillsPluginsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsPluginsListOutput represents the skills plugins list output type.
type SkillsPluginsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []SkillsPluginsListOutputItems    `json:"items"`
	Pagination SkillsPluginsListOutputPagination `json:"pagination"`
}

// MapSkillsPluginsListOutputFromJSON deserializes JSON data into a SkillsPluginsListOutput.
func MapSkillsPluginsListOutputFromJSON(data []byte) (*SkillsPluginsListOutput, error) {
	var v SkillsPluginsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsListOutputToJSON serializes a SkillsPluginsListOutput to JSON.
func MapSkillsPluginsListOutputToJSON(v *SkillsPluginsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsListQueryCreatedAt - Only return results whose skill plugin creation time is in this range. Set `gt`, `lt`, or both.
type SkillsPluginsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsListQueryUpdatedAt - Only return results whose skill plugin last update time is in this range. Set `gt`, `lt`, or both.
type SkillsPluginsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsListQuery represents the skills plugins list query type.
type SkillsPluginsListQuery struct {
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
	// Id - Only return skill plugins with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillMarketplaceId - Only return skill plugins in one of these marketplaces.
	SkillMarketplaceId *any `json:"skill_marketplace_id,omitempty"`
	// Status - Only return skill plugins with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Category - Only return skill plugins in this category.
	Category *string `json:"category,omitempty"`
	// Search - Only return skill plugins matching this search text.
	Search *string `json:"search,omitempty"`
	// SkillConfigurationId - Only return skill plugins using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// SkillGroupId - Only return plugins whose skills are synced from these skill groups.
	SkillGroupId *any `json:"skill_group_id,omitempty"`
	// CreatedAt - Only return results whose skill plugin creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsPluginsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill plugin last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SkillsPluginsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsPluginsListQueryFromJSON deserializes JSON data into a SkillsPluginsListQuery.
func MapSkillsPluginsListQueryFromJSON(data []byte) (*SkillsPluginsListQuery, error) {
	var v SkillsPluginsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsListQueryToJSON serializes a SkillsPluginsListQuery to JSON.
func MapSkillsPluginsListQueryToJSON(v *SkillsPluginsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
