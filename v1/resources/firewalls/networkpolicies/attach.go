package networkpolicies

import (
	"encoding/json"
	"time"
)

// FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRulesPorts represents the firewalls network policies attach output network policies rules ports type.
type FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRules represents the firewalls network policies attach output network policies rules type.
type FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRules struct {
	Object      string                                                           `json:"object"`
	Id          string                                                           `json:"id"`
	Effect      string                                                           `json:"effect"`
	Direction   string                                                           `json:"direction"`
	Cidrs       []string                                                         `json:"cidrs"`
	Description *string                                                          `json:"description,omitempty"`
	Enabled     bool                                                             `json:"enabled"`
	Priority    float64                                                          `json:"priority"`
	Ports       *[]FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsNetworkPoliciesAttachOutputNetworkPolicies represents the firewalls network policies attach output network policies type.
type FirewallsNetworkPoliciesAttachOutputNetworkPolicies struct {
	Object  string                                                     `json:"object"`
	Id      string                                                     `json:"id"`
	Name    string                                                     `json:"name"`
	Version float64                                                    `json:"version"`
	Rules   []FirewallsNetworkPoliciesAttachOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsNetworkPoliciesAttachOutput represents the firewalls network policies attach output type.
type FirewallsNetworkPoliciesAttachOutput struct {
	Object          string                                                `json:"object"`
	Id              string                                                `json:"id"`
	Slug            string                                                `json:"slug"`
	Name            string                                                `json:"name"`
	Description     *string                                               `json:"description,omitempty"`
	Status          string                                                `json:"status"`
	NetworkId       string                                                `json:"network_id"`
	NetworkPolicies []FirewallsNetworkPoliciesAttachOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                                             `json:"created_at"`
	UpdatedAt       time.Time                                             `json:"updated_at"`
	ArchivedAt      *time.Time                                            `json:"archived_at,omitempty"`
}

// MapFirewallsNetworkPoliciesAttachOutputFromJSON deserializes JSON data into a FirewallsNetworkPoliciesAttachOutput.
func MapFirewallsNetworkPoliciesAttachOutputFromJSON(data []byte) (*FirewallsNetworkPoliciesAttachOutput, error) {
	var v FirewallsNetworkPoliciesAttachOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsNetworkPoliciesAttachOutputToJSON serializes a FirewallsNetworkPoliciesAttachOutput to JSON.
func MapFirewallsNetworkPoliciesAttachOutputToJSON(v *FirewallsNetworkPoliciesAttachOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsNetworkPoliciesAttachBody represents the firewalls network policies attach body type.
type FirewallsNetworkPoliciesAttachBody struct {
	NetworkPolicyId string   `json:"network_policy_id"`
	Position        *float64 `json:"position,omitempty"`
}

// MapFirewallsNetworkPoliciesAttachBodyFromJSON deserializes JSON data into a FirewallsNetworkPoliciesAttachBody.
func MapFirewallsNetworkPoliciesAttachBodyFromJSON(data []byte) (*FirewallsNetworkPoliciesAttachBody, error) {
	var v FirewallsNetworkPoliciesAttachBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsNetworkPoliciesAttachBodyToJSON serializes a FirewallsNetworkPoliciesAttachBody to JSON.
func MapFirewallsNetworkPoliciesAttachBodyToJSON(v *FirewallsNetworkPoliciesAttachBody) ([]byte, error) {
	return json.Marshal(v)
}
