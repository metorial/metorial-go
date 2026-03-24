package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/instances"
)

// InstancesEndpoint provides access to management endpoints for listing and retrieving instances.
type InstancesEndpoint struct {
	client *endpoint.Client
}

// NewInstancesEndpoint creates a new InstancesEndpoint.
func NewInstancesEndpoint(client *endpoint.Client) *InstancesEndpoint {
	return &InstancesEndpoint{client: client}
}

// Get retrieves metadata and configuration details for a specific instance.
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

// List lists all instances within the organization that the authenticated actor has access to.
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
