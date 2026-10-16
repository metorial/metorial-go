package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsDeleteOutputFirewall represents the integration firewall assignments delete output firewall type.
type IntegrationFirewallAssignmentsDeleteOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsDeleteOutputTarget represents the integration firewall assignments delete output target type.
type IntegrationFirewallAssignmentsDeleteOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsDeleteOutput represents the integration firewall assignments delete output type.
type IntegrationFirewallAssignmentsDeleteOutput struct {
	Object     string                                             `json:"object"`
	Id         string                                             `json:"id"`
	TargetType string                                             `json:"target_type"`
	Firewall   IntegrationFirewallAssignmentsDeleteOutputFirewall `json:"firewall"`
	Target     *IntegrationFirewallAssignmentsDeleteOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                                          `json:"created_at"`
}

// MapIntegrationFirewallAssignmentsDeleteOutputFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsDeleteOutput.
func MapIntegrationFirewallAssignmentsDeleteOutputFromJSON(data []byte) (*IntegrationFirewallAssignmentsDeleteOutput, error) {
	var v IntegrationFirewallAssignmentsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsDeleteOutputToJSON serializes a IntegrationFirewallAssignmentsDeleteOutput to JSON.
func MapIntegrationFirewallAssignmentsDeleteOutputToJSON(v *IntegrationFirewallAssignmentsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
