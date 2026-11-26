package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsListOutputItemsNetworkPoliciesRulesPorts represents the firewalls list output items network policies rules ports type.
type FirewallsListOutputItemsNetworkPoliciesRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// FirewallsListOutputItemsNetworkPoliciesRules represents the firewalls list output items network policies rules type.
type FirewallsListOutputItemsNetworkPoliciesRules struct {
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
	Ports *[]FirewallsListOutputItemsNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsListOutputItemsNetworkPolicies - Network policy attached to the firewall.
type FirewallsListOutputItemsNetworkPolicies struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Version number of the policy. Increments with every rule change.
	Version float64 `json:"version"`
	// Rules - Rules in the attached network policy.
	Rules []FirewallsListOutputItemsNetworkPoliciesRules `json:"rules"`
}

// FirewallsListOutputItems represents the firewalls list output items type.
type FirewallsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Status - Current lifecycle status of the firewall. `archived` once you archive the firewall.
	Status string `json:"status"`
	// NetworkPolicies - Network policies attached to the firewall, in position order.
	NetworkPolicies []FirewallsListOutputItemsNetworkPolicies `json:"network_policies"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// FirewallsListOutputPagination represents the firewalls list output pagination type.
type FirewallsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// FirewallsListOutput represents the firewalls list output type.
type FirewallsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                        `json:"object"`
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

// FirewallsListQueryCreatedAt - Only return results whose firewall creation time is in this range. Set `gt`, `lt`, or both.
type FirewallsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// FirewallsListQueryUpdatedAt - Only return results whose firewall last update time is in this range. Set `gt`, `lt`, or both.
type FirewallsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// FirewallsListQuery represents the firewalls list query type.
type FirewallsListQuery struct {
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
	// Search - Filter firewalls by name or description.
	Search *string `json:"search,omitempty"`
	// Id - Only return firewalls with one of these IDs.
	Id *any `json:"id,omitempty"`
	// Slug - Only return firewalls with one of these slugs.
	Slug *any `json:"slug,omitempty"`
	// Status - Only return firewalls with one of these statuses.
	Status *any `json:"status,omitempty"`
	// EnclaveId - Only return firewalls applied to one of these enclaves.
	EnclaveId *any `json:"enclave_id,omitempty"`
	// ProviderId - Only return firewalls applied to one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// NetworkPolicyId - Only return firewalls with one of these network policies attached.
	NetworkPolicyId *any `json:"network_policy_id,omitempty"`
	// MachineId - Only return firewalls applied to one of these sandbox machines.
	MachineId *any `json:"machine_id,omitempty"`
	// CreatedAt - Only return results whose firewall creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *FirewallsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose firewall last update time is in this range. Set `gt`, `lt`, or both.
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
