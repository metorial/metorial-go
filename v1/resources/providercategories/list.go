package providercategories

import (
	"encoding/json"
	"time"
)

// ProviderCategoriesListOutputItems represents the provider categories list output items type.
type ProviderCategoriesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Description of providers in this category
	Description string `json:"description"`
	// Slug - URL-safe name for the category.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderCategoriesListOutputPagination represents the provider categories list output pagination type.
type ProviderCategoriesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderCategoriesListOutput represents the provider categories list output type.
type ProviderCategoriesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                 `json:"object"`
	Items      []ProviderCategoriesListOutputItems    `json:"items"`
	Pagination ProviderCategoriesListOutputPagination `json:"pagination"`
}

// MapProviderCategoriesListOutputFromJSON deserializes JSON data into a ProviderCategoriesListOutput.
func MapProviderCategoriesListOutputFromJSON(data []byte) (*ProviderCategoriesListOutput, error) {
	var v ProviderCategoriesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderCategoriesListOutputToJSON serializes a ProviderCategoriesListOutput to JSON.
func MapProviderCategoriesListOutputToJSON(v *ProviderCategoriesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderCategoriesListQuery represents the provider categories list query type.
type ProviderCategoriesListQuery struct {
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

// MapProviderCategoriesListQueryFromJSON deserializes JSON data into a ProviderCategoriesListQuery.
func MapProviderCategoriesListQueryFromJSON(data []byte) (*ProviderCategoriesListQuery, error) {
	var v ProviderCategoriesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderCategoriesListQueryToJSON serializes a ProviderCategoriesListQuery to JSON.
func MapProviderCategoriesListQueryToJSON(v *ProviderCategoriesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
