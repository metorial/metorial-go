package firewallbindings

import (
	"encoding/json"
	"time"
)

// FirewallBindingsCreateOutputFirewall represents the firewall bindings create output firewall type.
type FirewallBindingsCreateOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// FirewallBindingsCreateOutputTarget represents the firewall bindings create output target type.
type FirewallBindingsCreateOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// FirewallBindingsCreateOutput represents the firewall bindings create output type.
type FirewallBindingsCreateOutput struct {
	Object     string                               `json:"object"`
	Id         string                               `json:"id"`
	TargetType string                               `json:"target_type"`
	Firewall   FirewallBindingsCreateOutputFirewall `json:"firewall"`
	Target     *FirewallBindingsCreateOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                            `json:"created_at"`
}

// MapFirewallBindingsCreateOutputFromJSON deserializes JSON data into a FirewallBindingsCreateOutput.
func MapFirewallBindingsCreateOutputFromJSON(data []byte) (*FirewallBindingsCreateOutput, error) {
	var v FirewallBindingsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsCreateOutputToJSON serializes a FirewallBindingsCreateOutput to JSON.
func MapFirewallBindingsCreateOutputToJSON(v *FirewallBindingsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// FirewallBindingsCreateBody represents the firewall bindings create body type.
type FirewallBindingsCreateBody struct {
	FirewallId string  `json:"firewall_id"`
	TargetType string  `json:"target_type"`
	EnclaveId  *string `json:"enclave_id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
	NetworkId  *string `json:"network_id,omitempty"`
}

// MapFirewallBindingsCreateBodyFromJSON deserializes JSON data into a FirewallBindingsCreateBody.
func MapFirewallBindingsCreateBodyFromJSON(data []byte) (*FirewallBindingsCreateBody, error) {
	var v FirewallBindingsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsCreateBodyToJSON serializes a FirewallBindingsCreateBody to JSON.
func MapFirewallBindingsCreateBodyToJSON(v *FirewallBindingsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
