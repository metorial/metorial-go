package links

import (
	"encoding/json"
	"time"
)

// FilesLinksCreateOutput represents the files links create output type.
type FilesLinksCreateOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The links's unique identifier
	Id string `json:"id"`
	// FileId - The file's unique identifier
	FileId string `json:"file_id"`
	// Url - The file's public URL
	Url string `json:"url"`
	// CreatedAt - The links's creation date
	CreatedAt time.Time `json:"created_at"`
	// ExpiresAt - The file's expiration date
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
	FileId    string     `json:"file_id"`
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
