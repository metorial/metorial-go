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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsImportsListOutput represents the skills imports list output type.
type SkillsImportsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Id     *any     `json:"id,omitempty"`
	Status *any     `json:"status,omitempty"`
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
