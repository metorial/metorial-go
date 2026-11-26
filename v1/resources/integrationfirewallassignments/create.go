package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsCreateOutputFirewall - Firewall applied to the target.
type IntegrationFirewallAssignmentsCreateOutputFirewall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsCreateOutputTarget - Target the firewall applies to. `null` when the target no longer exists.
type IntegrationFirewallAssignmentsCreateOutputTarget struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Kind of target the firewall applies to.
	Type string `json:"type"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsCreateOutput represents the integration firewall assignments create output type.
type IntegrationFirewallAssignmentsCreateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// Firewall - Firewall applied to the target.
	Firewall IntegrationFirewallAssignmentsCreateOutputFirewall `json:"firewall"`
	// Target - Target the firewall applies to. `null` when the target no longer exists.
	Target *IntegrationFirewallAssignmentsCreateOutputTarget `json:"target,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
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
	// FirewallId - Firewall to apply to the target.
	FirewallId string `json:"firewall_id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// EnclaveId - Enclave to apply the firewall to. Required when `target_type` is `enclave`.
	EnclaveId *string `json:"enclave_id,omitempty"`
	// ProviderId - Provider to apply the firewall to. Required when `target_type` is `provider`.
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
