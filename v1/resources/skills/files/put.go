package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesPutOutputFile represents the skills files put output file type.
type SkillsFilesPutOutputFile struct {
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

// SkillsFilesPutOutput represents the skills files put output type.
type SkillsFilesPutOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The kind of change that was applied.
	Type string `json:"type"`
	// Path - The resulting path of the file. For moves, this is the new path.
	Path string `json:"path"`
	// PreviousPath - The path the file was moved from. Only set for moves.
	PreviousPath *string                   `json:"previous_path,omitempty"`
	File         *SkillsFilesPutOutputFile `json:"file,omitempty"`
}

// MapSkillsFilesPutOutputFromJSON deserializes JSON data into a SkillsFilesPutOutput.
func MapSkillsFilesPutOutputFromJSON(data []byte) (*SkillsFilesPutOutput, error) {
	var v SkillsFilesPutOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesPutOutputToJSON serializes a SkillsFilesPutOutput to JSON.
func MapSkillsFilesPutOutputToJSON(v *SkillsFilesPutOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsFilesPutBody represents the skills files put body type.
type SkillsFilesPutBody struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// Content - The new text content of the file. Markdown files are stored as documents, other files as plain text files.
	Content *string `json:"content,omitempty"`
	// FileId - The ID of a previously uploaded file to place at the path. Use this for binary files. Files that belong to a document cannot be attached; use content to write markdown documents.
	FileId *string `json:"file_id,omitempty"`
	// IfMatch - Only apply the change if the file still has this etag. Pass null to require that no file exists at the path yet. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// MapSkillsFilesPutBodyFromJSON deserializes JSON data into a SkillsFilesPutBody.
func MapSkillsFilesPutBodyFromJSON(data []byte) (*SkillsFilesPutBody, error) {
	var v SkillsFilesPutBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesPutBodyToJSON serializes a SkillsFilesPutBody to JSON.
func MapSkillsFilesPutBodyToJSON(v *SkillsFilesPutBody) ([]byte, error) {
	return json.Marshal(v)
}
