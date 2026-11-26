package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers/versions"
)

// ProvidersVersionsEndpoint provides access to a provider version is a specific release of a provider. Each version has its own tools, auth methods, and config schema. A deployment can lock to one version or follow the provider's current version.
type ProvidersVersionsEndpoint struct {
	client *endpoint.Client
}

// NewProvidersVersionsEndpoint creates a new ProvidersVersionsEndpoint.
func NewProvidersVersionsEndpoint(client *endpoint.Client) *ProvidersVersionsEndpoint {
	return &ProvidersVersionsEndpoint{client: client}
}

// ProvidersVersionsEndpointListParams contains optional query parameters for List.
type ProvidersVersionsEndpointListParams struct {
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
	// Id - Filter by version ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// Search - Match a provider version by its exact ID
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider version creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider version last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of provider versions.
func (e *ProvidersVersionsEndpoint) List(params *ProvidersVersionsEndpointListParams) (*versions.ProvidersVersionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-versions"},
		Query: query,
	}
	var result versions.ProvidersVersionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider version by ID.
func (e *ProvidersVersionsEndpoint) Get(providerVersionId string) (*versions.ProvidersVersionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-versions", providerVersionId},
	}
	var result versions.ProvidersVersionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
