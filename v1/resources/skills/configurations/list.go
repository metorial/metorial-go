package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsListOutputItems represents the skills configurations list output items type.
type SkillsConfigurationsListOutputItems struct {
	// Object - String representing the object's type
	Object                      string   `json:"object"`
	Id                          string   `json:"id"`
	IsDefault                   bool     `json:"is_default"`
	AllowScripts                bool     `json:"allow_scripts"`
	AllowedFileExtensions       []string `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool     `json:"allow_non_standard_directories"`
	// SecurityReviewPolicy - How security reviews affect distribution: `warn` flags failing skills, `block` withholds them from syncs and exports
	SecurityReviewPolicy string     `json:"security_review_policy"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}

// SkillsConfigurationsListOutputPagination represents the skills configurations list output pagination type.
type SkillsConfigurationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SkillsConfigurationsListOutput represents the skills configurations list output type.
type SkillsConfigurationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                   `json:"object"`
	Items      []SkillsConfigurationsListOutputItems    `json:"items"`
	Pagination SkillsConfigurationsListOutputPagination `json:"pagination"`
}

// MapSkillsConfigurationsListOutputFromJSON deserializes JSON data into a SkillsConfigurationsListOutput.
func MapSkillsConfigurationsListOutputFromJSON(data []byte) (*SkillsConfigurationsListOutput, error) {
	var v SkillsConfigurationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsListOutputToJSON serializes a SkillsConfigurationsListOutput to JSON.
func MapSkillsConfigurationsListOutputToJSON(v *SkillsConfigurationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsConfigurationsListQuery represents the skills configurations list query type.
type SkillsConfigurationsListQuery struct {
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

// MapSkillsConfigurationsListQueryFromJSON deserializes JSON data into a SkillsConfigurationsListQuery.
func MapSkillsConfigurationsListQueryFromJSON(data []byte) (*SkillsConfigurationsListQuery, error) {
	var v SkillsConfigurationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsConfigurationsListQueryToJSON serializes a SkillsConfigurationsListQuery to JSON.
func MapSkillsConfigurationsListQueryToJSON(v *SkillsConfigurationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
