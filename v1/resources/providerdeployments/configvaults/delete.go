package configvaults

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigVaultsDeleteOutputDeployment represents the provider deployments config vaults delete output deployment type.
type ProviderDeploymentsConfigVaultsDeleteOutputDeployment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// IsDefault - Used when a request doesn't name a deployment.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this deployment belongs to.
	ProviderId string `json:"provider_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// ProviderDeploymentsConfigVaultsDeleteOutput represents the provider deployments config vaults delete output type.
type ProviderDeploymentsConfigVaultsDeleteOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Archived once the vault is archived. Archived vaults stay listable.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ProviderId - Provider this vault belongs to.
	ProviderId string                                                 `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigVaultsDeleteOutputDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigVaultsDeleteOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsDeleteOutput.
func MapProviderDeploymentsConfigVaultsDeleteOutputFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsDeleteOutput, error) {
	var v ProviderDeploymentsConfigVaultsDeleteOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsDeleteOutputToJSON serializes a ProviderDeploymentsConfigVaultsDeleteOutput to JSON.
func MapProviderDeploymentsConfigVaultsDeleteOutputToJSON(v *ProviderDeploymentsConfigVaultsDeleteOutput) ([]byte, error) {
	return json.Marshal(v)
}
