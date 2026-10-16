package access

import (
	"encoding/json"
	"time"
)

// SandboxMachinesAccessPortsOutputUrls represents the sandbox machines access ports output urls type.
type SandboxMachinesAccessPortsOutputUrls struct {
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

// SandboxMachinesAccessPortsOutput represents the sandbox machines access ports output type.
type SandboxMachinesAccessPortsOutput struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Token - The short-lived token that authorizes access. It is only shown once.
	Token string `json:"token"`
	// Scope - What the token grants access to
	Scope string                               `json:"scope"`
	Ports *any                                 `json:"ports,omitempty"`
	Urls  SandboxMachinesAccessPortsOutputUrls `json:"urls"`
	// ExpiresAt - When the token expires
	ExpiresAt time.Time `json:"expires_at"`
}

// MapSandboxMachinesAccessPortsOutputFromJSON deserializes JSON data into a SandboxMachinesAccessPortsOutput.
func MapSandboxMachinesAccessPortsOutputFromJSON(data []byte) (*SandboxMachinesAccessPortsOutput, error) {
	var v SandboxMachinesAccessPortsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessPortsOutputToJSON serializes a SandboxMachinesAccessPortsOutput to JSON.
func MapSandboxMachinesAccessPortsOutputToJSON(v *SandboxMachinesAccessPortsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesAccessPortsBody represents the sandbox machines access ports body type.
type SandboxMachinesAccessPortsBody struct {
	// Ports - The ports the token grants access to, or "*" for all ports
	Ports any `json:"ports"`
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// MapSandboxMachinesAccessPortsBodyFromJSON deserializes JSON data into a SandboxMachinesAccessPortsBody.
func MapSandboxMachinesAccessPortsBodyFromJSON(data []byte) (*SandboxMachinesAccessPortsBody, error) {
	var v SandboxMachinesAccessPortsBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesAccessPortsBodyToJSON serializes a SandboxMachinesAccessPortsBody to JSON.
func MapSandboxMachinesAccessPortsBodyToJSON(v *SandboxMachinesAccessPortsBody) ([]byte, error) {
	return json.Marshal(v)
}
