package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesBatchOutputItemsFile represents the skills files batch output items file type.
type SkillsFilesBatchOutputItemsFile struct {
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

// SkillsFilesBatchOutputItems represents the skills files batch output items type.
type SkillsFilesBatchOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The kind of change that was applied.
	Type string `json:"type"`
	// Path - The resulting path of the file. For moves, this is the new path.
	Path string `json:"path"`
	// PreviousPath - The path the file was moved from. Only set for moves.
	PreviousPath *string                          `json:"previous_path,omitempty"`
	File         *SkillsFilesBatchOutputItemsFile `json:"file,omitempty"`
}

// SkillsFilesBatchOutputPagination represents the skills files batch output pagination type.
type SkillsFilesBatchOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SkillsFilesBatchOutput represents the skills files batch output type.
type SkillsFilesBatchOutput struct {
	Items      []SkillsFilesBatchOutputItems    `json:"items"`
	Pagination SkillsFilesBatchOutputPagination `json:"pagination"`
}

// MapSkillsFilesBatchOutputFromJSON deserializes JSON data into a SkillsFilesBatchOutput.
func MapSkillsFilesBatchOutputFromJSON(data []byte) (*SkillsFilesBatchOutput, error) {
	var v SkillsFilesBatchOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesBatchOutputToJSON serializes a SkillsFilesBatchOutput to JSON.
func MapSkillsFilesBatchOutputToJSON(v *SkillsFilesBatchOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsFilesBatchBody represents the skills files batch body type.
type SkillsFilesBatchBody struct {
	// Changes - The changes to apply, in order.
	Changes []any `json:"changes"`
}

// MapSkillsFilesBatchBodyFromJSON deserializes JSON data into a SkillsFilesBatchBody.
func MapSkillsFilesBatchBodyFromJSON(data []byte) (*SkillsFilesBatchBody, error) {
	var v SkillsFilesBatchBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesBatchBodyToJSON serializes a SkillsFilesBatchBody to JSON.
func MapSkillsFilesBatchBodyToJSON(v *SkillsFilesBatchBody) ([]byte, error) {
	return json.Marshal(v)
}
