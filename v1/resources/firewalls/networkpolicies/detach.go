package networkpolicies

import (
	"encoding/json"
	"time"
)

// FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRulesPorts represents the firewalls network policies detach output network policies rules ports type.
type FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRules represents the firewalls network policies detach output network policies rules type.
type FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRules struct {
	Object      string                                                           `json:"object"`
	Id          string                                                           `json:"id"`
	Effect      string                                                           `json:"effect"`
	Direction   string                                                           `json:"direction"`
	Cidrs       []string                                                         `json:"cidrs"`
	Description *string                                                          `json:"description,omitempty"`
	Enabled     bool                                                             `json:"enabled"`
	Priority    float64                                                          `json:"priority"`
	Ports       *[]FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsNetworkPoliciesDetachOutputNetworkPolicies represents the firewalls network policies detach output network policies type.
type FirewallsNetworkPoliciesDetachOutputNetworkPolicies struct {
	Object  string                                                     `json:"object"`
	Id      string                                                     `json:"id"`
	Name    string                                                     `json:"name"`
	Version float64                                                    `json:"version"`
	Rules   []FirewallsNetworkPoliciesDetachOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsNetworkPoliciesDetachOutput represents the firewalls network policies detach output type.
type FirewallsNetworkPoliciesDetachOutput struct {
	Object          string                                                `json:"object"`
	Id              string                                                `json:"id"`
	Slug            string                                                `json:"slug"`
	Name            string                                                `json:"name"`
	Description     *string                                               `json:"description,omitempty"`
	Status          string                                                `json:"status"`
	NetworkId       string                                                `json:"network_id"`
	NetworkPolicies []FirewallsNetworkPoliciesDetachOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                                             `json:"created_at"`
	UpdatedAt       time.Time                                             `json:"updated_at"`
	ArchivedAt      *time.Time                                            `json:"archived_at,omitempty"`
}

// MapFirewallsNetworkPoliciesDetachOutputFromJSON deserializes JSON data into a FirewallsNetworkPoliciesDetachOutput.
func MapFirewallsNetworkPoliciesDetachOutputFromJSON(data []byte) (*FirewallsNetworkPoliciesDetachOutput, error) {
	var v FirewallsNetworkPoliciesDetachOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsNetworkPoliciesDetachOutputToJSON serializes a FirewallsNetworkPoliciesDetachOutput to JSON.
func MapFirewallsNetworkPoliciesDetachOutputToJSON(v *FirewallsNetworkPoliciesDetachOutput) ([]byte, error) {
	return json.Marshal(v)
}
