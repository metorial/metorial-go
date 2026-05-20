package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/auth/app"
)

// PortalsAuthAppEndpoint provides access to manage the Ares-backed authentication configuration for a portal.
type PortalsAuthAppEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAuthAppEndpoint creates a new PortalsAuthAppEndpoint.
func NewPortalsAuthAppEndpoint(client *endpoint.Client) *PortalsAuthAppEndpoint {
	return &PortalsAuthAppEndpoint{client: client}
}

// PortalsAuthAppEndpointUpdateBody contains the request body for Update.
type PortalsAuthAppEndpointUpdateBody struct {
	EmailWhitelist *[]string `json:"email_whitelist,omitempty"`
}

// Get returns the Ares app configuration for a portal.
func (e *PortalsAuthAppEndpoint) Get(instanceId string, portalId string) (*app.PortalsAuthAppGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "auth", "app"},
	}
	var result app.PortalsAuthAppGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates the portal auth app configuration stored on the portal surface.
func (e *PortalsAuthAppEndpoint) Update(instanceId string, portalId string, body *PortalsAuthAppEndpointUpdateBody) (*app.PortalsAuthAppUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "auth", "app"},
		Body: body,
	}
	var result app.PortalsAuthAppUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
