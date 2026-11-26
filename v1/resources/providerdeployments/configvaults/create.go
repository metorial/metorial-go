package configvaults

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigVaultsCreateOutputDeployment represents the provider deployments config vaults create output deployment type.
type ProviderDeploymentsConfigVaultsCreateOutputDeployment struct {
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

// ProviderDeploymentsConfigVaultsCreateOutput represents the provider deployments config vaults create output type.
type ProviderDeploymentsConfigVaultsCreateOutput struct {
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
	Deployment *ProviderDeploymentsConfigVaultsCreateOutputDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigVaultsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsCreateOutput.
func MapProviderDeploymentsConfigVaultsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsCreateOutput, error) {
	var v ProviderDeploymentsConfigVaultsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsCreateOutputToJSON serializes a ProviderDeploymentsConfigVaultsCreateOutput to JSON.
func MapProviderDeploymentsConfigVaultsCreateOutputToJSON(v *ProviderDeploymentsConfigVaultsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigVaultsCreateBody represents the provider deployments config vaults create body type.
type ProviderDeploymentsConfigVaultsCreateBody struct {
	// ProviderId - Provider to create the vault for.
	ProviderId string `json:"provider_id"`
	// ProviderDeploymentId - Deployment the vault is locked to.
	ProviderDeploymentId *string `json:"provider_deployment_id,omitempty"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config vault. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Value - Secure configuration values to store in the vault
	Value map[string]any `json:"value"`
}

// MapProviderDeploymentsConfigVaultsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsCreateBody.
func MapProviderDeploymentsConfigVaultsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsCreateBody, error) {
	var v ProviderDeploymentsConfigVaultsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsCreateBodyToJSON serializes a ProviderDeploymentsConfigVaultsCreateBody to JSON.
func MapProviderDeploymentsConfigVaultsCreateBodyToJSON(v *ProviderDeploymentsConfigVaultsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
