package rules

import (
	"encoding/json"
)

// NetworkPoliciesRulesCreateOutputPorts represents the network policies rules create output ports type.
type NetworkPoliciesRulesCreateOutputPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesRulesCreateOutput - Rule in the network policy.
type NetworkPoliciesRulesCreateOutput struct {
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
	Ports *[]NetworkPoliciesRulesCreateOutputPorts `json:"ports,omitempty"`
}

// MapNetworkPoliciesRulesCreateOutputFromJSON deserializes JSON data into a NetworkPoliciesRulesCreateOutput.
func MapNetworkPoliciesRulesCreateOutputFromJSON(data []byte) (*NetworkPoliciesRulesCreateOutput, error) {
	var v NetworkPoliciesRulesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesRulesCreateOutputToJSON serializes a NetworkPoliciesRulesCreateOutput to JSON.
func MapNetworkPoliciesRulesCreateOutputToJSON(v *NetworkPoliciesRulesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesRulesCreateBodyPorts represents the network policies rules create body ports type.
type NetworkPoliciesRulesCreateBodyPorts struct {
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesRulesCreateBody represents the network policies rules create body type.
type NetworkPoliciesRulesCreateBody struct {
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
	Ports *[]NetworkPoliciesRulesCreateBodyPorts `json:"ports,omitempty"`
}

// MapNetworkPoliciesRulesCreateBodyFromJSON deserializes JSON data into a NetworkPoliciesRulesCreateBody.
func MapNetworkPoliciesRulesCreateBodyFromJSON(data []byte) (*NetworkPoliciesRulesCreateBody, error) {
	var v NetworkPoliciesRulesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesRulesCreateBodyToJSON serializes a NetworkPoliciesRulesCreateBody to JSON.
func MapNetworkPoliciesRulesCreateBodyToJSON(v *NetworkPoliciesRulesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
