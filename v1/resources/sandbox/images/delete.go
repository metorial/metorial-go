package images

import (
	"encoding/json"
	"time"
)

// SandboxImagesDeleteOutput represents the sandbox images delete output type.
type SandboxImagesDeleteOutput struct {
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

// MapSandboxImagesDeleteOutputFromJSON deserializes JSON data into a SandboxImagesDeleteOutput.
func MapSandboxImagesDeleteOutputFromJSON(data []byte) (*SandboxImagesDeleteOutput, error) {
	var v SandboxImagesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesDeleteOutputToJSON serializes a SandboxImagesDeleteOutput to JSON.
func MapSandboxImagesDeleteOutputToJSON(v *SandboxImagesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
