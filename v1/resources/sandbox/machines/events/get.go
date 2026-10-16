package events

import (
	"encoding/json"
	"time"
)

// SandboxMachinesEventsGetOutputMachine represents the sandbox machines events get output machine type.
type SandboxMachinesEventsGetOutputMachine struct {
	FromStatus *string `json:"from_status,omitempty"`
	ToStatus   string  `json:"to_status"`
}

// SandboxMachinesEventsGetOutputInstance represents the sandbox machines events get output instance type.
type SandboxMachinesEventsGetOutputInstance struct {
	FromStatus *string `json:"from_status,omitempty"`
	ToStatus   string  `json:"to_status"`
}

// SandboxMachinesEventsGetOutputError represents the sandbox machines events get output error type.
type SandboxMachinesEventsGetOutputError struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesEventsGetOutput represents the sandbox machines events get output type.
type SandboxMachinesEventsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The event's unique identifier
	Id string `json:"id"`
	// MachineId - The machine the event happened on
	MachineId string `json:"machine_id"`
	// MachineInstanceId - The instance the event happened on, for instance-level events
	MachineInstanceId *string `json:"machine_instance_id,omitempty"`
	// Generation - The generation of the instance the event happened on
	Generation *float64 `json:"generation,omitempty"`
	// Type - The event's type
	Type string `json:"type"`
	// Source - What caused the event
	Source   string                                  `json:"source"`
	Machine  *SandboxMachinesEventsGetOutputMachine  `json:"machine,omitempty"`
	Instance *SandboxMachinesEventsGetOutputInstance `json:"instance,omitempty"`
	// OperationId - The operation the event belongs to
	OperationId *string `json:"operation_id,omitempty"`
	// GuestErrorId - The guest error recorded with the event
	GuestErrorId *string                              `json:"guest_error_id,omitempty"`
	Error        *SandboxMachinesEventsGetOutputError `json:"error,omitempty"`
	OccurredAt   time.Time                            `json:"occurred_at"`
}

// MapSandboxMachinesEventsGetOutputFromJSON deserializes JSON data into a SandboxMachinesEventsGetOutput.
func MapSandboxMachinesEventsGetOutputFromJSON(data []byte) (*SandboxMachinesEventsGetOutput, error) {
	var v SandboxMachinesEventsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesEventsGetOutputToJSON serializes a SandboxMachinesEventsGetOutput to JSON.
func MapSandboxMachinesEventsGetOutputToJSON(v *SandboxMachinesEventsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
