package endpoints

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerissues"
)

// ProviderIssuesEndpoint provides access to provider issues aggregate repeated provider errors, authentication failures and ProtoGuard warnings into one record per distinct problem.
type ProviderIssuesEndpoint struct {
	client *endpoint.Client
}

// NewProviderIssuesEndpoint creates a new ProviderIssuesEndpoint.
func NewProviderIssuesEndpoint(client *endpoint.Client) *ProviderIssuesEndpoint {
	return &ProviderIssuesEndpoint{client: client}
}

// ProviderIssuesEndpointListParams contains optional query parameters for List.
type ProviderIssuesEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	// Type - Filter by issue type(s)
	Type *any `json:"type,omitempty"`
	// Severity - Filter by severity
	Severity *any `json:"severity,omitempty"`
	// Id - Filter by provider issue ID(s)
	Id *any `json:"id,omitempty"`
	// ProviderId - Filter by provider ID(s)
	ProviderId *any `json:"provider_id,omitempty"`
	// LastSeenAt - Filter provider issue last occurrence time by date range
	LastSeenAt *map[string]any `json:"last_seen_at,omitempty"`
	// CreatedAt - Filter provider issue creation time by date range
	CreatedAt *map[string]any `json:"created_at,omitempty"`
	// Search - Search by issue ID, code, subtype or message, or by the name of the related provider
	Search *string `json:"search,omitempty"`
}

// List returns a paginated list of provider issues. Only issue types the caller has read access to are included.
func (e *ProviderIssuesEndpoint) List(params *ProviderIssuesEndpointListParams) (*providerissues.ProviderIssuesListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"provider-issues"},
		Query: query,
	}
	var result providerissues.ProviderIssuesListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get retrieves a specific provider issue.
func (e *ProviderIssuesEndpoint) Get(providerIssueId string) (*providerissues.ProviderIssuesGetOutput, error) {
	req := &endpoint.Request{
		Path: []string{"provider-issues", providerIssueId},
	}
	var result providerissues.ProviderIssuesGetOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
