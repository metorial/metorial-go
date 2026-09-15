package imports

import (
	"encoding/json"
	"time"
)

// SkillsImportsGetOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsImportsGetOutputSource struct {
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

// SkillsImportsGetOutputItemsSkill represents the skills imports get output items skill type.
type SkillsImportsGetOutputItemsSkill struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// SkillsImportsGetOutputItems represents the skills imports get output items type.
type SkillsImportsGetOutputItems struct {
	Object      string                            `json:"object"`
	Id          string                            `json:"id"`
	Status      string                            `json:"status"`
	Path        string                            `json:"path"`
	Error       *string                           `json:"error,omitempty"`
	Skill       *SkillsImportsGetOutputItemsSkill `json:"skill,omitempty"`
	StartedAt   *time.Time                        `json:"started_at,omitempty"`
	CompletedAt *time.Time                        `json:"completed_at,omitempty"`
	CreatedAt   time.Time                         `json:"created_at"`
}

// SkillsImportsGetOutput represents the skills imports get output type.
type SkillsImportsGetOutput struct {
	Object      string                        `json:"object"`
	Id          string                        `json:"id"`
	Status      string                        `json:"status"`
	Source      SkillsImportsGetOutputSource  `json:"source"`
	Error       *string                       `json:"error,omitempty"`
	Items       []SkillsImportsGetOutputItems `json:"items"`
	StartedAt   *time.Time                    `json:"started_at,omitempty"`
	CompletedAt *time.Time                    `json:"completed_at,omitempty"`
	CreatedAt   time.Time                     `json:"created_at"`
}

// MapSkillsImportsGetOutputFromJSON deserializes JSON data into a SkillsImportsGetOutput.
func MapSkillsImportsGetOutputFromJSON(data []byte) (*SkillsImportsGetOutput, error) {
	var v SkillsImportsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsImportsGetOutputToJSON serializes a SkillsImportsGetOutput to JSON.
func MapSkillsImportsGetOutputToJSON(v *SkillsImportsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
