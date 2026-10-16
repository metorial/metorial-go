package instances

import (
	"encoding/json"
	"time"
)

// SandboxMachinesInstancesListOutputItemsImage represents the sandbox machines instances list output items image type.
type SandboxMachinesInstancesListOutputItemsImage struct {
	// Id - The image the instance booted from
	Id *string `json:"id,omitempty"`
	// VersionId - The image version the instance booted from
	VersionId *string `json:"version_id,omitempty"`
}

// SandboxMachinesInstancesListOutputItemsSpec represents the sandbox machines instances list output items spec type.
type SandboxMachinesInstancesListOutputItemsSpec struct {
	// Id - The spec the instance was created with
	Id *string `json:"id,omitempty"`
	// Identifier - The spec's human-readable identifier
	Identifier string `json:"identifier"`
	// MemoryGib - The instance's memory in GiB
	MemoryGib float64 `json:"memory_gib"`
	// Cpus - The instance's CPU count
	Cpus float64 `json:"cpus"`
}

// SandboxMachinesInstancesListOutputItemsOperation - The operation currently acting on the instance, if any
type SandboxMachinesInstancesListOutputItemsOperation struct {
	Id          string    `json:"id"`
	Kind        string    `json:"kind"`
	RequestedAt time.Time `json:"requested_at"`
}

// SandboxMachinesInstancesListOutputItemsMounts represents the sandbox machines instances list output items mounts type.
type SandboxMachinesInstancesListOutputItemsMounts struct {
	DiskId    string `json:"disk_id"`
	MountPath string `json:"mount_path"`
	// Status - The mount status
	Status string `json:"status"`
}

// SandboxMachinesInstancesListOutputItemsRuntimeSetup represents the sandbox machines instances list output items runtime setup type.
type SandboxMachinesInstancesListOutputItemsRuntimeSetup struct {
	Id         string  `json:"id"`
	ExitCode   float64 `json:"exit_code"`
	DurationMs float64 `json:"duration_ms"`
}

// SandboxMachinesInstancesListOutputItemsRuntimeEntrypoint represents the sandbox machines instances list output items runtime entrypoint type.
type SandboxMachinesInstancesListOutputItemsRuntimeEntrypoint struct {
	StartedAfterMs *float64 `json:"started_after_ms,omitempty"`
	ExitCode       *float64 `json:"exit_code,omitempty"`
	Signal         *string  `json:"signal,omitempty"`
}

// SandboxMachinesInstancesListOutputItemsRuntimeError - The error reported by the workload runtime, if any
type SandboxMachinesInstancesListOutputItemsRuntimeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SandboxMachinesInstancesListOutputItemsRuntime - What the workload runtime inside the instance reported
type SandboxMachinesInstancesListOutputItemsRuntime struct {
	// State - The state of the workload runtime
	State      string                                                    `json:"state"`
	Setup      []SandboxMachinesInstancesListOutputItemsRuntimeSetup     `json:"setup"`
	Entrypoint *SandboxMachinesInstancesListOutputItemsRuntimeEntrypoint `json:"entrypoint,omitempty"`
	// Error - The error reported by the workload runtime, if any
	Error *SandboxMachinesInstancesListOutputItemsRuntimeError `json:"error,omitempty"`
}

// SandboxMachinesInstancesListOutputItemsFailure represents the sandbox machines instances list output items failure type.
type SandboxMachinesInstancesListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesInstancesListOutputItems represents the sandbox machines instances list output items type.
type SandboxMachinesInstancesListOutputItems struct {
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
	CreationReason *string                                      `json:"creation_reason,omitempty"`
	Image          SandboxMachinesInstancesListOutputItemsImage `json:"image"`
	Spec           SandboxMachinesInstancesListOutputItemsSpec  `json:"spec"`
	// Operation - The operation currently acting on the instance, if any
	Operation *SandboxMachinesInstancesListOutputItemsOperation `json:"operation,omitempty"`
	// Mounts - The disks mounted into the instance
	Mounts []SandboxMachinesInstancesListOutputItemsMounts `json:"mounts"`
	// DisksReady - Whether every required disk is mounted
	DisksReady bool `json:"disks_ready"`
	// Runtime - What the workload runtime inside the instance reported
	Runtime *SandboxMachinesInstancesListOutputItemsRuntime `json:"runtime,omitempty"`
	Failure *SandboxMachinesInstancesListOutputItemsFailure `json:"failure,omitempty"`
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

// SandboxMachinesInstancesListOutputPagination represents the sandbox machines instances list output pagination type.
type SandboxMachinesInstancesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxMachinesInstancesListOutput represents the sandbox machines instances list output type.
type SandboxMachinesInstancesListOutput struct {
	Items      []SandboxMachinesInstancesListOutputItems    `json:"items"`
	Pagination SandboxMachinesInstancesListOutputPagination `json:"pagination"`
}

// MapSandboxMachinesInstancesListOutputFromJSON deserializes JSON data into a SandboxMachinesInstancesListOutput.
func MapSandboxMachinesInstancesListOutputFromJSON(data []byte) (*SandboxMachinesInstancesListOutput, error) {
	var v SandboxMachinesInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesInstancesListOutputToJSON serializes a SandboxMachinesInstancesListOutput to JSON.
func MapSandboxMachinesInstancesListOutputToJSON(v *SandboxMachinesInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesInstancesListQuery represents the sandbox machines instances list query type.
type SandboxMachinesInstancesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Status - Filter by instance status
	Status *any `json:"status,omitempty"`
}

// MapSandboxMachinesInstancesListQueryFromJSON deserializes JSON data into a SandboxMachinesInstancesListQuery.
func MapSandboxMachinesInstancesListQueryFromJSON(data []byte) (*SandboxMachinesInstancesListQuery, error) {
	var v SandboxMachinesInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesInstancesListQueryToJSON serializes a SandboxMachinesInstancesListQuery to JSON.
func MapSandboxMachinesInstancesListQueryToJSON(v *SandboxMachinesInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
