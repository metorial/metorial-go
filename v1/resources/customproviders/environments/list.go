package environments

import (
	"encoding/json"
	"time"
)

// CustomProvidersEnvironmentsListOutputItems represents the custom providers environments list output items type.
type CustomProvidersEnvironmentsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// CustomProviderId - Custom provider this environment belongs to.
	CustomProviderId string `json:"custom_provider_id"`
	// ProviderId - Associated provider, by ID. `null` when there is none.
	ProviderId *string `json:"provider_id,omitempty"`
	// CurrentProviderVersionId - Version currently live in this environment. `null` when no version is live.
	CurrentProviderVersionId *string `json:"current_provider_version_id,omitempty"`
	// InstanceId - ID of the instance this environment deploys into.
	InstanceId string `json:"instance_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// CustomProvidersEnvironmentsListOutputPagination represents the custom providers environments list output pagination type.
type CustomProvidersEnvironmentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// CustomProvidersEnvironmentsListOutput represents the custom providers environments list output type.
type CustomProvidersEnvironmentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                          `json:"object"`
	Items      []CustomProvidersEnvironmentsListOutputItems    `json:"items"`
	Pagination CustomProvidersEnvironmentsListOutputPagination `json:"pagination"`
}

// MapCustomProvidersEnvironmentsListOutputFromJSON deserializes JSON data into a CustomProvidersEnvironmentsListOutput.
func MapCustomProvidersEnvironmentsListOutputFromJSON(data []byte) (*CustomProvidersEnvironmentsListOutput, error) {
	var v CustomProvidersEnvironmentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersEnvironmentsListOutputToJSON serializes a CustomProvidersEnvironmentsListOutput to JSON.
func MapCustomProvidersEnvironmentsListOutputToJSON(v *CustomProvidersEnvironmentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// CustomProvidersEnvironmentsListQueryCreatedAt - Only return results whose custom provider environment creation time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersEnvironmentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersEnvironmentsListQueryUpdatedAt - Only return results whose custom provider environment last update time is in this range. Set `gt`, `lt`, or both.
type CustomProvidersEnvironmentsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// CustomProvidersEnvironmentsListQuery represents the custom providers environments list query type.
type CustomProvidersEnvironmentsListQuery struct {
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
	// Id - Only return environments with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CustomProviderVersionId - Only return environments running one of these versions.
	CustomProviderVersionId *any `json:"custom_provider_version_id,omitempty"`
	// CustomProviderId - Only return environments of one of these custom providers.
	CustomProviderId *any `json:"custom_provider_id,omitempty"`
	// CreatedAt - Only return results whose custom provider environment creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *CustomProvidersEnvironmentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose custom provider environment last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *CustomProvidersEnvironmentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapCustomProvidersEnvironmentsListQueryFromJSON deserializes JSON data into a CustomProvidersEnvironmentsListQuery.
func MapCustomProvidersEnvironmentsListQueryFromJSON(data []byte) (*CustomProvidersEnvironmentsListQuery, error) {
	var v CustomProvidersEnvironmentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapCustomProvidersEnvironmentsListQueryToJSON serializes a CustomProvidersEnvironmentsListQuery to JSON.
func MapCustomProvidersEnvironmentsListQueryToJSON(v *CustomProvidersEnvironmentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
