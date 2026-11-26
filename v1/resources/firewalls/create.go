package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsCreateOutputNetworkPoliciesRulesPorts represents the firewalls create output network policies rules ports type.
type FirewallsCreateOutputNetworkPoliciesRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// FirewallsCreateOutputNetworkPoliciesRules represents the firewalls create output network policies rules type.
type FirewallsCreateOutputNetworkPoliciesRules struct {
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
	Ports *[]FirewallsCreateOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsCreateOutputNetworkPolicies - Network policy attached to the firewall.
type FirewallsCreateOutputNetworkPolicies struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Version number of the policy. Increments with every rule change.
	Version float64 `json:"version"`
	// Rules - Rules in the attached network policy.
	Rules []FirewallsCreateOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsCreateOutput represents the firewalls create output type.
type FirewallsCreateOutput struct {
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
	NetworkPolicies []FirewallsCreateOutputNetworkPolicies `json:"network_policies"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapFirewallsCreateOutputFromJSON deserializes JSON data into a FirewallsCreateOutput.
func MapFirewallsCreateOutputFromJSON(data []byte) (*FirewallsCreateOutput, error) {
	var v FirewallsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsCreateOutputToJSON serializes a FirewallsCreateOutput to JSON.
func MapFirewallsCreateOutputToJSON(v *FirewallsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsCreateBody represents the firewalls create body type.
type FirewallsCreateBody struct {
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance. Generated from the name when omitted.
	Slug *string `json:"slug,omitempty"`
	// NetworkPolicyIds - Network policies to attach to the firewall, in order.
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// MapFirewallsCreateBodyFromJSON deserializes JSON data into a FirewallsCreateBody.
func MapFirewallsCreateBodyFromJSON(data []byte) (*FirewallsCreateBody, error) {
	var v FirewallsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsCreateBodyToJSON serializes a FirewallsCreateBody to JSON.
func MapFirewallsCreateBodyToJSON(v *FirewallsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
