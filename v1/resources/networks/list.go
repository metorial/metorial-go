package networks

import (
	"encoding/json"
	"time"
)

// NetworksListOutputItemsPublicIps represents the networks list output items public ips type.
type NetworksListOutputItemsPublicIps struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Ip        string    `json:"ip"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NetworksListOutputItems represents the networks list output items type.
type NetworksListOutputItems struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
	PublicIps   []NetworksListOutputItemsPublicIps `json:"public_ips"`
}

// NetworksListOutputPagination represents the networks list output pagination type.
type NetworksListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// NetworksListOutput represents the networks list output type.
type NetworksListOutput struct {
	Items      []NetworksListOutputItems    `json:"items"`
	Pagination NetworksListOutputPagination `json:"pagination"`
}

// MapNetworksListOutputFromJSON deserializes JSON data into a NetworksListOutput.
func MapNetworksListOutputFromJSON(data []byte) (*NetworksListOutput, error) {
	var v NetworksListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworksListOutputToJSON serializes a NetworksListOutput to JSON.
func MapNetworksListOutputToJSON(v *NetworksListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworksListQueryCreatedAt - Filter network creation time by date range
type NetworksListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for network creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for network creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworksListQueryUpdatedAt - Filter network last update time by date range
type NetworksListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for network last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for network last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworksListQuery represents the networks list query type.
type NetworksListQuery struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	EnclaveId  *any     `json:"enclave_id,omitempty"`
	// CreatedAt - Filter network creation time by date range
	CreatedAt *NetworksListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter network last update time by date range
	UpdatedAt *NetworksListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapNetworksListQueryFromJSON deserializes JSON data into a NetworksListQuery.
func MapNetworksListQueryFromJSON(data []byte) (*NetworksListQuery, error) {
	var v NetworksListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworksListQueryToJSON serializes a NetworksListQuery to JSON.
func MapNetworksListQueryToJSON(v *NetworksListQuery) ([]byte, error) {
	return json.Marshal(v)
}
