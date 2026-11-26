package links

import (
	"encoding/json"
	"time"
)

// FilesLinksDeleteOutput represents the files links delete output type.
type FilesLinksDeleteOutput struct {
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
