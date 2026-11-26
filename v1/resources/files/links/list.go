package links

import (
	"encoding/json"
	"time"
)

// FilesLinksListOutputItems represents the files links list output items type.
type FilesLinksListOutputItems struct {
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

// FilesLinksListOutputPagination represents the files links list output pagination type.
type FilesLinksListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// FilesLinksListOutput represents the files links list output type.
type FilesLinksListOutput struct {
	// Object - String representing the object’s type.
	Object     string                         `json:"object"`
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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// FileId - Only return links for this file.
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
