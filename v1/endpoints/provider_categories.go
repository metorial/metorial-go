package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providercategories"
)

// ProviderCategoriesEndpoint provides access to a provider category is a grouping of providers by function.
type ProviderCategoriesEndpoint struct {
	client *endpoint.Client
}

// NewProviderCategoriesEndpoint creates a new ProviderCategoriesEndpoint.
func NewProviderCategoriesEndpoint(client *endpoint.Client) *ProviderCategoriesEndpoint {
	return &ProviderCategoriesEndpoint{client: client}
}

// ProviderCategoriesEndpointListParams contains optional query parameters for List.
type ProviderCategoriesEndpointListParams struct {
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
	// Id - Filter by category ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderListingId - Filter by provider listing ID(s)
	ProviderListingId *any `json:"provider_listing_id,omitempty"`
}

// List returns a paginated list of provider categories.
func (e *ProviderCategoriesEndpoint) List(params *ProviderCategoriesEndpointListParams) (*providercategories.ProviderCategoriesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-categories"},
		Query: query,
	}
	var result providercategories.ProviderCategoriesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider category by ID.
func (e *ProviderCategoriesEndpoint) Get(providerCategoryId string) (*providercategories.ProviderCategoriesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-categories", providerCategoryId},
	}
	var result providercategories.ProviderCategoriesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
