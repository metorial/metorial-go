package versions

import (
	"encoding/json"
	"time"
)

// ProvidersVersionsListOutputItems represents the providers versions list output items type.
type ProvidersVersionsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release label for the version.
	Version string `json:"version"`
	// ProviderId - Provider this version belongs to.
	ProviderId string `json:"provider_id"`
	// IsCurrent - The version deployments follow unless locked to another.
	IsCurrent bool `json:"is_current"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SpecificationId - Specification this version belongs to.
	SpecificationId *string `json:"specification_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProvidersVersionsListOutputPagination represents the providers versions list output pagination type.
type ProvidersVersionsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProvidersVersionsListOutput represents the providers versions list output type.
type ProvidersVersionsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                `json:"object"`
	Items      []ProvidersVersionsListOutputItems    `json:"items"`
	Pagination ProvidersVersionsListOutputPagination `json:"pagination"`
}

// MapProvidersVersionsListOutputFromJSON deserializes JSON data into a ProvidersVersionsListOutput.
func MapProvidersVersionsListOutputFromJSON(data []byte) (*ProvidersVersionsListOutput, error) {
	var v ProvidersVersionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersVersionsListOutputToJSON serializes a ProvidersVersionsListOutput to JSON.
func MapProvidersVersionsListOutputToJSON(v *ProvidersVersionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProvidersVersionsListQueryCreatedAt - Only return results whose provider version creation time is in this range. Set `gt`, `lt`, or both.
type ProvidersVersionsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProvidersVersionsListQueryUpdatedAt - Only return results whose provider version last update time is in this range. Set `gt`, `lt`, or both.
type ProvidersVersionsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProvidersVersionsListQuery represents the providers versions list query type.
type ProvidersVersionsListQuery struct {
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
	CreatedAt *ProvidersVersionsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose provider version last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *ProvidersVersionsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapProvidersVersionsListQueryFromJSON deserializes JSON data into a ProvidersVersionsListQuery.
func MapProvidersVersionsListQueryFromJSON(data []byte) (*ProvidersVersionsListQuery, error) {
	var v ProvidersVersionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProvidersVersionsListQueryToJSON serializes a ProvidersVersionsListQuery to JSON.
func MapProvidersVersionsListQueryToJSON(v *ProvidersVersionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
