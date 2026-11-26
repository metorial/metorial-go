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
	// Outcome - `active` while the operation runs. `superseded` when a newer operation replaced it.
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
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxMachinesOperationsListOutput represents the sandbox machines operations list output type.
type SandboxMachinesOperationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                        `json:"object"`
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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
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
