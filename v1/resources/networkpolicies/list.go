package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesListOutputItemsRulesPorts represents the network policies list output items rules ports type.
type NetworkPoliciesListOutputItemsRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesListOutputItemsRules - Rule in the network policy.
type NetworkPoliciesListOutputItemsRules struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Effect - `allow` permits matching traffic. `deny` blocks it.
	Effect string `json:"effect"`
	// Direction - `ingress` matches incoming traffic. `egress` matches outgoing traffic.
	Direction string `json:"direction"`
	// Cidrs - CIDR ranges the rule matches.
	Cidrs []string `json:"cidrs"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Enabled - Whether the rule is enforced. Disabled rules are kept but ignored.
	Enabled bool `json:"enabled"`
	// Priority - Priority of the rule. Higher values are evaluated first.
	Priority float64 `json:"priority"`
	// Ports - Port ranges the rule matches. Omitted matches all ports.
	Ports *[]NetworkPoliciesListOutputItemsRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesListOutputItems represents the network policies list output items type.
type NetworkPoliciesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Status - Current lifecycle status of the network policy. `archived` once you archive the policy.
	Status string `json:"status"`
	// Version - Version number of the policy. Increments with every rule change.
	Version float64 `json:"version"`
	// Rules - Rules in the current version of the network policy.
	Rules []NetworkPoliciesListOutputItemsRules `json:"rules"`
	// FirewallIds - Firewalls this policy is attached to.
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// NetworkPoliciesListOutputPagination represents the network policies list output pagination type.
type NetworkPoliciesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// NetworkPoliciesListOutput represents the network policies list output type.
type NetworkPoliciesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
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

// NetworkPoliciesListQueryCreatedAt - Only return results whose network policy creation time is in this range. Set `gt`, `lt`, or both.
type NetworkPoliciesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworkPoliciesListQueryUpdatedAt - Only return results whose network policy last update time is in this range. Set `gt`, `lt`, or both.
type NetworkPoliciesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// NetworkPoliciesListQuery represents the network policies list query type.
type NetworkPoliciesListQuery struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Id - Only return network policies with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Status - Only return network policies with one of these statuses.
	Status *any `json:"status,omitempty"`
	// FirewallId - Only return network policies attached to one of these firewalls.
	FirewallId *any `json:"firewall_id,omitempty"`
	// Search - Filter network policies by name or description.
	Search *string `json:"search,omitempty"`
	// CreatedAt - Only return results whose network policy creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *NetworkPoliciesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose network policy last update time is in this range. Set `gt`, `lt`, or both.
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
