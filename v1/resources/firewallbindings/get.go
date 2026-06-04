package firewallbindings

import (
	"encoding/json"
	"time"
)

// FirewallBindingsGetOutputFirewall represents the firewall bindings get output firewall type.
type FirewallBindingsGetOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// FirewallBindingsGetOutputTarget represents the firewall bindings get output target type.
type FirewallBindingsGetOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// FirewallBindingsGetOutput represents the firewall bindings get output type.
type FirewallBindingsGetOutput struct {
	Object     string                            `json:"object"`
	Id         string                            `json:"id"`
	TargetType string                            `json:"target_type"`
	Firewall   FirewallBindingsGetOutputFirewall `json:"firewall"`
	Target     *FirewallBindingsGetOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                         `json:"created_at"`
}

// MapFirewallBindingsGetOutputFromJSON deserializes JSON data into a FirewallBindingsGetOutput.
func MapFirewallBindingsGetOutputFromJSON(data []byte) (*FirewallBindingsGetOutput, error) {
	var v FirewallBindingsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsGetOutputToJSON serializes a FirewallBindingsGetOutput to JSON.
func MapFirewallBindingsGetOutputToJSON(v *FirewallBindingsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
