package builds

import (
	"encoding/json"
	"time"
)

// SandboxImagesBuildsCreateOutputFailure represents the sandbox images builds create output failure type.
type SandboxImagesBuildsCreateOutputFailure struct {
	// Code - A stable, machine-readable code describing why the build failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxImagesBuildsCreateOutputRegions represents the sandbox images builds create output regions type.
type SandboxImagesBuildsCreateOutputRegions struct {
	// Region - The region identifier
	Region string `json:"region"`
	// Status - The build's progress in that region
	Status string `json:"status"`
}

// SandboxImagesBuildsCreateOutput represents the sandbox images builds create output type.
type SandboxImagesBuildsCreateOutput struct {
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
	Status  string                                  `json:"status"`
	Failure *SandboxImagesBuildsCreateOutputFailure `json:"failure,omitempty"`
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
	Regions []SandboxImagesBuildsCreateOutputRegions `json:"regions"`
	// HasLogs - Whether build logs are available
	HasLogs bool `json:"has_logs"`
	// ContextUploadedAt - When the build context finished uploading
	ContextUploadedAt *time.Time `json:"context_uploaded_at,omitempty"`
	BuiltAt           *time.Time `json:"built_at,omitempty"`
	EndedAt           *time.Time `json:"ended_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
}

// MapSandboxImagesBuildsCreateOutputFromJSON deserializes JSON data into a SandboxImagesBuildsCreateOutput.
func MapSandboxImagesBuildsCreateOutputFromJSON(data []byte) (*SandboxImagesBuildsCreateOutput, error) {
	var v SandboxImagesBuildsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesBuildsCreateOutputToJSON serializes a SandboxImagesBuildsCreateOutput to JSON.
func MapSandboxImagesBuildsCreateOutputToJSON(v *SandboxImagesBuildsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesBuildsCreateBody represents the sandbox images builds create body type.
type SandboxImagesBuildsCreateBody struct {
	// FileId - The file holding the build context as a zip, tar, or tar.gz archive. Limited to 50 MB.
	FileId string `json:"file_id"`
	// DockerfilePath - The path of the Dockerfile inside the context
	DockerfilePath *string `json:"dockerfile_path,omitempty"`
	// ContextPath - The directory inside the archive to build from
	ContextPath *string `json:"context_path,omitempty"`
	// Target - The build stage to build
	Target *string `json:"target,omitempty"`
	// BuildArgs - Build arguments passed to the Dockerfile
	BuildArgs *map[string]string `json:"build_args,omitempty"`
	// TimeoutSeconds - How long the build may run before it is cancelled
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
	// DiskSizeBytes - The size of disks created from the image
	DiskSizeBytes *float64 `json:"disk_size_bytes,omitempty"`
}

// MapSandboxImagesBuildsCreateBodyFromJSON deserializes JSON data into a SandboxImagesBuildsCreateBody.
func MapSandboxImagesBuildsCreateBodyFromJSON(data []byte) (*SandboxImagesBuildsCreateBody, error) {
	var v SandboxImagesBuildsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesBuildsCreateBodyToJSON serializes a SandboxImagesBuildsCreateBody to JSON.
func MapSandboxImagesBuildsCreateBodyToJSON(v *SandboxImagesBuildsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
