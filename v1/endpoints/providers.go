package endpoints

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
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
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
func (e *ProvidersEndpoint) List(params *ProvidersEndpointListParams) (*providers.ProvidersListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"providers"},
		Query: query,
	}
	var result providers.ProvidersListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider by ID.
func (e *ProvidersEndpoint) Get(providerId string) (*providers.ProvidersGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"providers", providerId},
	}
	var result providers.ProvidersGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
