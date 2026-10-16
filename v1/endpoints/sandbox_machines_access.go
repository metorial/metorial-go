package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/access"
)

// SandboxMachinesAccessEndpoint provides access to create and operate sandbox machines. A machine boots from an image, can mount disks, and moves between running, paused, and hibernated states.
type SandboxMachinesAccessEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesAccessEndpoint creates a new SandboxMachinesAccessEndpoint.
func NewSandboxMachinesAccessEndpoint(client *endpoint.Client) *SandboxMachinesAccessEndpoint {
	return &SandboxMachinesAccessEndpoint{client: client}
}

// SandboxMachinesAccessEndpointShellBody contains the request body for Shell.
type SandboxMachinesAccessEndpointShellBody struct {
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// SandboxMachinesAccessEndpointFilesBody contains the request body for Files.
type SandboxMachinesAccessEndpointFilesBody struct {
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// SandboxMachinesAccessEndpointPortsBody contains the request body for Ports.
type SandboxMachinesAccessEndpointPortsBody struct {
	// Ports - The ports the token grants access to, or "*" for all ports
	Ports any `json:"ports"`
	// TtlSeconds - How long the token stays valid, in seconds
	TtlSeconds *float64 `json:"ttl_seconds,omitempty"`
}

// Shell creates a short-lived token for an interactive shell in the machine.
func (e *SandboxMachinesAccessEndpoint) Shell(machineId string, body *SandboxMachinesAccessEndpointShellBody) (*access.SandboxMachinesAccessShellOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "access", "shell"},
		Body: body,
	}
	var result access.SandboxMachinesAccessShellOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Files creates a short-lived token for the machine's file service.
func (e *SandboxMachinesAccessEndpoint) Files(machineId string, body *SandboxMachinesAccessEndpointFilesBody) (*access.SandboxMachinesAccessFilesOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "access", "files"},
		Body: body,
	}
	var result access.SandboxMachinesAccessFilesOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ports creates a short-lived token for reaching ports inside the machine.
func (e *SandboxMachinesAccessEndpoint) Ports(machineId string, body *SandboxMachinesAccessEndpointPortsBody) (*access.SandboxMachinesAccessPortsOutput, error) {
	req := &endpoint.Request{
		Path: []string{"sandbox", "machines", machineId, "access", "ports"},
		Body: body,
	}
	var result access.SandboxMachinesAccessPortsOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
