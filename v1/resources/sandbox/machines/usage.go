package machines

import (
	"encoding/json"
	"time"
)

// SandboxMachinesUsageOutputItems represents the sandbox machines usage output items type.
type SandboxMachinesUsageOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - The sample's unique identifier
	Id string `json:"id"`
	// MachineInstanceId - The instance the sample was taken from
	MachineInstanceId string `json:"machine_instance_id"`
	// Timestamp - When the sample was taken
	Timestamp time.Time `json:"timestamp"`
	// CpuPercent - CPU utilization across all cores, in percent
	CpuPercent float64 `json:"cpu_percent"`
	// CoresUsed - The number of cores in use
	CoresUsed float64 `json:"cores_used"`
	// CoresAvailable - The number of cores available
	CoresAvailable float64 `json:"cores_available"`
	// MemoryUsedBytes - Memory in use, in bytes
	MemoryUsedBytes float64 `json:"memory_used_bytes"`
	// MemoryTotalBytes - Total memory, in bytes
	MemoryTotalBytes float64 `json:"memory_total_bytes"`
	// NetworkRxBytes - Total bytes received over the network
	NetworkRxBytes float64 `json:"network_rx_bytes"`
	// NetworkTxBytes - Total bytes sent over the network
	NetworkTxBytes float64 `json:"network_tx_bytes"`
	// DiskReadBytes - Total bytes read from disk
	DiskReadBytes float64 `json:"disk_read_bytes"`
	// DiskWriteBytes - Total bytes written to disk
	DiskWriteBytes float64 `json:"disk_write_bytes"`
}

// SandboxMachinesUsageOutputPagination represents the sandbox machines usage output pagination type.
type SandboxMachinesUsageOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// SandboxMachinesUsageOutput represents the sandbox machines usage output type.
type SandboxMachinesUsageOutput struct {
	Items      []SandboxMachinesUsageOutputItems    `json:"items"`
	Pagination SandboxMachinesUsageOutputPagination `json:"pagination"`
}

// MapSandboxMachinesUsageOutputFromJSON deserializes JSON data into a SandboxMachinesUsageOutput.
func MapSandboxMachinesUsageOutputFromJSON(data []byte) (*SandboxMachinesUsageOutput, error) {
	var v SandboxMachinesUsageOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUsageOutputToJSON serializes a SandboxMachinesUsageOutput to JSON.
func MapSandboxMachinesUsageOutputToJSON(v *SandboxMachinesUsageOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesUsageQuery represents the sandbox machines usage query type.
type SandboxMachinesUsageQuery struct {
	Limit             *float64 `json:"limit,omitempty"`
	After             *string  `json:"after,omitempty"`
	Before            *string  `json:"before,omitempty"`
	Cursor            *string  `json:"cursor,omitempty"`
	Order             *string  `json:"order,omitempty"`
	MachineInstanceId *any     `json:"machine_instance_id,omitempty"`
}

// MapSandboxMachinesUsageQueryFromJSON deserializes JSON data into a SandboxMachinesUsageQuery.
func MapSandboxMachinesUsageQueryFromJSON(data []byte) (*SandboxMachinesUsageQuery, error) {
	var v SandboxMachinesUsageQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUsageQueryToJSON serializes a SandboxMachinesUsageQuery to JSON.
func MapSandboxMachinesUsageQueryToJSON(v *SandboxMachinesUsageQuery) ([]byte, error) {
	return json.Marshal(v)
}
