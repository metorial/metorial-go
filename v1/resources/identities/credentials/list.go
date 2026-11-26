package credentials

import (
	"encoding/json"
	"time"
)

// IdentitiesCredentialsListOutputItems represents the identities credentials list output items type.
type IdentitiesCredentialsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the credential. `archived` once you archive the credential.
	Status string `json:"status"`
	// IdentityId - Identity that owns this credential.
	IdentityId string `json:"identity_id"`
	// ProviderId - Provider this credential holds access for.
	ProviderId string `json:"provider_id"`
	// DeploymentId - Provider deployment used by this credential.
	DeploymentId *string `json:"deployment_id,omitempty"`
	// ConfigId - Provider config used by this credential.
	ConfigId *string `json:"config_id,omitempty"`
	// AuthConfigId - Provider auth config used by this credential.
	AuthConfigId *string `json:"auth_config_id,omitempty"`
	// DelegationConfigId - Delegation config applied to this credential.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesCredentialsListOutputPagination represents the identities credentials list output pagination type.
type IdentitiesCredentialsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentitiesCredentialsListOutput represents the identities credentials list output type.
type IdentitiesCredentialsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                    `json:"object"`
	Items      []IdentitiesCredentialsListOutputItems    `json:"items"`
	Pagination IdentitiesCredentialsListOutputPagination `json:"pagination"`
}

// MapIdentitiesCredentialsListOutputFromJSON deserializes JSON data into a IdentitiesCredentialsListOutput.
func MapIdentitiesCredentialsListOutputFromJSON(data []byte) (*IdentitiesCredentialsListOutput, error) {
	var v IdentitiesCredentialsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesCredentialsListOutputToJSON serializes a IdentitiesCredentialsListOutput to JSON.
func MapIdentitiesCredentialsListOutputToJSON(v *IdentitiesCredentialsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesCredentialsListQueryCreatedAt - Only return results whose identity credential creation time is in this range. Set `gt`, `lt`, or both.
type IdentitiesCredentialsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesCredentialsListQueryUpdatedAt - Only return results whose identity credential last update time is in this range. Set `gt`, `lt`, or both.
type IdentitiesCredentialsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesCredentialsListQuery represents the identities credentials list query type.
type IdentitiesCredentialsListQuery struct {
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
	// Status - Only return identity credentials with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity credentials with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identity credentials owned by one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Only return identity credentials owned by one of these identity actors.
	ActorId *any `json:"actor_id,omitempty"`
	// ConsumerProfileId - Only return identity credentials owned by one of these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// IdentityId - Only return identity credentials attached to one of these identities.
	IdentityId *any `json:"identity_id,omitempty"`
	// ProviderId - Only return identity credentials for one of these providers.
	ProviderId *any `json:"provider_id,omitempty"`
	// ProviderDeploymentId - Only return identity credentials attached to one of these provider deployments.
	ProviderDeploymentId *any `json:"provider_deployment_id,omitempty"`
	// ProviderConfigId - Only return identity credentials attached to one of these provider configs.
	ProviderConfigId *any `json:"provider_config_id,omitempty"`
	// ProviderAuthConfigId - Only return identity credentials attached to one of these provider auth configs.
	ProviderAuthConfigId *any `json:"provider_auth_config_id,omitempty"`
	// CreatedAt - Only return results whose identity credential creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IdentitiesCredentialsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity credential last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentitiesCredentialsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentitiesCredentialsListQueryFromJSON deserializes JSON data into a IdentitiesCredentialsListQuery.
func MapIdentitiesCredentialsListQueryFromJSON(data []byte) (*IdentitiesCredentialsListQuery, error) {
	var v IdentitiesCredentialsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesCredentialsListQueryToJSON serializes a IdentitiesCredentialsListQuery to JSON.
func MapIdentitiesCredentialsListQueryToJSON(v *IdentitiesCredentialsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
