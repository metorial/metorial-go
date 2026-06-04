package rules

import (
	"encoding/json"
)

// NetworkPoliciesRulesUpdateOutputPorts represents the network policies rules update output ports type.
type NetworkPoliciesRulesUpdateOutputPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesRulesUpdateOutput represents the network policies rules update output type.
type NetworkPoliciesRulesUpdateOutput struct {
	Object      string                                   `json:"object"`
	Id          string                                   `json:"id"`
	Effect      string                                   `json:"effect"`
	Direction   string                                   `json:"direction"`
	Cidrs       []string                                 `json:"cidrs"`
	Description *string                                  `json:"description,omitempty"`
	Enabled     bool                                     `json:"enabled"`
	Priority    float64                                  `json:"priority"`
	Ports       *[]NetworkPoliciesRulesUpdateOutputPorts `json:"ports,omitempty"`
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
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// NetworkPoliciesRulesUpdateBody represents the network policies rules update body type.
type NetworkPoliciesRulesUpdateBody struct {
	Effect      string                                 `json:"effect"`
	Direction   string                                 `json:"direction"`
	Cidrs       []string                               `json:"cidrs"`
	Description *string                                `json:"description,omitempty"`
	Enabled     bool                                   `json:"enabled"`
	Priority    float64                                `json:"priority"`
	Ports       *[]NetworkPoliciesRulesUpdateBodyPorts `json:"ports,omitempty"`
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
