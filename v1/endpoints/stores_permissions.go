package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/stores/permissions"
)

// StoresPermissionsEndpoint provides access to a store is a shared collection of files and documents. Participants get read or write access to its contents.
type StoresPermissionsEndpoint struct {
	client *endpoint.Client
}

// NewStoresPermissionsEndpoint creates a new StoresPermissionsEndpoint.
func NewStoresPermissionsEndpoint(client *endpoint.Client) *StoresPermissionsEndpoint {
	return &StoresPermissionsEndpoint{client: client}
}

// Get returns what the caller can do with a store.
func (e *StoresPermissionsEndpoint) Get(storeId string) (*permissions.StoresPermissionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"stores", storeId, "permissions"},
	}
	var result permissions.StoresPermissionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
