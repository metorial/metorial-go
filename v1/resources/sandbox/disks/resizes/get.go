package resizes

import (
	"encoding/json"
	"time"
)

// SandboxDisksResizesGetOutputFailure represents the sandbox disks resizes get output failure type.
type SandboxDisksResizesGetOutputFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxDisksResizesGetOutput represents the sandbox disks resizes get output type.
type SandboxDisksResizesGetOutput struct {
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
	MachineIds  []string                             `json:"machine_ids"`
	Failure     *SandboxDisksResizesGetOutputFailure `json:"failure,omitempty"`
	CreatedAt   time.Time                            `json:"created_at"`
	UpdatedAt   time.Time                            `json:"updated_at"`
	CompletedAt *time.Time                           `json:"completed_at,omitempty"`
}

// MapSandboxDisksResizesGetOutputFromJSON deserializes JSON data into a SandboxDisksResizesGetOutput.
func MapSandboxDisksResizesGetOutputFromJSON(data []byte) (*SandboxDisksResizesGetOutput, error) {
	var v SandboxDisksResizesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxDisksResizesGetOutputToJSON serializes a SandboxDisksResizesGetOutput to JSON.
func MapSandboxDisksResizesGetOutputToJSON(v *SandboxDisksResizesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
