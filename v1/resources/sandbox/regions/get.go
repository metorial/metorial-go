package regions

import (
	"encoding/json"
)

// SandboxRegionsGetOutputArea represents the sandbox regions get output area type.
type SandboxRegionsGetOutputArea struct {
	// Id - The geographic area's identifier
	Id string `json:"id"`
	// Name - The geographic area's name
	Name string `json:"name"`
}

// SandboxRegionsGetOutput represents the sandbox regions get output type.
type SandboxRegionsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The region's identifier, used wherever a region is requested
	Id string `json:"id"`
	// Name - The region's display name
	Name string                      `json:"name"`
	Area SandboxRegionsGetOutputArea `json:"area"`
	// IsDefault - Whether this is the default sandbox region for this Metorial region
	IsDefault bool `json:"is_default"`
	// Availability - Whether the region is currently available for new machines
	Availability string `json:"availability"`
}

// MapSandboxRegionsGetOutputFromJSON deserializes JSON data into a SandboxRegionsGetOutput.
func MapSandboxRegionsGetOutputFromJSON(data []byte) (*SandboxRegionsGetOutput, error) {
	var v SandboxRegionsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxRegionsGetOutputToJSON serializes a SandboxRegionsGetOutput to JSON.
func MapSandboxRegionsGetOutputToJSON(v *SandboxRegionsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
