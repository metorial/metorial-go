package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers/tools"
)

// ProvidersToolsEndpoint provides access to a tool is a single action a provider can perform, like `search_issues` or `send_message`. Agents call tools over MCP, and each provider version has its own set.
type ProvidersToolsEndpoint struct {
	client *endpoint.Client
}

// NewProvidersToolsEndpoint creates a new ProvidersToolsEndpoint.
func NewProvidersToolsEndpoint(client *endpoint.Client) *ProvidersToolsEndpoint {
	return &ProvidersToolsEndpoint{client: client}
}

// ProvidersToolsEndpointListParams contains optional query parameters for List.
type ProvidersToolsEndpointListParams struct {
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
	// ProviderVersionId - Provider version to list tools for.
	ProviderVersionId string `json:"provider_version_id"`
	// ProviderAuthMethodId - Filter to tools compatible with one auth method ID, or the common compatible subset for multiple auth method IDs.
	ProviderAuthMethodId *any `json:"provider_auth_method_id,omitempty"`
}

// List returns a paginated list of provider tools for a provider version.
func (e *ProvidersToolsEndpoint) List(params *ProvidersToolsEndpointListParams) (*tools.ProvidersToolsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-tools"},
		Query: query,
	}
	var result tools.ProvidersToolsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider tool by ID.
func (e *ProvidersToolsEndpoint) Get(providerToolId string) (*tools.ProvidersToolsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-tools", providerToolId},
	}
	var result tools.ProvidersToolsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
