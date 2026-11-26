package errors

import (
	"encoding/json"
	"time"
)

// SandboxMachinesErrorsListOutputItems represents the sandbox machines errors list output items type.
type SandboxMachinesErrorsListOutputItems struct {
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

// SandboxMachinesErrorsListOutputPagination represents the sandbox machines errors list output pagination type.
type SandboxMachinesErrorsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// SandboxMachinesErrorsListOutput represents the sandbox machines errors list output type.
type SandboxMachinesErrorsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []SandboxMachinesErrorsListOutputItems    `json:"items"`
	Pagination SandboxMachinesErrorsListOutputPagination `json:"pagination"`
}

// MapSandboxMachinesErrorsListOutputFromJSON deserializes JSON data into a SandboxMachinesErrorsListOutput.
func MapSandboxMachinesErrorsListOutputFromJSON(data []byte) (*SandboxMachinesErrorsListOutput, error) {
	var v SandboxMachinesErrorsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesErrorsListOutputToJSON serializes a SandboxMachinesErrorsListOutput to JSON.
func MapSandboxMachinesErrorsListOutputToJSON(v *SandboxMachinesErrorsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesErrorsListQuery represents the sandbox machines errors list query type.
type SandboxMachinesErrorsListQuery struct {
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
	// MachineInstanceId - Only errors of these instances
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
}

// MapSandboxMachinesErrorsListQueryFromJSON deserializes JSON data into a SandboxMachinesErrorsListQuery.
func MapSandboxMachinesErrorsListQueryFromJSON(data []byte) (*SandboxMachinesErrorsListQuery, error) {
	var v SandboxMachinesErrorsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesErrorsListQueryToJSON serializes a SandboxMachinesErrorsListQuery to JSON.
func MapSandboxMachinesErrorsListQueryToJSON(v *SandboxMachinesErrorsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
