package consumergroups

import (
	"encoding/json"
	"time"
)

// PortalsConsumerGroupsListOutputItems represents the portals consumer groups list output items type.
type PortalsConsumerGroupsListOutputItems struct {
	Object      string    `json:"object"`
	Id          string    `json:"id"`
	Status      string    `json:"status"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	IsDefault   bool      `json:"is_default"`
	SsoGroupIds []string  `json:"sso_group_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// PortalsConsumerGroupsListOutputPagination represents the portals consumer groups list output pagination type.
type PortalsConsumerGroupsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsConsumerGroupsListOutput represents the portals consumer groups list output type.
type PortalsConsumerGroupsListOutput struct {
	Items      []PortalsConsumerGroupsListOutputItems    `json:"items"`
	Pagination PortalsConsumerGroupsListOutputPagination `json:"pagination"`
}

// MapPortalsConsumerGroupsListOutputFromJSON deserializes JSON data into a PortalsConsumerGroupsListOutput.
func MapPortalsConsumerGroupsListOutputFromJSON(data []byte) (*PortalsConsumerGroupsListOutput, error) {
	var v PortalsConsumerGroupsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsListOutputToJSON serializes a PortalsConsumerGroupsListOutput to JSON.
func MapPortalsConsumerGroupsListOutputToJSON(v *PortalsConsumerGroupsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsConsumerGroupsListQuery represents the portals consumer groups list query type.
type PortalsConsumerGroupsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Status *any     `json:"status,omitempty"`
	Search *string  `json:"search,omitempty"`
}

// MapPortalsConsumerGroupsListQueryFromJSON deserializes JSON data into a PortalsConsumerGroupsListQuery.
func MapPortalsConsumerGroupsListQueryFromJSON(data []byte) (*PortalsConsumerGroupsListQuery, error) {
	var v PortalsConsumerGroupsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsConsumerGroupsListQueryToJSON serializes a PortalsConsumerGroupsListQuery to JSON.
func MapPortalsConsumerGroupsListQueryToJSON(v *PortalsConsumerGroupsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
