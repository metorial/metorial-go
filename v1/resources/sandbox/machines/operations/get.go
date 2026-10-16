package operations

import (
	"encoding/json"
	"time"
)

// SandboxMachinesOperationsGetOutputFailure represents the sandbox machines operations get output failure type.
type SandboxMachinesOperationsGetOutputFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesOperationsGetOutputInstanceOperations represents the sandbox machines operations get output instance operations type.
type SandboxMachinesOperationsGetOutputInstanceOperations struct {
	Id                string `json:"id"`
	Kind              string `json:"kind"`
	MachineInstanceId string `json:"machine_instance_id"`
}

// SandboxMachinesOperationsGetOutput represents the sandbox machines operations get output type.
type SandboxMachinesOperationsGetOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The operation's unique identifier
	Id string `json:"id"`
	// MachineId - The machine the operation acts on
	MachineId string `json:"machine_id"`
	// Kind - What the operation does to the machine. Update records a configuration change; upgrade covers the interruption while a disk of the machine is resized.
	Kind string `json:"kind"`
	// Trigger - What requested the operation
	Trigger string `json:"trigger"`
	// Outcome - The operation's outcome
	Outcome string                                     `json:"outcome"`
	Failure *SandboxMachinesOperationsGetOutputFailure `json:"failure,omitempty"`
	// FromMachineInstanceId - The instance the machine was running on when the operation began
	FromMachineInstanceId *string `json:"from_machine_instance_id,omitempty"`
	// SupersededById - The operation that replaced this one
	SupersededById *string `json:"superseded_by_id,omitempty"`
	// InstanceOperations - The per-instance steps the operation consisted of
	InstanceOperations []SandboxMachinesOperationsGetOutputInstanceOperations `json:"instance_operations"`
	RequestedAt        time.Time                                              `json:"requested_at"`
	CompletedAt        *time.Time                                             `json:"completed_at,omitempty"`
}

// MapSandboxMachinesOperationsGetOutputFromJSON deserializes JSON data into a SandboxMachinesOperationsGetOutput.
func MapSandboxMachinesOperationsGetOutputFromJSON(data []byte) (*SandboxMachinesOperationsGetOutput, error) {
	var v SandboxMachinesOperationsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesOperationsGetOutputToJSON serializes a SandboxMachinesOperationsGetOutput to JSON.
func MapSandboxMachinesOperationsGetOutputToJSON(v *SandboxMachinesOperationsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
