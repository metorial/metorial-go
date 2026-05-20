package configurations

import (
	"encoding/json"
	"time"
)

// SkillsConfigurationsListOutputItems represents the skills configurations list output items type.
type SkillsConfigurationsListOutputItems struct {
	// Object - String representing the object's type
	Object                      string     `json:"object"`
	Id                          string     `json:"id"`
	IsDefault                   bool       `json:"is_default"`
	AllowScripts                bool       `json:"allow_scripts"`
	AllowedFileExtensions       []string   `json:"allowed_file_extensions"`
	AllowNonStandardDirectories bool       `json:"allow_non_standard_directories"`
	DeletedAt                   *time.Time `json:"deleted_at,omitempty"`
	CreatedAt                   time.Time  `json:"created_at"`
	UpdatedAt                   time.Time  `json:"updated_at"`
}

// SkillsConfigurationsListOutputPagination represents the skills configurations list output pagination type.
type SkillsConfigurationsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsConfigurationsListOutput represents the skills configurations list output type.
type SkillsConfigurationsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
