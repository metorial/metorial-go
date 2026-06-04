package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesGetOutputRulesPorts represents the network policies get output rules ports type.
type NetworkPoliciesGetOutputRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesGetOutputRules represents the network policies get output rules type.
type NetworkPoliciesGetOutputRules struct {
	Object      string                                `json:"object"`
	Id          string                                `json:"id"`
	Effect      string                                `json:"effect"`
	Direction   string                                `json:"direction"`
	Cidrs       []string                              `json:"cidrs"`
	Description *string                               `json:"description,omitempty"`
	Enabled     bool                                  `json:"enabled"`
	Priority    float64                               `json:"priority"`
	Ports       *[]NetworkPoliciesGetOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesGetOutput represents the network policies get output type.
type NetworkPoliciesGetOutput struct {
	Object      string                          `json:"object"`
	Id          string                          `json:"id"`
	Name        string                          `json:"name"`
	Description *string                         `json:"description,omitempty"`
	Status      string                          `json:"status"`
	Version     float64                         `json:"version"`
	Rules       []NetworkPoliciesGetOutputRules `json:"rules"`
	FirewallIds *[]string                       `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                       `json:"created_at"`
	UpdatedAt   time.Time                       `json:"updated_at"`
	ArchivedAt  *time.Time                      `json:"archived_at,omitempty"`
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
