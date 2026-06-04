package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesListOutputItemsRulesPorts represents the network policies list output items rules ports type.
type NetworkPoliciesListOutputItemsRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesListOutputItemsRules represents the network policies list output items rules type.
type NetworkPoliciesListOutputItemsRules struct {
	Object      string                                      `json:"object"`
	Id          string                                      `json:"id"`
	Effect      string                                      `json:"effect"`
	Direction   string                                      `json:"direction"`
	Cidrs       []string                                    `json:"cidrs"`
	Description *string                                     `json:"description,omitempty"`
	Enabled     bool                                        `json:"enabled"`
	Priority    float64                                     `json:"priority"`
	Ports       *[]NetworkPoliciesListOutputItemsRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesListOutputItems represents the network policies list output items type.
type NetworkPoliciesListOutputItems struct {
	Object      string                                `json:"object"`
	Id          string                                `json:"id"`
	Name        string                                `json:"name"`
	Description *string                               `json:"description,omitempty"`
	Status      string                                `json:"status"`
	Version     float64                               `json:"version"`
	Rules       []NetworkPoliciesListOutputItemsRules `json:"rules"`
	FirewallIds *[]string                             `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                             `json:"created_at"`
	UpdatedAt   time.Time                             `json:"updated_at"`
	ArchivedAt  *time.Time                            `json:"archived_at,omitempty"`
}

// NetworkPoliciesListOutputPagination represents the network policies list output pagination type.
type NetworkPoliciesListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// NetworkPoliciesListOutput represents the network policies list output type.
type NetworkPoliciesListOutput struct {
	Items      []NetworkPoliciesListOutputItems    `json:"items"`
	Pagination NetworkPoliciesListOutputPagination `json:"pagination"`
}

// MapNetworkPoliciesListOutputFromJSON deserializes JSON data into a NetworkPoliciesListOutput.
func MapNetworkPoliciesListOutputFromJSON(data []byte) (*NetworkPoliciesListOutput, error) {
	var v NetworkPoliciesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesListOutputToJSON serializes a NetworkPoliciesListOutput to JSON.
func MapNetworkPoliciesListOutputToJSON(v *NetworkPoliciesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesListQueryCreatedAt - Filter network policy creation time by date range
type NetworkPoliciesListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for network policy creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for network policy creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworkPoliciesListQueryUpdatedAt - Filter network policy last update time by date range
type NetworkPoliciesListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for network policy last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for network policy last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworkPoliciesListQuery represents the network policies list query type.
type NetworkPoliciesListQuery struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	Status     *any     `json:"status,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	Search     *string  `json:"search,omitempty"`
	// CreatedAt - Filter network policy creation time by date range
	CreatedAt *NetworkPoliciesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter network policy last update time by date range
	UpdatedAt *NetworkPoliciesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapNetworkPoliciesListQueryFromJSON deserializes JSON data into a NetworkPoliciesListQuery.
func MapNetworkPoliciesListQueryFromJSON(data []byte) (*NetworkPoliciesListQuery, error) {
	var v NetworkPoliciesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesListQueryToJSON serializes a NetworkPoliciesListQuery to JSON.
func MapNetworkPoliciesListQueryToJSON(v *NetworkPoliciesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
