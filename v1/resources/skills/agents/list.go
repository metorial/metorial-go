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
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsAgentsListOutput represents the skills agents list output type.
type SkillsAgentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
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
	// IncludeArchived - Include archived agents in the results.
	IncludeArchived *bool `json:"include_archived,omitempty"`
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
