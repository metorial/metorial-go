package imports

import (
	"encoding/json"
	"time"
)

// SkillsImportsCreateOutputSource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsImportsCreateOutputSource struct {
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

// SkillsImportsCreateOutputItemsSkill represents the skills imports create output items skill type.
type SkillsImportsCreateOutputItemsSkill struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

// SkillsImportsCreateOutputItems represents the skills imports create output items type.
type SkillsImportsCreateOutputItems struct {
	Object      string                               `json:"object"`
	Id          string                               `json:"id"`
	Status      string                               `json:"status"`
	Path        string                               `json:"path"`
	Error       *string                              `json:"error,omitempty"`
	Skill       *SkillsImportsCreateOutputItemsSkill `json:"skill,omitempty"`
	StartedAt   *time.Time                           `json:"started_at,omitempty"`
	CompletedAt *time.Time                           `json:"completed_at,omitempty"`
	CreatedAt   time.Time                            `json:"created_at"`
}

// SkillsImportsCreateOutput represents the skills imports create output type.
type SkillsImportsCreateOutput struct {
	Object      string                           `json:"object"`
	Id          string                           `json:"id"`
	Status      string                           `json:"status"`
	Source      SkillsImportsCreateOutputSource  `json:"source"`
	Error       *string                          `json:"error,omitempty"`
	Items       []SkillsImportsCreateOutputItems `json:"items"`
	StartedAt   *time.Time                       `json:"started_at,omitempty"`
	CompletedAt *time.Time                       `json:"completed_at,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
}

// MapSkillsImportsCreateOutputFromJSON deserializes JSON data into a SkillsImportsCreateOutput.
func MapSkillsImportsCreateOutputFromJSON(data []byte) (*SkillsImportsCreateOutput, error) {
	var v SkillsImportsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsImportsCreateOutputToJSON serializes a SkillsImportsCreateOutput to JSON.
func MapSkillsImportsCreateOutputToJSON(v *SkillsImportsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsImportsCreateBodySource represents one of several possible types.
// This is a union type - only one set of fields will be populated.
type SkillsImportsCreateBodySource struct {
	// Type - Import from an uploaded file.
	Type *string `json:"type,omitempty"`
	// RepositoryUrl - Public repository URL to import from.
	RepositoryUrl *string `json:"repository_url,omitempty"`
	// Ref - Branch, tag, or commit to import. Defaults to the default branch.
	Ref *string `json:"ref,omitempty"`
	// RepositoryId - Connected repository to import from, by ID.
	RepositoryId *string `json:"repository_id,omitempty"`
	// Path - Directory within the repository to import.
	Path *string `json:"path,omitempty"`
	// FileId - Uploaded file to import, by ID.
	FileId *string `json:"file_id,omitempty"`
}

// SkillsImportsCreateBody represents the skills imports create body type.
type SkillsImportsCreateBody struct {
	// Source - Where to import skills from.
	Source SkillsImportsCreateBodySource `json:"source"`
	// SkillGroupId - The ID of a skill group that every imported skill is added to.
	SkillGroupId *string `json:"skill_group_id,omitempty"`
}

// MapSkillsImportsCreateBodyFromJSON deserializes JSON data into a SkillsImportsCreateBody.
func MapSkillsImportsCreateBodyFromJSON(data []byte) (*SkillsImportsCreateBody, error) {
	var v SkillsImportsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsImportsCreateBodyToJSON serializes a SkillsImportsCreateBody to JSON.
func MapSkillsImportsCreateBodyToJSON(v *SkillsImportsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
