package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/permissions"
)

// StoresPermissionsEndpoint provides access to create and manage instance stores backed by Cargo.
type StoresPermissionsEndpoint struct {
	client *endpoint.Client
}

// NewStoresPermissionsEndpoint creates a new StoresPermissionsEndpoint.
func NewStoresPermissionsEndpoint(client *endpoint.Client) *StoresPermissionsEndpoint {
	return &StoresPermissionsEndpoint{client: client}
}

// Get returns the effective Cargo permissions for the current actor on a specific store.
func (e *StoresPermissionsEndpoint) Get(instanceId string, storeId string) (*permissions.StoresPermissionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "stores", storeId, "permissions"},
	}
	var result permissions.StoresPermissionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
