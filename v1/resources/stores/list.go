package stores

import (
	"encoding/json"
	"time"
)

// StoresListOutputItems represents the stores list output items type.
type StoresListOutputItems struct {
	// Object - String representing the object's type
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Access    string    `json:"access"`
	ItemCount float64   `json:"item_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// StoresListOutputPagination represents the stores list output pagination type.
type StoresListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// StoresListOutput represents the stores list output type.
type StoresListOutput struct {
	Items      []StoresListOutputItems    `json:"items"`
	Pagination StoresListOutputPagination `json:"pagination"`
}

// MapStoresListOutputFromJSON deserializes JSON data into a StoresListOutput.
func MapStoresListOutputFromJSON(data []byte) (*StoresListOutput, error) {
	var v StoresListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresListOutputToJSON serializes a StoresListOutput to JSON.
func MapStoresListOutputToJSON(v *StoresListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// StoresListQueryCreatedAt - Filter Filter by creation time by date range
type StoresListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for Filter by creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresListQueryUpdatedAt - Filter Filter by update time by date range
type StoresListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for Filter by update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for Filter by update time
	Lt *time.Time `json:"lt,omitempty"`
}

// StoresListQuery represents the stores list query type.
type StoresListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Id - Filter by store ID
	Id *any `json:"id,omitempty"`
	// CreatedAt - Filter Filter by creation time by date range
	CreatedAt *StoresListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter Filter by update time by date range
	UpdatedAt *StoresListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapStoresListQueryFromJSON deserializes JSON data into a StoresListQuery.
func MapStoresListQueryFromJSON(data []byte) (*StoresListQuery, error) {
	var v StoresListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapStoresListQueryToJSON serializes a StoresListQuery to JSON.
func MapStoresListQueryToJSON(v *StoresListQuery) ([]byte, error) {
	return json.Marshal(v)
}
