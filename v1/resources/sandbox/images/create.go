package images

import (
	"encoding/json"
	"time"
)

// SandboxImagesCreateOutput represents the sandbox images create output type.
type SandboxImagesCreateOutput struct {
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

// MapSandboxImagesCreateOutputFromJSON deserializes JSON data into a SandboxImagesCreateOutput.
func MapSandboxImagesCreateOutputFromJSON(data []byte) (*SandboxImagesCreateOutput, error) {
	var v SandboxImagesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesCreateOutputToJSON serializes a SandboxImagesCreateOutput to JSON.
func MapSandboxImagesCreateOutputToJSON(v *SandboxImagesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesCreateBody represents the sandbox images create body type.
type SandboxImagesCreateBody struct {
	// Name - A name for the image
	Name *string `json:"name,omitempty"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapSandboxImagesCreateBodyFromJSON deserializes JSON data into a SandboxImagesCreateBody.
func MapSandboxImagesCreateBodyFromJSON(data []byte) (*SandboxImagesCreateBody, error) {
	var v SandboxImagesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesCreateBodyToJSON serializes a SandboxImagesCreateBody to JSON.
func MapSandboxImagesCreateBodyToJSON(v *SandboxImagesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
