package integrationenclaves

import (
	"encoding/json"
	"time"
)

// IntegrationEnclavesGetOutputEnclaveEnvironment - Environment the enclave runs in.
type IntegrationEnclavesGetOutputEnclaveEnvironment struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Type - `metorial` for Metorial-hosted environments. `outpost` for environments in your own infrastructure.
	Type string `json:"type"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
}

// IntegrationEnclavesGetOutput represents the integration enclaves get output type.
type IntegrationEnclavesGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// InternetGatewayId - Internet gateway the enclave uses.
	InternetGatewayId string `json:"internet_gateway_id"`
	// ProviderDeploymentId - Provider deployment running in the enclave.
	ProviderDeploymentId string `json:"provider_deployment_id"`
	// EnclaveEnvironment - Environment the enclave runs in.
	EnclaveEnvironment IntegrationEnclavesGetOutputEnclaveEnvironment `json:"enclave_environment"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// LastUsedAt - When the enclave was last used. `null` when never used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
}

// MapIntegrationEnclavesGetOutputFromJSON deserializes JSON data into a IntegrationEnclavesGetOutput.
func MapIntegrationEnclavesGetOutputFromJSON(data []byte) (*IntegrationEnclavesGetOutput, error) {
	var v IntegrationEnclavesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationEnclavesGetOutputToJSON serializes a IntegrationEnclavesGetOutput to JSON.
func MapIntegrationEnclavesGetOutputToJSON(v *IntegrationEnclavesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
