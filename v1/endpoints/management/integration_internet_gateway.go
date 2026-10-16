package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationinternetgateway"
)

// IntegrationInternetGatewayEndpoint provides access to read the internet gateway that integration enclaves use to reach and be reached from the internet.
type IntegrationInternetGatewayEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationInternetGatewayEndpoint creates a new IntegrationInternetGatewayEndpoint.
func NewIntegrationInternetGatewayEndpoint(client *endpoint.Client) *IntegrationInternetGatewayEndpoint {
	return &IntegrationInternetGatewayEndpoint{client: client}
}

// IntegrationInternetGatewayEndpointListNetworkLogsParams contains optional query parameters for ListNetworkLogs.
type IntegrationInternetGatewayEndpointListNetworkLogsParams struct {
	Direction       string   `json:"direction"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	Hostname        *any     `json:"hostname,omitempty"`
	Ip              *any     `json:"ip,omitempty"`
	From            *string  `json:"from,omitempty"`
	To              *string  `json:"to,omitempty"`
	IntervalMinutes *float64 `json:"interval_minutes,omitempty"`
}

// Get retrieves the integration internet gateway of the instance, including its public IPs.
func (e *IntegrationInternetGatewayEndpoint) Get(instanceId string) (*integrationinternetgateway.IntegrationInternetGatewayGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "integration-internet-gateway"},
	}
	var result integrationinternetgateway.IntegrationInternetGatewayGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListNetworkLogs returns ingress or egress network logs for integration enclaves in the instance environment.
func (e *IntegrationInternetGatewayEndpoint) ListNetworkLogs(instanceId string, params *IntegrationInternetGatewayEndpointListNetworkLogsParams) (*integrationinternetgateway.IntegrationInternetGatewayListNetworkLogsOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "integration-internet-gateway", "network-logs"},
		Query: query,
	}
	var result integrationinternetgateway.IntegrationInternetGatewayListNetworkLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
