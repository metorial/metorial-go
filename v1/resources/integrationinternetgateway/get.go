package integrationinternetgateway

import (
	"encoding/json"
	"time"
)

// IntegrationInternetGatewayGetOutputPublicIps represents the integration internet gateway get output public ips type.
type IntegrationInternetGatewayGetOutputPublicIps struct {
	Object    string    `json:"object"`
	Id        string    `json:"id"`
	Ip        string    `json:"ip"`
	Region    string    `json:"region"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationInternetGatewayGetOutput represents the integration internet gateway get output type.
type IntegrationInternetGatewayGetOutput struct {
	Object      string                                         `json:"object"`
	Id          string                                         `json:"id"`
	Name        string                                         `json:"name"`
	Description *string                                        `json:"description,omitempty"`
	CreatedAt   time.Time                                      `json:"created_at"`
	UpdatedAt   time.Time                                      `json:"updated_at"`
	PublicIps   []IntegrationInternetGatewayGetOutputPublicIps `json:"public_ips"`
}

// MapIntegrationInternetGatewayGetOutputFromJSON deserializes JSON data into a IntegrationInternetGatewayGetOutput.
func MapIntegrationInternetGatewayGetOutputFromJSON(data []byte) (*IntegrationInternetGatewayGetOutput, error) {
	var v IntegrationInternetGatewayGetOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationInternetGatewayGetOutputToJSON serializes a IntegrationInternetGatewayGetOutput to JSON.
func MapIntegrationInternetGatewayGetOutputToJSON(v *IntegrationInternetGatewayGetOutput) ([]byte, error) {
	return json.Marshal(v)
}
