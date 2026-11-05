package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesMoveOutputFile represents the skills files move output file type.
type SkillsFilesMoveOutputFile struct {
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

// SkillsFilesMoveOutput represents the skills files move output type.
type SkillsFilesMoveOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Type - The kind of change that was applied.
	Type string `json:"type"`
	// Path - The resulting path of the file. For moves, this is the new path.
	Path string `json:"path"`
	// PreviousPath - The path the file was moved from. Only set for moves.
	PreviousPath *string                    `json:"previous_path,omitempty"`
	File         *SkillsFilesMoveOutputFile `json:"file,omitempty"`
}

// MapSkillsFilesMoveOutputFromJSON deserializes JSON data into a SkillsFilesMoveOutput.
func MapSkillsFilesMoveOutputFromJSON(data []byte) (*SkillsFilesMoveOutput, error) {
	var v SkillsFilesMoveOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesMoveOutputToJSON serializes a SkillsFilesMoveOutput to JSON.
func MapSkillsFilesMoveOutputToJSON(v *SkillsFilesMoveOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsFilesMoveBody represents the skills files move body type.
type SkillsFilesMoveBody struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
	// ToPath - The new path of the file.
	ToPath string `json:"to_path"`
	// IfMatch - Only apply the change if the file still has this etag. Omit to skip the check.
	IfMatch *string `json:"if_match,omitempty"`
}

// MapSkillsFilesMoveBodyFromJSON deserializes JSON data into a SkillsFilesMoveBody.
func MapSkillsFilesMoveBodyFromJSON(data []byte) (*SkillsFilesMoveBody, error) {
	var v SkillsFilesMoveBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesMoveBodyToJSON serializes a SkillsFilesMoveBody to JSON.
func MapSkillsFilesMoveBodyToJSON(v *SkillsFilesMoveBody) ([]byte, error) {
	return json.Marshal(v)
}
