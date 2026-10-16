package integrationinternetgateway

import (
	"encoding/json"
)

// IntegrationInternetGatewayListNetworkLogsOutputRecords represents the integration internet gateway list network logs output records type.
type IntegrationInternetGatewayListNetworkLogsOutputRecords struct {
	Object      string  `json:"object"`
	Direction   string  `json:"direction"`
	EnclaveId   string  `json:"enclave_id"`
	BucketStart string  `json:"bucket_start"`
	Hostname    string  `json:"hostname"`
	Ip          string  `json:"ip"`
	Port        float64 `json:"port"`
	Count       float64 `json:"count"`
	Result      *string `json:"result,omitempty"`
	FirstSeenAt string  `json:"first_seen_at"`
	LastSeenAt  string  `json:"last_seen_at"`
}

// IntegrationInternetGatewayListNetworkLogsOutput represents the integration internet gateway list network logs output type.
type IntegrationInternetGatewayListNetworkLogsOutput struct {
	Object     string                                                   `json:"object"`
	Direction  string                                                   `json:"direction"`
	EnclaveIds []string                                                 `json:"enclave_ids"`
	Records    []IntegrationInternetGatewayListNetworkLogsOutputRecords `json:"records"`
}

// MapIntegrationInternetGatewayListNetworkLogsOutputFromJSON deserializes JSON data into a IntegrationInternetGatewayListNetworkLogsOutput.
func MapIntegrationInternetGatewayListNetworkLogsOutputFromJSON(data []byte) (*IntegrationInternetGatewayListNetworkLogsOutput, error) {
	var v IntegrationInternetGatewayListNetworkLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationInternetGatewayListNetworkLogsOutputToJSON serializes a IntegrationInternetGatewayListNetworkLogsOutput to JSON.
func MapIntegrationInternetGatewayListNetworkLogsOutputToJSON(v *IntegrationInternetGatewayListNetworkLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IntegrationInternetGatewayListNetworkLogsQuery represents the integration internet gateway list network logs query type.
type IntegrationInternetGatewayListNetworkLogsQuery struct {
	Direction       string   `json:"direction"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	Hostname        *any     `json:"hostname,omitempty"`
	Ip              *any     `json:"ip,omitempty"`
	From            *string  `json:"from,omitempty"`
	To              *string  `json:"to,omitempty"`
	IntervalMinutes *float64 `json:"interval_minutes,omitempty"`
}

// MapIntegrationInternetGatewayListNetworkLogsQueryFromJSON deserializes JSON data into a IntegrationInternetGatewayListNetworkLogsQuery.
func MapIntegrationInternetGatewayListNetworkLogsQueryFromJSON(data []byte) (*IntegrationInternetGatewayListNetworkLogsQuery, error) {
	var v IntegrationInternetGatewayListNetworkLogsQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIntegrationInternetGatewayListNetworkLogsQueryToJSON serializes a IntegrationInternetGatewayListNetworkLogsQuery to JSON.
func MapIntegrationInternetGatewayListNetworkLogsQueryToJSON(v *IntegrationInternetGatewayListNetworkLogsQuery) ([]byte, error) {
	return json.Marshal(v)
}
