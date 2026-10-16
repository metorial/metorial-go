package builds

import (
	"encoding/json"
	"time"
)

// SandboxImagesBuildsListOutputItemsFailure represents the sandbox images builds list output items failure type.
type SandboxImagesBuildsListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing why the build failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxImagesBuildsListOutputItemsRegions represents the sandbox images builds list output items regions type.
type SandboxImagesBuildsListOutputItemsRegions struct {
	// Region - The region identifier
	Region string `json:"region"`
	// Status - The build's progress in that region
	Status string `json:"status"`
}

// SandboxImagesBuildsListOutputItems represents the sandbox images builds list output items type.
type SandboxImagesBuildsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The build's unique identifier
	Id string `json:"id"`
	// ImageId - The image the build produces a version of
	ImageId string `json:"image_id"`
	// VersionId - The image version the build produces
	VersionId string `json:"version_id"`
	// VersionNumber - The number of the produced version within its image
	VersionNumber *float64 `json:"version_number,omitempty"`
	// FileId - The uploaded file holding the build context
	FileId *string `json:"file_id,omitempty"`
	// Status - The build's status
	Status  string                                     `json:"status"`
	Failure *SandboxImagesBuildsListOutputItemsFailure `json:"failure,omitempty"`
	// DockerfilePath - The Dockerfile used, relative to the build context. Null for builds started before build settings were recorded.
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// ContextPath - The directory inside the uploaded archive used as the build context. Null for builds started before build settings were recorded.
	ContextPath *string `json:"context_path,omitempty"`
	// Target - The Dockerfile stage that was built
	Target *string `json:"target,omitempty"`
	// BuildArgNames - The names of the build arguments passed to the build. Their values are not kept. Null for builds started before build settings were recorded.
	BuildArgNames *[]string `json:"build_arg_names,omitempty"`
	// TimeoutSeconds - How long the build may run. Null for builds started before build settings were recorded.
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
	// DiskSizeBytes - The size of disks created from the image
	DiskSizeBytes float64 `json:"disk_size_bytes"`
	// Regions - Build and replication progress per region
	Regions []SandboxImagesBuildsListOutputItemsRegions `json:"regions"`
	// HasLogs - Whether build logs are available
	HasLogs bool `json:"has_logs"`
	// ContextUploadedAt - When the build context finished uploading
	ContextUploadedAt *time.Time `json:"context_uploaded_at,omitempty"`
	BuiltAt           *time.Time `json:"built_at,omitempty"`
	EndedAt           *time.Time `json:"ended_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
}

// SandboxImagesBuildsListOutputPagination represents the sandbox images builds list output pagination type.
type SandboxImagesBuildsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxImagesBuildsListOutput represents the sandbox images builds list output type.
type SandboxImagesBuildsListOutput struct {
	Items      []SandboxImagesBuildsListOutputItems    `json:"items"`
	Pagination SandboxImagesBuildsListOutputPagination `json:"pagination"`
}

// MapSandboxImagesBuildsListOutputFromJSON deserializes JSON data into a SandboxImagesBuildsListOutput.
func MapSandboxImagesBuildsListOutputFromJSON(data []byte) (*SandboxImagesBuildsListOutput, error) {
	var v SandboxImagesBuildsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesBuildsListOutputToJSON serializes a SandboxImagesBuildsListOutput to JSON.
func MapSandboxImagesBuildsListOutputToJSON(v *SandboxImagesBuildsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesBuildsListQuery represents the sandbox images builds list query type.
type SandboxImagesBuildsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// ImageVersionId - Filter by image version
	ImageVersionId *any `json:"image_version_id,omitempty"`
	// Status - Filter by build status
	Status *any `json:"status,omitempty"`
}

// MapSandboxImagesBuildsListQueryFromJSON deserializes JSON data into a SandboxImagesBuildsListQuery.
func MapSandboxImagesBuildsListQueryFromJSON(data []byte) (*SandboxImagesBuildsListQuery, error) {
	var v SandboxImagesBuildsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesBuildsListQueryToJSON serializes a SandboxImagesBuildsListQuery to JSON.
func MapSandboxImagesBuildsListQueryToJSON(v *SandboxImagesBuildsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
