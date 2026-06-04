package networkpolicies

import (
	"encoding/json"
	"time"
)

// NetworkPoliciesUpdateOutputRulesPorts represents the network policies update output rules ports type.
type NetworkPoliciesUpdateOutputRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// NetworkPoliciesUpdateOutputRules represents the network policies update output rules type.
type NetworkPoliciesUpdateOutputRules struct {
	Object      string                                   `json:"object"`
	Id          string                                   `json:"id"`
	Effect      string                                   `json:"effect"`
	Direction   string                                   `json:"direction"`
	Cidrs       []string                                 `json:"cidrs"`
	Description *string                                  `json:"description,omitempty"`
	Enabled     bool                                     `json:"enabled"`
	Priority    float64                                  `json:"priority"`
	Ports       *[]NetworkPoliciesUpdateOutputRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesUpdateOutput represents the network policies update output type.
type NetworkPoliciesUpdateOutput struct {
	Object      string                             `json:"object"`
	Id          string                             `json:"id"`
	Name        string                             `json:"name"`
	Description *string                            `json:"description,omitempty"`
	Status      string                             `json:"status"`
	Version     float64                            `json:"version"`
	Rules       []NetworkPoliciesUpdateOutputRules `json:"rules"`
	FirewallIds *[]string                          `json:"firewall_ids,omitempty"`
	CreatedAt   time.Time                          `json:"created_at"`
	UpdatedAt   time.Time                          `json:"updated_at"`
	ArchivedAt  *time.Time                         `json:"archived_at,omitempty"`
}

// MapNetworkPoliciesUpdateOutputFromJSON deserializes JSON data into a NetworkPoliciesUpdateOutput.
func MapNetworkPoliciesUpdateOutputFromJSON(data []byte) (*NetworkPoliciesUpdateOutput, error) {
	var v NetworkPoliciesUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesUpdateOutputToJSON serializes a NetworkPoliciesUpdateOutput to JSON.
func MapNetworkPoliciesUpdateOutputToJSON(v *NetworkPoliciesUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworkPoliciesUpdateBodyRulesPorts represents the network policies update body rules ports type.
type NetworkPoliciesUpdateBodyRulesPorts struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

// NetworkPoliciesUpdateBodyRules represents the network policies update body rules type.
type NetworkPoliciesUpdateBodyRules struct {
	Effect      string                                 `json:"effect"`
	Direction   string                                 `json:"direction"`
	Cidrs       []string                               `json:"cidrs"`
	Description *string                                `json:"description,omitempty"`
	Enabled     bool                                   `json:"enabled"`
	Priority    float64                                `json:"priority"`
	Ports       *[]NetworkPoliciesUpdateBodyRulesPorts `json:"ports,omitempty"`
}

// NetworkPoliciesUpdateBody represents the network policies update body type.
type NetworkPoliciesUpdateBody struct {
	Name        *string                           `json:"name,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Rules       *[]NetworkPoliciesUpdateBodyRules `json:"rules,omitempty"`
}

// MapNetworkPoliciesUpdateBodyFromJSON deserializes JSON data into a NetworkPoliciesUpdateBody.
func MapNetworkPoliciesUpdateBodyFromJSON(data []byte) (*NetworkPoliciesUpdateBody, error) {
	var v NetworkPoliciesUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworkPoliciesUpdateBodyToJSON serializes a NetworkPoliciesUpdateBody to JSON.
func MapNetworkPoliciesUpdateBodyToJSON(v *NetworkPoliciesUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
