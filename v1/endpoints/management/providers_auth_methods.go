package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers/authmethods"
)

// ProvidersAuthMethodsEndpoint provides access to a provider auth method is one way to authenticate with a provider, such as OAuth, an API token, or custom credentials. A provider version can offer multiple auth methods.
type ProvidersAuthMethodsEndpoint struct {
	client *endpoint.Client
}

// NewProvidersAuthMethodsEndpoint creates a new ProvidersAuthMethodsEndpoint.
func NewProvidersAuthMethodsEndpoint(client *endpoint.Client) *ProvidersAuthMethodsEndpoint {
	return &ProvidersAuthMethodsEndpoint{client: client}
}

// ProvidersAuthMethodsEndpointListParams contains optional query parameters for List.
type ProvidersAuthMethodsEndpointListParams struct {
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
	// ProviderVersionId - Provider version to list auth methods for.
	ProviderVersionId string `json:"provider_version_id"`
	// Adapter - Only return auth methods supporting this adapter.
	Adapter *string `json:"adapter,omitempty"`
}

// List returns a paginated list of provider auth methods.
func (e *ProvidersAuthMethodsEndpoint) List(instanceId string, params *ProvidersAuthMethodsEndpointListParams) (*authmethods.ProvidersAuthMethodsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-auth-methods"},
		Query: query,
	}
	var result authmethods.ProvidersAuthMethodsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider auth method by ID.
func (e *ProvidersAuthMethodsEndpoint) Get(instanceId string, providerAuthMethodId string) (*authmethods.ProvidersAuthMethodsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-auth-methods", providerAuthMethodId},
	}
	var result authmethods.ProvidersAuthMethodsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
