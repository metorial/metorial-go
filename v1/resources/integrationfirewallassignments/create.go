package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsCreateOutputFirewall represents the integration firewall assignments create output firewall type.
type IntegrationFirewallAssignmentsCreateOutputFirewall struct {
	Object string `json:"object"`
	Id     string `json:"id"`
	Slug   string `json:"slug"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsCreateOutputTarget represents the integration firewall assignments create output target type.
type IntegrationFirewallAssignmentsCreateOutputTarget struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

// IntegrationFirewallAssignmentsCreateOutput represents the integration firewall assignments create output type.
type IntegrationFirewallAssignmentsCreateOutput struct {
	Object     string                                             `json:"object"`
	Id         string                                             `json:"id"`
	TargetType string                                             `json:"target_type"`
	Firewall   IntegrationFirewallAssignmentsCreateOutputFirewall `json:"firewall"`
	Target     *IntegrationFirewallAssignmentsCreateOutputTarget  `json:"target,omitempty"`
	CreatedAt  time.Time                                          `json:"created_at"`
}

// MapIntegrationFirewallAssignmentsCreateOutputFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsCreateOutput.
func MapIntegrationFirewallAssignmentsCreateOutputFromJSON(data []byte) (*IntegrationFirewallAssignmentsCreateOutput, error) {
	var v IntegrationFirewallAssignmentsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsCreateOutputToJSON serializes a IntegrationFirewallAssignmentsCreateOutput to JSON.
func MapIntegrationFirewallAssignmentsCreateOutputToJSON(v *IntegrationFirewallAssignmentsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationFirewallAssignmentsCreateBody represents the integration firewall assignments create body type.
type IntegrationFirewallAssignmentsCreateBody struct {
	FirewallId string  `json:"firewall_id"`
	TargetType string  `json:"target_type"`
	EnclaveId  *string `json:"enclave_id,omitempty"`
	ProviderId *string `json:"provider_id,omitempty"`
}

// MapIntegrationFirewallAssignmentsCreateBodyFromJSON deserializes JSON data into a IntegrationFirewallAssignmentsCreateBody.
func MapIntegrationFirewallAssignmentsCreateBodyFromJSON(data []byte) (*IntegrationFirewallAssignmentsCreateBody, error) {
	var v IntegrationFirewallAssignmentsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationFirewallAssignmentsCreateBodyToJSON serializes a IntegrationFirewallAssignmentsCreateBody to JSON.
func MapIntegrationFirewallAssignmentsCreateBodyToJSON(v *IntegrationFirewallAssignmentsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
