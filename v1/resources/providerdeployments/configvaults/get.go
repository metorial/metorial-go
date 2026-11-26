package configvaults

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigVaultsGetOutputDeployment represents the provider deployments config vaults get output deployment type.
type ProviderDeploymentsConfigVaultsGetOutputDeployment struct {
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

// ProviderDeploymentsConfigVaultsGetOutput represents the provider deployments config vaults get output type.
type ProviderDeploymentsConfigVaultsGetOutput struct {
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
	ProviderId string                                              `json:"provider_id"`
	Deployment *ProviderDeploymentsConfigVaultsGetOutputDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigVaultsGetOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsGetOutput.
func MapProviderDeploymentsConfigVaultsGetOutputFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsGetOutput, error) {
	var v ProviderDeploymentsConfigVaultsGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsGetOutputToJSON serializes a ProviderDeploymentsConfigVaultsGetOutput to JSON.
func MapProviderDeploymentsConfigVaultsGetOutputToJSON(v *ProviderDeploymentsConfigVaultsGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
