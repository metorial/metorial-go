package rules

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesRulesDeleteOutputRulesPorts represents the network policies rules delete output rules ports type.
type NetworkPoliciesRulesDeleteOutputRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesRulesDeleteOutputRules represents the network policies rules delete output rules type.
type NetworkPoliciesRulesDeleteOutputRules struct {
	Object      string                                        `json:"object"`
	Id          string                                        `json:"id"`
	Effect      string                                        `json:"effect"`
	Direction   string                                        `json:"direction"`
	Cidrs       []string                                      `json:"cidrs"`
	Description *string                                       `json:"description,omitempty"`
	Enabled     bool                                          `json:"enabled"`
	Priority    float64                                       `json:"priority"`
	Ports       *[]NetworkPoliciesRulesDeleteOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesRulesDeleteOutput represents the network policies rules delete output type.
type NetworkPoliciesRulesDeleteOutput struct {
	Object      string                                  `json:"object"`
	Id          string                                  `json:"id"`
	Name        string                                  `json:"name"`
	Description *string                                 `json:"description,omitempty"`
	Status      string                                  `json:"status"`
	Version     float64                                 `json:"version"`
	Rules       []NetworkPoliciesRulesDeleteOutputRules `json:"rules"`
	FirewallIds *[]string                               `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                               `json:"created_at"`
	UpdatedAt   time.Time                               `json:"updated_at"`
	ArchivedAt  *time.Time                              `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesRulesDeleteOutputFromJSON deserializes JSON data into a NetworkPoliciesRulesDeleteOutput.
func MapNetworkPoliciesRulesDeleteOutputFromJSON(data []byte) (*NetworkPoliciesRulesDeleteOutput, error) {
	var v NetworkPoliciesRulesDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesRulesDeleteOutputToJSON serializes a NetworkPoliciesRulesDeleteOutput to JSON.
func MapNetworkPoliciesRulesDeleteOutputToJSON(v *NetworkPoliciesRulesDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
