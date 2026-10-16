package images

import (
	"encoding/json"
	"time"
)

// SandboxImagesListOutputItems represents the sandbox images list output items type.
type SandboxImagesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Metadata - User-defined metadata stored in Metorial
	Metadata map[string]any `json:"metadata"`
	// Id - The image's unique identifier
	Id string `json:"id"`
	// Owner - Whether you built the image or Metorial provides it to everyone
	Owner string `json:"owner"`
	// Status - The image's status. Images are pending until their first build succeeds.
	Status string `json:"status"`
	// Name - The image's name
	Name *string `json:"name,omitempty"`
	// Identifier - The image's human-readable identifier, set for global images
	Identifier *string `json:"identifier,omitempty"`
	// Regions - The identifiers of the regions the image is available in
	Regions []string `json:"regions"`
	// LatestVersionId - The most recent version of the image
	LatestVersionId *string `json:"latest_version_id,omitempty"`
	// LatestVersionNumber - The version number of the most recent active image version
	LatestVersionNumber *float64   `json:"latest_version_number,omitempty"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	DeletedAt           *time.Time `json:"deleted_at,omitempty"`
}

// SandboxImagesListOutputPagination represents the sandbox images list output pagination type.
type SandboxImagesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxImagesListOutput represents the sandbox images list output type.
type SandboxImagesListOutput struct {
	Items      []SandboxImagesListOutputItems    `json:"items"`
	Pagination SandboxImagesListOutputPagination `json:"pagination"`
}

// MapSandboxImagesListOutputFromJSON deserializes JSON data into a SandboxImagesListOutput.
func MapSandboxImagesListOutputFromJSON(data []byte) (*SandboxImagesListOutput, error) {
	var v SandboxImagesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesListOutputToJSON serializes a SandboxImagesListOutput to JSON.
func MapSandboxImagesListOutputToJSON(v *SandboxImagesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesListQuery represents the sandbox images list query type.
type SandboxImagesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Owner - Whether to list your own images, global images, or both (the default)
	Owner *any `json:"owner,omitempty"`
	// Status - Filter by image status
	Status *any `json:"status,omitempty"`
	// Search - Search by name or identifier
	Search *string `json:"search,omitempty"`
}

// MapSandboxImagesListQueryFromJSON deserializes JSON data into a SandboxImagesListQuery.
func MapSandboxImagesListQueryFromJSON(data []byte) (*SandboxImagesListQuery, error) {
	var v SandboxImagesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesListQueryToJSON serializes a SandboxImagesListQuery to JSON.
func MapSandboxImagesListQueryToJSON(v *SandboxImagesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
