package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/networkpolicies/rules"
)

// NetworkPoliciesRulesEndpoint provides access to manage reusable network policy definitions and their rules.
type NetworkPoliciesRulesEndpoint struct {
	client *endpoint.Client
}

// NewNetworkPoliciesRulesEndpoint creates a new NetworkPoliciesRulesEndpoint.
func NewNetworkPoliciesRulesEndpoint(client *endpoint.Client) *NetworkPoliciesRulesEndpoint {
	return &NetworkPoliciesRulesEndpoint{client: client}
}

// NetworkPoliciesRulesEndpointCreateBody contains the request body for Create.
type NetworkPoliciesRulesEndpointCreateBody struct {
	Effect      string            `json:"effect"`
	Direction   string            `json:"direction"`
	Cidrs       []string          `json:"cidrs"`
	Description *string           `json:"description,omitempty"`
	Enabled     bool              `json:"enabled"`
	Priority    float64           `json:"priority"`
	Ports       *[]map[string]any `json:"ports,omitempty"`
}

// NetworkPoliciesRulesEndpointUpdateBody contains the request body for Update.
type NetworkPoliciesRulesEndpointUpdateBody struct {
	Effect      string            `json:"effect"`
	Direction   string            `json:"direction"`
	Cidrs       []string          `json:"cidrs"`
	Description *string           `json:"description,omitempty"`
	Enabled     bool              `json:"enabled"`
	Priority    float64           `json:"priority"`
	Ports       *[]map[string]any `json:"ports,omitempty"`
}

// Create adds a rule to a network policy.
func (e *NetworkPoliciesRulesEndpoint) Create(instanceId string, networkPolicyId string, body *NetworkPoliciesRulesEndpointCreateBody) (*rules.NetworkPoliciesRulesCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId, "rules"},
		Body: body,
	}
	var result rules.NetworkPoliciesRulesCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update updates a rule on a network policy.
func (e *NetworkPoliciesRulesEndpoint) Update(instanceId string, networkPolicyId string, ruleId string, body *NetworkPoliciesRulesEndpointUpdateBody) (*rules.NetworkPoliciesRulesUpdateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId, "rules", ruleId},
		Body: body,
	}
	var result rules.NetworkPoliciesRulesUpdateOutput
	if err := e.client.Patch(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete removes a rule from a network policy.
func (e *NetworkPoliciesRulesEndpoint) Delete(instanceId string, networkPolicyId string, ruleId string) (*rules.NetworkPoliciesRulesDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "network-policies", networkPolicyId, "rules", ruleId},
	}
	var result rules.NetworkPoliciesRulesDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
