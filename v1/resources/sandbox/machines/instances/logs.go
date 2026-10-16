package instances

import (
	"encoding/json"
	"time"
)

// SandboxMachinesInstancesLogsOutputLines represents the sandbox machines instances logs output lines type.
type SandboxMachinesInstancesLogsOutputLines struct {
	// Index - The line's position in the log
	Index float64 `json:"index"`
	// Source - Where the line was written
	Source string `json:"source"`
	// Timestamp - When the line was written
	Timestamp time.Time `json:"timestamp"`
	// Text - The line
	Text string `json:"text"`
}

// SandboxMachinesInstancesLogsOutput represents the sandbox machines instances logs output type.
type SandboxMachinesInstancesLogsOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Lines - The requested lines
	Lines []SandboxMachinesInstancesLogsOutputLines `json:"lines"`
	// NextCursor - The cursor to pass to fetch the next page
	NextCursor *float64 `json:"next_cursor,omitempty"`
	// HasMore - Whether more lines are available
	HasMore bool `json:"has_more"`
	// LineCount - The total number of lines written so far
	LineCount float64 `json:"line_count"`
	// Status - Whether more lines can still be written
	Status string `json:"status"`
	// Truncated - Whether older lines were dropped to save space
	Truncated bool `json:"truncated"`
}

// MapSandboxMachinesInstancesLogsOutputFromJSON deserializes JSON data into a SandboxMachinesInstancesLogsOutput.
func MapSandboxMachinesInstancesLogsOutputFromJSON(data []byte) (*SandboxMachinesInstancesLogsOutput, error) {
	var v SandboxMachinesInstancesLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesInstancesLogsOutputToJSON serializes a SandboxMachinesInstancesLogsOutput to JSON.
func MapSandboxMachinesInstancesLogsOutputToJSON(v *SandboxMachinesInstancesLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesInstancesLogsQuery represents the sandbox machines instances logs query type.
type SandboxMachinesInstancesLogsQuery struct {
	// Direction - Whether to read from the cursor toward newer or older lines
	Direction *string `json:"direction,omitempty"`
	// Cursor - The line index to start from
	Cursor *float64 `json:"cursor,omitempty"`
	// Limit - The maximum number of lines to return
	Limit *float64 `json:"limit,omitempty"`
}

// MapSandboxMachinesInstancesLogsQueryFromJSON deserializes JSON data into a SandboxMachinesInstancesLogsQuery.
func MapSandboxMachinesInstancesLogsQueryFromJSON(data []byte) (*SandboxMachinesInstancesLogsQuery, error) {
	var v SandboxMachinesInstancesLogsQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesInstancesLogsQueryToJSON serializes a SandboxMachinesInstancesLogsQuery to JSON.
func MapSandboxMachinesInstancesLogsQueryToJSON(v *SandboxMachinesInstancesLogsQuery) ([]byte, error) {
	return json.Marshal(v)
}
