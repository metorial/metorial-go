package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providers/specifications"
)

// ProvidersSpecificationsEndpoint provides access to a provider specification is the set of tools, auth methods, and configuration fields a provider version offers.
type ProvidersSpecificationsEndpoint struct {
	client *endpoint.Client
}

// NewProvidersSpecificationsEndpoint creates a new ProvidersSpecificationsEndpoint.
func NewProvidersSpecificationsEndpoint(client *endpoint.Client) *ProvidersSpecificationsEndpoint {
	return &ProvidersSpecificationsEndpoint{client: client}
}

// ProvidersSpecificationsEndpointListParams contains optional query parameters for List.
type ProvidersSpecificationsEndpointListParams struct {
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
	// Id - Filter by specification ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Filter by provider version ID(s)
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ProviderDeploymentId - Filter by provider deployment ID(s)
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Filter by provider config ID(s)
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// Search - Match a provider specification by its exact ID
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose provider specification creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider specification last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of provider specifications.
func (e *ProvidersSpecificationsEndpoint) List(instanceId string, params *ProvidersSpecificationsEndpointListParams) (*specifications.ProvidersSpecificationsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-specifications"},
		Query: query,
	}
	var result specifications.ProvidersSpecificationsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider specification by ID.
func (e *ProvidersSpecificationsEndpoint) Get(instanceId string, providerSpecificationId string) (*specifications.ProvidersSpecificationsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-specifications", providerSpecificationId},
	}
	var result specifications.ProvidersSpecificationsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
