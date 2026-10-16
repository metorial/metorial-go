package integrationenclaves

import (
	"encoding/json"
	"time"
)

// IntegrationEnclavesGetOutputEnclaveEnvironment represents the integration enclaves get output enclave environment type.
type IntegrationEnclavesGetOutputEnclaveEnvironment struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// IntegrationEnclavesGetOutput represents the integration enclaves get output type.
type IntegrationEnclavesGetOutput struct {
	Object               string                                         `json:"object"`
	Id                   string                                         `json:"id"`
	Slug                 string                                         `json:"slug"`
	Name                 string                                         `json:"name"`
	Description          *string                                        `json:"description,omitempty"`
	InternetGatewayId    string                                         `json:"internet_gateway_id"`
	ProviderDeploymentId string                                         `json:"provider_deployment_id"`
	EnclaveEnvironment   IntegrationEnclavesGetOutputEnclaveEnvironment `json:"enclave_environment"`
	CreatedAt            time.Time                                      `json:"created_at"`
	LastUsedAt           *time.Time                                     `json:"last_used_at,omitempty"`
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
