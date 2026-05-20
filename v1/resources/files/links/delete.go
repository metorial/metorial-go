package links

import (
	"encoding/json"
	"time"
)

// FilesLinksDeleteOutput represents the files links delete output type.
type FilesLinksDeleteOutput struct {
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

// MapFilesLinksDeleteOutputFromJSON deserializes JSON data into a FilesLinksDeleteOutput.
func MapFilesLinksDeleteOutputFromJSON(data []byte) (*FilesLinksDeleteOutput, error) {
	var v FilesLinksDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksDeleteOutputToJSON serializes a FilesLinksDeleteOutput to JSON.
func MapFilesLinksDeleteOutputToJSON(v *FilesLinksDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
