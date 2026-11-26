package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesUpdateOutputRulesPorts represents the network policies update output rules ports type.
type NetworkPoliciesUpdateOutputRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesUpdateOutputRules - Rule in the network policy.
type NetworkPoliciesUpdateOutputRules struct {
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
	Ports *[]NetworkPoliciesUpdateOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesUpdateOutput represents the network policies update output type.
type NetworkPoliciesUpdateOutput struct {
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
	Rules []NetworkPoliciesUpdateOutputRules `json:"rules"`
	// FirewallIds - Firewalls this policy is attached to.
	FirewallIds *[]string `json:"firewall_ids,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesUpdateOutputFromJSON deserializes JSON data into a NetworkPoliciesUpdateOutput.
func MapNetworkPoliciesUpdateOutputFromJSON(data []byte) (*NetworkPoliciesUpdateOutput, error) {
	var v NetworkPoliciesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesUpdateOutputToJSON serializes a NetworkPoliciesUpdateOutput to JSON.
func MapNetworkPoliciesUpdateOutputToJSON(v *NetworkPoliciesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesUpdateBodyRulesPorts represents the network policies update body rules ports type.
type NetworkPoliciesUpdateBodyRulesPorts struct {
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesUpdateBodyRules represents the network policies update body rules type.
type NetworkPoliciesUpdateBodyRules struct {
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
	Ports *[]NetworkPoliciesUpdateBodyRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesUpdateBody represents the network policies update body type.
type NetworkPoliciesUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Rules - Rules in the network policy. Replaces existing rules.
	Rules *[]NetworkPoliciesUpdateBodyRules `json:"rules,omitempty"`
}

// MapNetworkPoliciesUpdateBodyFromJSON deserializes JSON data into a NetworkPoliciesUpdateBody.
func MapNetworkPoliciesUpdateBodyFromJSON(data []byte) (*NetworkPoliciesUpdateBody, error) {
	var v NetworkPoliciesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesUpdateBodyToJSON serializes a NetworkPoliciesUpdateBody to JSON.
func MapNetworkPoliciesUpdateBodyToJSON(v *NetworkPoliciesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
