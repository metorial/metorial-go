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
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxDisksMountsListOutput represents the sandbox disks mounts list output type.
type SandboxDisksMountsListOutput struct {
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
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
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
