package delegationconfigs

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationConfigsDeleteOutput represents the identities delegation configs delete output type.
type IdentitiesDelegationConfigsDeleteOutput struct {
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

// MapIdentitiesDelegationConfigsDeleteOutputFromJSON deserializes JSON data into a IdentitiesDelegationConfigsDeleteOutput.
func MapIdentitiesDelegationConfigsDeleteOutputFromJSON(data []byte) (*IdentitiesDelegationConfigsDeleteOutput, error) {
	var v IdentitiesDelegationConfigsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationConfigsDeleteOutputToJSON serializes a IdentitiesDelegationConfigsDeleteOutput to JSON.
func MapIdentitiesDelegationConfigsDeleteOutputToJSON(v *IdentitiesDelegationConfigsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
