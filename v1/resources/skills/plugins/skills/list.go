package skills

import (
	"encoding/json"
	"time"
)

// SkillsPluginsSkillsListOutputItems represents the skills plugins skills list output items type.
type SkillsPluginsSkillsListOutputItems struct {
	Object               string          `json:"object"`
	Id                   string          `json:"id"`
	Identifier           string          `json:"identifier"`
	Status               string          `json:"status"`
	ClientName           *string         `json:"client_name,omitempty"`
	ClientDescription    *string         `json:"client_description,omitempty"`
	ClientMetadata       *map[string]any `json:"client_metadata,omitempty"`
	License              *string         `json:"license,omitempty"`
	Compatibility        *string         `json:"compatibility,omitempty"`
	SkillConfigurationId *string         `json:"skill_configuration_id,omitempty"`
	SkillId              string          `json:"skill_id"`
	CreatedAt            time.Time       `json:"created_at"`
	UpdatedAt            time.Time       `json:"updated_at"`
}

// SkillsPluginsSkillsListOutputPagination represents the skills plugins skills list output pagination type.
type SkillsPluginsSkillsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsPluginsSkillsListOutput represents the skills plugins skills list output type.
type SkillsPluginsSkillsListOutput struct {
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

// SkillsPluginsSkillsListQueryCreatedAt - Filter skill plugin skill creation time by date range
type SkillsPluginsSkillsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for skill plugin skill creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill plugin skill creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsSkillsListQueryUpdatedAt - Filter skill plugin skill last update time by date range
type SkillsPluginsSkillsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for skill plugin skill last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for skill plugin skill last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// SkillsPluginsSkillsListQuery represents the skills plugins skills list query type.
type SkillsPluginsSkillsListQuery struct {
	Limit                *float64 `json:"limit,omitempty"`
	After                *string  `json:"after,omitempty"`
	Before               *string  `json:"before,omitempty"`
	Cursor               *string  `json:"cursor,omitempty"`
	Order                *string  `json:"order,omitempty"`
	Id                   *any     `json:"id,omitempty"`
	SkillId              *any     `json:"skill_id,omitempty"`
	Status               *any     `json:"status,omitempty"`
	SkillConfigurationId *any     `json:"skill_configuration_id,omitempty"`
	// CreatedAt - Filter skill plugin skill creation time by date range
	CreatedAt *SkillsPluginsSkillsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter skill plugin skill last update time by date range
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
