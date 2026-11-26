package instances

import (
	"encoding/json"
	"time"
)

// AgentsInstancesListOutputItemsAgentClient - Client this instance connects with.
type AgentsInstancesListOutputItemsAgentClient struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the client connects.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// LastConnectedAt - When the client last connected. `null` when never connected.
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
}

// AgentsInstancesListOutputItems represents the agents instances list output items type.
type AgentsInstancesListOutputItems struct {
	// Object - String representing the object’s type.
	Object string `json:"object"`
	// Id - Unique ID.
	Id string `json:"id"`
	// Type - How the agent instance connects and runs.
	Type string `json:"type"`
	// Name - Display name.
	Name string `json:"name"`
	// Version - Client version. `null` when the instance reports no version.
	Version *string `json:"version,omitempty"`
	// Description - Free-text description.
	Description *string `json:"description,omitempty"`
	// AgentId - Agent this instance belongs to.
	AgentId string `json:"agent_id"`
	// AgentClient - Client this instance connects with.
	AgentClient *AgentsInstancesListOutputItemsAgentClient `json:"agent_client,omitempty"`
	// CreatedAt - Creation time.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt - Last update time.
	UpdatedAt time.Time `json:"updated_at"`
	// LastConnectedAt - When the instance last connected. `null` when never connected.
	LastConnectedAt *time.Time `json:"last_connected_at,omitempty"`
}

// AgentsInstancesListOutputPagination represents the agents instances list output pagination type.
type AgentsInstancesListOutputPagination struct {
	// HasMoreBefore - Whether there are items before this page.
	HasMoreBefore bool `json:"has_more_before"`
	// HasMoreAfter - Whether there are items after this page.
	HasMoreAfter bool `json:"has_more_after"`
	// Before - Cursor for the previous page, returned by lists that page by cursor. Pass it as `before`.
	Before *string `json:"before,omitempty"`
	// After - Cursor for the next page, returned by lists that page by cursor. Pass it as `after`.
	After *string `json:"after,omitempty"`
}

// AgentsInstancesListOutput represents the agents instances list output type.
type AgentsInstancesListOutput struct {
	// Object - String representing the object’s type.
	Object     string                              `json:"object"`
	Items      []AgentsInstancesListOutputItems    `json:"items"`
	Pagination AgentsInstancesListOutputPagination `json:"pagination"`
}

// MapAgentsInstancesListOutputFromJSON deserializes JSON data into a AgentsInstancesListOutput.
func MapAgentsInstancesListOutputFromJSON(data []byte) (*AgentsInstancesListOutput, error) {
	var v AgentsInstancesListOutput
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsInstancesListOutputToJSON serializes a AgentsInstancesListOutput to JSON.
func MapAgentsInstancesListOutputToJSON(v *AgentsInstancesListOutput) ([]byte, error) {
	return json.Marshal(v)
}

// AgentsInstancesListQueryCreatedAt - Only return results whose agent instance creation time is in this range. Set `gt`, `lt`, or both.
type AgentsInstancesListQueryCreatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsInstancesListQueryUpdatedAt - Only return results whose agent instance last update time is in this range. Set `gt`, `lt`, or both.
type AgentsInstancesListQueryUpdatedAt struct {
	// Gt - Exclusive lower bound.
	Gt *time.Time `json:"gt,omitempty"`
	// Lt - Exclusive upper bound.
	Lt *time.Time `json:"lt,omitempty"`
}

// AgentsInstancesListQuery represents the agents instances list query type.
type AgentsInstancesListQuery struct {
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
	// Type - Only return agent instances with one of these types.
	Type *any `json:"type,omitempty"`
	// Id - Only return agent instances with one of these IDs.
	Id *any `json:"id,omitempty"`
	// AgentClientId - Only return agent instances linked to one of these agent clients.
	AgentClientId *any `json:"agent_client_id,omitempty"`
	// CreatedAt - Only return results whose agent instance creation time is in this range. Set `gt`, `lt`, or both.
	CreatedAt *AgentsInstancesListQueryCreatedAt `json:"created_at,omitempty"`
	// UpdatedAt - Only return results whose agent instance last update time is in this range. Set `gt`, `lt`, or both.
	UpdatedAt *AgentsInstancesListQueryUpdatedAt `json:"updated_at,omitempty"`
}

// MapAgentsInstancesListQueryFromJSON deserializes JSON data into a AgentsInstancesListQuery.
func MapAgentsInstancesListQueryFromJSON(data []byte) (*AgentsInstancesListQuery, error) {
	var v AgentsInstancesListQuery
	if err := json.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

// MapAgentsInstancesListQueryToJSON serializes a AgentsInstancesListQuery to JSON.
func MapAgentsInstancesListQueryToJSON(v *AgentsInstancesListQuery) ([]byte, error) {
	return json.Marshal(v)
}
