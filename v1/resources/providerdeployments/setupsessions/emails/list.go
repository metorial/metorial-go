package emails

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsSetupSessionsEmailsListOutputItems represents the provider deployments setup sessions emails list output items type.
type ProviderDeploymentsSetupSessionsEmailsListOutputItems struct {
	// Object - String representing the object's type
	Object string `json:"object"`
	// Id - Unique setup session email identifier
	Id string `json:"id"`
	// Deduplicated - True when this recipient was already emailed this setup link in the current fixed 10-minute window, so no new email was queued and the earlier record is returned instead. Only set by the create endpoint; always false when listing.
	Deduplicated bool `json:"deduplicated"`
	// SetupSessionId - ID of the setup session whose link was emailed
	SetupSessionId string `json:"setup_session_id"`
	// RecipientEmail - Email address the setup link was queued for delivery to, lowercased. Null when the email was addressed to an organization member and the caller lacks the `organization.member:read` scope; `organization_member_id` identifies the member instead.
	RecipientEmail *string `json:"recipient_email,omitempty"`
	// OrganizationMemberId - ID of the organization member the email was addressed to, or null when it was addressed to a plain email address
	OrganizationMemberId *string `json:"organization_member_id,omitempty"`
	// CreatedAt - Timestamp when the email was queued for delivery
	CreatedAt time.Time `json:"created_at"`
}

// ProviderDeploymentsSetupSessionsEmailsListOutputPagination represents the provider deployments setup sessions emails list output pagination type.
type ProviderDeploymentsSetupSessionsEmailsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProviderDeploymentsSetupSessionsEmailsListOutput represents the provider deployments setup sessions emails list output type.
type ProviderDeploymentsSetupSessionsEmailsListOutput struct {
	Items      []ProviderDeploymentsSetupSessionsEmailsListOutputItems    `json:"items"`
	Pagination ProviderDeploymentsSetupSessionsEmailsListOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsSetupSessionsEmailsListOutputFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsEmailsListOutput.
func MapProviderDeploymentsSetupSessionsEmailsListOutputFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsEmailsListOutput, error) {
	var v ProviderDeploymentsSetupSessionsEmailsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsEmailsListOutputToJSON serializes a ProviderDeploymentsSetupSessionsEmailsListOutput to JSON.
func MapProviderDeploymentsSetupSessionsEmailsListOutputToJSON(v *ProviderDeploymentsSetupSessionsEmailsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsSetupSessionsEmailsListQuery represents the provider deployments setup sessions emails list query type.
type ProviderDeploymentsSetupSessionsEmailsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
}

// MapProviderDeploymentsSetupSessionsEmailsListQueryFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsEmailsListQuery.
func MapProviderDeploymentsSetupSessionsEmailsListQueryFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsEmailsListQuery, error) {
	var v ProviderDeploymentsSetupSessionsEmailsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsEmailsListQueryToJSON serializes a ProviderDeploymentsSetupSessionsEmailsListQuery to JSON.
func MapProviderDeploymentsSetupSessionsEmailsListQueryToJSON(v *ProviderDeploymentsSetupSessionsEmailsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
