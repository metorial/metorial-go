package imagebuilds

import (
	"encoding/json"
	"time"
)

// SandboxImageBuildsLogsOutputLines represents the sandbox image builds logs output lines type.
type SandboxImageBuildsLogsOutputLines struct {
	// Index - The line's position in the log
	Index float64 `json:"index"`
	// Source - Where the line was written
	Source string `json:"source"`
	// Timestamp - When the line was written
	Timestamp time.Time `json:"timestamp"`
	// Text - The line
	Text string `json:"text"`
}

// SandboxImageBuildsLogsOutput represents the sandbox image builds logs output type.
type SandboxImageBuildsLogsOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Lines - The requested lines
	Lines []SandboxImageBuildsLogsOutputLines `json:"lines"`
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

// MapSandboxImageBuildsLogsOutputFromJSON deserializes JSON data into a SandboxImageBuildsLogsOutput.
func MapSandboxImageBuildsLogsOutputFromJSON(data []byte) (*SandboxImageBuildsLogsOutput, error) {
	var v SandboxImageBuildsLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImageBuildsLogsOutputToJSON serializes a SandboxImageBuildsLogsOutput to JSON.
func MapSandboxImageBuildsLogsOutputToJSON(v *SandboxImageBuildsLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxImageBuildsLogsQuery represents the sandbox image builds logs query type.
type SandboxImageBuildsLogsQuery struct {
	// Direction - Whether to read from the cursor toward newer or older lines
	Direction *string `json:"direction,omitempty"`
	// Cursor - The line index to start from
	Cursor *float64 `json:"cursor,omitempty"`
	// Limit - The maximum number of lines to return
	Limit *float64 `json:"limit,omitempty"`
}

// MapSandboxImageBuildsLogsQueryFromJSON deserializes JSON data into a SandboxImageBuildsLogsQuery.
func MapSandboxImageBuildsLogsQueryFromJSON(data []byte) (*SandboxImageBuildsLogsQuery, error) {
	var v SandboxImageBuildsLogsQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxImageBuildsLogsQueryToJSON serializes a SandboxImageBuildsLogsQuery to JSON.
func MapSandboxImageBuildsLogsQueryToJSON(v *SandboxImageBuildsLogsQuery) ([]byte, error) {
	return json.Marshal(v)
}
