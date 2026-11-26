package identities

import (
	"encoding/json"
	"time"
)

// IdentitiesListOutputItemsOwnerActor represents the identities list output items owner actor type.
type IdentitiesListOutputItemsOwnerActor struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - Type of actor that owns or participates in identities.
	Type string `json:"type"`
	// Status - Current lifecycle status of the identity actor. `archived` once you archive the actor.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// AgentId - Linked agent identifier when this actor represents an agent.
	AgentId *string `json:"agent_id,omitempty"`
	// ConsumerId - Identifier of the consumer that owns this actor when its type is consumer.
	ConsumerId *string `json:"consumer_id,omitempty"`
	// DefaultIdentityId - Identifier of the actor's default identity, which is its oldest active identity.
	DefaultIdentityId *string `json:"default_identity_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesListOutputItemsOwner - Actor that owns the identity.
type IdentitiesListOutputItemsOwner struct {
	// Type - Owner type for the identity.
	Type  string                              `json:"type"`
	Actor IdentitiesListOutputItemsOwnerActor `json:"actor"`
}

// IdentitiesListOutputItemsCredentials represents the identities list output items credentials type.
type IdentitiesListOutputItemsCredentials struct {
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

// IdentitiesListOutputItems represents the identities list output items type.
type IdentitiesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the identity. `archived` once you archive the identity.
	Status string `json:"status"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// Owner - Actor that owns the identity.
	Owner IdentitiesListOutputItemsOwner `json:"owner"`
	// Credentials - Credentials currently attached to the identity.
	Credentials []IdentitiesListOutputItemsCredentials `json:"credentials"`
	// DelegationConfigId - Default delegation config applied to the identity. `null` when none is set.
	DelegationConfigId *string `json:"delegation_config_id,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesListOutputPagination represents the identities list output pagination type.
type IdentitiesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentitiesListOutput represents the identities list output type.
type IdentitiesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                         `json:"object"`
	Items      []IdentitiesListOutputItems    `json:"items"`
	Pagination IdentitiesListOutputPagination `json:"pagination"`
}

// MapIdentitiesListOutputFromJSON deserializes JSON data into a IdentitiesListOutput.
func MapIdentitiesListOutputFromJSON(data []byte) (*IdentitiesListOutput, error) {
	var v IdentitiesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesListOutputToJSON serializes a IdentitiesListOutput to JSON.
func MapIdentitiesListOutputToJSON(v *IdentitiesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesListQueryCreatedAt - Only return results whose identity creation time is in this range. Set `gt`, `lt`, or both.
type IdentitiesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesListQueryUpdatedAt - Only return results whose identity last update time is in this range. Set `gt`, `lt`, or both.
type IdentitiesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesListQuery represents the identities list query type.
type IdentitiesListQuery struct {
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
	// Search - Filter identities by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identities with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identities with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identities owned by one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ActorId - Only return identities owned by one of these identity actors.
	ActorId *any `json:"actor_id,omitempty"`
	// CreatedAt - Only return results whose identity creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IdentitiesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentitiesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentitiesListQueryFromJSON deserializes JSON data into a IdentitiesListQuery.
func MapIdentitiesListQueryFromJSON(data []byte) (*IdentitiesListQuery, error) {
	var v IdentitiesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesListQueryToJSON serializes a IdentitiesListQuery to JSON.
func MapIdentitiesListQueryToJSON(v *IdentitiesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
