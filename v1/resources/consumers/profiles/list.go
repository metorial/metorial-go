package profiles

import (
	"encoding/json"
	"time"
)

// ConsumersProfilesListOutputItemsGroupsGroup represents the consumers profiles list output items groups group type.
type ConsumersProfilesListOutputItemsGroupsGroup struct {
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

// ConsumersProfilesListOutputItemsGroups represents the consumers profiles list output items groups type.
type ConsumersProfilesListOutputItemsGroups struct {
	Object      string                                      `json:"object"`
	Group       ConsumersProfilesListOutputItemsGroupsGroup `json:"group"`
	AssignedVia string                                      `json:"assigned_via"`
}

// ConsumersProfilesListOutputItems represents the consumers profiles list output items type.
type ConsumersProfilesListOutputItems struct {
	Object     string                                    `json:"object"`
	Id         string                                    `json:"id"`
	Name       string                                    `json:"name"`
	Email      string                                    `json:"email"`
	ImageUrl   string                                    `json:"image_url"`
	Groups     *[]ConsumersProfilesListOutputItemsGroups `json:"groups,omitempty"`
	ConsumerId string                                    `json:"consumer_id"`
	CreatedAt  time.Time                                 `json:"created_at"`
	UpdatedAt  time.Time                                 `json:"updated_at"`
}

// ConsumersProfilesListOutputPagination represents the consumers profiles list output pagination type.
type ConsumersProfilesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConsumersProfilesListOutput represents the consumers profiles list output type.
type ConsumersProfilesListOutput struct {
	Items      []ConsumersProfilesListOutputItems    `json:"items"`
	Pagination ConsumersProfilesListOutputPagination `json:"pagination"`
}

// MapConsumersProfilesListOutputFromJSON deserializes JSON data into a ConsumersProfilesListOutput.
func MapConsumersProfilesListOutputFromJSON(data []byte) (*ConsumersProfilesListOutput, error) {
	var v ConsumersProfilesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesListOutputToJSON serializes a ConsumersProfilesListOutput to JSON.
func MapConsumersProfilesListOutputToJSON(v *ConsumersProfilesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersProfilesListQuery represents the consumers profiles list query type.
type ConsumersProfilesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapConsumersProfilesListQueryFromJSON deserializes JSON data into a ConsumersProfilesListQuery.
func MapConsumersProfilesListQueryFromJSON(data []byte) (*ConsumersProfilesListQuery, error) {
	var v ConsumersProfilesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersProfilesListQueryToJSON serializes a ConsumersProfilesListQuery to JSON.
func MapConsumersProfilesListQueryToJSON(v *ConsumersProfilesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
