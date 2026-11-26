package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsListOutputItemsSkillImage - Image shown for the skill.
type SkillsGroupsItemsListOutputItemsSkillImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsItemsListOutputItemsSkill represents the skills groups items list output items skill type.
type SkillsGroupsItemsListOutputItemsSkill struct {
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
	Image SkillsGroupsItemsListOutputItemsSkillImage `json:"image"`
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

// SkillsGroupsItemsListOutputItems represents the skills groups items list output items type.
type SkillsGroupsItemsListOutputItems struct {
	Object       string                                `json:"object"`
	Id           string                                `json:"id"`
	Status       string                                `json:"status"`
	SkillGroupId string                                `json:"skill_group_id"`
	Skill        SkillsGroupsItemsListOutputItemsSkill `json:"skill"`
	CreatedAt    time.Time                             `json:"created_at"`
}

// SkillsGroupsItemsListOutputPagination represents the skills groups items list output pagination type.
type SkillsGroupsItemsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsGroupsItemsListOutput represents the skills groups items list output type.
type SkillsGroupsItemsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []SkillsGroupsItemsListOutputItems    `json:"items"`
	Pagination SkillsGroupsItemsListOutputPagination `json:"pagination"`
}

// MapSkillsGroupsItemsListOutputFromJSON deserializes JSON data into a SkillsGroupsItemsListOutput.
func MapSkillsGroupsItemsListOutputFromJSON(data []byte) (*SkillsGroupsItemsListOutput, error) {
	var v SkillsGroupsItemsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsListOutputToJSON serializes a SkillsGroupsItemsListOutput to JSON.
func MapSkillsGroupsItemsListOutputToJSON(v *SkillsGroupsItemsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsGroupsItemsListQueryCreatedAt - Only return results whose skill group item creation time is in this range. Set `gt`, `lt`, or both.
type SkillsGroupsItemsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsItemsListQuery represents the skills groups items list query type.
type SkillsGroupsItemsListQuery struct {
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
	// Status - Only return skill group items with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return skill group items with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return skill group items for one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// CreatedAt - Only return results whose skill group item creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsGroupsItemsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapSkillsGroupsItemsListQueryFromJSON deserializes JSON data into a SkillsGroupsItemsListQuery.
func MapSkillsGroupsItemsListQueryFromJSON(data []byte) (*SkillsGroupsItemsListQuery, error) {
	var v SkillsGroupsItemsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsItemsListQueryToJSON serializes a SkillsGroupsItemsListQuery to JSON.
func MapSkillsGroupsItemsListQueryToJSON(v *SkillsGroupsItemsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
