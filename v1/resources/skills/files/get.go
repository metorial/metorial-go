package files

import (
	"encoding/json"
	"time"
)

// SkillsFilesGetOutput represents the skills files get output type.
type SkillsFilesGetOutput struct {
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
	// Content - The text content of the file, for files of kind `document`.
	Content *string `json:"content,omitempty"`
	// DownloadUrl - A short-lived URL to download the file content, for files of kind `file`.
	DownloadUrl *string `json:"download_url,omitempty"`
}

// MapSkillsFilesGetOutputFromJSON deserializes JSON data into a SkillsFilesGetOutput.
func MapSkillsFilesGetOutputFromJSON(data []byte) (*SkillsFilesGetOutput, error) {
	var v SkillsFilesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesGetOutputToJSON serializes a SkillsFilesGetOutput to JSON.
func MapSkillsFilesGetOutputToJSON(v *SkillsFilesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SkillsFilesGetQuery represents the skills files get query type.
type SkillsFilesGetQuery struct {
	// Path - Path of the file inside the skill. A leading "/" is optional. Paths may contain letters, numbers, spaces and the characters . _ - ( ) $.
	Path string `json:"path"`
}

// MapSkillsFilesGetQueryFromJSON deserializes JSON data into a SkillsFilesGetQuery.
func MapSkillsFilesGetQueryFromJSON(data []byte) (*SkillsFilesGetQuery, error) {
	var v SkillsFilesGetQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSkillsFilesGetQueryToJSON serializes a SkillsFilesGetQuery to JSON.
func MapSkillsFilesGetQueryToJSON(v *SkillsFilesGetQuery) ([]byte, error) {
	return json.Marshal(v)
}
