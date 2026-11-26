package links

import (
	"encoding/json"
	"time"
)

// FilesLinksCreateOutput represents the files links create output type.
type FilesLinksCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// FileId - File this link downloads.
	FileId string `json:"file_id"`
	// Url - Public URL for downloading the file.
	Url string `json:"url"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - When the link stops working. `null` if it never expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapFilesLinksCreateOutputFromJSON deserializes JSON data into a FilesLinksCreateOutput.
func MapFilesLinksCreateOutputFromJSON(data []byte) (*FilesLinksCreateOutput, error) {
	var v FilesLinksCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksCreateOutputToJSON serializes a FilesLinksCreateOutput to JSON.
func MapFilesLinksCreateOutputToJSON(v *FilesLinksCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FilesLinksCreateBody represents the files links create body type.
type FilesLinksCreateBody struct {
	// FileId - File to share, by ID.
	FileId string `json:"file_id"`
	// ExpiresAt - When the link stops working. The link never expires when omitted.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// MapFilesLinksCreateBodyFromJSON deserializes JSON data into a FilesLinksCreateBody.
func MapFilesLinksCreateBodyFromJSON(data []byte) (*FilesLinksCreateBody, error) {
	var v FilesLinksCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksCreateBodyToJSON serializes a FilesLinksCreateBody to JSON.
func MapFilesLinksCreateBodyToJSON(v *FilesLinksCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
