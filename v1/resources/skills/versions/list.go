package versions

import (
	"encoding/json"
	"time"
)

// SkillsVersionsListOutputItems represents the skills versions list output items type.
type SkillsVersionsListOutputItems struct {
	Object         string    `json:"object"`
	Id             string    `json:"id"`
	SkillId        string    `json:"skill_id"`
	StoreId        string    `json:"store_id"`
	StoreVersionId string    `json:"store_version_id"`
	VersionNumber  float64   `json:"version_number"`
	CreatedAt      time.Time `json:"created_at"`
}

// SkillsVersionsListOutputPagination represents the skills versions list output pagination type.
type SkillsVersionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsVersionsListOutput represents the skills versions list output type.
type SkillsVersionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []SkillsVersionsListOutputItems    `json:"items"`
	Pagination SkillsVersionsListOutputPagination `json:"pagination"`
}

// MapSkillsVersionsListOutputFromJSON deserializes JSON data into a SkillsVersionsListOutput.
func MapSkillsVersionsListOutputFromJSON(data []byte) (*SkillsVersionsListOutput, error) {
	var v SkillsVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsVersionsListOutputToJSON serializes a SkillsVersionsListOutput to JSON.
func MapSkillsVersionsListOutputToJSON(v *SkillsVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsVersionsListQuery represents the skills versions list query type.
type SkillsVersionsListQuery struct {
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
}

// MapSkillsVersionsListQueryFromJSON deserializes JSON data into a SkillsVersionsListQuery.
func MapSkillsVersionsListQueryFromJSON(data []byte) (*SkillsVersionsListQuery, error) {
	var v SkillsVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsVersionsListQueryToJSON serializes a SkillsVersionsListQuery to JSON.
func MapSkillsVersionsListQueryToJSON(v *SkillsVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
