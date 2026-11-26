package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/networkpolicies/rules"
)

// NetworkPoliciesRulesEndpoint provides access to a network policy is a reusable set of allow and deny rules. Firewalls attach policies to enforce them.
type NetworkPoliciesRulesEndpoint struct {
	client *endpoint.Client
}

// NewNetworkPoliciesRulesEndpoint creates a new NetworkPoliciesRulesEndpoint.
func NewNetworkPoliciesRulesEndpoint(client *endpoint.Client) *NetworkPoliciesRulesEndpoint {
	return &NetworkPoliciesRulesEndpoint{client: client}
}

// NetworkPoliciesRulesEndpointCreateBody contains the request body for Create.
type NetworkPoliciesRulesEndpointCreateBody struct {
	// Effect - `allow` permits matching traffic. `deny` blocks it.
	Effect string `json:"effect"`
	// Direction - `ingress` matches incoming traffic. `egress` matches outgoing traffic.
	Direction string `json:"direction"`
	// Cidrs - CIDR ranges the rule matches.
	Cidrs []string `json:"cidrs"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Enabled - Whether the rule is enforced. Disabled rules are kept but ignored.
	Enabled bool `json:"enabled"`
	// Priority - Priority of the rule. Higher values are evaluated first.
	Priority float64 `json:"priority"`
	// Ports - Port ranges the rule matches. Omitted matches all ports.
	Ports *[]map[string]any `json:"ports,omitempty"`
}

// NetworkPoliciesRulesEndpointUpdateBody contains the request body for Update.
type NetworkPoliciesRulesEndpointUpdateBody struct {
	// Effect - `allow` permits matching traffic. `deny` blocks it.
	Effect string `json:"effect"`
	// Direction - `ingress` matches incoming traffic. `egress` matches outgoing traffic.
	Direction string `json:"direction"`
	// Cidrs - CIDR ranges the rule matches.
	Cidrs []string `json:"cidrs"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Enabled - Whether the rule is enforced. Disabled rules are kept but ignored.
	Enabled bool `json:"enabled"`
	// Priority - Priority of the rule. Higher values are evaluated first.
	Priority float64 `json:"priority"`
	// Ports - Port ranges the rule matches. Omitted matches all ports.
	Ports *[]map[string]any `json:"ports,omitempty"`
}

// Create adds a rule to a network policy.
func (e *NetworkPoliciesRulesEndpoint) Create(networkPolicyId string, body *NetworkPoliciesRulesEndpointCreateBody) (*rules.NetworkPoliciesRulesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId, "rules"},
		Body: body,
	}
	var result rules.NetworkPoliciesRulesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a rule on a network policy.
func (e *NetworkPoliciesRulesEndpoint) Update(networkPolicyId string, ruleId string, body *NetworkPoliciesRulesEndpointUpdateBody) (*rules.NetworkPoliciesRulesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId, "rules", ruleId},
		Body: body,
	}
	var result rules.NetworkPoliciesRulesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a rule from a network policy.
func (e *NetworkPoliciesRulesEndpoint) Delete(networkPolicyId string, ruleId string) (*rules.NetworkPoliciesRulesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"network-policies", networkPolicyId, "rules", ruleId},
	}
	var result rules.NetworkPoliciesRulesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
