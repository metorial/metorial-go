package disks

import (
	"encoding/json"
	"time"
)

// SandboxDisksUpdateOutputMachine - The machine the disk is attached to, if any
type SandboxDisksUpdateOutputMachine struct {
	// Id - The machine the disk is attached to
	Id string `json:"id"`
	// Name - The attached machine name
	Name *string `json:"name,omitempty"`
	// Slug - The attached machine slug
	Slug *string `json:"slug,omitempty"`
	// MountPath - Where the machine mounts the disk
	MountPath string `json:"mount_path"`
}

// SandboxDisksUpdateOutputActiveResize - The resize currently in progress, if any
type SandboxDisksUpdateOutputActiveResize struct {
	Id string `json:"id"`
	// Status - The status of the resize
	Status          string  `json:"status"`
	TargetSizeBytes float64 `json:"target_size_bytes"`
}

// SandboxDisksUpdateOutput represents the sandbox disks update output type.
type SandboxDisksUpdateOutput struct {
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
	Machine *SandboxDisksUpdateOutputMachine `json:"machine,omitempty"`
	// ActiveResize - The resize currently in progress, if any
	ActiveResize *SandboxDisksUpdateOutputActiveResize `json:"active_resize,omitempty"`
	CreatedAt    time.Time                             `json:"created_at"`
	UpdatedAt    time.Time                             `json:"updated_at"`
	DeletedAt    *time.Time                            `json:"deleted_at,omitempty"`
}

// MapSandboxDisksUpdateOutputFromJSON deserializes JSON data into a SandboxDisksUpdateOutput.
func MapSandboxDisksUpdateOutputFromJSON(data []byte) (*SandboxDisksUpdateOutput, error) {
	var v SandboxDisksUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksUpdateOutputToJSON serializes a SandboxDisksUpdateOutput to JSON.
func MapSandboxDisksUpdateOutputToJSON(v *SandboxDisksUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksUpdateBody represents the sandbox disks update body type.
type SandboxDisksUpdateBody struct {
	// Metadata - User-defined metadata stored in Metorial. Replaces existing metadata when updated.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Name - A name for the disk. Pass null to remove it.
	Name *string `json:"name,omitempty"`
}

// MapSandboxDisksUpdateBodyFromJSON deserializes JSON data into a SandboxDisksUpdateBody.
func MapSandboxDisksUpdateBodyFromJSON(data []byte) (*SandboxDisksUpdateBody, error) {
	var v SandboxDisksUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksUpdateBodyToJSON serializes a SandboxDisksUpdateBody to JSON.
func MapSandboxDisksUpdateBodyToJSON(v *SandboxDisksUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
