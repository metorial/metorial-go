package agents

import (
	"encoding/json"
	"time"
)

// SkillsAgentsListOutputItems represents the skills agents list output items type.
type SkillsAgentsListOutputItems struct {
	// Object - String representing the object's type
	Object      string     `json:"object"`
	Id          string     `json:"id"`
	SkillId     string     `json:"skill_id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Slug        string     `json:"slug"`
	Status      string     `json:"status"`
	StoreId     string     `json:"store_id"`
	StoreItemId *string    `json:"store_item_id,omitempty"`
	Path        *string    `json:"path,omitempty"`
	DocumentId  string     `json:"document_id"`
	ArchivedAt  *time.Time `json:"archived_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// SkillsAgentsListOutputPagination represents the skills agents list output pagination type.
type SkillsAgentsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsAgentsListOutput represents the skills agents list output type.
type SkillsAgentsListOutput struct {
	Items      []SkillsAgentsListOutputItems    `json:"items"`
	Pagination SkillsAgentsListOutputPagination `json:"pagination"`
}

// MapSkillsAgentsListOutputFromJSON deserializes JSON data into a SkillsAgentsListOutput.
func MapSkillsAgentsListOutputFromJSON(data []byte) (*SkillsAgentsListOutput, error) {
	var v SkillsAgentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsListOutputToJSON serializes a SkillsAgentsListOutput to JSON.
func MapSkillsAgentsListOutputToJSON(v *SkillsAgentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsAgentsListQuery represents the skills agents list query type.
type SkillsAgentsListQuery struct {
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	IncludeArchived *bool    `json:"include_archived,omitempty"`
}

// MapSkillsAgentsListQueryFromJSON deserializes JSON data into a SkillsAgentsListQuery.
func MapSkillsAgentsListQueryFromJSON(data []byte) (*SkillsAgentsListQuery, error) {
	var v SkillsAgentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsAgentsListQueryToJSON serializes a SkillsAgentsListQuery to JSON.
func MapSkillsAgentsListQueryToJSON(v *SkillsAgentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
