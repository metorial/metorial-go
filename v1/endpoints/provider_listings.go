package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerlistings"
)

// ProviderListingsEndpoint provides access to a provider listing is a provider with its marketplace details.
type ProviderListingsEndpoint struct {
	client *endpoint.Client
}

// NewProviderListingsEndpoint creates a new ProviderListingsEndpoint.
func NewProviderListingsEndpoint(client *endpoint.Client) *ProviderListingsEndpoint {
	return &ProviderListingsEndpoint{client: client}
}

// ProviderListingsEndpointListParams contains optional query parameters for List.
type ProviderListingsEndpointListParams struct {
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
	// Search - Search listings by name, description, or readme.
	Search *string `json:"search,omitempty"`
	// Id - Filter by listing ID(s).
	Id *any `json:"id,omitempty"`
	// ProviderCategoryId - Only return listings in these categories.
	ProviderCategoryId *any `json:"provider_category_id,omitempty"`
	// ProviderCollectionId - Only return listings in these collections.
	ProviderCollectionId *any `json:"provider_collection_id,omitempty"`
	// ProviderGroupId - Only return listings in these groups.
	ProviderGroupId *any `json:"provider_group_id,omitempty"`
	// PublisherId - Only return listings from these publishers.
	PublisherId *any `json:"publisher_id,omitempty"`
	// IsOwner - Only return listings published by your organization.
	IsOwner *bool `json:"is_owner,omitempty"`
	// IsPublic - Only return public listings.
	IsPublic *bool `json:"is_public,omitempty"`
	// IsVerified - Only return listings reviewed and trusted by Metorial.
	IsVerified *bool `json:"is_verified,omitempty"`
	// IsOfficial - Only return official listings.
	IsOfficial *bool `json:"is_official,omitempty"`
	// IsMetorial - Only return Metorial-maintained listings.
	IsMetorial *bool `json:"is_metorial,omitempty"`
	// CreatedAt - Only return results whose provider listing creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider listing last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *map[string]any `json:"updated_at,omitempty"`
}

// List returns a paginated list of provider listings.
func (e *ProviderListingsEndpoint) List(params *ProviderListingsEndpointListParams) (*providerlistings.ProviderListingsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-listings"},
		Query: query,
	}
	var result providerlistings.ProviderListingsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider listing by ID.
func (e *ProviderListingsEndpoint) Get(providerListingId string) (*providerlistings.ProviderListingsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-listings", providerListingId},
	}
	var result providerlistings.ProviderListingsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
