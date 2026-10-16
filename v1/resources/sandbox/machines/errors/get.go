package errors

import (
	"encoding/json"
	"time"
)

// SandboxMachinesErrorsGetOutput represents the sandbox machines errors get output type.
type SandboxMachinesErrorsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The error's unique identifier
	Id        string `json:"id"`
	MachineId string `json:"machine_id"`
	// MachineInstanceId - The instance the error occurred in
	MachineInstanceId string `json:"machine_instance_id"`
	// OperationId - The operation during which the error occurred
	OperationId string  `json:"operation_id"`
	EventId     *string `json:"event_id,omitempty"`
	// Kind - The kind of operation that was running
	Kind string `json:"kind"`
	// Stage - The stage of the operation at which the error occurred
	Stage string `json:"stage"`
	// Reason - Why the stage did not complete
	Reason string `json:"reason"`
	// Code - The error code reported from inside the machine
	Code string `json:"code"`
	// Message - The error message reported from inside the machine
	Message string `json:"message"`
	// ExitCode - The exit code of the failing process
	ExitCode *float64 `json:"exit_code,omitempty"`
	// DurationMs - How long the process ran
	DurationMs *float64 `json:"duration_ms,omitempty"`
	// Stdout - Standard output of the failing process
	Stdout     *string   `json:"stdout,omitempty"`
	OccurredAt time.Time `json:"occurred_at"`
}

// MapSandboxMachinesErrorsGetOutputFromJSON deserializes JSON data into a SandboxMachinesErrorsGetOutput.
func MapSandboxMachinesErrorsGetOutputFromJSON(data []byte) (*SandboxMachinesErrorsGetOutput, error) {
	var v SandboxMachinesErrorsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesErrorsGetOutputToJSON serializes a SandboxMachinesErrorsGetOutput to JSON.
func MapSandboxMachinesErrorsGetOutputToJSON(v *SandboxMachinesErrorsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
