package consumerprofiles

import (
	"encoding/json"
)

// PortalsConsumerProfilesListOutputPagination represents the portals consumer profiles list output pagination type.
type PortalsConsumerProfilesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// PortalsConsumerProfilesListOutput represents the portals consumer profiles list output type.
type PortalsConsumerProfilesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                      `json:"object"`
	Items      []map[string]any                            `json:"items"`
	Pagination PortalsConsumerProfilesListOutputPagination `json:"pagination"`
}

// MapPortalsConsumerProfilesListOutputFromJSON deserializes JSON data into a PortalsConsumerProfilesListOutput.
func MapPortalsConsumerProfilesListOutputFromJSON(data []byte) (*PortalsConsumerProfilesListOutput, error) {
	var v PortalsConsumerProfilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesListOutputToJSON serializes a PortalsConsumerProfilesListOutput to JSON.
func MapPortalsConsumerProfilesListOutputToJSON(v *PortalsConsumerProfilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerProfilesListQuery represents the portals consumer profiles list query type.
type PortalsConsumerProfilesListQuery struct {
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
	// Search - Search by name or email.
	Search *string `json:"search,omitempty"`
	// Email - Only return profiles with one of these emails.
	Email *any `json:"email,omitempty"`
	// ConsumerGroupId - Only return profiles in this consumer group.
	ConsumerGroupId *string `json:"consumer_group_id,omitempty"`
	// Source - Only return profiles created manually or through SSO.
	Source *string `json:"source,omitempty"`
	// Status - Only return profiles with one of these statuses.
	Status *any `json:"status,omitempty"`
}

// MapPortalsConsumerProfilesListQueryFromJSON deserializes JSON data into a PortalsConsumerProfilesListQuery.
func MapPortalsConsumerProfilesListQueryFromJSON(data []byte) (*PortalsConsumerProfilesListQuery, error) {
	var v PortalsConsumerProfilesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerProfilesListQueryToJSON serializes a PortalsConsumerProfilesListQuery to JSON.
func MapPortalsConsumerProfilesListQueryToJSON(v *PortalsConsumerProfilesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
