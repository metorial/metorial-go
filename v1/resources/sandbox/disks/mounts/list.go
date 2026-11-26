package mounts

import (
	"encoding/json"
	"time"
)

// SandboxDisksMountsListOutputItemsFailure represents the sandbox disks mounts list output items failure type.
type SandboxDisksMountsListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxDisksMountsListOutputItems represents the sandbox disks mounts list output items type.
type SandboxDisksMountsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The mount's unique identifier
	Id string `json:"id"`
	// DiskId - The mounted disk
	DiskId string `json:"disk_id"`
	// MachineId - The machine the disk was mounted for
	MachineId *string `json:"machine_id,omitempty"`
	// MachineInstanceId - The machine instance the disk was mounted into
	MachineInstanceId *string `json:"machine_instance_id,omitempty"`
	// MountPath - Where the disk was mounted
	MountPath string `json:"mount_path"`
	// Status - The mount's status
	Status     string                                    `json:"status"`
	Failure    *SandboxDisksMountsListOutputItemsFailure `json:"failure,omitempty"`
	MountedAt  *time.Time                                `json:"mounted_at,omitempty"`
	ReleasedAt *time.Time                                `json:"released_at,omitempty"`
	CreatedAt  time.Time                                 `json:"created_at"`
}

// SandboxDisksMountsListOutputPagination represents the sandbox disks mounts list output pagination type.
type SandboxDisksMountsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxDisksMountsListOutput represents the sandbox disks mounts list output type.
type SandboxDisksMountsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []SandboxDisksMountsListOutputItems    `json:"items"`
	Pagination SandboxDisksMountsListOutputPagination `json:"pagination"`
}

// MapSandboxDisksMountsListOutputFromJSON deserializes JSON data into a SandboxDisksMountsListOutput.
func MapSandboxDisksMountsListOutputFromJSON(data []byte) (*SandboxDisksMountsListOutput, error) {
	var v SandboxDisksMountsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksMountsListOutputToJSON serializes a SandboxDisksMountsListOutput to JSON.
func MapSandboxDisksMountsListOutputToJSON(v *SandboxDisksMountsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksMountsListQuery represents the sandbox disks mounts list query type.
type SandboxDisksMountsListQuery struct {
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

// MapSandboxDisksMountsListQueryFromJSON deserializes JSON data into a SandboxDisksMountsListQuery.
func MapSandboxDisksMountsListQueryFromJSON(data []byte) (*SandboxDisksMountsListQuery, error) {
	var v SandboxDisksMountsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksMountsListQueryToJSON serializes a SandboxDisksMountsListQuery to JSON.
func MapSandboxDisksMountsListQueryToJSON(v *SandboxDisksMountsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
