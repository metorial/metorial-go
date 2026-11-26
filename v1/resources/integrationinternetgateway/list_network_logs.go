package integrationinternetgateway

import (
	"encoding/json"
)

// IntegrationInternetGatewayListNetworkLogsOutputRecords - One bucket of network logs.
type IntegrationInternetGatewayListNetworkLogsOutputRecords struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Direction - Whether the record is an ingress or egress log.
	Direction string `json:"direction"`
	// EnclaveId - Enclave the record belongs to.
	EnclaveId string `json:"enclave_id"`
	// BucketStart - Start of the time bucket, as an ISO timestamp.
	BucketStart string `json:"bucket_start"`
	// Hostname - Hostname observed in the record.
	Hostname string `json:"hostname"`
	// Ip - IP address observed in the record.
	Ip string `json:"ip"`
	// Port - Port observed in the record.
	Port float64 `json:"port"`
	// Count - Requests folded into this record.
	Count float64 `json:"count"`
	// Result - Whether the traffic was allowed or denied. Only set for ingress logs.
	Result *string `json:"result,omitempty"`
	// FirstSeenAt - When matching traffic was first seen, as an ISO timestamp.
	FirstSeenAt string `json:"first_seen_at"`
	// LastSeenAt - When matching traffic was last seen, as an ISO timestamp.
	LastSeenAt string `json:"last_seen_at"`
}

// IntegrationInternetGatewayListNetworkLogsOutput represents the integration internet gateway list network logs output type.
type IntegrationInternetGatewayListNetworkLogsOutput struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Direction - Whether these are ingress or egress logs.
	Direction string `json:"direction"`
	// EnclaveIds - Enclaves covered by these logs.
	EnclaveIds []string `json:"enclave_ids"`
	// Records - Log records in the response.
	Records []IntegrationInternetGatewayListNetworkLogsOutputRecords `json:"records"`
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
