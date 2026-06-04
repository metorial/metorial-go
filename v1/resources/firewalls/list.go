package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsListOutputItemsNetworkPoliciesRulesPorts represents the firewalls list output items network policies rules ports type.
type FirewallsListOutputItemsNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsListOutputItemsNetworkPoliciesRules represents the firewalls list output items network policies rules type.
type FirewallsListOutputItemsNetworkPoliciesRules struct {
	Object      string                                               `json:"object"`
	Id          string                                               `json:"id"`
	Effect      string                                               `json:"effect"`
	Direction   string                                               `json:"direction"`
	Cidrs       []string                                             `json:"cidrs"`
	Description *string                                              `json:"description,omitempty"`
	Enabled     bool                                                 `json:"enabled"`
	Priority    float64                                              `json:"priority"`
	Ports       *[]FirewallsListOutputItemsNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsListOutputItemsNetworkPolicies represents the firewalls list output items network policies type.
type FirewallsListOutputItemsNetworkPolicies struct {
	Object  string                                         `json:"object"`
	Id      string                                         `json:"id"`
	Name    string                                         `json:"name"`
	Version float64                                        `json:"version"`
	Rules   []FirewallsListOutputItemsNetworkPoliciesRules `json:"rules"`
}

// FirewallsListOutputItems represents the firewalls list output items type.
type FirewallsListOutputItems struct {
	Object          string                                    `json:"object"`
	Id              string                                    `json:"id"`
	Slug            string                                    `json:"slug"`
	Name            string                                    `json:"name"`
	Description     *string                                   `json:"description,omitempty"`
	Status          string                                    `json:"status"`
	NetworkId       string                                    `json:"network_id"`
	NetworkPolicies []FirewallsListOutputItemsNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                                 `json:"created_at"`
	UpdatedAt       time.Time                                 `json:"updated_at"`
	ArchivedAt      *time.Time                                `json:"archived_at,omitempty"`
}

// FirewallsListOutputPagination represents the firewalls list output pagination type.
type FirewallsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// FirewallsListOutput represents the firewalls list output type.
type FirewallsListOutput struct {
	Items      []FirewallsListOutputItems    `json:"items"`
	Pagination FirewallsListOutputPagination `json:"pagination"`
}

// MapFirewallsListOutputFromJSON deserializes JSON data into a FirewallsListOutput.
func MapFirewallsListOutputFromJSON(data []byte) (*FirewallsListOutput, error) {
	var v FirewallsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsListOutputToJSON serializes a FirewallsListOutput to JSON.
func MapFirewallsListOutputToJSON(v *FirewallsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsListQueryCreatedAt - Filter firewall creation time by date range
type FirewallsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for firewall creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for firewall creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// FirewallsListQueryUpdatedAt - Filter firewall last update time by date range
type FirewallsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for firewall last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for firewall last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// FirewallsListQuery represents the firewalls list query type.
type FirewallsListQuery struct {
	Limit           *float64 `json:"limit,omitempty"`
	After           *string  `json:"after,omitempty"`
	Before          *string  `json:"before,omitempty"`
	Cursor          *string  `json:"cursor,omitempty"`
	Order           *string  `json:"order,omitempty"`
	Id              *any     `json:"id,omitempty"`
	Slug            *any     `json:"slug,omitempty"`
	Status          *any     `json:"status,omitempty"`
	NetworkId       *any     `json:"network_id,omitempty"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	ProviderId      *any     `json:"provider_id,omitempty"`
	NetworkPolicyId *any     `json:"network_policy_id,omitempty"`
	// CreatedAt - Filter firewall creation time by date range
	CreatedAt *FirewallsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter firewall last update time by date range
	UpdatedAt *FirewallsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapFirewallsListQueryFromJSON deserializes JSON data into a FirewallsListQuery.
func MapFirewallsListQueryFromJSON(data []byte) (*FirewallsListQuery, error) {
	var v FirewallsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsListQueryToJSON serializes a FirewallsListQuery to JSON.
func MapFirewallsListQueryToJSON(v *FirewallsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
