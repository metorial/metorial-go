package operations

import (
	"encoding/json"
	"time"
)

// SandboxMachinesOperationsListOutputItemsFailure represents the sandbox machines operations list output items failure type.
type SandboxMachinesOperationsListOutputItemsFailure struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesOperationsListOutputItemsInstanceOperations represents the sandbox machines operations list output items instance operations type.
type SandboxMachinesOperationsListOutputItemsInstanceOperations struct {
	Id                string `json:"id"`
	Kind              string `json:"kind"`
	MachineInstanceId string `json:"machine_instance_id"`
}

// SandboxMachinesOperationsListOutputItems represents the sandbox machines operations list output items type.
type SandboxMachinesOperationsListOutputItems struct {
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
	Outcome string                                           `json:"outcome"`
	Failure *SandboxMachinesOperationsListOutputItemsFailure `json:"failure,omitempty"`
	// FromMachineInstanceId - The instance the machine was running on when the operation began
	FromMachineInstanceId *string `json:"from_machine_instance_id,omitempty"`
	// SupersededById - The operation that replaced this one
	SupersededById *string `json:"superseded_by_id,omitempty"`
	// InstanceOperations - The per-instance steps the operation consisted of
	InstanceOperations []SandboxMachinesOperationsListOutputItemsInstanceOperations `json:"instance_operations"`
	RequestedAt        time.Time                                                    `json:"requested_at"`
	CompletedAt        *time.Time                                                   `json:"completed_at,omitempty"`
}

// SandboxMachinesOperationsListOutputPagination represents the sandbox machines operations list output pagination type.
type SandboxMachinesOperationsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxMachinesOperationsListOutput represents the sandbox machines operations list output type.
type SandboxMachinesOperationsListOutput struct {
	Items      []SandboxMachinesOperationsListOutputItems    `json:"items"`
	Pagination SandboxMachinesOperationsListOutputPagination `json:"pagination"`
}

// MapSandboxMachinesOperationsListOutputFromJSON deserializes JSON data into a SandboxMachinesOperationsListOutput.
func MapSandboxMachinesOperationsListOutputFromJSON(data []byte) (*SandboxMachinesOperationsListOutput, error) {
	var v SandboxMachinesOperationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesOperationsListOutputToJSON serializes a SandboxMachinesOperationsListOutput to JSON.
func MapSandboxMachinesOperationsListOutputToJSON(v *SandboxMachinesOperationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesOperationsListQuery represents the sandbox machines operations list query type.
type SandboxMachinesOperationsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// MachineInstanceId - Only operations that involved these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	Kind              *any `json:"kind,omitempty"`
	Trigger           *any `json:"trigger,omitempty"`
	Outcome           *any `json:"outcome,omitempty"`
}

// MapSandboxMachinesOperationsListQueryFromJSON deserializes JSON data into a SandboxMachinesOperationsListQuery.
func MapSandboxMachinesOperationsListQueryFromJSON(data []byte) (*SandboxMachinesOperationsListQuery, error) {
	var v SandboxMachinesOperationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesOperationsListQueryToJSON serializes a SandboxMachinesOperationsListQuery to JSON.
func MapSandboxMachinesOperationsListQueryToJSON(v *SandboxMachinesOperationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
