package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/sandbox/machines/usage"
)

// SandboxMachinesUsageEndpoint provides access to aggregated measured CPU, memory, network, disk I/O, and observed runtime. History is flushed approximately every minute and is not billing consumption.
type SandboxMachinesUsageEndpoint struct {
	client *endpoint.Client
}

// NewSandboxMachinesUsageEndpoint creates a new SandboxMachinesUsageEndpoint.
func NewSandboxMachinesUsageEndpoint(client *endpoint.Client) *SandboxMachinesUsageEndpoint {
	return &SandboxMachinesUsageEndpoint{client: client}
}

// SandboxMachinesUsageEndpointTimelineParams contains optional query parameters for Timeline.
type SandboxMachinesUsageEndpointTimelineParams struct {
	// From - Inclusive start of the requested range. History is retained for 30 days; older starts are clamped. Maximum requested span is 31 days.
	From string `json:"from"`
	// To - Exclusive end of the requested range. Must follow from and cannot be in the future. Bounds are rounded outward to complete UTC buckets.
	To string `json:"to"`
	// MachineInstanceId - Restrict usage to one or more executions of this machine.
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	// GroupBy - Group by machine execution. Omit to aggregate the whole machine.
	GroupBy *string `json:"group_by,omitempty"`
	// MaxSeries - Maximum individually reported groups, default 50. Remaining executions form an additional other group. Applies when grouping.
	MaxSeries *float64 `json:"max_series,omitempty"`
	// Interval - Bucket interval. Omit for automatic selection targeting about 500 points. At most 1500 points per series; intervals below one hour require a range within the last 7 days.
	Interval *string `json:"interval,omitempty"`
}

// SandboxMachinesUsageEndpointSummaryParams contains optional query parameters for Summary.
type SandboxMachinesUsageEndpointSummaryParams struct {
	// From - Inclusive start of the requested range. History is retained for 30 days; older starts are clamped. Maximum requested span is 31 days.
	From string `json:"from"`
	// To - Exclusive end of the requested range. Must follow from and cannot be in the future. Bounds are rounded outward to complete UTC buckets.
	To string `json:"to"`
	// MachineInstanceId - Restrict usage to one or more executions of this machine.
	MachineInstanceId *any `json:"machine_instance_id,omitempty"`
	// GroupBy - Group by machine execution. Omit to aggregate the whole machine.
	GroupBy *string `json:"group_by,omitempty"`
	// MaxSeries - Maximum individually reported groups, default 50. Remaining executions form an additional other group. Applies when grouping.
	MaxSeries *float64 `json:"max_series,omitempty"`
}

// Timeline returns chronological usage buckets. Minute data is retained for 7 days and hourly data for 30 days. Zero samples means no observations, not necessarily a stopped machine.
func (e *SandboxMachinesUsageEndpoint) Timeline(machineId string, params *SandboxMachinesUsageEndpointTimelineParams) (*usage.SandboxMachinesUsageTimelineOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines", machineId, "usage", "timeline"},
		Query: query,
	}
	var result usage.SandboxMachinesUsageTimelineOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Summary returns totals and averages for a range. Summary bounds align independently to minute or hour source buckets. Use a timeline’s effective bounds to compare its totals with a summary.
func (e *SandboxMachinesUsageEndpoint) Summary(machineId string, params *SandboxMachinesUsageEndpointSummaryParams) (*usage.SandboxMachinesUsageSummaryOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"sandbox", "machines", machineId, "usage", "summary"},
		Query: query,
	}
	var result usage.SandboxMachinesUsageSummaryOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
