package versions

import (
	"encoding/json"
	"time"
)

// SandboxImagesVersionsListOutputItems represents the sandbox images versions list output items type.
type SandboxImagesVersionsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The version's unique identifier
	Id string `json:"id"`
	// ImageId - The image the version belongs to
	ImageId string `json:"image_id"`
	// VersionNumber - The position of the version within its image, starting at 1. Null only briefly for versions that are still being numbered.
	VersionNumber *float64 `json:"version_number,omitempty"`
	// BuildId - The build that produced the version. Null for versions of global images.
	BuildId *string `json:"build_id,omitempty"`
	// Status - The version's status
	Status string `json:"status"`
	// Regions - The identifiers of the regions the version is available in
	Regions   []string  `json:"regions"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// SandboxImagesVersionsListOutputPagination represents the sandbox images versions list output pagination type.
type SandboxImagesVersionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxImagesVersionsListOutput represents the sandbox images versions list output type.
type SandboxImagesVersionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []SandboxImagesVersionsListOutputItems    `json:"items"`
	Pagination SandboxImagesVersionsListOutputPagination `json:"pagination"`
}

// MapSandboxImagesVersionsListOutputFromJSON deserializes JSON data into a SandboxImagesVersionsListOutput.
func MapSandboxImagesVersionsListOutputFromJSON(data []byte) (*SandboxImagesVersionsListOutput, error) {
	var v SandboxImagesVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesVersionsListOutputToJSON serializes a SandboxImagesVersionsListOutput to JSON.
func MapSandboxImagesVersionsListOutputToJSON(v *SandboxImagesVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesVersionsListQuery represents the sandbox images versions list query type.
type SandboxImagesVersionsListQuery struct {
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
}

// MapSandboxImagesVersionsListQueryFromJSON deserializes JSON data into a SandboxImagesVersionsListQuery.
func MapSandboxImagesVersionsListQueryFromJSON(data []byte) (*SandboxImagesVersionsListQuery, error) {
	var v SandboxImagesVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesVersionsListQueryToJSON serializes a SandboxImagesVersionsListQuery to JSON.
func MapSandboxImagesVersionsListQueryToJSON(v *SandboxImagesVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
