package rules

import (
	"encoding/json"
)

// NetworkPoliciesRulesUpdateOutputPorts represents the network policies rules update output ports type.
type NetworkPoliciesRulesUpdateOutputPorts struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesRulesUpdateOutput - Rule in the network policy.
type NetworkPoliciesRulesUpdateOutput struct {
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
	Ports *[]NetworkPoliciesRulesUpdateOutputPorts `json:"ports,omitempty"`
}

// MapNetworkPoliciesRulesUpdateOutputFromJSON deserializes JSON data into a NetworkPoliciesRulesUpdateOutput.
func MapNetworkPoliciesRulesUpdateOutputFromJSON(data []byte) (*NetworkPoliciesRulesUpdateOutput, error) {
	var v NetworkPoliciesRulesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesRulesUpdateOutputToJSON serializes a NetworkPoliciesRulesUpdateOutput to JSON.
func MapNetworkPoliciesRulesUpdateOutputToJSON(v *NetworkPoliciesRulesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesRulesUpdateBodyPorts represents the network policies rules update body ports type.
type NetworkPoliciesRulesUpdateBodyPorts struct {
	// From - First port in the range, inclusive.
	From float64 `json:"from"`
	// To - Last port in the range, inclusive.
	To float64 `json:"to"`
}

// NetworkPoliciesRulesUpdateBody represents the network policies rules update body type.
type NetworkPoliciesRulesUpdateBody struct {
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
	Ports *[]NetworkPoliciesRulesUpdateBodyPorts `json:"ports,omitempty"`
}

// MapNetworkPoliciesRulesUpdateBodyFromJSON deserializes JSON data into a NetworkPoliciesRulesUpdateBody.
func MapNetworkPoliciesRulesUpdateBodyFromJSON(data []byte) (*NetworkPoliciesRulesUpdateBody, error) {
	var v NetworkPoliciesRulesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesRulesUpdateBodyToJSON serializes a NetworkPoliciesRulesUpdateBody to JSON.
func MapNetworkPoliciesRulesUpdateBodyToJSON(v *NetworkPoliciesRulesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
