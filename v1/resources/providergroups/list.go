package providergroups

import (
	"encoding/json"
	"time"
)

// ProviderGroupsListOutputItems represents the provider groups list output items type.
type ProviderGroupsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name for the group.
	Slug string `json:"slug"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderGroupsListOutputPagination represents the provider groups list output pagination type.
type ProviderGroupsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderGroupsListOutput represents the provider groups list output type.
type ProviderGroupsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []ProviderGroupsListOutputItems    `json:"items"`
	Pagination ProviderGroupsListOutputPagination `json:"pagination"`
}

// MapProviderGroupsListOutputFromJSON deserializes JSON data into a ProviderGroupsListOutput.
func MapProviderGroupsListOutputFromJSON(data []byte) (*ProviderGroupsListOutput, error) {
	var v ProviderGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsListOutputToJSON serializes a ProviderGroupsListOutput to JSON.
func MapProviderGroupsListOutputToJSON(v *ProviderGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderGroupsListQueryCreatedAt - Only return results whose provider group creation time is in this range. Set `gt`, `lt`, or both.
type ProviderGroupsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderGroupsListQueryUpdatedAt - Only return results whose provider group last update time is in this range. Set `gt`, `lt`, or both.
type ProviderGroupsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderGroupsListQuery represents the provider groups list query type.
type ProviderGroupsListQuery struct {
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
	// Id - Filter by group ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderListingId - Filter by provider listing ID(s)
	ProviderListingId *any `json:"provider_listing_id,omitempty"`
	// CreatedAt - Only return results whose provider group creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderGroupsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider group last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProviderGroupsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProviderGroupsListQueryFromJSON deserializes JSON data into a ProviderGroupsListQuery.
func MapProviderGroupsListQueryFromJSON(data []byte) (*ProviderGroupsListQuery, error) {
	var v ProviderGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderGroupsListQueryToJSON serializes a ProviderGroupsListQuery to JSON.
func MapProviderGroupsListQueryToJSON(v *ProviderGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
