package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsDeleteOutputNetworkPoliciesRulesPorts represents the firewalls delete output network policies rules ports type.
type FirewallsDeleteOutputNetworkPoliciesRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// FirewallsDeleteOutputNetworkPoliciesRules represents the firewalls delete output network policies rules type.
type FirewallsDeleteOutputNetworkPoliciesRules struct {
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
	Ports *[]FirewallsDeleteOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsDeleteOutputNetworkPolicies - Network policy attached to the firewall.
type FirewallsDeleteOutputNetworkPolicies struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Version number of the policy. Increments with every rule change.
	Version float64 `json:"version"`
	// Rules - Rules in the attached network policy.
	Rules []FirewallsDeleteOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsDeleteOutput represents the firewalls delete output type.
type FirewallsDeleteOutput struct {
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
	NetworkPolicies []FirewallsDeleteOutputNetworkPolicies `json:"network_policies"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapFirewallsDeleteOutputFromJSON deserializes JSON data into a FirewallsDeleteOutput.
func MapFirewallsDeleteOutputFromJSON(data []byte) (*FirewallsDeleteOutput, error) {
	var v FirewallsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsDeleteOutputToJSON serializes a FirewallsDeleteOutput to JSON.
func MapFirewallsDeleteOutputToJSON(v *FirewallsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
