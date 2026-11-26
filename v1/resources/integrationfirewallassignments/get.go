package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsGetOutputFirewall - Firewall applied to the target.
type IntegrationFirewallAssignmentsGetOutputFirewall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsGetOutputTarget - Target the firewall applies to. `null` when the target no longer exists.
type IntegrationFirewallAssignmentsGetOutputTarget struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Kind of target the firewall applies to.
	Type string `json:"type"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsGetOutput represents the integration firewall assignments get output type.
type IntegrationFirewallAssignmentsGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// Firewall - Firewall applied to the target.
	Firewall IntegrationFirewallAssignmentsGetOutputFirewall `json:"firewall"`
	// Target - Target the firewall applies to. `null` when the target no longer exists.
	Target *IntegrationFirewallAssignmentsGetOutputTarget `json:"target,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
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
