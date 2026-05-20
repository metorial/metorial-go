package links

import (
	"encoding/json"
	"time"
)

// FilesLinksGetOutput represents the files links get output type.
type FilesLinksGetOutput struct {
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

// MapFilesLinksGetOutputFromJSON deserializes JSON data into a FilesLinksGetOutput.
func MapFilesLinksGetOutputFromJSON(data []byte) (*FilesLinksGetOutput, error) {
	var v FilesLinksGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksGetOutputToJSON serializes a FilesLinksGetOutput to JSON.
func MapFilesLinksGetOutputToJSON(v *FilesLinksGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
