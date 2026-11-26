package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsUpdateOutputNetworkPoliciesRulesPorts represents the firewalls update output network policies rules ports type.
type FirewallsUpdateOutputNetworkPoliciesRulesPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// FirewallsUpdateOutputNetworkPoliciesRules represents the firewalls update output network policies rules type.
type FirewallsUpdateOutputNetworkPoliciesRules struct {
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
	Ports *[]FirewallsUpdateOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsUpdateOutputNetworkPolicies - Network policy attached to the firewall.
type FirewallsUpdateOutputNetworkPolicies struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Version number of the policy. Increments with every rule change.
	Version float64 `json:"version"`
	// Rules - Rules in the attached network policy.
	Rules []FirewallsUpdateOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsUpdateOutput represents the firewalls update output type.
type FirewallsUpdateOutput struct {
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
	NetworkPolicies []FirewallsUpdateOutputNetworkPolicies `json:"network_policies"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// MapFirewallsUpdateOutputFromJSON deserializes JSON data into a FirewallsUpdateOutput.
func MapFirewallsUpdateOutputFromJSON(data []byte) (*FirewallsUpdateOutput, error) {
	var v FirewallsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsUpdateOutputToJSON serializes a FirewallsUpdateOutput to JSON.
func MapFirewallsUpdateOutputToJSON(v *FirewallsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsUpdateBody represents the firewalls update body type.
type FirewallsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance.
	Slug *string `json:"slug,omitempty"`
	// NetworkPolicyIds - Network policies attached to the firewall, in order. Replaces existing attachments.
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// MapFirewallsUpdateBodyFromJSON deserializes JSON data into a FirewallsUpdateBody.
func MapFirewallsUpdateBodyFromJSON(data []byte) (*FirewallsUpdateBody, error) {
	var v FirewallsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsUpdateBodyToJSON serializes a FirewallsUpdateBody to JSON.
func MapFirewallsUpdateBodyToJSON(v *FirewallsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
