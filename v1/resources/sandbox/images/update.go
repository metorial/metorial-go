package images

import (
	"encoding/json"
	"time"
)

// SandboxImagesUpdateOutput represents the sandbox images update output type.
type SandboxImagesUpdateOutput struct {
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

// MapSandboxImagesUpdateOutputFromJSON deserializes JSON data into a SandboxImagesUpdateOutput.
func MapSandboxImagesUpdateOutputFromJSON(data []byte) (*SandboxImagesUpdateOutput, error) {
	var v SandboxImagesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesUpdateOutputToJSON serializes a SandboxImagesUpdateOutput to JSON.
func MapSandboxImagesUpdateOutputToJSON(v *SandboxImagesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImagesUpdateBody represents the sandbox images update body type.
type SandboxImagesUpdateBody struct {
	// Name - A name for the image
	Name *string `json:"name,omitempty"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapSandboxImagesUpdateBodyFromJSON deserializes JSON data into a SandboxImagesUpdateBody.
func MapSandboxImagesUpdateBodyFromJSON(data []byte) (*SandboxImagesUpdateBody, error) {
	var v SandboxImagesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesUpdateBodyToJSON serializes a SandboxImagesUpdateBody to JSON.
func MapSandboxImagesUpdateBodyToJSON(v *SandboxImagesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
