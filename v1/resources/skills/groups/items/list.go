package items

import (
	"encoding/json"
	"time"
)

// SkillsGroupsItemsListOutputItemsSkill represents the skills groups items list output items skill type.
type SkillsGroupsItemsListOutputItemsSkill struct {
	Object            string          `json:"object"`
	Id                string          `json:"id"`
	Status            string          `json:"status"`
	Slug              string          `json:"slug"`
	Name              string          `json:"name"`
	Description       *string         `json:"description,omitempty"`
	ImageUrl          string          `json:"image_url"`
	ClientName        string          `json:"client_name"`
	ClientDescription *string         `json:"client_description,omitempty"`
	ClientMetadata    *map[string]any `json:"client_metadata,omitempty"`
	License           *string         `json:"license,omitempty"`
	Compatibility     *string         `json:"compatibility,omitempty"`
	Metadata          *map[string]any `json:"metadata,omitempty"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsGroupsItemsListOutput represents the skills groups items list output type.
type SkillsGroupsItemsListOutput struct {
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

// SkillsGroupsItemsListQueryCreatedAt - Filter skill group item creation time by date range
type SkillsGroupsItemsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill group item creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill group item creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsItemsListQuery represents the skills groups items list query type.
type SkillsGroupsItemsListQuery struct {
	Limit   *float64 `json:"limit,omitempty"`
	After   *string  `json:"after,omitempty"`
	Before  *string  `json:"before,omitempty"`
	Cursor  *string  `json:"cursor,omitempty"`
	Order   *string  `json:"order,omitempty"`
	Status  *any     `json:"status,omitempty"`
	Id      *any     `json:"id,omitempty"`
	SkillId *any     `json:"skill_id,omitempty"`
	// CreatedAt - Filter skill group item creation time by date range
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
