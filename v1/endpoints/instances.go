package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/instances"
)

// InstancesEndpoint provides access to an instance is an isolated environment within a Metorial project. Only the instances you can access are listed.
type InstancesEndpoint struct {
	client *endpoint.Client
}

// NewInstancesEndpoint creates a new InstancesEndpoint.
func NewInstancesEndpoint(client *endpoint.Client) *InstancesEndpoint {
	return &InstancesEndpoint{client: client}
}

// Get retrieves an instance.
func (e *InstancesEndpoint) Get(instanceId string) (*instances.InstancesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId},
	}
	var result instances.InstancesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns the instances in the organization that the caller can access.
func (e *InstancesEndpoint) List() (*instances.InstancesListOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances"},
	}
	var result instances.InstancesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
