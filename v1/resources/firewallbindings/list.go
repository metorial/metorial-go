package firewallbindings

import (
	"encoding/json"
	"time"
)

// FirewallBindingsListOutputItemsFirewall represents the firewall bindings list output items firewall type.
type FirewallBindingsListOutputItemsFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// FirewallBindingsListOutputItemsTarget represents the firewall bindings list output items target type.
type FirewallBindingsListOutputItemsTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// FirewallBindingsListOutputItems represents the firewall bindings list output items type.
type FirewallBindingsListOutputItems struct {
	Object     string                                  `json:"object"`
	Id         string                                  `json:"id"`
	TargetType string                                  `json:"target_type"`
	Firewall   FirewallBindingsListOutputItemsFirewall `json:"firewall"`
	Target     *FirewallBindingsListOutputItemsTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                               `json:"created_at"`
}

// FirewallBindingsListOutputPagination represents the firewall bindings list output pagination type.
type FirewallBindingsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// FirewallBindingsListOutput represents the firewall bindings list output type.
type FirewallBindingsListOutput struct {
	Items      []FirewallBindingsListOutputItems    `json:"items"`
	Pagination FirewallBindingsListOutputPagination `json:"pagination"`
}

// MapFirewallBindingsListOutputFromJSON deserializes JSON data into a FirewallBindingsListOutput.
func MapFirewallBindingsListOutputFromJSON(data []byte) (*FirewallBindingsListOutput, error) {
	var v FirewallBindingsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsListOutputToJSON serializes a FirewallBindingsListOutput to JSON.
func MapFirewallBindingsListOutputToJSON(v *FirewallBindingsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallBindingsListQueryCreatedAt - Filter firewall binding creation time by date range
type FirewallBindingsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for firewall binding creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for firewall binding creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// FirewallBindingsListQuery represents the firewall bindings list query type.
type FirewallBindingsListQuery struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	EnclaveId  *any     `json:"enclave_id,omitempty"`
	ProviderId *any     `json:"provider_id,omitempty"`
	NetworkId  *any     `json:"network_id,omitempty"`
	TargetType *any     `json:"target_type,omitempty"`
	// CreatedAt - Filter firewall binding creation time by date range
	CreatedAt *FirewallBindingsListQueryCreatedAt `json:"created_at,omitempty"`
}

// MapFirewallBindingsListQueryFromJSON deserializes JSON data into a FirewallBindingsListQuery.
func MapFirewallBindingsListQueryFromJSON(data []byte) (*FirewallBindingsListQuery, error) {
	var v FirewallBindingsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsListQueryToJSON serializes a FirewallBindingsListQuery to JSON.
func MapFirewallBindingsListQueryToJSON(v *FirewallBindingsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
