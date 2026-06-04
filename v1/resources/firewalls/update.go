package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsUpdateOutputNetworkPoliciesRulesPorts represents the firewalls update output network policies rules ports type.
type FirewallsUpdateOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsUpdateOutputNetworkPoliciesRules represents the firewalls update output network policies rules type.
type FirewallsUpdateOutputNetworkPoliciesRules struct {
	Object      string                                            `json:"object"`
	Id          string                                            `json:"id"`
	Effect      string                                            `json:"effect"`
	Direction   string                                            `json:"direction"`
	Cidrs       []string                                          `json:"cidrs"`
	Description *string                                           `json:"description,omitempty"`
	Enabled     bool                                              `json:"enabled"`
	Priority    float64                                           `json:"priority"`
	Ports       *[]FirewallsUpdateOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsUpdateOutputNetworkPolicies represents the firewalls update output network policies type.
type FirewallsUpdateOutputNetworkPolicies struct {
	Object  string                                      `json:"object"`
	Id      string                                      `json:"id"`
	Name    string                                      `json:"name"`
	Version float64                                     `json:"version"`
	Rules   []FirewallsUpdateOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsUpdateOutput represents the firewalls update output type.
type FirewallsUpdateOutput struct {
	Object          string                                 `json:"object"`
	Id              string                                 `json:"id"`
	Slug            string                                 `json:"slug"`
	Name            string                                 `json:"name"`
	Description     *string                                `json:"description,omitempty"`
	Status          string                                 `json:"status"`
	NetworkId       string                                 `json:"network_id"`
	NetworkPolicies []FirewallsUpdateOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                              `json:"created_at"`
	UpdatedAt       time.Time                              `json:"updated_at"`
	ArchivedAt      *time.Time                             `json:"archived_at,omitempty"`
}

// MapFirewallsUpdateOutputFromJSON deserializes JSON data into a FirewallsUpdateOutput.
func MapFirewallsUpdateOutputFromJSON(data []byte) (*FirewallsUpdateOutput, error) {
	var v FirewallsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsUpdateOutputToJSON serializes a FirewallsUpdateOutput to JSON.
func MapFirewallsUpdateOutputToJSON(v *FirewallsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsUpdateBody represents the firewalls update body type.
type FirewallsUpdateBody struct {
	Name             *string   `json:"name,omitempty"`
	Description      *string   `json:"description,omitempty"`
	Slug             *string   `json:"slug,omitempty"`
	NetworkPolicyIds *[]string `json:"network_policy_ids,omitempty"`
}

// MapFirewallsUpdateBodyFromJSON deserializes JSON data into a FirewallsUpdateBody.
func MapFirewallsUpdateBodyFromJSON(data []byte) (*FirewallsUpdateBody, error) {
	var v FirewallsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsUpdateBodyToJSON serializes a FirewallsUpdateBody to JSON.
func MapFirewallsUpdateBodyToJSON(v *FirewallsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
