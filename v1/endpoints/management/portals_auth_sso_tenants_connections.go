package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/portals/auth/ssotenants/connections"
)

// PortalsAuthSsoTenantsConnectionsEndpoint provides access to manage the Ares-backed authentication configuration for a portal.
type PortalsAuthSsoTenantsConnectionsEndpoint struct {
	client *endpoint.Client
}

// NewPortalsAuthSsoTenantsConnectionsEndpoint creates a new PortalsAuthSsoTenantsConnectionsEndpoint.
func NewPortalsAuthSsoTenantsConnectionsEndpoint(client *endpoint.Client) *PortalsAuthSsoTenantsConnectionsEndpoint {
	return &PortalsAuthSsoTenantsConnectionsEndpoint{client: client}
}

// PortalsAuthSsoTenantsConnectionsEndpointListParams contains optional query parameters for List.
type PortalsAuthSsoTenantsConnectionsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// List returns SSO connections that belong to a portal SSO tenant.
func (e *PortalsAuthSsoTenantsConnectionsEndpoint) List(instanceId string, portalId string, ssoTenantId string, params *PortalsAuthSsoTenantsConnectionsEndpointListParams) (*connections.PortalsAuthSsoTenantsConnectionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "portals", portalId, "auth", "sso-tenants", ssoTenantId, "connections"},
		Query: query,
	}
	var result connections.PortalsAuthSsoTenantsConnectionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
