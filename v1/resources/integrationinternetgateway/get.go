package integrationinternetgateway

import (
	"encoding/json"
	"time"
)

// IntegrationInternetGatewayGetOutputPublicIps represents the integration internet gateway get output public ips type.
type IntegrationInternetGatewayGetOutputPublicIps struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Ip - Public IP address. Masked as `x.x.x.x` unless your plan includes IP access.
	Ip string `json:"ip"`
	// Region - Region the IP belongs to.
	Region string `json:"region"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Creation time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IntegrationInternetGatewayGetOutput represents the integration internet gateway get output type.
type IntegrationInternetGatewayGetOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// PublicIps - Public IPs of the gateway. Masked as `x.x.x.x` unless your plan includes IP access.
	PublicIps []IntegrationInternetGatewayGetOutputPublicIps `json:"public_ips"`
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
