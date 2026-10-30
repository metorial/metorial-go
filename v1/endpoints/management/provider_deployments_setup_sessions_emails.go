package management

import (
	"github.com/metorial/metorial-go/v1/internal/endpoint"
	"github.com/metorial/metorial-go/v1/resources/providerdeployments/setupsessions/emails"
)

// ProviderDeploymentsSetupSessionsEmailsEndpoint provides access to a setup session tracks an in-progress OAuth flow, storing state during the redirect. On success, it creates an auth config with the user's access token.
type ProviderDeploymentsSetupSessionsEmailsEndpoint struct {
	client *endpoint.Client
}

// NewProviderDeploymentsSetupSessionsEmailsEndpoint creates a new ProviderDeploymentsSetupSessionsEmailsEndpoint.
func NewProviderDeploymentsSetupSessionsEmailsEndpoint(client *endpoint.Client) *ProviderDeploymentsSetupSessionsEmailsEndpoint {
	return &ProviderDeploymentsSetupSessionsEmailsEndpoint{client: client}
}

// ProviderDeploymentsSetupSessionsEmailsEndpointCreateBody contains the request body for Create.
type ProviderDeploymentsSetupSessionsEmailsEndpointCreateBody struct {
	// Emails - Email addresses to send the setup link to, at most 10
	Emails *[]string `json:"emails,omitempty"`
	// OrganizationMemberIds - Organization members to send the setup link to, at most 10. Each member is emailed at their organization profile email address, falling back to their account email when the profile has none.
	OrganizationMemberIds *[]string `json:"organization_member_ids,omitempty"`
}

// ProviderDeploymentsSetupSessionsEmailsEndpointListParams contains optional query parameters for List.
type ProviderDeploymentsSetupSessionsEmailsEndpointListParams struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// Create queues the link of an admin setup session for delivery to up to 10 recipients in total, given as email addresses and/or organization member IDs. Recipients are deduplicated case-insensitively and processed in order; if a later recipient fails, earlier ones may already have been queued, and retrying the request within the same window does not email them again. A recipient already emailed this setup link in the current fixed 10-minute window is not emailed again; their earlier record is returned with `deduplicated` set to true. A setup link can be emailed to at most 50 different recipients within 24 hours, and the setup links of an instance to at most 200 different recipients combined; a request that would exceed either limit is rejected with `admin_setup_email_limit_reached` before any email is queued. Passing `organization_member_ids` also requires the `organization.member:read` scope, since the response contains the members' email addresses.
func (e *ProviderDeploymentsSetupSessionsEmailsEndpoint) Create(instanceId string, providerSetupSessionId string, body *ProviderDeploymentsSetupSessionsEmailsEndpointCreateBody) (*emails.ProviderDeploymentsSetupSessionsEmailsCreateOutput, error) {
	req := &endpoint.Request{
		Path: []string{"instances", instanceId, "provider-setup-sessions", providerSetupSessionId, "emails"},
		Body: body,
	}
	var result emails.ProviderDeploymentsSetupSessionsEmailsCreateOutput
	if err := e.client.Post(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List returns a paginated list of the emails queued for delivery with the link of a provider setup session. Emails addressed to organization members have `recipient_email` set to null unless the caller also has the `organization.member:read` scope.
func (e *ProviderDeploymentsSetupSessionsEmailsEndpoint) List(instanceId string, providerSetupSessionId string, params *ProviderDeploymentsSetupSessionsEmailsEndpointListParams) (*emails.ProviderDeploymentsSetupSessionsEmailsListOutput, error) {
	var query map[string]any
	if params != nil {
		query = endpoint.StructToQuery(params)
	}
	req := &endpoint.Request{
		Path:  []string{"instances", instanceId, "provider-setup-sessions", providerSetupSessionId, "emails"},
		Query: query,
	}
	var result emails.ProviderDeploymentsSetupSessionsEmailsListOutput
	if err := e.client.Get(req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
