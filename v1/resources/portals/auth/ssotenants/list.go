package ssotenants

import (
	"encoding/json"
	"time"
)

// PortalsAuthSsoTenantsListOutputItemsCounts represents the portals auth sso tenants list output items counts type.
type PortalsAuthSsoTenantsListOutputItemsCounts struct {
	Connections float64 `json:"connections"`
}

// PortalsAuthSsoTenantsListOutputItems represents the portals auth sso tenants list output items type.
type PortalsAuthSsoTenantsListOutputItems struct {
	Object    string                                     `json:"object"`
	Id        string                                     `json:"id"`
	Name      string                                     `json:"name"`
	Status    string                                     `json:"status"`
	ClientId  string                                     `json:"client_id"`
	Counts    PortalsAuthSsoTenantsListOutputItemsCounts `json:"counts"`
	CreatedAt time.Time                                  `json:"created_at"`
	UpdatedAt time.Time                                  `json:"updated_at"`
}

// PortalsAuthSsoTenantsListOutputPagination represents the portals auth sso tenants list output pagination type.
type PortalsAuthSsoTenantsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsAuthSsoTenantsListOutput represents the portals auth sso tenants list output type.
type PortalsAuthSsoTenantsListOutput struct {
	Items      []PortalsAuthSsoTenantsListOutputItems    `json:"items"`
	Pagination PortalsAuthSsoTenantsListOutputPagination `json:"pagination"`
}

// MapPortalsAuthSsoTenantsListOutputFromJSON deserializes JSON data into a PortalsAuthSsoTenantsListOutput.
func MapPortalsAuthSsoTenantsListOutputFromJSON(data []byte) (*PortalsAuthSsoTenantsListOutput, error) {
	var v PortalsAuthSsoTenantsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsListOutputToJSON serializes a PortalsAuthSsoTenantsListOutput to JSON.
func MapPortalsAuthSsoTenantsListOutputToJSON(v *PortalsAuthSsoTenantsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAuthSsoTenantsListQuery represents the portals auth sso tenants list query type.
type PortalsAuthSsoTenantsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapPortalsAuthSsoTenantsListQueryFromJSON deserializes JSON data into a PortalsAuthSsoTenantsListQuery.
func MapPortalsAuthSsoTenantsListQueryFromJSON(data []byte) (*PortalsAuthSsoTenantsListQuery, error) {
	var v PortalsAuthSsoTenantsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsListQueryToJSON serializes a PortalsAuthSsoTenantsListQuery to JSON.
func MapPortalsAuthSsoTenantsListQueryToJSON(v *PortalsAuthSsoTenantsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
