package firewallbindings

import (
	"encoding/json"
	"time"
)

// FirewallBindingsDeleteOutputFirewall represents the firewall bindings delete output firewall type.
type FirewallBindingsDeleteOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// FirewallBindingsDeleteOutputTarget represents the firewall bindings delete output target type.
type FirewallBindingsDeleteOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// FirewallBindingsDeleteOutput represents the firewall bindings delete output type.
type FirewallBindingsDeleteOutput struct {
	Object     string                               `json:"object"`
	Id         string                               `json:"id"`
	TargetType string                               `json:"target_type"`
	Firewall   FirewallBindingsDeleteOutputFirewall `json:"firewall"`
	Target     *FirewallBindingsDeleteOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                            `json:"created_at"`
}

// MapFirewallBindingsDeleteOutputFromJSON deserializes JSON data into a FirewallBindingsDeleteOutput.
func MapFirewallBindingsDeleteOutputFromJSON(data []byte) (*FirewallBindingsDeleteOutput, error) {
	var v FirewallBindingsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapFirewallBindingsDeleteOutputToJSON serializes a FirewallBindingsDeleteOutput to JSON.
func MapFirewallBindingsDeleteOutputToJSON(v *FirewallBindingsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
