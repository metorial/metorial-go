package configvaults

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsConfigVaultsUpdateOutputDeployment represents the provider deployments config vaults update output deployment type.
type ProviderDeploymentsConfigVaultsUpdateOutputDeployment struct {
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

// ProviderDeploymentsConfigVaultsUpdateOutput represents the provider deployments config vaults update output type.
type ProviderDeploymentsConfigVaultsUpdateOutput struct {
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
	Deployment *ProviderDeploymentsConfigVaultsUpdateOutputDeployment `json:"deployment,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// MapProviderDeploymentsConfigVaultsUpdateOutputFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsUpdateOutput.
func MapProviderDeploymentsConfigVaultsUpdateOutputFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsUpdateOutput, error) {
	var v ProviderDeploymentsConfigVaultsUpdateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsUpdateOutputToJSON serializes a ProviderDeploymentsConfigVaultsUpdateOutput to JSON.
func MapProviderDeploymentsConfigVaultsUpdateOutputToJSON(v *ProviderDeploymentsConfigVaultsUpdateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsConfigVaultsUpdateBody represents the provider deployments config vaults update body type.
type ProviderDeploymentsConfigVaultsUpdateBody struct {
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data to identify or organize the provider config vault. Replaces existing metadata.
	Metadata *map[string]any `json:"metadata,omitempty"`
}

// MapProviderDeploymentsConfigVaultsUpdateBodyFromJSON deserializes JSON data into a ProviderDeploymentsConfigVaultsUpdateBody.
func MapProviderDeploymentsConfigVaultsUpdateBodyFromJSON(data []byte) (*ProviderDeploymentsConfigVaultsUpdateBody, error) {
	var v ProviderDeploymentsConfigVaultsUpdateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsConfigVaultsUpdateBodyToJSON serializes a ProviderDeploymentsConfigVaultsUpdateBody to JSON.
func MapProviderDeploymentsConfigVaultsUpdateBodyToJSON(v *ProviderDeploymentsConfigVaultsUpdateBody) ([]byte, error) {
	return json.Marshal(v)
}
