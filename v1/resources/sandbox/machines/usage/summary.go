package usage

import (
	"encoding/json"
	"time"
)

// SandboxMachinesUsageSummaryOutputGroupsGroup - Dimensions identifying this group.
type SandboxMachinesUsageSummaryOutputGroupsGroup struct {
	// MachineInstanceId - Execution identifier, or null for an ungrouped aggregate or the overflow group.
	MachineInstanceId *string `json:"machine_instance_id,omitempty"`
}

// SandboxMachinesUsageSummaryOutputGroupsTotals - Usage totals and averages for the effective range. Averages use the full range duration.
type SandboxMachinesUsageSummaryOutputGroupsTotals struct {
	// Samples - Number of observations. Zero means no telemetry was observed, not necessarily that the machine was stopped.
	Samples float64 `json:"samples"`
	// RuntimeSeconds - Observed runtime in seconds, not billable runtime.
	RuntimeSeconds float64 `json:"runtime_seconds"`
	// CpuCoreSeconds - CPU consumed in core seconds during the bucket or range.
	CpuCoreSeconds float64 `json:"cpu_core_seconds"`
	// AverageCoresUsed - CPU core seconds divided by the full bucket or range duration in seconds.
	AverageCoresUsed float64 `json:"average_cores_used"`
	// CpuPercent - CPU consumed as a percentage of observed allocated CPU capacity.
	CpuPercent float64 `json:"cpu_percent"`
	// PeakCoresUsed - Highest sampled number of cores in use.
	PeakCoresUsed float64 `json:"peak_cores_used"`
	// MemoryByteSeconds - Time-integrated used memory in byte seconds.
	MemoryByteSeconds float64 `json:"memory_byte_seconds"`
	// AverageMemoryUsedBytes - Used memory in bytes averaged over the full bucket or range duration.
	AverageMemoryUsedBytes float64 `json:"average_memory_used_bytes"`
	// PeakMemoryUsedBytes - Highest sampled memory usage in bytes.
	PeakMemoryUsedBytes float64 `json:"peak_memory_used_bytes"`
	// AllocatedMemoryByteSeconds - Time-integrated allocated memory in byte seconds.
	AllocatedMemoryByteSeconds float64 `json:"allocated_memory_byte_seconds"`
	// AverageMemoryAllocatedBytes - Allocated memory in bytes averaged over the full bucket or range duration.
	AverageMemoryAllocatedBytes float64 `json:"average_memory_allocated_bytes"`
	// NetworkRxBytes - Bytes received during the bucket or range; not a cumulative counter.
	NetworkRxBytes float64 `json:"network_rx_bytes"`
	// NetworkTxBytes - Bytes sent during the bucket or range; not a cumulative counter.
	NetworkTxBytes float64 `json:"network_tx_bytes"`
	// DiskReadBytes - Bytes read from disk during the bucket or range; not a cumulative counter.
	DiskReadBytes float64 `json:"disk_read_bytes"`
	// DiskWriteBytes - Bytes written to disk during the bucket or range; not a cumulative counter.
	DiskWriteBytes float64 `json:"disk_write_bytes"`
}

// SandboxMachinesUsageSummaryOutputGroups represents the sandbox machines usage summary output groups type.
type SandboxMachinesUsageSummaryOutputGroups struct {
	// Group - Dimensions identifying this group.
	Group SandboxMachinesUsageSummaryOutputGroupsGroup `json:"group"`
	// Other - Whether this group combines executions outside the top max_series groups, ranked by CPU consumption.
	Other bool `json:"other"`
	// Totals - Usage totals and averages for the effective range. Averages use the full range duration.
	Totals SandboxMachinesUsageSummaryOutputGroupsTotals `json:"totals"`
}

// SandboxMachinesUsageSummaryOutput represents the sandbox machines usage summary output type.
type SandboxMachinesUsageSummaryOutput struct {
	// Object - String representing the object's type.
	Object string `json:"object"`
	// MachineId - The machine whose resource usage is reported.
	MachineId string `json:"machine_id"`
	// From - Effective inclusive range start, clamped to retention and rounded down to a UTC bucket boundary.
	From time.Time `json:"from"`
	// To - Effective exclusive range end, rounded up to a UTC bucket boundary.
	To time.Time `json:"to"`
	// Resolution - Resolution of the retained source data, independently of the timeline interval.
	Resolution string `json:"resolution"`
	// GroupBy - Grouping dimensions. Empty means usage is aggregated across the machine.
	GroupBy []string `json:"group_by"`
	// Groups - Usage groups ordered by CPU consumption, with the overflow group last. Empty when no observations exist.
	Groups []SandboxMachinesUsageSummaryOutputGroups `json:"groups"`
}

// MapSandboxMachinesUsageSummaryOutputFromJSON deserializes JSON data into a SandboxMachinesUsageSummaryOutput.
func MapSandboxMachinesUsageSummaryOutputFromJSON(data []byte) (*SandboxMachinesUsageSummaryOutput, error) {
	var v SandboxMachinesUsageSummaryOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUsageSummaryOutputToJSON serializes a SandboxMachinesUsageSummaryOutput to JSON.
func MapSandboxMachinesUsageSummaryOutputToJSON(v *SandboxMachinesUsageSummaryOutput) ([]byte, error) {
	return json.Marshal(v)
}

// SandboxMachinesUsageSummaryQuery represents the sandbox machines usage summary query type.
type SandboxMachinesUsageSummaryQuery struct {
	// From - Inclusive start of the requested range. History is retained for 30 days; older starts are clamped. Maximum requested span is 31 days.
	From time.Time `json:"from"`
	// To - Exclusive end of the requested range. Must follow from and cannot be in the future. Bounds are rounded outward to complete UTC buckets.
	To time.Time `json:"to"`
	// MachineInstanceId - Restrict usage to one or more executions of this machine.
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	// GroupBy - Group by machine execution. Omit to aggregate the whole machine.
	GroupBy *string `json:"group_by,omitempty"`
	// MaxSeries - Maximum individually reported groups, default 50. Remaining executions form an additional other group. Applies when grouping.
	MaxSeries *float64 `json:"max_series,omitempty"`
}

// MapSandboxMachinesUsageSummaryQueryFromJSON deserializes JSON data into a SandboxMachinesUsageSummaryQuery.
func MapSandboxMachinesUsageSummaryQueryFromJSON(data []byte) (*SandboxMachinesUsageSummaryQuery, error) {
	var v SandboxMachinesUsageSummaryQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapSandboxMachinesUsageSummaryQueryToJSON serializes a SandboxMachinesUsageSummaryQuery to JSON.
func MapSandboxMachinesUsageSummaryQueryToJSON(v *SandboxMachinesUsageSummaryQuery) ([]byte, error) {
	return json.Marshal(v)
}
