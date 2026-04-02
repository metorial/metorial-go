package consumersurfaces

import (
	"encoding/json"
	"time"
)

// ConsumerSurfacesListOutputItemsAuth represents the consumer surfaces list output items auth type.
type ConsumerSurfacesListOutputItemsAuth struct {
	Object                     string  `json:"object"`
	SessionExpiryTimeInSeconds float64 `json:"session_expiry_time_in_seconds"`
}

// ConsumerSurfacesListOutputItems represents the consumer surfaces list output items type.
type ConsumerSurfacesListOutputItems struct {
	Object      string                              `json:"object"`
	Id          string                              `json:"id"`
	Status      string                              `json:"status"`
	Name        string                              `json:"name"`
	Description *string                             `json:"description,omitempty"`
	Auth        ConsumerSurfacesListOutputItemsAuth `json:"auth"`
	CreatedAt   time.Time                           `json:"created_at"`
	UpdatedAt   time.Time                           `json:"updated_at"`
}

// ConsumerSurfacesListOutputPagination represents the consumer surfaces list output pagination type.
type ConsumerSurfacesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ConsumerSurfacesListOutput represents the consumer surfaces list output type.
type ConsumerSurfacesListOutput struct {
	Items      []ConsumerSurfacesListOutputItems    `json:"items"`
	Pagination ConsumerSurfacesListOutputPagination `json:"pagination"`
}

// MapConsumerSurfacesListOutputFromJSON deserializes JSON data into a ConsumerSurfacesListOutput.
func MapConsumerSurfacesListOutputFromJSON(data []byte) (*ConsumerSurfacesListOutput, error) {
	var v ConsumerSurfacesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesListOutputToJSON serializes a ConsumerSurfacesListOutput to JSON.
func MapConsumerSurfacesListOutputToJSON(v *ConsumerSurfacesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ConsumerSurfacesListQuery represents the consumer surfaces list query type.
type ConsumerSurfacesListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapConsumerSurfacesListQueryFromJSON deserializes JSON data into a ConsumerSurfacesListQuery.
func MapConsumerSurfacesListQueryFromJSON(data []byte) (*ConsumerSurfacesListQuery, error) {
	var v ConsumerSurfacesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapConsumerSurfacesListQueryToJSON serializes a ConsumerSurfacesListQuery to JSON.
func MapConsumerSurfacesListQueryToJSON(v *ConsumerSurfacesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
