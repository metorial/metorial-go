package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsDeleteOutputNetworkPoliciesRulesPorts represents the firewalls delete output network policies rules ports type.
type FirewallsDeleteOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsDeleteOutputNetworkPoliciesRules represents the firewalls delete output network policies rules type.
type FirewallsDeleteOutputNetworkPoliciesRules struct {
	Object      string                                            `json:"object"`
	Id          string                                            `json:"id"`
	Effect      string                                            `json:"effect"`
	Direction   string                                            `json:"direction"`
	Cidrs       []string                                          `json:"cidrs"`
	Description *string                                           `json:"description,omitempty"`
	Enabled     bool                                              `json:"enabled"`
	Priority    float64                                           `json:"priority"`
	Ports       *[]FirewallsDeleteOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsDeleteOutputNetworkPolicies represents the firewalls delete output network policies type.
type FirewallsDeleteOutputNetworkPolicies struct {
	Object  string                                      `json:"object"`
	Id      string                                      `json:"id"`
	Name    string                                      `json:"name"`
	Version float64                                     `json:"version"`
	Rules   []FirewallsDeleteOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsDeleteOutput represents the firewalls delete output type.
type FirewallsDeleteOutput struct {
	Object          string                                 `json:"object"`
	Id              string                                 `json:"id"`
	Slug            string                                 `json:"slug"`
	Name            string                                 `json:"name"`
	Description     *string                                `json:"description,omitempty"`
	Status          string                                 `json:"status"`
	NetworkId       string                                 `json:"network_id"`
	NetworkPolicies []FirewallsDeleteOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                              `json:"created_at"`
	UpdatedAt       time.Time                              `json:"updated_at"`
	ArchivedAt      *time.Time                             `json:"archived_at,omitempty"`
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
