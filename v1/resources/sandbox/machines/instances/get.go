package instances

import (
	"encoding/json"
	"time"
)

// SandboxMachinesInstancesGetOutputImage represents the sandbox machines instances get output image type.
type SandboxMachinesInstancesGetOutputImage struct {
	// Id - The image the instance booted from
	Id *string `json:"id,omitempty"`
	// VersionId - The image version the instance booted from
	VersionId *string `json:"version_id,omitempty"`
}

// SandboxMachinesInstancesGetOutputSpec represents the sandbox machines instances get output spec type.
type SandboxMachinesInstancesGetOutputSpec struct {
	// Id - The spec the instance was created with
	Id *string `json:"id,omitempty"`
	// Identifier - The spec's human-readable identifier
	Identifier string `json:"identifier"`
	// MemoryGib - The instance's memory in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The instance's CPU count
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesInstancesGetOutputOperation - The operation currently acting on the instance, if any
type SandboxMachinesInstancesGetOutputOperation struct {
	Id          string    `json:"id"`
	Kind        string    `json:"kind"`
	RequestedAt time.Time `json:"requested_at"`
}

// SandboxMachinesInstancesGetOutputMounts represents the sandbox machines instances get output mounts type.
type SandboxMachinesInstancesGetOutputMounts struct {
	DiskId    string `json:"disk_id"`
	MountPath string `json:"mount_path"`
	// Status - The mount status
	Status string `json:"status"`
}

// SandboxMachinesInstancesGetOutputRuntimeSetup represents the sandbox machines instances get output runtime setup type.
type SandboxMachinesInstancesGetOutputRuntimeSetup struct {
	Id         string  `json:"id"`
	ExitCode   float64 `json:"exit_code"`
	DurationMs float64 `json:"duration_ms"`
}

// SandboxMachinesInstancesGetOutputRuntimeEntrypoint represents the sandbox machines instances get output runtime entrypoint type.
type SandboxMachinesInstancesGetOutputRuntimeEntrypoint struct {
	StartedAfterMs *float64 `json:"started_after_ms,omitempty"`
	ExitCode       *float64 `json:"exit_code,omitempty"`
	Signal         *string  `json:"signal,omitempty"`
}

// SandboxMachinesInstancesGetOutputRuntimeError - The error reported by the workload runtime, if any
type SandboxMachinesInstancesGetOutputRuntimeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SandboxMachinesInstancesGetOutputRuntime - What the workload runtime inside the instance reported
type SandboxMachinesInstancesGetOutputRuntime struct {
	// State - The state of the workload runtime
	State      string                                              `json:"state"`
	Setup      []SandboxMachinesInstancesGetOutputRuntimeSetup     `json:"setup"`
	Entrypoint *SandboxMachinesInstancesGetOutputRuntimeEntrypoint `json:"entrypoint,omitempty"`
	// Error - The error reported by the workload runtime, if any
	Error *SandboxMachinesInstancesGetOutputRuntimeError `json:"error,omitempty"`
}

// SandboxMachinesInstancesGetOutputFailure represents the sandbox machines instances get output failure type.
type SandboxMachinesInstancesGetOutputFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesInstancesGetOutput represents the sandbox machines instances get output type.
type SandboxMachinesInstancesGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The instance's unique identifier
	Id string `json:"id"`
	// MachineId - The machine the instance belongs to
	MachineId string `json:"machine_id"`
	// Generation - The instance number within its machine, counting from one
	Generation float64 `json:"generation"`
	// Status - The instance's detailed lifecycle status
	Status string `json:"status"`
	// Phase - The instance's coarse lifecycle phase
	Phase string `json:"phase"`
	// CreationReason - What caused the instance to be created
	CreationReason *string                                `json:"creation_reason,omitempty"`
	Image          SandboxMachinesInstancesGetOutputImage `json:"image"`
	Spec           SandboxMachinesInstancesGetOutputSpec  `json:"spec"`
	// Operation - The operation currently acting on the instance, if any
	Operation *SandboxMachinesInstancesGetOutputOperation `json:"operation,omitempty"`
	// Mounts - The disks mounted into the instance
	Mounts []SandboxMachinesInstancesGetOutputMounts `json:"mounts"`
	// DisksReady - Whether every required disk is mounted
	DisksReady bool `json:"disks_ready"`
	// Runtime - What the workload runtime inside the instance reported
	Runtime *SandboxMachinesInstancesGetOutputRuntime `json:"runtime,omitempty"`
	Failure *SandboxMachinesInstancesGetOutputFailure `json:"failure,omitempty"`
	// HasLogs - Whether logs are available for the instance
	HasLogs      bool       `json:"has_logs"`
	CreatedAt    time.Time  `json:"created_at"`
	StartedAt    *time.Time `json:"started_at,omitempty"`
	ReadyAt      *time.Time `json:"ready_at,omitempty"`
	PausedAt     *time.Time `json:"paused_at,omitempty"`
	TerminatedAt *time.Time `json:"terminated_at,omitempty"`
	// LastSeenAt - When the instance last reported being alive
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
}

// MapSandboxMachinesInstancesGetOutputFromJSON deserializes JSON data into a SandboxMachinesInstancesGetOutput.
func MapSandboxMachinesInstancesGetOutputFromJSON(data []byte) (*SandboxMachinesInstancesGetOutput, error) {
	var v SandboxMachinesInstancesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesInstancesGetOutputToJSON serializes a SandboxMachinesInstancesGetOutput to JSON.
func MapSandboxMachinesInstancesGetOutputToJSON(v *SandboxMachinesInstancesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
