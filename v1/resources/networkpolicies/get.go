package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesGetOutputRulesPorts represents the network policies get output rules ports type.
type NetworkPoliciesGetOutputRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesGetOutputRules - Rule in the network policy.
type NetworkPoliciesGetOutputRules struct {
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
	Ports *[]NetworkPoliciesGetOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesGetOutput represents the network policies get output type.
type NetworkPoliciesGetOutput struct {
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
	Rules []NetworkPoliciesGetOutputRules `json:"rules"`
	// FirewallIds - Firewalls this policy is attached to.
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesGetOutputFromJSON deserializes JSON data into a NetworkPoliciesGetOutput.
func MapNetworkPoliciesGetOutputFromJSON(data []byte) (*NetworkPoliciesGetOutput, error) {
	var v NetworkPoliciesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesGetOutputToJSON serializes a NetworkPoliciesGetOutput to JSON.
func MapNetworkPoliciesGetOutputToJSON(v *NetworkPoliciesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
