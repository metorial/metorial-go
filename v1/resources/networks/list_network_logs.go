package networks

import (
	"encoding/json"
)

// NetworksListNetworkLogsOutputRecords represents the networks list network logs output records type.
type NetworksListNetworkLogsOutputRecords struct {
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

// NetworksListNetworkLogsOutput represents the networks list network logs output type.
type NetworksListNetworkLogsOutput struct {
	Object     string                                 `json:"object"`
	Direction  string                                 `json:"direction"`
	EnclaveIds []string                               `json:"enclave_ids"`
	Records    []NetworksListNetworkLogsOutputRecords `json:"records"`
}

// MapNetworksListNetworkLogsOutputFromJSON deserializes JSON data into a NetworksListNetworkLogsOutput.
func MapNetworksListNetworkLogsOutputFromJSON(data []byte) (*NetworksListNetworkLogsOutput, error) {
	var v NetworksListNetworkLogsOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworksListNetworkLogsOutputToJSON serializes a NetworksListNetworkLogsOutput to JSON.
func MapNetworksListNetworkLogsOutputToJSON(v *NetworksListNetworkLogsOutput) ([]byte, error) {
	return json.Marshal(v)
}

// NetworksListNetworkLogsQuery represents the networks list network logs query type.
type NetworksListNetworkLogsQuery struct {
	Direction       string   `json:"direction"`
	EnclaveId       *any     `json:"enclave_id,omitempty"`
	Hostname        *any     `json:"hostname,omitempty"`
	Ip              *any     `json:"ip,omitempty"`
	From            *string  `json:"from,omitempty"`
	To              *string  `json:"to,omitempty"`
	IntervalMinutes *float64 `json:"interval_minutes,omitempty"`
}

// MapNetworksListNetworkLogsQueryFromJSON deserializes JSON data into a NetworksListNetworkLogsQuery.
func MapNetworksListNetworkLogsQueryFromJSON(data []byte) (*NetworksListNetworkLogsQuery, error) {
	var v NetworksListNetworkLogsQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapNetworksListNetworkLogsQueryToJSON serializes a NetworksListNetworkLogsQuery to JSON.
func MapNetworksListNetworkLogsQueryToJSON(v *NetworksListNetworkLogsQuery) ([]byte, error) {
	return json.Marshal(v)
}
