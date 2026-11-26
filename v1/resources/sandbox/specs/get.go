package specs

import (
	"encoding/json"
	"time"
)

// SandboxSpecsGetOutput represents the sandbox specs get output type.
type SandboxSpecsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The spec's unique identifier
	Id string `json:"id"`
	// Identifier - Stable size label shared across regions.
	Identifier string `json:"identifier"`
	// MemoryGib - The memory of machines using this spec, in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The number of CPUs of machines using this spec
	Cpus float64 `json:"cpus"`
	// Available - Whether your organization's plan allows creating machines with this spec
	Available bool `json:"available"`
	// Regions - The identifiers of the regions this spec is available in
	Regions []string `json:"regions"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// MapSandboxSpecsGetOutputFromJSON deserializes JSON data into a SandboxSpecsGetOutput.
func MapSandboxSpecsGetOutputFromJSON(data []byte) (*SandboxSpecsGetOutput, error) {
	var v SandboxSpecsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxSpecsGetOutputToJSON serializes a SandboxSpecsGetOutput to JSON.
func MapSandboxSpecsGetOutputToJSON(v *SandboxSpecsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
