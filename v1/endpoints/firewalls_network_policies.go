package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/firewalls/networkpolicies"
)

// FirewallsNetworkPoliciesEndpoint provides access to manage firewalls and their attached network policies.
type FirewallsNetworkPoliciesEndpoint struct {
	client *endpoint.Client
}

// NewFirewallsNetworkPoliciesEndpoint creates a new FirewallsNetworkPoliciesEndpoint.
func NewFirewallsNetworkPoliciesEndpoint(client *endpoint.Client) *FirewallsNetworkPoliciesEndpoint {
	return &FirewallsNetworkPoliciesEndpoint{client: client}
}

// FirewallsNetworkPoliciesEndpointAttachBody contains the request body for Attach.
type FirewallsNetworkPoliciesEndpointAttachBody struct {
	NetworkPolicyId string   `json:"network_policy_id"`
	Position        *float64 `json:"position,omitempty"`
}

// Attach attaches a network policy to a firewall.
func (e *FirewallsNetworkPoliciesEndpoint) Attach(firewallId string, body *FirewallsNetworkPoliciesEndpointAttachBody) (*networkpolicies.FirewallsNetworkPoliciesAttachOutput, error) {
	req := &endpoint.Request{
		Path: []string{"firewalls", firewallId, "network-policies"},
		Body: body,
	}
	var result networkpolicies.FirewallsNetworkPoliciesAttachOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Detach detaches a network policy from a firewall.
func (e *FirewallsNetworkPoliciesEndpoint) Detach(firewallId string, networkPolicyId string) (*networkpolicies.FirewallsNetworkPoliciesDetachOutput, error) {
	req := &endpoint.Request{
		Path: []string{"firewalls", firewallId, "network-policies", networkPolicyId},
	}
	var result networkpolicies.FirewallsNetworkPoliciesDetachOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
