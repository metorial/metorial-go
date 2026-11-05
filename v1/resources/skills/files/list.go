package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesListOutputItems represents the skills files list output items type.
type SkillsFilesListOutputItems struct {
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

// SkillsFilesListOutputPagination represents the skills files list output pagination type.
type SkillsFilesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsFilesListOutput represents the skills files list output type.
type SkillsFilesListOutput struct {
	Items      []SkillsFilesListOutputItems    `json:"items"`
	Pagination SkillsFilesListOutputPagination `json:"pagination"`
}

// MapSkillsFilesListOutputFromJSON deserializes JSON data into a SkillsFilesListOutput.
func MapSkillsFilesListOutputFromJSON(data []byte) (*SkillsFilesListOutput, error) {
	var v SkillsFilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesListOutputToJSON serializes a SkillsFilesListOutput to JSON.
func MapSkillsFilesListOutputToJSON(v *SkillsFilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}
