package identityactors

import (
	"encoding/json"
	"time"
)

// IdentityActorsListOutputItems represents the identity actors list output items type.
type IdentityActorsListOutputItems struct {
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

// IdentityActorsListOutputPagination represents the identity actors list output pagination type.
type IdentityActorsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentityActorsListOutput represents the identity actors list output type.
type IdentityActorsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                             `json:"object"`
	Items      []IdentityActorsListOutputItems    `json:"items"`
	Pagination IdentityActorsListOutputPagination `json:"pagination"`
}

// MapIdentityActorsListOutputFromJSON deserializes JSON data into a IdentityActorsListOutput.
func MapIdentityActorsListOutputFromJSON(data []byte) (*IdentityActorsListOutput, error) {
	var v IdentityActorsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsListOutputToJSON serializes a IdentityActorsListOutput to JSON.
func MapIdentityActorsListOutputToJSON(v *IdentityActorsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentityActorsListQueryCreatedAt - Only return results whose identity actor creation time is in this range. Set `gt`, `lt`, or both.
type IdentityActorsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentityActorsListQueryUpdatedAt - Only return results whose identity actor last update time is in this range. Set `gt`, `lt`, or both.
type IdentityActorsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentityActorsListQuery represents the identity actors list query type.
type IdentityActorsListQuery struct {
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
	// Search - Filter actors by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identity actors with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity actors with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentId - Only return identity actors linked to one of these agents.
	AgentId *any `json:"agent_id,omitempty"`
	// ConsumerProfileId - Only return identity actors linked to one of these consumer profiles.
	ConsumerProfileId *any `json:"consumer_profile_id,omitempty"`
	// ConsumerId - Only return identity actors linked to one of these consumers.
	ConsumerId *any `json:"consumer_id,omitempty"`
	// CreatedAt - Only return results whose identity actor creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IdentityActorsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity actor last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentityActorsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentityActorsListQueryFromJSON deserializes JSON data into a IdentityActorsListQuery.
func MapIdentityActorsListQueryFromJSON(data []byte) (*IdentityActorsListQuery, error) {
	var v IdentityActorsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentityActorsListQueryToJSON serializes a IdentityActorsListQuery to JSON.
func MapIdentityActorsListQueryToJSON(v *IdentityActorsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
