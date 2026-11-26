package delegationconfigs

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationConfigsUpdateOutput represents the identities delegation configs update output type.
type IdentitiesDelegationConfigsUpdateOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the delegation config. `archived` once you archive the config.
	Status string `json:"status"`
	// IsDefault - Delegations use this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SubDelegationBehavior - Whether delegates can delegate further. `require_consent` asks the owner for approval first.
	SubDelegationBehavior string `json:"sub_delegation_behavior"`
	// SubDelegationDepth - Maximum sub-delegation depth. `0` blocks sub-delegation.
	SubDelegationDepth float64 `json:"sub_delegation_depth"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapIdentitiesDelegationConfigsUpdateOutputFromJSON deserializes JSON data into a IdentitiesDelegationConfigsUpdateOutput.
func MapIdentitiesDelegationConfigsUpdateOutputFromJSON(data []byte) (*IdentitiesDelegationConfigsUpdateOutput, error) {
	var v IdentitiesDelegationConfigsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationConfigsUpdateOutputToJSON serializes a IdentitiesDelegationConfigsUpdateOutput to JSON.
func MapIdentitiesDelegationConfigsUpdateOutputToJSON(v *IdentitiesDelegationConfigsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationConfigsUpdateBody represents the identities delegation configs update body type.
type IdentitiesDelegationConfigsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the identity delegation config. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SubDelegationBehavior - Whether delegates can delegate further. `require_consent` asks the owner for approval first.
	SubDelegationBehavior *string `json:"sub_delegation_behavior,omitempty"`
	// SubDelegationDepth - Maximum sub-delegation depth. `0` blocks sub-delegation.
	SubDelegationDepth *float64 `json:"sub_delegation_depth,omitempty"`
}

// MapIdentitiesDelegationConfigsUpdateBodyFromJSON deserializes JSON data into a IdentitiesDelegationConfigsUpdateBody.
func MapIdentitiesDelegationConfigsUpdateBodyFromJSON(data []byte) (*IdentitiesDelegationConfigsUpdateBody, error) {
	var v IdentitiesDelegationConfigsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationConfigsUpdateBodyToJSON serializes a IdentitiesDelegationConfigsUpdateBody to JSON.
func MapIdentitiesDelegationConfigsUpdateBodyToJSON(v *IdentitiesDelegationConfigsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
