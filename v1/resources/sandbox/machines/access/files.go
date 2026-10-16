package access

import (
	"encoding/json"
	"time"
)

// SandboxMachinesAccessFilesOutputUrls represents the sandbox machines access files output urls type.
type SandboxMachinesAccessFilesOutputUrls struct {
	// Type - What the token grants access to
	Type string `json:"type"`
	// Shell - The WebSocket URL of the shell
	Shell *string `json:"shell,omitempty"`
	// Exec - The WebSocket URL for running commands. Shell tokens authorize it alongside the shell.
	Exec *string `json:"exec,omitempty"`
	// Files - The WebSocket URL of the file service
	Files *string `json:"files,omitempty"`
	// Ports - The URL of each requested port, keyed by port number
	Ports *map[string]string `json:"ports,omitempty"`
}

// SandboxMachinesAccessFilesOutput represents the sandbox machines access files output type.
type SandboxMachinesAccessFilesOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Token - The short-lived token that authorizes access. It is only shown once.
	Token string `json:"token"`
	// Scope - What the token grants access to
	Scope string                               `json:"scope"`
	Ports *any                                 `json:"ports,omitempty"`
	Urls  SandboxMachinesAccessFilesOutputUrls `json:"urls"`
	// ExpiresAt - When the token expires
	ExpiresAt time.Time `json:"expires_at"`
}

// MapSandboxMachinesAccessFilesOutputFromJSON deserializes JSON data into a SandboxMachinesAccessFilesOutput.
func MapSandboxMachinesAccessFilesOutputFromJSON(data []byte) (*SandboxMachinesAccessFilesOutput, error) {
	var v SandboxMachinesAccessFilesOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessFilesOutputToJSON serializes a SandboxMachinesAccessFilesOutput to JSON.
func MapSandboxMachinesAccessFilesOutputToJSON(v *SandboxMachinesAccessFilesOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesAccessFilesBody represents the sandbox machines access files body type.
type SandboxMachinesAccessFilesBody struct {
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// MapSandboxMachinesAccessFilesBodyFromJSON deserializes JSON data into a SandboxMachinesAccessFilesBody.
func MapSandboxMachinesAccessFilesBodyFromJSON(data []byte) (*SandboxMachinesAccessFilesBody, error) {
	var v SandboxMachinesAccessFilesBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessFilesBodyToJSON serializes a SandboxMachinesAccessFilesBody to JSON.
func MapSandboxMachinesAccessFilesBodyToJSON(v *SandboxMachinesAccessFilesBody) ([]byte, error) {
	return json.Marshal(v)
}
