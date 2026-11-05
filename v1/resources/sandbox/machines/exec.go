package machines

import (
	"encoding/json"
)

// SandboxMachinesExecOutput represents the sandbox machines exec output type.
type SandboxMachinesExecOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// ExitCode - The exit code of the command, or null if it was ended by a signal
	ExitCode *float64 `json:"exit_code,omitempty"`
	// Signal - The signal that ended the command, or null if it exited on its own
	Signal *string `json:"signal,omitempty"`
	// TimedOut - Whether the command was stopped for running longer than its timeout
	TimedOut bool `json:"timed_out"`
	// Stdout - The standard output of the command, decoded as UTF-8
	Stdout string `json:"stdout"`
	// Stderr - The standard error of the command, decoded as UTF-8
	Stderr string `json:"stderr"`
	// StdoutTruncated - Whether stdout was cut off after its first 1 MiB
	StdoutTruncated bool `json:"stdout_truncated"`
	// StderrTruncated - Whether stderr was cut off after its first 1 MiB
	StderrTruncated bool `json:"stderr_truncated"`
}

// MapSandboxMachinesExecOutputFromJSON deserializes JSON data into a SandboxMachinesExecOutput.
func MapSandboxMachinesExecOutputFromJSON(data []byte) (*SandboxMachinesExecOutput, error) {
	var v SandboxMachinesExecOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesExecOutputToJSON serializes a SandboxMachinesExecOutput to JSON.
func MapSandboxMachinesExecOutputToJSON(v *SandboxMachinesExecOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesExecBody represents the sandbox machines exec body type.
type SandboxMachinesExecBody struct {
	// Command - The command to run with /bin/sh -c
	Command string `json:"command"`
	// Cwd - The working directory to run the command in
	Cwd *string `json:"cwd,omitempty"`
	// Env - Environment variables to set for the command
	Env *map[string]string `json:"env,omitempty"`
	// TimeoutSeconds - How long the command may run before it is stopped, in seconds. Defaults to 60, at most 300.
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty"`
}

// MapSandboxMachinesExecBodyFromJSON deserializes JSON data into a SandboxMachinesExecBody.
func MapSandboxMachinesExecBodyFromJSON(data []byte) (*SandboxMachinesExecBody, error) {
	var v SandboxMachinesExecBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesExecBodyToJSON serializes a SandboxMachinesExecBody to JSON.
func MapSandboxMachinesExecBodyToJSON(v *SandboxMachinesExecBody) ([]byte, error) {
	return json.Marshal(v)
}
