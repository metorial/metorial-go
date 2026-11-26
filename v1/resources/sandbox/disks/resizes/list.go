package resizes

import (
	"encoding/json"
	"time"
)

// SandboxDisksResizesListOutputItemsFailure represents the sandbox disks resizes list output items failure type.
type SandboxDisksResizesListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxDisksResizesListOutputItems represents the sandbox disks resizes list output items type.
type SandboxDisksResizesListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The resize's unique identifier
	Id string `json:"id"`
	// DiskId - The disk being resized
	DiskId string `json:"disk_id"`
	// Status - The resize's status
	Status string `json:"status"`
	// OriginalSizeBytes - The disk size before the resize
	OriginalSizeBytes float64 `json:"original_size_bytes"`
	// TargetSizeBytes - The disk size the resize works toward
	TargetSizeBytes float64 `json:"target_size_bytes"`
	// MachineIds - The machines that are stopped and restarted for the resize
	MachineIds  []string                                   `json:"machine_ids"`
	Failure     *SandboxDisksResizesListOutputItemsFailure `json:"failure,omitempty"`
	CreatedAt   time.Time                                  `json:"created_at"`
	UpdatedAt   time.Time                                  `json:"updated_at"`
	CompletedAt *time.Time                                 `json:"completed_at,omitempty"`
}

// SandboxDisksResizesListOutputPagination represents the sandbox disks resizes list output pagination type.
type SandboxDisksResizesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxDisksResizesListOutput represents the sandbox disks resizes list output type.
type SandboxDisksResizesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []SandboxDisksResizesListOutputItems    `json:"items"`
	Pagination SandboxDisksResizesListOutputPagination `json:"pagination"`
}

// MapSandboxDisksResizesListOutputFromJSON deserializes JSON data into a SandboxDisksResizesListOutput.
func MapSandboxDisksResizesListOutputFromJSON(data []byte) (*SandboxDisksResizesListOutput, error) {
	var v SandboxDisksResizesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksResizesListOutputToJSON serializes a SandboxDisksResizesListOutput to JSON.
func MapSandboxDisksResizesListOutputToJSON(v *SandboxDisksResizesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksResizesListQuery represents the sandbox disks resizes list query type.
type SandboxDisksResizesListQuery struct {
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
	// Status - Filter by resize status
	Status *any `json:"status,omitempty"`
}

// MapSandboxDisksResizesListQueryFromJSON deserializes JSON data into a SandboxDisksResizesListQuery.
func MapSandboxDisksResizesListQueryFromJSON(data []byte) (*SandboxDisksResizesListQuery, error) {
	var v SandboxDisksResizesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksResizesListQueryToJSON serializes a SandboxDisksResizesListQuery to JSON.
func MapSandboxDisksResizesListQueryToJSON(v *SandboxDisksResizesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
