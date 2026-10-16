package disks

import (
	"encoding/json"
	"time"
)

// SandboxDisksGetOutputMachine - The machine the disk is attached to, if any
type SandboxDisksGetOutputMachine struct {
	// Id - The machine the disk is attached to
	Id string `json:"id"`
	// Name - The attached machine name
	Name *string `json:"name,omitempty"`
	// Slug - The attached machine slug
	Slug *string `json:"slug,omitempty"`
	// MountPath - Where the machine mounts the disk
	MountPath string `json:"mount_path"`
}

// SandboxDisksGetOutputActiveResize - The resize currently in progress, if any
type SandboxDisksGetOutputActiveResize struct {
	Id string `json:"id"`
	// Status - The status of the resize
	Status          string  `json:"status"`
	TargetSizeBytes float64 `json:"target_size_bytes"`
}

// SandboxDisksGetOutput represents the sandbox disks get output type.
type SandboxDisksGetOutput struct {
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
	Machine *SandboxDisksGetOutputMachine `json:"machine,omitempty"`
	// ActiveResize - The resize currently in progress, if any
	ActiveResize *SandboxDisksGetOutputActiveResize `json:"active_resize,omitempty"`
	CreatedAt    time.Time                          `json:"created_at"`
	UpdatedAt    time.Time                          `json:"updated_at"`
	DeletedAt    *time.Time                         `json:"deleted_at,omitempty"`
}

// MapSandboxDisksGetOutputFromJSON deserializes JSON data into a SandboxDisksGetOutput.
func MapSandboxDisksGetOutputFromJSON(data []byte) (*SandboxDisksGetOutput, error) {
	var v SandboxDisksGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksGetOutputToJSON serializes a SandboxDisksGetOutput to JSON.
func MapSandboxDisksGetOutputToJSON(v *SandboxDisksGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
