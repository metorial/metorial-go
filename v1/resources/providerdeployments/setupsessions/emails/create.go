package emails

import (
	"encoding/json"
	"time"
)

// ProviderDeploymentsSetupSessionsEmailsCreateOutputItems represents the provider deployments setup sessions emails create output items type.
type ProviderDeploymentsSetupSessionsEmailsCreateOutputItems struct {
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

// ProviderDeploymentsSetupSessionsEmailsCreateOutputPagination represents the provider deployments setup sessions emails create output pagination type.
type ProviderDeploymentsSetupSessionsEmailsCreateOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// ProviderDeploymentsSetupSessionsEmailsCreateOutput represents the provider deployments setup sessions emails create output type.
type ProviderDeploymentsSetupSessionsEmailsCreateOutput struct {
	Items      []ProviderDeploymentsSetupSessionsEmailsCreateOutputItems    `json:"items"`
	Pagination ProviderDeploymentsSetupSessionsEmailsCreateOutputPagination `json:"pagination"`
}

// MapProviderDeploymentsSetupSessionsEmailsCreateOutputFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsEmailsCreateOutput.
func MapProviderDeploymentsSetupSessionsEmailsCreateOutputFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsEmailsCreateOutput, error) {
	var v ProviderDeploymentsSetupSessionsEmailsCreateOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsEmailsCreateOutputToJSON serializes a ProviderDeploymentsSetupSessionsEmailsCreateOutput to JSON.
func MapProviderDeploymentsSetupSessionsEmailsCreateOutputToJSON(v *ProviderDeploymentsSetupSessionsEmailsCreateOutput) ([]byte, error) {
	return json.Marshal(v)
}

// ProviderDeploymentsSetupSessionsEmailsCreateBody represents the provider deployments setup sessions emails create body type.
type ProviderDeploymentsSetupSessionsEmailsCreateBody struct {
	// Emails - Email addresses to send the setup link to, at most 10
	Emails *[]string `json:"emails,omitempty"`
	// OrganizationMemberIds - Organization members to send the setup link to, at most 10. Each member is emailed at their organization profile email address, falling back to their account email when the profile has none.
	OrganizationMemberIds *[]string `json:"organization_member_ids,omitempty"`
}

// MapProviderDeploymentsSetupSessionsEmailsCreateBodyFromJSON deserializes JSON data into a ProviderDeploymentsSetupSessionsEmailsCreateBody.
func MapProviderDeploymentsSetupSessionsEmailsCreateBodyFromJSON(data []byte) (*ProviderDeploymentsSetupSessionsEmailsCreateBody, error) {
	var v ProviderDeploymentsSetupSessionsEmailsCreateBody
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapProviderDeploymentsSetupSessionsEmailsCreateBodyToJSON serializes a ProviderDeploymentsSetupSessionsEmailsCreateBody to JSON.
func MapProviderDeploymentsSetupSessionsEmailsCreateBodyToJSON(v *ProviderDeploymentsSetupSessionsEmailsCreateBody) ([]byte, error) {
	return json.Marshal(v)
}
