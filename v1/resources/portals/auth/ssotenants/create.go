package ssotenants

import (
	"encoding/json"
	"time"
)

// PortalsAuthSsoTenantsCreateOutputCounts represents the portals auth sso tenants create output counts type.
type PortalsAuthSsoTenantsCreateOutputCounts struct {
	Connections float64 `json:"connections"`
}

// PortalsAuthSsoTenantsCreateOutput represents the portals auth sso tenants create output type.
type PortalsAuthSsoTenantsCreateOutput struct {
	Object    string                                  `json:"object"`
	Id        string                                  `json:"id"`
	Name      string                                  `json:"name"`
	Status    string                                  `json:"status"`
	ClientId  string                                  `json:"client_id"`
	Counts    PortalsAuthSsoTenantsCreateOutputCounts `json:"counts"`
	CreatedAt time.Time                               `json:"created_at"`
	UpdatedAt time.Time                               `json:"updated_at"`
}

// MapPortalsAuthSsoTenantsCreateOutputFromJSON deserializes JSON data into a PortalsAuthSsoTenantsCreateOutput.
func MapPortalsAuthSsoTenantsCreateOutputFromJSON(data []byte) (*PortalsAuthSsoTenantsCreateOutput, error) {
	var v PortalsAuthSsoTenantsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsCreateOutputToJSON serializes a PortalsAuthSsoTenantsCreateOutput to JSON.
func MapPortalsAuthSsoTenantsCreateOutputToJSON(v *PortalsAuthSsoTenantsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// PortalsAuthSsoTenantsCreateBody represents the portals auth sso tenants create body type.
type PortalsAuthSsoTenantsCreateBody struct {
	Name string `json:"name"`
}

// MapPortalsAuthSsoTenantsCreateBodyFromJSON deserializes JSON data into a PortalsAuthSsoTenantsCreateBody.
func MapPortalsAuthSsoTenantsCreateBodyFromJSON(data []byte) (*PortalsAuthSsoTenantsCreateBody, error) {
	var v PortalsAuthSsoTenantsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapPortalsAuthSsoTenantsCreateBodyToJSON serializes a PortalsAuthSsoTenantsCreateBody to JSON.
func MapPortalsAuthSsoTenantsCreateBodyToJSON(v *PortalsAuthSsoTenantsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
