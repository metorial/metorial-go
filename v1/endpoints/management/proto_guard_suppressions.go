package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/protoguardsuppressions"
)

// ProtoGuardSuppressionsEndpoint provides access to a ProtoGuard suppression is a rule that hides known-benign findings. Matching findings no longer raise alerts.
type ProtoGuardSuppressionsEndpoint struct {
	client *endpoint.Client
}

// NewProtoGuardSuppressionsEndpoint creates a new ProtoGuardSuppressionsEndpoint.
func NewProtoGuardSuppressionsEndpoint(client *endpoint.Client) *ProtoGuardSuppressionsEndpoint {
	return &ProtoGuardSuppressionsEndpoint{client: client}
}

// ProtoGuardSuppressionsEndpointListParams contains optional query parameters for List.
type ProtoGuardSuppressionsEndpointListParams struct {
	// Limit - Page size, from 1 to 100. Defaults to 20.
	Limit *float64 `json:"limit,omitempty"`
	// After - Return the page after this point. Pass `pagination.after` if the previous response has it, otherwise the `id` of its last item.
	After *string `json:"after,omitempty"`
	// Before - Return the page before this point. Pass `pagination.before` if the previous response has it, otherwise the `id` of its first item.
	Before *string `json:"before,omitempty"`
	// Cursor - A `cur_` cursor from `pagination.after` or `pagination.before`. Alternative to `after` and `before`.
	Cursor *string `json:"cursor,omitempty"`
	// Order - Sort direction.
	Order *string `json:"order,omitempty"`
	// Status - Only return suppressions with these statuses. Defaults to active.
	Status *any `json:"status,omitempty"`
	// FilterId - Only return suppressions of these filters, by id or key.
	FilterId *any `json:"filter_id,omitempty"`
	// ProviderId - Only return suppressions scoped to these providers.
	ProviderId *any `json:"provider_id,omitempty"`
}

// ProtoGuardSuppressionsEndpointCreateBody contains the request body for Create.
type ProtoGuardSuppressionsEndpointCreateBody struct {
	// FilterId - The filter to suppress, by id or key.
	FilterId string `json:"filter_id"`
	// ProviderId - Only suppress findings from this provider.
	ProviderId *string `json:"provider_id,omitempty"`
	// ToolKey - Only suppress findings from this tool.
	ToolKey *string `json:"tool_key,omitempty"`
	// ContentHash - Only suppress findings with this content hash, as reported on a ProtoGuard alert finding.
	ContentHash *string `json:"content_hash,omitempty"`
	// Reason - Why the findings are considered benign.
	Reason *string `json:"reason,omitempty"`
	// ExpiresAt - When the suppression stops applying. Omit to never expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
}

// List returns a paginated list of ProtoGuard suppressions in the instance.
func (e *ProtoGuardSuppressionsEndpoint) List(instanceId string, params *ProtoGuardSuppressionsEndpointListParams) (*protoguardsuppressions.ProtoGuardSuppressionsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "protoguard-suppressions"},
		Query: query,
	}
	var result protoguardsuppressions.ProtoGuardSuppressionsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a ProtoGuard suppression.
func (e *ProtoGuardSuppressionsEndpoint) Get(instanceId string, suppressionId string) (*protoguardsuppressions.ProtoGuardSuppressionsGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-suppressions", suppressionId},
	}
	var result protoguardsuppressions.ProtoGuardSuppressionsGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create suppresses future findings of a filter, optionally narrowed to a provider, tool, or specific matched content.
func (e *ProtoGuardSuppressionsEndpoint) Create(instanceId string, body *ProtoGuardSuppressionsEndpointCreateBody) (*protoguardsuppressions.ProtoGuardSuppressionsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-suppressions"},
		Body: body,
	}
	var result protoguardsuppressions.ProtoGuardSuppressionsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Archive archives a ProtoGuard suppression so matching findings alert again.
func (e *ProtoGuardSuppressionsEndpoint) Archive(instanceId string, suppressionId string) (*protoguardsuppressions.ProtoGuardSuppressionsArchiveOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "protoguard-suppressions", suppressionId, "archive"},
	}
	var result protoguardsuppressions.ProtoGuardSuppressionsArchiveOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
