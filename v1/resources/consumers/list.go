package consumers

import (
	"encoding/json"
	"time"
)

// ConsumersListOutputItems represents the consumers list output items type.
type ConsumersListOutputItems struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConsumersListOutputPagination represents the consumers list output pagination type.
type ConsumersListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConsumersListOutput represents the consumers list output type.
type ConsumersListOutput struct {
	Items      []ConsumersListOutputItems    `json:"items"`
	Pagination ConsumersListOutputPagination `json:"pagination"`
}

// MapConsumersListOutputFromJSON deserializes JSON data into a ConsumersListOutput.
func MapConsumersListOutputFromJSON(data []byte) (*ConsumersListOutput, error) {
	var v ConsumersListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersListOutputToJSON serializes a ConsumersListOutput to JSON.
func MapConsumersListOutputToJSON(v *ConsumersListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumersListQuery represents the consumers list query type.
type ConsumersListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Id     *string  `json:"id,omitempty"`
}

// MapConsumersListQueryFromJSON deserializes JSON data into a ConsumersListQuery.
func MapConsumersListQueryFromJSON(data []byte) (*ConsumersListQuery, error) {
	var v ConsumersListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumersListQueryToJSON serializes a ConsumersListQuery to JSON.
func MapConsumersListQueryToJSON(v *ConsumersListQuery) ([]byte, error) {
	return json.Marshal(v)
}
