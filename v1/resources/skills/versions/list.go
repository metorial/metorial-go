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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsVersionsListOutput represents the skills versions list output type.
type SkillsVersionsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
