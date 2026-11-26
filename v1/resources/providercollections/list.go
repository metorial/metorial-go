package providercollections

import (
	"encoding/json"
	"time"
)

// ProviderCollectionsListOutputItems represents the provider collections list output items type.
type ProviderCollectionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description string `json:"description"`
	// Slug - URL-safe name for the collection.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderCollectionsListOutputPagination represents the provider collections list output pagination type.
type ProviderCollectionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderCollectionsListOutput represents the provider collections list output type.
type ProviderCollectionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                  `json:"object"`
	Items      []ProviderCollectionsListOutputItems    `json:"items"`
	Pagination ProviderCollectionsListOutputPagination `json:"pagination"`
}

// MapProviderCollectionsListOutputFromJSON deserializes JSON data into a ProviderCollectionsListOutput.
func MapProviderCollectionsListOutputFromJSON(data []byte) (*ProviderCollectionsListOutput, error) {
	var v ProviderCollectionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderCollectionsListOutputToJSON serializes a ProviderCollectionsListOutput to JSON.
func MapProviderCollectionsListOutputToJSON(v *ProviderCollectionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderCollectionsListQuery represents the provider collections list query type.
type ProviderCollectionsListQuery struct {
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

// MapProviderCollectionsListQueryFromJSON deserializes JSON data into a ProviderCollectionsListQuery.
func MapProviderCollectionsListQueryFromJSON(data []byte) (*ProviderCollectionsListQuery, error) {
	var v ProviderCollectionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderCollectionsListQueryToJSON serializes a ProviderCollectionsListQuery to JSON.
func MapProviderCollectionsListQueryToJSON(v *ProviderCollectionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
