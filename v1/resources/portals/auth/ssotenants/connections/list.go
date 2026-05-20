package connections

import (
	"encoding/json"
	"time"
)

// PortalsAuthSsoTenantsConnectionsListOutputItems represents the portals auth sso tenants connections list output items type.
type PortalsAuthSsoTenantsConnectionsListOutputItems struct {
	Object       string    `json:"object"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	ProviderType string    `json:"provider_type"`
	ProviderName *string   `json:"provider_name,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// PortalsAuthSsoTenantsConnectionsListOutputPagination represents the portals auth sso tenants connections list output pagination type.
type PortalsAuthSsoTenantsConnectionsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// PortalsAuthSsoTenantsConnectionsListOutput represents the portals auth sso tenants connections list output type.
type PortalsAuthSsoTenantsConnectionsListOutput struct {
	Items      []PortalsAuthSsoTenantsConnectionsListOutputItems    `json:"items"`
	Pagination PortalsAuthSsoTenantsConnectionsListOutputPagination `json:"pagination"`
}

// MapPortalsAuthSsoTenantsConnectionsListOutputFromJSON deserializes JSON data into a PortalsAuthSsoTenantsConnectionsListOutput.
func MapPortalsAuthSsoTenantsConnectionsListOutputFromJSON(data []byte) (*PortalsAuthSsoTenantsConnectionsListOutput, error) {
	var v PortalsAuthSsoTenantsConnectionsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsConnectionsListOutputToJSON serializes a PortalsAuthSsoTenantsConnectionsListOutput to JSON.
func MapPortalsAuthSsoTenantsConnectionsListOutputToJSON(v *PortalsAuthSsoTenantsConnectionsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAuthSsoTenantsConnectionsListQuery represents the portals auth sso tenants connections list query type.
type PortalsAuthSsoTenantsConnectionsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapPortalsAuthSsoTenantsConnectionsListQueryFromJSON deserializes JSON data into a PortalsAuthSsoTenantsConnectionsListQuery.
func MapPortalsAuthSsoTenantsConnectionsListQueryFromJSON(data []byte) (*PortalsAuthSsoTenantsConnectionsListQuery, error) {
	var v PortalsAuthSsoTenantsConnectionsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsConnectionsListQueryToJSON serializes a PortalsAuthSsoTenantsConnectionsListQuery to JSON.
func MapPortalsAuthSsoTenantsConnectionsListQueryToJSON(v *PortalsAuthSsoTenantsConnectionsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
