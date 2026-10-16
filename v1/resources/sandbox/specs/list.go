package specs

import (
	"encoding/json"
	"time"
)

// SandboxSpecsListOutputItems represents the sandbox specs list output items type.
type SandboxSpecsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The spec's unique identifier
	Id string `json:"id"`
	// Identifier - The spec's human-readable identifier
	Identifier string `json:"identifier"`
	// MemoryGib - The memory of machines using this spec, in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The number of CPUs of machines using this spec
	Cpus float64 `json:"cpus"`
	// Available - Whether your organization's plan allows creating machines with this spec
	Available bool `json:"available"`
	// Regions - The identifiers of the regions this spec is available in
	Regions []string `json:"regions"`
	// CreatedAt - The spec's creation date
	CreatedAt time.Time `json:"created_at"`
}

// SandboxSpecsListOutputPagination represents the sandbox specs list output pagination type.
type SandboxSpecsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxSpecsListOutput represents the sandbox specs list output type.
type SandboxSpecsListOutput struct {
	Items      []SandboxSpecsListOutputItems    `json:"items"`
	Pagination SandboxSpecsListOutputPagination `json:"pagination"`
}

// MapSandboxSpecsListOutputFromJSON deserializes JSON data into a SandboxSpecsListOutput.
func MapSandboxSpecsListOutputFromJSON(data []byte) (*SandboxSpecsListOutput, error) {
	var v SandboxSpecsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxSpecsListOutputToJSON serializes a SandboxSpecsListOutput to JSON.
func MapSandboxSpecsListOutputToJSON(v *SandboxSpecsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxSpecsListQuery represents the sandbox specs list query type.
type SandboxSpecsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapSandboxSpecsListQueryFromJSON deserializes JSON data into a SandboxSpecsListQuery.
func MapSandboxSpecsListQueryFromJSON(data []byte) (*SandboxSpecsListQuery, error) {
	var v SandboxSpecsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxSpecsListQueryToJSON serializes a SandboxSpecsListQuery to JSON.
func MapSandboxSpecsListQueryToJSON(v *SandboxSpecsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
