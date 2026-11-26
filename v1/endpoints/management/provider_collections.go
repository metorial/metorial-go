package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providercollections"
)

// ProviderCollectionsEndpoint provides access to a provider collection is a curated set of providers.
type ProviderCollectionsEndpoint struct {
	client *endpoint.Client
}

// NewProviderCollectionsEndpoint creates a new ProviderCollectionsEndpoint.
func NewProviderCollectionsEndpoint(client *endpoint.Client) *ProviderCollectionsEndpoint {
	return &ProviderCollectionsEndpoint{client: client}
}

// ProviderCollectionsEndpointListParams contains optional query parameters for List.
type ProviderCollectionsEndpointListParams struct {
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
	// Id - Filter by collection ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderListingId - Filter by provider listing ID(s)
	ProviderListingId *any `json:"provider_listing_id,omitempty"`
}

// List returns a paginated list of provider collections.
func (e *ProviderCollectionsEndpoint) List(instanceId string, params *ProviderCollectionsEndpointListParams) (*providercollections.ProviderCollectionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-collections"},
		Query: query,
	}
	var result providercollections.ProviderCollectionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider collection by ID.
func (e *ProviderCollectionsEndpoint) Get(instanceId string, providerCollectionId string) (*providercollections.ProviderCollectionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-collections", providerCollectionId},
	}
	var result providercollections.ProviderCollectionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
