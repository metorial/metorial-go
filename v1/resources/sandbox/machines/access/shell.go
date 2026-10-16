package access

import (
	"encoding/json"
	"time"
)

// SandboxMachinesAccessShellOutputUrls represents the sandbox machines access shell output urls type.
type SandboxMachinesAccessShellOutputUrls struct {
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

// SandboxMachinesAccessShellOutput represents the sandbox machines access shell output type.
type SandboxMachinesAccessShellOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Token - The short-lived token that authorizes access. It is only shown once.
	Token string `json:"token"`
	// Scope - What the token grants access to
	Scope string                               `json:"scope"`
	Ports *any                                 `json:"ports,omitempty"`
	Urls  SandboxMachinesAccessShellOutputUrls `json:"urls"`
	// ExpiresAt - When the token expires
	ExpiresAt time.Time `json:"expires_at"`
}

// MapSandboxMachinesAccessShellOutputFromJSON deserializes JSON data into a SandboxMachinesAccessShellOutput.
func MapSandboxMachinesAccessShellOutputFromJSON(data []byte) (*SandboxMachinesAccessShellOutput, error) {
	var v SandboxMachinesAccessShellOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessShellOutputToJSON serializes a SandboxMachinesAccessShellOutput to JSON.
func MapSandboxMachinesAccessShellOutputToJSON(v *SandboxMachinesAccessShellOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesAccessShellBody represents the sandbox machines access shell body type.
type SandboxMachinesAccessShellBody struct {
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// MapSandboxMachinesAccessShellBodyFromJSON deserializes JSON data into a SandboxMachinesAccessShellBody.
func MapSandboxMachinesAccessShellBodyFromJSON(data []byte) (*SandboxMachinesAccessShellBody, error) {
	var v SandboxMachinesAccessShellBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessShellBodyToJSON serializes a SandboxMachinesAccessShellBody to JSON.
func MapSandboxMachinesAccessShellBodyToJSON(v *SandboxMachinesAccessShellBody) ([]byte, error) {
	return json.Marshal(v)
}
