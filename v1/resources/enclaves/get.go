package enclaves

import (
	"encoding/json"
	"time"
)

// EnclavesGetOutputEnclaveEnvironment represents the enclaves get output enclave environment type.
type EnclavesGetOutputEnclaveEnvironment struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// EnclavesGetOutput represents the enclaves get output type.
type EnclavesGetOutput struct {
	Object               string                              `json:"object"`
	Id                   string                              `json:"id"`
	Slug                 string                              `json:"slug"`
	Name                 string                              `json:"name"`
	Description          *string                             `json:"description,omitempty"`
	NetworkId            string                              `json:"network_id"`
	ProviderDeploymentId string                              `json:"provider_deployment_id"`
	EnclaveEnvironment   EnclavesGetOutputEnclaveEnvironment `json:"enclave_environment"`
	CreatedAt            time.Time                           `json:"created_at"`
	LastUsedAt           *time.Time                          `json:"last_used_at,omitempty"`
}

// MapEnclavesGetOutputFromJSON deserializes JSON data into a EnclavesGetOutput.
func MapEnclavesGetOutputFromJSON(data []byte) (*EnclavesGetOutput, error) {
	var v EnclavesGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapEnclavesGetOutputToJSON serializes a EnclavesGetOutput to JSON.
func MapEnclavesGetOutputToJSON(v *EnclavesGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
