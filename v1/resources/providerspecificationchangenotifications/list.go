package providerspecificationchangenotifications

import (
	"encoding/json"
	"time"
)

// ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification - Specification on one side of the change.
type ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsToSpecification - Specification on one side of the change.
type ProviderSpecificationChangeNotificationsListOutputItemsToSpecification struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion - Provider version on one side of the change.
type ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion - Provider version on one side of the change.
type ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Version - Release identifier of the provider version.
	Version string `json:"version"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderSpecificationChangeNotificationsListOutputItems represents the provider specification change notifications list output items type.
type ProviderSpecificationChangeNotificationsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// ProviderId - Provider whose specification changed.
	ProviderId string `json:"provider_id"`
	// ProviderVersionId - Provider version whose specification changed.
	ProviderVersionId string `json:"provider_version_id"`
	// FromSpecification - Specification on one side of the change.
	FromSpecification *ProviderSpecificationChangeNotificationsListOutputItemsFromSpecification `json:"from_specification,omitempty"`
	// ToSpecification - Specification on one side of the change.
	ToSpecification *ProviderSpecificationChangeNotificationsListOutputItemsToSpecification `json:"to_specification,omitempty"`
	// FromProviderVersion - Provider version on one side of the change.
	FromProviderVersion *ProviderSpecificationChangeNotificationsListOutputItemsFromProviderVersion `json:"from_provider_version,omitempty"`
	// ToProviderVersion - Provider version on one side of the change.
	ToProviderVersion *ProviderSpecificationChangeNotificationsListOutputItemsToProviderVersion `json:"to_provider_version,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// ProviderSpecificationChangeNotificationsListOutputPagination represents the provider specification change notifications list output pagination type.
type ProviderSpecificationChangeNotificationsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// ProviderSpecificationChangeNotificationsListOutput represents the provider specification change notifications list output type.
type ProviderSpecificationChangeNotificationsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                                       `json:"object"`
	Items      []ProviderSpecificationChangeNotificationsListOutputItems    `json:"items"`
	Pagination ProviderSpecificationChangeNotificationsListOutputPagination `json:"pagination"`
}

// MapProviderSpecificationChangeNotificationsListOutputFromJSON deserializes JSON data into a ProviderSpecificationChangeNotificationsListOutput.
func MapProviderSpecificationChangeNotificationsListOutputFromJSON(data []byte) (*ProviderSpecificationChangeNotificationsListOutput, error) {
	var v ProviderSpecificationChangeNotificationsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderSpecificationChangeNotificationsListOutputToJSON serializes a ProviderSpecificationChangeNotificationsListOutput to JSON.
func MapProviderSpecificationChangeNotificationsListOutputToJSON(v *ProviderSpecificationChangeNotificationsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderSpecificationChangeNotificationsListQueryCreatedAt - Only return results whose provider specification change notification time is in this range. Set `gt`, `lt`, or both.
type ProviderSpecificationChangeNotificationsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// ProviderSpecificationChangeNotificationsListQuery represents the provider specification change notifications list query type.
type ProviderSpecificationChangeNotificationsListQuery struct {
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
	// Id - Only return provider specification change notifications with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Target - Only return provider specification change notifications with one of these targets.
	Target *any `json:"target,omitempty"`
	// ProviderId - Only return provider specification change notifications for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderVersionId - Only return provider specification change notifications for one of these provider versions.
	ProviderVersionId *any `json:"provider_version_id,omitempty"`
	// ProviderSpecificationId - Only return provider specification change notifications for one of these provider specifications.
	ProviderSpecificationId *any `json:"provider_specification_id,omitempty"`
	// CreatedAt - Only return results whose provider specification change notification time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *ProviderSpecificationChangeNotificationsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapProviderSpecificationChangeNotificationsListQueryFromJSON deserializes JSON data into a ProviderSpecificationChangeNotificationsListQuery.
func MapProviderSpecificationChangeNotificationsListQueryFromJSON(data []byte) (*ProviderSpecificationChangeNotificationsListQuery, error) {
	var v ProviderSpecificationChangeNotificationsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderSpecificationChangeNotificationsListQueryToJSON serializes a ProviderSpecificationChangeNotificationsListQuery to JSON.
func MapProviderSpecificationChangeNotificationsListQueryToJSON(v *ProviderSpecificationChangeNotificationsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
