package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerissues"
)

// ProviderIssuesEndpoint provides access to a provider issue is a group of repeated provider errors, authentication failures, or ProtoGuard warnings sharing one underlying problem.
type ProviderIssuesEndpoint struct {
	client *endpoint.Client
}

// NewProviderIssuesEndpoint creates a new ProviderIssuesEndpoint.
func NewProviderIssuesEndpoint(client *endpoint.Client) *ProviderIssuesEndpoint {
	return &ProviderIssuesEndpoint{client: client}
}

// ProviderIssuesEndpointListParams contains optional query parameters for List.
type ProviderIssuesEndpointListParams struct {
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
	// Type - Filter by issue type(s)
	Type *any `json:"type,omitempty"`
	// Severity - Filter by severity
	Severity *any `json:"severity,omitempty"`
	// Id - Filter by provider issue ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// LastSeenAt - Only return results whose provider issue last occurrence time is in this range. Set `gt`, `lt`, or both.
	LastSeenAt *map[string]any `json:"last_seen_at,omitempty"`
	// CreatedAt - Only return results whose provider issue creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// Search - Search by issue ID, code, subtype or message, or by the name of the related provider
	Search *string `json:"search,omitempty"`
}

// List returns a paginated list of provider issues. Only issue types the caller has read access to are included.
func (e *ProviderIssuesEndpoint) List(instanceId string, params *ProviderIssuesEndpointListParams) (*providerissues.ProviderIssuesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-issues"},
		Query: query,
	}
	var result providerissues.ProviderIssuesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider issue.
func (e *ProviderIssuesEndpoint) Get(instanceId string, providerIssueId string) (*providerissues.ProviderIssuesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-issues", providerIssueId},
	}
	var result providerissues.ProviderIssuesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
