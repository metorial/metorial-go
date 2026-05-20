package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/auth/ssotenants"
)

// PortalsAuthSsoTenantsEndpoint provides access to manage the Ares-backed authentication configuration for a portal.
type PortalsAuthSsoTenantsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAuthSsoTenantsEndpoint creates a new PortalsAuthSsoTenantsEndpoint.
func NewPortalsAuthSsoTenantsEndpoint(client *endpoint.Client) *PortalsAuthSsoTenantsEndpoint {
	return &PortalsAuthSsoTenantsEndpoint{client: client}
}

// PortalsAuthSsoTenantsEndpointListParams contains optional query parameters for List.
type PortalsAuthSsoTenantsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// PortalsAuthSsoTenantsEndpointCreateBody contains the request body for Create.
type PortalsAuthSsoTenantsEndpointCreateBody struct {
	Name string `json:"name"`
}

// List returns the SSO tenants configured for a portal Ares app.
func (e *PortalsAuthSsoTenantsEndpoint) List(instanceId string, portalId string, params *PortalsAuthSsoTenantsEndpointListParams) (*ssotenants.PortalsAuthSsoTenantsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "auth", "sso-tenants"},
		Query: query,
	}
	var result ssotenants.PortalsAuthSsoTenantsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates an SSO tenant for the portal Ares app.
func (e *PortalsAuthSsoTenantsEndpoint) Create(instanceId string, portalId string, body *PortalsAuthSsoTenantsEndpointCreateBody) (*ssotenants.PortalsAuthSsoTenantsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "auth", "sso-tenants"},
		Body: body,
	}
	var result ssotenants.PortalsAuthSsoTenantsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Setup creates an Ares setup URL for finishing portal SSO tenant configuration.
func (e *PortalsAuthSsoTenantsEndpoint) Setup(instanceId string, portalId string, ssoTenantId string) (*ssotenants.PortalsAuthSsoTenantsSetupOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "portals", portalId, "auth", "sso-tenants", ssoTenantId, "setup"},
	}
	var result ssotenants.PortalsAuthSsoTenantsSetupOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
