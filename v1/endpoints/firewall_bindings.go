package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/firewallbindings"
)

// FirewallBindingsEndpoint provides access to manage bindings that apply firewalls to enclaves, providers, or networks.
type FirewallBindingsEndpoint struct {
	client *endpoint.Client
}

// NewFirewallBindingsEndpoint creates a new FirewallBindingsEndpoint.
func NewFirewallBindingsEndpoint(client *endpoint.Client) *FirewallBindingsEndpoint {
	return &FirewallBindingsEndpoint{client: client}
}

// FirewallBindingsEndpointListParams contains optional query parameters for List.
type FirewallBindingsEndpointListParams struct {
	Limit      *float64 `json:"limit,omitempty"`
	After      *string  `json:"after,omitempty"`
	Before     *string  `json:"before,omitempty"`
	Cursor     *string  `json:"cursor,omitempty"`
	Order      *string  `json:"order,omitempty"`
	Id         *any     `json:"id,omitempty"`
	FirewallId *any     `json:"firewall_id,omitempty"`
	EnclaveId  *any     `json:"enclave_id,omitempty"`
	ProviderId *any     `json:"provider_id,omitempty"`
	NetworkId  *any     `json:"network_id,omitempty"`
	TargetType *any     `json:"target_type,omitempty"`
	// CreatedAt - Filter firewall binding creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
}

// FirewallBindingsEndpointCreateBody contains the request body for Create.
type FirewallBindingsEndpointCreateBody struct {
	FirewallId string  `json:"firewall_id"`
	TargetType string  `json:"target_type"`
	EnclaveId  *string `json:"enclave_id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
	NetworkId  *string `json:"network_id,omitempty"`
}

// List returns a paginated list of firewall bindings.
func (e *FirewallBindingsEndpoint) List(params *FirewallBindingsEndpointListParams) (*firewallbindings.FirewallBindingsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"firewall-bindings"},
		Query: query,
	}
	var result firewallbindings.FirewallBindingsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific firewall binding by ID.
func (e *FirewallBindingsEndpoint) Get(firewallBindingId string) (*firewallbindings.FirewallBindingsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"firewall-bindings", firewallBindingId},
	}
	var result firewallbindings.FirewallBindingsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create creates a binding that applies a firewall to a target.
func (e *FirewallBindingsEndpoint) Create(body *FirewallBindingsEndpointCreateBody) (*firewallbindings.FirewallBindingsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"firewall-bindings"},
		Body: body,
	}
	var result firewallbindings.FirewallBindingsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete deletes a firewall binding.
func (e *FirewallBindingsEndpoint) Delete(firewallBindingId string) (*firewallbindings.FirewallBindingsDeleteOutput, error) {
	req := &endpoint.Request{
		Path: []string{"firewall-bindings", firewallBindingId},
	}
	var result firewallbindings.FirewallBindingsDeleteOutput
	if err := e.client.Delete(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
