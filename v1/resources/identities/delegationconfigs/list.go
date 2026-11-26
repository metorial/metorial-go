package delegationconfigs

import (
	"encoding/json"
	"time"
)

// IdentitiesDelegationConfigsListOutputItems represents the identities delegation configs list output items type.
type IdentitiesDelegationConfigsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Status - Current lifecycle status of the delegation config. `archived` once you archive the config.
	Status string `json:"status"`
	// IsDefault - Delegations use this config when none is specified.
	IsDefault bool `json:"is_default"`
	// Name - Display name.
	Name *string `json:"name,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// SubDelegationBehavior - Whether delegates can delegate further. `require_consent` asks the owner for approval first.
	SubDelegationBehavior string `json:"sub_delegation_behavior"`
	// SubDelegationDepth - Maximum sub-delegation depth. `0` blocks sub-delegation.
	SubDelegationDepth float64 `json:"sub_delegation_depth"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
}

// IdentitiesDelegationConfigsListOutputPagination represents the identities delegation configs list output pagination type.
type IdentitiesDelegationConfigsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// IdentitiesDelegationConfigsListOutput represents the identities delegation configs list output type.
type IdentitiesDelegationConfigsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                                          `json:"object"`
	Items      []IdentitiesDelegationConfigsListOutputItems    `json:"items"`
	Pagination IdentitiesDelegationConfigsListOutputPagination `json:"pagination"`
}

// MapIdentitiesDelegationConfigsListOutputFromJSON deserializes JSON data into a IdentitiesDelegationConfigsListOutput.
func MapIdentitiesDelegationConfigsListOutputFromJSON(data []byte) (*IdentitiesDelegationConfigsListOutput, error) {
	var v IdentitiesDelegationConfigsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationConfigsListOutputToJSON serializes a IdentitiesDelegationConfigsListOutput to JSON.
func MapIdentitiesDelegationConfigsListOutputToJSON(v *IdentitiesDelegationConfigsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// IdentitiesDelegationConfigsListQueryCreatedAt - Only return results whose identity delegation config creation time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationConfigsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationConfigsListQueryUpdatedAt - Only return results whose identity delegation config last update time is in this range. Set `gt`, `lt`, or both.
type IdentitiesDelegationConfigsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// IdentitiesDelegationConfigsListQuery represents the identities delegation configs list query type.
type IdentitiesDelegationConfigsListQuery struct {
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
	// Search - Filter configs by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return identity delegation configs with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Id - Only return identity delegation configs with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose identity delegation config creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *IdentitiesDelegationConfigsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose identity delegation config last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *IdentitiesDelegationConfigsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapIdentitiesDelegationConfigsListQueryFromJSON deserializes JSON data into a IdentitiesDelegationConfigsListQuery.
func MapIdentitiesDelegationConfigsListQueryFromJSON(data []byte) (*IdentitiesDelegationConfigsListQuery, error) {
	var v IdentitiesDelegationConfigsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapIdentitiesDelegationConfigsListQueryToJSON serializes a IdentitiesDelegationConfigsListQuery to JSON.
func MapIdentitiesDelegationConfigsListQueryToJSON(v *IdentitiesDelegationConfigsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
