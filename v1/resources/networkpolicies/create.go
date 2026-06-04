package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesCreateOutputRulesPorts represents the network policies create output rules ports type.
type NetworkPoliciesCreateOutputRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesCreateOutputRules represents the network policies create output rules type.
type NetworkPoliciesCreateOutputRules struct {
	Object      string                                   `json:"object"`
	Id          string                                   `json:"id"`
	Effect      string                                   `json:"effect"`
	Direction   string                                   `json:"direction"`
	Cidrs       []string                                 `json:"cidrs"`
	Description *string                                  `json:"description,omitempty"`
	Enabled     bool                                     `json:"enabled"`
	Priority    float64                                  `json:"priority"`
	Ports       *[]NetworkPoliciesCreateOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesCreateOutput represents the network policies create output type.
type NetworkPoliciesCreateOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Status      string                             `json:"status"`
	Version     float64                            `json:"version"`
	Rules       []NetworkPoliciesCreateOutputRules `json:"rules"`
	FirewallIds *[]string                          `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
	ArchivedAt  *time.Time                         `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesCreateOutputFromJSON deserializes JSON data into a NetworkPoliciesCreateOutput.
func MapNetworkPoliciesCreateOutputFromJSON(data []byte) (*NetworkPoliciesCreateOutput, error) {
	var v NetworkPoliciesCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesCreateOutputToJSON serializes a NetworkPoliciesCreateOutput to JSON.
func MapNetworkPoliciesCreateOutputToJSON(v *NetworkPoliciesCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesCreateBodyRulesPorts represents the network policies create body rules ports type.
type NetworkPoliciesCreateBodyRulesPorts struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// NetworkPoliciesCreateBodyRules represents the network policies create body rules type.
type NetworkPoliciesCreateBodyRules struct {
	Effect      string                                 `json:"effect"`
	Direction   string                                 `json:"direction"`
	Cidrs       []string                               `json:"cidrs"`
	Description *string                                `json:"description,omitempty"`
	Enabled     bool                                   `json:"enabled"`
	Priority    float64                                `json:"priority"`
	Ports       *[]NetworkPoliciesCreateBodyRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesCreateBody represents the network policies create body type.
type NetworkPoliciesCreateBody struct {
	Name        string                            `json:"name"`
	Description *string                           `json:"description,omitempty"`
	Rules       *[]NetworkPoliciesCreateBodyRules `json:"rules,omitempty"`
}

// MapNetworkPoliciesCreateBodyFromJSON deserializes JSON data into a NetworkPoliciesCreateBody.
func MapNetworkPoliciesCreateBodyFromJSON(data []byte) (*NetworkPoliciesCreateBody, error) {
	var v NetworkPoliciesCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesCreateBodyToJSON serializes a NetworkPoliciesCreateBody to JSON.
func MapNetworkPoliciesCreateBodyToJSON(v *NetworkPoliciesCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
