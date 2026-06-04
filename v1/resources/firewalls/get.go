package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsGetOutputNetworkPoliciesRulesPorts represents the firewalls get output network policies rules ports type.
type FirewallsGetOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsGetOutputNetworkPoliciesRules represents the firewalls get output network policies rules type.
type FirewallsGetOutputNetworkPoliciesRules struct {
	Object      string                                         `json:"object"`
	Id          string                                         `json:"id"`
	Effect      string                                         `json:"effect"`
	Direction   string                                         `json:"direction"`
	Cidrs       []string                                       `json:"cidrs"`
	Description *string                                        `json:"description,omitempty"`
	Enabled     bool                                           `json:"enabled"`
	Priority    float64                                        `json:"priority"`
	Ports       *[]FirewallsGetOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsGetOutputNetworkPolicies represents the firewalls get output network policies type.
type FirewallsGetOutputNetworkPolicies struct {
	Object  string                                   `json:"object"`
	Id      string                                   `json:"id"`
	Name    string                                   `json:"name"`
	Version float64                                  `json:"version"`
	Rules   []FirewallsGetOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsGetOutput represents the firewalls get output type.
type FirewallsGetOutput struct {
	Object          string                              `json:"object"`
	Id              string                              `json:"id"`
	Slug            string                              `json:"slug"`
	Name            string                              `json:"name"`
	Description     *string                             `json:"description,omitempty"`
	Status          string                              `json:"status"`
	NetworkId       string                              `json:"network_id"`
	NetworkPolicies []FirewallsGetOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                           `json:"created_at"`
	UpdatedAt       time.Time                           `json:"updated_at"`
	ArchivedAt      *time.Time                          `json:"archived_at,omitempty"`
}

// MapFirewallsGetOutputFromJSON deserializes JSON data into a FirewallsGetOutput.
func MapFirewallsGetOutputFromJSON(data []byte) (*FirewallsGetOutput, error) {
	var v FirewallsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsGetOutputToJSON serializes a FirewallsGetOutput to JSON.
func MapFirewallsGetOutputToJSON(v *FirewallsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
