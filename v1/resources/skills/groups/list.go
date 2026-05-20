package groups

import (
	"encoding/json"
	"time"
)

// SkillsGroupsListOutputItemsSkills represents the skills groups list output items skills type.
type SkillsGroupsListOutputItemsSkills struct {
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

// SkillsGroupsListOutputItems represents the skills groups list output items type.
type SkillsGroupsListOutputItems struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Name        string                              `json:"name"`
	Description *string                             `json:"description,omitempty"`
	Metadata    *map[string]any                     `json:"metadata,omitempty"`
	Skills      []SkillsGroupsListOutputItemsSkills `json:"skills"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// SkillsGroupsListOutputPagination represents the skills groups list output pagination type.
type SkillsGroupsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsGroupsListOutput represents the skills groups list output type.
type SkillsGroupsListOutput struct {
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

// SkillsGroupsListQueryCreatedAt - Filter skill group creation time by date range
type SkillsGroupsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill group creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill group creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsListQueryUpdatedAt - Filter skill group last update time by date range
type SkillsGroupsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill group last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill group last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsGroupsListQuery represents the skills groups list query type.
type SkillsGroupsListQuery struct {
	Limit   *float64 `json:"limit,omitempty"`
	After   *string  `json:"after,omitempty"`
	Before  *string  `json:"before,omitempty"`
	Cursor  *string  `json:"cursor,omitempty"`
	Order   *string  `json:"order,omitempty"`
	Search  *string  `json:"search,omitempty"`
	Status  *any     `json:"status,omitempty"`
	Id      *any     `json:"id,omitempty"`
	SkillId *any     `json:"skill_id,omitempty"`
	// CreatedAt - Filter skill group creation time by date range
	CreatedAt *SkillsGroupsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill group last update time by date range
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
