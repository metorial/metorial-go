package disks

import (
	"encoding/json"
	"time"
)

// SandboxDisksResizeOutputFailure represents the sandbox disks resize output failure type.
type SandboxDisksResizeOutputFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxDisksResizeOutput represents the sandbox disks resize output type.
type SandboxDisksResizeOutput struct {
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
	MachineIds  []string                         `json:"machine_ids"`
	Failure     *SandboxDisksResizeOutputFailure `json:"failure,omitempty"`
	CreatedAt   time.Time                        `json:"created_at"`
	UpdatedAt   time.Time                        `json:"updated_at"`
	CompletedAt *time.Time                       `json:"completed_at,omitempty"`
}

// MapSandboxDisksResizeOutputFromJSON deserializes JSON data into a SandboxDisksResizeOutput.
func MapSandboxDisksResizeOutputFromJSON(data []byte) (*SandboxDisksResizeOutput, error) {
	var v SandboxDisksResizeOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksResizeOutputToJSON serializes a SandboxDisksResizeOutput to JSON.
func MapSandboxDisksResizeOutputToJSON(v *SandboxDisksResizeOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxDisksResizeBody represents the sandbox disks resize body type.
type SandboxDisksResizeBody struct {
	// SizeBytes - The new capacity of the disk in bytes
	SizeBytes float64 `json:"size_bytes"`
}

// MapSandboxDisksResizeBodyFromJSON deserializes JSON data into a SandboxDisksResizeBody.
func MapSandboxDisksResizeBodyFromJSON(data []byte) (*SandboxDisksResizeBody, error) {
	var v SandboxDisksResizeBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksResizeBodyToJSON serializes a SandboxDisksResizeBody to JSON.
func MapSandboxDisksResizeBodyToJSON(v *SandboxDisksResizeBody) ([]byte, error) {
	return json.Marshal(v)
}
