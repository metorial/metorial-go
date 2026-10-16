package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsGetOutputFirewall represents the integration firewall assignments get output firewall type.
type IntegrationFirewallAssignmentsGetOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsGetOutputTarget represents the integration firewall assignments get output target type.
type IntegrationFirewallAssignmentsGetOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsGetOutput represents the integration firewall assignments get output type.
type IntegrationFirewallAssignmentsGetOutput struct {
	Object     string                                          `json:"object"`
	Id         string                                          `json:"id"`
	TargetType string                                          `json:"target_type"`
	Firewall   IntegrationFirewallAssignmentsGetOutputFirewall `json:"firewall"`
	Target     *IntegrationFirewallAssignmentsGetOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                                       `json:"created_at"`
}

// MapIntegrationFirewallAssignmentsGetOutputFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsGetOutput.
func MapIntegrationFirewallAssignmentsGetOutputFromJSON(data []byte) (*IntegrationFirewallAssignmentsGetOutput, error) {
	var v IntegrationFirewallAssignmentsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsGetOutputToJSON serializes a IntegrationFirewallAssignmentsGetOutput to JSON.
func MapIntegrationFirewallAssignmentsGetOutputToJSON(v *IntegrationFirewallAssignmentsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
