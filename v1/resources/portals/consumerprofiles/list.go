package consumerprofiles

import (
	"encoding/json"
)

// PortalsConsumerProfilesListOutputPagination represents the portals consumer profiles list output pagination type.
type PortalsConsumerProfilesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsConsumerProfilesListOutput represents the portals consumer profiles list output type.
type PortalsConsumerProfilesListOutput struct {
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
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	Search          *string  `json:"search,omitempty"`
	ConsumerGroupId *string  `json:"consumer_group_id,omitempty"`
	Status          *any     `json:"status,omitempty"`
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
