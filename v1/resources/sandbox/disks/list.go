package disks

import (
	"encoding/json"
	"time"
)

// SandboxDisksListOutputItemsMachine - The machine the disk is attached to, if any
type SandboxDisksListOutputItemsMachine struct {
	// Id - The machine the disk is attached to
	Id string `json:"id"`
	// Name - The attached machine name
	Name *string `json:"name,omitempty"`
	// Slug - The attached machine slug
	Slug *string `json:"slug,omitempty"`
	// MountPath - Where the machine mounts the disk
	MountPath string `json:"mount_path"`
}

// SandboxDisksListOutputItemsActiveResize - The resize currently in progress, if any
type SandboxDisksListOutputItemsActiveResize struct {
	Id string `json:"id"`
	// Status - The status of the resize
	Status          string  `json:"status"`
	TargetSizeBytes float64 `json:"target_size_bytes"`
}

// SandboxDisksListOutputItems represents the sandbox disks list output items type.
type SandboxDisksListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Metadata - User-defined metadata stored in Metorial
	Metadata map[string]any `json:"metadata"`
	// Id - The disk's unique identifier
	Id string `json:"id"`
	// Type - Standard disks are managed by you. Root filesystem disks belong to a machine.
	Type string `json:"type"`
	// Name - The disk's name
	Name *string `json:"name,omitempty"`
	// Status - The disk's status. A disk that is still being cloned reports cloning until the copy completes.
	Status string `json:"status"`
	// Region - The identifier of the region the disk lives in
	Region string `json:"region"`
	// SizeBytes - The disk's capacity in bytes. Null only for a disk deleted before its size was recorded.
	SizeBytes *float64 `json:"size_bytes,omitempty"`
	// Machine - The machine the disk is attached to, if any
	Machine *SandboxDisksListOutputItemsMachine `json:"machine,omitempty"`
	// ActiveResize - The resize currently in progress, if any
	ActiveResize *SandboxDisksListOutputItemsActiveResize `json:"active_resize,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt - When this was deleted. `null` while active.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// SandboxDisksListOutputPagination represents the sandbox disks list output pagination type.
type SandboxDisksListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxDisksListOutput represents the sandbox disks list output type.
type SandboxDisksListOutput struct {
	// Object - String representing the object’s type.
	Object     string                           `json:"object"`
	Items      []SandboxDisksListOutputItems    `json:"items"`
	Pagination SandboxDisksListOutputPagination `json:"pagination"`
}

// MapSandboxDisksListOutputFromJSON deserializes JSON data into a SandboxDisksListOutput.
func MapSandboxDisksListOutputFromJSON(data []byte) (*SandboxDisksListOutput, error) {
	var v SandboxDisksListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksListOutputToJSON serializes a SandboxDisksListOutput to JSON.
func MapSandboxDisksListOutputToJSON(v *SandboxDisksListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksListQuery represents the sandbox disks list query type.
type SandboxDisksListQuery struct {
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
	// Id - Filter by resource IDs
	Id *any `json:"id,omitempty"`
	// Type - Filter by disk type
	Type *any `json:"type,omitempty"`
	// MachineId - Filter by the machines the disks are mounted into
	MachineId *any `json:"machine_id,omitempty"`
	// Region - Filter by the identifiers of the regions the disks live in
	Region *any `json:"region,omitempty"`
	// Attached - Only disks used by an active machine (true), as its root filesystem or a mount, or only disks no machine uses (false)
	Attached *bool `json:"attached,omitempty"`
	// Search - Search disks by name or ID
	Search *string `json:"search,omitempty"`
}

// MapSandboxDisksListQueryFromJSON deserializes JSON data into a SandboxDisksListQuery.
func MapSandboxDisksListQueryFromJSON(data []byte) (*SandboxDisksListQuery, error) {
	var v SandboxDisksListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksListQueryToJSON serializes a SandboxDisksListQuery to JSON.
func MapSandboxDisksListQueryToJSON(v *SandboxDisksListQuery) ([]byte, error) {
	return json.Marshal(v)
}
