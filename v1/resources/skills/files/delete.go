package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesDeleteOutputFile represents the skills files delete output file type.
type SkillsFilesDeleteOutputFile struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Path - Path of the file inside the skill, always starting with "/".
	Path string `json:"path"`
	// Kind - Markdown written with `content` is stored as a collaboratively editable document; everything else, including attached files, is stored as a file.
	Kind string `json:"kind"`
	// Size - Size of the file content in bytes.
	Size float64 `json:"size"`
	// Etag - Opaque version identifier of the file content. Pass it as `if_match` when writing to avoid overwriting changes made by others. The check is best-effort: a live collaborative edit landing between the check and the write is not detected. Listings do not flush pending live edits, so a listed etag can lag behind the one returned when fetching the file.
	Etag string `json:"etag"`
	// SkillId - The ID of the skill the file belongs to.
	SkillId string `json:"skill_id"`
	// StoreItemId - The ID of the store item backing this file.
	StoreItemId string `json:"store_item_id"`
	// DocumentId - The ID of the backing document, for files of kind `document`.
	DocumentId *string `json:"document_id,omitempty"`
	// FileId - The ID of the backing file, for files of kind `file`.
	FileId *string `json:"file_id,omitempty"`
	// UpdatedAt - When the file was last changed.
	UpdatedAt time.Time `json:"updated_at"`
}

// SkillsFilesDeleteOutput represents the skills files delete output type.
type SkillsFilesDeleteOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The kind of change that was applied.
	Type string `json:"type"`
	// Path - The resulting path of the file. For moves, this is the new path.
	Path string `json:"path"`
	// PreviousPath - The path the file was moved from. Only set for moves.
	PreviousPath *string                      `json:"previous_path,omitempty"`
	File         *SkillsFilesDeleteOutputFile `json:"file,omitempty"`
}

// MapSkillsFilesDeleteOutputFromJSON deserializes JSON data into a SkillsFilesDeleteOutput.
func MapSkillsFilesDeleteOutputFromJSON(data []byte) (*SkillsFilesDeleteOutput, error) {
	var v SkillsFilesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesDeleteOutputToJSON serializes a SkillsFilesDeleteOutput to JSON.
func MapSkillsFilesDeleteOutputToJSON(v *SkillsFilesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsFilesDeleteQuery represents the skills files delete query type.
type SkillsFilesDeleteQuery struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// IfMatch - Only apply the change if the file still has this etag. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// MapSkillsFilesDeleteQueryFromJSON deserializes JSON data into a SkillsFilesDeleteQuery.
func MapSkillsFilesDeleteQueryFromJSON(data []byte) (*SkillsFilesDeleteQuery, error) {
	var v SkillsFilesDeleteQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesDeleteQueryToJSON serializes a SkillsFilesDeleteQuery to JSON.
func MapSkillsFilesDeleteQueryToJSON(v *SkillsFilesDeleteQuery) ([]byte, error) {
	return json.Marshal(v)
}
