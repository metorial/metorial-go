package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsListOutputItemsImage - Image shown for the skill.
type SkillsGroupsListOutputItemsImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsListOutputItemsSkillsImage - Image shown for the skill.
type SkillsGroupsListOutputItemsSkillsImage struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Where the skill image comes from.
	Type string `json:"type"`
	// Icon - Library icon ID. `null` unless `type` is `icon`.
	Icon *string `json:"icon,omitempty"`
	// Color - Tile color. `null` unless `type` is `icon`.
	Color *string `json:"color,omitempty"`
}

// SkillsGroupsListOutputItemsSkills represents the skills groups list output items skills type.
type SkillsGroupsListOutputItemsSkills struct {
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
	Image SkillsGroupsListOutputItemsSkillsImage `json:"image"`
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

// SkillsGroupsListOutputItems represents the skills groups list output items type.
type SkillsGroupsListOutputItems struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ImageUrl    string          `json:"image_url"`
	// Image - Image shown for the skill.
	Image                        SkillsGroupsListOutputItemsImage    `json:"image"`
	AllowConsumerSkillAssignment bool                                `json:"allow_consumer_skill_assignment"`
	Skills                       []SkillsGroupsListOutputItemsSkills `json:"skills"`
	CreatedAt                    time.Time                           `json:"created_at"`
	UpdatedAt                    time.Time                           `json:"updated_at"`
}

// SkillsGroupsListOutputPagination represents the skills groups list output pagination type.
type SkillsGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsGroupsListOutput represents the skills groups list output type.
type SkillsGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
	Items      []SkillsGroupsListOutputItems    `json:"items"`
	Pagination SkillsGroupsListOutputPagination `json:"pagination"`
}

// MapSkillsGroupsListOutputFromJSON deserializes JSON data into a SkillsGroupsListOutput.
func MapSkillsGroupsListOutputFromJSON(data []byte) (*SkillsGroupsListOutput, error) {
	var v SkillsGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsListOutputToJSON serializes a SkillsGroupsListOutput to JSON.
func MapSkillsGroupsListOutputToJSON(v *SkillsGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsGroupsListQueryCreatedAt - Only return results whose skill group creation time is in this range. Set `gt`, `lt`, or both.
type SkillsGroupsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsListQueryUpdatedAt - Only return results whose skill group last update time is in this range. Set `gt`, `lt`, or both.
type SkillsGroupsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsListQuery represents the skills groups list query type.
type SkillsGroupsListQuery struct {
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
	// Search - Only return skill groups matching this search text.
	Search *string `json:"search,omitempty"`
	// Status - Only return skill groups with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return skill groups with these IDs.
	Id *any `json:"id,omitempty"`
	// SkillId - Only return skill groups containing one of these skills.
	SkillId *any `json:"skill_id,omitempty"`
	// CreatedAt - Only return results whose skill group creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *SkillsGroupsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose skill group last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *SkillsGroupsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapSkillsGroupsListQueryFromJSON deserializes JSON data into a SkillsGroupsListQuery.
func MapSkillsGroupsListQueryFromJSON(data []byte) (*SkillsGroupsListQuery, error) {
	var v SkillsGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsGroupsListQueryToJSON serializes a SkillsGroupsListQuery to JSON.
func MapSkillsGroupsListQueryToJSON(v *SkillsGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
