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
	// Name - Display name.
	Name string                            `json:"name"`
	Area SandboxRegionsListOutputItemsArea `json:"area"`
	// IsDefault - The default region for new machines in this Metorial region.
	IsDefault bool `json:"is_default"`
	// Availability - Whether the region is currently available for new machines
	Availability string `json:"availability"`
}

// SandboxRegionsListOutputPagination represents the sandbox regions list output pagination type.
type SandboxRegionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxRegionsListOutput represents the sandbox regions list output type.
type SandboxRegionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
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
