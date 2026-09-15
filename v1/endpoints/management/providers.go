package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers"
)

// ProvidersEndpoint provides access to a provider is a read-only template for an MCP server integration (like GitHub or Slack). To use a provider, create a deployment from it.
type ProvidersEndpoint struct {
	client *endpoint.Client
}

// NewProvidersEndpoint creates a new ProvidersEndpoint.
func NewProvidersEndpoint(client *endpoint.Client) *ProvidersEndpoint {
	return &ProvidersEndpoint{client: client}
}

// ProvidersEndpointListParams contains optional query parameters for List.
type ProvidersEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by provider ID(s)
	Id *any `json:"id,omitempty"`
	// Search - Search providers by name, description, or readme
	Search *string `json:"search,omitempty"`
	// AuthMethod - Filter by auth method — matches an auth method ID, auth method global ID, key, name, or type (oauth, token, service_account, custom)
	AuthMethod *any `json:"auth_method,omitempty"`
	// AuthSetup - Filter by auth setup status. "configured" matches providers with a token or custom auth method, or with auth credentials already configured. "not_configured" matches providers with only OAuth auth methods and no auth credentials configured.
	AuthSetup *any `json:"auth_setup,omitempty"`
}

// List returns a paginated list of providers.
func (e *ProvidersEndpoint) List(instanceId string, params *ProvidersEndpointListParams) (*providers.ProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "providers"},
		Query: query,
	}
	var result providers.ProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider by ID.
func (e *ProvidersEndpoint) Get(instanceId string, providerId string) (*providers.ProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "providers", providerId},
	}
	var result providers.ProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
