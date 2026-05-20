package profiles

import (
	"encoding/json"
)

// ConsumersProfilesListOutputPagination represents the consumers profiles list output pagination type.
type ConsumersProfilesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConsumersProfilesListOutput represents the consumers profiles list output type.
type ConsumersProfilesListOutput struct {
	Items      []map[string]any                      `json:"items"`
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
