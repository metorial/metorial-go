package links

import (
	"encoding/json"
	"time"
)

// FilesLinksGetOutput represents the files links get output type.
type FilesLinksGetOutput struct {
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
