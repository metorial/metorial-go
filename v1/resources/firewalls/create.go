package firewalls

import (
	"encoding/json"
	"time"
)

// FirewallsCreateOutputNetworkPoliciesRulesPorts represents the firewalls create output network policies rules ports type.
type FirewallsCreateOutputNetworkPoliciesRulesPorts struct {
	Object string  `json:"object"`
	From   float64 `json:"from"`
	To     float64 `json:"to"`
}

// FirewallsCreateOutputNetworkPoliciesRules represents the firewalls create output network policies rules type.
type FirewallsCreateOutputNetworkPoliciesRules struct {
	Object      string                                            `json:"object"`
	Id          string                                            `json:"id"`
	Effect      string                                            `json:"effect"`
	Direction   string                                            `json:"direction"`
	Cidrs       []string                                          `json:"cidrs"`
	Description *string                                           `json:"description,omitempty"`
	Enabled     bool                                              `json:"enabled"`
	Priority    float64                                           `json:"priority"`
	Ports       *[]FirewallsCreateOutputNetworkPoliciesRulesPorts `json:"ports,omitempty"`
}

// FirewallsCreateOutputNetworkPolicies represents the firewalls create output network policies type.
type FirewallsCreateOutputNetworkPolicies struct {
	Object  string                                      `json:"object"`
	Id      string                                      `json:"id"`
	Name    string                                      `json:"name"`
	Version float64                                     `json:"version"`
	Rules   []FirewallsCreateOutputNetworkPoliciesRules `json:"rules"`
}

// FirewallsCreateOutput represents the firewalls create output type.
type FirewallsCreateOutput struct {
	Object          string                                 `json:"object"`
	Id              string                                 `json:"id"`
	Slug            string                                 `json:"slug"`
	Name            string                                 `json:"name"`
	Description     *string                                `json:"description,omitempty"`
	Status          string                                 `json:"status"`
	NetworkId       string                                 `json:"network_id"`
	NetworkPolicies []FirewallsCreateOutputNetworkPolicies `json:"network_policies"`
	CreatedAt       time.Time                              `json:"created_at"`
	UpdatedAt       time.Time                              `json:"updated_at"`
	ArchivedAt      *time.Time                             `json:"archived_at,omitempty"`
}

// MapFirewallsCreateOutputFromJSON deserializes JSON data into a FirewallsCreateOutput.
func MapFirewallsCreateOutputFromJSON(data []byte) (*FirewallsCreateOutput, error) {
	var v FirewallsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsCreateOutputToJSON serializes a FirewallsCreateOutput to JSON.
func MapFirewallsCreateOutputToJSON(v *FirewallsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallsCreateBodyBindings represents the firewalls create body bindings type.
type FirewallsCreateBodyBindings struct {
	TargetType string  `json:"target_type"`
	EnclaveId  *string `json:"enclave_id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
	NetworkId  *string `json:"network_id,omitempty"`
}

// FirewallsCreateBody represents the firewalls create body type.
type FirewallsCreateBody struct {
	Name             string                         `json:"name"`
	Description      *string                        `json:"description,omitempty"`
	Slug             *string                        `json:"slug,omitempty"`
	NetworkId        string                         `json:"network_id"`
	Bindings         *[]FirewallsCreateBodyBindings `json:"bindings,omitempty"`
	NetworkPolicyIds *[]string                      `json:"network_policy_ids,omitempty"`
}

// MapFirewallsCreateBodyFromJSON deserializes JSON data into a FirewallsCreateBody.
func MapFirewallsCreateBodyFromJSON(data []byte) (*FirewallsCreateBody, error) {
	var v FirewallsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallsCreateBodyToJSON serializes a FirewallsCreateBody to JSON.
func MapFirewallsCreateBodyToJSON(v *FirewallsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
