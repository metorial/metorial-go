package versions

import (
	"encoding/json"
	"time"
)

// SandboxImagesVersionsGetOutput represents the sandbox images versions get output type.
type SandboxImagesVersionsGetOutput struct {
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

// MapSandboxImagesVersionsGetOutputFromJSON deserializes JSON data into a SandboxImagesVersionsGetOutput.
func MapSandboxImagesVersionsGetOutputFromJSON(data []byte) (*SandboxImagesVersionsGetOutput, error) {
	var v SandboxImagesVersionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImagesVersionsGetOutputToJSON serializes a SandboxImagesVersionsGetOutput to JSON.
func MapSandboxImagesVersionsGetOutputToJSON(v *SandboxImagesVersionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
