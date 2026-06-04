package networks

import (
	"encoding/json"
	"time"
)

// NetworksGetOutputPublicIps represents the networks get output public ips type.
type NetworksGetOutputPublicIps struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Ip        string    `json:"ip"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NetworksGetOutput represents the networks get output type.
type NetworksGetOutput struct {
	Object      string                       `json:"object"`
	Id          string                       `json:"id"`
	Name        string                       `json:"name"`
	Description *string                      `json:"description,omitempty"`
	CreatedAt   time.Time                    `json:"created_at"`
	UpdatedAt   time.Time                    `json:"updated_at"`
	PublicIps   []NetworksGetOutputPublicIps `json:"public_ips"`
}

// MapNetworksGetOutputFromJSON deserializes JSON data into a NetworksGetOutput.
func MapNetworksGetOutputFromJSON(data []byte) (*NetworksGetOutput, error) {
	var v NetworksGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworksGetOutputToJSON serializes a NetworksGetOutput to JSON.
func MapNetworksGetOutputToJSON(v *NetworksGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
