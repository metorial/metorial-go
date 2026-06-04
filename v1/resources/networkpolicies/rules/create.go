package rules

import (
	"encoding/json"
)

// NetworkPoliciesRulesCreateOutputPorts represents the network policies rules create output ports type.
type NetworkPoliciesRulesCreateOutputPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesRulesCreateOutput represents the network policies rules create output type.
type NetworkPoliciesRulesCreateOutput struct {
	Object      string                                   `json:"object"`
	Id          string                                   `json:"id"`
	Effect      string                                   `json:"effect"`
	Direction   string                                   `json:"direction"`
	Cidrs       []string                                 `json:"cidrs"`
	Description *string                                  `json:"description,omitempty"`
	Enabled     bool                                     `json:"enabled"`
	Priority    float64                                  `json:"priority"`
	Ports       *[]NetworkPoliciesRulesCreateOutputPorts `json:"ports,omitempty"`
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
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// NetworkPoliciesRulesCreateBody represents the network policies rules create body type.
type NetworkPoliciesRulesCreateBody struct {
	Effect      string                                 `json:"effect"`
	Direction   string                                 `json:"direction"`
	Cidrs       []string                               `json:"cidrs"`
	Description *string                                `json:"description,omitempty"`
	Enabled     bool                                   `json:"enabled"`
	Priority    float64                                `json:"priority"`
	Ports       *[]NetworkPoliciesRulesCreateBodyPorts `json:"ports,omitempty"`
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
