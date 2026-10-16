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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxImagesVersionsListOutput represents the sandbox images versions list output type.
type SandboxImagesVersionsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
