package events

import (
	"encoding/json"
	"time"
)

// SandboxMachinesEventsListOutputItemsMachine represents the sandbox machines events list output items machine type.
type SandboxMachinesEventsListOutputItemsMachine struct {
	FromStatus *string `json:"from_status,omitempty"`
	ToStatus   string  `json:"to_status"`
}

// SandboxMachinesEventsListOutputItemsInstance represents the sandbox machines events list output items instance type.
type SandboxMachinesEventsListOutputItemsInstance struct {
	FromStatus *string `json:"from_status,omitempty"`
	ToStatus   string  `json:"to_status"`
}

// SandboxMachinesEventsListOutputItemsError represents the sandbox machines events list output items error type.
type SandboxMachinesEventsListOutputItemsError struct {
	// Code - A stable, machine-readable code describing what failed
	Code string `json:"code"`
	// Message - A short, human-readable description of the failure
	Message string `json:"message"`
}

// SandboxMachinesEventsListOutputItems represents the sandbox machines events list output items type.
type SandboxMachinesEventsListOutputItems struct {
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
	Source   string                                        `json:"source"`
	Machine  *SandboxMachinesEventsListOutputItemsMachine  `json:"machine,omitempty"`
	Instance *SandboxMachinesEventsListOutputItemsInstance `json:"instance,omitempty"`
	// OperationId - The operation the event belongs to
	OperationId *string `json:"operation_id,omitempty"`
	// GuestErrorId - The guest error recorded with the event
	GuestErrorId *string                                    `json:"guest_error_id,omitempty"`
	Error        *SandboxMachinesEventsListOutputItemsError `json:"error,omitempty"`
	OccurredAt   time.Time                                  `json:"occurred_at"`
}

// SandboxMachinesEventsListOutputPagination represents the sandbox machines events list output pagination type.
type SandboxMachinesEventsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxMachinesEventsListOutput represents the sandbox machines events list output type.
type SandboxMachinesEventsListOutput struct {
	Items      []SandboxMachinesEventsListOutputItems    `json:"items"`
	Pagination SandboxMachinesEventsListOutputPagination `json:"pagination"`
}

// MapSandboxMachinesEventsListOutputFromJSON deserializes JSON data into a SandboxMachinesEventsListOutput.
func MapSandboxMachinesEventsListOutputFromJSON(data []byte) (*SandboxMachinesEventsListOutput, error) {
	var v SandboxMachinesEventsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesEventsListOutputToJSON serializes a SandboxMachinesEventsListOutput to JSON.
func MapSandboxMachinesEventsListOutputToJSON(v *SandboxMachinesEventsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesEventsListQuery represents the sandbox machines events list query type.
type SandboxMachinesEventsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Scope - Whether to include machine-level events, instance-level events, or both
	Scope *string `json:"scope,omitempty"`
	// MachineInstanceId - Only events of these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	// Type - Filter by event type
	Type *any `json:"type,omitempty"`
	// Source - Filter by what caused the event
	Source *any `json:"source,omitempty"`
	// Reverse - List the oldest events first
	Reverse *bool `json:"reverse,omitempty"`
}

// MapSandboxMachinesEventsListQueryFromJSON deserializes JSON data into a SandboxMachinesEventsListQuery.
func MapSandboxMachinesEventsListQueryFromJSON(data []byte) (*SandboxMachinesEventsListQuery, error) {
	var v SandboxMachinesEventsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesEventsListQueryToJSON serializes a SandboxMachinesEventsListQuery to JSON.
func MapSandboxMachinesEventsListQueryToJSON(v *SandboxMachinesEventsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
