package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsListOutputItemsSkillImage - Image shown for the skill.
type SkillsPluginsSkillsListOutputItemsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsPluginsSkillsListOutputItemsSkill represents the skills plugins skills list output items skill type.
type SkillsPluginsSkillsListOutputItemsSkill struct {
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
	Image SkillsPluginsSkillsListOutputItemsSkillImage `json:"image"`
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

// SkillsPluginsSkillsListOutputItems represents the skills plugins skills list output items type.
type SkillsPluginsSkillsListOutputItems struct {
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
	Skill                SkillsPluginsSkillsListOutputItemsSkill `json:"skill"`
	CreatedAt            time.Time                               `json:"created_at"`
	UpdatedAt            time.Time                               `json:"updated_at"`
}

// SkillsPluginsSkillsListOutputPagination represents the skills plugins skills list output pagination type.
type SkillsPluginsSkillsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsPluginsSkillsListOutput represents the skills plugins skills list output type.
type SkillsPluginsSkillsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []SkillsPluginsSkillsListOutputItems    `json:"items"`
	Pagination SkillsPluginsSkillsListOutputPagination `json:"pagination"`
}

// MapSkillsPluginsSkillsListOutputFromJSON deserializes JSON data into a SkillsPluginsSkillsListOutput.
func MapSkillsPluginsSkillsListOutputFromJSON(data []byte) (*SkillsPluginsSkillsListOutput, error) {
	var v SkillsPluginsSkillsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsListOutputToJSON serializes a SkillsPluginsSkillsListOutput to JSON.
func MapSkillsPluginsSkillsListOutputToJSON(v *SkillsPluginsSkillsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsPluginsSkillsListQueryCreatedAt - Only return results whose skill plugin skill creation time is in this range. Set `gt`, `lt`, or both.
type SkillsPluginsSkillsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsSkillsListQueryUpdatedAt - Only return results whose skill plugin skill last update time is in this range. Set `gt`, `lt`, or both.
type SkillsPluginsSkillsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsSkillsListQuery represents the skills plugins skills list query type.
type SkillsPluginsSkillsListQuery struct {
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
	// Id - Only return plugin skills with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return links for one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// Status - Only return plugin skills with one of these statuses.
	Status *any `json:"status,omitempty"`
	// SkillConfigurationId - Only return plugin skills using one of these skill configurations.
	SkillConfigurationId *any `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Only return results whose skill plugin skill creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsPluginsSkillsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill plugin skill last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SkillsPluginsSkillsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsPluginsSkillsListQueryFromJSON deserializes JSON data into a SkillsPluginsSkillsListQuery.
func MapSkillsPluginsSkillsListQueryFromJSON(data []byte) (*SkillsPluginsSkillsListQuery, error) {
	var v SkillsPluginsSkillsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsPluginsSkillsListQueryToJSON serializes a SkillsPluginsSkillsListQuery to JSON.
func MapSkillsPluginsSkillsListQueryToJSON(v *SkillsPluginsSkillsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
