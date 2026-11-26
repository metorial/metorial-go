package agents

import (
	"encoding/json"
	"time"
)

// AgentsListOutputItems represents the agents list output items type.
type AgentsListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the agent connects and runs.
	Type string `json:"type"`
	// Status - Current lifecycle status of the agent. `archived` once you archive the agent.
	Status string `json:"status"`
	// Name - Display name.
	Name string `json:"name"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// Slug - URL-safe name, unique within the instance.
	Slug string `json:"slug"`
	// Metadata - Custom data you attach to identify or organize this record.
	Metadata *map[string]any `json:"metadata,omitempty"`
	// ActorId - Identity actor linked to the agent.
	ActorId string `json:"actor_id"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// ArchivedAt - When this was archived. `null` while active.
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}

// AgentsListOutputPagination represents the agents list output pagination type.
type AgentsListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AgentsListOutput represents the agents list output type.
type AgentsListOutput struct {
	// Object - String representing the object’s type.
	Object     string                     `json:"object"`
	Items      []AgentsListOutputItems    `json:"items"`
	Pagination AgentsListOutputPagination `json:"pagination"`
}

// MapAgentsListOutputFromJSON deserializes JSON data into a AgentsListOutput.
func MapAgentsListOutputFromJSON(data []byte) (*AgentsListOutput, error) {
	var v AgentsListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsListOutputToJSON serializes a AgentsListOutput to JSON.
func MapAgentsListOutputToJSON(v *AgentsListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AgentsListQueryCreatedAt - Only return results whose agent creation time is in this range. Set `gt`, `lt`, or both.
type AgentsListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsListQueryUpdatedAt - Only return results whose agent last update time is in this range. Set `gt`, `lt`, or both.
type AgentsListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsListQuery represents the agents list query type.
type AgentsListQuery struct {
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
	// Search - Filter agents by name or description.
	Search *string `json:"search,omitempty"`
	// Status - Only return agents with one of these statuses.
	Status *any `json:"status,omitempty"`
	// Type - Only return agents with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return agents with one of these IDs.
	Id *any `json:"id,omitempty"`
	// CreatedAt - Only return results whose agent creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *AgentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose agent last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *AgentsListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapAgentsListQueryFromJSON deserializes JSON data into a AgentsListQuery.
func MapAgentsListQueryFromJSON(data []byte) (*AgentsListQuery, error) {
	var v AgentsListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsListQueryToJSON serializes a AgentsListQuery to JSON.
func MapAgentsListQueryToJSON(v *AgentsListQuery) ([]byte, error) {
	return json.Marshal(v)
}
