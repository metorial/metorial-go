package imports

import (
	"encoding/json"
	"time"
)

// SkillsImportsListOutputItemsSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsImportsListOutputItemsSource struct {
	Type           *string `json:"type,omitempty"`
	RepositoryUrl  *string `json:"repository_url,omitempty"`
	RepositoryName *string `json:"repository_name,omitempty"`
	Ref            *string `json:"ref,omitempty"`
	RepositoryId   *string `json:"repository_id,omitempty"`
	Path           *string `json:"path,omitempty"`
	FileId         *string `json:"file_id,omitempty"`
	FileName       *string `json:"file_name,omitempty"`
	Format         *string `json:"format,omitempty"`
}

// SkillsImportsListOutputItemsItemsSkill represents the skills imports list output items items skill type.
type SkillsImportsListOutputItemsItemsSkill struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// SkillsImportsListOutputItemsItems represents the skills imports list output items items type.
type SkillsImportsListOutputItemsItems struct {
	Object      string                                  `json:"object"`
	Id          string                                  `json:"id"`
	Status      string                                  `json:"status"`
	Path        string                                  `json:"path"`
	Error       *string                                 `json:"error,omitempty"`
	Skill       *SkillsImportsListOutputItemsItemsSkill `json:"skill,omitempty"`
	StartedAt   *time.Time                              `json:"started_at,omitempty"`
	CompletedAt *time.Time                              `json:"completed_at,omitempty"`
	CreatedAt   time.Time                               `json:"created_at"`
}

// SkillsImportsListOutputItems represents the skills imports list output items type.
type SkillsImportsListOutputItems struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Source      SkillsImportsListOutputItemsSource  `json:"source"`
	Error       *string                             `json:"error,omitempty"`
	Items       []SkillsImportsListOutputItemsItems `json:"items"`
	StartedAt   *time.Time                          `json:"started_at,omitempty"`
	CompletedAt *time.Time                          `json:"completed_at,omitempty"`
	CreatedAt   time.Time                           `json:"created_at"`
}

// SkillsImportsListOutputPagination represents the skills imports list output pagination type.
type SkillsImportsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsImportsListOutput represents the skills imports list output type.
type SkillsImportsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                            `json:"object"`
	Items      []SkillsImportsListOutputItems    `json:"items"`
	Pagination SkillsImportsListOutputPagination `json:"pagination"`
}

// MapSkillsImportsListOutputFromJSON deserializes JSON data into a SkillsImportsListOutput.
func MapSkillsImportsListOutputFromJSON(data []byte) (*SkillsImportsListOutput, error) {
	var v SkillsImportsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsImportsListOutputToJSON serializes a SkillsImportsListOutput to JSON.
func MapSkillsImportsListOutputToJSON(v *SkillsImportsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsImportsListQuery represents the skills imports list query type.
type SkillsImportsListQuery struct {
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
	// Id - Only return skill imports with these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return skill imports with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// MapSkillsImportsListQueryFromJSON deserializes JSON data into a SkillsImportsListQuery.
func MapSkillsImportsListQueryFromJSON(data []byte) (*SkillsImportsListQuery, error) {
	var v SkillsImportsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsImportsListQueryToJSON serializes a SkillsImportsListQuery to JSON.
func MapSkillsImportsListQueryToJSON(v *SkillsImportsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
