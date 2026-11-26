package integrationfirewallassignments

import (
	"encoding/json"
	"time"
)

// IntegrationFirewallAssignmentsDeleteOutputFirewall - Firewall applied to the target.
type IntegrationFirewallAssignmentsDeleteOutputFirewall struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsDeleteOutputTarget - Target the firewall applies to. `null` when the target no longer exists.
type IntegrationFirewallAssignmentsDeleteOutputTarget struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Type - Kind of target the firewall applies to.
	Type string `json:"type"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
}

// IntegrationFirewallAssignmentsDeleteOutput represents the integration firewall assignments delete output type.
type IntegrationFirewallAssignmentsDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// TargetType - What the firewall applies to.
	TargetType string `json:"target_type"`
	// Firewall - Firewall applied to the target.
	Firewall IntegrationFirewallAssignmentsDeleteOutputFirewall `json:"firewall"`
	// Target - Target the firewall applies to. `null` when the target no longer exists.
	Target *IntegrationFirewallAssignmentsDeleteOutputTarget `json:"target,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
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
