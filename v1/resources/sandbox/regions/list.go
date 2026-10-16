package regions

import (
	"encoding/json"
)

// SandboxRegionsListOutputItemsArea represents the sandbox regions list output items area type.
type SandboxRegionsListOutputItemsArea struct {
	// Id - The geographic area's identifier
	Id string `json:"id"`
	// Name - The geographic area's name
	Name string `json:"name"`
}

// SandboxRegionsListOutputItems represents the sandbox regions list output items type.
type SandboxRegionsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The region's identifier, used wherever a region is requested
	Id string `json:"id"`
	// Name - The region's display name
	Name string                            `json:"name"`
	Area SandboxRegionsListOutputItemsArea `json:"area"`
	// IsDefault - Whether this is the default sandbox region for this Metorial region
	IsDefault bool `json:"is_default"`
	// Availability - Whether the region is currently available for new machines
	Availability string `json:"availability"`
}

// SandboxRegionsListOutputPagination represents the sandbox regions list output pagination type.
type SandboxRegionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxRegionsListOutput represents the sandbox regions list output type.
type SandboxRegionsListOutput struct {
	Items      []SandboxRegionsListOutputItems    `json:"items"`
	Pagination SandboxRegionsListOutputPagination `json:"pagination"`
}

// MapSandboxRegionsListOutputFromJSON deserializes JSON data into a SandboxRegionsListOutput.
func MapSandboxRegionsListOutputFromJSON(data []byte) (*SandboxRegionsListOutput, error) {
	var v SandboxRegionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxRegionsListOutputToJSON serializes a SandboxRegionsListOutput to JSON.
func MapSandboxRegionsListOutputToJSON(v *SandboxRegionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxRegionsListQuery represents the sandbox regions list query type.
type SandboxRegionsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapSandboxRegionsListQueryFromJSON deserializes JSON data into a SandboxRegionsListQuery.
func MapSandboxRegionsListQueryFromJSON(data []byte) (*SandboxRegionsListQuery, error) {
	var v SandboxRegionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxRegionsListQueryToJSON serializes a SandboxRegionsListQuery to JSON.
func MapSandboxRegionsListQueryToJSON(v *SandboxRegionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
