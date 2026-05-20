package ssotenants

import (
	"encoding/json"
)

// PortalsAuthSsoTenantsSetupOutput represents the portals auth sso tenants setup output type.
type PortalsAuthSsoTenantsSetupOutput struct {
	Object string `json:"object"`
	Url    string `json:"url"`
}

// MapPortalsAuthSsoTenantsSetupOutputFromJSON deserializes JSON data into a PortalsAuthSsoTenantsSetupOutput.
func MapPortalsAuthSsoTenantsSetupOutputFromJSON(data []byte) (*PortalsAuthSsoTenantsSetupOutput, error) {
	var v PortalsAuthSsoTenantsSetupOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsSetupOutputToJSON serializes a PortalsAuthSsoTenantsSetupOutput to JSON.
func MapPortalsAuthSsoTenantsSetupOutputToJSON(v *PortalsAuthSsoTenantsSetupOutput) ([]byte, error) {
	return json.Marshal(v)
}
