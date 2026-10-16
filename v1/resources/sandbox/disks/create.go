package disks

import (
	"encoding/json"
	"time"
)

// SandboxDisksCreateOutputMachine - The machine the disk is attached to, if any
type SandboxDisksCreateOutputMachine struct {
	// Id - The machine the disk is attached to
	Id string `json:"id"`
	// Name - The attached machine name
	Name *string `json:"name,omitempty"`
	// Slug - The attached machine slug
	Slug *string `json:"slug,omitempty"`
	// MountPath - Where the machine mounts the disk
	MountPath string `json:"mount_path"`
}

// SandboxDisksCreateOutputActiveResize - The resize currently in progress, if any
type SandboxDisksCreateOutputActiveResize struct {
	Id string `json:"id"`
	// Status - The status of the resize
	Status          string  `json:"status"`
	TargetSizeBytes float64 `json:"target_size_bytes"`
}

// SandboxDisksCreateOutput represents the sandbox disks create output type.
type SandboxDisksCreateOutput struct {
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
	Machine *SandboxDisksCreateOutputMachine `json:"machine,omitempty"`
	// ActiveResize - The resize currently in progress, if any
	ActiveResize *SandboxDisksCreateOutputActiveResize `json:"active_resize,omitempty"`
	CreatedAt    time.Time                             `json:"created_at"`
	UpdatedAt    time.Time                             `json:"updated_at"`
	DeletedAt    *time.Time                            `json:"deleted_at,omitempty"`
}

// MapSandboxDisksCreateOutputFromJSON deserializes JSON data into a SandboxDisksCreateOutput.
func MapSandboxDisksCreateOutputFromJSON(data []byte) (*SandboxDisksCreateOutput, error) {
	var v SandboxDisksCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksCreateOutputToJSON serializes a SandboxDisksCreateOutput to JSON.
func MapSandboxDisksCreateOutputToJSON(v *SandboxDisksCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksCreateBody represents the sandbox disks create body type.
type SandboxDisksCreateBody struct {
	// Region - The identifier of the region to create the disk in
	Region string `json:"region"`
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the disk
	Name *string `json:"name,omitempty"`
	// SizeBytes - The capacity of the disk in bytes
	SizeBytes *float64 `json:"size_bytes,omitempty"`
}

// MapSandboxDisksCreateBodyFromJSON deserializes JSON data into a SandboxDisksCreateBody.
func MapSandboxDisksCreateBodyFromJSON(data []byte) (*SandboxDisksCreateBody, error) {
	var v SandboxDisksCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksCreateBodyToJSON serializes a SandboxDisksCreateBody to JSON.
func MapSandboxDisksCreateBodyToJSON(v *SandboxDisksCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
