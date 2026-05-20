package links

import (
	"encoding/json"
	"time"
)

// FilesLinksListOutputItems represents the files links list output items type.
type FilesLinksListOutputItems struct {
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

// FilesLinksListOutputPagination represents the files links list output pagination type.
type FilesLinksListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// FilesLinksListOutput represents the files links list output type.
type FilesLinksListOutput struct {
	Items      []FilesLinksListOutputItems    `json:"items"`
	Pagination FilesLinksListOutputPagination `json:"pagination"`
}

// MapFilesLinksListOutputFromJSON deserializes JSON data into a FilesLinksListOutput.
func MapFilesLinksListOutputFromJSON(data []byte) (*FilesLinksListOutput, error) {
	var v FilesLinksListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksListOutputToJSON serializes a FilesLinksListOutput to JSON.
func MapFilesLinksListOutputToJSON(v *FilesLinksListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FilesLinksListQuery represents the files links list query type.
type FilesLinksListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// FileId - Filter by file ID
	FileId *string `json:"file_id,omitempty"`
}

// MapFilesLinksListQueryFromJSON deserializes JSON data into a FilesLinksListQuery.
func MapFilesLinksListQueryFromJSON(data []byte) (*FilesLinksListQuery, error) {
	var v FilesLinksListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFilesLinksListQueryToJSON serializes a FilesLinksListQuery to JSON.
func MapFilesLinksListQueryToJSON(v *FilesLinksListQuery) ([]byte, error) {
	return json.Marshal(v)
}
