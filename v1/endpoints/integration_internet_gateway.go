package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/integrationinternetgateway"
)

// IntegrationInternetGatewayEndpoint provides access to an integration internet gateway is the shared network exit for integration enclaves in an instance. Enclaves use it to reach the internet and be reached from it.
type IntegrationInternetGatewayEndpoint struct {
	client *endpoint.Client
}

// NewIntegrationInternetGatewayEndpoint creates a new IntegrationInternetGatewayEndpoint.
func NewIntegrationInternetGatewayEndpoint(client *endpoint.Client) *IntegrationInternetGatewayEndpoint {
	return &IntegrationInternetGatewayEndpoint{client: client}
}

// IntegrationInternetGatewayEndpointListNetworkLogsParams contains optional query parameters for ListNetworkLogs.
type IntegrationInternetGatewayEndpointListNetworkLogsParams struct {
	// Direction - Whether to return ingress or egress logs.
	Direction string `json:"direction"`
	// EnclaveId - Only return logs for one of these enclaves.
	EnclaveId *any `json:"enclave_id,omitempty"`
	// Hostname - Only return logs for one of these hostnames.
	Hostname *any `json:"hostname,omitempty"`
	// Ip - Only return logs for one of these IPs.
	Ip *any `json:"ip,omitempty"`
	// From - Only return logs at or after this time, as an ISO timestamp.
	From *string `json:"from,omitempty"`
	// To - Only return logs at or before this time, as an ISO timestamp.
	To *string `json:"to,omitempty"`
	// IntervalMinutes - Bucket size in minutes.
	IntervalMinutes *float64 `json:"interval_minutes,omitempty"`
}

// Get retrieves the instance's integration internet gateway. Public IPs are masked as `x.x.x.x` unless your plan includes IP access.
func (e *IntegrationInternetGatewayEndpoint) Get() (*integrationinternetgateway.IntegrationInternetGatewayGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"integration-internet-gateway"},
	}
	var result integrationinternetgateway.IntegrationInternetGatewayGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListNetworkLogs returns network logs for integration enclaves in the instance. Set `direction` to choose ingress or egress logs.
func (e *IntegrationInternetGatewayEndpoint) ListNetworkLogs(params *IntegrationInternetGatewayEndpointListNetworkLogsParams) (*integrationinternetgateway.IntegrationInternetGatewayListNetworkLogsOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"integration-internet-gateway", "network-logs"},
		Query: query,
	}
	var result integrationinternetgateway.IntegrationInternetGatewayListNetworkLogsOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
