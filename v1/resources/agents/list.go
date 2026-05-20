package agents

import (
	"encoding/json"
	"time"
)

// AgentsListOutputItems represents the agents list output items type.
type AgentsListOutputItems struct {
	Object      string          `json:"object"`
	Id          string          `json:"id"`
	Type        string          `json:"type"`
	Status      string          `json:"status"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Slug        string          `json:"slug"`
	Metadata    *map[string]any `json:"metadata,omitempty"`
	ActorId     string          `json:"actor_id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	ArchivedAt  *time.Time      `json:"archived_at,omitempty"`
}

// AgentsListOutputPagination represents the agents list output pagination type.
type AgentsListOutputPagination struct {
	HasMoreBefore bool `json:"has_more_before"`
	HasMoreAfter  bool `json:"has_more_after"`
}

// AgentsListOutput represents the agents list output type.
type AgentsListOutput struct {
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

// AgentsListQueryCreatedAt - Filter agent creation time by date range
type AgentsListQueryCreatedAt struct {
	// Gt - Only include records after this timestamp for agent creation time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for agent creation time
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsListQueryUpdatedAt - Filter agent last update time by date range
type AgentsListQueryUpdatedAt struct {
	// Gt - Only include records after this timestamp for agent last update time
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Only include records before this timestamp for agent last update time
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsListQuery represents the agents list query type.
type AgentsListQuery struct {
	Limit  *float64 `json:"limit,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	Cursor *string  `json:"cursor,omitempty"`
	Order  *string  `json:"order,omitempty"`
	Search *string  `json:"search,omitempty"`
	Status *any     `json:"status,omitempty"`
	Type   *any     `json:"type,omitempty"`
	Id     *any     `json:"id,omitempty"`
	// CreatedAt - Filter agent creation time by date range
	CreatedAt *AgentsListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Filter agent last update time by date range
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
