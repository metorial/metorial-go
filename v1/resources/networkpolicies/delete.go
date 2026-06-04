package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesDeleteOutputRulesPorts represents the network policies delete output rules ports type.
type NetworkPoliciesDeleteOutputRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesDeleteOutputRules represents the network policies delete output rules type.
type NetworkPoliciesDeleteOutputRules struct {
	Object      string                                   `json:"object"`
	Id          string                                   `json:"id"`
	Effect      string                                   `json:"effect"`
	Direction   string                                   `json:"direction"`
	Cidrs       []string                                 `json:"cidrs"`
	Description *string                                  `json:"description,omitempty"`
	Enabled     bool                                     `json:"enabled"`
	Priority    float64                                  `json:"priority"`
	Ports       *[]NetworkPoliciesDeleteOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesDeleteOutput represents the network policies delete output type.
type NetworkPoliciesDeleteOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Status      string                             `json:"status"`
	Version     float64                            `json:"version"`
	Rules       []NetworkPoliciesDeleteOutputRules `json:"rules"`
	FirewallIds *[]string                          `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
	ArchivedAt  *time.Time                         `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesDeleteOutputFromJSON deserializes JSON data into a NetworkPoliciesDeleteOutput.
func MapNetworkPoliciesDeleteOutputFromJSON(data []byte) (*NetworkPoliciesDeleteOutput, error) {
	var v NetworkPoliciesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesDeleteOutputToJSON serializes a NetworkPoliciesDeleteOutput to JSON.
func MapNetworkPoliciesDeleteOutputToJSON(v *NetworkPoliciesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
